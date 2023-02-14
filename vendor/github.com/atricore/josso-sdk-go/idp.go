package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the idp dto to use as template
func (c *IdbusApiClient) CreateIdp(ida string, idp api.IdentityProviderDTO) (api.IdentityProviderDTO, error) {
	var result api.IdentityProviderDTO
	l := c.Logger()

	l.Debugf("createIdP : %s [%s]", *idp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdP(&idp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdP(ctx)
	req = req.StoreIdPReq(api.StoreIdPReq{IdOrName: &ida, Idp: &idp})
	res, _, err := c.apiClient.DefaultApi.CreateIdPExecute(req)
	if err != nil {
		c.logger.Errorf("createIdP. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdP. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no idp received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateIdp(ida string, idp api.IdentityProviderDTO) (api.IdentityProviderDTO, error) {
	var result api.IdentityProviderDTO
	l := c.Logger()

	l.Debugf("updateIdP. : %s [%s]", *idp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdP(&idp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdP(ctx)
	req = req.StoreIdPReq(api.StoreIdPReq{IdOrName: &ida, Idp: &idp})
	res, _, err := c.apiClient.DefaultApi.UpdateIdPExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdP. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdP. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no idp received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteIdp(ida string, idp string) (bool, error) {
	c.logger.Debugf("deleteIdp. %s [%s]", idp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdp") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdP(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &idp})
	res, _, err := c.apiClient.DefaultApi.DeleteIdPExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdp. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdp. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdp. Deleted %s : %t", idp, *res.Removed)

	return *res.Removed, err
}

// Gets an IdP based on the appliance name and idp name
func (c *IdbusApiClient) GetIdp(ida string, idp string) (api.IdentityProviderDTO, error) {
	c.logger.Debugf("getIdp. %s [%s]", idp, ida)
	var result api.IdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdP(ctx)
	req = req.GetIdPReq(api.GetIdPReq{IdOrName: &ida, Name: &idp})
	res, _, err := c.apiClient.DefaultApi.GetIdPExecute(req)
	if err != nil {
		c.logger.Errorf("getIdP. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdP. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("getIdP. NOT FOUND %s", idp)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("getIdP. %s found for ID/name %s", *result.Name, idp)
	} else {
		c.logger.Debugf("getIdP. not found for ID/name %s", idp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdps(ida string) ([]api.IdentityProviderDTO, error) {
	c.logger.Debugf("get idps: all [%s]", ida)
	var result []api.IdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdPs(ctx)
	req = req.GetIdPReq(api.GetIdPReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdPsExecute(req)
	if err != nil {
		c.logger.Errorf("getIdPs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Idps == nil {
		return result, nil
	}

	result = res.Idps

	return result, nil

}

func initIdP(idp *api.IdentityProviderDTO) {
	idp.AdditionalProperties = make(map[string]interface{})
	idp.AdditionalProperties["@c"] = ".IdentityProviderDTO"
}
