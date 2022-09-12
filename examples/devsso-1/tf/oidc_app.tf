
resource "iamtf_app_oidc" "app-oidc" {
  ida         = iamtf_identity_appliance.devsso-1.name
  name                      = "app-oidc"
  client_id                 = "my-client"
  client_secret             = "my-secret"
  client_authn              = "CLIENT_SECRET_BASIC"
  redirect_uris             = ["https://dev.atricore.com/oidc-client"]
  grant_types               = ["AUTHORIZATION_CODE"]
  response_types            = ["CODE"]
  response_modes            = ["QUERY"]
  signature_alg             = "RS256"

  idps = [
    iamtf_idp.idp-1.name
  ]

  depends_on = [
    iamtf_idp.idp-1
  ]

}