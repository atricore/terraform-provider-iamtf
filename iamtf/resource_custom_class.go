package iamtf

import (
	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func customClassSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Allows you to use a custom component for a given resource.  Components are installed as OSGi bundles in the server.  You can refer to a component instance or create a new instance based on its class",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fqcn": {
					Type:        schema.TypeString,
					Description: "component java FQCN. Refers to the OSGi component type or Java class to be instantiated",
					Required:    true,
				},
				"type": {
					Type:             schema.TypeString,
					Description:      "extension type: SERVICE (for OSGi service references) or INSTANCE (for creating a new instance). ",
					ValidateDiagFunc: stringInSlice([]string{"SERVICE", "INSTANCE"}),
					Optional:         true,
					Default:          "SERVICE",
				},
				"osgi_filter": {
					Type:        schema.TypeString,
					Description: "filter to locate the OSGi service (Only when extension type is SERVICE).",
					Optional:    true,
				},
				"property": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "list of configuration properties and its values (only when extension type is INSTANCE)",
					MinItems:    0,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name as the property",
								Required:    true,
							},
							"value": {
								Type:        schema.TypeString,
								Description: "Value as the property ",
								Required:    true,
							},
						},
					},
				},
			},
		},
	}
}

func convertCustomClassDTOToMapArr(cc *api.CustomClassDTO) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// If cc is null, return an empty map
	if cc == nil {
		return result, nil
	}

	et := "SERVICE"
	if !cc.GetOsgiService() {
		et = "INSTANCE"
	}

	var mapProps []map[string]interface{}
	for _, m := range cc.GetProperties() {
		mMap := make(map[string]interface{})
		mMap["name"] = m.GetName()
		mMap["value"] = m.GetValue()
		mapProps = append(mapProps, mMap)
	}

	cc_map := map[string]interface{}{
		"fqcn":        cc.GetFqcn(),
		"osgi_filter": cc.GetOsgiFilter(),
		"type":        et,
		"property":    mapProps,
	}
	result = append(result, cc_map)

	return result, nil
}

func convertCustomClassMapArrToDTO(cc_arr interface{}) (*api.CustomClassDTO, error) {
	var cc *api.CustomClassDTO
	tfMapLs, err := asTFMapSingle(cc_arr)
	if err != nil {
		return cc, err
	}
	// If map is empty, return nil
	if tfMapLs == nil || len(tfMapLs) == 0 {
		return cc, nil
	}

	ncc := api.NewCustomClassDTO()
	ncc.SetFqcn(api.AsString(tfMapLs["fqcn"], ""))
	ncc.SetOsgiFilter(api.AsString(tfMapLs["osgi_filter"], ""))

	if api.AsString(tfMapLs["type"], "") == "INSTANCE" {
		ncc.SetOsgiService(false)
	} else {
		ncc.SetOsgiService(true)
	}

	nm := tfMapLs["property"].(*schema.Set)
	var props []api.CustomClassPropertyDTO
	for _, v := range nm.List() {
		prop := v.(map[string]interface{})
		ccp := api.NewCustomClassPropertyDTO()
		ccp.SetName(prop["name"].(string))
		ccp.SetValue(prop["value"].(string))
		props = append(props, *ccp)
	}
	ncc.SetProperties(props)
	return ncc, nil
}
