package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIssExecenv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIssExecenvCreate,
		ReadContext:   resourceIssExecenvRead,
		UpdateContext: resourceIssExecenvUpdate,
		DeleteContext: resourceIssExecenvDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "resource name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IIS execution environment description",
			},
			"architecture": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: stringInSlice([]string{"32", "64"}),
				Description:      "IIS architecture. Values: 32, 64",
			},
			"activation_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(activation) Iss path",
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
			"isapi_extension_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IIS ISAPI filter URI (i.e. /josso)",
			},
		},
	}
}
func resourceIssExecenvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIssExecenvCreate", "ida", d.Get("ida").(string))

	iss, err := buildIssExecenvDTO(d)
	if err != nil {
		return diag.Errorf("failed to build issexeenv: %v", err)
	}
	l.Trace("resourceIssExecenvCreate", "ida", d.Get("ida").(string), "name", *iss.Name)

	a, err := getJossoClient(m).CreateIssExeEnv(d.Get("ida").(string), iss)
	if err != nil {
		l.Debug("resourceIssExecenvCreate %v", err)
		return diag.Errorf("failed to create issexeenv: %v", err)
	}

	if err = buildIssExecenvResource(d, a); err != nil {
		l.Debug("resourceIssExecenvCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIssExecenvCreate OK", "ida", d.Get("ida").(string), "name", *iss.Name)

	return nil
}

func resourceIssExecenvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIssExecenvRead", "ida", d.Get("ida").(string), "name", d.Id())
	iss, err := getJossoClient(m).GetIssExeEnv(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIssExecenvRead %v", err)
		return diag.Errorf("resourceIssExecenvRead: %v", err)
	}
	if iss.Name == nil || *iss.Name == "" {
		l.Debug("resourceIssExecenvRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIssExecenvResource(d, iss); err != nil {
		l.Debug("resourceIssExecenvRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIssExecenvRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIssExecenvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIssExecenvUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	iss, err := buildIssExecenvDTO(d)
	if err != nil {
		l.Debug("resourceIssExecenvUpdate %v", err)
		return diag.Errorf("failed to build issexeenv: %v", err)
	}

	a, err := getJossoClient(m).UpdateIssExeEnv(d.Get("ida").(string), iss)
	if err != nil {
		l.Debug("resourceIssExecenvUpdate %v", err)
		return diag.Errorf("failed to update issexeenv: %v", err)
	}

	if err = buildIssExecenvResource(d, a); err != nil {
		l.Debug("resourceIssExecenvUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIssExecenvUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIssExecenvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIssExecenvDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIssExeEnv(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIssExecenvDelete %v", err)
		return diag.Errorf("failed to delete issexeenv: %v", err)
	}

	l.Debug("resourceIssExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}
func buildIssExecenvDTO(d *schema.ResourceData) (api.WindowsIISExecutionEnvironmentDTO, error) {
	var err error
	dto := api.NewWindowsIISExecutionEnvironmentDTO()
	dto.Name = PtrSchemaStr(d, "name")
	dto.Description = PtrSchemaStr(d, "description")

	pid, err := architectureToPlatformId(d.Get("architecture").(string))
	if err != nil {
		return *dto, err
	}
	dto.PlatformId = &pid

	//dto.InstallDemoApps = PtrSchemaBool(d, "activation_install_samples")
	dto.InstallUri = PtrSchemaStr(d, "activation_path")
	dto.OverwriteOriginalSetup = PtrSchemaBool(d, "activation_override_setup")
	dto.Location = PtrSchemaStr(d, "activation_remote_target")
	dto.IsapiExtensionPath = PtrSchemaStr(d, "isapi_extension_path")
	return *dto, err
}

func buildIssExecenvResource(d *schema.ResourceData, dto api.WindowsIISExecutionEnvironmentDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("name", cli.StrDeref(dto.Name))
	_ = d.Set("description", cli.StrDeref(dto.Description))

	ver, err := platformIdToArchitecture(cli.StrDeref(dto.PlatformId))
	if err != nil {
		return err
	}
	_ = d.Set("architecture", ver)

	_ = d.Set("activation_path", cli.StrDeref(dto.InstallUri))
	_ = d.Set("activation_override_setup", cli.BoolDeref(dto.OverwriteOriginalSetup))
	_ = d.Set("activation_remote_target", cli.StrDeref(dto.Location))
	_ = d.Set("isapi_extension_path", cli.StrDeref(dto.IsapiExtensionPath))

	return nil
}

func platformIdToArchitecture(ver string) (string, error) {
	switch ver {
	case "iss-32":
		return "32", nil
	case "iss-64":
		return "64", nil
	}
	return "", fmt.Errorf("unknown architecture %s", ver)
}

func architectureToPlatformId(pid string) (string, error) {
	switch pid {
	case "32":
		return "iss-32", nil
	case "64":
		return "iss-64", nil
	}
	return "", fmt.Errorf("unknown architecture %s", pid)
}
