# ProtectionJobRequestBody

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
**IncrementalProtectionSlaTimeMins** | Pointer to **NullableInt64** | If specified, this setting is number of minutes that a Job Run of a CBT-based backup schedule is expected to complete, which is known as a Service-Level Agreement (SLA). A SLA violation is reported when the run time of a Job Run exceeds the SLA time period specified for this backup schedule. | [optional] 
**IncrementalProtectionStartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the time of day to start the CBT-based Protection Schedule. This is optional and only applicable if the Protection Policy defines a monthly or a daily CBT-based Protection Schedule. Default value is 02:00 AM. deprecated: true | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**IsDirectArchiveEnabled** | Pointer to **NullableBool** | Specifies if this is a direct archive backup job. | [optional] 
**IsNativeFormat** | Pointer to **NullableBool** | Specifies if native format should be used for archiving, applicable for only direct archive jobs. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies if the Protection Job is paused, which means that no new Job Runs are started but any existing Job Runs continue to execute. | [optional] 
**LeverageStorageSnapshots** | Pointer to **NullableBool** | Specifies whether to leverage the storage array based snapshots for this backup job. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, job will fallback to the default backup method. | [optional] 
**LeverageStorageSnapshotsForHyperflex** | Pointer to **NullableBool** | Specifies whether to leverage Hyperflex as the storage snapshot array | [optional] 
**Name** | **NullableString** | Specifies the name of the Protection Job. | 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the registered Protection Source that is the parent of the Objects that may be protected by this Job. For example when a vCenter Server is registered on a Cohesity Cluster, the Cohesity Cluster assigns a unique id to this field that represents the vCenter Server. | [optional] 
**PerformSourceSideDedup** | Pointer to **NullableBool** | Specifies whether source side dedupe should be performed or not. | [optional] 
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
**Timezone** | Pointer to **NullableString** | Specifies the timezone to use when calculating time for this Protection Job such as the Job start time. Specify the timezone in the following format: \&quot;Area/Location\&quot;, for example: \&quot;America/New_York\&quot;. | [optional] 
**UserSpecifiedTags** | Pointer to **[]string** | Tags associated with the job. User can specify tags/keywords that can indexed by Yoda and can be later searched in UI. For example, user can create a &#39;kPuppeteer&#39; job to backup Oracle DB for &#39;payroll&#39; department. User can specify following tags: &#39;payroll&#39;, &#39;Oracle_DB&#39;. | [optional] 
**ViewBoxId** | **NullableInt64** | Specifies the Storage Domain (View Box) id where this Job writes data. | 
**ViewName** | Pointer to **NullableString** | For a Remote Adapter &#39;kPuppeteer&#39; Job or a &#39;kView&#39; Job, this field specifies a View name that should be protected. Specify this field when creating a Protection Job for the first time for a View. If this field is specified, ParentSourceId, SourceIds, and ExcludeSourceIds should not be specified. | [optional] 
**VmTagIds** | Pointer to **[][]int64** | Array of Arrays of VMs Tags Ids that Specify VMs to Protect.  Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only &#39;Eng&#39; VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the &#39;Eng&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with &#39;Eng&#39; and &#39;East&#39; (an intersection). The outer array combines the list from the inner array with list of VMs tagged with &#39;West&#39; (a union). The list of resulting VMs are protected by this Job. | [optional] 

## Methods

### NewProtectionJobRequestBody

`func NewProtectionJobRequestBody(name NullableString, policyId NullableString, sourceSpecialParameters []SourceSpecialParameter, viewBoxId NullableInt64, ) *ProtectionJobRequestBody`

NewProtectionJobRequestBody instantiates a new ProtectionJobRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionJobRequestBodyWithDefaults

`func NewProtectionJobRequestBodyWithDefaults() *ProtectionJobRequestBody`

NewProtectionJobRequestBodyWithDefaults instantiates a new ProtectionJobRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeverageSanTransport

