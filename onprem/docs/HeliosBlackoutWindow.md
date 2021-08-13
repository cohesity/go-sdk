# HeliosBlackoutWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **NullableString** | Specifies a day in the week when no new Protection Group Runs should be started such as &#39;Sunday&#39;. Specifies a day in a week such as &#39;Sunday&#39;, &#39;Monday&#39;, etc. | 
**StartTime** | [**TimeOfDay**](TimeOfDay.md) |  | 
**EndTime** | [**TimeOfDay**](TimeOfDay.md) |  | 
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the blackout getting added. This field should only be set if policy is getting updated. | [optional] 

## Methods

### NewHeliosBlackoutWindow

`func NewHeliosBlackoutWindow(day NullableString, startTime TimeOfDay, endTime TimeOfDay, ) *HeliosBlackoutWindow`

NewHeliosBlackoutWindow instantiates a new HeliosBlackoutWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosBlackoutWindowWithDefaults

`func NewHeliosBlackoutWindowWithDefaults() *HeliosBlackoutWindow`

NewHeliosBlackoutWindowWithDefaults instantiates a new HeliosBlackoutWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *HeliosBlackoutWindow) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *HeliosBlackoutWindow) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *HeliosBlackoutWindow) SetDay(v string)`

SetDay sets Day field to given value.


### SetDayNil

`func (o *HeliosBlackoutWindow) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *HeliosBlackoutWindow) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetStartTime

`func (o *HeliosBlackoutWindow) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HeliosBlackoutWindow) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HeliosBlackoutWindow) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *HeliosBlackoutWindow) GetEndTime() TimeOfDay`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HeliosBlackoutWindow) GetEndTimeOk() (*TimeOfDay, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HeliosBlackoutWindow) SetEndTime(v TimeOfDay)`

SetEndTime sets EndTime field to given value.


### GetConfigId

`func (o *HeliosBlackoutWindow) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *HeliosBlackoutWindow) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *HeliosBlackoutWindow) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *HeliosBlackoutWindow) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *HeliosBlackoutWindow) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *HeliosBlackoutWindow) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


