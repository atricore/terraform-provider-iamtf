resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_self_service" "test" {
  description = "My self-services created"
  name        = "self-svc-rs-replace_with_uuid"
  ida = iamtf_identity_appliance.test.name
}
