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

// UpdateFeatureFlagParams Describes update feature flag request params.
type UpdateFeatureFlagParams struct {
	// Name of the feature flag that is to be updated.
	Name NullableString `json:"name,omitempty"`
	// Bool to specify if it's UI feature flag.
	IsUiFeature NullableBool `json:"isUiFeature,omitempty"`
	// New override status for the feature flag.
	IsApproved NullableBool `json:"isApproved,omitempty"`
	// Reason for updating the feature flag override status.
	Reason NullableString `json:"reason,omitempty"`
	// Clear any override status for the feature flag.
	Clear NullableBool `json:"clear,omitempty"`
	// Specifies the timestamp of override operation.
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
}

// NewUpdateFeatureFlagParams instantiates a new UpdateFeatureFlagParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFeatureFlagParams() *UpdateFeatureFlagParams {
	this := UpdateFeatureFlagParams{}
	return &this
}

// NewUpdateFeatureFlagParamsWithDefaults instantiates a new UpdateFeatureFlagParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFeatureFlagParamsWithDefaults() *UpdateFeatureFlagParams {
	this := UpdateFeatureFlagParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateFeatureFlagParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateFeatureFlagParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetName() {
	o.Name.Unset()
}

// GetIsUiFeature returns the IsUiFeature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetIsUiFeature() bool {
	if o == nil || o.IsUiFeature.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsUiFeature.Get()
}

// GetIsUiFeatureOk returns a tuple with the IsUiFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetIsUiFeatureOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsUiFeature.Get(), o.IsUiFeature.IsSet()
}

// HasIsUiFeature returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasIsUiFeature() bool {
	if o != nil && o.IsUiFeature.IsSet() {
		return true
	}

	return false
}

// SetIsUiFeature gets a reference to the given NullableBool and assigns it to the IsUiFeature field.
func (o *UpdateFeatureFlagParams) SetIsUiFeature(v bool) {
	o.IsUiFeature.Set(&v)
}
// SetIsUiFeatureNil sets the value for IsUiFeature to be an explicit nil
func (o *UpdateFeatureFlagParams) SetIsUiFeatureNil() {
	o.IsUiFeature.Set(nil)
}

// UnsetIsUiFeature ensures that no value is present for IsUiFeature, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetIsUiFeature() {
	o.IsUiFeature.Unset()
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetIsApproved() bool {
	if o == nil || o.IsApproved.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsApproved.Get()
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetIsApprovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsApproved.Get(), o.IsApproved.IsSet()
}

// HasIsApproved returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasIsApproved() bool {
	if o != nil && o.IsApproved.IsSet() {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given NullableBool and assigns it to the IsApproved field.
func (o *UpdateFeatureFlagParams) SetIsApproved(v bool) {
	o.IsApproved.Set(&v)
}
// SetIsApprovedNil sets the value for IsApproved to be an explicit nil
func (o *UpdateFeatureFlagParams) SetIsApprovedNil() {
	o.IsApproved.Set(nil)
}

// UnsetIsApproved ensures that no value is present for IsApproved, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetIsApproved() {
	o.IsApproved.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *UpdateFeatureFlagParams) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *UpdateFeatureFlagParams) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetReason() {
	o.Reason.Unset()
}

// GetClear returns the Clear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetClear() bool {
	if o == nil || o.Clear.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Clear.Get()
}

// GetClearOk returns a tuple with the Clear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetClearOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Clear.Get(), o.Clear.IsSet()
}

// HasClear returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasClear() bool {
	if o != nil && o.Clear.IsSet() {
		return true
	}

	return false
}

// SetClear gets a reference to the given NullableBool and assigns it to the Clear field.
func (o *UpdateFeatureFlagParams) SetClear(v bool) {
	o.Clear.Set(&v)
}
// SetClearNil sets the value for Clear to be an explicit nil
func (o *UpdateFeatureFlagParams) SetClearNil() {
	o.Clear.Set(nil)
}

// UnsetClear ensures that no value is present for Clear, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetClear() {
	o.Clear.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagParams) GetTimestamp() int64 {
	if o == nil || o.Timestamp.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureFlagParams) GetTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *UpdateFeatureFlagParams) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *UpdateFeatureFlagParams) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *UpdateFeatureFlagParams) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *UpdateFeatureFlagParams) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o UpdateFeatureFlagParams) MarshalJSON() ([]byte, error) {
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
	if o.Clear.IsSet() {
		toSerialize["clear"] = o.Clear.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFeatureFlagParams struct {
	value *UpdateFeatureFlagParams
	isSet bool
}

func (v NullableUpdateFeatureFlagParams) Get() *UpdateFeatureFlagParams {
	return v.value
}

func (v *NullableUpdateFeatureFlagParams) Set(val *UpdateFeatureFlagParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFeatureFlagParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFeatureFlagParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFeatureFlagParams(val *UpdateFeatureFlagParams) *NullableUpdateFeatureFlagParams {
	return &NullableUpdateFeatureFlagParams{value: val, isSet: true}
}

func (v NullableUpdateFeatureFlagParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFeatureFlagParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


