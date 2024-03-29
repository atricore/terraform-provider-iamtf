---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iamtf_idsource_db Resource - terraform-provider-iamtf"
subcategory: ""
description: |-
  
---

# iamtf_idsource_db (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connectionurl` (String) jdbc connection string
- `ida` (String) identity appliance name
- `name` (String) identiy source name
- `sql_credentials` (String) credentials query string. Must return a single row with columns: username, password, salt (optional)
- `sql_groups` (String) user groups query string.  Must return a single column with group names
- `sql_user_attrs` (String) user attributes query string. Must return a single row with columns: username, name, value
- `sql_username` (String) username query string. Used to retrieve the username from the DB

### Optional

- `acquire_increment` (Number) number of connections to aquire when incrementing the pool
- `connection_pool` (Boolean) enable a connection pool
- `description` (String) resource description
- `dml_reset_credential` (String) query string used to update the password credential
- `extension` (Block List) Allows you to use a custom component for a given resource.  Components are installed as OSGi bundles in the server.  You can refer to a component instance or create a new instance based on its class (see [below for nested schema](#nestedblock--extension))
- `idle_connection_test_period` (Number) dbidentitysource idleconnectiontestperiod
- `initial_pool_size` (Number) dbidentitysource initialpoolsize
- `jdbc_driver` (String) JDBC driver
- `max_idle_time` (Number) dbidentitysource maxidletime
- `max_pool_size` (Number) dbidentitysource maxpoolsize
- `min_pool_size` (Number) dbidentitysource minpoolsize
- `password` (String) connection password
- `sql_relay_credential` (String) query string to retrieve the credential/claim used to recover a password (i.e. email)
- `use_column_name_as_property_name` (Boolean) Use sql_user_attrs result-set column names as properties names
- `username` (String) connection username

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--extension"></a>
### Nested Schema for `extension`

Required:

- `fqcn` (String) component java FQCN. Refers to the OSGi component type or Java class to be instantiated

Optional:

- `osgi_filter` (String) filter to locate the OSGi service (Only when extension type is SERVICE).
- `property` (Block Set) list of configuration properties and its values (only when extension type is INSTANCE) (see [below for nested schema](#nestedblock--extension--property))
- `type` (String) extension type: SERVICE (for OSGi service references) or INSTANCE (for creating a new instance).

<a id="nestedblock--extension--property"></a>
### Nested Schema for `extension.property`

Required:

- `name` (String) Name as the property
- `value` (String) Value as the property


