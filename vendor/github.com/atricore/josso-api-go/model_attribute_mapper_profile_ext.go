package jossoappi

func NewAttriburteMapperProfileDTOInit(n string) *AttributeMapperProfileDTO {
	return newAttributeMapperProfileDTOInit(n, "CUSTOM")
}

func newAttributeMapperProfileDTOInit(n string, t string) *AttributeMapperProfileDTO {
	ba := NewAttributeMapperProfileDTO()
	ba.AdditionalProperties = make(map[string]interface{})
	ba.AdditionalProperties["@c"] = ".AttributeMapperProfileDTO"
	ba.SetName(n)
	ba.SetProfileType(t)

	return ba

}

func (amp *AttributeMapperProfileDTO) ToAttrProfile() (*AttributeProfileDTO, error) {

	m := NewAttributeProfileDTO()
	m.SetName(amp.GetName())
	m.SetProfileType(amp.GetProfileType())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".AttributeMapperProfileDTO"

	var attributeMaps []interface{}
	for _, am := range amp.GetAttributeMaps() {
		mapping := make(map[string]interface{})
		mapping["attrName"] = am.GetAttrName()
		mapping["reportedAttrNameFormat"] = am.GetReportedAttrNameFormat()
		mapping["reportedAttrName"] = am.GetReportedAttrName()
		attributeMaps = append(attributeMaps, mapping)
	}
	m.AdditionalProperties["attributeMaps"] = attributeMaps
	m.AdditionalProperties["includeNonMappedProperties"] = amp.GetIncludeNonMappedProperties()

	return m, nil

}
