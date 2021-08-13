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

// ProtectionSummary Specifies a summary of an object protection.
type ProtectionSummary struct {
	// Specifies the policy name for this group.
	PolicyName NullableString `json:"policyName,omitempty"`
	// Specifies the policy id for this protection.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Specifies the storage domain id of this protection. Format is clusterId:clusterIncarnationId:storageDomainId.
	StorageDomainId NullableString `json:"storageDomainId,omitempty"`
	// Specifies the status of last local back up run.
	LastBackupRunStatus NullableString `json:"lastBackupRunStatus,omitempty"`
	// Specifies the status of last archival run.
	LastArchivalRunStatus NullableString `json:"lastArchivalRunStatus,omitempty"`
	// Specifies the status of last replication run.
	LastReplicationRunStatus NullableString `json:"lastReplicationRunStatus,omitempty"`
	// Specifies if the sla is violated in last run.
	LastRunSlaViolated NullableBool `json:"lastRunSlaViolated,omitempty"`
}

// NewProtectionSummary instantiates a new ProtectionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionSummary() *ProtectionSummary {
	this := ProtectionSummary{}
	return &this
}

// NewProtectionSummaryWithDefaults instantiates a new ProtectionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionSummaryWithDefaults() *ProtectionSummary {
	this := ProtectionSummary{}
	return &this
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetPolicyName() string {
	if o == nil || o.PolicyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ProtectionSummary) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *ProtectionSummary) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *ProtectionSummary) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *ProtectionSummary) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ProtectionSummary) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *ProtectionSummary) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *ProtectionSummary) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *ProtectionSummary) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetStorageDomainId() string {
	if o == nil || o.StorageDomainId.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetStorageDomainIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *ProtectionSummary) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableString and assigns it to the StorageDomainId field.
func (o *ProtectionSummary) SetStorageDomainId(v string) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *ProtectionSummary) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *ProtectionSummary) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetLastBackupRunStatus returns the LastBackupRunStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetLastBackupRunStatus() string {
	if o == nil || o.LastBackupRunStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastBackupRunStatus.Get()
}

// GetLastBackupRunStatusOk returns a tuple with the LastBackupRunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetLastBackupRunStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastBackupRunStatus.Get(), o.LastBackupRunStatus.IsSet()
}

// HasLastBackupRunStatus returns a boolean if a field has been set.
func (o *ProtectionSummary) HasLastBackupRunStatus() bool {
	if o != nil && o.LastBackupRunStatus.IsSet() {
		return true
	}

	return false
}

// SetLastBackupRunStatus gets a reference to the given NullableString and assigns it to the LastBackupRunStatus field.
func (o *ProtectionSummary) SetLastBackupRunStatus(v string) {
	o.LastBackupRunStatus.Set(&v)
}
// SetLastBackupRunStatusNil sets the value for LastBackupRunStatus to be an explicit nil
func (o *ProtectionSummary) SetLastBackupRunStatusNil() {
	o.LastBackupRunStatus.Set(nil)
}

// UnsetLastBackupRunStatus ensures that no value is present for LastBackupRunStatus, not even an explicit nil
func (o *ProtectionSummary) UnsetLastBackupRunStatus() {
	o.LastBackupRunStatus.Unset()
}

// GetLastArchivalRunStatus returns the LastArchivalRunStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetLastArchivalRunStatus() string {
	if o == nil || o.LastArchivalRunStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastArchivalRunStatus.Get()
}

// GetLastArchivalRunStatusOk returns a tuple with the LastArchivalRunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetLastArchivalRunStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastArchivalRunStatus.Get(), o.LastArchivalRunStatus.IsSet()
}

// HasLastArchivalRunStatus returns a boolean if a field has been set.
func (o *ProtectionSummary) HasLastArchivalRunStatus() bool {
	if o != nil && o.LastArchivalRunStatus.IsSet() {
		return true
	}

	return false
}

// SetLastArchivalRunStatus gets a reference to the given NullableString and assigns it to the LastArchivalRunStatus field.
func (o *ProtectionSummary) SetLastArchivalRunStatus(v string) {
	o.LastArchivalRunStatus.Set(&v)
}
// SetLastArchivalRunStatusNil sets the value for LastArchivalRunStatus to be an explicit nil
func (o *ProtectionSummary) SetLastArchivalRunStatusNil() {
	o.LastArchivalRunStatus.Set(nil)
}

