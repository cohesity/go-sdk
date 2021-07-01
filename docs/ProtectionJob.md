# ProtectionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeverageSanTransport** | Pointer to **NullableBool** | If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network. | [optional] 
**AbortInBlackoutPeriod** | Pointer to **NullableBool** | If true, the Cohesity Cluster aborts any currently executing Job Runs of this Protection Job when a blackout period specified for this Job starts, even if the Job Run started before the blackout period began. If false, a Job Run continues to execute, if the Job Run started before the blackout period starts. | [optional] 
**AlertingConfig** | Pointer to [**AlertingConfig**](AlertingConfig.md) |  | [optional] 
**AlertingPolicy** | Pointer to **[]string** | Array of Job Events.  During Job Runs, the following Job Events are generated: 1) Job succeeds 2) Job fails 3) Job violates the SLA These Job Events can cause Alerts to be generated. &#39;kSuccess&#39; means the Protection Job succeeded. &#39;kFailure&#39; means the Protection Job failed. &#39;kSlaViolation&#39; means the Protection Job took longer than the time period specified in the SLA. | [optional] 
**CloudParameters** | Pointer to [**CloudParameters**](CloudParameters.md) |  | [optional] 
**ContinueOnQuiesceFailure** | Pointer to **NullableBool** | Whether to continue backing up on quiesce failure. | [optional] 
**CreateRemoteView** | Pointer to **NullableBool** | Specifies whether to create a remote view name to use for view overwrite. | [optional] 
**CreationTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the Protection Job was created. | [optional] 
**DataMigrationPolicy** | Pointer to [**DataMigrationPolicy**](DataMigrationPolicy.md) |  | [optional] 
**DedupDisabledSourceIds** | Pointer to **[]int64** | List of source ids for which source side dedup is disabled from the backup job. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a text description about the Protection Job. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the epoch time (in microseconds) after which the Protection Job becomes dormant. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type (such as kVMware or kSQL) of the Protection Source this Job is protecting. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**EnvironmentParameters** | Pointer to [**EnvironmentTypeJobParameters**](EnvironmentTypeJobParameters.md) |  | [optional] 
**ExcludeSourceIds** | Pointer to **[]int64** | Array of Excluded Source Objects.  List of Object ids from a Protection Source that should not be protected and are excluded from being backed up by the Protection Job. Leaf and non-leaf Objects may be in this list and an Object in this list must have an ancestor in the sourceId list. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of Arrays of VM Tag Ids that Specify VMs to Exclude.  Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the &#39;Former Employees&#39; VMs in the East and West but keep all the VMs for &#39;Former Employees&#39; in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the &#39;Former Employee&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;East&#39; (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;West&#39; (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job. | [optional] 
**FullProtectionSlaTimeMins** | Pointer to **NullableInt64** | If specified, this setting is number of minutes that a Job Run of a Full (no CBT) backup schedule is expected to complete, which is known as a Service-Level Agreement (SLA). A SLA violation is reported when the run time of a Job Run exceeds the SLA time period specified for this backup schedule. | [optional] 
**FullProtectionStartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the time of day to start the Full Protection Schedule. This is optional and only applicable if the Protection Policy defines a monthly or a daily Full (no CBT) Protection Schedule. Default value is 02:00 AM. deprecated: true | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies an id for the Protection Job. | [optional] 
**IncrementalProtectionSlaTimeMins** | Pointer to **NullableInt64** | If specified, this setting is number of minutes that a Job Run of a CBT-based backup schedule is expected to complete, which is known as a Service-Level Agreement (SLA). A SLA violation is reported when the run time of a Job Run exceeds the SLA time period specified for this backup schedule. | [optional] 
**IncrementalProtectionStartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the time of day to start the CBT-based Protection Schedule. This is optional and only applicable if the Protection Policy defines a monthly or a daily CBT-based Protection Schedule. Default value is 02:00 AM. deprecated: true | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**IsActive** | Pointer to **NullableBool** | Indicates if the current state of the Protection Job is Active or Inactive. When you create a Protection Job on a Primary Cluster with a replication schedule, the Cohesity Cluster creates an Inactive copy of the Job on the Remote Cluster. In addition, when an Active and running Job is deactivated, the Job becomes Inactive. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Equals &#39;true&#39; if the Protection Job was deleted but some Snapshots are still associated with this Job. If &#39;true&#39;, no new Job Runs are started by this Protection Job. | [optional] 
**IsDirectArchiveEnabled** | Pointer to **NullableBool** | Specifies if this is a direct archive backup job. | [optional] 
**IsNativeFormat** | Pointer to **NullableBool** | Specifies if native format should be used for archiving, applicable for only direct archive jobs. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Indicates if the Protection Job is paused, which means that no new Job Runs are started but any existing Job Runs continue to execute. | [optional] 
**LastRun** | Pointer to [**ProtectionRunInstance**](ProtectionRunInstance.md) |  | [optional] 
**LeverageStorageSnapshots** | Pointer to **NullableBool** | Specifies whether to leverage the storage array based snapshots for this backup job. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, job will fallback to the default backup method. | [optional] 
**LeverageStorageSnapshotsForHyperflex** | Pointer to **NullableBool** | Specifies whether to leverage Hyperflex as the storage snapshot array | [optional] 
**MissingEntities** | Pointer to [**[]ProtectionSource**](ProtectionSource.md) | Specifies Information about missing entities. | [optional] 
**ModificationTimeUsecs** | Pointer to **NullableInt64** | Specifies the last time this Job was updated. | [optional] 
**ModifiedByUser** | Pointer to **NullableString** | Specifies the last Cohesity user who updated this Job. | [optional] 
**Name** | **NullableString** | Specifies the name of the Protection Job. | 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the registered Protection Source that is the parent of the Objects that may be protected by this Job. For example when a vCenter Server is registered on a Cohesity Cluster, the Cohesity Cluster assigns a unique id to this field that represents the vCenter Server. | [optional] 
**PerformSourceSideDedup** | Pointer to **NullableBool** | Specifies whether source side dedupe should be performed or not. | [optional] 
**PolicyAppliedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time (in milliseconds) when the associated Policy was last applied to this Job. This is used to determine if the Policy has changed since it was last applied to this Job. | [optional] 
**PolicyId** | **NullableString** | Specifies the unique id of the Protection Policy associated with the Protection Job. The Policy provides retry settings, Protection Schedules, Priority, SLA, etc. The Job defines the Storage Domain (View Box), the Objects to Protect (if applicable), Start Time, Indexing settings, etc. | 
**PostBackupScript** | Pointer to [**NullableBackupScript**](BackupScript.md) | Specifies the script associated with the backup job. This field must be specified for &#39;kPhysical&#39; jobs. This script will be executed post backup run. | [optional] 
**PreBackupScript** | Pointer to [**NullableBackupScript**](BackupScript.md) | Specifies the script associated with the backup job. This field must be specified for &#39;kPhysical&#39; jobs. This script will be executed pre backup run. The &#39;remoteScript&#39; field will be used for remote adapter jobs and &#39;preBackupScript&#39; field will be used for &#39;kPhysical&#39; jobs. | [optional] 
**Priority** | Pointer to **NullableString** | Specifies the priority of execution for a Protection Job. Cohesity supports concurrent backups but if the number of Jobs exceeds the ability to process Jobs, the specified priority determines the execution Job priority. This field also specifies the replication priority. &#39;kLow&#39; indicates lowest execution priority for a Protection job. &#39;kMedium&#39; indicates medium execution priority for a Protection job. &#39;kHigh&#39; indicates highest execution priority for a Protection job. | [optional] 
**QosType** | Pointer to **NullableString** | Specifies the QoS policy type to use for this Protection Job. &#39;kBackupHDD&#39; indicates the Cohesity Cluster writes data directly to the HDD tier for this Protection Job. This is the recommended setting. &#39;kBackupSSD&#39; indicates the Cohesity Cluster writes data directly to the SSD tier for this Protection Job. Only specify this policy if you need fast ingest speed for a small number of Protection Jobs. | [optional] 
**Quiesce** | Pointer to **NullableBool** | Indicates if the App-Consistent option is enabled for this Job. If the option is enabled, the Cohesity Cluster quiesces the file system and applications before taking Application-Consistent Snapshots. VMware Tools must be installed on the guest Operating System. | [optional] 
**RemoteScript** | Pointer to [**NullableRemoteJobScript**](RemoteJobScript.md) | For a Remote Adapter &#39;kPuppeteer&#39; Job, this field specifies the settings about the remote script that will be executed by this Job. Only specify this field for Remote Adapter &#39;kPuppeteer&#39; Jobs. | [optional] 
**RemoteViewName** | Pointer to **NullableString** | Specifies the remote view name to use for view overwrite. | [optional] 
**SourceIds** | Pointer to **[]int64** | Array of Protected Source Objects.  Specifies the list of Object ids from the Protection Source to protect (or back up) by the Protection Job. An Object in this list may be descendant of another Object in this list. For example a Datacenter could be selected but its child Host excluded. However, a child VM under the Host could be explicitly selected to be protected. Both the Datacenter and the VM are listed. | [optional] 
**SourceSpecialParameters** | [**[]SourceSpecialParameter**](SourceSpecialParameter.md) | Array of Special Source Parameters.  Specifies additional settings that can apply to a subset of the Sources listed in the Protection Job. For example, you can specify a list of files and folders to protect instead of protecting the entire Physical Server. If this field&#39;s setting conflicts with environmentParameters, then this setting will be used. Specific volume selections must be passed in here to take effect. | 
**StartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the time of day to start the Protection Schedule. This is optional and only applicable if the Protection Policy defines a monthly or a daily Protection Schedule. Default value is 02:00 AM. | [optional] 
**SummaryStats** | Pointer to [**ProtectionJobSummaryStats**](ProtectionJobSummaryStats.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies the timezone to use when calculating time for this Protection Job such as the Job start time. Specify the timezone in the following format: \&quot;Area/Location\&quot;, for example: \&quot;America/New_York\&quot;. | [optional] 
**Uid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a global Protection Job id that is unique across Cohesity Clusters. | [optional] 
**UserSpecifiedTags** | Pointer to **[]string** | Tags associated with the job. User can specify tags/keywords that can indexed by Yoda and can be later searched in UI. For example, user can create a &#39;kPuppeteer&#39; job to backup Oracle DB for &#39;payroll&#39; department. User can specify following tags: &#39;payroll&#39;, &#39;Oracle_DB&#39;. | [optional] 
**ViewBoxId** | **NullableInt64** | Specifies the Storage Domain (View Box) id where this Job writes data. | 
**ViewName** | Pointer to **NullableString** | For a Remote Adapter &#39;kPuppeteer&#39; Job or a &#39;kView&#39; Job, this field specifies a View name that should be protected. Specify this field when creating a Protection Job for the first time for a View. If this field is specified, ParentSourceId, SourceIds, and ExcludeSourceIds should not be specified. | [optional] 
**VmTagIds** | Pointer to **[][]int64** | Array of Arrays of VMs Tags Ids that Specify VMs to Protect.  Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only &#39;Eng&#39; VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the &#39;Eng&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with &#39;Eng&#39; and &#39;East&#39; (an intersection). The outer array combines the list from the inner array with list of VMs tagged with &#39;West&#39; (a union). The list of resulting VMs are protected by this Job. | [optional] 

## Methods

### NewProtectionJob

`func NewProtectionJob(name NullableString, policyId NullableString, sourceSpecialParameters []SourceSpecialParameter, viewBoxId NullableInt64, ) *ProtectionJob`

NewProtectionJob instantiates a new ProtectionJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobWithDefaults

`func NewProtectionJobWithDefaults() *ProtectionJob`

NewProtectionJobWithDefaults instantiates a new ProtectionJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeverageSanTransport

`func (o *ProtectionJob) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *ProtectionJob) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *ProtectionJob) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *ProtectionJob) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *ProtectionJob) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *ProtectionJob) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetAbortInBlackoutPeriod

`func (o *ProtectionJob) GetAbortInBlackoutPeriod() bool`

GetAbortInBlackoutPeriod returns the AbortInBlackoutPeriod field if non-nil, zero value otherwise.

### GetAbortInBlackoutPeriodOk

`func (o *ProtectionJob) GetAbortInBlackoutPeriodOk() (*bool, bool)`

GetAbortInBlackoutPeriodOk returns a tuple with the AbortInBlackoutPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackoutPeriod

`func (o *ProtectionJob) SetAbortInBlackoutPeriod(v bool)`

SetAbortInBlackoutPeriod sets AbortInBlackoutPeriod field to given value.

### HasAbortInBlackoutPeriod

`func (o *ProtectionJob) HasAbortInBlackoutPeriod() bool`

HasAbortInBlackoutPeriod returns a boolean if a field has been set.

### SetAbortInBlackoutPeriodNil

`func (o *ProtectionJob) SetAbortInBlackoutPeriodNil(b bool)`

 SetAbortInBlackoutPeriodNil sets the value for AbortInBlackoutPeriod to be an explicit nil

### UnsetAbortInBlackoutPeriod
`func (o *ProtectionJob) UnsetAbortInBlackoutPeriod()`

UnsetAbortInBlackoutPeriod ensures that no value is present for AbortInBlackoutPeriod, not even an explicit nil
### GetAlertingConfig

`func (o *ProtectionJob) GetAlertingConfig() AlertingConfig`

GetAlertingConfig returns the AlertingConfig field if non-nil, zero value otherwise.

### GetAlertingConfigOk

`func (o *ProtectionJob) GetAlertingConfigOk() (*AlertingConfig, bool)`

GetAlertingConfigOk returns a tuple with the AlertingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingConfig

`func (o *ProtectionJob) SetAlertingConfig(v AlertingConfig)`

SetAlertingConfig sets AlertingConfig field to given value.

### HasAlertingConfig

`func (o *ProtectionJob) HasAlertingConfig() bool`

HasAlertingConfig returns a boolean if a field has been set.

### GetAlertingPolicy

`func (o *ProtectionJob) GetAlertingPolicy() []string`

GetAlertingPolicy returns the AlertingPolicy field if non-nil, zero value otherwise.

### GetAlertingPolicyOk

`func (o *ProtectionJob) GetAlertingPolicyOk() (*[]string, bool)`

GetAlertingPolicyOk returns a tuple with the AlertingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingPolicy

`func (o *ProtectionJob) SetAlertingPolicy(v []string)`

SetAlertingPolicy sets AlertingPolicy field to given value.

### HasAlertingPolicy

`func (o *ProtectionJob) HasAlertingPolicy() bool`

HasAlertingPolicy returns a boolean if a field has been set.

### SetAlertingPolicyNil

`func (o *ProtectionJob) SetAlertingPolicyNil(b bool)`

 SetAlertingPolicyNil sets the value for AlertingPolicy to be an explicit nil

### UnsetAlertingPolicy
`func (o *ProtectionJob) UnsetAlertingPolicy()`

UnsetAlertingPolicy ensures that no value is present for AlertingPolicy, not even an explicit nil
### GetCloudParameters

`func (o *ProtectionJob) GetCloudParameters() CloudParameters`

GetCloudParameters returns the CloudParameters field if non-nil, zero value otherwise.

### GetCloudParametersOk

`func (o *ProtectionJob) GetCloudParametersOk() (*CloudParameters, bool)`

GetCloudParametersOk returns a tuple with the CloudParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudParameters

`func (o *ProtectionJob) SetCloudParameters(v CloudParameters)`

SetCloudParameters sets CloudParameters field to given value.

### HasCloudParameters

`func (o *ProtectionJob) HasCloudParameters() bool`

HasCloudParameters returns a boolean if a field has been set.

### GetContinueOnQuiesceFailure

`func (o *ProtectionJob) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *ProtectionJob) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *ProtectionJob) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *ProtectionJob) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *ProtectionJob) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *ProtectionJob) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
### GetCreateRemoteView

