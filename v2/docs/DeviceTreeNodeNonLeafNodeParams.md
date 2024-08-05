# DeviceTreeNodeNonLeafNodeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildrenNodes** | Pointer to [**[]DeviceTreeNode**](DeviceTreeNode.md) | Specifies a list of children nodes. | [optional] 
**DeviceId** | Pointer to **NullableInt64** | Specifies the id of device. | [optional] 
**DeviceLength** | Pointer to **NullableInt64** | Specifies the length of device. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the children nodes combine type. | [optional] 

## Methods

### NewDeviceTreeNodeNonLeafNodeParams

`func NewDeviceTreeNodeNonLeafNodeParams() *DeviceTreeNodeNonLeafNodeParams`

NewDeviceTreeNodeNonLeafNodeParams instantiates a new DeviceTreeNodeNonLeafNodeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeNodeNonLeafNodeParamsWithDefaults

`func NewDeviceTreeNodeNonLeafNodeParamsWithDefaults() *DeviceTreeNodeNonLeafNodeParams`

NewDeviceTreeNodeNonLeafNodeParamsWithDefaults instantiates a new DeviceTreeNodeNonLeafNodeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildrenNodes

`func (o *DeviceTreeNodeNonLeafNodeParams) GetChildrenNodes() []DeviceTreeNode`

GetChildrenNodes returns the ChildrenNodes field if non-nil, zero value otherwise.

### GetChildrenNodesOk

`func (o *DeviceTreeNodeNonLeafNodeParams) GetChildrenNodesOk() (*[]DeviceTreeNode, bool)`

GetChildrenNodesOk returns a tuple with the ChildrenNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenNodes

`func (o *DeviceTreeNodeNonLeafNodeParams) SetChildrenNodes(v []DeviceTreeNode)`

SetChildrenNodes sets ChildrenNodes field to given value.

### HasChildrenNodes

`func (o *DeviceTreeNodeNonLeafNodeParams) HasChildrenNodes() bool`

HasChildrenNodes returns a boolean if a field has been set.

### SetChildrenNodesNil

`func (o *DeviceTreeNodeNonLeafNodeParams) SetChildrenNodesNil(b bool)`

 SetChildrenNodesNil sets the value for ChildrenNodes to be an explicit nil

### UnsetChildrenNodes
`func (o *DeviceTreeNodeNonLeafNodeParams) UnsetChildrenNodes()`

UnsetChildrenNodes ensures that no value is present for ChildrenNodes, not even an explicit nil
### GetDeviceId

`func (o *DeviceTreeNodeNonLeafNodeParams) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceTreeNodeNonLeafNodeParams) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceTreeNodeNonLeafNodeParams) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceTreeNodeNonLeafNodeParams) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceTreeNodeNonLeafNodeParams) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceTreeNodeNonLeafNodeParams) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceLength

`func (o *DeviceTreeNodeNonLeafNodeParams) GetDeviceLength() int64`

GetDeviceLength returns the DeviceLength field if non-nil, zero value otherwise.

### GetDeviceLengthOk

`func (o *DeviceTreeNodeNonLeafNodeParams) GetDeviceLengthOk() (*int64, bool)`

GetDeviceLengthOk returns a tuple with the DeviceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLength

`func (o *DeviceTreeNodeNonLeafNodeParams) SetDeviceLength(v int64)`

SetDeviceLength sets DeviceLength field to given value.

### HasDeviceLength

`func (o *DeviceTreeNodeNonLeafNodeParams) HasDeviceLength() bool`

HasDeviceLength returns a boolean if a field has been set.

### SetDeviceLengthNil

`func (o *DeviceTreeNodeNonLeafNodeParams) SetDeviceLengthNil(b bool)`

 SetDeviceLengthNil sets the value for DeviceLength to be an explicit nil

### UnsetDeviceLength
`func (o *DeviceTreeNodeNonLeafNodeParams) UnsetDeviceLength()`

UnsetDeviceLength ensures that no value is present for DeviceLength, not even an explicit nil
### GetType

`func (o *DeviceTreeNodeNonLeafNodeParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceTreeNodeNonLeafNodeParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceTreeNodeNonLeafNodeParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceTreeNodeNonLeafNodeParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeviceTreeNodeNonLeafNodeParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeviceTreeNodeNonLeafNodeParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


