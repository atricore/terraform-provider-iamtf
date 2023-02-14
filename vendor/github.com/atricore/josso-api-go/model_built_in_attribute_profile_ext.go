package jossoappi

func (bap BuiltInAttributeProfileDTO) ToAttrProfile() (*AttributeProfileDTO, error) {

	m := NewAttributeProfileDTO()
	m.SetName(bap.GetName())
	m.SetProfileType(bap.GetProfileType())

	m.AdditionalProperties = make(map[string]interface{})
	m.AdditionalProperties["@c"] = ".BuiltInAttributeProfileDTO"

	return m, nil

}

func NewBasicAttributeProfileDTOInit(n string) *BuiltInAttributeProfileDTO {
	return newBuiltInAttributeProfileDTOInit(n, "BASIC")
}

func NewJOSSOAttributeProfileDTOInit(n string) *BuiltInAttributeProfileDTO {
	return newBuiltInAttributeProfileDTOInit(n, "JOSSO")
}

func NewOneToOneAttributeProfileDTOInit(n string) *BuiltInAttributeProfileDTO {
	return newBuiltInAttributeProfileDTOInit(n, "ONE_TO_ONE")
}

func newBuiltInAttributeProfileDTOInit(n string, t string) *BuiltInAttributeProfileDTO {
	ba := NewBuiltInAttributeProfileDTO()
	ba.AdditionalProperties = make(map[string]interface{})
	ba.AdditionalProperties["@c"] = ".BuiltInAttributeProfileDTO"
	ba.SetName(n)
	ba.SetProfileType(t)

	return ba

}
