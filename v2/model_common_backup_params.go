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

// checks if the CommonBackupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBackupParams{}

// CommonBackupParams Specifies the common parameters for backup. These parameters are common across protection group and object protection.
type CommonBackupParams struct {
	// Specifies whether currently executing object backup should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false.
	AbortInBlackouts NullableBool `json:"abortInBlackouts,omitempty"`
	// Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	PolicyConfig *PolicyConfig `json:"policyConfig,omitempty"`
	// Specifies the unique id of the Protection Policy. The Policy settings will be attached with every object and will be used in backup.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Specifies the priority for the objects backup.
	Priority NullableString `json:"priority,omitempty"`
	// Specifies whether object backup will be written to HDD or SSD.
	QosPolicy NullableString `json:"qosPolicy,omitempty"`
	// Specifies whether to skip Rigel for backup or not.
	SkipRigelForBackup NullableBool `json:"skipRigelForBackup,omitempty"`
	// Specifies the SLA parameters for list of objects.
	Sla []SlaRule `json:"sla,omitempty"`
	StartTime *TimeOfDay `json:"startTime,omitempty"`
	// Specifies the Storage Domain (View Box) ID where the object backup will be taken. This is not required if Cloud archive direct is benig used.
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
}

// NewCommonBackupParams instantiates a new CommonBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBackupParams() *CommonBackupParams {
	this := CommonBackupParams{}
	return &this
}

// NewCommonBackupParamsWithDefaults instantiates a new CommonBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBackupParamsWithDefaults() *CommonBackupParams {
	this := CommonBackupParams{}
	return &this
}

// GetAbortInBlackouts returns the AbortInBlackouts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetAbortInBlackouts() bool {
	if o == nil || IsNil(o.AbortInBlackouts.Get()) {
		var ret bool
		return ret
	}
	return *o.AbortInBlackouts.Get()
}

// GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetAbortInBlackoutsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbortInBlackouts.Get(), o.AbortInBlackouts.IsSet()
}

// HasAbortInBlackouts returns a boolean if a field has been set.
func (o *CommonBackupParams) HasAbortInBlackouts() bool {
	if o != nil && o.AbortInBlackouts.IsSet() {
		return true
	}

	return false
}

// SetAbortInBlackouts gets a reference to the given NullableBool and assigns it to the AbortInBlackouts field.
func (o *CommonBackupParams) SetAbortInBlackouts(v bool) {
	o.AbortInBlackouts.Set(&v)
}
// SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil
func (o *CommonBackupParams) SetAbortInBlackoutsNil() {
	o.AbortInBlackouts.Set(nil)
}

// UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
func (o *CommonBackupParams) UnsetAbortInBlackouts() {
	o.AbortInBlackouts.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetEndTimeUsecs() int64 {
	if o == nil || IsNil(o.EndTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *CommonBackupParams) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *CommonBackupParams) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *CommonBackupParams) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *CommonBackupParams) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetPolicyConfig returns the PolicyConfig field value if set, zero value otherwise.
func (o *CommonBackupParams) GetPolicyConfig() PolicyConfig {
	if o == nil || IsNil(o.PolicyConfig) {
		var ret PolicyConfig
		return ret
	}
	return *o.PolicyConfig
}

// GetPolicyConfigOk returns a tuple with the PolicyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBackupParams) GetPolicyConfigOk() (*PolicyConfig, bool) {
	if o == nil || IsNil(o.PolicyConfig) {
		return nil, false
	}
	return o.PolicyConfig, true
}

// HasPolicyConfig returns a boolean if a field has been set.
func (o *CommonBackupParams) HasPolicyConfig() bool {
	if o != nil && !IsNil(o.PolicyConfig) {
		return true
	}

	return false
}

// SetPolicyConfig gets a reference to the given PolicyConfig and assigns it to the PolicyConfig field.
func (o *CommonBackupParams) SetPolicyConfig(v PolicyConfig) {
	o.PolicyConfig = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *CommonBackupParams) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *CommonBackupParams) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *CommonBackupParams) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *CommonBackupParams) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetPriority() string {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret string
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CommonBackupParams) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableString and assigns it to the Priority field.
func (o *CommonBackupParams) SetPriority(v string) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CommonBackupParams) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CommonBackupParams) UnsetPriority() {
	o.Priority.Unset()
}

// GetQosPolicy returns the QosPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetQosPolicy() string {
	if o == nil || IsNil(o.QosPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.QosPolicy.Get()
}

