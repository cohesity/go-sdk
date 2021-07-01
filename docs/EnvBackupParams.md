# EnvBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeBackupJobParams** | Pointer to [**ExchangeBackupJobParams**](ExchangeBackupJobParams.md) |  | [optional] 
**FileStubbingParams** | Pointer to [**FileStubbingParams**](FileStubbingParams.md) |  | [optional] 
**FileUptieringParams** | Pointer to [**FileUptieringParams**](FileUptieringParams.md) |  | [optional] 
**HypervBackupParams** | Pointer to [**HyperVBackupEnvParams**](HyperVBackupEnvParams.md) |  | [optional] 
**NasBackupParams** | Pointer to [**NasBackupParams**](NasBackupParams.md) |  | [optional] 
**NosqlBackupJobParams** | Pointer to [**NoSqlBackupJobParams**](NoSqlBackupJobParams.md) |  | [optional] 
**O365BackupParams** | Pointer to [**O365BackupEnvParams**](O365BackupEnvParams.md) |  | [optional] 
**OracleBackupJobParams** | Pointer to [**OracleBackupJobParams**](OracleBackupJobParams.md) |  | [optional] 
**OutlookBackupParams** | Pointer to [**OutlookBackupEnvParams**](OutlookBackupEnvParams.md) |  | [optional] 
**PhysicalBackupParams** | Pointer to [**PhysicalBackupEnvParams**](PhysicalBackupEnvParams.md) |  | [optional] 
**SnapshotManagerParams** | Pointer to [**SnapshotManagerParams**](SnapshotManagerParams.md) |  | [optional] 
**SqlBackupJobParams** | Pointer to [**SqlBackupJobParams**](SqlBackupJobParams.md) |  | [optional] 
**VmwareBackupParams** | Pointer to [**VMwareBackupEnvParams**](VMwareBackupEnvParams.md) |  | [optional] 

## Methods

### NewEnvBackupParams

`func NewEnvBackupParams() *EnvBackupParams`

NewEnvBackupParams instantiates a new EnvBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvBackupParamsWithDefaults

`func NewEnvBackupParamsWithDefaults() *EnvBackupParams`

NewEnvBackupParamsWithDefaults instantiates a new EnvBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeBackupJobParams

`func (o *EnvBackupParams) GetExchangeBackupJobParams() ExchangeBackupJobParams`

GetExchangeBackupJobParams returns the ExchangeBackupJobParams field if non-nil, zero value otherwise.

### GetExchangeBackupJobParamsOk

`func (o *EnvBackupParams) GetExchangeBackupJobParamsOk() (*ExchangeBackupJobParams, bool)`

GetExchangeBackupJobParamsOk returns a tuple with the ExchangeBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeBackupJobParams

`func (o *EnvBackupParams) SetExchangeBackupJobParams(v ExchangeBackupJobParams)`

SetExchangeBackupJobParams sets ExchangeBackupJobParams field to given value.

### HasExchangeBackupJobParams

`func (o *EnvBackupParams) HasExchangeBackupJobParams() bool`

HasExchangeBackupJobParams returns a boolean if a field has been set.

### GetFileStubbingParams

`func (o *EnvBackupParams) GetFileStubbingParams() FileStubbingParams`

GetFileStubbingParams returns the FileStubbingParams field if non-nil, zero value otherwise.

### GetFileStubbingParamsOk

`func (o *EnvBackupParams) GetFileStubbingParamsOk() (*FileStubbingParams, bool)`

GetFileStubbingParamsOk returns a tuple with the FileStubbingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStubbingParams

`func (o *EnvBackupParams) SetFileStubbingParams(v FileStubbingParams)`

SetFileStubbingParams sets FileStubbingParams field to given value.

### HasFileStubbingParams

`func (o *EnvBackupParams) HasFileStubbingParams() bool`

HasFileStubbingParams returns a boolean if a field has been set.

### GetFileUptieringParams

`func (o *EnvBackupParams) GetFileUptieringParams() FileUptieringParams`

GetFileUptieringParams returns the FileUptieringParams field if non-nil, zero value otherwise.

### GetFileUptieringParamsOk

`func (o *EnvBackupParams) GetFileUptieringParamsOk() (*FileUptieringParams, bool)`

GetFileUptieringParamsOk returns a tuple with the FileUptieringParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUptieringParams

