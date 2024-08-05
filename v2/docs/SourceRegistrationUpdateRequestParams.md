# SourceRegistrationUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a protection source. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | [optional] 
**Connections** | Pointer to [**[]ConnectionConfig**](ConnectionConfig.md) | Specfies the list of connections for the source. | [optional] 
**ConnectorGroupId** | Pointer to **NullableInt64** | Specifies the connector group id of connector groups. | [optional] 
**EncryptionKey** | Pointer to **NullableString** | Specifies the key that user has encrypted the credential with. | [optional] 
**Environment** | **NullableString** | Specifies the environment type of the Protection Source. | 
**IsInternalEncrypted** | Pointer to **NullableBool** | Specifies if credentials are encrypted by internal key. | [optional] 
**Name** | Pointer to **NullableString** | A user specified name for this source. | [optional] 
**AwsParams** | Pointer to [**AwsSourceRegistrationParams**](AwsSourceRegistrationParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureSourceRegistrationParams**](AzureSourceRegistrationParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraSourceRegistrationParams**](CassandraSourceRegistrationParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseSourceRegistrationParams**](CouchbaseSourceRegistrationParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileRegistrationParams**](ElastifileRegistrationParams.md) |  | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterSourceRegistrationParams**](ExperimentalAdapterSourceRegistrationParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeRegistrationParams**](FlashbladeRegistrationParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasRegistrationParams**](GenericNasRegistrationParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsRegistrationParams**](GpfsRegistrationParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseSourceRegistrationParams**](HbaseSourceRegistrationParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsSourceRegistrationParams**](HdfsSourceRegistrationParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveSourceRegistrationParams**](HiveSourceRegistrationParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVSourceRegistrationParams**](HyperVSourceRegistrationParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonRegistrationParams**](IsilonRegistrationParams.md) |  | [optional] 
**LastModifiedTimestampUsecs** | Pointer to **NullableInt64** | Specifies the last time this protection source was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection source was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | [optional] 
**MongodbParams** | Pointer to [**MongoDBSourceRegistrationParams**](MongoDBSourceRegistrationParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappRegistrationParams**](NetappRegistrationParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365SourceRegistrationParams**](Office365SourceRegistrationParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalSourceRegistrationParams**](PhysicalSourceRegistrationParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaSourceRegistrationParams**](SapHanaSourceRegistrationParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**SfdcSourceRegistrationParams**](SfdcSourceRegistrationParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaSourceRegistrationParams**](UdaSourceRegistrationParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareSourceRegistrationParams**](VmwareSourceRegistrationParams.md) |  | [optional] 

## Methods

### NewSourceRegistrationUpdateRequestParams

`func NewSourceRegistrationUpdateRequestParams(environment NullableString, ) *SourceRegistrationUpdateRequestParams`

NewSourceRegistrationUpdateRequestParams instantiates a new SourceRegistrationUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRegistrationUpdateRequestParamsWithDefaults

`func NewSourceRegistrationUpdateRequestParamsWithDefaults() *SourceRegistrationUpdateRequestParams`

NewSourceRegistrationUpdateRequestParamsWithDefaults instantiates a new SourceRegistrationUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfigs

`func (o *SourceRegistrationUpdateRequestParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *SourceRegistrationUpdateRequestParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *SourceRegistrationUpdateRequestParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *SourceRegistrationUpdateRequestParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *SourceRegistrationUpdateRequestParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *SourceRegistrationUpdateRequestParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetConnectionId

`func (o *SourceRegistrationUpdateRequestParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SourceRegistrationUpdateRequestParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SourceRegistrationUpdateRequestParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SourceRegistrationUpdateRequestParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *SourceRegistrationUpdateRequestParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *SourceRegistrationUpdateRequestParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnections

`func (o *SourceRegistrationUpdateRequestParams) GetConnections() []ConnectionConfig`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *SourceRegistrationUpdateRequestParams) GetConnectionsOk() (*[]ConnectionConfig, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *SourceRegistrationUpdateRequestParams) SetConnections(v []ConnectionConfig)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *SourceRegistrationUpdateRequestParams) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *SourceRegistrationUpdateRequestParams) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *SourceRegistrationUpdateRequestParams) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetConnectorGroupId

`func (o *SourceRegistrationUpdateRequestParams) GetConnectorGroupId() int64`

GetConnectorGroupId returns the ConnectorGroupId field if non-nil, zero value otherwise.

### GetConnectorGroupIdOk

`func (o *SourceRegistrationUpdateRequestParams) GetConnectorGroupIdOk() (*int64, bool)`

GetConnectorGroupIdOk returns a tuple with the ConnectorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroupId

`func (o *SourceRegistrationUpdateRequestParams) SetConnectorGroupId(v int64)`

SetConnectorGroupId sets ConnectorGroupId field to given value.

### HasConnectorGroupId

`func (o *SourceRegistrationUpdateRequestParams) HasConnectorGroupId() bool`

HasConnectorGroupId returns a boolean if a field has been set.

### SetConnectorGroupIdNil

