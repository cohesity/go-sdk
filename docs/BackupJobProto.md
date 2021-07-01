# BackupJobProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortInExclusionWindow** | Pointer to **NullableBool** | This field determines whether a backup run should be aborted when it hits an exclusion window (assuming that it was started earlier when it was not in an exclusion window). | [optional] 
**AlertingPolicy** | Pointer to [**AlertingPolicyProto**](AlertingPolicyProto.md) |  | [optional] 
**BackupQosPrincipal** | Pointer to **NullableInt32** | The backup QoS principal to use for the backup job. | [optional] 
**BackupSourceParams** | Pointer to [**[]BackupSourceParams**](BackupSourceParams.md) | This contains additional backup params that are applicable to sources that are captured as part of the backup job. NOTE: The sources could point to higher level entities (such as a \&quot;Cluster\&quot; in VMware environment), but the source params captured here will not be for the matching higher level entity, but instead be for leaf-level entities (such as VMs). | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Whether to continue backing up on quiesce failure. | [optional] 
**CreateRemoteView** | Pointer to **NullableBool** | If set to false, a remote view will not be created. If set to true and: 1) Remote view name is not provided by the user, a remote view is created with the same name as source view name. 2) Remote view name is provided by the user, a remote view is created with the given name. | [optional] 
**DedupDisabledSourceIdVec** | Pointer to **[]int64** | List of source ids for which source side dedup is disabled from the backup job. | [optional] 
**DeletionStatus** | Pointer to **NullableInt32** | Determines if the job (and associated backups) should be deleted. Once a job has been deleted, its status cannot be changed. | [optional] 
**Description** | Pointer to **NullableString** | Job description (as entered by the user). | [optional] 
**DrToCloudParams** | Pointer to [**BackupJobProtoDRToCloudParams**](BackupJobProtoDRToCloudParams.md) |  | [optional] 
**EhParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | The time (in usecs) after which no backup for the job should be scheduled. | [optional] 
**EnvBackupParams** | Pointer to [**EnvBackupParams**](EnvBackupParams.md) |  | [optional] 
**ExcludeSources** | Pointer to [**[]BackupJobProtoExcludeSource**](BackupJobProtoExcludeSource.md) | The list of sources to exclude from backups. These can have non-leaf-level entities, but it&#39;s up to the creator to ensure that a child of these sources hasn&#39;t been explicitly added to &#39;sources&#39;. | [optional] 
**ExcludeSourcesDEPRECATED** | Pointer to [**[]EntityProto**](EntityProto.md) | The list of sources to exclude from backups. These can have non-leaf-level entities, but it&#39;s up to the creator to ensure that a child of these sources hasn&#39;t been explicitly added to &#39;sources&#39;. | [optional] 
**ExclusionRanges** | Pointer to [**[]BackupJobProtoExclusionTimeRange**](BackupJobProtoExclusionTimeRange.md) | Do not run backups in these time-ranges. | [optional] 
**FullBackupJobPolicy** | Pointer to [**JobPolicyProto**](JobPolicyProto.md) |  | [optional] 
**FullBackupSlaTimeMins** | Pointer to **NullableInt64** | Same as &#39;sla_time_mins&#39; above, but applies to full backups. NOTE: This value is considered only for full backups that are excepted i.e either scheduled or the first full backup and not for full backups that happen as a result of incremental backup failure. | [optional] 
**GlobalIncludeExclude** | Pointer to [**PhysicalFileBackupParamsGlobalIncludeExclude**](PhysicalFileBackupParamsGlobalIncludeExclude.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicyProto**](IndexingPolicyProto.md) |  | [optional] 
**IsActive** | Pointer to **NullableBool** | Whether the backup job is active or not. Details about what an active job is can be found here: https://goo.gl/1mLvS3. | [optional] 
**IsContinuousSnapshottingEnabled** | Pointer to **NullableBool** | Indicates if Magneto should continue taking source snapshots even if there is a pending run. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Tracks whether the backup job has actually been deleted. | [optional] 
**IsDirectArchiveEnabled** | Pointer to **NullableBool** | This field is set to true if this is a direct archive backup job. | [optional] 
**IsDirectArchiveNativeFormatEnabled** | Pointer to **NullableBool** | This field is set to true if native format should be used for archiving. Applicable for only direct archive jobs. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Whether the backup job is paused. New backup runs are not scheduled for the paused backup job. Active run of a backup job (if any) is not impacted. | [optional] 
**IsRpoJob** | Pointer to **NullableBool** | Whether the backup job is an RPO policy job. These jobs are hidden from the user, and are created internally to have a backup schedule for the given source. | [optional] 
**JobCreationTimeUsecs** | Pointer to **NullableInt64** | Time when this job was first created. | [optional] 
**JobId** | Pointer to **NullableInt64** | A unique id for locally created jobs. This should only be used to identify jobs created on the local cluster. When Iris communicates with Magneto, Iris can continue to use this job_id field, which will always be assumed to refer to locally created jobs.  For remotely created jobs, the &#39;job_uid&#39; field should be used. The only time Iris should ever need to refer to a remote job is when restoring an object from a remote snapshot. In all such cases, Iris should use the job_uid field. | [optional] 
**JobPolicy** | Pointer to [**JobPolicyProto**](JobPolicyProto.md) |  | [optional] 
**JobUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**LastModificationTimeUsecs** | Pointer to **NullableInt64** | Time when this job description was last updated. | [optional] 
**LastPauseModificationTimeUsecs** | Pointer to **NullableInt64** | Time when the job was last paused or unpaused. | [optional] 
**LastPauseReason** | Pointer to **NullableInt32** | Last reason for pausing the backup job. Capturing the reason will help in resuming only the jobs that were paused because of a reason once the reason for pausing is not applicable. | [optional] 
**LastUpdatedUsername** | Pointer to **NullableString** | The user who modified the job most recently. | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | This is set to true by the user in order to backup the objects via a dedicated storage area network (SAN), as opposed to transport via LAN or management network. NOTE: Not all adapters support this method. Currently only VMware. | [optional] 
**LeverageStorageSnapshots** | Pointer to **NullableBool** | Whether to leverage the storage array based snapshots for this backup job. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, job will fallback to the default backup method. NOTE: This will be set for Pure snapshots. | [optional] 
**LeverageStorageSnapshotsForHyperflex** | Pointer to **NullableBool** | This is set to true by the user if hyperflex snapshots are requested NOTE: If this is set to true, then leverage_storage_snapshots above should be false. | [optional] 
**LogBackupJobPolicy** | Pointer to [**JobPolicyProto**](JobPolicyProto.md) |  | [optional] 
**MaxAllowedSourceSnapshots** | Pointer to **NullableInt32** | Determines the maximum number of source snapshots allowed for a given entity in the job. This is only applicable if &#39;is_continuous_snapshotting_enabled&#39; is set to true. | [optional] 
**Name** | Pointer to **NullableString** | The name of this backup job. This must be unique across all jobs. | [optional] 
**NumSnapshotsToKeepOnPrimary** | Pointer to **NullableInt64** | Specifies how many recent snapshots of each backed up entity to retain on the primary environment. If not specified, then snapshots will not be be deleted from the primary environment. NOTE: This is only applicable for certain environments like kPure. | [optional] 
**ParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PerformSourceSideDedup** | Pointer to **NullableBool** | Whether or not to perform source side dedup. | [optional] 
**PolicyAppliedTimeMsecs** | Pointer to **NullableInt64** | Epoch time in milliseconds when the policy was last applied to this job. This field will be used to determine whether a policy has changed after it was applied to a particular job. | [optional] 
**PolicyId** | Pointer to **NullableString** | Id of the policy being applied to the backup job. It is expected to be of the form \&quot;cluster_id:cluster_instance_id:local_identifier\&quot;. | [optional] 
**PolicyName** | Pointer to **NullableString** | The name of the policy referred to by policy_uid. This field can be stale and should not be relied upon for the latest name. | [optional] 
**PostBackupScript** | Pointer to [**BackupJobPreOrPostScript**](BackupJobPreOrPostScript.md) |  | [optional] 
**PreScript** | Pointer to [**BackupJobPreOrPostScript**](BackupJobPreOrPostScript.md) |  | [optional] 
**PrimaryJobUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**Priority** | Pointer to **NullableInt32** | The priority for the job. This is used at admission time - all admitted jobs are treated equally. This is also used to determine the Madrox replication priority. | [optional] 
**Quiesce** | Pointer to **NullableBool** | Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup. | [optional] 
**RemoteJobUids** | Pointer to [**[]UniversalIdProto**](UniversalIdProto.md) | The globally unique ids of all remote jobs that are linked to this job (because of incoming replications). This field will only be populated for locally created jobs. This field is populated only for the local(stub) job during incoming replications. In the most common case of one cluster replicating to another, this field will only have one entry (which is the id of the job on Tx side) and matches the primary_job_uid. This will have multiple entries in the following situation: A-&gt;B, A-&gt;C replication. The backup is failed over to B, and B now starts replicating to C. In this case, the stub job at C will have two entries. One is the job id from cluster A, and another is the local(stub) job uid from B. Also note that since the job originated from A, primary_job_uid for all the replicated instances of this job across multiple clusters will remain the same (which is equal to the job id from the original cluster A). | [optional] 
**RemoteViewName** | Pointer to **NullableString** | A human readable name of the remote view. A remote view is created with name overwriting the latest snapshot. | [optional] 
**RequiredFeatureVec** | Pointer to **[]string** | The features that are strictly required to be supported by the cluster of the backup job. This is currently used in the following cases: 1. Tx cluster looks at the Rx cluster&#39;s supported features and replicates the backup job only if all the features captured here are supported. 2. When performing remote restore of a backup job from an archival, this job will be retrieved only if the cluster supports all the features listed here. | [optional] 
**SlaTimeMins** | Pointer to **NullableInt64** | If specified, this variable determines the amount of time (after backup has started) in which backup is expected to finish for this job. An SLA violation is counted against this job if the amount of time taken exceeds this amount. | [optional] 
**SourceFilters** | Pointer to [**SourceFilters**](SourceFilters.md) |  | [optional] 
**Sources** | Pointer to [**[]BackupJobProtoBackupSource**](BackupJobProtoBackupSource.md) | The list of sources that should be backed up. A source in this list could be a descendant of another source in the list (this will be used when specifying override backup schedules). | [optional] 
**StartTime** | Pointer to [**Time**](Time.md) |  | [optional] 
**StubbingPolicy** | Pointer to [**StubbingPolicyProto**](StubbingPolicyProto.md) |  | [optional] 
**TagVec** | Pointer to **[]string** | Tags associated with the job. User can specify tags/keywords that can indexed by Yoda and can be later searched in UI. For example, user can create a &#39;kPuppeteer&#39; job to backup Oracle DB for &#39;payroll&#39; department. User can specify following tags: &#39;payroll&#39;, &#39;Oracle_DB&#39;. | [optional] 
**Timezone** | Pointer to **NullableString** | Timezone of the backup job. All time fields (i.e., TimeOfDay) in this backup job are stored wrt to this timezone.  The time zones have unique names of the form \&quot;Area/Location\&quot;, e.g. \&quot;America/New_York\&quot;. We are using \&quot;America/Los_Angeles\&quot; as a default value so as to be backward compatible with pre-2.7 code. | [optional] 
**TruncateLogs** | Pointer to **NullableBool** | Whether to truncate logs after a backup run. This is currently only relevant for full or incremental backups in a SQL environment. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this backup job corresponds to. | [optional] 
**UserInfo** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box to which data will be written. | [optional] 

