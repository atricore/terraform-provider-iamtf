package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceOidcRp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOidcRpCreate,
		ReadContext:   resourceOidcRpRead,
		UpdateContext: resourceOidcRpUpdate,
		DeleteContext: resourceOidcRpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OidcRp element ID",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliane name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "OidcRp name",
			},
			"description": {
				Type:        schema.TypeString,
				Description: "OidcRp description",
				Optional:    true,
			},
			"client_id": {
				Type:        schema.TypeString,
				Description: "client ID",
				Required:    true,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "client secret",
				Required:    true,
			},
			"client_authn": {
				Type:             schema.TypeString,
				Description:      "client authentication. Note: use 'NONE' will assume 'code challenge' is used",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "CLIENT_SECRET_BASIC", "CLIENT_SECRET_JWT", "CLIENT_SECRET_POST", "PRIVATE_KEY_JWT"}),
				Optional:         true,
				Default:          "CLIENT_SECRET_BASIC",
			},
			"redirect_uris": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see okta_app_oauth_redirect_uri for appending to this list in a decentralized way.",
			},
			"post_logout_redirect_uris": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "List of URIs for redirection after logout",
			},
			"response_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: stringInSlice([]string{"TOKEN", "CODE", "ID_TOKEN"}),
				},
				Required:    true,
				Description: "List of OIDC response type strings.",
			},
			"response_modes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: stringInSlice([]string{"QUERY", "JWT"}),
				},
				Required:    true,
				Description: "List of OIDC response type strings.",
			},
			"grant_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: stringInSlice([]string{"AUTHORIZATION_CODE", "REFRESH_TOKEN", "JWT_BEARER_PWD", "CLIENT_CREDENTIALS", "JWT_BEARER", "PASSWORD", "IMPLICIT", "SAML2_BEARER"}),
				},
				Required:    true,
				Description: "List of OIDC grant types. ",
			},
			"signature_alg": {
				Type:             schema.TypeString,
				Description:      "signature algorithm",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "HS256", "HS384", "HS512", "RS256", "RS384", "RS512"}),
				Optional:         true,
				Default:          "HS256",
			},
			"encryption_alg": {
				Type:             schema.TypeString,
				Description:      "encryption algorithm",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "RSA1_5", "A128KW", "A128GCMKW", "A192KW", "A192GCMKW", "A256KW", "A256GCMKW"}),
				Optional:         true,
				Default:          "NONE",
			},
			"encryption_method": {
				Type:             schema.TypeString,
				Description:      "encryption method",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM"}),
				Optional:         true,
				Default:          "NONE",
			},
			"idtoken_signature_alg": {
				Type:             schema.TypeString,
				Description:      "signature algorithm",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "HS256", "HS384", "HS512", "RS256", "RS384", "RS512"}),
				Optional:         true,
				Default:          "HS256",
			},
			"idtoken_encryption_alg": {
				Type:             schema.TypeString,
				Description:      "encryption algorithm",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "RSA1_5", "A128KW", "A128GCMKW", "A192KW", "A192GCMKW", "A256KW", "A256GCMKW"}),
				Optional:         true,
				Default:          "NONE",
			},
			"idtoken_encryption_method": {
				Type:             schema.TypeString,
				Description:      "encryption method",
				ValidateDiagFunc: stringInSlice([]string{"NONE", "A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM"}),
				Optional:         true,
				Default:          "NONE",
			},
			"idps": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "List of trusted IdPs. If empty, all IdPs in the appliance will be trusted",
			},
		},
	}
}

func resourceOidcRpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceOidcRpCreate", "ida", d.Get("ida").(string))

	oidcrp, err := buildOidcRpDTO(d)
	if err != nil {
		return diag.Errorf("failed to build oidcrp: %v", err)
	}
	l.Trace("resourceOidcRpCreate", "ida", d.Get("ida").(string), "name", *oidcrp.Name)

	a, err := getJossoClient(m).CreateOidcRp(d.Get("ida").(string), oidcrp)
	if err != nil {
		l.Debug("resourceOidcRpCreate %v", err)
		return diag.Errorf("failed to create oidcrp: %v", err)
	}

	if err = buildOidcRpResource(d, a); err != nil {
		l.Debug("resourceOidcRpCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceOidcRpCreate OK", "ida", d.Get("ida").(string), "name", *oidcrp.Name)

	return nil
}

func resourceOidcRpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceOidcRpRead", "ida", d.Get("ida").(string), "name", d.Id())
	oidcrp, err := getJossoClient(m).GetOidcRp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceOidcRpRead %v", err)
		return diag.Errorf("resourceOidcRpRead: %v", err)
	}
	if oidcrp.Name == nil || *oidcrp.Name == "" {
		l.Debug("resourceOidcRpRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildOidcRpResource(d, oidcrp); err != nil {
		l.Debug("resourceOidcRpRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceOidcRpRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceOidcRpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceOidcRpUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	oidcrp, err := buildOidcRpDTO(d)
	if err != nil {
		l.Debug("resourceOidcRpUpdate %v", err)
		return diag.Errorf("failed to build oidcrp: %v", err)
	}

	a, err := getJossoClient(m).UpdateOidcRp(d.Get("ida").(string), oidcrp)
	if err != nil {
		l.Debug("resourceOidcRpUpdate %v", err)
		return diag.Errorf("failed to update oidcrp: %v", err)
	}

	if err = buildOidcRpResource(d, a); err != nil {
		l.Debug("resourceOidcRpUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceOidcRpUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceOidcRpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceOidcRpDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteOidcRp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceOidcRpDelete %v", err)
		return diag.Errorf("failed to delete oidcrp: %v", err)
	}

	l.Debug("resourceOidcRpDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildOidcRpDTO(d *schema.ResourceData) (api.ExternalOpenIDConnectRelayingPartyDTO, error) {
	var err error
	dto := api.NewExternalOpenIDConnectRelayingPartyDTO()
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ClientSecret = PtrSchemaStr(d, "client_secret")
	dto.ClientAuthnMethod = PtrSchemaStr(d, "client_authn")

	ru := convertInterfaceToStringSetNullable(d.Get("redirect_uris"))
	dto.AuthorizedURIs = ru

	plru := convertInterfaceToStringSetNullable(d.Get("post_logout_redirect_uris"))
	dto.PostLogoutRedirectionURIs = plru

	rt := convertInterfaceToStringSetNullable(d.Get("response_types"))
	dto.ResponseTypes = rt

	/*rm := convertInterfaceToStringSetNullable(d.Get("response_modes"))
	dto.ResponseModes = &rm*/

	gt := convertInterfaceToStringSetNullable(d.Get("grant_types"))
	dto.Grants = gt
	dto.SigningAlg = PtrSchemaStr(d, "signature_alg")
	dto.EncryptionAlg = PtrSchemaStr(d, "encryption_alg")
	dto.EncryptionMethod = PtrSchemaStr(d, "encryption_method")
	dto.IdTokenSigningAlg = PtrSchemaStr(d, "idtoken_signature_alg")
	dto.IdTokenEncryptionAlg = PtrSchemaStr(d, "idtoken_encryption_alg")
	dto.IdTokenEncryptionMethod = PtrSchemaStr(d, "idtoken_encryption_method")

	idps := convertInterfaceToStringSetNullable(d.Get("idps"))
	// Connetions to an IDP are always 'B'
	dto.FederatedConnectionsB = convertStringArrToFederatedConnections(idps)

	return *dto, err
}

func buildOidcRpResource(d *schema.ResourceData, dto api.ExternalOpenIDConnectRelayingPartyDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("client_secret", cli.StrDeref(dto.ClientSecret))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("client_authn", cli.StrDeref(dto.ClientAuthnMethod))
	_ = setNonPrimitives(d, map[string]interface{}{
		"redirect_uris": convertStringSetToInterface(dto.AuthorizedURIs)})
	_ = setNonPrimitives(d, map[string]interface{}{
		"post_logout_redirect_uris": convertStringSetToInterface(dto.PostLogoutRedirectionURIs)})
	_ = setNonPrimitives(d, map[string]interface{}{
		"response_types": convertStringSetToInterface(dto.ResponseTypes)})
	/*_ = setNonPrimitives(d, map[string]interface{}{
		"response_modes": convertStringSetToInterface(*dto.ResponseModes),
	})*/
	_ = setNonPrimitives(d, map[string]interface{}{
		"grant_types": convertStringSetToInterface(dto.Grants)})
	_ = d.Set("signature_alg", cli.StrDeref(dto.SigningAlg))
	_ = d.Set("encryption_alg", cli.StrDeref(dto.EncryptionAlg))
	_ = d.Set("encryption_method", cli.StrDeref(dto.EncryptionMethod))
	_ = d.Set("idtoken_signature_alg", cli.StrDeref(dto.IdTokenSigningAlg))
	_ = d.Set("idtoken_encryption_alg", cli.StrDeref(dto.IdTokenEncryptionAlg))
	_ = d.Set("idtoken_encryption_method", cli.StrDeref(dto.IdTokenEncryptionMethod))

	// TODO : FC A?
	idps := convertFederatedConnectionsToStringArr(dto.FederatedConnectionsB)
	aggMap := map[string]interface{}{
		"idps": convertStringSetToInterface(idps),
	}

	_ = setNonPrimitives(d, aggMap)

	return nil
}
