# FilePathFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFilters** | Pointer to **[]string** | Array of Excluded File Path Filters.  Specifies filters to match files or directories that should be removed from the list of objects matching ProtectFilters. | [optional] 
**ProtectFilters** | Pointer to **[]string** | Array of Protected File Path Filters.  Specifies filters to match files or directories that should be protected. | [optional] 

## Methods

### NewFilePathFilter

`func NewFilePathFilter() *FilePathFilter`

NewFilePathFilter instantiates a new FilePathFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePathFilterWithDefaults

`func NewFilePathFilterWithDefaults() *FilePathFilter`

NewFilePathFilterWithDefaults instantiates a new FilePathFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFilters

`func (o *FilePathFilter) GetExcludeFilters() []string`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *FilePathFilter) GetExcludeFiltersOk() (*[]string, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *FilePathFilter) SetExcludeFilters(v []string)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *FilePathFilter) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *FilePathFilter) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *FilePathFilter) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetProtectFilters

`func (o *FilePathFilter) GetProtectFilters() []string`

GetProtectFilters returns the ProtectFilters field if non-nil, zero value otherwise.

### GetProtectFiltersOk

`func (o *FilePathFilter) GetProtectFiltersOk() (*[]string, bool)`

GetProtectFiltersOk returns a tuple with the ProtectFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectFilters

`func (o *FilePathFilter) SetProtectFilters(v []string)`

SetProtectFilters sets ProtectFilters field to given value.

### HasProtectFilters

`func (o *FilePathFilter) HasProtectFilters() bool`

HasProtectFilters returns a boolean if a field has been set.

### SetProtectFiltersNil

`func (o *FilePathFilter) SetProtectFiltersNil(b bool)`

 SetProtectFiltersNil sets the value for ProtectFilters to be an explicit nil

### UnsetProtectFilters
`func (o *FilePathFilter) UnsetProtectFilters()`

UnsetProtectFilters ensures that no value is present for ProtectFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


