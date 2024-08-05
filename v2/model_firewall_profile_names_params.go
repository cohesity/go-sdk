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

// checks if the FirewallProfileNamesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallProfileNamesParams{}

// FirewallProfileNamesParams Specifies the firewall profile names to be removed.
type FirewallProfileNamesParams struct {
	// Specifies the list of profile names to be removed.
	Names []string `json:"names"`
}

type _FirewallProfileNamesParams FirewallProfileNamesParams

// NewFirewallProfileNamesParams instantiates a new FirewallProfileNamesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallProfileNamesParams(names []string) *FirewallProfileNamesParams {
	this := FirewallProfileNamesParams{}
	this.Names = names
	return &this
}

// NewFirewallProfileNamesParamsWithDefaults instantiates a new FirewallProfileNamesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallProfileNamesParamsWithDefaults() *FirewallProfileNamesParams {
	this := FirewallProfileNamesParams{}
	return &this
}

// GetNames returns the Names field value
func (o *FirewallProfileNamesParams) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *FirewallProfileNamesParams) GetNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *FirewallProfileNamesParams) SetNames(v []string) {
	o.Names = v
}

func (o FirewallProfileNamesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallProfileNamesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["names"] = o.Names
	return toSerialize, nil
}

func (o *FirewallProfileNamesParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"names",
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

	varFirewallProfileNamesParams := _FirewallProfileNamesParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFirewallProfileNamesParams)

	if err != nil {
		return err
	}

	*o = FirewallProfileNamesParams(varFirewallProfileNamesParams)

	return err
}

type NullableFirewallProfileNamesParams struct {
	value *FirewallProfileNamesParams
	isSet bool
}

func (v NullableFirewallProfileNamesParams) Get() *FirewallProfileNamesParams {
	return v.value
}

func (v *NullableFirewallProfileNamesParams) Set(val *FirewallProfileNamesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallProfileNamesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallProfileNamesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallProfileNamesParams(val *FirewallProfileNamesParams) *NullableFirewallProfileNamesParams {
	return &NullableFirewallProfileNamesParams{value: val, isSet: true}
}

func (v NullableFirewallProfileNamesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallProfileNamesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


