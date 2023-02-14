package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the Iss dto to use as template
func (c *IdbusApiClient) CreateIssExeEnv(ida string, iss api.WindowsIISExecutionEnvironmentDTO) (api.WindowsIISExecutionEnvironmentDTO, error) {
	var result api.WindowsIISExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("createIssExeEnv : %s [%s]", *iss.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIssExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIss(&iss)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIisExecEnv(ctx)
	req = req.StoreIisExecEnvReq(api.StoreIisExecEnvReq{IdOrName: &ida, IisExecEnv: &iss})
	res, _, err := c.apiClient.DefaultApi.CreateIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("CreateIssExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateIssExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IisExecEnv == nil {
		return result, errors.New("no IssExeEnv received after creation")
	}

	result = *res.IisExecEnv

	return result, nil
}

func (c *IdbusApiClient) UpdateIssExeEnv(ida string, iss api.WindowsIISExecutionEnvironmentDTO) (api.WindowsIISExecutionEnvironmentDTO, error) {
	var result api.WindowsIISExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("updateIssExeEnv. : %s [%s]", *iss.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIssExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIss(&iss)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIisExecEnv(ctx)
	req = req.StoreIisExecEnvReq(api.StoreIisExecEnvReq{IdOrName: &ida, IisExecEnv: &iss})
	res, _, err := c.apiClient.DefaultApi.UpdateIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("updateIssExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIssExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IisExecEnv == nil {
		return result, errors.New("no iss received after update")
	}

	result = *res.IisExecEnv

	return result, nil
}

func (c *IdbusApiClient) DeleteIssExeEnv(ida string, Iss string) (bool, error) {
	c.logger.Debugf("deleteIssExeEnv. %s [%s]", Iss, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.Deletev") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIssExeEnv. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIisExecEnv(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &Iss})
	res, _, err := c.apiClient.DefaultApi.DeleteIisExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIssExeEnv. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIssExeEnv. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIssExeEnv. Deleted %s : %t", Iss, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIssExeEnv(ida string, Iss string) (api.WindowsIISExecutionEnvironmentDTO, error) {
	c.logger.Debugf("getIss. %s [%s]", Iss, ida)
	var result api.WindowsIISExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.getIss") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIisExecEnv(ctx)
	req = req.GetIisExecEnvReq(api.GetIisExecEnvReq{IdOrName: &ida, Name: &Iss})
	res, _, err := c.apiClient.DefaultApi.GetIisExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getIss. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIss. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IisExecEnv == nil {
		c.logger.Debugf("getIss. NOT FOUND %s", Iss)
		return result, nil
	}

	if res.IisExecEnv != nil {
		result = *res.IisExecEnv
		c.logger.Debugf("getIss. %s found for ID/name %s", *result.Name, Iss)
	} else {
		c.logger.Debugf("getIss. not found for ID/name %s", Iss)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIssExeEnvs(ida string) ([]api.WindowsIISExecutionEnvironmentDTO, error) {
	c.logger.Debugf("get Iss: all [%s]", ida)
	var result []api.WindowsIISExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIss") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIisExecEnvs(ctx)
	req = req.GetIisExecEnvReq(api.GetIisExecEnvReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIisExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("getIssExeEnvs. Error %v", err)
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

func initIss(Iss *api.WindowsIISExecutionEnvironmentDTO) {
	Iss.AdditionalProperties = make(map[string]interface{})
	Iss.AdditionalProperties["@c"] = ".WindowsIISExecutionEnvironmentDTO"
}
