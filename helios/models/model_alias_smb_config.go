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

// AliasSmbConfig Message defining SMB config for IRIS. SMB config contains SMB encryption flags, SMB discoverable flag and Share level permissions.
type AliasSmbConfig struct {
	// Whether the share is discoverable.
	DiscoveryEnabled NullableBool `json:"discoveryEnabled,omitempty"`
	// Whether SMB encryption is enabled for this share. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	// Whether to enforce encryption for all the sessions for this view. When enabled all unencrypted sessions are disallowed.
	EncryptionRequired NullableBool `json:"encryptionRequired,omitempty"`
	// Share level permissions.
	Permissions []SmbPermission `json:"permissions,omitempty"`
	// Specifies a list of super user sids.
	SuperUserSids []string `json:"superUserSids,omitempty"`
	// Indicate if offline file caching is supported.
	CachingEnabled NullableBool `json:"cachingEnabled,omitempty"`
	// Indicate if share level permission is cleared by user.
	IsShareLevelPermissionEmpty NullableBool `json:"isShareLevelPermissionEmpty,omitempty"`
	// Indicate the operation lock is enabled by this view.
	OplockEnabled NullableBool `json:"oplockEnabled,omitempty"`
	// Whether file open handles are persited to scribe to survive bridge process crash. When set to false, open handles will be kept in memory untill the current node has exclusive ticket for the entity handle. When the entity is opened from another node, the exclusive ticket would be revoked from the node. In revoke control flow, the current node would persist the state to scribe. On acquiring the exclusive ticket,another node would read the file open handles from scribe and resume the handling of operation.
	ContinuousAvailability NullableBool `json:"continuousAvailability,omitempty"`
}

// NewAliasSmbConfig instantiates a new AliasSmbConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAliasSmbConfig() *AliasSmbConfig {
	this := AliasSmbConfig{}
	return &this
}

// NewAliasSmbConfigWithDefaults instantiates a new AliasSmbConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasSmbConfigWithDefaults() *AliasSmbConfig {
	this := AliasSmbConfig{}
	return &this
}

// GetDiscoveryEnabled returns the DiscoveryEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetDiscoveryEnabled() bool {
	if o == nil || o.DiscoveryEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DiscoveryEnabled.Get()
}

// GetDiscoveryEnabledOk returns a tuple with the DiscoveryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetDiscoveryEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiscoveryEnabled.Get(), o.DiscoveryEnabled.IsSet()
}

// HasDiscoveryEnabled returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasDiscoveryEnabled() bool {
	if o != nil && o.DiscoveryEnabled.IsSet() {
		return true
	}

	return false
}

// SetDiscoveryEnabled gets a reference to the given NullableBool and assigns it to the DiscoveryEnabled field.
func (o *AliasSmbConfig) SetDiscoveryEnabled(v bool) {
	o.DiscoveryEnabled.Set(&v)
}
// SetDiscoveryEnabledNil sets the value for DiscoveryEnabled to be an explicit nil
func (o *AliasSmbConfig) SetDiscoveryEnabledNil() {
	o.DiscoveryEnabled.Set(nil)
}

// UnsetDiscoveryEnabled ensures that no value is present for DiscoveryEnabled, not even an explicit nil
func (o *AliasSmbConfig) UnsetDiscoveryEnabled() {
	o.DiscoveryEnabled.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *AliasSmbConfig) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *AliasSmbConfig) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *AliasSmbConfig) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetEncryptionRequired returns the EncryptionRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetEncryptionRequired() bool {
	if o == nil || o.EncryptionRequired.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionRequired.Get()
}

// GetEncryptionRequiredOk returns a tuple with the EncryptionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetEncryptionRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionRequired.Get(), o.EncryptionRequired.IsSet()
}

// HasEncryptionRequired returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasEncryptionRequired() bool {
	if o != nil && o.EncryptionRequired.IsSet() {
		return true
	}

	return false
}

