package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateIdpAzure(ida string, IdAzure api.AzureOpenIDConnectIdentityProviderDTO) (api.AzureOpenIDConnectIdentityProviderDTO, error) {
	var result api.AzureOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("createIdpAzure : %s [%s]", *IdAzure.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdAzure") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdAzure(&IdAzure)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdpAzure(ctx)
	req = req.StoreIdpAzureReq(api.StoreIdpAzureReq{IdOrName: &ida, Idp: &IdAzure})
	res, _, err := c.apiClient.DefaultApi.CreateIdpAzureExecute(req)
	if err != nil {
		c.logger.Errorf("createIdpAzure. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdpAzure. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdAzure received after creation")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) UpdateIdpAzure(ida string, IdAzure api.AzureOpenIDConnectIdentityProviderDTO) (api.AzureOpenIDConnectIdentityProviderDTO, error) {
	var result api.AzureOpenIDConnectIdentityProviderDTO
	l := c.Logger()

	l.Debugf("updateIdAzure. : %s [%s]", *IdAzure.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdpAzure") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdAzure(&IdAzure)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdpAzure(ctx)
	req = req.StoreIdpAzureReq(api.StoreIdpAzureReq{IdOrName: &ida, Idp: &IdAzure})
	res, _, err := c.apiClient.DefaultApi.UpdateIdpAzureExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdAzure. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdAzure. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Idp == nil {
		return result, errors.New("no IdAzure received after update")
	}

	result = *res.Idp

	return result, nil
}

func (c *IdbusApiClient) DeleteIdpAzure(ida string, IdAzure string) (bool, error) {
	c.logger.Debugf("deleteIdAzure. %s [%s]", IdAzure, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdAzure") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdAzure. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdpAzure(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &IdAzure})
	res, _, err := c.apiClient.DefaultApi.DeleteIdpAzureExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdAzure. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdAzure. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdAzure. Deleted %s : %t", IdAzure, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIdpAzure(ida string, IdAzure string) (api.AzureOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("getIdAzure. %s [%s]", IdAzure, ida)
	var result api.AzureOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdAzure") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpAzure(ctx)
	req = req.GetIdpAzureReq(api.GetIdpAzureReq{IdOrName: &ida, Name: &IdAzure})
	res, _, err := c.apiClient.DefaultApi.GetIdpAzureExecute(req)
	if err != nil {
		c.logger.Errorf("getIdAzure. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdAzure. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Idp == nil {
		c.logger.Debugf("getIdAzure. NOT FOUND %s", IdAzure)
		return result, nil
	}

	if res.Idp != nil {
		result = *res.Idp
		c.logger.Debugf("getIdAzure. %s found for ID/name %s", *result.Name, IdAzure)
	} else {
		c.logger.Debugf("getIdAzure. not found for ID/name %s", IdAzure)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdpAzures(ida string) ([]api.AzureOpenIDConnectIdentityProviderDTO, error) {
	c.logger.Debugf("get IdAzures: all [%s]", ida)
	var result []api.AzureOpenIDConnectIdentityProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdpAzures") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdpAzures(ctx)
	req = req.GetIdpAzureReq(api.GetIdpAzureReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdpAzuresExecute(req)
	if err != nil {
		c.logger.Errorf("getIdpAzures. Error %v", err)
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

func initIdAzure(IdAzure *api.AzureOpenIDConnectIdentityProviderDTO) {
	IdAzure.AdditionalProperties = make(map[string]interface{})
	IdAzure.AdditionalProperties["@c"] = ".AzureOpenIDConnectIdentityProviderDTO"
}
