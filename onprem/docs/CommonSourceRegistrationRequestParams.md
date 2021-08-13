# CommonSourceRegistrationRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Protection Source. | 
**IsInternalEncrypted** | Pointer to **NullableBool** | Specifies if credentials are encrypted by internal key. | [optional] 
**EncryptionKey** | Pointer to **NullableString** | Specifies the key that user has encrypted the credential with. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | [optional] 
**Connections** | Pointer to [**[]ConnectionConfig**](ConnectionConfig.md) | Specfies the list of connections for the source. | [optional] 

## Methods

### NewCommonSourceRegistrationRequestParams

`func NewCommonSourceRegistrationRequestParams(environment NullableString, ) *CommonSourceRegistrationRequestParams`

NewCommonSourceRegistrationRequestParams instantiates a new CommonSourceRegistrationRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSourceRegistrationRequestParamsWithDefaults

`func NewCommonSourceRegistrationRequestParamsWithDefaults() *CommonSourceRegistrationRequestParams`

NewCommonSourceRegistrationRequestParamsWithDefaults instantiates a new CommonSourceRegistrationRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonSourceRegistrationRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonSourceRegistrationRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonSourceRegistrationRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonSourceRegistrationRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonSourceRegistrationRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsInternalEncrypted

`func (o *CommonSourceRegistrationRequestParams) GetIsInternalEncrypted() bool`

GetIsInternalEncrypted returns the IsInternalEncrypted field if non-nil, zero value otherwise.

### GetIsInternalEncryptedOk

`func (o *CommonSourceRegistrationRequestParams) GetIsInternalEncryptedOk() (*bool, bool)`

GetIsInternalEncryptedOk returns a tuple with the IsInternalEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalEncrypted

`func (o *CommonSourceRegistrationRequestParams) SetIsInternalEncrypted(v bool)`

SetIsInternalEncrypted sets IsInternalEncrypted field to given value.

### HasIsInternalEncrypted

`func (o *CommonSourceRegistrationRequestParams) HasIsInternalEncrypted() bool`

HasIsInternalEncrypted returns a boolean if a field has been set.

### SetIsInternalEncryptedNil

`func (o *CommonSourceRegistrationRequestParams) SetIsInternalEncryptedNil(b bool)`

 SetIsInternalEncryptedNil sets the value for IsInternalEncrypted to be an explicit nil

### UnsetIsInternalEncrypted
`func (o *CommonSourceRegistrationRequestParams) UnsetIsInternalEncrypted()`

UnsetIsInternalEncrypted ensures that no value is present for IsInternalEncrypted, not even an explicit nil
### GetEncryptionKey

`func (o *CommonSourceRegistrationRequestParams) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CommonSourceRegistrationRequestParams) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CommonSourceRegistrationRequestParams) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *CommonSourceRegistrationRequestParams) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *CommonSourceRegistrationRequestParams) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *CommonSourceRegistrationRequestParams) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetConnectionId

`func (o *CommonSourceRegistrationRequestParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CommonSourceRegistrationRequestParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CommonSourceRegistrationRequestParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CommonSourceRegistrationRequestParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *CommonSourceRegistrationRequestParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CommonSourceRegistrationRequestParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnections

`func (o *CommonSourceRegistrationRequestParams) GetConnections() []ConnectionConfig`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CommonSourceRegistrationRequestParams) GetConnectionsOk() (*[]ConnectionConfig, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CommonSourceRegistrationRequestParams) SetConnections(v []ConnectionConfig)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *CommonSourceRegistrationRequestParams) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *CommonSourceRegistrationRequestParams) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *CommonSourceRegistrationRequestParams) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


