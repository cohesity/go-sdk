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

// MssqlNativeObjectProtectionParams Specifies the params to create a Native based MSSQL Object Protection.
type MssqlNativeObjectProtectionParams struct {
	// Specifies the list of objects to be protected.
	Objects []MssqlNativeObjectProtection `json:"objects"`
	// Specifies the number of streams to be used.
	NumStreams NullableInt32 `json:"numStreams,omitempty"`
	// Specifies the WithClause to be used.
	WithClause NullableString `json:"withClause,omitempty"`
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

// NewMssqlNativeObjectProtectionParams instantiates a new MssqlNativeObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMssqlNativeObjectProtectionParams(objects []MssqlNativeObjectProtection) *MssqlNativeObjectProtectionParams {
	this := MssqlNativeObjectProtectionParams{}
	this.Objects = objects
	return &this
}

// NewMssqlNativeObjectProtectionParamsWithDefaults instantiates a new MssqlNativeObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMssqlNativeObjectProtectionParamsWithDefaults() *MssqlNativeObjectProtectionParams {
	this := MssqlNativeObjectProtectionParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []MssqlNativeObjectProtection will be returned
func (o *MssqlNativeObjectProtectionParams) GetObjects() []MssqlNativeObjectProtection {
	if o == nil {
		var ret []MssqlNativeObjectProtection
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetObjectsOk() (*[]MssqlNativeObjectProtection, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *MssqlNativeObjectProtectionParams) SetObjects(v []MssqlNativeObjectProtection) {
	o.Objects = v
}

// GetNumStreams returns the NumStreams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetNumStreams() int32 {
	if o == nil || o.NumStreams.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumStreams.Get()
}

// GetNumStreamsOk returns a tuple with the NumStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetNumStreamsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumStreams.Get(), o.NumStreams.IsSet()
}

// HasNumStreams returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasNumStreams() bool {
	if o != nil && o.NumStreams.IsSet() {
		return true
	}

	return false
}

// SetNumStreams gets a reference to the given NullableInt32 and assigns it to the NumStreams field.
func (o *MssqlNativeObjectProtectionParams) SetNumStreams(v int32) {
	o.NumStreams.Set(&v)
}
// SetNumStreamsNil sets the value for NumStreams to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetNumStreamsNil() {
	o.NumStreams.Set(nil)
}

// UnsetNumStreams ensures that no value is present for NumStreams, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetNumStreams() {
	o.NumStreams.Unset()
}

// GetWithClause returns the WithClause field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetWithClause() string {
	if o == nil || o.WithClause.Get() == nil {
		var ret string
		return ret
	}
	return *o.WithClause.Get()
}

// GetWithClauseOk returns a tuple with the WithClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetWithClauseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WithClause.Get(), o.WithClause.IsSet()
}

// HasWithClause returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasWithClause() bool {
	if o != nil && o.WithClause.IsSet() {
		return true
	}

	return false
}

// SetWithClause gets a reference to the given NullableString and assigns it to the WithClause field.
func (o *MssqlNativeObjectProtectionParams) SetWithClause(v string) {
	o.WithClause.Set(&v)
}
// SetWithClauseNil sets the value for WithClause to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetWithClauseNil() {
	o.WithClause.Set(nil)
}

// UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetWithClause() {
	o.WithClause.Unset()
}

// GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetUserDbBackupPreferenceType() string {
	if o == nil || o.UserDbBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDbBackupPreferenceType.Get()
}

// GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetUserDbBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDbBackupPreferenceType.Get(), o.UserDbBackupPreferenceType.IsSet()
}

// HasUserDbBackupPreferenceType returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasUserDbBackupPreferenceType() bool {
	if o != nil && o.UserDbBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetUserDbBackupPreferenceType gets a reference to the given NullableString and assigns it to the UserDbBackupPreferenceType field.
func (o *MssqlNativeObjectProtectionParams) SetUserDbBackupPreferenceType(v string) {
	o.UserDbBackupPreferenceType.Set(&v)
}
// SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetUserDbBackupPreferenceTypeNil() {
	o.UserDbBackupPreferenceType.Set(nil)
}

// UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetUserDbBackupPreferenceType() {
	o.UserDbBackupPreferenceType.Unset()
}

// GetBackupSystemDbs returns the BackupSystemDbs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetBackupSystemDbs() bool {
	if o == nil || o.BackupSystemDbs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupSystemDbs.Get()
}

// GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetBackupSystemDbsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupSystemDbs.Get(), o.BackupSystemDbs.IsSet()
}

// HasBackupSystemDbs returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasBackupSystemDbs() bool {
	if o != nil && o.BackupSystemDbs.IsSet() {
		return true
	}

	return false
}

// SetBackupSystemDbs gets a reference to the given NullableBool and assigns it to the BackupSystemDbs field.
func (o *MssqlNativeObjectProtectionParams) SetBackupSystemDbs(v bool) {
	o.BackupSystemDbs.Set(&v)
}
// SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetBackupSystemDbsNil() {
	o.BackupSystemDbs.Set(nil)
}

// UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetBackupSystemDbs() {
	o.BackupSystemDbs.Unset()
}

// GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetUseAagPreferencesFromServer() bool {
	if o == nil || o.UseAagPreferencesFromServer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseAagPreferencesFromServer.Get()
}

// GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetUseAagPreferencesFromServerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseAagPreferencesFromServer.Get(), o.UseAagPreferencesFromServer.IsSet()
}

// HasUseAagPreferencesFromServer returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasUseAagPreferencesFromServer() bool {
	if o != nil && o.UseAagPreferencesFromServer.IsSet() {
		return true
	}

	return false
}

// SetUseAagPreferencesFromServer gets a reference to the given NullableBool and assigns it to the UseAagPreferencesFromServer field.
func (o *MssqlNativeObjectProtectionParams) SetUseAagPreferencesFromServer(v bool) {
	o.UseAagPreferencesFromServer.Set(&v)
}
// SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetUseAagPreferencesFromServerNil() {
	o.UseAagPreferencesFromServer.Set(nil)
}

// UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetUseAagPreferencesFromServer() {
	o.UseAagPreferencesFromServer.Unset()
}

// GetAagBackupPreferenceType returns the AagBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetAagBackupPreferenceType() string {
	if o == nil || o.AagBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AagBackupPreferenceType.Get()
}

// GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetAagBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AagBackupPreferenceType.Get(), o.AagBackupPreferenceType.IsSet()
}

// HasAagBackupPreferenceType returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasAagBackupPreferenceType() bool {
	if o != nil && o.AagBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetAagBackupPreferenceType gets a reference to the given NullableString and assigns it to the AagBackupPreferenceType field.
func (o *MssqlNativeObjectProtectionParams) SetAagBackupPreferenceType(v string) {
	o.AagBackupPreferenceType.Set(&v)
}
// SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetAagBackupPreferenceTypeNil() {
	o.AagBackupPreferenceType.Set(nil)
}

// UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetAagBackupPreferenceType() {
	o.AagBackupPreferenceType.Unset()
}

// GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetFullBackupsCopyOnly() bool {
	if o == nil || o.FullBackupsCopyOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FullBackupsCopyOnly.Get()
}

// GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetFullBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullBackupsCopyOnly.Get(), o.FullBackupsCopyOnly.IsSet()
}

// HasFullBackupsCopyOnly returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasFullBackupsCopyOnly() bool {
	if o != nil && o.FullBackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetFullBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the FullBackupsCopyOnly field.
func (o *MssqlNativeObjectProtectionParams) SetFullBackupsCopyOnly(v bool) {
	o.FullBackupsCopyOnly.Set(&v)
}
// SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil
func (o *MssqlNativeObjectProtectionParams) SetFullBackupsCopyOnlyNil() {
	o.FullBackupsCopyOnly.Set(nil)
}

// UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
func (o *MssqlNativeObjectProtectionParams) UnsetFullBackupsCopyOnly() {
	o.FullBackupsCopyOnly.Unset()
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *MssqlNativeObjectProtectionParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlNativeObjectProtectionParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *MssqlNativeObjectProtectionParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetExcludeFilters returns the ExcludeFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MssqlNativeObjectProtectionParams) GetExcludeFilters() []Filter {
	if o == nil  {
		var ret []Filter
		return ret
	}
	return o.ExcludeFilters
}

// GetExcludeFiltersOk returns a tuple with the ExcludeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MssqlNativeObjectProtectionParams) GetExcludeFiltersOk() (*[]Filter, bool) {
	if o == nil || o.ExcludeFilters == nil {
		return nil, false
	}
	return &o.ExcludeFilters, true
}

// HasExcludeFilters returns a boolean if a field has been set.
func (o *MssqlNativeObjectProtectionParams) HasExcludeFilters() bool {
	if o != nil && o.ExcludeFilters != nil {
		return true
	}

	return false
}

// SetExcludeFilters gets a reference to the given []Filter and assigns it to the ExcludeFilters field.
func (o *MssqlNativeObjectProtectionParams) SetExcludeFilters(v []Filter) {
	o.ExcludeFilters = v
}

func (o MssqlNativeObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.NumStreams.IsSet() {
		toSerialize["numStreams"] = o.NumStreams.Get()
	}
	if o.WithClause.IsSet() {
		toSerialize["withClause"] = o.WithClause.Get()
	}
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

type NullableMssqlNativeObjectProtectionParams struct {
	value *MssqlNativeObjectProtectionParams
	isSet bool
}

func (v NullableMssqlNativeObjectProtectionParams) Get() *MssqlNativeObjectProtectionParams {
	return v.value
}

func (v *NullableMssqlNativeObjectProtectionParams) Set(val *MssqlNativeObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMssqlNativeObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMssqlNativeObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMssqlNativeObjectProtectionParams(val *MssqlNativeObjectProtectionParams) *NullableMssqlNativeObjectProtectionParams {
	return &NullableMssqlNativeObjectProtectionParams{value: val, isSet: true}
}

func (v NullableMssqlNativeObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMssqlNativeObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


