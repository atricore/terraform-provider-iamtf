package jossoappi

import "strings"

func (o *AttributeMappingDTO) GetType() string {

	e := o.GetReportedAttrName()

	if strings.HasPrefix(e, "vt:") {
		return "exp"
	} else if strings.HasPrefix(e, "\"") {
		return "const"
	} else {
		return "claim"
	}

}

// Removes type marks from reported attribute name (i.e. vt:, "")
func (o *AttributeMappingDTO) GetMapping() string {

	t := o.GetType()
	s := o.GetReportedAttrName()

	if t == "exp" {
		return strings.TrimSuffix(s, "vt:")
	} else if t == "const" {
		return s[1 : len(s)-1]
	} else {
		return s
	}

}

// Converts a string to a mapping expresion based on the provider type.
// Type must be exp, const or claim
func ToAttributeMapping(t string, s string) string {
	if t == "exp" {
		return "vt:" + s
	} else if t == "const" {
		return "\"" + s + "\""
	} else {
		return s
	}
}
