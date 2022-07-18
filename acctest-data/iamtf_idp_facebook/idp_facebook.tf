resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_facebook" "test" {
  name                 = "idp_fb-replace_with_uuid"
  description          = "My Facebook IdP"
  client_id            = "my-client"
  client_secret        = "secret for idpfacebook"
  authz_token_service  = "https://www.facebook.com/dialog/oauth"
  access_token_service = "https://graph.facebook.com/oauth/access_token"
  scopes               = ["user_likes", "user_photos"]
  user_fields          = ["about", "age_range"]
  ida                  = iamtf_identity_appliance.test.name
}
