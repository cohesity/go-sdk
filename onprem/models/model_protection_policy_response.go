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

// ProtectionPolicyResponse Specifies the details about the Protection Policy.
type ProtectionPolicyResponse struct {
	// Specifies the name of the Protection Policy.
	Name NullableString `json:"name"`
	BackupPolicy BackupPolicy `json:"backupPolicy"`
	// Specifies the description of the Protection Policy.
	Description NullableString `json:"description,omitempty"`
	// List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod.
	BlackoutWindow []BlackoutWindow `json:"blackoutWindow,omitempty"`
	// Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
	ExtendedRetention []ExtendedRetentionPolicy `json:"extendedRetention,omitempty"`
	RemoteTargetPolicy *TargetsConfiguration `json:"remoteTargetPolicy,omitempty"`
	RetryOptions *RetryOptions `json:"retryOptions,omitempty"`
	// This field is now deprecated. Please use the DataLockConfig in the backup retention.
	DataLock NullableString `json:"dataLock,omitempty"`
	// Specifies a unique Policy id assigned by the Cohesity Cluster.
	Id NullableString `json:"id,omitempty"`
	// Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template.
	TemplateId NullableString `json:"templateId,omitempty"`
	// This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated.
	IsUsable NullableBool `json:"isUsable,omitempty"`
	// This field is set to true when policy is the replicated policy.
	IsReplicated NullableBool `json:"isReplicated,omitempty"`
	// Specifies the number of protection groups using the protection policy.
	NumProtectionGroups NullableInt64 `json:"numProtectionGroups,omitempty"`
	// Specifies the number of protected objects using the protection policy.
	NumProtectedObjects NullableInt64 `json:"numProtectedObjects,omitempty"`
}

// NewProtectionPolicyResponse instantiates a new ProtectionPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionPolicyResponse(name NullableString, backupPolicy BackupPolicy) *ProtectionPolicyResponse {
	this := ProtectionPolicyResponse{}
	this.Name = name
	this.BackupPolicy = backupPolicy
	return &this
}

// NewProtectionPolicyResponseWithDefaults instantiates a new ProtectionPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionPolicyResponseWithDefaults() *ProtectionPolicyResponse {
	this := ProtectionPolicyResponse{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProtectionPolicyResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ProtectionPolicyResponse) SetName(v string) {
	o.Name.Set(&v)
}

// GetBackupPolicy returns the BackupPolicy field value
func (o *ProtectionPolicyResponse) GetBackupPolicy() BackupPolicy {
	if o == nil {
		var ret BackupPolicy
		return ret
	}

	return o.BackupPolicy
}

// GetBackupPolicyOk returns a tuple with the BackupPolicy field value
// and a boolean to check if the value has been set.
func (o *ProtectionPolicyResponse) GetBackupPolicyOk() (*BackupPolicy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BackupPolicy, true
}

// SetBackupPolicy sets field value
func (o *ProtectionPolicyResponse) SetBackupPolicy(v BackupPolicy) {
	o.BackupPolicy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProtectionPolicyResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProtectionPolicyResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetBlackoutWindow returns the BlackoutWindow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetBlackoutWindow() []BlackoutWindow {
	if o == nil  {
		var ret []BlackoutWindow
		return ret
	}
	return o.BlackoutWindow
}

// GetBlackoutWindowOk returns a tuple with the BlackoutWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetBlackoutWindowOk() (*[]BlackoutWindow, bool) {
	if o == nil || o.BlackoutWindow == nil {
		return nil, false
	}
	return &o.BlackoutWindow, true
}

// HasBlackoutWindow returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasBlackoutWindow() bool {
	if o != nil && o.BlackoutWindow != nil {
		return true
	}

	return false
}

// SetBlackoutWindow gets a reference to the given []BlackoutWindow and assigns it to the BlackoutWindow field.
func (o *ProtectionPolicyResponse) SetBlackoutWindow(v []BlackoutWindow) {
	o.BlackoutWindow = v
}

// GetExtendedRetention returns the ExtendedRetention field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetExtendedRetention() []ExtendedRetentionPolicy {
	if o == nil  {
		var ret []ExtendedRetentionPolicy
		return ret
	}
	return o.ExtendedRetention
}

// GetExtendedRetentionOk returns a tuple with the ExtendedRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetExtendedRetentionOk() (*[]ExtendedRetentionPolicy, bool) {
	if o == nil || o.ExtendedRetention == nil {
		return nil, false
	}
	return &o.ExtendedRetention, true
}

// HasExtendedRetention returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasExtendedRetention() bool {
	if o != nil && o.ExtendedRetention != nil {
		return true
	}

	return false
}

