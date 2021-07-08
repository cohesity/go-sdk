# ErasureCodingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **NullableString** | Algorthm used for erasure coding. REED_SOLOMON indicates the algorithm used for erasure coding. LRC indicates the algorithm used for erasure coding. | [optional] 
**ErasureCodingEnabled** | Pointer to **NullableBool** | Specifies whether Erasure coding is enabled on the Storage Domain (View Box). | [optional] 
**InlineErasureCoding** | Pointer to **NullableBool** | Specifies if erasure coding should occur inline (as the data is being written). This field is only relevant if erasure coding is enabled. | [optional] 
**NumCodedStripes** | Pointer to **NullableInt32** | The number of coded stripes. | [optional] 
**NumDataStripes** | Pointer to **NullableInt32** | The number of stripes containing data. | [optional] 

## Methods

### NewErasureCodingInfo

`func NewErasureCodingInfo() *ErasureCodingInfo`

NewErasureCodingInfo instantiates a new ErasureCodingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErasureCodingInfoWithDefaults

`func NewErasureCodingInfoWithDefaults() *ErasureCodingInfo`

NewErasureCodingInfoWithDefaults instantiates a new ErasureCodingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *ErasureCodingInfo) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ErasureCodingInfo) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ErasureCodingInfo) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ErasureCodingInfo) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *ErasureCodingInfo) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *ErasureCodingInfo) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
### GetErasureCodingEnabled

`func (o *ErasureCodingInfo) GetErasureCodingEnabled() bool`

GetErasureCodingEnabled returns the ErasureCodingEnabled field if non-nil, zero value otherwise.

### GetErasureCodingEnabledOk

`func (o *ErasureCodingInfo) GetErasureCodingEnabledOk() (*bool, bool)`

GetErasureCodingEnabledOk returns a tuple with the ErasureCodingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodingEnabled

`func (o *ErasureCodingInfo) SetErasureCodingEnabled(v bool)`

SetErasureCodingEnabled sets ErasureCodingEnabled field to given value.

### HasErasureCodingEnabled

`func (o *ErasureCodingInfo) HasErasureCodingEnabled() bool`

HasErasureCodingEnabled returns a boolean if a field has been set.

### SetErasureCodingEnabledNil

`func (o *ErasureCodingInfo) SetErasureCodingEnabledNil(b bool)`

 SetErasureCodingEnabledNil sets the value for ErasureCodingEnabled to be an explicit nil

### UnsetErasureCodingEnabled
`func (o *ErasureCodingInfo) UnsetErasureCodingEnabled()`

UnsetErasureCodingEnabled ensures that no value is present for ErasureCodingEnabled, not even an explicit nil
### GetInlineErasureCoding

`func (o *ErasureCodingInfo) GetInlineErasureCoding() bool`

GetInlineErasureCoding returns the InlineErasureCoding field if non-nil, zero value otherwise.

### GetInlineErasureCodingOk

`func (o *ErasureCodingInfo) GetInlineErasureCodingOk() (*bool, bool)`

GetInlineErasureCodingOk returns a tuple with the InlineErasureCoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineErasureCoding

`func (o *ErasureCodingInfo) SetInlineErasureCoding(v bool)`

SetInlineErasureCoding sets InlineErasureCoding field to given value.

### HasInlineErasureCoding

`func (o *ErasureCodingInfo) HasInlineErasureCoding() bool`

HasInlineErasureCoding returns a boolean if a field has been set.

### SetInlineErasureCodingNil

`func (o *ErasureCodingInfo) SetInlineErasureCodingNil(b bool)`

 SetInlineErasureCodingNil sets the value for InlineErasureCoding to be an explicit nil

### UnsetInlineErasureCoding
`func (o *ErasureCodingInfo) UnsetInlineErasureCoding()`

UnsetInlineErasureCoding ensures that no value is present for InlineErasureCoding, not even an explicit nil
### GetNumCodedStripes

`func (o *ErasureCodingInfo) GetNumCodedStripes() int32`

GetNumCodedStripes returns the NumCodedStripes field if non-nil, zero value otherwise.

### GetNumCodedStripesOk

`func (o *ErasureCodingInfo) GetNumCodedStripesOk() (*int32, bool)`

GetNumCodedStripesOk returns a tuple with the NumCodedStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCodedStripes

`func (o *ErasureCodingInfo) SetNumCodedStripes(v int32)`

SetNumCodedStripes sets NumCodedStripes field to given value.

### HasNumCodedStripes

`func (o *ErasureCodingInfo) HasNumCodedStripes() bool`

HasNumCodedStripes returns a boolean if a field has been set.

### SetNumCodedStripesNil

`func (o *ErasureCodingInfo) SetNumCodedStripesNil(b bool)`

 SetNumCodedStripesNil sets the value for NumCodedStripes to be an explicit nil

### UnsetNumCodedStripes
`func (o *ErasureCodingInfo) UnsetNumCodedStripes()`

UnsetNumCodedStripes ensures that no value is present for NumCodedStripes, not even an explicit nil
### GetNumDataStripes

`func (o *ErasureCodingInfo) GetNumDataStripes() int32`

GetNumDataStripes returns the NumDataStripes field if non-nil, zero value otherwise.

### GetNumDataStripesOk

`func (o *ErasureCodingInfo) GetNumDataStripesOk() (*int32, bool)`

GetNumDataStripesOk returns a tuple with the NumDataStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDataStripes

`func (o *ErasureCodingInfo) SetNumDataStripes(v int32)`

SetNumDataStripes sets NumDataStripes field to given value.

### HasNumDataStripes

`func (o *ErasureCodingInfo) HasNumDataStripes() bool`

HasNumDataStripes returns a boolean if a field has been set.

### SetNumDataStripesNil

`func (o *ErasureCodingInfo) SetNumDataStripesNil(b bool)`

 SetNumDataStripesNil sets the value for NumDataStripes to be an explicit nil

### UnsetNumDataStripes
`func (o *ErasureCodingInfo) UnsetNumDataStripes()`

UnsetNumDataStripes ensures that no value is present for NumDataStripes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


