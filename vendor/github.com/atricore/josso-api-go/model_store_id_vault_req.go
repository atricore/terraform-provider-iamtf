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

// StoreIdVaultReq struct for StoreIdVaultReq
type StoreIdVaultReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	IdVault *EmbeddedIdentityVaultDTO `json:"idVault,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdVaultReq StoreIdVaultReq

// NewStoreIdVaultReq instantiates a new StoreIdVaultReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdVaultReq() *StoreIdVaultReq {
	this := StoreIdVaultReq{}
	return &this
}

// NewStoreIdVaultReqWithDefaults instantiates a new StoreIdVaultReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdVaultReqWithDefaults() *StoreIdVaultReq {
	this := StoreIdVaultReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *StoreIdVaultReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdVaultReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *StoreIdVaultReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *StoreIdVaultReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdVault returns the IdVault field value if set, zero value otherwise.
func (o *StoreIdVaultReq) GetIdVault() EmbeddedIdentityVaultDTO {
	if o == nil || isNil(o.IdVault) {
		var ret EmbeddedIdentityVaultDTO
		return ret
	}
	return *o.IdVault
}

// GetIdVaultOk returns a tuple with the IdVault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdVaultReq) GetIdVaultOk() (*EmbeddedIdentityVaultDTO, bool) {
	if o == nil || isNil(o.IdVault) {
    return nil, false
	}
	return o.IdVault, true
}

// HasIdVault returns a boolean if a field has been set.
func (o *StoreIdVaultReq) HasIdVault() bool {
	if o != nil && !isNil(o.IdVault) {
		return true
	}

	return false
}

// SetIdVault gets a reference to the given EmbeddedIdentityVaultDTO and assigns it to the IdVault field.
func (o *StoreIdVaultReq) SetIdVault(v EmbeddedIdentityVaultDTO) {
	o.IdVault = &v
}

func (o StoreIdVaultReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.IdVault) {
		toSerialize["idVault"] = o.IdVault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreIdVaultReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdVaultReq := _StoreIdVaultReq{}

	if err = json.Unmarshal(bytes, &varStoreIdVaultReq); err == nil {
		*o = StoreIdVaultReq(varStoreIdVaultReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idVault")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIdVaultReq struct {
	value *StoreIdVaultReq
	isSet bool
}

func (v NullableStoreIdVaultReq) Get() *StoreIdVaultReq {
	return v.value
}

func (v *NullableStoreIdVaultReq) Set(val *StoreIdVaultReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdVaultReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdVaultReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdVaultReq(val *StoreIdVaultReq) *NullableStoreIdVaultReq {
	return &NullableStoreIdVaultReq{value: val, isSet: true}
}

func (v NullableStoreIdVaultReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdVaultReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


