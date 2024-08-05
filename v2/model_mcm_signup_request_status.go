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

// checks if the McmSignupRequestStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McmSignupRequestStatus{}

// McmSignupRequestStatus MCM signup request status type.
type McmSignupRequestStatus struct {
	// MCM signup request status type.
	Type *string `json:"type,omitempty"`
}

// NewMcmSignupRequestStatus instantiates a new McmSignupRequestStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcmSignupRequestStatus() *McmSignupRequestStatus {
	this := McmSignupRequestStatus{}
	return &this
}

// NewMcmSignupRequestStatusWithDefaults instantiates a new McmSignupRequestStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcmSignupRequestStatusWithDefaults() *McmSignupRequestStatus {
	this := McmSignupRequestStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *McmSignupRequestStatus) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McmSignupRequestStatus) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *McmSignupRequestStatus) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *McmSignupRequestStatus) SetType(v string) {
	o.Type = &v
}

func (o McmSignupRequestStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McmSignupRequestStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMcmSignupRequestStatus struct {
	value *McmSignupRequestStatus
	isSet bool
}

func (v NullableMcmSignupRequestStatus) Get() *McmSignupRequestStatus {
	return v.value
}

func (v *NullableMcmSignupRequestStatus) Set(val *McmSignupRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMcmSignupRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMcmSignupRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcmSignupRequestStatus(val *McmSignupRequestStatus) *NullableMcmSignupRequestStatus {
	return &NullableMcmSignupRequestStatus{value: val, isSet: true}
}

func (v NullableMcmSignupRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcmSignupRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


