# SearchGraphNodesResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraphNodes** | Pointer to [**[]SearchGraphNodesResponseParamsGraphNodesInner**](SearchGraphNodesResponseParamsGraphNodesInner.md) | Specifies list of graph nodes. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 

## Methods

### NewSearchGraphNodesResponseParams

`func NewSearchGraphNodesResponseParams() *SearchGraphNodesResponseParams`

NewSearchGraphNodesResponseParams instantiates a new SearchGraphNodesResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchGraphNodesResponseParamsWithDefaults

`func NewSearchGraphNodesResponseParamsWithDefaults() *SearchGraphNodesResponseParams`

NewSearchGraphNodesResponseParamsWithDefaults instantiates a new SearchGraphNodesResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphNodes

`func (o *SearchGraphNodesResponseParams) GetGraphNodes() []SearchGraphNodesResponseParamsGraphNodesInner`

GetGraphNodes returns the GraphNodes field if non-nil, zero value otherwise.

### GetGraphNodesOk

`func (o *SearchGraphNodesResponseParams) GetGraphNodesOk() (*[]SearchGraphNodesResponseParamsGraphNodesInner, bool)`

GetGraphNodesOk returns a tuple with the GraphNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphNodes

`func (o *SearchGraphNodesResponseParams) SetGraphNodes(v []SearchGraphNodesResponseParamsGraphNodesInner)`

SetGraphNodes sets GraphNodes field to given value.

### HasGraphNodes

`func (o *SearchGraphNodesResponseParams) HasGraphNodes() bool`

HasGraphNodes returns a boolean if a field has been set.

### SetGraphNodesNil

`func (o *SearchGraphNodesResponseParams) SetGraphNodesNil(b bool)`

 SetGraphNodesNil sets the value for GraphNodes to be an explicit nil

### UnsetGraphNodes
`func (o *SearchGraphNodesResponseParams) UnsetGraphNodes()`

UnsetGraphNodes ensures that no value is present for GraphNodes, not even an explicit nil
### GetPaginationCookie

`func (o *SearchGraphNodesResponseParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *SearchGraphNodesResponseParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *SearchGraphNodesResponseParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *SearchGraphNodesResponseParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *SearchGraphNodesResponseParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *SearchGraphNodesResponseParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