// SetEncryptionRequired gets a reference to the given NullableBool and assigns it to the EncryptionRequired field.
func (o *AliasSmbConfig) SetEncryptionRequired(v bool) {
	o.EncryptionRequired.Set(&v)
}
// SetEncryptionRequiredNil sets the value for EncryptionRequired to be an explicit nil
func (o *AliasSmbConfig) SetEncryptionRequiredNil() {
	o.EncryptionRequired.Set(nil)
}

// UnsetEncryptionRequired ensures that no value is present for EncryptionRequired, not even an explicit nil
func (o *AliasSmbConfig) UnsetEncryptionRequired() {
	o.EncryptionRequired.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetPermissions() []SmbPermission {
	if o == nil  {
		var ret []SmbPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetPermissionsOk() (*[]SmbPermission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []SmbPermission and assigns it to the Permissions field.
func (o *AliasSmbConfig) SetPermissions(v []SmbPermission) {
	o.Permissions = v
}

// GetSuperUserSids returns the SuperUserSids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetSuperUserSids() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SuperUserSids
}

// GetSuperUserSidsOk returns a tuple with the SuperUserSids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetSuperUserSidsOk() (*[]string, bool) {
	if o == nil || o.SuperUserSids == nil {
		return nil, false
	}
	return &o.SuperUserSids, true
}

// HasSuperUserSids returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasSuperUserSids() bool {
	if o != nil && o.SuperUserSids != nil {
		return true
	}

	return false
}

// SetSuperUserSids gets a reference to the given []string and assigns it to the SuperUserSids field.
func (o *AliasSmbConfig) SetSuperUserSids(v []string) {
	o.SuperUserSids = v
}

// GetCachingEnabled returns the CachingEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetCachingEnabled() bool {
	if o == nil || o.CachingEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CachingEnabled.Get()
}

// GetCachingEnabledOk returns a tuple with the CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetCachingEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CachingEnabled.Get(), o.CachingEnabled.IsSet()
}

// HasCachingEnabled returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasCachingEnabled() bool {
	if o != nil && o.CachingEnabled.IsSet() {
		return true
	}

	return false
}

// SetCachingEnabled gets a reference to the given NullableBool and assigns it to the CachingEnabled field.
func (o *AliasSmbConfig) SetCachingEnabled(v bool) {
	o.CachingEnabled.Set(&v)
}
// SetCachingEnabledNil sets the value for CachingEnabled to be an explicit nil
func (o *AliasSmbConfig) SetCachingEnabledNil() {
	o.CachingEnabled.Set(nil)
}

// UnsetCachingEnabled ensures that no value is present for CachingEnabled, not even an explicit nil
func (o *AliasSmbConfig) UnsetCachingEnabled() {
	o.CachingEnabled.Unset()
}

// GetIsShareLevelPermissionEmpty returns the IsShareLevelPermissionEmpty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetIsShareLevelPermissionEmpty() bool {
	if o == nil || o.IsShareLevelPermissionEmpty.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsShareLevelPermissionEmpty.Get()
}

// GetIsShareLevelPermissionEmptyOk returns a tuple with the IsShareLevelPermissionEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetIsShareLevelPermissionEmptyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsShareLevelPermissionEmpty.Get(), o.IsShareLevelPermissionEmpty.IsSet()
}

// HasIsShareLevelPermissionEmpty returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasIsShareLevelPermissionEmpty() bool {
	if o != nil && o.IsShareLevelPermissionEmpty.IsSet() {
		return true
	}

	return false
}

// SetIsShareLevelPermissionEmpty gets a reference to the given NullableBool and assigns it to the IsShareLevelPermissionEmpty field.
func (o *AliasSmbConfig) SetIsShareLevelPermissionEmpty(v bool) {
	o.IsShareLevelPermissionEmpty.Set(&v)
}
// SetIsShareLevelPermissionEmptyNil sets the value for IsShareLevelPermissionEmpty to be an explicit nil
func (o *AliasSmbConfig) SetIsShareLevelPermissionEmptyNil() {
	o.IsShareLevelPermissionEmpty.Set(nil)
}

// UnsetIsShareLevelPermissionEmpty ensures that no value is present for IsShareLevelPermissionEmpty, not even an explicit nil
func (o *AliasSmbConfig) UnsetIsShareLevelPermissionEmpty() {
	o.IsShareLevelPermissionEmpty.Unset()
}

