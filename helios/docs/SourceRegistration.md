# SourceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Source Registration ID. This can be used to retrieve, edit or delete the source registration. | [optional] [readonly] 
**SourceId** | Pointer to **NullableInt64** | ID of top level source object discovered after the registration. | [optional] [readonly] 
**SourceInfo** | Pointer to [**Object**](Object.md) |  | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. | [optional] 
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

### NewSourceRegistration

`func NewSourceRegistration() *SourceRegistration`

NewSourceRegistration instantiates a new SourceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRegistrationWithDefaults

`func NewSourceRegistrationWithDefaults() *SourceRegistration`

NewSourceRegistrationWithDefaults instantiates a new SourceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceRegistration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceRegistration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceRegistration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SourceRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SourceRegistration) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SourceRegistration) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceId

`func (o *SourceRegistration) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceRegistration) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceRegistration) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SourceRegistration) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *SourceRegistration) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *SourceRegistration) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceInfo

`func (o *SourceRegistration) GetSourceInfo() Object`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *SourceRegistration) GetSourceInfoOk() (*Object, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *SourceRegistration) SetSourceInfo(v Object)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *SourceRegistration) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetEnvironment

`func (o *SourceRegistration) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SourceRegistration) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SourceRegistration) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SourceRegistration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *SourceRegistration) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SourceRegistration) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetConnectionId

`func (o *SourceRegistration) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SourceRegistration) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SourceRegistration) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SourceRegistration) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *SourceRegistration) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *SourceRegistration) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnections

`func (o *SourceRegistration) GetConnections() []ConnectionConfig`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *SourceRegistration) GetConnectionsOk() (*[]ConnectionConfig, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *SourceRegistration) SetConnections(v []ConnectionConfig)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *SourceRegistration) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *SourceRegistration) SetConnectionsNil(b bool)`

 SetConnectionsNil sets the value for Connections to be an explicit nil

### UnsetConnections
`func (o *SourceRegistration) UnsetConnections()`

UnsetConnections ensures that no value is present for Connections, not even an explicit nil
### GetVmwareParams

`func (o *SourceRegistration) GetVmwareParams() VmwareSourceRegistrationParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *SourceRegistration) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *SourceRegistration) SetVmwareParams(v VmwareSourceRegistrationParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *SourceRegistration) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *SourceRegistration) GetPhysicalParams() PhysicalSourceRegistrationParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *SourceRegistration) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *SourceRegistration) SetPhysicalParams(v PhysicalSourceRegistrationParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *SourceRegistration) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *SourceRegistration) GetGenericNasParams() GenericNasRegistrationParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *SourceRegistration) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *SourceRegistration) SetGenericNasParams(v GenericNasRegistrationParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *SourceRegistration) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *SourceRegistration) GetIsilonParams() IsilonRegistrationParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *SourceRegistration) GetIsilonParamsOk() (*IsilonRegistrationParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *SourceRegistration) SetIsilonParams(v IsilonRegistrationParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *SourceRegistration) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *SourceRegistration) GetNetappParams() NetappRegistrationParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *SourceRegistration) GetNetappParamsOk() (*NetappRegistrationParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *SourceRegistration) SetNetappParams(v NetappRegistrationParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *SourceRegistration) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *SourceRegistration) GetElastifileParams() ElastifileRegistrationParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *SourceRegistration) GetElastifileParamsOk() (*ElastifileRegistrationParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *SourceRegistration) SetElastifileParams(v ElastifileRegistrationParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *SourceRegistration) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *SourceRegistration) GetFlashbladeParams() FlashbladeRegistrationParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *SourceRegistration) GetFlashbladeParamsOk() (*FlashbladeRegistrationParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *SourceRegistration) SetFlashbladeParams(v FlashbladeRegistrationParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *SourceRegistration) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *SourceRegistration) GetGpfsParams() GpfsRegistrationParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *SourceRegistration) GetGpfsParamsOk() (*GpfsRegistrationParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *SourceRegistration) SetGpfsParams(v GpfsRegistrationParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *SourceRegistration) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *SourceRegistration) GetCassandraParams() CassandraSourceRegistrationParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SourceRegistration) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SourceRegistration) SetCassandraParams(v CassandraSourceRegistrationParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SourceRegistration) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *SourceRegistration) GetMongodbParams() MongoDBSourceRegistrationParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *SourceRegistration) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *SourceRegistration) SetMongodbParams(v MongoDBSourceRegistrationParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *SourceRegistration) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *SourceRegistration) GetCouchbaseParams() CouchbaseSourceRegistrationParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *SourceRegistration) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *SourceRegistration) SetCouchbaseParams(v CouchbaseSourceRegistrationParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *SourceRegistration) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *SourceRegistration) GetHdfsParams() HdfsSourceRegistrationParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *SourceRegistration) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *SourceRegistration) SetHdfsParams(v HdfsSourceRegistrationParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *SourceRegistration) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *SourceRegistration) GetHbaseParams() HbaseSourceRegistrationParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *SourceRegistration) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *SourceRegistration) SetHbaseParams(v HbaseSourceRegistrationParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *SourceRegistration) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *SourceRegistration) GetHiveParams() HiveSourceRegistrationParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *SourceRegistration) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *SourceRegistration) SetHiveParams(v HiveSourceRegistrationParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *SourceRegistration) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *SourceRegistration) GetUdaParams() UdaSourceRegistrationParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *SourceRegistration) GetUdaParamsOk() (*UdaSourceRegistrationParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *SourceRegistration) SetUdaParams(v UdaSourceRegistrationParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *SourceRegistration) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *SourceRegistration) GetOffice365Params() Office365SourceRegistrationParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *SourceRegistration) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *SourceRegistration) SetOffice365Params(v Office365SourceRegistrationParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *SourceRegistration) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *SourceRegistration) GetAwsParams() AwsSourceRegistrationParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *SourceRegistration) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *SourceRegistration) SetAwsParams(v AwsSourceRegistrationParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *SourceRegistration) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *SourceRegistration) GetHypervParams() HyperVSourceRegistrationParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *SourceRegistration) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *SourceRegistration) SetHypervParams(v HyperVSourceRegistrationParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *SourceRegistration) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


