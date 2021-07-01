# InfectedFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the infected files. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | This cookie can be used in the succeeding call to list infected files to get the next set of infected files. If set to nil, it means that there&#39;s no more results that the server could provide. | [optional] 

## Methods

### NewInfectedFiles

`func NewInfectedFiles() *InfectedFiles`

NewInfectedFiles instantiates a new InfectedFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFilesWithDefaults

`func NewInfectedFilesWithDefaults() *InfectedFiles`

NewInfectedFilesWithDefaults instantiates a new InfectedFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedFiles

`func (o *InfectedFiles) GetInfectedFiles() []InfectedFile`

GetInfectedFiles returns the InfectedFiles field if non-nil, zero value otherwise.

### GetInfectedFilesOk

`func (o *InfectedFiles) GetInfectedFilesOk() (*[]InfectedFile, bool)`

GetInfectedFilesOk returns a tuple with the InfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFiles

`func (o *InfectedFiles) SetInfectedFiles(v []InfectedFile)`

SetInfectedFiles sets InfectedFiles field to given value.

### HasInfectedFiles

`func (o *InfectedFiles) HasInfectedFiles() bool`

HasInfectedFiles returns a boolean if a field has been set.

### SetInfectedFilesNil

`func (o *InfectedFiles) SetInfectedFilesNil(b bool)`

 SetInfectedFilesNil sets the value for InfectedFiles to be an explicit nil

### UnsetInfectedFiles
`func (o *InfectedFiles) UnsetInfectedFiles()`

UnsetInfectedFiles ensures that no value is present for InfectedFiles, not even an explicit nil
### GetPaginationCookie

`func (o *InfectedFiles) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *InfectedFiles) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *InfectedFiles) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *InfectedFiles) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *InfectedFiles) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *InfectedFiles) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


