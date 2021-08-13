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

// CommonMSSQLProtectionGroupParams Specifies the common params to create a MSSQL Protection Group.
type CommonMSSQLProtectionGroupParams struct {
	// Specifies the preference type for backing up user databases on the host.
	UserDbBackupPreferenceType NullableString `json:"userDbBackupPreferenceType,omitempty"`
	// Specifies whether to backup system databases. If not specified then parameter is set to true.
	BackupSystemDbs NullableBool `json:"backupSystemDbs,omitempty"`
	// Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used.
	UseAagPreferencesFromServer NullableBool `json:"useAagPreferencesFromServer,omitempty"`
	// Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases.
	AagBackupPreferenceType NullableString `json:"aagBackupPreferenceType,omitempty"`
	// Specifies whether full backups should be copy-only.
	FullBackupsCopyOnly NullableBool `json:"fullBackupsCopyOnly,omitempty"`
	PrePostScript *PrePostScriptParams `json:"prePostScript,omitempty"`
	// Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
	ExcludeFilters []Filter `json:"excludeFilters,omitempty"`
}

// NewCommonMSSQLProtectionGroupParams instantiates a new CommonMSSQLProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonMSSQLProtectionGroupParams() *CommonMSSQLProtectionGroupParams {
	this := CommonMSSQLProtectionGroupParams{}
	return &this
}

// NewCommonMSSQLProtectionGroupParamsWithDefaults instantiates a new CommonMSSQLProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonMSSQLProtectionGroupParamsWithDefaults() *CommonMSSQLProtectionGroupParams {
	this := CommonMSSQLProtectionGroupParams{}
	return &this
}

// GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceType() string {
	if o == nil || o.UserDbBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDbBackupPreferenceType.Get()
}

// GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDbBackupPreferenceType.Get(), o.UserDbBackupPreferenceType.IsSet()
}

// HasUserDbBackupPreferenceType returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasUserDbBackupPreferenceType() bool {
	if o != nil && o.UserDbBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetUserDbBackupPreferenceType gets a reference to the given NullableString and assigns it to the UserDbBackupPreferenceType field.
func (o *CommonMSSQLProtectionGroupParams) SetUserDbBackupPreferenceType(v string) {
	o.UserDbBackupPreferenceType.Set(&v)
}
// SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetUserDbBackupPreferenceTypeNil() {
	o.UserDbBackupPreferenceType.Set(nil)
}

// UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetUserDbBackupPreferenceType() {
	o.UserDbBackupPreferenceType.Unset()
}

// GetBackupSystemDbs returns the BackupSystemDbs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbs() bool {
	if o == nil || o.BackupSystemDbs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupSystemDbs.Get()
}

// GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupSystemDbs.Get(), o.BackupSystemDbs.IsSet()
}

// HasBackupSystemDbs returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasBackupSystemDbs() bool {
	if o != nil && o.BackupSystemDbs.IsSet() {
		return true
	}

	return false
}

// SetBackupSystemDbs gets a reference to the given NullableBool and assigns it to the BackupSystemDbs field.
func (o *CommonMSSQLProtectionGroupParams) SetBackupSystemDbs(v bool) {
	o.BackupSystemDbs.Set(&v)
}
// SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetBackupSystemDbsNil() {
	o.BackupSystemDbs.Set(nil)
}

// UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetBackupSystemDbs() {
	o.BackupSystemDbs.Unset()
}

// GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServer() bool {
	if o == nil || o.UseAagPreferencesFromServer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseAagPreferencesFromServer.Get()
}

// GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseAagPreferencesFromServer.Get(), o.UseAagPreferencesFromServer.IsSet()
}

// HasUseAagPreferencesFromServer returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasUseAagPreferencesFromServer() bool {
	if o != nil && o.UseAagPreferencesFromServer.IsSet() {
		return true
	}

	return false
}

// SetUseAagPreferencesFromServer gets a reference to the given NullableBool and assigns it to the UseAagPreferencesFromServer field.
func (o *CommonMSSQLProtectionGroupParams) SetUseAagPreferencesFromServer(v bool) {
	o.UseAagPreferencesFromServer.Set(&v)
}
// SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetUseAagPreferencesFromServerNil() {
	o.UseAagPreferencesFromServer.Set(nil)
}

// UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetUseAagPreferencesFromServer() {
	o.UseAagPreferencesFromServer.Unset()
}

