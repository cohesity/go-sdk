# BandwidthThrottling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RateLimitBytesPerSec** | Pointer to **NullableInt64** | Specifies the maximum allowed data transfer rate between the local Cluster and remote Clusters. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies a time zone for the specified time period. | [optional] 
**BandwidthLimitOverrides** | Pointer to [**[]BandwidthThrottlingOverride**](BandwidthThrottlingOverride.md) | Specifies the max rate limit at which we upload the data. | [optional] 

## Methods

### NewBandwidthThrottling

`func NewBandwidthThrottling() *BandwidthThrottling`

NewBandwidthThrottling instantiates a new BandwidthThrottling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthThrottlingWithDefaults

`func NewBandwidthThrottlingWithDefaults() *BandwidthThrottling`

NewBandwidthThrottlingWithDefaults instantiates a new BandwidthThrottling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRateLimitBytesPerSec

`func (o *BandwidthThrottling) GetRateLimitBytesPerSec() int64`

GetRateLimitBytesPerSec returns the RateLimitBytesPerSec field if non-nil, zero value otherwise.

### GetRateLimitBytesPerSecOk

`func (o *BandwidthThrottling) GetRateLimitBytesPerSecOk() (*int64, bool)`

GetRateLimitBytesPerSecOk returns a tuple with the RateLimitBytesPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitBytesPerSec

`func (o *BandwidthThrottling) SetRateLimitBytesPerSec(v int64)`

SetRateLimitBytesPerSec sets RateLimitBytesPerSec field to given value.

### HasRateLimitBytesPerSec

`func (o *BandwidthThrottling) HasRateLimitBytesPerSec() bool`

HasRateLimitBytesPerSec returns a boolean if a field has been set.

### SetRateLimitBytesPerSecNil

`func (o *BandwidthThrottling) SetRateLimitBytesPerSecNil(b bool)`

 SetRateLimitBytesPerSecNil sets the value for RateLimitBytesPerSec to be an explicit nil

### UnsetRateLimitBytesPerSec
`func (o *BandwidthThrottling) UnsetRateLimitBytesPerSec()`

UnsetRateLimitBytesPerSec ensures that no value is present for RateLimitBytesPerSec, not even an explicit nil
### GetTimezone

`func (o *BandwidthThrottling) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BandwidthThrottling) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BandwidthThrottling) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BandwidthThrottling) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *BandwidthThrottling) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *BandwidthThrottling) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetBandwidthLimitOverrides

`func (o *BandwidthThrottling) GetBandwidthLimitOverrides() []BandwidthThrottlingOverride`

GetBandwidthLimitOverrides returns the BandwidthLimitOverrides field if non-nil, zero value otherwise.

### GetBandwidthLimitOverridesOk

`func (o *BandwidthThrottling) GetBandwidthLimitOverridesOk() (*[]BandwidthThrottlingOverride, bool)`

GetBandwidthLimitOverridesOk returns a tuple with the BandwidthLimitOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimitOverrides

`func (o *BandwidthThrottling) SetBandwidthLimitOverrides(v []BandwidthThrottlingOverride)`

SetBandwidthLimitOverrides sets BandwidthLimitOverrides field to given value.

### HasBandwidthLimitOverrides

`func (o *BandwidthThrottling) HasBandwidthLimitOverrides() bool`

HasBandwidthLimitOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