`func (o *ProtectionJobRequestBody) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *ProtectionJobRequestBody) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *ProtectionJobRequestBody) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *ProtectionJobRequestBody) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *ProtectionJobRequestBody) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *ProtectionJobRequestBody) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetAbortInBlackoutPeriod

`func (o *ProtectionJobRequestBody) GetAbortInBlackoutPeriod() bool`

GetAbortInBlackoutPeriod returns the AbortInBlackoutPeriod field if non-nil, zero value otherwise.

### GetAbortInBlackoutPeriodOk

`func (o *ProtectionJobRequestBody) GetAbortInBlackoutPeriodOk() (*bool, bool)`

GetAbortInBlackoutPeriodOk returns a tuple with the AbortInBlackoutPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackoutPeriod

`func (o *ProtectionJobRequestBody) SetAbortInBlackoutPeriod(v bool)`

SetAbortInBlackoutPeriod sets AbortInBlackoutPeriod field to given value.

### HasAbortInBlackoutPeriod

`func (o *ProtectionJobRequestBody) HasAbortInBlackoutPeriod() bool`

HasAbortInBlackoutPeriod returns a boolean if a field has been set.

### SetAbortInBlackoutPeriodNil

`func (o *ProtectionJobRequestBody) SetAbortInBlackoutPeriodNil(b bool)`

 SetAbortInBlackoutPeriodNil sets the value for AbortInBlackoutPeriod to be an explicit nil

### UnsetAbortInBlackoutPeriod
`func (o *ProtectionJobRequestBody) UnsetAbortInBlackoutPeriod()`

UnsetAbortInBlackoutPeriod ensures that no value is present for AbortInBlackoutPeriod, not even an explicit nil
### GetAlertingConfig

`func (o *ProtectionJobRequestBody) GetAlertingConfig() AlertingConfig`

GetAlertingConfig returns the AlertingConfig field if non-nil, zero value otherwise.

### GetAlertingConfigOk

`func (o *ProtectionJobRequestBody) GetAlertingConfigOk() (*AlertingConfig, bool)`

GetAlertingConfigOk returns a tuple with the AlertingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingConfig

`func (o *ProtectionJobRequestBody) SetAlertingConfig(v AlertingConfig)`

SetAlertingConfig sets AlertingConfig field to given value.

### HasAlertingConfig

`func (o *ProtectionJobRequestBody) HasAlertingConfig() bool`

HasAlertingConfig returns a boolean if a field has been set.

### GetAlertingPolicy

`func (o *ProtectionJobRequestBody) GetAlertingPolicy() []string`

GetAlertingPolicy returns the AlertingPolicy field if non-nil, zero value otherwise.

### GetAlertingPolicyOk

`func (o *ProtectionJobRequestBody) GetAlertingPolicyOk() (*[]string, bool)`

GetAlertingPolicyOk returns a tuple with the AlertingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingPolicy

`func (o *ProtectionJobRequestBody) SetAlertingPolicy(v []string)`

SetAlertingPolicy sets AlertingPolicy field to given value.

### HasAlertingPolicy

`func (o *ProtectionJobRequestBody) HasAlertingPolicy() bool`

HasAlertingPolicy returns a boolean if a field has been set.

### SetAlertingPolicyNil

`func (o *ProtectionJobRequestBody) SetAlertingPolicyNil(b bool)`

 SetAlertingPolicyNil sets the value for AlertingPolicy to be an explicit nil

### UnsetAlertingPolicy
`func (o *ProtectionJobRequestBody) UnsetAlertingPolicy()`

UnsetAlertingPolicy ensures that no value is present for AlertingPolicy, not even an explicit nil
### GetCloudParameters

`func (o *ProtectionJobRequestBody) GetCloudParameters() CloudParameters`

GetCloudParameters returns the CloudParameters field if non-nil, zero value otherwise.

### GetCloudParametersOk

`func (o *ProtectionJobRequestBody) GetCloudParametersOk() (*CloudParameters, bool)`

GetCloudParametersOk returns a tuple with the CloudParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudParameters

`func (o *ProtectionJobRequestBody) SetCloudParameters(v CloudParameters)`

SetCloudParameters sets CloudParameters field to given value.

### HasCloudParameters

`func (o *ProtectionJobRequestBody) HasCloudParameters() bool`

HasCloudParameters returns a boolean if a field has been set.

### GetContinueOnQuiesceFailure

`func (o *ProtectionJobRequestBody) GetContinueOnQuiesceFailure() bool`

GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field if non-nil, zero value otherwise.

### GetContinueOnQuiesceFailureOk

`func (o *ProtectionJobRequestBody) GetContinueOnQuiesceFailureOk() (*bool, bool)`

GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnQuiesceFailure

`func (o *ProtectionJobRequestBody) SetContinueOnQuiesceFailure(v bool)`

SetContinueOnQuiesceFailure sets ContinueOnQuiesceFailure field to given value.

### HasContinueOnQuiesceFailure

`func (o *ProtectionJobRequestBody) HasContinueOnQuiesceFailure() bool`

HasContinueOnQuiesceFailure returns a boolean if a field has been set.

### SetContinueOnQuiesceFailureNil

`func (o *ProtectionJobRequestBody) SetContinueOnQuiesceFailureNil(b bool)`

 SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil

### UnsetContinueOnQuiesceFailure
`func (o *ProtectionJobRequestBody) UnsetContinueOnQuiesceFailure()`

UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
### GetCreateRemoteView

`func (o *ProtectionJobRequestBody) GetCreateRemoteView() bool`

GetCreateRemoteView returns the CreateRemoteView field if non-nil, zero value otherwise.

### GetCreateRemoteViewOk

`func (o *ProtectionJobRequestBody) GetCreateRemoteViewOk() (*bool, bool)`

GetCreateRemoteViewOk returns a tuple with the CreateRemoteView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRemoteView

`func (o *ProtectionJobRequestBody) SetCreateRemoteView(v bool)`

SetCreateRemoteView sets CreateRemoteView field to given value.

### HasCreateRemoteView

`func (o *ProtectionJobRequestBody) HasCreateRemoteView() bool`

HasCreateRemoteView returns a boolean if a field has been set.

### SetCreateRemoteViewNil

`func (o *ProtectionJobRequestBody) SetCreateRemoteViewNil(b bool)`

 SetCreateRemoteViewNil sets the value for CreateRemoteView to be an explicit nil

### UnsetCreateRemoteView
`func (o *ProtectionJobRequestBody) UnsetCreateRemoteView()`

UnsetCreateRemoteView ensures that no value is present for CreateRemoteView, not even an explicit nil
### GetDataMigrationPolicy

`func (o *ProtectionJobRequestBody) GetDataMigrationPolicy() DataMigrationPolicy`

GetDataMigrationPolicy returns the DataMigrationPolicy field if non-nil, zero value otherwise.

### GetDataMigrationPolicyOk

`func (o *ProtectionJobRequestBody) GetDataMigrationPolicyOk() (*DataMigrationPolicy, bool)`

GetDataMigrationPolicyOk returns a tuple with the DataMigrationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationPolicy

`func (o *ProtectionJobRequestBody) SetDataMigrationPolicy(v DataMigrationPolicy)`

SetDataMigrationPolicy sets DataMigrationPolicy field to given value.

### HasDataMigrationPolicy

`func (o *ProtectionJobRequestBody) HasDataMigrationPolicy() bool`

HasDataMigrationPolicy returns a boolean if a field has been set.

### GetDedupDisabledSourceIds

`func (o *ProtectionJobRequestBody) GetDedupDisabledSourceIds() []int64`

GetDedupDisabledSourceIds returns the DedupDisabledSourceIds field if non-nil, zero value otherwise.

### GetDedupDisabledSourceIdsOk

`func (o *ProtectionJobRequestBody) GetDedupDisabledSourceIdsOk() (*[]int64, bool)`

GetDedupDisabledSourceIdsOk returns a tuple with the DedupDisabledSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupDisabledSourceIds

`func (o *ProtectionJobRequestBody) SetDedupDisabledSourceIds(v []int64)`

SetDedupDisabledSourceIds sets DedupDisabledSourceIds field to given value.

### HasDedupDisabledSourceIds

`func (o *ProtectionJobRequestBody) HasDedupDisabledSourceIds() bool`

HasDedupDisabledSourceIds returns a boolean if a field has been set.

### SetDedupDisabledSourceIdsNil

`func (o *ProtectionJobRequestBody) SetDedupDisabledSourceIdsNil(b bool)`

 SetDedupDisabledSourceIdsNil sets the value for DedupDisabledSourceIds to be an explicit nil

### UnsetDedupDisabledSourceIds
`func (o *ProtectionJobRequestBody) UnsetDedupDisabledSourceIds()`

UnsetDedupDisabledSourceIds ensures that no value is present for DedupDisabledSourceIds, not even an explicit nil
### GetDescription

`func (o *ProtectionJobRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionJobRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionJobRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionJobRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionJobRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionJobRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectionJobRequestBody) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionJobRequestBody) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionJobRequestBody) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionJobRequestBody) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionJobRequestBody) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionJobRequestBody) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *ProtectionJobRequestBody) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionJobRequestBody) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionJobRequestBody) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionJobRequestBody) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionJobRequestBody) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionJobRequestBody) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvironmentParameters

`func (o *ProtectionJobRequestBody) GetEnvironmentParameters() EnvironmentTypeJobParameters`

GetEnvironmentParameters returns the EnvironmentParameters field if non-nil, zero value otherwise.

### GetEnvironmentParametersOk

`func (o *ProtectionJobRequestBody) GetEnvironmentParametersOk() (*EnvironmentTypeJobParameters, bool)`

GetEnvironmentParametersOk returns a tuple with the EnvironmentParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParameters

`func (o *ProtectionJobRequestBody) SetEnvironmentParameters(v EnvironmentTypeJobParameters)`

SetEnvironmentParameters sets EnvironmentParameters field to given value.

### HasEnvironmentParameters

`func (o *ProtectionJobRequestBody) HasEnvironmentParameters() bool`

HasEnvironmentParameters returns a boolean if a field has been set.

### GetExcludeSourceIds

`func (o *ProtectionJobRequestBody) GetExcludeSourceIds() []int64`

GetExcludeSourceIds returns the ExcludeSourceIds field if non-nil, zero value otherwise.

### GetExcludeSourceIdsOk

`func (o *ProtectionJobRequestBody) GetExcludeSourceIdsOk() (*[]int64, bool)`

GetExcludeSourceIdsOk returns a tuple with the ExcludeSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSourceIds

`func (o *ProtectionJobRequestBody) SetExcludeSourceIds(v []int64)`

SetExcludeSourceIds sets ExcludeSourceIds field to given value.

### HasExcludeSourceIds

`func (o *ProtectionJobRequestBody) HasExcludeSourceIds() bool`

HasExcludeSourceIds returns a boolean if a field has been set.

