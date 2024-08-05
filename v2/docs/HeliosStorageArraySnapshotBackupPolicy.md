# HeliosStorageArraySnapshotBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**HeliosRetention**](HeliosRetention.md) |  | 
**Schedule** | [**HeliosStorageArraySnapshotSchedule**](HeliosStorageArraySnapshotSchedule.md) |  | 

## Methods

### NewHeliosStorageArraySnapshotBackupPolicy

`func NewHeliosStorageArraySnapshotBackupPolicy(retention HeliosRetention, schedule HeliosStorageArraySnapshotSchedule, ) *HeliosStorageArraySnapshotBackupPolicy`

NewHeliosStorageArraySnapshotBackupPolicy instantiates a new HeliosStorageArraySnapshotBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosStorageArraySnapshotBackupPolicyWithDefaults

`func NewHeliosStorageArraySnapshotBackupPolicyWithDefaults() *HeliosStorageArraySnapshotBackupPolicy`

NewHeliosStorageArraySnapshotBackupPolicyWithDefaults instantiates a new HeliosStorageArraySnapshotBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *HeliosStorageArraySnapshotBackupPolicy) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosStorageArraySnapshotBackupPolicy) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosStorageArraySnapshotBackupPolicy) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.


### GetSchedule

`func (o *HeliosStorageArraySnapshotBackupPolicy) GetSchedule() HeliosStorageArraySnapshotSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosStorageArraySnapshotBackupPolicy) GetScheduleOk() (*HeliosStorageArraySnapshotSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosStorageArraySnapshotBackupPolicy) SetSchedule(v HeliosStorageArraySnapshotSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


