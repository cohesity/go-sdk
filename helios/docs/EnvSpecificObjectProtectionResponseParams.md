# EnvSpecificObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment for current object. | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectProtectionResponseParams**](VmwareObjectProtectionResponseParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionResponseParams**](GenericNasObjectProtectionResponseParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionResponseParams**](GpfsObjectProtectionResponseParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionResponseParams**](ElastifileObjectProtectionResponseParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionResponseParams**](NetappObjectProtectionResponseParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionResponseParams**](IsilonObjectProtectionResponseParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionResponseParams**](FlashbladeObjectProtectionResponseParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionResponseParams**](MssqlObjectProtectionResponseParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionResponseParams**](OracleObjectProtectionResponseParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionResponseParams**](Office365ObjectProtectionResponseParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsObjectProtectionResponseParams**](AwsObjectProtectionResponseParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionResponseParams**](HyperVObjectProtectionResponseParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionResponseParams**](PhysicalObjectProtectionResponseParams.md) |  | [optional] 

## Methods

### NewEnvSpecificObjectProtectionResponseParams

`func NewEnvSpecificObjectProtectionResponseParams() *EnvSpecificObjectProtectionResponseParams`

NewEnvSpecificObjectProtectionResponseParams instantiates a new EnvSpecificObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvSpecificObjectProtectionResponseParamsWithDefaults

`func NewEnvSpecificObjectProtectionResponseParamsWithDefaults() *EnvSpecificObjectProtectionResponseParams`

NewEnvSpecificObjectProtectionResponseParamsWithDefaults instantiates a new EnvSpecificObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *EnvSpecificObjectProtectionResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvSpecificObjectProtectionResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvSpecificObjectProtectionResponseParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvSpecificObjectProtectionResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvSpecificObjectProtectionResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetVmwareParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetVmwareParams() VmwareObjectProtectionResponseParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetVmwareParamsOk() (*VmwareObjectProtectionResponseParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetVmwareParams(v VmwareObjectProtectionResponseParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetGenericNasParams() GenericNasObjectProtectionResponseParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetGenericNasParamsOk() (*GenericNasObjectProtectionResponseParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetGenericNasParams(v GenericNasObjectProtectionResponseParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetGpfsParams() GpfsObjectProtectionResponseParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetGpfsParamsOk() (*GpfsObjectProtectionResponseParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetGpfsParams(v GpfsObjectProtectionResponseParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetElastifileParams() ElastifileObjectProtectionResponseParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetElastifileParamsOk() (*ElastifileObjectProtectionResponseParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetElastifileParams(v ElastifileObjectProtectionResponseParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetNetappParams() NetappObjectProtectionResponseParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetNetappParamsOk() (*NetappObjectProtectionResponseParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetNetappParams(v NetappObjectProtectionResponseParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetIsilonParams() IsilonObjectProtectionResponseParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetIsilonParamsOk() (*IsilonObjectProtectionResponseParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetIsilonParams(v IsilonObjectProtectionResponseParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetFlashbladeParams() FlashbladeObjectProtectionResponseParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionResponseParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetFlashbladeParams(v FlashbladeObjectProtectionResponseParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetMssqlParams() MssqlObjectProtectionResponseParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetMssqlParamsOk() (*MssqlObjectProtectionResponseParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetMssqlParams(v MssqlObjectProtectionResponseParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetOracleParams() OracleObjectProtectionResponseParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetOracleParamsOk() (*OracleObjectProtectionResponseParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetOracleParams(v OracleObjectProtectionResponseParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *EnvSpecificObjectProtectionResponseParams) GetOffice365Params() Office365ObjectProtectionResponseParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetOffice365ParamsOk() (*Office365ObjectProtectionResponseParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *EnvSpecificObjectProtectionResponseParams) SetOffice365Params(v Office365ObjectProtectionResponseParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *EnvSpecificObjectProtectionResponseParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetAwsParams() AwsObjectProtectionResponseParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetAwsParamsOk() (*AwsObjectProtectionResponseParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetAwsParams(v AwsObjectProtectionResponseParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetHypervParams() HyperVObjectProtectionResponseParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetHypervParamsOk() (*HyperVObjectProtectionResponseParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetHypervParams(v HyperVObjectProtectionResponseParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *EnvSpecificObjectProtectionResponseParams) GetPhysicalParams() PhysicalObjectProtectionResponseParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *EnvSpecificObjectProtectionResponseParams) GetPhysicalParamsOk() (*PhysicalObjectProtectionResponseParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *EnvSpecificObjectProtectionResponseParams) SetPhysicalParams(v PhysicalObjectProtectionResponseParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *EnvSpecificObjectProtectionResponseParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


