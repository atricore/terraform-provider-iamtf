
resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_google" "test" {
  name                 = "idp-google-replace_with_uuid"
  description          = "My Google IdP updated"
  client_id            = "my-client"
  client_secret        = "client-secret"
  authz_token_service  = "http://accounts.google.com/o/oauth2/auth"
  access_token_service = "http://accounts.google.com/o/oauth2/token"
  google_apps_domain   = "apps domain"
  ida                  = iamtf_identity_appliance.test.name

}
