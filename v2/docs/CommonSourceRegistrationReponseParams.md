# CommonSourceRegistrationReponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationStatus** | Pointer to **NullableString** | Specifies the status of the authentication during the registration of a Protection Source. &#39;Pending&#39; indicates the authentication is in progress. &#39;Scheduled&#39; indicates the authentication is scheduled. &#39;Finished&#39; indicates the authentication is completed. &#39;RefreshInProgress&#39; indicates the refresh is in progress. | [optional] [readonly] 
**LastRefreshedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the source was last refreshed in milliseconds. | [optional] [readonly] 
**RegistrationTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the source was registered in milliseconds | [optional] [readonly] 
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a protection source. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. | [optional] 
**Connections** | Pointer to [**[]ConnectionConfig**](ConnectionConfig.md) | Specfies the list of connections for the source. | [optional] 
**ConnectorGroupId** | Pointer to **NullableInt64** | Specifies the connector group id of connector groups. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source. | [optional] 
**Id** | Pointer to **NullableInt64** | Source Registration ID. This can be used to retrieve, edit or delete the source registration. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The user specified name for this source. | [optional] 
**SourceId** | Pointer to **NullableInt64** | ID of top level source object discovered after the registration. | [optional] [readonly] 
**SourceInfo** | Pointer to [**Object**](Object.md) |  | [optional] 

## Methods

### NewCommonSourceRegistrationReponseParams

`func NewCommonSourceRegistrationReponseParams() *CommonSourceRegistrationReponseParams`

NewCommonSourceRegistrationReponseParams instantiates a new CommonSourceRegistrationReponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSourceRegistrationReponseParamsWithDefaults

`func NewCommonSourceRegistrationReponseParamsWithDefaults() *CommonSourceRegistrationReponseParams`

NewCommonSourceRegistrationReponseParamsWithDefaults instantiates a new CommonSourceRegistrationReponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationStatus

`func (o *CommonSourceRegistrationReponseParams) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *CommonSourceRegistrationReponseParams) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *CommonSourceRegistrationReponseParams) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *CommonSourceRegistrationReponseParams) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### SetAuthenticationStatusNil

`func (o *CommonSourceRegistrationReponseParams) SetAuthenticationStatusNil(b bool)`

 SetAuthenticationStatusNil sets the value for AuthenticationStatus to be an explicit nil

### UnsetAuthenticationStatus
`func (o *CommonSourceRegistrationReponseParams) UnsetAuthenticationStatus()`

UnsetAuthenticationStatus ensures that no value is present for AuthenticationStatus, not even an explicit nil
### GetLastRefreshedTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) GetLastRefreshedTimeMsecs() int64`

GetLastRefreshedTimeMsecs returns the LastRefreshedTimeMsecs field if non-nil, zero value otherwise.

### GetLastRefreshedTimeMsecsOk

`func (o *CommonSourceRegistrationReponseParams) GetLastRefreshedTimeMsecsOk() (*int64, bool)`

GetLastRefreshedTimeMsecsOk returns a tuple with the LastRefreshedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) SetLastRefreshedTimeMsecs(v int64)`

SetLastRefreshedTimeMsecs sets LastRefreshedTimeMsecs field to given value.

### HasLastRefreshedTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) HasLastRefreshedTimeMsecs() bool`

HasLastRefreshedTimeMsecs returns a boolean if a field has been set.

### SetLastRefreshedTimeMsecsNil

`func (o *CommonSourceRegistrationReponseParams) SetLastRefreshedTimeMsecsNil(b bool)`

 SetLastRefreshedTimeMsecsNil sets the value for LastRefreshedTimeMsecs to be an explicit nil

### UnsetLastRefreshedTimeMsecs
`func (o *CommonSourceRegistrationReponseParams) UnsetLastRefreshedTimeMsecs()`

UnsetLastRefreshedTimeMsecs ensures that no value is present for LastRefreshedTimeMsecs, not even an explicit nil
### GetRegistrationTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) GetRegistrationTimeMsecs() int64`

GetRegistrationTimeMsecs returns the RegistrationTimeMsecs field if non-nil, zero value otherwise.

### GetRegistrationTimeMsecsOk

`func (o *CommonSourceRegistrationReponseParams) GetRegistrationTimeMsecsOk() (*int64, bool)`

