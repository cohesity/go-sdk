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

// ProtectionGroupAlertingPolicy Specifies a policy for alerting users of the status of a Protection Group.
type ProtectionGroupAlertingPolicy struct {
	// Specifies the run status for which the user would like to receive alerts.
	BackupRunStatus []string `json:"backupRunStatus"`
	// Specifies a list of targets to receive the alerts.
	AlertTargets *[]AlertTarget `json:"alertTargets,omitempty"`
}

// NewProtectionGroupAlertingPolicy instantiates a new ProtectionGroupAlertingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionGroupAlertingPolicy(backupRunStatus []string) *ProtectionGroupAlertingPolicy {
	this := ProtectionGroupAlertingPolicy{}
	this.BackupRunStatus = backupRunStatus
	return &this
}

// NewProtectionGroupAlertingPolicyWithDefaults instantiates a new ProtectionGroupAlertingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionGroupAlertingPolicyWithDefaults() *ProtectionGroupAlertingPolicy {
	this := ProtectionGroupAlertingPolicy{}
	return &this
}

// GetBackupRunStatus returns the BackupRunStatus field value
func (o *ProtectionGroupAlertingPolicy) GetBackupRunStatus() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackupRunStatus
}

// GetBackupRunStatusOk returns a tuple with the BackupRunStatus field value
// and a boolean to check if the value has been set.
func (o *ProtectionGroupAlertingPolicy) GetBackupRunStatusOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupRunStatus, true
}

// SetBackupRunStatus sets field value
func (o *ProtectionGroupAlertingPolicy) SetBackupRunStatus(v []string) {
	o.BackupRunStatus = v
}

// GetAlertTargets returns the AlertTargets field value if set, zero value otherwise.
func (o *ProtectionGroupAlertingPolicy) GetAlertTargets() []AlertTarget {
	if o == nil || o.AlertTargets == nil {
		var ret []AlertTarget
		return ret
	}
	return *o.AlertTargets
}

// GetAlertTargetsOk returns a tuple with the AlertTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectionGroupAlertingPolicy) GetAlertTargetsOk() (*[]AlertTarget, bool) {
	if o == nil || o.AlertTargets == nil {
		return nil, false
	}
	return o.AlertTargets, true
}

// HasAlertTargets returns a boolean if a field has been set.
func (o *ProtectionGroupAlertingPolicy) HasAlertTargets() bool {
	if o != nil && o.AlertTargets != nil {
		return true
	}

	return false
}

// SetAlertTargets gets a reference to the given []AlertTarget and assigns it to the AlertTargets field.
func (o *ProtectionGroupAlertingPolicy) SetAlertTargets(v []AlertTarget) {
	o.AlertTargets = &v
}

func (o ProtectionGroupAlertingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["backupRunStatus"] = o.BackupRunStatus
	}
	if o.AlertTargets != nil {
		toSerialize["alertTargets"] = o.AlertTargets
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionGroupAlertingPolicy struct {
	value *ProtectionGroupAlertingPolicy
	isSet bool
}

func (v NullableProtectionGroupAlertingPolicy) Get() *ProtectionGroupAlertingPolicy {
	return v.value
}

func (v *NullableProtectionGroupAlertingPolicy) Set(val *ProtectionGroupAlertingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionGroupAlertingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionGroupAlertingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionGroupAlertingPolicy(val *ProtectionGroupAlertingPolicy) *NullableProtectionGroupAlertingPolicy {
	return &NullableProtectionGroupAlertingPolicy{value: val, isSet: true}
}

func (v NullableProtectionGroupAlertingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionGroupAlertingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ProtectionGroupAlertingPolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}