package cli

import (
	"context"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) GetInfo() (api.GetServerInfoRes, error) {
	c.logger.Debug("get Info: all")
	var result api.GetServerInfoRes

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetInfo") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetInfo(ctx)
	req = req.GetServerInfoReq(api.GetServerInfoReq{})
	res, _, err := c.apiClient.DefaultApi.GetInfoExecute(req)
	if err != nil {
		c.logger.Errorf("GetBundlesReq. Error %v", err)
		return result, err
	}

	return *res, nil

}
