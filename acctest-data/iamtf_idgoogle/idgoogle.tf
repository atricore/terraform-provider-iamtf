
resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idgoogle" "test" {
  name                 = "idGoogle-replace_with_uuid"
  element_id           = ""
  description          = "My idGoogle"
  location             = "http://localhost:8081"
  client_id            = "my-client"
  server_key           = "server-key for idpfacebook"
  authz_token_service  = "https://accounts.google.com:443/o/oauth2/auth"
  access_token_service = "https://accounts.google.com:443/o/oauth2/token"
  google_apps_domain   = "Google Suite"
  ida                  = iamtf_identity_appliance.test.name
}