`func (o *EnvBackupParams) SetFileUptieringParams(v FileUptieringParams)`

SetFileUptieringParams sets FileUptieringParams field to given value.

### HasFileUptieringParams

`func (o *EnvBackupParams) HasFileUptieringParams() bool`

HasFileUptieringParams returns a boolean if a field has been set.

### GetHypervBackupParams

`func (o *EnvBackupParams) GetHypervBackupParams() HyperVBackupEnvParams`

GetHypervBackupParams returns the HypervBackupParams field if non-nil, zero value otherwise.

### GetHypervBackupParamsOk

`func (o *EnvBackupParams) GetHypervBackupParamsOk() (*HyperVBackupEnvParams, bool)`

GetHypervBackupParamsOk returns a tuple with the HypervBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervBackupParams

`func (o *EnvBackupParams) SetHypervBackupParams(v HyperVBackupEnvParams)`

SetHypervBackupParams sets HypervBackupParams field to given value.

### HasHypervBackupParams

`func (o *EnvBackupParams) HasHypervBackupParams() bool`

HasHypervBackupParams returns a boolean if a field has been set.

### GetNasBackupParams

`func (o *EnvBackupParams) GetNasBackupParams() NasBackupParams`

GetNasBackupParams returns the NasBackupParams field if non-nil, zero value otherwise.

### GetNasBackupParamsOk

`func (o *EnvBackupParams) GetNasBackupParamsOk() (*NasBackupParams, bool)`

GetNasBackupParamsOk returns a tuple with the NasBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasBackupParams

`func (o *EnvBackupParams) SetNasBackupParams(v NasBackupParams)`

SetNasBackupParams sets NasBackupParams field to given value.

### HasNasBackupParams

`func (o *EnvBackupParams) HasNasBackupParams() bool`

HasNasBackupParams returns a boolean if a field has been set.

### GetNosqlBackupJobParams

`func (o *EnvBackupParams) GetNosqlBackupJobParams() NoSqlBackupJobParams`

GetNosqlBackupJobParams returns the NosqlBackupJobParams field if non-nil, zero value otherwise.

### GetNosqlBackupJobParamsOk

`func (o *EnvBackupParams) GetNosqlBackupJobParamsOk() (*NoSqlBackupJobParams, bool)`

GetNosqlBackupJobParamsOk returns a tuple with the NosqlBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlBackupJobParams

`func (o *EnvBackupParams) SetNosqlBackupJobParams(v NoSqlBackupJobParams)`

SetNosqlBackupJobParams sets NosqlBackupJobParams field to given value.

### HasNosqlBackupJobParams

`func (o *EnvBackupParams) HasNosqlBackupJobParams() bool`

HasNosqlBackupJobParams returns a boolean if a field has been set.

### GetO365BackupParams

`func (o *EnvBackupParams) GetO365BackupParams() O365BackupEnvParams`

GetO365BackupParams returns the O365BackupParams field if non-nil, zero value otherwise.

### GetO365BackupParamsOk

`func (o *EnvBackupParams) GetO365BackupParamsOk() (*O365BackupEnvParams, bool)`

GetO365BackupParamsOk returns a tuple with the O365BackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365BackupParams

`func (o *EnvBackupParams) SetO365BackupParams(v O365BackupEnvParams)`

SetO365BackupParams sets O365BackupParams field to given value.

### HasO365BackupParams

`func (o *EnvBackupParams) HasO365BackupParams() bool`

HasO365BackupParams returns a boolean if a field has been set.

### GetOracleBackupJobParams

`func (o *EnvBackupParams) GetOracleBackupJobParams() OracleBackupJobParams`

GetOracleBackupJobParams returns the OracleBackupJobParams field if non-nil, zero value otherwise.

### GetOracleBackupJobParamsOk

`func (o *EnvBackupParams) GetOracleBackupJobParamsOk() (*OracleBackupJobParams, bool)`

GetOracleBackupJobParamsOk returns a tuple with the OracleBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleBackupJobParams

`func (o *EnvBackupParams) SetOracleBackupJobParams(v OracleBackupJobParams)`

SetOracleBackupJobParams sets OracleBackupJobParams field to given value.

### HasOracleBackupJobParams

`func (o *EnvBackupParams) HasOracleBackupJobParams() bool`

