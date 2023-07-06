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

// ExecEnvContainerDTO struct for ExecEnvContainerDTO
type ExecEnvContainerDTO struct {
	Captive *bool `json:"captive,omitempty"`
	ExecEnv *ExecutionEnvironmentDTO `json:"execEnv,omitempty"`
	Location *string `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecEnvContainerDTO ExecEnvContainerDTO

// NewExecEnvContainerDTO instantiates a new ExecEnvContainerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecEnvContainerDTO() *ExecEnvContainerDTO {
	this := ExecEnvContainerDTO{}
	return &this
}

// NewExecEnvContainerDTOWithDefaults instantiates a new ExecEnvContainerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecEnvContainerDTOWithDefaults() *ExecEnvContainerDTO {
	this := ExecEnvContainerDTO{}
	return &this
}

// GetCaptive returns the Captive field value if set, zero value otherwise.
func (o *ExecEnvContainerDTO) GetCaptive() bool {
	if o == nil || isNil(o.Captive) {
		var ret bool
		return ret
	}
	return *o.Captive
}

// GetCaptiveOk returns a tuple with the Captive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecEnvContainerDTO) GetCaptiveOk() (*bool, bool) {
	if o == nil || isNil(o.Captive) {
    return nil, false
	}
	return o.Captive, true
}

// HasCaptive returns a boolean if a field has been set.
func (o *ExecEnvContainerDTO) HasCaptive() bool {
	if o != nil && !isNil(o.Captive) {
		return true
	}

	return false
}

// SetCaptive gets a reference to the given bool and assigns it to the Captive field.
func (o *ExecEnvContainerDTO) SetCaptive(v bool) {
	o.Captive = &v
}

// GetExecEnv returns the ExecEnv field value if set, zero value otherwise.
func (o *ExecEnvContainerDTO) GetExecEnv() ExecutionEnvironmentDTO {
	if o == nil || isNil(o.ExecEnv) {
		var ret ExecutionEnvironmentDTO
		return ret
	}
	return *o.ExecEnv
}

// GetExecEnvOk returns a tuple with the ExecEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecEnvContainerDTO) GetExecEnvOk() (*ExecutionEnvironmentDTO, bool) {
	if o == nil || isNil(o.ExecEnv) {
    return nil, false
	}
	return o.ExecEnv, true
}

// HasExecEnv returns a boolean if a field has been set.
func (o *ExecEnvContainerDTO) HasExecEnv() bool {
	if o != nil && !isNil(o.ExecEnv) {
		return true
	}

	return false
}

// SetExecEnv gets a reference to the given ExecutionEnvironmentDTO and assigns it to the ExecEnv field.
func (o *ExecEnvContainerDTO) SetExecEnv(v ExecutionEnvironmentDTO) {
	o.ExecEnv = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ExecEnvContainerDTO) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecEnvContainerDTO) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ExecEnvContainerDTO) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ExecEnvContainerDTO) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExecEnvContainerDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecEnvContainerDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExecEnvContainerDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExecEnvContainerDTO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExecEnvContainerDTO) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecEnvContainerDTO) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExecEnvContainerDTO) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExecEnvContainerDTO) SetType(v string) {
	o.Type = &v
}

func (o ExecEnvContainerDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Captive) {
		toSerialize["captive"] = o.Captive
	}
	if !isNil(o.ExecEnv) {
		toSerialize["execEnv"] = o.ExecEnv
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExecEnvContainerDTO) UnmarshalJSON(bytes []byte) (err error) {
	varExecEnvContainerDTO := _ExecEnvContainerDTO{}

	if err = json.Unmarshal(bytes, &varExecEnvContainerDTO); err == nil {
		*o = ExecEnvContainerDTO(varExecEnvContainerDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "captive")
		delete(additionalProperties, "execEnv")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecEnvContainerDTO struct {
	value *ExecEnvContainerDTO
	isSet bool
}

func (v NullableExecEnvContainerDTO) Get() *ExecEnvContainerDTO {
	return v.value
}

func (v *NullableExecEnvContainerDTO) Set(val *ExecEnvContainerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableExecEnvContainerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableExecEnvContainerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecEnvContainerDTO(val *ExecEnvContainerDTO) *NullableExecEnvContainerDTO {
	return &NullableExecEnvContainerDTO{value: val, isSet: true}
}

func (v NullableExecEnvContainerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecEnvContainerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


