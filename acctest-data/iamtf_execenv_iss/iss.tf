resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}


resource "iamtf_execenv_iss" "test" {
    name                              = "iss-replace_with_uuid"
    description                       = "Iss Iss-Exect-Env"
    architecture                      = "32"
    activation_remote_target          = "http://remote-josso:8081"
    ida                               = iamtf_identity_appliance.test.name
    isapi_extension_path              = "/isapi/sso.a"      
}