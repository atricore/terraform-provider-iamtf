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

// TomcatExecutionEnvironmentDTO struct for TomcatExecutionEnvironmentDTO
type TomcatExecutionEnvironmentDTO struct {
	Activations []ActivationDTO `json:"activations,omitempty"`
	Active *bool `json:"active,omitempty"`
	BindingLocation *LocationDTO `json:"bindingLocation,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	InstallDemoApps *bool `json:"installDemoApps,omitempty"`
	InstallUri *string `json:"installUri,omitempty"`
	Location *string `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	OverwriteOriginalSetup *bool `json:"overwriteOriginalSetup,omitempty"`
	PlatformId *string `json:"platformId,omitempty"`
	TargetJDK *string `json:"targetJDK,omitempty"`
	Type *string `json:"type,omitempty"`
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TomcatExecutionEnvironmentDTO TomcatExecutionEnvironmentDTO

// NewTomcatExecutionEnvironmentDTO instantiates a new TomcatExecutionEnvironmentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTomcatExecutionEnvironmentDTO() *TomcatExecutionEnvironmentDTO {
	this := TomcatExecutionEnvironmentDTO{}
	return &this
}

// NewTomcatExecutionEnvironmentDTOWithDefaults instantiates a new TomcatExecutionEnvironmentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTomcatExecutionEnvironmentDTOWithDefaults() *TomcatExecutionEnvironmentDTO {
	this := TomcatExecutionEnvironmentDTO{}
	return &this
}

// GetActivations returns the Activations field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetActivations() []ActivationDTO {
	if o == nil || isNil(o.Activations) {
		var ret []ActivationDTO
		return ret
	}
	return o.Activations
}

// GetActivationsOk returns a tuple with the Activations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetActivationsOk() ([]ActivationDTO, bool) {
	if o == nil || isNil(o.Activations) {
    return nil, false
	}
	return o.Activations, true
}

// HasActivations returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasActivations() bool {
	if o != nil && !isNil(o.Activations) {
		return true
	}

	return false
}

// SetActivations gets a reference to the given []ActivationDTO and assigns it to the Activations field.
func (o *TomcatExecutionEnvironmentDTO) SetActivations(v []ActivationDTO) {
	o.Activations = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TomcatExecutionEnvironmentDTO) SetActive(v bool) {
	o.Active = &v
}

// GetBindingLocation returns the BindingLocation field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetBindingLocation() LocationDTO {
	if o == nil || isNil(o.BindingLocation) {
		var ret LocationDTO
		return ret
	}
	return *o.BindingLocation
}

// GetBindingLocationOk returns a tuple with the BindingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetBindingLocationOk() (*LocationDTO, bool) {
	if o == nil || isNil(o.BindingLocation) {
    return nil, false
	}
	return o.BindingLocation, true
}

// HasBindingLocation returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasBindingLocation() bool {
	if o != nil && !isNil(o.BindingLocation) {
		return true
	}

	return false
}

// SetBindingLocation gets a reference to the given LocationDTO and assigns it to the BindingLocation field.
func (o *TomcatExecutionEnvironmentDTO) SetBindingLocation(v LocationDTO) {
	o.BindingLocation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TomcatExecutionEnvironmentDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TomcatExecutionEnvironmentDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetElementId() string {
	if o == nil || isNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetElementIdOk() (*string, bool) {
	if o == nil || isNil(o.ElementId) {
    return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasElementId() bool {
	if o != nil && !isNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *TomcatExecutionEnvironmentDTO) SetElementId(v string) {
	o.ElementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TomcatExecutionEnvironmentDTO) SetId(v int64) {
	o.Id = &v
}

// GetInstallDemoApps returns the InstallDemoApps field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetInstallDemoApps() bool {
	if o == nil || isNil(o.InstallDemoApps) {
		var ret bool
		return ret
	}
	return *o.InstallDemoApps
}

// GetInstallDemoAppsOk returns a tuple with the InstallDemoApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetInstallDemoAppsOk() (*bool, bool) {
	if o == nil || isNil(o.InstallDemoApps) {
    return nil, false
	}
	return o.InstallDemoApps, true
}

// HasInstallDemoApps returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasInstallDemoApps() bool {
	if o != nil && !isNil(o.InstallDemoApps) {
		return true
	}

	return false
}

// SetInstallDemoApps gets a reference to the given bool and assigns it to the InstallDemoApps field.
func (o *TomcatExecutionEnvironmentDTO) SetInstallDemoApps(v bool) {
	o.InstallDemoApps = &v
}

// GetInstallUri returns the InstallUri field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetInstallUri() string {
	if o == nil || isNil(o.InstallUri) {
		var ret string
		return ret
	}
	return *o.InstallUri
}

// GetInstallUriOk returns a tuple with the InstallUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetInstallUriOk() (*string, bool) {
	if o == nil || isNil(o.InstallUri) {
    return nil, false
	}
	return o.InstallUri, true
}

// HasInstallUri returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasInstallUri() bool {
	if o != nil && !isNil(o.InstallUri) {
		return true
	}

	return false
}

// SetInstallUri gets a reference to the given string and assigns it to the InstallUri field.
func (o *TomcatExecutionEnvironmentDTO) SetInstallUri(v string) {
	o.InstallUri = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *TomcatExecutionEnvironmentDTO) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TomcatExecutionEnvironmentDTO) SetName(v string) {
	o.Name = &v
}

