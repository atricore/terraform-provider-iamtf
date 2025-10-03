package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceExtSaml2Sp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceExtSaml2SpCreate,
		ReadContext:   resourceExtSaml2SpRead,
		UpdateContext: resourceExtSaml2SpUpdate,
		DeleteContext: resourceExtSaml2SpDelete,
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
			"idp": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "List of IDP resource names trusted by this SP",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "name of the trusted IdP",
						},
						"is_preferred": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "identifies this IdP as the preferred one (only one IdP must be set to preferred)",
						},
					},
				},
			},
		},
	}
}

func resourceExtSaml2SpCreate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceExtSaml2SpCreate", "ida", d.Get("ida").(string))

	extsaml2sp, err := buildExtSaml2SpDTO(d)
	if err != nil {
		return diag.Errorf("failed to build extsaml2sp: %v", err)
	}
	l.Trace("resourceExtSaml2SpCreate", "ida", d.Get("ida").(string), "name", *extsaml2sp.Name)

	a, err := getJossoClient(m).CreateExtSaml2Sp(d.Get("ida").(string), extsaml2sp)
	if err != nil {
		l.Debug("resourceExtSaml2SpCreate %v", err)
		return diag.Errorf("failed to create extsaml2sp: %v", err)
	}

	if err = buildExtSaml2SpResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceExtSaml2SpCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceExtSaml2SpCreate OK", "ida", d.Get("ida").(string), "name", *extsaml2sp.Name)

	return nil
}

func resourceExtSaml2SpRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceExtSaml2SpRead", "ida", idaName, "name", d.Id())
	extsaml2sp, err := getJossoClient(m).GetExtSaml2Sp(idaName, d.Id())
	if err != nil {
		l.Debug("resourceExtSaml2SpRead %v", err)
		return diag.Errorf("resourceExtSaml2SpRead: %v", err)
	}
	if extsaml2sp.Name == nil || *extsaml2sp.Name == "" {
		l.Debug("resourceExtSaml2SpRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildExtSaml2SpResource(idaName, d, extsaml2sp); err != nil {
		l.Debug("resourceExtSaml2SpRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceExtSaml2SpRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceExtSaml2SpUpdate(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceExtSaml2SpUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	extsaml2sp, err := buildExtSaml2SpDTO(d)
	if err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.Errorf("failed to build extsaml2sp: %v", err)
	}

	a, err := getJossoClient(m).UpdateExtSaml2Sp(d.Get("ida").(string), extsaml2sp)
	if err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.Errorf("failed to update extsaml2sp: %v", err)
	}

	if err = buildExtSaml2SpResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceExtSaml2SpUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceExtSaml2SpDelete(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceExtSaml2SpDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteExtSaml2Sp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceExtSaml2SpDelete %v", err)
		return diag.Errorf("failed to delete extsaml2sp: %v", err)
	}

	l.Debug("resourceExtSaml2SpDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildExtSaml2SpDTO(d *schema.ResourceData) (api.ExternalSaml2ServiceProviderDTO, error) {
	var err error
	dto := api.NewExternalSaml2ServiceProviderDTO()
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	m := api.NewResourceDTO()
	m.SetValue(*PtrSchemaStr(d, "metadata"))
	m.SetName(fmt.Sprintf("%s-md", dto.GetName()))
	dto.SetMetadata(*m)

	// Federated connections / idps
	// SP side of federated connection is for the SP
	dto.FederatedConnectionsB, err = convertExtSsamlSp_IdPFederatedConnectionsMapArrToDTOs(
		dto,
		d,
		d.Get("idp"),
	)
	if err != nil {
		return *dto, err
	}

	return *dto, err
}

func buildExtSaml2SpResource(
	idaName string,
	d *schema.ResourceData,
	dto api.ExternalSaml2ServiceProviderDTO,
) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))

	m := dto.GetMetadata()
	_ = d.Set("metadata", m.GetValue())

	// Federated connections / idps
	idps, err := convertExtSaml2SP_IdPFederatedConnectionsToMapArr(dto.FederatedConnectionsB)
	if err != nil {
		return err
	}
	_ = d.Set("idp", idps)

	return nil
}

func convertExtSaml2SP_IdPFederatedConnectionsToMapArr(
	fcs []api.FederatedConnectionDTO,
) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	for _, fc := range fcs {

		idpChannel, err := fc.GetIDPChannel()
		if err != nil {
			return result, err
		}
		idp_map := map[string]interface{}{
			"name":         fc.GetName(),
			"is_preferred": idpChannel.GetPreferred(),
		}
		result = append(result, idp_map)

	}

	return result, nil
}

func convertExtSsamlSp_IdPFederatedConnectionsMapArrToDTOs(
	sp *api.ExternalSaml2ServiceProviderDTO,
	d *schema.ResourceData,
	idp interface{},
) ([]api.FederatedConnectionDTO, error) {
	result := make([]api.FederatedConnectionDTO, 0)
	ls, ok := idp.([]interface{})
	if !ok {
		return result, fmt.Errorf("invalid type: %T", idp)
	}

	for _, v := range ls {
		// 1. For each IdP(terraform), create a FederatedConnection
		// Store all connections in sp.FederatedConnectionsB array.
		// The idp name is to be used as the federated connection name
		// 2. Each connection will have a FederatedChannel,
		// Create an IDPChannel and use convertion function to get FederatecChannel
		// Store the FederatedChannel in the federatedConnection.federatedChannelB member/element/var
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
		idpChannel := api.NewIdentityProviderChannelDTO()
		// Assing values for preferred option
		idpChannel.SetPreferred(m["is_preferred"].(bool))

		idpChannel.SetOverrideProviderSetup(false)
		c.SetIDPChannel(idpChannel)
		result = append(result, *c)
	}
	return result, nil
}
