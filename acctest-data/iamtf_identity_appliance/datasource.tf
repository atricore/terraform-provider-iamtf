resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  location            = "http://localhost:8081"  
  description         = "Appliance #replace_with_uuid"
}

data "iamtf_identity_appliance" "test" {
  name        = iamtf_identity_appliance.test.name
  namespace   = iamtf_identity_appliance.test.namespace
  description = iamtf_identity_appliance.test.description
  location    = iamtf_identity_appliance.test.location
}