GetRegistrationTimeMsecsOk returns a tuple with the RegistrationTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) SetRegistrationTimeMsecs(v int64)`

SetRegistrationTimeMsecs sets RegistrationTimeMsecs field to given value.

### HasRegistrationTimeMsecs

`func (o *CommonSourceRegistrationReponseParams) HasRegistrationTimeMsecs() bool`

HasRegistrationTimeMsecs returns a boolean if a field has been set.

### SetRegistrationTimeMsecsNil

`func (o *CommonSourceRegistrationReponseParams) SetRegistrationTimeMsecsNil(b bool)`

 SetRegistrationTimeMsecsNil sets the value for RegistrationTimeMsecs to be an explicit nil

### UnsetRegistrationTimeMsecs
`func (o *CommonSourceRegistrationReponseParams) UnsetRegistrationTimeMsecs()`

UnsetRegistrationTimeMsecs ensures that no value is present for RegistrationTimeMsecs, not even an explicit nil
### GetAdvancedConfigs

`func (o *CommonSourceRegistrationReponseParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *CommonSourceRegistrationReponseParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *CommonSourceRegistrationReponseParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *CommonSourceRegistrationReponseParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *CommonSourceRegistrationReponseParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *CommonSourceRegistrationReponseParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetConnectionId

`func (o *CommonSourceRegistrationReponseParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CommonSourceRegistrationReponseParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CommonSourceRegistrationReponseParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CommonSourceRegistrationReponseParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *CommonSourceRegistrationReponseParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CommonSourceRegistrationReponseParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnections

`func (o *CommonSourceRegistrationReponseParams) GetConnections() []ConnectionConfig`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CommonSourceRegistrationReponseParams) GetConnectionsOk() (*[]ConnectionConfig, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CommonSourceRegistrationReponseParams) SetConnections(v []ConnectionConfig)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *CommonSourceRegistrationReponseParams) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *CommonSourceRegistrationReponseParams) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *CommonSourceRegistrationReponseParams) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetConnectorGroupId

`func (o *CommonSourceRegistrationReponseParams) GetConnectorGroupId() int64`

GetConnectorGroupId returns the ConnectorGroupId field if non-nil, zero value otherwise.

### GetConnectorGroupIdOk

`func (o *CommonSourceRegistrationReponseParams) GetConnectorGroupIdOk() (*int64, bool)`

GetConnectorGroupIdOk returns a tuple with the ConnectorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroupId

`func (o *CommonSourceRegistrationReponseParams) SetConnectorGroupId(v int64)`

SetConnectorGroupId sets ConnectorGroupId field to given value.

### HasConnectorGroupId

`func (o *CommonSourceRegistrationReponseParams) HasConnectorGroupId() bool`

HasConnectorGroupId returns a boolean if a field has been set.

### SetConnectorGroupIdNil

`func (o *CommonSourceRegistrationReponseParams) SetConnectorGroupIdNil(b bool)`

 SetConnectorGroupIdNil sets the value for ConnectorGroupId to be an explicit nil

### UnsetConnectorGroupId
`func (o *CommonSourceRegistrationReponseParams) UnsetConnectorGroupId()`

UnsetConnectorGroupId ensures that no value is present for ConnectorGroupId, not even an explicit nil
### GetEnvironment

`func (o *CommonSourceRegistrationReponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonSourceRegistrationReponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonSourceRegistrationReponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CommonSourceRegistrationReponseParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CommonSourceRegistrationReponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonSourceRegistrationReponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *CommonSourceRegistrationReponseParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonSourceRegistrationReponseParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonSourceRegistrationReponseParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonSourceRegistrationReponseParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonSourceRegistrationReponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonSourceRegistrationReponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonSourceRegistrationReponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonSourceRegistrationReponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonSourceRegistrationReponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonSourceRegistrationReponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonSourceRegistrationReponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonSourceRegistrationReponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *CommonSourceRegistrationReponseParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CommonSourceRegistrationReponseParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CommonSourceRegistrationReponseParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CommonSourceRegistrationReponseParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *CommonSourceRegistrationReponseParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *CommonSourceRegistrationReponseParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceInfo

`func (o *CommonSourceRegistrationReponseParams) GetSourceInfo() Object`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *CommonSourceRegistrationReponseParams) GetSourceInfoOk() (*Object, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *CommonSourceRegistrationReponseParams) SetSourceInfo(v Object)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *CommonSourceRegistrationReponseParams) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


