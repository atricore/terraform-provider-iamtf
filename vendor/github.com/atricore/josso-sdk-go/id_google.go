package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateIdpGoogle(ida string, IdGoogle api.GoogleOpenIDConnectIdentityProviderDTO) (api.GoogleOpenIDConnectIdentityProviderDTO, error) {
	var result api.GoogleOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("createIdpGoogle : %s [%s]", *IdGoogle.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdGoogle") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdGoogle(&IdGoogle)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdpGoogle(ctx)
	req = req.StoreIdpGoogleReq(api.StoreIdpGoogleReq{IdOrName: &ida, Idp: &IdGoogle})
	res, _, err := c.apiClient.DefaultApi.CreateIdpGoogleExecute(req)
	if err != nil {
		c.logger.Errorf("createIdpGoogle. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdpGoogle. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdGoogle received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateIdpGoogle(ida string, IdGoogle api.GoogleOpenIDConnectIdentityProviderDTO) (api.GoogleOpenIDConnectIdentityProviderDTO, error) {
	var result api.GoogleOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("updateIdGoogle. : %s [%s]", *IdGoogle.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdpGoogle") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdGoogle(&IdGoogle)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdpGoogle(ctx)
	req = req.StoreIdpGoogleReq(api.StoreIdpGoogleReq{IdOrName: &ida, Idp: &IdGoogle})
	res, _, err := c.apiClient.DefaultApi.UpdateIdpGoogleExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdGoogle. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdGoogle. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdGoogle received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteIdpGoogle(ida string, IdGoogle string) (bool, error) {
	c.logger.Debugf("deleteIdGoogle. %s [%s]", IdGoogle, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdGoogle") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdGoogle. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdpGoogle(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &IdGoogle})
	res, _, err := c.apiClient.DefaultApi.DeleteIdpGoogleExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdGoogle. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdGoogle. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdGoogle. Deleted %s : %t", IdGoogle, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIdpGoogle(ida string, IdGoogle string) (api.GoogleOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("getIdGoogle. %s [%s]", IdGoogle, ida)
	var result api.GoogleOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdGoogle") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpGoogle(ctx)
	req = req.GetIdpGoogleReq(api.GetIdpGoogleReq{IdOrName: &ida, Name: &IdGoogle})
	res, _, err := c.apiClient.DefaultApi.GetIdpGoogleExecute(req)
	if err != nil {
		c.logger.Errorf("getIdGoogle. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdGoogle. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("getIdGoogle. NOT FOUND %s", IdGoogle)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("getIdGoogle. %s found for ID/name %s", *result.Name, IdGoogle)
	} else {
		c.logger.Debugf("getIdGoogle. not found for ID/name %s", IdGoogle)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdpGoogles(ida string) ([]api.GoogleOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("get IdGoogles: all [%s]", ida)
	var result []api.GoogleOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIIdpGoogles") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpGoogles(ctx)
	req = req.GetIdpGoogleReq(api.GetIdpGoogleReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdpGooglesExecute(req)
	if err != nil {
		c.logger.Errorf("getIIdpGoogles. Error %v", err)
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

func initIdGoogle(IdGoogle *api.GoogleOpenIDConnectIdentityProviderDTO) {
	IdGoogle.AdditionalProperties = make(map[string]interface{})
	IdGoogle.AdditionalProperties["@c"] = ".GoogleOpenIDConnectIdentityProviderDTO"
}
