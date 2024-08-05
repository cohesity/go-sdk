# InterfaceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the interface group. | 
**NetworkParams** | Pointer to [**InterfaceGroupNetworkParams**](InterfaceGroupNetworkParams.md) |  | [optional] 
**NodeInterfaceParams** | [**[]NodeInterfaceParams**](NodeInterfaceParams.md) | Node and interface parameters. | 
**Type** | **string** | Type of the interface group. | 
**Id** | Pointer to **int32** | Id of the interface group. | [optional] 

## Methods

### NewInterfaceGroup

`func NewInterfaceGroup(name string, nodeInterfaceParams []NodeInterfaceParams, type_ string, ) *InterfaceGroup`

NewInterfaceGroup instantiates a new InterfaceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceGroupWithDefaults

`func NewInterfaceGroupWithDefaults() *InterfaceGroup`

NewInterfaceGroupWithDefaults instantiates a new InterfaceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkParams

`func (o *InterfaceGroup) GetNetworkParams() InterfaceGroupNetworkParams`

GetNetworkParams returns the NetworkParams field if non-nil, zero value otherwise.

### GetNetworkParamsOk

`func (o *InterfaceGroup) GetNetworkParamsOk() (*InterfaceGroupNetworkParams, bool)`

GetNetworkParamsOk returns a tuple with the NetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParams

`func (o *InterfaceGroup) SetNetworkParams(v InterfaceGroupNetworkParams)`

SetNetworkParams sets NetworkParams field to given value.

### HasNetworkParams

`func (o *InterfaceGroup) HasNetworkParams() bool`

HasNetworkParams returns a boolean if a field has been set.

### GetNodeInterfaceParams

`func (o *InterfaceGroup) GetNodeInterfaceParams() []NodeInterfaceParams`

GetNodeInterfaceParams returns the NodeInterfaceParams field if non-nil, zero value otherwise.

### GetNodeInterfaceParamsOk

`func (o *InterfaceGroup) GetNodeInterfaceParamsOk() (*[]NodeInterfaceParams, bool)`

GetNodeInterfaceParamsOk returns a tuple with the NodeInterfaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInterfaceParams

`func (o *InterfaceGroup) SetNodeInterfaceParams(v []NodeInterfaceParams)`

SetNodeInterfaceParams sets NodeInterfaceParams field to given value.


### GetType

`func (o *InterfaceGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceGroup) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InterfaceGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InterfaceGroup) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


