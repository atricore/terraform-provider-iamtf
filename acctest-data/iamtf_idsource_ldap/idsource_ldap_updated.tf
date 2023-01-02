resource "iamtf_identity_appliance" "test" {
  name                = "testacc-replace_with_uuid"
  namespace           = "com.atricore.idbus.test.testacc.replace_with_uuid"
  description         = "Appliance #replace_with_uuid"
  location            = "http://localhost:8081"  
}

resource "iamtf_idsource_ldap" "test" {
    ida             = iamtf_identity_appliance.test.name
    name            = "idvault-replace_with_uuid"
    provider_url    = "ldaps://localhost:10636"
    username        = "uid=manager,ou=system"
    password        = "changeme"
    users_ctx_dn    = "dc=example1,dc=com,ou=IAM,ou=People"
    userid_attr     = "uid1"
    groups_ctx_dn   = "dc=example1,dc=com,ou=IAM,ou=Groups"
    groupid_attr    = "cn1"
    groupmember_attr = "uniquemember1"
    
    user_attributes {
        attribute = "cn"
        claim = "first_name"
    }
    
    user_attributes {
        attribute = "sn"
        claim = "last_name"
    }
  extension {
    fqcn           = "fqcn change"
    osgi_filter    = "osgi_filter change"
    type = "INSTANCE"
    property {
      name  = "name change"
      value = "value change"
    }
  }
}