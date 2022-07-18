resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_facebook" "test" {
  name                 = "idp_fb-replace_with_uuid"
  description          = "My Facebook IdP #updated"
  client_id            = "my-client-1"
  client_secret        = "secret for idpfacebook updated"
  authz_token_service  = "http://www.facebook.com/dialog/oauth1"
  access_token_service = "http://graph.facebook.com/oauth/access_token1"
  scopes               = ["user_likes"]
  user_fields          = ["about", "gender"]
  ida                  = iamtf_identity_appliance.test.name
}
