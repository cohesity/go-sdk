# IncrementalBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**IncrementalSchedule**](IncrementalSchedule.md) |  | 

## Methods

### NewIncrementalBackupPolicy

`func NewIncrementalBackupPolicy(schedule IncrementalSchedule, ) *IncrementalBackupPolicy`

NewIncrementalBackupPolicy instantiates a new IncrementalBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncrementalBackupPolicyWithDefaults

`func NewIncrementalBackupPolicyWithDefaults() *IncrementalBackupPolicy`

NewIncrementalBackupPolicyWithDefaults instantiates a new IncrementalBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *IncrementalBackupPolicy) GetSchedule() IncrementalSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *IncrementalBackupPolicy) GetScheduleOk() (*IncrementalSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *IncrementalBackupPolicy) SetSchedule(v IncrementalSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


