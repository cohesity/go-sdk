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

// checks if the CapacityByTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapacityByTier{}

// CapacityByTier CapacityByTier provides the physical capacity in bytes of each storage tier.
type CapacityByTier struct {
	// maxPhysicalCapacityBytesTier is the maximum physical capacity in bytes of the storage tier.
	MaxPhysicalCapacityBytesTier NullableInt64 `json:"maxPhysicalCapacityBytesTier,omitempty"`
	// StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. 'kPCIeSSD' indicates storage tier type of Pci Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Hard Disk Drive. 'kCLOUD' indicates storage tier type of Cloud.
	StorageTier NullableString `json:"storageTier,omitempty"`
}

// NewCapacityByTier instantiates a new CapacityByTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityByTier() *CapacityByTier {
	this := CapacityByTier{}
	return &this
}

// NewCapacityByTierWithDefaults instantiates a new CapacityByTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityByTierWithDefaults() *CapacityByTier {
	this := CapacityByTier{}
	return &this
}

// GetMaxPhysicalCapacityBytesTier returns the MaxPhysicalCapacityBytesTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapacityByTier) GetMaxPhysicalCapacityBytesTier() int64 {
	if o == nil || IsNil(o.MaxPhysicalCapacityBytesTier.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxPhysicalCapacityBytesTier.Get()
}

// GetMaxPhysicalCapacityBytesTierOk returns a tuple with the MaxPhysicalCapacityBytesTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapacityByTier) GetMaxPhysicalCapacityBytesTierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPhysicalCapacityBytesTier.Get(), o.MaxPhysicalCapacityBytesTier.IsSet()
}

// HasMaxPhysicalCapacityBytesTier returns a boolean if a field has been set.
func (o *CapacityByTier) HasMaxPhysicalCapacityBytesTier() bool {
	if o != nil && o.MaxPhysicalCapacityBytesTier.IsSet() {
		return true
	}

	return false
}

// SetMaxPhysicalCapacityBytesTier gets a reference to the given NullableInt64 and assigns it to the MaxPhysicalCapacityBytesTier field.
func (o *CapacityByTier) SetMaxPhysicalCapacityBytesTier(v int64) {
	o.MaxPhysicalCapacityBytesTier.Set(&v)
}
// SetMaxPhysicalCapacityBytesTierNil sets the value for MaxPhysicalCapacityBytesTier to be an explicit nil
func (o *CapacityByTier) SetMaxPhysicalCapacityBytesTierNil() {
	o.MaxPhysicalCapacityBytesTier.Set(nil)
}

// UnsetMaxPhysicalCapacityBytesTier ensures that no value is present for MaxPhysicalCapacityBytesTier, not even an explicit nil
func (o *CapacityByTier) UnsetMaxPhysicalCapacityBytesTier() {
	o.MaxPhysicalCapacityBytesTier.Unset()
}

// GetStorageTier returns the StorageTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapacityByTier) GetStorageTier() string {
	if o == nil || IsNil(o.StorageTier.Get()) {
		var ret string
		return ret
	}
	return *o.StorageTier.Get()
}

// GetStorageTierOk returns a tuple with the StorageTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapacityByTier) GetStorageTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageTier.Get(), o.StorageTier.IsSet()
}

// HasStorageTier returns a boolean if a field has been set.
func (o *CapacityByTier) HasStorageTier() bool {
	if o != nil && o.StorageTier.IsSet() {
		return true
	}

	return false
}

// SetStorageTier gets a reference to the given NullableString and assigns it to the StorageTier field.
func (o *CapacityByTier) SetStorageTier(v string) {
	o.StorageTier.Set(&v)
}
// SetStorageTierNil sets the value for StorageTier to be an explicit nil
func (o *CapacityByTier) SetStorageTierNil() {
	o.StorageTier.Set(nil)
}

// UnsetStorageTier ensures that no value is present for StorageTier, not even an explicit nil
func (o *CapacityByTier) UnsetStorageTier() {
	o.StorageTier.Unset()
}

func (o CapacityByTier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapacityByTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxPhysicalCapacityBytesTier.IsSet() {
		toSerialize["maxPhysicalCapacityBytesTier"] = o.MaxPhysicalCapacityBytesTier.Get()
	}
	if o.StorageTier.IsSet() {
		toSerialize["storageTier"] = o.StorageTier.Get()
	}
	return toSerialize, nil
}

type NullableCapacityByTier struct {
	value *CapacityByTier
	isSet bool
}

func (v NullableCapacityByTier) Get() *CapacityByTier {
	return v.value
}

func (v *NullableCapacityByTier) Set(val *CapacityByTier) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityByTier) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityByTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityByTier(val *CapacityByTier) *NullableCapacityByTier {
	return &NullableCapacityByTier{value: val, isSet: true}
}

func (v NullableCapacityByTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityByTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


