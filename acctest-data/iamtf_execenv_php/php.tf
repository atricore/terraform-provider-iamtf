resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}


resource "iamtf_execenv_php" "test" {
    name                              = "php-replace_with_uuid"
    description                       = "Php Php-Exect-Env"
    activation_remote_target          = "http://remote-josso:8081"
    activation_install_samples        = true
    activation_path                   = "/opt/atricore/josso-ee-2/Php-Exect-Env"
    activation_override_setup         = true
    ida                               = iamtf_identity_appliance.test.name
    environment                       = "STANDARD"
}