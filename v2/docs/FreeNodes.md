# FreeNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]FreeNodeInformation**](FreeNodeInformation.md) | Specifies the list of free nodes. | [optional] 

## Methods

### NewFreeNodes

`func NewFreeNodes() *FreeNodes`

NewFreeNodes instantiates a new FreeNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeNodesWithDefaults

`func NewFreeNodesWithDefaults() *FreeNodes`

NewFreeNodesWithDefaults instantiates a new FreeNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *FreeNodes) GetNodes() []FreeNodeInformation`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FreeNodes) GetNodesOk() (*[]FreeNodeInformation, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FreeNodes) SetNodes(v []FreeNodeInformation)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *FreeNodes) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *FreeNodes) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *FreeNodes) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


