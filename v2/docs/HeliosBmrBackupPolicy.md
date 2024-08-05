# HeliosBmrBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 
**Schedule** | Pointer to [**HeliosBmrSchedule**](HeliosBmrSchedule.md) |  | [optional] 

## Methods

### NewHeliosBmrBackupPolicy

`func NewHeliosBmrBackupPolicy() *HeliosBmrBackupPolicy`

NewHeliosBmrBackupPolicy instantiates a new HeliosBmrBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosBmrBackupPolicyWithDefaults

`func NewHeliosBmrBackupPolicyWithDefaults() *HeliosBmrBackupPolicy`

NewHeliosBmrBackupPolicyWithDefaults instantiates a new HeliosBmrBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *HeliosBmrBackupPolicy) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosBmrBackupPolicy) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosBmrBackupPolicy) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosBmrBackupPolicy) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *HeliosBmrBackupPolicy) GetSchedule() HeliosBmrSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosBmrBackupPolicy) GetScheduleOk() (*HeliosBmrSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosBmrBackupPolicy) SetSchedule(v HeliosBmrSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HeliosBmrBackupPolicy) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