`func (o *ProtectionJob) GetCreateRemoteView() bool`

GetCreateRemoteView returns the CreateRemoteView field if non-nil, zero value otherwise.

### GetCreateRemoteViewOk

`func (o *ProtectionJob) GetCreateRemoteViewOk() (*bool, bool)`

GetCreateRemoteViewOk returns a tuple with the CreateRemoteView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRemoteView

`func (o *ProtectionJob) SetCreateRemoteView(v bool)`

SetCreateRemoteView sets CreateRemoteView field to given value.

### HasCreateRemoteView

`func (o *ProtectionJob) HasCreateRemoteView() bool`

HasCreateRemoteView returns a boolean if a field has been set.

### SetCreateRemoteViewNil

`func (o *ProtectionJob) SetCreateRemoteViewNil(b bool)`

 SetCreateRemoteViewNil sets the value for CreateRemoteView to be an explicit nil

### UnsetCreateRemoteView
`func (o *ProtectionJob) UnsetCreateRemoteView()`

UnsetCreateRemoteView ensures that no value is present for CreateRemoteView, not even an explicit nil
### GetCreationTimeUsecs

`func (o *ProtectionJob) GetCreationTimeUsecs() int64`

GetCreationTimeUsecs returns the CreationTimeUsecs field if non-nil, zero value otherwise.

