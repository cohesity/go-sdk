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

// checks if the HeliosPolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosPolicyResponse{}

// HeliosPolicyResponse Specifies the details about the Policy.
type HeliosPolicyResponse struct {
	BackupPolicy *HeliosBackupPolicy `json:"backupPolicy,omitempty"`
	// List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod.
	BlackoutWindow []HeliosBlackoutWindow `json:"blackoutWindow,omitempty"`
	// Specifies the cluster to which this policy belongs. This required is only for type OnPremPolicy. The format is clusterId:clusterIncarnationId.
	ClusterIdentifier NullableString `json:"clusterIdentifier,omitempty" validate:"regexp=^([0-9]+:[0-9]+)$"`
	// This field is now deprecated. Please use the DataLockConfig in the backup retention.
	DataLock NullableString `json:"dataLock,omitempty"`
	// Specifies the description of the Protection Policy.
	Description NullableString `json:"description,omitempty"`
	// Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
	ExtendedRetention []HeliosExtendedRetentionPolicy `json:"extendedRetention,omitempty"`
	// Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature.
	IsCBSEnabled NullableBool `json:"isCBSEnabled,omitempty"`
	// Specifies the name of the Protection Policy.
	Name NullableString `json:"name"`
	RemoteTargetPolicy *HeliosTargetsConfiguration `json:"remoteTargetPolicy,omitempty"`
	RetryOptions *HeliosRetryOptions `json:"retryOptions,omitempty"`
	// Specifies the tenants which have access to this object.
	TenantIds []*string `json:"tenantIds,omitempty"`
	// Specifies the type of the Protection Policy to be created on Helios.
	Type NullableString `json:"type"`
	// Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases.
	Version NullableInt32 `json:"version,omitempty"`
	// Specifies a unique policy id assigned by the Helios.
	Id NullableString `json:"id,omitempty"`
	// In case of global policy response, specifies the number of policies linked to this global policy on the cluster.
	NumLinkedPolicies NullableInt64 `json:"numLinkedPolicies,omitempty"`
	// Specifies the number of object protections using the protection policy.
	NumObjectProtections NullableInt64 `json:"numObjectProtections,omitempty"`
}

type _HeliosPolicyResponse HeliosPolicyResponse

// NewHeliosPolicyResponse instantiates a new HeliosPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosPolicyResponse(name NullableString, type_ NullableString) *HeliosPolicyResponse {
	this := HeliosPolicyResponse{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewHeliosPolicyResponseWithDefaults instantiates a new HeliosPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosPolicyResponseWithDefaults() *HeliosPolicyResponse {
	this := HeliosPolicyResponse{}
	return &this
}

// GetBackupPolicy returns the BackupPolicy field value if set, zero value otherwise.
func (o *HeliosPolicyResponse) GetBackupPolicy() HeliosBackupPolicy {
	if o == nil || IsNil(o.BackupPolicy) {
		var ret HeliosBackupPolicy
		return ret
	}
	return *o.BackupPolicy
}

// GetBackupPolicyOk returns a tuple with the BackupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosPolicyResponse) GetBackupPolicyOk() (*HeliosBackupPolicy, bool) {
	if o == nil || IsNil(o.BackupPolicy) {
		return nil, false
	}
	return o.BackupPolicy, true
}

// HasBackupPolicy returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasBackupPolicy() bool {
	if o != nil && !IsNil(o.BackupPolicy) {
		return true
	}

	return false
}

// SetBackupPolicy gets a reference to the given HeliosBackupPolicy and assigns it to the BackupPolicy field.
func (o *HeliosPolicyResponse) SetBackupPolicy(v HeliosBackupPolicy) {
	o.BackupPolicy = &v
}

// GetBlackoutWindow returns the BlackoutWindow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetBlackoutWindow() []HeliosBlackoutWindow {
	if o == nil {
		var ret []HeliosBlackoutWindow
		return ret
	}
	return o.BlackoutWindow
}

// GetBlackoutWindowOk returns a tuple with the BlackoutWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetBlackoutWindowOk() ([]HeliosBlackoutWindow, bool) {
	if o == nil || IsNil(o.BlackoutWindow) {
		return nil, false
	}
	return o.BlackoutWindow, true
}