HasOracleBackupJobParams returns a boolean if a field has been set.

### GetOutlookBackupParams

`func (o *EnvBackupParams) GetOutlookBackupParams() OutlookBackupEnvParams`

GetOutlookBackupParams returns the OutlookBackupParams field if non-nil, zero value otherwise.

### GetOutlookBackupParamsOk

`func (o *EnvBackupParams) GetOutlookBackupParamsOk() (*OutlookBackupEnvParams, bool)`

GetOutlookBackupParamsOk returns a tuple with the OutlookBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookBackupParams

`func (o *EnvBackupParams) SetOutlookBackupParams(v OutlookBackupEnvParams)`

SetOutlookBackupParams sets OutlookBackupParams field to given value.

### HasOutlookBackupParams

`func (o *EnvBackupParams) HasOutlookBackupParams() bool`

HasOutlookBackupParams returns a boolean if a field has been set.

### GetPhysicalBackupParams

`func (o *EnvBackupParams) GetPhysicalBackupParams() PhysicalBackupEnvParams`

GetPhysicalBackupParams returns the PhysicalBackupParams field if non-nil, zero value otherwise.

### GetPhysicalBackupParamsOk

`func (o *EnvBackupParams) GetPhysicalBackupParamsOk() (*PhysicalBackupEnvParams, bool)`

GetPhysicalBackupParamsOk returns a tuple with the PhysicalBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBackupParams

`func (o *EnvBackupParams) SetPhysicalBackupParams(v PhysicalBackupEnvParams)`

SetPhysicalBackupParams sets PhysicalBackupParams field to given value.

### HasPhysicalBackupParams

`func (o *EnvBackupParams) HasPhysicalBackupParams() bool`

HasPhysicalBackupParams returns a boolean if a field has been set.

### GetSnapshotManagerParams

`func (o *EnvBackupParams) GetSnapshotManagerParams() SnapshotManagerParams`

GetSnapshotManagerParams returns the SnapshotManagerParams field if non-nil, zero value otherwise.

### GetSnapshotManagerParamsOk

`func (o *EnvBackupParams) GetSnapshotManagerParamsOk() (*SnapshotManagerParams, bool)`

GetSnapshotManagerParamsOk returns a tuple with the SnapshotManagerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotManagerParams

`func (o *EnvBackupParams) SetSnapshotManagerParams(v SnapshotManagerParams)`

SetSnapshotManagerParams sets SnapshotManagerParams field to given value.

### HasSnapshotManagerParams

`func (o *EnvBackupParams) HasSnapshotManagerParams() bool`

HasSnapshotManagerParams returns a boolean if a field has been set.

### GetSqlBackupJobParams

`func (o *EnvBackupParams) GetSqlBackupJobParams() SqlBackupJobParams`

GetSqlBackupJobParams returns the SqlBackupJobParams field if non-nil, zero value otherwise.

### GetSqlBackupJobParamsOk

`func (o *EnvBackupParams) GetSqlBackupJobParamsOk() (*SqlBackupJobParams, bool)`

GetSqlBackupJobParamsOk returns a tuple with the SqlBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlBackupJobParams

`func (o *EnvBackupParams) SetSqlBackupJobParams(v SqlBackupJobParams)`

SetSqlBackupJobParams sets SqlBackupJobParams field to given value.

### HasSqlBackupJobParams

`func (o *EnvBackupParams) HasSqlBackupJobParams() bool`

HasSqlBackupJobParams returns a boolean if a field has been set.

### GetVmwareBackupParams

`func (o *EnvBackupParams) GetVmwareBackupParams() VMwareBackupEnvParams`

GetVmwareBackupParams returns the VmwareBackupParams field if non-nil, zero value otherwise.

### GetVmwareBackupParamsOk

`func (o *EnvBackupParams) GetVmwareBackupParamsOk() (*VMwareBackupEnvParams, bool)`

GetVmwareBackupParamsOk returns a tuple with the VmwareBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareBackupParams

`func (o *EnvBackupParams) SetVmwareBackupParams(v VMwareBackupEnvParams)`

SetVmwareBackupParams sets VmwareBackupParams field to given value.

### HasVmwareBackupParams

`func (o *EnvBackupParams) HasVmwareBackupParams() bool`

HasVmwareBackupParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


