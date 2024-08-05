# CreateRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the Recovery. | 
**SnapshotEnvironment** | **string** | Specifies the type of environment of snapshots for which the Recovery has to be performed. | 
**AcropolisParams** | Pointer to [**RecoverAcropolisParams**](RecoverAcropolisParams.md) |  | [optional] 
**AwsParams** | Pointer to [**RecoverAwsParams**](RecoverAwsParams.md) |  | [optional] 
**AzureParams** | Pointer to [**RecoverAzureParams**](RecoverAzureParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraParams**](CassandraParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseParams**](CouchbaseParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**RecoverElastifileParams**](RecoverElastifileParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**RecoverExchangeParams**](RecoverExchangeParams.md) |  | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterParams**](ExperimentalAdapterParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**RecoverFlashbladeParams**](RecoverFlashbladeParams.md) |  | [optional] 
**GcpParams** | Pointer to [**RecoverGcpParams**](RecoverGcpParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**RecoverGenericNasParams**](RecoverGenericNasParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**RecoverGpfsParams**](RecoverGpfsParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseParams**](HbaseParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsParams**](HdfsParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveParams**](HiveParams.md) |  | [optional] 
**HypervParams** | Pointer to [**RecoverHyperVParams**](RecoverHyperVParams.md) |  | [optional] 
**IbmFlashSystemParams** | Pointer to [**RecoverPureParams**](RecoverPureParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**RecoverIsilonParams**](RecoverIsilonParams.md) |  | [optional] 
**KubernetesParams** | Pointer to [**RecoverKubernetesParams**](RecoverKubernetesParams.md) |  | [optional] 
**KvmParams** | Pointer to [**RecoverKvmParams**](RecoverKvmParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongodbParams**](MongodbParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**RecoverSqlParams**](RecoverSqlParams.md) |  | [optional] 
**NetappParams** | Pointer to [**RecoverNetappParams**](RecoverNetappParams.md) |  | [optional] 
**Office365Params** | Pointer to [**RecoverO365Params**](RecoverO365Params.md) |  | [optional] 
**OracleParams** | Pointer to [**RecoverOracleParams**](RecoverOracleParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**RecoverPhysicalParams**](RecoverPhysicalParams.md) |  | [optional] 
**PureParams** | Pointer to [**RecoverPureParams**](RecoverPureParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaParams**](SapHanaParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**RecoverSalesforceParams**](RecoverSalesforceParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaParams**](UdaParams.md) |  | [optional] 
**ViewParams** | Pointer to [**RecoverViewParams**](RecoverViewParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**RecoverVmwareParams**](RecoverVmwareParams.md) |  | [optional] 

## Methods

### NewCreateRecoveryRequest

`func NewCreateRecoveryRequest(name NullableString, snapshotEnvironment string, ) *CreateRecoveryRequest`

NewCreateRecoveryRequest instantiates a new CreateRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecoveryRequestWithDefaults

`func NewCreateRecoveryRequestWithDefaults() *CreateRecoveryRequest`

NewCreateRecoveryRequestWithDefaults instantiates a new CreateRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRecoveryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRecoveryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRecoveryRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateRecoveryRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateRecoveryRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSnapshotEnvironment

`func (o *CreateRecoveryRequest) GetSnapshotEnvironment() string`

GetSnapshotEnvironment returns the SnapshotEnvironment field if non-nil, zero value otherwise.

### GetSnapshotEnvironmentOk

`func (o *CreateRecoveryRequest) GetSnapshotEnvironmentOk() (*string, bool)`

GetSnapshotEnvironmentOk returns a tuple with the SnapshotEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnvironment

`func (o *CreateRecoveryRequest) SetSnapshotEnvironment(v string)`

SetSnapshotEnvironment sets SnapshotEnvironment field to given value.


### GetAcropolisParams

