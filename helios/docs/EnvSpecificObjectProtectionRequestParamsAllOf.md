# EnvSpecificObjectProtectionRequestParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmwareParams** | Pointer to [**VmwareObjectProtectionRequestParams**](VmwareObjectProtectionRequestParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionRequestParams**](GenericNasObjectProtectionRequestParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionRequestParams**](GpfsObjectProtectionRequestParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionRequestParams**](ElastifileObjectProtectionRequestParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionRequestParams**](NetappObjectProtectionRequestParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionRequestParams**](IsilonObjectProtectionRequestParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionRequestParams**](FlashbladeObjectProtectionRequestParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionRequestParams**](MssqlObjectProtectionRequestParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionRequestParams**](OracleObjectProtectionRequestParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionRequestParams**](Office365ObjectProtectionRequestParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsObjectProtectionRequestParams**](AwsObjectProtectionRequestParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionRequestParams**](HyperVObjectProtectionRequestParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionRequestParams**](PhysicalObjectProtectionRequestParams.md) |  | [optional] 

## Methods

### NewEnvSpecificObjectProtectionRequestParamsAllOf

`func NewEnvSpecificObjectProtectionRequestParamsAllOf() *EnvSpecificObjectProtectionRequestParamsAllOf`

NewEnvSpecificObjectProtectionRequestParamsAllOf instantiates a new EnvSpecificObjectProtectionRequestParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvSpecificObjectProtectionRequestParamsAllOfWithDefaults

`func NewEnvSpecificObjectProtectionRequestParamsAllOfWithDefaults() *EnvSpecificObjectProtectionRequestParamsAllOf`

NewEnvSpecificObjectProtectionRequestParamsAllOfWithDefaults instantiates a new EnvSpecificObjectProtectionRequestParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetVmwareParams() VmwareObjectProtectionRequestParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetVmwareParamsOk() (*VmwareObjectProtectionRequestParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetVmwareParams(v VmwareObjectProtectionRequestParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetGenericNasParams() GenericNasObjectProtectionRequestParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetGenericNasParamsOk() (*GenericNasObjectProtectionRequestParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetGenericNasParams(v GenericNasObjectProtectionRequestParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetGpfsParams() GpfsObjectProtectionRequestParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetGpfsParamsOk() (*GpfsObjectProtectionRequestParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetGpfsParams(v GpfsObjectProtectionRequestParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetElastifileParams() ElastifileObjectProtectionRequestParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetElastifileParamsOk() (*ElastifileObjectProtectionRequestParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetElastifileParams(v ElastifileObjectProtectionRequestParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetNetappParams() NetappObjectProtectionRequestParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetNetappParamsOk() (*NetappObjectProtectionRequestParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetNetappParams(v NetappObjectProtectionRequestParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetIsilonParams() IsilonObjectProtectionRequestParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetIsilonParamsOk() (*IsilonObjectProtectionRequestParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetIsilonParams(v IsilonObjectProtectionRequestParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetFlashbladeParams() FlashbladeObjectProtectionRequestParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionRequestParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetFlashbladeParams(v FlashbladeObjectProtectionRequestParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetMssqlParams() MssqlObjectProtectionRequestParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetMssqlParamsOk() (*MssqlObjectProtectionRequestParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetMssqlParams(v MssqlObjectProtectionRequestParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetOracleParams() OracleObjectProtectionRequestParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetOracleParamsOk() (*OracleObjectProtectionRequestParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetOracleParams(v OracleObjectProtectionRequestParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetOffice365Params() Office365ObjectProtectionRequestParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetOffice365ParamsOk() (*Office365ObjectProtectionRequestParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetOffice365Params(v Office365ObjectProtectionRequestParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetAwsParams() AwsObjectProtectionRequestParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetAwsParamsOk() (*AwsObjectProtectionRequestParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetAwsParams(v AwsObjectProtectionRequestParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetHypervParams() HyperVObjectProtectionRequestParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetHypervParamsOk() (*HyperVObjectProtectionRequestParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetHypervParams(v HyperVObjectProtectionRequestParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetPhysicalParams() PhysicalObjectProtectionRequestParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) GetPhysicalParamsOk() (*PhysicalObjectProtectionRequestParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) SetPhysicalParams(v PhysicalObjectProtectionRequestParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *EnvSpecificObjectProtectionRequestParamsAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


