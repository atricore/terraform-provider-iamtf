package iamtf

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIdentityAppliance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdentityApplianceCreate,
		ReadContext:   resourceIdentityApplianceRead,
		UpdateContext: resourceIdentityApplianceUpdate,
		DeleteContext: resourceIdentityApplianceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This MUST be a unique value, use only letters, numbers and hyphen (-)",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This MUST be a unique value, conforming JAVA package naming conventions. For example: **com.mycompany.appliance1.dev**",
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Provide a description for your identity appliance.",
				Optional:    true,
			},
			"location": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The location is the public base URL for the appliance, for example: https://mysso.mycompany.com, or http://localhost:8081",
				ValidateDiagFunc: locationIsValid(),
			},
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reserved for internal use",
			},
			"bundles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "list of additional OSGi bundles this appliance requires",
			},
		},
	}
}

func resourceIdentityApplianceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	c := getJossoClient(m)

	// Build DTO
	ida, err := buildIdentityAppliance(d)
	if err != nil {
		return diag.Errorf("failed to build identity appliance: %v", err)
	}

	l.Trace("resourceIdentityApplianceCreate", "name", *ida.Name)

	// Invoke API
	a, err := c.CreateAppliance(ida)
	if err != nil {
		return diag.Errorf("failed to create identity applinace: %v", err)
	}

	// Reade new resource
	if err = buildIdentityApplianceResource(d.Get("ida").(string), d, &a); err != nil {
		return diag.FromErr(err)
	}

	l.Debug("resourceIdentityApplianceCreate OK", "name", *ida.Name)

	return nil
}

func resourceIdentityApplianceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}
	ida, err := getJossoClient(m).GetAppliance(idaName, d.Id())
	if err != nil {
		return diag.Errorf("failed to get identity appliance: %v", err)
	}
	if ida.Name == nil || *ida.Name == "" {
		d.SetId("")
		return nil
	}
	if err = buildIdentityApplianceResource(idaName, d, &ida); err != nil {
		return diag.FromErr(err)
	}

	getLogger(m).Debug("resourceIdentityApplianceRead OK", "name", *ida.Name)

	return nil
}

func resourceIdentityApplianceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	cli := getJossoClient(m)

	// Build DTO
	ida, err := buildIdentityAppliance(d)
	if err != nil {
		return diag.Errorf("failed to build identity appliance: %v", err)
	}

	l.Trace("resourceIdentityApplianceUpdate", "name", *ida.Name)

	// Invoke API
	_, err = cli.UpdateAppliance(ida)
	if err != nil {
		return diag.Errorf("failed to update identity applinace: %v", err)
	}

	l.Debug("resourceIdentityApplianceUpdate OK", "name", *ida.Name)

	// Reade new resource
	return resourceIdentityApplianceRead(ctx, d, m)
}

func resourceIdentityApplianceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var deleted bool
	var err error

	ID := d.Id()
	if deleted, err = getJossoClient(m).DeleteAppliance(ID); err != nil {
		return diag.Errorf("failed to delete appliance [%s]: %v", ID, err)
	}
	if !deleted {
		return diag.Errorf("appliance %s NOT deleted", ID)
	}

	getLogger(m).Debug("resourceIdentityApplianceDelete OK", "name", ID)

	return nil
}

// Builds an IdentiyApplianceDefinitionDTO from the resource data
func buildIdentityAppliance(d *schema.ResourceData) (api.IdentityApplianceDefinitionDTO, error) {
	var err error
	a := api.NewIdentityApplianceDefinitionDTO()

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	a.Id = &id
	a.Name = PtrSchemaStr(d, "name")
	a.Namespace = PtrSchemaStr(d, "namespace")

	ru := convertInterfaceToStringSetNullable(d.Get("bundles"))
	a.RequiredBundles = ru

	// Add '/IDBUS/APPLIANCE-NAME' to location
	l := PtrSchemaStr(d, "location")
	location, err := cli.StrToLocation(fmt.Sprintf("%s/IDBUS/%s", *l, strings.ToUpper(a.GetName())))
	if err != nil {
		return *a, err
	}
	a.Location = location

	a.Description = PtrSchemaStr(d, "description")
	// TODO : More attributes

	// IDP Selector

	// Branding

	// Configuration ?! (i.e. configuration file?!)

	return *a, err
}

// Builds a resource data from IdentiyApplianceDefinitionDTO
func buildIdentityApplianceResource(idaName string, d *schema.ResourceData, iam *api.IdentityApplianceDefinitionDTO) error {

	id := strconv.FormatInt(cli.Int64Deref(iam.Id), 10)
	d.SetId(id)
	_ = d.Set("ida", idaName)
	_ = d.Set("element_id", cli.StrDeref(iam.ElementId))
	_ = d.Set("name", cli.StrDeref(iam.Name))
	_ = d.Set("namespace", cli.StrDeref(iam.Namespace))
	_ = d.Set("description", cli.StrDeref(iam.Description))
	_ = setNonPrimitives(d, map[string]interface{}{
		"bundles": convertStringSetToInterface(iam.GetRequiredBundles())})

	// Clear
	iam.Location.SetContext("")
	iam.Location.SetUri("")
	_ = d.Set("location", cli.LocationToStr(iam.Location))
	// TODO : More attributes

	return nil
}

func locationIsValid() schema.SchemaValidateDiagFunc {
	return func(i interface{}, k cty.Path) diag.Diagnostics {
		v, ok := i.(string)
		if !ok {
			return diag.Errorf("expected type of %v to be string", k)
		}
		l, err := cli.StrToLocation(v)
		if err != nil {
			return diag.Errorf("location [%s] is invalid %v", v, err)
		}

		if l.GetContext() != "" || l.GetUri() != "" {
			return diag.Errorf("location [%s] is invalid. Path is not allowed", v)
		}

		return nil
	}
}
