package jossoappi

// Creates a new resource DTO with name, display name and value
func NewResourceDTOInit(n string, d string, v string) *ResourceDTO {
	s := NewResourceDTO()
	s.AdditionalProperties = map[string]interface{}{}
	s.AdditionalProperties["@c"] = ".ResourceDTO"
	s.SetName(n)
	s.SetDisplayName(d)
	s.SetValue(v)
	return s
}
