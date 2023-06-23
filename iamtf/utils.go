package iamtf

import (
	"fmt"
	"reflect"
	"strings"

	cli "github.com/atricore/josso-sdk-go"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api "github.com/atricore/josso-api-go"
)

// This builds a space separetd string with the values of the given schema resource (TypeSet -> TypeString)
func PtrSchemaAsSpacedList(d *schema.ResourceData, p string) (bool, string) {
	m := d.Get(p).(*schema.Set)

	if m == nil {
		return false, ""
	}

	ls := m.List()
	if len(ls) < 1 {
		return false, ""
	}

	var sb strings.Builder
	for _, v := range ls {
		sb.WriteString(v.(string))
		sb.WriteString(" ")
	}
	r := sb.String()
	return true, r[:len(r)-1]
}

func SpacedListToSet(strLs string) (bool, *schema.Set) {

	if strLs == "" {
		return false, nil
	}

	ls := strings.Split(strLs, " ")
	if len(ls) < 1 {
		return false, nil
	}
	return true, convertStringSetToInterface(ls)
}

// Gets the value of the resources property as a string pointer
func PtrSchemaStr(d *schema.ResourceData, p string) *string {
	a := d.Get(p)
	var v string
	switch a.(type) {
	case nil:
		return nil
	default:
		v = a.(string)
	}
	return &v
}
func PtrSchemaStrPointer(d *schema.ResourceData, p *string) *string {
	a := d.Get(*p)
	var v *string
	switch a.(type) {
	case nil:
		return nil
	default:
		v = a.(*string)
	}
	return v
}

func PtrSchemaInt(d *schema.ResourceData, p string) *int {
	a := d.Get(p)
	var v int
	switch a.(type) {
	case nil:
		return nil
	default:
		v = a.(int)
	}
	return &v
}

func PtrSchemaInt32(d *schema.ResourceData, p string) *int32 {
	a := d.Get(p)
	var v int32
	switch a.(type) {
	case nil:
		return nil
	default:
		v = int32(a.(int))
	}
	return &v
}

func PtrSchemaInt64(d *schema.ResourceData, p string) *int64 {
	a := d.Get(p)
	var v int64
	switch a.(type) {
	case nil:
		return nil
	default:
		v = int64(a.(int))
	}
	return &v
}

func PtrSchemaBool(d *schema.ResourceData, p string) *bool {
	a := d.Get(p)
	var v bool
	switch a.(type) {
	case nil:
		return &v
	default:
		v = a.(bool)
	}
	return &v
}

func PtrSchemaLocation(d *schema.ResourceData, p string) (*api.LocationDTO, error) {
	l := d.Get(p).(string)
	return cli.StrToLocation(l)
}

func getJossoClient(meta interface{}) *cli.IdbusApiClient {
	return meta.(*Config).apiClient
}

func getLogger(meta interface{}) hclog.Logger {
	return meta.(*Config).logger
}

func convertStringArrToIdLookups(str []string) []api.IdentityLookupDTO {
	arr := make([]api.IdentityLookupDTO, len(str))
	for i, s := range str {
		name := s
		arr[i] = api.IdentityLookupDTO{Name: &name}
		arr[i].AdditionalProperties = make(map[string]interface{})
		arr[i].AdditionalProperties["@c"] = ".IdentityLookupDTO"
	}
	return arr
}

func convertIdLookupsToStringArr(dtos []api.IdentityLookupDTO) []string {
	arr := make([]string, len(dtos))
	for i, dto := range dtos {
		arr[i] = *dto.Name
	}
	return arr
}

// --------------------------------------------

// Takes the first element of the array and returns its value as a map[string]interface{}
func asTFMapSingle(arr interface{}) (map[string]interface{}, error) {

	var result map[string]interface{}

	ls, err := asTFMap(arr, 1)
	if err != nil {
		return result, err
	}

	if len(ls) == 0 {
		return result, nil
	}

	result, ok := ls[0].(map[string]interface{}) //
	if !ok {
		return result, fmt.Errorf("invalid argument value type %s", reflect.TypeOf(ls[0]))
	}

	return result, nil

}

func asTFMapAll(arr interface{}) ([]interface{}, error) {
	return asTFMap(arr, 0)
}