## Methods

### NewBackupJobProto

`func NewBackupJobProto() *BackupJobProto`

NewBackupJobProto instantiates a new BackupJobProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobProtoWithDefaults

`func NewBackupJobProtoWithDefaults() *BackupJobProto`

NewBackupJobProtoWithDefaults instantiates a new BackupJobProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInExclusionWindow

`func (o *BackupJobProto) GetAbortInExclusionWindow() bool`

GetAbortInExclusionWindow returns the AbortInExclusionWindow field if non-nil, zero value otherwise.

### GetAbortInExclusionWindowOk

`func (o *BackupJobProto) GetAbortInExclusionWindowOk() (*bool, bool)`

GetAbortInExclusionWindowOk returns a tuple with the AbortInExclusionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInExclusionWindow

`func (o *BackupJobProto) SetAbortInExclusionWindow(v bool)`

SetAbortInExclusionWindow sets AbortInExclusionWindow field to given value.

### HasAbortInExclusionWindow

`func (o *BackupJobProto) HasAbortInExclusionWindow() bool`

HasAbortInExclusionWindow returns a boolean if a field has been set.

### SetAbortInExclusionWindowNil

`func (o *BackupJobProto) SetAbortInExclusionWindowNil(b bool)`

 SetAbortInExclusionWindowNil sets the value for AbortInExclusionWindow to be an explicit nil

### UnsetAbortInExclusionWindow
`func (o *BackupJobProto) UnsetAbortInExclusionWindow()`

UnsetAbortInExclusionWindow ensures that no value is present for AbortInExclusionWindow, not even an explicit nil
### GetAlertingPolicy

`func (o *BackupJobProto) GetAlertingPolicy() AlertingPolicyProto`

GetAlertingPolicy returns the AlertingPolicy field if non-nil, zero value otherwise.

### GetAlertingPolicyOk

`func (o *BackupJobProto) GetAlertingPolicyOk() (*AlertingPolicyProto, bool)`

GetAlertingPolicyOk returns a tuple with the AlertingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingPolicy

`func (o *BackupJobProto) SetAlertingPolicy(v AlertingPolicyProto)`

SetAlertingPolicy sets AlertingPolicy field to given value.

### HasAlertingPolicy

`func (o *BackupJobProto) HasAlertingPolicy() bool`

HasAlertingPolicy returns a boolean if a field has been set.

### GetBackupQosPrincipal

`func (o *BackupJobProto) GetBackupQosPrincipal() int32`

GetBackupQosPrincipal returns the BackupQosPrincipal field if non-nil, zero value otherwise.

### GetBackupQosPrincipalOk

`func (o *BackupJobProto) GetBackupQosPrincipalOk() (*int32, bool)`

GetBackupQosPrincipalOk returns a tuple with the BackupQosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupQosPrincipal

`func (o *BackupJobProto) SetBackupQosPrincipal(v int32)`

SetBackupQosPrincipal sets BackupQosPrincipal field to given value.

### HasBackupQosPrincipal

`func (o *BackupJobProto) HasBackupQosPrincipal() bool`

HasBackupQosPrincipal returns a boolean if a field has been set.

### SetBackupQosPrincipalNil

`func (o *BackupJobProto) SetBackupQosPrincipalNil(b bool)`

 SetBackupQosPrincipalNil sets the value for BackupQosPrincipal to be an explicit nil

### UnsetBackupQosPrincipal
`func (o *BackupJobProto) UnsetBackupQosPrincipal()`

UnsetBackupQosPrincipal ensures that no value is present for BackupQosPrincipal, not even an explicit nil
### GetBackupSourceParams

`func (o *BackupJobProto) GetBackupSourceParams() []BackupSourceParams`

GetBackupSourceParams returns the BackupSourceParams field if non-nil, zero value otherwise.

### GetBackupSourceParamsOk

`func (o *BackupJobProto) GetBackupSourceParamsOk() (*[]BackupSourceParams, bool)`

GetBackupSourceParamsOk returns a tuple with the BackupSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSourceParams

`func (o *BackupJobProto) SetBackupSourceParams(v []BackupSourceParams)`

SetBackupSourceParams sets BackupSourceParams field to given value.

### HasBackupSourceParams

`func (o *BackupJobProto) HasBackupSourceParams() bool`

HasBackupSourceParams returns a boolean if a field has been set.

### SetBackupSourceParamsNil

