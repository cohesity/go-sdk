# AdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClonedDbBackupStatus** | Pointer to **NullableString** | Whether to report error if SQL database is cloned. | [optional] 
**DbBackupIfNotOnlineStatus** | Pointer to **NullableString** | Whether to report error if SQL database is not online. | [optional] 
**LogChainBreakAutoTriggerOobIncrBackup** | Pointer to **NullableBool** | If set to true, out of band incremental backup will be started when the log chain is broken and it would be started at the end of the log backup. Default value is false. | [optional] 
**MissingDbBackupStatus** | Pointer to **NullableString** | Fail the backup job when the database is missing. The database may be missing if it is deleted or corrupted. | [optional] 
**NewDatabaseAutoTriggerOobIncrBackup** | Pointer to **NullableBool** | If set to true, out of band incremental backup will be triggered when a new database is found and it would be started at the end of the log backup. Default value is false. | [optional] 
**OfflineRestoringDbBackupStatus** | Pointer to **NullableString** | Fail the backup job when database is offline or restoring. | [optional] 
**ReadOnlyDbBackupStatus** | Pointer to **NullableString** | Whether to skip backup for read-only SQL databases. | [optional] 
**ReportAllNonAutoprotectDbErrors** | Pointer to **NullableString** | Whether to report error for all dbs in non-autoprotect jobs. | [optional] 

## Methods

### NewAdvancedSettings

`func NewAdvancedSettings() *AdvancedSettings`

NewAdvancedSettings instantiates a new AdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSettingsWithDefaults

`func NewAdvancedSettingsWithDefaults() *AdvancedSettings`

NewAdvancedSettingsWithDefaults instantiates a new AdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClonedDbBackupStatus

`func (o *AdvancedSettings) GetClonedDbBackupStatus() string`

GetClonedDbBackupStatus returns the ClonedDbBackupStatus field if non-nil, zero value otherwise.

### GetClonedDbBackupStatusOk

`func (o *AdvancedSettings) GetClonedDbBackupStatusOk() (*string, bool)`

GetClonedDbBackupStatusOk returns a tuple with the ClonedDbBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedDbBackupStatus

`func (o *AdvancedSettings) SetClonedDbBackupStatus(v string)`

SetClonedDbBackupStatus sets ClonedDbBackupStatus field to given value.

### HasClonedDbBackupStatus

`func (o *AdvancedSettings) HasClonedDbBackupStatus() bool`

HasClonedDbBackupStatus returns a boolean if a field has been set.

### SetClonedDbBackupStatusNil

`func (o *AdvancedSettings) SetClonedDbBackupStatusNil(b bool)`

 SetClonedDbBackupStatusNil sets the value for ClonedDbBackupStatus to be an explicit nil

### UnsetClonedDbBackupStatus
`func (o *AdvancedSettings) UnsetClonedDbBackupStatus()`

UnsetClonedDbBackupStatus ensures that no value is present for ClonedDbBackupStatus, not even an explicit nil
### GetDbBackupIfNotOnlineStatus

`func (o *AdvancedSettings) GetDbBackupIfNotOnlineStatus() string`

GetDbBackupIfNotOnlineStatus returns the DbBackupIfNotOnlineStatus field if non-nil, zero value otherwise.

### GetDbBackupIfNotOnlineStatusOk

`func (o *AdvancedSettings) GetDbBackupIfNotOnlineStatusOk() (*string, bool)`

GetDbBackupIfNotOnlineStatusOk returns a tuple with the DbBackupIfNotOnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbBackupIfNotOnlineStatus

`func (o *AdvancedSettings) SetDbBackupIfNotOnlineStatus(v string)`

SetDbBackupIfNotOnlineStatus sets DbBackupIfNotOnlineStatus field to given value.

### HasDbBackupIfNotOnlineStatus

`func (o *AdvancedSettings) HasDbBackupIfNotOnlineStatus() bool`

HasDbBackupIfNotOnlineStatus returns a boolean if a field has been set.

### SetDbBackupIfNotOnlineStatusNil

`func (o *AdvancedSettings) SetDbBackupIfNotOnlineStatusNil(b bool)`

 SetDbBackupIfNotOnlineStatusNil sets the value for DbBackupIfNotOnlineStatus to be an explicit nil

### UnsetDbBackupIfNotOnlineStatus
`func (o *AdvancedSettings) UnsetDbBackupIfNotOnlineStatus()`

UnsetDbBackupIfNotOnlineStatus ensures that no value is present for DbBackupIfNotOnlineStatus, not even an explicit nil
### GetLogChainBreakAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) GetLogChainBreakAutoTriggerOobIncrBackup() bool`

