package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Gets an IdP based on the appliance name and idp name
func (c *IdbusApiClient) GetIdSource(ida string, idsource string) (api.IdSourceContainerDTO, error) {
	c.logger.Debugf("getIdSource. %s [%s]", idsource, ida)
	var result api.IdSourceContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdSource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSource(ctx)
	req = req.GetIdSourceReq(api.GetIdSourceReq{IdOrName: &ida, Name: &idsource})
	res, _, err := c.apiClient.DefaultApi.GetIdSourceExecute(req)
	if err != nil {
		c.logger.Errorf("getIdSource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdSource. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IdSource == nil {
		c.logger.Debugf("getIdSource. NOT FOUND %s", idsource)
		return result, nil
	}

	result = res.GetIdSource()

	return result, nil

}

func (c *IdbusApiClient) GetIdSources(ida string) ([]api.IdSourceContainerDTO, error) {
	c.logger.Debugf("getIdSources: all [%s]", ida)
	var result []api.IdSourceContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdSources") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSources(ctx)
	req = req.GetIdSourcesReq(api.GetIdSourcesReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdSourcesExecute(req)
	if err != nil {
		c.logger.Errorf("getIdSources. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.IdSources == nil {
		return result, nil
	}

	result = res.IdSources

	return result, nil

}
