package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the Php dto to use as template
func (c *IdbusApiClient) CreatePhpExeEnv(ida string, php api.PHPExecutionEnvironmentDTO) (api.PHPExecutionEnvironmentDTO, error) {
	var result api.PHPExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("createPhpExeEnv : %s [%s]", *php.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.createPhpExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initPhp(&php)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreatePhpExecEnv(ctx)
	req = req.StorePhpExecEnvReq(api.StorePhpExecEnvReq{IdOrName: &ida, PhpExecEnv: &php})
	res, _, err := c.apiClient.DefaultApi.CreatePhpExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("CreatePhpExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreatePhpExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.PhpExecEnv == nil {
		return result, errors.New("no PhpExeEnv received after creation")
	}

	result = *res.PhpExecEnv

	return result, nil
}

func (c *IdbusApiClient) UpdatePhpExeEnv(ida string, php api.PHPExecutionEnvironmentDTO) (api.PHPExecutionEnvironmentDTO, error) {
	var result api.PHPExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("updatePhpExeEnv. : %s [%s]", *php.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdatePhpExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initPhp(&php)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdatePhpExecEnv(ctx)
	req = req.StorePhpExecEnvReq(api.StorePhpExecEnvReq{IdOrName: &ida, PhpExecEnv: &php})
	res, _, err := c.apiClient.DefaultApi.UpdatePhpExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("updatePhpExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updatePhpExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.PhpExecEnv == nil {
		return result, errors.New("no php received after update")
	}

	result = *res.PhpExecEnv

	return result, nil
}

func (c *IdbusApiClient) DeletePhpExeEnv(ida string, Php string) (bool, error) {
	c.logger.Debugf("deletePhpExeEnv. %s [%s]", Php, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.Deletev") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deletePhpExeEnv. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeletePhpExecEnv(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &Php})
	res, _, err := c.apiClient.DefaultApi.DeletePhpExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("deletePhpExeEnv. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deletePhpExeEnv. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deletePhpExeEnv. Deleted %s : %t", Php, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetPhpExeEnv(ida string, Php string) (api.PHPExecutionEnvironmentDTO, error) {
	c.logger.Debugf("getPhp. %s [%s]", Php, ida)
	var result api.PHPExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.getPhp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetPhpExecEnv(ctx)
	req = req.GetPhpExecEnvReq(api.GetPhpExecEnvReq{IdOrName: &ida, Name: &Php})
	res, _, err := c.apiClient.DefaultApi.GetPhpExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getPhp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getPhp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.PhpExecEnv == nil {
		c.logger.Debugf("getPhp. NOT FOUND %s", Php)
		return result, nil
	}

	if res.PhpExecEnv != nil {
		result = *res.PhpExecEnv
		c.logger.Debugf("getPhp. %s found for ID/name %s", *result.Name, Php)
	} else {
		c.logger.Debugf("getPhp. not found for ID/name %s", Php)
	}

	return result, nil

}

func (c *IdbusApiClient) GetPhpExeEnvs(ida string) ([]api.PHPExecutionEnvironmentDTO, error) {
	c.logger.Debugf("get Php: all [%s]", ida)
	var result []api.PHPExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetPhp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetPhpExecEnvs(ctx)
	req = req.GetPhpExecEnvReq(api.GetPhpExecEnvReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetPhpExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("getPhpExeEnvs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.PhpExecEnv == nil {
		return result, nil
	}

	result = res.PhpExecEnv

	return result, nil

}

func initPhp(Php *api.PHPExecutionEnvironmentDTO) {
	Php.AdditionalProperties = make(map[string]interface{})
	Php.AdditionalProperties["@c"] = ".PHPExecutionEnvironmentDTO"
}
