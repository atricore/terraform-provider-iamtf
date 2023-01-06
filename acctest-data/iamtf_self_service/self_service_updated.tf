resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_self_service" "test" {
  description = "My idGoogle updated"
  element_id  = "element id updated"
  location    = "http://localhost:8081"
  name        = "idp-google-replace_with_uuid"
  secret      = "top secret"

  #  not soported by server
  #   service_connection {
  #     description = "description updated sc"
  #     element_id  = "element id updated sc"
  #     name        = "name updated sc"
  #   }
  ida = iamtf_identity_appliance.test.name
}
