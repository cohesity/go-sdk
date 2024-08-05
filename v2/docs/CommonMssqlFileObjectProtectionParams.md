# CommonMssqlFileObjectProtectionParams

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

## Methods

### NewCommonMssqlFileObjectProtectionParams

`func NewCommonMssqlFileObjectProtectionParams() *CommonMssqlFileObjectProtectionParams`

NewCommonMssqlFileObjectProtectionParams instantiates a new CommonMssqlFileObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMssqlFileObjectProtectionParamsWithDefaults

`func NewCommonMssqlFileObjectProtectionParamsWithDefaults() *CommonMssqlFileObjectProtectionParams`

NewCommonMssqlFileObjectProtectionParamsWithDefaults instantiates a new CommonMssqlFileObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) GetAagBackupPreferenceType() string`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *CommonMssqlFileObjectProtectionParams) GetAagBackupPreferenceTypeOk() (*string, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) SetAagBackupPreferenceType(v string)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *CommonMssqlFileObjectProtectionParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *CommonMssqlFileObjectProtectionParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetAdvancedSettings

`func (o *CommonMssqlFileObjectProtectionParams) GetAdvancedSettings() AdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *CommonMssqlFileObjectProtectionParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *CommonMssqlFileObjectProtectionParams) SetAdvancedSettings(v AdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *CommonMssqlFileObjectProtectionParams) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetBackupSystemDbs

