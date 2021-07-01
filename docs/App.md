# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | Specifies unique id allocated by the AppStore. | [optional] 
**Clusters** | Pointer to [**[]ClusterInfo**](ClusterInfo.md) | Specifies the list of clusters on which the app is installed for a particular account Id. | [optional] 
**DownloadProgressPct** | Pointer to **NullableFloat64** | Specifies app download progress percentage. | [optional] 
**ExternalIpRequired** | Pointer to **NullableBool** | Specifies if an external ip is required for the app. | [optional] 
**ExternalNetworks** | Pointer to [**[]ExternalNetworkInfo**](ExternalNetworkInfo.md) | List of external network information available for the app. | [optional] 
**InstallState** | Pointer to **NullableString** | Specifies app installation status. Specifies status of the app installation. kNotInstalled - App yet to be installed. kInstallInProgress - App installation is in progress. kInstalled - App is installed successfully and can be launched. kInstallFailed - App installation failed. kUninstallInProgress - App uninstallation is in progress. kUninstallFailed - App uninstallation failed. kDownloadNotStarted - App download has not started. kDownloadInProgress - App download in progress. kDownloadComplete - App download completed. kDownloadFailed - App download failed. | [optional] 
**InstallTime** | Pointer to **NullableInt64** | Specifies timestamp when the app was installed. | [optional] 
**InstanceSizes** | Pointer to **[]string** | List of applicable instance size specifications (e.g. small/medium/large) for the app. Used to determine container resources. | [optional] 
**IsLatest** | Pointer to **NullableBool** | Specifies whether the app currently installed on all clusters is the latest version or not. | [optional] 
**LatestVersion** | Pointer to **NullableInt64** | Specifies application version assigned by the AppStore for the latest version of an app. | [optional] 
**Metadata** | Pointer to [**AppMetadata**](AppMetadata.md) |  | [optional] 
**RequiredPrivileges** | Pointer to **[]string** | Specifies privileges that are required for this app. App privilege information.  Specifies privileges that are required for this app. kReadAccess - App needs views for read access. kReadWriteAccess - App needs views for Read/write access. kManagementAccess - App needs management access via iris API. kAutoMountAccess - Whether to allow auto-mounting all the views. kUnrestrictedAppUIAccess - Whether app requires unrestricted UI access (i.e. without passing app access token in URL). kAuditLogViewReadAccess - Whether app requires read access to the internal audit log view. kProtectedObjectAccess - Whether app requires read access to protected objects. | [optional] 
**UninstallTime** | Pointer to **NullableInt64** | Specifies timestamp when the app was uninstalled. | [optional] 
**Version** | Pointer to **NullableInt64** | Specifies application version assigned by the AppStore. | [optional] 
**VmNameInfoList** | Pointer to [**[]VmNameInfo**](VmNameInfo.md) | List of vm name info objects. | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *App) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *App) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *App) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *App) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *App) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *App) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetClusters

`func (o *App) GetClusters() []ClusterInfo`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *App) GetClustersOk() (*[]ClusterInfo, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *App) SetClusters(v []ClusterInfo)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *App) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### SetClustersNil

`func (o *App) SetClustersNil(b bool)`

 SetClustersNil sets the value for Clusters to be an explicit nil

### UnsetClusters
`func (o *App) UnsetClusters()`

UnsetClusters ensures that no value is present for Clusters, not even an explicit nil
### GetDownloadProgressPct

`func (o *App) GetDownloadProgressPct() float64`

GetDownloadProgressPct returns the DownloadProgressPct field if non-nil, zero value otherwise.

### GetDownloadProgressPctOk

`func (o *App) GetDownloadProgressPctOk() (*float64, bool)`

GetDownloadProgressPctOk returns a tuple with the DownloadProgressPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadProgressPct

`func (o *App) SetDownloadProgressPct(v float64)`

SetDownloadProgressPct sets DownloadProgressPct field to given value.

### HasDownloadProgressPct

`func (o *App) HasDownloadProgressPct() bool`

HasDownloadProgressPct returns a boolean if a field has been set.

### SetDownloadProgressPctNil

`func (o *App) SetDownloadProgressPctNil(b bool)`

 SetDownloadProgressPctNil sets the value for DownloadProgressPct to be an explicit nil

### UnsetDownloadProgressPct
`func (o *App) UnsetDownloadProgressPct()`

UnsetDownloadProgressPct ensures that no value is present for DownloadProgressPct, not even an explicit nil
### GetExternalIpRequired

`func (o *App) GetExternalIpRequired() bool`

