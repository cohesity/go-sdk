# McmSourceRegistrationRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Protection Source. | 
**IsInternalEncrypted** | Pointer to **NullableBool** | Specifies if credentials are encrypted by internal key. | [optional] 
**EncryptionKey** | Pointer to **NullableString** | Specifies the key that user has encrypted the credential with. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | [optional] 
**Connections** | Pointer to [**[]ConnectionConfig**](ConnectionConfig.md) | Specfies the list of connections for the source. | [optional] 
**VmwareParams** | Pointer to [**VmwareSourceRegistrationParams**](VmwareSourceRegistrationParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalSourceRegistrationParams**](PhysicalSourceRegistrationParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasRegistrationParams**](GenericNasRegistrationParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonRegistrationParams**](IsilonRegistrationParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappRegistrationParams**](NetappRegistrationParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileRegistrationParams**](ElastifileRegistrationParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeRegistrationParams**](FlashbladeRegistrationParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsRegistrationParams**](GpfsRegistrationParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraSourceRegistrationParams**](CassandraSourceRegistrationParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDBSourceRegistrationParams**](MongoDBSourceRegistrationParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseSourceRegistrationParams**](CouchbaseSourceRegistrationParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsSourceRegistrationParams**](HdfsSourceRegistrationParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseSourceRegistrationParams**](HbaseSourceRegistrationParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveSourceRegistrationParams**](HiveSourceRegistrationParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaSourceRegistrationParams**](UdaSourceRegistrationParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365SourceRegistrationParams**](Office365SourceRegistrationParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsSourceRegistrationParams**](AwsSourceRegistrationParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVSourceRegistrationParams**](HyperVSourceRegistrationParams.md) |  | [optional] 

## Methods

### NewMcmSourceRegistrationRequestParams

`func NewMcmSourceRegistrationRequestParams(environment NullableString, ) *McmSourceRegistrationRequestParams`

NewMcmSourceRegistrationRequestParams instantiates a new McmSourceRegistrationRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourceRegistrationRequestParamsWithDefaults

`func NewMcmSourceRegistrationRequestParamsWithDefaults() *McmSourceRegistrationRequestParams`

NewMcmSourceRegistrationRequestParamsWithDefaults instantiates a new McmSourceRegistrationRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *McmSourceRegistrationRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmSourceRegistrationRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmSourceRegistrationRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *McmSourceRegistrationRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmSourceRegistrationRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsInternalEncrypted

`func (o *McmSourceRegistrationRequestParams) GetIsInternalEncrypted() bool`

GetIsInternalEncrypted returns the IsInternalEncrypted field if non-nil, zero value otherwise.

### GetIsInternalEncryptedOk

`func (o *McmSourceRegistrationRequestParams) GetIsInternalEncryptedOk() (*bool, bool)`

GetIsInternalEncryptedOk returns a tuple with the IsInternalEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalEncrypted

`func (o *McmSourceRegistrationRequestParams) SetIsInternalEncrypted(v bool)`

SetIsInternalEncrypted sets IsInternalEncrypted field to given value.

### HasIsInternalEncrypted

`func (o *McmSourceRegistrationRequestParams) HasIsInternalEncrypted() bool`

HasIsInternalEncrypted returns a boolean if a field has been set.

### SetIsInternalEncryptedNil

`func (o *McmSourceRegistrationRequestParams) SetIsInternalEncryptedNil(b bool)`

 SetIsInternalEncryptedNil sets the value for IsInternalEncrypted to be an explicit nil

### UnsetIsInternalEncrypted
`func (o *McmSourceRegistrationRequestParams) UnsetIsInternalEncrypted()`

UnsetIsInternalEncrypted ensures that no value is present for IsInternalEncrypted, not even an explicit nil
### GetEncryptionKey

`func (o *McmSourceRegistrationRequestParams) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *McmSourceRegistrationRequestParams) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *McmSourceRegistrationRequestParams) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *McmSourceRegistrationRequestParams) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *McmSourceRegistrationRequestParams) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *McmSourceRegistrationRequestParams) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetConnectionId

`func (o *McmSourceRegistrationRequestParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *McmSourceRegistrationRequestParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *McmSourceRegistrationRequestParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *McmSourceRegistrationRequestParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *McmSourceRegistrationRequestParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *McmSourceRegistrationRequestParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnections

`func (o *McmSourceRegistrationRequestParams) GetConnections() []ConnectionConfig`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *McmSourceRegistrationRequestParams) GetConnectionsOk() (*[]ConnectionConfig, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *McmSourceRegistrationRequestParams) SetConnections(v []ConnectionConfig)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *McmSourceRegistrationRequestParams) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *McmSourceRegistrationRequestParams) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *McmSourceRegistrationRequestParams) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetVmwareParams

