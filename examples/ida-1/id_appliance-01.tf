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
