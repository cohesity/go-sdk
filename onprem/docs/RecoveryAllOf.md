# RecoveryAllOf

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
**Office365Params** | Pointer to [**RecoverO365Params**](RecoverO365Params.md) |  | [optional] 
**KubernetesParams** | Pointer to [**RecoverKubernetesParams**](RecoverKubernetesParams.md) |  | [optional] 
**OracleParams** | Pointer to [**RecoverOracleParams**](RecoverOracleParams.md) |  | [optional] 
**ViewParams** | Pointer to [**RecoverViewParams**](RecoverViewParams.md) |  | [optional] 

## Methods

### NewRecoveryAllOf

`func NewRecoveryAllOf() *RecoveryAllOf`

NewRecoveryAllOf instantiates a new RecoveryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryAllOfWithDefaults

`func NewRecoveryAllOfWithDefaults() *RecoveryAllOf`

NewRecoveryAllOfWithDefaults instantiates a new RecoveryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *RecoveryAllOf) GetVmwareParams() RecoverVmwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *RecoveryAllOf) GetVmwareParamsOk() (*RecoverVmwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *RecoveryAllOf) SetVmwareParams(v RecoverVmwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *RecoveryAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *RecoveryAllOf) GetAwsParams() RecoverAwsParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *RecoveryAllOf) GetAwsParamsOk() (*RecoverAwsParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *RecoveryAllOf) SetAwsParams(v RecoverAwsParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *RecoveryAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *RecoveryAllOf) GetGcpParams() RecoverGcpParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *RecoveryAllOf) GetGcpParamsOk() (*RecoverGcpParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *RecoveryAllOf) SetGcpParams(v RecoverGcpParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *RecoveryAllOf) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *RecoveryAllOf) GetAzureParams() RecoverAzureParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *RecoveryAllOf) GetAzureParamsOk() (*RecoverAzureParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *RecoveryAllOf) SetAzureParams(v RecoverAzureParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *RecoveryAllOf) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *RecoveryAllOf) GetKvmParams() RecoverKvmParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *RecoveryAllOf) GetKvmParamsOk() (*RecoverKvmParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *RecoveryAllOf) SetKvmParams(v RecoverKvmParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *RecoveryAllOf) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### GetAcropolisParams

