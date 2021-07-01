# BandwidthLimitOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesPerSecond** | Pointer to **NullableInt64** | Specifies the value to override the regular maximum bandwidth rate (rateLimitBytesPerSec) for the specified time period. The value is specified in bytes per second. | [optional] 
**IoRate** | Pointer to **NullableInt32** | Specifies the value to override the default IO rate for the specified time period. | [optional] 
**TimePeriods** | Pointer to [**TimeOfAWeek**](TimeOfAWeek.md) |  | [optional] 

## Methods

### NewBandwidthLimitOverride

`func NewBandwidthLimitOverride() *BandwidthLimitOverride`

NewBandwidthLimitOverride instantiates a new BandwidthLimitOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthLimitOverrideWithDefaults

`func NewBandwidthLimitOverrideWithDefaults() *BandwidthLimitOverride`

NewBandwidthLimitOverrideWithDefaults instantiates a new BandwidthLimitOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesPerSecond

`func (o *BandwidthLimitOverride) GetBytesPerSecond() int64`

GetBytesPerSecond returns the BytesPerSecond field if non-nil, zero value otherwise.

### GetBytesPerSecondOk

`func (o *BandwidthLimitOverride) GetBytesPerSecondOk() (*int64, bool)`

GetBytesPerSecondOk returns a tuple with the BytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesPerSecond

`func (o *BandwidthLimitOverride) SetBytesPerSecond(v int64)`

SetBytesPerSecond sets BytesPerSecond field to given value.

### HasBytesPerSecond

`func (o *BandwidthLimitOverride) HasBytesPerSecond() bool`

HasBytesPerSecond returns a boolean if a field has been set.

### SetBytesPerSecondNil

`func (o *BandwidthLimitOverride) SetBytesPerSecondNil(b bool)`

 SetBytesPerSecondNil sets the value for BytesPerSecond to be an explicit nil

### UnsetBytesPerSecond
`func (o *BandwidthLimitOverride) UnsetBytesPerSecond()`

UnsetBytesPerSecond ensures that no value is present for BytesPerSecond, not even an explicit nil
### GetIoRate

`func (o *BandwidthLimitOverride) GetIoRate() int32`

GetIoRate returns the IoRate field if non-nil, zero value otherwise.

### GetIoRateOk

`func (o *BandwidthLimitOverride) GetIoRateOk() (*int32, bool)`

GetIoRateOk returns a tuple with the IoRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRate

`func (o *BandwidthLimitOverride) SetIoRate(v int32)`

SetIoRate sets IoRate field to given value.

### HasIoRate

`func (o *BandwidthLimitOverride) HasIoRate() bool`

HasIoRate returns a boolean if a field has been set.

### SetIoRateNil

`func (o *BandwidthLimitOverride) SetIoRateNil(b bool)`

 SetIoRateNil sets the value for IoRate to be an explicit nil

### UnsetIoRate
`func (o *BandwidthLimitOverride) UnsetIoRate()`

UnsetIoRate ensures that no value is present for IoRate, not even an explicit nil
### GetTimePeriods

`func (o *BandwidthLimitOverride) GetTimePeriods() TimeOfAWeek`

GetTimePeriods returns the TimePeriods field if non-nil, zero value otherwise.

### GetTimePeriodsOk

`func (o *BandwidthLimitOverride) GetTimePeriodsOk() (*TimeOfAWeek, bool)`

GetTimePeriodsOk returns a tuple with the TimePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriods

`func (o *BandwidthLimitOverride) SetTimePeriods(v TimeOfAWeek)`

SetTimePeriods sets TimePeriods field to given value.

### HasTimePeriods

`func (o *BandwidthLimitOverride) HasTimePeriods() bool`

HasTimePeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


