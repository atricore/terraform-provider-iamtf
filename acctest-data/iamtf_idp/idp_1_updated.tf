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

resource "iamtf_idp" "test" {

  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "idp_1-replace_with_uuid"            // Required, no default

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }

  saml2 {                         // Min 0, Max 1
    want_authn_req_signed = false // Optional, default false
    want_req_signed       = false // Optional, default false
    message_ttl           = 301   // Optional, computed (server will provide)
    message_ttl_tolerance = 302   // Optional , computed (server will provide)

    signature_hash    = "SHA256" // Optional, default SHA256, valid vlaues SHA1, SHA256, SHA384, SHA512
    encrypt_algorithm = "NONE"   // Optional, default NONE, valid values AES-128, AES-256, AES-3DES
  }

  sp {
    name = "idp-replace_with_uuid"
    saml2 {
      message_ttl           = 303
      want_authn_req_signed = false
      encrypt_algorithm = "AES-128" 
    }
  }

  authn_basic {
    priority          = 0         // Required, default 0 (should be unique)
    pwd_hash          = "SHA-256" // Optional, default SHA-256, valid values: NONE, CRYPT, BCRYPT, SHA-512, SHA-256, SHA-1, MD5
    pwd_encoding      = "BASE64"  // Otional, default BASE64, valid values: NONE, HEX, BASE64
  }

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

  // VERY IMPORTANT TO ADD ALL DEPENDENCIES
  depends_on = [
    iamtf_idvault.test1
  ]

}

resource "iamtf_execenv_tomcat" "tc85" {
  ida  = iamtf_identity_appliance.test.name // Required, no default
  name        = "tc85"
  description = "Tomcat 8.5"
  version     = "8.5"
  depends_on  = [iamtf_idp.test]
}

resource "iamtf_app_agent" "partnerapp1" {
  ida  = iamtf_identity_appliance.test.name // Required, no default
  name         = "partnerapp-replace_with_uuid"
  app_location = "http://localhost:8080/partnerapp"

  keystore {
    resource = filebase64("../../acctest-data/sp.p12")
    password = "changeme"
  }

  idp {
    name         = iamtf_idp.test.name
    is_preferred = true
  }

  exec_env = iamtf_execenv_tomcat.tc85.name

  depends_on = [
    iamtf_idp.test, iamtf_execenv_tomcat.tc85
  ]

}