// GetOplockEnabled returns the OplockEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetOplockEnabled() bool {
	if o == nil || o.OplockEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OplockEnabled.Get()
}

// GetOplockEnabledOk returns a tuple with the OplockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetOplockEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OplockEnabled.Get(), o.OplockEnabled.IsSet()
}

// HasOplockEnabled returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasOplockEnabled() bool {
	if o != nil && o.OplockEnabled.IsSet() {
		return true
	}

	return false
}

// SetOplockEnabled gets a reference to the given NullableBool and assigns it to the OplockEnabled field.
func (o *AliasSmbConfig) SetOplockEnabled(v bool) {
	o.OplockEnabled.Set(&v)
}
// SetOplockEnabledNil sets the value for OplockEnabled to be an explicit nil
func (o *AliasSmbConfig) SetOplockEnabledNil() {
	o.OplockEnabled.Set(nil)
}

// UnsetOplockEnabled ensures that no value is present for OplockEnabled, not even an explicit nil
func (o *AliasSmbConfig) UnsetOplockEnabled() {
	o.OplockEnabled.Unset()
}

// GetContinuousAvailability returns the ContinuousAvailability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AliasSmbConfig) GetContinuousAvailability() bool {
	if o == nil || o.ContinuousAvailability.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinuousAvailability.Get()
}

// GetContinuousAvailabilityOk returns a tuple with the ContinuousAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AliasSmbConfig) GetContinuousAvailabilityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinuousAvailability.Get(), o.ContinuousAvailability.IsSet()
}

// HasContinuousAvailability returns a boolean if a field has been set.
func (o *AliasSmbConfig) HasContinuousAvailability() bool {
	if o != nil && o.ContinuousAvailability.IsSet() {
		return true
	}

	return false
}

// SetContinuousAvailability gets a reference to the given NullableBool and assigns it to the ContinuousAvailability field.
func (o *AliasSmbConfig) SetContinuousAvailability(v bool) {
	o.ContinuousAvailability.Set(&v)
}
// SetContinuousAvailabilityNil sets the value for ContinuousAvailability to be an explicit nil
func (o *AliasSmbConfig) SetContinuousAvailabilityNil() {
	o.ContinuousAvailability.Set(nil)
}

// UnsetContinuousAvailability ensures that no value is present for ContinuousAvailability, not even an explicit nil
func (o *AliasSmbConfig) UnsetContinuousAvailability() {
	o.ContinuousAvailability.Unset()
}

func (o AliasSmbConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiscoveryEnabled.IsSet() {
		toSerialize["discoveryEnabled"] = o.DiscoveryEnabled.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if o.EncryptionRequired.IsSet() {
		toSerialize["encryptionRequired"] = o.EncryptionRequired.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.SuperUserSids != nil {
		toSerialize["superUserSids"] = o.SuperUserSids
	}
	if o.CachingEnabled.IsSet() {
		toSerialize["cachingEnabled"] = o.CachingEnabled.Get()
	}
	if o.IsShareLevelPermissionEmpty.IsSet() {
		toSerialize["isShareLevelPermissionEmpty"] = o.IsShareLevelPermissionEmpty.Get()
	}
	if o.OplockEnabled.IsSet() {
		toSerialize["oplockEnabled"] = o.OplockEnabled.Get()
	}
	if o.ContinuousAvailability.IsSet() {
		toSerialize["continuousAvailability"] = o.ContinuousAvailability.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAliasSmbConfig struct {
	value *AliasSmbConfig
	isSet bool
}

func (v NullableAliasSmbConfig) Get() *AliasSmbConfig {
	return v.value
}

func (v *NullableAliasSmbConfig) Set(val *AliasSmbConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAliasSmbConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAliasSmbConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAliasSmbConfig(val *AliasSmbConfig) *NullableAliasSmbConfig {
	return &NullableAliasSmbConfig{value: val, isSet: true}
}

func (v NullableAliasSmbConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAliasSmbConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


