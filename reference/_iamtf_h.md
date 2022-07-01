
The JOSSO terreaform plugin allows you to manage identity appliances and appliance elements as terraform resources in a JOSSO server.

**main.tf**
```
terraform {
  required_providers {
    josso = {
      version = "~> 0.1.4"
      source  = "atricore.com/iam/josso"
    }
  }
}

```

You can configure the plugin directly in your terraform descriptor, as follows. 

**provider.tf**

```
provider "josso" {
  org_name      = "my company"
  endpoint      = "http://localhost:8081/atricore-rest/services"
  client_id     = "idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7"
  client_secret = "changeme"
}
```

You can also use environment valirables, and set minimun configuration in your plugin descriptor:

```
export JOSSO_API_CLIENT_ID=idbus-f2f7244e-bbce-44ca-8b33-f5c0bde339f7
export JOSSO_API_CLIENT_SECRET=changeme
export JOSSO_API_ENDPOINT=http://localhost:8081/atricore-rest/services
```

```
provider "josso" {
    org_name = "my company"
}
```
