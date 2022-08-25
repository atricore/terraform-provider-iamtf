

resource "iamtf_identity_appliance" "devsso-1" {
  name        = "devsso-1"
  namespace   = "com.atricore.devsso"
  description = "Appliance #1"
  location    = "https://devsso.atricore.com"
}
