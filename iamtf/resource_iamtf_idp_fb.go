package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIdFacebook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdFacebookCreate,
		ReadContext:   resourceIdFacebookRead,
		UpdateContext: resourceIdFacebookUpdate,
		DeleteContext: resourceIdFacebookDelete,
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
				Description: "facebook application id",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "facebook application secret",
			},
			"authz_token_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "facebook authorization token endpoint",
			},
			"access_token_service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "facebook access token endpoint",
			},
			"scopes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "facebook premissions. These will be added to **email** and **public_profile**",
			},
			"user_fields": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "facebook user fields. These will be added to **id**, **name**, **email**, **first_name**, **last_name**",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceIdFacebookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdFacebookCreate", "ida", d.Get("ida").(string))

	idFacebook, err := buildIdFacebookDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idFacebook: %v", err)
	}
	l.Trace("resourceIdFacebookCreate", "ida", d.Get("ida").(string), "name", *idFacebook.Name)

	a, err := getJossoClient(m).CreateIdFacebook(d.Get("ida").(string), idFacebook)
	if err != nil {
		l.Debug("resourceIdFacebookCreate %v", err)
		return diag.Errorf("failed to create idFacebook: %v", err)
	}

	if err = buildIdFacebookResource(d, a); err != nil {
		l.Debug("resourceIdFacebookCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdFacebookCreate OK", "ida", d.Get("ida").(string), "name", *idFacebook.Name)

	return nil
}
func resourceIdFacebookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdFacebookRead", "ida", d.Get("ida").(string), "name", d.Id())
	idFacebook, err := getJossoClient(m).GetIdpFacebook(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdFacebookRead %v", err)
		return diag.Errorf("resourceIdFacebookRead: %v", err)
	}
	if idFacebook.Name == nil || *idFacebook.Name == "" {
		l.Debug("resourceIdFacebookRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdFacebookResource(d, idFacebook); err != nil {
		l.Debug("resourceIdFacebookRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdFacebookRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdFacebookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdFacebookUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idFacebook, err := buildIdFacebookDTO(d)
	if err != nil {
		l.Debug("resourceIdFacebookUpdate %v", err)
		return diag.Errorf("failed to build idFacebook: %v", err)
	}

	a, err := getJossoClient(m).UpdateIdpFacebook(d.Get("ida").(string), idFacebook)
	if err != nil {
		l.Debug("resourceIdFacebookUpdate %v", err)
		return diag.Errorf("failed to update idFacebook: %v", err)
	}

	if err = buildIdFacebookResource(d, a); err != nil {
		l.Debug("resourceIdFacebookUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdFacebookUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdFacebookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdFacebookDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdpFacebook(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdFacebookDelete %v", err)
		return diag.Errorf("failed to delete idFacebook: %v", err)
	}

	l.Debug("resourceIdFacebookDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdFacebookDTO(d *schema.ResourceData) (api.FacebookOpenIDConnectIdentityProviderDTO, error) {
	var err error
	dto := api.NewFacebookOpenIDConnectIdentityProviderDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ClientSecret = PtrSchemaStr(d, "client_secret")

	dto.AccessTokenService, err = PtrSchemaLocation(d, "access_token_service")
	if err != nil {
		return *dto, err
	}
	dto.AuthzTokenService, err = PtrSchemaLocation(d, "authz_token_service")
	if err != nil {
		return *dto, err
	}

	// list to space separated values
	if notEmpty, s := PtrSchemaAsSpacedList(d, "scopes"); notEmpty {
		dto.SetScopes(s)
	}

	if notEmpty, s := PtrSchemaAsSpacedList(d, "user_fields"); notEmpty {
		dto.SetUserFields(s)
	}

	return *dto, err
}

func buildIdFacebookResource(d *schema.ResourceData, dto api.FacebookOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("client_secret", cli.StrDeref(dto.ClientSecret))
	_ = d.Set("authz_token_service", cli.LocationToStr(dto.AuthzTokenService))
	_ = d.Set("access_token_service", cli.LocationToStr(dto.AccessTokenService))

	// space separated values to list

	if notEmtpy, ls := SpacedListToSet(dto.GetScopes()); notEmtpy {
		_ = d.Set("scopes", ls)
	}
	if notEmtpy, ls := SpacedListToSet(dto.GetUserFields()); notEmtpy {
		_ = d.Set("user_fields", ls)
	}

	return nil
}
