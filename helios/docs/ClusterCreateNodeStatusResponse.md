# ClusterCreateNodeStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableInt64** | Specifies the id of the node. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the Ip address of the node. | [optional] 
**Message** | Pointer to **NullableString** | Specifies optional message related to node status. | [optional] 

## Methods

### NewClusterCreateNodeStatusResponse

`func NewClusterCreateNodeStatusResponse() *ClusterCreateNodeStatusResponse`

NewClusterCreateNodeStatusResponse instantiates a new ClusterCreateNodeStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateNodeStatusResponseWithDefaults

`func NewClusterCreateNodeStatusResponseWithDefaults() *ClusterCreateNodeStatusResponse`

NewClusterCreateNodeStatusResponseWithDefaults instantiates a new ClusterCreateNodeStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ClusterCreateNodeStatusResponse) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ClusterCreateNodeStatusResponse) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ClusterCreateNodeStatusResponse) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ClusterCreateNodeStatusResponse) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *ClusterCreateNodeStatusResponse) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *ClusterCreateNodeStatusResponse) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *ClusterCreateNodeStatusResponse) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ClusterCreateNodeStatusResponse) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ClusterCreateNodeStatusResponse) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *ClusterCreateNodeStatusResponse) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *ClusterCreateNodeStatusResponse) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ClusterCreateNodeStatusResponse) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetMessage

`func (o *ClusterCreateNodeStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterCreateNodeStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterCreateNodeStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterCreateNodeStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterCreateNodeStatusResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterCreateNodeStatusResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


