package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceidGoogle() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceidGoogleCreate,
		ReadContext:   resourceidGoogleRead,
		UpdateContext: resourceidGoogleUpdate,
		DeleteContext: resourceidGoogleDelete,
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
				Description: "google application id",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "google application secret",
			},
			"authz_token_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "google authorization token endpoint",
			},
			"access_token_service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "google access token endpoint",
			},
			"google_apps_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "google suite domain ",
			},
			"google_apis": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "APIs to be accessed with the authorization token",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceidGoogleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceidGoogleCreate", "ida", d.Get("ida").(string))

	idGoogle, err := buildidGoogleDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idGoogle: %v", err)
	}
	l.Trace("resourceidGoogleCreate", "ida", d.Get("ida").(string), "name", *idGoogle.Name)

	a, err := getJossoClient(m).CreateIdpGoogle(d.Get("ida").(string), idGoogle)
	if err != nil {
		l.Debug("resourceidGoogleCreate %v", err)
		return diag.Errorf("failed to create idGoogle: %v", err)
	}

	if err = buildidGoogleResource(d, a); err != nil {
		l.Debug("resourceidGoogleCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceidGoogleCreate OK", "ida", d.Get("ida").(string), "name", *idGoogle.Name)

	return nil
}
func resourceidGoogleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceidGoogleRead", "ida", d.Get("ida").(string), "name", d.Id())
	idGoogle, err := getJossoClient(m).GetIdpGoogle(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceidGoogleRead %v", err)
		return diag.Errorf("resourceidGoogleRead: %v", err)
	}
	if idGoogle.Name == nil || *idGoogle.Name == "" {
		l.Debug("resourceidGoogleRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildidGoogleResource(d, idGoogle); err != nil {
		l.Debug("resourceidGoogleRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceidGoogleRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceidGoogleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceidGoogleUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idGoogle, err := buildidGoogleDTO(d)
	if err != nil {
		l.Debug("resourceidGoogleUpdate %v", err)
		return diag.Errorf("failed to build idGoogle: %v", err)
	}

	a, err := getJossoClient(m).UpdateIdpGoogle(d.Get("ida").(string), idGoogle)
	if err != nil {
		l.Debug("resourceidGoogleUpdate %v", err)
		return diag.Errorf("failed to update idGoogle: %v", err)
	}

	if err = buildidGoogleResource(d, a); err != nil {
		l.Debug("resourceidGoogleUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceidGoogleUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceidGoogleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceidGoogleDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdpFacebook(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceidGoogleDelete %v", err)
		return diag.Errorf("failed to delete idGoogle: %v", err)
	}

	l.Debug("resourceidGoogleDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildidGoogleDTO(d *schema.ResourceData) (api.GoogleOpenIDConnectIdentityProviderDTO, error) {
	var err error
	dto := api.NewGoogleOpenIDConnectIdentityProviderDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ClientSecret = PtrSchemaStr(d, "client_secret")
	//	dto.ServerKey = PtrSchemaStr(d, "server_key")
	dto.GoogleAppsDomain = PtrSchemaStr(d, "google_apps_domain")

	dto.AccessTokenService, err = PtrSchemaLocation(d, "access_token_service")
	if err != nil {
		return *dto, err
	}
	dto.AuthzTokenService, err = PtrSchemaLocation(d, "authz_token_service")
	if err != nil {
		return *dto, err
	}

	if notEmpty, s := PtrSchemaAsSpacedList(d, "google_apis"); notEmpty {
		dto.SetScopes(s)
	}

	return *dto, err
}

func buildidGoogleResource(d *schema.ResourceData, dto api.GoogleOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("client_secret", cli.StrDeref(dto.ClientSecret))
	_ = d.Set("authz_token_service", cli.LocationToStr(dto.AuthzTokenService))
	_ = d.Set("access_token_service", cli.LocationToStr(dto.AccessTokenService))
	_ = d.Set("google_apps_domain", cli.StrDeref(dto.GoogleAppsDomain))

	if notEmtpy, ls := SpacedListToSet(dto.GetScopes()); notEmtpy {
		_ = d.Set("google_apis", ls)
	}

	return nil
}
