# DeviceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntermediateNode** | Pointer to [**DeviceTreeDetails**](DeviceTreeDetails.md) |  | [optional] 
**LeafNode** | Pointer to [**FilePartitionBlock**](FilePartitionBlock.md) |  | [optional] 

## Methods

### NewDeviceNode

`func NewDeviceNode() *DeviceNode`

NewDeviceNode instantiates a new DeviceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceNodeWithDefaults

`func NewDeviceNodeWithDefaults() *DeviceNode`

NewDeviceNodeWithDefaults instantiates a new DeviceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntermediateNode

`func (o *DeviceNode) GetIntermediateNode() DeviceTreeDetails`

GetIntermediateNode returns the IntermediateNode field if non-nil, zero value otherwise.

### GetIntermediateNodeOk

`func (o *DeviceNode) GetIntermediateNodeOk() (*DeviceTreeDetails, bool)`

GetIntermediateNodeOk returns a tuple with the IntermediateNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateNode

`func (o *DeviceNode) SetIntermediateNode(v DeviceTreeDetails)`

SetIntermediateNode sets IntermediateNode field to given value.

### HasIntermediateNode

`func (o *DeviceNode) HasIntermediateNode() bool`

HasIntermediateNode returns a boolean if a field has been set.

### GetLeafNode

`func (o *DeviceNode) GetLeafNode() FilePartitionBlock`

GetLeafNode returns the LeafNode field if non-nil, zero value otherwise.

### GetLeafNodeOk

`func (o *DeviceNode) GetLeafNodeOk() (*FilePartitionBlock, bool)`

GetLeafNodeOk returns a tuple with the LeafNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNode

`func (o *DeviceNode) SetLeafNode(v FilePartitionBlock)`

SetLeafNode sets LeafNode field to given value.

### HasLeafNode

`func (o *DeviceNode) HasLeafNode() bool`

HasLeafNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


