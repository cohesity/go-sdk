# MSSQLVolumeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]MSSQLVolumeProtectionGroupObjectParams**](MSSQLVolumeProtectionGroupObjectParams.md) | Specifies the list of object ids to be protected. | 
**IncrementalBackupAfterRestart** | Pointer to **NullableBool** | Specifies whether or to perform incremental backups the first time after a server restarts. By default, a full backup will be performed. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**BackupDbVolumesOnly** | Pointer to **NullableBool** | Specifies whether to only backup volumes on which the specified databases reside. If not specified (default), all the volumes of the host will be protected. | [optional] 
**AdditionalHostParams** | Pointer to [**[]MSSQLVolumeProtectionGroupHostParams**](MSSQLVolumeProtectionGroupHostParams.md) | Specifies settings which are to be applied to specific host containers in this protection group. | [optional] 
**UserDbBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up user databases on the host. | [optional] 
**BackupSystemDbs** | Pointer to **NullableBool** | Specifies whether to backup system databases. If not specified then parameter is set to true. | [optional] 
**UseAagPreferencesFromServer** | Pointer to **NullableBool** | Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used. | [optional] 
**AagBackupPreferenceType** | Pointer to **NullableString** | Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases. | [optional] 
**FullBackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether full backups should be copy-only. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**ExcludeFilters** | Pointer to [**[]Filter**](Filter.md) | Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters. | [optional] 

## Methods

### NewMSSQLVolumeProtectionGroupParams

`func NewMSSQLVolumeProtectionGroupParams(objects []MSSQLVolumeProtectionGroupObjectParams, ) *MSSQLVolumeProtectionGroupParams`

NewMSSQLVolumeProtectionGroupParams instantiates a new MSSQLVolumeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLVolumeProtectionGroupParamsWithDefaults

`func NewMSSQLVolumeProtectionGroupParamsWithDefaults() *MSSQLVolumeProtectionGroupParams`

NewMSSQLVolumeProtectionGroupParamsWithDefaults instantiates a new MSSQLVolumeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MSSQLVolumeProtectionGroupParams) GetObjects() []MSSQLVolumeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MSSQLVolumeProtectionGroupParams) GetObjectsOk() (*[]MSSQLVolumeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MSSQLVolumeProtectionGroupParams) SetObjects(v []MSSQLVolumeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MSSQLVolumeProtectionGroupParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MSSQLVolumeProtectionGroupParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParams) GetIncrementalBackupAfterRestart() bool`

GetIncrementalBackupAfterRestart returns the IncrementalBackupAfterRestart field if non-nil, zero value otherwise.

### GetIncrementalBackupAfterRestartOk

`func (o *MSSQLVolumeProtectionGroupParams) GetIncrementalBackupAfterRestartOk() (*bool, bool)`

GetIncrementalBackupAfterRestartOk returns a tuple with the IncrementalBackupAfterRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParams) SetIncrementalBackupAfterRestart(v bool)`

SetIncrementalBackupAfterRestart sets IncrementalBackupAfterRestart field to given value.

### HasIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParams) HasIncrementalBackupAfterRestart() bool`

HasIncrementalBackupAfterRestart returns a boolean if a field has been set.

### SetIncrementalBackupAfterRestartNil

`func (o *MSSQLVolumeProtectionGroupParams) SetIncrementalBackupAfterRestartNil(b bool)`

 SetIncrementalBackupAfterRestartNil sets the value for IncrementalBackupAfterRestart to be an explicit nil

### UnsetIncrementalBackupAfterRestart
`func (o *MSSQLVolumeProtectionGroupParams) UnsetIncrementalBackupAfterRestart()`

UnsetIncrementalBackupAfterRestart ensures that no value is present for IncrementalBackupAfterRestart, not even an explicit nil
### GetIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *MSSQLVolumeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParams) GetBackupDbVolumesOnly() bool`

GetBackupDbVolumesOnly returns the BackupDbVolumesOnly field if non-nil, zero value otherwise.

### GetBackupDbVolumesOnlyOk

`func (o *MSSQLVolumeProtectionGroupParams) GetBackupDbVolumesOnlyOk() (*bool, bool)`

GetBackupDbVolumesOnlyOk returns a tuple with the BackupDbVolumesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParams) SetBackupDbVolumesOnly(v bool)`

SetBackupDbVolumesOnly sets BackupDbVolumesOnly field to given value.

### HasBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParams) HasBackupDbVolumesOnly() bool`

HasBackupDbVolumesOnly returns a boolean if a field has been set.

### SetBackupDbVolumesOnlyNil

`func (o *MSSQLVolumeProtectionGroupParams) SetBackupDbVolumesOnlyNil(b bool)`

 SetBackupDbVolumesOnlyNil sets the value for BackupDbVolumesOnly to be an explicit nil

### UnsetBackupDbVolumesOnly
`func (o *MSSQLVolumeProtectionGroupParams) UnsetBackupDbVolumesOnly()`

UnsetBackupDbVolumesOnly ensures that no value is present for BackupDbVolumesOnly, not even an explicit nil
### GetAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParams) GetAdditionalHostParams() []MSSQLVolumeProtectionGroupHostParams`

GetAdditionalHostParams returns the AdditionalHostParams field if non-nil, zero value otherwise.

### GetAdditionalHostParamsOk

`func (o *MSSQLVolumeProtectionGroupParams) GetAdditionalHostParamsOk() (*[]MSSQLVolumeProtectionGroupHostParams, bool)`

GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParams) SetAdditionalHostParams(v []MSSQLVolumeProtectionGroupHostParams)`

SetAdditionalHostParams sets AdditionalHostParams field to given value.

### HasAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParams) HasAdditionalHostParams() bool`

