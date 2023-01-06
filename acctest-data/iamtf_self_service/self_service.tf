resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_self_service" "test" {
  description = "My idGoogle"
  element_id  = "element id"
  location    = "http://localhost:8081"
  name        = "idp-google-replace_with_uuid"
  secret      = "secret"

  #  not soported by server
  #   service_connection {
  #     description = "description sc"
  #     element_id  = "element id sc"
  #     name        = "name sc"
  #   }
  ida = iamtf_identity_appliance.test.name
}
