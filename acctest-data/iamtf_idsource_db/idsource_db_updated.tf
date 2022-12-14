resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idsource_db" "test" {

  name = "dbid-replace_with_uuid"
  ida  = iamtf_identity_appliance.test.name

  connectionurl = "jdbc:mysql:localhost/%s?create=true"
  jdbc_driver   = "org.mysql.driver"

  description = "SSO Users (Mysql DB)"
  username    = "usr-dbid"
  password    = "pdw-dbid"

  # DB pool
  connection_pool             = true
  acquire_increment           = 1
  idle_connection_test_period = 1
  initial_pool_size           = 10
  max_idle_time               = 15
  max_pool_size               = 20
  min_pool_size               = 1


  # SQL to retrieve user information
  sql_username         = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN1 = ?"
  sql_user_attrs       = "LASTNAME1"
  sql_credentials      = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN1 = ?"
  sql_relay_credential = "n/a1"
  sql_groups           = "SELECT R.ROLE FROM JOSSO_ROLE1 R"
  dml_reset_credential = "1"

  use_column_name_as_property_name = true

  extension {
    fqcn           = "fqcn change"
    osgi_filter    = "osgi_filter change"
    extension_type = "INSTANCE"
    properties {
      name  = "name change"
      value = "value change"
    }
  }
}