// HasBlackoutWindow returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasBlackoutWindow() bool {
	if o != nil && !IsNil(o.BlackoutWindow) {
		return true
	}

	return false
}

// SetBlackoutWindow gets a reference to the given []HeliosBlackoutWindow and assigns it to the BlackoutWindow field.
func (o *HeliosPolicyResponse) SetBlackoutWindow(v []HeliosBlackoutWindow) {
	o.BlackoutWindow = v
}

// GetClusterIdentifier returns the ClusterIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetClusterIdentifier() string {
	if o == nil || IsNil(o.ClusterIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ClusterIdentifier.Get()
}

// GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetClusterIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterIdentifier.Get(), o.ClusterIdentifier.IsSet()
}

// HasClusterIdentifier returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasClusterIdentifier() bool {
	if o != nil && o.ClusterIdentifier.IsSet() {
		return true
	}

	return false
}

// SetClusterIdentifier gets a reference to the given NullableString and assigns it to the ClusterIdentifier field.
func (o *HeliosPolicyResponse) SetClusterIdentifier(v string) {
	o.ClusterIdentifier.Set(&v)
}
// SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil
func (o *HeliosPolicyResponse) SetClusterIdentifierNil() {
	o.ClusterIdentifier.Set(nil)
}

// UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetClusterIdentifier() {
	o.ClusterIdentifier.Unset()
}

// GetDataLock returns the DataLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetDataLock() string {
	if o == nil || IsNil(o.DataLock.Get()) {
		var ret string
		return ret
	}
	return *o.DataLock.Get()
}

// GetDataLockOk returns a tuple with the DataLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetDataLockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataLock.Get(), o.DataLock.IsSet()
}

// HasDataLock returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasDataLock() bool {
	if o != nil && o.DataLock.IsSet() {
		return true
	}

	return false
}

// SetDataLock gets a reference to the given NullableString and assigns it to the DataLock field.
func (o *HeliosPolicyResponse) SetDataLock(v string) {
	o.DataLock.Set(&v)
}
// SetDataLockNil sets the value for DataLock to be an explicit nil
func (o *HeliosPolicyResponse) SetDataLockNil() {
	o.DataLock.Set(nil)
}

// UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetDataLock() {
	o.DataLock.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *HeliosPolicyResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *HeliosPolicyResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetExtendedRetention returns the ExtendedRetention field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetExtendedRetention() []HeliosExtendedRetentionPolicy {
	if o == nil {
		var ret []HeliosExtendedRetentionPolicy
		return ret
	}
	return o.ExtendedRetention
}

// GetExtendedRetentionOk returns a tuple with the ExtendedRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetExtendedRetentionOk() ([]HeliosExtendedRetentionPolicy, bool) {
	if o == nil || IsNil(o.ExtendedRetention) {
		return nil, false
	}
	return o.ExtendedRetention, true
}

// HasExtendedRetention returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasExtendedRetention() bool {
	if o != nil && !IsNil(o.ExtendedRetention) {
		return true
	}

	return false
}

// SetExtendedRetention gets a reference to the given []HeliosExtendedRetentionPolicy and assigns it to the ExtendedRetention field.
func (o *HeliosPolicyResponse) SetExtendedRetention(v []HeliosExtendedRetentionPolicy) {
	o.ExtendedRetention = v
}

// GetIsCBSEnabled returns the IsCBSEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetIsCBSEnabled() bool {
	if o == nil || IsNil(o.IsCBSEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCBSEnabled.Get()
}

// GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetIsCBSEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCBSEnabled.Get(), o.IsCBSEnabled.IsSet()
}

// HasIsCBSEnabled returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasIsCBSEnabled() bool {
	if o != nil && o.IsCBSEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsCBSEnabled gets a reference to the given NullableBool and assigns it to the IsCBSEnabled field.
func (o *HeliosPolicyResponse) SetIsCBSEnabled(v bool) {
	o.IsCBSEnabled.Set(&v)
}
// SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil
func (o *HeliosPolicyResponse) SetIsCBSEnabledNil() {
	o.IsCBSEnabled.Set(nil)
}

// UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetIsCBSEnabled() {
	o.IsCBSEnabled.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HeliosPolicyResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *HeliosPolicyResponse) SetName(v string) {
	o.Name.Set(&v)
}

// GetRemoteTargetPolicy returns the RemoteTargetPolicy field value if set, zero value otherwise.
func (o *HeliosPolicyResponse) GetRemoteTargetPolicy() HeliosTargetsConfiguration {
	if o == nil || IsNil(o.RemoteTargetPolicy) {
		var ret HeliosTargetsConfiguration
		return ret
	}
	return *o.RemoteTargetPolicy
}

// GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosPolicyResponse) GetRemoteTargetPolicyOk() (*HeliosTargetsConfiguration, bool) {
	if o == nil || IsNil(o.RemoteTargetPolicy) {
		return nil, false
	}
	return o.RemoteTargetPolicy, true
}

// HasRemoteTargetPolicy returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasRemoteTargetPolicy() bool {
	if o != nil && !IsNil(o.RemoteTargetPolicy) {
		return true
	}

	return false
}

// SetRemoteTargetPolicy gets a reference to the given HeliosTargetsConfiguration and assigns it to the RemoteTargetPolicy field.
func (o *HeliosPolicyResponse) SetRemoteTargetPolicy(v HeliosTargetsConfiguration) {
	o.RemoteTargetPolicy = &v
}

// GetRetryOptions returns the RetryOptions field value if set, zero value otherwise.
func (o *HeliosPolicyResponse) GetRetryOptions() HeliosRetryOptions {
	if o == nil || IsNil(o.RetryOptions) {
		var ret HeliosRetryOptions
		return ret
	}
	return *o.RetryOptions
}

// GetRetryOptionsOk returns a tuple with the RetryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosPolicyResponse) GetRetryOptionsOk() (*HeliosRetryOptions, bool) {
	if o == nil || IsNil(o.RetryOptions) {
		return nil, false
	}
	return o.RetryOptions, true
}

// HasRetryOptions returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasRetryOptions() bool {
	if o != nil && !IsNil(o.RetryOptions) {
		return true
	}

	return false
}

