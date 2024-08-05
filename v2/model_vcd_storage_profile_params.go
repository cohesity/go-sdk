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

// checks if the VcdStorageProfileParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VcdStorageProfileParams{}

// VcdStorageProfileParams Specifies the parameters of a VCD storage profile.
type VcdStorageProfileParams struct {
	// Specifies the name of the storage profile.
	Name NullableString `json:"name,omitempty"`
	// Specifies the UUID assigned by the VCD to the storage profile.
	VcdUuid NullableString `json:"vcdUuid"`
}

type _VcdStorageProfileParams VcdStorageProfileParams

// NewVcdStorageProfileParams instantiates a new VcdStorageProfileParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcdStorageProfileParams(vcdUuid NullableString) *VcdStorageProfileParams {
	this := VcdStorageProfileParams{}
	this.VcdUuid = vcdUuid
	return &this
}

// NewVcdStorageProfileParamsWithDefaults instantiates a new VcdStorageProfileParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcdStorageProfileParamsWithDefaults() *VcdStorageProfileParams {
	this := VcdStorageProfileParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VcdStorageProfileParams) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VcdStorageProfileParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VcdStorageProfileParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VcdStorageProfileParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VcdStorageProfileParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VcdStorageProfileParams) UnsetName() {
	o.Name.Unset()
}

// GetVcdUuid returns the VcdUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VcdStorageProfileParams) GetVcdUuid() string {
	if o == nil || o.VcdUuid.Get() == nil {
		var ret string
		return ret
	}

	return *o.VcdUuid.Get()
}

// GetVcdUuidOk returns a tuple with the VcdUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VcdStorageProfileParams) GetVcdUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcdUuid.Get(), o.VcdUuid.IsSet()
}

// SetVcdUuid sets field value
func (o *VcdStorageProfileParams) SetVcdUuid(v string) {
	o.VcdUuid.Set(&v)
}

func (o VcdStorageProfileParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VcdStorageProfileParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["vcdUuid"] = o.VcdUuid.Get()
	return toSerialize, nil
}

func (o *VcdStorageProfileParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vcdUuid",
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

	varVcdStorageProfileParams := _VcdStorageProfileParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVcdStorageProfileParams)

	if err != nil {
		return err
	}

	*o = VcdStorageProfileParams(varVcdStorageProfileParams)

	return err
}

type NullableVcdStorageProfileParams struct {
	value *VcdStorageProfileParams
	isSet bool
}

func (v NullableVcdStorageProfileParams) Get() *VcdStorageProfileParams {
	return v.value
}

func (v *NullableVcdStorageProfileParams) Set(val *VcdStorageProfileParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVcdStorageProfileParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVcdStorageProfileParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcdStorageProfileParams(val *VcdStorageProfileParams) *NullableVcdStorageProfileParams {
	return &NullableVcdStorageProfileParams{value: val, isSet: true}
}

func (v NullableVcdStorageProfileParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcdStorageProfileParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


