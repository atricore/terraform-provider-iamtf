package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new Josso1Resource in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateJossoresource(ida string, jossors api.JOSSO1ResourceDTO) (api.JOSSO1ResourceDTO, error) {
	var result api.JOSSO1ResourceDTO
	l := c.Logger()

	l.Debugf("CreateJossoresource : %s [%s]", *jossors.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateJossoresource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initJosso1Resource(&jossors)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateJossoRs(ctx)
	req = req.StoreJossoRsReq(api.StoreJossoRsReq{IdOrName: &ida, Resource: &jossors})
	res, _, err := c.apiClient.DefaultApi.CreateJossoRsExecute(req)
	if err != nil {
		c.logger.Errorf("CreateJossoresource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateJossoresource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no Josso1Resource received after creation")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) UpdateJosso1Resource(ida string, sp api.JOSSO1ResourceDTO) (api.JOSSO1ResourceDTO, error) {
	var result api.JOSSO1ResourceDTO
	l := c.Logger()

	l.Debugf("UpdateJosso1Resource. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateJosso1Resource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initJosso1Resource(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateJossoRs(ctx)
	req = req.StoreJossoRsReq(api.StoreJossoRsReq{IdOrName: &ida, Resource: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateJossoRsExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateJosso1Resource. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateJosso1Resource. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Resource == nil {
		return result, errors.New("no Josso1Resource received after update")
	}

	result = *res.Resource

	return result, nil
}

func (c *IdbusApiClient) DeleteJosso1Resource(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteJosso1Resource. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteJosso1Resource") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteJosso1Resource. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteJossoRs(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteJossoRsExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIntSaml2Ss. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIntSaml2Ss. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetJosso1Resource(ida string, resource string) (api.JOSSO1ResourceDTO, error) {
	c.logger.Debugf("GetJosso1Resource. %s [%s]", resource, ida)
	var result api.JOSSO1ResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetJosso1Resource") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetJossoRs(ctx)
	req = req.GetJossoRsReq(api.GetJossoRsReq{IdOrName: &ida, Name: &resource})
	res, _, err := c.apiClient.DefaultApi.GetJossoRsExecute(req)
	if err != nil {
		c.logger.Errorf("GetJosso1Resource. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetJosso1Resource. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Resource == nil {
		c.logger.Debugf("GetJosso1Resource. NOT FOUND %s", resource)
		return result, nil
	}

	if res.Resource != nil {
		result = *res.Resource
		c.logger.Debugf("GetJosso1Resource. %s found for ID/name %s", *result.Name, resource)
	} else {
		c.logger.Debugf("GetJosso1Resource. not found for ID/name %s", resource)
	}

	return result, nil

}

func (c *IdbusApiClient) GetJosso1Resources(ida string) ([]api.JOSSO1ResourceDTO, error) {
	c.logger.Debugf("get Josso1Resources: all [%s]", ida)
	var result []api.JOSSO1ResourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetJosso1Resources") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetJossoRss(ctx)
	req = req.GetJossoRsReq(api.GetJossoRsReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetJossoRssExecute(req)
	if err != nil {
		c.logger.Errorf("getJosso1Resources. Error %v", err)
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

func initJosso1Resource(Josso1Resource *api.JOSSO1ResourceDTO) {
	Josso1Resource.AdditionalProperties = make(map[string]interface{})
	Josso1Resource.AdditionalProperties["@c"] = ".JOSSO1ResourceDTO"
}