`func (o *CreateRecoveryRequest) GetAcropolisParams() RecoverAcropolisParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *CreateRecoveryRequest) GetAcropolisParamsOk() (*RecoverAcropolisParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *CreateRecoveryRequest) SetAcropolisParams(v RecoverAcropolisParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *CreateRecoveryRequest) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *CreateRecoveryRequest) GetAwsParams() RecoverAwsParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CreateRecoveryRequest) GetAwsParamsOk() (*RecoverAwsParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CreateRecoveryRequest) SetAwsParams(v RecoverAwsParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CreateRecoveryRequest) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CreateRecoveryRequest) GetAzureParams() RecoverAzureParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CreateRecoveryRequest) GetAzureParamsOk() (*RecoverAzureParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CreateRecoveryRequest) SetAzureParams(v RecoverAzureParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CreateRecoveryRequest) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *CreateRecoveryRequest) GetCassandraParams() CassandraParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *CreateRecoveryRequest) GetCassandraParamsOk() (*CassandraParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *CreateRecoveryRequest) SetCassandraParams(v CassandraParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *CreateRecoveryRequest) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *CreateRecoveryRequest) GetCouchbaseParams() CouchbaseParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *CreateRecoveryRequest) GetCouchbaseParamsOk() (*CouchbaseParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *CreateRecoveryRequest) SetCouchbaseParams(v CouchbaseParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *CreateRecoveryRequest) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *CreateRecoveryRequest) GetElastifileParams() RecoverElastifileParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *CreateRecoveryRequest) GetElastifileParamsOk() (*RecoverElastifileParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *CreateRecoveryRequest) SetElastifileParams(v RecoverElastifileParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *CreateRecoveryRequest) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *CreateRecoveryRequest) GetExchangeParams() RecoverExchangeParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *CreateRecoveryRequest) GetExchangeParamsOk() (*RecoverExchangeParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *CreateRecoveryRequest) SetExchangeParams(v RecoverExchangeParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *CreateRecoveryRequest) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *CreateRecoveryRequest) GetExperimentalAdapterParams() ExperimentalAdapterParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *CreateRecoveryRequest) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *CreateRecoveryRequest) SetExperimentalAdapterParams(v ExperimentalAdapterParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *CreateRecoveryRequest) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *CreateRecoveryRequest) GetFlashbladeParams() RecoverFlashbladeParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *CreateRecoveryRequest) GetFlashbladeParamsOk() (*RecoverFlashbladeParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *CreateRecoveryRequest) SetFlashbladeParams(v RecoverFlashbladeParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *CreateRecoveryRequest) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *CreateRecoveryRequest) GetGcpParams() RecoverGcpParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *CreateRecoveryRequest) GetGcpParamsOk() (*RecoverGcpParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *CreateRecoveryRequest) SetGcpParams(v RecoverGcpParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *CreateRecoveryRequest) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *CreateRecoveryRequest) GetGenericNasParams() RecoverGenericNasParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *CreateRecoveryRequest) GetGenericNasParamsOk() (*RecoverGenericNasParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *CreateRecoveryRequest) SetGenericNasParams(v RecoverGenericNasParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *CreateRecoveryRequest) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *CreateRecoveryRequest) GetGpfsParams() RecoverGpfsParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *CreateRecoveryRequest) GetGpfsParamsOk() (*RecoverGpfsParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *CreateRecoveryRequest) SetGpfsParams(v RecoverGpfsParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *CreateRecoveryRequest) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *CreateRecoveryRequest) GetHbaseParams() HbaseParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *CreateRecoveryRequest) GetHbaseParamsOk() (*HbaseParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *CreateRecoveryRequest) SetHbaseParams(v HbaseParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *CreateRecoveryRequest) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *CreateRecoveryRequest) GetHdfsParams() HdfsParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *CreateRecoveryRequest) GetHdfsParamsOk() (*HdfsParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *CreateRecoveryRequest) SetHdfsParams(v HdfsParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *CreateRecoveryRequest) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *CreateRecoveryRequest) GetHiveParams() HiveParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *CreateRecoveryRequest) GetHiveParamsOk() (*HiveParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *CreateRecoveryRequest) SetHiveParams(v HiveParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *CreateRecoveryRequest) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *CreateRecoveryRequest) GetHypervParams() RecoverHyperVParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *CreateRecoveryRequest) GetHypervParamsOk() (*RecoverHyperVParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *CreateRecoveryRequest) SetHypervParams(v RecoverHyperVParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *CreateRecoveryRequest) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIbmFlashSystemParams

`func (o *CreateRecoveryRequest) GetIbmFlashSystemParams() RecoverPureParams`

GetIbmFlashSystemParams returns the IbmFlashSystemParams field if non-nil, zero value otherwise.

### GetIbmFlashSystemParamsOk

`func (o *CreateRecoveryRequest) GetIbmFlashSystemParamsOk() (*RecoverPureParams, bool)`

GetIbmFlashSystemParamsOk returns a tuple with the IbmFlashSystemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmFlashSystemParams

`func (o *CreateRecoveryRequest) SetIbmFlashSystemParams(v RecoverPureParams)`

SetIbmFlashSystemParams sets IbmFlashSystemParams field to given value.

### HasIbmFlashSystemParams

`func (o *CreateRecoveryRequest) HasIbmFlashSystemParams() bool`

HasIbmFlashSystemParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *CreateRecoveryRequest) GetIsilonParams() RecoverIsilonParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *CreateRecoveryRequest) GetIsilonParamsOk() (*RecoverIsilonParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *CreateRecoveryRequest) SetIsilonParams(v RecoverIsilonParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *CreateRecoveryRequest) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *CreateRecoveryRequest) GetKubernetesParams() RecoverKubernetesParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *CreateRecoveryRequest) GetKubernetesParamsOk() (*RecoverKubernetesParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *CreateRecoveryRequest) SetKubernetesParams(v RecoverKubernetesParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *CreateRecoveryRequest) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *CreateRecoveryRequest) GetKvmParams() RecoverKvmParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *CreateRecoveryRequest) GetKvmParamsOk() (*RecoverKvmParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *CreateRecoveryRequest) SetKvmParams(v RecoverKvmParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *CreateRecoveryRequest) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *CreateRecoveryRequest) GetMongodbParams() MongodbParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *CreateRecoveryRequest) GetMongodbParamsOk() (*MongodbParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *CreateRecoveryRequest) SetMongodbParams(v MongodbParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *CreateRecoveryRequest) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *CreateRecoveryRequest) GetMssqlParams() RecoverSqlParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *CreateRecoveryRequest) GetMssqlParamsOk() (*RecoverSqlParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *CreateRecoveryRequest) SetMssqlParams(v RecoverSqlParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *CreateRecoveryRequest) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *CreateRecoveryRequest) GetNetappParams() RecoverNetappParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *CreateRecoveryRequest) GetNetappParamsOk() (*RecoverNetappParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *CreateRecoveryRequest) SetNetappParams(v RecoverNetappParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *CreateRecoveryRequest) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *CreateRecoveryRequest) GetOffice365Params() RecoverO365Params`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *CreateRecoveryRequest) GetOffice365ParamsOk() (*RecoverO365Params, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *CreateRecoveryRequest) SetOffice365Params(v RecoverO365Params)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *CreateRecoveryRequest) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *CreateRecoveryRequest) GetOracleParams() RecoverOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *CreateRecoveryRequest) GetOracleParamsOk() (*RecoverOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *CreateRecoveryRequest) SetOracleParams(v RecoverOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *CreateRecoveryRequest) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *CreateRecoveryRequest) GetPhysicalParams() RecoverPhysicalParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *CreateRecoveryRequest) GetPhysicalParamsOk() (*RecoverPhysicalParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *CreateRecoveryRequest) SetPhysicalParams(v RecoverPhysicalParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *CreateRecoveryRequest) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetPureParams

