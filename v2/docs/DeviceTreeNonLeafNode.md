# DeviceTreeNonLeafNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildrenNodes** | Pointer to [**[]DeviceTreeNode**](DeviceTreeNode.md) | Specifies a list of children nodes. | [optional] 
**DeviceId** | Pointer to **NullableInt64** | Specifies the id of device. | [optional] 
**DeviceLength** | Pointer to **NullableInt64** | Specifies the length of device. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the children nodes combine type. | [optional] 

## Methods

### NewDeviceTreeNonLeafNode

`func NewDeviceTreeNonLeafNode() *DeviceTreeNonLeafNode`

NewDeviceTreeNonLeafNode instantiates a new DeviceTreeNonLeafNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeNonLeafNodeWithDefaults

`func NewDeviceTreeNonLeafNodeWithDefaults() *DeviceTreeNonLeafNode`

NewDeviceTreeNonLeafNodeWithDefaults instantiates a new DeviceTreeNonLeafNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildrenNodes

`func (o *DeviceTreeNonLeafNode) GetChildrenNodes() []DeviceTreeNode`

GetChildrenNodes returns the ChildrenNodes field if non-nil, zero value otherwise.

### GetChildrenNodesOk

`func (o *DeviceTreeNonLeafNode) GetChildrenNodesOk() (*[]DeviceTreeNode, bool)`

GetChildrenNodesOk returns a tuple with the ChildrenNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenNodes

`func (o *DeviceTreeNonLeafNode) SetChildrenNodes(v []DeviceTreeNode)`

SetChildrenNodes sets ChildrenNodes field to given value.

### HasChildrenNodes

`func (o *DeviceTreeNonLeafNode) HasChildrenNodes() bool`

HasChildrenNodes returns a boolean if a field has been set.

### SetChildrenNodesNil

`func (o *DeviceTreeNonLeafNode) SetChildrenNodesNil(b bool)`

 SetChildrenNodesNil sets the value for ChildrenNodes to be an explicit nil

### UnsetChildrenNodes
`func (o *DeviceTreeNonLeafNode) UnsetChildrenNodes()`

UnsetChildrenNodes ensures that no value is present for ChildrenNodes, not even an explicit nil
### GetDeviceId

`func (o *DeviceTreeNonLeafNode) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceTreeNonLeafNode) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceTreeNonLeafNode) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceTreeNonLeafNode) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceTreeNonLeafNode) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceTreeNonLeafNode) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceLength

`func (o *DeviceTreeNonLeafNode) GetDeviceLength() int64`

GetDeviceLength returns the DeviceLength field if non-nil, zero value otherwise.

### GetDeviceLengthOk

`func (o *DeviceTreeNonLeafNode) GetDeviceLengthOk() (*int64, bool)`

GetDeviceLengthOk returns a tuple with the DeviceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLength

`func (o *DeviceTreeNonLeafNode) SetDeviceLength(v int64)`

SetDeviceLength sets DeviceLength field to given value.

### HasDeviceLength

`func (o *DeviceTreeNonLeafNode) HasDeviceLength() bool`

HasDeviceLength returns a boolean if a field has been set.

### SetDeviceLengthNil

`func (o *DeviceTreeNonLeafNode) SetDeviceLengthNil(b bool)`

 SetDeviceLengthNil sets the value for DeviceLength to be an explicit nil

### UnsetDeviceLength
`func (o *DeviceTreeNonLeafNode) UnsetDeviceLength()`

UnsetDeviceLength ensures that no value is present for DeviceLength, not even an explicit nil
### GetType

`func (o *DeviceTreeNonLeafNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceTreeNonLeafNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceTreeNonLeafNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceTreeNonLeafNode) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeviceTreeNonLeafNode) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeviceTreeNonLeafNode) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


