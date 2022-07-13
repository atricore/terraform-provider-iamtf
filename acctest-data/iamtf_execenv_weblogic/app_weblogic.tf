resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}


resource "iamtf_execenv_weblogic" "test" {
  active                   = false
  description              = "wl description "
  display_name             = "wl displayName"
  domain                   = "wl domain"
  install_demo_apps        = false
  install_uri              = "wl installUri"
  name                     = "wl-replace_with_uuid"
  overwrite_original_setup = false
  version                  = "14"
  target_jdk               = "Local"
  type                     = "wl type"
  ida                      = iamtf_identity_appliance.test.name
}
