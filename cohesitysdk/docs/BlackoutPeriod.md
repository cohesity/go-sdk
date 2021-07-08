# BlackoutPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specified the Id for a snapshot copy policy. This is generated when the policy is created. | [optional] 
**Day** | Pointer to **NullableString** | Blackout Day.  Specifies a day in the week when no new Job Runs should be started such as &#39;kSunday&#39;. If not set, the time range applies to all days. Specifies a day in a week such as &#39;kSunday&#39;, &#39;kMonday&#39;, etc. | [optional] 
**EndTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the end time of the blackout time range. | [optional] 
**StartTime** | Pointer to [**NullableTimeOfDay**](TimeOfDay.md) | Specifies the start time of the blackout time range. | [optional] 

## Methods

### NewBlackoutPeriod

`func NewBlackoutPeriod() *BlackoutPeriod`

NewBlackoutPeriod instantiates a new BlackoutPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutPeriodWithDefaults

`func NewBlackoutPeriodWithDefaults() *BlackoutPeriod`

NewBlackoutPeriodWithDefaults instantiates a new BlackoutPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlackoutPeriod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlackoutPeriod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlackoutPeriod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlackoutPeriod) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BlackoutPeriod) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BlackoutPeriod) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDay

`func (o *BlackoutPeriod) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BlackoutPeriod) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BlackoutPeriod) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *BlackoutPeriod) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *BlackoutPeriod) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *BlackoutPeriod) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetEndTime

`func (o *BlackoutPeriod) GetEndTime() TimeOfDay`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BlackoutPeriod) GetEndTimeOk() (*TimeOfDay, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BlackoutPeriod) SetEndTime(v TimeOfDay)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BlackoutPeriod) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *BlackoutPeriod) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *BlackoutPeriod) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *BlackoutPeriod) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BlackoutPeriod) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BlackoutPeriod) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BlackoutPeriod) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BlackoutPeriod) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BlackoutPeriod) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


