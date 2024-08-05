# RecoverCassandraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a recovery job. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 
**IsLiveTableRestore** | Pointer to **NullableBool** | Specifies whether the current recovery operation is a live table restore operation. | [optional] 
**IsSystemKeyspaceRestore** | Pointer to **NullableBool** | Specifies whether the current recovery operation is a system keyspace restore operation. | [optional] 
**LogRestoreDirectory** | Pointer to **NullableString** | Specifies the directory for restoring the logs. | [optional] 
**RecoverPrivileges** | Pointer to **NullableBool** | Specifies whether recover/skip roles and permissions. | [optional] 
**RestartAtUsecs** | Pointer to **NullableInt64** | Specifies the time in Unix epoch timestamp in microseconds at which the Cassandra services are to be restarted. | [optional] 
**RestartCommand** | Pointer to **NullableString** | Specifies the command to restart Cassandra services after the point in time recovery. | [optional] 
**RestartImmediately** | Pointer to **NullableBool** | Specifies whether to restart Cassandra services immediately after the point in time recovery. | [optional] 
**RestartServices** | Pointer to **NullableBool** | Specifies whether to restart Cassandra services after the point in time recovery. | [optional] 
**RestartServicesTaskId** | Pointer to **NullableInt64** | Specifies the Id of the task required to restart Cassandra services. | [optional] [readonly] 
**RunPreChecks** | Pointer to **NullableBool** | Specifies Whether to run checks before the recovery. E.x if there is sufficient space in the destination cluster for the recovery to succeed. | [optional] 
**SelectedDataCenters** | Pointer to **[]string** | Selected Data centers for this cluster. | [optional] 
**Snapshots** | [**[]RecoverCassandraSnapshotParams**](RecoverCassandraSnapshotParams.md) | Specifies the local snapshot ids and other details of the Objects to be recovered. | 
**StagingDirectoryList** | Pointer to **[]string** | Specifies the directory on the primary to copy the files which are to be uploaded using destination sstableloader. | [optional] 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 

## Methods

### NewRecoverCassandraParams

`func NewRecoverCassandraParams(snapshots []RecoverCassandraSnapshotParams, ) *RecoverCassandraParams`

NewRecoverCassandraParams instantiates a new RecoverCassandraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCassandraParamsWithDefaults

`func NewRecoverCassandraParamsWithDefaults() *RecoverCassandraParams`

NewRecoverCassandraParamsWithDefaults instantiates a new RecoverCassandraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfigs

