resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idsource_db" "test" {
    acquireincrement                  = 1
    admin                             = "usr-dbid"
    connectionurl                     = "jdbc:mysql:localhost/%s?create=true"
    credentialsquerystring            = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN = ?"
    description                       = "Description"
    drivername                        = "org.mysql.driver"
    idleconnectiontestperiod          = 1
    initialpoolsize                   = 10
    maxidletime                       = 15
    maxpoolsize                       = 20
    minpoolsize                       = 1
    name                              = "dbid-replace_with_uuid"
    password                          = "pdw-dbid"
    pooleddatasource                  = true
    relaycredentialquerystring        = "n/a"
    resetcredentialdml                = ""
    rolesquerystring                  = "SELECT R.ROLE FROM JOSSO_ROLE R"
    usecolumnnamesaspropertynames     = true
    userpropertiesquerystring         = "LASTNAME"
    userquerystring                   = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN = ?"
    ida                               = iamtf_identity_appliance.test.name
}