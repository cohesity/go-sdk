# DeduplicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Specifies whether deduplication is enabled on this Storage Domain. If enabled, cohesity cluster will eliminate duplicate blocks and thus reducing the amount of storage space. | [optional] 
**InlineEnabled** | Pointer to **NullableBool** | Specifies if inline deduplication is enabled. This field is appliciable only if deduplicationEnabled is set to true. | [optional] 

## Methods

### NewDeduplicationParams

`func NewDeduplicationParams() *DeduplicationParams`

NewDeduplicationParams instantiates a new DeduplicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeduplicationParamsWithDefaults

`func NewDeduplicationParamsWithDefaults() *DeduplicationParams`

NewDeduplicationParamsWithDefaults instantiates a new DeduplicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeduplicationParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeduplicationParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeduplicationParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeduplicationParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *DeduplicationParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *DeduplicationParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInlineEnabled

`func (o *DeduplicationParams) GetInlineEnabled() bool`

GetInlineEnabled returns the InlineEnabled field if non-nil, zero value otherwise.

### GetInlineEnabledOk

`func (o *DeduplicationParams) GetInlineEnabledOk() (*bool, bool)`

GetInlineEnabledOk returns a tuple with the InlineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineEnabled

`func (o *DeduplicationParams) SetInlineEnabled(v bool)`

SetInlineEnabled sets InlineEnabled field to given value.

### HasInlineEnabled

`func (o *DeduplicationParams) HasInlineEnabled() bool`

HasInlineEnabled returns a boolean if a field has been set.

### SetInlineEnabledNil

`func (o *DeduplicationParams) SetInlineEnabledNil(b bool)`

 SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil

### UnsetInlineEnabled
`func (o *DeduplicationParams) UnsetInlineEnabled()`

UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


