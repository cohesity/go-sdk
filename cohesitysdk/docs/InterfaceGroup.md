# InterfaceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Interface group Id.  Specifies the id of the interface group. | [optional] 
**ModelInterfaceLists** | Pointer to [**[]ProductModelInterfaceTuple**](ProductModelInterfaceTuple.md) | Specifies the product model and interface lists. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the interface group. | [optional] 
**NetworkParams** | Pointer to [**NetworkParams**](NetworkParams.md) |  | [optional] 

## Methods

### NewInterfaceGroup

`func NewInterfaceGroup() *InterfaceGroup`

NewInterfaceGroup instantiates a new InterfaceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceGroupWithDefaults

`func NewInterfaceGroupWithDefaults() *InterfaceGroup`

NewInterfaceGroupWithDefaults instantiates a new InterfaceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### SetIdNil

`func (o *InterfaceGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InterfaceGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModelInterfaceLists

`func (o *InterfaceGroup) GetModelInterfaceLists() []ProductModelInterfaceTuple`

GetModelInterfaceLists returns the ModelInterfaceLists field if non-nil, zero value otherwise.

### GetModelInterfaceListsOk

`func (o *InterfaceGroup) GetModelInterfaceListsOk() (*[]ProductModelInterfaceTuple, bool)`

GetModelInterfaceListsOk returns a tuple with the ModelInterfaceLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInterfaceLists

`func (o *InterfaceGroup) SetModelInterfaceLists(v []ProductModelInterfaceTuple)`

SetModelInterfaceLists sets ModelInterfaceLists field to given value.

### HasModelInterfaceLists

`func (o *InterfaceGroup) HasModelInterfaceLists() bool`

HasModelInterfaceLists returns a boolean if a field has been set.

### SetModelInterfaceListsNil

`func (o *InterfaceGroup) SetModelInterfaceListsNil(b bool)`

 SetModelInterfaceListsNil sets the value for ModelInterfaceLists to be an explicit nil

### UnsetModelInterfaceLists
`func (o *InterfaceGroup) UnsetModelInterfaceLists()`

UnsetModelInterfaceLists ensures that no value is present for ModelInterfaceLists, not even an explicit nil
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

### HasName

`func (o *InterfaceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InterfaceGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InterfaceGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkParams

`func (o *InterfaceGroup) GetNetworkParams() NetworkParams`

GetNetworkParams returns the NetworkParams field if non-nil, zero value otherwise.

### GetNetworkParamsOk

`func (o *InterfaceGroup) GetNetworkParamsOk() (*NetworkParams, bool)`

GetNetworkParamsOk returns a tuple with the NetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParams

`func (o *InterfaceGroup) SetNetworkParams(v NetworkParams)`

SetNetworkParams sets NetworkParams field to given value.

### HasNetworkParams

`func (o *InterfaceGroup) HasNetworkParams() bool`

HasNetworkParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


