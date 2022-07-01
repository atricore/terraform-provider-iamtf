package iamtf

import (
	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"

	"github.com/hashicorp/go-hclog"
)

type (
	// Config contains our provider schema values and JOSSO client
	Config struct {
		orgName   string
		clientId  string
		secret    string
		endpoint  string
		logLevel  int32
		logger    hclog.Logger
		trace     bool
		apiClient *cli.IdbusApiClient
	}
)

func (c *Config) Logger() hclog.Logger {
	return c.logger
}

func (c *Config) loadAndValidate() error {

	// Configure logger
	c.logger = hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Level(c.logLevel),
		TimeFormat: "2006/01/02 03:04:05",
	})

	l := ProviderLogger{wrapped: c.logger}

	if apiClient == nil {
		c.logger.Debug("loadAndValidate. creating IAM.tf client", "trace", c.trace)
		apiClient = cli.NewIdbusApiClient(l, c.trace)
	}

	// We reuse the API client
	c.apiClient = apiClient

	if err := c.apiClient.RegisterServer(&cli.IdbusServer{

		Config: &api.ServerConfiguration{
			URL:         c.endpoint,
			Description: "IAM.tf server at " + c.orgName,
		},
		Credentials: &cli.ServerCredentials{
			ClientId: c.clientId,
			Secret:   c.secret,
		},
	}, ""); err != nil {
		return err
	}

	return c.apiClient.Authn()

}
