# IndexingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrefixes** | Pointer to **[]string** | Array of Indexed Directories.  Specifies a list of directories to index. | [optional] 
**DenyPrefixes** | Pointer to **[]string** | Array of Excluded Directories.  Specifies a list of directories to exclude from indexing. | [optional] 
**DisableIndexing** | Pointer to **NullableBool** | Specifies if the files found in an Object (such as a VM) should be indexed. If false (the default), files are indexed. | [optional] 

## Methods

### NewIndexingPolicy

`func NewIndexingPolicy() *IndexingPolicy`

NewIndexingPolicy instantiates a new IndexingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingPolicyWithDefaults

`func NewIndexingPolicyWithDefaults() *IndexingPolicy`

NewIndexingPolicyWithDefaults instantiates a new IndexingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPrefixes

`func (o *IndexingPolicy) GetAllowPrefixes() []string`

GetAllowPrefixes returns the AllowPrefixes field if non-nil, zero value otherwise.

### GetAllowPrefixesOk

`func (o *IndexingPolicy) GetAllowPrefixesOk() (*[]string, bool)`

GetAllowPrefixesOk returns a tuple with the AllowPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrefixes

`func (o *IndexingPolicy) SetAllowPrefixes(v []string)`

SetAllowPrefixes sets AllowPrefixes field to given value.

### HasAllowPrefixes

`func (o *IndexingPolicy) HasAllowPrefixes() bool`

HasAllowPrefixes returns a boolean if a field has been set.

### SetAllowPrefixesNil

`func (o *IndexingPolicy) SetAllowPrefixesNil(b bool)`

 SetAllowPrefixesNil sets the value for AllowPrefixes to be an explicit nil

### UnsetAllowPrefixes
`func (o *IndexingPolicy) UnsetAllowPrefixes()`

UnsetAllowPrefixes ensures that no value is present for AllowPrefixes, not even an explicit nil
### GetDenyPrefixes

`func (o *IndexingPolicy) GetDenyPrefixes() []string`

GetDenyPrefixes returns the DenyPrefixes field if non-nil, zero value otherwise.

### GetDenyPrefixesOk

`func (o *IndexingPolicy) GetDenyPrefixesOk() (*[]string, bool)`

GetDenyPrefixesOk returns a tuple with the DenyPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPrefixes

`func (o *IndexingPolicy) SetDenyPrefixes(v []string)`

SetDenyPrefixes sets DenyPrefixes field to given value.

### HasDenyPrefixes

`func (o *IndexingPolicy) HasDenyPrefixes() bool`

HasDenyPrefixes returns a boolean if a field has been set.

### SetDenyPrefixesNil

`func (o *IndexingPolicy) SetDenyPrefixesNil(b bool)`

 SetDenyPrefixesNil sets the value for DenyPrefixes to be an explicit nil

### UnsetDenyPrefixes
`func (o *IndexingPolicy) UnsetDenyPrefixes()`

UnsetDenyPrefixes ensures that no value is present for DenyPrefixes, not even an explicit nil
### GetDisableIndexing

`func (o *IndexingPolicy) GetDisableIndexing() bool`

GetDisableIndexing returns the DisableIndexing field if non-nil, zero value otherwise.

### GetDisableIndexingOk

`func (o *IndexingPolicy) GetDisableIndexingOk() (*bool, bool)`

GetDisableIndexingOk returns a tuple with the DisableIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIndexing

`func (o *IndexingPolicy) SetDisableIndexing(v bool)`

SetDisableIndexing sets DisableIndexing field to given value.

### HasDisableIndexing

`func (o *IndexingPolicy) HasDisableIndexing() bool`

HasDisableIndexing returns a boolean if a field has been set.

### SetDisableIndexingNil

`func (o *IndexingPolicy) SetDisableIndexingNil(b bool)`

 SetDisableIndexingNil sets the value for DisableIndexing to be an explicit nil

### UnsetDisableIndexing
`func (o *IndexingPolicy) UnsetDisableIndexing()`

UnsetDisableIndexing ensures that no value is present for DisableIndexing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


