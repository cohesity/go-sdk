# BackupRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment type that the task is protecting. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**Error** | Pointer to **NullableString** | Specifies if an error occurred (if any) while running this task. This field is populated when the status is equal to &#39;kFailure&#39;. | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the Job Run that ran the backup task and the copy tasks. | [optional] 
**Message** | Pointer to **NullableString** | Specifies a message after finishing the task successfully. This field is optionally populated when the status is equal to &#39;kSuccess&#39;. | [optional] 
**MetadataDeleted** | Pointer to **NullableBool** | Specifies if the metadata and snapshots associated with this Job Run have been deleted. This field is set to true when the snapshots, which are marked for deletion, are removed by garbage collection. The associated metadata is also deleted. | [optional] 
**Quiesced** | Pointer to **NullableBool** | Specifies if app-consistent snapshot was captured. This field is set to true, if an app-consistent snapshot was taken by quiescing applications and the file system before taking a backup. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of backup such as &#39;kRegular&#39;, &#39;kFull&#39;, &#39;kLog&#39; or &#39;kSystem&#39;. &#39;kRegular&#39; indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time. | [optional] 
**SlaViolated** | Pointer to **NullableBool** | Specifies if the SLA was violated for the Job Run. This field is set to true, if time to complete the Job Run is longer than the SLA specified. This field is populated when the status is set to &#39;kSuccess&#39; or &#39;kFailure&#39;. | [optional] 
**SnapshotsDeleted** | Pointer to **NullableBool** | Specifies if backup snapshots associated with this Job Run have been marked for deletion because of the retention settings in the Policy or if they were manually deleted from the Cohesity Dashboard. | [optional] 
**SnapshotsDeletedTimeUsecs** | Pointer to **NullableInt64** | Specifies if backup snapshots associated with this Job Run have been marked for deletion because of the retention settings in the Policy or if they were manually deleted from the Cohesity Dashboard. | [optional] 
**SourceBackupStatus** | Pointer to [**[]SourceBackupStatus**](SourceBackupStatus.md) | Array of Source Object Backup Status.  Specifies the status of backing up each source objects (such as VMs) associated with the job. | [optional] 
**Stats** | Pointer to [**ProtectionJobRunStats**](ProtectionJobRunStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of Backup task such as &#39;kRunning&#39;, &#39;kSuccess&#39;, &#39;kFailure&#39; etc. &#39;kAccepted&#39; indicates the task is queued to run but not yet running. &#39;kRunning&#39; indicates the task is running. &#39;kCanceling&#39; indicates a request to cancel the task has occurred but the task is not yet canceled. &#39;kCanceled&#39; indicates the task has been canceled. &#39;kSuccess&#39; indicates the task was successful. &#39;kFailure&#39; indicates the task failed. &#39;kWarning&#39; indicates the task has finished with warning. &#39;kOnHold&#39; indicates the task is kept onHold. &#39;kMissed&#39; indicates the task is missed. | [optional] 
**Warnings** | Pointer to **[]string** | Array of Warnings.  Specifies the warnings that occurred (if any) while running this task. | [optional] 
**WormRetentionType** | Pointer to **NullableString** | Specifies WORM retention type for the snapshot as given by the policy. When a WORM retention type is specified, the snapshot will be kept until the maximum of the snapshot retention time. During that time, the snapshot cannot be deleted. &#39;kNone&#39; implies there is no WORM retention set. &#39;kCompliance&#39; implies WORM retention is set for compliance reason. &#39;kAdministrative&#39; implies WORM retention is set for administrative purposes. | [optional] 

## Methods

### NewBackupRun

`func NewBackupRun() *BackupRun`

NewBackupRun instantiates a new BackupRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunWithDefaults

`func NewBackupRunWithDefaults() *BackupRun`

NewBackupRunWithDefaults instantiates a new BackupRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *BackupRun) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BackupRun) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BackupRun) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BackupRun) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BackupRun) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BackupRun) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetError

`func (o *BackupRun) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackupRun) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackupRun) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BackupRun) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BackupRun) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BackupRun) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetJobRunId

`func (o *BackupRun) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *BackupRun) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *BackupRun) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *BackupRun) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *BackupRun) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *BackupRun) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetMessage

`func (o *BackupRun) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupRun) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupRun) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BackupRun) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BackupRun) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BackupRun) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMetadataDeleted

`func (o *BackupRun) GetMetadataDeleted() bool`

GetMetadataDeleted returns the MetadataDeleted field if non-nil, zero value otherwise.

### GetMetadataDeletedOk

`func (o *BackupRun) GetMetadataDeletedOk() (*bool, bool)`

GetMetadataDeletedOk returns a tuple with the MetadataDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataDeleted

`func (o *BackupRun) SetMetadataDeleted(v bool)`

SetMetadataDeleted sets MetadataDeleted field to given value.

### HasMetadataDeleted

`func (o *BackupRun) HasMetadataDeleted() bool`

HasMetadataDeleted returns a boolean if a field has been set.

### SetMetadataDeletedNil

`func (o *BackupRun) SetMetadataDeletedNil(b bool)`

 SetMetadataDeletedNil sets the value for MetadataDeleted to be an explicit nil

### UnsetMetadataDeleted
`func (o *BackupRun) UnsetMetadataDeleted()`

UnsetMetadataDeleted ensures that no value is present for MetadataDeleted, not even an explicit nil
### GetQuiesced

`func (o *BackupRun) GetQuiesced() bool`

GetQuiesced returns the Quiesced field if non-nil, zero value otherwise.

### GetQuiescedOk

`func (o *BackupRun) GetQuiescedOk() (*bool, bool)`

GetQuiescedOk returns a tuple with the Quiesced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesced

`func (o *BackupRun) SetQuiesced(v bool)`

SetQuiesced sets Quiesced field to given value.

### HasQuiesced

`func (o *BackupRun) HasQuiesced() bool`

HasQuiesced returns a boolean if a field has been set.

### SetQuiescedNil

`func (o *BackupRun) SetQuiescedNil(b bool)`

 SetQuiescedNil sets the value for Quiesced to be an explicit nil

### UnsetQuiesced
`func (o *BackupRun) UnsetQuiesced()`

UnsetQuiesced ensures that no value is present for Quiesced, not even an explicit nil
### GetRunType

`func (o *BackupRun) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *BackupRun) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *BackupRun) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *BackupRun) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *BackupRun) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *BackupRun) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetSlaViolated

`func (o *BackupRun) GetSlaViolated() bool`

GetSlaViolated returns the SlaViolated field if non-nil, zero value otherwise.

### GetSlaViolatedOk

`func (o *BackupRun) GetSlaViolatedOk() (*bool, bool)`

GetSlaViolatedOk returns a tuple with the SlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaViolated

`func (o *BackupRun) SetSlaViolated(v bool)`

SetSlaViolated sets SlaViolated field to given value.

### HasSlaViolated

`func (o *BackupRun) HasSlaViolated() bool`

HasSlaViolated returns a boolean if a field has been set.

### SetSlaViolatedNil

`func (o *BackupRun) SetSlaViolatedNil(b bool)`

 SetSlaViolatedNil sets the value for SlaViolated to be an explicit nil

### UnsetSlaViolated
`func (o *BackupRun) UnsetSlaViolated()`

UnsetSlaViolated ensures that no value is present for SlaViolated, not even an explicit nil
### GetSnapshotsDeleted

`func (o *BackupRun) GetSnapshotsDeleted() bool`

GetSnapshotsDeleted returns the SnapshotsDeleted field if non-nil, zero value otherwise.

### GetSnapshotsDeletedOk

`func (o *BackupRun) GetSnapshotsDeletedOk() (*bool, bool)`

GetSnapshotsDeletedOk returns a tuple with the SnapshotsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsDeleted

`func (o *BackupRun) SetSnapshotsDeleted(v bool)`

SetSnapshotsDeleted sets SnapshotsDeleted field to given value.

### HasSnapshotsDeleted

`func (o *BackupRun) HasSnapshotsDeleted() bool`

HasSnapshotsDeleted returns a boolean if a field has been set.

### SetSnapshotsDeletedNil

`func (o *BackupRun) SetSnapshotsDeletedNil(b bool)`

 SetSnapshotsDeletedNil sets the value for SnapshotsDeleted to be an explicit nil

### UnsetSnapshotsDeleted
`func (o *BackupRun) UnsetSnapshotsDeleted()`

