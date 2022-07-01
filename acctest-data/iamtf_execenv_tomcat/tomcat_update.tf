resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}


resource "iamtf_execenv_tomcat" "test" {
    name                              = "tc-replace_with_uuid"
    description                       = "Tomcat Tomcat-Exect-Env-updated"
    version                           = "9"
    activation_remote_target          = "http://remote-josso:8082"
    activation_install_samples        = false
    activation_path                   = "/opt/atricore/josso-ee-2/Tomcat-Exect-Env/updated"
    activation_override_setup         = false
    ida                               = iamtf_identity_appliance.test.name
      

}