`func (o *SourceRegistrationUpdateRequestParams) SetConnectorGroupIdNil(b bool)`

 SetConnectorGroupIdNil sets the value for ConnectorGroupId to be an explicit nil

### UnsetConnectorGroupId
`func (o *SourceRegistrationUpdateRequestParams) UnsetConnectorGroupId()`

UnsetConnectorGroupId ensures that no value is present for ConnectorGroupId, not even an explicit nil
### GetEncryptionKey

`func (o *SourceRegistrationUpdateRequestParams) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *SourceRegistrationUpdateRequestParams) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *SourceRegistrationUpdateRequestParams) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *SourceRegistrationUpdateRequestParams) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *SourceRegistrationUpdateRequestParams) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *SourceRegistrationUpdateRequestParams) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetEnvironment

`func (o *SourceRegistrationUpdateRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SourceRegistrationUpdateRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SourceRegistrationUpdateRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *SourceRegistrationUpdateRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SourceRegistrationUpdateRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsInternalEncrypted

`func (o *SourceRegistrationUpdateRequestParams) GetIsInternalEncrypted() bool`

GetIsInternalEncrypted returns the IsInternalEncrypted field if non-nil, zero value otherwise.

### GetIsInternalEncryptedOk

`func (o *SourceRegistrationUpdateRequestParams) GetIsInternalEncryptedOk() (*bool, bool)`

GetIsInternalEncryptedOk returns a tuple with the IsInternalEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalEncrypted

`func (o *SourceRegistrationUpdateRequestParams) SetIsInternalEncrypted(v bool)`

SetIsInternalEncrypted sets IsInternalEncrypted field to given value.

### HasIsInternalEncrypted

`func (o *SourceRegistrationUpdateRequestParams) HasIsInternalEncrypted() bool`

HasIsInternalEncrypted returns a boolean if a field has been set.

### SetIsInternalEncryptedNil

`func (o *SourceRegistrationUpdateRequestParams) SetIsInternalEncryptedNil(b bool)`

 SetIsInternalEncryptedNil sets the value for IsInternalEncrypted to be an explicit nil

### UnsetIsInternalEncrypted
`func (o *SourceRegistrationUpdateRequestParams) UnsetIsInternalEncrypted()`

UnsetIsInternalEncrypted ensures that no value is present for IsInternalEncrypted, not even an explicit nil
### GetName

`func (o *SourceRegistrationUpdateRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceRegistrationUpdateRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceRegistrationUpdateRequestParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceRegistrationUpdateRequestParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SourceRegistrationUpdateRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SourceRegistrationUpdateRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAwsParams

`func (o *SourceRegistrationUpdateRequestParams) GetAwsParams() AwsSourceRegistrationParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *SourceRegistrationUpdateRequestParams) SetAwsParams(v AwsSourceRegistrationParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *SourceRegistrationUpdateRequestParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *SourceRegistrationUpdateRequestParams) GetAzureParams() AzureSourceRegistrationParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetAzureParamsOk() (*AzureSourceRegistrationParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *SourceRegistrationUpdateRequestParams) SetAzureParams(v AzureSourceRegistrationParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *SourceRegistrationUpdateRequestParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *SourceRegistrationUpdateRequestParams) GetCassandraParams() CassandraSourceRegistrationParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SourceRegistrationUpdateRequestParams) SetCassandraParams(v CassandraSourceRegistrationParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SourceRegistrationUpdateRequestParams) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *SourceRegistrationUpdateRequestParams) GetCouchbaseParams() CouchbaseSourceRegistrationParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *SourceRegistrationUpdateRequestParams) SetCouchbaseParams(v CouchbaseSourceRegistrationParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *SourceRegistrationUpdateRequestParams) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *SourceRegistrationUpdateRequestParams) GetElastifileParams() ElastifileRegistrationParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetElastifileParamsOk() (*ElastifileRegistrationParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *SourceRegistrationUpdateRequestParams) SetElastifileParams(v ElastifileRegistrationParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *SourceRegistrationUpdateRequestParams) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *SourceRegistrationUpdateRequestParams) GetExperimentalAdapterParams() ExperimentalAdapterSourceRegistrationParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterSourceRegistrationParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *SourceRegistrationUpdateRequestParams) SetExperimentalAdapterParams(v ExperimentalAdapterSourceRegistrationParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *SourceRegistrationUpdateRequestParams) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *SourceRegistrationUpdateRequestParams) GetFlashbladeParams() FlashbladeRegistrationParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetFlashbladeParamsOk() (*FlashbladeRegistrationParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *SourceRegistrationUpdateRequestParams) SetFlashbladeParams(v FlashbladeRegistrationParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *SourceRegistrationUpdateRequestParams) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *SourceRegistrationUpdateRequestParams) GetGenericNasParams() GenericNasRegistrationParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *SourceRegistrationUpdateRequestParams) SetGenericNasParams(v GenericNasRegistrationParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *SourceRegistrationUpdateRequestParams) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *SourceRegistrationUpdateRequestParams) GetGpfsParams() GpfsRegistrationParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetGpfsParamsOk() (*GpfsRegistrationParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *SourceRegistrationUpdateRequestParams) SetGpfsParams(v GpfsRegistrationParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *SourceRegistrationUpdateRequestParams) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *SourceRegistrationUpdateRequestParams) GetHbaseParams() HbaseSourceRegistrationParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *SourceRegistrationUpdateRequestParams) SetHbaseParams(v HbaseSourceRegistrationParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *SourceRegistrationUpdateRequestParams) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *SourceRegistrationUpdateRequestParams) GetHdfsParams() HdfsSourceRegistrationParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *SourceRegistrationUpdateRequestParams) SetHdfsParams(v HdfsSourceRegistrationParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *SourceRegistrationUpdateRequestParams) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *SourceRegistrationUpdateRequestParams) GetHiveParams() HiveSourceRegistrationParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *SourceRegistrationUpdateRequestParams) SetHiveParams(v HiveSourceRegistrationParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *SourceRegistrationUpdateRequestParams) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *SourceRegistrationUpdateRequestParams) GetHypervParams() HyperVSourceRegistrationParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *SourceRegistrationUpdateRequestParams) SetHypervParams(v HyperVSourceRegistrationParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *SourceRegistrationUpdateRequestParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *SourceRegistrationUpdateRequestParams) GetIsilonParams() IsilonRegistrationParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetIsilonParamsOk() (*IsilonRegistrationParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *SourceRegistrationUpdateRequestParams) SetIsilonParams(v IsilonRegistrationParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *SourceRegistrationUpdateRequestParams) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetLastModifiedTimestampUsecs

