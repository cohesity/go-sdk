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

// checks if the PrivateNetworkInfoRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateNetworkInfoRegion{}

// PrivateNetworkInfoRegion Specifies the region of the virtual network.
type PrivateNetworkInfoRegion struct {
	// Specifies the id of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
}

type _PrivateNetworkInfoRegion PrivateNetworkInfoRegion

// NewPrivateNetworkInfoRegion instantiates a new PrivateNetworkInfoRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkInfoRegion(id NullableInt64) *PrivateNetworkInfoRegion {
	this := PrivateNetworkInfoRegion{}
	this.Id = id
	return &this
}

// NewPrivateNetworkInfoRegionWithDefaults instantiates a new PrivateNetworkInfoRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkInfoRegionWithDefaults() *PrivateNetworkInfoRegion {
	this := PrivateNetworkInfoRegion{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *PrivateNetworkInfoRegion) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateNetworkInfoRegion) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *PrivateNetworkInfoRegion) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateNetworkInfoRegion) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateNetworkInfoRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrivateNetworkInfoRegion) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrivateNetworkInfoRegion) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PrivateNetworkInfoRegion) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrivateNetworkInfoRegion) UnsetName() {
	o.Name.Unset()
}

func (o PrivateNetworkInfoRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateNetworkInfoRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *PrivateNetworkInfoRegion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPrivateNetworkInfoRegion := _PrivateNetworkInfoRegion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateNetworkInfoRegion)

	if err != nil {
		return err
	}

	*o = PrivateNetworkInfoRegion(varPrivateNetworkInfoRegion)

	return err
}

type NullablePrivateNetworkInfoRegion struct {
	value *PrivateNetworkInfoRegion
	isSet bool
}

func (v NullablePrivateNetworkInfoRegion) Get() *PrivateNetworkInfoRegion {
	return v.value
}

func (v *NullablePrivateNetworkInfoRegion) Set(val *PrivateNetworkInfoRegion) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkInfoRegion) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkInfoRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkInfoRegion(val *PrivateNetworkInfoRegion) *NullablePrivateNetworkInfoRegion {
	return &NullablePrivateNetworkInfoRegion{value: val, isSet: true}
}

func (v NullablePrivateNetworkInfoRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkInfoRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


