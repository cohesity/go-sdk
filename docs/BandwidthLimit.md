# BandwidthLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthLimitOverrides** | Pointer to [**[]BandwidthLimitOverride**](BandwidthLimitOverride.md) | Array of Override Bandwidth Limits.  Specifies a list of override bandwidth limits and time periods when those limits override the rateLimitBytesPerSec limit. If overlapping time periods are specified, the last one in the array takes precedence. | [optional] 
**IoRate** | Pointer to **NullableInt32** | Specifies the default IO Rate of the throttling schedule. This value is internally mapped to some notion of how many resources a process should be consuming. | [optional] 
**RateLimitBytesPerSec** | Pointer to **NullableInt64** | Specifies the maximum allowed data transfer rate between the local Cluster and remote Clusters. The value is specified in bytes per second. If not set, the data transfer rate is not limited. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies a time zone for the specified time period. The time zone is defined in the following format: \&quot;Area/Location\&quot;, for example: \&quot;America/New_York\&quot;. | [optional] 

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

### GetBandwidthLimitOverrides

`func (o *BandwidthLimit) GetBandwidthLimitOverrides() []BandwidthLimitOverride`

GetBandwidthLimitOverrides returns the BandwidthLimitOverrides field if non-nil, zero value otherwise.

### GetBandwidthLimitOverridesOk

`func (o *BandwidthLimit) GetBandwidthLimitOverridesOk() (*[]BandwidthLimitOverride, bool)`

GetBandwidthLimitOverridesOk returns a tuple with the BandwidthLimitOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimitOverrides

`func (o *BandwidthLimit) SetBandwidthLimitOverrides(v []BandwidthLimitOverride)`

SetBandwidthLimitOverrides sets BandwidthLimitOverrides field to given value.

### HasBandwidthLimitOverrides

`func (o *BandwidthLimit) HasBandwidthLimitOverrides() bool`

HasBandwidthLimitOverrides returns a boolean if a field has been set.

### SetBandwidthLimitOverridesNil

`func (o *BandwidthLimit) SetBandwidthLimitOverridesNil(b bool)`

 SetBandwidthLimitOverridesNil sets the value for BandwidthLimitOverrides to be an explicit nil

### UnsetBandwidthLimitOverrides
`func (o *BandwidthLimit) UnsetBandwidthLimitOverrides()`

UnsetBandwidthLimitOverrides ensures that no value is present for BandwidthLimitOverrides, not even an explicit nil
### GetIoRate

`func (o *BandwidthLimit) GetIoRate() int32`

GetIoRate returns the IoRate field if non-nil, zero value otherwise.

### GetIoRateOk

`func (o *BandwidthLimit) GetIoRateOk() (*int32, bool)`

GetIoRateOk returns a tuple with the IoRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoRate

`func (o *BandwidthLimit) SetIoRate(v int32)`

SetIoRate sets IoRate field to given value.

### HasIoRate

`func (o *BandwidthLimit) HasIoRate() bool`

HasIoRate returns a boolean if a field has been set.

### SetIoRateNil

`func (o *BandwidthLimit) SetIoRateNil(b bool)`

 SetIoRateNil sets the value for IoRate to be an explicit nil

### UnsetIoRate
`func (o *BandwidthLimit) UnsetIoRate()`

UnsetIoRate ensures that no value is present for IoRate, not even an explicit nil
### GetRateLimitBytesPerSec

`func (o *BandwidthLimit) GetRateLimitBytesPerSec() int64`

GetRateLimitBytesPerSec returns the RateLimitBytesPerSec field if non-nil, zero value otherwise.

### GetRateLimitBytesPerSecOk

`func (o *BandwidthLimit) GetRateLimitBytesPerSecOk() (*int64, bool)`

GetRateLimitBytesPerSecOk returns a tuple with the RateLimitBytesPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitBytesPerSec

`func (o *BandwidthLimit) SetRateLimitBytesPerSec(v int64)`

SetRateLimitBytesPerSec sets RateLimitBytesPerSec field to given value.

### HasRateLimitBytesPerSec

`func (o *BandwidthLimit) HasRateLimitBytesPerSec() bool`

HasRateLimitBytesPerSec returns a boolean if a field has been set.

### SetRateLimitBytesPerSecNil

`func (o *BandwidthLimit) SetRateLimitBytesPerSecNil(b bool)`

 SetRateLimitBytesPerSecNil sets the value for RateLimitBytesPerSec to be an explicit nil

### UnsetRateLimitBytesPerSec
`func (o *BandwidthLimit) UnsetRateLimitBytesPerSec()`

UnsetRateLimitBytesPerSec ensures that no value is present for RateLimitBytesPerSec, not even an explicit nil
### GetTimezone

`func (o *BandwidthLimit) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BandwidthLimit) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BandwidthLimit) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BandwidthLimit) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *BandwidthLimit) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *BandwidthLimit) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


