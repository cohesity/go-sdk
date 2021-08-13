# SupportChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **NullableBool** | Specifies id the support channel is enabled. | 
**EndTimeUsecs** | **NullableInt64** | Specifies the support channel expiry time. | 

## Methods

### NewSupportChannelConfig

`func NewSupportChannelConfig(isEnabled NullableBool, endTimeUsecs NullableInt64, ) *SupportChannelConfig`

NewSupportChannelConfig instantiates a new SupportChannelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportChannelConfigWithDefaults

`func NewSupportChannelConfigWithDefaults() *SupportChannelConfig`

NewSupportChannelConfigWithDefaults instantiates a new SupportChannelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SupportChannelConfig) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SupportChannelConfig) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SupportChannelConfig) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### SetIsEnabledNil

`func (o *SupportChannelConfig) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *SupportChannelConfig) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetEndTimeUsecs

`func (o *SupportChannelConfig) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *SupportChannelConfig) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *SupportChannelConfig) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.


### SetEndTimeUsecsNil

`func (o *SupportChannelConfig) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *SupportChannelConfig) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


