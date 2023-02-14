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

// GetTomcatExecEnvReq struct for GetTomcatExecEnvReq
type GetTomcatExecEnvReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetTomcatExecEnvReq GetTomcatExecEnvReq

// NewGetTomcatExecEnvReq instantiates a new GetTomcatExecEnvReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTomcatExecEnvReq() *GetTomcatExecEnvReq {
	this := GetTomcatExecEnvReq{}
	return &this
}

// NewGetTomcatExecEnvReqWithDefaults instantiates a new GetTomcatExecEnvReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTomcatExecEnvReqWithDefaults() *GetTomcatExecEnvReq {
	this := GetTomcatExecEnvReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *GetTomcatExecEnvReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTomcatExecEnvReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *GetTomcatExecEnvReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *GetTomcatExecEnvReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetTomcatExecEnvReq) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTomcatExecEnvReq) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetTomcatExecEnvReq) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetTomcatExecEnvReq) SetName(v string) {
	o.Name = &v
}

func (o GetTomcatExecEnvReq) MarshalJSON() ([]byte, error) {
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

func (o *GetTomcatExecEnvReq) UnmarshalJSON(bytes []byte) (err error) {
	varGetTomcatExecEnvReq := _GetTomcatExecEnvReq{}

	if err = json.Unmarshal(bytes, &varGetTomcatExecEnvReq); err == nil {
		*o = GetTomcatExecEnvReq(varGetTomcatExecEnvReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetTomcatExecEnvReq struct {
	value *GetTomcatExecEnvReq
	isSet bool
}

func (v NullableGetTomcatExecEnvReq) Get() *GetTomcatExecEnvReq {
	return v.value
}

func (v *NullableGetTomcatExecEnvReq) Set(val *GetTomcatExecEnvReq) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTomcatExecEnvReq) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTomcatExecEnvReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTomcatExecEnvReq(val *GetTomcatExecEnvReq) *NullableGetTomcatExecEnvReq {
	return &NullableGetTomcatExecEnvReq{value: val, isSet: true}
}

func (v NullableGetTomcatExecEnvReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTomcatExecEnvReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

