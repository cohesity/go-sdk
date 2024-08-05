# SearchFileRequestParamsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the files. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all files name are matched with the prefix string. | [optional] 
**SourceEnvironments** | Pointer to **[]string** | Specifies a list of the source environments. Only files from these types of source will be returned. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of file types. Only files within the given types will be returned. | [optional] 

## Methods

### NewSearchFileRequestParamsBase

`func NewSearchFileRequestParamsBase() *SearchFileRequestParamsBase`

NewSearchFileRequestParamsBase instantiates a new SearchFileRequestParamsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFileRequestParamsBaseWithDefaults

`func NewSearchFileRequestParamsBaseWithDefaults() *SearchFileRequestParamsBase`

NewSearchFileRequestParamsBaseWithDefaults instantiates a new SearchFileRequestParamsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *SearchFileRequestParamsBase) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *SearchFileRequestParamsBase) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *SearchFileRequestParamsBase) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *SearchFileRequestParamsBase) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### SetSearchStringNil

`func (o *SearchFileRequestParamsBase) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *SearchFileRequestParamsBase) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSourceEnvironments

`func (o *SearchFileRequestParamsBase) GetSourceEnvironments() []string`

GetSourceEnvironments returns the SourceEnvironments field if non-nil, zero value otherwise.

### GetSourceEnvironmentsOk

`func (o *SearchFileRequestParamsBase) GetSourceEnvironmentsOk() (*[]string, bool)`

GetSourceEnvironmentsOk returns a tuple with the SourceEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironments

`func (o *SearchFileRequestParamsBase) SetSourceEnvironments(v []string)`

SetSourceEnvironments sets SourceEnvironments field to given value.

### HasSourceEnvironments

`func (o *SearchFileRequestParamsBase) HasSourceEnvironments() bool`

HasSourceEnvironments returns a boolean if a field has been set.

### SetSourceEnvironmentsNil

`func (o *SearchFileRequestParamsBase) SetSourceEnvironmentsNil(b bool)`

 SetSourceEnvironmentsNil sets the value for SourceEnvironments to be an explicit nil

### UnsetSourceEnvironments
`func (o *SearchFileRequestParamsBase) UnsetSourceEnvironments()`

UnsetSourceEnvironments ensures that no value is present for SourceEnvironments, not even an explicit nil
### GetTypes

`func (o *SearchFileRequestParamsBase) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchFileRequestParamsBase) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchFileRequestParamsBase) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchFileRequestParamsBase) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchFileRequestParamsBase) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchFileRequestParamsBase) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


