package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new SelfServiceResource in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateSelfServiceResource(ida string, SelfService api.SelfServicesResourceDTO) (api.SelfServicesResourceDTO, error) {
	var result api.SelfServicesResourceDTO
	l := c.Logger()

	l.Debugf("CreateSelfServiceResource : %s [%s]", *SelfService.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateSelfServiceResource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initSServiceResource(&SelfService)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateSelfSvcs(ctx)
	req = req.StoreSelfSvcRsReq(api.StoreSelfSvcRsReq{IdOrName: &ida, Resource: &SelfService})
	res, _, err := c.apiClient.DefaultApi.CreateSelfSvcsExecute(req)
	if err != nil {
		c.logger.Errorf("CreateSelfServiceResource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateSelfServiceResource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no SelfServiceResource received after creation")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) UpdateSelfServiceResource(ida string, sp api.SelfServicesResourceDTO) (api.SelfServicesResourceDTO, error) {
	var result api.SelfServicesResourceDTO
	l := c.Logger()

	l.Debugf("UpdateSelfServiceResource. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.updateSelfServiceResource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initSServiceResource(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateSelfSvcs(ctx)
	req = req.StoreSelfSvcRsReq(api.StoreSelfSvcRsReq{IdOrName: &ida, Resource: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateSelfSvcsExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateSelfServiceResource. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateSelfServiceResource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no SelfSeServiceResourcerviceResource received after update")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) DeleteSelfServiceResource(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteSelfServiceResource. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteSelfServiceResource") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteSelfServiceResource. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteSelfSvcs(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteSelfSvcsExecute(req)

	if err != nil {
		c.logger.Errorf("deleteSelfServiceResource. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteSelfServiceResource. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteSelfServiceResource. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetSelfServiceResource(ida string, resource string) (api.SelfServicesResourceDTO, error) {
	c.logger.Debugf("GetSelfServiceResource. %s [%s]", resource, ida)
	var result api.SelfServicesResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetSelfServiceResource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetSelfSvcs(ctx)
	req = req.GetSelfSvcRsReq(api.GetSelfSvcRsReq{IdOrName: &ida, Name: &resource})
	res, _, err := c.apiClient.DefaultApi.GetSelfSvcsExecute(req)
	if err != nil {
		c.logger.Errorf("GetSelfServiceResource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetSelfServiceResource. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Resource == nil {
		c.logger.Debugf("GetSelfServiceResource. NOT FOUND %s", resource)
		return result, nil
	}

	if res.Resource != nil {
		result = *res.Resource
		c.logger.Debugf("GetSelfServiceResource. %s found for ID/name %s", *result.Name, resource)
	} else {
		c.logger.Debugf("GetSelfServiceResource. not found for ID/name %s", resource)
	}

	return result, nil

}

func (c *IdbusApiClient) GetSelfServiceResources(ida string) ([]api.SelfServicesResourceDTO, error) {
	c.logger.Debugf("get SelfServiceResources: all [%s]", ida)
	var result []api.SelfServicesResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetSelfServiceResources") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetSelfSvcss(ctx)
	req = req.GetSelfSvcRsReq(api.GetSelfSvcRsReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetSelfSvcssExecute(req)
	if err != nil {
		c.logger.Errorf("getSelfServiceResources. Error %v", err)
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

func initSServiceResource(SelfServiceResource *api.SelfServicesResourceDTO) {
	SelfServiceResource.AdditionalProperties = make(map[string]interface{})
	SelfServiceResource.AdditionalProperties["@c"] = ".SelfServicesResourceDTO"
}
