resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}


resource "iamtf_execenv_weblogic" "test" {
  active                   = true
  description              = "wl description updated "
  display_name             = "wl displayName updated"
  domain                   = "wl domain updated"
  install_demo_apps        = true
  install_uri              = "wl installUri updated"
  name                     = "wl-replace_with_uuid updated"
  overwrite_original_setup = true
  version                  = "12"
  target_jdk               = "Remote"
  type                     = "wl type updated"
  ida                      = iamtf_identity_appliance.test.name
}
