# NodeGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpInstance** | Pointer to [**BgpInstance**](BgpInstance.md) |  | [optional] 
**DnsServersInfo** | Pointer to [**DnsServersInfo**](DnsServersInfo.md) |  | [optional] 
**Id** | Pointer to **NullableInt32** | Id of the node group. | [optional] 
**Name** | **string** | Specifies the name of the Node Group. | 
**NodeIds** | Pointer to **[]int64** | List of Node Ids that are part of this node group. | [optional] 
**SubnetInfo** | Pointer to [**SubnetInfo**](SubnetInfo.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | Type of the node group. | [optional] 

## Methods

### NewNodeGroupRequest

`func NewNodeGroupRequest(name string, ) *NodeGroupRequest`

NewNodeGroupRequest instantiates a new NodeGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupRequestWithDefaults

`func NewNodeGroupRequestWithDefaults() *NodeGroupRequest`

NewNodeGroupRequestWithDefaults instantiates a new NodeGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpInstance

`func (o *NodeGroupRequest) GetBgpInstance() BgpInstance`

GetBgpInstance returns the BgpInstance field if non-nil, zero value otherwise.

### GetBgpInstanceOk

`func (o *NodeGroupRequest) GetBgpInstanceOk() (*BgpInstance, bool)`

GetBgpInstanceOk returns a tuple with the BgpInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpInstance

`func (o *NodeGroupRequest) SetBgpInstance(v BgpInstance)`

SetBgpInstance sets BgpInstance field to given value.

### HasBgpInstance

`func (o *NodeGroupRequest) HasBgpInstance() bool`

HasBgpInstance returns a boolean if a field has been set.

### GetDnsServersInfo

`func (o *NodeGroupRequest) GetDnsServersInfo() DnsServersInfo`

GetDnsServersInfo returns the DnsServersInfo field if non-nil, zero value otherwise.

### GetDnsServersInfoOk

`func (o *NodeGroupRequest) GetDnsServersInfoOk() (*DnsServersInfo, bool)`

GetDnsServersInfoOk returns a tuple with the DnsServersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServersInfo

`func (o *NodeGroupRequest) SetDnsServersInfo(v DnsServersInfo)`

SetDnsServersInfo sets DnsServersInfo field to given value.

### HasDnsServersInfo

`func (o *NodeGroupRequest) HasDnsServersInfo() bool`

HasDnsServersInfo returns a boolean if a field has been set.

### GetId

`func (o *NodeGroupRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeGroupRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeGroupRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeGroupRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeGroupRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *NodeGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNodeIds

`func (o *NodeGroupRequest) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *NodeGroupRequest) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *NodeGroupRequest) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *NodeGroupRequest) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *NodeGroupRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *NodeGroupRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetSubnetInfo

`func (o *NodeGroupRequest) GetSubnetInfo() SubnetInfo`

GetSubnetInfo returns the SubnetInfo field if non-nil, zero value otherwise.

### GetSubnetInfoOk

`func (o *NodeGroupRequest) GetSubnetInfoOk() (*SubnetInfo, bool)`

GetSubnetInfoOk returns a tuple with the SubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetInfo

`func (o *NodeGroupRequest) SetSubnetInfo(v SubnetInfo)`

SetSubnetInfo sets SubnetInfo field to given value.

### HasSubnetInfo

`func (o *NodeGroupRequest) HasSubnetInfo() bool`

HasSubnetInfo returns a boolean if a field has been set.

### GetType

`func (o *NodeGroupRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeGroupRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeGroupRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *NodeGroupRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NodeGroupRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NodeGroupRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


