The following is an example of an identity appliance definition.  The resource **name** and name property must match.  The **namespace** must be unique in the server and follow java package naming convention.  (use characters and dots). The **location** is the base public URL to access all the services.

```
resource "iamtf_identity_appliance" "my-ida" {
  name        = "my-ida"
  namespace   = "com.mycompany.myida"
  description = "My identity appliance!"
  location    = "https://mysso.mycompany.com"
}

```

JOSSO will create all service URLs under the two paths: **/IDBUS** and **/IDBUS-UI**.

In our example, these are thebase locations:

* https://mysso.mycompany.com/IDBUS/MY-IDA
* https://mysso.mycompany.com/IDBUS-UI/MY-IDA

:::tip
You should configure your load balancer or reverse proxy to **ONLY forward requests using these paths to JOSSO servers**.
:::