// SetExtendedRetention gets a reference to the given []ExtendedRetentionPolicy and assigns it to the ExtendedRetention field.
func (o *ProtectionPolicyResponse) SetExtendedRetention(v []ExtendedRetentionPolicy) {
	o.ExtendedRetention = v
}

// GetRemoteTargetPolicy returns the RemoteTargetPolicy field value if set, zero value otherwise.
func (o *ProtectionPolicyResponse) GetRemoteTargetPolicy() TargetsConfiguration {
	if o == nil || o.RemoteTargetPolicy == nil {
		var ret TargetsConfiguration
		return ret
	}
	return *o.RemoteTargetPolicy
}

// GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectionPolicyResponse) GetRemoteTargetPolicyOk() (*TargetsConfiguration, bool) {
	if o == nil || o.RemoteTargetPolicy == nil {
		return nil, false
	}
	return o.RemoteTargetPolicy, true
}

// HasRemoteTargetPolicy returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasRemoteTargetPolicy() bool {
	if o != nil && o.RemoteTargetPolicy != nil {
		return true
	}

	return false
}

// SetRemoteTargetPolicy gets a reference to the given TargetsConfiguration and assigns it to the RemoteTargetPolicy field.
func (o *ProtectionPolicyResponse) SetRemoteTargetPolicy(v TargetsConfiguration) {
	o.RemoteTargetPolicy = &v
}

// GetRetryOptions returns the RetryOptions field value if set, zero value otherwise.
func (o *ProtectionPolicyResponse) GetRetryOptions() RetryOptions {
	if o == nil || o.RetryOptions == nil {
		var ret RetryOptions
		return ret
	}
	return *o.RetryOptions
}

// GetRetryOptionsOk returns a tuple with the RetryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectionPolicyResponse) GetRetryOptionsOk() (*RetryOptions, bool) {
	if o == nil || o.RetryOptions == nil {
		return nil, false
	}
	return o.RetryOptions, true
}

// HasRetryOptions returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasRetryOptions() bool {
	if o != nil && o.RetryOptions != nil {
		return true
	}

	return false
}

// SetRetryOptions gets a reference to the given RetryOptions and assigns it to the RetryOptions field.
func (o *ProtectionPolicyResponse) SetRetryOptions(v RetryOptions) {
	o.RetryOptions = &v
}

// GetDataLock returns the DataLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetDataLock() string {
	if o == nil || o.DataLock.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataLock.Get()
}

// GetDataLockOk returns a tuple with the DataLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetDataLockOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataLock.Get(), o.DataLock.IsSet()
}

// HasDataLock returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasDataLock() bool {
	if o != nil && o.DataLock.IsSet() {
		return true
	}

	return false
}

// SetDataLock gets a reference to the given NullableString and assigns it to the DataLock field.
func (o *ProtectionPolicyResponse) SetDataLock(v string) {
	o.DataLock.Set(&v)
}
// SetDataLockNil sets the value for DataLock to be an explicit nil
func (o *ProtectionPolicyResponse) SetDataLockNil() {
	o.DataLock.Set(nil)
}

// UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetDataLock() {
	o.DataLock.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ProtectionPolicyResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProtectionPolicyResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetId() {
	o.Id.Unset()
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *ProtectionPolicyResponse) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *ProtectionPolicyResponse) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetIsUsable returns the IsUsable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetIsUsable() bool {
	if o == nil || o.IsUsable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsUsable.Get()
}

// GetIsUsableOk returns a tuple with the IsUsable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetIsUsableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsUsable.Get(), o.IsUsable.IsSet()
}

// HasIsUsable returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasIsUsable() bool {
	if o != nil && o.IsUsable.IsSet() {
		return true
	}

	return false
}

// SetIsUsable gets a reference to the given NullableBool and assigns it to the IsUsable field.
func (o *ProtectionPolicyResponse) SetIsUsable(v bool) {
	o.IsUsable.Set(&v)
}
// SetIsUsableNil sets the value for IsUsable to be an explicit nil
func (o *ProtectionPolicyResponse) SetIsUsableNil() {
	o.IsUsable.Set(nil)
}

// UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetIsUsable() {
	o.IsUsable.Unset()
}

// GetIsReplicated returns the IsReplicated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetIsReplicated() bool {
	if o == nil || o.IsReplicated.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsReplicated.Get()
}

// GetIsReplicatedOk returns a tuple with the IsReplicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetIsReplicatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsReplicated.Get(), o.IsReplicated.IsSet()
}

// HasIsReplicated returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasIsReplicated() bool {
	if o != nil && o.IsReplicated.IsSet() {
		return true
	}

	return false
}

