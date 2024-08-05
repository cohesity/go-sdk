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

// checks if the CountByTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountByTier{}

// CountByTier CountByTier provides the disk count of each storage tier.
type CountByTier struct {
	// DiskCount is the disk number of the storage tier.
	DiskCount NullableInt64 `json:"diskCount,omitempty"`
	// StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. 'kPCIeSSD' indicates storage tier type of Pci Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Hard Disk Drive. 'kCLOUD' indicates storage tier type of Cloud.
	StorageTier NullableString `json:"storageTier,omitempty"`
}

// NewCountByTier instantiates a new CountByTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountByTier() *CountByTier {
	this := CountByTier{}
	return &this
}

// NewCountByTierWithDefaults instantiates a new CountByTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountByTierWithDefaults() *CountByTier {
	this := CountByTier{}
	return &this
}

// GetDiskCount returns the DiskCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountByTier) GetDiskCount() int64 {
	if o == nil || IsNil(o.DiskCount.Get()) {
		var ret int64
		return ret
	}
	return *o.DiskCount.Get()
}

// GetDiskCountOk returns a tuple with the DiskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountByTier) GetDiskCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskCount.Get(), o.DiskCount.IsSet()
}

// HasDiskCount returns a boolean if a field has been set.
func (o *CountByTier) HasDiskCount() bool {
	if o != nil && o.DiskCount.IsSet() {
		return true
	}

	return false
}

// SetDiskCount gets a reference to the given NullableInt64 and assigns it to the DiskCount field.
func (o *CountByTier) SetDiskCount(v int64) {
	o.DiskCount.Set(&v)
}
// SetDiskCountNil sets the value for DiskCount to be an explicit nil
func (o *CountByTier) SetDiskCountNil() {
	o.DiskCount.Set(nil)
}

// UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
func (o *CountByTier) UnsetDiskCount() {
	o.DiskCount.Unset()
}

// GetStorageTier returns the StorageTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountByTier) GetStorageTier() string {
	if o == nil || IsNil(o.StorageTier.Get()) {
		var ret string
		return ret
	}
	return *o.StorageTier.Get()
}

// GetStorageTierOk returns a tuple with the StorageTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountByTier) GetStorageTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageTier.Get(), o.StorageTier.IsSet()
}

// HasStorageTier returns a boolean if a field has been set.
func (o *CountByTier) HasStorageTier() bool {
	if o != nil && o.StorageTier.IsSet() {
		return true
	}

	return false
}

// SetStorageTier gets a reference to the given NullableString and assigns it to the StorageTier field.
func (o *CountByTier) SetStorageTier(v string) {
	o.StorageTier.Set(&v)
}
// SetStorageTierNil sets the value for StorageTier to be an explicit nil
func (o *CountByTier) SetStorageTierNil() {
	o.StorageTier.Set(nil)
}

// UnsetStorageTier ensures that no value is present for StorageTier, not even an explicit nil
func (o *CountByTier) UnsetStorageTier() {
	o.StorageTier.Unset()
}

func (o CountByTier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountByTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskCount.IsSet() {
		toSerialize["diskCount"] = o.DiskCount.Get()
	}
	if o.StorageTier.IsSet() {
		toSerialize["storageTier"] = o.StorageTier.Get()
	}
	return toSerialize, nil
}

type NullableCountByTier struct {
	value *CountByTier
	isSet bool
}

func (v NullableCountByTier) Get() *CountByTier {
	return v.value
}

func (v *NullableCountByTier) Set(val *CountByTier) {
	v.value = val
	v.isSet = true
}

func (v NullableCountByTier) IsSet() bool {
	return v.isSet
}

func (v *NullableCountByTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountByTier(val *CountByTier) *NullableCountByTier {
	return &NullableCountByTier{value: val, isSet: true}
}

func (v NullableCountByTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountByTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


