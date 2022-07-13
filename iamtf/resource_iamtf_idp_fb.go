package iamtf

import (
	"context"
	"fmt"

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
				Computed:    true,
				Optional:    true,
				Description: "idfacebook name",
			},
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook elementId",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook description",
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook location",
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook clientId",
			},
			"server_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook serverKey",
			},
			//"Authorization endpoint": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//	Description: "idfacebook ",
			//},
			//"Token endpoint": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "idfacebook ",
			//},
			//"Permissions": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "idfacebook ",
			//},
			"user_fields": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "idfacebook userFields",
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
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Description = PtrSchemaStr(d, "description")

	_, e := d.GetOk("location")
	if e {
		dto.Location, err = PtrSchemaLocation(d, "location")
		if err != nil {
			return *dto, fmt.Errorf("invalid location %s", err)
		}
	}
	dto.ClientId = PtrSchemaStr(d, "client_id")

	dto.ServerKey = PtrSchemaStr(d, "server_key")
	//a.ClientId = PtrSchemaStr(d, "Authorization endpoint")
	//a.ClientId = PtrSchemaStr(d, "Token endpoint")
	//a.ClientId = PtrSchemaStr(d, "Permissions")
	dto.UserFields = PtrSchemaStr(d, "user_fields")
	return *dto, err
}

func buildIdFacebookResource(d *schema.ResourceData, dto api.FacebookOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("location", cli.LocationToStr(dto.Location))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("server_key", cli.StrDeref(dto.ServerKey))
	//	_ = d.Set("Authorization endpoint", cli.StrDeref(idFacebook.Name))
	//	_ = d.Set("Token endpoint", cli.StrDeref(idFacebook.Name))
	//	_ = d.Set("Permissions", cli.StrDeref(idFacebook.Name))
	_ = d.Set("user_fields", cli.StrDeref(dto.UserFields))

	return nil
}
