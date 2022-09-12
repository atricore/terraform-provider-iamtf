
resource "iamtf_idsource_ldap" "ad-users" {
    ida             = iamtf_identity_appliance.devsso-1.name
    name            = "ad-users"
    provider_url    = "ldap://openldap:1389"
    username        = "cn=admin,dc=devsso1,dc=atricore,dc=com"
    password        = "secret"

    // Groups
    groups_ctx_dn   = "ou=groups,dc=devsso1,dc=atricore,dc=com"
    groupid_attr    = "cn"
    groupmember_attr = "member"
    
    // Users
    users_ctx_dn    = "ou=users,dc=devsso1,dc=atricore,dc=com"
    userid_attr     = "uid"
    
    // User attributes
    user_attributes {
        attribute = "cn"
        claim = "first_name"
    }

    user_attributes {
        attribute = "sn"
        claim = "last_name"
    }
    
}