# SearchGraphNodesResponseParamsGraphNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**NullableError**](Error.md) | Specifies the error information when node information could not be fetched. In case of error graphNodeInfo will only have the node id. | [optional] 
**GraphNodeInfo** | Pointer to [**NullableGraphNode**](GraphNode.md) | Specifies fetched node information. | [optional] 

## Methods

### NewSearchGraphNodesResponseParamsGraphNodesInner

`func NewSearchGraphNodesResponseParamsGraphNodesInner() *SearchGraphNodesResponseParamsGraphNodesInner`

NewSearchGraphNodesResponseParamsGraphNodesInner instantiates a new SearchGraphNodesResponseParamsGraphNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchGraphNodesResponseParamsGraphNodesInnerWithDefaults

`func NewSearchGraphNodesResponseParamsGraphNodesInnerWithDefaults() *SearchGraphNodesResponseParamsGraphNodesInner`

NewSearchGraphNodesResponseParamsGraphNodesInnerWithDefaults instantiates a new SearchGraphNodesResponseParamsGraphNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SearchGraphNodesResponseParamsGraphNodesInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetGraphNodeInfo

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) GetGraphNodeInfo() GraphNode`

GetGraphNodeInfo returns the GraphNodeInfo field if non-nil, zero value otherwise.

### GetGraphNodeInfoOk

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) GetGraphNodeInfoOk() (*GraphNode, bool)`

GetGraphNodeInfoOk returns a tuple with the GraphNodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphNodeInfo

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) SetGraphNodeInfo(v GraphNode)`

SetGraphNodeInfo sets GraphNodeInfo field to given value.

### HasGraphNodeInfo

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) HasGraphNodeInfo() bool`

HasGraphNodeInfo returns a boolean if a field has been set.

### SetGraphNodeInfoNil

`func (o *SearchGraphNodesResponseParamsGraphNodesInner) SetGraphNodeInfoNil(b bool)`

 SetGraphNodeInfoNil sets the value for GraphNodeInfo to be an explicit nil

### UnsetGraphNodeInfo
`func (o *SearchGraphNodesResponseParamsGraphNodesInner) UnsetGraphNodeInfo()`

UnsetGraphNodeInfo ensures that no value is present for GraphNodeInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


