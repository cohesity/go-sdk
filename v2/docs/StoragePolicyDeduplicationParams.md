# StoragePolicyDeduplicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Specifies whether deduplication is enabled on this Storage Domain. If enabled, cohesity cluster will eliminate duplicate blocks and thus reducing the amount of storage space. | [optional] 
**InlineEnabled** | Pointer to **NullableBool** | Specifies if inline deduplication is enabled. This field is appliciable only if deduplicationEnabled is set to true. | [optional] 

## Methods

### NewStoragePolicyDeduplicationParams

`func NewStoragePolicyDeduplicationParams() *StoragePolicyDeduplicationParams`

NewStoragePolicyDeduplicationParams instantiates a new StoragePolicyDeduplicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyDeduplicationParamsWithDefaults

`func NewStoragePolicyDeduplicationParamsWithDefaults() *StoragePolicyDeduplicationParams`

NewStoragePolicyDeduplicationParamsWithDefaults instantiates a new StoragePolicyDeduplicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StoragePolicyDeduplicationParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoragePolicyDeduplicationParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoragePolicyDeduplicationParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StoragePolicyDeduplicationParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *StoragePolicyDeduplicationParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *StoragePolicyDeduplicationParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInlineEnabled

`func (o *StoragePolicyDeduplicationParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *StoragePolicyDeduplicationParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *StoragePolicyDeduplicationParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *StoragePolicyDeduplicationParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *StoragePolicyDeduplicationParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *StoragePolicyDeduplicationParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


