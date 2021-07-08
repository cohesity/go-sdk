# ClusterConfigProtoStoragePolicyOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableInlineDedupAndCompression** | Pointer to **NullableBool** | If this is set to true, we will not do inline dedup and compression even if deduplicate_compress_delay_secs is set to 0 in the view box&#39;s storage policy. | [optional] 

## Methods

### NewClusterConfigProtoStoragePolicyOverride

`func NewClusterConfigProtoStoragePolicyOverride() *ClusterConfigProtoStoragePolicyOverride`

NewClusterConfigProtoStoragePolicyOverride instantiates a new ClusterConfigProtoStoragePolicyOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigProtoStoragePolicyOverrideWithDefaults

`func NewClusterConfigProtoStoragePolicyOverrideWithDefaults() *ClusterConfigProtoStoragePolicyOverride`

NewClusterConfigProtoStoragePolicyOverrideWithDefaults instantiates a new ClusterConfigProtoStoragePolicyOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableInlineDedupAndCompression

`func (o *ClusterConfigProtoStoragePolicyOverride) GetDisableInlineDedupAndCompression() bool`

GetDisableInlineDedupAndCompression returns the DisableInlineDedupAndCompression field if non-nil, zero value otherwise.

### GetDisableInlineDedupAndCompressionOk

`func (o *ClusterConfigProtoStoragePolicyOverride) GetDisableInlineDedupAndCompressionOk() (*bool, bool)`

GetDisableInlineDedupAndCompressionOk returns a tuple with the DisableInlineDedupAndCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInlineDedupAndCompression

`func (o *ClusterConfigProtoStoragePolicyOverride) SetDisableInlineDedupAndCompression(v bool)`

SetDisableInlineDedupAndCompression sets DisableInlineDedupAndCompression field to given value.

### HasDisableInlineDedupAndCompression

`func (o *ClusterConfigProtoStoragePolicyOverride) HasDisableInlineDedupAndCompression() bool`

HasDisableInlineDedupAndCompression returns a boolean if a field has been set.

### SetDisableInlineDedupAndCompressionNil

`func (o *ClusterConfigProtoStoragePolicyOverride) SetDisableInlineDedupAndCompressionNil(b bool)`

 SetDisableInlineDedupAndCompressionNil sets the value for DisableInlineDedupAndCompression to be an explicit nil

### UnsetDisableInlineDedupAndCompression
`func (o *ClusterConfigProtoStoragePolicyOverride) UnsetDisableInlineDedupAndCompression()`

UnsetDisableInlineDedupAndCompression ensures that no value is present for DisableInlineDedupAndCompression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


