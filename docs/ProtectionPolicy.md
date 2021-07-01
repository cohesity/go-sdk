# ProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutPeriods** | Pointer to [**[]BlackoutPeriod**](BlackoutPeriod.md) | Array of Blackout Periods.  If specified, this field defines black periods when new Job Runs are not started. If a Job Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | [optional] 
**CdpSchedulingPolicy** | Pointer to [**SchedulingPolicy**](SchedulingPolicy.md) |  | [optional] 
**CloudDeployPolicies** | Pointer to [**[]SnapshotCloudCopyPolicy**](SnapshotCloudCopyPolicy.md) | Array of Cloud Deploy Policies.  Specifies settings for copying Snapshots to Cloud. CloudDeploy target where backup snapshots may be converted and stored. It also defines the retention of copied Snapshots on the Cloud. | [optional] 
**DatalockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DatalockConfigLog** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DatalockConfigSystem** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies how many days to retain Snapshots on the Cohesity Cluster. | [optional] 
**DaysToKeepLog** | Pointer to **NullableInt64** | Specifies the number of days to retain log run if Log Schedule exists. | [optional] 
**DaysToKeepSystem** | Pointer to **NullableInt64** | Specifies the number of days to retain system backups made for bare metal recovery. This field is applicable if systemSchedulingPolicy is specified. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Protection Policy. | [optional] 
**ExtendedRetentionPolicies** | Pointer to [**[]ExtendedRetentionPolicy**](ExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 
**FullSchedulingPolicy** | Pointer to [**NullableSchedulingPolicy**](SchedulingPolicy.md) | Specifies the Full (no CBT) backup schedule of a Protection Job and how long Snapshots captured by this schedule are retained on the Cohesity Cluster. | [optional] 
**Id** | Pointer to **NullableString** | Specifies a unique Policy id assigned by the Cohesity Cluster. | [optional] 
**IncrementalSchedulingPolicy** | Pointer to [**NullableSchedulingPolicy**](SchedulingPolicy.md) | Specifies the CBT-based backup schedule of a Protection Job and how long Snapshots captured by this schedule are retained on the Cohesity Cluster. | [optional] 
**IsUsable** | Pointer to **NullableBool** | Specifies if the policy can be used to create a job. | [optional] 
**LastModifiedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time (in milliseconds) when the Protection Policy was last modified. | [optional] 
**LogSchedulingPolicy** | Pointer to [**SchedulingPolicy**](SchedulingPolicy.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Protection Policy. | [optional] 
**NumLinkedPolicies** | Pointer to **NullableInt64** | Species the number of policies linked to a global policy. | [optional] 
**ParentPolicyId** | Pointer to **NullableString** | Specifies the parent global policy Id. This must be specified when creating a policy from global policy template. | [optional] 
**Retries** | Pointer to **NullableInt32** | Specifies the number of times to retry capturing Snapshots before the Job Run fails. | [optional] 
**RetryIntervalMins** | Pointer to **NullableInt32** | Specifies the number of minutes before retrying a failed Protection Job. | [optional] 
**RpoPolicySettings** | Pointer to [**RpoPolicySettings**](RpoPolicySettings.md) |  | [optional] 
**SkipIntervalMins** | Pointer to **NullableInt32** | Specifies the period of time before skipping the execution of new Job Runs if an existing queued Job Run of the same Protection Job has not started. For example if this field is set to 30 minutes and a Job Run is scheduled to start at 5:00 AM every day but does not start due to conflicts (such as too many Jobs are running). If the new Job Run does not start by 5:30AM, the Cohesity Cluster will skip the new Job Run. If the original Job Run completes before 5:30AM the next day, a new Job Run is created and starts executing. This field is optional. | [optional] 
**SnapshotArchivalCopyPolicies** | Pointer to [**[]SnapshotArchivalCopyPolicy**](SnapshotArchivalCopyPolicy.md) | Array of External Targets.  Specifies settings for copying Snapshots to  Archival External Targets (such as AWS or Tape). It also defines the retention of copied Snapshots on an External Targets such as AWS and Tape. | [optional] 
**SnapshotReplicationCopyPolicies** | Pointer to [**[]SnapshotReplicationCopyPolicy**](SnapshotReplicationCopyPolicy.md) | Array of Remote Clusters.  Specifies settings for copying Snapshots to Remote Clusters. It also defines the retention of copied Snapshots on a Remote Cluster. | [optional] 
**SystemSchedulingPolicy** | Pointer to [**SchedulingPolicy**](SchedulingPolicy.md) |  | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies which organizations have been assigned this policy. This value is only populated for the cluster admin for now. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the protection policy. &#39;kRegular&#39; means a regular Protection Policy. &#39;kRPO&#39; means an RPO Protection Policy. | [optional] 
**WormRetentionType** | Pointer to **NullableString** | Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Jobs using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. This field is deprecated. Use DataLockConfig for incremental runs, DataLockConfigLog for log runs, DataLockConfigSystem for BMR runs, and DataLockConfig in extended retention and for copy targets config. deprecated: true &#39;kNone&#39; implies there is no WORM retention set. &#39;kCompliance&#39; implies WORM retention is set for compliance reason. &#39;kAdministrative&#39; implies WORM retention is set for administrative purposes. | [optional] 

## Methods

### NewProtectionPolicy

`func NewProtectionPolicy() *ProtectionPolicy`

NewProtectionPolicy instantiates a new ProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyWithDefaults

`func NewProtectionPolicyWithDefaults() *ProtectionPolicy`

NewProtectionPolicyWithDefaults instantiates a new ProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutPeriods

`func (o *ProtectionPolicy) GetBlackoutPeriods() []BlackoutPeriod`

GetBlackoutPeriods returns the BlackoutPeriods field if non-nil, zero value otherwise.

### GetBlackoutPeriodsOk

`func (o *ProtectionPolicy) GetBlackoutPeriodsOk() (*[]BlackoutPeriod, bool)`

GetBlackoutPeriodsOk returns a tuple with the BlackoutPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutPeriods

`func (o *ProtectionPolicy) SetBlackoutPeriods(v []BlackoutPeriod)`

SetBlackoutPeriods sets BlackoutPeriods field to given value.

### HasBlackoutPeriods

`func (o *ProtectionPolicy) HasBlackoutPeriods() bool`

HasBlackoutPeriods returns a boolean if a field has been set.

### SetBlackoutPeriodsNil

`func (o *ProtectionPolicy) SetBlackoutPeriodsNil(b bool)`

 SetBlackoutPeriodsNil sets the value for BlackoutPeriods to be an explicit nil

### UnsetBlackoutPeriods
`func (o *ProtectionPolicy) UnsetBlackoutPeriods()`

UnsetBlackoutPeriods ensures that no value is present for BlackoutPeriods, not even an explicit nil
### GetCdpSchedulingPolicy

`func (o *ProtectionPolicy) GetCdpSchedulingPolicy() SchedulingPolicy`

GetCdpSchedulingPolicy returns the CdpSchedulingPolicy field if non-nil, zero value otherwise.

### GetCdpSchedulingPolicyOk

`func (o *ProtectionPolicy) GetCdpSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetCdpSchedulingPolicyOk returns a tuple with the CdpSchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpSchedulingPolicy

`func (o *ProtectionPolicy) SetCdpSchedulingPolicy(v SchedulingPolicy)`

SetCdpSchedulingPolicy sets CdpSchedulingPolicy field to given value.

### HasCdpSchedulingPolicy

`func (o *ProtectionPolicy) HasCdpSchedulingPolicy() bool`

HasCdpSchedulingPolicy returns a boolean if a field has been set.

### GetCloudDeployPolicies

`func (o *ProtectionPolicy) GetCloudDeployPolicies() []SnapshotCloudCopyPolicy`

GetCloudDeployPolicies returns the CloudDeployPolicies field if non-nil, zero value otherwise.

### GetCloudDeployPoliciesOk

`func (o *ProtectionPolicy) GetCloudDeployPoliciesOk() (*[]SnapshotCloudCopyPolicy, bool)`

GetCloudDeployPoliciesOk returns a tuple with the CloudDeployPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployPolicies

`func (o *ProtectionPolicy) SetCloudDeployPolicies(v []SnapshotCloudCopyPolicy)`

SetCloudDeployPolicies sets CloudDeployPolicies field to given value.

### HasCloudDeployPolicies

`func (o *ProtectionPolicy) HasCloudDeployPolicies() bool`

HasCloudDeployPolicies returns a boolean if a field has been set.

### SetCloudDeployPoliciesNil

`func (o *ProtectionPolicy) SetCloudDeployPoliciesNil(b bool)`

 SetCloudDeployPoliciesNil sets the value for CloudDeployPolicies to be an explicit nil

### UnsetCloudDeployPolicies
`func (o *ProtectionPolicy) UnsetCloudDeployPolicies()`

UnsetCloudDeployPolicies ensures that no value is present for CloudDeployPolicies, not even an explicit nil
### GetDatalockConfig

`func (o *ProtectionPolicy) GetDatalockConfig() DataLockConfig`

GetDatalockConfig returns the DatalockConfig field if non-nil, zero value otherwise.

### GetDatalockConfigOk

`func (o *ProtectionPolicy) GetDatalockConfigOk() (*DataLockConfig, bool)`

GetDatalockConfigOk returns a tuple with the DatalockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfig

`func (o *ProtectionPolicy) SetDatalockConfig(v DataLockConfig)`

SetDatalockConfig sets DatalockConfig field to given value.

### HasDatalockConfig

`func (o *ProtectionPolicy) HasDatalockConfig() bool`

HasDatalockConfig returns a boolean if a field has been set.

### GetDatalockConfigLog

`func (o *ProtectionPolicy) GetDatalockConfigLog() DataLockConfig`

GetDatalockConfigLog returns the DatalockConfigLog field if non-nil, zero value otherwise.

### GetDatalockConfigLogOk

`func (o *ProtectionPolicy) GetDatalockConfigLogOk() (*DataLockConfig, bool)`

GetDatalockConfigLogOk returns a tuple with the DatalockConfigLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfigLog

`func (o *ProtectionPolicy) SetDatalockConfigLog(v DataLockConfig)`

SetDatalockConfigLog sets DatalockConfigLog field to given value.

### HasDatalockConfigLog

`func (o *ProtectionPolicy) HasDatalockConfigLog() bool`

HasDatalockConfigLog returns a boolean if a field has been set.

### GetDatalockConfigSystem

`func (o *ProtectionPolicy) GetDatalockConfigSystem() DataLockConfig`

GetDatalockConfigSystem returns the DatalockConfigSystem field if non-nil, zero value otherwise.

### GetDatalockConfigSystemOk

`func (o *ProtectionPolicy) GetDatalockConfigSystemOk() (*DataLockConfig, bool)`

GetDatalockConfigSystemOk returns a tuple with the DatalockConfigSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfigSystem

`func (o *ProtectionPolicy) SetDatalockConfigSystem(v DataLockConfig)`

SetDatalockConfigSystem sets DatalockConfigSystem field to given value.

### HasDatalockConfigSystem

`func (o *ProtectionPolicy) HasDatalockConfigSystem() bool`

HasDatalockConfigSystem returns a boolean if a field has been set.

### GetDaysToKeep

`func (o *ProtectionPolicy) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *ProtectionPolicy) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *ProtectionPolicy) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *ProtectionPolicy) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *ProtectionPolicy) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *ProtectionPolicy) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetDaysToKeepLog

