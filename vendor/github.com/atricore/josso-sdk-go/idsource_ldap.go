package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

func (c *IdbusApiClient) CreateIdSourceLdap(ida string, idSourceLdap api.LdapIdentitySourceDTO) (api.LdapIdentitySourceDTO, error) {
	var result api.LdapIdentitySourceDTO
	l := c.Logger()

	l.Debugf("createIdSourceLdap : %s [%s]", *idSourceLdap.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateIdSourceLdap") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdSourceLdap(&idSourceLdap)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateIdSourceLdap(ctx)
	req = req.StoreIdSourceLdapReq(api.StoreIdSourceLdapReq{IdOrName: &ida, IdSourceLdap: &idSourceLdap})
	res, _, err := c.apiClient.DefaultApi.CreateIdSourceLdapExecute(req)
	if err != nil {
		c.logger.Errorf("createIdSourceLdap. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("createIdSourceLdap. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdSourceLdap == nil {
		return result, errors.New("no idSourceLdap received after creation")
	}

	result = *res.IdSourceLdap

	return result, nil
}

func (c *IdbusApiClient) UpdateIdSourceLdap(ida string, idSourceLdap api.LdapIdentitySourceDTO) (api.LdapIdentitySourceDTO, error) {
	var result api.LdapIdentitySourceDTO
	l := c.Logger()

	l.Debugf("updateIdSourceLdap. : %s [%s]", *idSourceLdap.Name, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateIdSourceLdap") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initIdSourceLdap(&idSourceLdap)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateIdSourceLdap(ctx)
	req = req.StoreIdSourceLdapReq(api.StoreIdSourceLdapReq{IdOrName: &ida, IdSourceLdap: &idSourceLdap})
	res, _, err := c.apiClient.DefaultApi.UpdateIdSourceLdapExecute(req)
	if err != nil {
		c.logger.Errorf("updateIdSourceLdap. Error %v", err)
		return result, err

	}

	if res.Error != nil {
		msg := buildErrorMsg(*res.Error, res.ValidationErrors)
		c.logger.Errorf("updateIdSourceLdap. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.IdSourceLdap == nil {
		return result, errors.New("no idSourceLdap received after update")
	}

	result = *res.IdSourceLdap

	return result, nil
}

func (c *IdbusApiClient) DeleteIdSourceLdap(ida string, idSourceLdap string) (bool, error) {
	c.logger.Debugf("deleteIdSourceLdap. %s [%s]", idSourceLdap, ida)
	sc, err := c.IdbusServerForOperation("DefaultApiService.DeleteIdSourceLdap") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteIdSourceLdap. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteIdSourceLdap(ctx)
	req = req.DeleteReq(api.DeleteReq{IdOrName: &ida, Name: &idSourceLdap})
	res, _, err := c.apiClient.DefaultApi.DeleteIdSourceLdapExecute(req)

	if err != nil {
		c.logger.Errorf("deleteIdSourceLdap. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteIdSourceLdap. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteIdSourceLdap. Deleted %s : %t", idSourceLdap, *res.Removed)

	return *res.Removed, err
}

func (c *IdbusApiClient) GetIdSourceLdap(ida string, idSourceLdap string) (api.LdapIdentitySourceDTO, error) {
	c.logger.Debugf("getIdSourceLdap. %s [%s]", idSourceLdap, ida)
	var result api.LdapIdentitySourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdSourceLdap") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSourceLdap(ctx)
	req = req.GetIdSourceLdapReq(api.GetIdSourceLdapReq{IdOrName: &ida, Name: &idSourceLdap})
	res, _, err := c.apiClient.DefaultApi.GetIdSourceLdapExecute(req)
	if err != nil {
		c.logger.Errorf("getIdSourceLdap. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("getIdSourceLdap. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.IdSourceLdap == nil {
		c.logger.Debugf("getIdSourceLdap. NOT FOUND %s", idSourceLdap)
		return result, nil
	}

	if res.IdSourceLdap != nil {
		result = *res.IdSourceLdap
		c.logger.Debugf("getIdSourceLdap. %s found for ID/name %s", *result.Name, idSourceLdap)
	} else {
		c.logger.Debugf("getIdSourceLdap. not found for ID/name %s", idSourceLdap)
	}

	return result, nil

}

func (c *IdbusApiClient) GetIdSourceLdaps(ida string) ([]api.LdapIdentitySourceDTO, error) {
	c.logger.Debugf("get idSourceLdaps: all [%s]", ida)
	var result []api.LdapIdentitySourceDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetIdSourceLdaps") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetIdSourceLdaps(ctx)
	req = req.GetIdSourceLdapReq(api.GetIdSourceLdapReq{IdOrName: &ida})
	res, _, err := c.apiClient.DefaultApi.GetIdSourceLdapsExecute(req)
	if err != nil {
		c.logger.Errorf("getIdSourceLdaps. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.IdSourceLdaps == nil {
		return result, nil
	}

	result = res.IdSourceLdaps

	return result, nil

}

func initIdSourceLdap(idSourceLdap *api.LdapIdentitySourceDTO) {
	idSourceLdap.AdditionalProperties = make(map[string]interface{})
	idSourceLdap.AdditionalProperties["@c"] = ".LdapIdentitySourceDTO"
}
