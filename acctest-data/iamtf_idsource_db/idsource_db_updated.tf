resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idsource_db" "test" {
    acquireincrement                  = 2
    admin                             = "usr-dbid"
    connectionurl                     = "jdbc:mysql:localhost/%s?create=true"
    credentialsquerystring            = "SELECT PASSWORD FROM JOSSO_USER WHERE LOGIN = ?"
    description                       = "Descrip"
    drivername                        = "org.mysql.driver"
    idleconnectiontestperiod          = 2
    initialpoolsize                   = 11
    maxidletime                       = 16
    maxpoolsize                       = 21
    minpoolsize                       = 2
    name                              = "dbid-replace_with_uuid"
    password                          = "pdw-dbid"
    pooleddatasource                  = false
    relaycredentialquerystring        = "n/a"
    resetcredentialdml                = ""
    rolesquerystring                  = "SELECT R.ROLE FROM JOSSO_USER U"
    usecolumnnamesaspropertynames     = false
    userpropertiesquerystring         = "EMAIL"
    userquerystring                   = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN = ?"
    ida                               = iamtf_identity_appliance.test.name
}