`func (o *ProtectionPolicy) GetDaysToKeepLog() int64`

GetDaysToKeepLog returns the DaysToKeepLog field if non-nil, zero value otherwise.

### GetDaysToKeepLogOk

`func (o *ProtectionPolicy) GetDaysToKeepLogOk() (*int64, bool)`

GetDaysToKeepLogOk returns a tuple with the DaysToKeepLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeepLog

`func (o *ProtectionPolicy) SetDaysToKeepLog(v int64)`

SetDaysToKeepLog sets DaysToKeepLog field to given value.

### HasDaysToKeepLog

`func (o *ProtectionPolicy) HasDaysToKeepLog() bool`

HasDaysToKeepLog returns a boolean if a field has been set.

### SetDaysToKeepLogNil

`func (o *ProtectionPolicy) SetDaysToKeepLogNil(b bool)`

 SetDaysToKeepLogNil sets the value for DaysToKeepLog to be an explicit nil

### UnsetDaysToKeepLog
`func (o *ProtectionPolicy) UnsetDaysToKeepLog()`

UnsetDaysToKeepLog ensures that no value is present for DaysToKeepLog, not even an explicit nil
### GetDaysToKeepSystem

`func (o *ProtectionPolicy) GetDaysToKeepSystem() int64`

GetDaysToKeepSystem returns the DaysToKeepSystem field if non-nil, zero value otherwise.