`func (o *CommonMssqlFileObjectProtectionParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *CommonMssqlFileObjectProtectionParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *CommonMssqlFileObjectProtectionParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *CommonMssqlFileObjectProtectionParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *CommonMssqlFileObjectProtectionParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *CommonMssqlFileObjectProtectionParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetExcludeFilters

`func (o *CommonMssqlFileObjectProtectionParams) GetExcludeFilters() []Filter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *CommonMssqlFileObjectProtectionParams) GetExcludeFiltersOk() (*[]Filter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *CommonMssqlFileObjectProtectionParams) SetExcludeFilters(v []Filter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *CommonMssqlFileObjectProtectionParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *CommonMssqlFileObjectProtectionParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *CommonMssqlFileObjectProtectionParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetFullBackupsCopyOnly

`func (o *CommonMssqlFileObjectProtectionParams) GetFullBackupsCopyOnly() bool`

GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field if non-nil, zero value otherwise.

### GetFullBackupsCopyOnlyOk

`func (o *CommonMssqlFileObjectProtectionParams) GetFullBackupsCopyOnlyOk() (*bool, bool)`

GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupsCopyOnly

`func (o *CommonMssqlFileObjectProtectionParams) SetFullBackupsCopyOnly(v bool)`

SetFullBackupsCopyOnly sets FullBackupsCopyOnly field to given value.

### HasFullBackupsCopyOnly

`func (o *CommonMssqlFileObjectProtectionParams) HasFullBackupsCopyOnly() bool`

HasFullBackupsCopyOnly returns a boolean if a field has been set.

### SetFullBackupsCopyOnlyNil

`func (o *CommonMssqlFileObjectProtectionParams) SetFullBackupsCopyOnlyNil(b bool)`

 SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil

### UnsetFullBackupsCopyOnly
`func (o *CommonMssqlFileObjectProtectionParams) UnsetFullBackupsCopyOnly()`

UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
### GetLogBackupNumStreams

`func (o *CommonMssqlFileObjectProtectionParams) GetLogBackupNumStreams() int32`

GetLogBackupNumStreams returns the LogBackupNumStreams field if non-nil, zero value otherwise.

### GetLogBackupNumStreamsOk

`func (o *CommonMssqlFileObjectProtectionParams) GetLogBackupNumStreamsOk() (*int32, bool)`

GetLogBackupNumStreamsOk returns a tuple with the LogBackupNumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupNumStreams

`func (o *CommonMssqlFileObjectProtectionParams) SetLogBackupNumStreams(v int32)`

SetLogBackupNumStreams sets LogBackupNumStreams field to given value.

### HasLogBackupNumStreams

`func (o *CommonMssqlFileObjectProtectionParams) HasLogBackupNumStreams() bool`

HasLogBackupNumStreams returns a boolean if a field has been set.

### SetLogBackupNumStreamsNil

`func (o *CommonMssqlFileObjectProtectionParams) SetLogBackupNumStreamsNil(b bool)`

 SetLogBackupNumStreamsNil sets the value for LogBackupNumStreams to be an explicit nil

### UnsetLogBackupNumStreams
`func (o *CommonMssqlFileObjectProtectionParams) UnsetLogBackupNumStreams()`

UnsetLogBackupNumStreams ensures that no value is present for LogBackupNumStreams, not even an explicit nil
### GetLogBackupWithClause

`func (o *CommonMssqlFileObjectProtectionParams) GetLogBackupWithClause() string`

GetLogBackupWithClause returns the LogBackupWithClause field if non-nil, zero value otherwise.

### GetLogBackupWithClauseOk

`func (o *CommonMssqlFileObjectProtectionParams) GetLogBackupWithClauseOk() (*string, bool)`

GetLogBackupWithClauseOk returns a tuple with the LogBackupWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupWithClause

`func (o *CommonMssqlFileObjectProtectionParams) SetLogBackupWithClause(v string)`

SetLogBackupWithClause sets LogBackupWithClause field to given value.

### HasLogBackupWithClause

`func (o *CommonMssqlFileObjectProtectionParams) HasLogBackupWithClause() bool`

HasLogBackupWithClause returns a boolean if a field has been set.

### SetLogBackupWithClauseNil

`func (o *CommonMssqlFileObjectProtectionParams) SetLogBackupWithClauseNil(b bool)`

 SetLogBackupWithClauseNil sets the value for LogBackupWithClause to be an explicit nil

### UnsetLogBackupWithClause
`func (o *CommonMssqlFileObjectProtectionParams) UnsetLogBackupWithClause()`

UnsetLogBackupWithClause ensures that no value is present for LogBackupWithClause, not even an explicit nil
### GetPrePostScript

`func (o *CommonMssqlFileObjectProtectionParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *CommonMssqlFileObjectProtectionParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *CommonMssqlFileObjectProtectionParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *CommonMssqlFileObjectProtectionParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetUseAagPreferencesFromServer

`func (o *CommonMssqlFileObjectProtectionParams) GetUseAagPreferencesFromServer() bool`

GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromServerOk

`func (o *CommonMssqlFileObjectProtectionParams) GetUseAagPreferencesFromServerOk() (*bool, bool)`

GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromServer

`func (o *CommonMssqlFileObjectProtectionParams) SetUseAagPreferencesFromServer(v bool)`

SetUseAagPreferencesFromServer sets UseAagPreferencesFromServer field to given value.

### HasUseAagPreferencesFromServer

`func (o *CommonMssqlFileObjectProtectionParams) HasUseAagPreferencesFromServer() bool`

HasUseAagPreferencesFromServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromServerNil

`func (o *CommonMssqlFileObjectProtectionParams) SetUseAagPreferencesFromServerNil(b bool)`

 SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil

### UnsetUseAagPreferencesFromServer
`func (o *CommonMssqlFileObjectProtectionParams) UnsetUseAagPreferencesFromServer()`

UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
### GetUserDbBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) GetUserDbBackupPreferenceType() string`

GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field if non-nil, zero value otherwise.

### GetUserDbBackupPreferenceTypeOk

`func (o *CommonMssqlFileObjectProtectionParams) GetUserDbBackupPreferenceTypeOk() (*string, bool)`

GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) SetUserDbBackupPreferenceType(v string)`

SetUserDbBackupPreferenceType sets UserDbBackupPreferenceType field to given value.

### HasUserDbBackupPreferenceType

`func (o *CommonMssqlFileObjectProtectionParams) HasUserDbBackupPreferenceType() bool`

HasUserDbBackupPreferenceType returns a boolean if a field has been set.

### SetUserDbBackupPreferenceTypeNil

`func (o *CommonMssqlFileObjectProtectionParams) SetUserDbBackupPreferenceTypeNil(b bool)`

 SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil

### UnsetUserDbBackupPreferenceType
`func (o *CommonMssqlFileObjectProtectionParams) UnsetUserDbBackupPreferenceType()`

UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


