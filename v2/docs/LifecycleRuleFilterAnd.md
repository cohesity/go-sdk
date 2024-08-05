# LifecycleRuleFilterAnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **NullableString** | Specifies a Prefix identifying one or more objects to which the rule applies. | [optional] 
**Tags** | Pointer to [**[]LifecycleRuleFilterTag**](LifecycleRuleFilterTag.md) | Specifies the tag in the object&#39;s tag set to which the rule applies. All of these tags must exist in the object&#39;s tag set in order for the rule to apply. | [optional] 

## Methods

### NewLifecycleRuleFilterAnd

`func NewLifecycleRuleFilterAnd() *LifecycleRuleFilterAnd`

NewLifecycleRuleFilterAnd instantiates a new LifecycleRuleFilterAnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleFilterAndWithDefaults

`func NewLifecycleRuleFilterAndWithDefaults() *LifecycleRuleFilterAnd`

NewLifecycleRuleFilterAndWithDefaults instantiates a new LifecycleRuleFilterAnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *LifecycleRuleFilterAnd) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LifecycleRuleFilterAnd) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LifecycleRuleFilterAnd) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LifecycleRuleFilterAnd) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *LifecycleRuleFilterAnd) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *LifecycleRuleFilterAnd) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetTags

`func (o *LifecycleRuleFilterAnd) GetTags() []LifecycleRuleFilterTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LifecycleRuleFilterAnd) GetTagsOk() (*[]LifecycleRuleFilterTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LifecycleRuleFilterAnd) SetTags(v []LifecycleRuleFilterTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LifecycleRuleFilterAnd) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *LifecycleRuleFilterAnd) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *LifecycleRuleFilterAnd) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


