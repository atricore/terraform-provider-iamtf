package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new ExtSaml2Sp in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateExtSaml2Sp(ida string, extsp api.ExternalSaml2ServiceProviderDTO) (api.ExternalSaml2ServiceProviderDTO, error) {
	var result api.ExternalSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("createExtSaml2Sp : %s [%s]", *extsp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateExtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ExtSaml2Sp(&extsp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateExtSaml2Sp(ctx)
	req = req.StoreExtSaml2SpReq(api.StoreExtSaml2SpReq{IdOrName: &ida, Sp: &extsp})
	res, _, err := c.apiClient.DefaultApi.CreateExtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("CreateExtSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateExtSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no sp received after creation")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) UpdateExtSaml2Sp(ida string, sp api.ExternalSaml2ServiceProviderDTO) (api.ExternalSaml2ServiceProviderDTO, error) {
	var result api.ExternalSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("UpdateExtSaml2Sp. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateExtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ExtSaml2Sp(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateExtSaml2Sp(ctx)
	req = req.StoreExtSaml2SpReq(api.StoreExtSaml2SpReq{IdOrName: &ida, Sp: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateExtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateExtSaml2Sp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateExtSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no sp received after update")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) DeleteExtSaml2Sp(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteExtSaml2Sp. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteExtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteExtSaml2Sp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteExtSaml2Sp(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteExtSaml2SpExecute(req)

	if err != nil {
		c.logger.Errorf("deletesp. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deletesp. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deletesp. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetExtSaml2Sp(ida string, sp string) (api.ExternalSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("GetExtSaml2Sp. %s [%s]", sp, ida)
	var result api.ExternalSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetExtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetExtSaml2Sp(ctx)
	req = req.GetExtSaml2SpReq(api.GetExtSaml2SpReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.GetExtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("GetExtSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetExtSaml2Sp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Sp == nil {
		c.logger.Debugf("GetExtSaml2Sp. NOT FOUND %s", sp)
		return result, nil
	}

	if res.Sp != nil {
		result = *res.Sp
		c.logger.Debugf("GetExtSaml2Sp. %s found for ID/name %s", *result.Name, sp)
	} else {
		c.logger.Debugf("GetExtSaml2Sp. not found for ID/name %s", sp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetExtSaml2Sps(ida string) ([]api.ExternalSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("get ExtSaml2Sps: all [%s]", ida)
	var result []api.ExternalSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetExtSaml2Sps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetExtSaml2Sps(ctx)
	req = req.GetExtSaml2SpReq(api.GetExtSaml2SpReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetExtSaml2SpsExecute(req)
	if err != nil {
		c.logger.Errorf("getExtSaml2Sps. Error %v", err)
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

func ExtSaml2Sp(ExtSaml2Sp *api.ExternalSaml2ServiceProviderDTO) {
	ExtSaml2Sp.AdditionalProperties = make(map[string]interface{})
	ExtSaml2Sp.AdditionalProperties["@c"] = ".ExternalSaml2ServiceProviderDTO"
}
