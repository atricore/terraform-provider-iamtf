#### Example 1

```
  authn_basic {
    priority          = 0         
    pwd_hash          = "SHA-256"
    pwd_encoding      = "BASE64"
  }
```

#### Example 2

```
  authn_basic {
    priority          = 0         
    pwd_hash          = "SHA-256"
    pwd_encoding      = "BASE64"
    crypt_salt_length = 0
    salt_prefix       = "sp1235"
    salt_suffix       = "sf5432"

    saml_authn_ctx = "urn:oasis:names:tc:SAML:2.0:ac:classes:Password"

  }
```