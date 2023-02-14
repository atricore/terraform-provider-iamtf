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

// FederatedConnectionDTO struct for FederatedConnectionDTO
type FederatedConnectionDTO struct {
	ChannelA *FederatedChannelDTO `json:"channelA,omitempty"`
	ChannelB *FederatedChannelDTO `json:"channelB,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	RoleA *FederatedProviderDTO `json:"roleA,omitempty"`
	RoleB *FederatedProviderDTO `json:"roleB,omitempty"`
	Waypoints []PointDTO `json:"waypoints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederatedConnectionDTO FederatedConnectionDTO

// NewFederatedConnectionDTO instantiates a new FederatedConnectionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederatedConnectionDTO() *FederatedConnectionDTO {
	this := FederatedConnectionDTO{}
	return &this
}

// NewFederatedConnectionDTOWithDefaults instantiates a new FederatedConnectionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederatedConnectionDTOWithDefaults() *FederatedConnectionDTO {
	this := FederatedConnectionDTO{}
	return &this
}

// GetChannelA returns the ChannelA field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetChannelA() FederatedChannelDTO {
	if o == nil || isNil(o.ChannelA) {
		var ret FederatedChannelDTO
		return ret
	}
	return *o.ChannelA
}

// GetChannelAOk returns a tuple with the ChannelA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetChannelAOk() (*FederatedChannelDTO, bool) {
	if o == nil || isNil(o.ChannelA) {
    return nil, false
	}
	return o.ChannelA, true
}

// HasChannelA returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasChannelA() bool {
	if o != nil && !isNil(o.ChannelA) {
		return true
	}

	return false
}

// SetChannelA gets a reference to the given FederatedChannelDTO and assigns it to the ChannelA field.
func (o *FederatedConnectionDTO) SetChannelA(v FederatedChannelDTO) {
	o.ChannelA = &v
}

// GetChannelB returns the ChannelB field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetChannelB() FederatedChannelDTO {
	if o == nil || isNil(o.ChannelB) {
		var ret FederatedChannelDTO
		return ret
	}
	return *o.ChannelB
}

// GetChannelBOk returns a tuple with the ChannelB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetChannelBOk() (*FederatedChannelDTO, bool) {
	if o == nil || isNil(o.ChannelB) {
    return nil, false
	}
	return o.ChannelB, true
}

// HasChannelB returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasChannelB() bool {
	if o != nil && !isNil(o.ChannelB) {
		return true
	}

	return false
}

// SetChannelB gets a reference to the given FederatedChannelDTO and assigns it to the ChannelB field.
func (o *FederatedConnectionDTO) SetChannelB(v FederatedChannelDTO) {
	o.ChannelB = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FederatedConnectionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *FederatedConnectionDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FederatedConnectionDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FederatedConnectionDTO) SetName(v string) {
	o.Name = &v
}

// GetRoleA returns the RoleA field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetRoleA() FederatedProviderDTO {
	if o == nil || isNil(o.RoleA) {
		var ret FederatedProviderDTO
		return ret
	}
	return *o.RoleA
}

// GetRoleAOk returns a tuple with the RoleA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetRoleAOk() (*FederatedProviderDTO, bool) {
	if o == nil || isNil(o.RoleA) {
    return nil, false
	}
	return o.RoleA, true
}

// HasRoleA returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasRoleA() bool {
	if o != nil && !isNil(o.RoleA) {
		return true
	}

	return false
}

// SetRoleA gets a reference to the given FederatedProviderDTO and assigns it to the RoleA field.
func (o *FederatedConnectionDTO) SetRoleA(v FederatedProviderDTO) {
	o.RoleA = &v
}

// GetRoleB returns the RoleB field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetRoleB() FederatedProviderDTO {
	if o == nil || isNil(o.RoleB) {
		var ret FederatedProviderDTO
		return ret
	}
	return *o.RoleB
}

// GetRoleBOk returns a tuple with the RoleB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetRoleBOk() (*FederatedProviderDTO, bool) {
	if o == nil || isNil(o.RoleB) {
    return nil, false
	}
	return o.RoleB, true
}

// HasRoleB returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasRoleB() bool {
	if o != nil && !isNil(o.RoleB) {
		return true
	}

	return false
}

// SetRoleB gets a reference to the given FederatedProviderDTO and assigns it to the RoleB field.
func (o *FederatedConnectionDTO) SetRoleB(v FederatedProviderDTO) {
	o.RoleB = &v
}

// GetWaypoints returns the Waypoints field value if set, zero value otherwise.
func (o *FederatedConnectionDTO) GetWaypoints() []PointDTO {
	if o == nil || isNil(o.Waypoints) {
		var ret []PointDTO
		return ret
	}
	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedConnectionDTO) GetWaypointsOk() ([]PointDTO, bool) {
	if o == nil || isNil(o.Waypoints) {
    return nil, false
	}
	return o.Waypoints, true
}

// HasWaypoints returns a boolean if a field has been set.
func (o *FederatedConnectionDTO) HasWaypoints() bool {
	if o != nil && !isNil(o.Waypoints) {
		return true
	}

	return false
}

// SetWaypoints gets a reference to the given []PointDTO and assigns it to the Waypoints field.
func (o *FederatedConnectionDTO) SetWaypoints(v []PointDTO) {
	o.Waypoints = v
}

func (o FederatedConnectionDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ChannelA) {
		toSerialize["channelA"] = o.ChannelA
	}
	if !isNil(o.ChannelB) {
		toSerialize["channelB"] = o.ChannelB
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.RoleA) {
		toSerialize["roleA"] = o.RoleA
	}
	if !isNil(o.RoleB) {
		toSerialize["roleB"] = o.RoleB
	}
	if !isNil(o.Waypoints) {
		toSerialize["waypoints"] = o.Waypoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FederatedConnectionDTO) UnmarshalJSON(bytes []byte) (err error) {
	varFederatedConnectionDTO := _FederatedConnectionDTO{}

	if err = json.Unmarshal(bytes, &varFederatedConnectionDTO); err == nil {
		*o = FederatedConnectionDTO(varFederatedConnectionDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "channelA")
		delete(additionalProperties, "channelB")
		delete(additionalProperties, "description")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "roleA")
		delete(additionalProperties, "roleB")
		delete(additionalProperties, "waypoints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederatedConnectionDTO struct {
	value *FederatedConnectionDTO
	isSet bool
}

func (v NullableFederatedConnectionDTO) Get() *FederatedConnectionDTO {
	return v.value
}

func (v *NullableFederatedConnectionDTO) Set(val *FederatedConnectionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFederatedConnectionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFederatedConnectionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederatedConnectionDTO(val *FederatedConnectionDTO) *NullableFederatedConnectionDTO {
	return &NullableFederatedConnectionDTO{value: val, isSet: true}
}

func (v NullableFederatedConnectionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederatedConnectionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


