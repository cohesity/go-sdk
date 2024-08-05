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

// checks if the CommonMSSQLProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonMSSQLProtectionGroupParams{}

// CommonMSSQLProtectionGroupParams Specifies the common params to create a MSSQL Protection Group.
type CommonMSSQLProtectionGroupParams struct {
	// Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases.
	AagBackupPreferenceType NullableString `json:"aagBackupPreferenceType,omitempty"`
	AdvancedSettings *AdvancedSettings `json:"advancedSettings,omitempty"`
	// Specifies whether to backup system databases. If not specified then parameter is set to true.
	BackupSystemDbs NullableBool `json:"backupSystemDbs,omitempty"`
	// Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
	ExcludeFilters []Filter `json:"excludeFilters,omitempty"`
	// Specifies whether full backups should be copy-only.
	FullBackupsCopyOnly NullableBool `json:"fullBackupsCopyOnly,omitempty"`
	// Specifies the number of streams to be used for log backups.
	LogBackupNumStreams NullableInt32 `json:"logBackupNumStreams,omitempty"`
	// Specifies the WithClause to be used for log backups.
	LogBackupWithClause NullableString `json:"logBackupWithClause,omitempty"`
	PrePostScript *PrePostScriptParams `json:"prePostScript,omitempty"`
	// Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used.
	UseAagPreferencesFromServer NullableBool `json:"useAagPreferencesFromServer,omitempty"`
	// Specifies the preference type for backing up user databases on the host.
	UserDbBackupPreferenceType NullableString `json:"userDbBackupPreferenceType,omitempty"`
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

// GetAagBackupPreferenceType returns the AagBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceType() string {
	if o == nil || IsNil(o.AagBackupPreferenceType.Get()) {
		var ret string
		return ret
	}
	return *o.AagBackupPreferenceType.Get()
}

// GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool) {
	if o == nil {
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

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *CommonMSSQLProtectionGroupParams) GetAdvancedSettings() AdvancedSettings {
	if o == nil || IsNil(o.AdvancedSettings) {
		var ret AdvancedSettings
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMSSQLProtectionGroupParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool) {
	if o == nil || IsNil(o.AdvancedSettings) {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasAdvancedSettings() bool {
	if o != nil && !IsNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given AdvancedSettings and assigns it to the AdvancedSettings field.
func (o *CommonMSSQLProtectionGroupParams) SetAdvancedSettings(v AdvancedSettings) {
	o.AdvancedSettings = &v
}

// GetBackupSystemDbs returns the BackupSystemDbs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbs() bool {
	if o == nil || IsNil(o.BackupSystemDbs.Get()) {
		var ret bool
		return ret
	}
	return *o.BackupSystemDbs.Get()
}

// GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool) {
	if o == nil {
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

// GetExcludeFilters returns the ExcludeFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetExcludeFilters() []Filter {
	if o == nil {
		var ret []Filter
		return ret
	}
	return o.ExcludeFilters
}

// GetExcludeFiltersOk returns a tuple with the ExcludeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetExcludeFiltersOk() ([]Filter, bool) {
	if o == nil || IsNil(o.ExcludeFilters) {
		return nil, false
	}
	return o.ExcludeFilters, true
}

// HasExcludeFilters returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasExcludeFilters() bool {
	if o != nil && !IsNil(o.ExcludeFilters) {
		return true
	}

	return false
}

// SetExcludeFilters gets a reference to the given []Filter and assigns it to the ExcludeFilters field.
func (o *CommonMSSQLProtectionGroupParams) SetExcludeFilters(v []Filter) {
	o.ExcludeFilters = v
}

// GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnly() bool {
	if o == nil || IsNil(o.FullBackupsCopyOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.FullBackupsCopyOnly.Get()
}

// GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil {
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

// GetLogBackupNumStreams returns the LogBackupNumStreams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetLogBackupNumStreams() int32 {
	if o == nil || IsNil(o.LogBackupNumStreams.Get()) {
		var ret int32
		return ret
	}
	return *o.LogBackupNumStreams.Get()
}

// GetLogBackupNumStreamsOk returns a tuple with the LogBackupNumStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetLogBackupNumStreamsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogBackupNumStreams.Get(), o.LogBackupNumStreams.IsSet()
}

// HasLogBackupNumStreams returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasLogBackupNumStreams() bool {
	if o != nil && o.LogBackupNumStreams.IsSet() {
		return true
	}

	return false
}

