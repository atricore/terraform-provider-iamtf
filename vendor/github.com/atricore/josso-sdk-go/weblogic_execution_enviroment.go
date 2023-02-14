package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the wlogic dto to use as template
func (c *IdbusApiClient) CreateWLogic(ida string, wlogic api.WeblogicExecutionEnvironmentDTO) (api.WeblogicExecutionEnvironmentDTO, error) {
	var result api.WeblogicExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("createwlogicExeEnv : %s [%s]", *wlogic.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreatewlogicExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initWLogic(&wlogic)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateWeblogicExecEnv(ctx)
	req = req.StoreWeblogicExecEnvReq(api.StoreWeblogicExecEnvReq{IdOrName: &ida, ExecEnv: &wlogic})
	res, _, err := c.apiClient.DefaultApi.CreateWeblogicExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("CreatewlogicExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreatewlogicExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.ExecEnv == nil {
		return result, errors.New("no wlogicExeEnv received after creation")
	}

	result = *res.ExecEnv

	return result, nil
}

func (c *IdbusApiClient) UpdateWLogic(ida string, wlogic api.WeblogicExecutionEnvironmentDTO) (api.WeblogicExecutionEnvironmentDTO, error) {
	var result api.WeblogicExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("updatewlogicExeEnv. : %s [%s]", *wlogic.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdatewlogicExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initWLogic(&wlogic)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateWeblogicExecEnv(ctx)
	req = req.StoreWeblogicExecEnvReq(api.StoreWeblogicExecEnvReq{IdOrName: &ida, ExecEnv: &wlogic})
	res, _, err := c.apiClient.DefaultApi.UpdateWeblogicExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("updatewlogicExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updatewlogicExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.ExecEnv == nil {
		return result, errors.New("no wlogic received after update")
	}

	result = *res.ExecEnv

	return result, nil
}

func (c *IdbusApiClient) DeleteWLogic(ida string, wlogic string) (bool, error) {
	c.logger.Debugf("deletewlogicExeEnv. %s [%s]", wlogic, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.Deletev") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deletewlogicExeEnv. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteWeblogicExecEnv(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &wlogic})
	res, _, err := c.apiClient.DefaultApi.DeleteWeblogicExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("deletewlogicExeEnv. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deletewlogicExeEnv. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deletewlogicExeEnv. Deleted %s : %t", wlogic, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetWebLogic(ida string, wlogic string) (api.WeblogicExecutionEnvironmentDTO, error) {
	c.logger.Debugf("getwlogic. %s [%s]", wlogic, ida)
	var result api.WeblogicExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.getwlogic") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetWeblogicExecEnv(ctx)
	req = req.GetWeblogicExecEnvReq(api.GetWeblogicExecEnvReq{IdOrName: &ida, Name: &wlogic})
	res, _, err := c.apiClient.DefaultApi.GetWeblogicExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getwlogic. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getwlogic. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.ExecEnv == nil {
		c.logger.Debugf("getwlogic. NOT FOUND %s", wlogic)
		return result, nil
	}

	if res.ExecEnv != nil {
		result = *res.ExecEnv
		c.logger.Debugf("getwlogic. %s found for ID/name %s", *result.Name, wlogic)
	} else {
		c.logger.Debugf("getwlogic. not found for ID/name %s", wlogic)
	}

	return result, nil

}

func (c *IdbusApiClient) GetWebLogics(ida string) ([]api.WeblogicExecutionEnvironmentDTO, error) {
	c.logger.Debugf("get wlogic: all [%s]", ida)
	var result []api.WeblogicExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.Getwlogic") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetWeblogicExecEnvs(ctx)
	req = req.GetWeblogicExecEnvReq(api.GetWeblogicExecEnvReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetWeblogicExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("getwlogicExeEnvs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.ExecEnv == nil {
		return result, nil
	}

	result = res.ExecEnv

	return result, nil

}

func initWLogic(WLogic *api.WeblogicExecutionEnvironmentDTO) {
	WLogic.AdditionalProperties = make(map[string]interface{})
	WLogic.AdditionalProperties["@c"] = ".WeblogicExecutionEnvironmentDTO"
}
