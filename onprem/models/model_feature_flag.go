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

// FeatureFlag Describes the feature flag struct.
type FeatureFlag struct {
	// Name of the feature flag.
	Name NullableString `json:"name,omitempty"`
	// Bool to denote if it's a UI feature.
	IsUiFeature NullableBool `json:"isUiFeature,omitempty"`
	// Bool to approval status of feature flag.
	IsApproved NullableBool `json:"isApproved,omitempty"`
	// Reason for the feature flag override status.
	Reason NullableString `json:"reason,omitempty"`
	// Timestamp in secs when the override is done.
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
}

// NewFeatureFlag instantiates a new FeatureFlag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureFlag() *FeatureFlag {
	this := FeatureFlag{}
	return &this
}

// NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureFlagWithDefaults() *FeatureFlag {
	this := FeatureFlag{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlag) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureFlag) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FeatureFlag) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FeatureFlag) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FeatureFlag) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FeatureFlag) UnsetName() {
	o.Name.Unset()
}

// GetIsUiFeature returns the IsUiFeature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlag) GetIsUiFeature() bool {
	if o == nil || o.IsUiFeature.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsUiFeature.Get()
}

// GetIsUiFeatureOk returns a tuple with the IsUiFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureFlag) GetIsUiFeatureOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsUiFeature.Get(), o.IsUiFeature.IsSet()
}

// HasIsUiFeature returns a boolean if a field has been set.
func (o *FeatureFlag) HasIsUiFeature() bool {
	if o != nil && o.IsUiFeature.IsSet() {
		return true
	}

	return false
}

// SetIsUiFeature gets a reference to the given NullableBool and assigns it to the IsUiFeature field.
func (o *FeatureFlag) SetIsUiFeature(v bool) {
	o.IsUiFeature.Set(&v)
}
// SetIsUiFeatureNil sets the value for IsUiFeature to be an explicit nil
func (o *FeatureFlag) SetIsUiFeatureNil() {
	o.IsUiFeature.Set(nil)
}

// UnsetIsUiFeature ensures that no value is present for IsUiFeature, not even an explicit nil
func (o *FeatureFlag) UnsetIsUiFeature() {
	o.IsUiFeature.Unset()
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlag) GetIsApproved() bool {
	if o == nil || o.IsApproved.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsApproved.Get()
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureFlag) GetIsApprovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsApproved.Get(), o.IsApproved.IsSet()
}

// HasIsApproved returns a boolean if a field has been set.
func (o *FeatureFlag) HasIsApproved() bool {
	if o != nil && o.IsApproved.IsSet() {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given NullableBool and assigns it to the IsApproved field.
func (o *FeatureFlag) SetIsApproved(v bool) {
	o.IsApproved.Set(&v)
}
// SetIsApprovedNil sets the value for IsApproved to be an explicit nil
func (o *FeatureFlag) SetIsApprovedNil() {
	o.IsApproved.Set(nil)
}

// UnsetIsApproved ensures that no value is present for IsApproved, not even an explicit nil
func (o *FeatureFlag) UnsetIsApproved() {
	o.IsApproved.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlag) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureFlag) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *FeatureFlag) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *FeatureFlag) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *FeatureFlag) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *FeatureFlag) UnsetReason() {
	o.Reason.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlag) GetTimestamp() int64 {
	if o == nil || o.Timestamp.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureFlag) GetTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FeatureFlag) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *FeatureFlag) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *FeatureFlag) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *FeatureFlag) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o FeatureFlag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsUiFeature.IsSet() {
		toSerialize["isUiFeature"] = o.IsUiFeature.Get()
	}
	if o.IsApproved.IsSet() {
		toSerialize["isApproved"] = o.IsApproved.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureFlag struct {
	value *FeatureFlag
	isSet bool
}

func (v NullableFeatureFlag) Get() *FeatureFlag {
	return v.value
}

func (v *NullableFeatureFlag) Set(val *FeatureFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureFlag(val *FeatureFlag) *NullableFeatureFlag {
	return &NullableFeatureFlag{value: val, isSet: true}
}

func (v NullableFeatureFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o FeatureFlag) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}