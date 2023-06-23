package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the IIS dto to use as template
func (c *IdbusApiClient) CreateIISExeEnv(ida string, iss api.WindowsIISExecutionEnvironmentDTO) (api.WindowsIISExecutionEnvironmentDTO, error) {
	var result api.WindowsIISExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("createIISExeEnv : %s [%s]", *iss.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIISExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIIS(&iss)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIisExecEnv(ctx)
	req = req.StoreIisExecEnvReq(api.StoreIisExecEnvReq{IdOrName: &ida, IisExecEnv: &iss})
	res, _, err := c.apiClient.DefaultApi.CreateIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("CreateIISExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateIISExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IisExecEnv == nil {
		return result, errors.New("no IISExeEnv received after creation")
	}

	result = *res.IisExecEnv

	return result, nil
}

func (c *IdbusApiClient) UpdateIISExeEnv(ida string, iss api.WindowsIISExecutionEnvironmentDTO) (api.WindowsIISExecutionEnvironmentDTO, error) {
	var result api.WindowsIISExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("updateIISExeEnv. : %s [%s]", *iss.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIISExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIIS(&iss)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIisExecEnv(ctx)
	req = req.StoreIisExecEnvReq(api.StoreIisExecEnvReq{IdOrName: &ida, IisExecEnv: &iss})
	res, _, err := c.apiClient.DefaultApi.UpdateIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("updateIISExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIISExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IisExecEnv == nil {
		return result, errors.New("no iss received after update")
	}

	result = *res.IisExecEnv

	return result, nil
}

func (c *IdbusApiClient) DeleteIISExeEnv(ida string, IIS string) (bool, error) {
	c.logger.Debugf("deleteIISExeEnv. %s [%s]", IIS, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.Deletev") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIISExeEnv. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIisExecEnv(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &IIS})
	res, _, err := c.apiClient.DefaultApi.DeleteIisExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIISExeEnv. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIISExeEnv. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIISExeEnv. Deleted %s : %t", IIS, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIISExeEnv(ida string, IIS string) (api.WindowsIISExecutionEnvironmentDTO, error) {
	c.logger.Debugf("getIIS. %s [%s]", IIS, ida)
	var result api.WindowsIISExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.getIIS") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIisExecEnv(ctx)
	req = req.GetIisExecEnvReq(api.GetIisExecEnvReq{IdOrName: &ida, Name: &IIS})
	res, _, err := c.apiClient.DefaultApi.GetIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getIIS. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIIS. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IisExecEnv == nil {
		c.logger.Debugf("getIIS. NOT FOUND %s", IIS)
		return result, nil
	}

	if res.IisExecEnv != nil {
		result = *res.IisExecEnv
		c.logger.Debugf("getIIS. %s found for ID/name %s", *result.Name, IIS)
	} else {
		c.logger.Debugf("getIIS. not found for ID/name %s", IIS)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIISExeEnvs(ida string) ([]api.WindowsIISExecutionEnvironmentDTO, error) {
	c.logger.Debugf("get IIS: all [%s]", ida)
	var result []api.WindowsIISExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIIS") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIisExecEnvs(ctx)
	req = req.GetIisExecEnvReq(api.GetIisExecEnvReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIisExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("getIISExeEnvs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.IisExecEnv == nil {
		return result, nil
	}

	result = res.IisExecEnv

	return result, nil

}

func initIIS(IIS *api.WindowsIISExecutionEnvironmentDTO) {
	IIS.AdditionalProperties = make(map[string]interface{})
	IIS.AdditionalProperties["@c"] = ".WindowsIISExecutionEnvironmentDTO"
}
