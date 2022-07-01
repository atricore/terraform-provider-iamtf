package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	sdk "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func ResourceIdP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdPCreate,
		ReadContext:   resourceIdPRead,
		UpdateContext: resourceIdPUpdate,
		DeleteContext: resourceIdPDelete,
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
				Description: "idp name, unique in the appliance scope",
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

			// ui properties
			"branding": {
				Type:        schema.TypeString,
				Description: "the name of the UI branding plugin installed in JOSSO",
				Default:     "josso25-branding",
				Optional:    true,
			},
			"dashboard_url": {
				Type:        schema.TypeString,
				Description: "URL to an external resource that can handle user UI",
				Optional:    true,
			},
			"error_binding": {
				Type:             schema.TypeString,
				Description:      "how the IDP reports errors, works combinded with **dashboard_url**",
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
			"max_sessions_per_user": {
				Type:        schema.TypeInt,
				Description: "Max number of sessions per user, -1 unbounded. This will limit the amount of simutaneous SSO sessions a user can create.  Works in combination with **destroy_previous_session**.",
				Optional:    true,
				Default:     -1,
			},
			"destroy_previous_session": {
				Type:        schema.TypeBool,
				Description: "If the max sessions per user is reached, JOSSO can destroy previously crated sessions (default), or prevent new logins",
				Optional:    true,
				Default:     true,
			},

			// SAML
			"saml2": idpSamlSchema(),

			"sp": spConnectionSchema(),

			// OAUTH2
			"oauth2": {
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
						"pwdless_authn_enabled": {
							Type:        schema.TypeBool,
							Description: "passwordless authentication enabled. Usefull for one-click logins",
							Optional:    true,
							Default:     false,
						},
						"pwdless_authn_subject": {
							Type:        schema.TypeString,
							Description: "message subject used during one-click login",
							Optional:    true,
						},
						"pwdless_authn_template": {
							Type:        schema.TypeString,
							Description: "name of the message template sent to the user during one-click login",
							Optional:    true,
						},
						"pwdless_authn_to": {
							Type:        schema.TypeString,
							Description: "passwordless authn subject TO",
							Optional:    true,
						},
						"pwdless_authn_from": {
							Type:        schema.TypeString,
							Description: "passwordless authn subject FROM",
							Optional:    true,
						},
					},
				},
			},

			// OAuth 2
			"oidc": {
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
					},
				},
			},

			"authn_bind_ldap": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    0,
				Description: "LDAP bind authentication settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"initial_ctx_factory": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Java JNDI initial context factory",
							Default:     "com.sun.jndi.ldap.LdapCtxFactory",
						},
						"provider_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "LDAP server connection url: ldaps://localhost:636",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "credential to connect to the LDAP server",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "credential to connect to the LDAP server",
						},
						"authentication": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"none", "strong", "simple"}),
							Default:          "simple",
							Description:      "credential to connect to the LDAP server",
						},
						"password_policy": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"none", "ldap-rfc-draft"}),
							Default:          "none",
							Description:      "Support LDAP password policy management.",
						},
						"perform_dn_search": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Perform a user search by DN before authentiation",
						},
						"users_ctx_dn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "DN to search for users",
						},
						"userid_attr": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "uid",
							Description: "LDAP attribute containing a user identifier",
						},
						"saml_authn_ctx": {
							Type:        schema.TypeString,
							Description: "password encoding algorithm",
							Optional:    true,
							ValidateDiagFunc: stringInSlice([]string{
								"urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
								"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport"}),
							Default: "urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
						},
						"search_scope": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"base", "one", "subtree", "chidlren"}),
							Default:          "subtree",
							Description:      "LDAP search scope",
						},
						"referrals": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"follow", "ignore"}),
							Default:          "follow",
							Description:      "how to process referrals in a directory node",
						},
						"operational_attrs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Require LDAP operational attributes (useful for LDAP password policy management)",
						},
					},
				},
			},

			"authn_basic": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    0,
				Description: "Basic authentication settings. JOSSO will verify user provided credentials (username, password) with stored values in an identity source",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"pwd_hash": {
							Type:             schema.TypeString,
							Description:      "password hashing algorithm, valid values are:\n* NONE (NOT recommended!)\n* CRYPT (LDAP only)\n* BCRYPT\n* SHA-512\n* SHA-256\n* SHA-1\n* MD5 (NOT recommended!)",
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"NONE", "CRYPT", "BCRYPT", "SHA-512", "SHA-256", "SHA-1", "MD5"}),
							Default:          "SHA-256",
						},
						"pwd_encoding": {
							Type:             schema.TypeString,
							Description:      "password encoding algorithm, valid values are:\n* NONE\n * BASE64\n * HEX",
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"NONE", "BASE64", "HEX"}),
							Default:          "BASE64",
						},
						"crypt_salt_lenght": {
							Type:             schema.TypeInt,
							Description:      "crypt salt lenght (in bytes: 0, 8, 16, 24, 32, 48, 64, 128, 256)",
							Optional:         true,
							Default:          0,
							ValidateDiagFunc: intInSlice([]int{0, 8, 16, 24, 32, 48, 64, 128, 256}),
						},
						"salt_prefix": {
							Type:        schema.TypeString,
							Description: "fixed salt prefix for password hashing",
							Optional:    true,
						},
						"salt_suffix": {
							Type:        schema.TypeString,
							Description: "fixed salt suffix for password hashing",
							Optional:    true,
						},
						"saml_authn_ctx": {
							Type:        schema.TypeString,
							Description: "reported SAML2 password authentication context. Some proivders required a specific value.  Valid values are: \n* urn:oasis:names:tc:SAML:2.0:ac:classes:Password (default)\n * urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport\n",
							Optional:    true,
							ValidateDiagFunc: stringInSlice([]string{
								"urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
								"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport"}),
							Default: "urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
						},
					},
				},
			},

			// TODO : "authn_wia"
			// TODO : "authn_client_cert"
			// TODO : "authn_oauth2_pre"

			"id_sources": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "list of identity sources used by the IDP.  At least one is required.",
			},
		},
	}
}

func resourceIdPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdPCreate", "ida", d.Get("ida").(string))

	idp, err := buildIdpDTO(d) //
	if err != nil {
		return diag.Errorf("failed to build idp: %v", err)
	}
	l.Trace("resourceIdPCreate", "ida", d.Get("ida").(string), "name", *idp.Name)

	idp, err = getJossoClient(m).CreateIdp(d.Get("ida").(string), idp)
	if err != nil {
		l.Debug("resourceIdPCreate %v", err)
		return diag.Errorf("failed to create idp: %v", err)
	}

	if err = buildIdPResource(d, idp); err != nil {
		l.Debug("resourceIdPCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdPCreate OK", "ida", d.Get("ida").(string), "name", *idp.Name)

	return nil
}
func resourceIdPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdPRead", "ida", d.Get("ida").(string), "name", d.Id())
	idp, err := getJossoClient(m).GetIdp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdPRead %v", err)
		return diag.Errorf("resourceIdPRead: %v", err)
	}
	if idp.Name == nil || *idp.Name == "" {
		l.Debug("resourceIdPRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdPResource(d, idp); err != nil {
		l.Debug("resourceIdPRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdPRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdPUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idp, err := buildIdpDTO(d)
	if err != nil {
		l.Debug("resourceIdPUpdate %v", err)
		return diag.Errorf("failed to build idp: %v", err)
	}

	idp, err = getJossoClient(m).UpdateIdp(d.Get("ida").(string), idp)
	if err != nil {
		l.Debug("resourceIdPUpdate %v", err)
		return diag.Errorf("failed to update idp: %v", err)
	}

	if err = buildIdPResource(d, idp); err != nil {
		l.Debug("resourceIdPUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdPUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdPDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdPDelete %v", err)
		return diag.Errorf("failed to delete idp: %v", err)
	}

	l.Debug("resourceIdPDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdpDTO(d *schema.ResourceData) (api.IdentityProviderDTO, error) {
	var err, errWrap error
	idp := api.NewIdentityProviderDTO()

	idp.Name = PtrSchemaStr(d, "name")
	idp.ElementId = PtrSchemaStr(d, "elementi_d")
	idp.Description = PtrSchemaStr(d, "description")

	// ui properties
	idp.UserDashboardBranding = PtrSchemaStr(d, "branding")
	idp.DashboardUrl = PtrSchemaStr(d, "dashboard_url")
	idp.ErrorBinding = PtrSchemaStr(d, "error_binding")

	// session properties
	idp.SsoSessionTimeout = PtrSchemaInt32(d, "session_timeout")
	idp.MaxSessionsPerUser = PtrSchemaInt32(d, "max_sessions_per_user")
	idp.DestroyPreviousSession = PtrSchemaBool(d, "destroy_previous_session")

	// IDP Configuration
	ks, err := convertKeystoreMapArrToDTO(idp.GetName(), d.Get("keystore"))
	if err != nil {
		errWrap = errors.Wrap(err, "keystore")
	}

	cfg := api.NewSamlR2IDPConfigDTOInit()
	cfg.SetSigner(*ks)
	cfg.SetEncrypter(*ks)
	cfg.SetUseSampleStore(false)
	cfg.SetUseSystemStore(false)

	idp.SetSamlR2IDPConfig(cfg)

	// Other sections
	err = convertIdPSaml2MapArrToDTO(d.Get("saml2"), idp) //
	if err != nil {
		errWrap = errors.Wrap(err, "saml2")
	}

	// Attribute profile
	af := api.NewBasicAttributeProfileDTOInit(fmt.Sprintf("%s-attr", idp.GetName()))
	ap, err := af.ToAttrProfile()
	if err != nil {
		errWrap = errors.Wrap(err, "attrs")
	}
	idp.SetAttributeProfile(*ap)

	// IDP side of federated connection is for the SP
	idp.FederatedConnectionsA, err = convertSPFederatedConnectionsMapArrToDTOs(idp, d, d.Get("sp"))
	if err != nil {
		errWrap = errors.Wrap(err, "sp")
	}

	err = convertOAuth2MapArrToDTO(d.Get("oauth2"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "oauth2")
	}

	err = convertOidcMapArrToDTO(d.Get("oidc"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "oidc")
	}

	// TODO : convertAuthnLdapBindMapArrToDTO()

	err = convertAuthnBasicMapArrToDTO(d.Get("authn_basic"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_basic")
	}

	err = convertAuthnOAuth2PreMapArrToDTO(d.Get("authn_oauth2_pre"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_oauth2_pre")
	}

	// Id sources
	id_sources := convertInterfaceToStringSetNullable(d.Get("id_sources"))
	idp.IdentityLookups = convertStringArrToIdLookups(id_sources)

	return *idp, errWrap
}

func buildIdPResource(d *schema.ResourceData, idp api.IdentityProviderDTO) error {
	d.SetId(sdk.StrDeref(idp.Name))
	_ = d.Set("name", sdk.StrDeref(idp.Name))
	_ = d.Set("element_id", sdk.StrDeref(idp.ElementId))
	_ = d.Set("description", sdk.StrDeref(idp.Description))

	_ = d.Set("branding", sdk.StrDeref(idp.UserDashboardBranding))
	_ = d.Set("dashboard_url", sdk.StrDeref(idp.DashboardUrl))
	_ = d.Set("error_binding", sdk.StrDeref(idp.ErrorBinding))

	_ = d.Set("session_timeout", sdk.Int32Deref(idp.SsoSessionTimeout))
	_ = d.Set("max_sessions_per_user", sdk.Int32Deref(idp.MaxSessionsPerUser))
	_ = d.Set("destroy_previous_session", sdk.BoolDeref(idp.DestroyPreviousSession))

	cfg, err := idp.GetSamlR2IDPConfig()
	if err != nil {
		return err
	}

	ks := cfg.GetSigner()
	ks_m, err := convertKeystoreDTOToMapArr(&ks)
	if err != nil {
		return err
	}

	_ = d.Set("keystore", ks_m)

	saml2_m, err := convertIdPSaml2DTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("saml2", saml2_m)

	// "sp" list
	sps, err := convertSPFederatedConnectionsToMapArr(idp.FederatedConnectionsA)
	if err != nil {
		return err
	}
	_ = d.Set("sp", sps)

	oauth2_m, err := convertOAuth2DTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("oauth2", oauth2_m)

	oidc_m, err := convertOidcDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("oidc", oidc_m)

	// TODO convertAuthnBasicDTOToMapArr

	basic_authn, err := convertAuthnBasicDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_basic", basic_authn)

	// TODO : Get from additional properties !?
	ids := convertIdLookupsToStringArr(idp.IdentityLookups)
	aggMap := map[string]interface{}{
		"id_sources": convertStringSetToInterface(ids),
	}

	err = setNonPrimitives(d, aggMap)

	return err
}

// --------------------------------------------------------------------

// --------------------------------------------------------------------

func convertAuthnBasicMapArrToDTO(authn_basic_arr interface{}, idp *api.IdentityProviderDTO) error {
	m, err := asTFMapSingle(authn_basic_arr)
	if err != nil {
		return err
	}

	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	ba := api.NewBasicAuthenticationDTOInit()

	ba.SetName(api.AsString(m["name"], "basic-authn"))
	ba.SetPriority(api.AsInt32(m["priority"], 0))
	ba.SetHashAlgorithm(api.AsString(m["pwd_hash"], "SHA-256"))
	ba.SetHashEncoding(api.AsString(m["pwd_encoding"], "BASE64"))
	ba.SetSaltLength(api.AsInt32(m["crypt_salt_lenght"], 0))
	ba.SetSaltPrefix(api.AsString(m["salt_prefix"], ""))
	ba.SetSaltSuffix(api.AsString(m["salt_suffix"], ""))
	ba.SetSimpleAuthnSaml2AuthnCtxClass(api.AsString(m["saml_authn_ctx"], "urn:oasis:names:tc:SAML:2.0:ac:classes:Password"))
	ba.SetEnabled(true)
	idp.AddBasicAuthn(ba)

	return nil

}

func convertAuthnBasicDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	bas, err := idp.GetBasicAuthns()
	if err != nil {
		return result, err
	}

	if len(bas) != 1 {
		return result, fmt.Errorf("too many basic authn defines for idp %s, %d", idp.GetName(), len(bas))
	}

	authn_basic := bas[0]

	authn_basic_map := map[string]interface{}{
		"priority":          authn_basic.GetPriority(),
		"pwd_hash":          authn_basic.GetHashAlgorithm(),
		"pwd_encoding":      authn_basic.GetHashEncoding(),
		"crypt_salt_lenght": authn_basic.GetSaltLength(),
		"salt_prefix":       authn_basic.GetSaltPrefix(),
		"salt_suffix":       authn_basic.GetSaltSuffix(),
		"saml_authn_ctx":    authn_basic.GetSimpleAuthnSaml2AuthnCtxClass(),
	}

	result = append(result, authn_basic_map)

	return result, nil

}

// --------------------------------------------------------------------

// This takes a TF map and creates the corresponding DTOs
// AuthenticationMechanismDTO -> DelegatedAuthenticationDTO -> AuthenticationServiceDTO
func convertAuthnOAuth2PreMapArrToDTO(authn_oauth2_pre interface{}, idp *api.IdentityProviderDTO) error {

	// 0. Search for authn mechanisms that have a service of type AuthenticationServiceDTO
	// user idp.GetOAuth2PreAuthns() : this func looks for all authns, that have a delegated authn containing an authn svc that isOauthPreAuthSvc() is true

	// 1. Create AuthenticationMechanismDTO (am), key attribute is priority
	// Priority and name from TF map (authn_oauth2_pre)
	// name = idp.name + '-oauth2_pre_authn-scheme'

	// 2. Create DelegatedAuthenticationDTO (da)
	// 3. Inject/Set da into am

	// 4. Create OAuth2PreAuthnServiceDTO (oaSvc)
	// name = idp.name + '-oauth2_pre_authn-svc'
	// More properties from TF map (authn_oauth2_pre)

	// 5. Conver oaSvc -> AuthenticationServiceDTO (aSvc)
	// 6. Inject oaSvc into da

	return nil
}

func convertAuthnOAuth2PreDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	// Ivert conversion from func above

	// 1. Create new map

	// 2. Populate with DTOs values
	return nil, nil
}

// --------------------------------------------------------------------

func convertOidcMapArrToDTO(oidc_arr interface{}, idp *api.IdentityProviderDTO) error {
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

	return nil

}

func convertOidcDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	oidc_map := map[string]interface{}{
		"enabled":          idp.GetOpenIdEnabled(),
		"access_token_ttl": int(idp.GetOidcAccessTokenTimeToLive()),
		"authz_code_ttl":   int(idp.GetOidcAuthzCodeTimeToLive()),
		"id_token_ttl":     int(idp.GetOidcIdTokenTimeToLive()),
	}
	result = append(result, oidc_map)

	return result, nil
}

// --------------------------------------------------------------------

func convertOAuth2MapArrToDTO(oauth2_arr interface{}, idp *api.IdentityProviderDTO) error {
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

	idp.SetPwdlessAuthnEnabled(oauth2_map["pwdless_authn_enabled"].(bool))
	idp.SetPwdlessAuthnSubject(oauth2_map["pwdless_authn_subject"].(string))
	idp.SetPwdlessAuthnTemplate(oauth2_map["pwdless_authn_template"].(string))
	idp.SetPwdlessAuthnTo(oauth2_map["pwdless_authn_to"].(string))
	idp.SetPwdlessAuthnFrom(oauth2_map["pwdless_authn_from"].(string))

	return nil
}

func convertOAuth2DTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	oauth2_map := map[string]interface{}{
		"enabled":                   idp.GetOauth2Enabled(),
		"shared_key":                idp.GetOauth2Key(),
		"token_validity":            int64(idp.GetOauth2TokenValidity()),
		"rememberme_token_validity": idp.GetOauth2RememberMeTokenValidity(),
		"pwdless_authn_enabled":     idp.GetPwdlessAuthnEnabled(),
		"pwdless_authn_subject":     idp.GetPwdlessAuthnSubject(),
		"pwdless_authn_template":    idp.GetPwdlessAuthnTemplate(),
		"pwdless_authn_to":          idp.GetPwdlessAuthnTo(),
		"pwdless_authn_from":        idp.GetPwdlessAuthnFrom(),
	}
	result = append(result, oauth2_map)

	return result, nil
}
