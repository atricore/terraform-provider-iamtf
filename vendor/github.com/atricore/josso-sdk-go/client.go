package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	api "github.com/atricore/josso-api-go"
)

const (
	DEFAULT_SVR = "__default__"
)

type (
	IdbusApiClient struct {
		apiClient *api.APIClient
		config    *api.Configuration
		servers   map[string]*ServerConnection
		logger    Logger
	}

	IdbusServer struct {
		Config      *api.ServerConfiguration
		Credentials *ServerCredentials
	}

	ServerCredentials struct {
		ClientId string
		Secret   string
	}

	ServerConnection struct {
		Authn  ServerAuthn
		Server *IdbusServer
	}

	ServerAuthn struct {
		AccessToken  string
		RefreshToken string
	}
)

func NewIdbusApiClientWithDefaults() *IdbusApiClient {
	return NewIdbusApiClient(&DefaultLogger{debug: true}, false)
}

func NewIdbusApiClient(l Logger, trace bool) *IdbusApiClient {
	l.Debugf("newIdbusApiClient TRACE: %t", trace)

	if trace {
		log.Print("Using client TRACE ON")
	}
	cfg := config(trace)
	cli := &IdbusApiClient{
		config:    cfg,
		apiClient: api.NewAPIClient(cfg),
		servers:   make(map[string]*ServerConnection),
		logger:    l,
	}
	return cli
}

func (c *IdbusApiClient) Logger() Logger {
	return c.logger
}

/*
* Register a new server
 */
func (c *IdbusApiClient) RegisterServer(svr *IdbusServer, operation string) error {

	key := operation
	if key == "" {
		key = DEFAULT_SVR
	}
	c.logger.Tracef("registering server %s", svr.Config.URL)

	// We replace configuration if the server is already registerd for the URL
	sc := ServerConnection{
		Server: svr,
	}
	if ok := c.servers[key]; ok != nil {
		c.logger.Tracef("replacing server registration")
		found := false
		for _, sc := range c.apiClient.GetConfig().Servers {
			if sc.URL == svr.Config.URL {
				c.logger.Tracef("replacing server configuration for %s", sc.URL)
				sc.Description = svr.Config.Description
				sc.Variables = svr.Config.Variables
				found = true
				break
			}
		}
		if !found {
			c.logger.Errorf("server registered, but config not found for %s", key)
			return fmt.Errorf("server registered, but config not found for %s", key)
		}
	} else {
		c.logger.Tracef("adding server configuration for %s", svr.Config.URL)
		c.apiClient.GetConfig().Servers = append(c.apiClient.GetConfig().Servers, *svr.Config)

	}
	c.servers[key] = &sc

	if operation != "" {
		scs := c.apiClient.GetConfig().OperationServers[operation]
		scs = append(scs, *svr.Config)
		c.apiClient.GetConfig().OperationServers[operation] = scs
	}

	return nil
}

func (c *IdbusApiClient) Authn() error {

	sc, err := c.IdbusServerForOperation("DefaultApiService.SignOn") // Also hard-coded in generated openapi
	if err != nil {
		return err
	}

	c.logger.Tracef("authn: %s secret found: %t",
		sc.Server.Credentials.ClientId,
		sc.Server.Credentials.Secret != "")

	req := c.apiClient.DefaultApi.SignOn(context.Background())
	req = req.OIDCSignOnRequest(api.OIDCSignOnRequest{
		ClientId: &sc.Server.Credentials.ClientId,
		Secret:   &sc.Server.Credentials.Secret})

	res, _, err := c.apiClient.DefaultApi.SignOnExecute(req)
	if err != nil {
		return fmt.Errorf("cannot authenticate with IAMTF/JOSSO server [%s]: %v", sc.Server.Config.URL, err)
	}

	sc.Authn.AccessToken = *res.AccessToken
	sc.Authn.RefreshToken = *res.RefreshToken

	return nil

}

func ServerVersion(cfg *IdbusServer) (string, error) {

	c1, err := CreateClient(cfg, false)
	if err != nil {
		return "", fmt.Errorf("cannot get version from IAMTF/JOSSO server [%s]: %v", cfg.Config.URL, err)
	}

	sc, err := c1.IdbusServerForOperation("DefaultApiService.ServerVersion") // Also hard-coded in generated openapi
	if err != nil {
		return "", err
	}

	req := c1.apiClient.DefaultApi.Version(context.Background())
	req = req.ServerVersionRequest(api.ServerVersionRequest{})

	res, _, err := c1.apiClient.DefaultApi.VersionExecute(req)
	if err != nil {
		return "", fmt.Errorf("cannot get version from IAMTF/JOSSO server [%s]: %v", sc.Server.Config.URL, err)
	}

	if res.Version == nil {
		return "", fmt.Errorf("cannot get version from IAMTF/JOSSO server [%s]: response is nil", sc.Server.Config.URL)
	}

	return *res.Version, nil

}

func GetServerConfigFromEnv() (*IdbusServer, error) {

	clientSecret := os.Getenv("JOSSO_API_SECRET")
	clientId := os.Getenv("JOSSO_API_CLIENT_ID")
	endpoint := os.Getenv("JOSSO_API_ENDPOINT")
	if clientSecret == "" || clientId == "" || endpoint == "" {
		return nil, errors.New("JOSSO variables must be set for acceptance tests")
	}

	s := IdbusServer{
		Config: &api.ServerConfiguration{
			URL:         endpoint,
			Description: "JOSSO Test server",
		},
		Credentials: &ServerCredentials{
			ClientId: clientId,
			Secret:   clientSecret,
		},
	}
	return &s, nil
}

func CreateClient(s *IdbusServer, authn bool) (*IdbusApiClient, error) {

	trace, err := GetenvBool("JOSSO_API_TRACE")
	if err != nil {
		trace = false
	}

	l := DefaultLogger{debug: trace}

	c := NewIdbusApiClient(&l, trace)
	err = c.RegisterServer(s, "")
	if err != nil {
		return nil, err
	}

	if authn {
		err = c.Authn()
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Create default configuration
func config(debug bool) *api.Configuration {
	return &api.Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            debug,
		Servers:          api.ServerConfigurations{},
		OperationServers: make(map[string]api.ServerConfigurations), // Servers for specific operations
	}
}

func (c *IdbusApiClient) IdbusServerForOperation(operation string) (*ServerConnection, error) {
	sc, ok := c.servers[operation]
	if ok {
		return sc, nil
	} else {
		return c.servers[DEFAULT_SVR], nil
	}
}
