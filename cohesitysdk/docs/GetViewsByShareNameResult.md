# GetViewsByShareNameResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaginationCookie** | Pointer to **NullableString** | If set, i.e. there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetViewsByShare. | [optional] 
**SharesList** | Pointer to [**[]Share**](Share.md) | Array of Views and Aliases by Share name. Specifies the list of Views returned in this response. | [optional] 

## Methods

### NewGetViewsByShareNameResult

`func NewGetViewsByShareNameResult() *GetViewsByShareNameResult`

NewGetViewsByShareNameResult instantiates a new GetViewsByShareNameResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewsByShareNameResultWithDefaults

`func NewGetViewsByShareNameResultWithDefaults() *GetViewsByShareNameResult`

NewGetViewsByShareNameResultWithDefaults instantiates a new GetViewsByShareNameResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaginationCookie

`func (o *GetViewsByShareNameResult) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *GetViewsByShareNameResult) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *GetViewsByShareNameResult) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *GetViewsByShareNameResult) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *GetViewsByShareNameResult) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *GetViewsByShareNameResult) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetSharesList

`func (o *GetViewsByShareNameResult) GetSharesList() []Share`

GetSharesList returns the SharesList field if non-nil, zero value otherwise.

### GetSharesListOk

`func (o *GetViewsByShareNameResult) GetSharesListOk() (*[]Share, bool)`

GetSharesListOk returns a tuple with the SharesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesList

`func (o *GetViewsByShareNameResult) SetSharesList(v []Share)`

SetSharesList sets SharesList field to given value.

### HasSharesList

`func (o *GetViewsByShareNameResult) HasSharesList() bool`

HasSharesList returns a boolean if a field has been set.

### SetSharesListNil

`func (o *GetViewsByShareNameResult) SetSharesListNil(b bool)`

 SetSharesListNil sets the value for SharesList to be an explicit nil

### UnsetSharesList
`func (o *GetViewsByShareNameResult) UnsetSharesList()`

UnsetSharesList ensures that no value is present for SharesList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


