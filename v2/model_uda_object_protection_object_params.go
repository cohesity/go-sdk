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

// checks if the UdaObjectProtectionObjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdaObjectProtectionObjectParams{}

// UdaObjectProtectionObjectParams Specifies the object parameters to create an Universal Data Adapter Object Protection.
type UdaObjectProtectionObjectParams struct {
	// Specifies the ids of the objects to be excluded in the Object Protection. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds []int64 `json:"excludeObjectIds,omitempty"`
	// Specifies the id of the object being protected. This can be a leaf level or non leaf level object.
	Id int64 `json:"id"`
}

type _UdaObjectProtectionObjectParams UdaObjectProtectionObjectParams

// NewUdaObjectProtectionObjectParams instantiates a new UdaObjectProtectionObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaObjectProtectionObjectParams(id int64) *UdaObjectProtectionObjectParams {
	this := UdaObjectProtectionObjectParams{}
	this.Id = id
	return &this
}

// NewUdaObjectProtectionObjectParamsWithDefaults instantiates a new UdaObjectProtectionObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaObjectProtectionObjectParamsWithDefaults() *UdaObjectProtectionObjectParams {
	this := UdaObjectProtectionObjectParams{}
	return &this
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionObjectParams) GetExcludeObjectIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionObjectParams) GetExcludeObjectIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *UdaObjectProtectionObjectParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *UdaObjectProtectionObjectParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = v
}

// GetId returns the Id field value
func (o *UdaObjectProtectionObjectParams) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UdaObjectProtectionObjectParams) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UdaObjectProtectionObjectParams) SetId(v int64) {
	o.Id = v
}

func (o UdaObjectProtectionObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdaObjectProtectionObjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UdaObjectProtectionObjectParams) UnmarshalJSON(data []byte) (err error) {
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

	varUdaObjectProtectionObjectParams := _UdaObjectProtectionObjectParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUdaObjectProtectionObjectParams)

	if err != nil {
		return err
	}

	*o = UdaObjectProtectionObjectParams(varUdaObjectProtectionObjectParams)

	return err
}

type NullableUdaObjectProtectionObjectParams struct {
	value *UdaObjectProtectionObjectParams
	isSet bool
}

func (v NullableUdaObjectProtectionObjectParams) Get() *UdaObjectProtectionObjectParams {
	return v.value
}

func (v *NullableUdaObjectProtectionObjectParams) Set(val *UdaObjectProtectionObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaObjectProtectionObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaObjectProtectionObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaObjectProtectionObjectParams(val *UdaObjectProtectionObjectParams) *NullableUdaObjectProtectionObjectParams {
	return &NullableUdaObjectProtectionObjectParams{value: val, isSet: true}
}

func (v NullableUdaObjectProtectionObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaObjectProtectionObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


