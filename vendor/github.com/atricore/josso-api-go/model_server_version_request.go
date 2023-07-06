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

// ServerVersionRequest struct for ServerVersionRequest
type ServerVersionRequest struct {
	Server *ServerContext `json:"server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerVersionRequest ServerVersionRequest

// NewServerVersionRequest instantiates a new ServerVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerVersionRequest() *ServerVersionRequest {
	this := ServerVersionRequest{}
	return &this
}

// NewServerVersionRequestWithDefaults instantiates a new ServerVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerVersionRequestWithDefaults() *ServerVersionRequest {
	this := ServerVersionRequest{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ServerVersionRequest) GetServer() ServerContext {
	if o == nil || isNil(o.Server) {
		var ret ServerContext
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVersionRequest) GetServerOk() (*ServerContext, bool) {
	if o == nil || isNil(o.Server) {
    return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ServerVersionRequest) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given ServerContext and assigns it to the Server field.
func (o *ServerVersionRequest) SetServer(v ServerContext) {
	o.Server = &v
}

func (o ServerVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerVersionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varServerVersionRequest := _ServerVersionRequest{}

	if err = json.Unmarshal(bytes, &varServerVersionRequest); err == nil {
		*o = ServerVersionRequest(varServerVersionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerVersionRequest struct {
	value *ServerVersionRequest
	isSet bool
}

func (v NullableServerVersionRequest) Get() *ServerVersionRequest {
	return v.value
}

func (v *NullableServerVersionRequest) Set(val *ServerVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerVersionRequest(val *ServerVersionRequest) *NullableServerVersionRequest {
	return &NullableServerVersionRequest{value: val, isSet: true}
}

func (v NullableServerVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


