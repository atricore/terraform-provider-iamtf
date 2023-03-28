package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePhpExecenv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhpExecenvCreate,
		ReadContext:   resourcePhpExecenvRead,
		UpdateContext: resourcePhpExecenvUpdate,
		DeleteContext: resourcePhpExecenvDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Execution environment name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Execution environment description",
			},
			"activation_install_samples": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "(activation) install samples",
			},
			"activation_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(activation) PHP install path",
			},
			"activation_remote_target": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(activation) Activate using remote JOSSO server ",
			},
			"activation_override_setup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "(activation) Override agent setup",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliane name",
			},
			"environment": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringInSlice([]string{"STANDARD", "DRUPAL"}),
				Description:      "PHP install type",
			},
		},
	}
}
func resourcePhpExecenvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourcePhpExecenvCreate", "ida", d.Get("ida").(string))

	php, err := buildPhpExecenvDTO(d)
	if err != nil {
		return diag.Errorf("failed to build phpexeenv: %v", err)
	}
	l.Trace("resourcePhpExecenvCreate", "ida", d.Get("ida").(string), "name", *php.Name)

	a, err := getJossoClient(m).CreatePhpExeEnv(d.Get("ida").(string), php)
	if err != nil {
		l.Debug("resourcePhpExecenvCreate %v", err)
		return diag.Errorf("failed to create phpexeenv: %v", err)
	}

	if err = buildPhpExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourcePhpExecenvCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourcePhpExecenvCreate OK", "ida", d.Get("ida").(string), "name", *php.Name)

	return nil
}

func resourcePhpExecenvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourcePhpExecenvRead", "ida", idaName, "name", d.Id())
	php, err := getJossoClient(m).GetPhpExeEnv(idaName, d.Id())
	if err != nil {
		l.Debug("resourcePhpExecenvRead %v", err)
		return diag.Errorf("resourcePhpExecenvRead: %v", err)
	}
	if php.Name == nil || *php.Name == "" {
		l.Debug("resourcePhpExecenvRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildPhpExecenvResource(idaName, d, php); err != nil {
		l.Debug("resourcePhpExecenvRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourcePhpExecenvRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourcePhpExecenvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourcePhpExecenvUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	php, err := buildPhpExecenvDTO(d)
	if err != nil {
		l.Debug("resourcePhpExecenvUpdate %v", err)
		return diag.Errorf("failed to build phpexeenv: %v", err)
	}

	a, err := getJossoClient(m).UpdatePhpExeEnv(d.Get("ida").(string), php)
	if err != nil {
		l.Debug("resourcePhpExecenvUpdate %v", err)
		return diag.Errorf("failed to update phpexeenv: %v", err)
	}

	if err = buildPhpExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourcePhpExecenvUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourcePhpExecenvUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourcePhpExecenvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourcePhpExecenvDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeletePhpExeEnv(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourcePhpExecenvDelete %v", err)
		return diag.Errorf("failed to delete phpexeenv: %v", err)
	}

	l.Debug("resourcePhpExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}
func buildPhpExecenvDTO(d *schema.ResourceData) (api.PHPExecutionEnvironmentDTO, error) {
	var err error
	dto := api.NewPHPExecutionEnvironmentDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	dto.InstallDemoApps = PtrSchemaBool(d, "activation_install_samples")
	dto.InstallUri = PtrSchemaStr(d, "activation_path")
	dto.OverwriteOriginalSetup = PtrSchemaBool(d, "activation_override_setup")
	dto.Location = PtrSchemaStr(d, "activation_remote_target")
	dto.PhpEnvironmentType = PtrSchemaStr(d, "environment")

	return *dto, err
}

func buildPhpExecenvResource(idaName string, d *schema.ResourceData, dto api.PHPExecutionEnvironmentDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	_ = d.Set("activation_install_samples", cli.BoolDeref(dto.InstallDemoApps))
	_ = d.Set("activation_path", cli.StrDeref(dto.InstallUri))
	_ = d.Set("activation_override_setup", cli.BoolDeref(dto.OverwriteOriginalSetup))
	_ = d.Set("activation_remote_target", cli.StrDeref(dto.Location))
	_ = d.Set("environment", cli.StrDeref(dto.PhpEnvironmentType))

	return nil
}
