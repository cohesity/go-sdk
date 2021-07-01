# SchedulingPolicyProtoRPOSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RpoIntervalMins** | Pointer to **NullableInt64** | If this field is set, then at any point, a recovery point should be available not older than the given interval mins. | [optional] 

## Methods

### NewSchedulingPolicyProtoRPOSchedule

`func NewSchedulingPolicyProtoRPOSchedule() *SchedulingPolicyProtoRPOSchedule`

NewSchedulingPolicyProtoRPOSchedule instantiates a new SchedulingPolicyProtoRPOSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingPolicyProtoRPOScheduleWithDefaults

`func NewSchedulingPolicyProtoRPOScheduleWithDefaults() *SchedulingPolicyProtoRPOSchedule`

NewSchedulingPolicyProtoRPOScheduleWithDefaults instantiates a new SchedulingPolicyProtoRPOSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpoIntervalMins

`func (o *SchedulingPolicyProtoRPOSchedule) GetRpoIntervalMins() int64`

GetRpoIntervalMins returns the RpoIntervalMins field if non-nil, zero value otherwise.

### GetRpoIntervalMinsOk

`func (o *SchedulingPolicyProtoRPOSchedule) GetRpoIntervalMinsOk() (*int64, bool)`

GetRpoIntervalMinsOk returns a tuple with the RpoIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoIntervalMins

`func (o *SchedulingPolicyProtoRPOSchedule) SetRpoIntervalMins(v int64)`

SetRpoIntervalMins sets RpoIntervalMins field to given value.

### HasRpoIntervalMins

`func (o *SchedulingPolicyProtoRPOSchedule) HasRpoIntervalMins() bool`

HasRpoIntervalMins returns a boolean if a field has been set.

### SetRpoIntervalMinsNil

`func (o *SchedulingPolicyProtoRPOSchedule) SetRpoIntervalMinsNil(b bool)`

 SetRpoIntervalMinsNil sets the value for RpoIntervalMins to be an explicit nil

### UnsetRpoIntervalMins
`func (o *SchedulingPolicyProtoRPOSchedule) UnsetRpoIntervalMins()`

UnsetRpoIntervalMins ensures that no value is present for RpoIntervalMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