`func (o *BackupJobProto) SetBackupSourceParamsNil(b bool)`

 SetBackupSourceParamsNil sets the value for BackupSourceParams to be an explicit nil

### UnsetBackupSourceParams
`func (o *BackupJobProto) UnsetBackupSourceParams()`

UnsetBackupSourceParams ensures that no value is present for BackupSourceParams, not even an explicit nil
### GetContinueOnQuiesceFailure

`func (o *BackupJobProto) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *BackupJobProto) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *BackupJobProto) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *BackupJobProto) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *BackupJobProto) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *BackupJobProto) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
### GetCreateRemoteView

`func (o *BackupJobProto) GetCreateRemoteView() bool`

GetCreateRemoteView returns the CreateRemoteView field if non-nil, zero value otherwise.

### GetCreateRemoteViewOk

`func (o *BackupJobProto) GetCreateRemoteViewOk() (*bool, bool)`

GetCreateRemoteViewOk returns a tuple with the CreateRemoteView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRemoteView

`func (o *BackupJobProto) SetCreateRemoteView(v bool)`

SetCreateRemoteView sets CreateRemoteView field to given value.

### HasCreateRemoteView

`func (o *BackupJobProto) HasCreateRemoteView() bool`

HasCreateRemoteView returns a boolean if a field has been set.

### SetCreateRemoteViewNil

`func (o *BackupJobProto) SetCreateRemoteViewNil(b bool)`

 SetCreateRemoteViewNil sets the value for CreateRemoteView to be an explicit nil

### UnsetCreateRemoteView
`func (o *BackupJobProto) UnsetCreateRemoteView()`

UnsetCreateRemoteView ensures that no value is present for CreateRemoteView, not even an explicit nil
### GetDedupDisabledSourceIdVec

`func (o *BackupJobProto) GetDedupDisabledSourceIdVec() []int64`

GetDedupDisabledSourceIdVec returns the DedupDisabledSourceIdVec field if non-nil, zero value otherwise.

### GetDedupDisabledSourceIdVecOk

`func (o *BackupJobProto) GetDedupDisabledSourceIdVecOk() (*[]int64, bool)`

GetDedupDisabledSourceIdVecOk returns a tuple with the DedupDisabledSourceIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupDisabledSourceIdVec

`func (o *BackupJobProto) SetDedupDisabledSourceIdVec(v []int64)`

SetDedupDisabledSourceIdVec sets DedupDisabledSourceIdVec field to given value.

### HasDedupDisabledSourceIdVec

`func (o *BackupJobProto) HasDedupDisabledSourceIdVec() bool`

HasDedupDisabledSourceIdVec returns a boolean if a field has been set.

### SetDedupDisabledSourceIdVecNil

`func (o *BackupJobProto) SetDedupDisabledSourceIdVecNil(b bool)`

 SetDedupDisabledSourceIdVecNil sets the value for DedupDisabledSourceIdVec to be an explicit nil

### UnsetDedupDisabledSourceIdVec
`func (o *BackupJobProto) UnsetDedupDisabledSourceIdVec()`

UnsetDedupDisabledSourceIdVec ensures that no value is present for DedupDisabledSourceIdVec, not even an explicit nil
### GetDeletionStatus

`func (o *BackupJobProto) GetDeletionStatus() int32`

GetDeletionStatus returns the DeletionStatus field if non-nil, zero value otherwise.

### GetDeletionStatusOk

`func (o *BackupJobProto) GetDeletionStatusOk() (*int32, bool)`

GetDeletionStatusOk returns a tuple with the DeletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionStatus

`func (o *BackupJobProto) SetDeletionStatus(v int32)`

SetDeletionStatus sets DeletionStatus field to given value.

### HasDeletionStatus

`func (o *BackupJobProto) HasDeletionStatus() bool`

HasDeletionStatus returns a boolean if a field has been set.

### SetDeletionStatusNil

`func (o *BackupJobProto) SetDeletionStatusNil(b bool)`

 SetDeletionStatusNil sets the value for DeletionStatus to be an explicit nil

### UnsetDeletionStatus
`func (o *BackupJobProto) UnsetDeletionStatus()`

UnsetDeletionStatus ensures that no value is present for DeletionStatus, not even an explicit nil
### GetDescription

`func (o *BackupJobProto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupJobProto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupJobProto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupJobProto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupJobProto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupJobProto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDrToCloudParams

`func (o *BackupJobProto) GetDrToCloudParams() BackupJobProtoDRToCloudParams`

GetDrToCloudParams returns the DrToCloudParams field if non-nil, zero value otherwise.

### GetDrToCloudParamsOk

`func (o *BackupJobProto) GetDrToCloudParamsOk() (*BackupJobProtoDRToCloudParams, bool)`

GetDrToCloudParamsOk returns a tuple with the DrToCloudParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrToCloudParams

`func (o *BackupJobProto) SetDrToCloudParams(v BackupJobProtoDRToCloudParams)`

SetDrToCloudParams sets DrToCloudParams field to given value.

### HasDrToCloudParams

`func (o *BackupJobProto) HasDrToCloudParams() bool`

HasDrToCloudParams returns a boolean if a field has been set.

### GetEhParentSource

`func (o *BackupJobProto) GetEhParentSource() EntityProto`

GetEhParentSource returns the EhParentSource field if non-nil, zero value otherwise.

### GetEhParentSourceOk

`func (o *BackupJobProto) GetEhParentSourceOk() (*EntityProto, bool)`

GetEhParentSourceOk returns a tuple with the EhParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEhParentSource

`func (o *BackupJobProto) SetEhParentSource(v EntityProto)`

SetEhParentSource sets EhParentSource field to given value.

### HasEhParentSource

`func (o *BackupJobProto) HasEhParentSource() bool`

HasEhParentSource returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *BackupJobProto) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupJobProto) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupJobProto) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupJobProto) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupJobProto) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupJobProto) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvBackupParams

`func (o *BackupJobProto) GetEnvBackupParams() EnvBackupParams`

GetEnvBackupParams returns the EnvBackupParams field if non-nil, zero value otherwise.

### GetEnvBackupParamsOk

`func (o *BackupJobProto) GetEnvBackupParamsOk() (*EnvBackupParams, bool)`

GetEnvBackupParamsOk returns a tuple with the EnvBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvBackupParams

`func (o *BackupJobProto) SetEnvBackupParams(v EnvBackupParams)`

SetEnvBackupParams sets EnvBackupParams field to given value.

### HasEnvBackupParams

`func (o *BackupJobProto) HasEnvBackupParams() bool`

HasEnvBackupParams returns a boolean if a field has been set.

### GetExcludeSources

`func (o *BackupJobProto) GetExcludeSources() []BackupJobProtoExcludeSource`

GetExcludeSources returns the ExcludeSources field if non-nil, zero value otherwise.

### GetExcludeSourcesOk

`func (o *BackupJobProto) GetExcludeSourcesOk() (*[]BackupJobProtoExcludeSource, bool)`

GetExcludeSourcesOk returns a tuple with the ExcludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSources

`func (o *BackupJobProto) SetExcludeSources(v []BackupJobProtoExcludeSource)`

SetExcludeSources sets ExcludeSources field to given value.

### HasExcludeSources

`func (o *BackupJobProto) HasExcludeSources() bool`

HasExcludeSources returns a boolean if a field has been set.

### SetExcludeSourcesNil

`func (o *BackupJobProto) SetExcludeSourcesNil(b bool)`

 SetExcludeSourcesNil sets the value for ExcludeSources to be an explicit nil

### UnsetExcludeSources
`func (o *BackupJobProto) UnsetExcludeSources()`

UnsetExcludeSources ensures that no value is present for ExcludeSources, not even an explicit nil
### GetExcludeSourcesDEPRECATED

`func (o *BackupJobProto) GetExcludeSourcesDEPRECATED() []EntityProto`

GetExcludeSourcesDEPRECATED returns the ExcludeSourcesDEPRECATED field if non-nil, zero value otherwise.

### GetExcludeSourcesDEPRECATEDOk

`func (o *BackupJobProto) GetExcludeSourcesDEPRECATEDOk() (*[]EntityProto, bool)`

