# StaticRouteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description of the Static Route. | [optional] 
**DestinationNetwork** | **NullableString** | Specifies the destination network of the Static Route. | 
**Id** | Pointer to **NullableString** | Specifies the unique identifier for the route. | [optional] [readonly] 
**Interface** | Pointer to **NullableString** | Specifies the network interface name to use for communicating with the destination network. | [optional] 
**InterfaceGroup** | **NullableString** | Specifies the network interfaces name to use for communicating with the destination network. | 
**Mtu** | Pointer to **NullableInt64** | Specifies MTU setting per route. | [optional] 
**NextHop** | **NullableString** | Specifies the next hop to the destination network. | 
**NodeGroupName** | Pointer to **NullableString** | Specifies the network node group to represent a group of nodes. | [optional] 

## Methods

### NewStaticRouteParams

`func NewStaticRouteParams(destinationNetwork NullableString, interfaceGroup NullableString, nextHop NullableString, ) *StaticRouteParams`

NewStaticRouteParams instantiates a new StaticRouteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRouteParamsWithDefaults

`func NewStaticRouteParamsWithDefaults() *StaticRouteParams`

NewStaticRouteParamsWithDefaults instantiates a new StaticRouteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StaticRouteParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticRouteParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticRouteParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticRouteParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StaticRouteParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StaticRouteParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationNetwork

`func (o *StaticRouteParams) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *StaticRouteParams) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *StaticRouteParams) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.


### SetDestinationNetworkNil

`func (o *StaticRouteParams) SetDestinationNetworkNil(b bool)`

 SetDestinationNetworkNil sets the value for DestinationNetwork to be an explicit nil

### UnsetDestinationNetwork
`func (o *StaticRouteParams) UnsetDestinationNetwork()`

UnsetDestinationNetwork ensures that no value is present for DestinationNetwork, not even an explicit nil
### GetId

`func (o *StaticRouteParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticRouteParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticRouteParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StaticRouteParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StaticRouteParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StaticRouteParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInterface

`func (o *StaticRouteParams) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *StaticRouteParams) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *StaticRouteParams) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *StaticRouteParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *StaticRouteParams) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *StaticRouteParams) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetInterfaceGroup

`func (o *StaticRouteParams) GetInterfaceGroup() string`

GetInterfaceGroup returns the InterfaceGroup field if non-nil, zero value otherwise.

### GetInterfaceGroupOk

`func (o *StaticRouteParams) GetInterfaceGroupOk() (*string, bool)`

GetInterfaceGroupOk returns a tuple with the InterfaceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroup

`func (o *StaticRouteParams) SetInterfaceGroup(v string)`

SetInterfaceGroup sets InterfaceGroup field to given value.


### SetInterfaceGroupNil

`func (o *StaticRouteParams) SetInterfaceGroupNil(b bool)`

 SetInterfaceGroupNil sets the value for InterfaceGroup to be an explicit nil

### UnsetInterfaceGroup
`func (o *StaticRouteParams) UnsetInterfaceGroup()`

UnsetInterfaceGroup ensures that no value is present for InterfaceGroup, not even an explicit nil
### GetMtu

`func (o *StaticRouteParams) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *StaticRouteParams) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *StaticRouteParams) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *StaticRouteParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *StaticRouteParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *StaticRouteParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetNextHop

`func (o *StaticRouteParams) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *StaticRouteParams) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *StaticRouteParams) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.


### SetNextHopNil

`func (o *StaticRouteParams) SetNextHopNil(b bool)`

 SetNextHopNil sets the value for NextHop to be an explicit nil

### UnsetNextHop
`func (o *StaticRouteParams) UnsetNextHop()`

UnsetNextHop ensures that no value is present for NextHop, not even an explicit nil
### GetNodeGroupName

`func (o *StaticRouteParams) GetNodeGroupName() string`

GetNodeGroupName returns the NodeGroupName field if non-nil, zero value otherwise.

### GetNodeGroupNameOk

`func (o *StaticRouteParams) GetNodeGroupNameOk() (*string, bool)`

GetNodeGroupNameOk returns a tuple with the NodeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupName

`func (o *StaticRouteParams) SetNodeGroupName(v string)`

SetNodeGroupName sets NodeGroupName field to given value.

### HasNodeGroupName

`func (o *StaticRouteParams) HasNodeGroupName() bool`

HasNodeGroupName returns a boolean if a field has been set.

### SetNodeGroupNameNil

`func (o *StaticRouteParams) SetNodeGroupNameNil(b bool)`

 SetNodeGroupNameNil sets the value for NodeGroupName to be an explicit nil

### UnsetNodeGroupName
`func (o *StaticRouteParams) UnsetNodeGroupName()`

UnsetNodeGroupName ensures that no value is present for NodeGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


