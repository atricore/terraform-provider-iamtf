package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSelfService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSelfServiceCreate,
		ReadContext:   resourceSelfServiceRead,
		UpdateContext: resourceSelfServiceUpdate,
		DeleteContext: resourceSelfServiceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceSelfServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceSelfServiceCreate", "ida", d.Get("ida").(string))

	selfService, err := buildSelfServiceDTO(d)
	if err != nil {
		return diag.Errorf("failed to build selfService: %v", err)
	}
	l.Trace("resourceSelfServiceCreate", "ida", d.Get("ida").(string), "name", *selfService.Name)

	a, err := getJossoClient(m).CreateSelfServiceResource(d.Get("ida").(string), selfService)
	if err != nil {
		l.Debug("resourceSelfServiceCreate %v", err)
		return diag.Errorf("failed to create selfService: %v", err)
	}

	if err = buildSelfServiceResource(d, a); err != nil {
		l.Debug("resourceSelfServiceCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceSelfServiceCreate OK", "ida", d.Get("ida").(string), "name", *selfService.Name)

	return nil
}
func resourceSelfServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceSelfServiceRead", "ida", d.Get("ida").(string), "name", d.Id())
	selfService, err := getJossoClient(m).GetSelfServiceResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceSelfServiceRead %v", err)
		return diag.Errorf("resourceSelfServiceRead: %v", err)
	}
	if selfService.Name == nil || *selfService.Name == "" {
		l.Debug("resourceSelfServiceRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildSelfServiceResource(d, selfService); err != nil {
		l.Debug("resourceSelfServiceRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceSelfServiceRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceSelfServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceSelfServiceUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	selfService, err := buildSelfServiceDTO(d)
	if err != nil {
		l.Debug("resourceSelfServiceUpdate %v", err)
		return diag.Errorf("failed to build selfService: %v", err)
	}

	a, err := getJossoClient(m).UpdateSelfServiceResource(d.Get("ida").(string), selfService)
	if err != nil {
		l.Debug("resourceSelfServiceUpdate %v", err)
		return diag.Errorf("failed to update selfService: %v", err)
	}

	if err = buildSelfServiceResource(d, a); err != nil {
		l.Debug("resourceSelfServiceUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceSelfServiceUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceSelfServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceSelfServiceDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteSelfServiceResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceSelfServiceDelete %v", err)
		return diag.Errorf("failed to delete selfService: %v", err)
	}

	l.Debug("resourceSelfServiceDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildSelfServiceDTO(d *schema.ResourceData) (api.SelfServicesResourceDTO, error) {
	var err error
	dto := api.NewSelfServicesResourceDTO()
	dto.Description = PtrSchemaStr(d, "description")
	dto.Name = PtrSchemaStr(d, "name")

	return *dto, err
}

func buildSelfServiceResource(d *schema.ResourceData, dto api.SelfServicesResourceDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	return nil
}
