/*
Atricore Console :: Remote : API

# Atricore Console API

API version: 1.5.0-SNAPSHOT
Contact: sgonzalez@atricore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jossoappi

import (
	"encoding/json"
)

// GetIntSaml2SpReq struct for GetIntSaml2SpReq
type GetIntSaml2SpReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIntSaml2SpReq GetIntSaml2SpReq

// NewGetIntSaml2SpReq instantiates a new GetIntSaml2SpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIntSaml2SpReq() *GetIntSaml2SpReq {
	this := GetIntSaml2SpReq{}
	return &this
}

// NewGetIntSaml2SpReqWithDefaults instantiates a new GetIntSaml2SpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIntSaml2SpReqWithDefaults() *GetIntSaml2SpReq {
	this := GetIntSaml2SpReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *GetIntSaml2SpReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *GetIntSaml2SpReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *GetIntSaml2SpReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIntSaml2SpReq) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpReq) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIntSaml2SpReq) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIntSaml2SpReq) SetName(v string) {
	o.Name = &v
}

func (o GetIntSaml2SpReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIntSaml2SpReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetIntSaml2SpReq := _GetIntSaml2SpReq{}

	if err = json.Unmarshal(bytes, &varGetIntSaml2SpReq); err == nil {
		*o = GetIntSaml2SpReq(varGetIntSaml2SpReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIntSaml2SpReq struct {
	value *GetIntSaml2SpReq
	isSet bool
}

func (v NullableGetIntSaml2SpReq) Get() *GetIntSaml2SpReq {
	return v.value
}

func (v *NullableGetIntSaml2SpReq) Set(val *GetIntSaml2SpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIntSaml2SpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIntSaml2SpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIntSaml2SpReq(val *GetIntSaml2SpReq) *NullableGetIntSaml2SpReq {
	return &NullableGetIntSaml2SpReq{value: val, isSet: true}
}

func (v NullableGetIntSaml2SpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIntSaml2SpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


