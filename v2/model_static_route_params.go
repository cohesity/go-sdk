/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StaticRouteParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticRouteParams{}

// StaticRouteParams Specifies the static route parameters.
type StaticRouteParams struct {
	// Specifies a description of the Static Route.
	Description NullableString `json:"description,omitempty"`
	// Specifies the destination network of the Static Route.
	DestinationNetwork NullableString `json:"destinationNetwork"`
	// Specifies the unique identifier for the route.
	Id NullableString `json:"id,omitempty"`
	// Specifies the network interface name to use for communicating with the destination network.
	Interface NullableString `json:"interface,omitempty"`
	// Specifies the network interfaces name to use for communicating with the destination network.
	InterfaceGroup NullableString `json:"interfaceGroup"`
	// Specifies MTU setting per route.
	Mtu NullableInt64 `json:"mtu,omitempty"`
	// Specifies the next hop to the destination network.
	NextHop NullableString `json:"nextHop"`
	// Specifies the network node group to represent a group of nodes.
	NodeGroupName NullableString `json:"nodeGroupName,omitempty"`
}

type _StaticRouteParams StaticRouteParams

// NewStaticRouteParams instantiates a new StaticRouteParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticRouteParams(destinationNetwork NullableString, interfaceGroup NullableString, nextHop NullableString) *StaticRouteParams {
	this := StaticRouteParams{}
	this.DestinationNetwork = destinationNetwork
	this.InterfaceGroup = interfaceGroup
	this.NextHop = nextHop
	return &this
}

// NewStaticRouteParamsWithDefaults instantiates a new StaticRouteParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticRouteParamsWithDefaults() *StaticRouteParams {
	this := StaticRouteParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticRouteParams) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StaticRouteParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StaticRouteParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StaticRouteParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StaticRouteParams) UnsetDescription() {
	o.Description.Unset()
}

// GetDestinationNetwork returns the DestinationNetwork field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StaticRouteParams) GetDestinationNetwork() string {
	if o == nil || o.DestinationNetwork.Get() == nil {
		var ret string
		return ret
	}

	return *o.DestinationNetwork.Get()
}

// GetDestinationNetworkOk returns a tuple with the DestinationNetwork field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetDestinationNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationNetwork.Get(), o.DestinationNetwork.IsSet()
}

// SetDestinationNetwork sets field value
func (o *StaticRouteParams) SetDestinationNetwork(v string) {
	o.DestinationNetwork.Set(&v)
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticRouteParams) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *StaticRouteParams) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *StaticRouteParams) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *StaticRouteParams) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *StaticRouteParams) UnsetId() {
	o.Id.Unset()
}

// GetInterface returns the Interface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticRouteParams) GetInterface() string {
	if o == nil || IsNil(o.Interface.Get()) {
		var ret string
		return ret
	}
	return *o.Interface.Get()
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interface.Get(), o.Interface.IsSet()
}

// HasInterface returns a boolean if a field has been set.
func (o *StaticRouteParams) HasInterface() bool {
	if o != nil && o.Interface.IsSet() {
		return true
	}

	return false
}

// SetInterface gets a reference to the given NullableString and assigns it to the Interface field.
func (o *StaticRouteParams) SetInterface(v string) {
	o.Interface.Set(&v)
}
// SetInterfaceNil sets the value for Interface to be an explicit nil
func (o *StaticRouteParams) SetInterfaceNil() {
	o.Interface.Set(nil)
}

// UnsetInterface ensures that no value is present for Interface, not even an explicit nil
func (o *StaticRouteParams) UnsetInterface() {
	o.Interface.Unset()
}

// GetInterfaceGroup returns the InterfaceGroup field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StaticRouteParams) GetInterfaceGroup() string {
	if o == nil || o.InterfaceGroup.Get() == nil {
		var ret string
		return ret
	}

	return *o.InterfaceGroup.Get()
}

// GetInterfaceGroupOk returns a tuple with the InterfaceGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetInterfaceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterfaceGroup.Get(), o.InterfaceGroup.IsSet()
}

// SetInterfaceGroup sets field value
func (o *StaticRouteParams) SetInterfaceGroup(v string) {
	o.InterfaceGroup.Set(&v)
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticRouteParams) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int64
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetMtuOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *StaticRouteParams) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt64 and assigns it to the Mtu field.
func (o *StaticRouteParams) SetMtu(v int64) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *StaticRouteParams) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *StaticRouteParams) UnsetMtu() {
	o.Mtu.Unset()
}

// GetNextHop returns the NextHop field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StaticRouteParams) GetNextHop() string {
	if o == nil || o.NextHop.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextHop.Get()
}

// GetNextHopOk returns a tuple with the NextHop field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetNextHopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextHop.Get(), o.NextHop.IsSet()
}

// SetNextHop sets field value
func (o *StaticRouteParams) SetNextHop(v string) {
	o.NextHop.Set(&v)
}

// GetNodeGroupName returns the NodeGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StaticRouteParams) GetNodeGroupName() string {
	if o == nil || IsNil(o.NodeGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.NodeGroupName.Get()
}

// GetNodeGroupNameOk returns a tuple with the NodeGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StaticRouteParams) GetNodeGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeGroupName.Get(), o.NodeGroupName.IsSet()
}

// HasNodeGroupName returns a boolean if a field has been set.
func (o *StaticRouteParams) HasNodeGroupName() bool {
	if o != nil && o.NodeGroupName.IsSet() {
		return true
	}

	return false
}

// SetNodeGroupName gets a reference to the given NullableString and assigns it to the NodeGroupName field.
func (o *StaticRouteParams) SetNodeGroupName(v string) {
	o.NodeGroupName.Set(&v)
}
// SetNodeGroupNameNil sets the value for NodeGroupName to be an explicit nil
func (o *StaticRouteParams) SetNodeGroupNameNil() {
	o.NodeGroupName.Set(nil)
}

// UnsetNodeGroupName ensures that no value is present for NodeGroupName, not even an explicit nil
func (o *StaticRouteParams) UnsetNodeGroupName() {
	o.NodeGroupName.Unset()
}

func (o StaticRouteParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticRouteParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["destinationNetwork"] = o.DestinationNetwork.Get()
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Interface.IsSet() {
		toSerialize["interface"] = o.Interface.Get()
	}
	toSerialize["interfaceGroup"] = o.InterfaceGroup.Get()
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	toSerialize["nextHop"] = o.NextHop.Get()
	if o.NodeGroupName.IsSet() {
		toSerialize["nodeGroupName"] = o.NodeGroupName.Get()
	}
	return toSerialize, nil
}

func (o *StaticRouteParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destinationNetwork",
		"interfaceGroup",
		"nextHop",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStaticRouteParams := _StaticRouteParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticRouteParams)

	if err != nil {
		return err
	}

	*o = StaticRouteParams(varStaticRouteParams)

	return err
}

type NullableStaticRouteParams struct {
	value *StaticRouteParams
	isSet bool
}

func (v NullableStaticRouteParams) Get() *StaticRouteParams {
	return v.value
}

func (v *NullableStaticRouteParams) Set(val *StaticRouteParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticRouteParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticRouteParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticRouteParams(val *StaticRouteParams) *NullableStaticRouteParams {
	return &NullableStaticRouteParams{value: val, isSet: true}
}

func (v NullableStaticRouteParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticRouteParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


