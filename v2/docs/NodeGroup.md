# NodeGroup

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

### NewNodeGroup

`func NewNodeGroup(name string, ) *NodeGroup`

NewNodeGroup instantiates a new NodeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupWithDefaults

`func NewNodeGroupWithDefaults() *NodeGroup`

NewNodeGroupWithDefaults instantiates a new NodeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpInstance

`func (o *NodeGroup) GetBgpInstance() BgpInstance`

GetBgpInstance returns the BgpInstance field if non-nil, zero value otherwise.

### GetBgpInstanceOk

`func (o *NodeGroup) GetBgpInstanceOk() (*BgpInstance, bool)`

GetBgpInstanceOk returns a tuple with the BgpInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpInstance

`func (o *NodeGroup) SetBgpInstance(v BgpInstance)`

SetBgpInstance sets BgpInstance field to given value.

### HasBgpInstance

`func (o *NodeGroup) HasBgpInstance() bool`

HasBgpInstance returns a boolean if a field has been set.

### GetDnsServersInfo

`func (o *NodeGroup) GetDnsServersInfo() DnsServersInfo`

GetDnsServersInfo returns the DnsServersInfo field if non-nil, zero value otherwise.

### GetDnsServersInfoOk

`func (o *NodeGroup) GetDnsServersInfoOk() (*DnsServersInfo, bool)`

GetDnsServersInfoOk returns a tuple with the DnsServersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServersInfo

`func (o *NodeGroup) SetDnsServersInfo(v DnsServersInfo)`

SetDnsServersInfo sets DnsServersInfo field to given value.

### HasDnsServersInfo

`func (o *NodeGroup) HasDnsServersInfo() bool`

HasDnsServersInfo returns a boolean if a field has been set.

### GetId

`func (o *NodeGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NodeGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *NodeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNodeIds

`func (o *NodeGroup) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *NodeGroup) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *NodeGroup) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *NodeGroup) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *NodeGroup) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *NodeGroup) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetSubnetInfo

`func (o *NodeGroup) GetSubnetInfo() SubnetInfo`

GetSubnetInfo returns the SubnetInfo field if non-nil, zero value otherwise.

### GetSubnetInfoOk

`func (o *NodeGroup) GetSubnetInfoOk() (*SubnetInfo, bool)`

GetSubnetInfoOk returns a tuple with the SubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetInfo

`func (o *NodeGroup) SetSubnetInfo(v SubnetInfo)`

SetSubnetInfo sets SubnetInfo field to given value.

### HasSubnetInfo

`func (o *NodeGroup) HasSubnetInfo() bool`

HasSubnetInfo returns a boolean if a field has been set.

### GetType

`func (o *NodeGroup) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeGroup) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeGroup) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *NodeGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NodeGroup) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NodeGroup) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