`func (o *RecoverCassandraParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *RecoverCassandraParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *RecoverCassandraParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *RecoverCassandraParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *RecoverCassandraParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *RecoverCassandraParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetBandwidthMBPS

`func (o *RecoverCassandraParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *RecoverCassandraParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *RecoverCassandraParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *RecoverCassandraParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *RecoverCassandraParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *RecoverCassandraParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *RecoverCassandraParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverCassandraParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverCassandraParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverCassandraParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverCassandraParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverCassandraParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetOverwrite

`func (o *RecoverCassandraParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverCassandraParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverCassandraParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverCassandraParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverCassandraParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverCassandraParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRecoverTo

`func (o *RecoverCassandraParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverCassandraParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverCassandraParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverCassandraParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverCassandraParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverCassandraParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetWarnings

`func (o *RecoverCassandraParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverCassandraParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverCassandraParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverCassandraParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverCassandraParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverCassandraParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetIsLiveTableRestore

`func (o *RecoverCassandraParams) GetIsLiveTableRestore() bool`

GetIsLiveTableRestore returns the IsLiveTableRestore field if non-nil, zero value otherwise.

### GetIsLiveTableRestoreOk

`func (o *RecoverCassandraParams) GetIsLiveTableRestoreOk() (*bool, bool)`

GetIsLiveTableRestoreOk returns a tuple with the IsLiveTableRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLiveTableRestore

`func (o *RecoverCassandraParams) SetIsLiveTableRestore(v bool)`

SetIsLiveTableRestore sets IsLiveTableRestore field to given value.

### HasIsLiveTableRestore

`func (o *RecoverCassandraParams) HasIsLiveTableRestore() bool`

HasIsLiveTableRestore returns a boolean if a field has been set.

### SetIsLiveTableRestoreNil

`func (o *RecoverCassandraParams) SetIsLiveTableRestoreNil(b bool)`

 SetIsLiveTableRestoreNil sets the value for IsLiveTableRestore to be an explicit nil

### UnsetIsLiveTableRestore
`func (o *RecoverCassandraParams) UnsetIsLiveTableRestore()`

UnsetIsLiveTableRestore ensures that no value is present for IsLiveTableRestore, not even an explicit nil
### GetIsSystemKeyspaceRestore

`func (o *RecoverCassandraParams) GetIsSystemKeyspaceRestore() bool`

GetIsSystemKeyspaceRestore returns the IsSystemKeyspaceRestore field if non-nil, zero value otherwise.

### GetIsSystemKeyspaceRestoreOk

`func (o *RecoverCassandraParams) GetIsSystemKeyspaceRestoreOk() (*bool, bool)`

GetIsSystemKeyspaceRestoreOk returns a tuple with the IsSystemKeyspaceRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemKeyspaceRestore

`func (o *RecoverCassandraParams) SetIsSystemKeyspaceRestore(v bool)`

SetIsSystemKeyspaceRestore sets IsSystemKeyspaceRestore field to given value.

### HasIsSystemKeyspaceRestore

`func (o *RecoverCassandraParams) HasIsSystemKeyspaceRestore() bool`

HasIsSystemKeyspaceRestore returns a boolean if a field has been set.

### SetIsSystemKeyspaceRestoreNil

`func (o *RecoverCassandraParams) SetIsSystemKeyspaceRestoreNil(b bool)`

 SetIsSystemKeyspaceRestoreNil sets the value for IsSystemKeyspaceRestore to be an explicit nil

### UnsetIsSystemKeyspaceRestore
`func (o *RecoverCassandraParams) UnsetIsSystemKeyspaceRestore()`

UnsetIsSystemKeyspaceRestore ensures that no value is present for IsSystemKeyspaceRestore, not even an explicit nil
### GetLogRestoreDirectory

`func (o *RecoverCassandraParams) GetLogRestoreDirectory() string`

GetLogRestoreDirectory returns the LogRestoreDirectory field if non-nil, zero value otherwise.

### GetLogRestoreDirectoryOk

`func (o *RecoverCassandraParams) GetLogRestoreDirectoryOk() (*string, bool)`

GetLogRestoreDirectoryOk returns a tuple with the LogRestoreDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRestoreDirectory

`func (o *RecoverCassandraParams) SetLogRestoreDirectory(v string)`

SetLogRestoreDirectory sets LogRestoreDirectory field to given value.

### HasLogRestoreDirectory

`func (o *RecoverCassandraParams) HasLogRestoreDirectory() bool`

HasLogRestoreDirectory returns a boolean if a field has been set.

### SetLogRestoreDirectoryNil

`func (o *RecoverCassandraParams) SetLogRestoreDirectoryNil(b bool)`

 SetLogRestoreDirectoryNil sets the value for LogRestoreDirectory to be an explicit nil

### UnsetLogRestoreDirectory
`func (o *RecoverCassandraParams) UnsetLogRestoreDirectory()`

UnsetLogRestoreDirectory ensures that no value is present for LogRestoreDirectory, not even an explicit nil
### GetRecoverPrivileges

`func (o *RecoverCassandraParams) GetRecoverPrivileges() bool`

GetRecoverPrivileges returns the RecoverPrivileges field if non-nil, zero value otherwise.

### GetRecoverPrivilegesOk

`func (o *RecoverCassandraParams) GetRecoverPrivilegesOk() (*bool, bool)`

GetRecoverPrivilegesOk returns a tuple with the RecoverPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPrivileges

`func (o *RecoverCassandraParams) SetRecoverPrivileges(v bool)`

SetRecoverPrivileges sets RecoverPrivileges field to given value.

### HasRecoverPrivileges

`func (o *RecoverCassandraParams) HasRecoverPrivileges() bool`

HasRecoverPrivileges returns a boolean if a field has been set.

### SetRecoverPrivilegesNil

`func (o *RecoverCassandraParams) SetRecoverPrivilegesNil(b bool)`

 SetRecoverPrivilegesNil sets the value for RecoverPrivileges to be an explicit nil

### UnsetRecoverPrivileges
`func (o *RecoverCassandraParams) UnsetRecoverPrivileges()`

UnsetRecoverPrivileges ensures that no value is present for RecoverPrivileges, not even an explicit nil
### GetRestartAtUsecs

`func (o *RecoverCassandraParams) GetRestartAtUsecs() int64`

GetRestartAtUsecs returns the RestartAtUsecs field if non-nil, zero value otherwise.

### GetRestartAtUsecsOk

`func (o *RecoverCassandraParams) GetRestartAtUsecsOk() (*int64, bool)`

GetRestartAtUsecsOk returns a tuple with the RestartAtUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartAtUsecs

`func (o *RecoverCassandraParams) SetRestartAtUsecs(v int64)`

SetRestartAtUsecs sets RestartAtUsecs field to given value.

### HasRestartAtUsecs

`func (o *RecoverCassandraParams) HasRestartAtUsecs() bool`

HasRestartAtUsecs returns a boolean if a field has been set.

### SetRestartAtUsecsNil

`func (o *RecoverCassandraParams) SetRestartAtUsecsNil(b bool)`

 SetRestartAtUsecsNil sets the value for RestartAtUsecs to be an explicit nil

### UnsetRestartAtUsecs
`func (o *RecoverCassandraParams) UnsetRestartAtUsecs()`

UnsetRestartAtUsecs ensures that no value is present for RestartAtUsecs, not even an explicit nil
### GetRestartCommand

`func (o *RecoverCassandraParams) GetRestartCommand() string`

GetRestartCommand returns the RestartCommand field if non-nil, zero value otherwise.

### GetRestartCommandOk

`func (o *RecoverCassandraParams) GetRestartCommandOk() (*string, bool)`

GetRestartCommandOk returns a tuple with the RestartCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCommand

`func (o *RecoverCassandraParams) SetRestartCommand(v string)`

SetRestartCommand sets RestartCommand field to given value.

### HasRestartCommand

`func (o *RecoverCassandraParams) HasRestartCommand() bool`

HasRestartCommand returns a boolean if a field has been set.

### SetRestartCommandNil

`func (o *RecoverCassandraParams) SetRestartCommandNil(b bool)`

 SetRestartCommandNil sets the value for RestartCommand to be an explicit nil

### UnsetRestartCommand
`func (o *RecoverCassandraParams) UnsetRestartCommand()`

UnsetRestartCommand ensures that no value is present for RestartCommand, not even an explicit nil
### GetRestartImmediately

`func (o *RecoverCassandraParams) GetRestartImmediately() bool`

GetRestartImmediately returns the RestartImmediately field if non-nil, zero value otherwise.

### GetRestartImmediatelyOk

`func (o *RecoverCassandraParams) GetRestartImmediatelyOk() (*bool, bool)`

GetRestartImmediatelyOk returns a tuple with the RestartImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartImmediately

`func (o *RecoverCassandraParams) SetRestartImmediately(v bool)`

SetRestartImmediately sets RestartImmediately field to given value.

### HasRestartImmediately

`func (o *RecoverCassandraParams) HasRestartImmediately() bool`

HasRestartImmediately returns a boolean if a field has been set.

### SetRestartImmediatelyNil

`func (o *RecoverCassandraParams) SetRestartImmediatelyNil(b bool)`

 SetRestartImmediatelyNil sets the value for RestartImmediately to be an explicit nil

### UnsetRestartImmediately
`func (o *RecoverCassandraParams) UnsetRestartImmediately()`

UnsetRestartImmediately ensures that no value is present for RestartImmediately, not even an explicit nil
### GetRestartServices

`func (o *RecoverCassandraParams) GetRestartServices() bool`

GetRestartServices returns the RestartServices field if non-nil, zero value otherwise.

### GetRestartServicesOk

`func (o *RecoverCassandraParams) GetRestartServicesOk() (*bool, bool)`

GetRestartServicesOk returns a tuple with the RestartServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartServices

`func (o *RecoverCassandraParams) SetRestartServices(v bool)`

SetRestartServices sets RestartServices field to given value.

### HasRestartServices

`func (o *RecoverCassandraParams) HasRestartServices() bool`

HasRestartServices returns a boolean if a field has been set.

### SetRestartServicesNil

`func (o *RecoverCassandraParams) SetRestartServicesNil(b bool)`

 SetRestartServicesNil sets the value for RestartServices to be an explicit nil

### UnsetRestartServices
`func (o *RecoverCassandraParams) UnsetRestartServices()`

UnsetRestartServices ensures that no value is present for RestartServices, not even an explicit nil
### GetRestartServicesTaskId

`func (o *RecoverCassandraParams) GetRestartServicesTaskId() int64`

GetRestartServicesTaskId returns the RestartServicesTaskId field if non-nil, zero value otherwise.

### GetRestartServicesTaskIdOk

`func (o *RecoverCassandraParams) GetRestartServicesTaskIdOk() (*int64, bool)`

GetRestartServicesTaskIdOk returns a tuple with the RestartServicesTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartServicesTaskId

`func (o *RecoverCassandraParams) SetRestartServicesTaskId(v int64)`

SetRestartServicesTaskId sets RestartServicesTaskId field to given value.

### HasRestartServicesTaskId

`func (o *RecoverCassandraParams) HasRestartServicesTaskId() bool`

HasRestartServicesTaskId returns a boolean if a field has been set.

### SetRestartServicesTaskIdNil

`func (o *RecoverCassandraParams) SetRestartServicesTaskIdNil(b bool)`

 SetRestartServicesTaskIdNil sets the value for RestartServicesTaskId to be an explicit nil

### UnsetRestartServicesTaskId
`func (o *RecoverCassandraParams) UnsetRestartServicesTaskId()`

UnsetRestartServicesTaskId ensures that no value is present for RestartServicesTaskId, not even an explicit nil
### GetRunPreChecks

`func (o *RecoverCassandraParams) GetRunPreChecks() bool`

GetRunPreChecks returns the RunPreChecks field if non-nil, zero value otherwise.

### GetRunPreChecksOk

`func (o *RecoverCassandraParams) GetRunPreChecksOk() (*bool, bool)`

GetRunPreChecksOk returns a tuple with the RunPreChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPreChecks

`func (o *RecoverCassandraParams) SetRunPreChecks(v bool)`

SetRunPreChecks sets RunPreChecks field to given value.

### HasRunPreChecks

`func (o *RecoverCassandraParams) HasRunPreChecks() bool`

HasRunPreChecks returns a boolean if a field has been set.

### SetRunPreChecksNil

`func (o *RecoverCassandraParams) SetRunPreChecksNil(b bool)`

 SetRunPreChecksNil sets the value for RunPreChecks to be an explicit nil

### UnsetRunPreChecks
`func (o *RecoverCassandraParams) UnsetRunPreChecks()`

UnsetRunPreChecks ensures that no value is present for RunPreChecks, not even an explicit nil
### GetSelectedDataCenters

`func (o *RecoverCassandraParams) GetSelectedDataCenters() []string`

GetSelectedDataCenters returns the SelectedDataCenters field if non-nil, zero value otherwise.

### GetSelectedDataCentersOk

`func (o *RecoverCassandraParams) GetSelectedDataCentersOk() (*[]string, bool)`

GetSelectedDataCentersOk returns a tuple with the SelectedDataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDataCenters

`func (o *RecoverCassandraParams) SetSelectedDataCenters(v []string)`

SetSelectedDataCenters sets SelectedDataCenters field to given value.

### HasSelectedDataCenters

`func (o *RecoverCassandraParams) HasSelectedDataCenters() bool`

HasSelectedDataCenters returns a boolean if a field has been set.

### GetSnapshots

`func (o *RecoverCassandraParams) GetSnapshots() []RecoverCassandraSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverCassandraParams) GetSnapshotsOk() (*[]RecoverCassandraSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverCassandraParams) SetSnapshots(v []RecoverCassandraSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverCassandraParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverCassandraParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetStagingDirectoryList

`func (o *RecoverCassandraParams) GetStagingDirectoryList() []string`

GetStagingDirectoryList returns the StagingDirectoryList field if non-nil, zero value otherwise.

### GetStagingDirectoryListOk

`func (o *RecoverCassandraParams) GetStagingDirectoryListOk() (*[]string, bool)`

GetStagingDirectoryListOk returns a tuple with the StagingDirectoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDirectoryList

`func (o *RecoverCassandraParams) SetStagingDirectoryList(v []string)`

SetStagingDirectoryList sets StagingDirectoryList field to given value.

### HasStagingDirectoryList

`func (o *RecoverCassandraParams) HasStagingDirectoryList() bool`

HasStagingDirectoryList returns a boolean if a field has been set.

### GetSuffix

`func (o *RecoverCassandraParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverCassandraParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverCassandraParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverCassandraParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverCassandraParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverCassandraParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


