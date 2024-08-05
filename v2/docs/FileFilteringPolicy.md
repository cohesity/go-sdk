# FileFilteringPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeList** | Pointer to **[]string** | Specifies the list of excluded files for this protection Protection Group. Exclude filters have a higher priority than include filters. | [optional] 
**IncludeList** | Pointer to **[]string** | Specifies the list of included files for this Protection Group. | [optional] 

## Methods

### NewFileFilteringPolicy

`func NewFileFilteringPolicy() *FileFilteringPolicy`

NewFileFilteringPolicy instantiates a new FileFilteringPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFilteringPolicyWithDefaults

`func NewFileFilteringPolicyWithDefaults() *FileFilteringPolicy`

NewFileFilteringPolicyWithDefaults instantiates a new FileFilteringPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeList

`func (o *FileFilteringPolicy) GetExcludeList() []string`

GetExcludeList returns the ExcludeList field if non-nil, zero value otherwise.

### GetExcludeListOk

`func (o *FileFilteringPolicy) GetExcludeListOk() (*[]string, bool)`

GetExcludeListOk returns a tuple with the ExcludeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeList

`func (o *FileFilteringPolicy) SetExcludeList(v []string)`

SetExcludeList sets ExcludeList field to given value.

### HasExcludeList

`func (o *FileFilteringPolicy) HasExcludeList() bool`

HasExcludeList returns a boolean if a field has been set.

### GetIncludeList

`func (o *FileFilteringPolicy) GetIncludeList() []string`

GetIncludeList returns the IncludeList field if non-nil, zero value otherwise.

### GetIncludeListOk

`func (o *FileFilteringPolicy) GetIncludeListOk() (*[]string, bool)`

GetIncludeListOk returns a tuple with the IncludeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeList

`func (o *FileFilteringPolicy) SetIncludeList(v []string)`

SetIncludeList sets IncludeList field to given value.

### HasIncludeList

`func (o *FileFilteringPolicy) HasIncludeList() bool`

HasIncludeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


