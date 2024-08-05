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

// checks if the IsilonDataTieringParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsilonDataTieringParams{}

// IsilonDataTieringParams Specifies the parameters which are specific to Isilon related Protection Groups.
type IsilonDataTieringParams struct {
	// Specifies the objects to be included in the Protection Group.
	Objects []ProtectionObjectInput `json:"objects"`
	// Specifies the id of the root of data tiering source.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
}

type _IsilonDataTieringParams IsilonDataTieringParams

// NewIsilonDataTieringParams instantiates a new IsilonDataTieringParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsilonDataTieringParams(objects []ProtectionObjectInput) *IsilonDataTieringParams {
	this := IsilonDataTieringParams{}
	this.Objects = objects
	return &this
}

// NewIsilonDataTieringParamsWithDefaults instantiates a new IsilonDataTieringParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsilonDataTieringParamsWithDefaults() *IsilonDataTieringParams {
	this := IsilonDataTieringParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *IsilonDataTieringParams) GetObjects() []ProtectionObjectInput {
	if o == nil {
		var ret []ProtectionObjectInput
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *IsilonDataTieringParams) GetObjectsOk() ([]ProtectionObjectInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *IsilonDataTieringParams) SetObjects(v []ProtectionObjectInput) {
	o.Objects = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonDataTieringParams) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonDataTieringParams) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *IsilonDataTieringParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *IsilonDataTieringParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *IsilonDataTieringParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *IsilonDataTieringParams) UnsetSourceId() {
	o.SourceId.Unset()
}

func (o IsilonDataTieringParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsilonDataTieringParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objects"] = o.Objects
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	return toSerialize, nil
}

func (o *IsilonDataTieringParams) UnmarshalJSON(data []byte) (err error) {
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

	varIsilonDataTieringParams := _IsilonDataTieringParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIsilonDataTieringParams)

	if err != nil {
		return err
	}

	*o = IsilonDataTieringParams(varIsilonDataTieringParams)

	return err
}

type NullableIsilonDataTieringParams struct {
	value *IsilonDataTieringParams
	isSet bool
}

func (v NullableIsilonDataTieringParams) Get() *IsilonDataTieringParams {
	return v.value
}

func (v *NullableIsilonDataTieringParams) Set(val *IsilonDataTieringParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIsilonDataTieringParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIsilonDataTieringParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsilonDataTieringParams(val *IsilonDataTieringParams) *NullableIsilonDataTieringParams {
	return &NullableIsilonDataTieringParams{value: val, isSet: true}
}

func (v NullableIsilonDataTieringParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsilonDataTieringParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


