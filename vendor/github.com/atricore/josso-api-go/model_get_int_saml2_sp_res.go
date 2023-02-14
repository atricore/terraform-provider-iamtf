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

// GetIntSaml2SpRes struct for GetIntSaml2SpRes
type GetIntSaml2SpRes struct {
	Config *SamlR2SPConfigDTO `json:"config,omitempty"`
	Error *string `json:"error,omitempty"`
	Sp *InternalSaml2ServiceProviderDTO `json:"sp,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIntSaml2SpRes GetIntSaml2SpRes

// NewGetIntSaml2SpRes instantiates a new GetIntSaml2SpRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIntSaml2SpRes() *GetIntSaml2SpRes {
	this := GetIntSaml2SpRes{}
	return &this
}

// NewGetIntSaml2SpResWithDefaults instantiates a new GetIntSaml2SpRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIntSaml2SpResWithDefaults() *GetIntSaml2SpRes {
	this := GetIntSaml2SpRes{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GetIntSaml2SpRes) GetConfig() SamlR2SPConfigDTO {
	if o == nil || isNil(o.Config) {
		var ret SamlR2SPConfigDTO
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpRes) GetConfigOk() (*SamlR2SPConfigDTO, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *GetIntSaml2SpRes) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SamlR2SPConfigDTO and assigns it to the Config field.
func (o *GetIntSaml2SpRes) SetConfig(v SamlR2SPConfigDTO) {
	o.Config = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIntSaml2SpRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIntSaml2SpRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIntSaml2SpRes) SetError(v string) {
	o.Error = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *GetIntSaml2SpRes) GetSp() InternalSaml2ServiceProviderDTO {
	if o == nil || isNil(o.Sp) {
		var ret InternalSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpRes) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool) {
	if o == nil || isNil(o.Sp) {
    return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *GetIntSaml2SpRes) HasSp() bool {
	if o != nil && !isNil(o.Sp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given InternalSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *GetIntSaml2SpRes) SetSp(v InternalSaml2ServiceProviderDTO) {
	o.Sp = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIntSaml2SpRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIntSaml2SpRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIntSaml2SpRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIntSaml2SpRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIntSaml2SpRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Sp) {
		toSerialize["sp"] = o.Sp
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIntSaml2SpRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIntSaml2SpRes := _GetIntSaml2SpRes{}

	if err = json.Unmarshal(bytes, &varGetIntSaml2SpRes); err == nil {
		*o = GetIntSaml2SpRes(varGetIntSaml2SpRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "error")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIntSaml2SpRes struct {
	value *GetIntSaml2SpRes
	isSet bool
}

func (v NullableGetIntSaml2SpRes) Get() *GetIntSaml2SpRes {
	return v.value
}

func (v *NullableGetIntSaml2SpRes) Set(val *GetIntSaml2SpRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIntSaml2SpRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIntSaml2SpRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIntSaml2SpRes(val *GetIntSaml2SpRes) *NullableGetIntSaml2SpRes {
	return &NullableGetIntSaml2SpRes{value: val, isSet: true}
}

func (v NullableGetIntSaml2SpRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIntSaml2SpRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

