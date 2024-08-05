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

// checks if the VmwareConnectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmwareConnectionParams{}

// VmwareConnectionParams Specifies the parameters to connect to a seed node and fetch information from its config file.
type VmwareConnectionParams struct {
	// Specifies the VMware Source type.
	Type NullableString `json:"type"`
	VcdParams *VcdConnectionParams `json:"vcdParams,omitempty"`
}

type _VmwareConnectionParams VmwareConnectionParams

// NewVmwareConnectionParams instantiates a new VmwareConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareConnectionParams(type_ NullableString) *VmwareConnectionParams {
	this := VmwareConnectionParams{}
	this.Type = type_
	return &this
}

// NewVmwareConnectionParamsWithDefaults instantiates a new VmwareConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareConnectionParamsWithDefaults() *VmwareConnectionParams {
	this := VmwareConnectionParams{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VmwareConnectionParams) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareConnectionParams) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *VmwareConnectionParams) SetType(v string) {
	o.Type.Set(&v)
}

// GetVcdParams returns the VcdParams field value if set, zero value otherwise.
func (o *VmwareConnectionParams) GetVcdParams() VcdConnectionParams {
	if o == nil || IsNil(o.VcdParams) {
		var ret VcdConnectionParams
		return ret
	}
	return *o.VcdParams
}

// GetVcdParamsOk returns a tuple with the VcdParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareConnectionParams) GetVcdParamsOk() (*VcdConnectionParams, bool) {
	if o == nil || IsNil(o.VcdParams) {
		return nil, false
	}
	return o.VcdParams, true
}

// HasVcdParams returns a boolean if a field has been set.
func (o *VmwareConnectionParams) HasVcdParams() bool {
	if o != nil && !IsNil(o.VcdParams) {
		return true
	}

	return false
}

// SetVcdParams gets a reference to the given VcdConnectionParams and assigns it to the VcdParams field.
func (o *VmwareConnectionParams) SetVcdParams(v VcdConnectionParams) {
	o.VcdParams = &v
}

func (o VmwareConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmwareConnectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type.Get()
	if !IsNil(o.VcdParams) {
		toSerialize["vcdParams"] = o.VcdParams
	}
	return toSerialize, nil
}

func (o *VmwareConnectionParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varVmwareConnectionParams := _VmwareConnectionParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmwareConnectionParams)

	if err != nil {
		return err
	}

	*o = VmwareConnectionParams(varVmwareConnectionParams)

	return err
}

type NullableVmwareConnectionParams struct {
	value *VmwareConnectionParams
	isSet bool
}

func (v NullableVmwareConnectionParams) Get() *VmwareConnectionParams {
	return v.value
}

func (v *NullableVmwareConnectionParams) Set(val *VmwareConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareConnectionParams(val *VmwareConnectionParams) *NullableVmwareConnectionParams {
	return &NullableVmwareConnectionParams{value: val, isSet: true}
}

func (v NullableVmwareConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