### SetExcludeSourceIdsNil

`func (o *ProtectionJobRequestBody) SetExcludeSourceIdsNil(b bool)`

 SetExcludeSourceIdsNil sets the value for ExcludeSourceIds to be an explicit nil

### UnsetExcludeSourceIds
`func (o *ProtectionJobRequestBody) UnsetExcludeSourceIds()`

UnsetExcludeSourceIds ensures that no value is present for ExcludeSourceIds, not even an explicit nil
### GetExcludeVmTagIds

`func (o *ProtectionJobRequestBody) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *ProtectionJobRequestBody) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *ProtectionJobRequestBody) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *ProtectionJobRequestBody) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### SetExcludeVmTagIdsNil

`func (o *ProtectionJobRequestBody) SetExcludeVmTagIdsNil(b bool)`

 SetExcludeVmTagIdsNil sets the value for ExcludeVmTagIds to be an explicit nil

### UnsetExcludeVmTagIds
`func (o *ProtectionJobRequestBody) UnsetExcludeVmTagIds()`

UnsetExcludeVmTagIds ensures that no value is present for ExcludeVmTagIds, not even an explicit nil
### GetFullProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) GetFullProtectionSlaTimeMins() int64`

GetFullProtectionSlaTimeMins returns the FullProtectionSlaTimeMins field if non-nil, zero value otherwise.

### GetFullProtectionSlaTimeMinsOk

`func (o *ProtectionJobRequestBody) GetFullProtectionSlaTimeMinsOk() (*int64, bool)`

GetFullProtectionSlaTimeMinsOk returns a tuple with the FullProtectionSlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) SetFullProtectionSlaTimeMins(v int64)`

SetFullProtectionSlaTimeMins sets FullProtectionSlaTimeMins field to given value.

### HasFullProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) HasFullProtectionSlaTimeMins() bool`

HasFullProtectionSlaTimeMins returns a boolean if a field has been set.

### SetFullProtectionSlaTimeMinsNil

`func (o *ProtectionJobRequestBody) SetFullProtectionSlaTimeMinsNil(b bool)`

 SetFullProtectionSlaTimeMinsNil sets the value for FullProtectionSlaTimeMins to be an explicit nil

### UnsetFullProtectionSlaTimeMins
`func (o *ProtectionJobRequestBody) UnsetFullProtectionSlaTimeMins()`

UnsetFullProtectionSlaTimeMins ensures that no value is present for FullProtectionSlaTimeMins, not even an explicit nil
### GetFullProtectionStartTime

`func (o *ProtectionJobRequestBody) GetFullProtectionStartTime() TimeOfDay`

GetFullProtectionStartTime returns the FullProtectionStartTime field if non-nil, zero value otherwise.

### GetFullProtectionStartTimeOk

`func (o *ProtectionJobRequestBody) GetFullProtectionStartTimeOk() (*TimeOfDay, bool)`

GetFullProtectionStartTimeOk returns a tuple with the FullProtectionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProtectionStartTime

`func (o *ProtectionJobRequestBody) SetFullProtectionStartTime(v TimeOfDay)`

SetFullProtectionStartTime sets FullProtectionStartTime field to given value.

### HasFullProtectionStartTime

`func (o *ProtectionJobRequestBody) HasFullProtectionStartTime() bool`

HasFullProtectionStartTime returns a boolean if a field has been set.

### SetFullProtectionStartTimeNil

`func (o *ProtectionJobRequestBody) SetFullProtectionStartTimeNil(b bool)`

 SetFullProtectionStartTimeNil sets the value for FullProtectionStartTime to be an explicit nil

### UnsetFullProtectionStartTime
`func (o *ProtectionJobRequestBody) UnsetFullProtectionStartTime()`

UnsetFullProtectionStartTime ensures that no value is present for FullProtectionStartTime, not even an explicit nil
### GetIncrementalProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) GetIncrementalProtectionSlaTimeMins() int64`

GetIncrementalProtectionSlaTimeMins returns the IncrementalProtectionSlaTimeMins field if non-nil, zero value otherwise.

### GetIncrementalProtectionSlaTimeMinsOk

`func (o *ProtectionJobRequestBody) GetIncrementalProtectionSlaTimeMinsOk() (*int64, bool)`