GetExcludeSourcesDEPRECATEDOk returns a tuple with the ExcludeSourcesDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSourcesDEPRECATED

`func (o *BackupJobProto) SetExcludeSourcesDEPRECATED(v []EntityProto)`

SetExcludeSourcesDEPRECATED sets ExcludeSourcesDEPRECATED field to given value.

### HasExcludeSourcesDEPRECATED

`func (o *BackupJobProto) HasExcludeSourcesDEPRECATED() bool`

HasExcludeSourcesDEPRECATED returns a boolean if a field has been set.

### SetExcludeSourcesDEPRECATEDNil

`func (o *BackupJobProto) SetExcludeSourcesDEPRECATEDNil(b bool)`

 SetExcludeSourcesDEPRECATEDNil sets the value for ExcludeSourcesDEPRECATED to be an explicit nil

### UnsetExcludeSourcesDEPRECATED
`func (o *BackupJobProto) UnsetExcludeSourcesDEPRECATED()`

UnsetExcludeSourcesDEPRECATED ensures that no value is present for ExcludeSourcesDEPRECATED, not even an explicit nil
### GetExclusionRanges

`func (o *BackupJobProto) GetExclusionRanges() []BackupJobProtoExclusionTimeRange`

GetExclusionRanges returns the ExclusionRanges field if non-nil, zero value otherwise.

### GetExclusionRangesOk

`func (o *BackupJobProto) GetExclusionRangesOk() (*[]BackupJobProtoExclusionTimeRange, bool)`

GetExclusionRangesOk returns a tuple with the ExclusionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRanges

`func (o *BackupJobProto) SetExclusionRanges(v []BackupJobProtoExclusionTimeRange)`

SetExclusionRanges sets ExclusionRanges field to given value.

### HasExclusionRanges

`func (o *BackupJobProto) HasExclusionRanges() bool`

HasExclusionRanges returns a boolean if a field has been set.

### SetExclusionRangesNil

`func (o *BackupJobProto) SetExclusionRangesNil(b bool)`

 SetExclusionRangesNil sets the value for ExclusionRanges to be an explicit nil

### UnsetExclusionRanges
`func (o *BackupJobProto) UnsetExclusionRanges()`

UnsetExclusionRanges ensures that no value is present for ExclusionRanges, not even an explicit nil
### GetFullBackupJobPolicy

`func (o *BackupJobProto) GetFullBackupJobPolicy() JobPolicyProto`

GetFullBackupJobPolicy returns the FullBackupJobPolicy field if non-nil, zero value otherwise.

### GetFullBackupJobPolicyOk

`func (o *BackupJobProto) GetFullBackupJobPolicyOk() (*JobPolicyProto, bool)`

GetFullBackupJobPolicyOk returns a tuple with the FullBackupJobPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupJobPolicy

`func (o *BackupJobProto) SetFullBackupJobPolicy(v JobPolicyProto)`

SetFullBackupJobPolicy sets FullBackupJobPolicy field to given value.

### HasFullBackupJobPolicy

`func (o *BackupJobProto) HasFullBackupJobPolicy() bool`

HasFullBackupJobPolicy returns a boolean if a field has been set.

### GetFullBackupSlaTimeMins

`func (o *BackupJobProto) GetFullBackupSlaTimeMins() int64`

GetFullBackupSlaTimeMins returns the FullBackupSlaTimeMins field if non-nil, zero value otherwise.

### GetFullBackupSlaTimeMinsOk

`func (o *BackupJobProto) GetFullBackupSlaTimeMinsOk() (*int64, bool)`

GetFullBackupSlaTimeMinsOk returns a tuple with the FullBackupSlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupSlaTimeMins

`func (o *BackupJobProto) SetFullBackupSlaTimeMins(v int64)`

SetFullBackupSlaTimeMins sets FullBackupSlaTimeMins field to given value.

### HasFullBackupSlaTimeMins

`func (o *BackupJobProto) HasFullBackupSlaTimeMins() bool`

HasFullBackupSlaTimeMins returns a boolean if a field has been set.

### SetFullBackupSlaTimeMinsNil

`func (o *BackupJobProto) SetFullBackupSlaTimeMinsNil(b bool)`

 SetFullBackupSlaTimeMinsNil sets the value for FullBackupSlaTimeMins to be an explicit nil

### UnsetFullBackupSlaTimeMins
`func (o *BackupJobProto) UnsetFullBackupSlaTimeMins()`

UnsetFullBackupSlaTimeMins ensures that no value is present for FullBackupSlaTimeMins, not even an explicit nil
### GetGlobalIncludeExclude

`func (o *BackupJobProto) GetGlobalIncludeExclude() PhysicalFileBackupParamsGlobalIncludeExclude`

GetGlobalIncludeExclude returns the GlobalIncludeExclude field if non-nil, zero value otherwise.

### GetGlobalIncludeExcludeOk

`func (o *BackupJobProto) GetGlobalIncludeExcludeOk() (*PhysicalFileBackupParamsGlobalIncludeExclude, bool)`

GetGlobalIncludeExcludeOk returns a tuple with the GlobalIncludeExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIncludeExclude

`func (o *BackupJobProto) SetGlobalIncludeExclude(v PhysicalFileBackupParamsGlobalIncludeExclude)`

SetGlobalIncludeExclude sets GlobalIncludeExclude field to given value.

### HasGlobalIncludeExclude

`func (o *BackupJobProto) HasGlobalIncludeExclude() bool`

HasGlobalIncludeExclude returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *BackupJobProto) GetIndexingPolicy() IndexingPolicyProto`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *BackupJobProto) GetIndexingPolicyOk() (*IndexingPolicyProto, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *BackupJobProto) SetIndexingPolicy(v IndexingPolicyProto)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *BackupJobProto) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetIsActive

`func (o *BackupJobProto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BackupJobProto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BackupJobProto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BackupJobProto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *BackupJobProto) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *BackupJobProto) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsContinuousSnapshottingEnabled

`func (o *BackupJobProto) GetIsContinuousSnapshottingEnabled() bool`

GetIsContinuousSnapshottingEnabled returns the IsContinuousSnapshottingEnabled field if non-nil, zero value otherwise.

### GetIsContinuousSnapshottingEnabledOk

`func (o *BackupJobProto) GetIsContinuousSnapshottingEnabledOk() (*bool, bool)`

GetIsContinuousSnapshottingEnabledOk returns a tuple with the IsContinuousSnapshottingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContinuousSnapshottingEnabled

`func (o *BackupJobProto) SetIsContinuousSnapshottingEnabled(v bool)`

SetIsContinuousSnapshottingEnabled sets IsContinuousSnapshottingEnabled field to given value.

### HasIsContinuousSnapshottingEnabled

`func (o *BackupJobProto) HasIsContinuousSnapshottingEnabled() bool`

HasIsContinuousSnapshottingEnabled returns a boolean if a field has been set.

### SetIsContinuousSnapshottingEnabledNil

`func (o *BackupJobProto) SetIsContinuousSnapshottingEnabledNil(b bool)`

 SetIsContinuousSnapshottingEnabledNil sets the value for IsContinuousSnapshottingEnabled to be an explicit nil

### UnsetIsContinuousSnapshottingEnabled
`func (o *BackupJobProto) UnsetIsContinuousSnapshottingEnabled()`

UnsetIsContinuousSnapshottingEnabled ensures that no value is present for IsContinuousSnapshottingEnabled, not even an explicit nil
### GetIsDeleted

