# MSSQLFileProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AagBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases. | [optional] 
**AdvancedSettings** | Pointer to [**AdvancedSettings**](AdvancedSettings.md) |  | [optional] 
**BackupSystemDbs** | Pointer to **NullableBool** | Specifies whether to backup system databases. If not specified then parameter is set to true. | [optional] 
**ExcludeFilters** | Pointer to [**[]Filter**](Filter.md) | Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters. | [optional] 
**FullBackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether full backups should be copy-only. | [optional] 
**LogBackupNumStreams** | Pointer to **NullableInt32** | Specifies the number of streams to be used for log backups. | [optional] 
**LogBackupWithClause** | Pointer to **NullableString** | Specifies the WithClause to be used for log backups. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**UseAagPreferencesFromServer** | Pointer to **NullableBool** | Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used. | [optional] 
**UserDbBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up user databases on the host. | [optional] 
**AdditionalHostParams** | Pointer to [**[]MSSQLFileProtectionGroupHostParams**](MSSQLFileProtectionGroupHostParams.md) | Specifies settings which are to be applied to specific host containers in this protection group. | [optional] 
**Objects** | [**[]MSSQLFileProtectionGroupObjectParams**](MSSQLFileProtectionGroupObjectParams.md) | Specifies the list of object params to be protected. | 
**PerformSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to perform source side deduplication on this Protection Group. | [optional] 

## Methods

### NewMSSQLFileProtectionGroupParams

`func NewMSSQLFileProtectionGroupParams(objects []MSSQLFileProtectionGroupObjectParams, ) *MSSQLFileProtectionGroupParams`

NewMSSQLFileProtectionGroupParams instantiates a new MSSQLFileProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLFileProtectionGroupParamsWithDefaults

`func NewMSSQLFileProtectionGroupParamsWithDefaults() *MSSQLFileProtectionGroupParams`

NewMSSQLFileProtectionGroupParamsWithDefaults instantiates a new MSSQLFileProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) GetAagBackupPreferenceType() string`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *MSSQLFileProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) SetAagBackupPreferenceType(v string)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *MSSQLFileProtectionGroupParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *MSSQLFileProtectionGroupParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetAdvancedSettings

`func (o *MSSQLFileProtectionGroupParams) GetAdvancedSettings() AdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *MSSQLFileProtectionGroupParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *MSSQLFileProtectionGroupParams) SetAdvancedSettings(v AdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *MSSQLFileProtectionGroupParams) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetBackupSystemDbs

`func (o *MSSQLFileProtectionGroupParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *MSSQLFileProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *MSSQLFileProtectionGroupParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *MSSQLFileProtectionGroupParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *MSSQLFileProtectionGroupParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *MSSQLFileProtectionGroupParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetExcludeFilters

`func (o *MSSQLFileProtectionGroupParams) GetExcludeFilters() []Filter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *MSSQLFileProtectionGroupParams) GetExcludeFiltersOk() (*[]Filter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *MSSQLFileProtectionGroupParams) SetExcludeFilters(v []Filter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *MSSQLFileProtectionGroupParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *MSSQLFileProtectionGroupParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *MSSQLFileProtectionGroupParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetFullBackupsCopyOnly

`func (o *MSSQLFileProtectionGroupParams) GetFullBackupsCopyOnly() bool`

GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field if non-nil, zero value otherwise.

### GetFullBackupsCopyOnlyOk

`func (o *MSSQLFileProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool)`

GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupsCopyOnly

`func (o *MSSQLFileProtectionGroupParams) SetFullBackupsCopyOnly(v bool)`

SetFullBackupsCopyOnly sets FullBackupsCopyOnly field to given value.

### HasFullBackupsCopyOnly

`func (o *MSSQLFileProtectionGroupParams) HasFullBackupsCopyOnly() bool`

HasFullBackupsCopyOnly returns a boolean if a field has been set.

### SetFullBackupsCopyOnlyNil

`func (o *MSSQLFileProtectionGroupParams) SetFullBackupsCopyOnlyNil(b bool)`

 SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil

### UnsetFullBackupsCopyOnly
`func (o *MSSQLFileProtectionGroupParams) UnsetFullBackupsCopyOnly()`

UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
### GetLogBackupNumStreams

`func (o *MSSQLFileProtectionGroupParams) GetLogBackupNumStreams() int32`

GetLogBackupNumStreams returns the LogBackupNumStreams field if non-nil, zero value otherwise.

### GetLogBackupNumStreamsOk

`func (o *MSSQLFileProtectionGroupParams) GetLogBackupNumStreamsOk() (*int32, bool)`

GetLogBackupNumStreamsOk returns a tuple with the LogBackupNumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupNumStreams

`func (o *MSSQLFileProtectionGroupParams) SetLogBackupNumStreams(v int32)`

SetLogBackupNumStreams sets LogBackupNumStreams field to given value.

### HasLogBackupNumStreams

`func (o *MSSQLFileProtectionGroupParams) HasLogBackupNumStreams() bool`

HasLogBackupNumStreams returns a boolean if a field has been set.

### SetLogBackupNumStreamsNil

`func (o *MSSQLFileProtectionGroupParams) SetLogBackupNumStreamsNil(b bool)`

 SetLogBackupNumStreamsNil sets the value for LogBackupNumStreams to be an explicit nil

### UnsetLogBackupNumStreams
`func (o *MSSQLFileProtectionGroupParams) UnsetLogBackupNumStreams()`

UnsetLogBackupNumStreams ensures that no value is present for LogBackupNumStreams, not even an explicit nil
### GetLogBackupWithClause

`func (o *MSSQLFileProtectionGroupParams) GetLogBackupWithClause() string`

GetLogBackupWithClause returns the LogBackupWithClause field if non-nil, zero value otherwise.

### GetLogBackupWithClauseOk

`func (o *MSSQLFileProtectionGroupParams) GetLogBackupWithClauseOk() (*string, bool)`

GetLogBackupWithClauseOk returns a tuple with the LogBackupWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupWithClause

`func (o *MSSQLFileProtectionGroupParams) SetLogBackupWithClause(v string)`

SetLogBackupWithClause sets LogBackupWithClause field to given value.

### HasLogBackupWithClause

`func (o *MSSQLFileProtectionGroupParams) HasLogBackupWithClause() bool`

HasLogBackupWithClause returns a boolean if a field has been set.

### SetLogBackupWithClauseNil

`func (o *MSSQLFileProtectionGroupParams) SetLogBackupWithClauseNil(b bool)`

 SetLogBackupWithClauseNil sets the value for LogBackupWithClause to be an explicit nil

### UnsetLogBackupWithClause
`func (o *MSSQLFileProtectionGroupParams) UnsetLogBackupWithClause()`

UnsetLogBackupWithClause ensures that no value is present for LogBackupWithClause, not even an explicit nil
### GetPrePostScript

`func (o *MSSQLFileProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *MSSQLFileProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *MSSQLFileProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *MSSQLFileProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetUseAagPreferencesFromServer

`func (o *MSSQLFileProtectionGroupParams) GetUseAagPreferencesFromServer() bool`

GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromServerOk

`func (o *MSSQLFileProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool)`

GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromServer

`func (o *MSSQLFileProtectionGroupParams) SetUseAagPreferencesFromServer(v bool)`

SetUseAagPreferencesFromServer sets UseAagPreferencesFromServer field to given value.

### HasUseAagPreferencesFromServer

`func (o *MSSQLFileProtectionGroupParams) HasUseAagPreferencesFromServer() bool`

HasUseAagPreferencesFromServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromServerNil

`func (o *MSSQLFileProtectionGroupParams) SetUseAagPreferencesFromServerNil(b bool)`

 SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil

### UnsetUseAagPreferencesFromServer
`func (o *MSSQLFileProtectionGroupParams) UnsetUseAagPreferencesFromServer()`

UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
### GetUserDbBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) GetUserDbBackupPreferenceType() string`

GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field if non-nil, zero value otherwise.

### GetUserDbBackupPreferenceTypeOk

`func (o *MSSQLFileProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool)`

GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) SetUserDbBackupPreferenceType(v string)`

SetUserDbBackupPreferenceType sets UserDbBackupPreferenceType field to given value.

### HasUserDbBackupPreferenceType

`func (o *MSSQLFileProtectionGroupParams) HasUserDbBackupPreferenceType() bool`

HasUserDbBackupPreferenceType returns a boolean if a field has been set.

### SetUserDbBackupPreferenceTypeNil

`func (o *MSSQLFileProtectionGroupParams) SetUserDbBackupPreferenceTypeNil(b bool)`

 SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil

### UnsetUserDbBackupPreferenceType
`func (o *MSSQLFileProtectionGroupParams) UnsetUserDbBackupPreferenceType()`

UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
### GetAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParams) GetAdditionalHostParams() []MSSQLFileProtectionGroupHostParams`

GetAdditionalHostParams returns the AdditionalHostParams field if non-nil, zero value otherwise.

### GetAdditionalHostParamsOk

`func (o *MSSQLFileProtectionGroupParams) GetAdditionalHostParamsOk() (*[]MSSQLFileProtectionGroupHostParams, bool)`

GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParams) SetAdditionalHostParams(v []MSSQLFileProtectionGroupHostParams)`

SetAdditionalHostParams sets AdditionalHostParams field to given value.

### HasAdditionalHostParams

`func (o *MSSQLFileProtectionGroupParams) HasAdditionalHostParams() bool`

HasAdditionalHostParams returns a boolean if a field has been set.

### GetObjects

`func (o *MSSQLFileProtectionGroupParams) GetObjects() []MSSQLFileProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MSSQLFileProtectionGroupParams) GetObjectsOk() (*[]MSSQLFileProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MSSQLFileProtectionGroupParams) SetObjects(v []MSSQLFileProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MSSQLFileProtectionGroupParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MSSQLFileProtectionGroupParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParams) GetPerformSourceSideDeduplication() bool`

GetPerformSourceSideDeduplication returns the PerformSourceSideDeduplication field if non-nil, zero value otherwise.

### GetPerformSourceSideDeduplicationOk

`func (o *MSSQLFileProtectionGroupParams) GetPerformSourceSideDeduplicationOk() (*bool, bool)`

GetPerformSourceSideDeduplicationOk returns a tuple with the PerformSourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParams) SetPerformSourceSideDeduplication(v bool)`

SetPerformSourceSideDeduplication sets PerformSourceSideDeduplication field to given value.

### HasPerformSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupParams) HasPerformSourceSideDeduplication() bool`

HasPerformSourceSideDeduplication returns a boolean if a field has been set.

### SetPerformSourceSideDeduplicationNil

`func (o *MSSQLFileProtectionGroupParams) SetPerformSourceSideDeduplicationNil(b bool)`

 SetPerformSourceSideDeduplicationNil sets the value for PerformSourceSideDeduplication to be an explicit nil

### UnsetPerformSourceSideDeduplication
`func (o *MSSQLFileProtectionGroupParams) UnsetPerformSourceSideDeduplication()`

UnsetPerformSourceSideDeduplication ensures that no value is present for PerformSourceSideDeduplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