### GetCreationTimeUsecsOk

`func (o *ProtectionJob) GetCreationTimeUsecsOk() (*int64, bool)`

GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeUsecs

`func (o *ProtectionJob) SetCreationTimeUsecs(v int64)`

SetCreationTimeUsecs sets CreationTimeUsecs field to given value.

### HasCreationTimeUsecs

`func (o *ProtectionJob) HasCreationTimeUsecs() bool`

HasCreationTimeUsecs returns a boolean if a field has been set.

### SetCreationTimeUsecsNil

`func (o *ProtectionJob) SetCreationTimeUsecsNil(b bool)`

 SetCreationTimeUsecsNil sets the value for CreationTimeUsecs to be an explicit nil

### UnsetCreationTimeUsecs
`func (o *ProtectionJob) UnsetCreationTimeUsecs()`

UnsetCreationTimeUsecs ensures that no value is present for CreationTimeUsecs, not even an explicit nil
### GetDataMigrationPolicy

`func (o *ProtectionJob) GetDataMigrationPolicy() DataMigrationPolicy`

GetDataMigrationPolicy returns the DataMigrationPolicy field if non-nil, zero value otherwise.

### GetDataMigrationPolicyOk

`func (o *ProtectionJob) GetDataMigrationPolicyOk() (*DataMigrationPolicy, bool)`

GetDataMigrationPolicyOk returns a tuple with the DataMigrationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationPolicy

`func (o *ProtectionJob) SetDataMigrationPolicy(v DataMigrationPolicy)`

SetDataMigrationPolicy sets DataMigrationPolicy field to given value.

### HasDataMigrationPolicy

`func (o *ProtectionJob) HasDataMigrationPolicy() bool`

HasDataMigrationPolicy returns a boolean if a field has been set.

### GetDedupDisabledSourceIds

