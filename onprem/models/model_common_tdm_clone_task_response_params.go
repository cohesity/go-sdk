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

// CommonTdmCloneTaskResponseParams Specifies the common response params for a TDM clone task.
type CommonTdmCloneTaskResponseParams struct {
	// Specifies the environment of the TDM Clone task.
	Environment NullableString `json:"environment"`
	// Specifies the ID of an existing protection group, which should start protecting this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupName should not be specified.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies the name of a new protection group, which should be created to protect this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupId should not be specified.
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// Specifies the ID of the policy, which should be used to protect this clone. This is useful for automatic snapshots. This must be specified if either of protectionGroupId and protectionGroupName is specified.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Specifies the details of the snapshot used for cloning.
	Snapshot NullableObjectSnapshot `json:"snapshot,omitempty"`
	// Specifies the details of the parent object of the clone.
	Parent NullableObjectSummary `json:"parent,omitempty"`
	// Specifies the details of the target, where the clone is created.
	Target NullableObjectSummary `json:"target,omitempty"`
	// Specifies the details of the view, which is used for the clone.
	View NullableViewParams `json:"view,omitempty"`
}

// NewCommonTdmCloneTaskResponseParams instantiates a new CommonTdmCloneTaskResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTdmCloneTaskResponseParams(environment NullableString) *CommonTdmCloneTaskResponseParams {
	this := CommonTdmCloneTaskResponseParams{}
	this.Environment = environment
	return &this
}

// NewCommonTdmCloneTaskResponseParamsWithDefaults instantiates a new CommonTdmCloneTaskResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTdmCloneTaskResponseParamsWithDefaults() *CommonTdmCloneTaskResponseParams {
	this := CommonTdmCloneTaskResponseParams{}
	return &this
}

// GetEnvironment returns the Environment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTdmCloneTaskResponseParams) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// SetEnvironment sets field value
func (o *CommonTdmCloneTaskResponseParams) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *CommonTdmCloneTaskResponseParams) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *CommonTdmCloneTaskResponseParams) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *CommonTdmCloneTaskResponseParams) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetSnapshot() ObjectSnapshot {
	if o == nil || o.Snapshot.Get() == nil {
		var ret ObjectSnapshot
		return ret
	}
	return *o.Snapshot.Get()
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetSnapshotOk() (*ObjectSnapshot, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Snapshot.Get(), o.Snapshot.IsSet()
}

// HasSnapshot returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasSnapshot() bool {
	if o != nil && o.Snapshot.IsSet() {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given NullableObjectSnapshot and assigns it to the Snapshot field.
func (o *CommonTdmCloneTaskResponseParams) SetSnapshot(v ObjectSnapshot) {
	o.Snapshot.Set(&v)
}
// SetSnapshotNil sets the value for Snapshot to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetSnapshotNil() {
	o.Snapshot.Set(nil)
}

// UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetSnapshot() {
	o.Snapshot.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetParent() ObjectSummary {
	if o == nil || o.Parent.Get() == nil {
		var ret ObjectSummary
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetParentOk() (*ObjectSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableObjectSummary and assigns it to the Parent field.
func (o *CommonTdmCloneTaskResponseParams) SetParent(v ObjectSummary) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetParent() {
	o.Parent.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetTarget() ObjectSummary {
	if o == nil || o.Target.Get() == nil {
		var ret ObjectSummary
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetTargetOk() (*ObjectSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableObjectSummary and assigns it to the Target field.
func (o *CommonTdmCloneTaskResponseParams) SetTarget(v ObjectSummary) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetTarget() {
	o.Target.Unset()
}

// GetView returns the View field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmCloneTaskResponseParams) GetView() ViewParams {
	if o == nil || o.View.Get() == nil {
		var ret ViewParams
		return ret
	}
	return *o.View.Get()
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmCloneTaskResponseParams) GetViewOk() (*ViewParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.View.Get(), o.View.IsSet()
}

// HasView returns a boolean if a field has been set.
func (o *CommonTdmCloneTaskResponseParams) HasView() bool {
	if o != nil && o.View.IsSet() {
		return true
	}

	return false
}

// SetView gets a reference to the given NullableViewParams and assigns it to the View field.
func (o *CommonTdmCloneTaskResponseParams) SetView(v ViewParams) {
	o.View.Set(&v)
}
// SetViewNil sets the value for View to be an explicit nil
func (o *CommonTdmCloneTaskResponseParams) SetViewNil() {
	o.View.Set(nil)
}

// UnsetView ensures that no value is present for View, not even an explicit nil
func (o *CommonTdmCloneTaskResponseParams) UnsetView() {
	o.View.Unset()
}

func (o CommonTdmCloneTaskResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.Snapshot.IsSet() {
		toSerialize["snapshot"] = o.Snapshot.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	if o.View.IsSet() {
		toSerialize["view"] = o.View.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTdmCloneTaskResponseParams struct {
	value *CommonTdmCloneTaskResponseParams
	isSet bool
}

func (v NullableCommonTdmCloneTaskResponseParams) Get() *CommonTdmCloneTaskResponseParams {
	return v.value
}

func (v *NullableCommonTdmCloneTaskResponseParams) Set(val *CommonTdmCloneTaskResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTdmCloneTaskResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTdmCloneTaskResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTdmCloneTaskResponseParams(val *CommonTdmCloneTaskResponseParams) *NullableCommonTdmCloneTaskResponseParams {
	return &NullableCommonTdmCloneTaskResponseParams{value: val, isSet: true}
}

func (v NullableCommonTdmCloneTaskResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTdmCloneTaskResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonTdmCloneTaskResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}