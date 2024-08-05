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

// checks if the IsilonObjectProtectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsilonObjectProtectionParams{}

// IsilonObjectProtectionParams Specifies the parameters which are specific to Isilon object protection.
type IsilonObjectProtectionParams struct {
	ContinuousSnapshots *ContinuousSnapshotParams `json:"continuousSnapshots,omitempty"`
	// Specifies the protocol of the NAS device being backed up.
	Protocol NullableString `json:"protocol,omitempty"`
	// Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup.
	UseChangelist NullableBool `json:"useChangelist,omitempty"`
}

// NewIsilonObjectProtectionParams instantiates a new IsilonObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsilonObjectProtectionParams() *IsilonObjectProtectionParams {
	this := IsilonObjectProtectionParams{}
	return &this
}

// NewIsilonObjectProtectionParamsWithDefaults instantiates a new IsilonObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsilonObjectProtectionParamsWithDefaults() *IsilonObjectProtectionParams {
	this := IsilonObjectProtectionParams{}
	return &this
}

// GetContinuousSnapshots returns the ContinuousSnapshots field value if set, zero value otherwise.
func (o *IsilonObjectProtectionParams) GetContinuousSnapshots() ContinuousSnapshotParams {
	if o == nil || IsNil(o.ContinuousSnapshots) {
		var ret ContinuousSnapshotParams
		return ret
	}
	return *o.ContinuousSnapshots
}

// GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonObjectProtectionParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool) {
	if o == nil || IsNil(o.ContinuousSnapshots) {
		return nil, false
	}
	return o.ContinuousSnapshots, true
}

// HasContinuousSnapshots returns a boolean if a field has been set.
func (o *IsilonObjectProtectionParams) HasContinuousSnapshots() bool {
	if o != nil && !IsNil(o.ContinuousSnapshots) {
		return true
	}

	return false
}

// SetContinuousSnapshots gets a reference to the given ContinuousSnapshotParams and assigns it to the ContinuousSnapshots field.
func (o *IsilonObjectProtectionParams) SetContinuousSnapshots(v ContinuousSnapshotParams) {
	o.ContinuousSnapshots = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionParams) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionParams) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *IsilonObjectProtectionParams) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *IsilonObjectProtectionParams) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *IsilonObjectProtectionParams) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *IsilonObjectProtectionParams) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetUseChangelist returns the UseChangelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonObjectProtectionParams) GetUseChangelist() bool {
	if o == nil || IsNil(o.UseChangelist.Get()) {
		var ret bool
		return ret
	}
	return *o.UseChangelist.Get()
}

// GetUseChangelistOk returns a tuple with the UseChangelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonObjectProtectionParams) GetUseChangelistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseChangelist.Get(), o.UseChangelist.IsSet()
}

// HasUseChangelist returns a boolean if a field has been set.
func (o *IsilonObjectProtectionParams) HasUseChangelist() bool {
	if o != nil && o.UseChangelist.IsSet() {
		return true
	}

	return false
}

// SetUseChangelist gets a reference to the given NullableBool and assigns it to the UseChangelist field.
func (o *IsilonObjectProtectionParams) SetUseChangelist(v bool) {
	o.UseChangelist.Set(&v)
}
// SetUseChangelistNil sets the value for UseChangelist to be an explicit nil
func (o *IsilonObjectProtectionParams) SetUseChangelistNil() {
	o.UseChangelist.Set(nil)
}

// UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil
func (o *IsilonObjectProtectionParams) UnsetUseChangelist() {
	o.UseChangelist.Unset()
}

func (o IsilonObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsilonObjectProtectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContinuousSnapshots) {
		toSerialize["continuousSnapshots"] = o.ContinuousSnapshots
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.UseChangelist.IsSet() {
		toSerialize["useChangelist"] = o.UseChangelist.Get()
	}
	return toSerialize, nil
}

type NullableIsilonObjectProtectionParams struct {
	value *IsilonObjectProtectionParams
	isSet bool
}

func (v NullableIsilonObjectProtectionParams) Get() *IsilonObjectProtectionParams {
	return v.value
}

func (v *NullableIsilonObjectProtectionParams) Set(val *IsilonObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIsilonObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIsilonObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsilonObjectProtectionParams(val *IsilonObjectProtectionParams) *NullableIsilonObjectProtectionParams {
	return &NullableIsilonObjectProtectionParams{value: val, isSet: true}
}

func (v NullableIsilonObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsilonObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


