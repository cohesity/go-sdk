# EnvSpecificObjectProtectionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment for current object. | [optional] 
**AwsParams** | Pointer to [**AwsObjectProtectionRequestParams**](AwsObjectProtectionRequestParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureObjectProtectionRequestParams**](AzureObjectProtectionRequestParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionRequestParams**](ElastifileObjectProtectionRequestParams.md) |  | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterObjectProtectionRequestParams**](ExperimentalAdapterObjectProtectionRequestParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionRequestParams**](FlashbladeObjectProtectionRequestParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionRequestParams**](GenericNasObjectProtectionRequestParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionRequestParams**](GpfsObjectProtectionRequestParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionRequestParams**](HyperVObjectProtectionRequestParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionRequestParams**](IsilonObjectProtectionRequestParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionRequestParams**](MssqlObjectProtectionRequestParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionRequestParams**](NetappObjectProtectionRequestParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionRequestParams**](Office365ObjectProtectionRequestParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionRequestParams**](OracleObjectProtectionRequestParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionRequestParams**](PhysicalObjectProtectionRequestParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaObjectProtectionRequestParams**](SapHanaObjectProtectionRequestParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**SfdcObjectProtectionRequestParams**](SfdcObjectProtectionRequestParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaObjectProtectionRequestParams**](UdaObjectProtectionRequestParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectProtectionRequestParams**](VmwareObjectProtectionRequestParams.md) |  | [optional] 

## Methods

### NewEnvSpecificObjectProtectionRequestParams

`func NewEnvSpecificObjectProtectionRequestParams() *EnvSpecificObjectProtectionRequestParams`

NewEnvSpecificObjectProtectionRequestParams instantiates a new EnvSpecificObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvSpecificObjectProtectionRequestParamsWithDefaults

`func NewEnvSpecificObjectProtectionRequestParamsWithDefaults() *EnvSpecificObjectProtectionRequestParams`

NewEnvSpecificObjectProtectionRequestParamsWithDefaults instantiates a new EnvSpecificObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *EnvSpecificObjectProtectionRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvSpecificObjectProtectionRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvSpecificObjectProtectionRequestParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvSpecificObjectProtectionRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvSpecificObjectProtectionRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetAwsParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetAwsParams() AwsObjectProtectionRequestParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetAwsParamsOk() (*AwsObjectProtectionRequestParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetAwsParams(v AwsObjectProtectionRequestParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetAzureParams() AzureObjectProtectionRequestParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetAzureParamsOk() (*AzureObjectProtectionRequestParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetAzureParams(v AzureObjectProtectionRequestParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetElastifileParams() ElastifileObjectProtectionRequestParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetElastifileParamsOk() (*ElastifileObjectProtectionRequestParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetElastifileParams(v ElastifileObjectProtectionRequestParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetExperimentalAdapterParams() ExperimentalAdapterObjectProtectionRequestParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterObjectProtectionRequestParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetExperimentalAdapterParams(v ExperimentalAdapterObjectProtectionRequestParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetFlashbladeParams() FlashbladeObjectProtectionRequestParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionRequestParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetFlashbladeParams(v FlashbladeObjectProtectionRequestParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetGenericNasParams() GenericNasObjectProtectionRequestParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetGenericNasParamsOk() (*GenericNasObjectProtectionRequestParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetGenericNasParams(v GenericNasObjectProtectionRequestParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetGpfsParams() GpfsObjectProtectionRequestParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetGpfsParamsOk() (*GpfsObjectProtectionRequestParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetGpfsParams(v GpfsObjectProtectionRequestParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetHypervParams() HyperVObjectProtectionRequestParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetHypervParamsOk() (*HyperVObjectProtectionRequestParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetHypervParams(v HyperVObjectProtectionRequestParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetIsilonParams() IsilonObjectProtectionRequestParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetIsilonParamsOk() (*IsilonObjectProtectionRequestParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetIsilonParams(v IsilonObjectProtectionRequestParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetMssqlParams() MssqlObjectProtectionRequestParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetMssqlParamsOk() (*MssqlObjectProtectionRequestParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetMssqlParams(v MssqlObjectProtectionRequestParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetNetappParams() NetappObjectProtectionRequestParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetNetappParamsOk() (*NetappObjectProtectionRequestParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetNetappParams(v NetappObjectProtectionRequestParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParams) GetOffice365Params() Office365ObjectProtectionRequestParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetOffice365ParamsOk() (*Office365ObjectProtectionRequestParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParams) SetOffice365Params(v Office365ObjectProtectionRequestParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetOracleParams() OracleObjectProtectionRequestParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetOracleParamsOk() (*OracleObjectProtectionRequestParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetOracleParams(v OracleObjectProtectionRequestParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetPhysicalParams() PhysicalObjectProtectionRequestParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetPhysicalParamsOk() (*PhysicalObjectProtectionRequestParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetPhysicalParams(v PhysicalObjectProtectionRequestParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetSapHanaParams() SapHanaObjectProtectionRequestParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetSapHanaParamsOk() (*SapHanaObjectProtectionRequestParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetSapHanaParams(v SapHanaObjectProtectionRequestParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetSfdcParams() SfdcObjectProtectionRequestParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetSfdcParamsOk() (*SfdcObjectProtectionRequestParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetSfdcParams(v SfdcObjectProtectionRequestParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetUdaParams() UdaObjectProtectionRequestParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetUdaParamsOk() (*UdaObjectProtectionRequestParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetUdaParams(v UdaObjectProtectionRequestParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParams) GetVmwareParams() VmwareObjectProtectionRequestParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *EnvSpecificObjectProtectionRequestParams) GetVmwareParamsOk() (*VmwareObjectProtectionRequestParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParams) SetVmwareParams(v VmwareObjectProtectionRequestParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