`func (o *McmSourceRegistrationRequestParams) GetVmwareParams() VmwareSourceRegistrationParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *McmSourceRegistrationRequestParams) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *McmSourceRegistrationRequestParams) SetVmwareParams(v VmwareSourceRegistrationParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *McmSourceRegistrationRequestParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *McmSourceRegistrationRequestParams) GetPhysicalParams() PhysicalSourceRegistrationParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *McmSourceRegistrationRequestParams) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *McmSourceRegistrationRequestParams) SetPhysicalParams(v PhysicalSourceRegistrationParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *McmSourceRegistrationRequestParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *McmSourceRegistrationRequestParams) GetGenericNasParams() GenericNasRegistrationParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *McmSourceRegistrationRequestParams) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *McmSourceRegistrationRequestParams) SetGenericNasParams(v GenericNasRegistrationParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *McmSourceRegistrationRequestParams) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *McmSourceRegistrationRequestParams) GetIsilonParams() IsilonRegistrationParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *McmSourceRegistrationRequestParams) GetIsilonParamsOk() (*IsilonRegistrationParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *McmSourceRegistrationRequestParams) SetIsilonParams(v IsilonRegistrationParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *McmSourceRegistrationRequestParams) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *McmSourceRegistrationRequestParams) GetNetappParams() NetappRegistrationParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *McmSourceRegistrationRequestParams) GetNetappParamsOk() (*NetappRegistrationParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *McmSourceRegistrationRequestParams) SetNetappParams(v NetappRegistrationParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *McmSourceRegistrationRequestParams) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *McmSourceRegistrationRequestParams) GetElastifileParams() ElastifileRegistrationParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *McmSourceRegistrationRequestParams) GetElastifileParamsOk() (*ElastifileRegistrationParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *McmSourceRegistrationRequestParams) SetElastifileParams(v ElastifileRegistrationParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *McmSourceRegistrationRequestParams) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *McmSourceRegistrationRequestParams) GetFlashbladeParams() FlashbladeRegistrationParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *McmSourceRegistrationRequestParams) GetFlashbladeParamsOk() (*FlashbladeRegistrationParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *McmSourceRegistrationRequestParams) SetFlashbladeParams(v FlashbladeRegistrationParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *McmSourceRegistrationRequestParams) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *McmSourceRegistrationRequestParams) GetGpfsParams() GpfsRegistrationParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *McmSourceRegistrationRequestParams) GetGpfsParamsOk() (*GpfsRegistrationParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *McmSourceRegistrationRequestParams) SetGpfsParams(v GpfsRegistrationParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *McmSourceRegistrationRequestParams) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *McmSourceRegistrationRequestParams) GetCassandraParams() CassandraSourceRegistrationParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *McmSourceRegistrationRequestParams) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *McmSourceRegistrationRequestParams) SetCassandraParams(v CassandraSourceRegistrationParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *McmSourceRegistrationRequestParams) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *McmSourceRegistrationRequestParams) GetMongodbParams() MongoDBSourceRegistrationParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *McmSourceRegistrationRequestParams) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *McmSourceRegistrationRequestParams) SetMongodbParams(v MongoDBSourceRegistrationParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *McmSourceRegistrationRequestParams) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *McmSourceRegistrationRequestParams) GetCouchbaseParams() CouchbaseSourceRegistrationParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *McmSourceRegistrationRequestParams) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *McmSourceRegistrationRequestParams) SetCouchbaseParams(v CouchbaseSourceRegistrationParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *McmSourceRegistrationRequestParams) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *McmSourceRegistrationRequestParams) GetHdfsParams() HdfsSourceRegistrationParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *McmSourceRegistrationRequestParams) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *McmSourceRegistrationRequestParams) SetHdfsParams(v HdfsSourceRegistrationParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *McmSourceRegistrationRequestParams) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *McmSourceRegistrationRequestParams) GetHbaseParams() HbaseSourceRegistrationParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *McmSourceRegistrationRequestParams) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *McmSourceRegistrationRequestParams) SetHbaseParams(v HbaseSourceRegistrationParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *McmSourceRegistrationRequestParams) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *McmSourceRegistrationRequestParams) GetHiveParams() HiveSourceRegistrationParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *McmSourceRegistrationRequestParams) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *McmSourceRegistrationRequestParams) SetHiveParams(v HiveSourceRegistrationParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *McmSourceRegistrationRequestParams) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *McmSourceRegistrationRequestParams) GetUdaParams() UdaSourceRegistrationParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *McmSourceRegistrationRequestParams) GetUdaParamsOk() (*UdaSourceRegistrationParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *McmSourceRegistrationRequestParams) SetUdaParams(v UdaSourceRegistrationParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *McmSourceRegistrationRequestParams) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *McmSourceRegistrationRequestParams) GetOffice365Params() Office365SourceRegistrationParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *McmSourceRegistrationRequestParams) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *McmSourceRegistrationRequestParams) SetOffice365Params(v Office365SourceRegistrationParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *McmSourceRegistrationRequestParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *McmSourceRegistrationRequestParams) GetAwsParams() AwsSourceRegistrationParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *McmSourceRegistrationRequestParams) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *McmSourceRegistrationRequestParams) SetAwsParams(v AwsSourceRegistrationParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *McmSourceRegistrationRequestParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *McmSourceRegistrationRequestParams) GetHypervParams() HyperVSourceRegistrationParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *McmSourceRegistrationRequestParams) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *McmSourceRegistrationRequestParams) SetHypervParams(v HyperVSourceRegistrationParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *McmSourceRegistrationRequestParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