// GetQosPolicyOk returns a tuple with the QosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetQosPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosPolicy.Get(), o.QosPolicy.IsSet()
}

// HasQosPolicy returns a boolean if a field has been set.
func (o *CommonBackupParams) HasQosPolicy() bool {
	if o != nil && o.QosPolicy.IsSet() {
		return true
	}

	return false
}

// SetQosPolicy gets a reference to the given NullableString and assigns it to the QosPolicy field.
func (o *CommonBackupParams) SetQosPolicy(v string) {
	o.QosPolicy.Set(&v)
}
// SetQosPolicyNil sets the value for QosPolicy to be an explicit nil
func (o *CommonBackupParams) SetQosPolicyNil() {
	o.QosPolicy.Set(nil)
}

// UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
func (o *CommonBackupParams) UnsetQosPolicy() {
	o.QosPolicy.Unset()
}

// GetSkipRigelForBackup returns the SkipRigelForBackup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetSkipRigelForBackup() bool {
	if o == nil || IsNil(o.SkipRigelForBackup.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipRigelForBackup.Get()
}

// GetSkipRigelForBackupOk returns a tuple with the SkipRigelForBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetSkipRigelForBackupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipRigelForBackup.Get(), o.SkipRigelForBackup.IsSet()
}

// HasSkipRigelForBackup returns a boolean if a field has been set.
func (o *CommonBackupParams) HasSkipRigelForBackup() bool {
	if o != nil && o.SkipRigelForBackup.IsSet() {
		return true
	}

	return false
}

// SetSkipRigelForBackup gets a reference to the given NullableBool and assigns it to the SkipRigelForBackup field.
func (o *CommonBackupParams) SetSkipRigelForBackup(v bool) {
	o.SkipRigelForBackup.Set(&v)
}
// SetSkipRigelForBackupNil sets the value for SkipRigelForBackup to be an explicit nil
func (o *CommonBackupParams) SetSkipRigelForBackupNil() {
	o.SkipRigelForBackup.Set(nil)
}

// UnsetSkipRigelForBackup ensures that no value is present for SkipRigelForBackup, not even an explicit nil
func (o *CommonBackupParams) UnsetSkipRigelForBackup() {
	o.SkipRigelForBackup.Unset()
}

// GetSla returns the Sla field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetSla() []SlaRule {
	if o == nil {
		var ret []SlaRule
		return ret
	}
	return o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetSlaOk() ([]SlaRule, bool) {
	if o == nil || IsNil(o.Sla) {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *CommonBackupParams) HasSla() bool {
	if o != nil && !IsNil(o.Sla) {
		return true
	}

	return false
}

// SetSla gets a reference to the given []SlaRule and assigns it to the Sla field.
func (o *CommonBackupParams) SetSla(v []SlaRule) {
	o.Sla = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CommonBackupParams) GetStartTime() TimeOfDay {
	if o == nil || IsNil(o.StartTime) {
		var ret TimeOfDay
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBackupParams) GetStartTimeOk() (*TimeOfDay, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CommonBackupParams) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given TimeOfDay and assigns it to the StartTime field.
func (o *CommonBackupParams) SetStartTime(v TimeOfDay) {
	o.StartTime = &v
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonBackupParams) GetStorageDomainId() int64 {
	if o == nil || IsNil(o.StorageDomainId.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonBackupParams) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *CommonBackupParams) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *CommonBackupParams) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *CommonBackupParams) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *CommonBackupParams) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

func (o CommonBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBackupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AbortInBlackouts.IsSet() {
		toSerialize["abortInBlackouts"] = o.AbortInBlackouts.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if !IsNil(o.PolicyConfig) {
		toSerialize["policyConfig"] = o.PolicyConfig
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.QosPolicy.IsSet() {
		toSerialize["qosPolicy"] = o.QosPolicy.Get()
	}
	if o.SkipRigelForBackup.IsSet() {
		toSerialize["skipRigelForBackup"] = o.SkipRigelForBackup.Get()
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

type NullableCommonBackupParams struct {
	value *CommonBackupParams
	isSet bool
}

func (v NullableCommonBackupParams) Get() *CommonBackupParams {
	return v.value
}

func (v *NullableCommonBackupParams) Set(val *CommonBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBackupParams(val *CommonBackupParams) *NullableCommonBackupParams {
	return &NullableCommonBackupParams{value: val, isSet: true}
}

func (v NullableCommonBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


