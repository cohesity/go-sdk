# VaultProviderStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRate** | Pointer to **NullableInt64** | Specifies the relative change of size of entities on the vault. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name. | [optional] 
**ReadBandwidth** | Pointer to **NullableInt64** | Specifies the average read bandwidth over the last 24 hours. | [optional] 
**StatsByEnv** | Pointer to [**[]VaultProviderStatsByEnv**](VaultProviderStatsByEnv.md) | Specifies the stats by environments. | [optional] 
**VaultGroup** | Pointer to **NullableString** | Specifies the cloud vendor type. | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the Vault id. | [optional] 
**VaultType** | Pointer to **NullableString** | Specifies the External Target type. | [optional] 
**Vaultname** | Pointer to **NullableString** | Specifies the Vault name. | [optional] 
**WriteBandwidth** | Pointer to **NullableInt64** | Specifies the average write bandwidth over the last 24 hours. | [optional] 

## Methods

### NewVaultProviderStatsInfo

`func NewVaultProviderStatsInfo() *VaultProviderStatsInfo`

NewVaultProviderStatsInfo instantiates a new VaultProviderStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultProviderStatsInfoWithDefaults

`func NewVaultProviderStatsInfoWithDefaults() *VaultProviderStatsInfo`

NewVaultProviderStatsInfoWithDefaults instantiates a new VaultProviderStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRate

`func (o *VaultProviderStatsInfo) GetChangeRate() int64`

GetChangeRate returns the ChangeRate field if non-nil, zero value otherwise.

### GetChangeRateOk

`func (o *VaultProviderStatsInfo) GetChangeRateOk() (*int64, bool)`

GetChangeRateOk returns a tuple with the ChangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRate

`func (o *VaultProviderStatsInfo) SetChangeRate(v int64)`

SetChangeRate sets ChangeRate field to given value.

### HasChangeRate

`func (o *VaultProviderStatsInfo) HasChangeRate() bool`

HasChangeRate returns a boolean if a field has been set.

### SetChangeRateNil

`func (o *VaultProviderStatsInfo) SetChangeRateNil(b bool)`

 SetChangeRateNil sets the value for ChangeRate to be an explicit nil

### UnsetChangeRate
`func (o *VaultProviderStatsInfo) UnsetChangeRate()`

UnsetChangeRate ensures that no value is present for ChangeRate, not even an explicit nil
### GetClusterId

`func (o *VaultProviderStatsInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *VaultProviderStatsInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *VaultProviderStatsInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *VaultProviderStatsInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *VaultProviderStatsInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *VaultProviderStatsInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *VaultProviderStatsInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *VaultProviderStatsInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *VaultProviderStatsInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *VaultProviderStatsInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *VaultProviderStatsInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *VaultProviderStatsInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *VaultProviderStatsInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VaultProviderStatsInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VaultProviderStatsInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VaultProviderStatsInfo) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *VaultProviderStatsInfo) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *VaultProviderStatsInfo) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetReadBandwidth

`func (o *VaultProviderStatsInfo) GetReadBandwidth() int64`

GetReadBandwidth returns the ReadBandwidth field if non-nil, zero value otherwise.

### GetReadBandwidthOk

`func (o *VaultProviderStatsInfo) GetReadBandwidthOk() (*int64, bool)`

GetReadBandwidthOk returns a tuple with the ReadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidth

`func (o *VaultProviderStatsInfo) SetReadBandwidth(v int64)`

SetReadBandwidth sets ReadBandwidth field to given value.

### HasReadBandwidth

`func (o *VaultProviderStatsInfo) HasReadBandwidth() bool`

HasReadBandwidth returns a boolean if a field has been set.

### SetReadBandwidthNil

`func (o *VaultProviderStatsInfo) SetReadBandwidthNil(b bool)`

 SetReadBandwidthNil sets the value for ReadBandwidth to be an explicit nil

### UnsetReadBandwidth
`func (o *VaultProviderStatsInfo) UnsetReadBandwidth()`

UnsetReadBandwidth ensures that no value is present for ReadBandwidth, not even an explicit nil
### GetStatsByEnv

`func (o *VaultProviderStatsInfo) GetStatsByEnv() []VaultProviderStatsByEnv`

GetStatsByEnv returns the StatsByEnv field if non-nil, zero value otherwise.

### GetStatsByEnvOk

`func (o *VaultProviderStatsInfo) GetStatsByEnvOk() (*[]VaultProviderStatsByEnv, bool)`

GetStatsByEnvOk returns a tuple with the StatsByEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnv

`func (o *VaultProviderStatsInfo) SetStatsByEnv(v []VaultProviderStatsByEnv)`

SetStatsByEnv sets StatsByEnv field to given value.

### HasStatsByEnv

`func (o *VaultProviderStatsInfo) HasStatsByEnv() bool`

HasStatsByEnv returns a boolean if a field has been set.

### GetVaultGroup

`func (o *VaultProviderStatsInfo) GetVaultGroup() string`

GetVaultGroup returns the VaultGroup field if non-nil, zero value otherwise.

### GetVaultGroupOk

`func (o *VaultProviderStatsInfo) GetVaultGroupOk() (*string, bool)`

GetVaultGroupOk returns a tuple with the VaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultGroup

`func (o *VaultProviderStatsInfo) SetVaultGroup(v string)`

SetVaultGroup sets VaultGroup field to given value.

### HasVaultGroup

`func (o *VaultProviderStatsInfo) HasVaultGroup() bool`

HasVaultGroup returns a boolean if a field has been set.

### SetVaultGroupNil

`func (o *VaultProviderStatsInfo) SetVaultGroupNil(b bool)`

 SetVaultGroupNil sets the value for VaultGroup to be an explicit nil

### UnsetVaultGroup
`func (o *VaultProviderStatsInfo) UnsetVaultGroup()`

UnsetVaultGroup ensures that no value is present for VaultGroup, not even an explicit nil
### GetVaultId

`func (o *VaultProviderStatsInfo) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *VaultProviderStatsInfo) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *VaultProviderStatsInfo) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *VaultProviderStatsInfo) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *VaultProviderStatsInfo) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *VaultProviderStatsInfo) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
### GetVaultType

