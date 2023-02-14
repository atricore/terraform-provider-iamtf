package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new SharePointResource in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateSharePointresource(ida string, sharepoint api.SharepointResourceDTO) (api.SharepointResourceDTO, error) {
	var result api.SharepointResourceDTO
	l := c.Logger()

	l.Debugf("CreateSharePointresource : %s [%s]", *sharepoint.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateSharePointresource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initSPointResource(&sharepoint)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateSharepointRs(ctx)
	req = req.StoreSharepointRsReq(api.StoreSharepointRsReq{IdOrName: &ida, Resource: &sharepoint})
	res, _, err := c.apiClient.DefaultApi.CreateSharepointRsExecute(req)
	if err != nil {
		c.logger.Errorf("CreateSharePointresource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateSharePointresource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no SharePointResource received after creation")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) UpdateSharePointResource(ida string, sp api.SharepointResourceDTO) (api.SharepointResourceDTO, error) {
	var result api.SharepointResourceDTO
	l := c.Logger()

	l.Debugf("UpdateSharePointResource. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.updateSharePointResource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initSPointResource(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateSharepointRs(ctx)
	req = req.StoreSharepointRsReq(api.StoreSharepointRsReq{IdOrName: &ida, Resource: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateSharepointRsExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateSharePointResource. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateSharePointResource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no SharePointResource received after update")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) DeleteSharePointResource(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteSharePointResource. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteSharePointResource") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteSharePointResource. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteSharepointRs(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteSharepointRsExecute(req)

	if err != nil {
		c.logger.Errorf("deleteSharePointResource. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteSharePointResource. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteSharePointResource. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetSharePointResource(ida string, resource string) (api.SharepointResourceDTO, error) {
	c.logger.Debugf("GetSharePointResource. %s [%s]", resource, ida)
	var result api.SharepointResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetSharePointResource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetSharepointRs(ctx)
	req = req.GetSharepointRsReq(api.GetSharepointRsReq{IdOrName: &ida, Name: &resource})
	res, _, err := c.apiClient.DefaultApi.GetSharepointRsExecute(req)
	if err != nil {
		c.logger.Errorf("GetSharePointResource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetSharePointResource. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Resource == nil {
		c.logger.Debugf("GetSharePointResource. NOT FOUND %s", resource)
		return result, nil
	}

	if res.Resource != nil {
		result = *res.Resource
		c.logger.Debugf("GetSharePointResource. %s found for ID/name %s", *result.Name, resource)
	} else {
		c.logger.Debugf("GetSharePointResource. not found for ID/name %s", resource)
	}

	return result, nil

}

func (c *IdbusApiClient) GetSharePointResources(ida string) ([]api.SharepointResourceDTO, error) {
	c.logger.Debugf("get SharePointResources: all [%s]", ida)
	var result []api.SharepointResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetSharePointResources") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetSharepointRss(ctx)
	req = req.GetSharepointRsReq(api.GetSharepointRsReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetSharepointRssExecute(req)
	if err != nil {
		c.logger.Errorf("getSharePointResources. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Resources == nil {
		return result, nil
	}

	result = res.Resources

	return result, nil

}

func initSPointResource(SharePointResource *api.SharepointResourceDTO) {
	SharePointResource.AdditionalProperties = make(map[string]interface{})
	SharePointResource.AdditionalProperties["@c"] = ".SharepointResourceDTO"
}