// GetAagBackupPreferenceType returns the AagBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceType() string {
	if o == nil || o.AagBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AagBackupPreferenceType.Get()
}

// GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AagBackupPreferenceType.Get(), o.AagBackupPreferenceType.IsSet()
}

// HasAagBackupPreferenceType returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasAagBackupPreferenceType() bool {
	if o != nil && o.AagBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetAagBackupPreferenceType gets a reference to the given NullableString and assigns it to the AagBackupPreferenceType field.
func (o *CommonMSSQLProtectionGroupParams) SetAagBackupPreferenceType(v string) {
	o.AagBackupPreferenceType.Set(&v)
}
// SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetAagBackupPreferenceTypeNil() {
	o.AagBackupPreferenceType.Set(nil)
}

// UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetAagBackupPreferenceType() {
	o.AagBackupPreferenceType.Unset()
}

// GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnly() bool {
	if o == nil || o.FullBackupsCopyOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FullBackupsCopyOnly.Get()
}

// GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullBackupsCopyOnly.Get(), o.FullBackupsCopyOnly.IsSet()
}

// HasFullBackupsCopyOnly returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasFullBackupsCopyOnly() bool {
	if o != nil && o.FullBackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetFullBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the FullBackupsCopyOnly field.
func (o *CommonMSSQLProtectionGroupParams) SetFullBackupsCopyOnly(v bool) {
	o.FullBackupsCopyOnly.Set(&v)
}
// SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetFullBackupsCopyOnlyNil() {
	o.FullBackupsCopyOnly.Set(nil)
}

// UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetFullBackupsCopyOnly() {
	o.FullBackupsCopyOnly.Unset()
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *CommonMSSQLProtectionGroupParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMSSQLProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *CommonMSSQLProtectionGroupParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetExcludeFilters returns the ExcludeFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetExcludeFilters() []Filter {
	if o == nil  {
		var ret []Filter
		return ret
	}
	return o.ExcludeFilters
}

// GetExcludeFiltersOk returns a tuple with the ExcludeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetExcludeFiltersOk() (*[]Filter, bool) {
	if o == nil || o.ExcludeFilters == nil {
		return nil, false
	}
	return &o.ExcludeFilters, true
}

// HasExcludeFilters returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasExcludeFilters() bool {
	if o != nil && o.ExcludeFilters != nil {
		return true
	}

	return false
}

// SetExcludeFilters gets a reference to the given []Filter and assigns it to the ExcludeFilters field.
func (o *CommonMSSQLProtectionGroupParams) SetExcludeFilters(v []Filter) {
	o.ExcludeFilters = v
}

func (o CommonMSSQLProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserDbBackupPreferenceType.IsSet() {
		toSerialize["userDbBackupPreferenceType"] = o.UserDbBackupPreferenceType.Get()
	}
	if o.BackupSystemDbs.IsSet() {
		toSerialize["backupSystemDbs"] = o.BackupSystemDbs.Get()
	}
	if o.UseAagPreferencesFromServer.IsSet() {
		toSerialize["useAagPreferencesFromServer"] = o.UseAagPreferencesFromServer.Get()
	}
	if o.AagBackupPreferenceType.IsSet() {
		toSerialize["aagBackupPreferenceType"] = o.AagBackupPreferenceType.Get()
	}
	if o.FullBackupsCopyOnly.IsSet() {
		toSerialize["fullBackupsCopyOnly"] = o.FullBackupsCopyOnly.Get()
	}
	if o.PrePostScript != nil {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if o.ExcludeFilters != nil {
		toSerialize["excludeFilters"] = o.ExcludeFilters
	}
	return json.Marshal(toSerialize)
}

type NullableCommonMSSQLProtectionGroupParams struct {
	value *CommonMSSQLProtectionGroupParams
	isSet bool
}

func (v NullableCommonMSSQLProtectionGroupParams) Get() *CommonMSSQLProtectionGroupParams {
	return v.value
}

func (v *NullableCommonMSSQLProtectionGroupParams) Set(val *CommonMSSQLProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonMSSQLProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonMSSQLProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonMSSQLProtectionGroupParams(val *CommonMSSQLProtectionGroupParams) *NullableCommonMSSQLProtectionGroupParams {
	return &NullableCommonMSSQLProtectionGroupParams{value: val, isSet: true}
}

func (v NullableCommonMSSQLProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonMSSQLProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonMSSQLProtectionGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}