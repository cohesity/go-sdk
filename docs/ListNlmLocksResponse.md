# ListNlmLocksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies an opaque string to pass to get the next set of NLM locks. If null is returned, this response is the last set of NLM locks. | [optional] 
**FilesNlmLocks** | Pointer to [**[]FileNlmLocks**](FileNlmLocks.md) | Specifies the list of NLM locks. | [optional] 

## Methods

### NewListNlmLocksResponse

`func NewListNlmLocksResponse() *ListNlmLocksResponse`

NewListNlmLocksResponse instantiates a new ListNlmLocksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNlmLocksResponseWithDefaults

`func NewListNlmLocksResponseWithDefaults() *ListNlmLocksResponse`

NewListNlmLocksResponseWithDefaults instantiates a new ListNlmLocksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *ListNlmLocksResponse) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *ListNlmLocksResponse) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *ListNlmLocksResponse) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *ListNlmLocksResponse) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *ListNlmLocksResponse) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *ListNlmLocksResponse) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetFilesNlmLocks

`func (o *ListNlmLocksResponse) GetFilesNlmLocks() []FileNlmLocks`

GetFilesNlmLocks returns the FilesNlmLocks field if non-nil, zero value otherwise.

### GetFilesNlmLocksOk

`func (o *ListNlmLocksResponse) GetFilesNlmLocksOk() (*[]FileNlmLocks, bool)`

GetFilesNlmLocksOk returns a tuple with the FilesNlmLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesNlmLocks

`func (o *ListNlmLocksResponse) SetFilesNlmLocks(v []FileNlmLocks)`

SetFilesNlmLocks sets FilesNlmLocks field to given value.

### HasFilesNlmLocks

`func (o *ListNlmLocksResponse) HasFilesNlmLocks() bool`

HasFilesNlmLocks returns a boolean if a field has been set.

### SetFilesNlmLocksNil

`func (o *ListNlmLocksResponse) SetFilesNlmLocksNil(b bool)`

 SetFilesNlmLocksNil sets the value for FilesNlmLocks to be an explicit nil

### UnsetFilesNlmLocks
`func (o *ListNlmLocksResponse) UnsetFilesNlmLocks()`

UnsetFilesNlmLocks ensures that no value is present for FilesNlmLocks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


