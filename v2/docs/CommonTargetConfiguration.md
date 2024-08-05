# CommonTargetConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRunType** | Pointer to **NullableString** | Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full. | [optional] 
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated. | [optional] 
**CopyOnRunSuccess** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. &lt;br&gt; If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. &lt;br&gt; If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group. | [optional] 
**LogRetention** | Pointer to [**LogRetention**](LogRetention.md) |  | [optional] 
**Retention** | [**Retention**](Retention.md) |  | 
**RunTimeouts** | Pointer to [**[]CancellationTimeoutParams**](CancellationTimeoutParams.md) | Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.). | [optional] 
**Schedule** | [**TargetSchedule**](TargetSchedule.md) |  | 

## Methods

### NewCommonTargetConfiguration

`func NewCommonTargetConfiguration(retention Retention, schedule TargetSchedule, ) *CommonTargetConfiguration`

NewCommonTargetConfiguration instantiates a new CommonTargetConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTargetConfigurationWithDefaults

`func NewCommonTargetConfigurationWithDefaults() *CommonTargetConfiguration`

NewCommonTargetConfigurationWithDefaults instantiates a new CommonTargetConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRunType

`func (o *CommonTargetConfiguration) GetBackupRunType() string`

GetBackupRunType returns the BackupRunType field if non-nil, zero value otherwise.

### GetBackupRunTypeOk

`func (o *CommonTargetConfiguration) GetBackupRunTypeOk() (*string, bool)`

GetBackupRunTypeOk returns a tuple with the BackupRunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunType

`func (o *CommonTargetConfiguration) SetBackupRunType(v string)`

SetBackupRunType sets BackupRunType field to given value.

### HasBackupRunType

`func (o *CommonTargetConfiguration) HasBackupRunType() bool`

HasBackupRunType returns a boolean if a field has been set.

### SetBackupRunTypeNil

`func (o *CommonTargetConfiguration) SetBackupRunTypeNil(b bool)`

 SetBackupRunTypeNil sets the value for BackupRunType to be an explicit nil

### UnsetBackupRunType
`func (o *CommonTargetConfiguration) UnsetBackupRunType()`

UnsetBackupRunType ensures that no value is present for BackupRunType, not even an explicit nil
### GetConfigId

`func (o *CommonTargetConfiguration) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *CommonTargetConfiguration) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *CommonTargetConfiguration) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *CommonTargetConfiguration) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *CommonTargetConfiguration) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *CommonTargetConfiguration) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetCopyOnRunSuccess

`func (o *CommonTargetConfiguration) GetCopyOnRunSuccess() bool`

GetCopyOnRunSuccess returns the CopyOnRunSuccess field if non-nil, zero value otherwise.

### GetCopyOnRunSuccessOk

`func (o *CommonTargetConfiguration) GetCopyOnRunSuccessOk() (*bool, bool)`

GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnRunSuccess

`func (o *CommonTargetConfiguration) SetCopyOnRunSuccess(v bool)`

SetCopyOnRunSuccess sets CopyOnRunSuccess field to given value.

### HasCopyOnRunSuccess

`func (o *CommonTargetConfiguration) HasCopyOnRunSuccess() bool`

HasCopyOnRunSuccess returns a boolean if a field has been set.

### SetCopyOnRunSuccessNil

`func (o *CommonTargetConfiguration) SetCopyOnRunSuccessNil(b bool)`

 SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil

### UnsetCopyOnRunSuccess
`func (o *CommonTargetConfiguration) UnsetCopyOnRunSuccess()`

UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
### GetLogRetention

`func (o *CommonTargetConfiguration) GetLogRetention() LogRetention`

GetLogRetention returns the LogRetention field if non-nil, zero value otherwise.

### GetLogRetentionOk

`func (o *CommonTargetConfiguration) GetLogRetentionOk() (*LogRetention, bool)`

GetLogRetentionOk returns a tuple with the LogRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetention

`func (o *CommonTargetConfiguration) SetLogRetention(v LogRetention)`

SetLogRetention sets LogRetention field to given value.

### HasLogRetention

`func (o *CommonTargetConfiguration) HasLogRetention() bool`

HasLogRetention returns a boolean if a field has been set.

### GetRetention

`func (o *CommonTargetConfiguration) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CommonTargetConfiguration) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CommonTargetConfiguration) SetRetention(v Retention)`

SetRetention sets Retention field to given value.


### GetRunTimeouts

`func (o *CommonTargetConfiguration) GetRunTimeouts() []CancellationTimeoutParams`

GetRunTimeouts returns the RunTimeouts field if non-nil, zero value otherwise.

### GetRunTimeoutsOk

`func (o *CommonTargetConfiguration) GetRunTimeoutsOk() (*[]CancellationTimeoutParams, bool)`

GetRunTimeoutsOk returns a tuple with the RunTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeouts

`func (o *CommonTargetConfiguration) SetRunTimeouts(v []CancellationTimeoutParams)`

SetRunTimeouts sets RunTimeouts field to given value.

### HasRunTimeouts

`func (o *CommonTargetConfiguration) HasRunTimeouts() bool`

HasRunTimeouts returns a boolean if a field has been set.

### SetRunTimeoutsNil

`func (o *CommonTargetConfiguration) SetRunTimeoutsNil(b bool)`

 SetRunTimeoutsNil sets the value for RunTimeouts to be an explicit nil

### UnsetRunTimeouts
`func (o *CommonTargetConfiguration) UnsetRunTimeouts()`

UnsetRunTimeouts ensures that no value is present for RunTimeouts, not even an explicit nil
### GetSchedule

`func (o *CommonTargetConfiguration) GetSchedule() TargetSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CommonTargetConfiguration) GetScheduleOk() (*TargetSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CommonTargetConfiguration) SetSchedule(v TargetSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


