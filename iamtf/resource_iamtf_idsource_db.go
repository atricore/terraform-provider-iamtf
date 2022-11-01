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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identiy source name",
			},
			"ida": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "identity appliance name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "resource description",
			},

			// Connection
			"connectionurl": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "jdbc connection string",
			},
			"jdbc_driver": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "JDBC driver",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "connection username",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "connection password",
			},

			// SQL queries

			"sql_relay_credential": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource relaycredentialquerystring",
			},
			"dml_reset_credential": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "dbidentitysource resetcredentialdml",
			},
			"sql_groups": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "user groups query string.  Must return a single column with group names",
			},
			"sql_credentials": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "credentials query string. Must return a single row with columns: username, password, salt (optional)",
			},
			"sql_user_attrs": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "user attributes query string. Must return a single row with columns: username, name, value",
			},
			"use_column_name_as_property_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Use sql_user_attrs resultset column names as properties names",
			},
			"sql_username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "username query string. Used to retrieve the username from the DB",
			},

			// Connection pool

			"idle_connection_test_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "dbidentitysource idleconnectiontestperiod",
			},
			"acquireincrement": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "number of connections to aquire when incrementing the pool",
			},
			"initialpoolsize": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "dbidentitysource initialpoolsize",
			},
			"maxidletime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "dbidentitysource maxidletime",
			},
			"maxpoolsize": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "dbidentitysource maxpoolsize",
			},
			"minpoolsize": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "dbidentitysource minpoolsize",
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

	dto.Name = PtrSchemaStr(d, "name")

	dto.Admin = PtrSchemaStr(d, "username")
	dto.Password = PtrSchemaStr(d, "password")

	dto.ConnectionUrl = PtrSchemaStr(d, "connectionurl")
	dto.CredentialsQueryString = PtrSchemaStr(d, "sql_credentials")
	dto.Description = PtrSchemaStr(d, "description")
	dto.DriverName = PtrSchemaStr(d, "drivername")
	dto.IdleConnectionTestPeriod = PtrSchemaInt32(d, "idleconnectiontestperiod")
	dto.InitialPoolSize = PtrSchemaInt32(d, "initialpoolsize")
	dto.MaxIdleTime = PtrSchemaInt32(d, "maxidletime")
	dto.MaxPoolSize = PtrSchemaInt32(d, "maxpoolsize")
	dto.MinPoolSize = PtrSchemaInt32(d, "minpoolsize")

	dto.AcquireIncrement = PtrSchemaInt32(d, "acquireincrement")
	dto.PooledDatasource = PtrSchemaBool(d, "pooleddatasource")
	dto.RelayCredentialQueryString = PtrSchemaStr(d, "sql_relay_credential")
	dto.ResetCredentialDml = PtrSchemaStr(d, "dml_reset_credential")
	dto.RolesQueryString = PtrSchemaStr(d, "sql_groups")
	dto.UseColumnNamesAsPropertyNames = PtrSchemaBool(d, "use_column_name_as_property_name")
	dto.UserPropertiesQueryString = PtrSchemaStr(d, "sql_user_attrs")
	dto.UserQueryString = PtrSchemaStr(d, "sql_username")

	return *dto, err
}

func buildDbIdSourceResource(d *schema.ResourceData, dto api.DbIdentitySourceDTO) error {
	d.SetId(sdk.StrDeref(dto.Name))
	_ = d.Set("acquireincrement", dto.GetAcquireIncrement())
	_ = d.Set("username", dto.GetAdmin())
	_ = d.Set("connectionurl", dto.GetConnectionUrl())
	_ = d.Set("sql_credentials", dto.GetCredentialsQueryString())
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
	_ = d.Set("sql_relay_credential", dto.GetRelayCredentialQueryString())
	_ = d.Set("dml_reset_credential", dto.GetResetCredentialDml())
	_ = d.Set("sql_groups", dto.GetRolesQueryString())
	_ = d.Set("use_column_name_as_property_name", dto.GetUseColumnNamesAsPropertyNames())
	_ = d.Set("sql_user_attrs", dto.GetUserPropertiesQueryString())
	_ = d.Set("sql_username", dto.GetUserQueryString())

	return nil
}