### GetDaysToKeepSystemOk

`func (o *ProtectionPolicy) GetDaysToKeepSystemOk() (*int64, bool)`

GetDaysToKeepSystemOk returns a tuple with the DaysToKeepSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeepSystem

`func (o *ProtectionPolicy) SetDaysToKeepSystem(v int64)`

SetDaysToKeepSystem sets DaysToKeepSystem field to given value.

### HasDaysToKeepSystem

`func (o *ProtectionPolicy) HasDaysToKeepSystem() bool`

HasDaysToKeepSystem returns a boolean if a field has been set.

### SetDaysToKeepSystemNil

`func (o *ProtectionPolicy) SetDaysToKeepSystemNil(b bool)`

 SetDaysToKeepSystemNil sets the value for DaysToKeepSystem to be an explicit nil

### UnsetDaysToKeepSystem
`func (o *ProtectionPolicy) UnsetDaysToKeepSystem()`

UnsetDaysToKeepSystem ensures that no value is present for DaysToKeepSystem, not even an explicit nil
### GetDescription

`func (o *ProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetentionPolicies

`func (o *ProtectionPolicy) GetExtendedRetentionPolicies() []ExtendedRetentionPolicy`

GetExtendedRetentionPolicies returns the ExtendedRetentionPolicies field if non-nil, zero value otherwise.

### GetExtendedRetentionPoliciesOk

`func (o *ProtectionPolicy) GetExtendedRetentionPoliciesOk() (*[]ExtendedRetentionPolicy, bool)`

GetExtendedRetentionPoliciesOk returns a tuple with the ExtendedRetentionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetentionPolicies

`func (o *ProtectionPolicy) SetExtendedRetentionPolicies(v []ExtendedRetentionPolicy)`

SetExtendedRetentionPolicies sets ExtendedRetentionPolicies field to given value.

### HasExtendedRetentionPolicies

`func (o *ProtectionPolicy) HasExtendedRetentionPolicies() bool`

HasExtendedRetentionPolicies returns a boolean if a field has been set.

### SetExtendedRetentionPoliciesNil

`func (o *ProtectionPolicy) SetExtendedRetentionPoliciesNil(b bool)`

 SetExtendedRetentionPoliciesNil sets the value for ExtendedRetentionPolicies to be an explicit nil

### UnsetExtendedRetentionPolicies
`func (o *ProtectionPolicy) UnsetExtendedRetentionPolicies()`

UnsetExtendedRetentionPolicies ensures that no value is present for ExtendedRetentionPolicies, not even an explicit nil
### GetFullSchedulingPolicy

`func (o *ProtectionPolicy) GetFullSchedulingPolicy() SchedulingPolicy`

GetFullSchedulingPolicy returns the FullSchedulingPolicy field if non-nil, zero value otherwise.

### GetFullSchedulingPolicyOk

`func (o *ProtectionPolicy) GetFullSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetFullSchedulingPolicyOk returns a tuple with the FullSchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSchedulingPolicy

`func (o *ProtectionPolicy) SetFullSchedulingPolicy(v SchedulingPolicy)`

SetFullSchedulingPolicy sets FullSchedulingPolicy field to given value.

### HasFullSchedulingPolicy

`func (o *ProtectionPolicy) HasFullSchedulingPolicy() bool`

HasFullSchedulingPolicy returns a boolean if a field has been set.

### SetFullSchedulingPolicyNil

`func (o *ProtectionPolicy) SetFullSchedulingPolicyNil(b bool)`

 SetFullSchedulingPolicyNil sets the value for FullSchedulingPolicy to be an explicit nil

### UnsetFullSchedulingPolicy
`func (o *ProtectionPolicy) UnsetFullSchedulingPolicy()`

UnsetFullSchedulingPolicy ensures that no value is present for FullSchedulingPolicy, not even an explicit nil
### GetId

`func (o *ProtectionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncrementalSchedulingPolicy

`func (o *ProtectionPolicy) GetIncrementalSchedulingPolicy() SchedulingPolicy`

GetIncrementalSchedulingPolicy returns the IncrementalSchedulingPolicy field if non-nil, zero value otherwise.

### GetIncrementalSchedulingPolicyOk

`func (o *ProtectionPolicy) GetIncrementalSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetIncrementalSchedulingPolicyOk returns a tuple with the IncrementalSchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSchedulingPolicy

