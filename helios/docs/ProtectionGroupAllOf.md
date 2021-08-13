# ProtectionGroupAllOf

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
**ElastifileParams** | Pointer to [**NullableElastifileProtectionGroupParams**](ElastifileProtectionGroupParams.md) |  | [optional] 
**RemoteAdapterParams** | Pointer to [**NullableRemoteAdapterProtectionGroupParams**](RemoteAdapterProtectionGroupParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**ExchangeProtectionGroupParams**](ExchangeProtectionGroupParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraProtectionGroupParams**](CassandraProtectionGroupParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDBProtectionGroupParams**](MongoDBProtectionGroupParams.md) |  | [optional] 
**HiveParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsProtectionGroupParams**](HdfsProtectionGroupParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaProtectionGroupParams**](UdaProtectionGroupParams.md) |  | [optional] 

## Methods

### NewProtectionGroupAllOf

`func NewProtectionGroupAllOf() *ProtectionGroupAllOf`

NewProtectionGroupAllOf instantiates a new ProtectionGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionGroupAllOfWithDefaults

`func NewProtectionGroupAllOfWithDefaults() *ProtectionGroupAllOf`

NewProtectionGroupAllOfWithDefaults instantiates a new ProtectionGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *ProtectionGroupAllOf) GetVmwareParams() VmwareProtectionGroupParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ProtectionGroupAllOf) GetVmwareParamsOk() (*VmwareProtectionGroupParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ProtectionGroupAllOf) SetVmwareParams(v VmwareProtectionGroupParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ProtectionGroupAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetAcropolisParams

`func (o *ProtectionGroupAllOf) GetAcropolisParams() AcropolisProtectionGroupParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *ProtectionGroupAllOf) GetAcropolisParamsOk() (*AcropolisProtectionGroupParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *ProtectionGroupAllOf) SetAcropolisParams(v AcropolisProtectionGroupParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *ProtectionGroupAllOf) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *ProtectionGroupAllOf) GetKubernetesParams() KubernetesProtectionGroupParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *ProtectionGroupAllOf) GetKubernetesParamsOk() (*KubernetesProtectionGroupParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *ProtectionGroupAllOf) SetKubernetesParams(v KubernetesProtectionGroupParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *ProtectionGroupAllOf) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ProtectionGroupAllOf) GetMssqlParams() MSSQLProtectionGroupParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ProtectionGroupAllOf) GetMssqlParamsOk() (*MSSQLProtectionGroupParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ProtectionGroupAllOf) SetMssqlParams(v MSSQLProtectionGroupParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ProtectionGroupAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ProtectionGroupAllOf) GetOracleParams() OracleProtectionGroupParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ProtectionGroupAllOf) GetOracleParamsOk() (*OracleProtectionGroupParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ProtectionGroupAllOf) SetOracleParams(v OracleProtectionGroupParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ProtectionGroupAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetViewParams

`func (o *ProtectionGroupAllOf) GetViewParams() ViewProtectionGroupParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *ProtectionGroupAllOf) GetViewParamsOk() (*ViewProtectionGroupParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *ProtectionGroupAllOf) SetViewParams(v ViewProtectionGroupParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *ProtectionGroupAllOf) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### SetViewParamsNil

`func (o *ProtectionGroupAllOf) SetViewParamsNil(b bool)`

 SetViewParamsNil sets the value for ViewParams to be an explicit nil

### UnsetViewParams
`func (o *ProtectionGroupAllOf) UnsetViewParams()`

UnsetViewParams ensures that no value is present for ViewParams, not even an explicit nil
### GetPureParams

