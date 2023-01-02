provider "iamtf" {
  org_name      = "atricore"
  endpoint      = "http://localhost:8081/atricore-rest/services"
  client_id     = "idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
  client_secret = "7oUHlv(HLT%vxK4L"
}

resource "iamtf_identity_appliance" "ida-1" {
  name        = "ida-1"
  namespace   = "com.atricore.idbus.testacc.ida01"
  description = "Appliance #1"
  location    = "http://localhost:8081"
}

resource "iamtf_idvault" "sso-users" {
  ida  = iamtf_identity_appliance.ida-1.name
  name = "sso-users"
}

resource "iamtf_idp" "idp" {
  ida  = iamtf_identity_appliance.ida-1.name
  name = "idp"

  keystore {
    resource = filebase64("./idp.p12")
    password = "changeme"
  }

  authn_bind_ldap {
    priority          = 0
    provider_url      = "ldap://localhost:389"
    username          = "cn=admin,dc=mycompany,dc=com"
    password          = "chageme"
    authentication    = "strong"
    password_policy   = "none"
    perform_dn_search = false
    users_ctx_dn      = "ou=People,dc=mycompany,dc=com"
    userid_attr       = "uid"

    saml_authn_ctx    = "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport"
    referrals         = "ignore"
    operational_attrs = true

  }

  id_sources = [iamtf_idvault.sso-users.name]
  depends_on = [
    iamtf_idvault.sso-users
  ]

}

resource "iamtf_execenv_tomcat" "tc85" {
  ida         = iamtf_identity_appliance.ida-1.name
  name        = "tc85"
  description = "Tomcat 8.5"
  version     = "8.5"
  depends_on  = [iamtf_idp.idp]
}
resource "iamtf_idsource_db" "test" {

  name = "dbid-replace_with_uuid"
  ida  = iamtf_identity_appliance.ida-1.name

   connectionurl = "jdbc:mysql:localhost/%s?create=true"
  jdbc_driver    = "org.mysql.driver"

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
  sql_username         = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN = ?"
  sql_user_attrs       = "LASTNAME"
  sql_credentials      = "SELECT USERNAME FROM JOSSO_USER WHERE LOGIN = ?"
  sql_relay_credential = "n/a"
  sql_groups           = "SELECT R.ROLE FROM JOSSO_ROLE R"
  dml_reset_credential = ""

  use_column_name_as_property_name = true


  extension{
    fqcn = "fqcn test"
    extension_type="SERVICE"
    osgi_filter="osgi_filter test"
  }
}

resource "iamtf_app_agent" "partnerapp1" {
  ida          = iamtf_identity_appliance.ida-1.name
  name         = "partnerapp1"
  app_location = "http://localhost:8080/partnerapp-1"

  keystore {
    resource = filebase64("./sp.p12")
    password = "changeme"
    key_password = "secret"
  }

  idp {
    name         = iamtf_idp.idp.name
    is_preferred = true
  }

  exec_env = iamtf_execenv_tomcat.tc85.name

  depends_on = [
    iamtf_idp.idp, iamtf_execenv_tomcat.tc85
  ]

}
resource "iamtf_idsource_ldap" "test" {
    ida             = iamtf_identity_appliance.ida-1.name
    name            = "idvault-replace_with_uuid"
    provider_url    = "ldap://localhost:10389"
    username        = "uid=admin,ou=system"
    password        = "secret"
    users_ctx_dn    = "dc=example,dc=com,ou=IAM,ou=People"
    userid_attr     = "uid"
    groups_ctx_dn   = "dc=example,dc=com,ou=IAM,ou=Groups"
    groupid_attr    = "cn"
    groupmember_attr = "uniquemember"
    
    user_attributes {
        attribute = "cn"
        claim = "first_name"
    }

    user_attributes {
        attribute = "sn"
        claim = "last_name"
    }
    extension{
        fqcn = "fqcn test"
        osgi_filter="osgi_filter test"
        type="SERVICE"
        property {
            name    = "name test"
            value = "value test"
    }
  }  
}