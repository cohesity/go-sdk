# ConnectorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalParams** | Pointer to [**AdditionalConnectorParams**](AdditionalConnectorParams.md) |  | [optional] 
**AgentEndpoint** | Pointer to **NullableString** | For some of the environments connection to endpoint is done through an agent. This captures the agent endpoint information. | [optional] 
**AgentPort** | Pointer to **NullableInt32** | Optional agent port to use when connecting to the server. If this is not specified, then environment specific default port will be used. | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**Endpoint** | Pointer to **NullableString** | The endpoint URL of the environment (such as the address of the vCenter instance for a VMware environment, etc). | [optional] 
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**HostType** | Pointer to **NullableInt32** | The host environment type. This is set for kPhysical type environment. | [optional] 
**Id** | Pointer to **NullableInt64** | A unique id associated with this connector params. This is a convenience field and is used to maintain an index to different connection params. This is generated at the time when the source is registered with Magneto. | [optional] 
**PopulateSubnetForAllClusterNodes** | Pointer to **NullableBool** | If set to true, inter agent communcation will be enabled and for every GetAgentInfo call we will fill subnet information of all the nodes in clustered entity. | [optional] 
**Port** | Pointer to **NullableInt32** | Optional port to use when connecting to the server. If this is not specified, then environment specific default port will be used. | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant_id for the environment. This is used to remotely access connectors and executors via bifrost. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment to connect to. | [optional] 
**Version** | Pointer to **NullableInt64** | A version that is associated with the params. This is updated anytime any of the params change. This is used to discard older connector params. | [optional] 

## Methods

### NewConnectorParams

`func NewConnectorParams() *ConnectorParams`

NewConnectorParams instantiates a new ConnectorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorParamsWithDefaults

`func NewConnectorParamsWithDefaults() *ConnectorParams`

NewConnectorParamsWithDefaults instantiates a new ConnectorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalParams

`func (o *ConnectorParams) GetAdditionalParams() AdditionalConnectorParams`

GetAdditionalParams returns the AdditionalParams field if non-nil, zero value otherwise.

### GetAdditionalParamsOk

`func (o *ConnectorParams) GetAdditionalParamsOk() (*AdditionalConnectorParams, bool)`

GetAdditionalParamsOk returns a tuple with the AdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParams

`func (o *ConnectorParams) SetAdditionalParams(v AdditionalConnectorParams)`

SetAdditionalParams sets AdditionalParams field to given value.

### HasAdditionalParams

`func (o *ConnectorParams) HasAdditionalParams() bool`

HasAdditionalParams returns a boolean if a field has been set.

### GetAgentEndpoint

`func (o *ConnectorParams) GetAgentEndpoint() string`

GetAgentEndpoint returns the AgentEndpoint field if non-nil, zero value otherwise.

### GetAgentEndpointOk

`func (o *ConnectorParams) GetAgentEndpointOk() (*string, bool)`

GetAgentEndpointOk returns a tuple with the AgentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEndpoint

`func (o *ConnectorParams) SetAgentEndpoint(v string)`

SetAgentEndpoint sets AgentEndpoint field to given value.

### HasAgentEndpoint

`func (o *ConnectorParams) HasAgentEndpoint() bool`

HasAgentEndpoint returns a boolean if a field has been set.

### SetAgentEndpointNil

`func (o *ConnectorParams) SetAgentEndpointNil(b bool)`

 SetAgentEndpointNil sets the value for AgentEndpoint to be an explicit nil

### UnsetAgentEndpoint
`func (o *ConnectorParams) UnsetAgentEndpoint()`

UnsetAgentEndpoint ensures that no value is present for AgentEndpoint, not even an explicit nil
### GetAgentPort

`func (o *ConnectorParams) GetAgentPort() int32`

GetAgentPort returns the AgentPort field if non-nil, zero value otherwise.

### GetAgentPortOk

`func (o *ConnectorParams) GetAgentPortOk() (*int32, bool)`

GetAgentPortOk returns a tuple with the AgentPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPort

`func (o *ConnectorParams) SetAgentPort(v int32)`

SetAgentPort sets AgentPort field to given value.

### HasAgentPort

`func (o *ConnectorParams) HasAgentPort() bool`

HasAgentPort returns a boolean if a field has been set.

### SetAgentPortNil

`func (o *ConnectorParams) SetAgentPortNil(b bool)`

 SetAgentPortNil sets the value for AgentPort to be an explicit nil

### UnsetAgentPort
`func (o *ConnectorParams) UnsetAgentPort()`

UnsetAgentPort ensures that no value is present for AgentPort, not even an explicit nil
### GetCredentials

`func (o *ConnectorParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectorParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectorParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConnectorParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoint

`func (o *ConnectorParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ConnectorParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ConnectorParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ConnectorParams) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ConnectorParams) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ConnectorParams) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetEntity

`func (o *ConnectorParams) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ConnectorParams) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ConnectorParams) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ConnectorParams) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHostType

`func (o *ConnectorParams) GetHostType() int32`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ConnectorParams) GetHostTypeOk() (*int32, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ConnectorParams) SetHostType(v int32)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *ConnectorParams) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *ConnectorParams) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *ConnectorParams) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetId

`func (o *ConnectorParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConnectorParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConnectorParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPopulateSubnetForAllClusterNodes

`func (o *ConnectorParams) GetPopulateSubnetForAllClusterNodes() bool`

GetPopulateSubnetForAllClusterNodes returns the PopulateSubnetForAllClusterNodes field if non-nil, zero value otherwise.

### GetPopulateSubnetForAllClusterNodesOk

`func (o *ConnectorParams) GetPopulateSubnetForAllClusterNodesOk() (*bool, bool)`

GetPopulateSubnetForAllClusterNodesOk returns a tuple with the PopulateSubnetForAllClusterNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateSubnetForAllClusterNodes

`func (o *ConnectorParams) SetPopulateSubnetForAllClusterNodes(v bool)`

SetPopulateSubnetForAllClusterNodes sets PopulateSubnetForAllClusterNodes field to given value.

### HasPopulateSubnetForAllClusterNodes

`func (o *ConnectorParams) HasPopulateSubnetForAllClusterNodes() bool`

HasPopulateSubnetForAllClusterNodes returns a boolean if a field has been set.

### SetPopulateSubnetForAllClusterNodesNil

`func (o *ConnectorParams) SetPopulateSubnetForAllClusterNodesNil(b bool)`

 SetPopulateSubnetForAllClusterNodesNil sets the value for PopulateSubnetForAllClusterNodes to be an explicit nil

### UnsetPopulateSubnetForAllClusterNodes
`func (o *ConnectorParams) UnsetPopulateSubnetForAllClusterNodes()`

UnsetPopulateSubnetForAllClusterNodes ensures that no value is present for PopulateSubnetForAllClusterNodes, not even an explicit nil
### GetPort

`func (o *ConnectorParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectorParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectorParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ConnectorParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ConnectorParams) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ConnectorParams) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTenantId

`func (o *ConnectorParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectorParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectorParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ConnectorParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ConnectorParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ConnectorParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *ConnectorParams) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorParams) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorParams) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ConnectorParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ConnectorParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVersion

`func (o *ConnectorParams) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorParams) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorParams) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConnectorParams) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConnectorParams) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


