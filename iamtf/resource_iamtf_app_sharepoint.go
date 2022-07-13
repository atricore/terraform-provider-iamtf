package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSharePoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSharePointCreate,
		ReadContext:   resourceSharePointRead,
		UpdateContext: resourceSharePointUpdate,
		DeleteContext: resourceSharePointDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"slo_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sts_signing_cert_subject": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"slo_location_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},
			"sts_encrypting_cert_subject": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ida": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "identity appliane name",
			},
		},
	}
}

func resourceSharePointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceSharePointCreate", "ida", d.Get("ida").(string))

	SharePoint, err := buildSharePointDTO(d)
	if err != nil {
		return diag.Errorf("failed to build SharePoint: %v", err)
	}
	l.Trace("resourceSharePointCreate", "ida", d.Get("ida").(string), "name", *SharePoint.Name)

	a, err := getJossoClient(m).CreateSharePointresource(d.Get("ida").(string), SharePoint)
	if err != nil {
		l.Debug("resourceSharePointCreate %v", err)
		return diag.Errorf("failed to create SharePoint: %v", err)
	}

	if err = buildSharePointResource(d, a); err != nil {
		l.Debug("resourceSharePointCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceSharePointCreate OK", "ida", d.Get("ida").(string), "name", *SharePoint.Name)

	return nil
}

func resourceSharePointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceSharePointRead", "ida", d.Get("ida").(string), "name", d.Id())
	SharePoint, err := getJossoClient(m).GetSharePointResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceSharePointRead %v", err)
		return diag.Errorf("resourceSharePointRead: %v", err)
	}
	if SharePoint.Name == nil || *SharePoint.Name == "" {
		l.Debug("resourceSharePointRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildSharePointResource(d, SharePoint); err != nil {
		l.Debug("resourceSharePointRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceSharePointRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceSharePointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceSharePointUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	SharePoint, err := buildSharePointDTO(d)
	if err != nil {
		l.Debug("resourceSharePointUpdate %v", err)
		return diag.Errorf("failed to build SharePoint: %v", err)
	}

	a, err := getJossoClient(m).UpdateSharePointResource(d.Get("ida").(string), SharePoint)
	if err != nil {
		l.Debug("resourceSharePointUpdate %v", err)
		return diag.Errorf("failed to update SharePoint: %v", err)
	}

	if err = buildSharePointResource(d, a); err != nil {
		l.Debug("resourceSharePointUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceSharePointUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceSharePointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceSharePointDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteSharePointResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceSharePointDelete %v", err)
		return diag.Errorf("failed to delete SharePoint: %v", err)
	}

	l.Debug("resourceSharePointDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildSharePointDTO(d *schema.ResourceData) (api.SharepointResourceDTO, error) {
	var err error
	dto := api.NewSharepointResourceDTO()

	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	_, e := d.GetOk("slo_location")
	if e {
		dto.SloLocation, err = PtrSchemaLocation(d, "slo_location")
		if err != nil {
			return *dto, fmt.Errorf("invalid slo_location %s", err)
		}
	}
	dto.SloLocationEnabled = &e

	dto.StsSigningCertSubject = PtrSchemaStr(d, "sts_signing_cert_subject")
	dto.StsEncryptingCertSubject = PtrSchemaStr(d, "sts_encrypting_cert_subject")

	return *dto, err
}

func buildSharePointResource(d *schema.ResourceData, dto api.SharepointResourceDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("slo_location", cli.LocationToStr(dto.SloLocation))
	_ = d.Set("sts_signing_cert_subject", cli.StrDeref(dto.StsSigningCertSubject))
	_ = d.Set("sts_encrypting_cert_subject", cli.StrDeref(dto.StsEncryptingCertSubject))

	return nil
}
