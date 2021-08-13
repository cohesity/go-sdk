# McmSourceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] [readonly] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] [readonly] 
**RegionId** | Pointer to **NullableString** | Specifies the region id. | [optional] [readonly] 
**Id** | Pointer to **NullableString** | Source Registration ID. This can be used to retrieve, edit or delete the source registration. | [optional] [readonly] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. | [optional] 
**SourceId** | Pointer to **NullableString** | ID of top level source object discovered after the registration. | [optional] [readonly] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source. | [optional] 
**VmwareParams** | Pointer to [**VmwareSourceRegistrationParams**](VmwareSourceRegistrationParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalSourceRegistrationParams**](PhysicalSourceRegistrationParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasRegistrationParams**](GenericNasRegistrationParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonRegistrationParams**](IsilonRegistrationParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappRegistrationParams**](NetappRegistrationParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraSourceRegistrationParams**](CassandraSourceRegistrationParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDBSourceRegistrationParams**](MongoDBSourceRegistrationParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseSourceRegistrationParams**](CouchbaseSourceRegistrationParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsSourceRegistrationParams**](HdfsSourceRegistrationParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseSourceRegistrationParams**](HbaseSourceRegistrationParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveSourceRegistrationParams**](HiveSourceRegistrationParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsSourceRegistrationParams**](AwsSourceRegistrationParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365SourceRegistrationParams**](Office365SourceRegistrationParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVSourceRegistrationParams**](HyperVSourceRegistrationParams.md) |  | [optional] 

## Methods

### NewMcmSourceRegistration

`func NewMcmSourceRegistration() *McmSourceRegistration`

NewMcmSourceRegistration instantiates a new McmSourceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourceRegistrationWithDefaults

`func NewMcmSourceRegistrationWithDefaults() *McmSourceRegistration`

NewMcmSourceRegistrationWithDefaults instantiates a new McmSourceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *McmSourceRegistration) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmSourceRegistration) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmSourceRegistration) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmSourceRegistration) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmSourceRegistration) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmSourceRegistration) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmSourceRegistration) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmSourceRegistration) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmSourceRegistration) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmSourceRegistration) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmSourceRegistration) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmSourceRegistration) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *McmSourceRegistration) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmSourceRegistration) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmSourceRegistration) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmSourceRegistration) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmSourceRegistration) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmSourceRegistration) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetId

`func (o *McmSourceRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmSourceRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmSourceRegistration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McmSourceRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmSourceRegistration) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmSourceRegistration) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetConnectionId

`func (o *McmSourceRegistration) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *McmSourceRegistration) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *McmSourceRegistration) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *McmSourceRegistration) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *McmSourceRegistration) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *McmSourceRegistration) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetSourceId

`func (o *McmSourceRegistration) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *McmSourceRegistration) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *McmSourceRegistration) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *McmSourceRegistration) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *McmSourceRegistration) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *McmSourceRegistration) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetEnvironment

`func (o *McmSourceRegistration) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmSourceRegistration) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmSourceRegistration) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McmSourceRegistration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *McmSourceRegistration) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmSourceRegistration) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetVmwareParams

`func (o *McmSourceRegistration) GetVmwareParams() VmwareSourceRegistrationParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *McmSourceRegistration) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *McmSourceRegistration) SetVmwareParams(v VmwareSourceRegistrationParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *McmSourceRegistration) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *McmSourceRegistration) GetPhysicalParams() PhysicalSourceRegistrationParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *McmSourceRegistration) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *McmSourceRegistration) SetPhysicalParams(v PhysicalSourceRegistrationParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *McmSourceRegistration) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *McmSourceRegistration) GetGenericNasParams() GenericNasRegistrationParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *McmSourceRegistration) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *McmSourceRegistration) SetGenericNasParams(v GenericNasRegistrationParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *McmSourceRegistration) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *McmSourceRegistration) GetIsilonParams() IsilonRegistrationParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *McmSourceRegistration) GetIsilonParamsOk() (*IsilonRegistrationParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *McmSourceRegistration) SetIsilonParams(v IsilonRegistrationParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *McmSourceRegistration) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *McmSourceRegistration) GetNetappParams() NetappRegistrationParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *McmSourceRegistration) GetNetappParamsOk() (*NetappRegistrationParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *McmSourceRegistration) SetNetappParams(v NetappRegistrationParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *McmSourceRegistration) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *McmSourceRegistration) GetCassandraParams() CassandraSourceRegistrationParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *McmSourceRegistration) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *McmSourceRegistration) SetCassandraParams(v CassandraSourceRegistrationParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *McmSourceRegistration) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *McmSourceRegistration) GetMongodbParams() MongoDBSourceRegistrationParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *McmSourceRegistration) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *McmSourceRegistration) SetMongodbParams(v MongoDBSourceRegistrationParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *McmSourceRegistration) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *McmSourceRegistration) GetCouchbaseParams() CouchbaseSourceRegistrationParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *McmSourceRegistration) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *McmSourceRegistration) SetCouchbaseParams(v CouchbaseSourceRegistrationParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *McmSourceRegistration) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *McmSourceRegistration) GetHdfsParams() HdfsSourceRegistrationParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *McmSourceRegistration) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *McmSourceRegistration) SetHdfsParams(v HdfsSourceRegistrationParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *McmSourceRegistration) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *McmSourceRegistration) GetHbaseParams() HbaseSourceRegistrationParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *McmSourceRegistration) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *McmSourceRegistration) SetHbaseParams(v HbaseSourceRegistrationParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *McmSourceRegistration) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *McmSourceRegistration) GetHiveParams() HiveSourceRegistrationParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *McmSourceRegistration) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *McmSourceRegistration) SetHiveParams(v HiveSourceRegistrationParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *McmSourceRegistration) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *McmSourceRegistration) GetAwsParams() AwsSourceRegistrationParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *McmSourceRegistration) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *McmSourceRegistration) SetAwsParams(v AwsSourceRegistrationParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *McmSourceRegistration) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *McmSourceRegistration) GetOffice365Params() Office365SourceRegistrationParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *McmSourceRegistration) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *McmSourceRegistration) SetOffice365Params(v Office365SourceRegistrationParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *McmSourceRegistration) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetHypervParams

`func (o *McmSourceRegistration) GetHypervParams() HyperVSourceRegistrationParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *McmSourceRegistration) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *McmSourceRegistration) SetHypervParams(v HyperVSourceRegistrationParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *McmSourceRegistration) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


