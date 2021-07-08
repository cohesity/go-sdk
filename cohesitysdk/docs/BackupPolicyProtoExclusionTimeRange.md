# BackupPolicyProtoExclusionTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **NullableInt32** | If the day is not set, the time range applies to all days. | [optional] 
**EndTime** | Pointer to [**Time**](Time.md) |  | [optional] 
**StartTime** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewBackupPolicyProtoExclusionTimeRange

`func NewBackupPolicyProtoExclusionTimeRange() *BackupPolicyProtoExclusionTimeRange`

NewBackupPolicyProtoExclusionTimeRange instantiates a new BackupPolicyProtoExclusionTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyProtoExclusionTimeRangeWithDefaults

`func NewBackupPolicyProtoExclusionTimeRangeWithDefaults() *BackupPolicyProtoExclusionTimeRange`

NewBackupPolicyProtoExclusionTimeRangeWithDefaults instantiates a new BackupPolicyProtoExclusionTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *BackupPolicyProtoExclusionTimeRange) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupPolicyProtoExclusionTimeRange) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupPolicyProtoExclusionTimeRange) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupPolicyProtoExclusionTimeRange) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *BackupPolicyProtoExclusionTimeRange) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *BackupPolicyProtoExclusionTimeRange) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetEndTime

`func (o *BackupPolicyProtoExclusionTimeRange) GetEndTime() Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BackupPolicyProtoExclusionTimeRange) GetEndTimeOk() (*Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BackupPolicyProtoExclusionTimeRange) SetEndTime(v Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BackupPolicyProtoExclusionTimeRange) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupPolicyProtoExclusionTimeRange) GetStartTime() Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupPolicyProtoExclusionTimeRange) GetStartTimeOk() (*Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupPolicyProtoExclusionTimeRange) SetStartTime(v Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupPolicyProtoExclusionTimeRange) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


