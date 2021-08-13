# RigelClusterConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]RigelClusterNode**](RigelClusterNode.md) | Specifies the Nodes present in this Cluster. | [optional] 
**DataplaneEndpoint** | Pointer to **NullableString** | Specifies the endpoint of the dataplane cluster which is associated with this rigel. | [optional] 
**ClaimToken** | Pointer to **NullableString** | Specifies the token which is used to claim this Cluster to Helios. | [optional] [readonly] 

## Methods

### NewRigelClusterConfigParams

`func NewRigelClusterConfigParams() *RigelClusterConfigParams`

NewRigelClusterConfigParams instantiates a new RigelClusterConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelClusterConfigParamsWithDefaults

`func NewRigelClusterConfigParamsWithDefaults() *RigelClusterConfigParams`

NewRigelClusterConfigParamsWithDefaults instantiates a new RigelClusterConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *RigelClusterConfigParams) GetNodes() []RigelClusterNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *RigelClusterConfigParams) GetNodesOk() (*[]RigelClusterNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *RigelClusterConfigParams) SetNodes(v []RigelClusterNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *RigelClusterConfigParams) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *RigelClusterConfigParams) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *RigelClusterConfigParams) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetDataplaneEndpoint

`func (o *RigelClusterConfigParams) GetDataplaneEndpoint() string`

GetDataplaneEndpoint returns the DataplaneEndpoint field if non-nil, zero value otherwise.

### GetDataplaneEndpointOk

`func (o *RigelClusterConfigParams) GetDataplaneEndpointOk() (*string, bool)`

GetDataplaneEndpointOk returns a tuple with the DataplaneEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplaneEndpoint

`func (o *RigelClusterConfigParams) SetDataplaneEndpoint(v string)`

SetDataplaneEndpoint sets DataplaneEndpoint field to given value.

### HasDataplaneEndpoint

`func (o *RigelClusterConfigParams) HasDataplaneEndpoint() bool`

HasDataplaneEndpoint returns a boolean if a field has been set.

### SetDataplaneEndpointNil

`func (o *RigelClusterConfigParams) SetDataplaneEndpointNil(b bool)`

 SetDataplaneEndpointNil sets the value for DataplaneEndpoint to be an explicit nil

### UnsetDataplaneEndpoint
`func (o *RigelClusterConfigParams) UnsetDataplaneEndpoint()`

UnsetDataplaneEndpoint ensures that no value is present for DataplaneEndpoint, not even an explicit nil
### GetClaimToken

`func (o *RigelClusterConfigParams) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *RigelClusterConfigParams) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *RigelClusterConfigParams) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.

### HasClaimToken

`func (o *RigelClusterConfigParams) HasClaimToken() bool`

HasClaimToken returns a boolean if a field has been set.

### SetClaimTokenNil

`func (o *RigelClusterConfigParams) SetClaimTokenNil(b bool)`

 SetClaimTokenNil sets the value for ClaimToken to be an explicit nil

### UnsetClaimToken
`func (o *RigelClusterConfigParams) UnsetClaimToken()`

UnsetClaimToken ensures that no value is present for ClaimToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