`func (o *ProtectionPolicy) SetIncrementalSchedulingPolicy(v SchedulingPolicy)`

SetIncrementalSchedulingPolicy sets IncrementalSchedulingPolicy field to given value.

### HasIncrementalSchedulingPolicy

`func (o *ProtectionPolicy) HasIncrementalSchedulingPolicy() bool`

HasIncrementalSchedulingPolicy returns a boolean if a field has been set.

### SetIncrementalSchedulingPolicyNil

`func (o *ProtectionPolicy) SetIncrementalSchedulingPolicyNil(b bool)`

 SetIncrementalSchedulingPolicyNil sets the value for IncrementalSchedulingPolicy to be an explicit nil

### UnsetIncrementalSchedulingPolicy
`func (o *ProtectionPolicy) UnsetIncrementalSchedulingPolicy()`

UnsetIncrementalSchedulingPolicy ensures that no value is present for IncrementalSchedulingPolicy, not even an explicit nil
### GetIsUsable

`func (o *ProtectionPolicy) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *ProtectionPolicy) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *ProtectionPolicy) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.

### HasIsUsable

`func (o *ProtectionPolicy) HasIsUsable() bool`

HasIsUsable returns a boolean if a field has been set.

### SetIsUsableNil

`func (o *ProtectionPolicy) SetIsUsableNil(b bool)`

 SetIsUsableNil sets the value for IsUsable to be an explicit nil

### UnsetIsUsable
`func (o *ProtectionPolicy) UnsetIsUsable()`

UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil
### GetLastModifiedTimeMsecs

`func (o *ProtectionPolicy) GetLastModifiedTimeMsecs() int64`

GetLastModifiedTimeMsecs returns the LastModifiedTimeMsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimeMsecsOk

`func (o *ProtectionPolicy) GetLastModifiedTimeMsecsOk() (*int64, bool)`

GetLastModifiedTimeMsecsOk returns a tuple with the LastModifiedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimeMsecs

`func (o *ProtectionPolicy) SetLastModifiedTimeMsecs(v int64)`

SetLastModifiedTimeMsecs sets LastModifiedTimeMsecs field to given value.

### HasLastModifiedTimeMsecs

`func (o *ProtectionPolicy) HasLastModifiedTimeMsecs() bool`

HasLastModifiedTimeMsecs returns a boolean if a field has been set.

### SetLastModifiedTimeMsecsNil

`func (o *ProtectionPolicy) SetLastModifiedTimeMsecsNil(b bool)`

 SetLastModifiedTimeMsecsNil sets the value for LastModifiedTimeMsecs to be an explicit nil

### UnsetLastModifiedTimeMsecs
`func (o *ProtectionPolicy) UnsetLastModifiedTimeMsecs()`

UnsetLastModifiedTimeMsecs ensures that no value is present for LastModifiedTimeMsecs, not even an explicit nil
### GetLogSchedulingPolicy

`func (o *ProtectionPolicy) GetLogSchedulingPolicy() SchedulingPolicy`

GetLogSchedulingPolicy returns the LogSchedulingPolicy field if non-nil, zero value otherwise.

### GetLogSchedulingPolicyOk

`func (o *ProtectionPolicy) GetLogSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetLogSchedulingPolicyOk returns a tuple with the LogSchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSchedulingPolicy

`func (o *ProtectionPolicy) SetLogSchedulingPolicy(v SchedulingPolicy)`

SetLogSchedulingPolicy sets LogSchedulingPolicy field to given value.

### HasLogSchedulingPolicy

`func (o *ProtectionPolicy) HasLogSchedulingPolicy() bool`

HasLogSchedulingPolicy returns a boolean if a field has been set.

### GetName

`func (o *ProtectionPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectionPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectionPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumLinkedPolicies

`func (o *ProtectionPolicy) GetNumLinkedPolicies() int64`

GetNumLinkedPolicies returns the NumLinkedPolicies field if non-nil, zero value otherwise.

### GetNumLinkedPoliciesOk

`func (o *ProtectionPolicy) GetNumLinkedPoliciesOk() (*int64, bool)`

GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLinkedPolicies

`func (o *ProtectionPolicy) SetNumLinkedPolicies(v int64)`

SetNumLinkedPolicies sets NumLinkedPolicies field to given value.

### HasNumLinkedPolicies

`func (o *ProtectionPolicy) HasNumLinkedPolicies() bool`

HasNumLinkedPolicies returns a boolean if a field has been set.

### SetNumLinkedPoliciesNil

`func (o *ProtectionPolicy) SetNumLinkedPoliciesNil(b bool)`

 SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil

### UnsetNumLinkedPolicies
`func (o *ProtectionPolicy) UnsetNumLinkedPolicies()`

UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
### GetParentPolicyId

`func (o *ProtectionPolicy) GetParentPolicyId() string`

GetParentPolicyId returns the ParentPolicyId field if non-nil, zero value otherwise.

### GetParentPolicyIdOk

`func (o *ProtectionPolicy) GetParentPolicyIdOk() (*string, bool)`

GetParentPolicyIdOk returns a tuple with the ParentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPolicyId

`func (o *ProtectionPolicy) SetParentPolicyId(v string)`

SetParentPolicyId sets ParentPolicyId field to given value.

### HasParentPolicyId

`func (o *ProtectionPolicy) HasParentPolicyId() bool`

HasParentPolicyId returns a boolean if a field has been set.

### SetParentPolicyIdNil

`func (o *ProtectionPolicy) SetParentPolicyIdNil(b bool)`

 SetParentPolicyIdNil sets the value for ParentPolicyId to be an explicit nil

### UnsetParentPolicyId
`func (o *ProtectionPolicy) UnsetParentPolicyId()`

UnsetParentPolicyId ensures that no value is present for ParentPolicyId, not even an explicit nil
### GetRetries

`func (o *ProtectionPolicy) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ProtectionPolicy) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ProtectionPolicy) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ProtectionPolicy) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *ProtectionPolicy) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *ProtectionPolicy) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRetryIntervalMins

`func (o *ProtectionPolicy) GetRetryIntervalMins() int32`

GetRetryIntervalMins returns the RetryIntervalMins field if non-nil, zero value otherwise.

### GetRetryIntervalMinsOk

`func (o *ProtectionPolicy) GetRetryIntervalMinsOk() (*int32, bool)`

GetRetryIntervalMinsOk returns a tuple with the RetryIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryIntervalMins

`func (o *ProtectionPolicy) SetRetryIntervalMins(v int32)`

SetRetryIntervalMins sets RetryIntervalMins field to given value.

### HasRetryIntervalMins

`func (o *ProtectionPolicy) HasRetryIntervalMins() bool`

HasRetryIntervalMins returns a boolean if a field has been set.

### SetRetryIntervalMinsNil

