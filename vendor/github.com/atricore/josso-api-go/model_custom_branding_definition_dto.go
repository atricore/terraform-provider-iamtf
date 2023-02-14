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

// CustomBrandingDefinitionDTO struct for CustomBrandingDefinitionDTO
type CustomBrandingDefinitionDTO struct {
	BundleSymbolicName *string `json:"bundleSymbolicName,omitempty"`
	BundleUri *string `json:"bundleUri,omitempty"`
	CustomOpenIdAppClazz *string `json:"customOpenIdAppClazz,omitempty"`
	CustomSsoAppClazz *string `json:"customSsoAppClazz,omitempty"`
	CustomSsoIdPAppClazz *string `json:"customSsoIdPAppClazz,omitempty"`
	DefaultLocale *string `json:"defaultLocale,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Type *string `json:"type,omitempty"`
	WebBrandingId *string `json:"webBrandingId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomBrandingDefinitionDTO CustomBrandingDefinitionDTO

// NewCustomBrandingDefinitionDTO instantiates a new CustomBrandingDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBrandingDefinitionDTO() *CustomBrandingDefinitionDTO {
	this := CustomBrandingDefinitionDTO{}
	return &this
}

// NewCustomBrandingDefinitionDTOWithDefaults instantiates a new CustomBrandingDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBrandingDefinitionDTOWithDefaults() *CustomBrandingDefinitionDTO {
	this := CustomBrandingDefinitionDTO{}
	return &this
}

// GetBundleSymbolicName returns the BundleSymbolicName field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetBundleSymbolicName() string {
	if o == nil || isNil(o.BundleSymbolicName) {
		var ret string
		return ret
	}
	return *o.BundleSymbolicName
}

// GetBundleSymbolicNameOk returns a tuple with the BundleSymbolicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetBundleSymbolicNameOk() (*string, bool) {
	if o == nil || isNil(o.BundleSymbolicName) {
    return nil, false
	}
	return o.BundleSymbolicName, true
}

// HasBundleSymbolicName returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasBundleSymbolicName() bool {
	if o != nil && !isNil(o.BundleSymbolicName) {
		return true
	}

	return false
}

// SetBundleSymbolicName gets a reference to the given string and assigns it to the BundleSymbolicName field.
func (o *CustomBrandingDefinitionDTO) SetBundleSymbolicName(v string) {
	o.BundleSymbolicName = &v
}

// GetBundleUri returns the BundleUri field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetBundleUri() string {
	if o == nil || isNil(o.BundleUri) {
		var ret string
		return ret
	}
	return *o.BundleUri
}

// GetBundleUriOk returns a tuple with the BundleUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetBundleUriOk() (*string, bool) {
	if o == nil || isNil(o.BundleUri) {
    return nil, false
	}
	return o.BundleUri, true
}

// HasBundleUri returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasBundleUri() bool {
	if o != nil && !isNil(o.BundleUri) {
		return true
	}

	return false
}

// SetBundleUri gets a reference to the given string and assigns it to the BundleUri field.
func (o *CustomBrandingDefinitionDTO) SetBundleUri(v string) {
	o.BundleUri = &v
}

// GetCustomOpenIdAppClazz returns the CustomOpenIdAppClazz field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetCustomOpenIdAppClazz() string {
	if o == nil || isNil(o.CustomOpenIdAppClazz) {
		var ret string
		return ret
	}
	return *o.CustomOpenIdAppClazz
}

// GetCustomOpenIdAppClazzOk returns a tuple with the CustomOpenIdAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetCustomOpenIdAppClazzOk() (*string, bool) {
	if o == nil || isNil(o.CustomOpenIdAppClazz) {
    return nil, false
	}
	return o.CustomOpenIdAppClazz, true
}

// HasCustomOpenIdAppClazz returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasCustomOpenIdAppClazz() bool {
	if o != nil && !isNil(o.CustomOpenIdAppClazz) {
		return true
	}

	return false
}

// SetCustomOpenIdAppClazz gets a reference to the given string and assigns it to the CustomOpenIdAppClazz field.
func (o *CustomBrandingDefinitionDTO) SetCustomOpenIdAppClazz(v string) {
	o.CustomOpenIdAppClazz = &v
}

// GetCustomSsoAppClazz returns the CustomSsoAppClazz field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetCustomSsoAppClazz() string {
	if o == nil || isNil(o.CustomSsoAppClazz) {
		var ret string
		return ret
	}
	return *o.CustomSsoAppClazz
}

// GetCustomSsoAppClazzOk returns a tuple with the CustomSsoAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetCustomSsoAppClazzOk() (*string, bool) {
	if o == nil || isNil(o.CustomSsoAppClazz) {
    return nil, false
	}
	return o.CustomSsoAppClazz, true
}

// HasCustomSsoAppClazz returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasCustomSsoAppClazz() bool {
	if o != nil && !isNil(o.CustomSsoAppClazz) {
		return true
	}

	return false
}

// SetCustomSsoAppClazz gets a reference to the given string and assigns it to the CustomSsoAppClazz field.
func (o *CustomBrandingDefinitionDTO) SetCustomSsoAppClazz(v string) {
	o.CustomSsoAppClazz = &v
}

// GetCustomSsoIdPAppClazz returns the CustomSsoIdPAppClazz field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetCustomSsoIdPAppClazz() string {
	if o == nil || isNil(o.CustomSsoIdPAppClazz) {
		var ret string
		return ret
	}
	return *o.CustomSsoIdPAppClazz
}

