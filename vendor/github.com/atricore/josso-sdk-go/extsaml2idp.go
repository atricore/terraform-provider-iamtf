package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IdPSaml2 in the provided identity appliance. It receives the appliance name or id and the IdP dto to use as template
func (c *IdbusApiClient) CreateIdPSaml2(ida string, extIdP api.ExternalSaml2IdentityProviderDTO) (api.ExternalSaml2IdentityProviderDTO, error) {
	var result api.ExternalSaml2IdentityProviderDTO
	l := c.Logger()

	l.Debugf("createIdPSaml2 : %s [%s]", *extIdP.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdPSaml2") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	IdPSaml2(&extIdP)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdpSaml2(ctx)
	req = req.StoreIdPSaml2Req(api.StoreIdPSaml2Req{IdOrName: &ida, Idp: &extIdP})
	res, _, err := c.apiClient.DefaultApi.CreateIdpSaml2Execute(req)
	if err != nil {
		c.logger.Errorf("CreateIdPSaml2. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateIdPSaml2. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdP received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateIdPSaml2(ida string, IdP api.ExternalSaml2IdentityProviderDTO) (api.ExternalSaml2IdentityProviderDTO, error) {
	var result api.ExternalSaml2IdentityProviderDTO
	l := c.Logger()

	l.Debugf("UpdateIdPSaml2. : %s [%s]", *IdP.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdPSaml2") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	IdPSaml2(&IdP)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdpSaml2(ctx)
	req = req.StoreIdPSaml2Req(api.StoreIdPSaml2Req{IdOrName: &ida, Idp: &IdP})
	res, _, err := c.apiClient.DefaultApi.UpdateIdpSaml2Execute(req)
	if err != nil {
		c.logger.Errorf("UpdateIdPSaml2. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateIdPSaml2. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdP received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteIdPSaml2(ida string, IdP string) (bool, error) {
	c.logger.Debugf("deleteIdPSaml2. %s [%s]", IdP, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteIdPSaml2") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdPSaml2. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdpSaml2(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &IdP})
	res, _, err := c.apiClient.DefaultApi.DeleteIdpSaml2Execute(req)

	if err != nil {
		c.logger.Errorf("deleteIdP. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdP. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdP. Deleted %s : %t", IdP, *res.Removed)

	return *res.Removed, err
}

// Gets an IdP based on the appliance name and IdP name
func (c *IdbusApiClient) GetIdPSaml2(ida string, IdP string) (api.ExternalSaml2IdentityProviderDTO, error) {
	c.logger.Debugf("GetIdPSaml2. %s [%s]", IdP, ida)
	var result api.ExternalSaml2IdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdPSaml2") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpSaml2(ctx)
	req = req.GetIdPSaml2Req(api.GetIdPSaml2Req{IdOrName: &ida, Name: &IdP})
	res, _, err := c.apiClient.DefaultApi.GetIdpSaml2Execute(req)
	if err != nil {
		c.logger.Errorf("GetIdPSaml2. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetIdPSaml2. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("GetIdPSaml2. NOT FOUND %s", IdP)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("GetIdPSaml2. %s found for ID/name %s", *result.Name, IdP)
	} else {
		c.logger.Debugf("GetIdPSaml2. not found for ID/name %s", IdP)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdPSaml2s(ida string) ([]api.ExternalSaml2IdentityProviderDTO, error) {
	c.logger.Debugf("get IdPSaml2s: all [%s]", ida)
	var result []api.ExternalSaml2IdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdPSaml2s") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpSaml2s(ctx)
	req = req.GetIdPSaml2Req(api.GetIdPSaml2Req{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdpSaml2sExecute(req)
	if err != nil {
		c.logger.Errorf("getIdPSaml2s. Error %v", err)
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

func IdPSaml2(IdPSaml2 *api.ExternalSaml2IdentityProviderDTO) {
	IdPSaml2.AdditionalProperties = make(map[string]interface{})
	IdPSaml2.AdditionalProperties["@c"] = ".ExternalSaml2IdentityProviderDTO"
}
