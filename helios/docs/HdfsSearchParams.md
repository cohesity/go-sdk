# HdfsSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **NullableString** | Specifies the search string to search the HDFS Folders and Files. | 
**HdfsTypes** | **[]string** | Specifies types as Folders or Files or both to be searched. | 

## Methods

### NewHdfsSearchParams

`func NewHdfsSearchParams(searchString NullableString, hdfsTypes []string, ) *HdfsSearchParams`

NewHdfsSearchParams instantiates a new HdfsSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsSearchParamsWithDefaults

`func NewHdfsSearchParamsWithDefaults() *HdfsSearchParams`

NewHdfsSearchParamsWithDefaults instantiates a new HdfsSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *HdfsSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HdfsSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HdfsSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HdfsSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HdfsSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetHdfsTypes

`func (o *HdfsSearchParams) GetHdfsTypes() []string`

GetHdfsTypes returns the HdfsTypes field if non-nil, zero value otherwise.

### GetHdfsTypesOk

`func (o *HdfsSearchParams) GetHdfsTypesOk() (*[]string, bool)`

GetHdfsTypesOk returns a tuple with the HdfsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsTypes

`func (o *HdfsSearchParams) SetHdfsTypes(v []string)`

SetHdfsTypes sets HdfsTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