`func (o *ProtectionJob) GetDedupDisabledSourceIds() []int64`

GetDedupDisabledSourceIds returns the DedupDisabledSourceIds field if non-nil, zero value otherwise.

### GetDedupDisabledSourceIdsOk

`func (o *ProtectionJob) GetDedupDisabledSourceIdsOk() (*[]int64, bool)`

GetDedupDisabledSourceIdsOk returns a tuple with the DedupDisabledSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupDisabledSourceIds

`func (o *ProtectionJob) SetDedupDisabledSourceIds(v []int64)`

SetDedupDisabledSourceIds sets DedupDisabledSourceIds field to given value.

### HasDedupDisabledSourceIds

`func (o *ProtectionJob) HasDedupDisabledSourceIds() bool`

HasDedupDisabledSourceIds returns a boolean if a field has been set.

### SetDedupDisabledSourceIdsNil

`func (o *ProtectionJob) SetDedupDisabledSourceIdsNil(b bool)`

 SetDedupDisabledSourceIdsNil sets the value for DedupDisabledSourceIds to be an explicit nil

### UnsetDedupDisabledSourceIds
`func (o *ProtectionJob) UnsetDedupDisabledSourceIds()`

UnsetDedupDisabledSourceIds ensures that no value is present for DedupDisabledSourceIds, not even an explicit nil
### GetDescription

`func (o *ProtectionJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionJob) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionJob) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectionJob) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionJob) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionJob) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionJob) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionJob) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionJob) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *ProtectionJob) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionJob) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionJob) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionJob) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionJob) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionJob) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvironmentParameters

`func (o *ProtectionJob) GetEnvironmentParameters() EnvironmentTypeJobParameters`

GetEnvironmentParameters returns the EnvironmentParameters field if non-nil, zero value otherwise.

### GetEnvironmentParametersOk

`func (o *ProtectionJob) GetEnvironmentParametersOk() (*EnvironmentTypeJobParameters, bool)`

GetEnvironmentParametersOk returns a tuple with the EnvironmentParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParameters

`func (o *ProtectionJob) SetEnvironmentParameters(v EnvironmentTypeJobParameters)`

SetEnvironmentParameters sets EnvironmentParameters field to given value.

### HasEnvironmentParameters

`func (o *ProtectionJob) HasEnvironmentParameters() bool`

HasEnvironmentParameters returns a boolean if a field has been set.

### GetExcludeSourceIds

`func (o *ProtectionJob) GetExcludeSourceIds() []int64`

GetExcludeSourceIds returns the ExcludeSourceIds field if non-nil, zero value otherwise.

### GetExcludeSourceIdsOk

`func (o *ProtectionJob) GetExcludeSourceIdsOk() (*[]int64, bool)`

GetExcludeSourceIdsOk returns a tuple with the ExcludeSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSourceIds

`func (o *ProtectionJob) SetExcludeSourceIds(v []int64)`

SetExcludeSourceIds sets ExcludeSourceIds field to given value.

### HasExcludeSourceIds

`func (o *ProtectionJob) HasExcludeSourceIds() bool`

HasExcludeSourceIds returns a boolean if a field has been set.

### SetExcludeSourceIdsNil

`func (o *ProtectionJob) SetExcludeSourceIdsNil(b bool)`

 SetExcludeSourceIdsNil sets the value for ExcludeSourceIds to be an explicit nil

### UnsetExcludeSourceIds
`func (o *ProtectionJob) UnsetExcludeSourceIds()`

UnsetExcludeSourceIds ensures that no value is present for ExcludeSourceIds, not even an explicit nil
### GetExcludeVmTagIds

`func (o *ProtectionJob) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *ProtectionJob) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *ProtectionJob) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *ProtectionJob) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### SetExcludeVmTagIdsNil

`func (o *ProtectionJob) SetExcludeVmTagIdsNil(b bool)`

 SetExcludeVmTagIdsNil sets the value for ExcludeVmTagIds to be an explicit nil

### UnsetExcludeVmTagIds
`func (o *ProtectionJob) UnsetExcludeVmTagIds()`

UnsetExcludeVmTagIds ensures that no value is present for ExcludeVmTagIds, not even an explicit nil
### GetFullProtectionSlaTimeMins

`func (o *ProtectionJob) GetFullProtectionSlaTimeMins() int64`

GetFullProtectionSlaTimeMins returns the FullProtectionSlaTimeMins field if non-nil, zero value otherwise.

### GetFullProtectionSlaTimeMinsOk

`func (o *ProtectionJob) GetFullProtectionSlaTimeMinsOk() (*int64, bool)`

GetFullProtectionSlaTimeMinsOk returns a tuple with the FullProtectionSlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProtectionSlaTimeMins

`func (o *ProtectionJob) SetFullProtectionSlaTimeMins(v int64)`

SetFullProtectionSlaTimeMins sets FullProtectionSlaTimeMins field to given value.

### HasFullProtectionSlaTimeMins

`func (o *ProtectionJob) HasFullProtectionSlaTimeMins() bool`

HasFullProtectionSlaTimeMins returns a boolean if a field has been set.

### SetFullProtectionSlaTimeMinsNil

`func (o *ProtectionJob) SetFullProtectionSlaTimeMinsNil(b bool)`

 SetFullProtectionSlaTimeMinsNil sets the value for FullProtectionSlaTimeMins to be an explicit nil

### UnsetFullProtectionSlaTimeMins
`func (o *ProtectionJob) UnsetFullProtectionSlaTimeMins()`

UnsetFullProtectionSlaTimeMins ensures that no value is present for FullProtectionSlaTimeMins, not even an explicit nil
### GetFullProtectionStartTime

`func (o *ProtectionJob) GetFullProtectionStartTime() TimeOfDay`

GetFullProtectionStartTime returns the FullProtectionStartTime field if non-nil, zero value otherwise.

### GetFullProtectionStartTimeOk

`func (o *ProtectionJob) GetFullProtectionStartTimeOk() (*TimeOfDay, bool)`

GetFullProtectionStartTimeOk returns a tuple with the FullProtectionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProtectionStartTime

`func (o *ProtectionJob) SetFullProtectionStartTime(v TimeOfDay)`

SetFullProtectionStartTime sets FullProtectionStartTime field to given value.

### HasFullProtectionStartTime

`func (o *ProtectionJob) HasFullProtectionStartTime() bool`

HasFullProtectionStartTime returns a boolean if a field has been set.