HasAdditionalHostParams returns a boolean if a field has been set.

### GetUserDbBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) GetUserDbBackupPreferenceType() string`

GetUserDbBackupPreferenceType returns the UserDbBackupPreferenceType field if non-nil, zero value otherwise.

### GetUserDbBackupPreferenceTypeOk

`func (o *MSSQLVolumeProtectionGroupParams) GetUserDbBackupPreferenceTypeOk() (*string, bool)`

GetUserDbBackupPreferenceTypeOk returns a tuple with the UserDbBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) SetUserDbBackupPreferenceType(v string)`

SetUserDbBackupPreferenceType sets UserDbBackupPreferenceType field to given value.

### HasUserDbBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) HasUserDbBackupPreferenceType() bool`

HasUserDbBackupPreferenceType returns a boolean if a field has been set.

### SetUserDbBackupPreferenceTypeNil

`func (o *MSSQLVolumeProtectionGroupParams) SetUserDbBackupPreferenceTypeNil(b bool)`

 SetUserDbBackupPreferenceTypeNil sets the value for UserDbBackupPreferenceType to be an explicit nil

### UnsetUserDbBackupPreferenceType
`func (o *MSSQLVolumeProtectionGroupParams) UnsetUserDbBackupPreferenceType()`

UnsetUserDbBackupPreferenceType ensures that no value is present for UserDbBackupPreferenceType, not even an explicit nil
### GetBackupSystemDbs

`func (o *MSSQLVolumeProtectionGroupParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *MSSQLVolumeProtectionGroupParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *MSSQLVolumeProtectionGroupParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *MSSQLVolumeProtectionGroupParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *MSSQLVolumeProtectionGroupParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *MSSQLVolumeProtectionGroupParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetUseAagPreferencesFromServer

`func (o *MSSQLVolumeProtectionGroupParams) GetUseAagPreferencesFromServer() bool`

GetUseAagPreferencesFromServer returns the UseAagPreferencesFromServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromServerOk

`func (o *MSSQLVolumeProtectionGroupParams) GetUseAagPreferencesFromServerOk() (*bool, bool)`

GetUseAagPreferencesFromServerOk returns a tuple with the UseAagPreferencesFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromServer

`func (o *MSSQLVolumeProtectionGroupParams) SetUseAagPreferencesFromServer(v bool)`

SetUseAagPreferencesFromServer sets UseAagPreferencesFromServer field to given value.

### HasUseAagPreferencesFromServer

`func (o *MSSQLVolumeProtectionGroupParams) HasUseAagPreferencesFromServer() bool`

HasUseAagPreferencesFromServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromServerNil

`func (o *MSSQLVolumeProtectionGroupParams) SetUseAagPreferencesFromServerNil(b bool)`

 SetUseAagPreferencesFromServerNil sets the value for UseAagPreferencesFromServer to be an explicit nil

### UnsetUseAagPreferencesFromServer
`func (o *MSSQLVolumeProtectionGroupParams) UnsetUseAagPreferencesFromServer()`

UnsetUseAagPreferencesFromServer ensures that no value is present for UseAagPreferencesFromServer, not even an explicit nil
### GetAagBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) GetAagBackupPreferenceType() string`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *MSSQLVolumeProtectionGroupParams) GetAagBackupPreferenceTypeOk() (*string, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) SetAagBackupPreferenceType(v string)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *MSSQLVolumeProtectionGroupParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *MSSQLVolumeProtectionGroupParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *MSSQLVolumeProtectionGroupParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetFullBackupsCopyOnly

`func (o *MSSQLVolumeProtectionGroupParams) GetFullBackupsCopyOnly() bool`

GetFullBackupsCopyOnly returns the FullBackupsCopyOnly field if non-nil, zero value otherwise.

### GetFullBackupsCopyOnlyOk

`func (o *MSSQLVolumeProtectionGroupParams) GetFullBackupsCopyOnlyOk() (*bool, bool)`

GetFullBackupsCopyOnlyOk returns a tuple with the FullBackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupsCopyOnly

`func (o *MSSQLVolumeProtectionGroupParams) SetFullBackupsCopyOnly(v bool)`

SetFullBackupsCopyOnly sets FullBackupsCopyOnly field to given value.

### HasFullBackupsCopyOnly

`func (o *MSSQLVolumeProtectionGroupParams) HasFullBackupsCopyOnly() bool`

HasFullBackupsCopyOnly returns a boolean if a field has been set.

### SetFullBackupsCopyOnlyNil

`func (o *MSSQLVolumeProtectionGroupParams) SetFullBackupsCopyOnlyNil(b bool)`

 SetFullBackupsCopyOnlyNil sets the value for FullBackupsCopyOnly to be an explicit nil

### UnsetFullBackupsCopyOnly
`func (o *MSSQLVolumeProtectionGroupParams) UnsetFullBackupsCopyOnly()`

UnsetFullBackupsCopyOnly ensures that no value is present for FullBackupsCopyOnly, not even an explicit nil
### GetPrePostScript

`func (o *MSSQLVolumeProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *MSSQLVolumeProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *MSSQLVolumeProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *MSSQLVolumeProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetExcludeFilters

`func (o *MSSQLVolumeProtectionGroupParams) GetExcludeFilters() []Filter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *MSSQLVolumeProtectionGroupParams) GetExcludeFiltersOk() (*[]Filter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *MSSQLVolumeProtectionGroupParams) SetExcludeFilters(v []Filter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *MSSQLVolumeProtectionGroupParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *MSSQLVolumeProtectionGroupParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *MSSQLVolumeProtectionGroupParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


