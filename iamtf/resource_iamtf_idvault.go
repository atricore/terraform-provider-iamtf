package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIdVault() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIdVaultCreate,
		ReadContext:   resourceIdVaultRead,
		UpdateContext: resourceIdVaultUpdate,
		DeleteContext: resourceIdVaultDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"element_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "element ID",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "id vault name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "id vault description",
			},
			"connector": {
				Type:        schema.TypeString,
				Description: "identity connector",
				Default:     "connector-default",
				Optional:    true,
			},
		},
	}
}

func resourceIdVaultCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIdVaultCreate", "ida", d.Get("ida").(string))

	idVault, err := buildIdVaultDTO(d)
	if err != nil {
		return diag.Errorf("failed to build idVault: %v", err)
	}
	l.Trace("resourceIdVaultCreate", "ida", d.Get("ida").(string), "name", *idVault.Name)

	a, err := getJossoClient(m).CreateIdVault(d.Get("ida").(string), idVault)
	if err != nil {
		l.Debug("resourceIdVaultCreate %v", err)
		return diag.Errorf("failed to create idVault: %v", err)
	}

	if err = buildIdVaultResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdVaultCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIdVaultCreate OK", "ida", d.Get("ida").(string), "name", *idVault.Name)

	return nil
}
func resourceIdVaultRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIdVaultRead", "ida", idaName, "name", d.Id())
	idVault, err := getJossoClient(m).GetIdVault(idaName, d.Id())
	if err != nil {
		l.Debug("resourceIdVaultRead %v", err)
		return diag.Errorf("resourceIdVaultRead: %v", err)
	}
	if idVault.Name == nil || *idVault.Name == "" {
		l.Debug("resourceIdVaultRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIdVaultResource(idaName, d, idVault); err != nil {
		l.Debug("resourceIdVaultRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIdVaultRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceIdVaultUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIdVaultUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	idVault, err := buildIdVaultDTO(d)
	if err != nil {
		l.Debug("resourceIdVaultUpdate %v", err)
		return diag.Errorf("failed to build idVault: %v", err)
	}

	a, err := getJossoClient(m).UpdateIdVault(d.Get("ida").(string), idVault)
	if err != nil {
		l.Debug("resourceIdVaultUpdate %v", err)
		return diag.Errorf("failed to update idVault: %v", err)
	}

	if err = buildIdVaultResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIdVaultUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIdVaultUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIdVaultDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIdVaultDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIdVault(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIdVaultDelete %v", err)
		return diag.Errorf("failed to delete idVault: %v", err)
	}

	l.Debug("resourceIdVaultDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildIdVaultDTO(d *schema.ResourceData) (api.EmbeddedIdentityVaultDTO, error) {
	var err error
	a := api.NewEmbeddedIdentityVaultDTO()
	a.ElementId = PtrSchemaStr(d, "element_id")
	a.Name = PtrSchemaStr(d, "name")
	a.Description = PtrSchemaStr(d, "description")
	a.IdentityConnectorName = PtrSchemaStr(d, "connector")

	return *a, err
}

func buildIdVaultResource(idaName string, d *schema.ResourceData, idVault api.EmbeddedIdentityVaultDTO) error {
	d.SetId(cli.StrDeref(idVault.Name))
	_ = d.Set("ida", idaName)
	_ = d.Set("element_id", cli.StrDeref(idVault.ElementId))
	_ = d.Set("name", cli.StrDeref(idVault.Name))
	_ = d.Set("description", cli.StrDeref(idVault.Description))
	_ = d.Set("connector", cli.StrDeref(idVault.IdentityConnectorName))

	return nil
}
