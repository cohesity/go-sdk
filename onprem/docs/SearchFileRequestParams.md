# SearchFileRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the files. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all files name are matched with the prefix string. | [optional] 
**Types** | Pointer to **[]string** | Specifies a list of file types. Only files within the given types will be returned. | [optional] 
**SourceEnvironments** | Pointer to **[]string** | Specifies a list of the source environments. Only files from these types of source will be returned. | [optional] 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 
**ObjectIds** | Pointer to **[]int64** | Specifies a list of object ids. Only files found in these objects will be returned. | [optional] 

## Methods

### NewSearchFileRequestParams

`func NewSearchFileRequestParams() *SearchFileRequestParams`

NewSearchFileRequestParams instantiates a new SearchFileRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFileRequestParamsWithDefaults

`func NewSearchFileRequestParamsWithDefaults() *SearchFileRequestParams`

NewSearchFileRequestParamsWithDefaults instantiates a new SearchFileRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *SearchFileRequestParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *SearchFileRequestParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *SearchFileRequestParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *SearchFileRequestParams) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### SetSearchStringNil

`func (o *SearchFileRequestParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *SearchFileRequestParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetTypes

`func (o *SearchFileRequestParams) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchFileRequestParams) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchFileRequestParams) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchFileRequestParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchFileRequestParams) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchFileRequestParams) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetSourceEnvironments

`func (o *SearchFileRequestParams) GetSourceEnvironments() []string`

GetSourceEnvironments returns the SourceEnvironments field if non-nil, zero value otherwise.

### GetSourceEnvironmentsOk

`func (o *SearchFileRequestParams) GetSourceEnvironmentsOk() (*[]string, bool)`

GetSourceEnvironmentsOk returns a tuple with the SourceEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironments

`func (o *SearchFileRequestParams) SetSourceEnvironments(v []string)`

SetSourceEnvironments sets SourceEnvironments field to given value.

### HasSourceEnvironments

`func (o *SearchFileRequestParams) HasSourceEnvironments() bool`

HasSourceEnvironments returns a boolean if a field has been set.

### SetSourceEnvironmentsNil

`func (o *SearchFileRequestParams) SetSourceEnvironmentsNil(b bool)`

 SetSourceEnvironmentsNil sets the value for SourceEnvironments to be an explicit nil

### UnsetSourceEnvironments
`func (o *SearchFileRequestParams) UnsetSourceEnvironments()`

UnsetSourceEnvironments ensures that no value is present for SourceEnvironments, not even an explicit nil
### GetSourceIds

`func (o *SearchFileRequestParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *SearchFileRequestParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *SearchFileRequestParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *SearchFileRequestParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *SearchFileRequestParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *SearchFileRequestParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil
### GetObjectIds

`func (o *SearchFileRequestParams) GetObjectIds() []int64`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *SearchFileRequestParams) GetObjectIdsOk() (*[]int64, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *SearchFileRequestParams) SetObjectIds(v []int64)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *SearchFileRequestParams) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### SetObjectIdsNil

`func (o *SearchFileRequestParams) SetObjectIdsNil(b bool)`

 SetObjectIdsNil sets the value for ObjectIds to be an explicit nil

### UnsetObjectIds
`func (o *SearchFileRequestParams) UnsetObjectIds()`

UnsetObjectIds ensures that no value is present for ObjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


