This is one of the key resources available in JOSSO.  It allows you to define an identity provider. JOSSO supports multiple identity protocols, you can connect any SP using: 

* OpenID Connect 1.0
* OAuth 2.0
* SAML 2.0
* JOSSO 1.0
* Custom

You can mix heterogeneouse service providers that have different requireminets and connect them with a single IDP. For example, you can share identity between SAML and OIDC applications transparently, providing users with a single authentication experience.

To configure an IDP you must provide an **identity source**, an **authentication mechanism** and a **keystore**.  Then you need to define each **service provider**, and reference the trusted IDP.

In JOSSO you can have multiple identity providers running in a single identity appliance.

Let's take a look at the following example:

```
resource "iamtf_idp" "idp" {
  
  ida  = iamtf_identity_appliance.ida-1.name
  
  name = "idp"

  keystore {
    resource = filebase64("./idp.p12")
    password = "changeme"
  }

  id_sources = [iamtf_idvault.sso-users.name]

  authn_basic {
    priority          = 0         
    pwd_hash          = "SHA-256"
    pwd_encoding      = "BASE64"
  }

  depends_on = [
    iamtf_idvault.sso-users
  ]

}
```

