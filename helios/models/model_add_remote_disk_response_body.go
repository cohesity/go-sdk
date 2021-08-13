/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// AddRemoteDiskResponseBody Specifies the response of creating remote disk.
type AddRemoteDiskResponseBody struct {
	// Specifies a list of remote disk creating status.
	RemoteDisks []CreateRemoteDiskStatus `json:"remoteDisks,omitempty"`
}

// NewAddRemoteDiskResponseBody instantiates a new AddRemoteDiskResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRemoteDiskResponseBody() *AddRemoteDiskResponseBody {
	this := AddRemoteDiskResponseBody{}
	return &this
}

// NewAddRemoteDiskResponseBodyWithDefaults instantiates a new AddRemoteDiskResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRemoteDiskResponseBodyWithDefaults() *AddRemoteDiskResponseBody {
	this := AddRemoteDiskResponseBody{}
	return &this
}

// GetRemoteDisks returns the RemoteDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddRemoteDiskResponseBody) GetRemoteDisks() []CreateRemoteDiskStatus {
	if o == nil  {
		var ret []CreateRemoteDiskStatus
		return ret
	}
	return o.RemoteDisks
}

// GetRemoteDisksOk returns a tuple with the RemoteDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddRemoteDiskResponseBody) GetRemoteDisksOk() (*[]CreateRemoteDiskStatus, bool) {
	if o == nil || o.RemoteDisks == nil {
		return nil, false
	}
	return &o.RemoteDisks, true
}

// HasRemoteDisks returns a boolean if a field has been set.
func (o *AddRemoteDiskResponseBody) HasRemoteDisks() bool {
	if o != nil && o.RemoteDisks != nil {
		return true
	}

	return false
}

// SetRemoteDisks gets a reference to the given []CreateRemoteDiskStatus and assigns it to the RemoteDisks field.
func (o *AddRemoteDiskResponseBody) SetRemoteDisks(v []CreateRemoteDiskStatus) {
	o.RemoteDisks = v
}

func (o AddRemoteDiskResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteDisks != nil {
		toSerialize["remoteDisks"] = o.RemoteDisks
	}
	return json.Marshal(toSerialize)
}

type NullableAddRemoteDiskResponseBody struct {
	value *AddRemoteDiskResponseBody
	isSet bool
}

func (v NullableAddRemoteDiskResponseBody) Get() *AddRemoteDiskResponseBody {
	return v.value
}

func (v *NullableAddRemoteDiskResponseBody) Set(val *AddRemoteDiskResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRemoteDiskResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRemoteDiskResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRemoteDiskResponseBody(val *AddRemoteDiskResponseBody) *NullableAddRemoteDiskResponseBody {
	return &NullableAddRemoteDiskResponseBody{value: val, isSet: true}
}

func (v NullableAddRemoteDiskResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRemoteDiskResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


