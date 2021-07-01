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

// DeviceTreeDetails Specifies a logical volume stored as a tree where the leaves are the blocks of partitions and intermediate nodes are assembled by combining nodes using one of the following modes: linear layout, striped, mirrored, RAID etc. A deviceTree is a block device formed by combining one or more Devices using a combining strategy.
type DeviceTreeDetails struct {
	// Specifies how to combine the children of this node. The combining strategy for child devices. Some of these strategies imply constraint on the number of child devices. e.g. RAID5 will have 5 children. 'LINEAR' indicates children are juxtaposed to form this device. 'STRIPE' indicates children are striped. 'MIRROR' indicates children are mirrored. 'RAID5' 'RAID6' 'ZERO' 'THIN' 'THINPOOL' 'SNAPSHOT' 'CACHE' 'CACHEPOOL'
	CombineMethod NullableString `json:"combineMethod,omitempty"`
	// Specifies the length of this device. This number should match the length that is calculated from the children and combining method.
	DeviceLength NullableInt64 `json:"deviceLength,omitempty"`
	// Specifies the children of this node in the device tree.
	DeviceNodes []DeviceNode `json:"deviceNodes,omitempty"`
	// Specifies the size of the striped data if the data is striped.
	StripeSize NullableInt32 `json:"stripeSize,omitempty"`
}

// NewDeviceTreeDetails instantiates a new DeviceTreeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTreeDetails() *DeviceTreeDetails {
	this := DeviceTreeDetails{}
	return &this
}

// NewDeviceTreeDetailsWithDefaults instantiates a new DeviceTreeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTreeDetailsWithDefaults() *DeviceTreeDetails {
	this := DeviceTreeDetails{}
	return &this
}

// GetCombineMethod returns the CombineMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTreeDetails) GetCombineMethod() string {
	if o == nil || o.CombineMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.CombineMethod.Get()
}

// GetCombineMethodOk returns a tuple with the CombineMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTreeDetails) GetCombineMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CombineMethod.Get(), o.CombineMethod.IsSet()
}

// HasCombineMethod returns a boolean if a field has been set.
func (o *DeviceTreeDetails) HasCombineMethod() bool {
	if o != nil && o.CombineMethod.IsSet() {
		return true
	}

	return false
}

// SetCombineMethod gets a reference to the given NullableString and assigns it to the CombineMethod field.
func (o *DeviceTreeDetails) SetCombineMethod(v string) {
	o.CombineMethod.Set(&v)
}
// SetCombineMethodNil sets the value for CombineMethod to be an explicit nil
func (o *DeviceTreeDetails) SetCombineMethodNil() {
	o.CombineMethod.Set(nil)
}

// UnsetCombineMethod ensures that no value is present for CombineMethod, not even an explicit nil
func (o *DeviceTreeDetails) UnsetCombineMethod() {
	o.CombineMethod.Unset()
}

// GetDeviceLength returns the DeviceLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTreeDetails) GetDeviceLength() int64 {
	if o == nil || o.DeviceLength.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DeviceLength.Get()
}

// GetDeviceLengthOk returns a tuple with the DeviceLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTreeDetails) GetDeviceLengthOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceLength.Get(), o.DeviceLength.IsSet()
}

// HasDeviceLength returns a boolean if a field has been set.
func (o *DeviceTreeDetails) HasDeviceLength() bool {
	if o != nil && o.DeviceLength.IsSet() {
		return true
	}

	return false
}

// SetDeviceLength gets a reference to the given NullableInt64 and assigns it to the DeviceLength field.
func (o *DeviceTreeDetails) SetDeviceLength(v int64) {
	o.DeviceLength.Set(&v)
}
// SetDeviceLengthNil sets the value for DeviceLength to be an explicit nil
func (o *DeviceTreeDetails) SetDeviceLengthNil() {
	o.DeviceLength.Set(nil)
}

// UnsetDeviceLength ensures that no value is present for DeviceLength, not even an explicit nil
func (o *DeviceTreeDetails) UnsetDeviceLength() {
	o.DeviceLength.Unset()
}

// GetDeviceNodes returns the DeviceNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTreeDetails) GetDeviceNodes() []DeviceNode {
	if o == nil  {
		var ret []DeviceNode
		return ret
	}
	return o.DeviceNodes
}

// GetDeviceNodesOk returns a tuple with the DeviceNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTreeDetails) GetDeviceNodesOk() (*[]DeviceNode, bool) {
	if o == nil || o.DeviceNodes == nil {
		return nil, false
	}
	return &o.DeviceNodes, true
}

// HasDeviceNodes returns a boolean if a field has been set.
func (o *DeviceTreeDetails) HasDeviceNodes() bool {
	if o != nil && o.DeviceNodes != nil {
		return true
	}

	return false
}

// SetDeviceNodes gets a reference to the given []DeviceNode and assigns it to the DeviceNodes field.
func (o *DeviceTreeDetails) SetDeviceNodes(v []DeviceNode) {
	o.DeviceNodes = v
}

// GetStripeSize returns the StripeSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTreeDetails) GetStripeSize() int32 {
	if o == nil || o.StripeSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StripeSize.Get()
}

// GetStripeSizeOk returns a tuple with the StripeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTreeDetails) GetStripeSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StripeSize.Get(), o.StripeSize.IsSet()
}

// HasStripeSize returns a boolean if a field has been set.
func (o *DeviceTreeDetails) HasStripeSize() bool {
	if o != nil && o.StripeSize.IsSet() {
		return true
	}

	return false
}

// SetStripeSize gets a reference to the given NullableInt32 and assigns it to the StripeSize field.
func (o *DeviceTreeDetails) SetStripeSize(v int32) {
	o.StripeSize.Set(&v)
}
// SetStripeSizeNil sets the value for StripeSize to be an explicit nil
func (o *DeviceTreeDetails) SetStripeSizeNil() {
	o.StripeSize.Set(nil)
}

// UnsetStripeSize ensures that no value is present for StripeSize, not even an explicit nil
func (o *DeviceTreeDetails) UnsetStripeSize() {
	o.StripeSize.Unset()
}

func (o DeviceTreeDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CombineMethod.IsSet() {
		toSerialize["combineMethod"] = o.CombineMethod.Get()
	}
	if o.DeviceLength.IsSet() {
		toSerialize["deviceLength"] = o.DeviceLength.Get()
	}
	if o.DeviceNodes != nil {
		toSerialize["deviceNodes"] = o.DeviceNodes
	}
	if o.StripeSize.IsSet() {
		toSerialize["stripeSize"] = o.StripeSize.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceTreeDetails struct {
	value *DeviceTreeDetails
	isSet bool
}

func (v NullableDeviceTreeDetails) Get() *DeviceTreeDetails {
	return v.value
}

func (v *NullableDeviceTreeDetails) Set(val *DeviceTreeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTreeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTreeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTreeDetails(val *DeviceTreeDetails) *NullableDeviceTreeDetails {
	return &NullableDeviceTreeDetails{value: val, isSet: true}
}

func (v NullableDeviceTreeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTreeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


