package iamtf

import (
	"context"
	"os"

	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	identityAppliance = "iamtf_identity_appliance"
	idp               = "iamtf_idp"
	idVault           = "iamtf_idvault"
	idSourceLdap      = "iamtf_idsource_ldap"
	oidcRp            = "iamtf_app_oidc"
	extSaml2Sp        = "iamtf_app_saml2"
	tomcat            = "iamtf_execenv_tomcat"
	josso1Rs          = "iamtf_app_agent"
	dbidsource        = "iamtf_idsource_db"
	iss               = "iamtf_execenv_iss"
	php               = "iamtf_execenv_php"
	spoint            = "iamtf_app_sharepoint"
	weblogic          = "iamtf_execenv_weblogic"
	idFacebook        = "iamtf_idfacebook"
	azure             = "iamtf_idazure"
	google            = "iamtf_idgoogle"
)

var (
	apiClient *cli.IdbusApiClient
)

// Provider establishes a client connection to JOSSO server
// determined by its schema string values
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JOSSO_ORG_NAME", nil),
				Description: "Organization using JOSSO. Supports configuration from environment variable **JOSSO_ORG_NAME**",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JOSSO_API_ENDPOINT", nil),
				Description: "JOSSO Server endpoint, for example: http://localhost:8081/atricore-rest/services/iam-deploy. Supports configuration from environment variable **JOSSO_API_ENDPOINT**",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JOSSO_API_CLIENT_ID", nil),
				Description: "client identifier used to connect to the JOSSO server. Supports configuration from environment variable **JOSSO_API_CLIENT_ID**",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JOSSO_API_SECRET", nil),
				Description: "Secret used to connect to the JOSSO server. Supports configuration from environment variable **JOSSO_API_SECRET**",
			},
			"trace": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JOSSO_API_TRACE", nil),
				Description: "Trace API traffic (See also TF_LOG and TF_PROVIDER_LOG).  Supports configuration from environment variable **JOSSO_API_TRACE**",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			identityAppliance: ResourceIdentityAppliance(),
			idp:               ResourceIdP(),
			idVault:           ResourceIdVault(),
			idSourceLdap:      ResourceIdSourceLdap(),
			oidcRp:            ResourceOidcRp(),
			extSaml2Sp:        ResourceExtSaml2Sp(),
			tomcat:            ResourceTomcatExecenv(),
			josso1Rs:          ResourceJosso1Re(),
			dbidsource:        ResourcedbidSource(),
			iss:               ResourceIssExecenv(),
			php:               ResourcePhpExecenv(),
			spoint:            ResourceSharePoint(),
			weblogic:          ResourceWebLogicExecenv(),
			idFacebook:        ResourceIdFacebook(),
			azure:             ResourceidAzure(),
			google:            ResourceidGoogle(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			identityAppliance: dataSourceIdentityAppliance(),
		},
		ConfigureContextFunc: providerConfigure,
	}

}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	tfLog := os.Getenv("TF_LOG")
	tfLogProvider := os.Getenv("TF_LOG_PROVIDER")

	// Default level
	logLevel := hclog.LevelFromString("INFO")

	if tfLogProvider != "" {
		logLevel = hclog.LevelFromString(tfLogProvider)
	} else if tfLog != "" {
		logLevel = hclog.LevelFromString(tfLog)
	}

	config := Config{
		orgName:  d.Get("org_name").(string),
		clientId: d.Get("client_id").(string),
		secret:   d.Get("client_secret").(string),
		endpoint: d.Get("endpoint").(string),
		trace:    d.Get("trace").(bool),
		logLevel: int32(logLevel),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, diag.Errorf("[ERROR] Error initializing IAM.tf API cient: %v", err)
	}
	return &config, nil
}
