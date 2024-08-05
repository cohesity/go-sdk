# EnvSpecificObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsParams** | Pointer to [**AwsObjectProtectionUpdateRequestParams**](AwsObjectProtectionUpdateRequestParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureObjectProtectionUpdateRequestParams**](AzureObjectProtectionUpdateRequestParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionUpdateRequestParams**](ElastifileObjectProtectionUpdateRequestParams.md) |  | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment for current object. | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterObjectProtectionUpdateRequestParams**](ExperimentalAdapterObjectProtectionUpdateRequestParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionUpdateRequestParams**](FlashbladeObjectProtectionUpdateRequestParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionUpdateRequestParams**](GenericNasObjectProtectionUpdateRequestParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionUpdateRequestParams**](GpfsObjectProtectionUpdateRequestParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionUpdateRequestParams**](HyperVObjectProtectionUpdateRequestParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionUpdateRequestParams**](IsilonObjectProtectionUpdateRequestParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionUpdateRequestParams**](MssqlObjectProtectionUpdateRequestParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionUpdateRequestParams**](NetappObjectProtectionUpdateRequestParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionUpdateRequestParams**](Office365ObjectProtectionUpdateRequestParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionUpdateRequestParams**](OracleObjectProtectionUpdateRequestParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionUpdateRequestParams**](PhysicalObjectProtectionUpdateRequestParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaObjectProtectionUpdateRequestParams**](SapHanaObjectProtectionUpdateRequestParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**SfdcObjectProtectionUpdateRequestParams**](SfdcObjectProtectionUpdateRequestParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaObjectProtectionUpdateRequestParams**](UdaObjectProtectionUpdateRequestParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectProtectionUpdateRequestParams**](VmwareObjectProtectionUpdateRequestParams.md) |  | [optional] 

## Methods

### NewEnvSpecificObjectProtectionUpdateRequestParams

`func NewEnvSpecificObjectProtectionUpdateRequestParams() *EnvSpecificObjectProtectionUpdateRequestParams`

NewEnvSpecificObjectProtectionUpdateRequestParams instantiates a new EnvSpecificObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvSpecificObjectProtectionUpdateRequestParamsWithDefaults

`func NewEnvSpecificObjectProtectionUpdateRequestParamsWithDefaults() *EnvSpecificObjectProtectionUpdateRequestParams`

NewEnvSpecificObjectProtectionUpdateRequestParamsWithDefaults instantiates a new EnvSpecificObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetAwsParams() AwsObjectProtectionUpdateRequestParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetAwsParamsOk() (*AwsObjectProtectionUpdateRequestParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetAwsParams(v AwsObjectProtectionUpdateRequestParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetAzureParams() AzureObjectProtectionUpdateRequestParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetAzureParamsOk() (*AzureObjectProtectionUpdateRequestParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetAzureParams(v AzureObjectProtectionUpdateRequestParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetElastifileParams() ElastifileObjectProtectionUpdateRequestParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetElastifileParamsOk() (*ElastifileObjectProtectionUpdateRequestParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetElastifileParams(v ElastifileObjectProtectionUpdateRequestParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetEnvironment

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvSpecificObjectProtectionUpdateRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetExperimentalAdapterParams() ExperimentalAdapterObjectProtectionUpdateRequestParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterObjectProtectionUpdateRequestParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetExperimentalAdapterParams(v ExperimentalAdapterObjectProtectionUpdateRequestParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetFlashbladeParams() FlashbladeObjectProtectionUpdateRequestParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionUpdateRequestParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetFlashbladeParams(v FlashbladeObjectProtectionUpdateRequestParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetGenericNasParams() GenericNasObjectProtectionUpdateRequestParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetGenericNasParamsOk() (*GenericNasObjectProtectionUpdateRequestParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetGenericNasParams(v GenericNasObjectProtectionUpdateRequestParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetGpfsParams() GpfsObjectProtectionUpdateRequestParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetGpfsParamsOk() (*GpfsObjectProtectionUpdateRequestParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetGpfsParams(v GpfsObjectProtectionUpdateRequestParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetHypervParams() HyperVObjectProtectionUpdateRequestParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetHypervParamsOk() (*HyperVObjectProtectionUpdateRequestParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetHypervParams(v HyperVObjectProtectionUpdateRequestParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetIsilonParams() IsilonObjectProtectionUpdateRequestParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetIsilonParamsOk() (*IsilonObjectProtectionUpdateRequestParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetIsilonParams(v IsilonObjectProtectionUpdateRequestParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetMssqlParams() MssqlObjectProtectionUpdateRequestParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetMssqlParamsOk() (*MssqlObjectProtectionUpdateRequestParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetMssqlParams(v MssqlObjectProtectionUpdateRequestParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetNetappParams() NetappObjectProtectionUpdateRequestParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetNetappParamsOk() (*NetappObjectProtectionUpdateRequestParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetNetappParams(v NetappObjectProtectionUpdateRequestParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetOffice365Params() Office365ObjectProtectionUpdateRequestParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetOffice365ParamsOk() (*Office365ObjectProtectionUpdateRequestParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetOffice365Params(v Office365ObjectProtectionUpdateRequestParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetOracleParams() OracleObjectProtectionUpdateRequestParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetOracleParamsOk() (*OracleObjectProtectionUpdateRequestParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetOracleParams(v OracleObjectProtectionUpdateRequestParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetPhysicalParams() PhysicalObjectProtectionUpdateRequestParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetPhysicalParamsOk() (*PhysicalObjectProtectionUpdateRequestParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetPhysicalParams(v PhysicalObjectProtectionUpdateRequestParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetSapHanaParams() SapHanaObjectProtectionUpdateRequestParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetSapHanaParamsOk() (*SapHanaObjectProtectionUpdateRequestParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetSapHanaParams(v SapHanaObjectProtectionUpdateRequestParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetSfdcParams() SfdcObjectProtectionUpdateRequestParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetSfdcParamsOk() (*SfdcObjectProtectionUpdateRequestParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetSfdcParams(v SfdcObjectProtectionUpdateRequestParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetUdaParams() UdaObjectProtectionUpdateRequestParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetUdaParamsOk() (*UdaObjectProtectionUpdateRequestParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetUdaParams(v UdaObjectProtectionUpdateRequestParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetVmwareParams() VmwareObjectProtectionUpdateRequestParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) GetVmwareParamsOk() (*VmwareObjectProtectionUpdateRequestParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) SetVmwareParams(v VmwareObjectProtectionUpdateRequestParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *EnvSpecificObjectProtectionUpdateRequestParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