`func (o *ProtectionGroupAllOf) GetPureParams() PureProtectionGroupParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *ProtectionGroupAllOf) GetPureParamsOk() (*PureProtectionGroupParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *ProtectionGroupAllOf) SetPureParams(v PureProtectionGroupParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *ProtectionGroupAllOf) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### SetPureParamsNil

`func (o *ProtectionGroupAllOf) SetPureParamsNil(b bool)`

 SetPureParamsNil sets the value for PureParams to be an explicit nil

### UnsetPureParams
`func (o *ProtectionGroupAllOf) UnsetPureParams()`

UnsetPureParams ensures that no value is present for PureParams, not even an explicit nil
### GetNimbleParams

`func (o *ProtectionGroupAllOf) GetNimbleParams() NimbleProtectionGroupParams`

GetNimbleParams returns the NimbleParams field if non-nil, zero value otherwise.

### GetNimbleParamsOk

`func (o *ProtectionGroupAllOf) GetNimbleParamsOk() (*NimbleProtectionGroupParams, bool)`

GetNimbleParamsOk returns a tuple with the NimbleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleParams

`func (o *ProtectionGroupAllOf) SetNimbleParams(v NimbleProtectionGroupParams)`

SetNimbleParams sets NimbleParams field to given value.

### HasNimbleParams

`func (o *ProtectionGroupAllOf) HasNimbleParams() bool`

HasNimbleParams returns a boolean if a field has been set.

### SetNimbleParamsNil

`func (o *ProtectionGroupAllOf) SetNimbleParamsNil(b bool)`

 SetNimbleParamsNil sets the value for NimbleParams to be an explicit nil

### UnsetNimbleParams
`func (o *ProtectionGroupAllOf) UnsetNimbleParams()`

UnsetNimbleParams ensures that no value is present for NimbleParams, not even an explicit nil
### GetHypervParams

`func (o *ProtectionGroupAllOf) GetHypervParams() HyperVProtectionGroupParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *ProtectionGroupAllOf) GetHypervParamsOk() (*HyperVProtectionGroupParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *ProtectionGroupAllOf) SetHypervParams(v HyperVProtectionGroupParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *ProtectionGroupAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *ProtectionGroupAllOf) GetAwsParams() AwsProtectionGroupParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ProtectionGroupAllOf) GetAwsParamsOk() (*AwsProtectionGroupParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ProtectionGroupAllOf) SetAwsParams(v AwsProtectionGroupParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ProtectionGroupAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *ProtectionGroupAllOf) GetAzureParams() AzureProtectionGroupParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *ProtectionGroupAllOf) GetAzureParamsOk() (*AzureProtectionGroupParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *ProtectionGroupAllOf) SetAzureParams(v AzureProtectionGroupParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *ProtectionGroupAllOf) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *ProtectionGroupAllOf) GetGcpParams() GcpProtectionGroupParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *ProtectionGroupAllOf) GetGcpParamsOk() (*GcpProtectionGroupParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *ProtectionGroupAllOf) SetGcpParams(v GcpProtectionGroupParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *ProtectionGroupAllOf) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *ProtectionGroupAllOf) GetKvmParams() KvmProtectionGroupParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *ProtectionGroupAllOf) GetKvmParamsOk() (*KvmProtectionGroupParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *ProtectionGroupAllOf) SetKvmParams(v KvmProtectionGroupParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *ProtectionGroupAllOf) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### SetKvmParamsNil

`func (o *ProtectionGroupAllOf) SetKvmParamsNil(b bool)`

 SetKvmParamsNil sets the value for KvmParams to be an explicit nil

### UnsetKvmParams
`func (o *ProtectionGroupAllOf) UnsetKvmParams()`

UnsetKvmParams ensures that no value is present for KvmParams, not even an explicit nil
### GetPhysicalParams