GetLogChainBreakAutoTriggerOobIncrBackup returns the LogChainBreakAutoTriggerOobIncrBackup field if non-nil, zero value otherwise.

### GetLogChainBreakAutoTriggerOobIncrBackupOk

`func (o *AdvancedSettings) GetLogChainBreakAutoTriggerOobIncrBackupOk() (*bool, bool)`

GetLogChainBreakAutoTriggerOobIncrBackupOk returns a tuple with the LogChainBreakAutoTriggerOobIncrBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogChainBreakAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) SetLogChainBreakAutoTriggerOobIncrBackup(v bool)`

SetLogChainBreakAutoTriggerOobIncrBackup sets LogChainBreakAutoTriggerOobIncrBackup field to given value.

### HasLogChainBreakAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) HasLogChainBreakAutoTriggerOobIncrBackup() bool`

HasLogChainBreakAutoTriggerOobIncrBackup returns a boolean if a field has been set.

### SetLogChainBreakAutoTriggerOobIncrBackupNil

`func (o *AdvancedSettings) SetLogChainBreakAutoTriggerOobIncrBackupNil(b bool)`

 SetLogChainBreakAutoTriggerOobIncrBackupNil sets the value for LogChainBreakAutoTriggerOobIncrBackup to be an explicit nil

### UnsetLogChainBreakAutoTriggerOobIncrBackup
`func (o *AdvancedSettings) UnsetLogChainBreakAutoTriggerOobIncrBackup()`

UnsetLogChainBreakAutoTriggerOobIncrBackup ensures that no value is present for LogChainBreakAutoTriggerOobIncrBackup, not even an explicit nil
### GetMissingDbBackupStatus

`func (o *AdvancedSettings) GetMissingDbBackupStatus() string`

GetMissingDbBackupStatus returns the MissingDbBackupStatus field if non-nil, zero value otherwise.

### GetMissingDbBackupStatusOk

`func (o *AdvancedSettings) GetMissingDbBackupStatusOk() (*string, bool)`

GetMissingDbBackupStatusOk returns a tuple with the MissingDbBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingDbBackupStatus

`func (o *AdvancedSettings) SetMissingDbBackupStatus(v string)`

SetMissingDbBackupStatus sets MissingDbBackupStatus field to given value.

### HasMissingDbBackupStatus

`func (o *AdvancedSettings) HasMissingDbBackupStatus() bool`

HasMissingDbBackupStatus returns a boolean if a field has been set.

### SetMissingDbBackupStatusNil

`func (o *AdvancedSettings) SetMissingDbBackupStatusNil(b bool)`

 SetMissingDbBackupStatusNil sets the value for MissingDbBackupStatus to be an explicit nil

### UnsetMissingDbBackupStatus
`func (o *AdvancedSettings) UnsetMissingDbBackupStatus()`

UnsetMissingDbBackupStatus ensures that no value is present for MissingDbBackupStatus, not even an explicit nil
### GetNewDatabaseAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) GetNewDatabaseAutoTriggerOobIncrBackup() bool`

GetNewDatabaseAutoTriggerOobIncrBackup returns the NewDatabaseAutoTriggerOobIncrBackup field if non-nil, zero value otherwise.

### GetNewDatabaseAutoTriggerOobIncrBackupOk

`func (o *AdvancedSettings) GetNewDatabaseAutoTriggerOobIncrBackupOk() (*bool, bool)`