GetExternalIpRequired returns the ExternalIpRequired field if non-nil, zero value otherwise.

### GetExternalIpRequiredOk

`func (o *App) GetExternalIpRequiredOk() (*bool, bool)`

GetExternalIpRequiredOk returns a tuple with the ExternalIpRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpRequired

`func (o *App) SetExternalIpRequired(v bool)`

SetExternalIpRequired sets ExternalIpRequired field to given value.

### HasExternalIpRequired

`func (o *App) HasExternalIpRequired() bool`

HasExternalIpRequired returns a boolean if a field has been set.

### SetExternalIpRequiredNil

`func (o *App) SetExternalIpRequiredNil(b bool)`

 SetExternalIpRequiredNil sets the value for ExternalIpRequired to be an explicit nil

### UnsetExternalIpRequired
`func (o *App) UnsetExternalIpRequired()`

UnsetExternalIpRequired ensures that no value is present for ExternalIpRequired, not even an explicit nil
### GetExternalNetworks

`func (o *App) GetExternalNetworks() []ExternalNetworkInfo`

GetExternalNetworks returns the ExternalNetworks field if non-nil, zero value otherwise.

### GetExternalNetworksOk

`func (o *App) GetExternalNetworksOk() (*[]ExternalNetworkInfo, bool)`

GetExternalNetworksOk returns a tuple with the ExternalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetworks

`func (o *App) SetExternalNetworks(v []ExternalNetworkInfo)`

SetExternalNetworks sets ExternalNetworks field to given value.

### HasExternalNetworks

`func (o *App) HasExternalNetworks() bool`

HasExternalNetworks returns a boolean if a field has been set.

### SetExternalNetworksNil

`func (o *App) SetExternalNetworksNil(b bool)`

 SetExternalNetworksNil sets the value for ExternalNetworks to be an explicit nil

### UnsetExternalNetworks
`func (o *App) UnsetExternalNetworks()`

UnsetExternalNetworks ensures that no value is present for ExternalNetworks, not even an explicit nil
### GetInstallState

`func (o *App) GetInstallState() string`

GetInstallState returns the InstallState field if non-nil, zero value otherwise.

### GetInstallStateOk

`func (o *App) GetInstallStateOk() (*string, bool)`

GetInstallStateOk returns a tuple with the InstallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallState

`func (o *App) SetInstallState(v string)`

SetInstallState sets InstallState field to given value.

### HasInstallState

`func (o *App) HasInstallState() bool`

HasInstallState returns a boolean if a field has been set.

### SetInstallStateNil

`func (o *App) SetInstallStateNil(b bool)`

 SetInstallStateNil sets the value for InstallState to be an explicit nil

### UnsetInstallState
`func (o *App) UnsetInstallState()`

UnsetInstallState ensures that no value is present for InstallState, not even an explicit nil
### GetInstallTime

`func (o *App) GetInstallTime() int64`

GetInstallTime returns the InstallTime field if non-nil, zero value otherwise.

### GetInstallTimeOk

`func (o *App) GetInstallTimeOk() (*int64, bool)`

GetInstallTimeOk returns a tuple with the InstallTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallTime

`func (o *App) SetInstallTime(v int64)`

SetInstallTime sets InstallTime field to given value.

### HasInstallTime

`func (o *App) HasInstallTime() bool`

HasInstallTime returns a boolean if a field has been set.

### SetInstallTimeNil

`func (o *App) SetInstallTimeNil(b bool)`

 SetInstallTimeNil sets the value for InstallTime to be an explicit nil

### UnsetInstallTime
`func (o *App) UnsetInstallTime()`

UnsetInstallTime ensures that no value is present for InstallTime, not even an explicit nil
### GetInstanceSizes

`func (o *App) GetInstanceSizes() []string`

GetInstanceSizes returns the InstanceSizes field if non-nil, zero value otherwise.

### GetInstanceSizesOk

`func (o *App) GetInstanceSizesOk() (*[]string, bool)`

GetInstanceSizesOk returns a tuple with the InstanceSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizes

`func (o *App) SetInstanceSizes(v []string)`

SetInstanceSizes sets InstanceSizes field to given value.

### HasInstanceSizes

`func (o *App) HasInstanceSizes() bool`

HasInstanceSizes returns a boolean if a field has been set.

### SetInstanceSizesNil

`func (o *App) SetInstanceSizesNil(b bool)`

 SetInstanceSizesNil sets the value for InstanceSizes to be an explicit nil

