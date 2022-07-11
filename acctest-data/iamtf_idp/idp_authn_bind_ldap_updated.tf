resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idvault" "test1" {
  ida  = iamtf_identity_appliance.test.name
  name = "idvault1-replace_with_uuid"
}

resource "iamtf_idp" "test" {

  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "idp-replace_with_uuid"            // Required, no default

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
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

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

  // VERY IMPORTANT TO ADD ALL DEPENDENCIES
  depends_on = [
    iamtf_idvault.test1
  ]

}
