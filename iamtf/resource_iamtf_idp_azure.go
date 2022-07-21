package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceidAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceidAzureCreate,
		ReadContext:   resourceidAzureRead,
		UpdateContext: resourceidAzureUpdate,
		DeleteContext: resourceidAzureDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IdP name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IdP description",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "azure app client id",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "azure app secret",
			},
			"server_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "azure app server key",
			},
			"authz_token_service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "token endpoint: https://login.microsoft.com/{tenant}/oauth2/v2.0/authorize",
			},
			"access_token_service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "authorization endpiont: https://login.microsoft.com/{tenant}/oauth2/v2.0/token",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceidAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceidAzureCreate", "ida", d.Get("ida").(string))

	idAzure, err := buildidAzureDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idAzure: %v", err)
	}
	l.Trace("resourceidAzureCreate", "ida", d.Get("ida").(string), "name", *idAzure.Name)

	a, err := getJossoClient(m).CreateIdpAzure(d.Get("ida").(string), idAzure)
	if err != nil {
		l.Debug("resourceidAzureCreate %v", err)
		return diag.Errorf("failed to create idAzure: %v", err)
	}

	if err = buildidAzureResource(d, a); err != nil {
		l.Debug("resourceidAzureCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceidAzureCreate OK", "ida", d.Get("ida").(string), "name", *idAzure.Name)

	return nil
}
func resourceidAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceidAzureRead", "ida", d.Get("ida").(string), "name", d.Id())
	idAzure, err := getJossoClient(m).GetIdpAzure(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceidAzureRead %v", err)
		return diag.Errorf("resourceidAzureRead: %v", err)
	}
	if idAzure.Name == nil || *idAzure.Name == "" {
		l.Debug("resourceidAzureRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildidAzureResource(d, idAzure); err != nil {
		l.Debug("resourceidAzureRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceidAzureRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceidAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceidAzureUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idAzure, err := buildidAzureDTO(d)
	if err != nil {
		l.Debug("resourceidAzureUpdate %v", err)
		return diag.Errorf("failed to build idAzure: %v", err)
	}

	a, err := getJossoClient(m).UpdateIdpAzure(d.Get("ida").(string), idAzure)
	if err != nil {
		l.Debug("resourceidAzureUpdate %v", err)
		return diag.Errorf("failed to update idAzure: %v", err)
	}

	if err = buildidAzureResource(d, a); err != nil {
		l.Debug("resourceidAzureUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceidAzureUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceidAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceidAzureDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdpGoogle(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceidAzureDelete %v", err)
		return diag.Errorf("failed to delete idAzure: %v", err)
	}

	l.Debug("resourceidAzureDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildidAzureDTO(d *schema.ResourceData) (api.AzureOpenIDConnectIdentityProviderDTO, error) {
	var err error
	dto := api.NewAzureOpenIDConnectIdentityProviderDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.ClientId = PtrSchemaStr(d, "client_id")
	dto.ClientSecret = PtrSchemaStr(d, "client_secret")
	dto.ServerKey = PtrSchemaStr(d, "server_key")
	dto.AuthzTokenService, err = PtrSchemaLocation(d, "authz_token_service")
	if err != nil {
		return *dto, fmt.Errorf("invalid authz_token_service %s", err)
	}
	dto.AccessTokenService, err = PtrSchemaLocation(d, "access_token_service")
	if err != nil {
		return *dto, fmt.Errorf("invalid access_token_service %s", err)
	}
	return *dto, err
}

func buildidAzureResource(d *schema.ResourceData, dto api.AzureOpenIDConnectIdentityProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("client_id", cli.StrDeref(dto.ClientId))
	_ = d.Set("client_secret", cli.StrDeref(dto.ClientSecret))
	_ = d.Set("server_key", cli.StrDeref(dto.ServerKey))
	_ = d.Set("authz_token_service", cli.LocationToStr(dto.AuthzTokenService))
	_ = d.Set("access_token_service", cli.LocationToStr(dto.AccessTokenService))

	return nil
}
