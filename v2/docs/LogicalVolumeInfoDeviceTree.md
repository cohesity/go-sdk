# LogicalVolumeInfoDeviceTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLeaf** | Pointer to **NullableBool** | Specifies if the node is a leaf node. | [optional] 
**LeafNodeParams** | Pointer to [**DeviceTreeNodeLeafNodeParams**](DeviceTreeNodeLeafNodeParams.md) |  | [optional] 
**NonLeafNodeParams** | Pointer to [**DeviceTreeNodeNonLeafNodeParams**](DeviceTreeNodeNonLeafNodeParams.md) |  | [optional] 

## Methods

### NewLogicalVolumeInfoDeviceTree

`func NewLogicalVolumeInfoDeviceTree() *LogicalVolumeInfoDeviceTree`

NewLogicalVolumeInfoDeviceTree instantiates a new LogicalVolumeInfoDeviceTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalVolumeInfoDeviceTreeWithDefaults

`func NewLogicalVolumeInfoDeviceTreeWithDefaults() *LogicalVolumeInfoDeviceTree`

NewLogicalVolumeInfoDeviceTreeWithDefaults instantiates a new LogicalVolumeInfoDeviceTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLeaf

`func (o *LogicalVolumeInfoDeviceTree) GetIsLeaf() bool`

GetIsLeaf returns the IsLeaf field if non-nil, zero value otherwise.

### GetIsLeafOk

`func (o *LogicalVolumeInfoDeviceTree) GetIsLeafOk() (*bool, bool)`

GetIsLeafOk returns a tuple with the IsLeaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeaf

`func (o *LogicalVolumeInfoDeviceTree) SetIsLeaf(v bool)`

SetIsLeaf sets IsLeaf field to given value.

### HasIsLeaf

`func (o *LogicalVolumeInfoDeviceTree) HasIsLeaf() bool`

HasIsLeaf returns a boolean if a field has been set.

### SetIsLeafNil

`func (o *LogicalVolumeInfoDeviceTree) SetIsLeafNil(b bool)`

 SetIsLeafNil sets the value for IsLeaf to be an explicit nil

### UnsetIsLeaf
`func (o *LogicalVolumeInfoDeviceTree) UnsetIsLeaf()`

UnsetIsLeaf ensures that no value is present for IsLeaf, not even an explicit nil
### GetLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) GetLeafNodeParams() DeviceTreeNodeLeafNodeParams`

GetLeafNodeParams returns the LeafNodeParams field if non-nil, zero value otherwise.

### GetLeafNodeParamsOk

`func (o *LogicalVolumeInfoDeviceTree) GetLeafNodeParamsOk() (*DeviceTreeNodeLeafNodeParams, bool)`

GetLeafNodeParamsOk returns a tuple with the LeafNodeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) SetLeafNodeParams(v DeviceTreeNodeLeafNodeParams)`

SetLeafNodeParams sets LeafNodeParams field to given value.

### HasLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) HasLeafNodeParams() bool`

HasLeafNodeParams returns a boolean if a field has been set.

### GetNonLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) GetNonLeafNodeParams() DeviceTreeNodeNonLeafNodeParams`

GetNonLeafNodeParams returns the NonLeafNodeParams field if non-nil, zero value otherwise.

### GetNonLeafNodeParamsOk

`func (o *LogicalVolumeInfoDeviceTree) GetNonLeafNodeParamsOk() (*DeviceTreeNodeNonLeafNodeParams, bool)`

GetNonLeafNodeParamsOk returns a tuple with the NonLeafNodeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) SetNonLeafNodeParams(v DeviceTreeNodeNonLeafNodeParams)`

SetNonLeafNodeParams sets NonLeafNodeParams field to given value.

### HasNonLeafNodeParams

`func (o *LogicalVolumeInfoDeviceTree) HasNonLeafNodeParams() bool`

HasNonLeafNodeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