// GetOverwriteOriginalSetup returns the OverwriteOriginalSetup field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetOverwriteOriginalSetup() bool {
	if o == nil || isNil(o.OverwriteOriginalSetup) {
		var ret bool
		return ret
	}
	return *o.OverwriteOriginalSetup
}

// GetOverwriteOriginalSetupOk returns a tuple with the OverwriteOriginalSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetOverwriteOriginalSetupOk() (*bool, bool) {
	if o == nil || isNil(o.OverwriteOriginalSetup) {
    return nil, false
	}
	return o.OverwriteOriginalSetup, true
}

// HasOverwriteOriginalSetup returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasOverwriteOriginalSetup() bool {
	if o != nil && !isNil(o.OverwriteOriginalSetup) {
		return true
	}

	return false
}

// SetOverwriteOriginalSetup gets a reference to the given bool and assigns it to the OverwriteOriginalSetup field.
func (o *TomcatExecutionEnvironmentDTO) SetOverwriteOriginalSetup(v bool) {
	o.OverwriteOriginalSetup = &v
}

// GetPlatformId returns the PlatformId field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetPlatformId() string {
	if o == nil || isNil(o.PlatformId) {
		var ret string
		return ret
	}
	return *o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetPlatformIdOk() (*string, bool) {
	if o == nil || isNil(o.PlatformId) {
    return nil, false
	}
	return o.PlatformId, true
}

// HasPlatformId returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasPlatformId() bool {
	if o != nil && !isNil(o.PlatformId) {
		return true
	}

	return false
}

// SetPlatformId gets a reference to the given string and assigns it to the PlatformId field.
func (o *TomcatExecutionEnvironmentDTO) SetPlatformId(v string) {
	o.PlatformId = &v
}

// GetTargetJDK returns the TargetJDK field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetTargetJDK() string {
	if o == nil || isNil(o.TargetJDK) {
		var ret string
		return ret
	}
	return *o.TargetJDK
}

// GetTargetJDKOk returns a tuple with the TargetJDK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetTargetJDKOk() (*string, bool) {
	if o == nil || isNil(o.TargetJDK) {
    return nil, false
	}
	return o.TargetJDK, true
}

// HasTargetJDK returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasTargetJDK() bool {
	if o != nil && !isNil(o.TargetJDK) {
		return true
	}

	return false
}

// SetTargetJDK gets a reference to the given string and assigns it to the TargetJDK field.
func (o *TomcatExecutionEnvironmentDTO) SetTargetJDK(v string) {
	o.TargetJDK = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TomcatExecutionEnvironmentDTO) SetType(v string) {
	o.Type = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetX() float64 {
	if o == nil || isNil(o.X) {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetXOk() (*float64, bool) {
	if o == nil || isNil(o.X) {
    return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *TomcatExecutionEnvironmentDTO) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *TomcatExecutionEnvironmentDTO) GetY() float64 {
	if o == nil || isNil(o.Y) {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TomcatExecutionEnvironmentDTO) GetYOk() (*float64, bool) {
	if o == nil || isNil(o.Y) {
    return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *TomcatExecutionEnvironmentDTO) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *TomcatExecutionEnvironmentDTO) SetY(v float64) {
	o.Y = &v
}

func (o TomcatExecutionEnvironmentDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Activations) {
		toSerialize["activations"] = o.Activations
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.BindingLocation) {
		toSerialize["bindingLocation"] = o.BindingLocation
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.ElementId) {
		toSerialize["elementId"] = o.ElementId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.InstallDemoApps) {
		toSerialize["installDemoApps"] = o.InstallDemoApps
	}
	if !isNil(o.InstallUri) {
		toSerialize["installUri"] = o.InstallUri
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OverwriteOriginalSetup) {
		toSerialize["overwriteOriginalSetup"] = o.OverwriteOriginalSetup
	}
	if !isNil(o.PlatformId) {
		toSerialize["platformId"] = o.PlatformId
	}
	if !isNil(o.TargetJDK) {
		toSerialize["targetJDK"] = o.TargetJDK
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TomcatExecutionEnvironmentDTO) UnmarshalJSON(bytes []byte) (err error) {
	varTomcatExecutionEnvironmentDTO := _TomcatExecutionEnvironmentDTO{}

	if err = json.Unmarshal(bytes, &varTomcatExecutionEnvironmentDTO); err == nil {
		*o = TomcatExecutionEnvironmentDTO(varTomcatExecutionEnvironmentDTO)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activations")
		delete(additionalProperties, "active")
		delete(additionalProperties, "bindingLocation")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "installDemoApps")
		delete(additionalProperties, "installUri")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "overwriteOriginalSetup")
		delete(additionalProperties, "platformId")
		delete(additionalProperties, "targetJDK")
		delete(additionalProperties, "type")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTomcatExecutionEnvironmentDTO struct {
	value *TomcatExecutionEnvironmentDTO
	isSet bool
}

func (v NullableTomcatExecutionEnvironmentDTO) Get() *TomcatExecutionEnvironmentDTO {
	return v.value
}

func (v *NullableTomcatExecutionEnvironmentDTO) Set(val *TomcatExecutionEnvironmentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTomcatExecutionEnvironmentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTomcatExecutionEnvironmentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTomcatExecutionEnvironmentDTO(val *TomcatExecutionEnvironmentDTO) *NullableTomcatExecutionEnvironmentDTO {
	return &NullableTomcatExecutionEnvironmentDTO{value: val, isSet: true}
}

func (v NullableTomcatExecutionEnvironmentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTomcatExecutionEnvironmentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


