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

// NodeNetworkInterfaces Specifies the network interfaces present on a specific Node in a Cluster.
type NodeNetworkInterfaces struct {
	// Specifies the serial number of Chassis.
	ChassisSerial NullableString `json:"chassisSerial,omitempty"`
	// Specifies the list of network interfaces present on this Node.
	Interfaces []NetworkInterface `json:"interfaces,omitempty"`
	// Specifies an optional message describing the result of the request pertaining to this Node.
	Message NullableString `json:"message,omitempty"`
	// Specifies the ID of the Node.
	NodeId NullableInt64 `json:"nodeId,omitempty"`
	// Specifies the slot number the Node is located in.
	Slot NullableInt64 `json:"slot,omitempty"`
}

// NewNodeNetworkInterfaces instantiates a new NodeNetworkInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeNetworkInterfaces() *NodeNetworkInterfaces {
	this := NodeNetworkInterfaces{}
	return &this
}

// NewNodeNetworkInterfacesWithDefaults instantiates a new NodeNetworkInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeNetworkInterfacesWithDefaults() *NodeNetworkInterfaces {
	this := NodeNetworkInterfaces{}
	return &this
}

// GetChassisSerial returns the ChassisSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeNetworkInterfaces) GetChassisSerial() string {
	if o == nil || o.ChassisSerial.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChassisSerial.Get()
}

// GetChassisSerialOk returns a tuple with the ChassisSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInterfaces) GetChassisSerialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChassisSerial.Get(), o.ChassisSerial.IsSet()
}

// HasChassisSerial returns a boolean if a field has been set.
func (o *NodeNetworkInterfaces) HasChassisSerial() bool {
	if o != nil && o.ChassisSerial.IsSet() {
		return true
	}

	return false
}

// SetChassisSerial gets a reference to the given NullableString and assigns it to the ChassisSerial field.
func (o *NodeNetworkInterfaces) SetChassisSerial(v string) {
	o.ChassisSerial.Set(&v)
}
// SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil
func (o *NodeNetworkInterfaces) SetChassisSerialNil() {
	o.ChassisSerial.Set(nil)
}

// UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
func (o *NodeNetworkInterfaces) UnsetChassisSerial() {
	o.ChassisSerial.Unset()
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeNetworkInterfaces) GetInterfaces() []NetworkInterface {
	if o == nil  {
		var ret []NetworkInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInterfaces) GetInterfacesOk() (*[]NetworkInterface, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return &o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *NodeNetworkInterfaces) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []NetworkInterface and assigns it to the Interfaces field.
func (o *NodeNetworkInterfaces) SetInterfaces(v []NetworkInterface) {
	o.Interfaces = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeNetworkInterfaces) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInterfaces) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *NodeNetworkInterfaces) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *NodeNetworkInterfaces) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *NodeNetworkInterfaces) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *NodeNetworkInterfaces) UnsetMessage() {
	o.Message.Unset()
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeNetworkInterfaces) GetNodeId() int64 {
	if o == nil || o.NodeId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInterfaces) GetNodeIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *NodeNetworkInterfaces) HasNodeId() bool {
	if o != nil && o.NodeId.IsSet() {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given NullableInt64 and assigns it to the NodeId field.
func (o *NodeNetworkInterfaces) SetNodeId(v int64) {
	o.NodeId.Set(&v)
}
// SetNodeIdNil sets the value for NodeId to be an explicit nil
func (o *NodeNetworkInterfaces) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
func (o *NodeNetworkInterfaces) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetSlot returns the Slot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeNetworkInterfaces) GetSlot() int64 {
	if o == nil || o.Slot.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Slot.Get()
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInterfaces) GetSlotOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Slot.Get(), o.Slot.IsSet()
}

// HasSlot returns a boolean if a field has been set.
func (o *NodeNetworkInterfaces) HasSlot() bool {
	if o != nil && o.Slot.IsSet() {
		return true
	}

	return false
}

// SetSlot gets a reference to the given NullableInt64 and assigns it to the Slot field.
func (o *NodeNetworkInterfaces) SetSlot(v int64) {
	o.Slot.Set(&v)
}
// SetSlotNil sets the value for Slot to be an explicit nil
func (o *NodeNetworkInterfaces) SetSlotNil() {
	o.Slot.Set(nil)
}

// UnsetSlot ensures that no value is present for Slot, not even an explicit nil
func (o *NodeNetworkInterfaces) UnsetSlot() {
	o.Slot.Unset()
}

func (o NodeNetworkInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChassisSerial.IsSet() {
		toSerialize["chassisSerial"] = o.ChassisSerial.Get()
	}
	if o.Interfaces != nil {
		toSerialize["interfaces"] = o.Interfaces
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.Slot.IsSet() {
		toSerialize["slot"] = o.Slot.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNodeNetworkInterfaces struct {
	value *NodeNetworkInterfaces
	isSet bool
}

func (v NullableNodeNetworkInterfaces) Get() *NodeNetworkInterfaces {
	return v.value
}

func (v *NullableNodeNetworkInterfaces) Set(val *NodeNetworkInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeNetworkInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeNetworkInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeNetworkInterfaces(val *NodeNetworkInterfaces) *NullableNodeNetworkInterfaces {
	return &NullableNodeNetworkInterfaces{value: val, isSet: true}
}

func (v NullableNodeNetworkInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeNetworkInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


