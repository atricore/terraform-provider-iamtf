provider "iamtf" {
  org_name      = "atricore"
  endpoint      = "http://localhost:8081/atricore-rest/services"
  client_id     = "idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
  client_secret = "7oUHlv(HLT%vxK4L"
}

resource "iamtf_identity_appliance" "devsso-1" {
  name        = "devsso-1"
  namespace   = "com.atricore.devsso"
  description = "Appliance #1"
  location    = "https://devsso.atricore.com"
}
