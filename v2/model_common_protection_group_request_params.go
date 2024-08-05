/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonProtectionGroupRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonProtectionGroupRequestParams{}

// CommonProtectionGroupRequestParams Specifies the parameters which are common between all Protection Group requests.
type CommonProtectionGroupRequestParams struct {
	// Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'pauseInBlackouts' is set to true.
	AbortInBlackouts NullableBool `json:"abortInBlackouts,omitempty"`
	// Specifies the advanced configuration for a protection job.
	AdvancedConfigs []KeyValuePair `json:"advancedConfigs,omitempty"`
	AlertPolicy *ProtectionGroupAlertingPolicy `json:"alertPolicy,omitempty"`
	// Specifies a description of the Protection Group.
	Description NullableString `json:"description,omitempty"`
	// Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specifies the environment type of the Protection Group.
	Environment NullableString `json:"environment"`
	// Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted.
	IsPaused NullableBool `json:"isPaused,omitempty"`
	// Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.
	LastModifiedTimestampUsecs NullableInt64 `json:"lastModifiedTimestampUsecs,omitempty"`
	// Specifies the name of the Protection Group.
	Name NullableString `json:"name"`
	// Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true.
	PauseInBlackouts NullableBool `json:"pauseInBlackouts,omitempty"`
	// Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc.
	PolicyId NullableString `json:"policyId"`
	// Specifies the priority of the Protection Group.
	Priority NullableString `json:"priority,omitempty"`
	// Specifies whether the Protection Group will be written to HDD or SSD.
	QosPolicy NullableString `json:"qosPolicy,omitempty"`
	// Specifies the SLA parameters for this Protection Group.
	Sla []SlaRule `json:"sla,omitempty"`
	StartTime *TimeOfDay `json:"startTime,omitempty"`
	// Specifies the Storage Domain (View Box) ID where this Protection Group writes data.
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
}

type _CommonProtectionGroupRequestParams CommonProtectionGroupRequestParams

// NewCommonProtectionGroupRequestParams instantiates a new CommonProtectionGroupRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonProtectionGroupRequestParams(environment NullableString, name NullableString, policyId NullableString) *CommonProtectionGroupRequestParams {
	this := CommonProtectionGroupRequestParams{}
	this.Environment = environment
	this.Name = name
	this.PolicyId = policyId
	return &this
}

// NewCommonProtectionGroupRequestParamsWithDefaults instantiates a new CommonProtectionGroupRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonProtectionGroupRequestParamsWithDefaults() *CommonProtectionGroupRequestParams {
	this := CommonProtectionGroupRequestParams{}
	return &this
}

// GetAbortInBlackouts returns the AbortInBlackouts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetAbortInBlackouts() bool {
	if o == nil || IsNil(o.AbortInBlackouts.Get()) {
		var ret bool
		return ret
	}
	return *o.AbortInBlackouts.Get()
}

// GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetAbortInBlackoutsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbortInBlackouts.Get(), o.AbortInBlackouts.IsSet()
}

// HasAbortInBlackouts returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasAbortInBlackouts() bool {
	if o != nil && o.AbortInBlackouts.IsSet() {
		return true
	}

	return false
}

// SetAbortInBlackouts gets a reference to the given NullableBool and assigns it to the AbortInBlackouts field.
func (o *CommonProtectionGroupRequestParams) SetAbortInBlackouts(v bool) {
	o.AbortInBlackouts.Set(&v)
}
// SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetAbortInBlackoutsNil() {
	o.AbortInBlackouts.Set(nil)
}

// UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetAbortInBlackouts() {
	o.AbortInBlackouts.Unset()
}

// GetAdvancedConfigs returns the AdvancedConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetAdvancedConfigs() []KeyValuePair {
	if o == nil {
		var ret []KeyValuePair
		return ret
	}
	return o.AdvancedConfigs
}

// GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetAdvancedConfigsOk() ([]KeyValuePair, bool) {
	if o == nil || IsNil(o.AdvancedConfigs) {
		return nil, false
	}
	return o.AdvancedConfigs, true
}

// HasAdvancedConfigs returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasAdvancedConfigs() bool {
	if o != nil && !IsNil(o.AdvancedConfigs) {
		return true
	}

	return false
}

// SetAdvancedConfigs gets a reference to the given []KeyValuePair and assigns it to the AdvancedConfigs field.
func (o *CommonProtectionGroupRequestParams) SetAdvancedConfigs(v []KeyValuePair) {
	o.AdvancedConfigs = v
}

