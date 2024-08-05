# StoragePolicyCompressionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineEnabled** | Pointer to **NullableBool** | Specifies whether inline compression is enabled. This field is appliciable only if inlineDeduplicationEnabled is set to true and compression is enabled. | [optional] 
**Type** | Pointer to **NullableString** | Specifies copmpression type for a Storage Domain. | [optional] 

## Methods

### NewStoragePolicyCompressionParams

`func NewStoragePolicyCompressionParams() *StoragePolicyCompressionParams`

NewStoragePolicyCompressionParams instantiates a new StoragePolicyCompressionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyCompressionParamsWithDefaults

`func NewStoragePolicyCompressionParamsWithDefaults() *StoragePolicyCompressionParams`

NewStoragePolicyCompressionParamsWithDefaults instantiates a new StoragePolicyCompressionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineEnabled

`func (o *StoragePolicyCompressionParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *StoragePolicyCompressionParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *StoragePolicyCompressionParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *StoragePolicyCompressionParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *StoragePolicyCompressionParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *StoragePolicyCompressionParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
### GetType

`func (o *StoragePolicyCompressionParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoragePolicyCompressionParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoragePolicyCompressionParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoragePolicyCompressionParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *StoragePolicyCompressionParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StoragePolicyCompressionParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


