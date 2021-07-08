# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description of the Static Route. | [optional] 
**DestNetwork** | Pointer to **NullableString** | Destination network.  Specifies the destination network of the Static Route. overrideDescription: true | [optional] 
**IfName** | Pointer to **NullableString** | Specifies the network interfaces name to use for communicating with the destination network. | [optional] 
**IfaceGroupName** | Pointer to **NullableString** | Specifies the network interfaces group or interface group with vlan id to use for communicating with the destination network. | [optional] 
**NextHop** | Pointer to **NullableString** | Specifies the next hop to the destination network. | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Route) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Route) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Route) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestNetwork

`func (o *Route) GetDestNetwork() string`

GetDestNetwork returns the DestNetwork field if non-nil, zero value otherwise.

### GetDestNetworkOk

`func (o *Route) GetDestNetworkOk() (*string, bool)`

GetDestNetworkOk returns a tuple with the DestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestNetwork

`func (o *Route) SetDestNetwork(v string)`

SetDestNetwork sets DestNetwork field to given value.

### HasDestNetwork

`func (o *Route) HasDestNetwork() bool`

HasDestNetwork returns a boolean if a field has been set.

### SetDestNetworkNil

`func (o *Route) SetDestNetworkNil(b bool)`

 SetDestNetworkNil sets the value for DestNetwork to be an explicit nil

### UnsetDestNetwork
`func (o *Route) UnsetDestNetwork()`

UnsetDestNetwork ensures that no value is present for DestNetwork, not even an explicit nil
### GetIfName

`func (o *Route) GetIfName() string`

GetIfName returns the IfName field if non-nil, zero value otherwise.

### GetIfNameOk

`func (o *Route) GetIfNameOk() (*string, bool)`

GetIfNameOk returns a tuple with the IfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfName

`func (o *Route) SetIfName(v string)`

SetIfName sets IfName field to given value.

### HasIfName

`func (o *Route) HasIfName() bool`

HasIfName returns a boolean if a field has been set.

### SetIfNameNil

`func (o *Route) SetIfNameNil(b bool)`

 SetIfNameNil sets the value for IfName to be an explicit nil

### UnsetIfName
`func (o *Route) UnsetIfName()`

UnsetIfName ensures that no value is present for IfName, not even an explicit nil
### GetIfaceGroupName

`func (o *Route) GetIfaceGroupName() string`

GetIfaceGroupName returns the IfaceGroupName field if non-nil, zero value otherwise.

### GetIfaceGroupNameOk

`func (o *Route) GetIfaceGroupNameOk() (*string, bool)`

GetIfaceGroupNameOk returns a tuple with the IfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfaceGroupName

`func (o *Route) SetIfaceGroupName(v string)`

SetIfaceGroupName sets IfaceGroupName field to given value.

### HasIfaceGroupName

`func (o *Route) HasIfaceGroupName() bool`

HasIfaceGroupName returns a boolean if a field has been set.

### SetIfaceGroupNameNil

`func (o *Route) SetIfaceGroupNameNil(b bool)`

 SetIfaceGroupNameNil sets the value for IfaceGroupName to be an explicit nil

### UnsetIfaceGroupName
`func (o *Route) UnsetIfaceGroupName()`

UnsetIfaceGroupName ensures that no value is present for IfaceGroupName, not even an explicit nil
### GetNextHop

`func (o *Route) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *Route) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *Route) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *Route) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### SetNextHopNil

`func (o *Route) SetNextHopNil(b bool)`

 SetNextHopNil sets the value for NextHop to be an explicit nil

### UnsetNextHop
`func (o *Route) UnsetNextHop()`

UnsetNextHop ensures that no value is present for NextHop, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


