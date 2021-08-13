/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// ReplicaFailoverObject Specifies the object paring of replicated object and failover object created after restore.
type ReplicaFailoverObject struct {
	// Specifies the object Id existing on the replciation cluster.
	ReplicaObjectId NullableInt64 `json:"replicaObjectId,omitempty"`
	// Specifies the corrosponding object id of the failover object.
	FailoverObjectId NullableInt64 `json:"failoverObjectId,omitempty"`
}

// NewReplicaFailoverObject instantiates a new ReplicaFailoverObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaFailoverObject() *ReplicaFailoverObject {
	this := ReplicaFailoverObject{}
	return &this
}

// NewReplicaFailoverObjectWithDefaults instantiates a new ReplicaFailoverObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaFailoverObjectWithDefaults() *ReplicaFailoverObject {
	this := ReplicaFailoverObject{}
	return &this
}

// GetReplicaObjectId returns the ReplicaObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicaFailoverObject) GetReplicaObjectId() int64 {
	if o == nil || o.ReplicaObjectId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaObjectId.Get()
}

// GetReplicaObjectIdOk returns a tuple with the ReplicaObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaFailoverObject) GetReplicaObjectIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReplicaObjectId.Get(), o.ReplicaObjectId.IsSet()
}

// HasReplicaObjectId returns a boolean if a field has been set.
func (o *ReplicaFailoverObject) HasReplicaObjectId() bool {
	if o != nil && o.ReplicaObjectId.IsSet() {
		return true
	}

	return false
}

// SetReplicaObjectId gets a reference to the given NullableInt64 and assigns it to the ReplicaObjectId field.
func (o *ReplicaFailoverObject) SetReplicaObjectId(v int64) {
	o.ReplicaObjectId.Set(&v)
}
// SetReplicaObjectIdNil sets the value for ReplicaObjectId to be an explicit nil
func (o *ReplicaFailoverObject) SetReplicaObjectIdNil() {
	o.ReplicaObjectId.Set(nil)
}

// UnsetReplicaObjectId ensures that no value is present for ReplicaObjectId, not even an explicit nil
func (o *ReplicaFailoverObject) UnsetReplicaObjectId() {
	o.ReplicaObjectId.Unset()
}

// GetFailoverObjectId returns the FailoverObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicaFailoverObject) GetFailoverObjectId() int64 {
	if o == nil || o.FailoverObjectId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FailoverObjectId.Get()
}

// GetFailoverObjectIdOk returns a tuple with the FailoverObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicaFailoverObject) GetFailoverObjectIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailoverObjectId.Get(), o.FailoverObjectId.IsSet()
}

// HasFailoverObjectId returns a boolean if a field has been set.
func (o *ReplicaFailoverObject) HasFailoverObjectId() bool {
	if o != nil && o.FailoverObjectId.IsSet() {
		return true
	}

	return false
}

// SetFailoverObjectId gets a reference to the given NullableInt64 and assigns it to the FailoverObjectId field.
func (o *ReplicaFailoverObject) SetFailoverObjectId(v int64) {
	o.FailoverObjectId.Set(&v)
}
// SetFailoverObjectIdNil sets the value for FailoverObjectId to be an explicit nil
func (o *ReplicaFailoverObject) SetFailoverObjectIdNil() {
	o.FailoverObjectId.Set(nil)
}

// UnsetFailoverObjectId ensures that no value is present for FailoverObjectId, not even an explicit nil
func (o *ReplicaFailoverObject) UnsetFailoverObjectId() {
	o.FailoverObjectId.Unset()
}

func (o ReplicaFailoverObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReplicaObjectId.IsSet() {
		toSerialize["replicaObjectId"] = o.ReplicaObjectId.Get()
	}
	if o.FailoverObjectId.IsSet() {
		toSerialize["failoverObjectId"] = o.FailoverObjectId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReplicaFailoverObject struct {
	value *ReplicaFailoverObject
	isSet bool
}

func (v NullableReplicaFailoverObject) Get() *ReplicaFailoverObject {
	return v.value
}

func (v *NullableReplicaFailoverObject) Set(val *ReplicaFailoverObject) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaFailoverObject) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaFailoverObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaFailoverObject(val *ReplicaFailoverObject) *NullableReplicaFailoverObject {
	return &NullableReplicaFailoverObject{value: val, isSet: true}
}

func (v NullableReplicaFailoverObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaFailoverObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ReplicaFailoverObject) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}