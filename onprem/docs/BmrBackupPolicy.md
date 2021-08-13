# BmrBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**BmrSchedule**](BmrSchedule.md) |  | 
**Retention** | [**Retention**](Retention.md) |  | 

## Methods

### NewBmrBackupPolicy

`func NewBmrBackupPolicy(schedule BmrSchedule, retention Retention, ) *BmrBackupPolicy`

NewBmrBackupPolicy instantiates a new BmrBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBmrBackupPolicyWithDefaults

`func NewBmrBackupPolicyWithDefaults() *BmrBackupPolicy`

NewBmrBackupPolicyWithDefaults instantiates a new BmrBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *BmrBackupPolicy) GetSchedule() BmrSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BmrBackupPolicy) GetScheduleOk() (*BmrSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BmrBackupPolicy) SetSchedule(v BmrSchedule)`

SetSchedule sets Schedule field to given value.


### GetRetention

`func (o *BmrBackupPolicy) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *BmrBackupPolicy) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *BmrBackupPolicy) SetRetention(v Retention)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


