resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_google" "test" {
  name                 = "idGoogle-updated-replace_with_uuid"
  element_id           = ""
  description          = "My idGoogle updated"
  location             = "http://localhost:8082"
  client_id            = "my-client updated"
  server_key           = "server-key for idpfacebook updated"
  authz_token_service  = "http://accounts.google.com:443/o/oauth2/auth"
  access_token_service = "http://accounts.google.com:443/o/oauth2/token"
  google_apps_domain   = "apps domain"
  ida                  = iamtf_identity_appliance.test.name
}
