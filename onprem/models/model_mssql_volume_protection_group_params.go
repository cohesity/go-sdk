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

// MSSQLVolumeProtectionGroupParams Specifies the params to create a Volume based MSSQL Protection Group.
type MSSQLVolumeProtectionGroupParams struct {
	// Specifies the list of object ids to be protected.
	Objects []MSSQLVolumeProtectionGroupObjectParams `json:"objects"`
	// Specifies whether or to perform incremental backups the first time after a server restarts. By default, a full backup will be performed.
	IncrementalBackupAfterRestart NullableBool `json:"incrementalBackupAfterRestart,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies whether to only backup volumes on which the specified databases reside. If not specified (default), all the volumes of the host will be protected.
	BackupDbVolumesOnly NullableBool `json:"backupDbVolumesOnly,omitempty"`
	// Specifies settings which are to be applied to specific host containers in this protection group.
	AdditionalHostParams *[]MSSQLVolumeProtectionGroupHostParams `json:"additionalHostParams,omitempty"`
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

// NewMSSQLVolumeProtectionGroupParams instantiates a new MSSQLVolumeProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSSQLVolumeProtectionGroupParams(objects []MSSQLVolumeProtectionGroupObjectParams) *MSSQLVolumeProtectionGroupParams {
	this := MSSQLVolumeProtectionGroupParams{}
	this.Objects = objects
	return &this
}

// NewMSSQLVolumeProtectionGroupParamsWithDefaults instantiates a new MSSQLVolumeProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSSQLVolumeProtectionGroupParamsWithDefaults() *MSSQLVolumeProtectionGroupParams {
	this := MSSQLVolumeProtectionGroupParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []MSSQLVolumeProtectionGroupObjectParams will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetObjects() []MSSQLVolumeProtectionGroupObjectParams {
	if o == nil {
		var ret []MSSQLVolumeProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetObjectsOk() (*[]MSSQLVolumeProtectionGroupObjectParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *MSSQLVolumeProtectionGroupParams) SetObjects(v []MSSQLVolumeProtectionGroupObjectParams) {
	o.Objects = v
}

// GetIncrementalBackupAfterRestart returns the IncrementalBackupAfterRestart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetIncrementalBackupAfterRestart() bool {
	if o == nil || o.IncrementalBackupAfterRestart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncrementalBackupAfterRestart.Get()
}

// GetIncrementalBackupAfterRestartOk returns a tuple with the IncrementalBackupAfterRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetIncrementalBackupAfterRestartOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncrementalBackupAfterRestart.Get(), o.IncrementalBackupAfterRestart.IsSet()
}

// HasIncrementalBackupAfterRestart returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasIncrementalBackupAfterRestart() bool {
	if o != nil && o.IncrementalBackupAfterRestart.IsSet() {
		return true
	}

	return false
}

// SetIncrementalBackupAfterRestart gets a reference to the given NullableBool and assigns it to the IncrementalBackupAfterRestart field.
func (o *MSSQLVolumeProtectionGroupParams) SetIncrementalBackupAfterRestart(v bool) {
	o.IncrementalBackupAfterRestart.Set(&v)
}
// SetIncrementalBackupAfterRestartNil sets the value for IncrementalBackupAfterRestart to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetIncrementalBackupAfterRestartNil() {
	o.IncrementalBackupAfterRestart.Set(nil)
}

// UnsetIncrementalBackupAfterRestart ensures that no value is present for IncrementalBackupAfterRestart, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetIncrementalBackupAfterRestart() {
	o.IncrementalBackupAfterRestart.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *MSSQLVolumeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLVolumeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *MSSQLVolumeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetBackupDbVolumesOnly returns the BackupDbVolumesOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetBackupDbVolumesOnly() bool {
	if o == nil || o.BackupDbVolumesOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupDbVolumesOnly.Get()
}

// GetBackupDbVolumesOnlyOk returns a tuple with the BackupDbVolumesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetBackupDbVolumesOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupDbVolumesOnly.Get(), o.BackupDbVolumesOnly.IsSet()
}

// HasBackupDbVolumesOnly returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasBackupDbVolumesOnly() bool {
	if o != nil && o.BackupDbVolumesOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupDbVolumesOnly gets a reference to the given NullableBool and assigns it to the BackupDbVolumesOnly field.
func (o *MSSQLVolumeProtectionGroupParams) SetBackupDbVolumesOnly(v bool) {
	o.BackupDbVolumesOnly.Set(&v)
}
// SetBackupDbVolumesOnlyNil sets the value for BackupDbVolumesOnly to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetBackupDbVolumesOnlyNil() {
	o.BackupDbVolumesOnly.Set(nil)
}

// UnsetBackupDbVolumesOnly ensures that no value is present for BackupDbVolumesOnly, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetBackupDbVolumesOnly() {
	o.BackupDbVolumesOnly.Unset()
}

// GetAdditionalHostParams returns the AdditionalHostParams field value if set, zero value otherwise.
func (o *MSSQLVolumeProtectionGroupParams) GetAdditionalHostParams() []MSSQLVolumeProtectionGroupHostParams {
	if o == nil || o.AdditionalHostParams == nil {
		var ret []MSSQLVolumeProtectionGroupHostParams
		return ret
	}
	return *o.AdditionalHostParams
}

// GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLVolumeProtectionGroupParams) GetAdditionalHostParamsOk() (*[]MSSQLVolumeProtectionGroupHostParams, bool) {
	if o == nil || o.AdditionalHostParams == nil {
		return nil, false
	}
	return o.AdditionalHostParams, true
}

// HasAdditionalHostParams returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasAdditionalHostParams() bool {
	if o != nil && o.AdditionalHostParams != nil {
		return true
	}

	return false
}

// SetAdditionalHostParams gets a reference to the given []MSSQLVolumeProtectionGroupHostParams and assigns it to the AdditionalHostParams field.
func (o *MSSQLVolumeProtectionGroupParams) SetAdditionalHostParams(v []MSSQLVolumeProtectionGroupHostParams) {
	o.AdditionalHostParams = &v
}

// GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetUserDbBackupPreferenceType() string {
	if o == nil || o.UserDbBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDbBackupPreferenceType.Get()
}

// GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDbBackupPreferenceType.Get(), o.UserDbBackupPreferenceType.IsSet()
}

// HasUserDbBackupPreferenceType returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasUserDbBackupPreferenceType() bool {
	if o != nil && o.UserDbBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetUserDbBackupPreferenceType gets a reference to the given NullableString and assigns it to the UserDbBackupPreferenceType field.
func (o *MSSQLVolumeProtectionGroupParams) SetUserDbBackupPreferenceType(v string) {
	o.UserDbBackupPreferenceType.Set(&v)
}
// SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetUserDbBackupPreferenceTypeNil() {
	o.UserDbBackupPreferenceType.Set(nil)
}

// UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetUserDbBackupPreferenceType() {
	o.UserDbBackupPreferenceType.Unset()
}

// GetBackupSystemDbs returns the BackupSystemDbs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetBackupSystemDbs() bool {
	if o == nil || o.BackupSystemDbs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupSystemDbs.Get()
}

// GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupSystemDbs.Get(), o.BackupSystemDbs.IsSet()
}

// HasBackupSystemDbs returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasBackupSystemDbs() bool {
	if o != nil && o.BackupSystemDbs.IsSet() {
		return true
	}

	return false
}

// SetBackupSystemDbs gets a reference to the given NullableBool and assigns it to the BackupSystemDbs field.
func (o *MSSQLVolumeProtectionGroupParams) SetBackupSystemDbs(v bool) {
	o.BackupSystemDbs.Set(&v)
}
// SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetBackupSystemDbsNil() {
	o.BackupSystemDbs.Set(nil)
}

// UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetBackupSystemDbs() {
	o.BackupSystemDbs.Unset()
}

// GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetUseAagPreferencesFromServer() bool {
	if o == nil || o.UseAagPreferencesFromServer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseAagPreferencesFromServer.Get()
}

// GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseAagPreferencesFromServer.Get(), o.UseAagPreferencesFromServer.IsSet()
}

// HasUseAagPreferencesFromServer returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasUseAagPreferencesFromServer() bool {
	if o != nil && o.UseAagPreferencesFromServer.IsSet() {
		return true
	}

	return false
}

// SetUseAagPreferencesFromServer gets a reference to the given NullableBool and assigns it to the UseAagPreferencesFromServer field.
func (o *MSSQLVolumeProtectionGroupParams) SetUseAagPreferencesFromServer(v bool) {
	o.UseAagPreferencesFromServer.Set(&v)
}
// SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetUseAagPreferencesFromServerNil() {
	o.UseAagPreferencesFromServer.Set(nil)
}

// UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetUseAagPreferencesFromServer() {
	o.UseAagPreferencesFromServer.Unset()
}

// GetAagBackupPreferenceType returns the AagBackupPreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetAagBackupPreferenceType() string {
	if o == nil || o.AagBackupPreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AagBackupPreferenceType.Get()
}

// GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AagBackupPreferenceType.Get(), o.AagBackupPreferenceType.IsSet()
}

// HasAagBackupPreferenceType returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasAagBackupPreferenceType() bool {
	if o != nil && o.AagBackupPreferenceType.IsSet() {
		return true
	}

	return false
}

// SetAagBackupPreferenceType gets a reference to the given NullableString and assigns it to the AagBackupPreferenceType field.
func (o *MSSQLVolumeProtectionGroupParams) SetAagBackupPreferenceType(v string) {
	o.AagBackupPreferenceType.Set(&v)
}
// SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetAagBackupPreferenceTypeNil() {
	o.AagBackupPreferenceType.Set(nil)
}

// UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetAagBackupPreferenceType() {
	o.AagBackupPreferenceType.Unset()
}

// GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetFullBackupsCopyOnly() bool {
	if o == nil || o.FullBackupsCopyOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FullBackupsCopyOnly.Get()
}

// GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullBackupsCopyOnly.Get(), o.FullBackupsCopyOnly.IsSet()
}

// HasFullBackupsCopyOnly returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasFullBackupsCopyOnly() bool {
	if o != nil && o.FullBackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetFullBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the FullBackupsCopyOnly field.
func (o *MSSQLVolumeProtectionGroupParams) SetFullBackupsCopyOnly(v bool) {
	o.FullBackupsCopyOnly.Set(&v)
}
// SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) SetFullBackupsCopyOnlyNil() {
	o.FullBackupsCopyOnly.Set(nil)
}

// UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParams) UnsetFullBackupsCopyOnly() {
	o.FullBackupsCopyOnly.Unset()
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *MSSQLVolumeProtectionGroupParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLVolumeProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *MSSQLVolumeProtectionGroupParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetExcludeFilters returns the ExcludeFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParams) GetExcludeFilters() []Filter {
	if o == nil  {
		var ret []Filter
		return ret
	}
	return o.ExcludeFilters
}

// GetExcludeFiltersOk returns a tuple with the ExcludeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParams) GetExcludeFiltersOk() (*[]Filter, bool) {
	if o == nil || o.ExcludeFilters == nil {
		return nil, false
	}
	return &o.ExcludeFilters, true
}

// HasExcludeFilters returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParams) HasExcludeFilters() bool {
	if o != nil && o.ExcludeFilters != nil {
		return true
	}

	return false
}

// SetExcludeFilters gets a reference to the given []Filter and assigns it to the ExcludeFilters field.
func (o *MSSQLVolumeProtectionGroupParams) SetExcludeFilters(v []Filter) {
	o.ExcludeFilters = v
}

func (o MSSQLVolumeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.IncrementalBackupAfterRestart.IsSet() {
		toSerialize["incrementalBackupAfterRestart"] = o.IncrementalBackupAfterRestart.Get()
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.BackupDbVolumesOnly.IsSet() {
		toSerialize["backupDbVolumesOnly"] = o.BackupDbVolumesOnly.Get()
	}
	if o.AdditionalHostParams != nil {
		toSerialize["additionalHostParams"] = o.AdditionalHostParams
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

type NullableMSSQLVolumeProtectionGroupParams struct {
	value *MSSQLVolumeProtectionGroupParams
	isSet bool
}

func (v NullableMSSQLVolumeProtectionGroupParams) Get() *MSSQLVolumeProtectionGroupParams {
	return v.value
}

func (v *NullableMSSQLVolumeProtectionGroupParams) Set(val *MSSQLVolumeProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMSSQLVolumeProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMSSQLVolumeProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSSQLVolumeProtectionGroupParams(val *MSSQLVolumeProtectionGroupParams) *NullableMSSQLVolumeProtectionGroupParams {
	return &NullableMSSQLVolumeProtectionGroupParams{value: val, isSet: true}
}

func (v NullableMSSQLVolumeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSSQLVolumeProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o MSSQLVolumeProtectionGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}