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

// GetIdpAzuresRes struct for GetIdpAzuresRes
type GetIdpAzuresRes struct {
	Error *string `json:"error,omitempty"`
	Idps []AzureOpenIDConnectIdentityProviderDTO `json:"idps,omitempty"`
	ValidationErrors []string `json:"validationErrors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIdpAzuresRes GetIdpAzuresRes

// NewGetIdpAzuresRes instantiates a new GetIdpAzuresRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdpAzuresRes() *GetIdpAzuresRes {
	this := GetIdpAzuresRes{}
	return &this
}

// NewGetIdpAzuresResWithDefaults instantiates a new GetIdpAzuresRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdpAzuresResWithDefaults() *GetIdpAzuresRes {
	this := GetIdpAzuresRes{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetIdpAzuresRes) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpAzuresRes) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetIdpAzuresRes) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetIdpAzuresRes) SetError(v string) {
	o.Error = &v
}

// GetIdps returns the Idps field value if set, zero value otherwise.
func (o *GetIdpAzuresRes) GetIdps() []AzureOpenIDConnectIdentityProviderDTO {
	if o == nil || isNil(o.Idps) {
		var ret []AzureOpenIDConnectIdentityProviderDTO
		return ret
	}
	return o.Idps
}

// GetIdpsOk returns a tuple with the Idps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpAzuresRes) GetIdpsOk() ([]AzureOpenIDConnectIdentityProviderDTO, bool) {
	if o == nil || isNil(o.Idps) {
    return nil, false
	}
	return o.Idps, true
}

// HasIdps returns a boolean if a field has been set.
func (o *GetIdpAzuresRes) HasIdps() bool {
	if o != nil && !isNil(o.Idps) {
		return true
	}

	return false
}

// SetIdps gets a reference to the given []AzureOpenIDConnectIdentityProviderDTO and assigns it to the Idps field.
func (o *GetIdpAzuresRes) SetIdps(v []AzureOpenIDConnectIdentityProviderDTO) {
	o.Idps = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *GetIdpAzuresRes) GetValidationErrors() []string {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdpAzuresRes) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *GetIdpAzuresRes) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *GetIdpAzuresRes) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o GetIdpAzuresRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Idps) {
		toSerialize["idps"] = o.Idps
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetIdpAzuresRes) UnmarshalJSON(bytes []byte) (err error) {
	varGetIdpAzuresRes := _GetIdpAzuresRes{}

	if err = json.Unmarshal(bytes, &varGetIdpAzuresRes); err == nil {
		*o = GetIdpAzuresRes(varGetIdpAzuresRes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "idps")
		delete(additionalProperties, "validationErrors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIdpAzuresRes struct {
	value *GetIdpAzuresRes
	isSet bool
}

func (v NullableGetIdpAzuresRes) Get() *GetIdpAzuresRes {
	return v.value
}

func (v *NullableGetIdpAzuresRes) Set(val *GetIdpAzuresRes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdpAzuresRes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdpAzuresRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdpAzuresRes(val *GetIdpAzuresRes) *NullableGetIdpAzuresRes {
	return &NullableGetIdpAzuresRes{value: val, isSet: true}
}

func (v NullableGetIdpAzuresRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdpAzuresRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


