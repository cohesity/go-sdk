# HDFSOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsTypes** | **[]string** | Specifies types as Folders or Files or both to be searched. | 
**SearchString** | **NullableString** | Specifies the search string to search the HDFS Folders and Files. | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewHDFSOnPremSearchParams

`func NewHDFSOnPremSearchParams(hdfsTypes []string, searchString NullableString, ) *HDFSOnPremSearchParams`

NewHDFSOnPremSearchParams instantiates a new HDFSOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHDFSOnPremSearchParamsWithDefaults

`func NewHDFSOnPremSearchParamsWithDefaults() *HDFSOnPremSearchParams`

NewHDFSOnPremSearchParamsWithDefaults instantiates a new HDFSOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsTypes

`func (o *HDFSOnPremSearchParams) GetHdfsTypes() []string`

GetHdfsTypes returns the HdfsTypes field if non-nil, zero value otherwise.

### GetHdfsTypesOk

`func (o *HDFSOnPremSearchParams) GetHdfsTypesOk() (*[]string, bool)`

GetHdfsTypesOk returns a tuple with the HdfsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsTypes

`func (o *HDFSOnPremSearchParams) SetHdfsTypes(v []string)`

SetHdfsTypes sets HdfsTypes field to given value.


### GetSearchString

`func (o *HDFSOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HDFSOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HDFSOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HDFSOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HDFSOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSourceIds

`func (o *HDFSOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *HDFSOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *HDFSOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *HDFSOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *HDFSOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *HDFSOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


