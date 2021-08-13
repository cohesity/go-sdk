# CommonSourceRegistrationReponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Source Registration ID. This can be used to retrieve, edit or delete the source registration. | [optional] [readonly] 
**SourceId** | Pointer to **NullableInt64** | ID of top level source object discovered after the registration. | [optional] [readonly] 
**SourceInfo** | Pointer to [**Object**](Object.md) |  | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. | [optional] 
**Connections** | Pointer to [**[]ConnectionConfig**](ConnectionConfig.md) | Specfies the list of connections for the source. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


