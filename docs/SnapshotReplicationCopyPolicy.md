# SnapshotReplicationCopyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specified the Id for a snapshot copy policy. This is generated when the policy is created. | [optional] 
**CloudTarget** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 
**CopyPartial** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Job Run or the first partially successful Job Run occurring at the start of the replication schedule. If true, Snapshots are copied from the first Job Run occurring at the start of the replication schedule, even if first Job Run was not completely successful i.e. Snapshots were not captured for all Objects in the Job. If false, Snapshots are copied from the first Job Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Job were successfully captured. | [optional] 
**DatalockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies the number of days to retain copied Snapshots on the target. | [optional] 
**Multiplier** | Pointer to **NullableInt32** | Specifies a factor to multiply the periodicity by, to determine the copy schedule. For example if set to 2 and the periodicity is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied. | [optional] 
**Periodicity** | Pointer to **NullableString** | Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. &#39;kEvery&#39; means that the Snapshot copy occurs after the number of Job Runs equals the number specified in the multiplier. &#39;kHour&#39; means that the Snapshot copy occurs hourly at the frequency set in the multiplier, for example if multiplier is 2, the copy occurs every 2 hours. &#39;kDay&#39; means that the Snapshot copy occurs daily at the frequency set in the multiplier. &#39;kWeek&#39; means that the Snapshot copy occurs weekly at the frequency set in the multiplier. &#39;kMonth&#39; means that the Snapshot copy occurs monthly at the frequency set in the multiplier. &#39;kYear&#39; means that the Snapshot copy occurs yearly at the frequency set in the multiplier. | [optional] 
**Target** | Pointer to [**NullableReplicationTargetSettings**](ReplicationTargetSettings.md) | Specifies the replication target to copy the Snapshots to. | [optional] 

## Methods

### NewSnapshotReplicationCopyPolicy

`func NewSnapshotReplicationCopyPolicy() *SnapshotReplicationCopyPolicy`

NewSnapshotReplicationCopyPolicy instantiates a new SnapshotReplicationCopyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotReplicationCopyPolicyWithDefaults

`func NewSnapshotReplicationCopyPolicyWithDefaults() *SnapshotReplicationCopyPolicy`

NewSnapshotReplicationCopyPolicyWithDefaults instantiates a new SnapshotReplicationCopyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotReplicationCopyPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotReplicationCopyPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotReplicationCopyPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotReplicationCopyPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SnapshotReplicationCopyPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SnapshotReplicationCopyPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCloudTarget

`func (o *SnapshotReplicationCopyPolicy) GetCloudTarget() CloudDeployTargetDetails`

GetCloudTarget returns the CloudTarget field if non-nil, zero value otherwise.

### GetCloudTargetOk

`func (o *SnapshotReplicationCopyPolicy) GetCloudTargetOk() (*CloudDeployTargetDetails, bool)`

GetCloudTargetOk returns a tuple with the CloudTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTarget

`func (o *SnapshotReplicationCopyPolicy) SetCloudTarget(v CloudDeployTargetDetails)`

SetCloudTarget sets CloudTarget field to given value.

### HasCloudTarget

`func (o *SnapshotReplicationCopyPolicy) HasCloudTarget() bool`

HasCloudTarget returns a boolean if a field has been set.

### GetCopyPartial

`func (o *SnapshotReplicationCopyPolicy) GetCopyPartial() bool`

GetCopyPartial returns the CopyPartial field if non-nil, zero value otherwise.

### GetCopyPartialOk

`func (o *SnapshotReplicationCopyPolicy) GetCopyPartialOk() (*bool, bool)`

GetCopyPartialOk returns a tuple with the CopyPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPartial

`func (o *SnapshotReplicationCopyPolicy) SetCopyPartial(v bool)`

SetCopyPartial sets CopyPartial field to given value.

### HasCopyPartial

`func (o *SnapshotReplicationCopyPolicy) HasCopyPartial() bool`

HasCopyPartial returns a boolean if a field has been set.

### SetCopyPartialNil

`func (o *SnapshotReplicationCopyPolicy) SetCopyPartialNil(b bool)`

 SetCopyPartialNil sets the value for CopyPartial to be an explicit nil

### UnsetCopyPartial
`func (o *SnapshotReplicationCopyPolicy) UnsetCopyPartial()`

UnsetCopyPartial ensures that no value is present for CopyPartial, not even an explicit nil
### GetDatalockConfig

`func (o *SnapshotReplicationCopyPolicy) GetDatalockConfig() DataLockConfig`

GetDatalockConfig returns the DatalockConfig field if non-nil, zero value otherwise.

### GetDatalockConfigOk

`func (o *SnapshotReplicationCopyPolicy) GetDatalockConfigOk() (*DataLockConfig, bool)`

GetDatalockConfigOk returns a tuple with the DatalockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfig

`func (o *SnapshotReplicationCopyPolicy) SetDatalockConfig(v DataLockConfig)`

SetDatalockConfig sets DatalockConfig field to given value.

### HasDatalockConfig

`func (o *SnapshotReplicationCopyPolicy) HasDatalockConfig() bool`

HasDatalockConfig returns a boolean if a field has been set.

### GetDaysToKeep

`func (o *SnapshotReplicationCopyPolicy) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *SnapshotReplicationCopyPolicy) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *SnapshotReplicationCopyPolicy) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *SnapshotReplicationCopyPolicy) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *SnapshotReplicationCopyPolicy) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *SnapshotReplicationCopyPolicy) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetMultiplier

`func (o *SnapshotReplicationCopyPolicy) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *SnapshotReplicationCopyPolicy) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *SnapshotReplicationCopyPolicy) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *SnapshotReplicationCopyPolicy) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *SnapshotReplicationCopyPolicy) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *SnapshotReplicationCopyPolicy) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
### GetPeriodicity

`func (o *SnapshotReplicationCopyPolicy) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *SnapshotReplicationCopyPolicy) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *SnapshotReplicationCopyPolicy) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *SnapshotReplicationCopyPolicy) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *SnapshotReplicationCopyPolicy) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *SnapshotReplicationCopyPolicy) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil
### GetTarget

`func (o *SnapshotReplicationCopyPolicy) GetTarget() ReplicationTargetSettings`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SnapshotReplicationCopyPolicy) GetTargetOk() (*ReplicationTargetSettings, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SnapshotReplicationCopyPolicy) SetTarget(v ReplicationTargetSettings)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SnapshotReplicationCopyPolicy) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *SnapshotReplicationCopyPolicy) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *SnapshotReplicationCopyPolicy) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


