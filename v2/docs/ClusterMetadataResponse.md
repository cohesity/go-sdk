# ClusterMetadataResponse

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

### NewClusterMetadataResponse

`func NewClusterMetadataResponse() *ClusterMetadataResponse`

NewClusterMetadataResponse instantiates a new ClusterMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetadataResponseWithDefaults

`func NewClusterMetadataResponseWithDefaults() *ClusterMetadataResponse`

NewClusterMetadataResponseWithDefaults instantiates a new ClusterMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAuditLogConfig

`func (o *ClusterMetadataResponse) GetClusterAuditLogConfig() ClusterAuditLogConfig`

GetClusterAuditLogConfig returns the ClusterAuditLogConfig field if non-nil, zero value otherwise.

### GetClusterAuditLogConfigOk

`func (o *ClusterMetadataResponse) GetClusterAuditLogConfigOk() (*ClusterAuditLogConfig, bool)`

GetClusterAuditLogConfigOk returns a tuple with the ClusterAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAuditLogConfig

`func (o *ClusterMetadataResponse) SetClusterAuditLogConfig(v ClusterAuditLogConfig)`

SetClusterAuditLogConfig sets ClusterAuditLogConfig field to given value.

### HasClusterAuditLogConfig

`func (o *ClusterMetadataResponse) HasClusterAuditLogConfig() bool`

HasClusterAuditLogConfig returns a boolean if a field has been set.

### GetClusterDeploymentType

`func (o *ClusterMetadataResponse) GetClusterDeploymentType() string`

GetClusterDeploymentType returns the ClusterDeploymentType field if non-nil, zero value otherwise.

### GetClusterDeploymentTypeOk

`func (o *ClusterMetadataResponse) GetClusterDeploymentTypeOk() (*string, bool)`

GetClusterDeploymentTypeOk returns a tuple with the ClusterDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDeploymentType

`func (o *ClusterMetadataResponse) SetClusterDeploymentType(v string)`

SetClusterDeploymentType sets ClusterDeploymentType field to given value.

### HasClusterDeploymentType

`func (o *ClusterMetadataResponse) HasClusterDeploymentType() bool`

HasClusterDeploymentType returns a boolean if a field has been set.

### SetClusterDeploymentTypeNil

`func (o *ClusterMetadataResponse) SetClusterDeploymentTypeNil(b bool)`

 SetClusterDeploymentTypeNil sets the value for ClusterDeploymentType to be an explicit nil

### UnsetClusterDeploymentType
`func (o *ClusterMetadataResponse) UnsetClusterDeploymentType()`

UnsetClusterDeploymentType ensures that no value is present for ClusterDeploymentType, not even an explicit nil
### GetClusterSize

`func (o *ClusterMetadataResponse) GetClusterSize() string`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *ClusterMetadataResponse) GetClusterSizeOk() (*string, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *ClusterMetadataResponse) SetClusterSize(v string)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *ClusterMetadataResponse) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### SetClusterSizeNil

`func (o *ClusterMetadataResponse) SetClusterSizeNil(b bool)`

 SetClusterSizeNil sets the value for ClusterSize to be an explicit nil

### UnsetClusterSize
`func (o *ClusterMetadataResponse) UnsetClusterSize()`

UnsetClusterSize ensures that no value is present for ClusterSize, not even an explicit nil
### GetCohesionClusterParams

`func (o *ClusterMetadataResponse) GetCohesionClusterParams() CohesionClusterConfigParams`

GetCohesionClusterParams returns the CohesionClusterParams field if non-nil, zero value otherwise.

### GetCohesionClusterParamsOk

`func (o *ClusterMetadataResponse) GetCohesionClusterParamsOk() (*CohesionClusterConfigParams, bool)`

GetCohesionClusterParamsOk returns a tuple with the CohesionClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesionClusterParams

`func (o *ClusterMetadataResponse) SetCohesionClusterParams(v CohesionClusterConfigParams)`

SetCohesionClusterParams sets CohesionClusterParams field to given value.

### HasCohesionClusterParams

`func (o *ClusterMetadataResponse) HasCohesionClusterParams() bool`

HasCohesionClusterParams returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterMetadataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterMetadataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterMetadataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterMetadataResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterMetadataResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterMetadataResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableEncryption

`func (o *ClusterMetadataResponse) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *ClusterMetadataResponse) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *ClusterMetadataResponse) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.

### HasEnableEncryption

`func (o *ClusterMetadataResponse) HasEnableEncryption() bool`

HasEnableEncryption returns a boolean if a field has been set.

### SetEnableEncryptionNil

`func (o *ClusterMetadataResponse) SetEnableEncryptionNil(b bool)`

 SetEnableEncryptionNil sets the value for EnableEncryption to be an explicit nil

### UnsetEnableEncryption
`func (o *ClusterMetadataResponse) UnsetEnableEncryption()`

UnsetEnableEncryption ensures that no value is present for EnableEncryption, not even an explicit nil
### GetFileServicesAuditLogConfig

`func (o *ClusterMetadataResponse) GetFileServicesAuditLogConfig() AuditLogConfig`

GetFileServicesAuditLogConfig returns the FileServicesAuditLogConfig field if non-nil, zero value otherwise.

### GetFileServicesAuditLogConfigOk

`func (o *ClusterMetadataResponse) GetFileServicesAuditLogConfigOk() (*AuditLogConfig, bool)`

GetFileServicesAuditLogConfigOk returns a tuple with the FileServicesAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesAuditLogConfig

`func (o *ClusterMetadataResponse) SetFileServicesAuditLogConfig(v AuditLogConfig)`

SetFileServicesAuditLogConfig sets FileServicesAuditLogConfig field to given value.

### HasFileServicesAuditLogConfig

`func (o *ClusterMetadataResponse) HasFileServicesAuditLogConfig() bool`

HasFileServicesAuditLogConfig returns a boolean if a field has been set.

### GetId

`func (o *ClusterMetadataResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterMetadataResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterMetadataResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterMetadataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClusterMetadataResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClusterMetadataResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncarnationId

`func (o *ClusterMetadataResponse) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *ClusterMetadataResponse) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *ClusterMetadataResponse) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *ClusterMetadataResponse) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *ClusterMetadataResponse) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *ClusterMetadataResponse) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
### GetLocalTenantId

`func (o *ClusterMetadataResponse) GetLocalTenantId() string`

GetLocalTenantId returns the LocalTenantId field if non-nil, zero value otherwise.

### GetLocalTenantIdOk

`func (o *ClusterMetadataResponse) GetLocalTenantIdOk() (*string, bool)`

GetLocalTenantIdOk returns a tuple with the LocalTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTenantId

`func (o *ClusterMetadataResponse) SetLocalTenantId(v string)`

SetLocalTenantId sets LocalTenantId field to given value.

### HasLocalTenantId

`func (o *ClusterMetadataResponse) HasLocalTenantId() bool`

HasLocalTenantId returns a boolean if a field has been set.

### SetLocalTenantIdNil

`func (o *ClusterMetadataResponse) SetLocalTenantIdNil(b bool)`

 SetLocalTenantIdNil sets the value for LocalTenantId to be an explicit nil

### UnsetLocalTenantId
`func (o *ClusterMetadataResponse) UnsetLocalTenantId()`

UnsetLocalTenantId ensures that no value is present for LocalTenantId, not even an explicit nil
### GetMetadata

`func (o *ClusterMetadataResponse) GetMetadata() ClusterMetadataRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterMetadataResponse) GetMetadataOk() (*ClusterMetadataRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterMetadataResponse) SetMetadata(v ClusterMetadataRequest)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterMetadataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ClusterMetadataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterMetadataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterMetadataResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterMetadataResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClusterMetadataResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClusterMetadataResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkConfig

`func (o *ClusterMetadataResponse) GetNetworkConfig() ClusterCreateNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *ClusterMetadataResponse) GetNetworkConfigOk() (*ClusterCreateNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *ClusterMetadataResponse) SetNetworkConfig(v ClusterCreateNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *ClusterMetadataResponse) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### GetProxyServerConfig

`func (o *ClusterMetadataResponse) GetProxyServerConfig() ClusterProxyServerConfig`

GetProxyServerConfig returns the ProxyServerConfig field if non-nil, zero value otherwise.

### GetProxyServerConfigOk

`func (o *ClusterMetadataResponse) GetProxyServerConfigOk() (*ClusterProxyServerConfig, bool)`

GetProxyServerConfigOk returns a tuple with the ProxyServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServerConfig

`func (o *ClusterMetadataResponse) SetProxyServerConfig(v ClusterProxyServerConfig)`

SetProxyServerConfig sets ProxyServerConfig field to given value.

### HasProxyServerConfig

`func (o *ClusterMetadataResponse) HasProxyServerConfig() bool`

HasProxyServerConfig returns a boolean if a field has been set.

### SetProxyServerConfigNil

`func (o *ClusterMetadataResponse) SetProxyServerConfigNil(b bool)`

 SetProxyServerConfigNil sets the value for ProxyServerConfig to be an explicit nil

### UnsetProxyServerConfig
`func (o *ClusterMetadataResponse) UnsetProxyServerConfig()`

UnsetProxyServerConfig ensures that no value is present for ProxyServerConfig, not even an explicit nil
### GetRegionId

`func (o *ClusterMetadataResponse) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ClusterMetadataResponse) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ClusterMetadataResponse) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ClusterMetadataResponse) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *ClusterMetadataResponse) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *ClusterMetadataResponse) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetRigelClusterParams

`func (o *ClusterMetadataResponse) GetRigelClusterParams() RigelClusterConfigParams`

GetRigelClusterParams returns the RigelClusterParams field if non-nil, zero value otherwise.

### GetRigelClusterParamsOk

`func (o *ClusterMetadataResponse) GetRigelClusterParamsOk() (*RigelClusterConfigParams, bool)`

GetRigelClusterParamsOk returns a tuple with the RigelClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelClusterParams

`func (o *ClusterMetadataResponse) SetRigelClusterParams(v RigelClusterConfigParams)`

SetRigelClusterParams sets RigelClusterParams field to given value.

### HasRigelClusterParams

`func (o *ClusterMetadataResponse) HasRigelClusterParams() bool`

HasRigelClusterParams returns a boolean if a field has been set.

### SetRigelClusterParamsNil

`func (o *ClusterMetadataResponse) SetRigelClusterParamsNil(b bool)`

 SetRigelClusterParamsNil sets the value for RigelClusterParams to be an explicit nil

### UnsetRigelClusterParams
`func (o *ClusterMetadataResponse) UnsetRigelClusterParams()`

UnsetRigelClusterParams ensures that no value is present for RigelClusterParams, not even an explicit nil
### GetSwVersion

`func (o *ClusterMetadataResponse) GetSwVersion() string`

GetSwVersion returns the SwVersion field if non-nil, zero value otherwise.

### GetSwVersionOk

`func (o *ClusterMetadataResponse) GetSwVersionOk() (*string, bool)`

GetSwVersionOk returns a tuple with the SwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwVersion

`func (o *ClusterMetadataResponse) SetSwVersion(v string)`

SetSwVersion sets SwVersion field to given value.

### HasSwVersion

`func (o *ClusterMetadataResponse) HasSwVersion() bool`

HasSwVersion returns a boolean if a field has been set.

### SetSwVersionNil

`func (o *ClusterMetadataResponse) SetSwVersionNil(b bool)`

 SetSwVersionNil sets the value for SwVersion to be an explicit nil

### UnsetSwVersion
`func (o *ClusterMetadataResponse) UnsetSwVersion()`

UnsetSwVersion ensures that no value is present for SwVersion, not even an explicit nil
### GetTenantId

`func (o *ClusterMetadataResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ClusterMetadataResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ClusterMetadataResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ClusterMetadataResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ClusterMetadataResponse) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ClusterMetadataResponse) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTieringAuditLogConfig

`func (o *ClusterMetadataResponse) GetTieringAuditLogConfig() AuditLogConfig`

GetTieringAuditLogConfig returns the TieringAuditLogConfig field if non-nil, zero value otherwise.

### GetTieringAuditLogConfigOk

`func (o *ClusterMetadataResponse) GetTieringAuditLogConfigOk() (*AuditLogConfig, bool)`

GetTieringAuditLogConfigOk returns a tuple with the TieringAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringAuditLogConfig

`func (o *ClusterMetadataResponse) SetTieringAuditLogConfig(v AuditLogConfig)`

SetTieringAuditLogConfig sets TieringAuditLogConfig field to given value.

### HasTieringAuditLogConfig

`func (o *ClusterMetadataResponse) HasTieringAuditLogConfig() bool`

HasTieringAuditLogConfig returns a boolean if a field has been set.

### GetType

`func (o *ClusterMetadataResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterMetadataResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterMetadataResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterMetadataResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ClusterMetadataResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ClusterMetadataResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewsGlobalSettings

`func (o *ClusterMetadataResponse) GetViewsGlobalSettings() ViewsGlobalSettings`

GetViewsGlobalSettings returns the ViewsGlobalSettings field if non-nil, zero value otherwise.

### GetViewsGlobalSettingsOk

`func (o *ClusterMetadataResponse) GetViewsGlobalSettingsOk() (*ViewsGlobalSettings, bool)`

GetViewsGlobalSettingsOk returns a tuple with the ViewsGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsGlobalSettings

`func (o *ClusterMetadataResponse) SetViewsGlobalSettings(v ViewsGlobalSettings)`

SetViewsGlobalSettings sets ViewsGlobalSettings field to given value.

### HasViewsGlobalSettings

`func (o *ClusterMetadataResponse) HasViewsGlobalSettings() bool`

HasViewsGlobalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