### SetFullProtectionStartTimeNil

`func (o *ProtectionJob) SetFullProtectionStartTimeNil(b bool)`

 SetFullProtectionStartTimeNil sets the value for FullProtectionStartTime to be an explicit nil

### UnsetFullProtectionStartTime
`func (o *ProtectionJob) UnsetFullProtectionStartTime()`

UnsetFullProtectionStartTime ensures that no value is present for FullProtectionStartTime, not even an explicit nil
### GetId

`func (o *ProtectionJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionJob) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionJob) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionJob) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncrementalProtectionSlaTimeMins

`func (o *ProtectionJob) GetIncrementalProtectionSlaTimeMins() int64`

GetIncrementalProtectionSlaTimeMins returns the IncrementalProtectionSlaTimeMins field if non-nil, zero value otherwise.

### GetIncrementalProtectionSlaTimeMinsOk

`func (o *ProtectionJob) GetIncrementalProtectionSlaTimeMinsOk() (*int64, bool)`

GetIncrementalProtectionSlaTimeMinsOk returns a tuple with the IncrementalProtectionSlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalProtectionSlaTimeMins

`func (o *ProtectionJob) SetIncrementalProtectionSlaTimeMins(v int64)`

SetIncrementalProtectionSlaTimeMins sets IncrementalProtectionSlaTimeMins field to given value.

### HasIncrementalProtectionSlaTimeMins

`func (o *ProtectionJob) HasIncrementalProtectionSlaTimeMins() bool`

HasIncrementalProtectionSlaTimeMins returns a boolean if a field has been set.

### SetIncrementalProtectionSlaTimeMinsNil

`func (o *ProtectionJob) SetIncrementalProtectionSlaTimeMinsNil(b bool)`

 SetIncrementalProtectionSlaTimeMinsNil sets the value for IncrementalProtectionSlaTimeMins to be an explicit nil

### UnsetIncrementalProtectionSlaTimeMins
`func (o *ProtectionJob) UnsetIncrementalProtectionSlaTimeMins()`

UnsetIncrementalProtectionSlaTimeMins ensures that no value is present for IncrementalProtectionSlaTimeMins, not even an explicit nil
### GetIncrementalProtectionStartTime

`func (o *ProtectionJob) GetIncrementalProtectionStartTime() TimeOfDay`

GetIncrementalProtectionStartTime returns the IncrementalProtectionStartTime field if non-nil, zero value otherwise.

### GetIncrementalProtectionStartTimeOk

`func (o *ProtectionJob) GetIncrementalProtectionStartTimeOk() (*TimeOfDay, bool)`

GetIncrementalProtectionStartTimeOk returns a tuple with the IncrementalProtectionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalProtectionStartTime

`func (o *ProtectionJob) SetIncrementalProtectionStartTime(v TimeOfDay)`

SetIncrementalProtectionStartTime sets IncrementalProtectionStartTime field to given value.

### HasIncrementalProtectionStartTime

`func (o *ProtectionJob) HasIncrementalProtectionStartTime() bool`

HasIncrementalProtectionStartTime returns a boolean if a field has been set.

### SetIncrementalProtectionStartTimeNil

`func (o *ProtectionJob) SetIncrementalProtectionStartTimeNil(b bool)`

 SetIncrementalProtectionStartTimeNil sets the value for IncrementalProtectionStartTime to be an explicit nil

### UnsetIncrementalProtectionStartTime
`func (o *ProtectionJob) UnsetIncrementalProtectionStartTime()`

UnsetIncrementalProtectionStartTime ensures that no value is present for IncrementalProtectionStartTime, not even an explicit nil
### GetIndexingPolicy

`func (o *ProtectionJob) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *ProtectionJob) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *ProtectionJob) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *ProtectionJob) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetIsActive

`func (o *ProtectionJob) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProtectionJob) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProtectionJob) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProtectionJob) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProtectionJob) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProtectionJob) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsDeleted

`func (o *ProtectionJob) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProtectionJob) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProtectionJob) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProtectionJob) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ProtectionJob) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ProtectionJob) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetIsDirectArchiveEnabled

`func (o *ProtectionJob) GetIsDirectArchiveEnabled() bool`

GetIsDirectArchiveEnabled returns the IsDirectArchiveEnabled field if non-nil, zero value otherwise.

### GetIsDirectArchiveEnabledOk

`func (o *ProtectionJob) GetIsDirectArchiveEnabledOk() (*bool, bool)`

GetIsDirectArchiveEnabledOk returns a tuple with the IsDirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectArchiveEnabled

`func (o *ProtectionJob) SetIsDirectArchiveEnabled(v bool)`

SetIsDirectArchiveEnabled sets IsDirectArchiveEnabled field to given value.

### HasIsDirectArchiveEnabled

`func (o *ProtectionJob) HasIsDirectArchiveEnabled() bool`

HasIsDirectArchiveEnabled returns a boolean if a field has been set.

### SetIsDirectArchiveEnabledNil

`func (o *ProtectionJob) SetIsDirectArchiveEnabledNil(b bool)`

 SetIsDirectArchiveEnabledNil sets the value for IsDirectArchiveEnabled to be an explicit nil

### UnsetIsDirectArchiveEnabled
`func (o *ProtectionJob) UnsetIsDirectArchiveEnabled()`

UnsetIsDirectArchiveEnabled ensures that no value is present for IsDirectArchiveEnabled, not even an explicit nil
### GetIsNativeFormat

`func (o *ProtectionJob) GetIsNativeFormat() bool`

GetIsNativeFormat returns the IsNativeFormat field if non-nil, zero value otherwise.

### GetIsNativeFormatOk

`func (o *ProtectionJob) GetIsNativeFormatOk() (*bool, bool)`

GetIsNativeFormatOk returns a tuple with the IsNativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeFormat

`func (o *ProtectionJob) SetIsNativeFormat(v bool)`

SetIsNativeFormat sets IsNativeFormat field to given value.

### HasIsNativeFormat

`func (o *ProtectionJob) HasIsNativeFormat() bool`

HasIsNativeFormat returns a boolean if a field has been set.

### SetIsNativeFormatNil

`func (o *ProtectionJob) SetIsNativeFormatNil(b bool)`

 SetIsNativeFormatNil sets the value for IsNativeFormat to be an explicit nil

### UnsetIsNativeFormat
`func (o *ProtectionJob) UnsetIsNativeFormat()`

