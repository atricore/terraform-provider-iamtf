package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIdpOidc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdpOidcCreate,
		ReadContext:   resourceIdpOidcRead,
		UpdateContext: resourceIdpOidcUpdate,
		DeleteContext: resourceIdpOidcDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "idp name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idp description",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "OIDC client id",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "OIDC client secret",
			},
			"server_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "OIDC server key (optional)",
			},
			"issuer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "OIDC issuer identifier",
			},
			"load_metadata": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "automatically load OIDC metadata from provider's well-known endpoint",
			},
			"authz_token_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "OIDC authorization endpoint (optional when load_metadata is true)",
			},
			"access_token_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "OIDC token endpoint (optional when load_metadata is true)",
			},
			"scopes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "OIDC scopes to request",
			},
			"user_fields": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "User fields to request from userinfo endpoint",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceIdpOidcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdpOidcCreate", "ida", d.Get("ida").(string))

	idpOidc, err := buildIdpOidcDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idpOidc: %v", err)
	}
	l.Trace("resourceIdpOidcCreate", "ida", d.Get("ida").(string), "name", *idpOidc.Name)

	a, err := getJossoClient(m).CreateOidcIdp(d.Get("ida").(string), idpOidc)
	if err != nil {
		l.Debug("resourceIdpOidcCreate %v", err)
		return diag.Errorf("failed to create idpOidc: %v", err)
	}

	if err = buildIdpOidcResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdpOidcCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdpOidcCreate OK", "ida", d.Get("ida").(string), "name", *idpOidc.Name)

	return nil
}

func resourceIdpOidcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIdpOidcRead", "ida", idaName, "name", d.Id())
	idpOidc, err := getJossoClient(m).GetOidcIdp(idaName, d.Id())
	if err != nil {
		l.Debug("resourceIdpOidcRead %v", err)
		return diag.Errorf("resourceIdpOidcRead: %v", err)
	}
	if idpOidc.Name == nil || *idpOidc.Name == "" {
		l.Debug("resourceIdpOidcRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdpOidcResource(idaName, d, idpOidc); err != nil {
		l.Debug("resourceIdpOidcRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdpOidcRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceIdpOidcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdpOidcUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idpOidc, err := buildIdpOidcDTO(d)
	if err != nil {
		l.Debug("resourceIdpOidcUpdate %v", err)
		return diag.Errorf("failed to build idpOidc: %v", err)
	}

	a, err := getJossoClient(m).UpdateOidcIdp(d.Get("ida").(string), idpOidc)
	if err != nil {
		l.Debug("resourceIdpOidcUpdate %v", err)
		return diag.Errorf("failed to update idpOidc: %v", err)
	}

	if err = buildIdpOidcResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdpOidcUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdpOidcUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdpOidcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdpOidcDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteOidcIdp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdpOidcDelete %v", err)
		return diag.Errorf("failed to delete idpOidc: %v", err)
	}

	l.Debug("resourceIdpOidcDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdpOidcDTO(d *schema.ResourceData) (api.GenericOpenIDConnectIdentityProviderDTO, error) {
	var err error
	dto := api.NewGenericOpenIDConnectIdentityProviderDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ClientSecret = PtrSchemaStr(d, "client_secret")
	dto.ServerKey = PtrSchemaStr(d, "server_key")
	dto.Issuer = PtrSchemaStr(d, "issuer")

	// Set LoadMetadata (defaults to true)
	loadMetadata := d.Get("load_metadata").(bool)
	dto.LoadMetadata = &loadMetadata

	// These are optional when load_metadata is true
	if _, ok := d.GetOk("access_token_service"); ok {
		dto.AccessTokenService, err = PtrSchemaLocation(d, "access_token_service")
		if err != nil {
			return *dto, err
		}
	}
	if _, ok := d.GetOk("authz_token_service"); ok {
		dto.AuthzTokenService, err = PtrSchemaLocation(d, "authz_token_service")
		if err != nil {
			return *dto, err
		}
	}

	if notEmpty, s := PtrSchemaAsSpacedList(d, "scopes"); notEmpty {
		dto.SetScopes(s)
	}

	if notEmpty, s := PtrSchemaAsSpacedList(d, "user_fields"); notEmpty {
		dto.SetUserFields(s)
	}

	return *dto, err
}

func buildIdpOidcResource(idaName string, d *schema.ResourceData, dto api.GenericOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("client_secret", cli.StrDeref(dto.ClientSecret))
	_ = d.Set("server_key", cli.StrDeref(dto.ServerKey))
	_ = d.Set("issuer", cli.StrDeref(dto.Issuer))
	_ = d.Set("load_metadata", cli.BoolDeref(dto.LoadMetadata))
	_ = d.Set("authz_token_service", cli.LocationToStr(dto.AuthzTokenService))
	_ = d.Set("access_token_service", cli.LocationToStr(dto.AccessTokenService))

	if notEmpty, ls := SpacedListToSet(dto.GetScopes()); notEmpty {
		_ = d.Set("scopes", ls)
	}
	if notEmpty, ls := SpacedListToSet(dto.GetUserFields()); notEmpty {
		_ = d.Set("user_fields", ls)
	}

	return nil
}
