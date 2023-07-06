package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIdPSaml2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdPSaml2Create,
		ReadContext:   resourceIdPSaml2Read,
		UpdateContext: resourceIdPSaml2Update,
		DeleteContext: resourceIdPSaml2Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "internal element id",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "resource name",
			},
			"description": {
				Type:        schema.TypeString,
				Description: "SAML 2 service provider description",
				Optional:    true,
			},
			"metadata": {
				Type:        schema.TypeString,
				Description: "SAML 2 service provider XML metadata descriptor",
				Required:    true,
			},
		},
	}
}

func resourceIdPSaml2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdPSaml2Create", "ida", d.Get("ida").(string))

	idpSaml2, err := buildIdPSaml2DTO(d)
	if err != nil {
		return diag.Errorf("failed to build idpsaml2 (new) : %v", err)
	}
	l.Trace("resourceIdPSaml2Create", "ida", d.Get("ida").(string), "name", *idpSaml2.Name)

	a, err := getJossoClient(m).CreateIdPSaml2(d.Get("ida").(string), idpSaml2)
	if err != nil {
		l.Debug("resourceIdPSaml2Create %v", err)
		return diag.Errorf("failed to create idpsaml2: %v", err)
	}

	if err = buildIdPSaml2Resource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdPSaml2Create %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdPSaml2Create OK", "ida", d.Get("ida").(string), "name", *idpSaml2.Name)

	return nil

}

func resourceIdPSaml2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIdPSaml2Read", "ida", idaName, "name", d.Id())
	idpSaml2, err := getJossoClient(m).GetIdPSaml2(idaName, d.Id())
	if err != nil {
		l.Debug("resourceIdPSaml2Read %v", err)
		return diag.Errorf("resourceIdPSaml2Read: %v", err)
	}
	if idpSaml2.Name == nil || *idpSaml2.Name == "" {
		l.Debug("resourceIdPSaml2Read NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdPSaml2Resource(idaName, d, idpSaml2); err != nil {
		l.Debug("resourceIdPSaml2Read %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdPSaml2Read OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceIdPSaml2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdPSaml2Update", "ida", d.Get("ida").(string), "name", d.Id())

	idpsaml2, err := buildIdPSaml2DTO(d)
	if err != nil {
		l.Debug("resourceIdPSaml2Update %v", err)
		return diag.Errorf("failed to build idpsaml2 (upd): %v", err)
	}

	a, err := getJossoClient(m).UpdateIdPSaml2(d.Get("ida").(string), idpsaml2)
	if err != nil {
		l.Debug("resourceIdPSaml2Update %v", err)
		return diag.Errorf("failed to update idpsaml2: %v", err)
	}

	if err = buildIdPSaml2Resource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdPSaml2Update %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdPSaml2Update OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdPSaml2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdPSaml2Delete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdPSaml2(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdPSaml2Delete %v", err)
		return diag.Errorf("failed to delete idpsaml2: %v", err)
	}

	l.Debug("resourceIdPSaml2Delete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdPSaml2DTO(d *schema.ResourceData) (api.ExternalSaml2IdentityProviderDTO, error) {
	var err error
	dto := api.NewExternalSaml2IdentityProviderDTO()
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	m := api.NewResourceDTO()
	m.SetValue(*PtrSchemaStr(d, "metadata"))
	m.SetName(fmt.Sprintf("%s-md", dto.GetName()))
	dto.SetMetadata(*m)

	// Federated connections / idps
	// SP side of federated connection is for the SP
	/*
		dto.FederatedConnectionsA, err = convertIdPSaml2_SPFederatedConnectionsMapArrToDTOs(dto, d, d.Get("sp"))
		if err != nil {
			return *dto, err
		}
	*/

	return *dto, err
}

func buildIdPSaml2Resource(idaName string, d *schema.ResourceData, dto api.ExternalSaml2IdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))

	m := dto.GetMetadata()
	_ = d.Set("metadata", m.GetValue())

	// Federated connections / idps
	/*
		sps, err := convertIdPSaml2_SPFederatedConnectionsToMapArr(dto.FederatedConnectionsB)
		if err != nil {
			return err
		}
		_ = d.Set("sp", sps)
	*/

	return nil
}

func convertIdPSaml2_SPFederatedConnectionsToMapArr(fcs []api.FederatedConnectionDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	for _, fc := range fcs {

		idp_map := map[string]interface{}{
			"name": fc.GetName(),
		}
		result = append(result, idp_map)

	}

	return result, nil
}

func convertIdPSaml2_SPFederatedConnectionsMapArrToDTOs(idpSaml2 *api.ExternalSaml2IdentityProviderDTO, d *schema.ResourceData, sp interface{}) ([]api.FederatedConnectionDTO, error) {
	result := make([]api.FederatedConnectionDTO, 0)
	ls, ok := sp.([]interface{})
	if !ok {
		return result, fmt.Errorf("invalid type: %T", sp)
	}

	for _, v := range ls {
		// 1. For each SP(terraform), create a FederatedConnection
		// Store all connections in idpSaml2.FederatedConnectionsA array.
		// The sp name is to be used as the federated connection name
		// 2. Each connection will have a FederatedChannel,
		// Create an SPChannel and use convertion function to get FederatecChannel
		// Store the FederatedChannel in the federatedConnection.federatedChannelA member/element/var
		m, ok := v.(map[string]interface{})
		if !ok {
			return result, fmt.Errorf("invalid element type: %T", v)
		}

		// build new federatedconnectionDTO
		c := api.NewFederatedConnectionDTO()
		c.AdditionalProperties = make(map[string]interface{})

		c.SetName(m["name"].(string))
		c.AdditionalProperties["@c"] = ".FederatedConnectionDTO"

		// from federatedconnectionDTO.Channelb values
		spChannel := api.NewInternalSaml2ServiceProviderChannelDTO()
		// Assing values for preferred option
		//spChannel.SetPreferred(m["is_preferred"].(bool))
		spChannel.SetOverrideProviderSetup(false)
		c.SetSPChannel(spChannel)
		result = append(result, *c)
	}
	return result, nil
}
