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

// GetOidcRpsRes struct for GetOidcRpsRes
type GetOidcRpsRes struct {
	Error *string `json:"error,omitempty"`
	OidcRps []ExternalOpenIDConnectRelayingPartyDTO `json:"oidcRps,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOidcRpsRes GetOidcRpsRes

// NewGetOidcRpsRes instantiates a new GetOidcRpsRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOidcRpsRes() *GetOidcRpsRes {
	this := GetOidcRpsRes{}
	return &this
}

// NewGetOidcRpsResWithDefaults instantiates a new GetOidcRpsRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOidcRpsResWithDefaults() *GetOidcRpsRes {
	this := GetOidcRpsRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetOidcRpsRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOidcRpsRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetOidcRpsRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetOidcRpsRes) SetError(v string) {
	o.Error = &v
}

// GetOidcRps returns the OidcRps field value if set, zero value otherwise.
func (o *GetOidcRpsRes) GetOidcRps() []ExternalOpenIDConnectRelayingPartyDTO {
	if o == nil || isNil(o.OidcRps) {
		var ret []ExternalOpenIDConnectRelayingPartyDTO
		return ret
	}
	return o.OidcRps
}

// GetOidcRpsOk returns a tuple with the OidcRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOidcRpsRes) GetOidcRpsOk() ([]ExternalOpenIDConnectRelayingPartyDTO, bool) {
	if o == nil || isNil(o.OidcRps) {
    return nil, false
	}
	return o.OidcRps, true
}

// HasOidcRps returns a boolean if a field has been set.
func (o *GetOidcRpsRes) HasOidcRps() bool {
	if o != nil && !isNil(o.OidcRps) {
		return true
	}

	return false
}

// SetOidcRps gets a reference to the given []ExternalOpenIDConnectRelayingPartyDTO and assigns it to the OidcRps field.
func (o *GetOidcRpsRes) SetOidcRps(v []ExternalOpenIDConnectRelayingPartyDTO) {
	o.OidcRps = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetOidcRpsRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOidcRpsRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetOidcRpsRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetOidcRpsRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetOidcRpsRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.OidcRps) {
		toSerialize["oidcRps"] = o.OidcRps
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetOidcRpsRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetOidcRpsRes := _GetOidcRpsRes{}

	if err = json.Unmarshal(bytes, &varGetOidcRpsRes); err == nil {
		*o = GetOidcRpsRes(varGetOidcRpsRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "oidcRps")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOidcRpsRes struct {
	value *GetOidcRpsRes
	isSet bool
}

func (v NullableGetOidcRpsRes) Get() *GetOidcRpsRes {
	return v.value
}

func (v *NullableGetOidcRpsRes) Set(val *GetOidcRpsRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOidcRpsRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOidcRpsRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOidcRpsRes(val *GetOidcRpsRes) *NullableGetOidcRpsRes {
	return &NullableGetOidcRpsRes{value: val, isSet: true}
}

func (v NullableGetOidcRpsRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOidcRpsRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


