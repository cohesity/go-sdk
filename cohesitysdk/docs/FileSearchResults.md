# FileSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]FileSearchResult**](FileSearchResult.md) | Array of Files and Folders.  Specifies the list of files and folders returned by this request that match the specified search and filter criteria. The number of files returned is limited by the pageCount field. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies cookie for resuming search if pagination is being used. For Librarian queries only. | [optional] 
**TotalCount** | Pointer to **NullableInt64** | Specifies the total number of files and folders that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 

## Methods

### NewFileSearchResults

`func NewFileSearchResults() *FileSearchResults`

NewFileSearchResults instantiates a new FileSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSearchResultsWithDefaults

`func NewFileSearchResultsWithDefaults() *FileSearchResults`

NewFileSearchResultsWithDefaults instantiates a new FileSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FileSearchResults) GetFiles() []FileSearchResult`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileSearchResults) GetFilesOk() (*[]FileSearchResult, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileSearchResults) SetFiles(v []FileSearchResult)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileSearchResults) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FileSearchResults) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FileSearchResults) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetPaginationCookie

`func (o *FileSearchResults) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *FileSearchResults) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *FileSearchResults) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *FileSearchResults) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *FileSearchResults) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *FileSearchResults) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetTotalCount

`func (o *FileSearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FileSearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FileSearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FileSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *FileSearchResults) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *FileSearchResults) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


