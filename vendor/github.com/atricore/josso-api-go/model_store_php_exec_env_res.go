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

// StorePhpExecEnvRes struct for StorePhpExecEnvRes
type StorePhpExecEnvRes struct {
	Error *string `json:"error,omitempty"`
	PhpExecEnv *PHPExecutionEnvironmentDTO `json:"phpExecEnv,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorePhpExecEnvRes StorePhpExecEnvRes

// NewStorePhpExecEnvRes instantiates a new StorePhpExecEnvRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorePhpExecEnvRes() *StorePhpExecEnvRes {
	this := StorePhpExecEnvRes{}
	return &this
}

// NewStorePhpExecEnvResWithDefaults instantiates a new StorePhpExecEnvRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorePhpExecEnvResWithDefaults() *StorePhpExecEnvRes {
	this := StorePhpExecEnvRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StorePhpExecEnvRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorePhpExecEnvRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StorePhpExecEnvRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StorePhpExecEnvRes) SetError(v string) {
	o.Error = &v
}

// GetPhpExecEnv returns the PhpExecEnv field value if set, zero value otherwise.
func (o *StorePhpExecEnvRes) GetPhpExecEnv() PHPExecutionEnvironmentDTO {
	if o == nil || isNil(o.PhpExecEnv) {
		var ret PHPExecutionEnvironmentDTO
		return ret
	}
	return *o.PhpExecEnv
}

// GetPhpExecEnvOk returns a tuple with the PhpExecEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorePhpExecEnvRes) GetPhpExecEnvOk() (*PHPExecutionEnvironmentDTO, bool) {
	if o == nil || isNil(o.PhpExecEnv) {
    return nil, false
	}
	return o.PhpExecEnv, true
}

// HasPhpExecEnv returns a boolean if a field has been set.
func (o *StorePhpExecEnvRes) HasPhpExecEnv() bool {
	if o != nil && !isNil(o.PhpExecEnv) {
		return true
	}

	return false
}

// SetPhpExecEnv gets a reference to the given PHPExecutionEnvironmentDTO and assigns it to the PhpExecEnv field.
func (o *StorePhpExecEnvRes) SetPhpExecEnv(v PHPExecutionEnvironmentDTO) {
	o.PhpExecEnv = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *StorePhpExecEnvRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorePhpExecEnvRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *StorePhpExecEnvRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *StorePhpExecEnvRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o StorePhpExecEnvRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.PhpExecEnv) {
		toSerialize["phpExecEnv"] = o.PhpExecEnv
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorePhpExecEnvRes) UnmarshalJSON(bytes []byte) (err error) {
	varStorePhpExecEnvRes := _StorePhpExecEnvRes{}

	if err = json.Unmarshal(bytes, &varStorePhpExecEnvRes); err == nil {
		*o = StorePhpExecEnvRes(varStorePhpExecEnvRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "phpExecEnv")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorePhpExecEnvRes struct {
	value *StorePhpExecEnvRes
	isSet bool
}

func (v NullableStorePhpExecEnvRes) Get() *StorePhpExecEnvRes {
	return v.value
}

func (v *NullableStorePhpExecEnvRes) Set(val *StorePhpExecEnvRes) {
	v.value = val
	v.isSet = true
}

func (v NullableStorePhpExecEnvRes) IsSet() bool {
	return v.isSet
}

func (v *NullableStorePhpExecEnvRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorePhpExecEnvRes(val *StorePhpExecEnvRes) *NullableStorePhpExecEnvRes {
	return &NullableStorePhpExecEnvRes{value: val, isSet: true}
}

func (v NullableStorePhpExecEnvRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorePhpExecEnvRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


