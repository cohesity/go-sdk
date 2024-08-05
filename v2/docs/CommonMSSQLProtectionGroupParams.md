# CommonMSSQLProtectionGroupParams

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

### NewCommonMSSQLProtectionGroupParams

`func NewCommonMSSQLProtectionGroupParams() *CommonMSSQLProtectionGroupParams`

NewCommonMSSQLProtectionGroupParams instantiates a new CommonMSSQLProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMSSQLProtectionGroupParamsWithDefaults

`func NewCommonMSSQLProtectionGroupParamsWithDefaults() *CommonMSSQLProtectionGroupParams`

NewCommonMSSQLProtectionGroupParamsWithDefaults instantiates a new CommonMSSQLProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceType() string`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *CommonMSSQLProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) SetAagBackupPreferenceType(v string)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *CommonMSSQLProtectionGroupParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *CommonMSSQLProtectionGroupParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetAdvancedSettings

`func (o *CommonMSSQLProtectionGroupParams) GetAdvancedSettings() AdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *CommonMSSQLProtectionGroupParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *CommonMSSQLProtectionGroupParams) SetAdvancedSettings(v AdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *CommonMSSQLProtectionGroupParams) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetBackupSystemDbs

`func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *CommonMSSQLProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *CommonMSSQLProtectionGroupParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *CommonMSSQLProtectionGroupParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *CommonMSSQLProtectionGroupParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *CommonMSSQLProtectionGroupParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetExcludeFilters

`func (o *CommonMSSQLProtectionGroupParams) GetExcludeFilters() []Filter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *CommonMSSQLProtectionGroupParams) GetExcludeFiltersOk() (*[]Filter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *CommonMSSQLProtectionGroupParams) SetExcludeFilters(v []Filter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *CommonMSSQLProtectionGroupParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *CommonMSSQLProtectionGroupParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *CommonMSSQLProtectionGroupParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetFullBackupsCopyOnly

`func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnly() bool`

GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field if non-nil, zero value otherwise.

### GetFullBackupsCopyOnlyOk

`func (o *CommonMSSQLProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool)`

GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupsCopyOnly

`func (o *CommonMSSQLProtectionGroupParams) SetFullBackupsCopyOnly(v bool)`

SetFullBackupsCopyOnly sets FullBackupsCopyOnly field to given value.

### HasFullBackupsCopyOnly

`func (o *CommonMSSQLProtectionGroupParams) HasFullBackupsCopyOnly() bool`

HasFullBackupsCopyOnly returns a boolean if a field has been set.

### SetFullBackupsCopyOnlyNil

`func (o *CommonMSSQLProtectionGroupParams) SetFullBackupsCopyOnlyNil(b bool)`

 SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil

### UnsetFullBackupsCopyOnly
`func (o *CommonMSSQLProtectionGroupParams) UnsetFullBackupsCopyOnly()`

UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
### GetLogBackupNumStreams

`func (o *CommonMSSQLProtectionGroupParams) GetLogBackupNumStreams() int32`

GetLogBackupNumStreams returns the LogBackupNumStreams field if non-nil, zero value otherwise.

### GetLogBackupNumStreamsOk

`func (o *CommonMSSQLProtectionGroupParams) GetLogBackupNumStreamsOk() (*int32, bool)`

GetLogBackupNumStreamsOk returns a tuple with the LogBackupNumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupNumStreams

`func (o *CommonMSSQLProtectionGroupParams) SetLogBackupNumStreams(v int32)`

SetLogBackupNumStreams sets LogBackupNumStreams field to given value.

### HasLogBackupNumStreams

`func (o *CommonMSSQLProtectionGroupParams) HasLogBackupNumStreams() bool`

HasLogBackupNumStreams returns a boolean if a field has been set.

### SetLogBackupNumStreamsNil

`func (o *CommonMSSQLProtectionGroupParams) SetLogBackupNumStreamsNil(b bool)`

 SetLogBackupNumStreamsNil sets the value for LogBackupNumStreams to be an explicit nil

### UnsetLogBackupNumStreams
`func (o *CommonMSSQLProtectionGroupParams) UnsetLogBackupNumStreams()`

UnsetLogBackupNumStreams ensures that no value is present for LogBackupNumStreams, not even an explicit nil
### GetLogBackupWithClause

`func (o *CommonMSSQLProtectionGroupParams) GetLogBackupWithClause() string`

GetLogBackupWithClause returns the LogBackupWithClause field if non-nil, zero value otherwise.

### GetLogBackupWithClauseOk

`func (o *CommonMSSQLProtectionGroupParams) GetLogBackupWithClauseOk() (*string, bool)`

GetLogBackupWithClauseOk returns a tuple with the LogBackupWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupWithClause

`func (o *CommonMSSQLProtectionGroupParams) SetLogBackupWithClause(v string)`

SetLogBackupWithClause sets LogBackupWithClause field to given value.

### HasLogBackupWithClause

`func (o *CommonMSSQLProtectionGroupParams) HasLogBackupWithClause() bool`

HasLogBackupWithClause returns a boolean if a field has been set.

### SetLogBackupWithClauseNil

`func (o *CommonMSSQLProtectionGroupParams) SetLogBackupWithClauseNil(b bool)`

 SetLogBackupWithClauseNil sets the value for LogBackupWithClause to be an explicit nil

### UnsetLogBackupWithClause
`func (o *CommonMSSQLProtectionGroupParams) UnsetLogBackupWithClause()`

UnsetLogBackupWithClause ensures that no value is present for LogBackupWithClause, not even an explicit nil
### GetPrePostScript

`func (o *CommonMSSQLProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *CommonMSSQLProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *CommonMSSQLProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *CommonMSSQLProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetUseAagPreferencesFromServer

`func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServer() bool`

GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromServerOk

`func (o *CommonMSSQLProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool)`

GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromServer

`func (o *CommonMSSQLProtectionGroupParams) SetUseAagPreferencesFromServer(v bool)`

SetUseAagPreferencesFromServer sets UseAagPreferencesFromServer field to given value.

### HasUseAagPreferencesFromServer

`func (o *CommonMSSQLProtectionGroupParams) HasUseAagPreferencesFromServer() bool`

HasUseAagPreferencesFromServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromServerNil

`func (o *CommonMSSQLProtectionGroupParams) SetUseAagPreferencesFromServerNil(b bool)`

 SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil

### UnsetUseAagPreferencesFromServer
`func (o *CommonMSSQLProtectionGroupParams) UnsetUseAagPreferencesFromServer()`

UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
### GetUserDbBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceType() string`

GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field if non-nil, zero value otherwise.

### GetUserDbBackupPreferenceTypeOk

`func (o *CommonMSSQLProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool)`

GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) SetUserDbBackupPreferenceType(v string)`

SetUserDbBackupPreferenceType sets UserDbBackupPreferenceType field to given value.

### HasUserDbBackupPreferenceType

`func (o *CommonMSSQLProtectionGroupParams) HasUserDbBackupPreferenceType() bool`

HasUserDbBackupPreferenceType returns a boolean if a field has been set.

### SetUserDbBackupPreferenceTypeNil

`func (o *CommonMSSQLProtectionGroupParams) SetUserDbBackupPreferenceTypeNil(b bool)`

 SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil

### UnsetUserDbBackupPreferenceType
`func (o *CommonMSSQLProtectionGroupParams) UnsetUserDbBackupPreferenceType()`

UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


