# QueryGraphNodesDiffParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Specifies the number of objects to be fetched for the specified pagination cookie. | [optional] 
**DiffTypes** | Pointer to **[]string** | Specifies an optional mask to filter only certain kinds of diffs. Supported diff types - Added/Modified/Deleted/Unmodified | [optional] 
**NodeFilter** | Pointer to [**GraphNodeFilterParams**](GraphNodeFilterParams.md) | Specifies the filter params for the node to be fetched. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies a cookie which can be passed in by the user in order to retrieve the next page of results. | [optional] 
**SessionId** | Pointer to **NullableString** | Specifies the id of the session for which diff of nodes has to be fetched. | [optional] 

## Methods

### NewQueryGraphNodesDiffParams

`func NewQueryGraphNodesDiffParams() *QueryGraphNodesDiffParams`

NewQueryGraphNodesDiffParams instantiates a new QueryGraphNodesDiffParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGraphNodesDiffParamsWithDefaults

`func NewQueryGraphNodesDiffParamsWithDefaults() *QueryGraphNodesDiffParams`

NewQueryGraphNodesDiffParamsWithDefaults instantiates a new QueryGraphNodesDiffParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *QueryGraphNodesDiffParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryGraphNodesDiffParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryGraphNodesDiffParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryGraphNodesDiffParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDiffTypes

`func (o *QueryGraphNodesDiffParams) GetDiffTypes() []string`

GetDiffTypes returns the DiffTypes field if non-nil, zero value otherwise.

### GetDiffTypesOk

`func (o *QueryGraphNodesDiffParams) GetDiffTypesOk() (*[]string, bool)`

GetDiffTypesOk returns a tuple with the DiffTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffTypes

`func (o *QueryGraphNodesDiffParams) SetDiffTypes(v []string)`

SetDiffTypes sets DiffTypes field to given value.

### HasDiffTypes

`func (o *QueryGraphNodesDiffParams) HasDiffTypes() bool`

HasDiffTypes returns a boolean if a field has been set.

### SetDiffTypesNil

`func (o *QueryGraphNodesDiffParams) SetDiffTypesNil(b bool)`

 SetDiffTypesNil sets the value for DiffTypes to be an explicit nil

### UnsetDiffTypes
`func (o *QueryGraphNodesDiffParams) UnsetDiffTypes()`

UnsetDiffTypes ensures that no value is present for DiffTypes, not even an explicit nil
### GetNodeFilter

`func (o *QueryGraphNodesDiffParams) GetNodeFilter() GraphNodeFilterParams`

GetNodeFilter returns the NodeFilter field if non-nil, zero value otherwise.

### GetNodeFilterOk

`func (o *QueryGraphNodesDiffParams) GetNodeFilterOk() (*GraphNodeFilterParams, bool)`

GetNodeFilterOk returns a tuple with the NodeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilter

`func (o *QueryGraphNodesDiffParams) SetNodeFilter(v GraphNodeFilterParams)`

SetNodeFilter sets NodeFilter field to given value.

### HasNodeFilter

`func (o *QueryGraphNodesDiffParams) HasNodeFilter() bool`

HasNodeFilter returns a boolean if a field has been set.

### GetPaginationCookie

`func (o *QueryGraphNodesDiffParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *QueryGraphNodesDiffParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *QueryGraphNodesDiffParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *QueryGraphNodesDiffParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *QueryGraphNodesDiffParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *QueryGraphNodesDiffParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetSessionId

`func (o *QueryGraphNodesDiffParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *QueryGraphNodesDiffParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *QueryGraphNodesDiffParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *QueryGraphNodesDiffParams) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *QueryGraphNodesDiffParams) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *QueryGraphNodesDiffParams) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


