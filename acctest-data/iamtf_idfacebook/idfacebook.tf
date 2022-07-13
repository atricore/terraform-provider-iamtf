resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_idfacebook" "test" {
  name        = "idFacebook-replace_with_uuid"
  element_id  = ""
  description = "My IdFacebook"
  location    = "http://localhost:8081"
  client_id   = "my-client"
  server_key  = "server-key for idpfacebook"
  //Authorization endpoint
  //Token endpoint
  //Permissions
  user_fields = "user-fields for idpfacebook"
  ida         = iamtf_identity_appliance.test.name
}
