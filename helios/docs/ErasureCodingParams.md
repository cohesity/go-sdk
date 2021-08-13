# ErasureCodingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies whether to enable erasure coding on a Storage Domain. | 
**Algorithm** | Pointer to **NullableString** | Specifies the algorithm used for erasure coding. | [optional] 
**NumDataStripes** | **NullableInt32** | Specifies the number of data stripes. | 
**NumCodedStripes** | **NullableInt32** | Specifies the number of coded stripes. | 
**InlineEnabled** | Pointer to **NullableBool** | Specifies whether inline erasure coding is enabled. This field is appliciable only if enabled is set to true. | [optional] 
**DelaySecs** | Pointer to **NullableInt32** | Specifies the time in seconds when erasure coding starts. | [optional] 
**UseContainer** | Pointer to **NullableBool** | Whether to encode chunk files as part of a container. This field is appliciable only if enabled is set to true. | [optional] 

## Methods

### NewErasureCodingParams

`func NewErasureCodingParams(enabled NullableBool, numDataStripes NullableInt32, numCodedStripes NullableInt32, ) *ErasureCodingParams`

NewErasureCodingParams instantiates a new ErasureCodingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErasureCodingParamsWithDefaults

`func NewErasureCodingParamsWithDefaults() *ErasureCodingParams`

NewErasureCodingParamsWithDefaults instantiates a new ErasureCodingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ErasureCodingParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ErasureCodingParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ErasureCodingParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ErasureCodingParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ErasureCodingParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetAlgorithm

`func (o *ErasureCodingParams) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ErasureCodingParams) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ErasureCodingParams) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ErasureCodingParams) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *ErasureCodingParams) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *ErasureCodingParams) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
### GetNumDataStripes

`func (o *ErasureCodingParams) GetNumDataStripes() int32`

GetNumDataStripes returns the NumDataStripes field if non-nil, zero value otherwise.

### GetNumDataStripesOk

`func (o *ErasureCodingParams) GetNumDataStripesOk() (*int32, bool)`

GetNumDataStripesOk returns a tuple with the NumDataStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDataStripes

`func (o *ErasureCodingParams) SetNumDataStripes(v int32)`

SetNumDataStripes sets NumDataStripes field to given value.


### SetNumDataStripesNil

`func (o *ErasureCodingParams) SetNumDataStripesNil(b bool)`

 SetNumDataStripesNil sets the value for NumDataStripes to be an explicit nil

### UnsetNumDataStripes
`func (o *ErasureCodingParams) UnsetNumDataStripes()`

UnsetNumDataStripes ensures that no value is present for NumDataStripes, not even an explicit nil
### GetNumCodedStripes

`func (o *ErasureCodingParams) GetNumCodedStripes() int32`

GetNumCodedStripes returns the NumCodedStripes field if non-nil, zero value otherwise.

### GetNumCodedStripesOk

`func (o *ErasureCodingParams) GetNumCodedStripesOk() (*int32, bool)`

GetNumCodedStripesOk returns a tuple with the NumCodedStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCodedStripes

`func (o *ErasureCodingParams) SetNumCodedStripes(v int32)`

SetNumCodedStripes sets NumCodedStripes field to given value.


### SetNumCodedStripesNil

`func (o *ErasureCodingParams) SetNumCodedStripesNil(b bool)`

 SetNumCodedStripesNil sets the value for NumCodedStripes to be an explicit nil

### UnsetNumCodedStripes
`func (o *ErasureCodingParams) UnsetNumCodedStripes()`

UnsetNumCodedStripes ensures that no value is present for NumCodedStripes, not even an explicit nil
### GetInlineEnabled

`func (o *ErasureCodingParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *ErasureCodingParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *ErasureCodingParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *ErasureCodingParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *ErasureCodingParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *ErasureCodingParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
### GetDelaySecs

`func (o *ErasureCodingParams) GetDelaySecs() int32`

GetDelaySecs returns the DelaySecs field if non-nil, zero value otherwise.

### GetDelaySecsOk

`func (o *ErasureCodingParams) GetDelaySecsOk() (*int32, bool)`

GetDelaySecsOk returns a tuple with the DelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySecs

`func (o *ErasureCodingParams) SetDelaySecs(v int32)`

SetDelaySecs sets DelaySecs field to given value.

### HasDelaySecs

`func (o *ErasureCodingParams) HasDelaySecs() bool`

HasDelaySecs returns a boolean if a field has been set.

### SetDelaySecsNil

`func (o *ErasureCodingParams) SetDelaySecsNil(b bool)`

 SetDelaySecsNil sets the value for DelaySecs to be an explicit nil

### UnsetDelaySecs
`func (o *ErasureCodingParams) UnsetDelaySecs()`

UnsetDelaySecs ensures that no value is present for DelaySecs, not even an explicit nil
### GetUseContainer

`func (o *ErasureCodingParams) GetUseContainer() bool`

GetUseContainer returns the UseContainer field if non-nil, zero value otherwise.

### GetUseContainerOk

`func (o *ErasureCodingParams) GetUseContainerOk() (*bool, bool)`

GetUseContainerOk returns a tuple with the UseContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseContainer

`func (o *ErasureCodingParams) SetUseContainer(v bool)`

SetUseContainer sets UseContainer field to given value.

### HasUseContainer

`func (o *ErasureCodingParams) HasUseContainer() bool`

HasUseContainer returns a boolean if a field has been set.

### SetUseContainerNil

`func (o *ErasureCodingParams) SetUseContainerNil(b bool)`

 SetUseContainerNil sets the value for UseContainer to be an explicit nil

### UnsetUseContainer
`func (o *ErasureCodingParams) UnsetUseContainer()`

UnsetUseContainer ensures that no value is present for UseContainer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


