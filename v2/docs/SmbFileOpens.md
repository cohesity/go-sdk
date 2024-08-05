# SmbFileOpens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFilePaths** | Pointer to [**[]SmbActiveFilePath**](SmbActiveFilePath.md) | Specifies the active opens for an SMB file in a view. | [optional] 
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie | [optional] 

## Methods

### NewSmbFileOpens

`func NewSmbFileOpens() *SmbFileOpens`

NewSmbFileOpens instantiates a new SmbFileOpens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbFileOpensWithDefaults

`func NewSmbFileOpensWithDefaults() *SmbFileOpens`

NewSmbFileOpensWithDefaults instantiates a new SmbFileOpens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFilePaths

`func (o *SmbFileOpens) GetActiveFilePaths() []SmbActiveFilePath`

GetActiveFilePaths returns the ActiveFilePaths field if non-nil, zero value otherwise.

### GetActiveFilePathsOk

`func (o *SmbFileOpens) GetActiveFilePathsOk() (*[]SmbActiveFilePath, bool)`

GetActiveFilePathsOk returns a tuple with the ActiveFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFilePaths

`func (o *SmbFileOpens) SetActiveFilePaths(v []SmbActiveFilePath)`

SetActiveFilePaths sets ActiveFilePaths field to given value.

### HasActiveFilePaths

`func (o *SmbFileOpens) HasActiveFilePaths() bool`

HasActiveFilePaths returns a boolean if a field has been set.

### SetActiveFilePathsNil

`func (o *SmbFileOpens) SetActiveFilePathsNil(b bool)`

 SetActiveFilePathsNil sets the value for ActiveFilePaths to be an explicit nil

### UnsetActiveFilePaths
`func (o *SmbFileOpens) UnsetActiveFilePaths()`

UnsetActiveFilePaths ensures that no value is present for ActiveFilePaths, not even an explicit nil
### GetCookie

`func (o *SmbFileOpens) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *SmbFileOpens) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *SmbFileOpens) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *SmbFileOpens) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *SmbFileOpens) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *SmbFileOpens) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