// GetAlertPolicy returns the AlertPolicy field value if set, zero value otherwise.
func (o *CommonProtectionGroupRequestParams) GetAlertPolicy() ProtectionGroupAlertingPolicy {
	if o == nil || IsNil(o.AlertPolicy) {
		var ret ProtectionGroupAlertingPolicy
		return ret
	}
	return *o.AlertPolicy
}

// GetAlertPolicyOk returns a tuple with the AlertPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonProtectionGroupRequestParams) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool) {
	if o == nil || IsNil(o.AlertPolicy) {
		return nil, false
	}
	return o.AlertPolicy, true
}

// HasAlertPolicy returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasAlertPolicy() bool {
	if o != nil && !IsNil(o.AlertPolicy) {
		return true
	}

	return false
}

// SetAlertPolicy gets a reference to the given ProtectionGroupAlertingPolicy and assigns it to the AlertPolicy field.
func (o *CommonProtectionGroupRequestParams) SetAlertPolicy(v ProtectionGroupAlertingPolicy) {
	o.AlertPolicy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CommonProtectionGroupRequestParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetDescription() {
	o.Description.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetEndTimeUsecs() int64 {
	if o == nil || IsNil(o.EndTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *CommonProtectionGroupRequestParams) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetEnvironment returns the Environment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonProtectionGroupRequestParams) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// SetEnvironment sets field value
func (o *CommonProtectionGroupRequestParams) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPaused.Get()
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetIsPausedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPaused.Get(), o.IsPaused.IsSet()
}

// HasIsPaused returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasIsPaused() bool {
	if o != nil && o.IsPaused.IsSet() {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given NullableBool and assigns it to the IsPaused field.
func (o *CommonProtectionGroupRequestParams) SetIsPaused(v bool) {
	o.IsPaused.Set(&v)
}
// SetIsPausedNil sets the value for IsPaused to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetIsPausedNil() {
	o.IsPaused.Set(nil)
}

// UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetIsPaused() {
	o.IsPaused.Unset()
}

// GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetLastModifiedTimestampUsecs() int64 {
	if o == nil || IsNil(o.LastModifiedTimestampUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.LastModifiedTimestampUsecs.Get()
}

// GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetLastModifiedTimestampUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedTimestampUsecs.Get(), o.LastModifiedTimestampUsecs.IsSet()
}

// HasLastModifiedTimestampUsecs returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasLastModifiedTimestampUsecs() bool {
	if o != nil && o.LastModifiedTimestampUsecs.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedTimestampUsecs gets a reference to the given NullableInt64 and assigns it to the LastModifiedTimestampUsecs field.
func (o *CommonProtectionGroupRequestParams) SetLastModifiedTimestampUsecs(v int64) {
	o.LastModifiedTimestampUsecs.Set(&v)
}
// SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetLastModifiedTimestampUsecsNil() {
	o.LastModifiedTimestampUsecs.Set(nil)
}

// UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetLastModifiedTimestampUsecs() {
	o.LastModifiedTimestampUsecs.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonProtectionGroupRequestParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CommonProtectionGroupRequestParams) SetName(v string) {
	o.Name.Set(&v)
}

// GetPauseInBlackouts returns the PauseInBlackouts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetPauseInBlackouts() bool {
	if o == nil || IsNil(o.PauseInBlackouts.Get()) {
		var ret bool
		return ret
	}
	return *o.PauseInBlackouts.Get()
}

// GetPauseInBlackoutsOk returns a tuple with the PauseInBlackouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetPauseInBlackoutsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PauseInBlackouts.Get(), o.PauseInBlackouts.IsSet()
}

// HasPauseInBlackouts returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasPauseInBlackouts() bool {
	if o != nil && o.PauseInBlackouts.IsSet() {
		return true
	}

	return false
}

// SetPauseInBlackouts gets a reference to the given NullableBool and assigns it to the PauseInBlackouts field.
func (o *CommonProtectionGroupRequestParams) SetPauseInBlackouts(v bool) {
	o.PauseInBlackouts.Set(&v)
}
// SetPauseInBlackoutsNil sets the value for PauseInBlackouts to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetPauseInBlackoutsNil() {
	o.PauseInBlackouts.Set(nil)
}

// UnsetPauseInBlackouts ensures that no value is present for PauseInBlackouts, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetPauseInBlackouts() {
	o.PauseInBlackouts.Unset()
}

