package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func ResourceJosso1Re() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppAgentCreate,
		ReadContext:   resourceAppAgentRead,
		UpdateContext: resourceAppAgentUpdate,
		DeleteContext: resourceAppAgentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identity appliance name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Application description",
			},
			"app_slo_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Application single logout location",
			},
			"sp_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service provider ID. The name of the SP that will be associated with the application",
			},
			"dashboard_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Application URL used to display error information (combined with error_binding)",
			},
			"app_location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "application location.  Base application URL, i.e. https://myapp.com",
			},
			"ignored_web_resources": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "list of URL patterns not subject to SSO control (space sperated)",
			},
			"default_resource": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "application default resource (after SSO/SLO) i.e. https://myapp.com/home",
			},
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "internal element ID",
			},
			"exec_env": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of the execution environment resource",
			},
			"error_binding": {
				Type:             schema.TypeString,
				Description:      "how errors are displayed to users (combined with dashboard_url)",
				ValidateDiagFunc: stringInSlice([]string{"JSON", "ARTIFACT", "GET"}),
				Optional:         true,
				Default:          "JSON",
			},
			"keystore": keystoreSchema(),
			"saml2":    spSamlSchema(),
			"idp":      idpConnectionSchema(),
		},
	}
}

func resourceAppAgentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("createAppAgent ", "ida", d.Get("ida").(string))

	// Create app elements: josso1resource and intsaml2sp
	josso1re, intsaml2sp, err := buildAppAgentDTO(d)
	if err != nil {
		return diag.Errorf("failed to build IntSaml2sp: %v", err)
	}

	// Store intsaml2sp in JOSSO server
	l.Trace("createIntSaml2sp ", "ida", d.Get("ida").(string), "name", *intsaml2sp.Name)
	a, err := getJossoClient(m).CreateIntSaml2Sp(d.Get("ida").(string), intsaml2sp)
	if err != nil {
		l.Debug("createIntSaml2sp %v", err)
		return diag.Errorf("failed to create IntSaml2sp: %v", err)
	}

	// Create a service connection between app and sp, no need to handle error
	josso1re.NewServiceConnection(intsaml2sp.GetName())

	// Store josso1resource in JOSSO server
	l.Trace("createJossoresource ", "ida", d.Get("ida").(string), "name", *intsaml2sp.Name)
	b, err := getJossoClient(m).CreateJossoresource(d.Get("ida").(string), josso1re)
	if err != nil {
		l.Debug("createJossoresource %v", err)
		return diag.Errorf("failed to create josso1re: %v", err)
	}

	// populate TF with received data
	if err = buildAppAgentResource(d.Get("ida").(string), d, b, a); err != nil {
		l.Debug("createAppAgent %v", err)
		return diag.FromErr(err)
	}

	return nil
}

func resourceAppAgentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIntSaml2spRead", "ida", idaName, "spname", *PtrSchemaStr(d, "sp_id"))
	sp, err := getJossoClient(m).GetIntSaml2Sp(idaName, *PtrSchemaStr(d, "sp_id"))
	if err != nil {
		l.Debug("resourceIntSaml2spRead %v", err)
		return diag.Errorf("resourceIntSaml2spRead: %v", err)
	}
	if sp.Name == nil || *sp.Name == "" {
		l.Debug("resourceIntSaml2spRead NOT FOUND")
		d.SetId("")
		return nil
	}
	l.Debug("resourceIntSaml2spRead OK", "ida", idaName, "name", d.Id())

	l.Trace("resourceJosso1ReRead", "ida", idaName, "name", d.Id())
	josso1re, err := getJossoClient(m).GetJosso1Resource(idaName, d.Id())
	if err != nil {
		l.Debug("resourceJosso1ReRead %v", err)
		return diag.Errorf("resourceJosso1ReRead: %v", err)
	}
	if josso1re.Name == nil || *josso1re.Name == "" {
		l.Debug("resourceJosso1ReRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildAppAgentResource(idaName, d, josso1re, sp); err != nil {
		l.Debug("resourceAppAgentRead %v", err)
		return diag.FromErr(err)
	}

	return nil
}

func resourceAppAgentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceAppAgentUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())
	josso1re, sp, err := buildAppAgentDTO(d)
	if err != nil {
		l.Debug("resourceAppAgentUpdate %v", err)
		return diag.Errorf("failed to build AppAgent: %v", err)
	}

	b, err := getJossoClient(m).UpdateIntSaml2Sp(d.Get("ida").(string), sp)
	if err != nil {
		l.Debug("resourceAppAgentUpdate/intsaml2sp %v", err)
		return diag.Errorf("failed to update IntSaml2sp: %v", err)
	}
	a, err := getJossoClient(m).UpdateJosso1Resource(d.Get("ida").(string), josso1re)
	if err != nil {
		l.Debug("resourceAppAgentUpdate/josso1re %v", err)
		return diag.Errorf("failed to update josso1re: %v", err)
	}

	if err = buildAppAgentResource(d.Get("ida").(string), d, a, b); err != nil {
		l.Debug("resourceAppAgentUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceAppAgentUpdate/intsaml2sp OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceAppAgentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceJosso1ReDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteJosso1Resource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceJosso1ReDelete %v", err)
		return diag.Errorf("failed to delete josso1re: %v", err)
	}
	l.Debug("resourceJosso1ReDelete OK", "ida", d.Get("ida").(string), "name", d.Id())
	d.Get("idps")
	l.Trace("resourceIntSaml2spExecenvDelete", "ida", d.Get("ida").(string), "name", d.Get("sp_id").(string))
	_, err = getJossoClient(m).DeleteIntSaml2Sp(d.Get("ida").(string), d.Get("sp_id").(string))
	if err != nil {
		l.Debug("resourceIntSaml2spExecenvDelete %v", err)
		return diag.Errorf("failed to delete saml2sp: %v", err)
	}

	l.Debug("resourceIntSaml2spExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildAppAgentDTO(d *schema.ResourceData) (api.JOSSO1ResourceDTO, api.InternalSaml2ServiceProviderDTO, error) {

	var err, errWrap error

	// JOSSO 1 Res
	josso1re := api.NewJOSSO1ResourceDTO()

	// SP
	sp := api.NewInternalSaml2ServiceProviderDTO()

	_, e := d.GetOk("app_slo_location")
	if e {
		josso1re.SloLocation, err = PtrSchemaLocation(d, "app_slo_location")
		if err != nil {
			return *josso1re, *sp, fmt.Errorf("invalid app_slo_location %s", err)
		}
	}
	josso1re.SloLocationEnabled = &e

	josso1re.PartnerAppLocation, err = PtrSchemaLocation(d, "app_location")
	if err != nil {
		return *josso1re, *sp, fmt.Errorf("invalid app_location %s", err)
	}
	ru := convertInterfaceToStringSetNullable(d.Get("ignored_web_resources"))
	josso1re.IgnoredWebResources = ru
	josso1re.DefaultResource = PtrSchemaStr(d, "default_resource")
	josso1re.Description = PtrSchemaStr(d, "description")
	josso1re.ElementId = PtrSchemaStr(d, "element_id")
	josso1re.Name = PtrSchemaStr(d, "name")

	josso1re.NewActivation(*PtrSchemaStr(d, "exec_env"))

	// --------------------------------------------------------
	// SP
	// --------------------------------------------------------
	// On create sp_id is empty, on update it has a valid value
	spName := PtrSchemaStr(d, "sp_id")
	if *spName == "" {
		// This is a create SP
		spName = PtrSchemaStr(d, "name")
		*spName = fmt.Sprintf("%s-sp", *spName)
	}
	sp.Name = spName

	sp.DashboardUrl = PtrSchemaStr(d, "dashboard_url")
	sp.Description = PtrSchemaStr(d, "description")
	sp.DisplayName = PtrSchemaStr(d, "version")
	sp.ErrorBinding = PtrSchemaStr(d, "error_binding")

	// SP Configuration
	ks, err := convertKeystoreMapArrToDTO(sp.GetName(), d.Get("keystore"))
	if err != nil {
		errWrap = errors.Wrap(err, "keystore")
	}

	cfg := api.NewSamlR2SPConfigDTOInit()
	cfg.SetSigner(*ks)
	cfg.SetEncrypter(*ks)
	cfg.SetUseSampleStore(false)
	cfg.SetUseSystemStore(false)

	sp.SetSamlR2SPConfig(cfg)

	// Some defaults

	// SAML2 settings
	err = convertSPSaml2MapArrToDTO(d.Get("saml2"), sp)
	if err != nil {
		errWrap = errors.Wrap(err, "saml2")
	}

	// SP side of federated connection is for the SP
	sp.FederatedConnectionsB, err = convertIdPFederatedConnectionsMapArrToDTOs(sp, d, d.Get("idp"))
	if err != nil {
		return *josso1re, *sp, err
	}

	// Copy preferred IDP channel values to SP
	_, idpChannel, err := getPreferredIdPChannel(sp)
	if err != nil {
		return *josso1re, *sp, err
	}
	if idpChannel == nil {
		return *josso1re, *sp, fmt.Errorf("iamtf_app_agent resource MUST have a preferred idp: %s", *josso1re.Name)
	}

	return *josso1re, *sp, errWrap
}