`func (o *ProtectionPolicy) SetRetryIntervalMinsNil(b bool)`

 SetRetryIntervalMinsNil sets the value for RetryIntervalMins to be an explicit nil

### UnsetRetryIntervalMins
`func (o *ProtectionPolicy) UnsetRetryIntervalMins()`

UnsetRetryIntervalMins ensures that no value is present for RetryIntervalMins, not even an explicit nil
### GetRpoPolicySettings

`func (o *ProtectionPolicy) GetRpoPolicySettings() RpoPolicySettings`

GetRpoPolicySettings returns the RpoPolicySettings field if non-nil, zero value otherwise.

### GetRpoPolicySettingsOk

`func (o *ProtectionPolicy) GetRpoPolicySettingsOk() (*RpoPolicySettings, bool)`

GetRpoPolicySettingsOk returns a tuple with the RpoPolicySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicySettings

`func (o *ProtectionPolicy) SetRpoPolicySettings(v RpoPolicySettings)`

SetRpoPolicySettings sets RpoPolicySettings field to given value.

### HasRpoPolicySettings

`func (o *ProtectionPolicy) HasRpoPolicySettings() bool`

HasRpoPolicySettings returns a boolean if a field has been set.

### GetSkipIntervalMins

`func (o *ProtectionPolicy) GetSkipIntervalMins() int32`

GetSkipIntervalMins returns the SkipIntervalMins field if non-nil, zero value otherwise.

### GetSkipIntervalMinsOk

`func (o *ProtectionPolicy) GetSkipIntervalMinsOk() (*int32, bool)`

GetSkipIntervalMinsOk returns a tuple with the SkipIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntervalMins

`func (o *ProtectionPolicy) SetSkipIntervalMins(v int32)`

SetSkipIntervalMins sets SkipIntervalMins field to given value.

### HasSkipIntervalMins

`func (o *ProtectionPolicy) HasSkipIntervalMins() bool`

HasSkipIntervalMins returns a boolean if a field has been set.

### SetSkipIntervalMinsNil

`func (o *ProtectionPolicy) SetSkipIntervalMinsNil(b bool)`

 SetSkipIntervalMinsNil sets the value for SkipIntervalMins to be an explicit nil

### UnsetSkipIntervalMins
`func (o *ProtectionPolicy) UnsetSkipIntervalMins()`

UnsetSkipIntervalMins ensures that no value is present for SkipIntervalMins, not even an explicit nil
### GetSnapshotArchivalCopyPolicies

`func (o *ProtectionPolicy) GetSnapshotArchivalCopyPolicies() []SnapshotArchivalCopyPolicy`

GetSnapshotArchivalCopyPolicies returns the SnapshotArchivalCopyPolicies field if non-nil, zero value otherwise.

### GetSnapshotArchivalCopyPoliciesOk

`func (o *ProtectionPolicy) GetSnapshotArchivalCopyPoliciesOk() (*[]SnapshotArchivalCopyPolicy, bool)`

GetSnapshotArchivalCopyPoliciesOk returns a tuple with the SnapshotArchivalCopyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotArchivalCopyPolicies

`func (o *ProtectionPolicy) SetSnapshotArchivalCopyPolicies(v []SnapshotArchivalCopyPolicy)`

SetSnapshotArchivalCopyPolicies sets SnapshotArchivalCopyPolicies field to given value.

### HasSnapshotArchivalCopyPolicies

`func (o *ProtectionPolicy) HasSnapshotArchivalCopyPolicies() bool`

HasSnapshotArchivalCopyPolicies returns a boolean if a field has been set.

### SetSnapshotArchivalCopyPoliciesNil

`func (o *ProtectionPolicy) SetSnapshotArchivalCopyPoliciesNil(b bool)`

 SetSnapshotArchivalCopyPoliciesNil sets the value for SnapshotArchivalCopyPolicies to be an explicit nil

### UnsetSnapshotArchivalCopyPolicies
`func (o *ProtectionPolicy) UnsetSnapshotArchivalCopyPolicies()`