`func (o *RecoveryAllOf) GetAcropolisParams() RecoverAcropolisParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *RecoveryAllOf) GetAcropolisParamsOk() (*RecoverAcropolisParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *RecoveryAllOf) SetAcropolisParams(v RecoverAcropolisParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *RecoveryAllOf) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *RecoveryAllOf) GetMssqlParams() RecoverSqlParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *RecoveryAllOf) GetMssqlParamsOk() (*RecoverSqlParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *RecoveryAllOf) SetMssqlParams(v RecoverSqlParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *RecoveryAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *RecoveryAllOf) GetNetappParams() RecoverNetappParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *RecoveryAllOf) GetNetappParamsOk() (*RecoverNetappParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *RecoveryAllOf) SetNetappParams(v RecoverNetappParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *RecoveryAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *RecoveryAllOf) GetGenericNasParams() RecoverGenericNasParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *RecoveryAllOf) GetGenericNasParamsOk() (*RecoverGenericNasParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *RecoveryAllOf) SetGenericNasParams(v RecoverGenericNasParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *RecoveryAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *RecoveryAllOf) GetIsilonParams() RecoverIsilonParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *RecoveryAllOf) GetIsilonParamsOk() (*RecoverIsilonParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *RecoveryAllOf) SetIsilonParams(v RecoverIsilonParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *RecoveryAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *RecoveryAllOf) GetFlashbladeParams() RecoverFlashbladeParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *RecoveryAllOf) GetFlashbladeParamsOk() (*RecoverFlashbladeParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *RecoveryAllOf) SetFlashbladeParams(v RecoverFlashbladeParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *RecoveryAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *RecoveryAllOf) GetElastifileParams() RecoverElastifileParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *RecoveryAllOf) GetElastifileParamsOk() (*RecoverElastifileParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *RecoveryAllOf) SetElastifileParams(v RecoverElastifileParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *RecoveryAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *RecoveryAllOf) GetGpfsParams() RecoverGpfsParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *RecoveryAllOf) GetGpfsParamsOk() (*RecoverGpfsParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *RecoveryAllOf) SetGpfsParams(v RecoverGpfsParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *RecoveryAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *RecoveryAllOf) GetPhysicalParams() RecoverPhysicalParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *RecoveryAllOf) GetPhysicalParamsOk() (*RecoverPhysicalParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *RecoveryAllOf) SetPhysicalParams(v RecoverPhysicalParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *RecoveryAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *RecoveryAllOf) GetHypervParams() RecoverHyperVParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *RecoveryAllOf) GetHypervParamsOk() (*RecoverHyperVParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *RecoveryAllOf) SetHypervParams(v RecoverHyperVParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *RecoveryAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *RecoveryAllOf) GetExchangeParams() RecoverExchangeParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *RecoveryAllOf) GetExchangeParamsOk() (*RecoverExchangeParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *RecoveryAllOf) SetExchangeParams(v RecoverExchangeParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *RecoveryAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *RecoveryAllOf) GetCassandraParams() CassandraParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *RecoveryAllOf) GetCassandraParamsOk() (*CassandraParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *RecoveryAllOf) SetCassandraParams(v CassandraParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *RecoveryAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *RecoveryAllOf) GetUdaParams() UdaParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *RecoveryAllOf) GetUdaParamsOk() (*UdaParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *RecoveryAllOf) SetUdaParams(v UdaParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *RecoveryAllOf) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *RecoveryAllOf) GetCouchbaseParams() CouchbaseParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *RecoveryAllOf) GetCouchbaseParamsOk() (*CouchbaseParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *RecoveryAllOf) SetCouchbaseParams(v CouchbaseParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *RecoveryAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *RecoveryAllOf) GetHbaseParams() HbaseParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *RecoveryAllOf) GetHbaseParamsOk() (*HbaseParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *RecoveryAllOf) SetHbaseParams(v HbaseParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *RecoveryAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *RecoveryAllOf) GetHdfsParams() HdfsParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *RecoveryAllOf) GetHdfsParamsOk() (*HdfsParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *RecoveryAllOf) SetHdfsParams(v HdfsParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *RecoveryAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *RecoveryAllOf) GetHiveParams() HiveParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *RecoveryAllOf) GetHiveParamsOk() (*HiveParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *RecoveryAllOf) SetHiveParams(v HiveParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *RecoveryAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *RecoveryAllOf) GetMongodbParams() MongodbParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *RecoveryAllOf) GetMongodbParamsOk() (*MongodbParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *RecoveryAllOf) SetMongodbParams(v MongodbParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *RecoveryAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetPureParams

`func (o *RecoveryAllOf) GetPureParams() RecoverPureParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *RecoveryAllOf) GetPureParamsOk() (*RecoverPureParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *RecoveryAllOf) SetPureParams(v RecoverPureParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *RecoveryAllOf) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *RecoveryAllOf) GetOffice365Params() RecoverO365Params`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *RecoveryAllOf) GetOffice365ParamsOk() (*RecoverO365Params, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *RecoveryAllOf) SetOffice365Params(v RecoverO365Params)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *RecoveryAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *RecoveryAllOf) GetKubernetesParams() RecoverKubernetesParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *RecoveryAllOf) GetKubernetesParamsOk() (*RecoverKubernetesParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *RecoveryAllOf) SetKubernetesParams(v RecoverKubernetesParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *RecoveryAllOf) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *RecoveryAllOf) GetOracleParams() RecoverOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *RecoveryAllOf) GetOracleParamsOk() (*RecoverOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *RecoveryAllOf) SetOracleParams(v RecoverOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *RecoveryAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetViewParams

`func (o *RecoveryAllOf) GetViewParams() RecoverViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *RecoveryAllOf) GetViewParamsOk() (*RecoverViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *RecoveryAllOf) SetViewParams(v RecoverViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *RecoveryAllOf) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


