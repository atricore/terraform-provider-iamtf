resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idfacebook" "test" {
  name        = "idFacebook-updated-replace_with_uuid"
  element_id  = ""
  description = "My IdFacebook updated"
  location    = "http://localhost:8082"
  client_id   = "my-client updated"
  server_key  = "server-key for idpfacebook updated"
  //Authorization endpoint
  //Token endpoint
  //Permissions
  user_fields = "user-fields for idpfacebook updated"
  ida         = iamtf_identity_appliance.test.name
}
