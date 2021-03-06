package iamtf

import (
	"context"

	api "github.com/atricore/josso-api-go"
	sdk "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcedbidSource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcedbidSourceCreate,
		ReadContext:   resourcedbidSourceRead,
		UpdateContext: resourcedbidSourceUpdate,
		DeleteContext: resourcedbidSourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"acquireincrement": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource name",
			},
			"admin": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource admin",
			},
			"connectionurl": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource connectionurl",
			},
			"credentialsquerystring": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource credentialsquerystring",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource description",
			},
			"drivername": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource drivername",
			},
			"idleconnectiontestperiod": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource idleconnectiontestperiod",
			},
			"initialpoolsize": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource initialpoolsize",
			},
			"maxidletime": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource maxidletime",
			},
			"maxpoolsize": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource maxpoolsize",
			},
			"minpoolsize": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "dbidentitysource minpoolsize",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource name",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource password",
			},
			"pooleddatasource": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "dbidentitysource pooleddatasource",
			},
			"relaycredentialquerystring": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource relaycredentialquerystring",
			},
			"resetcredentialdml": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource resetcredentialdml",
			},
			"rolesquerystring": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource rolesquerystring",
			},
			"usecolumnnamesaspropertynames": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "dbidentitysource usecolumnnamesaspropertynames",
			},
			"userpropertiesquerystring": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource userpropertiesquerystring",
			},
			"userquerystring": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource userquerystring",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource name",
			},
		},
	}
}

func resourcedbidSourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Debug("resourcedbidentitySourceCreate", "ida", d.Get("ida").(string))

	dbidentitySource, err := builddbidentitySourceDTO(d)
	if err != nil {
		return diag.Errorf("failed to build dbidentitySource: %v", err)
	}
	l.Trace("resourcedbidentitySourceCreate", "ida", d.Get("ida").(string), "name", *dbidentitySource.Name)

	a, err := getJossoClient(m).CreateDbIdentitySourceDTO(d.Get("ida").(string), dbidentitySource)
	if err != nil {
		l.Debug("resourcedbidentitySourceCreate %v", err)
		return diag.Errorf("failed to create dbidentitySource: %v", err)
	}

	if err = buildDbIdSourceResource(d, a); err != nil {
		l.Debug("resourcedbidentitySourceCreate %v", err)
		return diag.FromErr(err)
	}

	l.Debug("resourcedbidentitySourceCreate OK", "ida", d.Get("ida").(string), "name", *dbidentitySource.Name)

	return nil
}

func resourcedbidSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourcedbidentitySourceRead", "ida", d.Get("ida").(string), "name", d.Id())
	dbidentitySource, err := getJossoClient(m).GetDbIdentitySourceDTO(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourcedbidentitySourceRead %v", err)
		return diag.Errorf("resourcedbidentitySourceRead: %v", err)
	}
	if dbidentitySource.Name == nil || *dbidentitySource.Name == "" {
		l.Debug("resourcedbidentitySourceRead NOT FOUND")
		d.SetId("")
		return nil
	}
	if err = buildDbIdSourceResource(d, dbidentitySource); err != nil {
		l.Debug("resourcedbidentitySourceRead %v", err)
		return diag.FromErr(err)
	}
	l.Debug("resourcedbidentitySourceRead OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourcedbidSourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)
	l.Trace("resourcedbidentitySourceUpdate", "ida", d.Get("ida").(string), "name", d.Id())

	dbidentitySource, err := builddbidentitySourceDTO(d)
	if err != nil {
		l.Debug("resourcedbidentitySourceUpdate %v", err)
		return diag.Errorf("failed to build dbidentitySource: %v", err)
	}

	a, err := getJossoClient(m).UpdateDbIdentitySourceDTO(d.Get("ida").(string), dbidentitySource)
	if err != nil {
		l.Debug("resourcedbidentitySourceUpdate %v", err)
		return diag.Errorf("failed to update dbidentitySource: %v", err)
	}

	if err = buildDbIdSourceResource(d, a); err != nil {
		l.Debug("resourcedbidentitySourceUpdate %v", err)
		return diag.FromErr(err)
	}

	l.Trace("resourcedbidentitySourceUpdate OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func resourcedbidSourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l := getLogger(m)

	l.Trace("resourcedbidentitySourceDelete", "ida", d.Get("ida").(string), "name", d.Id())

	_, err := getJossoClient(m).DeleteDbIdentitySourceDTO(d.Get("ida").(string), d.Id())
	if err != nil {
		l.Debug("resourcedbidentitySourceDelete %v", err)
		return diag.Errorf("failed to delete dbidentitySource: %v", err)
	}

	l.Debug("resourcedbidentitySourceDelete OK", "ida", d.Get("ida").(string), "name", d.Id())

	return nil
}

func builddbidentitySourceDTO(d *schema.ResourceData) (api.DbIdentitySourceDTO, error) {
	var err error
	dto := api.NewDbIdentitySourceDTO()
	dto.AcquireIncrement = PtrSchemaInt32(d, "acquireincrement")
	dto.Admin = PtrSchemaStr(d, "admin")
	dto.ConnectionUrl = PtrSchemaStr(d, "connectionurl")
	dto.CredentialsQueryString = PtrSchemaStr(d, "credentialsquerystring")
	dto.Description = PtrSchemaStr(d, "description")
	dto.DriverName = PtrSchemaStr(d, "drivername")
	dto.IdleConnectionTestPeriod = PtrSchemaInt32(d, "idleconnectiontestperiod")
	dto.InitialPoolSize = PtrSchemaInt32(d, "initialpoolsize")
	dto.MaxIdleTime = PtrSchemaInt32(d, "maxidletime")
	dto.MaxPoolSize = PtrSchemaInt32(d, "maxpoolsize")
	dto.MinPoolSize = PtrSchemaInt32(d, "minpoolsize")
	dto.Name = PtrSchemaStr(d, "name")
	dto.Password = PtrSchemaStr(d, "password")
	dto.PooledDatasource = PtrSchemaBool(d, "pooleddatasource")
	dto.RelayCredentialQueryString = PtrSchemaStr(d, "relaycredentialquerystring")
	dto.ResetCredentialDml = PtrSchemaStr(d, "resetcredentialdml")
	dto.RolesQueryString = PtrSchemaStr(d, "rolesquerystring")
	dto.UseColumnNamesAsPropertyNames = PtrSchemaBool(d, "usecolumnnamesaspropertynames")
	dto.UserPropertiesQueryString = PtrSchemaStr(d, "userpropertiesquerystring")
	dto.UserQueryString = PtrSchemaStr(d, "userquerystring")

	return *dto, err
}

func buildDbIdSourceResource(d *schema.ResourceData, dto api.DbIdentitySourceDTO) error {
	d.SetId(sdk.StrDeref(dto.Name))
	_ = d.Set("acquireincrement", dto.GetAcquireIncrement())
	_ = d.Set("admin", dto.GetAdmin())
	_ = d.Set("connectionurl", dto.GetConnectionUrl())
	_ = d.Set("credentialsquerystring", dto.GetCredentialsQueryString())
	_ = d.Set("description", dto.GetDescription())
	_ = d.Set("drivername", dto.GetDriverName())
	_ = d.Set("idleconnectiontestperiod", dto.GetIdleConnectionTestPeriod())
	_ = d.Set("initialpoolsize", dto.GetInitialPoolSize())
	_ = d.Set("maxidletime", dto.GetMaxIdleTime())
	_ = d.Set("maxpoolsize", dto.GetMaxPoolSize())
	_ = d.Set("minpoolsize", dto.GetMinPoolSize())
	_ = d.Set("name", dto.GetName())
	_ = d.Set("password", dto.GetPassword())
	_ = d.Set("pooleddatasource", dto.GetPooledDatasource())
	_ = d.Set("relaycredentialquerystring", dto.GetRelayCredentialQueryString())
	_ = d.Set("resetcredentialdml", dto.GetResetCredentialDml())
	_ = d.Set("rolesquerystring", dto.GetRolesQueryString())
	_ = d.Set("usecolumnnamesaspropertynames", dto.GetUseColumnNamesAsPropertyNames())
	_ = d.Set("userpropertiesquerystring", dto.GetUserPropertiesQueryString())
	_ = d.Set("userquerystring", dto.GetUserQueryString())

	return nil
}
