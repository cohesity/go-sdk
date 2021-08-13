# IndexingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableIndexing** | **bool** | Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed. | 
**IncludePaths** | Pointer to **[]string** | Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/&lt;wildcard&gt;/AppData | [optional] 
**ExcludePaths** | Pointer to **[]string** | Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/&lt;wildcard&gt;/AppData | [optional] 

## Methods

### NewIndexingPolicy

`func NewIndexingPolicy(enableIndexing bool, ) *IndexingPolicy`

NewIndexingPolicy instantiates a new IndexingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingPolicyWithDefaults

`func NewIndexingPolicyWithDefaults() *IndexingPolicy`

NewIndexingPolicyWithDefaults instantiates a new IndexingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableIndexing

`func (o *IndexingPolicy) GetEnableIndexing() bool`

GetEnableIndexing returns the EnableIndexing field if non-nil, zero value otherwise.

### GetEnableIndexingOk

`func (o *IndexingPolicy) GetEnableIndexingOk() (*bool, bool)`

GetEnableIndexingOk returns a tuple with the EnableIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIndexing

`func (o *IndexingPolicy) SetEnableIndexing(v bool)`

SetEnableIndexing sets EnableIndexing field to given value.


### GetIncludePaths

`func (o *IndexingPolicy) GetIncludePaths() []string`

GetIncludePaths returns the IncludePaths field if non-nil, zero value otherwise.

### GetIncludePathsOk

`func (o *IndexingPolicy) GetIncludePathsOk() (*[]string, bool)`

GetIncludePathsOk returns a tuple with the IncludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaths

`func (o *IndexingPolicy) SetIncludePaths(v []string)`

SetIncludePaths sets IncludePaths field to given value.

### HasIncludePaths

`func (o *IndexingPolicy) HasIncludePaths() bool`

HasIncludePaths returns a boolean if a field has been set.

### SetIncludePathsNil

`func (o *IndexingPolicy) SetIncludePathsNil(b bool)`

 SetIncludePathsNil sets the value for IncludePaths to be an explicit nil

### UnsetIncludePaths
`func (o *IndexingPolicy) UnsetIncludePaths()`

UnsetIncludePaths ensures that no value is present for IncludePaths, not even an explicit nil
### GetExcludePaths

`func (o *IndexingPolicy) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *IndexingPolicy) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *IndexingPolicy) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *IndexingPolicy) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### SetExcludePathsNil

`func (o *IndexingPolicy) SetExcludePathsNil(b bool)`

 SetExcludePathsNil sets the value for ExcludePaths to be an explicit nil

### UnsetExcludePaths
`func (o *IndexingPolicy) UnsetExcludePaths()`

UnsetExcludePaths ensures that no value is present for ExcludePaths, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