UnsetSnapshotsDeleted ensures that no value is present for SnapshotsDeleted, not even an explicit nil
### GetSnapshotsDeletedTimeUsecs

`func (o *BackupRun) GetSnapshotsDeletedTimeUsecs() int64`

GetSnapshotsDeletedTimeUsecs returns the SnapshotsDeletedTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotsDeletedTimeUsecsOk

`func (o *BackupRun) GetSnapshotsDeletedTimeUsecsOk() (*int64, bool)`

GetSnapshotsDeletedTimeUsecsOk returns a tuple with the SnapshotsDeletedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsDeletedTimeUsecs

`func (o *BackupRun) SetSnapshotsDeletedTimeUsecs(v int64)`

SetSnapshotsDeletedTimeUsecs sets SnapshotsDeletedTimeUsecs field to given value.

### HasSnapshotsDeletedTimeUsecs

`func (o *BackupRun) HasSnapshotsDeletedTimeUsecs() bool`

HasSnapshotsDeletedTimeUsecs returns a boolean if a field has been set.

### SetSnapshotsDeletedTimeUsecsNil

`func (o *BackupRun) SetSnapshotsDeletedTimeUsecsNil(b bool)`

 SetSnapshotsDeletedTimeUsecsNil sets the value for SnapshotsDeletedTimeUsecs to be an explicit nil

### UnsetSnapshotsDeletedTimeUsecs
`func (o *BackupRun) UnsetSnapshotsDeletedTimeUsecs()`

UnsetSnapshotsDeletedTimeUsecs ensures that no value is present for SnapshotsDeletedTimeUsecs, not even an explicit nil
### GetSourceBackupStatus

`func (o *BackupRun) GetSourceBackupStatus() []SourceBackupStatus`

GetSourceBackupStatus returns the SourceBackupStatus field if non-nil, zero value otherwise.

### GetSourceBackupStatusOk

`func (o *BackupRun) GetSourceBackupStatusOk() (*[]SourceBackupStatus, bool)`

GetSourceBackupStatusOk returns a tuple with the SourceBackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupStatus

`func (o *BackupRun) SetSourceBackupStatus(v []SourceBackupStatus)`

SetSourceBackupStatus sets SourceBackupStatus field to given value.

### HasSourceBackupStatus

`func (o *BackupRun) HasSourceBackupStatus() bool`

HasSourceBackupStatus returns a boolean if a field has been set.

### SetSourceBackupStatusNil

`func (o *BackupRun) SetSourceBackupStatusNil(b bool)`

 SetSourceBackupStatusNil sets the value for SourceBackupStatus to be an explicit nil

### UnsetSourceBackupStatus
`func (o *BackupRun) UnsetSourceBackupStatus()`

UnsetSourceBackupStatus ensures that no value is present for SourceBackupStatus, not even an explicit nil
### GetStats

`func (o *BackupRun) GetStats() ProtectionJobRunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *BackupRun) GetStatsOk() (*ProtectionJobRunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *BackupRun) SetStats(v ProtectionJobRunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *BackupRun) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *BackupRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupRun) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupRun) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWarnings

`func (o *BackupRun) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *BackupRun) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *BackupRun) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *BackupRun) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *BackupRun) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *BackupRun) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetWormRetentionType

`func (o *BackupRun) GetWormRetentionType() string`

GetWormRetentionType returns the WormRetentionType field if non-nil, zero value otherwise.

### GetWormRetentionTypeOk

`func (o *BackupRun) GetWormRetentionTypeOk() (*string, bool)`

GetWormRetentionTypeOk returns a tuple with the WormRetentionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormRetentionType

`func (o *BackupRun) SetWormRetentionType(v string)`

SetWormRetentionType sets WormRetentionType field to given value.

### HasWormRetentionType

`func (o *BackupRun) HasWormRetentionType() bool`

HasWormRetentionType returns a boolean if a field has been set.

### SetWormRetentionTypeNil

`func (o *BackupRun) SetWormRetentionTypeNil(b bool)`

 SetWormRetentionTypeNil sets the value for WormRetentionType to be an explicit nil

### UnsetWormRetentionType
`func (o *BackupRun) UnsetWormRetentionType()`

UnsetWormRetentionType ensures that no value is present for WormRetentionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


