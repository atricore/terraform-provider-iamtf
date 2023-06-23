package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIISExecenv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIISExecenvCreate,
		ReadContext:   resourceIISExecenvRead,
		UpdateContext: resourceIISExecenvUpdate,
		DeleteContext: resourceIISExecenvDelete,
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
				Description: "(activation) IIS path",
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
func resourceIISExecenvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceIISExecenvCreate", "ida", d.Get("ida").(string))

	iis, err := buildIISExecenvDTO(d)
	if err != nil {
		return diag.Errorf("failed to build IIS exeenv: %v", err)
	}
	l.Trace("resourceIISExecenvCreate", "ida", d.Get("ida").(string), "name", *iis.Name)

	a, err := getJossoClient(m).CreateIISExeEnv(d.Get("ida").(string), iis)
	if err != nil {
		l.Debug("resourceIISExecenvCreate %v", err)
		return diag.Errorf("failed to create IIS exeenv: %v", err)
	}

	if err = buildIISExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIISExecenvCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceIISExecenvCreate OK", "ida", d.Get("ida").(string), "name", *iis.Name)

	return nil
}

func resourceIISExecenvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	idaName := d.Get("ida").(string)
	if idaName == "" {
		idaName = m.(*Config).appliance
	}

	l.Trace("resourceIISExecenvRead", "ida", idaName, "name", d.Id())
	iis, err := getJossoClient(m).GetIISExeEnv(idaName, d.Id())
	if err != nil {
		l.Debug("resourceIISExecenvRead %v", err)
		return diag.Errorf("resourceIISExecenvRead: %v", err)
	}
	if iis.Name == nil || *iis.Name == "" {
		l.Debug("resourceIISExecenvRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildIISExecenvResource(d.Get("ida").(string), d, iis); err != nil {
		l.Debug("resourceIISExecenvRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceIISExecenvRead OK", "ida", idaName, "name", d.Id())

	return nil
}

func resourceIISExecenvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceIISExecenvUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	iis, err := buildIISExecenvDTO(d)
	if err != nil {
		l.Debug("resourceIISExecenvUpdate %v", err)
		return diag.Errorf("failed to build IIS exeenv: %v", err)
	}

	a, err := getJossoClient(m).UpdateIISExeEnv(d.Get("ida").(string), iis)
	if err != nil {
		l.Debug("resourceIISExecenvUpdate %v", err)
		return diag.Errorf("failed to update IIS exeenv: %v", err)
	}

	if err = buildIISExecenvResource(d.Get("ida").(string), d, a); err != nil {
		l.Debug("resourceIISExecenvUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourceIISExecenvUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceIISExecenvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIISExecenvDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteIISExeEnv(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceIISExecenvDelete %v", err)
		return diag.Errorf("failed to delete IIS exeenv: %v", err)
	}

	l.Debug("resourceIISExecenvDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}
func buildIISExecenvDTO(d *schema.ResourceData) (api.WindowsIISExecutionEnvironmentDTO, error) {
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

func buildIISExecenvResource(idaName string, d *schema.ResourceData, dto api.WindowsIISExecutionEnvironmentDTO) error {
	d.SetId(cli.StrDeref(dto.Name))
	_ = d.Set("ida", idaName)
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
	case "iis-32", "iis32":
		return "32", nil
	case "iis-64", "iis64":
		return "64", nil
	}
	return "", fmt.Errorf("unknown architecture %s", ver)
}

func architectureToPlatformId(pid string) (string, error) {
	switch pid {
	case "32":
		return "iis32", nil
	case "64":
		return "iis64", nil
	}
	return "", fmt.Errorf("unknown architecture %s", pid)
}
