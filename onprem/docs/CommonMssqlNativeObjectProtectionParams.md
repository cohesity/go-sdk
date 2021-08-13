# CommonMssqlNativeObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumStreams** | Pointer to **NullableInt32** | Specifies the number of streams to be used. | [optional] 
**WithClause** | Pointer to **NullableString** | Specifies the WithClause to be used. | [optional] 
**UserDbBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up user databases on the host. | [optional] 
**BackupSystemDbs** | Pointer to **NullableBool** | Specifies whether to backup system databases. If not specified then parameter is set to true. | [optional] 
**UseAagPreferencesFromServer** | Pointer to **NullableBool** | Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used. | [optional] 
**AagBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases. | [optional] 
**FullBackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether full backups should be copy-only. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**ExcludeFilters** | Pointer to [**[]Filter**](Filter.md) | Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters. | [optional] 

## Methods

### NewCommonMssqlNativeObjectProtectionParams

`func NewCommonMssqlNativeObjectProtectionParams() *CommonMssqlNativeObjectProtectionParams`

NewCommonMssqlNativeObjectProtectionParams instantiates a new CommonMssqlNativeObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMssqlNativeObjectProtectionParamsWithDefaults

`func NewCommonMssqlNativeObjectProtectionParamsWithDefaults() *CommonMssqlNativeObjectProtectionParams`

NewCommonMssqlNativeObjectProtectionParamsWithDefaults instantiates a new CommonMssqlNativeObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumStreams

`func (o *CommonMssqlNativeObjectProtectionParams) GetNumStreams() int32`

GetNumStreams returns the NumStreams field if non-nil, zero value otherwise.

### GetNumStreamsOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetNumStreamsOk() (*int32, bool)`

GetNumStreamsOk returns a tuple with the NumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStreams

`func (o *CommonMssqlNativeObjectProtectionParams) SetNumStreams(v int32)`

SetNumStreams sets NumStreams field to given value.

### HasNumStreams

`func (o *CommonMssqlNativeObjectProtectionParams) HasNumStreams() bool`

HasNumStreams returns a boolean if a field has been set.

### SetNumStreamsNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetNumStreamsNil(b bool)`

 SetNumStreamsNil sets the value for NumStreams to be an explicit nil

### UnsetNumStreams
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetNumStreams()`

UnsetNumStreams ensures that no value is present for NumStreams, not even an explicit nil
### GetWithClause

`func (o *CommonMssqlNativeObjectProtectionParams) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *CommonMssqlNativeObjectProtectionParams) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *CommonMssqlNativeObjectProtectionParams) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil
### GetUserDbBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) GetUserDbBackupPreferenceType() string`

GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field if non-nil, zero value otherwise.

### GetUserDbBackupPreferenceTypeOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetUserDbBackupPreferenceTypeOk() (*string, bool)`

GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) SetUserDbBackupPreferenceType(v string)`

SetUserDbBackupPreferenceType sets UserDbBackupPreferenceType field to given value.

### HasUserDbBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) HasUserDbBackupPreferenceType() bool`

HasUserDbBackupPreferenceType returns a boolean if a field has been set.

### SetUserDbBackupPreferenceTypeNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetUserDbBackupPreferenceTypeNil(b bool)`

 SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil

### UnsetUserDbBackupPreferenceType
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetUserDbBackupPreferenceType()`

UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
### GetBackupSystemDbs

`func (o *CommonMssqlNativeObjectProtectionParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *CommonMssqlNativeObjectProtectionParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *CommonMssqlNativeObjectProtectionParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetUseAagPreferencesFromServer

`func (o *CommonMssqlNativeObjectProtectionParams) GetUseAagPreferencesFromServer() bool`

GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromServerOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetUseAagPreferencesFromServerOk() (*bool, bool)`

GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromServer

`func (o *CommonMssqlNativeObjectProtectionParams) SetUseAagPreferencesFromServer(v bool)`

SetUseAagPreferencesFromServer sets UseAagPreferencesFromServer field to given value.

### HasUseAagPreferencesFromServer

`func (o *CommonMssqlNativeObjectProtectionParams) HasUseAagPreferencesFromServer() bool`

HasUseAagPreferencesFromServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromServerNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetUseAagPreferencesFromServerNil(b bool)`

 SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil

### UnsetUseAagPreferencesFromServer
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetUseAagPreferencesFromServer()`

UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
### GetAagBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) GetAagBackupPreferenceType() string`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetAagBackupPreferenceTypeOk() (*string, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) SetAagBackupPreferenceType(v string)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *CommonMssqlNativeObjectProtectionParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetFullBackupsCopyOnly

`func (o *CommonMssqlNativeObjectProtectionParams) GetFullBackupsCopyOnly() bool`

GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field if non-nil, zero value otherwise.

### GetFullBackupsCopyOnlyOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetFullBackupsCopyOnlyOk() (*bool, bool)`

GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupsCopyOnly

`func (o *CommonMssqlNativeObjectProtectionParams) SetFullBackupsCopyOnly(v bool)`

SetFullBackupsCopyOnly sets FullBackupsCopyOnly field to given value.

### HasFullBackupsCopyOnly

`func (o *CommonMssqlNativeObjectProtectionParams) HasFullBackupsCopyOnly() bool`

HasFullBackupsCopyOnly returns a boolean if a field has been set.

### SetFullBackupsCopyOnlyNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetFullBackupsCopyOnlyNil(b bool)`

 SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil

### UnsetFullBackupsCopyOnly
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetFullBackupsCopyOnly()`

UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
### GetPrePostScript

`func (o *CommonMssqlNativeObjectProtectionParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *CommonMssqlNativeObjectProtectionParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *CommonMssqlNativeObjectProtectionParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetExcludeFilters

`func (o *CommonMssqlNativeObjectProtectionParams) GetExcludeFilters() []Filter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *CommonMssqlNativeObjectProtectionParams) GetExcludeFiltersOk() (*[]Filter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *CommonMssqlNativeObjectProtectionParams) SetExcludeFilters(v []Filter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *CommonMssqlNativeObjectProtectionParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *CommonMssqlNativeObjectProtectionParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *CommonMssqlNativeObjectProtectionParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


