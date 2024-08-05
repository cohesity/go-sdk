# GetDataAccessSessionsResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaginationCookie** | Pointer to **NullableString** | pecifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 
**Sessions** | Pointer to [**[]DataAccessSession**](DataAccessSession.md) | Specifies list of data access sessions. | [optional] 

## Methods

### NewGetDataAccessSessionsResponseParams

`func NewGetDataAccessSessionsResponseParams() *GetDataAccessSessionsResponseParams`

NewGetDataAccessSessionsResponseParams instantiates a new GetDataAccessSessionsResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDataAccessSessionsResponseParamsWithDefaults

`func NewGetDataAccessSessionsResponseParamsWithDefaults() *GetDataAccessSessionsResponseParams`

NewGetDataAccessSessionsResponseParamsWithDefaults instantiates a new GetDataAccessSessionsResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaginationCookie

`func (o *GetDataAccessSessionsResponseParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *GetDataAccessSessionsResponseParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *GetDataAccessSessionsResponseParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *GetDataAccessSessionsResponseParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *GetDataAccessSessionsResponseParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *GetDataAccessSessionsResponseParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetSessions

`func (o *GetDataAccessSessionsResponseParams) GetSessions() []DataAccessSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *GetDataAccessSessionsResponseParams) GetSessionsOk() (*[]DataAccessSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *GetDataAccessSessionsResponseParams) SetSessions(v []DataAccessSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *GetDataAccessSessionsResponseParams) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *GetDataAccessSessionsResponseParams) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *GetDataAccessSessionsResponseParams) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


