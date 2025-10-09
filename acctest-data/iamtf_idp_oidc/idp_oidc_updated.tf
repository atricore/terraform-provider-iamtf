resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_oidc" "test" {
  name                 = "idp-oidc-replace_with_uuid"
  description          = "My Generic OIDC IdP Updated"
  client_id            = "my-oidc-client-updated"
  client_secret        = "my-oidc-secret-updated"
  server_key           = "server-key-value"
  issuer               = "https://provider.example.com"
  load_metadata        = false
  # When load_metadata is false, endpoints must be explicit
  authz_token_service  = "https://provider.example.com/oauth2/v2/authorize"
  access_token_service = "https://provider.example.com/oauth2/v2/token"
  scopes               = ["openid", "profile", "email", "address"]
  user_fields          = ["sub", "name", "email", "picture"]
  ida                  = iamtf_identity_appliance.test.name
}
