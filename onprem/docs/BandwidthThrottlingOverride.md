# BandwidthThrottlingOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesPerSecond** | Pointer to **NullableInt64** | Specifies the value for the specified time period. The value is specified in bytes per second. | [optional] 
**TimePeriods** | Pointer to [**TimeOfAWeek**](TimeOfAWeek.md) |  | [optional] 

## Methods

### NewBandwidthThrottlingOverride

`func NewBandwidthThrottlingOverride() *BandwidthThrottlingOverride`

NewBandwidthThrottlingOverride instantiates a new BandwidthThrottlingOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthThrottlingOverrideWithDefaults

`func NewBandwidthThrottlingOverrideWithDefaults() *BandwidthThrottlingOverride`

NewBandwidthThrottlingOverrideWithDefaults instantiates a new BandwidthThrottlingOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesPerSecond

`func (o *BandwidthThrottlingOverride) GetBytesPerSecond() int64`

GetBytesPerSecond returns the BytesPerSecond field if non-nil, zero value otherwise.

### GetBytesPerSecondOk

`func (o *BandwidthThrottlingOverride) GetBytesPerSecondOk() (*int64, bool)`

GetBytesPerSecondOk returns a tuple with the BytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesPerSecond

`func (o *BandwidthThrottlingOverride) SetBytesPerSecond(v int64)`

SetBytesPerSecond sets BytesPerSecond field to given value.

### HasBytesPerSecond

`func (o *BandwidthThrottlingOverride) HasBytesPerSecond() bool`

HasBytesPerSecond returns a boolean if a field has been set.

### SetBytesPerSecondNil

`func (o *BandwidthThrottlingOverride) SetBytesPerSecondNil(b bool)`

 SetBytesPerSecondNil sets the value for BytesPerSecond to be an explicit nil

### UnsetBytesPerSecond
`func (o *BandwidthThrottlingOverride) UnsetBytesPerSecond()`

UnsetBytesPerSecond ensures that no value is present for BytesPerSecond, not even an explicit nil
### GetTimePeriods

`func (o *BandwidthThrottlingOverride) GetTimePeriods() TimeOfAWeek`

GetTimePeriods returns the TimePeriods field if non-nil, zero value otherwise.

### GetTimePeriodsOk

`func (o *BandwidthThrottlingOverride) GetTimePeriodsOk() (*TimeOfAWeek, bool)`

GetTimePeriodsOk returns a tuple with the TimePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriods

`func (o *BandwidthThrottlingOverride) SetTimePeriods(v TimeOfAWeek)`

SetTimePeriods sets TimePeriods field to given value.

### HasTimePeriods

`func (o *BandwidthThrottlingOverride) HasTimePeriods() bool`

HasTimePeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


