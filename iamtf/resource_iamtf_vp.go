package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	sdk "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// implement IdPRole interface
type VPIdPRole struct {
	vp *api.VirtualSaml2ServiceProviderDTO
}

// implement IdPRole interface for VPIdPRole struct
func (r VPIdPRole) GetName() string {
	return r.vp.GetName()
}
func (r VPIdPRole) GetWantAuthnRequestsSigned() bool {
	return r.vp.GetWantAuthnRequestsSigned()
}
func (r VPIdPRole) GetSignatureHash() string {
	return r.vp.GetIdpSignatureHash()
}
func (r VPIdPRole) GetEncryptAssertionAlgorithm() string {
	return r.vp.GetEncryptAssertionAlgorithm()
}
func (r VPIdPRole) GetMessageTtl() int32 {
	return r.vp.GetMessageTtl()
}
func (r VPIdPRole) GetMessageTtlTolerance() int32 {
	return r.vp.GetMessageTtlTolerance()
}

type VPSPRole struct {
	vp *api.VirtualSaml2ServiceProviderDTO
}

func (r VPSPRole) GetName() string {
	return r.vp.GetName()
}
func (r VPSPRole) GetSignAuthenticationRequests() bool {
	return r.vp.GetSignAuthenticationRequests()
}
func (r VPSPRole) GetIdentityMappingPolicy() api.IdentityMappingPolicyDTO {
	return r.vp.GetIdentityMappingPolicy()
}
func (r VPSPRole) GetAccountLinkagePolicy() api.AccountLinkagePolicyDTO {
	return r.vp.GetAccountLinkagePolicy()
}
func (r VPSPRole) GetWantAssertionSigned() bool {
	return r.vp.GetWantAssertionSigned()
}
func (r VPSPRole) GetSignatureHash() string {
	return r.vp.GetSpSignatureHash()
}
func (r VPSPRole) GetMessageTtl() int32 {
	return r.vp.GetMessageTtl()
}
func (r VPSPRole) GetMessageTtlTolerance() int32 {
	return r.vp.GetMessageTtlTolerance()
}

func ResourceVP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVPCreate,
		ReadContext:   resourceVPRead,
		UpdateContext: resourceVPUpdate,
		DeleteContext: resourceVPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliane name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vp name, unique in the appliance scope",
			},
			"element_id": {
				Type:        schema.TypeString,
				Description: "element id",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "provider description",
				Optional:    true,
			},

			// PKI
			"keystore": keystoreSchema(),

			"dashboard_url": {
				Type:        schema.TypeString,
				Description: "External user dashborad URL",
				Optional:    true,
			},
			"error_binding": {
				Type:             schema.TypeString,
				Description:      "how error information is encoded and shared with a custom user dashboard",
				ValidateDiagFunc: stringInSlice([]string{"JSON", "ARTIFACT", "GET"}),
				Default:          "JSON",
				Optional:         true,
			},
			// session properties
			"session_timeout": {
				Type:        schema.TypeInt,
				Description: "SSO session timeout (minutes, default 30)",
				Optional:    true,
				Default:     30,
			},
			// SAML
			"saml2_idp": idpSamlSchema(),
			"saml2_sp":  spSamlSchema(),

			"sp":  spConnectionSchema(),
			"idp": idpConnectionSchema(),

			// OAUTH2
			"oauth2_idp": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    0,
				Description: "OAuth2 protocol settings.  This is maily used by JOSSO internally, for SSO connetions OpenID Connect is the recommended protocol, which is a superset of OAuth2",

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Description: "enable OAuth2 protocol for this IDP",
							Optional:    true,
							Default:     false,
						},
						"shared_key": {
							Type:        schema.TypeString,
							Description: "IDP secret key to be shared with the OAuth client",
							Required:    true,
						},
						"token_validity": {
							Type:        schema.TypeInt,
							Description: "token validity (sec, default 300)",
							Optional:    true,
							Default:     300,
						},
						"rememberme_token_validity": {
							Type:        schema.TypeInt,
							Description: "remember me token validity (sec, default 43200)",
							Optional:    true,
							Default:     43200,
						},
					},
				},
			},

			// OAuth 2
			"oidc_idp": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    0,
				Description: "OpenID Connect protocol settings.  This is the recommended SSO protocol. You must combine this with **iamtf_app_odic** resources (Applications)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Description: "enable OIDC for this IDP",
							Optional:    true,
							Default:     false,
						},
						"access_token_ttl": {
							Type:        schema.TypeInt,
							Description: "access token time to live (sec)",
							Optional:    true,
							Computed:    true,
						},
						"authz_code_ttl": {
							Type:        schema.TypeInt,
							Description: "authorization code time to live (sec)",
							Optional:    true,
							Computed:    true,
						},
						"id_token_ttl": {
							Type:        schema.TypeInt,
							Description: "id token time to live (sec)",
							Optional:    true,
							Computed:    true,
						},
						"user_claims_in_access_token": {
							Type:        schema.TypeBool,
							Description: "include user claims in access token",
							Optional:    true,
							Default:     false,
						},
					},
				},
			},

			"id_sources": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "list of identity sources used by the IDP.  At least one is required.",
			},

			"attributes": idpAttributeProfileSchema(),
			"subject_id": {
				Type:             schema.TypeString,
				Description:      "subject identifier. valid values: **PRINCIPAL**, **EMAIL**, **ATTRIBUTE**, **CUSTOM**",
				ValidateDiagFunc: stringInSlice([]string{"PRINCIPAL", "EMAIL", "ATTRIBUTE", "CUSTOM"}),
				Default:          "PRINCIPAL",
				Optional:         true,
			},
			"subject_id_attr": {
				Type:        schema.TypeString,
				Description: "subject identifier attribute, only valid for **ATTRIBUTE** and **CUSTOM** subject identifier",
				Optional:    true,
			},
			"subject_authn_policies": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "todo add description for subject authens policies",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Todo",
							//ValidateDiagFunc: stringInSlice([]string{"ODO"}),
						},
					},
				},
			},
		},
	}
}

func resourceVPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceVPCreate", "ida", d.Get("ida").(string))

	vp, err := buildVPDTO(d) //
	if err != nil {
		return diag.Errorf("failed to build vp: %v", err)
	}
	l.Trace("resourceVPCreate", "ida", d.Get("ida").(string), "name", *vp.Name)

	vp, err = getJossoClient(m).CreateVirtSaml2Sp(d.Get("ida").(string), vp)
	if err != nil {
		l.Debug("resourceVPCreate %v", err)
		return diag.Errorf("failed to create vp: %v", err)
	}

	if err = buildVPResource(d.Get("ida").(string), d, vp); err != nil {
		l.Debug("resourceVPCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceVPCreate OK", "ida", d.Get("ida").(string), "name", *vp.Name)

	return nil
}
func resourceVPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceVPRead", "ida", d.Get(idaName), "name", d.Id())
	vp, err := getJossoClient(m).GetVirtSaml2Sp(idaName, d.Id())
	if err != nil {
		l.Debug("resourceVPRead %v", err)
		return diag.Errorf("resourceVPRead: %v", err)
	}
	if vp.Name == nil || *vp.Name == "" {
		l.Debug("resourceVPRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildVPResource(idaName, d, vp); err != nil {
		l.Debug("resourceVPRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceVPRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceVPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceVPUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	vp, err := buildVPDTO(d)
	if err != nil {
		l.Debug("resourceVPUpdate %v", err)
		return diag.Errorf("failed to build vp: %v", err)
	}

	vp, err = getJossoClient(m).UpdateVirtSaml2Sp(d.Get("ida").(string), vp)
	if err != nil {
		l.Debug("resourceVPUpdate %v", err)
		return diag.Errorf("failed to update vp: %v", err)
	}

	if err = buildVPResource(d.Get("ida").(string), d, vp); err != nil {
		l.Debug("resourceVPUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceVPUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceVPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceVPDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceVPDelete %v", err)
		return diag.Errorf("failed to delete vp: %v", err)
	}

	l.Debug("resourceVPDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildVPDTO(d *schema.ResourceData) (api.VirtualSaml2ServiceProviderDTO, error) {
	var err, errWrap error
	vp := api.NewVirtualSaml2ServiceProviderDTO()

	vp.Name = PtrSchemaStr(d, "name")
	vp.ElementId = PtrSchemaStr(d, "elementi_d")
	vp.Description = PtrSchemaStr(d, "description")

	// ui properties
	vp.DashboardUrl = PtrSchemaStr(d, "dashboard_url")
	vp.ErrorBinding = PtrSchemaStr(d, "error_binding")

	// session properties
	vp.SsoSessionTimeout = PtrSchemaInt32(d, "session_timeout")

	vp.SetSubjectNameIDPolicy(buildSubjectNameIdPolicy(d))

	// IDP Configuration
	ks, err := convertKeystoreMapArrToDTO(vp.GetName(), d.Get("keystore"))
	if err != nil {
		errWrap = errors.Wrap(err, "keystore")
	}

	cfg := api.NewSamlR2IDPConfigDTOInit()
	cfg.SetSigner(*ks)
	cfg.SetEncrypter(*ks)
	cfg.SetUseSampleStore(false)
	cfg.SetUseSystemStore(false)

	vp.SetSamlR2IDPConfig(cfg)

	// Other sections
	err = convertVPIdPSaml2MapArrToDTO(d.Get("saml2_idp"), vp) //
	if err != nil {
		errWrap = errors.Wrap(err, "saml2_idp")
	}

	err = convertVPSPSaml2MapArrToDTO(d.Get("saml2_sp"), vp) //
	if err != nil {
		errWrap = errors.Wrap(err, "saml2_sp")
	}

	// Attribute profile
	ap, err := convertAttributeProfileMapArrToDTOs(vp.GetName(), d.Get("attributes"))
	if err != nil {
		errWrap = errors.Wrap(err, "attributes")
	}
	vp.SetAttributeProfile(*ap)

	// IDP side of federated connection is for the SP
	vp.FederatedConnectionsA, err = convertSPFederatedConnectionsMapArrToDTOs(VPIdPRole{vp: vp}, d, d.Get("sp"))
	if err != nil {
		errWrap = errors.Wrap(err, "sp")
	}

	vp.FederatedConnectionsB, err = convertIdPFederatedConnectionsMapArrToDTOs(VPSPRole{vp: vp}, d, d.Get("idp"))
	if err != nil {
		errWrap = errors.Wrap(err, "idp")
	}

	err = convertVPOAuth2MapArrToDTO(d.Get("oauth2_idp"), vp)
	if err != nil {
		errWrap = errors.Wrap(err, "oauth2_idp")
	}

	err = convertVPOidcMapArrToDTO(d.Get("oidc_idp"), vp)
	if err != nil {
		errWrap = errors.Wrap(err, "oidc_idp")
	}

	// Id sources
	id_sources := convertInterfaceToStringSetNullable(d.Get("id_sources"))
	vp.IdentityLookups = convertStringArrToIdLookups(id_sources)

	subjectAuthen, err := convertSubjectAuthnPoliciesMapArrToDTO(d.Get("subject_authn_policies"))
	if err != nil {
		errWrap = errors.Wrap(err, "subject_authn_policies")
	}
	vp.SubjectAuthnPolicies = subjectAuthen

	return *vp, errWrap
}

func buildVPResource(idaName string, d *schema.ResourceData, vp api.VirtualSaml2ServiceProviderDTO) error {
	d.SetId(sdk.StrDeref(vp.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("name", sdk.StrDeref(vp.Name))
	_ = d.Set("element_id", sdk.StrDeref(vp.ElementId))
	_ = d.Set("description", sdk.StrDeref(vp.Description))

	_ = d.Set("dashboard_url", sdk.StrDeref(vp.DashboardUrl))
	_ = d.Set("error_binding", sdk.StrDeref(vp.ErrorBinding))

	_ = d.Set("session_timeout", sdk.Int32Deref(vp.SsoSessionTimeout))

	_ = d.Set("subject_id", sdk.StrDeref(vp.GetSubjectNameIDPolicy().Type))
	_ = d.Set("subject_id_attr", sdk.StrDeref(vp.GetSubjectNameIDPolicy().SubjectAttribute))

	cfg, err := vp.GetSamlR2IDPConfig()
	if err != nil {
		return err
	}

	ks := cfg.GetSigner()
	ks_m, err := convertKeystoreDTOToMapArr(&ks)
	if err != nil {
		return err
	}

	_ = d.Set("keystore", ks_m)

	saml2_idp_m, err := convertVPIdPSaml2DTOToMapArr(&vp)
	if err != nil {
		return err
	}
	_ = d.Set("saml2_idp", saml2_idp_m)

	saml2_sp_m, err := convertVPSPSaml2DTOToMapArr(&vp)
	if err != nil {
		return err
	}
	_ = d.Set("saml2_sp", saml2_sp_m)

	// "sp" list
	sps, err := convertSPFederatedConnectionDTOsToMapArr(vp.FederatedConnectionsA)
	if err != nil {
		return err
	}
	_ = d.Set("sp", sps)

	oauth2_m, err := convertVPOAuth2DTOToMapArr(&vp)
	if err != nil {
		return err
	}
	_ = d.Set("oauth2_idp", oauth2_m)

	oidc_m, err := convertVPOidcDTOToMapArr(&vp)
	if err != nil {
		return err
	}
	_ = d.Set("oidc_idp", oidc_m)

	attributes, err := convertAttributeProfileDTOToMapArr(vp.AttributeProfile)
	if err != nil {
		return err
	}
	_ = d.Set("attributes", attributes)

	ids := convertIdLookupsToStringArr(vp.IdentityLookups)
	aggMap := map[string]interface{}{
		"id_sources": convertStringSetToInterface(ids),
	}

	err = setNonPrimitives(d, aggMap)

	subjetAuthen, err := convertSubjectAuthnPoliciesDTOToMapArr(vp.SubjectAuthnPolicies)
	if err != nil {
		return err
	}
	_ = d.Set("subject_authn_policies", subjetAuthen)

	return err
}

func convertVPOAuth2DTOToMapArr(idp *api.VirtualSaml2ServiceProviderDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	oauth2_map := map[string]interface{}{
		"enabled":                   idp.GetOauth2Enabled(),
		"shared_key":                idp.GetOauth2Key(),
		"token_validity":            int64(idp.GetOauth2TokenValidity()),
		"rememberme_token_validity": idp.GetOauth2RememberMeTokenValidity(),
	}
	result = append(result, oauth2_map)

	return result, nil
}

func convertVPOAuth2MapArrToDTO(oauth2_arr interface{}, idp *api.VirtualSaml2ServiceProviderDTO) error {
	// Check that we have an array of any type (interface{})
	oauth2_map, err := asTFMapSingle(oauth2_arr)
	if err != nil {
		return err
	}

	if oauth2_map == nil {
		return nil
	}

	idp.SetOauth2Enabled(oauth2_map["enabled"].(bool))
	idp.SetOauth2Key(oauth2_map["shared_key"].(string))
	idp.SetOauth2TokenValidity(int64(oauth2_map["token_validity"].(int)))
	idp.SetOauth2RememberMeTokenValidity(int64(oauth2_map["rememberme_token_validity"].(int)))

	return nil
}

func convertVPOidcMapArrToDTO(oidc_arr interface{}, idp *api.VirtualSaml2ServiceProviderDTO) error {
	// Check that we have an array of any type (interface{})
	oidc_map, err := asTFMapSingle(oidc_arr)
	if err != nil {
		return err
	}

	if oidc_map == nil {
		return nil
	}

	idp.SetOpenIdEnabled(api.AsBool(oidc_map["enabled"], false))
	idp.SetOidcAccessTokenTimeToLive(int32(oidc_map["access_token_ttl"].(int)))
	idp.SetOidcAuthzCodeTimeToLive(int32(oidc_map["authz_code_ttl"].(int)))
	idp.SetOidcIdTokenTimeToLive(int32(oidc_map["id_token_ttl"].(int)))
	idp.SetOidcIncludeUserClaimsInAccessToken(bool(oidc_map["user_claims_in_access_token"].(bool)))

	return nil
}

func convertVPOidcDTOToMapArr(idp *api.VirtualSaml2ServiceProviderDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	oidc_map := map[string]interface{}{
		"enabled":                     idp.GetOpenIdEnabled(),
		"access_token_ttl":            int(idp.GetOidcAccessTokenTimeToLive()),
		"authz_code_ttl":              int(idp.GetOidcAuthzCodeTimeToLive()),
		"id_token_ttl":                int(idp.GetOidcIdTokenTimeToLive()),
		"user_claims_in_access_token": bool(idp.GetOidcIncludeUserClaimsInAccessToken()),
	}
	result = append(result, oidc_map)

	return result, nil
}

func convertVPIdPSaml2MapArrToDTO(saml2_arr interface{}, idp *api.VirtualSaml2ServiceProviderDTO) error {
	saml2_map, err := asTFMapSingle(saml2_arr)
	if err != nil {
		return err
	}

	if saml2_map == nil {
		return nil
	}

	idp.SetWantAuthnRequestsSigned(saml2_map["want_authn_req_signed"].(bool))
	idp.SetWantSignedRequests(saml2_map["want_req_signed"].(bool))
	idp.SetSignRequests(saml2_map["sign_reqs"].(bool))
	idp.SetIdpSignatureHash(saml2_map["signature_hash"].(string))
	idp.SetEncryptAssertionAlgorithm(saml2_map["encrypt_algorithm"].(string))
	//idp.SetEnableMetadataEndpoint(saml2_map["metadata_endpoint"].(bool))
	idp.SetEnableMetadataEndpoint(true)
	idp.SetMessageTtl(int32(saml2_map["message_ttl"].(int)))
	idp.SetMessageTtlTolerance(int32(saml2_map["message_ttl_tolerance"].(int)))

	return nil
}

func convertVPIdPSaml2DTOToMapArr(idp *api.VirtualSaml2ServiceProviderDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	saml2_map := map[string]interface{}{
		"want_authn_req_signed": idp.GetWantAuthnRequestsSigned(),
		"want_req_signed":       idp.GetWantSignedRequests(),
		"sign_reqs":             idp.GetSignRequests(),
		"signature_hash":        idp.GetIdpSignatureHash(),
		"encrypt_algorithm":     idp.GetEncryptAssertionAlgorithm(),
		//		"metadata_endpoint":     idp.GetEnableMetadataEndpoint(),
		"message_ttl":           int(idp.GetMessageTtl()),
		"message_ttl_tolerance": int(idp.GetMessageTtlTolerance()),
	}
	result = append(result, saml2_map)

	return result, nil
}

func convertVPSPSaml2MapArrToDTO(saml2_arr interface{}, sp *api.VirtualSaml2ServiceProviderDTO) error {
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
	sp.SetIdentityMappingPolicy(*im)

	sp.SetMessageTtl(api.AsInt32(m["message_ttl"], 300))
	sp.SetMessageTtlTolerance(api.AsInt32(m["message_ttl_tolerance"], 300))

	sp.SetSignRequests(api.AsBool(m["sign_requests"], false))
	sp.SetSignAuthenticationRequests(api.AsBool(m["sign_authentication_requests"], true))
	sp.SetSpSignatureHash(api.AsString(m["signature_hash"], "SHA-256"))
	sp.SetWantAssertionSigned(api.AsBool(m["want_assertion_signed"], false))

	return nil
}

func convertVPSPSaml2DTOToMapArr(sp *api.VirtualSaml2ServiceProviderDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)

	al := sp.GetAccountLinkagePolicy()
	im := sp.GetIdentityMappingPolicy()
	saml2_map := map[string]interface{}{
		"account_linkage":              al.GetLinkEmitterType(),
		"message_ttl":                  int(sp.GetMessageTtl()),
		"message_ttl_tolerance":        int(sp.GetMessageTtlTolerance()),
		"identity_mapping":             im.GetMappingType(),
		"sign_authentication_requests": sp.GetSignAuthenticationRequests(),
		"sign_requests":                sp.GetSignRequests(),
		"signature_hash":               sp.GetSpSignatureHash(),
		"want_assertion_signed":        sp.GetWantAssertionSigned(),
	}
	result = append(result, saml2_map)

	return result, nil
}
