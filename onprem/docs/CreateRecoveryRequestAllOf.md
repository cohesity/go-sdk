# CreateRecoveryRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmwareParams** | Pointer to [**RecoverVmwareParams**](RecoverVmwareParams.md) |  | [optional] 
**AwsParams** | Pointer to [**RecoverAwsParams**](RecoverAwsParams.md) |  | [optional] 
**GcpParams** | Pointer to [**RecoverGcpParams**](RecoverGcpParams.md) |  | [optional] 
**AzureParams** | Pointer to [**RecoverAzureParams**](RecoverAzureParams.md) |  | [optional] 
**KvmParams** | Pointer to [**RecoverKvmParams**](RecoverKvmParams.md) |  | [optional] 
**AcropolisParams** | Pointer to [**RecoverAcropolisParams**](RecoverAcropolisParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**RecoverSqlParams**](RecoverSqlParams.md) |  | [optional] 
**NetappParams** | Pointer to [**RecoverNetappParams**](RecoverNetappParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**RecoverGenericNasParams**](RecoverGenericNasParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**RecoverIsilonParams**](RecoverIsilonParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**RecoverFlashbladeParams**](RecoverFlashbladeParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**RecoverElastifileParams**](RecoverElastifileParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**RecoverGpfsParams**](RecoverGpfsParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**RecoverPhysicalParams**](RecoverPhysicalParams.md) |  | [optional] 
**HypervParams** | Pointer to [**RecoverHyperVParams**](RecoverHyperVParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**RecoverExchangeParams**](RecoverExchangeParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraParams**](CassandraParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaParams**](UdaParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseParams**](CouchbaseParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseParams**](HbaseParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsParams**](HdfsParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveParams**](HiveParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongodbParams**](MongodbParams.md) |  | [optional] 
**PureParams** | Pointer to [**RecoverPureParams**](RecoverPureParams.md) |  | [optional] 
**KubernetesParams** | Pointer to [**RecoverKubernetesParams**](RecoverKubernetesParams.md) |  | [optional] 
**Office365Params** | Pointer to [**RecoverO365Params**](RecoverO365Params.md) |  | [optional] 
**OracleParams** | Pointer to [**RecoverOracleParams**](RecoverOracleParams.md) |  | [optional] 

## Methods

### NewCreateRecoveryRequestAllOf

`func NewCreateRecoveryRequestAllOf() *CreateRecoveryRequestAllOf`

NewCreateRecoveryRequestAllOf instantiates a new CreateRecoveryRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecoveryRequestAllOfWithDefaults

`func NewCreateRecoveryRequestAllOfWithDefaults() *CreateRecoveryRequestAllOf`

NewCreateRecoveryRequestAllOfWithDefaults instantiates a new CreateRecoveryRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *CreateRecoveryRequestAllOf) GetVmwareParams() RecoverVmwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *CreateRecoveryRequestAllOf) GetVmwareParamsOk() (*RecoverVmwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *CreateRecoveryRequestAllOf) SetVmwareParams(v RecoverVmwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *CreateRecoveryRequestAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *CreateRecoveryRequestAllOf) GetAwsParams() RecoverAwsParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CreateRecoveryRequestAllOf) GetAwsParamsOk() (*RecoverAwsParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CreateRecoveryRequestAllOf) SetAwsParams(v RecoverAwsParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CreateRecoveryRequestAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *CreateRecoveryRequestAllOf) GetGcpParams() RecoverGcpParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *CreateRecoveryRequestAllOf) GetGcpParamsOk() (*RecoverGcpParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *CreateRecoveryRequestAllOf) SetGcpParams(v RecoverGcpParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *CreateRecoveryRequestAllOf) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CreateRecoveryRequestAllOf) GetAzureParams() RecoverAzureParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CreateRecoveryRequestAllOf) GetAzureParamsOk() (*RecoverAzureParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CreateRecoveryRequestAllOf) SetAzureParams(v RecoverAzureParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CreateRecoveryRequestAllOf) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *CreateRecoveryRequestAllOf) GetKvmParams() RecoverKvmParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *CreateRecoveryRequestAllOf) GetKvmParamsOk() (*RecoverKvmParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *CreateRecoveryRequestAllOf) SetKvmParams(v RecoverKvmParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *CreateRecoveryRequestAllOf) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### GetAcropolisParams

