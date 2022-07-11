resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}


resource "iamtf_app_sharepoint" "test" {
  slo_location                = "sharePoint sloLocation"
  description                 = "sharePoint description"
  name                        = "app-sharePoint-replace_with_uuid"
  sts_signing_cert_subject    = "sharePoint stsLocationCertSubject"
  slo_location_enabled        = false
  sts_encrypting_cert_subject = "sharePoint stsEncryptingCertSubject"
  ida                         = iamtf_identity_appliance.test.name

}
