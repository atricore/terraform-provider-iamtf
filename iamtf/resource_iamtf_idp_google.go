package iamtf

import (
	"context"
	"fmt"

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
				Computed:    true,
				Optional:    true,
				Description: "idGoogle name",
			},
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idGoogle elementId",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idGoogle description",
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idGoogle clientId",
			},
			"server_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idGoogle serverKey",
			},
			"authz_token_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idGoogle ",
			},
			"access_token_service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "idGoogle ",
			},
			"google_apps_domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "idGoogle ",
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
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ServerKey = PtrSchemaStr(d, "server_key")
	_, e := d.GetOk("location")
	if e {
		dto.AuthzTokenService, err = PtrSchemaLocation(d, "authz_token_service")
		if err != nil {
			return *dto, fmt.Errorf("invalid location %s", err)
		}
		dto.AccessTokenService, err = PtrSchemaLocation(d, "access_token_service")
		if err != nil {
			return *dto, fmt.Errorf("invalid location %s", err)
		}
	}
	dto.GoogleAppsDomain = PtrSchemaStr(d, "google_apps_domain")
	return *dto, err
}

func buildidGoogleResource(d *schema.ResourceData, dto api.GoogleOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("location", cli.LocationToStr(dto.Location))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("server_key", cli.StrDeref(dto.ServerKey))
	_ = d.Set("authz_token_service", cli.LocationToStr(dto.AuthzTokenService))
	_ = d.Set("access_token_service", cli.LocationToStr(dto.AccessTokenService))
	_ = d.Set("google_apps_domain", cli.StrDeref(dto.GoogleAppsDomain))

	return nil
}