// SetIsReplicated gets a reference to the given NullableBool and assigns it to the IsReplicated field.
func (o *ProtectionPolicyResponse) SetIsReplicated(v bool) {
	o.IsReplicated.Set(&v)
}
// SetIsReplicatedNil sets the value for IsReplicated to be an explicit nil
func (o *ProtectionPolicyResponse) SetIsReplicatedNil() {
	o.IsReplicated.Set(nil)
}

// UnsetIsReplicated ensures that no value is present for IsReplicated, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetIsReplicated() {
	o.IsReplicated.Unset()
}

// GetNumProtectionGroups returns the NumProtectionGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetNumProtectionGroups() int64 {
	if o == nil || o.NumProtectionGroups.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumProtectionGroups.Get()
}

// GetNumProtectionGroupsOk returns a tuple with the NumProtectionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetNumProtectionGroupsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumProtectionGroups.Get(), o.NumProtectionGroups.IsSet()
}

// HasNumProtectionGroups returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasNumProtectionGroups() bool {
	if o != nil && o.NumProtectionGroups.IsSet() {
		return true
	}

	return false
}

// SetNumProtectionGroups gets a reference to the given NullableInt64 and assigns it to the NumProtectionGroups field.
func (o *ProtectionPolicyResponse) SetNumProtectionGroups(v int64) {
	o.NumProtectionGroups.Set(&v)
}
// SetNumProtectionGroupsNil sets the value for NumProtectionGroups to be an explicit nil
func (o *ProtectionPolicyResponse) SetNumProtectionGroupsNil() {
	o.NumProtectionGroups.Set(nil)
}

// UnsetNumProtectionGroups ensures that no value is present for NumProtectionGroups, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetNumProtectionGroups() {
	o.NumProtectionGroups.Unset()
}

// GetNumProtectedObjects returns the NumProtectedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionPolicyResponse) GetNumProtectedObjects() int64 {
	if o == nil || o.NumProtectedObjects.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumProtectedObjects.Get()
}

// GetNumProtectedObjectsOk returns a tuple with the NumProtectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionPolicyResponse) GetNumProtectedObjectsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumProtectedObjects.Get(), o.NumProtectedObjects.IsSet()
}

// HasNumProtectedObjects returns a boolean if a field has been set.
func (o *ProtectionPolicyResponse) HasNumProtectedObjects() bool {
	if o != nil && o.NumProtectedObjects.IsSet() {
		return true
	}

	return false
}

// SetNumProtectedObjects gets a reference to the given NullableInt64 and assigns it to the NumProtectedObjects field.
func (o *ProtectionPolicyResponse) SetNumProtectedObjects(v int64) {
	o.NumProtectedObjects.Set(&v)
}
// SetNumProtectedObjectsNil sets the value for NumProtectedObjects to be an explicit nil
func (o *ProtectionPolicyResponse) SetNumProtectedObjectsNil() {
	o.NumProtectedObjects.Set(nil)
}

// UnsetNumProtectedObjects ensures that no value is present for NumProtectedObjects, not even an explicit nil
func (o *ProtectionPolicyResponse) UnsetNumProtectedObjects() {
	o.NumProtectedObjects.Unset()
}

func (o ProtectionPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["backupPolicy"] = o.BackupPolicy
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.BlackoutWindow != nil {
		toSerialize["blackoutWindow"] = o.BlackoutWindow
	}
	if o.ExtendedRetention != nil {
		toSerialize["extendedRetention"] = o.ExtendedRetention
	}
	if o.RemoteTargetPolicy != nil {
		toSerialize["remoteTargetPolicy"] = o.RemoteTargetPolicy
	}
	if o.RetryOptions != nil {
		toSerialize["retryOptions"] = o.RetryOptions
	}
	if o.DataLock.IsSet() {
		toSerialize["dataLock"] = o.DataLock.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.IsUsable.IsSet() {
		toSerialize["isUsable"] = o.IsUsable.Get()
	}
	if o.IsReplicated.IsSet() {
		toSerialize["isReplicated"] = o.IsReplicated.Get()
	}
	if o.NumProtectionGroups.IsSet() {
		toSerialize["numProtectionGroups"] = o.NumProtectionGroups.Get()
	}
	if o.NumProtectedObjects.IsSet() {
		toSerialize["numProtectedObjects"] = o.NumProtectedObjects.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionPolicyResponse struct {
	value *ProtectionPolicyResponse
	isSet bool
}

func (v NullableProtectionPolicyResponse) Get() *ProtectionPolicyResponse {
	return v.value
}

func (v *NullableProtectionPolicyResponse) Set(val *ProtectionPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionPolicyResponse(val *ProtectionPolicyResponse) *NullableProtectionPolicyResponse {
	return &NullableProtectionPolicyResponse{value: val, isSet: true}
}

func (v NullableProtectionPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ProtectionPolicyResponse) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}