func buildAppAgentResource(idaName string, d *schema.ResourceData, josso1 api.JOSSO1ResourceDTO, sp api.InternalSaml2ServiceProviderDTO) error {
	d.SetId(cli.StrDeref(josso1.Name))
	_ = d.Set("sp_id", cli.StrDeref(sp.Name))

	_ = d.Set("ida", idaName)
	// JOSSO 1 Res
	_ = d.Set("app_slo_location", cli.LocationToStr(josso1.SloLocation))
	_ = d.Set("app_location", cli.LocationToStr(josso1.PartnerAppLocation))
	_ = setNonPrimitives(d, map[string]interface{}{
		"ignored_web_resources": convertStringSetToInterface(josso1.IgnoredWebResources)})
	_ = d.Set("default_resource", cli.StrDeref(josso1.DefaultResource))
	_ = d.Set("description", cli.StrDeref(josso1.Description))
	_ = d.Set("element_id", cli.StrDeref(josso1.ElementId))
	_ = d.Set("name", cli.StrDeref(josso1.Name))

	_ = d.Set("exec_env", cli.StrDeref(josso1.Activation.Name))

	// SP
	_ = d.Set("dashboard_url", cli.StrDeref(sp.DashboardUrl))
	_ = d.Set("error_binding", cli.StrDeref(sp.ErrorBinding))

	saml2_m, err := convertSPSaml2DTOToMapArr(&sp)
	if err != nil {
		return err
	}
	_ = d.Set("saml2", saml2_m)

	idps, err := convertIdPFederatedConnectionsToMapArr(sp.FederatedConnectionsB)
	if err != nil {
		return err
	}
	_ = d.Set("idp", idps)

	return nil
}

func getPreferredIdPChannel(sp *api.InternalSaml2ServiceProviderDTO) (*api.FederatedConnectionDTO, *api.IdentityProviderChannelDTO, error) {

	//var idpChannel *api.IdentityProviderChannelDTO
	for _, fc := range sp.GetFederatedConnectionsB() {
		idpChannel, err := fc.GetIDPChannel()
		if err != nil {
			return nil, nil, err
		}

		if idpChannel.GetPreferred() {
			return &fc, idpChannel, nil
		}
	}
	return nil, nil, nil
}

func convertSPSaml2MapArrToDTO(saml2_arr interface{}, sp *api.InternalSaml2ServiceProviderDTO) error {
	m, err := asTFMapSingle(saml2_arr) //
	if err != nil {
		return err
	}

	if m == nil {
		return nil
	}

	// build new accountLinkagePolicyDTO
	al := api.NewAccountLinkagePolicyDTO()
	al.AdditionalProperties = make(map[string]interface{})
	al.AdditionalProperties["@c"] = ".AccountLinkagePolicyDTO"
	// TODO : support for custom mappings : al.SetName(api.AsStringDef(m["account_linkage_name"], "my-account-linkage", true))
	al.SetLinkEmitterType(api.AsStringDef(m["account_linkage"], "ONE_TO_ONE", true))
	sp.SetAccountLinkagePolicy(*al)

	// build new identityMappingPolicyDTO
	im := api.NewIdentityMappingPolicyDTO()
	im.AdditionalProperties = make(map[string]interface{})
	im.AdditionalProperties["@c"] = ".IdentityMappingPolicyDTO"
	// TODO : support for custom mappings : im.SetName(api.AsStringDef(m["identity_mapping_name"], "my-identity-mapping", true))
	im.SetMappingType(api.AsStringDef(m["identity_mapping"], "REMOTE", true))
	im.SetUseLocalId(api.AsBool(m["identity_mapping_localid"], false))
	sp.SetIdentityMappingPolicy(*im)

	sp.SetMessageTtl(api.AsInt32(m["message_ttl"], 300))
	sp.SetMessageTtlTolerance(api.AsInt32(m["message_ttl_tolerance"], 300))

	sp.SetSignRequests(api.AsBool(m["sign_requests"], false))
	sp.SetSignAuthenticationRequests(api.AsBool(m["sign_authentication_requests"], true))
	sp.SetSignatureHash(api.AsString(m["signature_hash"], "SHA-256"))
	sp.SetWantAssertionSigned(api.AsBool(m["want_assertion_signed"], false))

	b, err := convertMapArrToActiveBinding(m["bindings"])
	if err != nil {
		return err
	}

	sp.SetActiveBindings(b)

	return nil
}

func convertSPSaml2DTOToMapArr(sp *api.InternalSaml2ServiceProviderDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	al := sp.GetAccountLinkagePolicy()
	im := sp.GetIdentityMappingPolicy()

	bindings, err := convertActiveBindingToMapArr(sp.GetActiveBindings())
	if err != nil {
		return nil, err
	}

	saml2_map := map[string]interface{}{
		"account_linkage":              al.GetLinkEmitterType(),
		"message_ttl":                  int(sp.GetMessageTtl()),
		"message_ttl_tolerance":        int(sp.GetMessageTtlTolerance()),
		"identity_mapping":             im.GetMappingType(),
		"identity_mapping_localid":     im.GetUseLocalId(),
		"sign_authentication_requests": sp.GetSignAuthenticationRequests(),
		"sign_requests":                sp.GetSignRequests(),
		"signature_hash":               sp.GetSignatureHash(),
		"want_assertion_signed":        sp.GetWantAssertionSigned(),
		"bindings":                     bindings,
	}
	result = append(result, saml2_map)

	return result, nil
}
