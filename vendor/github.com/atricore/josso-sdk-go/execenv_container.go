package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Gets an IdP based on the appliance name and idp name
func (c *IdbusApiClient) GetExecEnv(ida string, idsource string) (api.ExecEnvContainerDTO, error) {
	c.logger.Debugf("getExecEnv. %s [%s]", idsource, ida)
	var result api.ExecEnvContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetExecEnv") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetExecEnv(ctx)
	req = req.GetExecEnvReq(api.GetExecEnvReq{IdOrName: &ida, Name: &idsource})
	res, _, err := c.apiClient.DefaultApi.GetExecEnvExecute(req)
	if err != nil {
		c.logger.Errorf("getExecEnv. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getExecEnv. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.ExecEnv == nil {
		c.logger.Debugf("getExecEnv. NOT FOUND %s", idsource)
		return result, nil
	}

	result = res.GetExecEnv()

	return result, nil

}

func (c *IdbusApiClient) GetExecEnvs(ida string) ([]api.ExecEnvContainerDTO, error) {
	c.logger.Debugf("getExecEnvs: all [%s]", ida)
	var result []api.ExecEnvContainerDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetExecEnvs") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetExecEnvs(ctx)
	req = req.GetApplianceReq(api.GetApplianceReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetExecEnvsExecute(req)
	if err != nil {
		c.logger.Errorf("getExecEnvs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.ExecEnvs == nil {
		return result, nil
	}

	result = res.ExecEnvs

	return result, nil

}
