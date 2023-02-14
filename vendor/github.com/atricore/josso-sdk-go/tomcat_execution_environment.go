package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the oidcRp dto to use as template
func (c *IdbusApiClient) CreateTomcatExeEnv(ida string, tc api.TomcatExecutionEnvironmentDTO) (api.TomcatExecutionEnvironmentDTO, error) {
	var result api.TomcatExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("createTomcatExeEnv : %s [%s]", *tc.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateTomcatExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initTomCat(&tc)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateTomcatExecEnv(ctx)
	req = req.StoreTomcatExecEnvReq(api.StoreTomcatExecEnvReq{IdOrName: &ida, TomcatExecEnv: &tc})
	res, _, err := c.apiClient.DefaultApi.CreateTomcatExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("createTomcatExeEnv. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createTomcatExeEnv. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.TomcatExecEnv == nil {
		return result, errors.New("no TomcatExeEnv received after creation")
	}

	result = *res.TomcatExecEnv

	return result, nil
}

func (c *IdbusApiClient) UpdateTomcatExeEnv(ida string, Tomcat api.TomcatExecutionEnvironmentDTO) (api.TomcatExecutionEnvironmentDTO, error) {
	var result api.TomcatExecutionEnvironmentDTO
	l := c.Logger()

	l.Debugf("updateTomcatExeEnv. : %s [%s]", *Tomcat.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateTomcatExeEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initTomCat(&Tomcat)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateTomcatExecEnv(ctx)
	req = req.StoreTomcatExecEnvReq(api.StoreTomcatExecEnvReq{IdOrName: &ida, TomcatExecEnv: &Tomcat})
	res, _, err := c.apiClient.DefaultApi.UpdateTomcatExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("updateTomcat. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateTomcat. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.TomcatExecEnv == nil {
		return result, errors.New("no tomcat received after update")
	}

	result = *res.TomcatExecEnv

	return result, nil
}

func (c *IdbusApiClient) DeleteTomcatExeEnv(ida string, Tomcat string) (bool, error) {
	c.logger.Debugf("deleteTomcat. %s [%s]", Tomcat, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.Deletev") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteTomcat. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteTomcatExecEnv(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &Tomcat})
	res, _, err := c.apiClient.DefaultApi.DeleteTomcatExecEnvExecute(req)

	if err != nil {
		c.logger.Errorf("deleteTomcat. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteTomcat. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteTomcat. Deleted %s : %t", Tomcat, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetTomcatExeEnv(ida string, Tomcat string) (api.TomcatExecutionEnvironmentDTO, error) {
	c.logger.Debugf("getTomcat. %s [%s]", Tomcat, ida)
	var result api.TomcatExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetTomcat") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetTomcatExecEnv(ctx)
	req = req.GetTomcatExecEnvReq(api.GetTomcatExecEnvReq{IdOrName: &ida, Name: &Tomcat})
	res, _, err := c.apiClient.DefaultApi.GetTomcatExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getTomcat. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getTomcat. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.TomcatExecEnv == nil {
		c.logger.Debugf("getTomcat. NOT FOUND %s", Tomcat)
		return result, nil
	}

	if res.TomcatExecEnv != nil {
		result = *res.TomcatExecEnv
		c.logger.Debugf("getTomcat. %s found for ID/name %s", *result.Name, Tomcat)
	} else {
		c.logger.Debugf("getTomcat. not found for ID/name %s", Tomcat)
	}

	return result, nil

}

func (c *IdbusApiClient) GetTomcatExeEnvs(ida string) ([]api.TomcatExecutionEnvironmentDTO, error) {
	c.logger.Debugf("get Tomcat: all [%s]", ida)
	var result []api.TomcatExecutionEnvironmentDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetTomcat") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetTomcatExecEnvs(ctx)
	req = req.GetTomcatExecEnvReq(api.GetTomcatExecEnvReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetTomcatExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("GetTomcatExeEnvs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.TomcatExecEnv == nil {
		return result, nil
	}

	result = res.TomcatExecEnv

	return result, nil

}

func initTomCat(Tomcat *api.TomcatExecutionEnvironmentDTO) {
	Tomcat.AdditionalProperties = make(map[string]interface{})
	Tomcat.AdditionalProperties["@c"] = ".TomcatExecutionEnvironmentDTO"
}
