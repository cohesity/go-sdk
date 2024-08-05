# ConnectivityCheckResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]NodeEndpointState**](NodeEndpointState.md) | Specifies the results of the nodes in the cluster. | [optional] 

## Methods

### NewConnectivityCheckResponseBody

`func NewConnectivityCheckResponseBody() *ConnectivityCheckResponseBody`

NewConnectivityCheckResponseBody instantiates a new ConnectivityCheckResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCheckResponseBodyWithDefaults

`func NewConnectivityCheckResponseBodyWithDefaults() *ConnectivityCheckResponseBody`

NewConnectivityCheckResponseBodyWithDefaults instantiates a new ConnectivityCheckResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ConnectivityCheckResponseBody) GetNodes() []NodeEndpointState`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ConnectivityCheckResponseBody) GetNodesOk() (*[]NodeEndpointState, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ConnectivityCheckResponseBody) SetNodes(v []NodeEndpointState)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ConnectivityCheckResponseBody) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *ConnectivityCheckResponseBody) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *ConnectivityCheckResponseBody) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


