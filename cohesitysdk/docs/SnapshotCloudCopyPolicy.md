# SnapshotCloudCopyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specified the Id for a snapshot copy policy. This is generated when the policy is created. | [optional] 
**CopyPartial** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Job Run or the first partially successful Job Run occurring at the start of the replication schedule. If true, Snapshots are copied from the first Job Run occurring at the start of the replication schedule, even if first Job Run was not completely successful i.e. Snapshots were not captured for all Objects in the Job. If false, Snapshots are copied from the first Job Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Job were successfully captured. | [optional] 
**DatalockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies the number of days to retain copied Snapshots on the target. | [optional] 
**Multiplier** | Pointer to **NullableInt32** | Specifies a factor to multiply the periodicity by, to determine the copy schedule. For example if set to 2 and the periodicity is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied. | [optional] 
**Periodicity** | Pointer to **NullableString** | Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. &#39;kEvery&#39; means that the Snapshot copy occurs after the number of Job Runs equals the number specified in the multiplier. &#39;kHour&#39; means that the Snapshot copy occurs hourly at the frequency set in the multiplier, for example if multiplier is 2, the copy occurs every 2 hours. &#39;kDay&#39; means that the Snapshot copy occurs daily at the frequency set in the multiplier. &#39;kWeek&#39; means that the Snapshot copy occurs weekly at the frequency set in the multiplier. &#39;kMonth&#39; means that the Snapshot copy occurs monthly at the frequency set in the multiplier. &#39;kYear&#39; means that the Snapshot copy occurs yearly at the frequency set in the multiplier. | [optional] 
**Target** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 

## Methods

### NewSnapshotCloudCopyPolicy

`func NewSnapshotCloudCopyPolicy() *SnapshotCloudCopyPolicy`

NewSnapshotCloudCopyPolicy instantiates a new SnapshotCloudCopyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCloudCopyPolicyWithDefaults

`func NewSnapshotCloudCopyPolicyWithDefaults() *SnapshotCloudCopyPolicy`

NewSnapshotCloudCopyPolicyWithDefaults instantiates a new SnapshotCloudCopyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotCloudCopyPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotCloudCopyPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotCloudCopyPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotCloudCopyPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SnapshotCloudCopyPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SnapshotCloudCopyPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCopyPartial

`func (o *SnapshotCloudCopyPolicy) GetCopyPartial() bool`

GetCopyPartial returns the CopyPartial field if non-nil, zero value otherwise.

### GetCopyPartialOk

`func (o *SnapshotCloudCopyPolicy) GetCopyPartialOk() (*bool, bool)`

GetCopyPartialOk returns a tuple with the CopyPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPartial

`func (o *SnapshotCloudCopyPolicy) SetCopyPartial(v bool)`

SetCopyPartial sets CopyPartial field to given value.

### HasCopyPartial

`func (o *SnapshotCloudCopyPolicy) HasCopyPartial() bool`

HasCopyPartial returns a boolean if a field has been set.

### SetCopyPartialNil

`func (o *SnapshotCloudCopyPolicy) SetCopyPartialNil(b bool)`

 SetCopyPartialNil sets the value for CopyPartial to be an explicit nil

### UnsetCopyPartial
`func (o *SnapshotCloudCopyPolicy) UnsetCopyPartial()`

UnsetCopyPartial ensures that no value is present for CopyPartial, not even an explicit nil
### GetDatalockConfig

`func (o *SnapshotCloudCopyPolicy) GetDatalockConfig() DataLockConfig`

GetDatalockConfig returns the DatalockConfig field if non-nil, zero value otherwise.

### GetDatalockConfigOk

`func (o *SnapshotCloudCopyPolicy) GetDatalockConfigOk() (*DataLockConfig, bool)`

GetDatalockConfigOk returns a tuple with the DatalockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfig

`func (o *SnapshotCloudCopyPolicy) SetDatalockConfig(v DataLockConfig)`

SetDatalockConfig sets DatalockConfig field to given value.

### HasDatalockConfig

`func (o *SnapshotCloudCopyPolicy) HasDatalockConfig() bool`

HasDatalockConfig returns a boolean if a field has been set.

### GetDaysToKeep

`func (o *SnapshotCloudCopyPolicy) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *SnapshotCloudCopyPolicy) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *SnapshotCloudCopyPolicy) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *SnapshotCloudCopyPolicy) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *SnapshotCloudCopyPolicy) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *SnapshotCloudCopyPolicy) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetMultiplier

`func (o *SnapshotCloudCopyPolicy) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *SnapshotCloudCopyPolicy) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *SnapshotCloudCopyPolicy) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *SnapshotCloudCopyPolicy) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *SnapshotCloudCopyPolicy) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *SnapshotCloudCopyPolicy) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
### GetPeriodicity

`func (o *SnapshotCloudCopyPolicy) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *SnapshotCloudCopyPolicy) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *SnapshotCloudCopyPolicy) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *SnapshotCloudCopyPolicy) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *SnapshotCloudCopyPolicy) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *SnapshotCloudCopyPolicy) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil
### GetTarget

`func (o *SnapshotCloudCopyPolicy) GetTarget() CloudDeployTargetDetails`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SnapshotCloudCopyPolicy) GetTargetOk() (*CloudDeployTargetDetails, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SnapshotCloudCopyPolicy) SetTarget(v CloudDeployTargetDetails)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SnapshotCloudCopyPolicy) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


