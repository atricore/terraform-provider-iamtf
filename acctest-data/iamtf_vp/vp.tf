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

resource "iamtf_idp" "idp-test" {

  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "idp-test-replace_with_uuid"       // Required, no default

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

  // VERY IMPORTANT TO ADD ALL DEPENDENCIES
  depends_on = [
    iamtf_idvault.test1
  ]

}

resource "iamtf_vp" "vp-test" {
  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "idp-test-replace_with_uuid"       // Required, no default

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
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
