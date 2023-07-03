package iamtf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIdentityAppliance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIdentityApplianceRead,
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identity Appliance element id",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identity Appliance name",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This MUST be a unique value, conforming JAVA package naming conventions. For example: **com.mycompany.appliance1.dev**",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identity Appliance location",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identity Appliance description",
			},
			"bundles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "list of additional OSGi bundles this appliance requires",
			},
			"branding": {
				Type:        schema.TypeString,
				Description: "the name of the UI branding plugin installed in JOSSO",
				Default:     "josso25-branding",
				Optional:    true,
			},
		},
	}
}

func dataSourceIdentityApplianceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	idOrName := d.Get("name").(string)

	iam, err := getJossoClient(m).GetAppliance(idOrName)
	if err != nil {
		return diag.Errorf("failed to get identity appliance: %v", err)
	}
	if iam.Name == nil || *iam.Name == "" {
		d.SetId("")
		return nil
	}
	if err = buildIdentityApplianceResource(idOrName, d, &iam); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
