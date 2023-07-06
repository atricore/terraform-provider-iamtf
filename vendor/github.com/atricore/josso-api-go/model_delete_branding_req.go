/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.5.1-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// DeleteBrandingReq struct for DeleteBrandingReq
type DeleteBrandingReq struct {
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteBrandingReq DeleteBrandingReq

// NewDeleteBrandingReq instantiates a new DeleteBrandingReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBrandingReq() *DeleteBrandingReq {
	this := DeleteBrandingReq{}
	return &this
}

// NewDeleteBrandingReqWithDefaults instantiates a new DeleteBrandingReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBrandingReqWithDefaults() *DeleteBrandingReq {
	this := DeleteBrandingReq{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeleteBrandingReq) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBrandingReq) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeleteBrandingReq) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeleteBrandingReq) SetName(v string) {
	o.Name = &v
}

func (o DeleteBrandingReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteBrandingReq) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteBrandingReq := _DeleteBrandingReq{}

	if err = json.Unmarshal(bytes, &varDeleteBrandingReq); err == nil {
		*o = DeleteBrandingReq(varDeleteBrandingReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteBrandingReq struct {
	value *DeleteBrandingReq
	isSet bool
}

func (v NullableDeleteBrandingReq) Get() *DeleteBrandingReq {
	return v.value
}

func (v *NullableDeleteBrandingReq) Set(val *DeleteBrandingReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBrandingReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBrandingReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBrandingReq(val *DeleteBrandingReq) *NullableDeleteBrandingReq {
	return &NullableDeleteBrandingReq{value: val, isSet: true}
}

func (v NullableDeleteBrandingReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBrandingReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


