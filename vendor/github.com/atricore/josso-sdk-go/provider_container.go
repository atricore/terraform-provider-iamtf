package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Gets an IdP based on the appliance name and idp name
func (c *IdbusApiClient) GetProvider(ida string, provider string) (api.ProviderContainerDTO, error) {
	c.logger.Debugf("getProvider. %s [%s]", provider, ida)
	var result api.ProviderContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetProvider") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetProvider(ctx)
	req = req.GetProviderReq(api.GetProviderReq{IdOrName: &ida, Name: &provider})
	res, _, err := c.apiClient.DefaultApi.GetProviderExecute(req)
	if err != nil {
		c.logger.Errorf("getProvider. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getProvider. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Provider == nil {
		c.logger.Debugf("getProvider. NOT FOUND %s", provider)
		return result, nil
	}

	result = res.GetProvider()

	return result, nil

}

func (c *IdbusApiClient) GetProviders(ida string) ([]api.ProviderContainerDTO, error) {
	c.logger.Debugf("getProviders: all [%s]", ida)
	var result []api.ProviderContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetProviders") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetProviders(ctx)
	req = req.GetApplianceReq(api.GetApplianceReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetProvidersExecute(req)
	if err != nil {
		c.logger.Errorf("getProviders. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Providers == nil {
		return result, nil
	}

	result = res.Providers

	return result, nil

}
