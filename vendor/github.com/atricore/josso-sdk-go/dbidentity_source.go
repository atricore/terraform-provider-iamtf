package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new DbIdentitySourceDTO in the provided identity appliance. It receives the appliance name or id and the Db dto to use as template
func (c *IdbusApiClient) CreateDbIdentitySourceDTO(ida string, intDbSource api.DbIdentitySourceDTO) (api.DbIdentitySourceDTO, error) {
	var result api.DbIdentitySourceDTO
	l := c.Logger()

	l.Debugf("CreateDbIdentitySource : %s [%s]", *intDbSource.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateDbIdentitySource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initDbIdentitySourceDTO(&intDbSource)

	//ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdSourceDb(ctx)
	req = req.StoreIdSourceDbReq(api.StoreIdSourceDbReq{IdOrName: &ida, IdSourceDb: &intDbSource})
	res, _, err := c.apiClient.DefaultApi.CreateIdSourceDbExecute(req)
	if err != nil {
		c.logger.Errorf("CreateDbIdentitySource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateDbIdentitySource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdSourceDb == nil {
		return result, errors.New("no DbIdentitySourceDTO received after creation")
	}

	result = *res.IdSourceDb

	return result, nil
}

func (c *IdbusApiClient) UpdateDbIdentitySourceDTO(ida string, intDbSource api.DbIdentitySourceDTO) (api.DbIdentitySourceDTO, error) {
	var result api.DbIdentitySourceDTO
	l := c.Logger()

	l.Debugf("UpdateDbIdentitySourceDTO. : %s [%s]", *intDbSource.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateDbIdentitySourceDTO") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initDbIdentitySourceDTO(&intDbSource)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdSourceDb(ctx)
	req = req.StoreIdSourceDbReq(api.StoreIdSourceDbReq{IdOrName: &ida, IdSourceDb: &intDbSource})
	res, _, err := c.apiClient.DefaultApi.UpdateIdSourceDbExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateDbIdentitySourceDTO. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateDbIdentitySourceDTO. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdSourceDb == nil {
		return result, errors.New("no DbIdentitySourceDTO received after update")
	}

	result = *res.IdSourceDb

	return result, nil
}

func (c *IdbusApiClient) DeleteDbIdentitySourceDTO(ida string, intDbSource string) (bool, error) {
	c.logger.Debugf("deleteDbIdentitySourceDTO. %s [%s]", intDbSource, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteDbIdentitySourceDTO") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteDbIdentitySourceDTO. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdSourceDb(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &intDbSource})
	res, _, err := c.apiClient.DefaultApi.DeleteIdSourceDbExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIntSaml2Ss. Deleted %s : %t", intDbSource, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and intDbSource name
func (c *IdbusApiClient) GetDbIdentitySourceDTO(ida string, intDbSource string) (api.DbIdentitySourceDTO, error) {
	c.logger.Debugf("GetDbIdentitySourceDTO. %s [%s]", intDbSource, ida)
	var result api.DbIdentitySourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetDbIdentitySourceDTO") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSourceDb(ctx)
	req = req.GetIdSourceDbReq(api.GetIdSourceDbReq{IdOrName: &ida, Name: &intDbSource})
	res, _, err := c.apiClient.DefaultApi.GetIdSourceDbExecute(req)
	if err != nil {
		c.logger.Errorf("GetDbIdentitySourceDTO. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetDbIdentitySourceDTO. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IdSourceDb == nil {
		c.logger.Debugf("GetDbIdentitySourceDTO. NOT FOUND %s", intDbSource)
		return result, nil
	}

	if res.IdSourceDb != nil {
		result = *res.IdSourceDb
		c.logger.Debugf("GetDbIdentitySourceDTO. %s found for ID/name %s", *result.Name, intDbSource)
	} else {
		c.logger.Debugf("GetDbIdentitySourceDTO. not found for ID/name %s", intDbSource)
	}

	return result, nil

}

func (c *IdbusApiClient) GetDbIdentitySourceDTOs(ida string) ([]api.DbIdentitySourceDTO, error) {
	c.logger.Debugf("get DbIdentitySourceDTOs: all [%s]", ida)
	var result []api.DbIdentitySourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetDbIdentitySourceDTOs") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSourceDbs(ctx)
	req = req.GetIdSourceDbReq(api.GetIdSourceDbReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdSourceDbsExecute(req)
	if err != nil {
		c.logger.Errorf("getDbIdentitySourceDTOs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.IdSourceDbs == nil {
		return result, nil
	}

	result = res.IdSourceDbs

	return result, nil

}

func initDbIdentitySourceDTO(DbIdentitySourceDTO *api.DbIdentitySourceDTO) {
	DbIdentitySourceDTO.AdditionalProperties = make(map[string]interface{})
	DbIdentitySourceDTO.AdditionalProperties["@c"] = ".DbIdentitySourceDTO"
}
