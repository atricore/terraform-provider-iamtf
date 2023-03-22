package iamtf

import (
	"errors"
	"fmt"
	"os"
	"testing"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testaccProvidersFactories map[string]func() (*schema.Provider, error)
	testaccProvider           *schema.Provider
)

func init() {
	testaccProvider = Provider()
	testaccProvidersFactories = map[string]func() (*schema.Provider, error){
		"iamtf": func() (*schema.Provider, error) {
			return testaccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	_ = Provider()
}

func testaccPreCheck(t *testing.T) {
	err := accPreCheck()
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func accPreCheck() error {

	c := jossoConfigEnv()
	if c.clientId == "" || c.secret == "" || c.endpoint == "" {
		return errors.New("JOSSO variables must be set for acceptance tests")
	}
	return nil
}

func jossoConfigEnv() *Config {
	return &Config{
		endpoint:  os.Getenv("JOSSO_API_ENDPOINT"),
		clientId:  os.Getenv("JOSSO_API_CLIENT_ID"),
		secret:    os.Getenv("JOSSO_API_SECRET"),
		appliance: os.Getenv("JOSSO_API_APPLIANCE"),
	}
}

func jossoConfig() (*Config, error) {

	logLevel := hclog.Trace
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Level(logLevel),
		TimeFormat: "2006/01/02 03:04:05",
		Output:     os.Stdout,
	})
	logger.Info("creating JOSSO TEST client")

	config := jossoConfigEnv()
	client := cli.NewIdbusApiClient(&ProviderLogger{wrapped: logger}, true)
	config.logger = logger
	config.logLevel = int32(logLevel)
	config.apiClient = client
	config.trace = true

	var err error

	err = client.RegisterServer(
		&cli.IdbusServer{
			Config: &api.ServerConfiguration{
				URL:         config.endpoint,
				Description: "JOSSO Test server",
			},
			Credentials: &cli.ServerCredentials{
				ClientId: config.clientId,
				Secret:   config.secret,
			},
		},
		"")
	if err != nil {
		return nil, fmt.Errorf("error creating JOSSO client: %v", err)
	}

	if err = config.loadAndValidate(); err != nil {
		return config, fmt.Errorf("error initializing JOSSO client: %v", err)
	}
	return config, nil
}
