/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the AzureDiskTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureDiskTag{}

// AzureDiskTag A key - value pair tag that may be specified on Azure Disks.
type AzureDiskTag struct {
	// Key of the tag specified on the azure disk.
	Key *string `json:"key,omitempty"`
	// Value of the tag specified on the azure disk.
	Value *string `json:"value,omitempty"`
}

// NewAzureDiskTag instantiates a new AzureDiskTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDiskTag() *AzureDiskTag {
	this := AzureDiskTag{}
	return &this
}

// NewAzureDiskTagWithDefaults instantiates a new AzureDiskTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDiskTagWithDefaults() *AzureDiskTag {
	this := AzureDiskTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AzureDiskTag) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDiskTag) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AzureDiskTag) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AzureDiskTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AzureDiskTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDiskTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AzureDiskTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AzureDiskTag) SetValue(v string) {
	o.Value = &v
}

func (o AzureDiskTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureDiskTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAzureDiskTag struct {
	value *AzureDiskTag
	isSet bool
}

func (v NullableAzureDiskTag) Get() *AzureDiskTag {
	return v.value
}

func (v *NullableAzureDiskTag) Set(val *AzureDiskTag) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDiskTag) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDiskTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDiskTag(val *AzureDiskTag) *NullableAzureDiskTag {
	return &NullableAzureDiskTag{value: val, isSet: true}
}

func (v NullableAzureDiskTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDiskTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


