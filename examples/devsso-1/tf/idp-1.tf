resource "iamtf_idp" "idp-1" {
  ida  = iamtf_identity_appliance.devsso-1.name
  name = "idp-1"

  keystore {
    resource = filebase64("./saml.p12")
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

  id_sources = [iamtf_idsource_ldap.ad-users.name]
  
  depends_on = [
    iamtf_idsource_ldap.ad-users
  ]

}