# SourceRegistrationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSourceRegistrationAllOf

`func NewSourceRegistrationAllOf() *SourceRegistrationAllOf`

NewSourceRegistrationAllOf instantiates a new SourceRegistrationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRegistrationAllOfWithDefaults

`func NewSourceRegistrationAllOfWithDefaults() *SourceRegistrationAllOf`

NewSourceRegistrationAllOfWithDefaults instantiates a new SourceRegistrationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *SourceRegistrationAllOf) GetVmwareParams() VmwareSourceRegistrationParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *SourceRegistrationAllOf) GetVmwareParamsOk() (*VmwareSourceRegistrationParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *SourceRegistrationAllOf) SetVmwareParams(v VmwareSourceRegistrationParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *SourceRegistrationAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *SourceRegistrationAllOf) GetPhysicalParams() PhysicalSourceRegistrationParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *SourceRegistrationAllOf) GetPhysicalParamsOk() (*PhysicalSourceRegistrationParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *SourceRegistrationAllOf) SetPhysicalParams(v PhysicalSourceRegistrationParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *SourceRegistrationAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *SourceRegistrationAllOf) GetGenericNasParams() GenericNasRegistrationParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *SourceRegistrationAllOf) GetGenericNasParamsOk() (*GenericNasRegistrationParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *SourceRegistrationAllOf) SetGenericNasParams(v GenericNasRegistrationParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *SourceRegistrationAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *SourceRegistrationAllOf) GetIsilonParams() IsilonRegistrationParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *SourceRegistrationAllOf) GetIsilonParamsOk() (*IsilonRegistrationParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *SourceRegistrationAllOf) SetIsilonParams(v IsilonRegistrationParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *SourceRegistrationAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *SourceRegistrationAllOf) GetNetappParams() NetappRegistrationParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *SourceRegistrationAllOf) GetNetappParamsOk() (*NetappRegistrationParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *SourceRegistrationAllOf) SetNetappParams(v NetappRegistrationParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *SourceRegistrationAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *SourceRegistrationAllOf) GetElastifileParams() ElastifileRegistrationParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *SourceRegistrationAllOf) GetElastifileParamsOk() (*ElastifileRegistrationParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *SourceRegistrationAllOf) SetElastifileParams(v ElastifileRegistrationParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *SourceRegistrationAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *SourceRegistrationAllOf) GetFlashbladeParams() FlashbladeRegistrationParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *SourceRegistrationAllOf) GetFlashbladeParamsOk() (*FlashbladeRegistrationParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *SourceRegistrationAllOf) SetFlashbladeParams(v FlashbladeRegistrationParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *SourceRegistrationAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *SourceRegistrationAllOf) GetGpfsParams() GpfsRegistrationParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *SourceRegistrationAllOf) GetGpfsParamsOk() (*GpfsRegistrationParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *SourceRegistrationAllOf) SetGpfsParams(v GpfsRegistrationParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *SourceRegistrationAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *SourceRegistrationAllOf) GetCassandraParams() CassandraSourceRegistrationParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SourceRegistrationAllOf) GetCassandraParamsOk() (*CassandraSourceRegistrationParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SourceRegistrationAllOf) SetCassandraParams(v CassandraSourceRegistrationParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SourceRegistrationAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *SourceRegistrationAllOf) GetMongodbParams() MongoDBSourceRegistrationParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *SourceRegistrationAllOf) GetMongodbParamsOk() (*MongoDBSourceRegistrationParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *SourceRegistrationAllOf) SetMongodbParams(v MongoDBSourceRegistrationParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *SourceRegistrationAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *SourceRegistrationAllOf) GetCouchbaseParams() CouchbaseSourceRegistrationParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *SourceRegistrationAllOf) GetCouchbaseParamsOk() (*CouchbaseSourceRegistrationParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *SourceRegistrationAllOf) SetCouchbaseParams(v CouchbaseSourceRegistrationParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *SourceRegistrationAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *SourceRegistrationAllOf) GetHdfsParams() HdfsSourceRegistrationParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *SourceRegistrationAllOf) GetHdfsParamsOk() (*HdfsSourceRegistrationParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *SourceRegistrationAllOf) SetHdfsParams(v HdfsSourceRegistrationParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *SourceRegistrationAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *SourceRegistrationAllOf) GetHbaseParams() HbaseSourceRegistrationParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *SourceRegistrationAllOf) GetHbaseParamsOk() (*HbaseSourceRegistrationParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *SourceRegistrationAllOf) SetHbaseParams(v HbaseSourceRegistrationParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *SourceRegistrationAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *SourceRegistrationAllOf) GetHiveParams() HiveSourceRegistrationParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *SourceRegistrationAllOf) GetHiveParamsOk() (*HiveSourceRegistrationParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *SourceRegistrationAllOf) SetHiveParams(v HiveSourceRegistrationParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *SourceRegistrationAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *SourceRegistrationAllOf) GetUdaParams() UdaSourceRegistrationParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *SourceRegistrationAllOf) GetUdaParamsOk() (*UdaSourceRegistrationParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *SourceRegistrationAllOf) SetUdaParams(v UdaSourceRegistrationParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *SourceRegistrationAllOf) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *SourceRegistrationAllOf) GetOffice365Params() Office365SourceRegistrationParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *SourceRegistrationAllOf) GetOffice365ParamsOk() (*Office365SourceRegistrationParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *SourceRegistrationAllOf) SetOffice365Params(v Office365SourceRegistrationParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *SourceRegistrationAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *SourceRegistrationAllOf) GetAwsParams() AwsSourceRegistrationParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *SourceRegistrationAllOf) GetAwsParamsOk() (*AwsSourceRegistrationParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *SourceRegistrationAllOf) SetAwsParams(v AwsSourceRegistrationParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *SourceRegistrationAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *SourceRegistrationAllOf) GetHypervParams() HyperVSourceRegistrationParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *SourceRegistrationAllOf) GetHypervParamsOk() (*HyperVSourceRegistrationParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *SourceRegistrationAllOf) SetHypervParams(v HyperVSourceRegistrationParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *SourceRegistrationAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


