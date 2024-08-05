# NodeInterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceName** | Pointer to **NullableString** | Name of the interface. | [optional] 
**NodeId** | **int64** | Node id. | 

## Methods

### NewNodeInterfaceParams

`func NewNodeInterfaceParams(nodeId int64, ) *NodeInterfaceParams`

NewNodeInterfaceParams instantiates a new NodeInterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInterfaceParamsWithDefaults

`func NewNodeInterfaceParamsWithDefaults() *NodeInterfaceParams`

NewNodeInterfaceParamsWithDefaults instantiates a new NodeInterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceName

`func (o *NodeInterfaceParams) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NodeInterfaceParams) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NodeInterfaceParams) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *NodeInterfaceParams) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *NodeInterfaceParams) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *NodeInterfaceParams) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetNodeId

`func (o *NodeInterfaceParams) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeInterfaceParams) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeInterfaceParams) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


