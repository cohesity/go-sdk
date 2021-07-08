# SourceFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeSourceFilterVec** | Pointer to [**[]SourceFiltersSourceFilter**](SourceFiltersSourceFilter.md) | This contains the list of exclude filters to be applied on the entities in the backup source. | [optional] 

## Methods

### NewSourceFilters

`func NewSourceFilters() *SourceFilters`

NewSourceFilters instantiates a new SourceFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceFiltersWithDefaults

`func NewSourceFiltersWithDefaults() *SourceFilters`

NewSourceFiltersWithDefaults instantiates a new SourceFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeSourceFilterVec

`func (o *SourceFilters) GetExcludeSourceFilterVec() []SourceFiltersSourceFilter`

GetExcludeSourceFilterVec returns the ExcludeSourceFilterVec field if non-nil, zero value otherwise.

### GetExcludeSourceFilterVecOk

`func (o *SourceFilters) GetExcludeSourceFilterVecOk() (*[]SourceFiltersSourceFilter, bool)`

GetExcludeSourceFilterVecOk returns a tuple with the ExcludeSourceFilterVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSourceFilterVec

`func (o *SourceFilters) SetExcludeSourceFilterVec(v []SourceFiltersSourceFilter)`

SetExcludeSourceFilterVec sets ExcludeSourceFilterVec field to given value.

### HasExcludeSourceFilterVec

`func (o *SourceFilters) HasExcludeSourceFilterVec() bool`

HasExcludeSourceFilterVec returns a boolean if a field has been set.

### SetExcludeSourceFilterVecNil

`func (o *SourceFilters) SetExcludeSourceFilterVecNil(b bool)`

 SetExcludeSourceFilterVecNil sets the value for ExcludeSourceFilterVec to be an explicit nil

### UnsetExcludeSourceFilterVec
`func (o *SourceFilters) UnsetExcludeSourceFilterVec()`

UnsetExcludeSourceFilterVec ensures that no value is present for ExcludeSourceFilterVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


