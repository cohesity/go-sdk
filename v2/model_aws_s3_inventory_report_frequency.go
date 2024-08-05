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

// checks if the AwsS3InventoryReportFrequency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsS3InventoryReportFrequency{}

// AwsS3InventoryReportFrequency Specifies the S3 inventory report frequency.
type AwsS3InventoryReportFrequency struct {
	// Specifies the S3 inventory report frequency.
	Enum *string `json:"enum,omitempty"`
}

// NewAwsS3InventoryReportFrequency instantiates a new AwsS3InventoryReportFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsS3InventoryReportFrequency() *AwsS3InventoryReportFrequency {
	this := AwsS3InventoryReportFrequency{}
	return &this
}

// NewAwsS3InventoryReportFrequencyWithDefaults instantiates a new AwsS3InventoryReportFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsS3InventoryReportFrequencyWithDefaults() *AwsS3InventoryReportFrequency {
	this := AwsS3InventoryReportFrequency{}
	return &this
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *AwsS3InventoryReportFrequency) GetEnum() string {
	if o == nil || IsNil(o.Enum) {
		var ret string
		return ret
	}
	return *o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3InventoryReportFrequency) GetEnumOk() (*string, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *AwsS3InventoryReportFrequency) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given string and assigns it to the Enum field.
func (o *AwsS3InventoryReportFrequency) SetEnum(v string) {
	o.Enum = &v
}

func (o AwsS3InventoryReportFrequency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsS3InventoryReportFrequency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enum) {
		toSerialize["enum"] = o.Enum
	}
	return toSerialize, nil
}

type NullableAwsS3InventoryReportFrequency struct {
	value *AwsS3InventoryReportFrequency
	isSet bool
}

func (v NullableAwsS3InventoryReportFrequency) Get() *AwsS3InventoryReportFrequency {
	return v.value
}

func (v *NullableAwsS3InventoryReportFrequency) Set(val *AwsS3InventoryReportFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsS3InventoryReportFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsS3InventoryReportFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsS3InventoryReportFrequency(val *AwsS3InventoryReportFrequency) *NullableAwsS3InventoryReportFrequency {
	return &NullableAwsS3InventoryReportFrequency{value: val, isSet: true}
}

func (v NullableAwsS3InventoryReportFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsS3InventoryReportFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


