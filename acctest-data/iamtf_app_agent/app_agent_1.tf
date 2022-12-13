resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idvault" "test1" {
  ida  = iamtf_identity_appliance.test.name
  name = "idvault-replace_with_uuid"
}

resource "iamtf_idp" "test1" {
  ida        = iamtf_identity_appliance.test.name
  name       = "idp-1-replace_with_uuid"
  id_sources = [iamtf_idvault.test1.name]

  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }


  authn_basic {
    priority          = 0         // Required, default 0 (should be unique)
    pwd_hash          = "SHA-256" // Optional, default SHA-256, valid values: NONE, CRYPT, BCRYPT, SHA-512, SHA-256, SHA-1, MD5
    pwd_encoding      = "BASE64"  // Otional, default BASE64, valid values: NONE, HEX, BASE64
    crypt_salt_length = 0         // Optional, default 0, valid values: multiples of 8 up to 256
    salt_prefix       = "sp"      // Optional, no default
    salt_suffix       = "sf"      // Optional, no default

    saml_authn_ctx = "urn:oasis:names:tc:SAML:2.0:ac:classes:Password" // Optional, default urn:oasis:names:tc:SAML:2.0:ac:classes:Password, valid values: urn:oasis:names:tc:SAML:2.0:ac:classes:Password, urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport
  }
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
    resource = filebase64("../../acctest-data/sp.p12")
    password = "changeme"
  }

  saml2 {
    message_ttl           = 400
    message_ttl_tolerance = 410

    sign_authentication_requests = true
    //sign_requests                 = true
    signature_hash        = "SHA-256"
    want_assertion_signed = true
    //want_slo_response_signed      = false

    // Use validation function to restrict possible values  
    account_linkage = "ONE_TO_ONE" // EMAIL, UID, ONE_TO_ONE, CUSTOM, Optional, Default = ONE_TO_ONE

    // Use validation function to restrict possible values
    identity_mapping = "REMOTE" // LOCAL, REMOTE, MERGED, CUSTOM, Optinal, Default REMOTE
  }

  idp {
    name         = iamtf_idp.test1.name
    is_preferred = true

    saml2 {
      signature_hash        = "SHA-512"
      message_ttl           = 401
      want_assertion_signed = false
      account_linkage       = "EMAIL"
    }

  }
}

