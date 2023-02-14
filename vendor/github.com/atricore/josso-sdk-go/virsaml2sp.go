package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new VirtSaml2Sp in the provided identity appliance. It receives the appliance name or id and the SP dto to use as template
func (c *IdbusApiClient) CreateVirtSaml2Sp(ida string, virsp api.VirtualSaml2ServiceProviderDTO) (api.VirtualSaml2ServiceProviderDTO, error) {
	var result api.VirtualSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("createVirtSaml2Sp : %s [%s]", *virsp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateVirtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	VirtSaml2Sp(&virsp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateVirtSaml2Sp(ctx)
	req = req.StoreVirtSaml2SpReq(api.StoreVirtSaml2SpReq{IdOrName: &ida, Sp: &virsp})
	res, _, err := c.apiClient.DefaultApi.CreateVirtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("CreateVirtSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("CreateVirtSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no Vir sp received after creation")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) UpdateVirtSaml2Sp(ida string, sp api.VirtualSaml2ServiceProviderDTO) (api.VirtualSaml2ServiceProviderDTO, error) {
	var result api.VirtualSaml2ServiceProviderDTO
	l := c.Logger()

	l.Debugf("UpdateVirtSaml2Sp. : %s [%s]", *sp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	VirtSaml2Sp(&sp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateVirtSaml2Sp(ctx)
	req = req.StoreVirtSaml2SpReq(api.StoreVirtSaml2SpReq{IdOrName: &ida, Sp: &sp})
	res, _, err := c.apiClient.DefaultApi.UpdateVirtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateVirtSaml2Sp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("UpdateVirtSaml2Sp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Sp == nil {
		return result, errors.New("no sp received after update")
	}

	result = *res.Sp

	return result, nil
}

func (c *IdbusApiClient) DeleteVirtSaml2Sp(ida string, sp string) (bool, error) {
	c.logger.Debugf("deleteVirtSaml2Sp. %s [%s]", sp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteVirtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteVirtSaml2Sp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteVirtSaml2Sp(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.DeleteVirtSaml2SpExecute(req)

	if err != nil {
		c.logger.Errorf("deletesp. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deletesp. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deletesp. Deleted %s : %t", sp, *res.Removed)

	return *res.Removed, err
}

// Gets an Sp based on the appliance name and sp name
func (c *IdbusApiClient) GetVirtSaml2Sp(ida string, sp string) (api.VirtualSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("GetVirtSaml2Sp. %s [%s]", sp, ida)
	var result api.VirtualSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetVirtSaml2Sp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetVirtSaml2Sp(ctx)
	req = req.GetVirtSaml2SpReq(api.GetVirtSaml2SpReq{IdOrName: &ida, Name: &sp})
	res, _, err := c.apiClient.DefaultApi.GetVirtSaml2SpExecute(req)
	if err != nil {
		c.logger.Errorf("GetVirtSaml2Sp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetVirtSaml2Sp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Sp == nil {
		c.logger.Debugf("GetVirtSaml2Sp. NOT FOUND %s", sp)
		return result, nil
	}

	if res.Sp != nil {
		result = *res.Sp
		c.logger.Debugf("GetVirtSaml2Sp. %s found for ID/name %s", *result.Name, sp)
	} else {
		c.logger.Debugf("GetVirtSaml2Sp. not found for ID/name %s", sp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetVirtSaml2Sps(ida string) ([]api.VirtualSaml2ServiceProviderDTO, error) {
	c.logger.Debugf("get VirtSaml2Sps: all [%s]", ida)
	var result []api.VirtualSaml2ServiceProviderDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetVirtSaml2Sps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetVirtSaml2Sps(ctx)
	req = req.GetVirtSaml2SpReq(api.GetVirtSaml2SpReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetVirtSaml2SpsExecute(req)
	if err != nil {
		c.logger.Errorf("getVirtSaml2Sps. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Sps == nil {
		return result, nil
	}

	result = res.Sps

	return result, nil

}

func VirtSaml2Sp(VirtSaml2Sp *api.VirtualSaml2ServiceProviderDTO) {
	VirtSaml2Sp.AdditionalProperties = make(map[string]interface{})
	VirtSaml2Sp.AdditionalProperties["@c"] = ".VirtualSaml2ServiceProviderDTO"
}
