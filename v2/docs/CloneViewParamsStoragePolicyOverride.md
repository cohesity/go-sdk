# CloneViewParamsStoragePolicyOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableDedup** | Pointer to **NullableBool** | If it is set to true, we would disable dedup for writes made in this view irrespective of the view box&#39;s storage policy. | [optional] 
**DisableInlineDedupAndCompression** | Pointer to **NullableBool** | If false, the inline deduplication and compression settings inherited from the Storage Domain (View Box) apply to this View. If true, both inline deduplication and compression are disabled for this View. This can only be set to true if inline deduplication is set for the Storage Domain (View Box). | [optional] 

## Methods

### NewCloneViewParamsStoragePolicyOverride

`func NewCloneViewParamsStoragePolicyOverride() *CloneViewParamsStoragePolicyOverride`

NewCloneViewParamsStoragePolicyOverride instantiates a new CloneViewParamsStoragePolicyOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneViewParamsStoragePolicyOverrideWithDefaults

`func NewCloneViewParamsStoragePolicyOverrideWithDefaults() *CloneViewParamsStoragePolicyOverride`

NewCloneViewParamsStoragePolicyOverrideWithDefaults instantiates a new CloneViewParamsStoragePolicyOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableDedup

`func (o *CloneViewParamsStoragePolicyOverride) GetDisableDedup() bool`

GetDisableDedup returns the DisableDedup field if non-nil, zero value otherwise.

### GetDisableDedupOk

`func (o *CloneViewParamsStoragePolicyOverride) GetDisableDedupOk() (*bool, bool)`

GetDisableDedupOk returns a tuple with the DisableDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDedup

`func (o *CloneViewParamsStoragePolicyOverride) SetDisableDedup(v bool)`

SetDisableDedup sets DisableDedup field to given value.

### HasDisableDedup

`func (o *CloneViewParamsStoragePolicyOverride) HasDisableDedup() bool`

HasDisableDedup returns a boolean if a field has been set.

### SetDisableDedupNil

`func (o *CloneViewParamsStoragePolicyOverride) SetDisableDedupNil(b bool)`

 SetDisableDedupNil sets the value for DisableDedup to be an explicit nil

### UnsetDisableDedup
`func (o *CloneViewParamsStoragePolicyOverride) UnsetDisableDedup()`

UnsetDisableDedup ensures that no value is present for DisableDedup, not even an explicit nil
### GetDisableInlineDedupAndCompression

`func (o *CloneViewParamsStoragePolicyOverride) GetDisableInlineDedupAndCompression() bool`

GetDisableInlineDedupAndCompression returns the DisableInlineDedupAndCompression field if non-nil, zero value otherwise.

### GetDisableInlineDedupAndCompressionOk

`func (o *CloneViewParamsStoragePolicyOverride) GetDisableInlineDedupAndCompressionOk() (*bool, bool)`

GetDisableInlineDedupAndCompressionOk returns a tuple with the DisableInlineDedupAndCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInlineDedupAndCompression

`func (o *CloneViewParamsStoragePolicyOverride) SetDisableInlineDedupAndCompression(v bool)`

SetDisableInlineDedupAndCompression sets DisableInlineDedupAndCompression field to given value.

### HasDisableInlineDedupAndCompression

`func (o *CloneViewParamsStoragePolicyOverride) HasDisableInlineDedupAndCompression() bool`

HasDisableInlineDedupAndCompression returns a boolean if a field has been set.

### SetDisableInlineDedupAndCompressionNil

`func (o *CloneViewParamsStoragePolicyOverride) SetDisableInlineDedupAndCompressionNil(b bool)`

 SetDisableInlineDedupAndCompressionNil sets the value for DisableInlineDedupAndCompression to be an explicit nil

### UnsetDisableInlineDedupAndCompression
`func (o *CloneViewParamsStoragePolicyOverride) UnsetDisableInlineDedupAndCompression()`

UnsetDisableInlineDedupAndCompression ensures that no value is present for DisableInlineDedupAndCompression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