GetIncrementalProtectionSlaTimeMinsOk returns a tuple with the IncrementalProtectionSlaTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) SetIncrementalProtectionSlaTimeMins(v int64)`

SetIncrementalProtectionSlaTimeMins sets IncrementalProtectionSlaTimeMins field to given value.

### HasIncrementalProtectionSlaTimeMins

`func (o *ProtectionJobRequestBody) HasIncrementalProtectionSlaTimeMins() bool`

HasIncrementalProtectionSlaTimeMins returns a boolean if a field has been set.

### SetIncrementalProtectionSlaTimeMinsNil

`func (o *ProtectionJobRequestBody) SetIncrementalProtectionSlaTimeMinsNil(b bool)`

 SetIncrementalProtectionSlaTimeMinsNil sets the value for IncrementalProtectionSlaTimeMins to be an explicit nil

### UnsetIncrementalProtectionSlaTimeMins
`func (o *ProtectionJobRequestBody) UnsetIncrementalProtectionSlaTimeMins()`

UnsetIncrementalProtectionSlaTimeMins ensures that no value is present for IncrementalProtectionSlaTimeMins, not even an explicit nil
### GetIncrementalProtectionStartTime

`func (o *ProtectionJobRequestBody) GetIncrementalProtectionStartTime() TimeOfDay`

GetIncrementalProtectionStartTime returns the IncrementalProtectionStartTime field if non-nil, zero value otherwise.

### GetIncrementalProtectionStartTimeOk

`func (o *ProtectionJobRequestBody) GetIncrementalProtectionStartTimeOk() (*TimeOfDay, bool)`

GetIncrementalProtectionStartTimeOk returns a tuple with the IncrementalProtectionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalProtectionStartTime

`func (o *ProtectionJobRequestBody) SetIncrementalProtectionStartTime(v TimeOfDay)`

SetIncrementalProtectionStartTime sets IncrementalProtectionStartTime field to given value.

### HasIncrementalProtectionStartTime

`func (o *ProtectionJobRequestBody) HasIncrementalProtectionStartTime() bool`

HasIncrementalProtectionStartTime returns a boolean if a field has been set.

### SetIncrementalProtectionStartTimeNil

`func (o *ProtectionJobRequestBody) SetIncrementalProtectionStartTimeNil(b bool)`

 SetIncrementalProtectionStartTimeNil sets the value for IncrementalProtectionStartTime to be an explicit nil

### UnsetIncrementalProtectionStartTime
`func (o *ProtectionJobRequestBody) UnsetIncrementalProtectionStartTime()`

UnsetIncrementalProtectionStartTime ensures that no value is present for IncrementalProtectionStartTime, not even an explicit nil
### GetIndexingPolicy

`func (o *ProtectionJobRequestBody) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *ProtectionJobRequestBody) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *ProtectionJobRequestBody) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *ProtectionJobRequestBody) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetIsDirectArchiveEnabled

`func (o *ProtectionJobRequestBody) GetIsDirectArchiveEnabled() bool`

GetIsDirectArchiveEnabled returns the IsDirectArchiveEnabled field if non-nil, zero value otherwise.

### GetIsDirectArchiveEnabledOk

`func (o *ProtectionJobRequestBody) GetIsDirectArchiveEnabledOk() (*bool, bool)`

GetIsDirectArchiveEnabledOk returns a tuple with the IsDirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectArchiveEnabled

`func (o *ProtectionJobRequestBody) SetIsDirectArchiveEnabled(v bool)`

SetIsDirectArchiveEnabled sets IsDirectArchiveEnabled field to given value.

### HasIsDirectArchiveEnabled

`func (o *ProtectionJobRequestBody) HasIsDirectArchiveEnabled() bool`

HasIsDirectArchiveEnabled returns a boolean if a field has been set.

### SetIsDirectArchiveEnabledNil

`func (o *ProtectionJobRequestBody) SetIsDirectArchiveEnabledNil(b bool)`

 SetIsDirectArchiveEnabledNil sets the value for IsDirectArchiveEnabled to be an explicit nil

### UnsetIsDirectArchiveEnabled
`func (o *ProtectionJobRequestBody) UnsetIsDirectArchiveEnabled()`

UnsetIsDirectArchiveEnabled ensures that no value is present for IsDirectArchiveEnabled, not even an explicit nil
### GetIsNativeFormat

`func (o *ProtectionJobRequestBody) GetIsNativeFormat() bool`

GetIsNativeFormat returns the IsNativeFormat field if non-nil, zero value otherwise.

### GetIsNativeFormatOk

`func (o *ProtectionJobRequestBody) GetIsNativeFormatOk() (*bool, bool)`

GetIsNativeFormatOk returns a tuple with the IsNativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeFormat

`func (o *ProtectionJobRequestBody) SetIsNativeFormat(v bool)`

SetIsNativeFormat sets IsNativeFormat field to given value.

### HasIsNativeFormat

`func (o *ProtectionJobRequestBody) HasIsNativeFormat() bool`

HasIsNativeFormat returns a boolean if a field has been set.

### SetIsNativeFormatNil

`func (o *ProtectionJobRequestBody) SetIsNativeFormatNil(b bool)`

 SetIsNativeFormatNil sets the value for IsNativeFormat to be an explicit nil

### UnsetIsNativeFormat
`func (o *ProtectionJobRequestBody) UnsetIsNativeFormat()`

UnsetIsNativeFormat ensures that no value is present for IsNativeFormat, not even an explicit nil
### GetIsPaused

