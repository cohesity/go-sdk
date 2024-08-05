# HeliosLogBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 
**Schedule** | Pointer to [**HeliosLogSchedule**](HeliosLogSchedule.md) |  | [optional] 

## Methods

### NewHeliosLogBackupPolicy

`func NewHeliosLogBackupPolicy() *HeliosLogBackupPolicy`

NewHeliosLogBackupPolicy instantiates a new HeliosLogBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosLogBackupPolicyWithDefaults

`func NewHeliosLogBackupPolicyWithDefaults() *HeliosLogBackupPolicy`

NewHeliosLogBackupPolicyWithDefaults instantiates a new HeliosLogBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *HeliosLogBackupPolicy) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosLogBackupPolicy) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosLogBackupPolicy) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosLogBackupPolicy) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *HeliosLogBackupPolicy) GetSchedule() HeliosLogSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosLogBackupPolicy) GetScheduleOk() (*HeliosLogSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosLogBackupPolicy) SetSchedule(v HeliosLogSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HeliosLogBackupPolicy) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


