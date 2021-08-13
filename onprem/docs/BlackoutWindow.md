# BlackoutWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **NullableString** | Specifies a day in the week when no new Protection Group Runs should be started such as &#39;Sunday&#39;. Specifies a day in a week such as &#39;Sunday&#39;, &#39;Monday&#39;, etc. | 
**StartTime** | [**TimeOfDay**](TimeOfDay.md) |  | 
**EndTime** | [**TimeOfDay**](TimeOfDay.md) |  | 
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated. | [optional] 

## Methods

### NewBlackoutWindow

`func NewBlackoutWindow(day NullableString, startTime TimeOfDay, endTime TimeOfDay, ) *BlackoutWindow`

NewBlackoutWindow instantiates a new BlackoutWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutWindowWithDefaults

`func NewBlackoutWindowWithDefaults() *BlackoutWindow`

NewBlackoutWindowWithDefaults instantiates a new BlackoutWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *BlackoutWindow) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BlackoutWindow) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BlackoutWindow) SetDay(v string)`

SetDay sets Day field to given value.


### SetDayNil

`func (o *BlackoutWindow) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *BlackoutWindow) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetStartTime

`func (o *BlackoutWindow) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BlackoutWindow) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BlackoutWindow) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *BlackoutWindow) GetEndTime() TimeOfDay`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BlackoutWindow) GetEndTimeOk() (*TimeOfDay, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BlackoutWindow) SetEndTime(v TimeOfDay)`

SetEndTime sets EndTime field to given value.


### GetConfigId

`func (o *BlackoutWindow) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *BlackoutWindow) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *BlackoutWindow) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *BlackoutWindow) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *BlackoutWindow) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *BlackoutWindow) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


