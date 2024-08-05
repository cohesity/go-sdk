# NodeEndpointState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckTimestampUsecs** | Pointer to **NullableInt64** | Specifies the time in Epoch in micro seconds when the check is performed. | [optional] 
**DnsServerReachability** | Pointer to [**EndpointConnectionState**](EndpointConnectionState.md) |  | [optional] 
**Endpoints** | Pointer to [**[]EndpointConnectionState**](EndpointConnectionState.md) | Specifies the results of the endpoints. | [optional] 
**GatewayReachability** | Pointer to [**EndpointConnectionState**](EndpointConnectionState.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Id of the node. | [optional] 
**NtpServerReachability** | Pointer to [**EndpointConnectionState**](EndpointConnectionState.md) |  | [optional] 

## Methods

### NewNodeEndpointState

`func NewNodeEndpointState() *NodeEndpointState`

NewNodeEndpointState instantiates a new NodeEndpointState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeEndpointStateWithDefaults

`func NewNodeEndpointStateWithDefaults() *NodeEndpointState`

NewNodeEndpointStateWithDefaults instantiates a new NodeEndpointState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTimestampUsecs

`func (o *NodeEndpointState) GetCheckTimestampUsecs() int64`

GetCheckTimestampUsecs returns the CheckTimestampUsecs field if non-nil, zero value otherwise.

### GetCheckTimestampUsecsOk

`func (o *NodeEndpointState) GetCheckTimestampUsecsOk() (*int64, bool)`

GetCheckTimestampUsecsOk returns a tuple with the CheckTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTimestampUsecs

`func (o *NodeEndpointState) SetCheckTimestampUsecs(v int64)`

SetCheckTimestampUsecs sets CheckTimestampUsecs field to given value.

### HasCheckTimestampUsecs

`func (o *NodeEndpointState) HasCheckTimestampUsecs() bool`

HasCheckTimestampUsecs returns a boolean if a field has been set.

### SetCheckTimestampUsecsNil

`func (o *NodeEndpointState) SetCheckTimestampUsecsNil(b bool)`

 SetCheckTimestampUsecsNil sets the value for CheckTimestampUsecs to be an explicit nil

### UnsetCheckTimestampUsecs
`func (o *NodeEndpointState) UnsetCheckTimestampUsecs()`

UnsetCheckTimestampUsecs ensures that no value is present for CheckTimestampUsecs, not even an explicit nil
### GetDnsServerReachability

`func (o *NodeEndpointState) GetDnsServerReachability() EndpointConnectionState`

GetDnsServerReachability returns the DnsServerReachability field if non-nil, zero value otherwise.

### GetDnsServerReachabilityOk

`func (o *NodeEndpointState) GetDnsServerReachabilityOk() (*EndpointConnectionState, bool)`

GetDnsServerReachabilityOk returns a tuple with the DnsServerReachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerReachability

`func (o *NodeEndpointState) SetDnsServerReachability(v EndpointConnectionState)`

SetDnsServerReachability sets DnsServerReachability field to given value.

### HasDnsServerReachability

`func (o *NodeEndpointState) HasDnsServerReachability() bool`

HasDnsServerReachability returns a boolean if a field has been set.

### GetEndpoints

`func (o *NodeEndpointState) GetEndpoints() []EndpointConnectionState`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NodeEndpointState) GetEndpointsOk() (*[]EndpointConnectionState, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *NodeEndpointState) SetEndpoints(v []EndpointConnectionState)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *NodeEndpointState) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *NodeEndpointState) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *NodeEndpointState) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetGatewayReachability

`func (o *NodeEndpointState) GetGatewayReachability() EndpointConnectionState`

GetGatewayReachability returns the GatewayReachability field if non-nil, zero value otherwise.

### GetGatewayReachabilityOk

`func (o *NodeEndpointState) GetGatewayReachabilityOk() (*EndpointConnectionState, bool)`

GetGatewayReachabilityOk returns a tuple with the GatewayReachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayReachability

`func (o *NodeEndpointState) SetGatewayReachability(v EndpointConnectionState)`

SetGatewayReachability sets GatewayReachability field to given value.

### HasGatewayReachability

`func (o *NodeEndpointState) HasGatewayReachability() bool`

HasGatewayReachability returns a boolean if a field has been set.

### GetId

`func (o *NodeEndpointState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeEndpointState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeEndpointState) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeEndpointState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeEndpointState) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeEndpointState) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNtpServerReachability

`func (o *NodeEndpointState) GetNtpServerReachability() EndpointConnectionState`

GetNtpServerReachability returns the NtpServerReachability field if non-nil, zero value otherwise.

### GetNtpServerReachabilityOk

`func (o *NodeEndpointState) GetNtpServerReachabilityOk() (*EndpointConnectionState, bool)`

GetNtpServerReachabilityOk returns a tuple with the NtpServerReachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerReachability

`func (o *NodeEndpointState) SetNtpServerReachability(v EndpointConnectionState)`

SetNtpServerReachability sets NtpServerReachability field to given value.

### HasNtpServerReachability

`func (o *NodeEndpointState) HasNtpServerReachability() bool`

HasNtpServerReachability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


