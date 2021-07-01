# StoragePolicyOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableInlineDedupAndCompression** | Pointer to **NullableBool** | If false, the inline deduplication and compression settings inherited from the Storage Domain (View Box) apply to this View. If true, both inline deduplication and compression are disabled for this View. This can only be set to true if inline deduplication is set for the Storage Domain (View Box). | [optional] 

## Methods

### NewStoragePolicyOverride

`func NewStoragePolicyOverride() *StoragePolicyOverride`

NewStoragePolicyOverride instantiates a new StoragePolicyOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyOverrideWithDefaults

`func NewStoragePolicyOverrideWithDefaults() *StoragePolicyOverride`

NewStoragePolicyOverrideWithDefaults instantiates a new StoragePolicyOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableInlineDedupAndCompression

`func (o *StoragePolicyOverride) GetDisableInlineDedupAndCompression() bool`

GetDisableInlineDedupAndCompression returns the DisableInlineDedupAndCompression field if non-nil, zero value otherwise.

### GetDisableInlineDedupAndCompressionOk

`func (o *StoragePolicyOverride) GetDisableInlineDedupAndCompressionOk() (*bool, bool)`

GetDisableInlineDedupAndCompressionOk returns a tuple with the DisableInlineDedupAndCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInlineDedupAndCompression

`func (o *StoragePolicyOverride) SetDisableInlineDedupAndCompression(v bool)`

SetDisableInlineDedupAndCompression sets DisableInlineDedupAndCompression field to given value.

### HasDisableInlineDedupAndCompression

`func (o *StoragePolicyOverride) HasDisableInlineDedupAndCompression() bool`

HasDisableInlineDedupAndCompression returns a boolean if a field has been set.

### SetDisableInlineDedupAndCompressionNil

`func (o *StoragePolicyOverride) SetDisableInlineDedupAndCompressionNil(b bool)`

 SetDisableInlineDedupAndCompressionNil sets the value for DisableInlineDedupAndCompression to be an explicit nil

### UnsetDisableInlineDedupAndCompression
`func (o *StoragePolicyOverride) UnsetDisableInlineDedupAndCompression()`

UnsetDisableInlineDedupAndCompression ensures that no value is present for DisableInlineDedupAndCompression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


