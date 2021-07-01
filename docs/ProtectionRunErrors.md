# ProtectionRunErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]RequestError**](RequestError.md) | Specifies the list of errors encountered by a task during a protection run. | [optional] 
**FileNames** | Pointer to **[]string** | Specifies the list of filenames with errors encountered by a task during a protection run. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies the cookie for next set of results. | [optional] 

## Methods

### NewProtectionRunErrors

`func NewProtectionRunErrors() *ProtectionRunErrors`

NewProtectionRunErrors instantiates a new ProtectionRunErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunErrorsWithDefaults

`func NewProtectionRunErrorsWithDefaults() *ProtectionRunErrors`

NewProtectionRunErrorsWithDefaults instantiates a new ProtectionRunErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ProtectionRunErrors) GetErrors() []RequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProtectionRunErrors) GetErrorsOk() (*[]RequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProtectionRunErrors) SetErrors(v []RequestError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProtectionRunErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ProtectionRunErrors) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ProtectionRunErrors) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetFileNames

`func (o *ProtectionRunErrors) GetFileNames() []string`

GetFileNames returns the FileNames field if non-nil, zero value otherwise.

### GetFileNamesOk

`func (o *ProtectionRunErrors) GetFileNamesOk() (*[]string, bool)`

GetFileNamesOk returns a tuple with the FileNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNames

`func (o *ProtectionRunErrors) SetFileNames(v []string)`

SetFileNames sets FileNames field to given value.

### HasFileNames

`func (o *ProtectionRunErrors) HasFileNames() bool`

HasFileNames returns a boolean if a field has been set.

### SetFileNamesNil

`func (o *ProtectionRunErrors) SetFileNamesNil(b bool)`

 SetFileNamesNil sets the value for FileNames to be an explicit nil

### UnsetFileNames
`func (o *ProtectionRunErrors) UnsetFileNames()`

UnsetFileNames ensures that no value is present for FileNames, not even an explicit nil
### GetPaginationCookie

`func (o *ProtectionRunErrors) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *ProtectionRunErrors) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *ProtectionRunErrors) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *ProtectionRunErrors) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *ProtectionRunErrors) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *ProtectionRunErrors) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


