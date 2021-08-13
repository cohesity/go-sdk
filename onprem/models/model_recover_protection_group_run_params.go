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

// RecoverProtectionGroupRunParams Specifies the Protection Group Run params to recover. All the VM's that are successfully backed up by specified Runs will be recovered.
type RecoverProtectionGroupRunParams struct {
	// Specifies the Protection Group Run id from which to recover VMs. All the VM's that are successfully protected by this Run will be recovered.
	ProtectionGroupRunId NullableString `json:"protectionGroupRunId"`
	// Specifies the Protection Group Instance id.
	ProtectionGroupInstanceId NullableInt64 `json:"protectionGroupInstanceId"`
	// Specifies the archival target id. If specified and Protection Group run has an archival snapshot then VMs are recovered from the specified archival snapshot. If not specified (default), VMs are recovered from local snapshot.
	ArchivalTargetId NullableInt64 `json:"archivalTargetId,omitempty"`
	// Specifies the local Protection Group id. In case of recovering a replication Run, this field should be provided with local Protection Group id.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
}

// NewRecoverProtectionGroupRunParams instantiates a new RecoverProtectionGroupRunParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverProtectionGroupRunParams(protectionGroupRunId NullableString, protectionGroupInstanceId NullableInt64) *RecoverProtectionGroupRunParams {
	this := RecoverProtectionGroupRunParams{}
	this.ProtectionGroupRunId = protectionGroupRunId
	this.ProtectionGroupInstanceId = protectionGroupInstanceId
	return &this
}

// NewRecoverProtectionGroupRunParamsWithDefaults instantiates a new RecoverProtectionGroupRunParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverProtectionGroupRunParamsWithDefaults() *RecoverProtectionGroupRunParams {
	this := RecoverProtectionGroupRunParams{}
	return &this
}

// GetProtectionGroupRunId returns the ProtectionGroupRunId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupRunId() string {
	if o == nil || o.ProtectionGroupRunId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProtectionGroupRunId.Get()
}

// GetProtectionGroupRunIdOk returns a tuple with the ProtectionGroupRunId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupRunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupRunId.Get(), o.ProtectionGroupRunId.IsSet()
}

// SetProtectionGroupRunId sets field value
func (o *RecoverProtectionGroupRunParams) SetProtectionGroupRunId(v string) {
	o.ProtectionGroupRunId.Set(&v)
}

// GetProtectionGroupInstanceId returns the ProtectionGroupInstanceId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupInstanceId() int64 {
	if o == nil || o.ProtectionGroupInstanceId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ProtectionGroupInstanceId.Get()
}

// GetProtectionGroupInstanceIdOk returns a tuple with the ProtectionGroupInstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupInstanceId.Get(), o.ProtectionGroupInstanceId.IsSet()
}

// SetProtectionGroupInstanceId sets field value
func (o *RecoverProtectionGroupRunParams) SetProtectionGroupInstanceId(v int64) {
	o.ProtectionGroupInstanceId.Set(&v)
}

// GetArchivalTargetId returns the ArchivalTargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverProtectionGroupRunParams) GetArchivalTargetId() int64 {
	if o == nil || o.ArchivalTargetId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ArchivalTargetId.Get()
}

// GetArchivalTargetIdOk returns a tuple with the ArchivalTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverProtectionGroupRunParams) GetArchivalTargetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArchivalTargetId.Get(), o.ArchivalTargetId.IsSet()
}

// HasArchivalTargetId returns a boolean if a field has been set.
func (o *RecoverProtectionGroupRunParams) HasArchivalTargetId() bool {
	if o != nil && o.ArchivalTargetId.IsSet() {
		return true
	}

	return false
}

// SetArchivalTargetId gets a reference to the given NullableInt64 and assigns it to the ArchivalTargetId field.
func (o *RecoverProtectionGroupRunParams) SetArchivalTargetId(v int64) {
	o.ArchivalTargetId.Set(&v)
}
// SetArchivalTargetIdNil sets the value for ArchivalTargetId to be an explicit nil
func (o *RecoverProtectionGroupRunParams) SetArchivalTargetIdNil() {
	o.ArchivalTargetId.Set(nil)
}

// UnsetArchivalTargetId ensures that no value is present for ArchivalTargetId, not even an explicit nil
func (o *RecoverProtectionGroupRunParams) UnsetArchivalTargetId() {
	o.ArchivalTargetId.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverProtectionGroupRunParams) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *RecoverProtectionGroupRunParams) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *RecoverProtectionGroupRunParams) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *RecoverProtectionGroupRunParams) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *RecoverProtectionGroupRunParams) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

func (o RecoverProtectionGroupRunParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protectionGroupRunId"] = o.ProtectionGroupRunId.Get()
	}
	if true {
		toSerialize["protectionGroupInstanceId"] = o.ProtectionGroupInstanceId.Get()
	}
	if o.ArchivalTargetId.IsSet() {
		toSerialize["archivalTargetId"] = o.ArchivalTargetId.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverProtectionGroupRunParams struct {
	value *RecoverProtectionGroupRunParams
	isSet bool
}

func (v NullableRecoverProtectionGroupRunParams) Get() *RecoverProtectionGroupRunParams {
	return v.value
}

func (v *NullableRecoverProtectionGroupRunParams) Set(val *RecoverProtectionGroupRunParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverProtectionGroupRunParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverProtectionGroupRunParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverProtectionGroupRunParams(val *RecoverProtectionGroupRunParams) *NullableRecoverProtectionGroupRunParams {
	return &NullableRecoverProtectionGroupRunParams{value: val, isSet: true}
}

func (v NullableRecoverProtectionGroupRunParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverProtectionGroupRunParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverProtectionGroupRunParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}