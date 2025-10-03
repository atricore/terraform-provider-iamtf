package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateOidcIdp(ida string, oidcIdp api.GenericOpenIDConnectIdentityProviderDTO) (api.GenericOpenIDConnectIdentityProviderDTO, error) {
	var result api.GenericOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("createOidcIdp : %s [%s]", *oidcIdp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateOidcIdp")
	if err != nil {
		return result, err
	}

	initOidcIdp(&oidcIdp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateOidcIdp(ctx)
	req = req.StoreOidcIdpReq(api.StoreOidcIdpReq{IdOrName: &ida, Idp: &oidcIdp})
	res, _, err := c.apiClient.DefaultApi.CreateOidcIdpExecute(req)
	if err != nil {
		c.logger.Errorf("createOidcIdp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createOidcIdp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no OIDC IDP received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateOidcIdp(ida string, oidcIdp api.GenericOpenIDConnectIdentityProviderDTO) (api.GenericOpenIDConnectIdentityProviderDTO, error) {
	var result api.GenericOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("updateOidcIdp. : %s [%s]", *oidcIdp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateOidcIdp")
	if err != nil {
		return result, err
	}

	initOidcIdp(&oidcIdp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateOidcIdp(ctx)
	req = req.StoreOidcIdpReq(api.StoreOidcIdpReq{IdOrName: &ida, Idp: &oidcIdp})
	res, _, err := c.apiClient.DefaultApi.UpdateOidcIdpExecute(req)
	if err != nil {
		c.logger.Errorf("updateOidcIdp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateOidcIdp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no OIDC IDP received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteOidcIdp(ida string, oidcIdp string) (bool, error) {
	c.logger.Debugf("deleteOidcIdp. %s [%s]", oidcIdp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteOidcIdp")
	if err != nil {
		c.logger.Errorf("deleteOidcIdp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteOidcIdp(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &oidcIdp})
	res, _, err := c.apiClient.DefaultApi.DeleteOidcIdpExecute(req)

	if err != nil {
		c.logger.Errorf("deleteOidcIdp. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteOidcIdp. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteOidcIdp. Deleted %s : %t", oidcIdp, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetOidcIdp(ida string, oidcIdp string) (api.GenericOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("getOidcIdp. %s [%s]", oidcIdp, ida)
	var result api.GenericOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetOidcIdp")
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetOidcIdp(ctx)
	req = req.GetOidcIdpReq(api.GetOidcIdpReq{IdOrName: &ida, Name: &oidcIdp})
	res, _, err := c.apiClient.DefaultApi.GetOidcIdpExecute(req)
	if err != nil {
		c.logger.Errorf("getOidcIdp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getOidcIdp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("getOidcIdp. NOT FOUND %s", oidcIdp)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("getOidcIdp. %s found for ID/name %s", *result.Name, oidcIdp)
	} else {
		c.logger.Debugf("getOidcIdp. not found for ID/name %s", oidcIdp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetOidcIdps(ida string) ([]api.GenericOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("get OidcIdps: all [%s]", ida)
	var result []api.GenericOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetOidcIdps")
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetOidcIdps(ctx)
	req = req.GetOidcIdpReq(api.GetOidcIdpReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetOidcIdpsExecute(req)
	if err != nil {
		c.logger.Errorf("getOidcIdps. Error %v", err)
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

func initOidcIdp(oidcIdp *api.GenericOpenIDConnectIdentityProviderDTO) {
	oidcIdp.AdditionalProperties = make(map[string]interface{})
	oidcIdp.AdditionalProperties["@c"] = ".GenericOpenIDConnectIdentityProviderDTO"
}
