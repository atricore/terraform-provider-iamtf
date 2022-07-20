resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_azure" "test" {
  name                 = "idp-az-replace_with_uuid"
  description          = "My Azure IdP updated"
  client_id            = "my-client updated"
  client_secret        = "server-key for azure updated"
  server_key           = "my-server-key updated"
  authz_token_service  = "http://login.microsoft.com/<change-me>/oauth2/v2.0/token1"
  access_token_service = "http://login.microsoft.com/<change-me>/oauth2/v2.0/authorize1"
  ida                  = iamtf_identity_appliance.test.name
}
