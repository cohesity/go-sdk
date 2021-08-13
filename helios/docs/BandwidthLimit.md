# BandwidthLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesPerSecond** | Pointer to **NullableInt64** | Specifies the value for the specified time period. The value is specified in bytes per second. | [optional] 
**TimePeriods** | Pointer to [**TimeOfAWeek**](TimeOfAWeek.md) |  | [optional] 

## Methods

### NewBandwidthLimit

`func NewBandwidthLimit() *BandwidthLimit`

NewBandwidthLimit instantiates a new BandwidthLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthLimitWithDefaults

`func NewBandwidthLimitWithDefaults() *BandwidthLimit`

NewBandwidthLimitWithDefaults instantiates a new BandwidthLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesPerSecond

`func (o *BandwidthLimit) GetBytesPerSecond() int64`

GetBytesPerSecond returns the BytesPerSecond field if non-nil, zero value otherwise.

### GetBytesPerSecondOk

`func (o *BandwidthLimit) GetBytesPerSecondOk() (*int64, bool)`

GetBytesPerSecondOk returns a tuple with the BytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesPerSecond

`func (o *BandwidthLimit) SetBytesPerSecond(v int64)`

SetBytesPerSecond sets BytesPerSecond field to given value.

### HasBytesPerSecond

`func (o *BandwidthLimit) HasBytesPerSecond() bool`

HasBytesPerSecond returns a boolean if a field has been set.

### SetBytesPerSecondNil

`func (o *BandwidthLimit) SetBytesPerSecondNil(b bool)`

 SetBytesPerSecondNil sets the value for BytesPerSecond to be an explicit nil

### UnsetBytesPerSecond
`func (o *BandwidthLimit) UnsetBytesPerSecond()`

UnsetBytesPerSecond ensures that no value is present for BytesPerSecond, not even an explicit nil
### GetTimePeriods

`func (o *BandwidthLimit) GetTimePeriods() TimeOfAWeek`

GetTimePeriods returns the TimePeriods field if non-nil, zero value otherwise.

### GetTimePeriodsOk

`func (o *BandwidthLimit) GetTimePeriodsOk() (*TimeOfAWeek, bool)`

GetTimePeriodsOk returns a tuple with the TimePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriods

`func (o *BandwidthLimit) SetTimePeriods(v TimeOfAWeek)`

SetTimePeriods sets TimePeriods field to given value.

### HasTimePeriods

`func (o *BandwidthLimit) HasTimePeriods() bool`

HasTimePeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


