provider "josso" {
  org_name = "atricore"
  endpoint = "http://localhost:8081/atricore-rest/services"
  client_id = "idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
  client_secret = "7oUHlv(HLT%vxK4L"
  username = "admin"
  password = "atricore"
}

resource "iamtf_identity_appliance" "ida2" {
  name                = "ida-2"
  namespace           = "com.atricore.idbus.test.ida2"
  description         = "Appliance #1"
  location            = "http://localhost:8081"  
}

resource "iamtf_idvault" "sso-users" {
    ida             = iamtf_identity_appliance.ida2.name
    name            = "sso-users"
}

resource "iamtf_idp" "idp1" {
    ida             = iamtf_identity_appliance.ida2.name
    name            = "idp-1"
    id_sources      = [iamtf_idvault.sso-users.name]

}


resource "iamtf_idp" "idp2" {
    ida             = iamtf_identity_appliance.ida2.name
    name            = "idp-2"
    id_sources      = [iamtf_idvault.sso-users.name]

}


resource "iamtf_idp" "idp3" {
    ida             = iamtf_identity_appliance.ida2.name
    name            = "idp-3"
    id_sources      = [iamtf_idvault.sso-users.name]

}


resource "iamtf_execenv_tomcat"  "tc85" {
    name                              = "tc85"
    description                       = "Tomcat 8.5"
    version                           = "tc85"
    plataformid                       = "tc85"
    ida                               = iamtf_identity_appliance.ida2.name
    depends_on                        = [iamtf_idp.idp1]
}

resource "iamtf_app_agent" "partnerapp1" {
    ida                             = iamtf_identity_appliance.ida2.name
    app_location                    = "http://localhost:8080/partnerapp"
    ignored_web_resources           = ["*.ico"]
    default_resource                = "/index.jsp"
    description                     = "Sample JEE partner app" 
    name                            = "app-1"
    enable_metadata_endpoint        = true
    error_binding                   = "JSON"
    message_ttl                     = 300
    message_ttl_tolerance           = 300
    dashboard_url                   = "data"
    
    //IdP channel
    sign_authentication_requests    = true
    sign_requests                   = true
    signature_hash                  = "SHA256"
    want_assertion_signed            = false
    want_slo_response_signed        = false
    preffered_idp                   = iamtf_idp.idp1.name

    idp { 
        name = iamtf_idp.idp2.name 
        sign_authentication_requests    = true
        sign_requests                   = true
        signature_hash                  = "SHA256"
        want_assertion_signed            = false
        want_slo_response_signed        = false
    }

    idp { 
        name = iamtf_idp.idp3.name 
    }


    exec_env = iamtf_execenv_tomcat.tc85.name

}


