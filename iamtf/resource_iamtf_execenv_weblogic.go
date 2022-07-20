package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWebLogicExecenv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebLogicExecenvCreate,
		ReadContext:   resourceWebLogicExecenvRead,
		UpdateContext: resourceWebLogicExecenvUpdate,
		DeleteContext: resourceWebLogicExecenvDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "execution environment ",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "execution environment description",
			},
			"version": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: stringInSlice([]string{"9.2", "11", "12", "14"}),
				Description:      "Weblogic version",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "weblogic domain",
			},
			"target_jdk": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "target jdk",
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
				Description: "(activation) Weblogic server path",
			},
			"activation_remote_target": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(activation) activate using remote JOSSO server ",
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
				Description: "identity appliane name",
			},
		},
	}
}
func resourceWebLogicExecenvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceWebLogicExecenvCreate", "ida", d.Get("ida").(string))

	wl, err := buildWebLogicExecenvDTO(d)
	if err != nil {
		return diag.Errorf("failed to build wlexeenv: %v", err)
	}
	l.Trace("resourceWebLogicExecenvCreate", "ida", d.Get("ida").(string), "name", *wl.Name)

	a, err := getJossoClient(m).CreateWLogic(d.Get("ida").(string), wl)
	if err != nil {
		l.Debug("resourceWebLogicExecenvCreate %v", err)
		return diag.Errorf("failed to create wlexeenv: %v", err)
	}

	if err = buildWebLogicExecenvResource(d, a); err != nil {
		l.Debug("resourceWebLogicExecenvCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceWebLogicExecenvCreate OK", "ida", d.Get("ida").(string), "name", *wl.Name)

	return nil
}

func resourceWebLogicExecenvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceWebLogicExecenvRead", "ida", d.Get("ida").(string), "name", d.Id())
	wl, err := getJossoClient(m).GetWebLogic(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceWebLogicExecenvRead %v", err)
		return diag.Errorf("resourceWebLogicExecenvRead: %v", err)
	}
	if wl.Name == nil || *wl.Name == "" {
		l.Debug("resourceWebLogicExecenvRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildWebLogicExecenvResource(d, wl); err != nil {
		l.Debug("resourceWebLogicExecenvRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceWebLogicExecenvRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceWebLogicExecenvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceWebLogicExecenvUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	wl, err := buildWebLogicExecenvDTO(d)
	if err != nil {
		l.Debug("resourceWebLogicExecenvUpdate %v", err)
		return diag.Errorf("failed to build wlexeenv: %v", err)
	}

	a, err := getJossoClient(m).UpdateWLogic(d.Get("ida").(string), wl)
	if err != nil {
		l.Debug("resourceWebLogicExecenvUpdate %v", err)
		return diag.Errorf("failed to update wlexeenv: %v", err)
	}

	if err = buildWebLogicExecenvResource(d, a); err != nil {
		l.Debug("resourceWebLogicExecenvUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceWebLogicExecenvUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceWebLogicExecenvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceWebLogicExecenvDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteWLogic(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceWebLogicExecenvDelete %v", err)
		return diag.Errorf("failed to delete wlexeenv: %v", err)
	}

	l.Debug("resourceWebLogicExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}
func buildWebLogicExecenvDTO(d *schema.ResourceData) (api.WeblogicExecutionEnvironmentDTO, error) {
	var err error
	dto := api.NewWeblogicExecutionEnvironmentDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")
	pid, err := versionToWLogic(d.Get("version").(string))
	if err != nil {
		return *dto, err
	}
	dto.PlatformId = &pid
	dto.Domain = PtrSchemaStr(d, "domain")
	dto.TargetJDK = PtrSchemaStr(d, "target_jdk")

	dto.InstallDemoApps = PtrSchemaBool(d, "activation_install_samples")
	dto.InstallUri = PtrSchemaStr(d, "activation_path")
	dto.OverwriteOriginalSetup = PtrSchemaBool(d, "activation_override_setup")
	dto.Location = PtrSchemaStr(d, "activation_remote_target")

	// TODO : dto.SetBindingLocation

	return *dto, err
}

func buildWebLogicExecenvResource(d *schema.ResourceData, dto api.WeblogicExecutionEnvironmentDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))
	ver, err := platformIdVersion(cli.StrDeref(dto.PlatformId))
	if err != nil {
		return err
	}
	_ = d.Set("version", ver)

	_ = d.Set("domain", cli.StrDeref(dto.Domain))
	_ = d.Set("target_jdk", cli.StrDeref(dto.TargetJDK))

	_ = d.Set("activation_install_samples", cli.BoolDeref(dto.InstallDemoApps))
	_ = d.Set("activation_path", cli.StrDeref(dto.InstallUri))
	_ = d.Set("activation_override_setup", cli.BoolDeref(dto.OverwriteOriginalSetup))
	_ = d.Set("activation_remote_target", cli.StrDeref(dto.Location))

	return nil
}

func platformIdVersion(ver string) (string, error) {
	switch ver {
	case "wl92":
		return "9.2", nil
	case "wl11":
		return "10", nil
	case "wl12":
		return "12", nil
	case "wl14":
		return "14", nil

	}
	return "", fmt.Errorf("unknown version %s", ver)
}

func versionToWLogic(pid string) (string, error) {
	switch pid {
	case "9.2":
		return "wl92", nil
	case "10":
		return "wl11", nil
	case "12":
		return "wl12", nil
	case "14":
		return "wl14", nil
	}
	return "", fmt.Errorf("unknown version %s", pid)
}
