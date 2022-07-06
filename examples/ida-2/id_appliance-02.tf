provider "iamtf" {
  org_name      = "atricore"
  endpoint      = "http://localhost:8081/atricore-rest/services"
  client_id     = "idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
  client_secret = "7oUHlv(HLT%vxK4L"
}

resource "iamtf_identity_appliance" "ida-2" {
  name        = "ida-2"
  namespace   = "com.atricore.idbus.testacc.ida02"
  description = "Appliance #2"
  location    = "http://localhost:8081/"
}

resource "iamtf_idvault" "sso-users" {
  ida  = iamtf_identity_appliance.ida-2.name
  name = "sso-users"
}

resource "iamtf_idp" "idp" {
  ida  = iamtf_identity_appliance.ida-2.name
  name = "idp"

  keystore {
    resource = filebase64("./idp.p12")
    password = "changeme"
  }

  id_sources = [iamtf_idvault.sso-users.name]
  depends_on = [
    iamtf_idvault.sso-users
  ]

}

resource "iamtf_app_oidc" "partnerapp2" {
  ida = iamtf_identity_appliance.ida-2.name

  name        = "partnerapp2"
  description = "OIDC App example"
  client_id   = "1234-5678-9012"
  client_secret = "changeme"
  client_authn = "CLIENT_SECRET_BASIC"
  redirect_uris = ["http://localhost:8080/oidc/"]

  idps = [
    iamtf_idp.idp.name
  ]

}
