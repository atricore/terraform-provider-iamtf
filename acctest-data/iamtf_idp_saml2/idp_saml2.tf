resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idp_saml2" "test1" {
  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "idp-replace_with_uuid"            // Required, no default
  metadata = filebase64("../../acctest-data/iamtf_app_saml2/md.xml")
}

resource "iamtf_execenv_tomcat" "test" {
  name                       = "tc-replace_with_uuid"
  description                = "Tomcat Tomcat-Exect-Env"
  version                    = "9"
  activation_remote_target   = "http://remote-josso:8081"
  activation_install_samples = true
  activation_path            = "/opt/atricore/josso-ee-2/Tomcat-Exect-Env"
  activation_override_setup  = true
  ida                        = iamtf_identity_appliance.test.name
  depends_on = [
    iamtf_idp.test1
  ]

}

resource "iamtf_app_agent" "test" {

  # Referenced resources MUST be provided as dependencies
  depends_on = [
    iamtf_idp.test1,
    iamtf_execenv_tomcat.test
  ]

  ida                      = iamtf_identity_appliance.test.name
  app_slo_location         = "http://myapp-replace_with_uuid:8080/partnerapp/slo"
  app_location             = "http://myapp-replace_with_uuid:8080/partnerapp"
  ignored_web_resources    = ["*.ico"]
  default_resource         = "http://myapp-replace_with_uuid:8080/partnerapp/home"
  description              = "desc app-a"
  name                     = "app-agent-replace_with_uuid"
  dashboard_url            = "http://myapp-replace_with_uuid:8080/partnerapp/dashboard"
  
  error_binding            = "JSON"

  exec_env = iamtf_execenv_tomcat.test.name

  keystore {
    resource          = filebase64("../../acctest-data/sp.p12")
    password = "changeme"
  }

  saml2 {
      message_ttl           = 400
      message_ttl_tolerance = 410

      sign_authentication_requests = true
      //sign_requests                 = true
      signature_hash       = "SHA-256"
      want_assertion_signed = true
      //want_slo_response_signed      = false

      // Use validation function to restrict possible values  
      account_linkage      = "ONE_TO_ONE" // EMAIL, UID, ONE_TO_ONE, CUSTOM, Optional, Default = ONE_TO_ONE

      // Use validation function to restrict possible values
      identity_mapping      = "REMOTE"        // LOCAL, REMOTE, MERGED, CUSTOM, Optinal, Default REMOTE

      bindings {
        http_post = true
        http_redirect = false
      }
  }

  idp {
    name = iamtf_idp_saml2.test1.name
    is_preferred = true
  }

}

