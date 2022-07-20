resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}


resource "iamtf_execenv_weblogic" "test" {
  name        = "wl-replace_with_uuid"
  description = "wl description"
  version     = "14"
  domain      = "wl_domain"
  target_jdk  = "jdk16"

  activation_remote_target   = "http://remote-josso:8081"
  activation_install_samples = true
  activation_path            = "/opt/Oracle/wl_server"
  activation_override_setup  = true

  ida                        = iamtf_identity_appliance.test.name
}
