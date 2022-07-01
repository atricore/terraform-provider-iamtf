resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}

resource "iamtf_idsource_ldap" "test" {
    ida             = iamtf_identity_appliance.test.name
    name            = "idvault-replace_with_uuid"
    provider_url    = "ldap://localhost:10389"
    username        = "uid=admin,ou=system"
    password        = "secret"
    users_ctx_dn    = "dc=example,dc=com,ou=IAM,ou=People"
    userid_attr     = "uid"
    groups_ctx_dn   = "dc=example,dc=com,ou=IAM,ou=Groups"
    groupid_attr    = "cn"
    groupmember_attr = "uniquemember"
    
    user_attributes {
        attribute = "cn"
        claim = "first_name"
    }

    user_attributes {
        attribute = "sn"
        claim = "last_name"
    }
    
}