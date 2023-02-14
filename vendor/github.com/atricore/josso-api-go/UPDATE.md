# Updating the YAML file

* Update **console-api-XXX-swagger.json** file
* Manually modify updated file as described bellow
* Run **make all**


::: tip
Some manual changes are required due to limitations/errors in the client generation tool.
:::

## ResourceDTO

Modify type, value should be:

```json
"value" : {
"type" : "string"
}
```

## CustomBarandingDefinitionDTO, StoreBrandingReq, GetBrandingRes

Modify type, resource should be:

```json
"resource" : {
  "type" : "string"
}
```

## IdentityProviderDTO

Modify type, add identity lookups property:

```json
"identityLookups" : {
            "uniqueItems" : true,
            "type" : "array",
            "items" : {
              "$ref" : "#/components/schemas/IdentityLookupDTO"
            }
          },
```

# Patching the file

## Generate patch file

```sh
diff -u console-api-1.5.0-SNAPSHOT-swagger.json console-api-1.5.0-SNAPSHOT-swagger.json.new > console.patch
```

## Remove all changes to abobe mentioned types

Remove changes to CustomBarandingDefinitionDTO, StoreBrandingReq, GetBrandingRes, ResourceDTO, IdentityProviderDTO

## Apply the patch

```sh
patch console-api-1.5.0-SNAPSHOT-swagger.json < console.patch
```