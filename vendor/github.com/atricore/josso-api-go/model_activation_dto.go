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

// ActivationDTO struct for ActivationDTO
type ActivationDTO struct {
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	ExecutionEnv *ExecutionEnvironmentDTO `json:"executionEnv,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Resource *ServiceResourceDTO `json:"resource,omitempty"`
	Sp *InternalSaml2ServiceProviderDTO `json:"sp,omitempty"`
	Waypoints []PointDTO `json:"waypoints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActivationDTO ActivationDTO

// NewActivationDTO instantiates a new ActivationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationDTO() *ActivationDTO {
	this := ActivationDTO{}
	return &this
}

// NewActivationDTOWithDefaults instantiates a new ActivationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationDTOWithDefaults() *ActivationDTO {
	this := ActivationDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActivationDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActivationDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActivationDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *ActivationDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *ActivationDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *ActivationDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetExecutionEnv returns the ExecutionEnv field value if set, zero value otherwise.
func (o *ActivationDTO) GetExecutionEnv() ExecutionEnvironmentDTO {
	if o == nil || isNil(o.ExecutionEnv) {
		var ret ExecutionEnvironmentDTO
		return ret
	}
	return *o.ExecutionEnv
}

// GetExecutionEnvOk returns a tuple with the ExecutionEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetExecutionEnvOk() (*ExecutionEnvironmentDTO, bool) {
	if o == nil || isNil(o.ExecutionEnv) {
    return nil, false
	}
	return o.ExecutionEnv, true
}

// HasExecutionEnv returns a boolean if a field has been set.
func (o *ActivationDTO) HasExecutionEnv() bool {
	if o != nil && !isNil(o.ExecutionEnv) {
		return true
	}

	return false
}

// SetExecutionEnv gets a reference to the given ExecutionEnvironmentDTO and assigns it to the ExecutionEnv field.
func (o *ActivationDTO) SetExecutionEnv(v ExecutionEnvironmentDTO) {
	o.ExecutionEnv = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivationDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivationDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ActivationDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActivationDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActivationDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActivationDTO) SetName(v string) {
	o.Name = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ActivationDTO) GetResource() ServiceResourceDTO {
	if o == nil || isNil(o.Resource) {
		var ret ServiceResourceDTO
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetResourceOk() (*ServiceResourceDTO, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ActivationDTO) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ServiceResourceDTO and assigns it to the Resource field.
func (o *ActivationDTO) SetResource(v ServiceResourceDTO) {
	o.Resource = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *ActivationDTO) GetSp() InternalSaml2ServiceProviderDTO {
	if o == nil || isNil(o.Sp) {
		var ret InternalSaml2ServiceProviderDTO
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetSpOk() (*InternalSaml2ServiceProviderDTO, bool) {
	if o == nil || isNil(o.Sp) {
    return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *ActivationDTO) HasSp() bool {
	if o != nil && !isNil(o.Sp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given InternalSaml2ServiceProviderDTO and assigns it to the Sp field.
func (o *ActivationDTO) SetSp(v InternalSaml2ServiceProviderDTO) {
	o.Sp = &v
}

// GetWaypoints returns the Waypoints field value if set, zero value otherwise.
func (o *ActivationDTO) GetWaypoints() []PointDTO {
	if o == nil || isNil(o.Waypoints) {
		var ret []PointDTO
		return ret
	}
	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationDTO) GetWaypointsOk() ([]PointDTO, bool) {
	if o == nil || isNil(o.Waypoints) {
    return nil, false
	}
	return o.Waypoints, true
}

// HasWaypoints returns a boolean if a field has been set.
func (o *ActivationDTO) HasWaypoints() bool {
	if o != nil && !isNil(o.Waypoints) {
		return true
	}

	return false
}

// SetWaypoints gets a reference to the given []PointDTO and assigns it to the Waypoints field.
func (o *ActivationDTO) SetWaypoints(v []PointDTO) {
	o.Waypoints = v
}

func (o ActivationDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.ExecutionEnv) {
		toSerialize["executionEnv"] = o.ExecutionEnv
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.Sp) {
		toSerialize["sp"] = o.Sp
	}
	if !isNil(o.Waypoints) {
		toSerialize["waypoints"] = o.Waypoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ActivationDTO) UnmarshalJSON(bytes []byte) (err error) {
	varActivationDTO := _ActivationDTO{}

	if err = json.Unmarshal(bytes, &varActivationDTO); err == nil {
		*o = ActivationDTO(varActivationDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "executionEnv")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "waypoints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActivationDTO struct {
	value *ActivationDTO
	isSet bool
}

func (v NullableActivationDTO) Get() *ActivationDTO {
	return v.value
}

func (v *NullableActivationDTO) Set(val *ActivationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationDTO(val *ActivationDTO) *NullableActivationDTO {
	return &NullableActivationDTO{value: val, isSet: true}
}

func (v NullableActivationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

