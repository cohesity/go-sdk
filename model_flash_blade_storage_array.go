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

// FlashBladeStorageArray Specifies information about a Pure Storage FlashBlade Array.
type FlashBladeStorageArray struct {
	// Specifies the total capacity in bytes of the Pure Storage FlashBlade Array.
	CapacityBytes NullableInt64 `json:"capacityBytes,omitempty"`
	// Specifies a unique id of a Pure Storage FlashBlade Array. The id is unique across Cohesity Clusters.
	Id NullableString `json:"id,omitempty"`
	// Specifies the network interfaces of the Pure Storage FlashBlade Array.
	Networks []FlashBladeNetworkInterface `json:"networks,omitempty"`
	// Specifies the space used for physical data in bytes.
	PhysicalUsedBytes NullableInt64 `json:"physicalUsedBytes,omitempty"`
	// Specifies the revision of the Pure Storage FlashBlade software.
	Revision NullableString `json:"revision,omitempty"`
	// Specifies the software version running on the Pure Storage FlashBlade Array.
	Version NullableString `json:"version,omitempty"`
}

// NewFlashBladeStorageArray instantiates a new FlashBladeStorageArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlashBladeStorageArray() *FlashBladeStorageArray {
	this := FlashBladeStorageArray{}
	return &this
}

// NewFlashBladeStorageArrayWithDefaults instantiates a new FlashBladeStorageArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlashBladeStorageArrayWithDefaults() *FlashBladeStorageArray {
	this := FlashBladeStorageArray{}
	return &this
}

// GetCapacityBytes returns the CapacityBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetCapacityBytes() int64 {
	if o == nil || o.CapacityBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CapacityBytes.Get()
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetCapacityBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CapacityBytes.Get(), o.CapacityBytes.IsSet()
}

// HasCapacityBytes returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasCapacityBytes() bool {
	if o != nil && o.CapacityBytes.IsSet() {
		return true
	}

	return false
}

// SetCapacityBytes gets a reference to the given NullableInt64 and assigns it to the CapacityBytes field.
func (o *FlashBladeStorageArray) SetCapacityBytes(v int64) {
	o.CapacityBytes.Set(&v)
}
// SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil
func (o *FlashBladeStorageArray) SetCapacityBytesNil() {
	o.CapacityBytes.Set(nil)
}

// UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
func (o *FlashBladeStorageArray) UnsetCapacityBytes() {
	o.CapacityBytes.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *FlashBladeStorageArray) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *FlashBladeStorageArray) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FlashBladeStorageArray) UnsetId() {
	o.Id.Unset()
}

// GetNetworks returns the Networks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetNetworks() []FlashBladeNetworkInterface {
	if o == nil  {
		var ret []FlashBladeNetworkInterface
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetNetworksOk() (*[]FlashBladeNetworkInterface, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return &o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []FlashBladeNetworkInterface and assigns it to the Networks field.
func (o *FlashBladeStorageArray) SetNetworks(v []FlashBladeNetworkInterface) {
	o.Networks = v
}

// GetPhysicalUsedBytes returns the PhysicalUsedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetPhysicalUsedBytes() int64 {
	if o == nil || o.PhysicalUsedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PhysicalUsedBytes.Get()
}

// GetPhysicalUsedBytesOk returns a tuple with the PhysicalUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetPhysicalUsedBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhysicalUsedBytes.Get(), o.PhysicalUsedBytes.IsSet()
}

// HasPhysicalUsedBytes returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasPhysicalUsedBytes() bool {
	if o != nil && o.PhysicalUsedBytes.IsSet() {
		return true
	}

	return false
}

// SetPhysicalUsedBytes gets a reference to the given NullableInt64 and assigns it to the PhysicalUsedBytes field.
func (o *FlashBladeStorageArray) SetPhysicalUsedBytes(v int64) {
	o.PhysicalUsedBytes.Set(&v)
}
// SetPhysicalUsedBytesNil sets the value for PhysicalUsedBytes to be an explicit nil
func (o *FlashBladeStorageArray) SetPhysicalUsedBytesNil() {
	o.PhysicalUsedBytes.Set(nil)
}

// UnsetPhysicalUsedBytes ensures that no value is present for PhysicalUsedBytes, not even an explicit nil
func (o *FlashBladeStorageArray) UnsetPhysicalUsedBytes() {
	o.PhysicalUsedBytes.Unset()
}

// GetRevision returns the Revision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetRevision() string {
	if o == nil || o.Revision.Get() == nil {
		var ret string
		return ret
	}
	return *o.Revision.Get()
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetRevisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Revision.Get(), o.Revision.IsSet()
}

// HasRevision returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasRevision() bool {
	if o != nil && o.Revision.IsSet() {
		return true
	}

	return false
}

// SetRevision gets a reference to the given NullableString and assigns it to the Revision field.
func (o *FlashBladeStorageArray) SetRevision(v string) {
	o.Revision.Set(&v)
}
// SetRevisionNil sets the value for Revision to be an explicit nil
func (o *FlashBladeStorageArray) SetRevisionNil() {
	o.Revision.Set(nil)
}

// UnsetRevision ensures that no value is present for Revision, not even an explicit nil
func (o *FlashBladeStorageArray) UnsetRevision() {
	o.Revision.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashBladeStorageArray) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashBladeStorageArray) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *FlashBladeStorageArray) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *FlashBladeStorageArray) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *FlashBladeStorageArray) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *FlashBladeStorageArray) UnsetVersion() {
	o.Version.Unset()
}

func (o FlashBladeStorageArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityBytes.IsSet() {
		toSerialize["capacityBytes"] = o.CapacityBytes.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.PhysicalUsedBytes.IsSet() {
		toSerialize["physicalUsedBytes"] = o.PhysicalUsedBytes.Get()
	}
	if o.Revision.IsSet() {
		toSerialize["revision"] = o.Revision.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFlashBladeStorageArray struct {
	value *FlashBladeStorageArray
	isSet bool
}

func (v NullableFlashBladeStorageArray) Get() *FlashBladeStorageArray {
	return v.value
}

func (v *NullableFlashBladeStorageArray) Set(val *FlashBladeStorageArray) {
	v.value = val
	v.isSet = true
}

func (v NullableFlashBladeStorageArray) IsSet() bool {
	return v.isSet
}

func (v *NullableFlashBladeStorageArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlashBladeStorageArray(val *FlashBladeStorageArray) *NullableFlashBladeStorageArray {
	return &NullableFlashBladeStorageArray{value: val, isSet: true}
}

func (v NullableFlashBladeStorageArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlashBladeStorageArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


