# ViewPinningConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies if view pinning is enabled. If set to true, view will be pinned to SSD and view data will be stored there. | 
**LastUpdatedTimestampSecs** | Pointer to **NullableInt64** | Specifies the timestamp when view pinning config is last updated. | [optional] [readonly] 
**PinnedTimeSecs** | Pointer to **NullableInt64** | Specifies the time to pin files after last access. | [optional] 

## Methods

### NewViewPinningConfig

`func NewViewPinningConfig(enabled NullableBool, ) *ViewPinningConfig`

NewViewPinningConfig instantiates a new ViewPinningConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPinningConfigWithDefaults

`func NewViewPinningConfigWithDefaults() *ViewPinningConfig`

NewViewPinningConfigWithDefaults instantiates a new ViewPinningConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ViewPinningConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ViewPinningConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ViewPinningConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ViewPinningConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ViewPinningConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLastUpdatedTimestampSecs

`func (o *ViewPinningConfig) GetLastUpdatedTimestampSecs() int64`

GetLastUpdatedTimestampSecs returns the LastUpdatedTimestampSecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampSecsOk

`func (o *ViewPinningConfig) GetLastUpdatedTimestampSecsOk() (*int64, bool)`

GetLastUpdatedTimestampSecsOk returns a tuple with the LastUpdatedTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestampSecs

`func (o *ViewPinningConfig) SetLastUpdatedTimestampSecs(v int64)`

SetLastUpdatedTimestampSecs sets LastUpdatedTimestampSecs field to given value.

### HasLastUpdatedTimestampSecs

`func (o *ViewPinningConfig) HasLastUpdatedTimestampSecs() bool`

HasLastUpdatedTimestampSecs returns a boolean if a field has been set.

### SetLastUpdatedTimestampSecsNil

`func (o *ViewPinningConfig) SetLastUpdatedTimestampSecsNil(b bool)`

 SetLastUpdatedTimestampSecsNil sets the value for LastUpdatedTimestampSecs to be an explicit nil

### UnsetLastUpdatedTimestampSecs
`func (o *ViewPinningConfig) UnsetLastUpdatedTimestampSecs()`

UnsetLastUpdatedTimestampSecs ensures that no value is present for LastUpdatedTimestampSecs, not even an explicit nil
### GetPinnedTimeSecs

`func (o *ViewPinningConfig) GetPinnedTimeSecs() int64`

GetPinnedTimeSecs returns the PinnedTimeSecs field if non-nil, zero value otherwise.

### GetPinnedTimeSecsOk

`func (o *ViewPinningConfig) GetPinnedTimeSecsOk() (*int64, bool)`

GetPinnedTimeSecsOk returns a tuple with the PinnedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedTimeSecs

`func (o *ViewPinningConfig) SetPinnedTimeSecs(v int64)`

SetPinnedTimeSecs sets PinnedTimeSecs field to given value.

### HasPinnedTimeSecs

`func (o *ViewPinningConfig) HasPinnedTimeSecs() bool`

HasPinnedTimeSecs returns a boolean if a field has been set.

### SetPinnedTimeSecsNil

`func (o *ViewPinningConfig) SetPinnedTimeSecsNil(b bool)`

 SetPinnedTimeSecsNil sets the value for PinnedTimeSecs to be an explicit nil

### UnsetPinnedTimeSecs
`func (o *ViewPinningConfig) UnsetPinnedTimeSecs()`

UnsetPinnedTimeSecs ensures that no value is present for PinnedTimeSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


