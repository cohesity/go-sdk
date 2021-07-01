# IndexingPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrefixes** | Pointer to **[]string** | List of directory prefixes to allow for indexing. | [optional] 
**DenyPrefixes** | Pointer to **[]string** | List of directory prefixes to filter out. | [optional] 
**DisableIndexing** | Pointer to **NullableBool** | If this field is set all the files in the VM will be filtered. | [optional] 

## Methods

### NewIndexingPolicyProto

`func NewIndexingPolicyProto() *IndexingPolicyProto`

NewIndexingPolicyProto instantiates a new IndexingPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingPolicyProtoWithDefaults

`func NewIndexingPolicyProtoWithDefaults() *IndexingPolicyProto`

NewIndexingPolicyProtoWithDefaults instantiates a new IndexingPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPrefixes

`func (o *IndexingPolicyProto) GetAllowPrefixes() []string`

GetAllowPrefixes returns the AllowPrefixes field if non-nil, zero value otherwise.

### GetAllowPrefixesOk

`func (o *IndexingPolicyProto) GetAllowPrefixesOk() (*[]string, bool)`

GetAllowPrefixesOk returns a tuple with the AllowPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrefixes

`func (o *IndexingPolicyProto) SetAllowPrefixes(v []string)`

SetAllowPrefixes sets AllowPrefixes field to given value.

### HasAllowPrefixes

`func (o *IndexingPolicyProto) HasAllowPrefixes() bool`

HasAllowPrefixes returns a boolean if a field has been set.

### SetAllowPrefixesNil

`func (o *IndexingPolicyProto) SetAllowPrefixesNil(b bool)`

 SetAllowPrefixesNil sets the value for AllowPrefixes to be an explicit nil

### UnsetAllowPrefixes
`func (o *IndexingPolicyProto) UnsetAllowPrefixes()`

UnsetAllowPrefixes ensures that no value is present for AllowPrefixes, not even an explicit nil
### GetDenyPrefixes

`func (o *IndexingPolicyProto) GetDenyPrefixes() []string`

GetDenyPrefixes returns the DenyPrefixes field if non-nil, zero value otherwise.

### GetDenyPrefixesOk

`func (o *IndexingPolicyProto) GetDenyPrefixesOk() (*[]string, bool)`

GetDenyPrefixesOk returns a tuple with the DenyPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPrefixes

`func (o *IndexingPolicyProto) SetDenyPrefixes(v []string)`

SetDenyPrefixes sets DenyPrefixes field to given value.

### HasDenyPrefixes

`func (o *IndexingPolicyProto) HasDenyPrefixes() bool`

HasDenyPrefixes returns a boolean if a field has been set.

### SetDenyPrefixesNil

`func (o *IndexingPolicyProto) SetDenyPrefixesNil(b bool)`

 SetDenyPrefixesNil sets the value for DenyPrefixes to be an explicit nil

### UnsetDenyPrefixes
`func (o *IndexingPolicyProto) UnsetDenyPrefixes()`

UnsetDenyPrefixes ensures that no value is present for DenyPrefixes, not even an explicit nil
### GetDisableIndexing

`func (o *IndexingPolicyProto) GetDisableIndexing() bool`

GetDisableIndexing returns the DisableIndexing field if non-nil, zero value otherwise.

### GetDisableIndexingOk

`func (o *IndexingPolicyProto) GetDisableIndexingOk() (*bool, bool)`

GetDisableIndexingOk returns a tuple with the DisableIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIndexing

`func (o *IndexingPolicyProto) SetDisableIndexing(v bool)`

SetDisableIndexing sets DisableIndexing field to given value.

### HasDisableIndexing

`func (o *IndexingPolicyProto) HasDisableIndexing() bool`

HasDisableIndexing returns a boolean if a field has been set.

### SetDisableIndexingNil

`func (o *IndexingPolicyProto) SetDisableIndexingNil(b bool)`

 SetDisableIndexingNil sets the value for DisableIndexing to be an explicit nil

### UnsetDisableIndexing
`func (o *IndexingPolicyProto) UnsetDisableIndexing()`

UnsetDisableIndexing ensures that no value is present for DisableIndexing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


