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

// checks if the AwsSubscriptionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsSubscriptionType{}

// AwsSubscriptionType AWS Subscription Type
type AwsSubscriptionType struct {
	// Specifies the AWS Subscription types.
	AwsSubscriptionType *string `json:"awsSubscriptionType,omitempty"`
}

// NewAwsSubscriptionType instantiates a new AwsSubscriptionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSubscriptionType() *AwsSubscriptionType {
	this := AwsSubscriptionType{}
	return &this
}

// NewAwsSubscriptionTypeWithDefaults instantiates a new AwsSubscriptionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSubscriptionTypeWithDefaults() *AwsSubscriptionType {
	this := AwsSubscriptionType{}
	return &this
}

// GetAwsSubscriptionType returns the AwsSubscriptionType field value if set, zero value otherwise.
func (o *AwsSubscriptionType) GetAwsSubscriptionType() string {
	if o == nil || IsNil(o.AwsSubscriptionType) {
		var ret string
		return ret
	}
	return *o.AwsSubscriptionType
}

// GetAwsSubscriptionTypeOk returns a tuple with the AwsSubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSubscriptionType) GetAwsSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSubscriptionType) {
		return nil, false
	}
	return o.AwsSubscriptionType, true
}

// HasAwsSubscriptionType returns a boolean if a field has been set.
func (o *AwsSubscriptionType) HasAwsSubscriptionType() bool {
	if o != nil && !IsNil(o.AwsSubscriptionType) {
		return true
	}

	return false
}

// SetAwsSubscriptionType gets a reference to the given string and assigns it to the AwsSubscriptionType field.
func (o *AwsSubscriptionType) SetAwsSubscriptionType(v string) {
	o.AwsSubscriptionType = &v
}

func (o AwsSubscriptionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsSubscriptionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsSubscriptionType) {
		toSerialize["awsSubscriptionType"] = o.AwsSubscriptionType
	}
	return toSerialize, nil
}

type NullableAwsSubscriptionType struct {
	value *AwsSubscriptionType
	isSet bool
}

func (v NullableAwsSubscriptionType) Get() *AwsSubscriptionType {
	return v.value
}

func (v *NullableAwsSubscriptionType) Set(val *AwsSubscriptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSubscriptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSubscriptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSubscriptionType(val *AwsSubscriptionType) *NullableAwsSubscriptionType {
	return &NullableAwsSubscriptionType{value: val, isSet: true}
}

func (v NullableAwsSubscriptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSubscriptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


