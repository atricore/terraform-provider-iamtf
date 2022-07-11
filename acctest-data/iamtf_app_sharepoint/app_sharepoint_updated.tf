resource "iamtf_identity_appliance" "test" {
  name        = "testacc-replace_with_uuid"
  namespace   = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description = "Appliance #replace_with_uuid"
  location    = "http://localhost:8081"
}

resource "iamtf_app_sharepoint" "test" {
  slo_location                = "sharePoint sloLocation updated"
  description                 = "sharePoint description updated"
  name                        = "app-sharePoint-replace_with_uuid updated"
  sts_signing_cert_subject    = "sharePoint stsLocationCertSubject updated"
  slo_location_enabled        = true
  sts_encrypting_cert_subject = "sharePoint stsEncryptingCertSubject updated"
  ida                         = iamtf_identity_appliance.test.name

}
