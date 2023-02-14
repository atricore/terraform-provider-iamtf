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

// StoreIdSourceLdapReq struct for StoreIdSourceLdapReq
type StoreIdSourceLdapReq struct {
	IdOrName *string `json:"idOrName,omitempty"`
	IdSourceLdap *LdapIdentitySourceDTO `json:"idSourceLdap,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreIdSourceLdapReq StoreIdSourceLdapReq

// NewStoreIdSourceLdapReq instantiates a new StoreIdSourceLdapReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreIdSourceLdapReq() *StoreIdSourceLdapReq {
	this := StoreIdSourceLdapReq{}
	return &this
}

// NewStoreIdSourceLdapReqWithDefaults instantiates a new StoreIdSourceLdapReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreIdSourceLdapReqWithDefaults() *StoreIdSourceLdapReq {
	this := StoreIdSourceLdapReq{}
	return &this
}

// GetIdOrName returns the IdOrName field value if set, zero value otherwise.
func (o *StoreIdSourceLdapReq) GetIdOrName() string {
	if o == nil || isNil(o.IdOrName) {
		var ret string
		return ret
	}
	return *o.IdOrName
}

// GetIdOrNameOk returns a tuple with the IdOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdSourceLdapReq) GetIdOrNameOk() (*string, bool) {
	if o == nil || isNil(o.IdOrName) {
    return nil, false
	}
	return o.IdOrName, true
}

// HasIdOrName returns a boolean if a field has been set.
func (o *StoreIdSourceLdapReq) HasIdOrName() bool {
	if o != nil && !isNil(o.IdOrName) {
		return true
	}

	return false
}

// SetIdOrName gets a reference to the given string and assigns it to the IdOrName field.
func (o *StoreIdSourceLdapReq) SetIdOrName(v string) {
	o.IdOrName = &v
}

// GetIdSourceLdap returns the IdSourceLdap field value if set, zero value otherwise.
func (o *StoreIdSourceLdapReq) GetIdSourceLdap() LdapIdentitySourceDTO {
	if o == nil || isNil(o.IdSourceLdap) {
		var ret LdapIdentitySourceDTO
		return ret
	}
	return *o.IdSourceLdap
}

// GetIdSourceLdapOk returns a tuple with the IdSourceLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreIdSourceLdapReq) GetIdSourceLdapOk() (*LdapIdentitySourceDTO, bool) {
	if o == nil || isNil(o.IdSourceLdap) {
    return nil, false
	}
	return o.IdSourceLdap, true
}

// HasIdSourceLdap returns a boolean if a field has been set.
func (o *StoreIdSourceLdapReq) HasIdSourceLdap() bool {
	if o != nil && !isNil(o.IdSourceLdap) {
		return true
	}

	return false
}

// SetIdSourceLdap gets a reference to the given LdapIdentitySourceDTO and assigns it to the IdSourceLdap field.
func (o *StoreIdSourceLdapReq) SetIdSourceLdap(v LdapIdentitySourceDTO) {
	o.IdSourceLdap = &v
}

func (o StoreIdSourceLdapReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdOrName) {
		toSerialize["idOrName"] = o.IdOrName
	}
	if !isNil(o.IdSourceLdap) {
		toSerialize["idSourceLdap"] = o.IdSourceLdap
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoreIdSourceLdapReq) UnmarshalJSON(bytes []byte) (err error) {
	varStoreIdSourceLdapReq := _StoreIdSourceLdapReq{}

	if err = json.Unmarshal(bytes, &varStoreIdSourceLdapReq); err == nil {
		*o = StoreIdSourceLdapReq(varStoreIdSourceLdapReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "idOrName")
		delete(additionalProperties, "idSourceLdap")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreIdSourceLdapReq struct {
	value *StoreIdSourceLdapReq
	isSet bool
}

func (v NullableStoreIdSourceLdapReq) Get() *StoreIdSourceLdapReq {
	return v.value
}

func (v *NullableStoreIdSourceLdapReq) Set(val *StoreIdSourceLdapReq) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreIdSourceLdapReq) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreIdSourceLdapReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreIdSourceLdapReq(val *StoreIdSourceLdapReq) *NullableStoreIdSourceLdapReq {
	return &NullableStoreIdSourceLdapReq{value: val, isSet: true}
}

func (v NullableStoreIdSourceLdapReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreIdSourceLdapReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

