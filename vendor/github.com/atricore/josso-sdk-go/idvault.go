package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateIdVault(ida string, idVault api.EmbeddedIdentityVaultDTO) (api.EmbeddedIdentityVaultDTO, error) {
	var result api.EmbeddedIdentityVaultDTO
	l := c.Logger()

	l.Debugf("createIdVault : %s [%s]", *idVault.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdVault") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdVault(&idVault)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdVault(ctx)
	req = req.StoreIdVaultReq(api.StoreIdVaultReq{IdOrName: &ida, IdVault: &idVault})
	res, _, err := c.apiClient.DefaultApi.CreateIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("createIdVault. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdVault. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdVault == nil {
		return result, errors.New("no idVault received after creation")
	}

	result = *res.IdVault

	return result, nil
}

func (c *IdbusApiClient) UpdateIdVault(ida string, idVault api.EmbeddedIdentityVaultDTO) (api.EmbeddedIdentityVaultDTO, error) {
	var result api.EmbeddedIdentityVaultDTO
	l := c.Logger()

	l.Debugf("updateIdVault. : %s [%s]", *idVault.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdVault") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdVault(&idVault)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdVault(ctx)
	req = req.StoreIdVaultReq(api.StoreIdVaultReq{IdOrName: &ida, IdVault: &idVault})
	res, _, err := c.apiClient.DefaultApi.UpdateIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdVault. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdVault. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdVault == nil {
		return result, errors.New("no idVault received after update")
	}

	result = *res.IdVault

	return result, nil
}

func (c *IdbusApiClient) DeleteIdVault(ida string, idVault string) (bool, error) {
	c.logger.Debugf("deleteIdVault. %s [%s]", idVault, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdVault") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdVault. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdVault(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &idVault})
	res, _, err := c.apiClient.DefaultApi.DeleteIdVaultExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdVault. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdVault. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdVault. Deleted %s : %t", idVault, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIdVault(ida string, idVault string) (api.EmbeddedIdentityVaultDTO, error) {
	c.logger.Debugf("getIdVault. %s [%s]", idVault, ida)
	var result api.EmbeddedIdentityVaultDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdVault") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdVault(ctx)
	req = req.GetIdVaultReq(api.GetIdVaultReq{IdOrName: &ida, Name: &idVault})
	res, _, err := c.apiClient.DefaultApi.GetIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("getIdVault. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdVault. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IdVault == nil {
		c.logger.Debugf("getIdVault. NOT FOUND %s", idVault)
		return result, nil
	}

	if res.IdVault != nil {
		result = *res.IdVault
		c.logger.Debugf("getIdVault. %s found for ID/name %s", *result.Name, idVault)
	} else {
		c.logger.Debugf("getIdVault. not found for ID/name %s", idVault)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdVaults(ida string) ([]api.EmbeddedIdentityVaultDTO, error) {
	c.logger.Debugf("get idVaults: all [%s]", ida)
	var result []api.EmbeddedIdentityVaultDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdVaults") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdVaults(ctx)
	req = req.GetIdVaultReq(api.GetIdVaultReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdVaultsExecute(req)
	if err != nil {
		c.logger.Errorf("getIdVaults. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.IdVaults == nil {
		return result, nil
	}

	result = res.IdVaults

	return result, nil

}

func initIdVault(idVault *api.EmbeddedIdentityVaultDTO) {
	idVault.AdditionalProperties = make(map[string]interface{})
	idVault.AdditionalProperties["@c"] = ".EmbeddedIdentityVaultDTO"
}
