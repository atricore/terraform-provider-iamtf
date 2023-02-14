package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new IDP in the provided identity appliance. It receives the appliance name or id and the oidcRp dto to use as template
func (c *IdbusApiClient) CreateOidcRp(ida string, oidcRp api.ExternalOpenIDConnectRelayingPartyDTO) (api.ExternalOpenIDConnectRelayingPartyDTO, error) {
	var result api.ExternalOpenIDConnectRelayingPartyDTO
	l := c.Logger()

	l.Debugf("createOidcRp : %s [%s]", *oidcRp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateOidcRp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initOidcRp(&oidcRp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateOidcRp(ctx)
	req = req.StoreOidcRpReq(api.StoreOidcRpReq{IdOrName: &ida, OidcRp: &oidcRp})
	res, _, err := c.apiClient.DefaultApi.CreateOidcRpExecute(req)
	if err != nil {
		c.logger.Errorf("createOidcRp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createOidcRp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.OidcRp == nil {
		return result, errors.New("no oidcRp received after creation")
	}

	result = *res.OidcRp

	return result, nil
}

func (c *IdbusApiClient) UpdateOidcRp(ida string, oidcRp api.ExternalOpenIDConnectRelayingPartyDTO) (api.ExternalOpenIDConnectRelayingPartyDTO, error) {
	var result api.ExternalOpenIDConnectRelayingPartyDTO
	l := c.Logger()

	l.Debugf("updateOidcRp. : %s [%s]", *oidcRp.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateOidcRp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initOidcRp(&oidcRp)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateOidcRp(ctx)
	req = req.StoreOidcRpReq(api.StoreOidcRpReq{IdOrName: &ida, OidcRp: &oidcRp})
	res, _, err := c.apiClient.DefaultApi.UpdateOidcRpExecute(req)
	if err != nil {
		c.logger.Errorf("updateOidcRp. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateOidcRp. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.OidcRp == nil {
		return result, errors.New("no oidcRp received after update")
	}

	result = *res.OidcRp

	return result, nil
}

func (c *IdbusApiClient) DeleteOidcRp(ida string, oidcRp string) (bool, error) {
	c.logger.Debugf("deleteOidcRp. %s [%s]", oidcRp, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteOidcRp") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteOidcRp. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteOidcRp(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &oidcRp})
	res, _, err := c.apiClient.DefaultApi.DeleteOidcRpExecute(req)

	if err != nil {
		c.logger.Errorf("deleteOidcRp. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteOidcRp. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteOidcRp. Deleted %s : %t", oidcRp, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetOidcRp(ida string, oidcRp string) (api.ExternalOpenIDConnectRelayingPartyDTO, error) {
	c.logger.Debugf("getOidcRp. %s [%s]", oidcRp, ida)
	var result api.ExternalOpenIDConnectRelayingPartyDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetOidcRp") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetOidcRp(ctx)
	req = req.GetOidcRpReq(api.GetOidcRpReq{IdOrName: &ida, Name: &oidcRp})
	res, _, err := c.apiClient.DefaultApi.GetOidcRpExecute(req)
	if err != nil {
		c.logger.Errorf("getOidcRp. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getOidcRp. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.OidcRp == nil {
		c.logger.Debugf("getOidcRp. NOT FOUND %s", oidcRp)
		return result, nil
	}

	if res.OidcRp != nil {
		result = *res.OidcRp
		c.logger.Debugf("getOidcRp. %s found for ID/name %s", *result.Name, oidcRp)
	} else {
		c.logger.Debugf("getOidcRp. not found for ID/name %s", oidcRp)
	}

	return result, nil

}

func (c *IdbusApiClient) GetOidcRps(ida string) ([]api.ExternalOpenIDConnectRelayingPartyDTO, error) {
	c.logger.Debugf("get oidcRps: all [%s]", ida)
	var result []api.ExternalOpenIDConnectRelayingPartyDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetOidcRps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetOidcRps(ctx)
	req = req.GetOidcRpReq(api.GetOidcRpReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetOidcRpsExecute(req)
	if err != nil {
		c.logger.Errorf("getOidcRps. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.OidcRps == nil {
		return result, nil
	}

	result = res.OidcRps

	return result, nil

}

func initOidcRp(oidcRp *api.ExternalOpenIDConnectRelayingPartyDTO) {
	oidcRp.AdditionalProperties = make(map[string]interface{})
	oidcRp.AdditionalProperties["@c"] = ".ExternalOpenIDConnectRelayingPartyDTO"
}