// UnsetLastArchivalRunStatus ensures that no value is present for LastArchivalRunStatus, not even an explicit nil
func (o *ProtectionSummary) UnsetLastArchivalRunStatus() {
	o.LastArchivalRunStatus.Unset()
}

// GetLastReplicationRunStatus returns the LastReplicationRunStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetLastReplicationRunStatus() string {
	if o == nil || o.LastReplicationRunStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastReplicationRunStatus.Get()
}

// GetLastReplicationRunStatusOk returns a tuple with the LastReplicationRunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetLastReplicationRunStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastReplicationRunStatus.Get(), o.LastReplicationRunStatus.IsSet()
}

// HasLastReplicationRunStatus returns a boolean if a field has been set.
func (o *ProtectionSummary) HasLastReplicationRunStatus() bool {
	if o != nil && o.LastReplicationRunStatus.IsSet() {
		return true
	}

	return false
}

// SetLastReplicationRunStatus gets a reference to the given NullableString and assigns it to the LastReplicationRunStatus field.
func (o *ProtectionSummary) SetLastReplicationRunStatus(v string) {
	o.LastReplicationRunStatus.Set(&v)
}
// SetLastReplicationRunStatusNil sets the value for LastReplicationRunStatus to be an explicit nil
func (o *ProtectionSummary) SetLastReplicationRunStatusNil() {
	o.LastReplicationRunStatus.Set(nil)
}

// UnsetLastReplicationRunStatus ensures that no value is present for LastReplicationRunStatus, not even an explicit nil
func (o *ProtectionSummary) UnsetLastReplicationRunStatus() {
	o.LastReplicationRunStatus.Unset()
}

// GetLastRunSlaViolated returns the LastRunSlaViolated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSummary) GetLastRunSlaViolated() bool {
	if o == nil || o.LastRunSlaViolated.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LastRunSlaViolated.Get()
}

// GetLastRunSlaViolatedOk returns a tuple with the LastRunSlaViolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSummary) GetLastRunSlaViolatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastRunSlaViolated.Get(), o.LastRunSlaViolated.IsSet()
}

// HasLastRunSlaViolated returns a boolean if a field has been set.
func (o *ProtectionSummary) HasLastRunSlaViolated() bool {
	if o != nil && o.LastRunSlaViolated.IsSet() {
		return true
	}

	return false
}

// SetLastRunSlaViolated gets a reference to the given NullableBool and assigns it to the LastRunSlaViolated field.
func (o *ProtectionSummary) SetLastRunSlaViolated(v bool) {
	o.LastRunSlaViolated.Set(&v)
}
// SetLastRunSlaViolatedNil sets the value for LastRunSlaViolated to be an explicit nil
func (o *ProtectionSummary) SetLastRunSlaViolatedNil() {
	o.LastRunSlaViolated.Set(nil)
}

// UnsetLastRunSlaViolated ensures that no value is present for LastRunSlaViolated, not even an explicit nil
func (o *ProtectionSummary) UnsetLastRunSlaViolated() {
	o.LastRunSlaViolated.Unset()
}

func (o ProtectionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicyName.IsSet() {
		toSerialize["policyName"] = o.PolicyName.Get()
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.LastBackupRunStatus.IsSet() {
		toSerialize["lastBackupRunStatus"] = o.LastBackupRunStatus.Get()
	}
	if o.LastArchivalRunStatus.IsSet() {
		toSerialize["lastArchivalRunStatus"] = o.LastArchivalRunStatus.Get()
	}
	if o.LastReplicationRunStatus.IsSet() {
		toSerialize["lastReplicationRunStatus"] = o.LastReplicationRunStatus.Get()
	}
	if o.LastRunSlaViolated.IsSet() {
		toSerialize["lastRunSlaViolated"] = o.LastRunSlaViolated.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionSummary struct {
	value *ProtectionSummary
	isSet bool
}

func (v NullableProtectionSummary) Get() *ProtectionSummary {
	return v.value
}

func (v *NullableProtectionSummary) Set(val *ProtectionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionSummary(val *ProtectionSummary) *NullableProtectionSummary {
	return &NullableProtectionSummary{value: val, isSet: true}
}

func (v NullableProtectionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