`func (o *ProtectionJobRequestBody) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectionJobRequestBody) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectionJobRequestBody) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectionJobRequestBody) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectionJobRequestBody) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectionJobRequestBody) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLeverageStorageSnapshots

`func (o *ProtectionJobRequestBody) GetLeverageStorageSnapshots() bool`

GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsOk

`func (o *ProtectionJobRequestBody) GetLeverageStorageSnapshotsOk() (*bool, bool)`

GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshots

`func (o *ProtectionJobRequestBody) SetLeverageStorageSnapshots(v bool)`

SetLeverageStorageSnapshots sets LeverageStorageSnapshots field to given value.

### HasLeverageStorageSnapshots

`func (o *ProtectionJobRequestBody) HasLeverageStorageSnapshots() bool`

HasLeverageStorageSnapshots returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsNil

`func (o *ProtectionJobRequestBody) SetLeverageStorageSnapshotsNil(b bool)`

 SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil

### UnsetLeverageStorageSnapshots
`func (o *ProtectionJobRequestBody) UnsetLeverageStorageSnapshots()`

UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
### GetLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJobRequestBody) GetLeverageStorageSnapshotsForHyperflex() bool`

GetLeverageStorageSnapshotsForHyperflex returns the LeverageStorageSnapshotsForHyperflex field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsForHyperflexOk

`func (o *ProtectionJobRequestBody) GetLeverageStorageSnapshotsForHyperflexOk() (*bool, bool)`

GetLeverageStorageSnapshotsForHyperflexOk returns a tuple with the LeverageStorageSnapshotsForHyperflex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJobRequestBody) SetLeverageStorageSnapshotsForHyperflex(v bool)`

SetLeverageStorageSnapshotsForHyperflex sets LeverageStorageSnapshotsForHyperflex field to given value.

### HasLeverageStorageSnapshotsForHyperflex

`func (o *ProtectionJobRequestBody) HasLeverageStorageSnapshotsForHyperflex() bool`

HasLeverageStorageSnapshotsForHyperflex returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsForHyperflexNil

`func (o *ProtectionJobRequestBody) SetLeverageStorageSnapshotsForHyperflexNil(b bool)`

 SetLeverageStorageSnapshotsForHyperflexNil sets the value for LeverageStorageSnapshotsForHyperflex to be an explicit nil

### UnsetLeverageStorageSnapshotsForHyperflex
`func (o *ProtectionJobRequestBody) UnsetLeverageStorageSnapshotsForHyperflex()`

UnsetLeverageStorageSnapshotsForHyperflex ensures that no value is present for LeverageStorageSnapshotsForHyperflex, not even an explicit nil
### GetName

`func (o *ProtectionJobRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionJobRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionJobRequestBody) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProtectionJobRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionJobRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceId

`func (o *ProtectionJobRequestBody) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *ProtectionJobRequestBody) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *ProtectionJobRequestBody) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *ProtectionJobRequestBody) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *ProtectionJobRequestBody) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *ProtectionJobRequestBody) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetPerformSourceSideDedup

`func (o *ProtectionJobRequestBody) GetPerformSourceSideDedup() bool`

GetPerformSourceSideDedup returns the PerformSourceSideDedup field if non-nil, zero value otherwise.

### GetPerformSourceSideDedupOk

`func (o *ProtectionJobRequestBody) GetPerformSourceSideDedupOk() (*bool, bool)`

GetPerformSourceSideDedupOk returns a tuple with the PerformSourceSideDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformSourceSideDedup

`func (o *ProtectionJobRequestBody) SetPerformSourceSideDedup(v bool)`

SetPerformSourceSideDedup sets PerformSourceSideDedup field to given value.

### HasPerformSourceSideDedup

`func (o *ProtectionJobRequestBody) HasPerformSourceSideDedup() bool`

HasPerformSourceSideDedup returns a boolean if a field has been set.

### SetPerformSourceSideDedupNil

`func (o *ProtectionJobRequestBody) SetPerformSourceSideDedupNil(b bool)`

 SetPerformSourceSideDedupNil sets the value for PerformSourceSideDedup to be an explicit nil

### UnsetPerformSourceSideDedup
`func (o *ProtectionJobRequestBody) UnsetPerformSourceSideDedup()`

UnsetPerformSourceSideDedup ensures that no value is present for PerformSourceSideDedup, not even an explicit nil
### GetPolicyId

`func (o *ProtectionJobRequestBody) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionJobRequestBody) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionJobRequestBody) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *ProtectionJobRequestBody) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionJobRequestBody) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPostBackupScript

`func (o *ProtectionJobRequestBody) GetPostBackupScript() BackupScript`

GetPostBackupScript returns the PostBackupScript field if non-nil, zero value otherwise.

### GetPostBackupScriptOk

`func (o *ProtectionJobRequestBody) GetPostBackupScriptOk() (*BackupScript, bool)`

GetPostBackupScriptOk returns a tuple with the PostBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBackupScript

`func (o *ProtectionJobRequestBody) SetPostBackupScript(v BackupScript)`

SetPostBackupScript sets PostBackupScript field to given value.

### HasPostBackupScript

`func (o *ProtectionJobRequestBody) HasPostBackupScript() bool`