UnsetSnapshotArchivalCopyPolicies ensures that no value is present for SnapshotArchivalCopyPolicies, not even an explicit nil
### GetSnapshotReplicationCopyPolicies

`func (o *ProtectionPolicy) GetSnapshotReplicationCopyPolicies() []SnapshotReplicationCopyPolicy`

GetSnapshotReplicationCopyPolicies returns the SnapshotReplicationCopyPolicies field if non-nil, zero value otherwise.

### GetSnapshotReplicationCopyPoliciesOk

`func (o *ProtectionPolicy) GetSnapshotReplicationCopyPoliciesOk() (*[]SnapshotReplicationCopyPolicy, bool)`

GetSnapshotReplicationCopyPoliciesOk returns a tuple with the SnapshotReplicationCopyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationCopyPolicies

`func (o *ProtectionPolicy) SetSnapshotReplicationCopyPolicies(v []SnapshotReplicationCopyPolicy)`

SetSnapshotReplicationCopyPolicies sets SnapshotReplicationCopyPolicies field to given value.

### HasSnapshotReplicationCopyPolicies

`func (o *ProtectionPolicy) HasSnapshotReplicationCopyPolicies() bool`

HasSnapshotReplicationCopyPolicies returns a boolean if a field has been set.

### SetSnapshotReplicationCopyPoliciesNil

`func (o *ProtectionPolicy) SetSnapshotReplicationCopyPoliciesNil(b bool)`

 SetSnapshotReplicationCopyPoliciesNil sets the value for SnapshotReplicationCopyPolicies to be an explicit nil

### UnsetSnapshotReplicationCopyPolicies
`func (o *ProtectionPolicy) UnsetSnapshotReplicationCopyPolicies()`

UnsetSnapshotReplicationCopyPolicies ensures that no value is present for SnapshotReplicationCopyPolicies, not even an explicit nil
### GetSystemSchedulingPolicy

`func (o *ProtectionPolicy) GetSystemSchedulingPolicy() SchedulingPolicy`

GetSystemSchedulingPolicy returns the SystemSchedulingPolicy field if non-nil, zero value otherwise.

### GetSystemSchedulingPolicyOk

`func (o *ProtectionPolicy) GetSystemSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetSystemSchedulingPolicyOk returns a tuple with the SystemSchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSchedulingPolicy

`func (o *ProtectionPolicy) SetSystemSchedulingPolicy(v SchedulingPolicy)`

SetSystemSchedulingPolicy sets SystemSchedulingPolicy field to given value.

### HasSystemSchedulingPolicy

`func (o *ProtectionPolicy) HasSystemSchedulingPolicy() bool`

HasSystemSchedulingPolicy returns a boolean if a field has been set.

### GetTenantIds

`func (o *ProtectionPolicy) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *ProtectionPolicy) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *ProtectionPolicy) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *ProtectionPolicy) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *ProtectionPolicy) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *ProtectionPolicy) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetType

`func (o *ProtectionPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtectionPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtectionPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtectionPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ProtectionPolicy) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProtectionPolicy) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWormRetentionType

`func (o *ProtectionPolicy) GetWormRetentionType() string`

GetWormRetentionType returns the WormRetentionType field if non-nil, zero value otherwise.

### GetWormRetentionTypeOk

`func (o *ProtectionPolicy) GetWormRetentionTypeOk() (*string, bool)`

GetWormRetentionTypeOk returns a tuple with the WormRetentionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormRetentionType

`func (o *ProtectionPolicy) SetWormRetentionType(v string)`

SetWormRetentionType sets WormRetentionType field to given value.

### HasWormRetentionType

`func (o *ProtectionPolicy) HasWormRetentionType() bool`

HasWormRetentionType returns a boolean if a field has been set.

### SetWormRetentionTypeNil

`func (o *ProtectionPolicy) SetWormRetentionTypeNil(b bool)`

 SetWormRetentionTypeNil sets the value for WormRetentionType to be an explicit nil

### UnsetWormRetentionType
`func (o *ProtectionPolicy) UnsetWormRetentionType()`

UnsetWormRetentionType ensures that no value is present for WormRetentionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


