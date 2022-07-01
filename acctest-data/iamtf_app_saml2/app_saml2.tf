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

  id_sources = [iamtf_idvault.test1.name]
  depends_on = [
    iamtf_idvault.test1
  ]
}

resource "iamtf_app_saml2" "test" {
  ida  = iamtf_identity_appliance.test.name // Required, no default
  name = "sp-replace_with_uuid"            // Required, no default
  description = "SP #replace_with_uuid"

  metadata = filebase64("../../acctest-data/iamtf_app_saml2/md.xml")

  depends_on = [
    iamtf_idp.test
  ]

}

