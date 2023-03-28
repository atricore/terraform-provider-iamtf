package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTomcatExecenv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTomcatExecenvCreate,
		ReadContext:   resourceTomcatExecenvRead,
		UpdateContext: resourceTomcatExecenvUpdate,
		DeleteContext: resourceTomcatExecenvDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "execution environment name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "execution environment description",
			},
			"version": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringInSlice([]string{"7", "8", "8.5", "9"}),
				Description:      "Tomcat version",
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
				Description: "(activation) Tomcat path",
			},
			"activation_remote_target": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(activation) activate using remote JOSSO server",
			},
			"activation_override_setup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "(activation) override agent setup",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
		},
	}
}
func resourceTomcatExecenvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceTomcatExecenvCreate", "ida", d.Get("ida").(string))

	tomcat, err := buildTomcatExecenvDTO(d)
	if err != nil {
		return diag.Errorf("failed to build tomcatexeenv: %v", err)
	}
	l.Trace("resourceTomcatExecenvCreate", "ida", d.Get("ida").(string), "name", *tomcat.Name)

	a, err := getJossoClient(m).CreateTomcatExeEnv(d.Get("ida").(string), tomcat)
	if err != nil {
		l.Debug("resourceTomcatExecenvCreate %v", err)
		return diag.Errorf("failed to create tomcatexeenv: %v", err)
	}

	if err = buildTomcatExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceTomcatExecenvCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceTomcatExecenvCreate OK", "ida", d.Get("ida").(string), "name", *tomcat.Name)

	return nil
}

func resourceTomcatExecenvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceTomcatExecenvRead", "ida", idaName, "name", d.Id())
	tomcat, err := getJossoClient(m).GetTomcatExeEnv(idaName, d.Id())
	if err != nil {
		l.Debug("resourceTomcatExecenvRead %v", err)
		return diag.Errorf("resourceTomcatExecenvRead: %v", err)
	}
	if tomcat.Name == nil || *tomcat.Name == "" {
		l.Debug("resourceTomcatExecenvRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildTomcatExecenvResource(idaName, d, tomcat); err != nil {
		l.Debug("resourceTomcatExecenvRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceTomcatExecenvRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceTomcatExecenvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceTomcatExecenvUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	tomcat, err := buildTomcatExecenvDTO(d)
	if err != nil {
		l.Debug("resourceTomcatExecenvUpdate %v", err)
		return diag.Errorf("failed to build tomcatexeenv: %v", err)
	}

	a, err := getJossoClient(m).UpdateTomcatExeEnv(d.Get("ida").(string), tomcat)
	if err != nil {
		l.Debug("resourceTomcatExecenvUpdate %v", err)
		return diag.Errorf("failed to update tomcatexeenv: %v", err)
	}

	if err = buildTomcatExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceTomcatExecenvUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceTomcatExecenvUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceTomcatExecenvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceTomcatExecenvDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteTomcatExeEnv(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceTomcatExecenvDelete %v", err)
		return diag.Errorf("failed to delete tomcatexeenv: %v", err)
	}

	l.Debug("resourceTomcatExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}
func buildTomcatExecenvDTO(d *schema.ResourceData) (api.TomcatExecutionEnvironmentDTO, error) {
	var err error
	dto := api.NewTomcatExecutionEnvironmentDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	pid, err := versionToPlatformId(d.Get("version").(string))
	if err != nil {
		return *dto, err
	}
	dto.PlatformId = &pid

	dto.InstallDemoApps = PtrSchemaBool(d, "activation_install_samples")
	dto.InstallUri = PtrSchemaStr(d, "activation_path")
	dto.OverwriteOriginalSetup = PtrSchemaBool(d, "activation_override_setup")
	dto.Location = PtrSchemaStr(d, "activation_remote_target")

	// TODO : dto.SetBindingLocation

	return *dto, err
}

func buildTomcatExecenvResource(idaName string, d *schema.ResourceData, dto api.TomcatExecutionEnvironmentDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))

	ver, err := platformIdToVersion(cli.StrDeref(dto.PlatformId))
	if err != nil {
		return err
	}
	_ = d.Set("version", ver)

	_ = d.Set("activation_install_samples", cli.BoolDeref(dto.InstallDemoApps))
	_ = d.Set("activation_path", cli.StrDeref(dto.InstallUri))
	_ = d.Set("activation_override_setup", cli.BoolDeref(dto.OverwriteOriginalSetup))
	_ = d.Set("activation_remote_target", cli.StrDeref(dto.Location))

	return nil
}

func platformIdToVersion(pid string) (string, error) {
	switch pid {
	case "tc50":
		return "5", nil
	case "tc55":
		return "5.5", nil
	case "tc60":
		return "6", nil
	case "tc70":
		return "7", nil
	case "tc80":
		return "8", nil
	case "tc85":
		return "8.5", nil
	case "tc90":
		return "9", nil
	}

	return "", fmt.Errorf("unknown platform-id %s", pid)

}

func versionToPlatformId(ver string) (string, error) {
	switch ver {
	case "5":
		return "tc50", nil
	case "5.5":
		return "tc55", nil
	case "6":
		return "tc60", nil
	case "7":
		return "tc70", nil
	case "8":
		return "tc80", nil
	case "8.5":
		return "tc85", nil
	case "9":
		return "tc90", nil

	}

	return "", fmt.Errorf("unknown version %s", ver)

}