### UnsetInstanceSizes
`func (o *App) UnsetInstanceSizes()`

UnsetInstanceSizes ensures that no value is present for InstanceSizes, not even an explicit nil
### GetIsLatest

`func (o *App) GetIsLatest() bool`

GetIsLatest returns the IsLatest field if non-nil, zero value otherwise.

### GetIsLatestOk

`func (o *App) GetIsLatestOk() (*bool, bool)`

GetIsLatestOk returns a tuple with the IsLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatest

`func (o *App) SetIsLatest(v bool)`

SetIsLatest sets IsLatest field to given value.

### HasIsLatest

`func (o *App) HasIsLatest() bool`

HasIsLatest returns a boolean if a field has been set.

### SetIsLatestNil

`func (o *App) SetIsLatestNil(b bool)`

 SetIsLatestNil sets the value for IsLatest to be an explicit nil

### UnsetIsLatest
`func (o *App) UnsetIsLatest()`

UnsetIsLatest ensures that no value is present for IsLatest, not even an explicit nil
### GetLatestVersion

`func (o *App) GetLatestVersion() int64`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *App) GetLatestVersionOk() (*int64, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *App) SetLatestVersion(v int64)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *App) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### SetLatestVersionNil

`func (o *App) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *App) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetMetadata

`func (o *App) GetMetadata() AppMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *App) GetMetadataOk() (*AppMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *App) SetMetadata(v AppMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *App) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRequiredPrivileges

`func (o *App) GetRequiredPrivileges() []string`

GetRequiredPrivileges returns the RequiredPrivileges field if non-nil, zero value otherwise.

### GetRequiredPrivilegesOk

`func (o *App) GetRequiredPrivilegesOk() (*[]string, bool)`

GetRequiredPrivilegesOk returns a tuple with the RequiredPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPrivileges

`func (o *App) SetRequiredPrivileges(v []string)`

SetRequiredPrivileges sets RequiredPrivileges field to given value.

### HasRequiredPrivileges

`func (o *App) HasRequiredPrivileges() bool`

HasRequiredPrivileges returns a boolean if a field has been set.

### SetRequiredPrivilegesNil

`func (o *App) SetRequiredPrivilegesNil(b bool)`

 SetRequiredPrivilegesNil sets the value for RequiredPrivileges to be an explicit nil

### UnsetRequiredPrivileges
`func (o *App) UnsetRequiredPrivileges()`

UnsetRequiredPrivileges ensures that no value is present for RequiredPrivileges, not even an explicit nil
### GetUninstallTime

`func (o *App) GetUninstallTime() int64`

GetUninstallTime returns the UninstallTime field if non-nil, zero value otherwise.

### GetUninstallTimeOk

`func (o *App) GetUninstallTimeOk() (*int64, bool)`

GetUninstallTimeOk returns a tuple with the UninstallTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstallTime

`func (o *App) SetUninstallTime(v int64)`

SetUninstallTime sets UninstallTime field to given value.

### HasUninstallTime

`func (o *App) HasUninstallTime() bool`

HasUninstallTime returns a boolean if a field has been set.

### SetUninstallTimeNil

`func (o *App) SetUninstallTimeNil(b bool)`

 SetUninstallTimeNil sets the value for UninstallTime to be an explicit nil

### UnsetUninstallTime
`func (o *App) UnsetUninstallTime()`

UnsetUninstallTime ensures that no value is present for UninstallTime, not even an explicit nil
### GetVersion

`func (o *App) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *App) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *App) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *App) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *App) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *App) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVmNameInfoList

`func (o *App) GetVmNameInfoList() []VmNameInfo`

GetVmNameInfoList returns the VmNameInfoList field if non-nil, zero value otherwise.

### GetVmNameInfoListOk

`func (o *App) GetVmNameInfoListOk() (*[]VmNameInfo, bool)`

GetVmNameInfoListOk returns a tuple with the VmNameInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNameInfoList

`func (o *App) SetVmNameInfoList(v []VmNameInfo)`

SetVmNameInfoList sets VmNameInfoList field to given value.

### HasVmNameInfoList

`func (o *App) HasVmNameInfoList() bool`

HasVmNameInfoList returns a boolean if a field has been set.

### SetVmNameInfoListNil

`func (o *App) SetVmNameInfoListNil(b bool)`

 SetVmNameInfoListNil sets the value for VmNameInfoList to be an explicit nil

### UnsetVmNameInfoList
`func (o *App) UnsetVmNameInfoList()`

UnsetVmNameInfoList ensures that no value is present for VmNameInfoList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


