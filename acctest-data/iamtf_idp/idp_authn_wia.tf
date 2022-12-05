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

  // Relative to the test running folder!
  keystore {
    resource = filebase64("../../acctest-data/idp.p12")
    password = "changeme"
  }

  authn_wia {
    priority                 = 0
    domain                   = "MYCOMPANY.COM"
    domain_controller        = "DC0.MYCOMPANY.COM"
    host                     = "sso.mycompany.com"
    port                     = 8081
    protocol                 = "https"
    service_name             = "jossosvc"
    service_class            = "tod"
    overwrite_kerberos_setup = true
    keytab = filebase64("../../acctest-data/josso2.keytab")

  }

  id_sources = [iamtf_idvault.test1.name] // Required, no default min 1, max unbounded

}
