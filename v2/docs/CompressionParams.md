# CompressionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineEnabled** | Pointer to **NullableBool** | Specifies whether inline compression is enabled. This field is appliciable only if inlineDeduplicationEnabled is set to true and compression is enabled. | [optional] 
**Type** | Pointer to **NullableString** | Specifies copmpression type for a Storage Domain. | [optional] 

## Methods

### NewCompressionParams

`func NewCompressionParams() *CompressionParams`

NewCompressionParams instantiates a new CompressionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompressionParamsWithDefaults

`func NewCompressionParamsWithDefaults() *CompressionParams`

NewCompressionParamsWithDefaults instantiates a new CompressionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineEnabled

`func (o *CompressionParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *CompressionParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *CompressionParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *CompressionParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *CompressionParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *CompressionParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
### GetType

`func (o *CompressionParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompressionParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompressionParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompressionParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CompressionParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CompressionParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


