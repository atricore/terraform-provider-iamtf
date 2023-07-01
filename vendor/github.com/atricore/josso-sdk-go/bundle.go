package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) GetOSGiBundles() ([]api.BundleDescr, error) {
	c.logger.Debug("get OSGiBundles: all")
	var result []api.BundleDescr

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetBundles") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetBundles(ctx)
	req = req.GetBundlesReq(api.GetBundlesReq{})
	res, _, err := c.apiClient.DefaultApi.GetBundlesExecute(req)
	if err != nil {
		c.logger.Errorf("GetBundlesReq. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Bundles == nil {
		return result, nil
	}

	result = res.Bundles

	return result, nil

}
