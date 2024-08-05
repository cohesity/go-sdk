# GetNlmLocksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. | [optional] 
**FileNlmLocks** | Pointer to [**[]FileNlmLocks**](FileNlmLocks.md) | Specifies the list of NLM locks. | [optional] 

## Methods

### NewGetNlmLocksResult

`func NewGetNlmLocksResult() *GetNlmLocksResult`

NewGetNlmLocksResult instantiates a new GetNlmLocksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNlmLocksResultWithDefaults

`func NewGetNlmLocksResultWithDefaults() *GetNlmLocksResult`

NewGetNlmLocksResultWithDefaults instantiates a new GetNlmLocksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *GetNlmLocksResult) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *GetNlmLocksResult) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *GetNlmLocksResult) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *GetNlmLocksResult) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *GetNlmLocksResult) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *GetNlmLocksResult) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetFileNlmLocks

`func (o *GetNlmLocksResult) GetFileNlmLocks() []FileNlmLocks`

GetFileNlmLocks returns the FileNlmLocks field if non-nil, zero value otherwise.

### GetFileNlmLocksOk

`func (o *GetNlmLocksResult) GetFileNlmLocksOk() (*[]FileNlmLocks, bool)`

GetFileNlmLocksOk returns a tuple with the FileNlmLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNlmLocks

`func (o *GetNlmLocksResult) SetFileNlmLocks(v []FileNlmLocks)`

SetFileNlmLocks sets FileNlmLocks field to given value.

### HasFileNlmLocks

`func (o *GetNlmLocksResult) HasFileNlmLocks() bool`

HasFileNlmLocks returns a boolean if a field has been set.

### SetFileNlmLocksNil

`func (o *GetNlmLocksResult) SetFileNlmLocksNil(b bool)`

 SetFileNlmLocksNil sets the value for FileNlmLocks to be an explicit nil

### UnsetFileNlmLocks
`func (o *GetNlmLocksResult) UnsetFileNlmLocks()`

UnsetFileNlmLocks ensures that no value is present for FileNlmLocks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