`func (o *CreateRecoveryRequestAllOf) GetAcropolisParams() RecoverAcropolisParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *CreateRecoveryRequestAllOf) GetAcropolisParamsOk() (*RecoverAcropolisParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *CreateRecoveryRequestAllOf) SetAcropolisParams(v RecoverAcropolisParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *CreateRecoveryRequestAllOf) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *CreateRecoveryRequestAllOf) GetMssqlParams() RecoverSqlParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *CreateRecoveryRequestAllOf) GetMssqlParamsOk() (*RecoverSqlParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *CreateRecoveryRequestAllOf) SetMssqlParams(v RecoverSqlParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *CreateRecoveryRequestAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *CreateRecoveryRequestAllOf) GetNetappParams() RecoverNetappParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *CreateRecoveryRequestAllOf) GetNetappParamsOk() (*RecoverNetappParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *CreateRecoveryRequestAllOf) SetNetappParams(v RecoverNetappParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *CreateRecoveryRequestAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *CreateRecoveryRequestAllOf) GetGenericNasParams() RecoverGenericNasParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *CreateRecoveryRequestAllOf) GetGenericNasParamsOk() (*RecoverGenericNasParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *CreateRecoveryRequestAllOf) SetGenericNasParams(v RecoverGenericNasParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *CreateRecoveryRequestAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *CreateRecoveryRequestAllOf) GetIsilonParams() RecoverIsilonParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *CreateRecoveryRequestAllOf) GetIsilonParamsOk() (*RecoverIsilonParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *CreateRecoveryRequestAllOf) SetIsilonParams(v RecoverIsilonParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *CreateRecoveryRequestAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *CreateRecoveryRequestAllOf) GetFlashbladeParams() RecoverFlashbladeParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *CreateRecoveryRequestAllOf) GetFlashbladeParamsOk() (*RecoverFlashbladeParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *CreateRecoveryRequestAllOf) SetFlashbladeParams(v RecoverFlashbladeParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *CreateRecoveryRequestAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *CreateRecoveryRequestAllOf) GetElastifileParams() RecoverElastifileParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *CreateRecoveryRequestAllOf) GetElastifileParamsOk() (*RecoverElastifileParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *CreateRecoveryRequestAllOf) SetElastifileParams(v RecoverElastifileParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *CreateRecoveryRequestAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *CreateRecoveryRequestAllOf) GetGpfsParams() RecoverGpfsParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *CreateRecoveryRequestAllOf) GetGpfsParamsOk() (*RecoverGpfsParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *CreateRecoveryRequestAllOf) SetGpfsParams(v RecoverGpfsParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *CreateRecoveryRequestAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *CreateRecoveryRequestAllOf) GetPhysicalParams() RecoverPhysicalParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *CreateRecoveryRequestAllOf) GetPhysicalParamsOk() (*RecoverPhysicalParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *CreateRecoveryRequestAllOf) SetPhysicalParams(v RecoverPhysicalParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *CreateRecoveryRequestAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *CreateRecoveryRequestAllOf) GetHypervParams() RecoverHyperVParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *CreateRecoveryRequestAllOf) GetHypervParamsOk() (*RecoverHyperVParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *CreateRecoveryRequestAllOf) SetHypervParams(v RecoverHyperVParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *CreateRecoveryRequestAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *CreateRecoveryRequestAllOf) GetExchangeParams() RecoverExchangeParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *CreateRecoveryRequestAllOf) GetExchangeParamsOk() (*RecoverExchangeParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *CreateRecoveryRequestAllOf) SetExchangeParams(v RecoverExchangeParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *CreateRecoveryRequestAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *CreateRecoveryRequestAllOf) GetCassandraParams() CassandraParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *CreateRecoveryRequestAllOf) GetCassandraParamsOk() (*CassandraParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *CreateRecoveryRequestAllOf) SetCassandraParams(v CassandraParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *CreateRecoveryRequestAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *CreateRecoveryRequestAllOf) GetUdaParams() UdaParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *CreateRecoveryRequestAllOf) GetUdaParamsOk() (*UdaParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *CreateRecoveryRequestAllOf) SetUdaParams(v UdaParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *CreateRecoveryRequestAllOf) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *CreateRecoveryRequestAllOf) GetCouchbaseParams() CouchbaseParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *CreateRecoveryRequestAllOf) GetCouchbaseParamsOk() (*CouchbaseParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *CreateRecoveryRequestAllOf) SetCouchbaseParams(v CouchbaseParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *CreateRecoveryRequestAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *CreateRecoveryRequestAllOf) GetHbaseParams() HbaseParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *CreateRecoveryRequestAllOf) GetHbaseParamsOk() (*HbaseParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *CreateRecoveryRequestAllOf) SetHbaseParams(v HbaseParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *CreateRecoveryRequestAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *CreateRecoveryRequestAllOf) GetHdfsParams() HdfsParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *CreateRecoveryRequestAllOf) GetHdfsParamsOk() (*HdfsParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *CreateRecoveryRequestAllOf) SetHdfsParams(v HdfsParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *CreateRecoveryRequestAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *CreateRecoveryRequestAllOf) GetHiveParams() HiveParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *CreateRecoveryRequestAllOf) GetHiveParamsOk() (*HiveParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *CreateRecoveryRequestAllOf) SetHiveParams(v HiveParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *CreateRecoveryRequestAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *CreateRecoveryRequestAllOf) GetMongodbParams() MongodbParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *CreateRecoveryRequestAllOf) GetMongodbParamsOk() (*MongodbParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *CreateRecoveryRequestAllOf) SetMongodbParams(v MongodbParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *CreateRecoveryRequestAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetPureParams

`func (o *CreateRecoveryRequestAllOf) GetPureParams() RecoverPureParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *CreateRecoveryRequestAllOf) GetPureParamsOk() (*RecoverPureParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *CreateRecoveryRequestAllOf) SetPureParams(v RecoverPureParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *CreateRecoveryRequestAllOf) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *CreateRecoveryRequestAllOf) GetKubernetesParams() RecoverKubernetesParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *CreateRecoveryRequestAllOf) GetKubernetesParamsOk() (*RecoverKubernetesParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *CreateRecoveryRequestAllOf) SetKubernetesParams(v RecoverKubernetesParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *CreateRecoveryRequestAllOf) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *CreateRecoveryRequestAllOf) GetOffice365Params() RecoverO365Params`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *CreateRecoveryRequestAllOf) GetOffice365ParamsOk() (*RecoverO365Params, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *CreateRecoveryRequestAllOf) SetOffice365Params(v RecoverO365Params)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *CreateRecoveryRequestAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *CreateRecoveryRequestAllOf) GetOracleParams() RecoverOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *CreateRecoveryRequestAllOf) GetOracleParamsOk() (*RecoverOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *CreateRecoveryRequestAllOf) SetOracleParams(v RecoverOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *CreateRecoveryRequestAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


