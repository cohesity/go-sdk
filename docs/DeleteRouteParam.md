# DeleteRouteParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestNetwork** | Pointer to **NullableString** | Destination network.  Specifies the destination network of the Static Route. overrideDescription: true | [optional] 
**IfName** | Pointer to **NullableString** | Specifies the network interfaces name to use for communicating with the destination network. | [optional] 
**IfaceGroupName** | Pointer to **NullableString** | Specifies the network interfaces group or vlan interface group to use for communicating with the destination network. | [optional] 

## Methods

### NewDeleteRouteParam

`func NewDeleteRouteParam() *DeleteRouteParam`

NewDeleteRouteParam instantiates a new DeleteRouteParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRouteParamWithDefaults

`func NewDeleteRouteParamWithDefaults() *DeleteRouteParam`

NewDeleteRouteParamWithDefaults instantiates a new DeleteRouteParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestNetwork

`func (o *DeleteRouteParam) GetDestNetwork() string`

GetDestNetwork returns the DestNetwork field if non-nil, zero value otherwise.

### GetDestNetworkOk

`func (o *DeleteRouteParam) GetDestNetworkOk() (*string, bool)`

GetDestNetworkOk returns a tuple with the DestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestNetwork

`func (o *DeleteRouteParam) SetDestNetwork(v string)`

SetDestNetwork sets DestNetwork field to given value.

### HasDestNetwork

`func (o *DeleteRouteParam) HasDestNetwork() bool`

HasDestNetwork returns a boolean if a field has been set.

### SetDestNetworkNil

`func (o *DeleteRouteParam) SetDestNetworkNil(b bool)`

 SetDestNetworkNil sets the value for DestNetwork to be an explicit nil

### UnsetDestNetwork
`func (o *DeleteRouteParam) UnsetDestNetwork()`

UnsetDestNetwork ensures that no value is present for DestNetwork, not even an explicit nil
### GetIfName

`func (o *DeleteRouteParam) GetIfName() string`

GetIfName returns the IfName field if non-nil, zero value otherwise.

### GetIfNameOk

`func (o *DeleteRouteParam) GetIfNameOk() (*string, bool)`

GetIfNameOk returns a tuple with the IfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfName

`func (o *DeleteRouteParam) SetIfName(v string)`

SetIfName sets IfName field to given value.

### HasIfName

`func (o *DeleteRouteParam) HasIfName() bool`

HasIfName returns a boolean if a field has been set.

### SetIfNameNil

`func (o *DeleteRouteParam) SetIfNameNil(b bool)`

 SetIfNameNil sets the value for IfName to be an explicit nil

### UnsetIfName
`func (o *DeleteRouteParam) UnsetIfName()`

UnsetIfName ensures that no value is present for IfName, not even an explicit nil
### GetIfaceGroupName

`func (o *DeleteRouteParam) GetIfaceGroupName() string`

GetIfaceGroupName returns the IfaceGroupName field if non-nil, zero value otherwise.

### GetIfaceGroupNameOk

`func (o *DeleteRouteParam) GetIfaceGroupNameOk() (*string, bool)`

GetIfaceGroupNameOk returns a tuple with the IfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfaceGroupName

`func (o *DeleteRouteParam) SetIfaceGroupName(v string)`

SetIfaceGroupName sets IfaceGroupName field to given value.

### HasIfaceGroupName

`func (o *DeleteRouteParam) HasIfaceGroupName() bool`

HasIfaceGroupName returns a boolean if a field has been set.

### SetIfaceGroupNameNil

`func (o *DeleteRouteParam) SetIfaceGroupNameNil(b bool)`

 SetIfaceGroupNameNil sets the value for IfaceGroupName to be an explicit nil

### UnsetIfaceGroupName
`func (o *DeleteRouteParam) UnsetIfaceGroupName()`

UnsetIfaceGroupName ensures that no value is present for IfaceGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