HasPostBackupScript returns a boolean if a field has been set.

### SetPostBackupScriptNil

`func (o *ProtectionJobRequestBody) SetPostBackupScriptNil(b bool)`

 SetPostBackupScriptNil sets the value for PostBackupScript to be an explicit nil

### UnsetPostBackupScript
`func (o *ProtectionJobRequestBody) UnsetPostBackupScript()`

UnsetPostBackupScript ensures that no value is present for PostBackupScript, not even an explicit nil
### GetPreBackupScript

`func (o *ProtectionJobRequestBody) GetPreBackupScript() BackupScript`

GetPreBackupScript returns the PreBackupScript field if non-nil, zero value otherwise.

### GetPreBackupScriptOk

`func (o *ProtectionJobRequestBody) GetPreBackupScriptOk() (*BackupScript, bool)`

GetPreBackupScriptOk returns a tuple with the PreBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBackupScript

`func (o *ProtectionJobRequestBody) SetPreBackupScript(v BackupScript)`

SetPreBackupScript sets PreBackupScript field to given value.

### HasPreBackupScript

`func (o *ProtectionJobRequestBody) HasPreBackupScript() bool`

HasPreBackupScript returns a boolean if a field has been set.

### SetPreBackupScriptNil

`func (o *ProtectionJobRequestBody) SetPreBackupScriptNil(b bool)`

 SetPreBackupScriptNil sets the value for PreBackupScript to be an explicit nil

### UnsetPreBackupScript
`func (o *ProtectionJobRequestBody) UnsetPreBackupScript()`

UnsetPreBackupScript ensures that no value is present for PreBackupScript, not even an explicit nil
### GetPriority

`func (o *ProtectionJobRequestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProtectionJobRequestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProtectionJobRequestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProtectionJobRequestBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ProtectionJobRequestBody) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ProtectionJobRequestBody) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosType

`func (o *ProtectionJobRequestBody) GetQosType() string`

GetQosType returns the QosType field if non-nil, zero value otherwise.

### GetQosTypeOk

`func (o *ProtectionJobRequestBody) GetQosTypeOk() (*string, bool)`

GetQosTypeOk returns a tuple with the QosType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosType

`func (o *ProtectionJobRequestBody) SetQosType(v string)`

SetQosType sets QosType field to given value.

### HasQosType

`func (o *ProtectionJobRequestBody) HasQosType() bool`

HasQosType returns a boolean if a field has been set.

### SetQosTypeNil

`func (o *ProtectionJobRequestBody) SetQosTypeNil(b bool)`

 SetQosTypeNil sets the value for QosType to be an explicit nil

### UnsetQosType
`func (o *ProtectionJobRequestBody) UnsetQosType()`

UnsetQosType ensures that no value is present for QosType, not even an explicit nil
### GetQuiesce

`func (o *ProtectionJobRequestBody) GetQuiesce() bool`

GetQuiesce returns the Quiesce field if non-nil, zero value otherwise.

### GetQuiesceOk

`func (o *ProtectionJobRequestBody) GetQuiesceOk() (*bool, bool)`

GetQuiesceOk returns a tuple with the Quiesce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesce

`func (o *ProtectionJobRequestBody) SetQuiesce(v bool)`

SetQuiesce sets Quiesce field to given value.

### HasQuiesce

`func (o *ProtectionJobRequestBody) HasQuiesce() bool`

HasQuiesce returns a boolean if a field has been set.

### SetQuiesceNil

`func (o *ProtectionJobRequestBody) SetQuiesceNil(b bool)`

 SetQuiesceNil sets the value for Quiesce to be an explicit nil

### UnsetQuiesce
`func (o *ProtectionJobRequestBody) UnsetQuiesce()`

UnsetQuiesce ensures that no value is present for Quiesce, not even an explicit nil
### GetRemoteScript

`func (o *ProtectionJobRequestBody) GetRemoteScript() RemoteJobScript`

GetRemoteScript returns the RemoteScript field if non-nil, zero value otherwise.

### GetRemoteScriptOk

`func (o *ProtectionJobRequestBody) GetRemoteScriptOk() (*RemoteJobScript, bool)`

GetRemoteScriptOk returns a tuple with the RemoteScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteScript

`func (o *ProtectionJobRequestBody) SetRemoteScript(v RemoteJobScript)`

SetRemoteScript sets RemoteScript field to given value.

### HasRemoteScript

`func (o *ProtectionJobRequestBody) HasRemoteScript() bool`

HasRemoteScript returns a boolean if a field has been set.

### SetRemoteScriptNil

`func (o *ProtectionJobRequestBody) SetRemoteScriptNil(b bool)`

 SetRemoteScriptNil sets the value for RemoteScript to be an explicit nil

### UnsetRemoteScript
`func (o *ProtectionJobRequestBody) UnsetRemoteScript()`

UnsetRemoteScript ensures that no value is present for RemoteScript, not even an explicit nil
### GetRemoteViewName

`func (o *ProtectionJobRequestBody) GetRemoteViewName() string`

GetRemoteViewName returns the RemoteViewName field if non-nil, zero value otherwise.

### GetRemoteViewNameOk

`func (o *ProtectionJobRequestBody) GetRemoteViewNameOk() (*string, bool)`

GetRemoteViewNameOk returns a tuple with the RemoteViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteViewName

`func (o *ProtectionJobRequestBody) SetRemoteViewName(v string)`

SetRemoteViewName sets RemoteViewName field to given value.

### HasRemoteViewName

`func (o *ProtectionJobRequestBody) HasRemoteViewName() bool`

HasRemoteViewName returns a boolean if a field has been set.

### SetRemoteViewNameNil

`func (o *ProtectionJobRequestBody) SetRemoteViewNameNil(b bool)`

 SetRemoteViewNameNil sets the value for RemoteViewName to be an explicit nil

### UnsetRemoteViewName
`func (o *ProtectionJobRequestBody) UnsetRemoteViewName()`

UnsetRemoteViewName ensures that no value is present for RemoteViewName, not even an explicit nil
### GetSourceIds

`func (o *ProtectionJobRequestBody) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *ProtectionJobRequestBody) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *ProtectionJobRequestBody) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *ProtectionJobRequestBody) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *ProtectionJobRequestBody) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *ProtectionJobRequestBody) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil
### GetSourceSpecialParameters

