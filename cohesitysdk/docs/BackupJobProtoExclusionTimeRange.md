# BackupJobProtoExclusionTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **NullableInt32** | If the day is not set, the time range applies to all days. | [optional] 
**EndTime** | Pointer to [**Time**](Time.md) |  | [optional] 
**StartTime** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewBackupJobProtoExclusionTimeRange

`func NewBackupJobProtoExclusionTimeRange() *BackupJobProtoExclusionTimeRange`

NewBackupJobProtoExclusionTimeRange instantiates a new BackupJobProtoExclusionTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobProtoExclusionTimeRangeWithDefaults

`func NewBackupJobProtoExclusionTimeRangeWithDefaults() *BackupJobProtoExclusionTimeRange`

NewBackupJobProtoExclusionTimeRangeWithDefaults instantiates a new BackupJobProtoExclusionTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *BackupJobProtoExclusionTimeRange) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupJobProtoExclusionTimeRange) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupJobProtoExclusionTimeRange) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupJobProtoExclusionTimeRange) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *BackupJobProtoExclusionTimeRange) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *BackupJobProtoExclusionTimeRange) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetEndTime

`func (o *BackupJobProtoExclusionTimeRange) GetEndTime() Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BackupJobProtoExclusionTimeRange) GetEndTimeOk() (*Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BackupJobProtoExclusionTimeRange) SetEndTime(v Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BackupJobProtoExclusionTimeRange) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupJobProtoExclusionTimeRange) GetStartTime() Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupJobProtoExclusionTimeRange) GetStartTimeOk() (*Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupJobProtoExclusionTimeRange) SetStartTime(v Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupJobProtoExclusionTimeRange) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


