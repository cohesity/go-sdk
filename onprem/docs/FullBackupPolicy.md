# FullBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**FullSchedule**](FullSchedule.md) |  | 

## Methods

### NewFullBackupPolicy

`func NewFullBackupPolicy(schedule FullSchedule, ) *FullBackupPolicy`

NewFullBackupPolicy instantiates a new FullBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullBackupPolicyWithDefaults

`func NewFullBackupPolicyWithDefaults() *FullBackupPolicy`

NewFullBackupPolicyWithDefaults instantiates a new FullBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *FullBackupPolicy) GetSchedule() FullSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *FullBackupPolicy) GetScheduleOk() (*FullSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *FullBackupPolicy) SetSchedule(v FullSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