`func (o *ProtectionJobRequestBody) GetSourceSpecialParameters() []SourceSpecialParameter`

GetSourceSpecialParameters returns the SourceSpecialParameters field if non-nil, zero value otherwise.

### GetSourceSpecialParametersOk

`func (o *ProtectionJobRequestBody) GetSourceSpecialParametersOk() (*[]SourceSpecialParameter, bool)`

GetSourceSpecialParametersOk returns a tuple with the SourceSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSpecialParameters

`func (o *ProtectionJobRequestBody) SetSourceSpecialParameters(v []SourceSpecialParameter)`

SetSourceSpecialParameters sets SourceSpecialParameters field to given value.


### SetSourceSpecialParametersNil

`func (o *ProtectionJobRequestBody) SetSourceSpecialParametersNil(b bool)`

 SetSourceSpecialParametersNil sets the value for SourceSpecialParameters to be an explicit nil

### UnsetSourceSpecialParameters
`func (o *ProtectionJobRequestBody) UnsetSourceSpecialParameters()`

UnsetSourceSpecialParameters ensures that no value is present for SourceSpecialParameters, not even an explicit nil
### GetStartTime

`func (o *ProtectionJobRequestBody) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProtectionJobRequestBody) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProtectionJobRequestBody) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProtectionJobRequestBody) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ProtectionJobRequestBody) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ProtectionJobRequestBody) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTimezone

`func (o *ProtectionJobRequestBody) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProtectionJobRequestBody) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProtectionJobRequestBody) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProtectionJobRequestBody) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ProtectionJobRequestBody) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ProtectionJobRequestBody) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetUserSpecifiedTags

`func (o *ProtectionJobRequestBody) GetUserSpecifiedTags() []string`

GetUserSpecifiedTags returns the UserSpecifiedTags field if non-nil, zero value otherwise.

### GetUserSpecifiedTagsOk

`func (o *ProtectionJobRequestBody) GetUserSpecifiedTagsOk() (*[]string, bool)`

GetUserSpecifiedTagsOk returns a tuple with the UserSpecifiedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSpecifiedTags

`func (o *ProtectionJobRequestBody) SetUserSpecifiedTags(v []string)`

SetUserSpecifiedTags sets UserSpecifiedTags field to given value.

### HasUserSpecifiedTags

`func (o *ProtectionJobRequestBody) HasUserSpecifiedTags() bool`

HasUserSpecifiedTags returns a boolean if a field has been set.

### SetUserSpecifiedTagsNil

`func (o *ProtectionJobRequestBody) SetUserSpecifiedTagsNil(b bool)`

 SetUserSpecifiedTagsNil sets the value for UserSpecifiedTags to be an explicit nil

### UnsetUserSpecifiedTags
`func (o *ProtectionJobRequestBody) UnsetUserSpecifiedTags()`

UnsetUserSpecifiedTags ensures that no value is present for UserSpecifiedTags, not even an explicit nil
### GetViewBoxId

`func (o *ProtectionJobRequestBody) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *ProtectionJobRequestBody) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *ProtectionJobRequestBody) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.


### SetViewBoxIdNil

`func (o *ProtectionJobRequestBody) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *ProtectionJobRequestBody) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *ProtectionJobRequestBody) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ProtectionJobRequestBody) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ProtectionJobRequestBody) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ProtectionJobRequestBody) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ProtectionJobRequestBody) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ProtectionJobRequestBody) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVmTagIds

`func (o *ProtectionJobRequestBody) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *ProtectionJobRequestBody) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *ProtectionJobRequestBody) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *ProtectionJobRequestBody) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *ProtectionJobRequestBody) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *ProtectionJobRequestBody) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


