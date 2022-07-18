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
    crypt_salt_lenght = 0         // Optional, default 0, valid values: multiples of 8 up to 256
    salt_prefix       = "sp"      // Optional, no default
    salt_suffix       = "sf"      // Optional, no default

    saml_authn_ctx = "urn:oasis:names:tc:SAML:2.0:ac:classes:Password" // Optional, default urn:oasis:names:tc:SAML:2.0:ac:classes:Password, valid values: urn:oasis:names:tc:SAML:2.0:ac:classes:Password, urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport
  }

  depends_on = [
    iamtf_idvault.test1
  ]
}


resource "iamtf_app_sharepoint" "test" {
  app_slo_location            = "http://myapp/slo.orig"
  description                 = "sharePoint description"
  name                        = "app-sharePoint-replace_with_uuid"
  sts_signing_cert_subject    = "sharePoint stsLocationCertSubject"
  sts_encrypting_cert_subject = "sharePoint stsEncryptingCertSubject"
  ida                         = iamtf_identity_appliance.test.name

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
  }

  depends_on = [
    iamtf_idp.test1
  ]

}