UnsetIsNativeFormat ensures that no value is present for IsNativeFormat, not even an explicit nil
### GetIsPaused

`func (o *ProtectionJob) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectionJob) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectionJob) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectionJob) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectionJob) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectionJob) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLastRun

`func (o *ProtectionJob) GetLastRun() ProtectionRunInstance`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ProtectionJob) GetLastRunOk() (*ProtectionRunInstance, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ProtectionJob) SetLastRun(v ProtectionRunInstance)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ProtectionJob) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLeverageStorageSnapshots

`func (o *ProtectionJob) GetLeverageStorageSnapshots() bool`

GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsOk

`func (o *ProtectionJob) GetLeverageStorageSnapshotsOk() (*bool, bool)`

GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshots

`func (o *ProtectionJob) SetLeverageStorageSnapshots(v bool)`

SetLeverageStorageSnapshots sets LeverageStorageSnapshots field to given value.

### HasLeverageStorageSnapshots

`func (o *ProtectionJob) HasLeverageStorageSnapshots() bool`

HasLeverageStorageSnapshots returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsNil

`func (o *ProtectionJob) SetLeverageStorageSnapshotsNil(b bool)`

 SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil

### UnsetLeverageStorageSnapshots
`func (o *ProtectionJob) UnsetLeverageStorageSnapshots()`

UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
### GetLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJob) GetLeverageStorageSnapshotsForHyperflex() bool`

GetLeverageStorageSnapshotsForHyperflex returns the LeverageStorageSnapshotsForHyperflex field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsForHyperflexOk

`func (o *ProtectionJob) GetLeverageStorageSnapshotsForHyperflexOk() (*bool, bool)`

GetLeverageStorageSnapshotsForHyperflexOk returns a tuple with the LeverageStorageSnapshotsForHyperflex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJob) SetLeverageStorageSnapshotsForHyperflex(v bool)`

SetLeverageStorageSnapshotsForHyperflex sets LeverageStorageSnapshotsForHyperflex field to given value.

### HasLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJob) HasLeverageStorageSnapshotsForHyperflex() bool`

HasLeverageStorageSnapshotsForHyperflex returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsForHyperflexNil

`func (o *ProtectionJob) SetLeverageStorageSnapshotsForHyperflexNil(b bool)`

 SetLeverageStorageSnapshotsForHyperflexNil sets the value for LeverageStorageSnapshotsForHyperflex to be an explicit nil

### UnsetLeverageStorageSnapshotsForHyperflex
`func (o *ProtectionJob) UnsetLeverageStorageSnapshotsForHyperflex()`

UnsetLeverageStorageSnapshotsForHyperflex ensures that no value is present for LeverageStorageSnapshotsForHyperflex, not even an explicit nil
### GetMissingEntities

`func (o *ProtectionJob) GetMissingEntities() []ProtectionSource`

GetMissingEntities returns the MissingEntities field if non-nil, zero value otherwise.

### GetMissingEntitiesOk

`func (o *ProtectionJob) GetMissingEntitiesOk() (*[]ProtectionSource, bool)`

GetMissingEntitiesOk returns a tuple with the MissingEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingEntities

`func (o *ProtectionJob) SetMissingEntities(v []ProtectionSource)`

SetMissingEntities sets MissingEntities field to given value.

### HasMissingEntities

`func (o *ProtectionJob) HasMissingEntities() bool`

HasMissingEntities returns a boolean if a field has been set.

### SetMissingEntitiesNil

`func (o *ProtectionJob) SetMissingEntitiesNil(b bool)`

 SetMissingEntitiesNil sets the value for MissingEntities to be an explicit nil

### UnsetMissingEntities
`func (o *ProtectionJob) UnsetMissingEntities()`

UnsetMissingEntities ensures that no value is present for MissingEntities, not even an explicit nil
### GetModificationTimeUsecs

`func (o *ProtectionJob) GetModificationTimeUsecs() int64`

GetModificationTimeUsecs returns the ModificationTimeUsecs field if non-nil, zero value otherwise.

### GetModificationTimeUsecsOk

`func (o *ProtectionJob) GetModificationTimeUsecsOk() (*int64, bool)`

GetModificationTimeUsecsOk returns a tuple with the ModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationTimeUsecs

`func (o *ProtectionJob) SetModificationTimeUsecs(v int64)`

SetModificationTimeUsecs sets ModificationTimeUsecs field to given value.

### HasModificationTimeUsecs

`func (o *ProtectionJob) HasModificationTimeUsecs() bool`

HasModificationTimeUsecs returns a boolean if a field has been set.

### SetModificationTimeUsecsNil

`func (o *ProtectionJob) SetModificationTimeUsecsNil(b bool)`

 SetModificationTimeUsecsNil sets the value for ModificationTimeUsecs to be an explicit nil

### UnsetModificationTimeUsecs
`func (o *ProtectionJob) UnsetModificationTimeUsecs()`

UnsetModificationTimeUsecs ensures that no value is present for ModificationTimeUsecs, not even an explicit nil
### GetModifiedByUser

`func (o *ProtectionJob) GetModifiedByUser() string`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *ProtectionJob) GetModifiedByUserOk() (*string, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *ProtectionJob) SetModifiedByUser(v string)`

SetModifiedByUser sets ModifiedByUser field to given value.

### HasModifiedByUser

`func (o *ProtectionJob) HasModifiedByUser() bool`

HasModifiedByUser returns a boolean if a field has been set.

### SetModifiedByUserNil

`func (o *ProtectionJob) SetModifiedByUserNil(b bool)`

 SetModifiedByUserNil sets the value for ModifiedByUser to be an explicit nil

### UnsetModifiedByUser
`func (o *ProtectionJob) UnsetModifiedByUser()`

UnsetModifiedByUser ensures that no value is present for ModifiedByUser, not even an explicit nil
### GetName

`func (o *ProtectionJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionJob) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProtectionJob) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionJob) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceId

`func (o *ProtectionJob) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *ProtectionJob) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *ProtectionJob) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *ProtectionJob) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *ProtectionJob) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *ProtectionJob) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetPerformSourceSideDedup

`func (o *ProtectionJob) GetPerformSourceSideDedup() bool`

GetPerformSourceSideDedup returns the PerformSourceSideDedup field if non-nil, zero value otherwise.

### GetPerformSourceSideDedupOk

`func (o *ProtectionJob) GetPerformSourceSideDedupOk() (*bool, bool)`

GetPerformSourceSideDedupOk returns a tuple with the PerformSourceSideDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDedup

`func (o *ProtectionJob) SetPerformSourceSideDedup(v bool)`

SetPerformSourceSideDedup sets PerformSourceSideDedup field to given value.

### HasPerformSourceSideDedup

`func (o *ProtectionJob) HasPerformSourceSideDedup() bool`

HasPerformSourceSideDedup returns a boolean if a field has been set.

### SetPerformSourceSideDedupNil

`func (o *ProtectionJob) SetPerformSourceSideDedupNil(b bool)`

 SetPerformSourceSideDedupNil sets the value for PerformSourceSideDedup to be an explicit nil

### UnsetPerformSourceSideDedup
`func (o *ProtectionJob) UnsetPerformSourceSideDedup()`

UnsetPerformSourceSideDedup ensures that no value is present for PerformSourceSideDedup, not even an explicit nil
### GetPolicyAppliedTimeMsecs

`func (o *ProtectionJob) GetPolicyAppliedTimeMsecs() int64`

GetPolicyAppliedTimeMsecs returns the PolicyAppliedTimeMsecs field if non-nil, zero value otherwise.

### GetPolicyAppliedTimeMsecsOk

`func (o *ProtectionJob) GetPolicyAppliedTimeMsecsOk() (*int64, bool)`

GetPolicyAppliedTimeMsecsOk returns a tuple with the PolicyAppliedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAppliedTimeMsecs

`func (o *ProtectionJob) SetPolicyAppliedTimeMsecs(v int64)`

SetPolicyAppliedTimeMsecs sets PolicyAppliedTimeMsecs field to given value.

### HasPolicyAppliedTimeMsecs

`func (o *ProtectionJob) HasPolicyAppliedTimeMsecs() bool`

HasPolicyAppliedTimeMsecs returns a boolean if a field has been set.

### SetPolicyAppliedTimeMsecsNil

`func (o *ProtectionJob) SetPolicyAppliedTimeMsecsNil(b bool)`

 SetPolicyAppliedTimeMsecsNil sets the value for PolicyAppliedTimeMsecs to be an explicit nil

### UnsetPolicyAppliedTimeMsecs
`func (o *ProtectionJob) UnsetPolicyAppliedTimeMsecs()`

UnsetPolicyAppliedTimeMsecs ensures that no value is present for PolicyAppliedTimeMsecs, not even an explicit nil
### GetPolicyId

`func (o *ProtectionJob) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionJob) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionJob) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *ProtectionJob) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionJob) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPostBackupScript

`func (o *ProtectionJob) GetPostBackupScript() BackupScript`

GetPostBackupScript returns the PostBackupScript field if non-nil, zero value otherwise.

### GetPostBackupScriptOk

`func (o *ProtectionJob) GetPostBackupScriptOk() (*BackupScript, bool)`

GetPostBackupScriptOk returns a tuple with the PostBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBackupScript

`func (o *ProtectionJob) SetPostBackupScript(v BackupScript)`

SetPostBackupScript sets PostBackupScript field to given value.

### HasPostBackupScript

`func (o *ProtectionJob) HasPostBackupScript() bool`

HasPostBackupScript returns a boolean if a field has been set.

### SetPostBackupScriptNil

`func (o *ProtectionJob) SetPostBackupScriptNil(b bool)`

 SetPostBackupScriptNil sets the value for PostBackupScript to be an explicit nil

### UnsetPostBackupScript
`func (o *ProtectionJob) UnsetPostBackupScript()`

UnsetPostBackupScript ensures that no value is present for PostBackupScript, not even an explicit nil
### GetPreBackupScript

`func (o *ProtectionJob) GetPreBackupScript() BackupScript`

GetPreBackupScript returns the PreBackupScript field if non-nil, zero value otherwise.

### GetPreBackupScriptOk

`func (o *ProtectionJob) GetPreBackupScriptOk() (*BackupScript, bool)`

GetPreBackupScriptOk returns a tuple with the PreBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBackupScript

`func (o *ProtectionJob) SetPreBackupScript(v BackupScript)`

SetPreBackupScript sets PreBackupScript field to given value.

### HasPreBackupScript

`func (o *ProtectionJob) HasPreBackupScript() bool`

HasPreBackupScript returns a boolean if a field has been set.

### SetPreBackupScriptNil

`func (o *ProtectionJob) SetPreBackupScriptNil(b bool)`

 SetPreBackupScriptNil sets the value for PreBackupScript to be an explicit nil

### UnsetPreBackupScript
`func (o *ProtectionJob) UnsetPreBackupScript()`

UnsetPreBackupScript ensures that no value is present for PreBackupScript, not even an explicit nil
### GetPriority

`func (o *ProtectionJob) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProtectionJob) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProtectionJob) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProtectionJob) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ProtectionJob) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ProtectionJob) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosType

`func (o *ProtectionJob) GetQosType() string`

GetQosType returns the QosType field if non-nil, zero value otherwise.

### GetQosTypeOk

`func (o *ProtectionJob) GetQosTypeOk() (*string, bool)`

GetQosTypeOk returns a tuple with the QosType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosType

`func (o *ProtectionJob) SetQosType(v string)`

SetQosType sets QosType field to given value.

### HasQosType

`func (o *ProtectionJob) HasQosType() bool`

HasQosType returns a boolean if a field has been set.

### SetQosTypeNil

`func (o *ProtectionJob) SetQosTypeNil(b bool)`

 SetQosTypeNil sets the value for QosType to be an explicit nil

### UnsetQosType
`func (o *ProtectionJob) UnsetQosType()`

UnsetQosType ensures that no value is present for QosType, not even an explicit nil
### GetQuiesce

`func (o *ProtectionJob) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *ProtectionJob) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *ProtectionJob) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *ProtectionJob) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### SetQuiesceNil

`func (o *ProtectionJob) SetQuiesceNil(b bool)`

 SetQuiesceNil sets the value for Quiesce to be an explicit nil

### UnsetQuiesce
`func (o *ProtectionJob) UnsetQuiesce()`

UnsetQuiesce ensures that no value is present for Quiesce, not even an explicit nil
### GetRemoteScript

`func (o *ProtectionJob) GetRemoteScript() RemoteJobScript`

GetRemoteScript returns the RemoteScript field if non-nil, zero value otherwise.