// SetLogBackupNumStreams gets a reference to the given NullableInt32 and assigns it to the LogBackupNumStreams field.
func (o *CommonMSSQLProtectionGroupParams) SetLogBackupNumStreams(v int32) {
	o.LogBackupNumStreams.Set(&v)
}
// SetLogBackupNumStreamsNil sets the value for LogBackupNumStreams to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetLogBackupNumStreamsNil() {
	o.LogBackupNumStreams.Set(nil)
}

// UnsetLogBackupNumStreams ensures that no value is present for LogBackupNumStreams, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetLogBackupNumStreams() {
	o.LogBackupNumStreams.Unset()
}

// GetLogBackupWithClause returns the LogBackupWithClause field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetLogBackupWithClause() string {
	if o == nil || IsNil(o.LogBackupWithClause.Get()) {
		var ret string
		return ret
	}
	return *o.LogBackupWithClause.Get()
}

// GetLogBackupWithClauseOk returns a tuple with the LogBackupWithClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetLogBackupWithClauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogBackupWithClause.Get(), o.LogBackupWithClause.IsSet()
}

// HasLogBackupWithClause returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasLogBackupWithClause() bool {
	if o != nil && o.LogBackupWithClause.IsSet() {
		return true
	}

	return false
}

// SetLogBackupWithClause gets a reference to the given NullableString and assigns it to the LogBackupWithClause field.
func (o *CommonMSSQLProtectionGroupParams) SetLogBackupWithClause(v string) {
	o.LogBackupWithClause.Set(&v)
}
// SetLogBackupWithClauseNil sets the value for LogBackupWithClause to be an explicit nil
func (o *CommonMSSQLProtectionGroupParams) SetLogBackupWithClauseNil() {
	o.LogBackupWithClause.Set(nil)
}

// UnsetLogBackupWithClause ensures that no value is present for LogBackupWithClause, not even an explicit nil
func (o *CommonMSSQLProtectionGroupParams) UnsetLogBackupWithClause() {
	o.LogBackupWithClause.Unset()
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *CommonMSSQLProtectionGroupParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || IsNil(o.PrePostScript) {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonMSSQLProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || IsNil(o.PrePostScript) {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *CommonMSSQLProtectionGroupParams) HasPrePostScript() bool {
	if o != nil && !IsNil(o.PrePostScript) {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *CommonMSSQLProtectionGroupParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServer() bool {
	if o == nil || IsNil(o.UseAagPreferencesFromServer.Get()) {
		var ret bool
		return ret
	}
	return *o.UseAagPreferencesFromServer.Get()
}

// GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool) {
	if o == nil {
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

// GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceType() string {
	if o == nil || IsNil(o.UserDbBackupPreferenceType.Get()) {
		var ret string
		return ret
	}
	return *o.UserDbBackupPreferenceType.Get()
}

// GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool) {
	if o == nil {
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

func (o CommonMSSQLProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonMSSQLProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AagBackupPreferenceType.IsSet() {
		toSerialize["aagBackupPreferenceType"] = o.AagBackupPreferenceType.Get()
	}
	if !IsNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	if o.BackupSystemDbs.IsSet() {
		toSerialize["backupSystemDbs"] = o.BackupSystemDbs.Get()
	}
	if o.ExcludeFilters != nil {
		toSerialize["excludeFilters"] = o.ExcludeFilters
	}
	if o.FullBackupsCopyOnly.IsSet() {
		toSerialize["fullBackupsCopyOnly"] = o.FullBackupsCopyOnly.Get()
	}
	if o.LogBackupNumStreams.IsSet() {
		toSerialize["logBackupNumStreams"] = o.LogBackupNumStreams.Get()
	}
	if o.LogBackupWithClause.IsSet() {
		toSerialize["logBackupWithClause"] = o.LogBackupWithClause.Get()
	}
	if !IsNil(o.PrePostScript) {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if o.UseAagPreferencesFromServer.IsSet() {
		toSerialize["useAagPreferencesFromServer"] = o.UseAagPreferencesFromServer.Get()
	}
	if o.UserDbBackupPreferenceType.IsSet() {
		toSerialize["userDbBackupPreferenceType"] = o.UserDbBackupPreferenceType.Get()
	}
	return toSerialize, nil
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


