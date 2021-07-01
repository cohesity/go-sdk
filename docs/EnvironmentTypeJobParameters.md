# EnvironmentTypeJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsSnapshotParameters** | Pointer to [**AwsSnapshotManagerParameters**](AwsSnapshotManagerParameters.md) |  | [optional] 
**ExchangeParameters** | Pointer to [**ExchangeEnvJobParameters**](ExchangeEnvJobParameters.md) |  | [optional] 
**HypervParameters** | Pointer to [**HypervEnvJobParameters**](HypervEnvJobParameters.md) |  | [optional] 
**NasParameters** | Pointer to [**NasEnvJobParameters**](NasEnvJobParameters.md) |  | [optional] 
**Office365Parameters** | Pointer to [**Office365EnvJobParameters**](Office365EnvJobParameters.md) |  | [optional] 
**OracleParameters** | Pointer to [**OracleEnvJobParameters**](OracleEnvJobParameters.md) |  | [optional] 
**PhysicalParameters** | Pointer to [**PhysicalEnvJobParameters**](PhysicalEnvJobParameters.md) |  | [optional] 
**PureParameters** | Pointer to [**SanEnvJobParameters**](SanEnvJobParameters.md) |  | [optional] 
**SqlParameters** | Pointer to [**SqlEnvJobParameters**](SqlEnvJobParameters.md) |  | [optional] 
**VmwareParameters** | Pointer to [**VmwareEnvJobParameters**](VmwareEnvJobParameters.md) |  | [optional] 

## Methods

### NewEnvironmentTypeJobParameters

`func NewEnvironmentTypeJobParameters() *EnvironmentTypeJobParameters`

NewEnvironmentTypeJobParameters instantiates a new EnvironmentTypeJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentTypeJobParametersWithDefaults

`func NewEnvironmentTypeJobParametersWithDefaults() *EnvironmentTypeJobParameters`

NewEnvironmentTypeJobParametersWithDefaults instantiates a new EnvironmentTypeJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsSnapshotParameters

`func (o *EnvironmentTypeJobParameters) GetAwsSnapshotParameters() AwsSnapshotManagerParameters`

GetAwsSnapshotParameters returns the AwsSnapshotParameters field if non-nil, zero value otherwise.

### GetAwsSnapshotParametersOk

`func (o *EnvironmentTypeJobParameters) GetAwsSnapshotParametersOk() (*AwsSnapshotManagerParameters, bool)`

GetAwsSnapshotParametersOk returns a tuple with the AwsSnapshotParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSnapshotParameters

`func (o *EnvironmentTypeJobParameters) SetAwsSnapshotParameters(v AwsSnapshotManagerParameters)`

SetAwsSnapshotParameters sets AwsSnapshotParameters field to given value.

### HasAwsSnapshotParameters

`func (o *EnvironmentTypeJobParameters) HasAwsSnapshotParameters() bool`

HasAwsSnapshotParameters returns a boolean if a field has been set.

### GetExchangeParameters

`func (o *EnvironmentTypeJobParameters) GetExchangeParameters() ExchangeEnvJobParameters`

GetExchangeParameters returns the ExchangeParameters field if non-nil, zero value otherwise.

### GetExchangeParametersOk

`func (o *EnvironmentTypeJobParameters) GetExchangeParametersOk() (*ExchangeEnvJobParameters, bool)`

GetExchangeParametersOk returns a tuple with the ExchangeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParameters

`func (o *EnvironmentTypeJobParameters) SetExchangeParameters(v ExchangeEnvJobParameters)`

SetExchangeParameters sets ExchangeParameters field to given value.

### HasExchangeParameters

`func (o *EnvironmentTypeJobParameters) HasExchangeParameters() bool`

HasExchangeParameters returns a boolean if a field has been set.

### GetHypervParameters

`func (o *EnvironmentTypeJobParameters) GetHypervParameters() HypervEnvJobParameters`

GetHypervParameters returns the HypervParameters field if non-nil, zero value otherwise.

### GetHypervParametersOk

`func (o *EnvironmentTypeJobParameters) GetHypervParametersOk() (*HypervEnvJobParameters, bool)`

GetHypervParametersOk returns a tuple with the HypervParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParameters

`func (o *EnvironmentTypeJobParameters) SetHypervParameters(v HypervEnvJobParameters)`

SetHypervParameters sets HypervParameters field to given value.

### HasHypervParameters

`func (o *EnvironmentTypeJobParameters) HasHypervParameters() bool`

HasHypervParameters returns a boolean if a field has been set.

### GetNasParameters

`func (o *EnvironmentTypeJobParameters) GetNasParameters() NasEnvJobParameters`

GetNasParameters returns the NasParameters field if non-nil, zero value otherwise.

### GetNasParametersOk

`func (o *EnvironmentTypeJobParameters) GetNasParametersOk() (*NasEnvJobParameters, bool)`

GetNasParametersOk returns a tuple with the NasParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasParameters

`func (o *EnvironmentTypeJobParameters) SetNasParameters(v NasEnvJobParameters)`

SetNasParameters sets NasParameters field to given value.

### HasNasParameters

`func (o *EnvironmentTypeJobParameters) HasNasParameters() bool`

HasNasParameters returns a boolean if a field has been set.

### GetOffice365Parameters

`func (o *EnvironmentTypeJobParameters) GetOffice365Parameters() Office365EnvJobParameters`

GetOffice365Parameters returns the Office365Parameters field if non-nil, zero value otherwise.

### GetOffice365ParametersOk

`func (o *EnvironmentTypeJobParameters) GetOffice365ParametersOk() (*Office365EnvJobParameters, bool)`

GetOffice365ParametersOk returns a tuple with the Office365Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Parameters

`func (o *EnvironmentTypeJobParameters) SetOffice365Parameters(v Office365EnvJobParameters)`

SetOffice365Parameters sets Office365Parameters field to given value.

### HasOffice365Parameters

`func (o *EnvironmentTypeJobParameters) HasOffice365Parameters() bool`

HasOffice365Parameters returns a boolean if a field has been set.

### GetOracleParameters

`func (o *EnvironmentTypeJobParameters) GetOracleParameters() OracleEnvJobParameters`

GetOracleParameters returns the OracleParameters field if non-nil, zero value otherwise.

### GetOracleParametersOk

`func (o *EnvironmentTypeJobParameters) GetOracleParametersOk() (*OracleEnvJobParameters, bool)`

GetOracleParametersOk returns a tuple with the OracleParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParameters

`func (o *EnvironmentTypeJobParameters) SetOracleParameters(v OracleEnvJobParameters)`

SetOracleParameters sets OracleParameters field to given value.

### HasOracleParameters

`func (o *EnvironmentTypeJobParameters) HasOracleParameters() bool`

HasOracleParameters returns a boolean if a field has been set.

### GetPhysicalParameters

`func (o *EnvironmentTypeJobParameters) GetPhysicalParameters() PhysicalEnvJobParameters`

GetPhysicalParameters returns the PhysicalParameters field if non-nil, zero value otherwise.

### GetPhysicalParametersOk

`func (o *EnvironmentTypeJobParameters) GetPhysicalParametersOk() (*PhysicalEnvJobParameters, bool)`

GetPhysicalParametersOk returns a tuple with the PhysicalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParameters

`func (o *EnvironmentTypeJobParameters) SetPhysicalParameters(v PhysicalEnvJobParameters)`

SetPhysicalParameters sets PhysicalParameters field to given value.

### HasPhysicalParameters

`func (o *EnvironmentTypeJobParameters) HasPhysicalParameters() bool`

HasPhysicalParameters returns a boolean if a field has been set.

### GetPureParameters

`func (o *EnvironmentTypeJobParameters) GetPureParameters() SanEnvJobParameters`

GetPureParameters returns the PureParameters field if non-nil, zero value otherwise.

### GetPureParametersOk

`func (o *EnvironmentTypeJobParameters) GetPureParametersOk() (*SanEnvJobParameters, bool)`

GetPureParametersOk returns a tuple with the PureParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParameters

`func (o *EnvironmentTypeJobParameters) SetPureParameters(v SanEnvJobParameters)`

SetPureParameters sets PureParameters field to given value.

### HasPureParameters

`func (o *EnvironmentTypeJobParameters) HasPureParameters() bool`

HasPureParameters returns a boolean if a field has been set.

### GetSqlParameters

`func (o *EnvironmentTypeJobParameters) GetSqlParameters() SqlEnvJobParameters`

GetSqlParameters returns the SqlParameters field if non-nil, zero value otherwise.

### GetSqlParametersOk

`func (o *EnvironmentTypeJobParameters) GetSqlParametersOk() (*SqlEnvJobParameters, bool)`

GetSqlParametersOk returns a tuple with the SqlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlParameters

`func (o *EnvironmentTypeJobParameters) SetSqlParameters(v SqlEnvJobParameters)`

SetSqlParameters sets SqlParameters field to given value.

### HasSqlParameters

`func (o *EnvironmentTypeJobParameters) HasSqlParameters() bool`

HasSqlParameters returns a boolean if a field has been set.

### GetVmwareParameters

`func (o *EnvironmentTypeJobParameters) GetVmwareParameters() VmwareEnvJobParameters`

GetVmwareParameters returns the VmwareParameters field if non-nil, zero value otherwise.

### GetVmwareParametersOk

`func (o *EnvironmentTypeJobParameters) GetVmwareParametersOk() (*VmwareEnvJobParameters, bool)`

GetVmwareParametersOk returns a tuple with the VmwareParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParameters

`func (o *EnvironmentTypeJobParameters) SetVmwareParameters(v VmwareEnvJobParameters)`

SetVmwareParameters sets VmwareParameters field to given value.

### HasVmwareParameters

`func (o *EnvironmentTypeJobParameters) HasVmwareParameters() bool`

HasVmwareParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