### GetRemoteScriptOk

`func (o *ProtectionJob) GetRemoteScriptOk() (*RemoteJobScript, bool)`

GetRemoteScriptOk returns a tuple with the RemoteScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteScript

`func (o *ProtectionJob) SetRemoteScript(v RemoteJobScript)`

SetRemoteScript sets RemoteScript field to given value.

### HasRemoteScript

`func (o *ProtectionJob) HasRemoteScript() bool`

HasRemoteScript returns a boolean if a field has been set.

### SetRemoteScriptNil

`func (o *ProtectionJob) SetRemoteScriptNil(b bool)`

 SetRemoteScriptNil sets the value for RemoteScript to be an explicit nil

### UnsetRemoteScript
`func (o *ProtectionJob) UnsetRemoteScript()`

UnsetRemoteScript ensures that no value is present for RemoteScript, not even an explicit nil
### GetRemoteViewName

`func (o *ProtectionJob) GetRemoteViewName() string`

GetRemoteViewName returns the RemoteViewName field if non-nil, zero value otherwise.

### GetRemoteViewNameOk

`func (o *ProtectionJob) GetRemoteViewNameOk() (*string, bool)`

GetRemoteViewNameOk returns a tuple with the RemoteViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteViewName

`func (o *ProtectionJob) SetRemoteViewName(v string)`

SetRemoteViewName sets RemoteViewName field to given value.

### HasRemoteViewName

`func (o *ProtectionJob) HasRemoteViewName() bool`

HasRemoteViewName returns a boolean if a field has been set.

### SetRemoteViewNameNil

`func (o *ProtectionJob) SetRemoteViewNameNil(b bool)`

 SetRemoteViewNameNil sets the value for RemoteViewName to be an explicit nil

### UnsetRemoteViewName
`func (o *ProtectionJob) UnsetRemoteViewName()`

UnsetRemoteViewName ensures that no value is present for RemoteViewName, not even an explicit nil
### GetSourceIds

`func (o *ProtectionJob) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *ProtectionJob) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *ProtectionJob) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *ProtectionJob) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *ProtectionJob) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *ProtectionJob) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil
### GetSourceSpecialParameters

`func (o *ProtectionJob) GetSourceSpecialParameters() []SourceSpecialParameter`

GetSourceSpecialParameters returns the SourceSpecialParameters field if non-nil, zero value otherwise.

### GetSourceSpecialParametersOk

`func (o *ProtectionJob) GetSourceSpecialParametersOk() (*[]SourceSpecialParameter, bool)`

GetSourceSpecialParametersOk returns a tuple with the SourceSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSpecialParameters

`func (o *ProtectionJob) SetSourceSpecialParameters(v []SourceSpecialParameter)`

SetSourceSpecialParameters sets SourceSpecialParameters field to given value.


### SetSourceSpecialParametersNil

`func (o *ProtectionJob) SetSourceSpecialParametersNil(b bool)`

 SetSourceSpecialParametersNil sets the value for SourceSpecialParameters to be an explicit nil

### UnsetSourceSpecialParameters
`func (o *ProtectionJob) UnsetSourceSpecialParameters()`

UnsetSourceSpecialParameters ensures that no value is present for SourceSpecialParameters, not even an explicit nil
### GetStartTime

`func (o *ProtectionJob) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProtectionJob) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProtectionJob) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProtectionJob) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ProtectionJob) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ProtectionJob) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetSummaryStats

`func (o *ProtectionJob) GetSummaryStats() ProtectionJobSummaryStats`

GetSummaryStats returns the SummaryStats field if non-nil, zero value otherwise.

### GetSummaryStatsOk

`func (o *ProtectionJob) GetSummaryStatsOk() (*ProtectionJobSummaryStats, bool)`

GetSummaryStatsOk returns a tuple with the SummaryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStats

`func (o *ProtectionJob) SetSummaryStats(v ProtectionJobSummaryStats)`

SetSummaryStats sets SummaryStats field to given value.

### HasSummaryStats

`func (o *ProtectionJob) HasSummaryStats() bool`

HasSummaryStats returns a boolean if a field has been set.

### GetTenantId

`func (o *ProtectionJob) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProtectionJob) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProtectionJob) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProtectionJob) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProtectionJob) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProtectionJob) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTimezone

`func (o *ProtectionJob) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProtectionJob) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProtectionJob) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProtectionJob) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ProtectionJob) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ProtectionJob) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetUid

`func (o *ProtectionJob) GetUid() UniversalId`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProtectionJob) GetUidOk() (*UniversalId, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProtectionJob) SetUid(v UniversalId)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProtectionJob) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ProtectionJob) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ProtectionJob) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetUserSpecifiedTags

`func (o *ProtectionJob) GetUserSpecifiedTags() []string`

GetUserSpecifiedTags returns the UserSpecifiedTags field if non-nil, zero value otherwise.

### GetUserSpecifiedTagsOk

`func (o *ProtectionJob) GetUserSpecifiedTagsOk() (*[]string, bool)`

GetUserSpecifiedTagsOk returns a tuple with the UserSpecifiedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSpecifiedTags

`func (o *ProtectionJob) SetUserSpecifiedTags(v []string)`

SetUserSpecifiedTags sets UserSpecifiedTags field to given value.

### HasUserSpecifiedTags

`func (o *ProtectionJob) HasUserSpecifiedTags() bool`

HasUserSpecifiedTags returns a boolean if a field has been set.

### SetUserSpecifiedTagsNil

`func (o *ProtectionJob) SetUserSpecifiedTagsNil(b bool)`

 SetUserSpecifiedTagsNil sets the value for UserSpecifiedTags to be an explicit nil

### UnsetUserSpecifiedTags
`func (o *ProtectionJob) UnsetUserSpecifiedTags()`

UnsetUserSpecifiedTags ensures that no value is present for UserSpecifiedTags, not even an explicit nil
### GetViewBoxId

`func (o *ProtectionJob) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *ProtectionJob) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *ProtectionJob) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.


### SetViewBoxIdNil

`func (o *ProtectionJob) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *ProtectionJob) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *ProtectionJob) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ProtectionJob) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ProtectionJob) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ProtectionJob) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ProtectionJob) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ProtectionJob) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVmTagIds

`func (o *ProtectionJob) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *ProtectionJob) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *ProtectionJob) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *ProtectionJob) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *ProtectionJob) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *ProtectionJob) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