// SetRetryOptions gets a reference to the given HeliosRetryOptions and assigns it to the RetryOptions field.
func (o *HeliosPolicyResponse) SetRetryOptions(v HeliosRetryOptions) {
	o.RetryOptions = &v
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise.
func (o *HeliosPolicyResponse) GetTenantIds() []*string {
	if o == nil || IsNil(o.TenantIds) {
		var ret []*string
		return ret
	}
	return o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosPolicyResponse) GetTenantIdsOk() ([]*string, bool) {
	if o == nil || IsNil(o.TenantIds) {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasTenantIds() bool {
	if o != nil && !IsNil(o.TenantIds) {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []*string and assigns it to the TenantIds field.
func (o *HeliosPolicyResponse) SetTenantIds(v []*string) {
	o.TenantIds = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HeliosPolicyResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *HeliosPolicyResponse) SetType(v string) {
	o.Type.Set(&v)
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetVersion() int32 {
	if o == nil || IsNil(o.Version.Get()) {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *HeliosPolicyResponse) SetVersion(v int32) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *HeliosPolicyResponse) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetVersion() {
	o.Version.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HeliosPolicyResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HeliosPolicyResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetId() {
	o.Id.Unset()
}

// GetNumLinkedPolicies returns the NumLinkedPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetNumLinkedPolicies() int64 {
	if o == nil || IsNil(o.NumLinkedPolicies.Get()) {
		var ret int64
		return ret
	}
	return *o.NumLinkedPolicies.Get()
}

// GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetNumLinkedPoliciesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumLinkedPolicies.Get(), o.NumLinkedPolicies.IsSet()
}

// HasNumLinkedPolicies returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasNumLinkedPolicies() bool {
	if o != nil && o.NumLinkedPolicies.IsSet() {
		return true
	}

	return false
}

// SetNumLinkedPolicies gets a reference to the given NullableInt64 and assigns it to the NumLinkedPolicies field.
func (o *HeliosPolicyResponse) SetNumLinkedPolicies(v int64) {
	o.NumLinkedPolicies.Set(&v)
}
// SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil
func (o *HeliosPolicyResponse) SetNumLinkedPoliciesNil() {
	o.NumLinkedPolicies.Set(nil)
}

// UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetNumLinkedPolicies() {
	o.NumLinkedPolicies.Unset()
}

// GetNumObjectProtections returns the NumObjectProtections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponse) GetNumObjectProtections() int64 {
	if o == nil || IsNil(o.NumObjectProtections.Get()) {
		var ret int64
		return ret
	}
	return *o.NumObjectProtections.Get()
}

// GetNumObjectProtectionsOk returns a tuple with the NumObjectProtections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponse) GetNumObjectProtectionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumObjectProtections.Get(), o.NumObjectProtections.IsSet()
}

// HasNumObjectProtections returns a boolean if a field has been set.
func (o *HeliosPolicyResponse) HasNumObjectProtections() bool {
	if o != nil && o.NumObjectProtections.IsSet() {
		return true
	}

	return false
}

// SetNumObjectProtections gets a reference to the given NullableInt64 and assigns it to the NumObjectProtections field.
func (o *HeliosPolicyResponse) SetNumObjectProtections(v int64) {
	o.NumObjectProtections.Set(&v)
}
// SetNumObjectProtectionsNil sets the value for NumObjectProtections to be an explicit nil
func (o *HeliosPolicyResponse) SetNumObjectProtectionsNil() {
	o.NumObjectProtections.Set(nil)
}

// UnsetNumObjectProtections ensures that no value is present for NumObjectProtections, not even an explicit nil
func (o *HeliosPolicyResponse) UnsetNumObjectProtections() {
	o.NumObjectProtections.Unset()
}

func (o HeliosPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupPolicy) {
		toSerialize["backupPolicy"] = o.BackupPolicy
	}
	if o.BlackoutWindow != nil {
		toSerialize["blackoutWindow"] = o.BlackoutWindow
	}
	if o.ClusterIdentifier.IsSet() {
		toSerialize["clusterIdentifier"] = o.ClusterIdentifier.Get()
	}
	if o.DataLock.IsSet() {
		toSerialize["dataLock"] = o.DataLock.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExtendedRetention != nil {
		toSerialize["extendedRetention"] = o.ExtendedRetention
	}
	if o.IsCBSEnabled.IsSet() {
		toSerialize["isCBSEnabled"] = o.IsCBSEnabled.Get()
	}
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.RemoteTargetPolicy) {
		toSerialize["remoteTargetPolicy"] = o.RemoteTargetPolicy
	}
	if !IsNil(o.RetryOptions) {
		toSerialize["retryOptions"] = o.RetryOptions
	}
	if !IsNil(o.TenantIds) {
		toSerialize["tenantIds"] = o.TenantIds
	}
	toSerialize["type"] = o.Type.Get()
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.NumLinkedPolicies.IsSet() {
		toSerialize["numLinkedPolicies"] = o.NumLinkedPolicies.Get()
	}
	if o.NumObjectProtections.IsSet() {
		toSerialize["numObjectProtections"] = o.NumObjectProtections.Get()
	}
	return toSerialize, nil
}

func (o *HeliosPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varHeliosPolicyResponse := _HeliosPolicyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeliosPolicyResponse)

	if err != nil {
		return err
	}

	*o = HeliosPolicyResponse(varHeliosPolicyResponse)

	return err
}

type NullableHeliosPolicyResponse struct {
	value *HeliosPolicyResponse
	isSet bool
}

func (v NullableHeliosPolicyResponse) Get() *HeliosPolicyResponse {
	return v.value
}

func (v *NullableHeliosPolicyResponse) Set(val *HeliosPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosPolicyResponse(val *HeliosPolicyResponse) *NullableHeliosPolicyResponse {
	return &NullableHeliosPolicyResponse{value: val, isSet: true}
}

func (v NullableHeliosPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


