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

// checks if the SfdcObjectProtectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SfdcObjectProtectionParams{}

// SfdcObjectProtectionParams Specifies the parameters that are specific to Sfdc Object Protection.
type SfdcObjectProtectionParams struct {
	// Specifies the objects to be included in the Object Protection.
	Objects []SfdcObjectProtectionObjectParams `json:"objects"`
}

type _SfdcObjectProtectionParams SfdcObjectProtectionParams

// NewSfdcObjectProtectionParams instantiates a new SfdcObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSfdcObjectProtectionParams(objects []SfdcObjectProtectionObjectParams) *SfdcObjectProtectionParams {
	this := SfdcObjectProtectionParams{}
	this.Objects = objects
	return &this
}

// NewSfdcObjectProtectionParamsWithDefaults instantiates a new SfdcObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSfdcObjectProtectionParamsWithDefaults() *SfdcObjectProtectionParams {
	this := SfdcObjectProtectionParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *SfdcObjectProtectionParams) GetObjects() []SfdcObjectProtectionObjectParams {
	if o == nil {
		var ret []SfdcObjectProtectionObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *SfdcObjectProtectionParams) GetObjectsOk() ([]SfdcObjectProtectionObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *SfdcObjectProtectionParams) SetObjects(v []SfdcObjectProtectionObjectParams) {
	o.Objects = v
}

func (o SfdcObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SfdcObjectProtectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *SfdcObjectProtectionParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objects",
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

	varSfdcObjectProtectionParams := _SfdcObjectProtectionParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSfdcObjectProtectionParams)

	if err != nil {
		return err
	}

	*o = SfdcObjectProtectionParams(varSfdcObjectProtectionParams)

	return err
}

type NullableSfdcObjectProtectionParams struct {
	value *SfdcObjectProtectionParams
	isSet bool
}

func (v NullableSfdcObjectProtectionParams) Get() *SfdcObjectProtectionParams {
	return v.value
}

func (v *NullableSfdcObjectProtectionParams) Set(val *SfdcObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSfdcObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSfdcObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSfdcObjectProtectionParams(val *SfdcObjectProtectionParams) *NullableSfdcObjectProtectionParams {
	return &NullableSfdcObjectProtectionParams{value: val, isSet: true}
}

func (v NullableSfdcObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSfdcObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


