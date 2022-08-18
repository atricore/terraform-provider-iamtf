
resource "iamtf_idsource_ldap" "ad-users" {
    ida             = iamtf_identity_appliance.devsso-1.name
    name            = "ad-users"
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