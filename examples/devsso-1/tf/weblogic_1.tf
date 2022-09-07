// Weblogic server definicion
resource "iamtf_execenv_weblogic" "wl" {
  ida         = iamtf_identity_appliance.devsso-1.name
  name        = "wl"
  description = "Weblogic 12"
  version     = "12"
  domain      = "/base_domain"

  depends_on = [iamtf_idp.idp-1]
}

// Weblogic application
resource "iamtf_app_agent" "partnerapp1" {
  ida          = iamtf_identity_appliance.devsso-1.name
  name         = "partnerapp1"

  // Application base location
  app_location = "https://dev.atricore.com/partnerapp"

  // Trusted identity providers
  idp {
    name         = iamtf_idp.idp-1.name
    is_preferred = true
  }
  
  // Weblogic server
  exec_env = iamtf_execenv_weblogic.wl.name

  keystore {
    resource     = filebase64("./saml.p12")
    password     = "changeme"
    key_password = "secret"
  }
  
  depends_on = [
    iamtf_idp.idp-1, iamtf_execenv_weblogic.wl
  ]

}
