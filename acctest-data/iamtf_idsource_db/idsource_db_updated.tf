resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idsource_db" "test" {

    name                              = "dbid-replace_with_uuid"
    ida                               = iamtf_identity_appliance.test.name

      connectionurl = "jdbc:mysql:localhost/%s?create=true"
  drivername    = "org.mysql.driver"

  description = "SSO Users (Mysql DB)"
  username    = "usr-dbid"
  password    = "pdw-dbid"

  # DB pool
  pooleddatasource         = true
  acquireincrement         = 1
  idleconnectiontestperiod = 1
  initialpoolsize          = 10
  maxidletime              = 15
  maxpoolsize              = 20
  minpoolsize              = 1


  # SQL to retrieve user information
  sql_user             = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN1 = ?"
  sql_user_attrs       = "LASTNAME1"
  sql_credentials      = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN1 = ?"
  sql_relay_credential = "n/a1"
  sql_groups           = "SELECT R.ROLE FROM JOSSO_ROLE1 R"
  dml_reset_credential = "1"

  use_column_name_as_property_name = true
}