// GetCustomSsoIdPAppClazzOk returns a tuple with the CustomSsoIdPAppClazz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetCustomSsoIdPAppClazzOk() (*string, bool) {
	if o == nil || isNil(o.CustomSsoIdPAppClazz) {
    return nil, false
	}
	return o.CustomSsoIdPAppClazz, true
}

// HasCustomSsoIdPAppClazz returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasCustomSsoIdPAppClazz() bool {
	if o != nil && !isNil(o.CustomSsoIdPAppClazz) {
		return true
	}

	return false
}

// SetCustomSsoIdPAppClazz gets a reference to the given string and assigns it to the CustomSsoIdPAppClazz field.
func (o *CustomBrandingDefinitionDTO) SetCustomSsoIdPAppClazz(v string) {
	o.CustomSsoIdPAppClazz = &v
}

// GetDefaultLocale returns the DefaultLocale field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetDefaultLocale() string {
	if o == nil || isNil(o.DefaultLocale) {
		var ret string
		return ret
	}
	return *o.DefaultLocale
}

// GetDefaultLocaleOk returns a tuple with the DefaultLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetDefaultLocaleOk() (*string, bool) {
	if o == nil || isNil(o.DefaultLocale) {
    return nil, false
	}
	return o.DefaultLocale, true
}

// HasDefaultLocale returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasDefaultLocale() bool {
	if o != nil && !isNil(o.DefaultLocale) {
		return true
	}

	return false
}

// SetDefaultLocale gets a reference to the given string and assigns it to the DefaultLocale field.
func (o *CustomBrandingDefinitionDTO) SetDefaultLocale(v string) {
	o.DefaultLocale = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomBrandingDefinitionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CustomBrandingDefinitionDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomBrandingDefinitionDTO) SetName(v string) {
	o.Name = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetResource() string {
	if o == nil || isNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetResourceOk() (*string, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *CustomBrandingDefinitionDTO) SetResource(v string) {
	o.Resource = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomBrandingDefinitionDTO) SetType(v string) {
	o.Type = &v
}

// GetWebBrandingId returns the WebBrandingId field value if set, zero value otherwise.
func (o *CustomBrandingDefinitionDTO) GetWebBrandingId() string {
	if o == nil || isNil(o.WebBrandingId) {
		var ret string
		return ret
	}
	return *o.WebBrandingId
}

// GetWebBrandingIdOk returns a tuple with the WebBrandingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBrandingDefinitionDTO) GetWebBrandingIdOk() (*string, bool) {
	if o == nil || isNil(o.WebBrandingId) {
    return nil, false
	}
	return o.WebBrandingId, true
}

// HasWebBrandingId returns a boolean if a field has been set.
func (o *CustomBrandingDefinitionDTO) HasWebBrandingId() bool {
	if o != nil && !isNil(o.WebBrandingId) {
		return true
	}

	return false
}

// SetWebBrandingId gets a reference to the given string and assigns it to the WebBrandingId field.
func (o *CustomBrandingDefinitionDTO) SetWebBrandingId(v string) {
	o.WebBrandingId = &v
}

func (o CustomBrandingDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BundleSymbolicName) {
		toSerialize["bundleSymbolicName"] = o.BundleSymbolicName
	}
	if !isNil(o.BundleUri) {
		toSerialize["bundleUri"] = o.BundleUri
	}
	if !isNil(o.CustomOpenIdAppClazz) {
		toSerialize["customOpenIdAppClazz"] = o.CustomOpenIdAppClazz
	}
	if !isNil(o.CustomSsoAppClazz) {
		toSerialize["customSsoAppClazz"] = o.CustomSsoAppClazz
	}
	if !isNil(o.CustomSsoIdPAppClazz) {
		toSerialize["customSsoIdPAppClazz"] = o.CustomSsoIdPAppClazz
	}
	if !isNil(o.DefaultLocale) {
		toSerialize["defaultLocale"] = o.DefaultLocale
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.WebBrandingId) {
		toSerialize["webBrandingId"] = o.WebBrandingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomBrandingDefinitionDTO) UnmarshalJSON(bytes []byte) (err error) {
	varCustomBrandingDefinitionDTO := _CustomBrandingDefinitionDTO{}

	if err = json.Unmarshal(bytes, &varCustomBrandingDefinitionDTO); err == nil {
		*o = CustomBrandingDefinitionDTO(varCustomBrandingDefinitionDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bundleSymbolicName")
		delete(additionalProperties, "bundleUri")
		delete(additionalProperties, "customOpenIdAppClazz")
		delete(additionalProperties, "customSsoAppClazz")
		delete(additionalProperties, "customSsoIdPAppClazz")
		delete(additionalProperties, "defaultLocale")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "type")
		delete(additionalProperties, "webBrandingId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomBrandingDefinitionDTO struct {
	value *CustomBrandingDefinitionDTO
	isSet bool
}

func (v NullableCustomBrandingDefinitionDTO) Get() *CustomBrandingDefinitionDTO {
	return v.value
}

func (v *NullableCustomBrandingDefinitionDTO) Set(val *CustomBrandingDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBrandingDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBrandingDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBrandingDefinitionDTO(val *CustomBrandingDefinitionDTO) *NullableCustomBrandingDefinitionDTO {
	return &NullableCustomBrandingDefinitionDTO{value: val, isSet: true}
}

func (v NullableCustomBrandingDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBrandingDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