// when maxLen is 0, it is umbounded
func asTFMap(arr interface{}, maxLen int) ([]interface{}, error) {

	ls, ok := arr.([]interface{})
	if !ok {
		return ls, fmt.Errorf("invalid argument type %s", reflect.TypeOf(arr))
	}

	if maxLen > 0 && len(ls) > maxLen {
		return ls, fmt.Errorf("invalid number of elements %d", len(ls))
	}

	return ls, nil
}

// --------------------------------------------------------------------

func convertStringArrToFederatedConnections(idps []string) []api.FederatedConnectionDTO {
	arr := make([]api.FederatedConnectionDTO, len(idps))
	for i, s := range idps {
		name := s
		arr[i] = *api.NewFederatedConnectionDTO()
		arr[i].Name = &name
		arr[i].AdditionalProperties = make(map[string]interface{})
		arr[i].AdditionalProperties["@c"] = ".FederatedConnectionDTO"
	}
	return arr
}

func convertFederatedConnectionsToStringArr(dtos []api.FederatedConnectionDTO) []string {
	arr := make([]string, len(dtos))
	for i, dto := range dtos {
		arr[i] = *dto.Name
	}
	return arr
}

func convertInterfaceToStringSet(set interface{}) []string {
	if set == nil {
		return make([]string, 0)
	}
	return convertInterfaceToStringArr(set.(*schema.Set).List())
}

func convertInterfaceToStringSetNullable(set interface{}) []string {
	if set == nil {
		return make([]string, 0)
	}
	return convertInterfaceToStringArrNullable(set.(*schema.Set).List())
}

func convertInterfaceToStringArrNullable(ls interface{}) []string {
	arr := convertInterfaceToStringArr(ls)

	if len(arr) < 1 {
		return nil
	}

	return arr
}

func convertInterfaceToStringArr(ls interface{}) []string {
	var arr []string
	lsArr, ok := ls.([]interface{})
	if ok {
		arr = make([]string, len(lsArr))
		for i, str := range lsArr {
			s := str.(string)
			arr[i] = s
		}
	}

	return arr

}

func convertStringSetToInterface(stringList []string) *schema.Set {
	arr := make([]interface{}, len(stringList))
	for i, str := range stringList {
		arr[i] = str
	}
	return schema.NewSet(schema.HashString, arr)
}

// The best practices states that aggregate types should have error handling (think non-primitive). This will not attempt to set nil values.
func setNonPrimitives(d *schema.ResourceData, valueMap map[string]interface{}) error {
	for k, v := range valueMap {
		if v != nil {
			if err := d.Set(k, v); err != nil {
				return fmt.Errorf("error setting %s for resource %s: %s", k, d.Id(), err)
			}
		}
	}
	return nil
}

func GetAsBool(d *schema.ResourceData, k string, v bool) bool {
	// Although deprecated, this is actually what we need. GetOK does not work when values are 'false'
	r, e := d.GetOkExists(k)
	if !e {
		return v
	}
	return r.(bool)
}

func GetAsString(d *schema.ResourceData, k string, v string) string {
	r, e := d.GetOk(k)
	if !e {
		return v
	}
	return r.(string)
}

func GetAsInt32(d *schema.ResourceData, k string, v int32) int32 {
	r, e := d.GetOk(k)
	if !e {
		return v
	}

	var result int32
	switch v := r.(type) {
	case int32:
		result = v
	case int:
		result = int32(v)
	case int64:
		result = int32(v)
	case float32:
		result = int32(v)
	case float64:
		result = int32(v)
	}

	return result
}

func GetAsIface(d *schema.ResourceData, k string, v interface{}) interface{} {
	r, e := d.GetOk(k)
	if !e {
		return v
	}
	return r
}

func convertSubjectAuthnPoliciesDTOToMapArr(ap []api.SubjectAuthenticationPolicyDTO) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	for _, sa := range ap {
		subjetMap := map[string]interface{}{
			"name": sa.GetName(),
		}

		result = append(result, subjetMap)
	}
	return result, nil
}

func convertSubjectAuthnPoliciesMapArrToDTO(ap_arr interface{}) ([]api.SubjectAuthenticationPolicyDTO, error) {
	var sap []api.SubjectAuthenticationPolicyDTO
	tfMapLs, err := asTFMapSingle(ap_arr)
	if err != nil {
		return sap, err
	}
	if tfMapLs == nil {
		return sap, err
	}

	nsap := api.NewSubjectAuthenticationPolicyDTO()
	nsap.SetName(api.AsString(tfMapLs["name"], ""))
	sap = append(sap, *nsap)
	return sap, nil
}
