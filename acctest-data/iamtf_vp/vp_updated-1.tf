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
    sign_reqs             = true  // Optional, default true
    message_ttl           = 301   // Optional, computed (server will provide)
    message_ttl_tolerance = 302   // Optional , computed (server will provide)

    signature_hash    = "SHA256" // Optional, default SHA256, valid vlaues SHA1, SHA256, SHA384, SHA512
    encrypt_algorithm = "NONE"   // Optional, default NONE, valid values AES128, AES256, AES3DES
  }
  authn_basic {
    priority          = 0         // Required, default 0 (should be unique)
    pwd_hash          = "SHA-256" // Optional, default SHA-256, valid values: NONE, CRYPT, BCRYPT, SHA-512, SHA-256, SHA-1, MD5
    pwd_encoding      = "BASE64"  // Otional, default BASE64, valid values: NONE, HEX, BASE64
  }

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

}

resource "iamtf_vp" "test" {

  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "vp-test-replace_with_uuid"        // Required, no default

  subject_id = "ATTRIBUTE"
  subject_id_attr = "mail"

  idp {
    name = iamtf_idp.test.name
  }

  //sp {
  //  name = "partnerapp-replace_with_uuid-sp"    
  //}

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }

  //id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

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
    name         = iamtf_vp.test.name
    is_preferred = true
  }

  exec_env = iamtf_execenv_tomcat.tc85.name

}