// GetPolicyId returns the PolicyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonProtectionGroupRequestParams) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// SetPolicyId sets field value
func (o *CommonProtectionGroupRequestParams) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetPriority() string {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret string
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableString and assigns it to the Priority field.
func (o *CommonProtectionGroupRequestParams) SetPriority(v string) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetPriority() {
	o.Priority.Unset()
}

// GetQosPolicy returns the QosPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetQosPolicy() string {
	if o == nil || IsNil(o.QosPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.QosPolicy.Get()
}

// GetQosPolicyOk returns a tuple with the QosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetQosPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosPolicy.Get(), o.QosPolicy.IsSet()
}

// HasQosPolicy returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasQosPolicy() bool {
	if o != nil && o.QosPolicy.IsSet() {
		return true
	}

	return false
}

// SetQosPolicy gets a reference to the given NullableString and assigns it to the QosPolicy field.
func (o *CommonProtectionGroupRequestParams) SetQosPolicy(v string) {
	o.QosPolicy.Set(&v)
}
// SetQosPolicyNil sets the value for QosPolicy to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetQosPolicyNil() {
	o.QosPolicy.Set(nil)
}

// UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetQosPolicy() {
	o.QosPolicy.Unset()
}

// GetSla returns the Sla field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetSla() []SlaRule {
	if o == nil {
		var ret []SlaRule
		return ret
	}
	return o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetSlaOk() ([]SlaRule, bool) {
	if o == nil || IsNil(o.Sla) {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasSla() bool {
	if o != nil && !IsNil(o.Sla) {
		return true
	}

	return false
}

// SetSla gets a reference to the given []SlaRule and assigns it to the Sla field.
func (o *CommonProtectionGroupRequestParams) SetSla(v []SlaRule) {
	o.Sla = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CommonProtectionGroupRequestParams) GetStartTime() TimeOfDay {
	if o == nil || IsNil(o.StartTime) {
		var ret TimeOfDay
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonProtectionGroupRequestParams) GetStartTimeOk() (*TimeOfDay, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given TimeOfDay and assigns it to the StartTime field.
func (o *CommonProtectionGroupRequestParams) SetStartTime(v TimeOfDay) {
	o.StartTime = &v
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonProtectionGroupRequestParams) GetStorageDomainId() int64 {
	if o == nil || IsNil(o.StorageDomainId.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonProtectionGroupRequestParams) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *CommonProtectionGroupRequestParams) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *CommonProtectionGroupRequestParams) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *CommonProtectionGroupRequestParams) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *CommonProtectionGroupRequestParams) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

func (o CommonProtectionGroupRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonProtectionGroupRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AbortInBlackouts.IsSet() {
		toSerialize["abortInBlackouts"] = o.AbortInBlackouts.Get()
	}
	if o.AdvancedConfigs != nil {
		toSerialize["advancedConfigs"] = o.AdvancedConfigs
	}
	if !IsNil(o.AlertPolicy) {
		toSerialize["alertPolicy"] = o.AlertPolicy
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	toSerialize["environment"] = o.Environment.Get()
	if o.IsPaused.IsSet() {
		toSerialize["isPaused"] = o.IsPaused.Get()
	}
	if o.LastModifiedTimestampUsecs.IsSet() {
		toSerialize["lastModifiedTimestampUsecs"] = o.LastModifiedTimestampUsecs.Get()
	}
	toSerialize["name"] = o.Name.Get()
	if o.PauseInBlackouts.IsSet() {
		toSerialize["pauseInBlackouts"] = o.PauseInBlackouts.Get()
	}
	toSerialize["policyId"] = o.PolicyId.Get()
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.QosPolicy.IsSet() {
		toSerialize["qosPolicy"] = o.QosPolicy.Get()
	}
	if o.Sla != nil {
		toSerialize["sla"] = o.Sla
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	return toSerialize, nil
}

func (o *CommonProtectionGroupRequestParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"name",
		"policyId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonProtectionGroupRequestParams := _CommonProtectionGroupRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonProtectionGroupRequestParams)

	if err != nil {
		return err
	}

	*o = CommonProtectionGroupRequestParams(varCommonProtectionGroupRequestParams)

	return err
}

type NullableCommonProtectionGroupRequestParams struct {
	value *CommonProtectionGroupRequestParams
	isSet bool
}

func (v NullableCommonProtectionGroupRequestParams) Get() *CommonProtectionGroupRequestParams {
	return v.value
}

func (v *NullableCommonProtectionGroupRequestParams) Set(val *CommonProtectionGroupRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonProtectionGroupRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonProtectionGroupRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonProtectionGroupRequestParams(val *CommonProtectionGroupRequestParams) *NullableCommonProtectionGroupRequestParams {
	return &NullableCommonProtectionGroupRequestParams{value: val, isSet: true}
}

func (v NullableCommonProtectionGroupRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonProtectionGroupRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


