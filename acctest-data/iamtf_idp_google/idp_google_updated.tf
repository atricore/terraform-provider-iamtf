resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_google" "test" {
  name                 = "idp-google-replace_with_uuid"
  description          = "My idGoogle updated"
  client_id            = "my-client updated"
  client_secret        = "client-secret-updated"
  authz_token_service  = "http://accounts.google.com/o/oauth2/auth1"
  access_token_service = "http://accounts.google.com/o/oauth2/token1"
  google_apps_domain   = "apps domain1"
  ida                  = iamtf_identity_appliance.test.name
}
