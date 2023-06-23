resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}


resource "iamtf_execenv_iis" "test" {
    name                              = "iis-replace_with_uuid"
    description                       = "IIS IIS-Exect-Env-updated"
    architecture                      = "64"
    activation_remote_target          = "http://remote-josso:8081"
    ida                               = iamtf_identity_appliance.test.name
    isapi_extension_path              = "/josso/agent.sso"
}