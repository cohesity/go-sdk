# QueryGraphNodesDiffResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffGraphNodes** | Pointer to [**[]QueryGraphNodesDiffResultDiffGraphNodesInner**](QueryGraphNodesDiffResultDiffGraphNodesInner.md) | Specifies the list of diff for all the nodes matching the filter | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 
**UnmodifiedGraphNodes** | Pointer to **[]string** | Specifies the list of all the graph node ids which are unmodified. | [optional] 

## Methods

### NewQueryGraphNodesDiffResult

`func NewQueryGraphNodesDiffResult() *QueryGraphNodesDiffResult`

NewQueryGraphNodesDiffResult instantiates a new QueryGraphNodesDiffResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGraphNodesDiffResultWithDefaults

`func NewQueryGraphNodesDiffResultWithDefaults() *QueryGraphNodesDiffResult`

NewQueryGraphNodesDiffResultWithDefaults instantiates a new QueryGraphNodesDiffResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffGraphNodes

`func (o *QueryGraphNodesDiffResult) GetDiffGraphNodes() []QueryGraphNodesDiffResultDiffGraphNodesInner`

GetDiffGraphNodes returns the DiffGraphNodes field if non-nil, zero value otherwise.

### GetDiffGraphNodesOk

`func (o *QueryGraphNodesDiffResult) GetDiffGraphNodesOk() (*[]QueryGraphNodesDiffResultDiffGraphNodesInner, bool)`

GetDiffGraphNodesOk returns a tuple with the DiffGraphNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffGraphNodes

`func (o *QueryGraphNodesDiffResult) SetDiffGraphNodes(v []QueryGraphNodesDiffResultDiffGraphNodesInner)`

SetDiffGraphNodes sets DiffGraphNodes field to given value.

### HasDiffGraphNodes

`func (o *QueryGraphNodesDiffResult) HasDiffGraphNodes() bool`

HasDiffGraphNodes returns a boolean if a field has been set.

### SetDiffGraphNodesNil

`func (o *QueryGraphNodesDiffResult) SetDiffGraphNodesNil(b bool)`

 SetDiffGraphNodesNil sets the value for DiffGraphNodes to be an explicit nil

### UnsetDiffGraphNodes
`func (o *QueryGraphNodesDiffResult) UnsetDiffGraphNodes()`

UnsetDiffGraphNodes ensures that no value is present for DiffGraphNodes, not even an explicit nil
### GetPaginationCookie

`func (o *QueryGraphNodesDiffResult) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *QueryGraphNodesDiffResult) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *QueryGraphNodesDiffResult) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *QueryGraphNodesDiffResult) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *QueryGraphNodesDiffResult) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *QueryGraphNodesDiffResult) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetUnmodifiedGraphNodes

`func (o *QueryGraphNodesDiffResult) GetUnmodifiedGraphNodes() []string`

GetUnmodifiedGraphNodes returns the UnmodifiedGraphNodes field if non-nil, zero value otherwise.

### GetUnmodifiedGraphNodesOk

`func (o *QueryGraphNodesDiffResult) GetUnmodifiedGraphNodesOk() (*[]string, bool)`

GetUnmodifiedGraphNodesOk returns a tuple with the UnmodifiedGraphNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmodifiedGraphNodes

`func (o *QueryGraphNodesDiffResult) SetUnmodifiedGraphNodes(v []string)`

SetUnmodifiedGraphNodes sets UnmodifiedGraphNodes field to given value.

### HasUnmodifiedGraphNodes

`func (o *QueryGraphNodesDiffResult) HasUnmodifiedGraphNodes() bool`

HasUnmodifiedGraphNodes returns a boolean if a field has been set.

### SetUnmodifiedGraphNodesNil

`func (o *QueryGraphNodesDiffResult) SetUnmodifiedGraphNodesNil(b bool)`

 SetUnmodifiedGraphNodesNil sets the value for UnmodifiedGraphNodes to be an explicit nil

### UnsetUnmodifiedGraphNodes
`func (o *QueryGraphNodesDiffResult) UnsetUnmodifiedGraphNodes()`

UnsetUnmodifiedGraphNodes ensures that no value is present for UnmodifiedGraphNodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


