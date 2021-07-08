# VaultStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on AWS vaults. | [optional] 
**AzureUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on Azure vaults. | [optional] 
**GcpUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on GCP vaults. | [optional] 
**NasUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on NAS vaults. | [optional] 
**OracleUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on Oracle vaults. | [optional] 
**QstarUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on QStar Tape vaults. | [optional] 
**S3cUsageBytes** | Pointer to **NullableInt64** | Specifies the usage on S3 Compatible vaults. | [optional] 
**VaultStatsList** | Pointer to [**[]VaultStatsInfo**](VaultStatsInfo.md) | Specifies the stats of all vaults on the cluster. | [optional] 

## Methods

### NewVaultStats

`func NewVaultStats() *VaultStats`

NewVaultStats instantiates a new VaultStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultStatsWithDefaults

`func NewVaultStatsWithDefaults() *VaultStats`

NewVaultStatsWithDefaults instantiates a new VaultStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsUsageBytes

`func (o *VaultStats) GetAwsUsageBytes() int64`

GetAwsUsageBytes returns the AwsUsageBytes field if non-nil, zero value otherwise.

### GetAwsUsageBytesOk

`func (o *VaultStats) GetAwsUsageBytesOk() (*int64, bool)`

GetAwsUsageBytesOk returns a tuple with the AwsUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUsageBytes

`func (o *VaultStats) SetAwsUsageBytes(v int64)`

SetAwsUsageBytes sets AwsUsageBytes field to given value.

### HasAwsUsageBytes

`func (o *VaultStats) HasAwsUsageBytes() bool`

HasAwsUsageBytes returns a boolean if a field has been set.

### SetAwsUsageBytesNil

`func (o *VaultStats) SetAwsUsageBytesNil(b bool)`

 SetAwsUsageBytesNil sets the value for AwsUsageBytes to be an explicit nil

### UnsetAwsUsageBytes
`func (o *VaultStats) UnsetAwsUsageBytes()`

UnsetAwsUsageBytes ensures that no value is present for AwsUsageBytes, not even an explicit nil
### GetAzureUsageBytes

`func (o *VaultStats) GetAzureUsageBytes() int64`

GetAzureUsageBytes returns the AzureUsageBytes field if non-nil, zero value otherwise.

### GetAzureUsageBytesOk

`func (o *VaultStats) GetAzureUsageBytesOk() (*int64, bool)`

GetAzureUsageBytesOk returns a tuple with the AzureUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureUsageBytes

`func (o *VaultStats) SetAzureUsageBytes(v int64)`

SetAzureUsageBytes sets AzureUsageBytes field to given value.

### HasAzureUsageBytes

`func (o *VaultStats) HasAzureUsageBytes() bool`

HasAzureUsageBytes returns a boolean if a field has been set.

### SetAzureUsageBytesNil

`func (o *VaultStats) SetAzureUsageBytesNil(b bool)`

 SetAzureUsageBytesNil sets the value for AzureUsageBytes to be an explicit nil

### UnsetAzureUsageBytes
`func (o *VaultStats) UnsetAzureUsageBytes()`

UnsetAzureUsageBytes ensures that no value is present for AzureUsageBytes, not even an explicit nil
### GetGcpUsageBytes

`func (o *VaultStats) GetGcpUsageBytes() int64`

GetGcpUsageBytes returns the GcpUsageBytes field if non-nil, zero value otherwise.

### GetGcpUsageBytesOk

`func (o *VaultStats) GetGcpUsageBytesOk() (*int64, bool)`

GetGcpUsageBytesOk returns a tuple with the GcpUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpUsageBytes

`func (o *VaultStats) SetGcpUsageBytes(v int64)`

SetGcpUsageBytes sets GcpUsageBytes field to given value.

### HasGcpUsageBytes

`func (o *VaultStats) HasGcpUsageBytes() bool`

HasGcpUsageBytes returns a boolean if a field has been set.

### SetGcpUsageBytesNil

`func (o *VaultStats) SetGcpUsageBytesNil(b bool)`

 SetGcpUsageBytesNil sets the value for GcpUsageBytes to be an explicit nil

### UnsetGcpUsageBytes
`func (o *VaultStats) UnsetGcpUsageBytes()`

UnsetGcpUsageBytes ensures that no value is present for GcpUsageBytes, not even an explicit nil
### GetNasUsageBytes

`func (o *VaultStats) GetNasUsageBytes() int64`

GetNasUsageBytes returns the NasUsageBytes field if non-nil, zero value otherwise.

### GetNasUsageBytesOk

`func (o *VaultStats) GetNasUsageBytesOk() (*int64, bool)`

GetNasUsageBytesOk returns a tuple with the NasUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasUsageBytes

