package iamtf

import (
	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func customClassSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Allows you to use a custom component for a given resource.  Componentse are installed as OSGi extensions.  You can refer to a component instance or create a new instance based on its class",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fqcn": {
					Type:        schema.TypeString,
					Description: "Component FQCN. Refers to the OSGi component class or Java class to be instantiated",
					Required:    true,
				},
				"osgiFilter": {
					Type:        schema.TypeString,
					Description: "filter to locate the OSGi component.",
					Required:    true,
				},
				"osgiService": {
					Type:        schema.TypeBool,
					Description: "TODO.",
					Required:    true,
				},
				// TODO : Must be a list
				"properties": {
					Type:        schema.TypeString,
					Description: "List of configuration properties and its values",
					Optional:    true,
					Required:    true,
				},
			},
		},
	}
}

func convertCustomClassDTOToMapArr(cc *api.CustomClassDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)
	cc_map := map[string]interface{}{
		"fqcn":        cc.GetFqcn(),
		"osgiFilter":  cc.GetOsgiFilter(),
		"osgiService": cc.GetOsgiService(),
		// TODO : Array of CustomClassPropertyDTO "properties":  cc.GetProperties(),
	}
	result = append(result, cc_map)

	return result, nil
}

func convertCustomClassMapArrToDTO(name string, cc_arr interface{}) (*api.CustomClassDTO, error) {
	var cc *api.CustomClassDTO
	cc = api.NewCustomClassDTO()
	cc.SetFqcn("fqcn")
	cc.SetOsgiFilter("osgiFilter")
	cc.SetOsgiService(false)
	// TODO : Array of CustomClassPropertyDTO  cc.setProperties("properties")
	return cc, nil

}
