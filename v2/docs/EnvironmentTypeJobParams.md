# EnvironmentTypeJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsSnapshotParams** | Pointer to [**AwsSnapshotManagerParams**](AwsSnapshotManagerParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**ExchangeEnvJobParams**](ExchangeEnvJobParams.md) |  | [optional] 
**ExternallyTriggeredJobParams** | Pointer to [**ExternallyTriggeredJobParams**](ExternallyTriggeredJobParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HypervEnvJobParams**](HypervEnvJobParams.md) |  | [optional] 
**NasParams** | Pointer to [**NasEnvJobParams**](NasEnvJobParams.md) |  | [optional] 
**Office365Params** | Pointer to [**O365EnvJobParams**](O365EnvJobParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleEnvJobParams**](OracleEnvJobParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalEnvJobParams**](PhysicalEnvJobParams.md) |  | [optional] 
**PureParams** | Pointer to [**SanEnvJobParams**](SanEnvJobParams.md) |  | [optional] 
**SqlParams** | Pointer to [**CommonMSSQLProtectionGroupParams**](CommonMSSQLProtectionGroupParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareEnvJobParams**](VmwareEnvJobParams.md) |  | [optional] 

## Methods

### NewEnvironmentTypeJobParams

`func NewEnvironmentTypeJobParams() *EnvironmentTypeJobParams`

NewEnvironmentTypeJobParams instantiates a new EnvironmentTypeJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentTypeJobParamsWithDefaults

`func NewEnvironmentTypeJobParamsWithDefaults() *EnvironmentTypeJobParams`

NewEnvironmentTypeJobParamsWithDefaults instantiates a new EnvironmentTypeJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsSnapshotParams

`func (o *EnvironmentTypeJobParams) GetAwsSnapshotParams() AwsSnapshotManagerParams`

GetAwsSnapshotParams returns the AwsSnapshotParams field if non-nil, zero value otherwise.

### GetAwsSnapshotParamsOk

`func (o *EnvironmentTypeJobParams) GetAwsSnapshotParamsOk() (*AwsSnapshotManagerParams, bool)`

GetAwsSnapshotParamsOk returns a tuple with the AwsSnapshotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSnapshotParams

`func (o *EnvironmentTypeJobParams) SetAwsSnapshotParams(v AwsSnapshotManagerParams)`

SetAwsSnapshotParams sets AwsSnapshotParams field to given value.

### HasAwsSnapshotParams

`func (o *EnvironmentTypeJobParams) HasAwsSnapshotParams() bool`

HasAwsSnapshotParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *EnvironmentTypeJobParams) GetExchangeParams() ExchangeEnvJobParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *EnvironmentTypeJobParams) GetExchangeParamsOk() (*ExchangeEnvJobParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *EnvironmentTypeJobParams) SetExchangeParams(v ExchangeEnvJobParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *EnvironmentTypeJobParams) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetExternallyTriggeredJobParams

`func (o *EnvironmentTypeJobParams) GetExternallyTriggeredJobParams() ExternallyTriggeredJobParams`

GetExternallyTriggeredJobParams returns the ExternallyTriggeredJobParams field if non-nil, zero value otherwise.

### GetExternallyTriggeredJobParamsOk

`func (o *EnvironmentTypeJobParams) GetExternallyTriggeredJobParamsOk() (*ExternallyTriggeredJobParams, bool)`

GetExternallyTriggeredJobParamsOk returns a tuple with the ExternallyTriggeredJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyTriggeredJobParams

`func (o *EnvironmentTypeJobParams) SetExternallyTriggeredJobParams(v ExternallyTriggeredJobParams)`

SetExternallyTriggeredJobParams sets ExternallyTriggeredJobParams field to given value.

### HasExternallyTriggeredJobParams

`func (o *EnvironmentTypeJobParams) HasExternallyTriggeredJobParams() bool`

HasExternallyTriggeredJobParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *EnvironmentTypeJobParams) GetHypervParams() HypervEnvJobParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *EnvironmentTypeJobParams) GetHypervParamsOk() (*HypervEnvJobParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *EnvironmentTypeJobParams) SetHypervParams(v HypervEnvJobParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *EnvironmentTypeJobParams) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetNasParams

`func (o *EnvironmentTypeJobParams) GetNasParams() NasEnvJobParams`

GetNasParams returns the NasParams field if non-nil, zero value otherwise.

### GetNasParamsOk

`func (o *EnvironmentTypeJobParams) GetNasParamsOk() (*NasEnvJobParams, bool)`

GetNasParamsOk returns a tuple with the NasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasParams

`func (o *EnvironmentTypeJobParams) SetNasParams(v NasEnvJobParams)`

SetNasParams sets NasParams field to given value.

### HasNasParams

`func (o *EnvironmentTypeJobParams) HasNasParams() bool`

HasNasParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *EnvironmentTypeJobParams) GetOffice365Params() O365EnvJobParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *EnvironmentTypeJobParams) GetOffice365ParamsOk() (*O365EnvJobParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *EnvironmentTypeJobParams) SetOffice365Params(v O365EnvJobParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *EnvironmentTypeJobParams) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *EnvironmentTypeJobParams) GetOracleParams() OracleEnvJobParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *EnvironmentTypeJobParams) GetOracleParamsOk() (*OracleEnvJobParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *EnvironmentTypeJobParams) SetOracleParams(v OracleEnvJobParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *EnvironmentTypeJobParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *EnvironmentTypeJobParams) GetPhysicalParams() PhysicalEnvJobParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *EnvironmentTypeJobParams) GetPhysicalParamsOk() (*PhysicalEnvJobParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *EnvironmentTypeJobParams) SetPhysicalParams(v PhysicalEnvJobParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *EnvironmentTypeJobParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetPureParams

`func (o *EnvironmentTypeJobParams) GetPureParams() SanEnvJobParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *EnvironmentTypeJobParams) GetPureParamsOk() (*SanEnvJobParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *EnvironmentTypeJobParams) SetPureParams(v SanEnvJobParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *EnvironmentTypeJobParams) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### GetSqlParams

`func (o *EnvironmentTypeJobParams) GetSqlParams() CommonMSSQLProtectionGroupParams`

GetSqlParams returns the SqlParams field if non-nil, zero value otherwise.

### GetSqlParamsOk

`func (o *EnvironmentTypeJobParams) GetSqlParamsOk() (*CommonMSSQLProtectionGroupParams, bool)`

GetSqlParamsOk returns a tuple with the SqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlParams

`func (o *EnvironmentTypeJobParams) SetSqlParams(v CommonMSSQLProtectionGroupParams)`

SetSqlParams sets SqlParams field to given value.

### HasSqlParams

`func (o *EnvironmentTypeJobParams) HasSqlParams() bool`

HasSqlParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *EnvironmentTypeJobParams) GetVmwareParams() VmwareEnvJobParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *EnvironmentTypeJobParams) GetVmwareParamsOk() (*VmwareEnvJobParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *EnvironmentTypeJobParams) SetVmwareParams(v VmwareEnvJobParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *EnvironmentTypeJobParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