GetNewDatabaseAutoTriggerOobIncrBackupOk returns a tuple with the NewDatabaseAutoTriggerOobIncrBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) SetNewDatabaseAutoTriggerOobIncrBackup(v bool)`

SetNewDatabaseAutoTriggerOobIncrBackup sets NewDatabaseAutoTriggerOobIncrBackup field to given value.

### HasNewDatabaseAutoTriggerOobIncrBackup

`func (o *AdvancedSettings) HasNewDatabaseAutoTriggerOobIncrBackup() bool`

HasNewDatabaseAutoTriggerOobIncrBackup returns a boolean if a field has been set.

### SetNewDatabaseAutoTriggerOobIncrBackupNil

`func (o *AdvancedSettings) SetNewDatabaseAutoTriggerOobIncrBackupNil(b bool)`

 SetNewDatabaseAutoTriggerOobIncrBackupNil sets the value for NewDatabaseAutoTriggerOobIncrBackup to be an explicit nil

### UnsetNewDatabaseAutoTriggerOobIncrBackup
`func (o *AdvancedSettings) UnsetNewDatabaseAutoTriggerOobIncrBackup()`

UnsetNewDatabaseAutoTriggerOobIncrBackup ensures that no value is present for NewDatabaseAutoTriggerOobIncrBackup, not even an explicit nil
### GetOfflineRestoringDbBackupStatus

`func (o *AdvancedSettings) GetOfflineRestoringDbBackupStatus() string`

GetOfflineRestoringDbBackupStatus returns the OfflineRestoringDbBackupStatus field if non-nil, zero value otherwise.

### GetOfflineRestoringDbBackupStatusOk

`func (o *AdvancedSettings) GetOfflineRestoringDbBackupStatusOk() (*string, bool)`

GetOfflineRestoringDbBackupStatusOk returns a tuple with the OfflineRestoringDbBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineRestoringDbBackupStatus

`func (o *AdvancedSettings) SetOfflineRestoringDbBackupStatus(v string)`

SetOfflineRestoringDbBackupStatus sets OfflineRestoringDbBackupStatus field to given value.

### HasOfflineRestoringDbBackupStatus

`func (o *AdvancedSettings) HasOfflineRestoringDbBackupStatus() bool`

HasOfflineRestoringDbBackupStatus returns a boolean if a field has been set.

### SetOfflineRestoringDbBackupStatusNil

`func (o *AdvancedSettings) SetOfflineRestoringDbBackupStatusNil(b bool)`

 SetOfflineRestoringDbBackupStatusNil sets the value for OfflineRestoringDbBackupStatus to be an explicit nil

### UnsetOfflineRestoringDbBackupStatus
`func (o *AdvancedSettings) UnsetOfflineRestoringDbBackupStatus()`

UnsetOfflineRestoringDbBackupStatus ensures that no value is present for OfflineRestoringDbBackupStatus, not even an explicit nil
### GetReadOnlyDbBackupStatus

`func (o *AdvancedSettings) GetReadOnlyDbBackupStatus() string`

GetReadOnlyDbBackupStatus returns the ReadOnlyDbBackupStatus field if non-nil, zero value otherwise.

### GetReadOnlyDbBackupStatusOk

`func (o *AdvancedSettings) GetReadOnlyDbBackupStatusOk() (*string, bool)`

GetReadOnlyDbBackupStatusOk returns a tuple with the ReadOnlyDbBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyDbBackupStatus

`func (o *AdvancedSettings) SetReadOnlyDbBackupStatus(v string)`

SetReadOnlyDbBackupStatus sets ReadOnlyDbBackupStatus field to given value.

### HasReadOnlyDbBackupStatus

`func (o *AdvancedSettings) HasReadOnlyDbBackupStatus() bool`

HasReadOnlyDbBackupStatus returns a boolean if a field has been set.

### SetReadOnlyDbBackupStatusNil

`func (o *AdvancedSettings) SetReadOnlyDbBackupStatusNil(b bool)`

 SetReadOnlyDbBackupStatusNil sets the value for ReadOnlyDbBackupStatus to be an explicit nil

### UnsetReadOnlyDbBackupStatus
`func (o *AdvancedSettings) UnsetReadOnlyDbBackupStatus()`

UnsetReadOnlyDbBackupStatus ensures that no value is present for ReadOnlyDbBackupStatus, not even an explicit nil
### GetReportAllNonAutoprotectDbErrors

`func (o *AdvancedSettings) GetReportAllNonAutoprotectDbErrors() string`

GetReportAllNonAutoprotectDbErrors returns the ReportAllNonAutoprotectDbErrors field if non-nil, zero value otherwise.

### GetReportAllNonAutoprotectDbErrorsOk

`func (o *AdvancedSettings) GetReportAllNonAutoprotectDbErrorsOk() (*string, bool)`

GetReportAllNonAutoprotectDbErrorsOk returns a tuple with the ReportAllNonAutoprotectDbErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportAllNonAutoprotectDbErrors

`func (o *AdvancedSettings) SetReportAllNonAutoprotectDbErrors(v string)`

SetReportAllNonAutoprotectDbErrors sets ReportAllNonAutoprotectDbErrors field to given value.

### HasReportAllNonAutoprotectDbErrors

`func (o *AdvancedSettings) HasReportAllNonAutoprotectDbErrors() bool`

HasReportAllNonAutoprotectDbErrors returns a boolean if a field has been set.

### SetReportAllNonAutoprotectDbErrorsNil

`func (o *AdvancedSettings) SetReportAllNonAutoprotectDbErrorsNil(b bool)`

 SetReportAllNonAutoprotectDbErrorsNil sets the value for ReportAllNonAutoprotectDbErrors to be an explicit nil

### UnsetReportAllNonAutoprotectDbErrors
`func (o *AdvancedSettings) UnsetReportAllNonAutoprotectDbErrors()`

UnsetReportAllNonAutoprotectDbErrors ensures that no value is present for ReportAllNonAutoprotectDbErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


