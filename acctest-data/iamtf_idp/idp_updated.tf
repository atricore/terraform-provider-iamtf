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

  branding      = "josso2-default-branding" // Optional, default josso25-branding
  dashboard_url = "1 dashboard_url"         // Optional, no default
  description   = "1 description"           // Optional, no default

  error_binding = "ARTIFACT" // Optional , default ARTIFACT

  session_timeout          = 31   // Optional, defualt 30
  max_sessions_per_user    = 303  // Optional, default -1
  destroy_previous_session = true // Optional, default false

  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }


  saml2 { // Min 0, Max 1

    want_authn_req_signed = false // Optional, default false
    want_req_signed       = false // Optional, default false
    sign_reqs             = true  // Optional, default true
    message_ttl           = 302   // Optional, computed (server will provide)
    message_ttl_tolerance = 303   // Optional , computed (server will provide)

    signature_hash    = "SHA256" // Optional, default SHA256, valid vlaues SHA1, SHA256, SHA384, SHA512
    encrypt_algorithm = "NONE"   // Optional, default NONE, valid values AES-128, AES-256, AES-3DES

  }

  oidc {                    // Min 0, Max 1
    access_token_ttl = 8082 // Optional, default 3600
    id_token_ttl     = 8082 // Optional, default 3600
    authz_code_ttl   = 8083 // Optional, default 300
    enabled          = true // Reduired, no default
  }

  oauth2 { // Min 0, Max 1

    enabled                   = true            // Required, no default
    shared_key                = "my secret key" // Required, no default
    rememberme_token_validity = 8080            // Optional, default 43200
    token_validity            = 8081            // Optional, default 300


    pwdless_authn_enabled  = true                       // Optinal , default false
    pwdless_authn_subject  = "1 pwdless_authn_subject"  // Optional, no default
    pwdless_authn_template = "1 pwdless_authn_template" // Optional, no default
    pwdless_authn_to       = "1 pwdless_authn_to"       // Optional, no default (todo, reqiured when enabled ?)
    pwdless_authn_from     = "1 pwdless_authn_to"       // Optional, no default (todo, reqiured when enabled ?)
  }

  authn_basic {
    priority          = 0         // Required, default 0 (should be unique)
    pwd_hash          = "SHA-512" // Optional, default SHA-256, valid values: NONE, CRYPT, BCRYPT, SHA-512, SHA-256, SHA-1, MD5
    pwd_encoding      = "HEX"     // Otional, default BASE64, valid values: NONE, HEX, BASE64
    crypt_salt_lenght = 8         // Optional, default 0, valid values: multiples of 8 up to 256
    salt_prefix       = "sp1"     // Optional, no default
    salt_suffix       = "sf1"     // Optional, no default

    saml_authn_ctx = "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport" // Optional, default urn:oasis:names:tc:SAML:2.0:ac:classes:Password, valid values: urn:oasis:names:tc:SAML:2.0:ac:classes:Password, urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport

  }

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

  // VERY IMPORTANT TO ADD ALL DEPENDENCIES
  depends_on = [
    iamtf_idvault.test1
  ]

}
