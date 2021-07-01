/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// CapacityByTier CapacityByTier provides the physical capacity in bytes of each storage tier.
type CapacityByTier struct {
	// StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. 'kPCIeSSD' indicates storage tier type of Pci Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Solid State Drive. 'kSATAHDD' indicates storage tier type of SATA Hard Disk Drive. 'kCLOUD' indicates storage tier type of Cloud.
	StorageTier NullableString `json:"storageTier,omitempty"`
	// TierMaxPhysicalCapacityBytes is the maximum physical capacity in bytes of the storage tier.
	TierMaxPhysicalCapacityBytes NullableInt64 `json:"tierMaxPhysicalCapacityBytes,omitempty"`
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

// GetStorageTier returns the StorageTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapacityByTier) GetStorageTier() string {
	if o == nil || o.StorageTier.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageTier.Get()
}

// GetStorageTierOk returns a tuple with the StorageTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapacityByTier) GetStorageTierOk() (*string, bool) {
	if o == nil  {
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

// GetTierMaxPhysicalCapacityBytes returns the TierMaxPhysicalCapacityBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapacityByTier) GetTierMaxPhysicalCapacityBytes() int64 {
	if o == nil || o.TierMaxPhysicalCapacityBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TierMaxPhysicalCapacityBytes.Get()
}

// GetTierMaxPhysicalCapacityBytesOk returns a tuple with the TierMaxPhysicalCapacityBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapacityByTier) GetTierMaxPhysicalCapacityBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TierMaxPhysicalCapacityBytes.Get(), o.TierMaxPhysicalCapacityBytes.IsSet()
}

// HasTierMaxPhysicalCapacityBytes returns a boolean if a field has been set.
func (o *CapacityByTier) HasTierMaxPhysicalCapacityBytes() bool {
	if o != nil && o.TierMaxPhysicalCapacityBytes.IsSet() {
		return true
	}

	return false
}

// SetTierMaxPhysicalCapacityBytes gets a reference to the given NullableInt64 and assigns it to the TierMaxPhysicalCapacityBytes field.
func (o *CapacityByTier) SetTierMaxPhysicalCapacityBytes(v int64) {
	o.TierMaxPhysicalCapacityBytes.Set(&v)
}
// SetTierMaxPhysicalCapacityBytesNil sets the value for TierMaxPhysicalCapacityBytes to be an explicit nil
func (o *CapacityByTier) SetTierMaxPhysicalCapacityBytesNil() {
	o.TierMaxPhysicalCapacityBytes.Set(nil)
}

// UnsetTierMaxPhysicalCapacityBytes ensures that no value is present for TierMaxPhysicalCapacityBytes, not even an explicit nil
func (o *CapacityByTier) UnsetTierMaxPhysicalCapacityBytes() {
	o.TierMaxPhysicalCapacityBytes.Unset()
}

func (o CapacityByTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageTier.IsSet() {
		toSerialize["storageTier"] = o.StorageTier.Get()
	}
	if o.TierMaxPhysicalCapacityBytes.IsSet() {
		toSerialize["tierMaxPhysicalCapacityBytes"] = o.TierMaxPhysicalCapacityBytes.Get()
	}
	return json.Marshal(toSerialize)
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


