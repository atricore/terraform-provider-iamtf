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
  name = "idp-replace_with_uuid"            // Required, no default

  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }

  attributes {
    profile                 = "CUSTOM"
    include_unmapped_claims = true
    map {
      name    = "first_name"
      mapping = "fname"
    }
    map {
      type    = "const"
      name    = "my-const"
      mapping = "my-const-value"
      format  = "BASIC"
    }
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

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

  // VERY IMPORTANT TO ADD ALL DEPENDENCIES
  depends_on = [
    iamtf_idvault.test1
  ]

}