`func (o *ProtectionGroupAllOf) GetPhysicalParams() PhysicalProtectionGroupParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ProtectionGroupAllOf) GetPhysicalParamsOk() (*PhysicalProtectionGroupParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ProtectionGroupAllOf) SetPhysicalParams(v PhysicalProtectionGroupParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ProtectionGroupAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetAdParams

`func (o *ProtectionGroupAllOf) GetAdParams() ADProtectionGroupParams`

GetAdParams returns the AdParams field if non-nil, zero value otherwise.

### GetAdParamsOk

`func (o *ProtectionGroupAllOf) GetAdParamsOk() (*ADProtectionGroupParams, bool)`

GetAdParamsOk returns a tuple with the AdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdParams

`func (o *ProtectionGroupAllOf) SetAdParams(v ADProtectionGroupParams)`

SetAdParams sets AdParams field to given value.

### HasAdParams

`func (o *ProtectionGroupAllOf) HasAdParams() bool`

HasAdParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *ProtectionGroupAllOf) GetOffice365Params() Office365ProtectionGroupParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *ProtectionGroupAllOf) GetOffice365ParamsOk() (*Office365ProtectionGroupParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *ProtectionGroupAllOf) SetOffice365Params(v Office365ProtectionGroupParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *ProtectionGroupAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetNetappParams

`func (o *ProtectionGroupAllOf) GetNetappParams() NetappProtectionGroupParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ProtectionGroupAllOf) GetNetappParamsOk() (*NetappProtectionGroupParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ProtectionGroupAllOf) SetNetappParams(v NetappProtectionGroupParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ProtectionGroupAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### SetNetappParamsNil

`func (o *ProtectionGroupAllOf) SetNetappParamsNil(b bool)`

 SetNetappParamsNil sets the value for NetappParams to be an explicit nil

### UnsetNetappParams
`func (o *ProtectionGroupAllOf) UnsetNetappParams()`

UnsetNetappParams ensures that no value is present for NetappParams, not even an explicit nil
### GetGenericNasParams

`func (o *ProtectionGroupAllOf) GetGenericNasParams() GenericNasProtectionGroupParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ProtectionGroupAllOf) GetGenericNasParamsOk() (*GenericNasProtectionGroupParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ProtectionGroupAllOf) SetGenericNasParams(v GenericNasProtectionGroupParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ProtectionGroupAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ProtectionGroupAllOf) GetIsilonParams() IsilonProtectionGroupParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ProtectionGroupAllOf) GetIsilonParamsOk() (*IsilonProtectionGroupParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ProtectionGroupAllOf) SetIsilonParams(v IsilonProtectionGroupParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ProtectionGroupAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### SetIsilonParamsNil

`func (o *ProtectionGroupAllOf) SetIsilonParamsNil(b bool)`

 SetIsilonParamsNil sets the value for IsilonParams to be an explicit nil

### UnsetIsilonParams
`func (o *ProtectionGroupAllOf) UnsetIsilonParams()`

UnsetIsilonParams ensures that no value is present for IsilonParams, not even an explicit nil
### GetFlashbladeParams

`func (o *ProtectionGroupAllOf) GetFlashbladeParams() FlashbladeProtectionGroupParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ProtectionGroupAllOf) GetFlashbladeParamsOk() (*FlashbladeProtectionGroupParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ProtectionGroupAllOf) SetFlashbladeParams(v FlashbladeProtectionGroupParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ProtectionGroupAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### SetFlashbladeParamsNil

`func (o *ProtectionGroupAllOf) SetFlashbladeParamsNil(b bool)`

 SetFlashbladeParamsNil sets the value for FlashbladeParams to be an explicit nil

### UnsetFlashbladeParams
`func (o *ProtectionGroupAllOf) UnsetFlashbladeParams()`

UnsetFlashbladeParams ensures that no value is present for FlashbladeParams, not even an explicit nil
### GetGpfsParams

`func (o *ProtectionGroupAllOf) GetGpfsParams() GpfsProtectionGroupParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ProtectionGroupAllOf) GetGpfsParamsOk() (*GpfsProtectionGroupParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ProtectionGroupAllOf) SetGpfsParams(v GpfsProtectionGroupParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ProtectionGroupAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### SetGpfsParamsNil

`func (o *ProtectionGroupAllOf) SetGpfsParamsNil(b bool)`

 SetGpfsParamsNil sets the value for GpfsParams to be an explicit nil

### UnsetGpfsParams
`func (o *ProtectionGroupAllOf) UnsetGpfsParams()`

UnsetGpfsParams ensures that no value is present for GpfsParams, not even an explicit nil
### GetElastifileParams

`func (o *ProtectionGroupAllOf) GetElastifileParams() ElastifileProtectionGroupParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ProtectionGroupAllOf) GetElastifileParamsOk() (*ElastifileProtectionGroupParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ProtectionGroupAllOf) SetElastifileParams(v ElastifileProtectionGroupParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ProtectionGroupAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### SetElastifileParamsNil

`func (o *ProtectionGroupAllOf) SetElastifileParamsNil(b bool)`

 SetElastifileParamsNil sets the value for ElastifileParams to be an explicit nil

### UnsetElastifileParams
`func (o *ProtectionGroupAllOf) UnsetElastifileParams()`

UnsetElastifileParams ensures that no value is present for ElastifileParams, not even an explicit nil
### GetRemoteAdapterParams

`func (o *ProtectionGroupAllOf) GetRemoteAdapterParams() RemoteAdapterProtectionGroupParams`

GetRemoteAdapterParams returns the RemoteAdapterParams field if non-nil, zero value otherwise.

### GetRemoteAdapterParamsOk

`func (o *ProtectionGroupAllOf) GetRemoteAdapterParamsOk() (*RemoteAdapterProtectionGroupParams, bool)`

GetRemoteAdapterParamsOk returns a tuple with the RemoteAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdapterParams

`func (o *ProtectionGroupAllOf) SetRemoteAdapterParams(v RemoteAdapterProtectionGroupParams)`

SetRemoteAdapterParams sets RemoteAdapterParams field to given value.

### HasRemoteAdapterParams

`func (o *ProtectionGroupAllOf) HasRemoteAdapterParams() bool`

HasRemoteAdapterParams returns a boolean if a field has been set.

### SetRemoteAdapterParamsNil

`func (o *ProtectionGroupAllOf) SetRemoteAdapterParamsNil(b bool)`

 SetRemoteAdapterParamsNil sets the value for RemoteAdapterParams to be an explicit nil

### UnsetRemoteAdapterParams
`func (o *ProtectionGroupAllOf) UnsetRemoteAdapterParams()`

UnsetRemoteAdapterParams ensures that no value is present for RemoteAdapterParams, not even an explicit nil
### GetExchangeParams

`func (o *ProtectionGroupAllOf) GetExchangeParams() ExchangeProtectionGroupParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *ProtectionGroupAllOf) GetExchangeParamsOk() (*ExchangeProtectionGroupParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *ProtectionGroupAllOf) SetExchangeParams(v ExchangeProtectionGroupParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *ProtectionGroupAllOf) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *ProtectionGroupAllOf) GetCassandraParams() CassandraProtectionGroupParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *ProtectionGroupAllOf) GetCassandraParamsOk() (*CassandraProtectionGroupParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *ProtectionGroupAllOf) SetCassandraParams(v CassandraProtectionGroupParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *ProtectionGroupAllOf) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *ProtectionGroupAllOf) GetCouchbaseParams() NoSqlProtectionGroupParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *ProtectionGroupAllOf) GetCouchbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *ProtectionGroupAllOf) SetCouchbaseParams(v NoSqlProtectionGroupParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *ProtectionGroupAllOf) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *ProtectionGroupAllOf) GetMongodbParams() MongoDBProtectionGroupParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *ProtectionGroupAllOf) GetMongodbParamsOk() (*MongoDBProtectionGroupParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *ProtectionGroupAllOf) SetMongodbParams(v MongoDBProtectionGroupParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *ProtectionGroupAllOf) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *ProtectionGroupAllOf) GetHiveParams() NoSqlProtectionGroupParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *ProtectionGroupAllOf) GetHiveParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *ProtectionGroupAllOf) SetHiveParams(v NoSqlProtectionGroupParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *ProtectionGroupAllOf) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *ProtectionGroupAllOf) GetHdfsParams() HdfsProtectionGroupParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *ProtectionGroupAllOf) GetHdfsParamsOk() (*HdfsProtectionGroupParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *ProtectionGroupAllOf) SetHdfsParams(v HdfsProtectionGroupParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *ProtectionGroupAllOf) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *ProtectionGroupAllOf) GetHbaseParams() NoSqlProtectionGroupParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *ProtectionGroupAllOf) GetHbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *ProtectionGroupAllOf) SetHbaseParams(v NoSqlProtectionGroupParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *ProtectionGroupAllOf) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *ProtectionGroupAllOf) GetUdaParams() UdaProtectionGroupParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *ProtectionGroupAllOf) GetUdaParamsOk() (*UdaProtectionGroupParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *ProtectionGroupAllOf) SetUdaParams(v UdaProtectionGroupParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *ProtectionGroupAllOf) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


