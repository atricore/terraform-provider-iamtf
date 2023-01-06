package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceselfService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceselfServiceCreate,
		ReadContext:   resourceselfServiceRead,
		UpdateContext: resourceselfServiceUpdate,
		DeleteContext: resourceselfServiceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"element_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "idp description",
			},
			"service_connection": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "todo add description for service connection",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Todo",
						},
						"element_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Todo",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Todo",
						},
					},
				},
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}

func resourceselfServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceselfServiceCreate", "ida", d.Get("ida").(string))

	selfService, err := buildselfServiceDTO(d)
	if err != nil {
		return diag.Errorf("failed to build selfService: %v", err)
	}
	l.Trace("resourceselfServiceCreate", "ida", d.Get("ida").(string), "name", *selfService.Name)

	a, err := getJossoClient(m).CreateSelfServiceresource(d.Get("ida").(string), selfService)
	if err != nil {
		l.Debug("resourceselfServiceCreate %v", err)
		return diag.Errorf("failed to create selfService: %v", err)
	}

	if err = buildselfServiceResource(d, a); err != nil {
		l.Debug("resourceselfServiceCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceselfServiceCreate OK", "ida", d.Get("ida").(string), "name", *selfService.Name)

	return nil
}
func resourceselfServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceselfServiceRead", "ida", d.Get("ida").(string), "name", d.Id())
	selfService, err := getJossoClient(m).GetSelfServiceResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceselfServiceRead %v", err)
		return diag.Errorf("resourceselfServiceRead: %v", err)
	}
	if selfService.Name == nil || *selfService.Name == "" {
		l.Debug("resourceselfServiceRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildselfServiceResource(d, selfService); err != nil {
		l.Debug("resourceselfServiceRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceselfServiceRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceselfServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceselfServiceUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	selfService, err := buildselfServiceDTO(d)
	if err != nil {
		l.Debug("resourceselfServiceUpdate %v", err)
		return diag.Errorf("failed to build selfService: %v", err)
	}

	a, err := getJossoClient(m).UpdateSelfServiceResource(d.Get("ida").(string), selfService)
	if err != nil {
		l.Debug("resourceselfServiceUpdate %v", err)
		return diag.Errorf("failed to update selfService: %v", err)
	}

	if err = buildselfServiceResource(d, a); err != nil {
		l.Debug("resourceselfServiceUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceselfServiceUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceselfServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceselfServiceDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteSelfServiceResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceselfServiceDelete %v", err)
		return diag.Errorf("failed to delete selfService: %v", err)
	}

	l.Debug("resourceselfServiceDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildselfServiceDTO(d *schema.ResourceData) (api.SelfServicesResourceDTO, error) {
	var err, errWrap error
	dto := api.NewSelfServicesResourceDTO()
	dto.Description = PtrSchemaStr(d, "description")
	dto.ElementId = PtrSchemaStr(d, "element_id")
	//dto.Id = PtrSchemaInt64(d, "id")
	dto.Location, err = PtrSchemaLocation(d, "location")
	if err != nil {
		return *dto, err
	}
	dto.Name = PtrSchemaStr(d, "name")
	dto.Secret = PtrSchemaStr(d, "secret")

	// not soported by server
	// service, err := convertServiceConnectionMapArrToDTO(d.Get("service_connection"))
	// if err != nil {
	// 	errWrap = errors.Wrap(err, "service_connection")
	// }
	// dto.ServiceConnection = service

	return *dto, errWrap
}

func buildselfServiceResource(d *schema.ResourceData, dto api.SelfServicesResourceDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("element_id", cli.StrDeref(dto.ElementId))
	_ = d.Set("location", cli.LocationToStr(dto.Location))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("secret", cli.StrDeref(dto.Secret))

	// not soported by server
	// service, err := convertServiceConnectionToMapArr(dto.ServiceConnection)
	// if err != nil {
	// 	return err
	// }
	// _ = d.Set("service_connection", service)

	return nil
}
func convertServiceConnectionToMapArr(sc *api.ServiceConnectionDTO) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// If cc is null, return an empty map
	if sc == nil {
		return result, nil
	}

	cc_map := map[string]interface{}{
		"description": sc.GetDescription(),
		"element_id":  sc.GetElementId(),
		"name":        sc.GetName(),
	}
	result = append(result, cc_map)

	return result, nil
}

func convertServiceConnectionMapArrToDTO(cc_arr interface{}) (*api.ServiceConnectionDTO, error) {
	var cc *api.ServiceConnectionDTO
	tfMapLs, err := asTFMapSingle(cc_arr)
	if err != nil {
		return cc, err
	}
	// If map is empty, return nil
	if tfMapLs == nil || len(tfMapLs) == 0 {
		return cc, nil
	}

	nsc := api.NewServiceConnectionDTO()
	nsc.SetDescription(api.AsString(tfMapLs["description"], ""))
	nsc.SetElementId(api.AsString(tfMapLs["element_id"], ""))
	nsc.SetName(api.AsString(tfMapLs["name"], ""))
	return nsc, nil
}