`func (o *BackupJobProto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BackupJobProto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BackupJobProto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BackupJobProto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *BackupJobProto) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *BackupJobProto) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetIsDirectArchiveEnabled

`func (o *BackupJobProto) GetIsDirectArchiveEnabled() bool`

GetIsDirectArchiveEnabled returns the IsDirectArchiveEnabled field if non-nil, zero value otherwise.

### GetIsDirectArchiveEnabledOk

`func (o *BackupJobProto) GetIsDirectArchiveEnabledOk() (*bool, bool)`

GetIsDirectArchiveEnabledOk returns a tuple with the IsDirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectArchiveEnabled

`func (o *BackupJobProto) SetIsDirectArchiveEnabled(v bool)`

SetIsDirectArchiveEnabled sets IsDirectArchiveEnabled field to given value.

### HasIsDirectArchiveEnabled

`func (o *BackupJobProto) HasIsDirectArchiveEnabled() bool`

HasIsDirectArchiveEnabled returns a boolean if a field has been set.

### SetIsDirectArchiveEnabledNil

`func (o *BackupJobProto) SetIsDirectArchiveEnabledNil(b bool)`

 SetIsDirectArchiveEnabledNil sets the value for IsDirectArchiveEnabled to be an explicit nil

### UnsetIsDirectArchiveEnabled
`func (o *BackupJobProto) UnsetIsDirectArchiveEnabled()`

UnsetIsDirectArchiveEnabled ensures that no value is present for IsDirectArchiveEnabled, not even an explicit nil
### GetIsDirectArchiveNativeFormatEnabled

`func (o *BackupJobProto) GetIsDirectArchiveNativeFormatEnabled() bool`

GetIsDirectArchiveNativeFormatEnabled returns the IsDirectArchiveNativeFormatEnabled field if non-nil, zero value otherwise.

### GetIsDirectArchiveNativeFormatEnabledOk

`func (o *BackupJobProto) GetIsDirectArchiveNativeFormatEnabledOk() (*bool, bool)`

GetIsDirectArchiveNativeFormatEnabledOk returns a tuple with the IsDirectArchiveNativeFormatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectArchiveNativeFormatEnabled

`func (o *BackupJobProto) SetIsDirectArchiveNativeFormatEnabled(v bool)`

SetIsDirectArchiveNativeFormatEnabled sets IsDirectArchiveNativeFormatEnabled field to given value.

### HasIsDirectArchiveNativeFormatEnabled

`func (o *BackupJobProto) HasIsDirectArchiveNativeFormatEnabled() bool`

HasIsDirectArchiveNativeFormatEnabled returns a boolean if a field has been set.

### SetIsDirectArchiveNativeFormatEnabledNil

`func (o *BackupJobProto) SetIsDirectArchiveNativeFormatEnabledNil(b bool)`

 SetIsDirectArchiveNativeFormatEnabledNil sets the value for IsDirectArchiveNativeFormatEnabled to be an explicit nil

### UnsetIsDirectArchiveNativeFormatEnabled
`func (o *BackupJobProto) UnsetIsDirectArchiveNativeFormatEnabled()`

UnsetIsDirectArchiveNativeFormatEnabled ensures that no value is present for IsDirectArchiveNativeFormatEnabled, not even an explicit nil
### GetIsPaused

`func (o *BackupJobProto) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *BackupJobProto) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *BackupJobProto) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *BackupJobProto) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *BackupJobProto) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *BackupJobProto) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsRpoJob

`func (o *BackupJobProto) GetIsRpoJob() bool`

GetIsRpoJob returns the IsRpoJob field if non-nil, zero value otherwise.

### GetIsRpoJobOk

`func (o *BackupJobProto) GetIsRpoJobOk() (*bool, bool)`

GetIsRpoJobOk returns a tuple with the IsRpoJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRpoJob

`func (o *BackupJobProto) SetIsRpoJob(v bool)`

SetIsRpoJob sets IsRpoJob field to given value.

### HasIsRpoJob

`func (o *BackupJobProto) HasIsRpoJob() bool`

HasIsRpoJob returns a boolean if a field has been set.

### SetIsRpoJobNil

`func (o *BackupJobProto) SetIsRpoJobNil(b bool)`

 SetIsRpoJobNil sets the value for IsRpoJob to be an explicit nil

### UnsetIsRpoJob
`func (o *BackupJobProto) UnsetIsRpoJob()`

UnsetIsRpoJob ensures that no value is present for IsRpoJob, not even an explicit nil
### GetJobCreationTimeUsecs

`func (o *BackupJobProto) GetJobCreationTimeUsecs() int64`

GetJobCreationTimeUsecs returns the JobCreationTimeUsecs field if non-nil, zero value otherwise.

### GetJobCreationTimeUsecsOk

`func (o *BackupJobProto) GetJobCreationTimeUsecsOk() (*int64, bool)`

GetJobCreationTimeUsecsOk returns a tuple with the JobCreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCreationTimeUsecs

`func (o *BackupJobProto) SetJobCreationTimeUsecs(v int64)`

SetJobCreationTimeUsecs sets JobCreationTimeUsecs field to given value.

### HasJobCreationTimeUsecs

`func (o *BackupJobProto) HasJobCreationTimeUsecs() bool`

HasJobCreationTimeUsecs returns a boolean if a field has been set.

### SetJobCreationTimeUsecsNil

`func (o *BackupJobProto) SetJobCreationTimeUsecsNil(b bool)`

 SetJobCreationTimeUsecsNil sets the value for JobCreationTimeUsecs to be an explicit nil

### UnsetJobCreationTimeUsecs
`func (o *BackupJobProto) UnsetJobCreationTimeUsecs()`

UnsetJobCreationTimeUsecs ensures that no value is present for JobCreationTimeUsecs, not even an explicit nil
### GetJobId

`func (o *BackupJobProto) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupJobProto) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupJobProto) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupJobProto) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *BackupJobProto) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *BackupJobProto) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobPolicy

`func (o *BackupJobProto) GetJobPolicy() JobPolicyProto`

GetJobPolicy returns the JobPolicy field if non-nil, zero value otherwise.

### GetJobPolicyOk

`func (o *BackupJobProto) GetJobPolicyOk() (*JobPolicyProto, bool)`

GetJobPolicyOk returns a tuple with the JobPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPolicy

`func (o *BackupJobProto) SetJobPolicy(v JobPolicyProto)`

SetJobPolicy sets JobPolicy field to given value.

### HasJobPolicy

`func (o *BackupJobProto) HasJobPolicy() bool`

HasJobPolicy returns a boolean if a field has been set.

### GetJobUid

`func (o *BackupJobProto) GetJobUid() UniversalIdProto`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *BackupJobProto) GetJobUidOk() (*UniversalIdProto, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *BackupJobProto) SetJobUid(v UniversalIdProto)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *BackupJobProto) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetLastModificationTimeUsecs

`func (o *BackupJobProto) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *BackupJobProto) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *BackupJobProto) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *BackupJobProto) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### SetLastModificationTimeUsecsNil

`func (o *BackupJobProto) SetLastModificationTimeUsecsNil(b bool)`

 SetLastModificationTimeUsecsNil sets the value for LastModificationTimeUsecs to be an explicit nil

### UnsetLastModificationTimeUsecs
`func (o *BackupJobProto) UnsetLastModificationTimeUsecs()`

UnsetLastModificationTimeUsecs ensures that no value is present for LastModificationTimeUsecs, not even an explicit nil
### GetLastPauseModificationTimeUsecs

`func (o *BackupJobProto) GetLastPauseModificationTimeUsecs() int64`

GetLastPauseModificationTimeUsecs returns the LastPauseModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastPauseModificationTimeUsecsOk

`func (o *BackupJobProto) GetLastPauseModificationTimeUsecsOk() (*int64, bool)`

GetLastPauseModificationTimeUsecsOk returns a tuple with the LastPauseModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPauseModificationTimeUsecs

`func (o *BackupJobProto) SetLastPauseModificationTimeUsecs(v int64)`

SetLastPauseModificationTimeUsecs sets LastPauseModificationTimeUsecs field to given value.

### HasLastPauseModificationTimeUsecs

`func (o *BackupJobProto) HasLastPauseModificationTimeUsecs() bool`

HasLastPauseModificationTimeUsecs returns a boolean if a field has been set.

### SetLastPauseModificationTimeUsecsNil

`func (o *BackupJobProto) SetLastPauseModificationTimeUsecsNil(b bool)`

 SetLastPauseModificationTimeUsecsNil sets the value for LastPauseModificationTimeUsecs to be an explicit nil

