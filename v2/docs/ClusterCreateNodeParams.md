# ClusterCreateNodeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **NullableInt64** | Specifies the node id of the node. | 
**NodeIp** | **NullableString** | Specifies the node ip address which is either in ipv4/ipv6 format. | 

## Methods

### NewClusterCreateNodeParams

`func NewClusterCreateNodeParams(nodeId NullableInt64, nodeIp NullableString, ) *ClusterCreateNodeParams`

NewClusterCreateNodeParams instantiates a new ClusterCreateNodeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateNodeParamsWithDefaults

`func NewClusterCreateNodeParamsWithDefaults() *ClusterCreateNodeParams`

NewClusterCreateNodeParamsWithDefaults instantiates a new ClusterCreateNodeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ClusterCreateNodeParams) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ClusterCreateNodeParams) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ClusterCreateNodeParams) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.


### SetNodeIdNil

`func (o *ClusterCreateNodeParams) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *ClusterCreateNodeParams) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *ClusterCreateNodeParams) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ClusterCreateNodeParams) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ClusterCreateNodeParams) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.


### SetNodeIpNil

`func (o *ClusterCreateNodeParams) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ClusterCreateNodeParams) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


