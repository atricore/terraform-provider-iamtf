resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_oidc" "test" {
  name          = "idp-oidc-replace_with_uuid"
  description   = "My Generic OIDC IdP"
  client_id     = "my-oidc-client"
  client_secret = "my-oidc-secret"
  issuer        = "https://provider.example.com"
  load_metadata = true
  # Endpoints are optional when load_metadata is true
  scopes      = ["openid", "profile", "email"]
  user_fields = ["sub", "name", "email"]
  ida         = iamtf_identity_appliance.test.name
}
