# StoragePolicyErasureCodingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelaySecs** | Pointer to **NullableInt32** | Specifies the time in seconds when erasure coding starts. | [optional] 
**Enabled** | **NullableBool** | Specifies whether to enable erasure coding on a Storage Domain. | 
**InlineEnabled** | Pointer to **NullableBool** | Specifies whether inline erasure coding is enabled. This field is appliciable only if enabled is set to true. | [optional] 
**NumCodedStripes** | **NullableInt32** | Specifies the number of coded stripes. | 
**NumDataStripes** | **NullableInt32** | Specifies the number of data stripes. | 

## Methods

### NewStoragePolicyErasureCodingParams

`func NewStoragePolicyErasureCodingParams(enabled NullableBool, numCodedStripes NullableInt32, numDataStripes NullableInt32, ) *StoragePolicyErasureCodingParams`

NewStoragePolicyErasureCodingParams instantiates a new StoragePolicyErasureCodingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyErasureCodingParamsWithDefaults

`func NewStoragePolicyErasureCodingParamsWithDefaults() *StoragePolicyErasureCodingParams`

NewStoragePolicyErasureCodingParamsWithDefaults instantiates a new StoragePolicyErasureCodingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelaySecs

`func (o *StoragePolicyErasureCodingParams) GetDelaySecs() int32`

GetDelaySecs returns the DelaySecs field if non-nil, zero value otherwise.

### GetDelaySecsOk

`func (o *StoragePolicyErasureCodingParams) GetDelaySecsOk() (*int32, bool)`

GetDelaySecsOk returns a tuple with the DelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySecs

`func (o *StoragePolicyErasureCodingParams) SetDelaySecs(v int32)`

SetDelaySecs sets DelaySecs field to given value.

### HasDelaySecs

`func (o *StoragePolicyErasureCodingParams) HasDelaySecs() bool`

HasDelaySecs returns a boolean if a field has been set.

### SetDelaySecsNil

`func (o *StoragePolicyErasureCodingParams) SetDelaySecsNil(b bool)`

 SetDelaySecsNil sets the value for DelaySecs to be an explicit nil

### UnsetDelaySecs
`func (o *StoragePolicyErasureCodingParams) UnsetDelaySecs()`

UnsetDelaySecs ensures that no value is present for DelaySecs, not even an explicit nil
### GetEnabled

`func (o *StoragePolicyErasureCodingParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoragePolicyErasureCodingParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoragePolicyErasureCodingParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *StoragePolicyErasureCodingParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *StoragePolicyErasureCodingParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInlineEnabled

`func (o *StoragePolicyErasureCodingParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *StoragePolicyErasureCodingParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *StoragePolicyErasureCodingParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *StoragePolicyErasureCodingParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *StoragePolicyErasureCodingParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *StoragePolicyErasureCodingParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
### GetNumCodedStripes

`func (o *StoragePolicyErasureCodingParams) GetNumCodedStripes() int32`

GetNumCodedStripes returns the NumCodedStripes field if non-nil, zero value otherwise.

### GetNumCodedStripesOk

`func (o *StoragePolicyErasureCodingParams) GetNumCodedStripesOk() (*int32, bool)`

GetNumCodedStripesOk returns a tuple with the NumCodedStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCodedStripes

`func (o *StoragePolicyErasureCodingParams) SetNumCodedStripes(v int32)`

SetNumCodedStripes sets NumCodedStripes field to given value.


### SetNumCodedStripesNil

`func (o *StoragePolicyErasureCodingParams) SetNumCodedStripesNil(b bool)`

 SetNumCodedStripesNil sets the value for NumCodedStripes to be an explicit nil

### UnsetNumCodedStripes
`func (o *StoragePolicyErasureCodingParams) UnsetNumCodedStripes()`

UnsetNumCodedStripes ensures that no value is present for NumCodedStripes, not even an explicit nil
### GetNumDataStripes

`func (o *StoragePolicyErasureCodingParams) GetNumDataStripes() int32`

GetNumDataStripes returns the NumDataStripes field if non-nil, zero value otherwise.

### GetNumDataStripesOk

`func (o *StoragePolicyErasureCodingParams) GetNumDataStripesOk() (*int32, bool)`

GetNumDataStripesOk returns a tuple with the NumDataStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDataStripes

`func (o *StoragePolicyErasureCodingParams) SetNumDataStripes(v int32)`

SetNumDataStripes sets NumDataStripes field to given value.


### SetNumDataStripesNil

`func (o *StoragePolicyErasureCodingParams) SetNumDataStripesNil(b bool)`

 SetNumDataStripesNil sets the value for NumDataStripes to be an explicit nil

### UnsetNumDataStripes
`func (o *StoragePolicyErasureCodingParams) UnsetNumDataStripes()`

UnsetNumDataStripes ensures that no value is present for NumDataStripes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


