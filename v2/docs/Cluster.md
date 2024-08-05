# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAuditLogConfig** | Pointer to [**ClusterAuditLogConfig**](ClusterAuditLogConfig.md) |  | [optional] 
**ClusterDeploymentType** | Pointer to **NullableString** | Type of Cluster Deployment. | [optional] 
**ClusterSize** | Pointer to **NullableString** | Specifies the size of the cloud platforms. | [optional] [readonly] 
**CohesionClusterParams** | Pointer to [**CohesionClusterConfigParams**](CohesionClusterConfigParams.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Description of the cluster. | [optional] 
**EnableEncryption** | Pointer to **NullableBool** | Specifies whether or not encryption is enabled. If encryption is enabled, all data on the Cluster will be encrypted. | [optional] [readonly] 
**FileServicesAuditLogConfig** | Pointer to [**AuditLogConfig**](AuditLogConfig.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the cluster id of the new cluster. | [optional] [readonly] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the new cluster. | [optional] [readonly] 
**LocalTenantId** | Pointer to **NullableString** | Specifies the local tenant id. Only applicable on Helios. | [optional] [readonly] 
**Metadata** | Pointer to [**ClusterMetadataRequest**](ClusterMetadataRequest.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name of the new cluster. | [optional] 
**NetworkConfig** | Pointer to [**ClusterCreateNetworkConfig**](ClusterCreateNetworkConfig.md) |  | [optional] 
**ProxyServerConfig** | Pointer to [**NullableClusterProxyServerConfig**](ClusterProxyServerConfig.md) |  | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id on which this cluster is present. Only applicable on Helios for DMaaS clusters. | [optional] [readonly] 
**RigelClusterParams** | Pointer to [**NullableRigelClusterConfigParams**](RigelClusterConfigParams.md) |  | [optional] 
**SwVersion** | Pointer to **NullableString** | Software version of the new cluster. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the globally unique tenant id. Only applicable on Helios. | [optional] [readonly] 
**TieringAuditLogConfig** | Pointer to [**AuditLogConfig**](AuditLogConfig.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the new cluster. | [optional] [readonly] 
**ViewsGlobalSettings** | Pointer to [**ViewsGlobalSettings**](ViewsGlobalSettings.md) |  | [optional] 

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

### GetClusterAuditLogConfig

`func (o *Cluster) GetClusterAuditLogConfig() ClusterAuditLogConfig`

GetClusterAuditLogConfig returns the ClusterAuditLogConfig field if non-nil, zero value otherwise.

### GetClusterAuditLogConfigOk

`func (o *Cluster) GetClusterAuditLogConfigOk() (*ClusterAuditLogConfig, bool)`

GetClusterAuditLogConfigOk returns a tuple with the ClusterAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAuditLogConfig

`func (o *Cluster) SetClusterAuditLogConfig(v ClusterAuditLogConfig)`

SetClusterAuditLogConfig sets ClusterAuditLogConfig field to given value.

### HasClusterAuditLogConfig

`func (o *Cluster) HasClusterAuditLogConfig() bool`

HasClusterAuditLogConfig returns a boolean if a field has been set.

### GetClusterDeploymentType

`func (o *Cluster) GetClusterDeploymentType() string`

GetClusterDeploymentType returns the ClusterDeploymentType field if non-nil, zero value otherwise.

### GetClusterDeploymentTypeOk

`func (o *Cluster) GetClusterDeploymentTypeOk() (*string, bool)`

GetClusterDeploymentTypeOk returns a tuple with the ClusterDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDeploymentType

`func (o *Cluster) SetClusterDeploymentType(v string)`

SetClusterDeploymentType sets ClusterDeploymentType field to given value.

### HasClusterDeploymentType

`func (o *Cluster) HasClusterDeploymentType() bool`

HasClusterDeploymentType returns a boolean if a field has been set.

### SetClusterDeploymentTypeNil

`func (o *Cluster) SetClusterDeploymentTypeNil(b bool)`

 SetClusterDeploymentTypeNil sets the value for ClusterDeploymentType to be an explicit nil

### UnsetClusterDeploymentType
`func (o *Cluster) UnsetClusterDeploymentType()`

UnsetClusterDeploymentType ensures that no value is present for ClusterDeploymentType, not even an explicit nil
### GetClusterSize

`func (o *Cluster) GetClusterSize() string`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *Cluster) GetClusterSizeOk() (*string, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *Cluster) SetClusterSize(v string)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *Cluster) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### SetClusterSizeNil

`func (o *Cluster) SetClusterSizeNil(b bool)`

 SetClusterSizeNil sets the value for ClusterSize to be an explicit nil

### UnsetClusterSize
`func (o *Cluster) UnsetClusterSize()`

UnsetClusterSize ensures that no value is present for ClusterSize, not even an explicit nil
### GetCohesionClusterParams

`func (o *Cluster) GetCohesionClusterParams() CohesionClusterConfigParams`

GetCohesionClusterParams returns the CohesionClusterParams field if non-nil, zero value otherwise.

### GetCohesionClusterParamsOk

`func (o *Cluster) GetCohesionClusterParamsOk() (*CohesionClusterConfigParams, bool)`

GetCohesionClusterParamsOk returns a tuple with the CohesionClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesionClusterParams

`func (o *Cluster) SetCohesionClusterParams(v CohesionClusterConfigParams)`

SetCohesionClusterParams sets CohesionClusterParams field to given value.

### HasCohesionClusterParams

`func (o *Cluster) HasCohesionClusterParams() bool`

HasCohesionClusterParams returns a boolean if a field has been set.

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
### GetFileServicesAuditLogConfig

`func (o *Cluster) GetFileServicesAuditLogConfig() AuditLogConfig`

GetFileServicesAuditLogConfig returns the FileServicesAuditLogConfig field if non-nil, zero value otherwise.

### GetFileServicesAuditLogConfigOk

`func (o *Cluster) GetFileServicesAuditLogConfigOk() (*AuditLogConfig, bool)`

GetFileServicesAuditLogConfigOk returns a tuple with the FileServicesAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesAuditLogConfig

`func (o *Cluster) SetFileServicesAuditLogConfig(v AuditLogConfig)`

SetFileServicesAuditLogConfig sets FileServicesAuditLogConfig field to given value.

### HasFileServicesAuditLogConfig

`func (o *Cluster) HasFileServicesAuditLogConfig() bool`

HasFileServicesAuditLogConfig returns a boolean if a field has been set.

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
### GetMetadata

`func (o *Cluster) GetMetadata() ClusterMetadataRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Cluster) GetMetadataOk() (*ClusterMetadataRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Cluster) SetMetadata(v ClusterMetadataRequest)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Cluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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
### GetTieringAuditLogConfig

`func (o *Cluster) GetTieringAuditLogConfig() AuditLogConfig`

GetTieringAuditLogConfig returns the TieringAuditLogConfig field if non-nil, zero value otherwise.

### GetTieringAuditLogConfigOk

`func (o *Cluster) GetTieringAuditLogConfigOk() (*AuditLogConfig, bool)`

GetTieringAuditLogConfigOk returns a tuple with the TieringAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringAuditLogConfig

`func (o *Cluster) SetTieringAuditLogConfig(v AuditLogConfig)`

SetTieringAuditLogConfig sets TieringAuditLogConfig field to given value.

### HasTieringAuditLogConfig

`func (o *Cluster) HasTieringAuditLogConfig() bool`

HasTieringAuditLogConfig returns a boolean if a field has been set.

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
### GetViewsGlobalSettings

`func (o *Cluster) GetViewsGlobalSettings() ViewsGlobalSettings`

GetViewsGlobalSettings returns the ViewsGlobalSettings field if non-nil, zero value otherwise.

### GetViewsGlobalSettingsOk

`func (o *Cluster) GetViewsGlobalSettingsOk() (*ViewsGlobalSettings, bool)`

GetViewsGlobalSettingsOk returns a tuple with the ViewsGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsGlobalSettings

`func (o *Cluster) SetViewsGlobalSettings(v ViewsGlobalSettings)`

SetViewsGlobalSettings sets ViewsGlobalSettings field to given value.

### HasViewsGlobalSettings

`func (o *Cluster) HasViewsGlobalSettings() bool`

HasViewsGlobalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


