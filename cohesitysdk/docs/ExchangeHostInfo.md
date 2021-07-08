# ExchangeHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentStatus** | Pointer to **NullableString** | Specifies the status of the agent on the Exchange host. Specifies the status of agent on Exchange Application Server. &#39;kSupported&#39; indicates the agent is supported for Exchange data protection. &#39;kUnSupported&#39; indicates the agent is not supported for Exchange data protection. &#39;kUpgrade&#39; indicates the agent of server need to be upgraded. | [optional] 
**Endpoint** | Pointer to **NullableString** | Specifies the endpoint of the Exchange host. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the guid of the Exchange host. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the display name of the Exchange host. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies the Protection source id of the Physical Host if the Exchange application is already registered on the physical host with the above endpoint. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the registration of the Exchange Host. Specifies the status of registration of Exchange Application Server. &#39;kUnknown&#39; indicates the status is not known. &#39;kHealthy&#39; indicates the status is healty and is registered as Exchange Server. &#39;kUnHealthy&#39; indicates the exchange application is registered on the physical server but it is unreachable now. &#39;kUnregistered&#39; indicates the server is not registered as physical source. &#39;kUnreachable&#39; indicates the server is not reachable from the cohesity cluster or the cohesity protection server is not installed on the exchange server. &#39;kDetached&#39; indicates the server is removed from the ExchangeDAG. | [optional] 

## Methods

### NewExchangeHostInfo

`func NewExchangeHostInfo() *ExchangeHostInfo`

NewExchangeHostInfo instantiates a new ExchangeHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeHostInfoWithDefaults

`func NewExchangeHostInfoWithDefaults() *ExchangeHostInfo`

NewExchangeHostInfoWithDefaults instantiates a new ExchangeHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentStatus

`func (o *ExchangeHostInfo) GetAgentStatus() string`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *ExchangeHostInfo) GetAgentStatusOk() (*string, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *ExchangeHostInfo) SetAgentStatus(v string)`

SetAgentStatus sets AgentStatus field to given value.

### HasAgentStatus

`func (o *ExchangeHostInfo) HasAgentStatus() bool`

HasAgentStatus returns a boolean if a field has been set.

### SetAgentStatusNil

`func (o *ExchangeHostInfo) SetAgentStatusNil(b bool)`

 SetAgentStatusNil sets the value for AgentStatus to be an explicit nil

### UnsetAgentStatus
`func (o *ExchangeHostInfo) UnsetAgentStatus()`

UnsetAgentStatus ensures that no value is present for AgentStatus, not even an explicit nil
### GetEndpoint

`func (o *ExchangeHostInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ExchangeHostInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ExchangeHostInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ExchangeHostInfo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ExchangeHostInfo) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ExchangeHostInfo) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetGuid

`func (o *ExchangeHostInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ExchangeHostInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ExchangeHostInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ExchangeHostInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ExchangeHostInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ExchangeHostInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *ExchangeHostInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeHostInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeHostInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeHostInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeHostInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeHostInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtectionSourceId

`func (o *ExchangeHostInfo) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *ExchangeHostInfo) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *ExchangeHostInfo) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *ExchangeHostInfo) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *ExchangeHostInfo) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *ExchangeHostInfo) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
### GetStatus

`func (o *ExchangeHostInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangeHostInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangeHostInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExchangeHostInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExchangeHostInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExchangeHostInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