`func (o *CreateRecoveryRequest) GetPureParams() RecoverPureParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *CreateRecoveryRequest) GetPureParamsOk() (*RecoverPureParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *CreateRecoveryRequest) SetPureParams(v RecoverPureParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *CreateRecoveryRequest) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *CreateRecoveryRequest) GetSapHanaParams() SapHanaParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *CreateRecoveryRequest) GetSapHanaParamsOk() (*SapHanaParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *CreateRecoveryRequest) SetSapHanaParams(v SapHanaParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *CreateRecoveryRequest) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *CreateRecoveryRequest) GetSfdcParams() RecoverSalesforceParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *CreateRecoveryRequest) GetSfdcParamsOk() (*RecoverSalesforceParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *CreateRecoveryRequest) SetSfdcParams(v RecoverSalesforceParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *CreateRecoveryRequest) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *CreateRecoveryRequest) GetUdaParams() UdaParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *CreateRecoveryRequest) GetUdaParamsOk() (*UdaParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *CreateRecoveryRequest) SetUdaParams(v UdaParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *CreateRecoveryRequest) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *CreateRecoveryRequest) GetViewParams() RecoverViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *CreateRecoveryRequest) GetViewParamsOk() (*RecoverViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *CreateRecoveryRequest) SetViewParams(v RecoverViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *CreateRecoveryRequest) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *CreateRecoveryRequest) GetVmwareParams() RecoverVmwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *CreateRecoveryRequest) GetVmwareParamsOk() (*RecoverVmwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *CreateRecoveryRequest) SetVmwareParams(v RecoverVmwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *CreateRecoveryRequest) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


