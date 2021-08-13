# CreateOrUpdateProtectionGroupRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmwareParams** | Pointer to [**VmwareProtectionGroupParams**](VmwareProtectionGroupParams.md) |  | [optional] 
**AcropolisParams** | Pointer to [**AcropolisProtectionGroupParams**](AcropolisProtectionGroupParams.md) |  | [optional] 
**KubernetesParams** | Pointer to [**KubernetesProtectionGroupParams**](KubernetesProtectionGroupParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MSSQLProtectionGroupParams**](MSSQLProtectionGroupParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleProtectionGroupParams**](OracleProtectionGroupParams.md) |  | [optional] 
**ViewParams** | Pointer to [**NullableViewProtectionGroupParams**](ViewProtectionGroupParams.md) |  | [optional] 
**PureParams** | Pointer to [**NullablePureProtectionGroupParams**](PureProtectionGroupParams.md) |  | [optional] 
**NimbleParams** | Pointer to [**NullableNimbleProtectionGroupParams**](NimbleProtectionGroupParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVProtectionGroupParams**](HyperVProtectionGroupParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsProtectionGroupParams**](AwsProtectionGroupParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureProtectionGroupParams**](AzureProtectionGroupParams.md) |  | [optional] 
**GcpParams** | Pointer to [**GcpProtectionGroupParams**](GcpProtectionGroupParams.md) |  | [optional] 
**KvmParams** | Pointer to [**NullableKvmProtectionGroupParams**](KvmProtectionGroupParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalProtectionGroupParams**](PhysicalProtectionGroupParams.md) |  | [optional] 
**AdParams** | Pointer to [**ADProtectionGroupParams**](ADProtectionGroupParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ProtectionGroupParams**](Office365ProtectionGroupParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NullableNetappProtectionGroupParams**](NetappProtectionGroupParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasProtectionGroupParams**](GenericNasProtectionGroupParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**NullableIsilonProtectionGroupParams**](IsilonProtectionGroupParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**NullableFlashbladeProtectionGroupParams**](FlashbladeProtectionGroupParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**NullableGpfsProtectionGroupParams**](GpfsProtectionGroupParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**NullableElastifileProtectionGroupParams**](ElastifileProtectionGroupParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraProtectionGroupParams**](CassandraProtectionGroupParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDBProtectionGroupParams**](MongoDBProtectionGroupParams.md) |  | [optional] 
**HiveParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsProtectionGroupParams**](HdfsProtectionGroupParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**RemoteAdapterParams** | Pointer to [**NullableRemoteAdapterProtectionGroupParams**](RemoteAdapterProtectionGroupParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**ExchangeProtectionGroupParams**](ExchangeProtectionGroupParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaProtectionGroupParams**](UdaProtectionGroupParams.md) |  | [optional] 

## Methods

### NewCreateOrUpdateProtectionGroupRequestAllOf

`func NewCreateOrUpdateProtectionGroupRequestAllOf() *CreateOrUpdateProtectionGroupRequestAllOf`

NewCreateOrUpdateProtectionGroupRequestAllOf instantiates a new CreateOrUpdateProtectionGroupRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateProtectionGroupRequestAllOfWithDefaults

`func NewCreateOrUpdateProtectionGroupRequestAllOfWithDefaults() *CreateOrUpdateProtectionGroupRequestAllOf`

NewCreateOrUpdateProtectionGroupRequestAllOfWithDefaults instantiates a new CreateOrUpdateProtectionGroupRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetVmwareParams() VmwareProtectionGroupParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetVmwareParamsOk() (*VmwareProtectionGroupParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetVmwareParams(v VmwareProtectionGroupParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAcropolisParams() AcropolisProtectionGroupParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAcropolisParamsOk() (*AcropolisProtectionGroupParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetAcropolisParams(v AcropolisProtectionGroupParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetKubernetesParams() KubernetesProtectionGroupParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetKubernetesParamsOk() (*KubernetesProtectionGroupParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetKubernetesParams(v KubernetesProtectionGroupParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetMssqlParams() MSSQLProtectionGroupParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetMssqlParamsOk() (*MSSQLProtectionGroupParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetMssqlParams(v MSSQLProtectionGroupParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetOracleParams() OracleProtectionGroupParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetOracleParamsOk() (*OracleProtectionGroupParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetOracleParams(v OracleProtectionGroupParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetViewParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetViewParams() ViewProtectionGroupParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetViewParamsOk() (*ViewProtectionGroupParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetViewParams(v ViewProtectionGroupParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### SetViewParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetViewParamsNil(b bool)`

 SetViewParamsNil sets the value for ViewParams to be an explicit nil

### UnsetViewParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetViewParams()`

UnsetViewParams ensures that no value is present for ViewParams, not even an explicit nil
### GetPureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetPureParams() PureProtectionGroupParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetPureParamsOk() (*PureProtectionGroupParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetPureParams(v PureProtectionGroupParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### SetPureParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetPureParamsNil(b bool)`

 SetPureParamsNil sets the value for PureParams to be an explicit nil

### UnsetPureParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetPureParams()`

UnsetPureParams ensures that no value is present for PureParams, not even an explicit nil
### GetNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetNimbleParams() NimbleProtectionGroupParams`

GetNimbleParams returns the NimbleParams field if non-nil, zero value otherwise.

### GetNimbleParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetNimbleParamsOk() (*NimbleProtectionGroupParams, bool)`

GetNimbleParamsOk returns a tuple with the NimbleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetNimbleParams(v NimbleProtectionGroupParams)`

SetNimbleParams sets NimbleParams field to given value.

### HasNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasNimbleParams() bool`

HasNimbleParams returns a boolean if a field has been set.

### SetNimbleParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetNimbleParamsNil(b bool)`

 SetNimbleParamsNil sets the value for NimbleParams to be an explicit nil

### UnsetNimbleParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetNimbleParams()`

UnsetNimbleParams ensures that no value is present for NimbleParams, not even an explicit nil
### GetHypervParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHypervParams() HyperVProtectionGroupParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHypervParamsOk() (*HyperVProtectionGroupParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetHypervParams(v HyperVProtectionGroupParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAwsParams() AwsProtectionGroupParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAwsParamsOk() (*AwsProtectionGroupParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetAwsParams(v AwsProtectionGroupParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAzureParams() AzureProtectionGroupParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAzureParamsOk() (*AzureProtectionGroupParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetAzureParams(v AzureProtectionGroupParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGcpParams() GcpProtectionGroupParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGcpParamsOk() (*GcpProtectionGroupParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetGcpParams(v GcpProtectionGroupParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetKvmParams() KvmProtectionGroupParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetKvmParamsOk() (*KvmProtectionGroupParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetKvmParams(v KvmProtectionGroupParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### SetKvmParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetKvmParamsNil(b bool)`

 SetKvmParamsNil sets the value for KvmParams to be an explicit nil

### UnsetKvmParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetKvmParams()`

UnsetKvmParams ensures that no value is present for KvmParams, not even an explicit nil
### GetPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetPhysicalParams() PhysicalProtectionGroupParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetPhysicalParamsOk() (*PhysicalProtectionGroupParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetPhysicalParams(v PhysicalProtectionGroupParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetAdParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAdParams() ADProtectionGroupParams`

GetAdParams returns the AdParams field if non-nil, zero value otherwise.

### GetAdParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetAdParamsOk() (*ADProtectionGroupParams, bool)`

GetAdParamsOk returns a tuple with the AdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetAdParams(v ADProtectionGroupParams)`

SetAdParams sets AdParams field to given value.

### HasAdParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasAdParams() bool`

HasAdParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetOffice365Params() Office365ProtectionGroupParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetOffice365ParamsOk() (*Office365ProtectionGroupParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetOffice365Params(v Office365ProtectionGroupParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetNetappParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetNetappParams() NetappProtectionGroupParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetNetappParamsOk() (*NetappProtectionGroupParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetNetappParams(v NetappProtectionGroupParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### SetNetappParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetNetappParamsNil(b bool)`

 SetNetappParamsNil sets the value for NetappParams to be an explicit nil

### UnsetNetappParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetNetappParams()`

UnsetNetappParams ensures that no value is present for NetappParams, not even an explicit nil
### GetGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGenericNasParams() GenericNasProtectionGroupParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGenericNasParamsOk() (*GenericNasProtectionGroupParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetGenericNasParams(v GenericNasProtectionGroupParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetIsilonParams() IsilonProtectionGroupParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetIsilonParamsOk() (*IsilonProtectionGroupParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetIsilonParams(v IsilonProtectionGroupParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### SetIsilonParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetIsilonParamsNil(b bool)`

 SetIsilonParamsNil sets the value for IsilonParams to be an explicit nil

### UnsetIsilonParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetIsilonParams()`

UnsetIsilonParams ensures that no value is present for IsilonParams, not even an explicit nil
### GetFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetFlashbladeParams() FlashbladeProtectionGroupParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetFlashbladeParamsOk() (*FlashbladeProtectionGroupParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetFlashbladeParams(v FlashbladeProtectionGroupParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### SetFlashbladeParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetFlashbladeParamsNil(b bool)`

 SetFlashbladeParamsNil sets the value for FlashbladeParams to be an explicit nil

### UnsetFlashbladeParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetFlashbladeParams()`

UnsetFlashbladeParams ensures that no value is present for FlashbladeParams, not even an explicit nil
### GetGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGpfsParams() GpfsProtectionGroupParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetGpfsParamsOk() (*GpfsProtectionGroupParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetGpfsParams(v GpfsProtectionGroupParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### SetGpfsParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetGpfsParamsNil(b bool)`

 SetGpfsParamsNil sets the value for GpfsParams to be an explicit nil

### UnsetGpfsParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetGpfsParams()`

UnsetGpfsParams ensures that no value is present for GpfsParams, not even an explicit nil
### GetCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetCouchbaseParams() NoSqlProtectionGroupParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetCouchbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetCouchbaseParams(v NoSqlProtectionGroupParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetElastifileParams() ElastifileProtectionGroupParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetElastifileParamsOk() (*ElastifileProtectionGroupParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetElastifileParams(v ElastifileProtectionGroupParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### SetElastifileParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetElastifileParamsNil(b bool)`

 SetElastifileParamsNil sets the value for ElastifileParams to be an explicit nil

### UnsetElastifileParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetElastifileParams()`

UnsetElastifileParams ensures that no value is present for ElastifileParams, not even an explicit nil
### GetCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetCassandraParams() CassandraProtectionGroupParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetCassandraParamsOk() (*CassandraProtectionGroupParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetCassandraParams(v CassandraProtectionGroupParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetMongodbParams() MongoDBProtectionGroupParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetMongodbParamsOk() (*MongoDBProtectionGroupParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetMongodbParams(v MongoDBProtectionGroupParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHiveParams() NoSqlProtectionGroupParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHiveParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetHiveParams(v NoSqlProtectionGroupParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHdfsParams() HdfsProtectionGroupParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHdfsParamsOk() (*HdfsProtectionGroupParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetHdfsParams(v HdfsProtectionGroupParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHbaseParams() NoSqlProtectionGroupParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetHbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetHbaseParams(v NoSqlProtectionGroupParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetRemoteAdapterParams() RemoteAdapterProtectionGroupParams`

GetRemoteAdapterParams returns the RemoteAdapterParams field if non-nil, zero value otherwise.

### GetRemoteAdapterParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetRemoteAdapterParamsOk() (*RemoteAdapterProtectionGroupParams, bool)`

GetRemoteAdapterParamsOk returns a tuple with the RemoteAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetRemoteAdapterParams(v RemoteAdapterProtectionGroupParams)`

SetRemoteAdapterParams sets RemoteAdapterParams field to given value.

### HasRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasRemoteAdapterParams() bool`

HasRemoteAdapterParams returns a boolean if a field has been set.

### SetRemoteAdapterParamsNil

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetRemoteAdapterParamsNil(b bool)`

 SetRemoteAdapterParamsNil sets the value for RemoteAdapterParams to be an explicit nil

### UnsetRemoteAdapterParams
`func (o *CreateOrUpdateProtectionGroupRequestAllOf) UnsetRemoteAdapterParams()`

UnsetRemoteAdapterParams ensures that no value is present for RemoteAdapterParams, not even an explicit nil
### GetExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetExchangeParams() ExchangeProtectionGroupParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetExchangeParamsOk() (*ExchangeProtectionGroupParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetExchangeParams(v ExchangeProtectionGroupParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetUdaParams() UdaProtectionGroupParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) GetUdaParamsOk() (*UdaProtectionGroupParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) SetUdaParams(v UdaProtectionGroupParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *CreateOrUpdateProtectionGroupRequestAllOf) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


