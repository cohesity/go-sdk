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

// checks if the StorageArraySnapshotMaxSpaceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageArraySnapshotMaxSpaceConfig{}

// StorageArraySnapshotMaxSpaceConfig Specifies max space threshold configuration that can used by snapshots to take storage snapshot.
type StorageArraySnapshotMaxSpaceConfig struct {
	// Specifies max space threshold can used by snapshots in percentage in volume/lun to take storage snapshot.
	MaxSnapshotSpacePercentage NullableInt32 `json:"maxSnapshotSpacePercentage,omitempty"`
}

// NewStorageArraySnapshotMaxSpaceConfig instantiates a new StorageArraySnapshotMaxSpaceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageArraySnapshotMaxSpaceConfig() *StorageArraySnapshotMaxSpaceConfig {
	this := StorageArraySnapshotMaxSpaceConfig{}
	return &this
}

// NewStorageArraySnapshotMaxSpaceConfigWithDefaults instantiates a new StorageArraySnapshotMaxSpaceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageArraySnapshotMaxSpaceConfigWithDefaults() *StorageArraySnapshotMaxSpaceConfig {
	this := StorageArraySnapshotMaxSpaceConfig{}
	return &this
}

// GetMaxSnapshotSpacePercentage returns the MaxSnapshotSpacePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageArraySnapshotMaxSpaceConfig) GetMaxSnapshotSpacePercentage() int32 {
	if o == nil || IsNil(o.MaxSnapshotSpacePercentage.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxSnapshotSpacePercentage.Get()
}

// GetMaxSnapshotSpacePercentageOk returns a tuple with the MaxSnapshotSpacePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageArraySnapshotMaxSpaceConfig) GetMaxSnapshotSpacePercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSnapshotSpacePercentage.Get(), o.MaxSnapshotSpacePercentage.IsSet()
}

// HasMaxSnapshotSpacePercentage returns a boolean if a field has been set.
func (o *StorageArraySnapshotMaxSpaceConfig) HasMaxSnapshotSpacePercentage() bool {
	if o != nil && o.MaxSnapshotSpacePercentage.IsSet() {
		return true
	}

	return false
}

// SetMaxSnapshotSpacePercentage gets a reference to the given NullableInt32 and assigns it to the MaxSnapshotSpacePercentage field.
func (o *StorageArraySnapshotMaxSpaceConfig) SetMaxSnapshotSpacePercentage(v int32) {
	o.MaxSnapshotSpacePercentage.Set(&v)
}
// SetMaxSnapshotSpacePercentageNil sets the value for MaxSnapshotSpacePercentage to be an explicit nil
func (o *StorageArraySnapshotMaxSpaceConfig) SetMaxSnapshotSpacePercentageNil() {
	o.MaxSnapshotSpacePercentage.Set(nil)
}

// UnsetMaxSnapshotSpacePercentage ensures that no value is present for MaxSnapshotSpacePercentage, not even an explicit nil
func (o *StorageArraySnapshotMaxSpaceConfig) UnsetMaxSnapshotSpacePercentage() {
	o.MaxSnapshotSpacePercentage.Unset()
}

func (o StorageArraySnapshotMaxSpaceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageArraySnapshotMaxSpaceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxSnapshotSpacePercentage.IsSet() {
		toSerialize["maxSnapshotSpacePercentage"] = o.MaxSnapshotSpacePercentage.Get()
	}
	return toSerialize, nil
}

type NullableStorageArraySnapshotMaxSpaceConfig struct {
	value *StorageArraySnapshotMaxSpaceConfig
	isSet bool
}

func (v NullableStorageArraySnapshotMaxSpaceConfig) Get() *StorageArraySnapshotMaxSpaceConfig {
	return v.value
}

func (v *NullableStorageArraySnapshotMaxSpaceConfig) Set(val *StorageArraySnapshotMaxSpaceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageArraySnapshotMaxSpaceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageArraySnapshotMaxSpaceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageArraySnapshotMaxSpaceConfig(val *StorageArraySnapshotMaxSpaceConfig) *NullableStorageArraySnapshotMaxSpaceConfig {
	return &NullableStorageArraySnapshotMaxSpaceConfig{value: val, isSet: true}
}

func (v NullableStorageArraySnapshotMaxSpaceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageArraySnapshotMaxSpaceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


