# QueryGraphNodesDiffResultDiffGraphNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**NullableError**](Error.md) | Error is set if there was an error in looking up information for the node either in current or base snapshot. In case of error graphNodeInfo will only have the node id. | [optional] 
**GraphNodeInfo** | Pointer to [**NullableDiffGraphNode**](DiffGraphNode.md) | Specifies the list of diff for all the nodes added/ deleted/modified. | [optional] 

## Methods

### NewQueryGraphNodesDiffResultDiffGraphNodesInner

`func NewQueryGraphNodesDiffResultDiffGraphNodesInner() *QueryGraphNodesDiffResultDiffGraphNodesInner`

NewQueryGraphNodesDiffResultDiffGraphNodesInner instantiates a new QueryGraphNodesDiffResultDiffGraphNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGraphNodesDiffResultDiffGraphNodesInnerWithDefaults

`func NewQueryGraphNodesDiffResultDiffGraphNodesInnerWithDefaults() *QueryGraphNodesDiffResultDiffGraphNodesInner`

NewQueryGraphNodesDiffResultDiffGraphNodesInnerWithDefaults instantiates a new QueryGraphNodesDiffResultDiffGraphNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetGraphNodeInfo

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) GetGraphNodeInfo() DiffGraphNode`

GetGraphNodeInfo returns the GraphNodeInfo field if non-nil, zero value otherwise.

### GetGraphNodeInfoOk

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) GetGraphNodeInfoOk() (*DiffGraphNode, bool)`

GetGraphNodeInfoOk returns a tuple with the GraphNodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphNodeInfo

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) SetGraphNodeInfo(v DiffGraphNode)`

SetGraphNodeInfo sets GraphNodeInfo field to given value.

### HasGraphNodeInfo

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) HasGraphNodeInfo() bool`

HasGraphNodeInfo returns a boolean if a field has been set.

### SetGraphNodeInfoNil

`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) SetGraphNodeInfoNil(b bool)`

 SetGraphNodeInfoNil sets the value for GraphNodeInfo to be an explicit nil

### UnsetGraphNodeInfo
`func (o *QueryGraphNodesDiffResultDiffGraphNodesInner) UnsetGraphNodeInfo()`

UnsetGraphNodeInfo ensures that no value is present for GraphNodeInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


