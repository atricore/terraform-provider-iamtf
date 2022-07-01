package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceExtSaml2Sp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceExtSaml2SpCreate,
		ReadContext:   resourceExtSaml2SpRead,
		UpdateContext: resourceExtSaml2SpUpdate,
		DeleteContext: resourceExtSaml2SpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SAML 2 sp element ID (internal use)",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliane name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SAML 2 service provider name",
			},
			"description": {
				Type:        schema.TypeString,
				Description: "SAML 2 service provider description",
				Optional:    true,
			},
			"metadata": {
				Type:        schema.TypeString,
				Description: "SAML 2 service provider XML metadata descriptor",
				Required:    true,
			},
		},
	}
}

func resourceExtSaml2SpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceExtSaml2SpCreate", "ida", d.Get("ida").(string))

	extsaml2sp, err := buildExtSaml2SpDTO(d)
	if err != nil {
		return diag.Errorf("failed to build extsaml2sp: %v", err)
	}
	l.Trace("resourceExtSaml2SpCreate", "ida", d.Get("ida").(string), "name", *extsaml2sp.Name)

	a, err := getJossoClient(m).CreateExtSaml2Sp(d.Get("ida").(string), extsaml2sp)
	if err != nil {
		l.Debug("resourceExtSaml2SpCreate %v", err)
		return diag.Errorf("failed to create extsaml2sp: %v", err)
	}

	if err = buildExtSaml2SpResource(d, a); err != nil {
		l.Debug("resourceExtSaml2SpCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceExtSaml2SpCreate OK", "ida", d.Get("ida").(string), "name", *extsaml2sp.Name)

	return nil

}

func resourceExtSaml2SpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceExtSaml2SpRead", "ida", d.Get("ida").(string), "name", d.Id())
	extsaml2sp, err := getJossoClient(m).GetExtSaml2Sp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceExtSaml2SpRead %v", err)
		return diag.Errorf("resourceExtSaml2SpRead: %v", err)
	}
	if extsaml2sp.Name == nil || *extsaml2sp.Name == "" {
		l.Debug("resourceExtSaml2SpRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildExtSaml2SpResource(d, extsaml2sp); err != nil {
		l.Debug("resourceExtSaml2SpRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceExtSaml2SpRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceExtSaml2SpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceExtSaml2SpUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	extsaml2sp, err := buildExtSaml2SpDTO(d)
	if err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.Errorf("failed to build extsaml2sp: %v", err)
	}

	a, err := getJossoClient(m).UpdateExtSaml2Sp(d.Get("ida").(string), extsaml2sp)
	if err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.Errorf("failed to update extsaml2sp: %v", err)
	}

	if err = buildExtSaml2SpResource(d, a); err != nil {
		l.Debug("resourceExtSaml2SpUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceExtSaml2SpUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceExtSaml2SpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceExtSaml2SpDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteExtSaml2Sp(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceExtSaml2SpDelete %v", err)
		return diag.Errorf("failed to delete extsaml2sp: %v", err)
	}

	l.Debug("resourceExtSaml2SpDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildExtSaml2SpDTO(d *schema.ResourceData) (api.ExternalSaml2ServiceProviderDTO, error) {
	var err error
	dto := api.NewExternalSaml2ServiceProviderDTO()
	dto.ElementId = PtrSchemaStr(d, "element_id")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	m := api.NewResourceDTO()
	m.SetValue(*PtrSchemaStr(d, "metadata"))
	m.SetName(fmt.Sprintf("%s-md", dto.GetName()))
	dto.SetMetadata(*m)

	return *dto, err
}

func buildExtSaml2SpResource(d *schema.ResourceData, dto api.ExternalSaml2ServiceProviderDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))

	m := dto.GetMetadata()
	_ = d.Set("metadata", m.GetValue())

	return nil
}
