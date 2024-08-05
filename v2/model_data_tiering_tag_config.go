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

// checks if the DataTieringTagConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataTieringTagConfig{}

// DataTieringTagConfig Array of data tiering tag objects.
type DataTieringTagConfig struct {
	// Specifies the ID of the data tiering analysis group.
	Id NullableString `json:"id,omitempty"`
	// Array of Tag objects.
	TagsInfo []DataTieringTagObject `json:"tagsInfo"`
}

type _DataTieringTagConfig DataTieringTagConfig

// NewDataTieringTagConfig instantiates a new DataTieringTagConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTieringTagConfig(tagsInfo []DataTieringTagObject) *DataTieringTagConfig {
	this := DataTieringTagConfig{}
	this.TagsInfo = tagsInfo
	return &this
}

// NewDataTieringTagConfigWithDefaults instantiates a new DataTieringTagConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTieringTagConfigWithDefaults() *DataTieringTagConfig {
	this := DataTieringTagConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringTagConfig) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringTagConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DataTieringTagConfig) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DataTieringTagConfig) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DataTieringTagConfig) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DataTieringTagConfig) UnsetId() {
	o.Id.Unset()
}

// GetTagsInfo returns the TagsInfo field value
// If the value is explicit nil, the zero value for []DataTieringTagObject will be returned
func (o *DataTieringTagConfig) GetTagsInfo() []DataTieringTagObject {
	if o == nil {
		var ret []DataTieringTagObject
		return ret
	}

	return o.TagsInfo
}

// GetTagsInfoOk returns a tuple with the TagsInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringTagConfig) GetTagsInfoOk() ([]DataTieringTagObject, bool) {
	if o == nil || IsNil(o.TagsInfo) {
		return nil, false
	}
	return o.TagsInfo, true
}

// SetTagsInfo sets field value
func (o *DataTieringTagConfig) SetTagsInfo(v []DataTieringTagObject) {
	o.TagsInfo = v
}

func (o DataTieringTagConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataTieringTagConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TagsInfo != nil {
		toSerialize["tagsInfo"] = o.TagsInfo
	}
	return toSerialize, nil
}

func (o *DataTieringTagConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tagsInfo",
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

	varDataTieringTagConfig := _DataTieringTagConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataTieringTagConfig)

	if err != nil {
		return err
	}

	*o = DataTieringTagConfig(varDataTieringTagConfig)

	return err
}

type NullableDataTieringTagConfig struct {
	value *DataTieringTagConfig
	isSet bool
}

func (v NullableDataTieringTagConfig) Get() *DataTieringTagConfig {
	return v.value
}

func (v *NullableDataTieringTagConfig) Set(val *DataTieringTagConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTieringTagConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTieringTagConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTieringTagConfig(val *DataTieringTagConfig) *NullableDataTieringTagConfig {
	return &NullableDataTieringTagConfig{value: val, isSet: true}
}

func (v NullableDataTieringTagConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTieringTagConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


