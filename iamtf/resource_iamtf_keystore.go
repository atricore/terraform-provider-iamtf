package iamtf

import (
	"fmt"

	api "github.com/atricore/josso-api-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func convertKeystoreDTOToMapArr(ks *api.KeystoreDTO) ([]map[string]interface{}, error) {

	result := make([]map[string]interface{}, 0)
	rs := ks.GetStore()
	ks_map := map[string]interface{}{
		"resource": rs.GetValue(),
		"password": ks.GetPassword(),
	}
	result = append(result, ks_map)

	return result, nil
}

func convertKeystoreMapArrToDTO(name string, ks_arr interface{}) (*api.KeystoreDTO, error) {
	var ks *api.KeystoreDTO
	ks_map, err := asTFMapSingle(ks_arr) //
	if err != nil {
		return ks, err
	}

	if ks_map == nil {
		return ks, nil
	}

	if ks_map["resource"] == nil {
		return ks, fmt.Errorf("resource value not present for %s", name)
	}
	if ks_map["password"] == nil {
		return ks, fmt.Errorf("password value not present for %s", name)
	}

	rs := api.NewResourceDTOInit(
		fmt.Sprintf("%s-ks", name),
		fmt.Sprintf("%s ks", name),
		ks_map["resource"].(string))
	rs.SetUri(fmt.Sprintf("%s-ks.p12", name))

	// Keystores MUST be pkcs12 format
	ks = api.NewKeystoreDTOInit(
		fmt.Sprintf("%s-ks", name),
		fmt.Sprintf("%s ks", name),
		rs)

	ks.SetKeystorePassOnly(true)
	ks.SetPassword(ks_map["password"].(string))
	ks.SetType("PKCS12")

	ks.SetCertificateAlias(api.AsString(ks_map["alias"], ""))
	ks.SetPrivateKeyPassword(api.AsString(ks_map["key_password"], ""))
	ks.SetPrivateKeyName(api.AsString(ks_map["alias"], ""))

	return ks, nil

}

func keystoreSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		Description: "Keystore configuration.  A single keystore containing the private key and certificate is supported.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resource": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "PKCS12 keystore in base64 format",
				},
				"password": {
					Description: "PKCS12 keystore password",
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
				},
				"alias": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Certificate and private key alias (optional)",
				},
				"key_password": {
					Description: "PKCS12 private key password (optional, the store password is used if not present)",
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
				},
			},
		},
	}

}
