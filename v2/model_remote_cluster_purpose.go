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

// checks if the RemoteClusterPurpose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteClusterPurpose{}

// RemoteClusterPurpose Specifies the purpose of Remote Cluster.
type RemoteClusterPurpose struct {
	// Specifies the purpose of Remote Cluster.
	Value *string `json:"value,omitempty"`
}

// NewRemoteClusterPurpose instantiates a new RemoteClusterPurpose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteClusterPurpose() *RemoteClusterPurpose {
	this := RemoteClusterPurpose{}
	return &this
}

// NewRemoteClusterPurposeWithDefaults instantiates a new RemoteClusterPurpose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterPurposeWithDefaults() *RemoteClusterPurpose {
	this := RemoteClusterPurpose{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RemoteClusterPurpose) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteClusterPurpose) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RemoteClusterPurpose) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RemoteClusterPurpose) SetValue(v string) {
	o.Value = &v
}

func (o RemoteClusterPurpose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteClusterPurpose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRemoteClusterPurpose struct {
	value *RemoteClusterPurpose
	isSet bool
}

func (v NullableRemoteClusterPurpose) Get() *RemoteClusterPurpose {
	return v.value
}

func (v *NullableRemoteClusterPurpose) Set(val *RemoteClusterPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteClusterPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteClusterPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteClusterPurpose(val *RemoteClusterPurpose) *NullableRemoteClusterPurpose {
	return &NullableRemoteClusterPurpose{value: val, isSet: true}
}

func (v NullableRemoteClusterPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteClusterPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


