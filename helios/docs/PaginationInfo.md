# PaginationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies a cookie which can be passed in by the user in order to retrieve the next page of results. | [optional] 

## Methods

### NewPaginationInfo

`func NewPaginationInfo() *PaginationInfo`

NewPaginationInfo instantiates a new PaginationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationInfoWithDefaults

`func NewPaginationInfoWithDefaults() *PaginationInfo`

NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *PaginationInfo) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *PaginationInfo) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *PaginationInfo) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *PaginationInfo) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *PaginationInfo) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *PaginationInfo) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