### UnsetLastPauseModificationTimeUsecs
`func (o *BackupJobProto) UnsetLastPauseModificationTimeUsecs()`

UnsetLastPauseModificationTimeUsecs ensures that no value is present for LastPauseModificationTimeUsecs, not even an explicit nil
### GetLastPauseReason

`func (o *BackupJobProto) GetLastPauseReason() int32`

GetLastPauseReason returns the LastPauseReason field if non-nil, zero value otherwise.

### GetLastPauseReasonOk

`func (o *BackupJobProto) GetLastPauseReasonOk() (*int32, bool)`

GetLastPauseReasonOk returns a tuple with the LastPauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPauseReason

`func (o *BackupJobProto) SetLastPauseReason(v int32)`

SetLastPauseReason sets LastPauseReason field to given value.

### HasLastPauseReason

`func (o *BackupJobProto) HasLastPauseReason() bool`

HasLastPauseReason returns a boolean if a field has been set.

### SetLastPauseReasonNil

`func (o *BackupJobProto) SetLastPauseReasonNil(b bool)`

 SetLastPauseReasonNil sets the value for LastPauseReason to be an explicit nil

### UnsetLastPauseReason
`func (o *BackupJobProto) UnsetLastPauseReason()`

UnsetLastPauseReason ensures that no value is present for LastPauseReason, not even an explicit nil
### GetLastUpdatedUsername

`func (o *BackupJobProto) GetLastUpdatedUsername() string`

GetLastUpdatedUsername returns the LastUpdatedUsername field if non-nil, zero value otherwise.

### GetLastUpdatedUsernameOk

`func (o *BackupJobProto) GetLastUpdatedUsernameOk() (*string, bool)`

GetLastUpdatedUsernameOk returns a tuple with the LastUpdatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedUsername

`func (o *BackupJobProto) SetLastUpdatedUsername(v string)`

SetLastUpdatedUsername sets LastUpdatedUsername field to given value.

### HasLastUpdatedUsername

`func (o *BackupJobProto) HasLastUpdatedUsername() bool`

HasLastUpdatedUsername returns a boolean if a field has been set.

### SetLastUpdatedUsernameNil

`func (o *BackupJobProto) SetLastUpdatedUsernameNil(b bool)`

 SetLastUpdatedUsernameNil sets the value for LastUpdatedUsername to be an explicit nil

### UnsetLastUpdatedUsername
`func (o *BackupJobProto) UnsetLastUpdatedUsername()`

UnsetLastUpdatedUsername ensures that no value is present for LastUpdatedUsername, not even an explicit nil
### GetLeverageSanTransport

`func (o *BackupJobProto) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *BackupJobProto) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *BackupJobProto) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *BackupJobProto) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *BackupJobProto) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *BackupJobProto) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetLeverageStorageSnapshots

`func (o *BackupJobProto) GetLeverageStorageSnapshots() bool`

GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsOk

`func (o *BackupJobProto) GetLeverageStorageSnapshotsOk() (*bool, bool)`

GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshots

`func (o *BackupJobProto) SetLeverageStorageSnapshots(v bool)`

SetLeverageStorageSnapshots sets LeverageStorageSnapshots field to given value.

### HasLeverageStorageSnapshots

`func (o *BackupJobProto) HasLeverageStorageSnapshots() bool`

HasLeverageStorageSnapshots returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsNil

`func (o *BackupJobProto) SetLeverageStorageSnapshotsNil(b bool)`

 SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil

### UnsetLeverageStorageSnapshots
`func (o *BackupJobProto) UnsetLeverageStorageSnapshots()`

UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
### GetLeverageStorageSnapshotsForHyperflex

`func (o *BackupJobProto) GetLeverageStorageSnapshotsForHyperflex() bool`

GetLeverageStorageSnapshotsForHyperflex returns the LeverageStorageSnapshotsForHyperflex field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsForHyperflexOk

`func (o *BackupJobProto) GetLeverageStorageSnapshotsForHyperflexOk() (*bool, bool)`

GetLeverageStorageSnapshotsForHyperflexOk returns a tuple with the LeverageStorageSnapshotsForHyperflex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshotsForHyperflex

`func (o *BackupJobProto) SetLeverageStorageSnapshotsForHyperflex(v bool)`

SetLeverageStorageSnapshotsForHyperflex sets LeverageStorageSnapshotsForHyperflex field to given value.

### HasLeverageStorageSnapshotsForHyperflex

`func (o *BackupJobProto) HasLeverageStorageSnapshotsForHyperflex() bool`

HasLeverageStorageSnapshotsForHyperflex returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsForHyperflexNil

`func (o *BackupJobProto) SetLeverageStorageSnapshotsForHyperflexNil(b bool)`

 SetLeverageStorageSnapshotsForHyperflexNil sets the value for LeverageStorageSnapshotsForHyperflex to be an explicit nil

### UnsetLeverageStorageSnapshotsForHyperflex
`func (o *BackupJobProto) UnsetLeverageStorageSnapshotsForHyperflex()`

UnsetLeverageStorageSnapshotsForHyperflex ensures that no value is present for LeverageStorageSnapshotsForHyperflex, not even an explicit nil
### GetLogBackupJobPolicy

`func (o *BackupJobProto) GetLogBackupJobPolicy() JobPolicyProto`

GetLogBackupJobPolicy returns the LogBackupJobPolicy field if non-nil, zero value otherwise.

### GetLogBackupJobPolicyOk

`func (o *BackupJobProto) GetLogBackupJobPolicyOk() (*JobPolicyProto, bool)`

GetLogBackupJobPolicyOk returns a tuple with the LogBackupJobPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupJobPolicy

`func (o *BackupJobProto) SetLogBackupJobPolicy(v JobPolicyProto)`

SetLogBackupJobPolicy sets LogBackupJobPolicy field to given value.

### HasLogBackupJobPolicy

`func (o *BackupJobProto) HasLogBackupJobPolicy() bool`

HasLogBackupJobPolicy returns a boolean if a field has been set.

### GetMaxAllowedSourceSnapshots

`func (o *BackupJobProto) GetMaxAllowedSourceSnapshots() int32`

GetMaxAllowedSourceSnapshots returns the MaxAllowedSourceSnapshots field if non-nil, zero value otherwise.

### GetMaxAllowedSourceSnapshotsOk

`func (o *BackupJobProto) GetMaxAllowedSourceSnapshotsOk() (*int32, bool)`

GetMaxAllowedSourceSnapshotsOk returns a tuple with the MaxAllowedSourceSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedSourceSnapshots

`func (o *BackupJobProto) SetMaxAllowedSourceSnapshots(v int32)`

SetMaxAllowedSourceSnapshots sets MaxAllowedSourceSnapshots field to given value.

### HasMaxAllowedSourceSnapshots

`func (o *BackupJobProto) HasMaxAllowedSourceSnapshots() bool`

HasMaxAllowedSourceSnapshots returns a boolean if a field has been set.

### SetMaxAllowedSourceSnapshotsNil

`func (o *BackupJobProto) SetMaxAllowedSourceSnapshotsNil(b bool)`

 SetMaxAllowedSourceSnapshotsNil sets the value for MaxAllowedSourceSnapshots to be an explicit nil

### UnsetMaxAllowedSourceSnapshots
`func (o *BackupJobProto) UnsetMaxAllowedSourceSnapshots()`

UnsetMaxAllowedSourceSnapshots ensures that no value is present for MaxAllowedSourceSnapshots, not even an explicit nil
### GetName

`func (o *BackupJobProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupJobProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupJobProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupJobProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupJobProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupJobProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumSnapshotsToKeepOnPrimary

`func (o *BackupJobProto) GetNumSnapshotsToKeepOnPrimary() int64`

GetNumSnapshotsToKeepOnPrimary returns the NumSnapshotsToKeepOnPrimary field if non-nil, zero value otherwise.

### GetNumSnapshotsToKeepOnPrimaryOk

`func (o *BackupJobProto) GetNumSnapshotsToKeepOnPrimaryOk() (*int64, bool)`

GetNumSnapshotsToKeepOnPrimaryOk returns a tuple with the NumSnapshotsToKeepOnPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshotsToKeepOnPrimary

`func (o *BackupJobProto) SetNumSnapshotsToKeepOnPrimary(v int64)`

SetNumSnapshotsToKeepOnPrimary sets NumSnapshotsToKeepOnPrimary field to given value.

### HasNumSnapshotsToKeepOnPrimary

`func (o *BackupJobProto) HasNumSnapshotsToKeepOnPrimary() bool`

HasNumSnapshotsToKeepOnPrimary returns a boolean if a field has been set.

### SetNumSnapshotsToKeepOnPrimaryNil

`func (o *BackupJobProto) SetNumSnapshotsToKeepOnPrimaryNil(b bool)`

 SetNumSnapshotsToKeepOnPrimaryNil sets the value for NumSnapshotsToKeepOnPrimary to be an explicit nil

### UnsetNumSnapshotsToKeepOnPrimary
`func (o *BackupJobProto) UnsetNumSnapshotsToKeepOnPrimary()`

UnsetNumSnapshotsToKeepOnPrimary ensures that no value is present for NumSnapshotsToKeepOnPrimary, not even an explicit nil
### GetParentSource

`func (o *BackupJobProto) GetParentSource() EntityProto`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *BackupJobProto) GetParentSourceOk() (*EntityProto, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *BackupJobProto) SetParentSource(v EntityProto)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *BackupJobProto) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPerformSourceSideDedup

`func (o *BackupJobProto) GetPerformSourceSideDedup() bool`

GetPerformSourceSideDedup returns the PerformSourceSideDedup field if non-nil, zero value otherwise.

### GetPerformSourceSideDedupOk

`func (o *BackupJobProto) GetPerformSourceSideDedupOk() (*bool, bool)`

GetPerformSourceSideDedupOk returns a tuple with the PerformSourceSideDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDedup

`func (o *BackupJobProto) SetPerformSourceSideDedup(v bool)`

SetPerformSourceSideDedup sets PerformSourceSideDedup field to given value.

### HasPerformSourceSideDedup

`func (o *BackupJobProto) HasPerformSourceSideDedup() bool`

HasPerformSourceSideDedup returns a boolean if a field has been set.

### SetPerformSourceSideDedupNil

`func (o *BackupJobProto) SetPerformSourceSideDedupNil(b bool)`

 SetPerformSourceSideDedupNil sets the value for PerformSourceSideDedup to be an explicit nil

### UnsetPerformSourceSideDedup
`func (o *BackupJobProto) UnsetPerformSourceSideDedup()`

UnsetPerformSourceSideDedup ensures that no value is present for PerformSourceSideDedup, not even an explicit nil
### GetPolicyAppliedTimeMsecs

`func (o *BackupJobProto) GetPolicyAppliedTimeMsecs() int64`

GetPolicyAppliedTimeMsecs returns the PolicyAppliedTimeMsecs field if non-nil, zero value otherwise.

### GetPolicyAppliedTimeMsecsOk

`func (o *BackupJobProto) GetPolicyAppliedTimeMsecsOk() (*int64, bool)`

GetPolicyAppliedTimeMsecsOk returns a tuple with the PolicyAppliedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAppliedTimeMsecs

`func (o *BackupJobProto) SetPolicyAppliedTimeMsecs(v int64)`

SetPolicyAppliedTimeMsecs sets PolicyAppliedTimeMsecs field to given value.

### HasPolicyAppliedTimeMsecs

`func (o *BackupJobProto) HasPolicyAppliedTimeMsecs() bool`

HasPolicyAppliedTimeMsecs returns a boolean if a field has been set.

### SetPolicyAppliedTimeMsecsNil

`func (o *BackupJobProto) SetPolicyAppliedTimeMsecsNil(b bool)`

 SetPolicyAppliedTimeMsecsNil sets the value for PolicyAppliedTimeMsecs to be an explicit nil

### UnsetPolicyAppliedTimeMsecs
`func (o *BackupJobProto) UnsetPolicyAppliedTimeMsecs()`

UnsetPolicyAppliedTimeMsecs ensures that no value is present for PolicyAppliedTimeMsecs, not even an explicit nil
### GetPolicyId

