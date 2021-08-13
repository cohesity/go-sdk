# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the cluster id of the new cluster. | [optional] [readonly] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the new cluster. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Name of the new cluster. | [optional] 
**Description** | Pointer to **NullableString** | Description of the cluster. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the new cluster. | [optional] [readonly] 
**LocalTenantId** | Pointer to **NullableString** | Specifies the local tenant id. Only applicable on Helios. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the globally unique tenant id. Only applicable on Helios. | [optional] [readonly] 
**RegionId** | Pointer to **NullableString** | Specifies the region id on which this cluster is present. Only applicable on Helios for DMaaS clusters. | [optional] [readonly] 
**RigelClusterParams** | Pointer to [**NullableRigelClusterConfigParams**](RigelClusterConfigParams.md) |  | [optional] 
**SwVersion** | Pointer to **NullableString** | Software version of the new cluster. | [optional] [readonly] 
**NetworkConfig** | Pointer to [**ClusterCreateNetworkConfig**](ClusterCreateNetworkConfig.md) |  | [optional] 
**ProxyServerConfig** | Pointer to [**NullableClusterProxyServerConfig**](ClusterProxyServerConfig.md) |  | [optional] 
**EnableEncryption** | Pointer to **NullableBool** | Specifies whether or not encryption is enabled. If encryption is enabled, all data on the Cluster will be encrypted. | [optional] [readonly] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Cluster) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Cluster) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncarnationId

`func (o *Cluster) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *Cluster) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *Cluster) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *Cluster) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *Cluster) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *Cluster) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Cluster) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Cluster) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Cluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Cluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *Cluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cluster) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Cluster) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Cluster) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLocalTenantId

`func (o *Cluster) GetLocalTenantId() string`

GetLocalTenantId returns the LocalTenantId field if non-nil, zero value otherwise.

### GetLocalTenantIdOk

`func (o *Cluster) GetLocalTenantIdOk() (*string, bool)`

GetLocalTenantIdOk returns a tuple with the LocalTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTenantId

`func (o *Cluster) SetLocalTenantId(v string)`

SetLocalTenantId sets LocalTenantId field to given value.

### HasLocalTenantId

`func (o *Cluster) HasLocalTenantId() bool`

HasLocalTenantId returns a boolean if a field has been set.

### SetLocalTenantIdNil

`func (o *Cluster) SetLocalTenantIdNil(b bool)`

 SetLocalTenantIdNil sets the value for LocalTenantId to be an explicit nil

### UnsetLocalTenantId
`func (o *Cluster) UnsetLocalTenantId()`

UnsetLocalTenantId ensures that no value is present for LocalTenantId, not even an explicit nil
### GetTenantId

`func (o *Cluster) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Cluster) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Cluster) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Cluster) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Cluster) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Cluster) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRegionId

`func (o *Cluster) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *Cluster) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *Cluster) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *Cluster) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *Cluster) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *Cluster) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetRigelClusterParams

`func (o *Cluster) GetRigelClusterParams() RigelClusterConfigParams`

GetRigelClusterParams returns the RigelClusterParams field if non-nil, zero value otherwise.

### GetRigelClusterParamsOk

`func (o *Cluster) GetRigelClusterParamsOk() (*RigelClusterConfigParams, bool)`

GetRigelClusterParamsOk returns a tuple with the RigelClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelClusterParams

`func (o *Cluster) SetRigelClusterParams(v RigelClusterConfigParams)`

SetRigelClusterParams sets RigelClusterParams field to given value.

### HasRigelClusterParams

`func (o *Cluster) HasRigelClusterParams() bool`

HasRigelClusterParams returns a boolean if a field has been set.

### SetRigelClusterParamsNil

`func (o *Cluster) SetRigelClusterParamsNil(b bool)`

 SetRigelClusterParamsNil sets the value for RigelClusterParams to be an explicit nil

### UnsetRigelClusterParams
`func (o *Cluster) UnsetRigelClusterParams()`

UnsetRigelClusterParams ensures that no value is present for RigelClusterParams, not even an explicit nil
### GetSwVersion

`func (o *Cluster) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *Cluster) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *Cluster) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *Cluster) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### SetSwVersionNil

`func (o *Cluster) SetSwVersionNil(b bool)`

 SetSwVersionNil sets the value for SwVersion to be an explicit nil

### UnsetSwVersion
`func (o *Cluster) UnsetSwVersion()`

UnsetSwVersion ensures that no value is present for SwVersion, not even an explicit nil
### GetNetworkConfig

`func (o *Cluster) GetNetworkConfig() ClusterCreateNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *Cluster) GetNetworkConfigOk() (*ClusterCreateNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *Cluster) SetNetworkConfig(v ClusterCreateNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *Cluster) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### GetProxyServerConfig

`func (o *Cluster) GetProxyServerConfig() ClusterProxyServerConfig`

GetProxyServerConfig returns the ProxyServerConfig field if non-nil, zero value otherwise.

### GetProxyServerConfigOk

`func (o *Cluster) GetProxyServerConfigOk() (*ClusterProxyServerConfig, bool)`

GetProxyServerConfigOk returns a tuple with the ProxyServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServerConfig

`func (o *Cluster) SetProxyServerConfig(v ClusterProxyServerConfig)`

SetProxyServerConfig sets ProxyServerConfig field to given value.

### HasProxyServerConfig

`func (o *Cluster) HasProxyServerConfig() bool`

HasProxyServerConfig returns a boolean if a field has been set.

### SetProxyServerConfigNil

`func (o *Cluster) SetProxyServerConfigNil(b bool)`

 SetProxyServerConfigNil sets the value for ProxyServerConfig to be an explicit nil

### UnsetProxyServerConfig
`func (o *Cluster) UnsetProxyServerConfig()`

UnsetProxyServerConfig ensures that no value is present for ProxyServerConfig, not even an explicit nil
### GetEnableEncryption

`func (o *Cluster) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *Cluster) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *Cluster) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.

### HasEnableEncryption

`func (o *Cluster) HasEnableEncryption() bool`

HasEnableEncryption returns a boolean if a field has been set.

### SetEnableEncryptionNil

`func (o *Cluster) SetEnableEncryptionNil(b bool)`

 SetEnableEncryptionNil sets the value for EnableEncryption to be an explicit nil

### UnsetEnableEncryption
`func (o *Cluster) UnsetEnableEncryption()`

UnsetEnableEncryption ensures that no value is present for EnableEncryption, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


