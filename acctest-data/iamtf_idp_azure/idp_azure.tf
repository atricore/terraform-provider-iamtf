resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_azure" "test" {
  name                 = "idp-az-replace_with_uuid"
  description          = "My Azure IdP"
  client_id            = "my-client"
  client_secret        = "server-key for azure"
  server_key           = "my-server-key"
  authz_token_service  = "http://login.microsoft.com/<change-me>/oauth2/v2.0/token"
  access_token_service = "http://login.microsoft.com/<change-me>/oauth2/v2.0/authorize"
  ida                  = iamtf_identity_appliance.test.name

}
