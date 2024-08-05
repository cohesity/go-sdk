# NodeGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeGroups** | Pointer to [**[]NodeGroup**](NodeGroup.md) | Specifies the details of a Node Group. | [optional] 

## Methods

### NewNodeGroupResponse

`func NewNodeGroupResponse() *NodeGroupResponse`

NewNodeGroupResponse instantiates a new NodeGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupResponseWithDefaults

`func NewNodeGroupResponseWithDefaults() *NodeGroupResponse`

NewNodeGroupResponseWithDefaults instantiates a new NodeGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeGroups

`func (o *NodeGroupResponse) GetNodeGroups() []NodeGroup`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *NodeGroupResponse) GetNodeGroupsOk() (*[]NodeGroup, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *NodeGroupResponse) SetNodeGroups(v []NodeGroup)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *NodeGroupResponse) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### SetNodeGroupsNil

`func (o *NodeGroupResponse) SetNodeGroupsNil(b bool)`

 SetNodeGroupsNil sets the value for NodeGroups to be an explicit nil

### UnsetNodeGroups
`func (o *NodeGroupResponse) UnsetNodeGroups()`

UnsetNodeGroups ensures that no value is present for NodeGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


