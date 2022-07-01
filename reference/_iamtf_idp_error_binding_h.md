The error binding allows you to customize and handle errors.  Depending on the binding type, JOSSO will encode errors differently and redirect the user to **dashboard_url** including the encoded error information.

* ARTIFACT: reserver for JOSSO internal usage.
* JSON: send error details encoded in JSON format.
* GET: send error details encoded in HTTP request query strings.