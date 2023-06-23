package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IntSaml2Sp in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateIntSaml2Sp(ida string, intsp api.InternalSaml2ServiceProviderDTO) (api.InternalSaml2ServiceProviderDTO, error) {
	var result api.InternalSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("createIntSaml2Sp : %s [%s]", *intsp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIntSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIntSaml2Sp(&intsp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIntSaml2Sp(ctx)
	req = req.StoreIntSaml2SpReq(api.StoreIntSaml2SpReq{IdOrName: &ida, Sp: &intsp})
	res, _, err := c.apiClient.DefaultApi.CreateIntSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("CreateIntSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateIntSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no IntSaml2Sp received after creation")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) UpdateIntSaml2Sp(ida string, sp api.InternalSaml2ServiceProviderDTO) (api.InternalSaml2ServiceProviderDTO, error) {
	var result api.InternalSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("UpdateIntSaml2Sp. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIntSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIntSaml2Sp(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIntSaml2Sp(ctx)
	req = req.StoreIntSaml2SpReq(api.StoreIntSaml2SpReq{IdOrName: &ida, Sp: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateIntSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateIntSaml2Sp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateIntSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no IntSaml2Sp received after update")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) DeleteIntSaml2Sp(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteIntSaml2Sp. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteIntSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIntSaml2Sp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIntSaml2Sp(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteIntSaml2SpExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIntSaml2Ss. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetIntSaml2Sp(ida string, sp string) (api.InternalSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("GetIntSaml2Sp. %s [%s]", sp, ida)
	var result api.InternalSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIntSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIntSaml2Sp(ctx)
	req = req.GetIntSaml2SpReq(api.GetIntSaml2SpReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.GetIntSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("GetIntSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetIntSaml2Sp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Sp == nil {
		c.logger.Debugf("GetIntSaml2Sp. NOT FOUND %s", sp)
		return result, nil
	}

	if res.Sp != nil {
		result = *res.Sp

		if res.Sp.Name == nil {
			return result, errors.New("no name received for IntSaml2Sp (unmarshalling error?)")
		}
		c.logger.Debugf("GetIntSaml2Sp. %s found for ID/name %s", *result.Name, sp)
	} else {
		c.logger.Debugf("GetIntSaml2Sp. not found for ID/name %s", sp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIntSaml2Sps(ida string) ([]api.InternalSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("get IntSaml2Sps: all [%s]", ida)
	var result []api.InternalSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIntSaml2Sps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIntSaml2Sps(ctx)
	req = req.GetIntSaml2SpReq(api.GetIntSaml2SpReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIntSaml2SpsExecute(req)
	if err != nil {
		c.logger.Errorf("getIntSaml2Sps. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Sps == nil {
		return result, nil
	}

	result = res.Sps

	return result, nil

}

func initIntSaml2Sp(IntSaml2Sp *api.InternalSaml2ServiceProviderDTO) {
	IntSaml2Sp.AdditionalProperties = make(map[string]interface{})
	IntSaml2Sp.AdditionalProperties["@c"] = ".InternalSaml2ServiceProviderDTO"
}
