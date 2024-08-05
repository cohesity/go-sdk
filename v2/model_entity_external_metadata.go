/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the EntityExternalMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityExternalMetadata{}

// EntityExternalMetadata Specifies the External metadata of an entity
type EntityExternalMetadata struct {
	MaintenanceModeConfig *MaintenanceModeConfig `json:"maintenanceModeConfig,omitempty"`
}

// NewEntityExternalMetadata instantiates a new EntityExternalMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityExternalMetadata() *EntityExternalMetadata {
	this := EntityExternalMetadata{}
	return &this
}

// NewEntityExternalMetadataWithDefaults instantiates a new EntityExternalMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityExternalMetadataWithDefaults() *EntityExternalMetadata {
	this := EntityExternalMetadata{}
	return &this
}

// GetMaintenanceModeConfig returns the MaintenanceModeConfig field value if set, zero value otherwise.
func (o *EntityExternalMetadata) GetMaintenanceModeConfig() MaintenanceModeConfig {
	if o == nil || IsNil(o.MaintenanceModeConfig) {
		var ret MaintenanceModeConfig
		return ret
	}
	return *o.MaintenanceModeConfig
}

// GetMaintenanceModeConfigOk returns a tuple with the MaintenanceModeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityExternalMetadata) GetMaintenanceModeConfigOk() (*MaintenanceModeConfig, bool) {
	if o == nil || IsNil(o.MaintenanceModeConfig) {
		return nil, false
	}
	return o.MaintenanceModeConfig, true
}

// HasMaintenanceModeConfig returns a boolean if a field has been set.
func (o *EntityExternalMetadata) HasMaintenanceModeConfig() bool {
	if o != nil && !IsNil(o.MaintenanceModeConfig) {
		return true
	}

	return false
}

// SetMaintenanceModeConfig gets a reference to the given MaintenanceModeConfig and assigns it to the MaintenanceModeConfig field.
func (o *EntityExternalMetadata) SetMaintenanceModeConfig(v MaintenanceModeConfig) {
	o.MaintenanceModeConfig = &v
}

func (o EntityExternalMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityExternalMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaintenanceModeConfig) {
		toSerialize["maintenanceModeConfig"] = o.MaintenanceModeConfig
	}
	return toSerialize, nil
}

type NullableEntityExternalMetadata struct {
	value *EntityExternalMetadata
	isSet bool
}

func (v NullableEntityExternalMetadata) Get() *EntityExternalMetadata {
	return v.value
}

func (v *NullableEntityExternalMetadata) Set(val *EntityExternalMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityExternalMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityExternalMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityExternalMetadata(val *EntityExternalMetadata) *NullableEntityExternalMetadata {
	return &NullableEntityExternalMetadata{value: val, isSet: true}
}

func (v NullableEntityExternalMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityExternalMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


