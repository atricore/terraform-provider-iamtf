resource "iamtf_execenv_weblogic" "wl" {
  ida         = iamtf_identity_appliance.devsso-1.name
  name        = "wl"
  description = "Weblogic 12"
  version     = "12"
  domain      = "/tmp"

  depends_on = [iamtf_idp.idp-1]
}

resource "iamtf_app_agent" "partnerapp1" {
  ida          = iamtf_identity_appliance.devsso-1.name
  name         = "partnerapp1"
  app_location = "http://localhost:8080/partnerapp-1"

  keystore {
    resource     = filebase64("./saml.p12")
    password     = "changeme"
    key_password = "secret"
  }

  idp {
    name         = iamtf_idp.idp-1.name
    is_preferred = true
  }

  exec_env = iamtf_execenv_weblogic.wl.name

  depends_on = [
    iamtf_idp.idp-1, iamtf_execenv_weblogic.wl
  ]

}
