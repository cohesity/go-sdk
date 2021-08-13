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

// RecoverMsGroupParams Specifies the parameters to recover Microsoft 365 Group.
type RecoverMsGroupParams struct {
	// Specifies a list of groups getting restored.
	MsGroups []MsGroupParam `json:"msGroups"`
	// Specifies whether or not all groups are restored to original location.
	RestoreToOriginal NullableBool `json:"restoreToOriginal,omitempty"`
	// Specifies target group nickname in case restoreToOriginal is false. This needs to be specifid when restoreToOriginal is false.
	TargetGroup NullableString `json:"targetGroup,omitempty"`
	// Specifies target group name in case restore_to_original is false. This needs to be specifid when restoreToOriginal is false. However, this will be ignored if restoring to alternate existing group (i.e. to a group the nickname of which is same as the one supplied by the end user).
	TargetGroupName NullableString `json:"targetGroupName,omitempty"`
	// Specifies whether to continue recovering other MS groups if one of MS groups failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewRecoverMsGroupParams instantiates a new RecoverMsGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverMsGroupParams(msGroups []MsGroupParam) *RecoverMsGroupParams {
	this := RecoverMsGroupParams{}
	this.MsGroups = msGroups
	return &this
}

// NewRecoverMsGroupParamsWithDefaults instantiates a new RecoverMsGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverMsGroupParamsWithDefaults() *RecoverMsGroupParams {
	this := RecoverMsGroupParams{}
	return &this
}

// GetMsGroups returns the MsGroups field value
// If the value is explicit nil, the zero value for []MsGroupParam will be returned
func (o *RecoverMsGroupParams) GetMsGroups() []MsGroupParam {
	if o == nil {
		var ret []MsGroupParam
		return ret
	}

	return o.MsGroups
}

// GetMsGroupsOk returns a tuple with the MsGroups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsGroupParams) GetMsGroupsOk() (*[]MsGroupParam, bool) {
	if o == nil || o.MsGroups == nil {
		return nil, false
	}
	return &o.MsGroups, true
}

// SetMsGroups sets field value
func (o *RecoverMsGroupParams) SetMsGroups(v []MsGroupParam) {
	o.MsGroups = v
}

// GetRestoreToOriginal returns the RestoreToOriginal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsGroupParams) GetRestoreToOriginal() bool {
	if o == nil || o.RestoreToOriginal.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RestoreToOriginal.Get()
}

// GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsGroupParams) GetRestoreToOriginalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreToOriginal.Get(), o.RestoreToOriginal.IsSet()
}

// HasRestoreToOriginal returns a boolean if a field has been set.
func (o *RecoverMsGroupParams) HasRestoreToOriginal() bool {
	if o != nil && o.RestoreToOriginal.IsSet() {
		return true
	}

	return false
}

// SetRestoreToOriginal gets a reference to the given NullableBool and assigns it to the RestoreToOriginal field.
func (o *RecoverMsGroupParams) SetRestoreToOriginal(v bool) {
	o.RestoreToOriginal.Set(&v)
}
// SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil
func (o *RecoverMsGroupParams) SetRestoreToOriginalNil() {
	o.RestoreToOriginal.Set(nil)
}

// UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
func (o *RecoverMsGroupParams) UnsetRestoreToOriginal() {
	o.RestoreToOriginal.Unset()
}

// GetTargetGroup returns the TargetGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsGroupParams) GetTargetGroup() string {
	if o == nil || o.TargetGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetGroup.Get()
}

// GetTargetGroupOk returns a tuple with the TargetGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsGroupParams) GetTargetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetGroup.Get(), o.TargetGroup.IsSet()
}

// HasTargetGroup returns a boolean if a field has been set.
func (o *RecoverMsGroupParams) HasTargetGroup() bool {
	if o != nil && o.TargetGroup.IsSet() {
		return true
	}

	return false
}

// SetTargetGroup gets a reference to the given NullableString and assigns it to the TargetGroup field.
func (o *RecoverMsGroupParams) SetTargetGroup(v string) {
	o.TargetGroup.Set(&v)
}
// SetTargetGroupNil sets the value for TargetGroup to be an explicit nil
func (o *RecoverMsGroupParams) SetTargetGroupNil() {
	o.TargetGroup.Set(nil)
}

// UnsetTargetGroup ensures that no value is present for TargetGroup, not even an explicit nil
func (o *RecoverMsGroupParams) UnsetTargetGroup() {
	o.TargetGroup.Unset()
}

// GetTargetGroupName returns the TargetGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsGroupParams) GetTargetGroupName() string {
	if o == nil || o.TargetGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetGroupName.Get()
}

// GetTargetGroupNameOk returns a tuple with the TargetGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsGroupParams) GetTargetGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetGroupName.Get(), o.TargetGroupName.IsSet()
}

// HasTargetGroupName returns a boolean if a field has been set.
func (o *RecoverMsGroupParams) HasTargetGroupName() bool {
	if o != nil && o.TargetGroupName.IsSet() {
		return true
	}

	return false
}

// SetTargetGroupName gets a reference to the given NullableString and assigns it to the TargetGroupName field.
func (o *RecoverMsGroupParams) SetTargetGroupName(v string) {
	o.TargetGroupName.Set(&v)
}
// SetTargetGroupNameNil sets the value for TargetGroupName to be an explicit nil
func (o *RecoverMsGroupParams) SetTargetGroupNameNil() {
	o.TargetGroupName.Set(nil)
}

// UnsetTargetGroupName ensures that no value is present for TargetGroupName, not even an explicit nil
func (o *RecoverMsGroupParams) UnsetTargetGroupName() {
	o.TargetGroupName.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMsGroupParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMsGroupParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverMsGroupParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverMsGroupParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverMsGroupParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverMsGroupParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o RecoverMsGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MsGroups != nil {
		toSerialize["msGroups"] = o.MsGroups
	}
	if o.RestoreToOriginal.IsSet() {
		toSerialize["restoreToOriginal"] = o.RestoreToOriginal.Get()
	}
	if o.TargetGroup.IsSet() {
		toSerialize["targetGroup"] = o.TargetGroup.Get()
	}
	if o.TargetGroupName.IsSet() {
		toSerialize["targetGroupName"] = o.TargetGroupName.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverMsGroupParams struct {
	value *RecoverMsGroupParams
	isSet bool
}

func (v NullableRecoverMsGroupParams) Get() *RecoverMsGroupParams {
	return v.value
}

func (v *NullableRecoverMsGroupParams) Set(val *RecoverMsGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverMsGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverMsGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverMsGroupParams(val *RecoverMsGroupParams) *NullableRecoverMsGroupParams {
	return &NullableRecoverMsGroupParams{value: val, isSet: true}
}

func (v NullableRecoverMsGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverMsGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverMsGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}