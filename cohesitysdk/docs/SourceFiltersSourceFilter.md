# SourceFiltersSourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRegex** | Pointer to **NullableBool** | If true, this implies &#39;source_filter&#39; is a regex filter. If false, it will be treated as wildcard/plain text filter. | [optional] 
**SourceFilter** | Pointer to **NullableString** | This contains the filter string. | [optional] 

## Methods

### NewSourceFiltersSourceFilter

`func NewSourceFiltersSourceFilter() *SourceFiltersSourceFilter`

NewSourceFiltersSourceFilter instantiates a new SourceFiltersSourceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceFiltersSourceFilterWithDefaults

`func NewSourceFiltersSourceFilterWithDefaults() *SourceFiltersSourceFilter`

NewSourceFiltersSourceFilterWithDefaults instantiates a new SourceFiltersSourceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRegex

`func (o *SourceFiltersSourceFilter) GetIsRegex() bool`

GetIsRegex returns the IsRegex field if non-nil, zero value otherwise.

### GetIsRegexOk

`func (o *SourceFiltersSourceFilter) GetIsRegexOk() (*bool, bool)`

GetIsRegexOk returns a tuple with the IsRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegex

`func (o *SourceFiltersSourceFilter) SetIsRegex(v bool)`

SetIsRegex sets IsRegex field to given value.

### HasIsRegex

`func (o *SourceFiltersSourceFilter) HasIsRegex() bool`

HasIsRegex returns a boolean if a field has been set.

### SetIsRegexNil

`func (o *SourceFiltersSourceFilter) SetIsRegexNil(b bool)`

 SetIsRegexNil sets the value for IsRegex to be an explicit nil

### UnsetIsRegex
`func (o *SourceFiltersSourceFilter) UnsetIsRegex()`

UnsetIsRegex ensures that no value is present for IsRegex, not even an explicit nil
### GetSourceFilter

`func (o *SourceFiltersSourceFilter) GetSourceFilter() string`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *SourceFiltersSourceFilter) GetSourceFilterOk() (*string, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *SourceFiltersSourceFilter) SetSourceFilter(v string)`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *SourceFiltersSourceFilter) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### SetSourceFilterNil

`func (o *SourceFiltersSourceFilter) SetSourceFilterNil(b bool)`

 SetSourceFilterNil sets the value for SourceFilter to be an explicit nil

### UnsetSourceFilter
`func (o *SourceFiltersSourceFilter) UnsetSourceFilter()`

UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


