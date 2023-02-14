package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateIdFacebook(ida string, IdFacebook api.FacebookOpenIDConnectIdentityProviderDTO) (api.FacebookOpenIDConnectIdentityProviderDTO, error) {
	var result api.FacebookOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("createIdFacebook : %s [%s]", *IdFacebook.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdFacebook") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdFacebook(&IdFacebook)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdpFacebook(ctx)
	req = req.StoreIdpFacebookReq(api.StoreIdpFacebookReq{IdOrName: &ida, Idp: &IdFacebook})
	res, _, err := c.apiClient.DefaultApi.CreateIdpFacebookExecute(req)
	if err != nil {
		c.logger.Errorf("createIdFacebook. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdFacebook. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdFacebook received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateIdpFacebook(ida string, IdFacebook api.FacebookOpenIDConnectIdentityProviderDTO) (api.FacebookOpenIDConnectIdentityProviderDTO, error) {
	var result api.FacebookOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("updateIdFacebook. : %s [%s]", *IdFacebook.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdpFacebook") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdFacebook(&IdFacebook)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdpFacebook(ctx)
	req = req.StoreIdpFacebookReq(api.StoreIdpFacebookReq{IdOrName: &ida, Idp: &IdFacebook})
	res, _, err := c.apiClient.DefaultApi.UpdateIdpFacebookExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdFacebook. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdFacebook. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdFacebook received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteIdpFacebook(ida string, IdFacebook string) (bool, error) {
	c.logger.Debugf("deleteIdFacebook. %s [%s]", IdFacebook, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdFacebook") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdFacebook. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdpFacebook(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &IdFacebook})
	res, _, err := c.apiClient.DefaultApi.DeleteIdpFacebookExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdFacebook. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdFacebook. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdFacebook. Deleted %s : %t", IdFacebook, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIdpFacebook(ida string, IdFacebook string) (api.FacebookOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("getIdFacebook. %s [%s]", IdFacebook, ida)
	var result api.FacebookOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdFacebook") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpFacebook(ctx)
	req = req.GetIdpFacebookReq(api.GetIdpFacebookReq{IdOrName: &ida, Name: &IdFacebook})
	res, _, err := c.apiClient.DefaultApi.GetIdpFacebookExecute(req)
	if err != nil {
		c.logger.Errorf("getIdFacebook. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdFacebook. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("getIdFacebook. NOT FOUND %s", IdFacebook)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("getIdFacebook. %s found for ID/name %s", *result.Name, IdFacebook)
	} else {
		c.logger.Debugf("getIdFacebook. not found for ID/name %s", IdFacebook)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdpFacebooks(ida string) ([]api.FacebookOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("get IdFacebooks: all [%s]", ida)
	var result []api.FacebookOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdpFacebooks") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpFacebooks(ctx)
	req = req.GetIdpFacebookReq(api.GetIdpFacebookReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdpFacebooksExecute(req)
	if err != nil {
		c.logger.Errorf("getIdpFacebooks. Error %v", err)
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

func initIdFacebook(IdFacebook *api.FacebookOpenIDConnectIdentityProviderDTO) {
	IdFacebook.AdditionalProperties = make(map[string]interface{})
	IdFacebook.AdditionalProperties["@c"] = ".FacebookOpenIDConnectIdentityProviderDTO"
}
