---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iamtf_app_agent Resource - terraform-provider-iamtf"
subcategory: ""
description: |-
  
---

# iamtf_app_agent (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_location` (String) application location.  Base application URL, i.e. https://myapp.com
- `exec_env` (String) name of the execution environment resource
- `ida` (String) Identity appliance name
- `keystore` (Block List, Min: 1) Keystore configuration.  A single keystore containing the private key and certificate is supported. (see [below for nested schema](#nestedblock--keystore))
- `name` (String) Application name

### Optional

- `app_slo_location` (String) Application single logout location
- `dashboard_url` (String) Application URL used to display error information (combined with error_binding)
- `default_resource` (String) application default resource (after SSO/SLO) i.e. https://myapp.com/home
- `description` (String) Application description
- `error_binding` (String) how errors are displayed to users (combined with dashboard_url)
- `idp` (Block List) SP to IDP SAML 2 settings (see [below for nested schema](#nestedblock--idp))
- `ignored_web_resources` (Set of String) list of URL patterns not subject to SSO control (space sperated)
- `saml2` (Block List) SP SAML 2 settings (see [below for nested schema](#nestedblock--saml2))

### Read-Only

- `element_id` (String) internal element ID
- `id` (String) The ID of this resource.
- `sp_id` (String) Service provider ID. The name of the SP that will be associated with the application

<a id="nestedblock--keystore"></a>
### Nested Schema for `keystore`

Required:

- `password` (String, Sensitive) PKCS12 keystore password
- `resource` (String) PKCS12 keystore in base64 format

Optional:

- `alias` (String) Ceertificate and private key alias (optional)
- `key_password` (String, Sensitive) PKCS12 private key password (optional, the store password is used if not present)


<a id="nestedblock--idp"></a>
### Nested Schema for `idp`

Required:

- `name` (String) name of the trusted IdP

Optional:

- `is_preferred` (Boolean) identifies this IdP as the preferred one (only one IdP must be set to preferred)
- `saml2` (Block List) SP SAML 2 settings (see [below for nested schema](#nestedblock--idp--saml2))

<a id="nestedblock--idp--saml2"></a>
### Nested Schema for `idp.saml2`

Optional:

- `account_linkage` (String) account linkage: which attribute to use as UID from the IdP.
- `bindings` (Block List, Max: 1) enabled SAML bindings (see [below for nested schema](#nestedblock--idp--saml2--bindings))
- `identity_mapping` (String) how the user identity should be mapped for this SP. LOCAL means that the user claims will be retrieved from an identity source connected to the SP.  REMOTE means that claims from the IdP will be used. MERGE is a mix of both claim sets (LOCAL and REMOTE)
- `message_ttl` (Number) SAML message time to live
- `message_ttl_tolerance` (Number) SAML message time to live tolerance
- `sign_authentication_requests` (Boolean) sign authentication requests issued to IdPs
- `sign_requests` (Boolean) sign requests issued to IdPs
- `signature_hash` (String) saml signature hash algorithm
- `want_assertion_signed` (Boolean) require signed assertions from IdPs

<a id="nestedblock--idp--saml2--bindings"></a>
### Nested Schema for `idp.saml2.bindings`

Optional:

- `artifact` (Boolean) use Artifact binding
- `http_post` (Boolean) use HTTP POST binding
- `http_redirect` (Boolean) use HTTP REDIRECT binding
- `local` (Boolean) use LOCAL binding
- `soap` (Boolean) use SOAP binding




<a id="nestedblock--saml2"></a>
### Nested Schema for `saml2`

Optional:

- `account_linkage` (String) account linkage: which attribute to use as UID from the IdP.
- `bindings` (Block List, Max: 1) enabled SAML bindings (see [below for nested schema](#nestedblock--saml2--bindings))
- `identity_mapping` (String) how the user identity should be mapped for this SP. LOCAL means that the user claims will be retrieved from an identity source connected to the SP.  REMOTE means that claims from the IdP will be used. MERGE is a mix of both claim sets (LOCAL and REMOTE)
- `message_ttl` (Number) SAML message time to live
- `message_ttl_tolerance` (Number) SAML message time to live tolerance
- `sign_authentication_requests` (Boolean) sign authentication requests issued to IdPs
- `sign_requests` (Boolean) sign requests issued to IdPs
- `signature_hash` (String) saml signature hash algorithm
- `want_assertion_signed` (Boolean) require signed assertions from IdPs

<a id="nestedblock--saml2--bindings"></a>
### Nested Schema for `saml2.bindings`

Optional:

- `artifact` (Boolean) use Artifact binding
- `http_post` (Boolean) use HTTP POST binding
- `http_redirect` (Boolean) use HTTP REDIRECT binding
- `local` (Boolean) use LOCAL binding
- `soap` (Boolean) use SOAP binding