`func (o *VaultStats) SetNasUsageBytes(v int64)`

SetNasUsageBytes sets NasUsageBytes field to given value.

### HasNasUsageBytes

`func (o *VaultStats) HasNasUsageBytes() bool`

HasNasUsageBytes returns a boolean if a field has been set.

### SetNasUsageBytesNil

`func (o *VaultStats) SetNasUsageBytesNil(b bool)`

 SetNasUsageBytesNil sets the value for NasUsageBytes to be an explicit nil

### UnsetNasUsageBytes
`func (o *VaultStats) UnsetNasUsageBytes()`

UnsetNasUsageBytes ensures that no value is present for NasUsageBytes, not even an explicit nil
### GetOracleUsageBytes

`func (o *VaultStats) GetOracleUsageBytes() int64`

GetOracleUsageBytes returns the OracleUsageBytes field if non-nil, zero value otherwise.

### GetOracleUsageBytesOk

`func (o *VaultStats) GetOracleUsageBytesOk() (*int64, bool)`

GetOracleUsageBytesOk returns a tuple with the OracleUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUsageBytes

`func (o *VaultStats) SetOracleUsageBytes(v int64)`

SetOracleUsageBytes sets OracleUsageBytes field to given value.

### HasOracleUsageBytes

`func (o *VaultStats) HasOracleUsageBytes() bool`

HasOracleUsageBytes returns a boolean if a field has been set.

### SetOracleUsageBytesNil

`func (o *VaultStats) SetOracleUsageBytesNil(b bool)`

 SetOracleUsageBytesNil sets the value for OracleUsageBytes to be an explicit nil

### UnsetOracleUsageBytes
`func (o *VaultStats) UnsetOracleUsageBytes()`

UnsetOracleUsageBytes ensures that no value is present for OracleUsageBytes, not even an explicit nil
### GetQstarUsageBytes

`func (o *VaultStats) GetQstarUsageBytes() int64`

GetQstarUsageBytes returns the QstarUsageBytes field if non-nil, zero value otherwise.

### GetQstarUsageBytesOk

`func (o *VaultStats) GetQstarUsageBytesOk() (*int64, bool)`

GetQstarUsageBytesOk returns a tuple with the QstarUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQstarUsageBytes

`func (o *VaultStats) SetQstarUsageBytes(v int64)`

SetQstarUsageBytes sets QstarUsageBytes field to given value.

### HasQstarUsageBytes

`func (o *VaultStats) HasQstarUsageBytes() bool`

HasQstarUsageBytes returns a boolean if a field has been set.

### SetQstarUsageBytesNil

`func (o *VaultStats) SetQstarUsageBytesNil(b bool)`

 SetQstarUsageBytesNil sets the value for QstarUsageBytes to be an explicit nil

### UnsetQstarUsageBytes
`func (o *VaultStats) UnsetQstarUsageBytes()`

UnsetQstarUsageBytes ensures that no value is present for QstarUsageBytes, not even an explicit nil
### GetS3cUsageBytes

`func (o *VaultStats) GetS3cUsageBytes() int64`

GetS3cUsageBytes returns the S3cUsageBytes field if non-nil, zero value otherwise.

### GetS3cUsageBytesOk

`func (o *VaultStats) GetS3cUsageBytesOk() (*int64, bool)`

GetS3cUsageBytesOk returns a tuple with the S3cUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3cUsageBytes

`func (o *VaultStats) SetS3cUsageBytes(v int64)`

SetS3cUsageBytes sets S3cUsageBytes field to given value.

### HasS3cUsageBytes

`func (o *VaultStats) HasS3cUsageBytes() bool`

HasS3cUsageBytes returns a boolean if a field has been set.

### SetS3cUsageBytesNil

`func (o *VaultStats) SetS3cUsageBytesNil(b bool)`

 SetS3cUsageBytesNil sets the value for S3cUsageBytes to be an explicit nil

### UnsetS3cUsageBytes
`func (o *VaultStats) UnsetS3cUsageBytes()`

UnsetS3cUsageBytes ensures that no value is present for S3cUsageBytes, not even an explicit nil
### GetVaultStatsList

`func (o *VaultStats) GetVaultStatsList() []VaultStatsInfo`

GetVaultStatsList returns the VaultStatsList field if non-nil, zero value otherwise.

### GetVaultStatsListOk

`func (o *VaultStats) GetVaultStatsListOk() (*[]VaultStatsInfo, bool)`

GetVaultStatsListOk returns a tuple with the VaultStatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultStatsList

`func (o *VaultStats) SetVaultStatsList(v []VaultStatsInfo)`

SetVaultStatsList sets VaultStatsList field to given value.

### HasVaultStatsList

`func (o *VaultStats) HasVaultStatsList() bool`

HasVaultStatsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


