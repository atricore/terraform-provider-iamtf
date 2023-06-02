package cli

import (
	"context"
	"errors"

	api "github.com/atricore/josso-api-go"
)

// Creates a new CustomBrandingDefinitionDTO in the provided identity appliance. It receives the appliance name or id and the Db dto to use as template
func (c *IdbusApiClient) CreateBrandingDefinitionDTO(
	name string,
	description string,
	bundleGroup string,
	bundleArtifact string,
	bundleVersion string,
	resource string,
	customSsoAppClazz string,
	customSsoIdPAppClazz string,
	customOpenIdAppClazz string,
) (api.CustomBrandingDefinitionDTO, error) {

	var result api.CustomBrandingDefinitionDTO
	l := c.Logger()

	l.Debugf("CreateCustomBrandingDefinition : %s", name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.CreateCustomBrandingDefinition") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initCustomBrandingDefinitionDTO(&result)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.CreateBranding(ctx)
	req = req.StoreBrandingReq(api.StoreBrandingReq{
		Name:                 &name,
		Description:          &description,
		BundleGroup:          &bundleGroup,
		BundleArtifact:       &bundleArtifact,
		BundleVersion:        &bundleVersion,
		Resource:             &resource,
		CustomOpenIdAppClazz: &customOpenIdAppClazz,
		CustomSsoAppClazz:    &customSsoAppClazz,
		CustomSsoIdPAppClazz: &customSsoIdPAppClazz,
	})
	res, _, err := c.apiClient.DefaultApi.CreateBrandingExecute(req)
	if err != nil {
		c.logger.Errorf("CreateCustomBrandingDefinition. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		msg := *res.Error
		c.logger.Errorf("CreateCustomBrandingDefinition. Error %s", msg)
		return result, errors.New(msg)
	}

	if res.Branding == nil {
		return result, errors.New("no CustomBrandingDefinitionDTO received after creation")
	}

	result = *res.Branding

	return result, nil
}

func (c *IdbusApiClient) UpdateCustomBrandingDefinitionDTO(
	name string,
	description string,
	bundleGroup string,
	bundleArtifact string,
	bundleVersion string,
	resource string,
	customSsoAppClazz string,
	customSsoIdPAppClazz string,
	customOpenIdAppClazz string,
) (api.CustomBrandingDefinitionDTO, error) {
	var result api.CustomBrandingDefinitionDTO
	l := c.Logger()

	l.Debugf("UpdateCustomBrandingDefinitionDTO. : %s ", name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.UpdateCustomBrandingDefinitionDTO") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	initCustomBrandingDefinitionDTO(&result)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.UpdateBranding(ctx)
	req = req.StoreBrandingReq(api.StoreBrandingReq{
		Name:                 &name,
		Description:          &description,
		BundleGroup:          &bundleGroup,
		BundleArtifact:       &bundleArtifact,
		BundleVersion:        &bundleVersion,
		Resource:             &resource,
		CustomOpenIdAppClazz: &customOpenIdAppClazz,
		CustomSsoAppClazz:    &customSsoAppClazz,
		CustomSsoIdPAppClazz: &customSsoIdPAppClazz,
	})
	res, _, err := c.apiClient.DefaultApi.UpdateBrandingExecute(req)
	if err != nil {
		c.logger.Errorf("UpdateCustomBrandingDefinitionDTO. Error %v", err)
		return result, err

	}

	if res.Branding == nil {
		return result, errors.New("no CustomBrandingDefinitionDTO received after update")
	}

	result = *res.Branding

	return result, nil
}

func (c *IdbusApiClient) DeleteCustomBrandingDefinitionDTO(name string) (bool, error) {
	c.logger.Debugf("deleteCustomBrandingDefinitionDTO. %s", name)
	sc, err := c.IdbusServerForOperation("DefaultApiService.deleteCustomBrandingDefinitionDTO") // Also hard-coded in generated client
	if err != nil {
		c.logger.Errorf("deleteCustomBrandingDefinitionDTO. Error %v", err)
		return false, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.DeleteBranding(ctx)
	req = req.DeleteBrandingReq(api.DeleteBrandingReq{Name: &name})
	res, _, err := c.apiClient.DefaultApi.DeleteBrandingExecute(req)

	if err != nil {
		c.logger.Errorf("deleteCustomBrandingDefinitionDTO. Error %v", err)
		return false, err
	}

	if res.Error != nil {
		c.logger.Errorf("deleteCustomBrandingDefinitionDTO. Error %v", *res.Error)
		return false, errors.New(*res.Error)
	}

	c.logger.Debugf("deleteCustomBrandingDefinitionDTO. Deleted %s : %t", name, *res.Removed)

	return *res.Removed, err
}

// Gets a custom branding definition DTO by name
func (c *IdbusApiClient) GetBrandingDefinitionDTO(name string) (api.CustomBrandingDefinitionDTO, error) {
	c.logger.Debugf("GetBrandingDefinitionDTO. %s", name)
	var result api.CustomBrandingDefinitionDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetBrandingDefinitionDTO") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetBranding(ctx)
	req = req.GetBrandingReq(api.GetBrandingReq{NameOrId: &name})
	res, _, err := c.apiClient.DefaultApi.GetBrandingExecute(req)
	if err != nil {
		c.logger.Errorf("GetBrandingDefinitionDTO. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		c.logger.Errorf("GetBrandingDefinitionDTO. Error %v", err)
		return result, errors.New(*res.Error)
	}

	if res.Name == nil {
		c.logger.Debugf("GetBrandingDefinitionDTO. NOT FOUND %s", name)
		return result, nil
	}

	// TODO : Populate result

	result.Name = res.Name
	result.Description = res.Description
	result.Type = res.Type

	return result, nil

}

func (c *IdbusApiClient) GetBrandingDefinitionDTOs() ([]api.CustomBrandingDefinitionDTO, error) {
	c.logger.Debugf("getBrandingDefinitionDTOs: all ")
	var result []api.CustomBrandingDefinitionDTO

	sc, err := c.IdbusServerForOperation("DefaultApiService.GetBrandingDefinitionDTOs") // Also hard-coded in generated client
	if err != nil {
		return result, err
	}

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, sc.Authn.AccessToken)
	req := c.apiClient.DefaultApi.GetAllBrandings(ctx)
	req = req.GetAllBrandingsReq(api.GetAllBrandingsReq{})
	res, _, err := c.apiClient.DefaultApi.GetAllBrandingsExecute(req)
	if err != nil {
		c.logger.Errorf("getBrandingDefinitionDTOs. Error %v", err)
		return result, err
	}

	if res.Error != nil {
		return result, errors.New(*res.Error)
	}

	if res.Brandings == nil {
		return result, nil
	}

	result = res.Brandings

	return result, nil

}

func initCustomBrandingDefinitionDTO(CustomBrandingDefinitionDTO *api.CustomBrandingDefinitionDTO) {
	CustomBrandingDefinitionDTO.AdditionalProperties = make(map[string]interface{})
	CustomBrandingDefinitionDTO.AdditionalProperties["@c"] = ".CustomBrandingDefinitionDTO"
}
