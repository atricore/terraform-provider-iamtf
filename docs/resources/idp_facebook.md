---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iamtf_idp_facebook Resource - terraform-provider-iamtf"
subcategory: ""
description: |-
  
---

# iamtf_idp_facebook (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_token_service` (String) facebook access token endpoint
- `client_id` (String) facebook application id
- `client_secret` (String) facebook application secret
- `ida` (String) identity appliance name
- `name` (String) idp name

### Optional

- `authz_token_service` (String) facebook authorization token endpoint
- `description` (String) idp description
- `scopes` (Set of String) facebook premissions. These will be added to **email** and **public_profile**
- `user_fields` (Set of String) facebook user fields. These will be added to **id**, **name**, **email**, **first_name**, **last_name**

### Read-Only

- `id` (String) The ID of this resource.

