package jossoappi

func (o *AttributeProfileDTO) ToAttributeMapperProfile() *AttributeMapperProfileDTO {

	amp := NewAttributeMapperProfileDTO()
	amp.SetName(o.GetName())
	amp.SetProfileType(o.GetProfileType())
	amp.SetElementId(o.GetElementId())

	amp.AdditionalProperties = make(map[string]interface{})
	amp.AdditionalProperties["@c"] = ".AttributeMapperProfileDTO"
	if o.AdditionalProperties["attributeMaps"] != nil {

		ls := o.AdditionalProperties["attributeMaps"].([]interface{})
		var attrMaps []AttributeMappingDTO
		for _, i := range ls {

			m := i.(map[string]interface{})

			attrMap := NewAttributeMappingDTO()
			attrMap.SetAttrName(AsString(m["attrName"], ""))
			attrMap.SetReportedAttrName(AsString(m["reportedAttrName"], ""))
			attrMap.SetReportedAttrNameFormat(AsString(m["reportedAttrNameFormat"], ""))
			attrMaps = append(attrMaps, *attrMap)
		}

		amp.SetAttributeMaps(attrMaps)
		amp.SetIncludeNonMappedProperties(AsBool(o.AdditionalProperties["includeNonMappedProperties"], true))
	}

	return amp
}
