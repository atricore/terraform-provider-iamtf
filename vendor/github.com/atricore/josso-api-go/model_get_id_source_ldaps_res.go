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

// GetIdSourceLdapsRes struct for GetIdSourceLdapsRes
type GetIdSourceLdapsRes struct {
	Error *string `json:"error,omitempty"`
	IdSourceLdaps []LdapIdentitySourceDTO `json:"idSourceLdaps,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdSourceLdapsRes GetIdSourceLdapsRes

// NewGetIdSourceLdapsRes instantiates a new GetIdSourceLdapsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdSourceLdapsRes() *GetIdSourceLdapsRes {
	this := GetIdSourceLdapsRes{}
	return &this
}

// NewGetIdSourceLdapsResWithDefaults instantiates a new GetIdSourceLdapsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdSourceLdapsResWithDefaults() *GetIdSourceLdapsRes {
	this := GetIdSourceLdapsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIdSourceLdapsRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapsRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIdSourceLdapsRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIdSourceLdapsRes) SetError(v string) {
	o.Error = &v
}

// GetIdSourceLdaps returns the IdSourceLdaps field value if set, zero value otherwise.
func (o *GetIdSourceLdapsRes) GetIdSourceLdaps() []LdapIdentitySourceDTO {
	if o == nil || isNil(o.IdSourceLdaps) {
		var ret []LdapIdentitySourceDTO
		return ret
	}
	return o.IdSourceLdaps
}

// GetIdSourceLdapsOk returns a tuple with the IdSourceLdaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapsRes) GetIdSourceLdapsOk() ([]LdapIdentitySourceDTO, bool) {
	if o == nil || isNil(o.IdSourceLdaps) {
    return nil, false
	}
	return o.IdSourceLdaps, true
}

// HasIdSourceLdaps returns a boolean if a field has been set.
func (o *GetIdSourceLdapsRes) HasIdSourceLdaps() bool {
	if o != nil && !isNil(o.IdSourceLdaps) {
		return true
	}

	return false
}

// SetIdSourceLdaps gets a reference to the given []LdapIdentitySourceDTO and assigns it to the IdSourceLdaps field.
func (o *GetIdSourceLdapsRes) SetIdSourceLdaps(v []LdapIdentitySourceDTO) {
	o.IdSourceLdaps = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIdSourceLdapsRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdSourceLdapsRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIdSourceLdapsRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIdSourceLdapsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIdSourceLdapsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.IdSourceLdaps) {
		toSerialize["idSourceLdaps"] = o.IdSourceLdaps
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIdSourceLdapsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdSourceLdapsRes := _GetIdSourceLdapsRes{}

	if err = json.Unmarshal(bytes, &varGetIdSourceLdapsRes); err == nil {
		*o = GetIdSourceLdapsRes(varGetIdSourceLdapsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "idSourceLdaps")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdSourceLdapsRes struct {
	value *GetIdSourceLdapsRes
	isSet bool
}

func (v NullableGetIdSourceLdapsRes) Get() *GetIdSourceLdapsRes {
	return v.value
}

func (v *NullableGetIdSourceLdapsRes) Set(val *GetIdSourceLdapsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdSourceLdapsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdSourceLdapsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdSourceLdapsRes(val *GetIdSourceLdapsRes) *NullableGetIdSourceLdapsRes {
	return &NullableGetIdSourceLdapsRes{value: val, isSet: true}
}

func (v NullableGetIdSourceLdapsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdSourceLdapsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