`func (o *VaultProviderStatsInfo) GetVaultType() string`

GetVaultType returns the VaultType field if non-nil, zero value otherwise.

### GetVaultTypeOk

`func (o *VaultProviderStatsInfo) GetVaultTypeOk() (*string, bool)`

GetVaultTypeOk returns a tuple with the VaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultType

`func (o *VaultProviderStatsInfo) SetVaultType(v string)`

SetVaultType sets VaultType field to given value.

### HasVaultType

`func (o *VaultProviderStatsInfo) HasVaultType() bool`

HasVaultType returns a boolean if a field has been set.

### SetVaultTypeNil

`func (o *VaultProviderStatsInfo) SetVaultTypeNil(b bool)`

 SetVaultTypeNil sets the value for VaultType to be an explicit nil

### UnsetVaultType
`func (o *VaultProviderStatsInfo) UnsetVaultType()`

UnsetVaultType ensures that no value is present for VaultType, not even an explicit nil
### GetVaultname

`func (o *VaultProviderStatsInfo) GetVaultname() string`

GetVaultname returns the Vaultname field if non-nil, zero value otherwise.

### GetVaultnameOk

`func (o *VaultProviderStatsInfo) GetVaultnameOk() (*string, bool)`

GetVaultnameOk returns a tuple with the Vaultname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultname

`func (o *VaultProviderStatsInfo) SetVaultname(v string)`

SetVaultname sets Vaultname field to given value.

### HasVaultname

`func (o *VaultProviderStatsInfo) HasVaultname() bool`

HasVaultname returns a boolean if a field has been set.

### SetVaultnameNil

`func (o *VaultProviderStatsInfo) SetVaultnameNil(b bool)`

 SetVaultnameNil sets the value for Vaultname to be an explicit nil

### UnsetVaultname
`func (o *VaultProviderStatsInfo) UnsetVaultname()`

UnsetVaultname ensures that no value is present for Vaultname, not even an explicit nil
### GetWriteBandwidth

`func (o *VaultProviderStatsInfo) GetWriteBandwidth() int64`

GetWriteBandwidth returns the WriteBandwidth field if non-nil, zero value otherwise.

### GetWriteBandwidthOk

`func (o *VaultProviderStatsInfo) GetWriteBandwidthOk() (*int64, bool)`

GetWriteBandwidthOk returns a tuple with the WriteBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidth

`func (o *VaultProviderStatsInfo) SetWriteBandwidth(v int64)`

SetWriteBandwidth sets WriteBandwidth field to given value.

### HasWriteBandwidth

`func (o *VaultProviderStatsInfo) HasWriteBandwidth() bool`

HasWriteBandwidth returns a boolean if a field has been set.

### SetWriteBandwidthNil

`func (o *VaultProviderStatsInfo) SetWriteBandwidthNil(b bool)`

 SetWriteBandwidthNil sets the value for WriteBandwidth to be an explicit nil

### UnsetWriteBandwidth
`func (o *VaultProviderStatsInfo) UnsetWriteBandwidth()`

UnsetWriteBandwidth ensures that no value is present for WriteBandwidth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


