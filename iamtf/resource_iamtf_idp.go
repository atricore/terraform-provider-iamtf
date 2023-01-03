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
			"max_sessions_per_user": {
				Type:        schema.TypeInt,
				Description: "Max number of sessions per user, -1 unbounded. This will limit the amount of simutaneous SSO sessions a user can create.  Works in combination with **destroy_previous_session**.",
				Optional:    true,
				Default:     -1,
			},
			"destroy_previous_session": {
				Type:        schema.TypeBool,
				Description: "JOSSO will logout an existing session to avoid exceeding the max session per user. if false, login will be  denied after reaching the max.",
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

			"authn_basic": {
				Type:        schema.TypeList,
				Optional:    true,
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
						"crypt_salt_length": {
							Type:             schema.TypeInt,
							Description:      "crypt salt length (in bytes: 0, 8, 16, 24, 32, 48, 64, 128, 256)",
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
						"extension": customClassSchema(),
					},
				},
			},
			"authn_bind_ldap": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "LDAP bind authentication settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"initial_ctx_factory": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Java JNDI initial context factory (default: com.sun.jndi.ldap.LdapCtxFactory)",
							Default:     "com.sun.jndi.ldap.LdapCtxFactory",
						},
						"provider_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "LDAP server connection url: ldaps://localhost:636",
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "username credential to connect to the LDAP server",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "secret credential to connect to the LDAP server",
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
							Description:      "Support LDAP password policy management. Values : none, ldap-rfc-draft (default none)",
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
							Description: "The Idp will provide the configured user identifier, ignoring the requested type(SAML 2)",
						},
						"saml_authn_ctx": {
							Type:        schema.TypeString,
							Description: "reported SAML 2 authentication context class",
							Optional:    true,
							ValidateDiagFunc: stringInSlice([]string{
								"urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
								"urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport"}),
							Default: "urn:oasis:names:tc:SAML:2.0:ac:classes:Password",
						},
						"search_scope": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"base", "one", "subtree", "children"}),
							Default:          "subtree",
							Description:      "LDAP search scope. Values : base, one, subtree, children",
						},
						"referrals": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: stringInSlice([]string{"follow", "ignore"}),
							Default:          "follow",
							Description:      "how to process referrals in a directory node.  Values: follow, ignore",
						},
						"operational_attrs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Require LDAP operational attributes (useful for LDAP password policy management)",
						},
						"extension": customClassSchema(),
					},
				},
			},

			"authn_client_cert": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Basic authentication settings. JOSSO will verify user provided credentials (username, password) with stored values in an identity source",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"clr_enabled": {
							Type:        schema.TypeBool,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"crl_refresh_seconds": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"crl_url": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"ocsp_enabled": {
							Type:        schema.TypeBool,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"ocsp_server": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"ocspserver": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"uid": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"extension": customClassSchema(),
					},
				},
			},
			"authn_wia": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Windows Integrated Authentication. JOSSO will verify identity by contacting a domain controller",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Description: "windows domain",
							Required:    true,
						},
						"domain_controller": {
							Type:        schema.TypeString,
							Description: "domain controller server",
							Required:    true,
						},
						"host": {
							Type:        schema.TypeString,
							Description: "JOSSO hostname",
							Required:    true,
						},
						"overwrite_kerberos_setup": {
							Type:        schema.TypeBool,
							Description: "override JOSSO kerberos configuration",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "JOSSO server port",
							Required:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "JOSSO server protocol (http/https)",
							Required:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Required:    true,
						},
						"service_class": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Required:    true,
						},
						"service_name": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Required:    true,
						},
						"keytab": {
							Type:        schema.TypeString,
							Description: "Kerberos keytab file",
							Required:    true,
						},
						"extension": customClassSchema(),
					},
				},
			},
			"authn_oauth2_pre": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Basic authentication settings. JOSSO will verify user provided credentials (username, password) with stored values in an identity source",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authn_service": {
							Type:        schema.TypeString,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"external_auth": {
							Type:        schema.TypeBool,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"remember_me": {
							Type:        schema.TypeBool,
							Description: "authentiacation priority compared to other mechanisms (ascening order)",
							Optional:    true,
							Computed:    true,
						},
						"extension": customClassSchema(),
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

func idpAttributeProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		MinItems:    0,
		MaxItems:    1,
		Description: "attributes mappings.  Define IdP claim mappings.",

		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"profile": {
					Type:             schema.TypeString,
					Description:      "Attribute profile to use",
					Optional:         true,
					ValidateDiagFunc: stringInSlice([]string{"JOSSO", "BASIC", "ONE_TO_ONE", "CUSTOM"}),
					Default:          "JOSSO",
				},
				"include_unmapped_claims": {
					Type:        schema.TypeBool,
					Description: "when using a custom profile, include unmapped claims",
					Optional:    true,
					Default:     true,
				},
				"map": {
					Type:        schema.TypeSet,
					Optional:    true,
					MinItems:    1,
					Description: "Custom attribute mappings",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "name of the new claim",
							},
							"type": {
								Type:             schema.TypeString,
								Description:      "mapping type.  claim if you want to man an existing claim to a new name.  const if you want to add a new constant value as claim.  exp if you want to create a claim using an expression",
								Optional:         true,
								Default:          "claim",
								ValidateDiagFunc: stringInSlice([]string{"claim", "const", "exp"}),
							},
							"mapping": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "value to us when doing the mapping, depending on the type this will be an existing claim name, a constant value or an expresison value",
							},
							"format": {
								Type:             schema.TypeString,
								Optional:         true,
								Default:          "BASIC",
								ValidateDiagFunc: stringInSlice([]string{"BASIC", "URN"}),
								Description:      "how the new claim name should be formatted.  Basic if the name will be used as is, URN to add a default SAML2 urn to the claim",
							},
						},
					},
				},
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
	ap, err := convertAttributeProfileMapArrToDTOs(idp.GetName(), d.Get("attributes"))
	if err != nil {
		errWrap = errors.Wrap(err, "attributes")
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

	err = convertAuthnBindLdapMapArrToDTO(d.Get("authn_bind_ldap"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_bind_ldap")
	}

	err = convertClientCertAuthnSvcMapArrToDTO(d.Get("authn_client_cert"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_client_cert")
	}

	err = convertWindowsIntegratedAuthnMapArrToDTO(d.Get("authn_wia"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_wia")
	}

	err = convertAuthnOAuth2PreMapArrToDTO(d.Get("authn_oauth2_pre"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_oauth2_pre")
	}

	err = convertAuthnBasicMapArrToDTO(d.Get("authn_basic"), idp)
	if err != nil {
		errWrap = errors.Wrap(err, "authn_basic")
	}

	// Id sources
	id_sources := convertInterfaceToStringSetNullable(d.Get("id_sources"))
	idp.IdentityLookups = convertStringArrToIdLookups(id_sources)

	subjectAuthen, err := convertSubjectAuthnPoliciesMapArrToDTO(d.Get("subject_authn_policies"))
	if err != nil {
		errWrap = errors.Wrap(err, "subject_authn_policies")
	}
	idp.SubjectAuthnPolicies = subjectAuthen

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

	authnbind, err := convertAuthnBindLdapDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_bind_ldap", authnbind)

	authclient, err := convertClientCertAuthnSvcDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_client_cert", authclient)

	wia, err := convertWindowsIntegratedAuthnDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_wia", wia)

	oauth2, err := convertAuthnOAuth2PreDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_oauth2_pre", oauth2)

	basic_authn, err := convertAuthnBasicDTOToMapArr(&idp)
	if err != nil {
		return err
	}
	_ = d.Set("authn_basic", basic_authn)

	attributes, err := convertAttributeProfileDTOToMapArr(idp.AttributeProfile)
	if err != nil {
		return err
	}
	_ = d.Set("attributes", attributes)

	ids := convertIdLookupsToStringArr(idp.IdentityLookups)
	aggMap := map[string]interface{}{
		"id_sources": convertStringSetToInterface(ids),
	}
	err = setNonPrimitives(d, aggMap)

	subjetAuthen, err := convertSubjectAuthnPoliciesDTOToMapArr(idp.SubjectAuthnPolicies)
	if err != nil {
		return err
	}
	_ = d.Set("subject_authn_policies", subjetAuthen)

	return err
}

// --------------------------------------------------------------------
func convertAuthnBasicMapArrToDTO(authn_basic_arr interface{}, idp *api.IdentityProviderDTO) error {
	tfMapLs, err := asTFMapAll(authn_basic_arr)
	if err != nil {
		return err
	}
	if len(tfMapLs) < 1 {
		return nil
	}

	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	for _, e := range tfMapLs {
		tfMap := e.(map[string]interface{})
		ba := api.NewBasicAuthenticationDTOInit()

		ba.SetName(api.AsString(tfMap["name"], "basic-authn"))
		ba.SetPriority(api.AsInt32(tfMap["priority"], 0))
		ba.SetHashAlgorithm(api.AsString(tfMap["pwd_hash"], "SHA-256"))
		ba.SetHashEncoding(api.AsString(tfMap["pwd_encoding"], "BASE64"))
		ba.SetSaltLength(api.AsInt32(tfMap["crypt_salt_length"], 0))
		ba.SetSaltPrefix(api.AsString(tfMap["salt_prefix"], ""))
		ba.SetSaltSuffix(api.AsString(tfMap["salt_suffix"], ""))
		ba.SetSimpleAuthnSaml2AuthnCtxClass(api.AsString(tfMap["saml_authn_ctx"], "urn:oasis:names:tc:SAML:2.0:ac:classes:Password"))
		ba.SetEnabled(true)

		// cc_dto, err := convertCustomClassMapArrToDTO(("extension"))
		// if err != nil {
		// 	err = errors.Wrap(err, "extension")
		// }
		// ba.SetCustomClass(*cc_dto)

		idp.AddBasicAuthn(ba)
	}

	return nil

}

func convertAuthnBasicDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {

	bas, err := idp.GetBasicAuthns()
	if err != nil {
		return nil, err
	}
	if len(bas) < 1 {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0)
	for _, ba := range bas {

		//customClass, err := convertCustomClassDTOToMapArr(bas.CustomClass)
		// if err != nil {
		// 	return nil, err
		// }
		authn_basic_map := map[string]interface{}{
			"priority":          ba.GetPriority(),
			"pwd_hash":          ba.GetHashAlgorithm(),
			"pwd_encoding":      ba.GetHashEncoding(),
			"crypt_salt_length": ba.GetSaltLength(),
			"salt_prefix":       ba.GetSaltPrefix(),
			"salt_suffix":       ba.GetSaltSuffix(),
			"saml_authn_ctx":    ba.GetSimpleAuthnSaml2AuthnCtxClass(),
			//"extension":         customClass,
		}
		result = append(result, authn_basic_map)
	}

	return result, nil

}

// --------------------------------------------------------------------
func convertAuthnBindLdapMapArrToDTO(authn_bind_arr interface{}, idp *api.IdentityProviderDTO) error {
	tfMapLs, err := asTFMapAll(authn_bind_arr)
	if err != nil {
		return err
	}
	if len(tfMapLs) < 1 {
		return nil
	}
	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	for _, e := range tfMapLs {
		tfMap := e.(map[string]interface{})
		das := api.NewDirectoryAuthenticationServiceDTO()
		das.SetInitialContextFactory(api.AsString(tfMap["initial_ctx_factory"], "com.sun.jndi.ldap.LdapCtxFactory"))
		das.SetProviderUrl(api.AsString(tfMap["provider_url"], "ldap://localhost:10389"))
		das.SetPerformDnSearch(api.AsBool(tfMap["perform_dn_search"], false))
		das.SetPasswordPolicy(api.AsString(tfMap["password_policy"], ""))
		das.SetSecurityAuthentication(api.AsString(tfMap["authentication"], "simple"))
		das.SetUsersCtxDN(api.AsString(tfMap["users_ctx_dn"], "dc=example,dc=com,ou=IAM,ou=People"))
		das.SetPrincipalUidAttributeID(api.AsString(tfMap["userid_attr"], "uid"))
		das.SetSecurityPrincipal(api.AsString(tfMap["username"], "uid=admin,ou=system"))
		das.SetSecurityCredential(api.AsString(tfMap["password"], "secret"))
		das.SetLdapSearchScope(api.AsString(tfMap["search_scope"], "subtree"))
		das.SetSimpleAuthnSaml2AuthnCtxClass(api.AsString(tfMap["saml_authn_ctx"], ""))
		das.SetReferrals(api.AsString(tfMap["referrals"], "follow"))
		das.SetIncludeOperationalAttributes(api.AsBool(tfMap["operational_attrs"], false))

		cc_dto, err := convertCustomClassMapArrToDTO(("extension"))
		if err != nil {
			err = errors.Wrap(err, "extension")
		}
		das.SetCustomClass(*cc_dto)

		idp.AddDirectoryAuthnSvc(das, api.AsInt32(tfMap["priority"], 0))
	}

	return nil

}

func convertAuthnBindLdapDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {

	authnMechanisms := idp.GetAuthenticationMechanisms()
	authnTfMapLs := make([]map[string]interface{}, 0)
	// For each authn mech
	for _, am := range authnMechanisms {

		authnSvc := am.GetDelegatedAuthentication().AuthnService
		if authnSvc == nil {
			continue
		}

		if authnSvc.IsDirectoryAuthnSvs() {
			dirAuthnSvc, err := authnSvc.ToDirectoryAuthnSvc()
			if err != nil {
				return nil, err
			}
			customClass, err := convertCustomClassDTOToMapArr(dirAuthnSvc.CustomClass)
			if err != nil {
				return nil, err
			}
			authnTfMap := map[string]interface{}{
				"priority":            am.GetPriority(),
				"initial_ctx_factory": dirAuthnSvc.GetInitialContextFactory(),
				"provider_url":        dirAuthnSvc.GetProviderUrl(),
				"perform_dn_search":   dirAuthnSvc.GetPerformDnSearch(),
				"password_policy":     dirAuthnSvc.GetPasswordPolicy(),
				"authentication":      dirAuthnSvc.GetSecurityAuthentication(),
				"users_ctx_dn":        dirAuthnSvc.GetUsersCtxDN(),
				"userid_attr":         dirAuthnSvc.GetPrincipalUidAttributeID(),
				"username":            dirAuthnSvc.GetSecurityPrincipal(),
				"password":            dirAuthnSvc.GetSecurityCredential(),
				"search_scope":        dirAuthnSvc.GetLdapSearchScope(),
				"saml_authn_ctx":      dirAuthnSvc.GetSimpleAuthnSaml2AuthnCtxClass(),
				"referrals":           dirAuthnSvc.GetReferrals(),
				"operational_attrs":   dirAuthnSvc.GetIncludeOperationalAttributes(),
				"extension":           customClass,
			}
			authnTfMapLs = append(authnTfMapLs, authnTfMap)
		}

	}

	return authnTfMapLs, nil

}

// --------------------------------------------------------------------

// This takes a TF map and creates the corresponding DTOs
// AuthenticationMechanismDTO -> DelegatedAuthenticationDTO -> AuthenticationServiceDTO
func convertClientCertAuthnSvcMapArrToDTO(client_cert interface{}, idp *api.IdentityProviderDTO) error {
	tfMapLs, err := asTFMapAll(client_cert)
	if err != nil {
		return err
	}
	if len(tfMapLs) < 1 {
		return nil
	}

	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	for _, e := range tfMapLs {
		tfMap := e.(map[string]interface{})
		cas := api.NewClientCertAuthnServiceDTO()
		cas.SetClrEnabled(api.AsBool(tfMap["clr_enabled"], false))
		cas.SetCrlRefreshSeconds(api.AsInt32(tfMap["crl_refresh_seconds"], 0))
		cas.SetCrlUrl(api.AsString(tfMap["crl_url"], ""))
		cas.SetOcspEnabled(api.AsBool(tfMap["ocsp_enabled"], false))
		cas.SetOcspServer(api.AsString(tfMap["ocsp_server"], ""))
		cas.SetOcspserver(api.AsString(tfMap["ocspserver"], ""))
		cas.SetUid(api.AsString(tfMap["uid"], ""))

		cc_dto, err := convertCustomClassMapArrToDTO(("extension"))
		if err != nil {
			err = errors.Wrap(err, "extension")
		}
		cas.SetCustomClass(*cc_dto)

		idp.AddClientCertAuthnSvs(cas, api.AsInt32(tfMap["priority"], 0))
	}

	return nil
}

func convertClientCertAuthnSvcDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	authnMechanisms := idp.GetAuthenticationMechanisms()
	authnTfMapLs := make([]map[string]interface{}, 0)
	// For each authn mech
	for _, am := range authnMechanisms {

		authnSvc := am.GetDelegatedAuthentication().AuthnService
		if authnSvc == nil {
			continue
		}

		if authnSvc.IsClientCertAuthnSvs() {
			clientCertAuthnSvc, err := authnSvc.ToClientCertAuthnSvc()
			if err != nil {
				return nil, err
			}
			customClass, err := convertCustomClassDTOToMapArr(clientCertAuthnSvc.CustomClass)
			if err != nil {
				return nil, err
			}
			authnTfMap := map[string]interface{}{
				"priority":            am.GetPriority(),
				"clr_enabled":         clientCertAuthnSvc.GetClrEnabled(),
				"crl_refresh_seconds": clientCertAuthnSvc.GetCrlRefreshSeconds(),
				"crl_url":             clientCertAuthnSvc.GetCrlUrl(),
				"ocsp_enabled":        clientCertAuthnSvc.GetOcspEnabled(),
				"ocsp_server":         clientCertAuthnSvc.GetOcspServer(),
				"ocspserver":          clientCertAuthnSvc.GetOcspserver(),
				"uid":                 clientCertAuthnSvc.GetUid(),
				"extension":           customClass,
			}
			authnTfMapLs = append(authnTfMapLs, authnTfMap)
		}

	}

	return authnTfMapLs, nil

}

// --------------------------------------------------------------------

// This takes a TF map and creates the corresponding DTOs
// AuthenticationMechanismDTO -> DelegatedAuthenticationDTO -> AuthenticationServiceDTO
func convertWindowsIntegratedAuthnMapArrToDTO(windows_integrated interface{}, idp *api.IdentityProviderDTO) error {

	tfMapLs, err := asTFMapAll(windows_integrated)
	if err != nil {
		return err
	}
	if len(tfMapLs) < 1 {
		return nil
	}

	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	for _, e := range tfMapLs {
		tfMap := e.(map[string]interface{})
		wia := api.NewWindowsIntegratedAuthenticationDTO()
		wia.SetDomain(api.AsString(tfMap["domain"], ""))
		wia.SetDomainController(api.AsString(tfMap["domain_controller"], ""))
		wia.SetHost(api.AsString(tfMap["host"], ""))
		wia.SetOverwriteKerberosSetup(api.AsBool(tfMap["overwrite_kerberos_setup"], false))
		wia.SetPort(api.AsInt32(tfMap["port"], 0))
		wia.SetProtocol(api.AsString(tfMap["protocol"], ""))
		wia.SetServiceClass(api.AsString(tfMap["service_class"], ""))
		wia.SetServiceName(api.AsString(tfMap["service_name"], ""))

		kt := api.NewResourceDTO()
		kt.SetValue(api.AsString(tfMap["keytab"], ""))
		wia.SetKeyTab(*kt)

		// cc_dto, err := convertCustomClassMapArrToDTO(("extension"))
		// if err != nil {
		// 	err = errors.Wrap(err, "extension")
		// }
		// wia.SetCustomClass(*cc_dto)

		idp.AddWindowsIntegratedAuthn(wia, api.AsInt32(tfMap["priority"], 0))
	}

	return nil
}

func convertWindowsIntegratedAuthnDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	authnMechanisms := idp.GetAuthenticationMechanisms()
	authnTfMapLs := make([]map[string]interface{}, 0)
	// For each authn mech
	for _, am := range authnMechanisms {

		authnSvc := am.GetDelegatedAuthentication().AuthnService
		if authnSvc == nil {
			continue
		}

		if authnSvc.IsWindowsIntegratedAuthn() {
			wiaAuthnSvc, err := authnSvc.ToWindowsIntegratedAuthn()
			if err != nil {
				return nil, err
			}
			// customClass, err := convertCustomClassDTOToMapArr(wiaAuthnSvc.CustomClass)
			// if err != nil {
			// 	return nil, err
			// }

			authnTfMap := map[string]interface{}{
				"priority":                 am.GetPriority(),
				"domain":                   wiaAuthnSvc.GetDomain(),
				"domain_controller":        wiaAuthnSvc.GetDomainController(),
				"host":                     wiaAuthnSvc.GetHost(),
				"overwrite_kerberos_setup": wiaAuthnSvc.GetOverwriteKerberosSetup(),
				"port":                     wiaAuthnSvc.GetPort(),
				"protocol":                 wiaAuthnSvc.GetProtocol(),
				"service_class":            wiaAuthnSvc.GetServiceClass(),
				"service_name":             wiaAuthnSvc.GetServiceName(),
				"keytab":                   wiaAuthnSvc.GetKeyTab().Value,
				// "extension":                customClass,
			}
			authnTfMapLs = append(authnTfMapLs, authnTfMap)
		}

	}

	return authnTfMapLs, nil

}

// --------------------------------------------------------------------

// This takes a TF map and creates the corresponding DTOs, injecting them into the IDP
// AuthenticationMechanismDTO -> DelegatedAuthenticationDTO -> AuthenticationServiceDTO
func convertAuthnOAuth2PreMapArrToDTO(authn_oauth2_pre interface{}, idp *api.IdentityProviderDTO) error {

	tfMapLs, err := asTFMapAll(authn_oauth2_pre)
	if err != nil {
		return err
	}
	if len(tfMapLs) < 1 {
		return nil
	}

	if idp.AuthenticationMechanisms == nil {
		idp.AuthenticationMechanisms = make([]api.AuthenticationMechanismDTO, 0)
	}

	for _, e := range tfMapLs {
		tfMap := e.(map[string]interface{})
		oauth2 := api.NewOAuth2PreAuthenticationServiceDTO()
		oauth2.SetAuthnService(api.AsString(tfMap["authn_service"], ""))
		oauth2.SetExternalAuth(api.AsBool(tfMap["external_auth"], false))
		oauth2.SetRememberMe(api.AsBool(tfMap["remember_me"], false))

		// cc_dto, err := convertCustomClassMapArrToDTO(("extension"))
		// if err != nil {
		// 	err = errors.Wrap(err, "extension")
		// }
		// oauth2.SetCustomClass(*cc_dto)

		idp.AddOauth2PreAuthnSvs(oauth2, api.AsInt32(tfMap["priority"], 0))
	}

	return nil
}

func convertAuthnOAuth2PreDTOToMapArr(idp *api.IdentityProviderDTO) ([]map[string]interface{}, error) {
	bas, err := idp.GetOauth2PreAuthnSvs()
	if err != nil {
		return nil, err
	}
	if len(bas) < 1 {
		return nil, nil
	}

	authnMechanisms := idp.GetAuthenticationMechanisms()
	authnTfMapLs := make([]map[string]interface{}, 0)
	// For each authn mech
	for _, am := range authnMechanisms {

		authnSvc := am.GetDelegatedAuthentication().AuthnService
		if authnSvc == nil {
			continue
		}

		if authnSvc.IsOauth2PreAuthnSvc() {
			oauth2svc, err := authnSvc.ToOauth2PreAuthnSvs()
			if err != nil {
				return nil, err
			}
			// customClass, err := convertCustomClassDTOToMapArr(oauth2svc.CustomClass)
			// if err != nil {
			// 	return nil, err
			// }
			authnTfMap := map[string]interface{}{
				"priority":      am.GetPriority(),
				"authn_service": oauth2svc.GetAuthnService(),
				"external_auth": oauth2svc.GetExternalAuth(),
				"remember_me":   oauth2svc.GetRememberMe(),
				// "extension":     customClass,
			}
			authnTfMapLs = append(authnTfMapLs, authnTfMap)
		}

	}

	return authnTfMapLs, nil
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

func convertAttributeProfileMapArrToDTOs(provider_name string, attrs interface{}) (*api.AttributeProfileDTO, error) {

	attrMap, err := asTFMapSingle(attrs)
	if err != nil {
		return nil, err
	}

	profile := api.AsString(attrMap["profile"], "JOSSO")
	include_unmapped_claims := api.AsBool(attrMap["include_unmapped_claims"], true)

	switch profile {
	case "JOSSO":
		af := api.NewJOSSOAttributeProfileDTOInit("josso-built-in")
		return af.ToAttrProfile()
	case "BASIC":
		af := api.NewBasicAttributeProfileDTOInit("basic-built-in")
		return af.ToAttrProfile()
	case "ONE_TO_ONE":
		af := api.NewOneToOneAttributeProfileDTOInit("one-to-one-built-in")
		return af.ToAttrProfile()
	case "CUSTOM":

		af := api.NewAttriburteMapperProfileDTOInit(fmt.Sprintf("%s-attr", provider_name))

		af.SetIncludeNonMappedProperties(include_unmapped_claims)
		var mappings []api.AttributeMappingDTO

		am := attrMap["map"].(*schema.Set)

		for _, v := range am.List() {

			mappingMap := v.(map[string]interface{})

			m := api.NewAttributeMappingDTO()
			m.SetAttrName(mappingMap["name"].(string))
			m.SetReportedAttrName(api.ToAttributeMapping(mappingMap["type"].(string), mappingMap["mapping"].(string)))
			m.SetReportedAttrNameFormat(api.AsString(mappingMap["format"], "BASIC"))

			mappings = append(mappings, *m)
		}

		af.SetAttributeMaps(mappings)
		return af.ToAttrProfile()
	}

	return nil, fmt.Errorf("invalid profile type %s\n", profile)
}

func convertAttributeProfileDTOToMapArr(ap *api.AttributeProfileDTO) ([]map[string]interface{}, error) {
	var r []map[string]interface{}

	apMap := make(map[string]interface{})
	apMap["profile"] = ap.GetProfileType()

	if ap.GetProfileType() == "CUSTOM" {
		amp := ap.ToAttributeMapperProfile()
		apMap["include_unmapped_claims"] = amp.GetIncludeNonMappedProperties()

		var maps []map[string]interface{}
		for _, m := range amp.GetAttributeMaps() {
			mMap := make(map[string]interface{})
			mMap["type"] = m.GetType()
			mMap["name"] = m.GetAttrName()

			mMap["mapping"] = m.GetMapping()
			mMap["format"] = m.GetReportedAttrNameFormat()
			maps = append(maps, mMap)
		}

		apMap["map"] = maps
	}

	r = append(r, apMap)

	return r, nil
}