`func (o *BackupJobProto) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *BackupJobProto) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *BackupJobProto) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *BackupJobProto) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *BackupJobProto) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *BackupJobProto) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *BackupJobProto) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *BackupJobProto) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *BackupJobProto) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *BackupJobProto) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *BackupJobProto) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *BackupJobProto) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPostBackupScript

`func (o *BackupJobProto) GetPostBackupScript() BackupJobPreOrPostScript`

GetPostBackupScript returns the PostBackupScript field if non-nil, zero value otherwise.

### GetPostBackupScriptOk

`func (o *BackupJobProto) GetPostBackupScriptOk() (*BackupJobPreOrPostScript, bool)`

GetPostBackupScriptOk returns a tuple with the PostBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBackupScript

`func (o *BackupJobProto) SetPostBackupScript(v BackupJobPreOrPostScript)`

SetPostBackupScript sets PostBackupScript field to given value.

### HasPostBackupScript

`func (o *BackupJobProto) HasPostBackupScript() bool`

HasPostBackupScript returns a boolean if a field has been set.

### GetPreScript

`func (o *BackupJobProto) GetPreScript() BackupJobPreOrPostScript`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *BackupJobProto) GetPreScriptOk() (*BackupJobPreOrPostScript, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *BackupJobProto) SetPreScript(v BackupJobPreOrPostScript)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *BackupJobProto) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPrimaryJobUid

`func (o *BackupJobProto) GetPrimaryJobUid() UniversalIdProto`

GetPrimaryJobUid returns the PrimaryJobUid field if non-nil, zero value otherwise.

### GetPrimaryJobUidOk

`func (o *BackupJobProto) GetPrimaryJobUidOk() (*UniversalIdProto, bool)`

GetPrimaryJobUidOk returns a tuple with the PrimaryJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryJobUid

`func (o *BackupJobProto) SetPrimaryJobUid(v UniversalIdProto)`

SetPrimaryJobUid sets PrimaryJobUid field to given value.

### HasPrimaryJobUid

`func (o *BackupJobProto) HasPrimaryJobUid() bool`

HasPrimaryJobUid returns a boolean if a field has been set.

### GetPriority

`func (o *BackupJobProto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BackupJobProto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BackupJobProto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BackupJobProto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *BackupJobProto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *BackupJobProto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQuiesce

`func (o *BackupJobProto) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *BackupJobProto) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *BackupJobProto) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *BackupJobProto) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### SetQuiesceNil

`func (o *BackupJobProto) SetQuiesceNil(b bool)`

 SetQuiesceNil sets the value for Quiesce to be an explicit nil

### UnsetQuiesce
`func (o *BackupJobProto) UnsetQuiesce()`

UnsetQuiesce ensures that no value is present for Quiesce, not even an explicit nil
### GetRemoteJobUids

`func (o *BackupJobProto) GetRemoteJobUids() []UniversalIdProto`

GetRemoteJobUids returns the RemoteJobUids field if non-nil, zero value otherwise.

### GetRemoteJobUidsOk

`func (o *BackupJobProto) GetRemoteJobUidsOk() (*[]UniversalIdProto, bool)`

GetRemoteJobUidsOk returns a tuple with the RemoteJobUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteJobUids

`func (o *BackupJobProto) SetRemoteJobUids(v []UniversalIdProto)`

SetRemoteJobUids sets RemoteJobUids field to given value.

### HasRemoteJobUids

`func (o *BackupJobProto) HasRemoteJobUids() bool`

HasRemoteJobUids returns a boolean if a field has been set.

### SetRemoteJobUidsNil

`func (o *BackupJobProto) SetRemoteJobUidsNil(b bool)`

 SetRemoteJobUidsNil sets the value for RemoteJobUids to be an explicit nil

### UnsetRemoteJobUids
`func (o *BackupJobProto) UnsetRemoteJobUids()`

UnsetRemoteJobUids ensures that no value is present for RemoteJobUids, not even an explicit nil
### GetRemoteViewName

`func (o *BackupJobProto) GetRemoteViewName() string`

GetRemoteViewName returns the RemoteViewName field if non-nil, zero value otherwise.

### GetRemoteViewNameOk

`func (o *BackupJobProto) GetRemoteViewNameOk() (*string, bool)`

GetRemoteViewNameOk returns a tuple with the RemoteViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteViewName

`func (o *BackupJobProto) SetRemoteViewName(v string)`

SetRemoteViewName sets RemoteViewName field to given value.

### HasRemoteViewName

`func (o *BackupJobProto) HasRemoteViewName() bool`

HasRemoteViewName returns a boolean if a field has been set.

### SetRemoteViewNameNil

`func (o *BackupJobProto) SetRemoteViewNameNil(b bool)`

 SetRemoteViewNameNil sets the value for RemoteViewName to be an explicit nil

### UnsetRemoteViewName
`func (o *BackupJobProto) UnsetRemoteViewName()`

UnsetRemoteViewName ensures that no value is present for RemoteViewName, not even an explicit nil
### GetRequiredFeatureVec

`func (o *BackupJobProto) GetRequiredFeatureVec() []string`

GetRequiredFeatureVec returns the RequiredFeatureVec field if non-nil, zero value otherwise.

### GetRequiredFeatureVecOk

`func (o *BackupJobProto) GetRequiredFeatureVecOk() (*[]string, bool)`

GetRequiredFeatureVecOk returns a tuple with the RequiredFeatureVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFeatureVec

`func (o *BackupJobProto) SetRequiredFeatureVec(v []string)`

SetRequiredFeatureVec sets RequiredFeatureVec field to given value.

### HasRequiredFeatureVec

`func (o *BackupJobProto) HasRequiredFeatureVec() bool`

HasRequiredFeatureVec returns a boolean if a field has been set.

### SetRequiredFeatureVecNil

`func (o *BackupJobProto) SetRequiredFeatureVecNil(b bool)`

 SetRequiredFeatureVecNil sets the value for RequiredFeatureVec to be an explicit nil

### UnsetRequiredFeatureVec
`func (o *BackupJobProto) UnsetRequiredFeatureVec()`

UnsetRequiredFeatureVec ensures that no value is present for RequiredFeatureVec, not even an explicit nil
### GetSlaTimeMins

`func (o *BackupJobProto) GetSlaTimeMins() int64`

GetSlaTimeMins returns the SlaTimeMins field if non-nil, zero value otherwise.

### GetSlaTimeMinsOk

`func (o *BackupJobProto) GetSlaTimeMinsOk() (*int64, bool)`

GetSlaTimeMinsOk returns a tuple with the SlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaTimeMins

`func (o *BackupJobProto) SetSlaTimeMins(v int64)`

SetSlaTimeMins sets SlaTimeMins field to given value.

### HasSlaTimeMins

`func (o *BackupJobProto) HasSlaTimeMins() bool`

HasSlaTimeMins returns a boolean if a field has been set.

### SetSlaTimeMinsNil

`func (o *BackupJobProto) SetSlaTimeMinsNil(b bool)`

 SetSlaTimeMinsNil sets the value for SlaTimeMins to be an explicit nil

### UnsetSlaTimeMins
`func (o *BackupJobProto) UnsetSlaTimeMins()`

UnsetSlaTimeMins ensures that no value is present for SlaTimeMins, not even an explicit nil
### GetSourceFilters

`func (o *BackupJobProto) GetSourceFilters() SourceFilters`

GetSourceFilters returns the SourceFilters field if non-nil, zero value otherwise.

### GetSourceFiltersOk

`func (o *BackupJobProto) GetSourceFiltersOk() (*SourceFilters, bool)`

GetSourceFiltersOk returns a tuple with the SourceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilters

`func (o *BackupJobProto) SetSourceFilters(v SourceFilters)`

SetSourceFilters sets SourceFilters field to given value.

### HasSourceFilters

`func (o *BackupJobProto) HasSourceFilters() bool`

HasSourceFilters returns a boolean if a field has been set.

### GetSources

`func (o *BackupJobProto) GetSources() []BackupJobProtoBackupSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BackupJobProto) GetSourcesOk() (*[]BackupJobProtoBackupSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BackupJobProto) SetSources(v []BackupJobProtoBackupSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BackupJobProto) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *BackupJobProto) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *BackupJobProto) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetStartTime

`func (o *BackupJobProto) GetStartTime() Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupJobProto) GetStartTimeOk() (*Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupJobProto) SetStartTime(v Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupJobProto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStubbingPolicy

`func (o *BackupJobProto) GetStubbingPolicy() StubbingPolicyProto`

GetStubbingPolicy returns the StubbingPolicy field if non-nil, zero value otherwise.

### GetStubbingPolicyOk

`func (o *BackupJobProto) GetStubbingPolicyOk() (*StubbingPolicyProto, bool)`

GetStubbingPolicyOk returns a tuple with the StubbingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubbingPolicy

`func (o *BackupJobProto) SetStubbingPolicy(v StubbingPolicyProto)`

SetStubbingPolicy sets StubbingPolicy field to given value.

### HasStubbingPolicy

`func (o *BackupJobProto) HasStubbingPolicy() bool`

HasStubbingPolicy returns a boolean if a field has been set.

### GetTagVec

`func (o *BackupJobProto) GetTagVec() []string`

GetTagVec returns the TagVec field if non-nil, zero value otherwise.

### GetTagVecOk

`func (o *BackupJobProto) GetTagVecOk() (*[]string, bool)`

GetTagVecOk returns a tuple with the TagVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagVec

`func (o *BackupJobProto) SetTagVec(v []string)`

SetTagVec sets TagVec field to given value.

### HasTagVec

`func (o *BackupJobProto) HasTagVec() bool`

HasTagVec returns a boolean if a field has been set.

### SetTagVecNil

`func (o *BackupJobProto) SetTagVecNil(b bool)`

 SetTagVecNil sets the value for TagVec to be an explicit nil

### UnsetTagVec
`func (o *BackupJobProto) UnsetTagVec()`

UnsetTagVec ensures that no value is present for TagVec, not even an explicit nil
### GetTimezone

`func (o *BackupJobProto) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BackupJobProto) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BackupJobProto) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BackupJobProto) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *BackupJobProto) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *BackupJobProto) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTruncateLogs

`func (o *BackupJobProto) GetTruncateLogs() bool`

GetTruncateLogs returns the TruncateLogs field if non-nil, zero value otherwise.

### GetTruncateLogsOk

`func (o *BackupJobProto) GetTruncateLogsOk() (*bool, bool)`

GetTruncateLogsOk returns a tuple with the TruncateLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogs

`func (o *BackupJobProto) SetTruncateLogs(v bool)`

SetTruncateLogs sets TruncateLogs field to given value.

### HasTruncateLogs

`func (o *BackupJobProto) HasTruncateLogs() bool`

HasTruncateLogs returns a boolean if a field has been set.

### SetTruncateLogsNil

`func (o *BackupJobProto) SetTruncateLogsNil(b bool)`

 SetTruncateLogsNil sets the value for TruncateLogs to be an explicit nil

### UnsetTruncateLogs
`func (o *BackupJobProto) UnsetTruncateLogs()`

UnsetTruncateLogs ensures that no value is present for TruncateLogs, not even an explicit nil
### GetType

`func (o *BackupJobProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupJobProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupJobProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *BackupJobProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BackupJobProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BackupJobProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserInfo

`func (o *BackupJobProto) GetUserInfo() UserInformation`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *BackupJobProto) GetUserInfoOk() (*UserInformation, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *BackupJobProto) SetUserInfo(v UserInformation)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *BackupJobProto) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetViewBoxId

`func (o *BackupJobProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *BackupJobProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *BackupJobProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *BackupJobProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *BackupJobProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *BackupJobProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


