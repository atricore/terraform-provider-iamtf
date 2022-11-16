package iamtf

import (
	"context"
	"fmt"

	api "github.com/atricore/josso-api-go"
	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
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
			"app_slo_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SLO location URL",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sharepoint application description",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "resource name",
			},
			"sts_signing_cert_subject": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "STS signing certificate subject",
			},
			"app_slo_location_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "enable application SLO location",
			},
			"sts_encrypting_cert_subject": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "STS encrypting certificate subject",
			},
			"ida": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "identity appliane name",
			},
			"sp_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SAML SP internal name",
			},
			"keystore": keystoreSchema(),
			"saml2":    spSamlSchema(),
			"idp":      idpConnectionSchema(),
		},
	}
}

func resourceSharePointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourceSharePointCreate", "ida", d.Get("ida").(string))

	josso1re, sp, err := buildSharePointDTO(d)
	if err != nil {
		return diag.Errorf("failed to build SharePoint: %v", err)
	}
	l.Trace("resourceSharePointCreate", "ida", d.Get("ida").(string), "name", *josso1re.Name)

	sp, err = getJossoClient(m).CreateIntSaml2Sp(d.Get("ida").(string), sp)
	if err != nil {
		l.Debug("resourceSharePointCreate %v", err)
		return diag.Errorf("failed to create Saml2 SP: %v", err)
	}

	josso1re.NewServiceConnection(sp.GetName())
	josso1re, err = getJossoClient(m).CreateSharePointresource(d.Get("ida").(string), josso1re)
	if err != nil {
		l.Debug("resourceSharePointCreate %v", err)
		return diag.Errorf("failed to create SharePoint: %v", err)
	}

	if err = buildSharePointResource(d, josso1re, sp); err != nil {
		l.Debug("resourceSharePointCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourceSharePointCreate OK", "ida", d.Get("ida").(string), "name", *josso1re.Name)

	return nil
}

func resourceSharePointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourceIntSaml2spRead", "ida", d.Get("ida").(string), "spname", *PtrSchemaStr(d, "sp_id"))
	sp, err := getJossoClient(m).GetIntSaml2Sp(d.Get("ida").(string), *PtrSchemaStr(d, "sp_id"))
	if err != nil {
		l.Debug("resourceIntSaml2spRead %v", err)
		return diag.Errorf("resourceIntSaml2spRead: %v", err)
	}
	if sp.Name == nil || *sp.Name == "" {
		l.Debug("resourceIntSaml2spRead NOT FOUND")
		d.SetId("")
		return nil
	}
	l.Debug("resourceIntSaml2spRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	l.Trace("resourceSharePointRead", "ida", d.Get("ida").(string), "name", d.Id())
	josso1re, err := getJossoClient(m).GetSharePointResource(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourceSharePointRead %v", err)
		return diag.Errorf("resourceSharePointRead: %v", err)
	}
	if josso1re.Name == nil || *josso1re.Name == "" {
		l.Debug("resourceSharePointRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildSharePointResource(d, josso1re, sp); err != nil {
		l.Debug("resourceSharePointRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourceSharePointRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourceSharePointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourceSharePointUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	josso1re, sp, err := buildSharePointDTO(d)
	if err != nil {
		l.Debug("resourceSharePointUpdate %v", err)
		return diag.Errorf("failed to build SharePoint: %v", err)
	}

	b, err := getJossoClient(m).UpdateIntSaml2Sp(d.Get("ida").(string), sp)
	if err != nil {
		l.Debug("resourceSharePointUpdate/intsaml2sp %v", err)
		return diag.Errorf("failed to update IntSaml2sp: %v", err)
	}
	a, err := getJossoClient(m).UpdateSharePointResource(d.Get("ida").(string), josso1re)
	if err != nil {
		l.Debug("resourceSharePointUpdate/josso1re %v", err)
		return diag.Errorf("failed to update josso1re: %v", err)
	}

	if err = buildSharePointResource(d, a, b); err != nil {
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
		return diag.Errorf("failed to delete josso1re: %v", err)
	}
	l.Debug("resourceSharePointDelete OK", "ida", d.Get("ida").(string), "name", d.Id())
	d.Get("idps")
	l.Trace("resourceIntSaml2spExecenvDelete", "ida", d.Get("ida").(string), "name", d.Get("sp_id").(string))
	_, err = getJossoClient(m).DeleteIntSaml2Sp(d.Get("ida").(string), d.Get("sp_id").(string))
	if err != nil {
		l.Debug("resourceIntSaml2spExecenvDelete %v", err)
		return diag.Errorf("failed to delete saml2sp: %v", err)
	}

	l.Debug("resourceSharePointDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func buildSharePointDTO(d *schema.ResourceData) (api.SharepointResourceDTO, api.InternalSaml2ServiceProviderDTO, error) {
	var err, errWrap error

	// Resource (sharepoint inherits JOSSO1Resource)
	josso1re := api.NewSharepointResourceDTO()
	// SP
	sp := api.NewInternalSaml2ServiceProviderDTO()

	josso1re.Name = PtrSchemaStr(d, "name")
	josso1re.Description = PtrSchemaStr(d, "description")

	_, e := d.GetOk("app_slo_location")
	if e {
		josso1re.SloLocation, err = PtrSchemaLocation(d, "app_slo_location")
		if err != nil {
			return *josso1re, *sp, fmt.Errorf("invalid app_slo_location %s", err)
		}
	}
	josso1re.SloLocationEnabled = &e

	josso1re.StsSigningCertSubject = PtrSchemaStr(d, "sts_signing_cert_subject")
	josso1re.StsEncryptingCertSubject = PtrSchemaStr(d, "sts_encrypting_cert_subject")

	// --------------------------------------------------------
	// SP
	// --------------------------------------------------------
	// On create sp_id is empty, on update it has a valid value
	spName := PtrSchemaStr(d, "sp_id")
	if *spName == "" {
		// This is a create SP
		spName = PtrSchemaStr(d, "name")
		*spName = fmt.Sprintf("%s-sp", *spName)
	}
	sp.Name = spName

	sp.DashboardUrl = PtrSchemaStr(d, "dashboard_url")
	sp.Description = PtrSchemaStr(d, "description")
	sp.DisplayName = PtrSchemaStr(d, "version")
	sp.ErrorBinding = PtrSchemaStr(d, "error_binding")

	// SP Configuration
	ks, err := convertKeystoreMapArrToDTO(sp.GetName(), d.Get("keystore"))
	if err != nil {
		errWrap = errors.Wrap(err, "keystore")
	}

	cfg := api.NewSamlR2SPConfigDTOInit()
	cfg.SetSigner(*ks)
	cfg.SetEncrypter(*ks)
	cfg.SetUseSampleStore(false)
	cfg.SetUseSystemStore(false)

	sp.SetSamlR2SPConfig(cfg)

	// Some defaults

	// SAML2 settings
	err = convertSPSaml2MapArrToDTO(d.Get("saml2"), sp)
	if err != nil {
		errWrap = errors.Wrap(err, "saml2")
	}

	// SP side of federated connection is for the SP
	sp.FederatedConnectionsB, err = convertIdPFederatedConnectionsMapArrToDTOs(sp, d, d.Get("idp"))
	if err != nil {
		return *josso1re, *sp, err
	}

	// Copy preferred IDP channel values to SP
	_, idpChannel, err := getPreferredIdPChannel(sp)
	if err != nil {
		return *josso1re, *sp, err
	}
	if idpChannel == nil {
		return *josso1re, *sp, fmt.Errorf("iamtf_app_agent resource MUST have a preferred idp: %s", *josso1re.Name)
	}

	return *josso1re, *sp, errWrap
}

func buildSharePointResource(d *schema.ResourceData, josso1re api.SharepointResourceDTO, sp api.InternalSaml2ServiceProviderDTO) error {
	d.SetId(cli.StrDeref(josso1re.Name))
	_ = d.Set("sp_id", cli.StrDeref(sp.Name))
	_ = d.Set("name", cli.StrDeref(josso1re.Name))
	_ = d.Set("description", cli.StrDeref(josso1re.Description))
	_ = d.Set("app_slo_location", cli.LocationToStr(josso1re.SloLocation))
	_ = d.Set("sts_signing_cert_subject", cli.StrDeref(josso1re.StsSigningCertSubject))
	_ = d.Set("sts_encrypting_cert_subject", cli.StrDeref(josso1re.StsEncryptingCertSubject))

	// Reuse iamtf_app_agent utils

	saml2_m, err := convertSPSaml2DTOToMapArr(&sp)
	if err != nil {
		return err
	}
	_ = d.Set("saml2", saml2_m)

	idps, err := convertIdPFederatedConnectionsToMapArr(sp.FederatedConnectionsB)
	if err != nil {
		return err
	}
	_ = d.Set("idp", idps)

	return nil
}
