package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new DbIdentityVaultDto in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateDbIdentityVault(ida string, intDbVault api.DbIdentityVaultDTO) (api.DbIdentityVaultDTO, error) {
	var result api.DbIdentityVaultDTO
	l := c.Logger()

	l.Debugf("CreateDbIdentityVault : %s [%s]", *intDbVault.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateDbIdentityVault") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initDbIdentityVaultDto(&intDbVault)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateDbIdVault(ctx)
	req = req.StoreDbIdVaultReq(api.StoreDbIdVaultReq{IdOrName: &ida, DbIdVault: &intDbVault})
	res, _, err := c.apiClient.DefaultApi.CreateDbIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("CreateDbIdentityVault. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateDbIdentityVault. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.DbIdVault == nil {
		return result, errors.New("no DbIdentityVaultDto received after creation")
	}

	result = *res.DbIdVault

	return result, nil
}

func (c *IdbusApiClient) UpdateDbIdentityVaultDTO(ida string, intDbVault api.DbIdentityVaultDTO) (api.DbIdentityVaultDTO, error) {
	var result api.DbIdentityVaultDTO
	l := c.Logger()

	l.Debugf("UpdateDbIdentityVaultDto. : %s [%s]", *intDbVault.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateDbIdentityVaultDTO") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initDbIdentityVaultDto(&intDbVault)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateDbIdVault(ctx)
	req = req.StoreDbIdVaultReq(api.StoreDbIdVaultReq{IdOrName: &ida, DbIdVault: &intDbVault})
	res, _, err := c.apiClient.DefaultApi.UpdateDbIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateDbIdentityVaultDto. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateDbIdentityVaultDto. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.DbIdVault == nil {
		return result, errors.New("no DbIdentityVaultDto received after update")
	}

	result = intDbVault

	return result, nil
}

func (c *IdbusApiClient) DeleteDbIdentityVaultDto(ida string, intDbVault string) (bool, error) {
	c.logger.Debugf("deleteDbIdentityVaultDto. %s [%s]", intDbVault, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteDbIdentityVaultDto") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteDbIdentityVaultDto. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteDbIdVault(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &intDbVault})
	res, _, err := c.apiClient.DefaultApi.DeleteDbIdVaultExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIntSaml2Ss. Deleted %s : %t", intDbVault, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and intDbVault name
func (c *IdbusApiClient) GetDbIdentityVaultDto(ida string, intDbVault string) (api.DbIdentityVaultDTO, error) {
	c.logger.Debugf("GetDbIdentityVaultDto. %s [%s]", intDbVault, ida)
	var result api.DbIdentityVaultDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetDbIdentityVaultDto") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetDbIdVault(ctx)
	req = req.GetDbIdVaultReq(api.GetDbIdVaultReq{IdOrName: &ida, Name: &intDbVault})
	res, _, err := c.apiClient.DefaultApi.GetDbIdVaultExecute(req)
	if err != nil {
		c.logger.Errorf("GetDbIdentityVaultDto. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetDbIdentityVaultDto. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.DbIdVault == nil {
		c.logger.Debugf("GetDbIdentityVaultDto. NOT FOUND %s", intDbVault)
		return result, nil
	}

	if res.DbIdVault != nil {
		result = *res.DbIdVault
		c.logger.Debugf("GetDbIdentityVaultDto. %s found for ID/name %s", *result.Name, intDbVault)
	} else {
		c.logger.Debugf("GetDbIdentityVaultDto. not found for ID/name %s", intDbVault)
	}

	return result, nil

}

func (c *IdbusApiClient) GetDbIdentityVaultDtos(ida string) ([]api.DbIdentityVaultDTO, error) {
	c.logger.Debugf("get DbIdentityVaultDtos: all [%s]", ida)
	var result []api.DbIdentityVaultDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetDbIdentityVaultDtos") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetDbIdVaults(ctx)
	req = req.GetDbIdVaultReq(api.GetDbIdVaultReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetDbIdVaultsExecute(req)
	if err != nil {
		c.logger.Errorf("getDbIdentityVaultDtos. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.DbIdVaults == nil {
		return result, nil
	}

	result = res.DbIdVaults

	return result, nil

}

func initDbIdentityVaultDto(DbIdentityVaultDto *api.DbIdentityVaultDTO) {
	DbIdentityVaultDto.AdditionalProperties = make(map[string]interface{})
	DbIdentityVaultDto.AdditionalProperties["@c"] = ".DbIdentityVaultDTO"
}
