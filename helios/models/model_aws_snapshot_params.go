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

// AwsSnapshotParams Specifies parameters of AWS type snapshots.
type AwsSnapshotParams struct {
	// Specifies the protection type of AWS snapshots.
	ProtectionType NullableString `json:"protectionType,omitempty"`
}

// NewAwsSnapshotParams instantiates a new AwsSnapshotParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSnapshotParams() *AwsSnapshotParams {
	this := AwsSnapshotParams{}
	return &this
}

// NewAwsSnapshotParamsWithDefaults instantiates a new AwsSnapshotParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSnapshotParamsWithDefaults() *AwsSnapshotParams {
	this := AwsSnapshotParams{}
	return &this
}

// GetProtectionType returns the ProtectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsSnapshotParams) GetProtectionType() string {
	if o == nil || o.ProtectionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionType.Get()
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSnapshotParams) GetProtectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionType.Get(), o.ProtectionType.IsSet()
}

// HasProtectionType returns a boolean if a field has been set.
func (o *AwsSnapshotParams) HasProtectionType() bool {
	if o != nil && o.ProtectionType.IsSet() {
		return true
	}

	return false
}

// SetProtectionType gets a reference to the given NullableString and assigns it to the ProtectionType field.
func (o *AwsSnapshotParams) SetProtectionType(v string) {
	o.ProtectionType.Set(&v)
}
// SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil
func (o *AwsSnapshotParams) SetProtectionTypeNil() {
	o.ProtectionType.Set(nil)
}

// UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
func (o *AwsSnapshotParams) UnsetProtectionType() {
	o.ProtectionType.Unset()
}

func (o AwsSnapshotParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionType.IsSet() {
		toSerialize["protectionType"] = o.ProtectionType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsSnapshotParams struct {
	value *AwsSnapshotParams
	isSet bool
}

func (v NullableAwsSnapshotParams) Get() *AwsSnapshotParams {
	return v.value
}

func (v *NullableAwsSnapshotParams) Set(val *AwsSnapshotParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSnapshotParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSnapshotParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSnapshotParams(val *AwsSnapshotParams) *NullableAwsSnapshotParams {
	return &NullableAwsSnapshotParams{value: val, isSet: true}
}

func (v NullableAwsSnapshotParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSnapshotParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