`func (o *SourceRegistrationUpdateRequestParams) GetLastModifiedTimestampUsecs() int64`

GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimestampUsecsOk

`func (o *SourceRegistrationUpdateRequestParams) GetLastModifiedTimestampUsecsOk() (*int64, bool)`

GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestampUsecs

`func (o *SourceRegistrationUpdateRequestParams) SetLastModifiedTimestampUsecs(v int64)`

SetLastModifiedTimestampUsecs sets LastModifiedTimestampUsecs field to given value.

### HasLastModifiedTimestampUsecs

`func (o *SourceRegistrationUpdateRequestParams) HasLastModifiedTimestampUsecs() bool`

HasLastModifiedTimestampUsecs returns a boolean if a field has been set.

### SetLastModifiedTimestampUsecsNil

`func (o *SourceRegistrationUpdateRequestParams) SetLastModifiedTimestampUsecsNil(b bool)`

 SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil

### UnsetLastModifiedTimestampUsecs
`func (o *SourceRegistrationUpdateRequestParams) UnsetLastModifiedTimestampUsecs()`

UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
### GetMongodbParams

`func (o *SourceRegistrationUpdateRequestParams) GetMongodbParams() MongoDBSourceRegistrationParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *SourceRegistrationUpdateRequestParams) SetMongodbParams(v MongoDBSourceRegistrationParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *SourceRegistrationUpdateRequestParams) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *SourceRegistrationUpdateRequestParams) GetNetappParams() NetappRegistrationParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetNetappParamsOk() (*NetappRegistrationParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *SourceRegistrationUpdateRequestParams) SetNetappParams(v NetappRegistrationParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *SourceRegistrationUpdateRequestParams) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *SourceRegistrationUpdateRequestParams) GetOffice365Params() Office365SourceRegistrationParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *SourceRegistrationUpdateRequestParams) SetOffice365Params(v Office365SourceRegistrationParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *SourceRegistrationUpdateRequestParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *SourceRegistrationUpdateRequestParams) GetPhysicalParams() PhysicalSourceRegistrationParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *SourceRegistrationUpdateRequestParams) SetPhysicalParams(v PhysicalSourceRegistrationParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *SourceRegistrationUpdateRequestParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *SourceRegistrationUpdateRequestParams) GetSapHanaParams() SapHanaSourceRegistrationParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetSapHanaParamsOk() (*SapHanaSourceRegistrationParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *SourceRegistrationUpdateRequestParams) SetSapHanaParams(v SapHanaSourceRegistrationParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *SourceRegistrationUpdateRequestParams) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *SourceRegistrationUpdateRequestParams) GetSfdcParams() SfdcSourceRegistrationParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetSfdcParamsOk() (*SfdcSourceRegistrationParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *SourceRegistrationUpdateRequestParams) SetSfdcParams(v SfdcSourceRegistrationParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *SourceRegistrationUpdateRequestParams) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *SourceRegistrationUpdateRequestParams) GetUdaParams() UdaSourceRegistrationParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetUdaParamsOk() (*UdaSourceRegistrationParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *SourceRegistrationUpdateRequestParams) SetUdaParams(v UdaSourceRegistrationParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *SourceRegistrationUpdateRequestParams) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *SourceRegistrationUpdateRequestParams) GetVmwareParams() VmwareSourceRegistrationParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *SourceRegistrationUpdateRequestParams) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *SourceRegistrationUpdateRequestParams) SetVmwareParams(v VmwareSourceRegistrationParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *SourceRegistrationUpdateRequestParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


