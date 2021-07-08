# ExtendedRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specified the Id for a snapshot copy policy. This is generated when the policy is created. | [optional] 
**BackupRunType** | Pointer to **NullableString** | The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is kFull. &#39;kRegular&#39; indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time. | [optional] 
**DatalockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies the number of days to retain copied Snapshots on the target. | [optional] 
**Multiplier** | Pointer to **NullableInt32** | Specifies a factor to multiply the periodicity by, to determine the copy schedule. For example if set to 2 and the periodicity is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied. | [optional] 
**Periodicity** | Pointer to **NullableString** | Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. &#39;kEvery&#39; means that the Snapshot copy occurs after the number of Job Runs equals the number specified in the multiplier. &#39;kHour&#39; means that the Snapshot copy occurs hourly at the frequency set in the multiplier, for example if multiplier is 2, the copy occurs every 2 hours. &#39;kDay&#39; means that the Snapshot copy occurs daily at the frequency set in the multiplier. &#39;kWeek&#39; means that the Snapshot copy occurs weekly at the frequency set in the multiplier. &#39;kMonth&#39; means that the Snapshot copy occurs monthly at the frequency set in the multiplier. &#39;kYear&#39; means that the Snapshot copy occurs yearly at the frequency set in the multiplier. | [optional] 

## Methods

### NewExtendedRetentionPolicy

`func NewExtendedRetentionPolicy() *ExtendedRetentionPolicy`

NewExtendedRetentionPolicy instantiates a new ExtendedRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedRetentionPolicyWithDefaults

`func NewExtendedRetentionPolicyWithDefaults() *ExtendedRetentionPolicy`

NewExtendedRetentionPolicyWithDefaults instantiates a new ExtendedRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedRetentionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedRetentionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedRetentionPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedRetentionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedRetentionPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedRetentionPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetBackupRunType

`func (o *ExtendedRetentionPolicy) GetBackupRunType() string`

GetBackupRunType returns the BackupRunType field if non-nil, zero value otherwise.

### GetBackupRunTypeOk

`func (o *ExtendedRetentionPolicy) GetBackupRunTypeOk() (*string, bool)`

GetBackupRunTypeOk returns a tuple with the BackupRunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunType

`func (o *ExtendedRetentionPolicy) SetBackupRunType(v string)`

SetBackupRunType sets BackupRunType field to given value.

### HasBackupRunType

`func (o *ExtendedRetentionPolicy) HasBackupRunType() bool`

HasBackupRunType returns a boolean if a field has been set.

### SetBackupRunTypeNil

`func (o *ExtendedRetentionPolicy) SetBackupRunTypeNil(b bool)`

 SetBackupRunTypeNil sets the value for BackupRunType to be an explicit nil

### UnsetBackupRunType
`func (o *ExtendedRetentionPolicy) UnsetBackupRunType()`

UnsetBackupRunType ensures that no value is present for BackupRunType, not even an explicit nil
### GetDatalockConfig

`func (o *ExtendedRetentionPolicy) GetDatalockConfig() DataLockConfig`

GetDatalockConfig returns the DatalockConfig field if non-nil, zero value otherwise.

### GetDatalockConfigOk

`func (o *ExtendedRetentionPolicy) GetDatalockConfigOk() (*DataLockConfig, bool)`

GetDatalockConfigOk returns a tuple with the DatalockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalockConfig

`func (o *ExtendedRetentionPolicy) SetDatalockConfig(v DataLockConfig)`

SetDatalockConfig sets DatalockConfig field to given value.

### HasDatalockConfig

`func (o *ExtendedRetentionPolicy) HasDatalockConfig() bool`

HasDatalockConfig returns a boolean if a field has been set.

### GetDaysToKeep

`func (o *ExtendedRetentionPolicy) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *ExtendedRetentionPolicy) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *ExtendedRetentionPolicy) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *ExtendedRetentionPolicy) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *ExtendedRetentionPolicy) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *ExtendedRetentionPolicy) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetMultiplier

`func (o *ExtendedRetentionPolicy) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *ExtendedRetentionPolicy) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *ExtendedRetentionPolicy) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *ExtendedRetentionPolicy) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *ExtendedRetentionPolicy) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *ExtendedRetentionPolicy) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
### GetPeriodicity

`func (o *ExtendedRetentionPolicy) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *ExtendedRetentionPolicy) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *ExtendedRetentionPolicy) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.

### HasPeriodicity

`func (o *ExtendedRetentionPolicy) HasPeriodicity() bool`

HasPeriodicity returns a boolean if a field has been set.

### SetPeriodicityNil

`func (o *ExtendedRetentionPolicy) SetPeriodicityNil(b bool)`

 SetPeriodicityNil sets the value for Periodicity to be an explicit nil

### UnsetPeriodicity
`func (o *ExtendedRetentionPolicy) UnsetPeriodicity()`

UnsetPeriodicity ensures that no value is present for Periodicity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


