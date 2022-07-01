resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}

resource "iamtf_idvault" "test" {
    ida             = iamtf_identity_appliance.test.name
    name            = "idvault-replace_with_uuid"
    description     = "My IdVault (upd)"
    connector       = "connector-2" 
}