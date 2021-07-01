# FileFstatResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableInt32** | Cookie is used for paginating results. If ReadVMDirResult is returning partial results, this field will be set. Supplying this cookie will resume listing from where this result left off. | [optional] 
**FstatInfo** | Pointer to [**FileStatInfo**](FileStatInfo.md) |  | [optional] 

## Methods

### NewFileFstatResult

`func NewFileFstatResult() *FileFstatResult`

NewFileFstatResult instantiates a new FileFstatResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFstatResultWithDefaults

`func NewFileFstatResultWithDefaults() *FileFstatResult`

NewFileFstatResultWithDefaults instantiates a new FileFstatResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *FileFstatResult) GetCookie() int32`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *FileFstatResult) GetCookieOk() (*int32, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *FileFstatResult) SetCookie(v int32)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *FileFstatResult) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *FileFstatResult) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *FileFstatResult) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetFstatInfo

`func (o *FileFstatResult) GetFstatInfo() FileStatInfo`

GetFstatInfo returns the FstatInfo field if non-nil, zero value otherwise.

### GetFstatInfoOk

`func (o *FileFstatResult) GetFstatInfoOk() (*FileStatInfo, bool)`

GetFstatInfoOk returns a tuple with the FstatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstatInfo

`func (o *FileFstatResult) SetFstatInfo(v FileStatInfo)`

SetFstatInfo sets FstatInfo field to given value.

### HasFstatInfo

`func (o *FileFstatResult) HasFstatInfo() bool`

HasFstatInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


