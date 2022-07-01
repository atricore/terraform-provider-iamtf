resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idvault" "test1" {
  ida  = iamtf_identity_appliance.test.name
  name = "idvault1-replace_with_uuid"
}

resource "iamtf_idp" "test1" {
  ida  = iamtf_identity_appliance.test.name
  name = "idp-1-replace_with_uuid"
  
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }
  
  authn_basic {
    pwd_hash     = "SHA-256"
    pwd_encoding = "BASE64"
  }

  id_sources = [iamtf_idvault.test1.name]

  depends_on = [
    iamtf_idvault.test1
  ]
}

resource "iamtf_app_oidc" "test" {
  ida                       = iamtf_identity_appliance.test.name
  name                      = "app-oidc-replace_with_uuid"
  client_id                 = "my-client"
  client_secret             = "my-secret"
  client_authn              = "CLIENT_SECRET_BASIC"
  redirect_uris             = ["http://localhost:8080/partnerapp", "http://localhost:8080/partnerapp/1"]
  grant_types               = ["AUTHORIZATION_CODE"]
  response_types            = ["CODE"]
  encryption_alg            = "NONE"
  encryption_method         = "NONE"
  idtoken_encryption_alg    = "NONE"
  idtoken_encryption_method = "NONE"
  idtoken_signature_alg     = "HS256"
  signature_alg             = "HS256"

  idps = [
    iamtf_idp.test1.name
  ]

  depends_on = [
    iamtf_idp.test1
  ]

}
