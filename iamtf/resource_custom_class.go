package iamtf

import (
	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func customClassSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "CustomClass settings",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fqcn": {
					Type:        schema.TypeString,
					Description: "Component FQCN.",
					Required:    true,
				},
				"osgiFilter": {
					Type:        schema.TypeString,
					Description: "TODO.",
					Required:    true,
				},
				"osgiService": {
					Type:        schema.TypeBool,
					Description: "TODO.",
					Required:    true,
				},
				"properties": {
					Type:        schema.TypeString,
					Description: "TODO.",
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
		"properties":  cc.GetProperties(),
	}
	result = append(result, cc_map)

	return result, nil
}

func convertCustomClassMapArrToDTO(name string, cc_arr interface{}) (*api.CustomClassDTO, error) {
	var cc *api.CustomClassDTO
	cc = api.NewCustomClassDTO()
	cc.SetFqcn()
	cc.SetOsgiFilter()
	cc.SetOsgiService(false)
	cc.SetProperties()
	return cc, nil

}
