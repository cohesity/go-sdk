# NodeInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the node. | [optional] 
**Interfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) | Specifies the list of network interfaces present on this Node. | [optional] 

## Methods

### NewNodeInterfaces

`func NewNodeInterfaces() *NodeInterfaces`

NewNodeInterfaces instantiates a new NodeInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInterfacesWithDefaults

`func NewNodeInterfacesWithDefaults() *NodeInterfaces`

NewNodeInterfacesWithDefaults instantiates a new NodeInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeInterfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeInterfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeInterfaces) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeInterfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeInterfaces) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeInterfaces) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInterfaces

`func (o *NodeInterfaces) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NodeInterfaces) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NodeInterfaces) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NodeInterfaces) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


