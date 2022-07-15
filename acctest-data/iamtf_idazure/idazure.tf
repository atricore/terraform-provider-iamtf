resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idazure" "test" {
  name                 = "idAzure-replace_with_uuid"
  element_id           = ""
  description          = "My idAzure"
  location             = "http://localhost:8081"
  client_id            = "my-client"
  server_key           = "server-key for idpfacebook"
  authz_token_service  = "https://login.microsoft.com:443/<change-me>/oauth2/v2.0/token"
  access_token_service = "https://login.microsoft.com:443/<change-me>//oauth2/v2.0/authorize"

  ida = iamtf_identity_appliance.test.name
}
