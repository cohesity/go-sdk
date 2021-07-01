# RecoverTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicFoldersParameters** | Pointer to [**PublicFoldersRestoreParameters**](PublicFoldersRestoreParameters.md) |  | [optional] 
**AcropolisParameters** | Pointer to [**AcropolisRestoreParameters**](AcropolisRestoreParameters.md) |  | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the Restore Task should continue when some operations on some objects fail. If true, the Cohesity Cluster ignores intermittent errors and restores as many objects as possible. | [optional] 
**DeployVmsToCloud** | Pointer to [**DeployVmsToCloud**](DeployVmsToCloud.md) |  | [optional] 
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the way data needs to be retrieved from the external target. This information will be filled in by Iris and Magneto will pass it along to the Icebox as it is to support bulk retrieval from Glacier. Specifies the type of Restore Task.  &#39;kStandard&#39; specifies retrievals that allow to access any of your archives within several hours. Standard retrievals typically complete within 3–5 hours. This is the default option for retrieval requests that do not specify the retrieval option. &#39;kBulk&#39; specifies retrievals that are Glacier’s lowest-cost retrieval option, which can be used to retrieve large amounts, even petabytes, of data inexpensively in a day. Bulk retrieval typically complete within 5–12 hours. &#39;kExpedited&#39; specifies retrievals that allows to quickly access your data when occasional urgent requests for a subset of archives are required. For all but the largest archives (250 MB+), data accessed using Expedited retrievals are typically made available within 1–5 minutes. | [optional] 
**HypervParameters** | Pointer to [**HypervRestoreParameters**](HypervRestoreParameters.md) |  | [optional] 
**KubernetesParameters** | Pointer to [**KubernetesRestoreParameters**](KubernetesRestoreParameters.md) |  | [optional] 
**MountParameters** | Pointer to [**MountVolumesParameters**](MountVolumesParameters.md) |  | [optional] 
**Name** | **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | 
**NewParentId** | Pointer to **NullableInt64** | Specify a new registered parent Protection Source. If specified the selected objects are cloned or recovered to this new Protection Source. If not specified, objects are cloned or recovered to the original Protection Source that was managing them. | [optional] 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**OneDriveParameters** | Pointer to [**OneDriveRestoreParameters**](OneDriveRestoreParameters.md) |  | [optional] 
**OutlookParameters** | Pointer to [**OutlookRestoreParameters**](OutlookRestoreParameters.md) |  | [optional] 
**RestoreViewParameters** | Pointer to [**UpdateViewParam**](UpdateViewParam.md) |  | [optional] 
**SharePointParameters** | Pointer to [**SharePointRestoreParameters**](SharePointRestoreParameters.md) |  | [optional] 
**Type** | **NullableString** | Specifies the type of Restore Task such as &#39;kRecoverVMs&#39; or &#39;kMountVolumes&#39;. &#39;kRecoverVMs&#39; specifies a Restore Task that recovers VMs. &#39;kMountVolumes&#39; specifies a Restore Task that mounts volumes to mount points. &#39;kRecoverNamespaces&#39; specifies a Restore Task that recovers Kubernetes namespaces. &#39;kMountFileVolume&#39; specifies a Restore Task that mounts a file volume. | 
**ViewName** | Pointer to **NullableString** | Specifie target view into which the objects are to be cloned when doing recovery for NAS. | [optional] 
**VirtualDiskRestoreParameters** | Pointer to [**VirtualDiskRestoreParameters**](VirtualDiskRestoreParameters.md) |  | [optional] 
**VlanParameters** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 
**VmwareParameters** | Pointer to [**VmwareRestoreParameters**](VmwareRestoreParameters.md) |  | [optional] 

## Methods

### NewRecoverTaskRequest

`func NewRecoverTaskRequest(name NullableString, type_ NullableString, ) *RecoverTaskRequest`

NewRecoverTaskRequest instantiates a new RecoverTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverTaskRequestWithDefaults

`func NewRecoverTaskRequestWithDefaults() *RecoverTaskRequest`

NewRecoverTaskRequestWithDefaults instantiates a new RecoverTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicFoldersParameters

`func (o *RecoverTaskRequest) GetPublicFoldersParameters() PublicFoldersRestoreParameters`

GetPublicFoldersParameters returns the PublicFoldersParameters field if non-nil, zero value otherwise.

### GetPublicFoldersParametersOk

`func (o *RecoverTaskRequest) GetPublicFoldersParametersOk() (*PublicFoldersRestoreParameters, bool)`

GetPublicFoldersParametersOk returns a tuple with the PublicFoldersParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFoldersParameters

`func (o *RecoverTaskRequest) SetPublicFoldersParameters(v PublicFoldersRestoreParameters)`

SetPublicFoldersParameters sets PublicFoldersParameters field to given value.

### HasPublicFoldersParameters

`func (o *RecoverTaskRequest) HasPublicFoldersParameters() bool`

HasPublicFoldersParameters returns a boolean if a field has been set.

### GetAcropolisParameters

`func (o *RecoverTaskRequest) GetAcropolisParameters() AcropolisRestoreParameters`

GetAcropolisParameters returns the AcropolisParameters field if non-nil, zero value otherwise.

### GetAcropolisParametersOk

`func (o *RecoverTaskRequest) GetAcropolisParametersOk() (*AcropolisRestoreParameters, bool)`

GetAcropolisParametersOk returns a tuple with the AcropolisParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParameters

`func (o *RecoverTaskRequest) SetAcropolisParameters(v AcropolisRestoreParameters)`

SetAcropolisParameters sets AcropolisParameters field to given value.

### HasAcropolisParameters

`func (o *RecoverTaskRequest) HasAcropolisParameters() bool`

HasAcropolisParameters returns a boolean if a field has been set.

### GetContinueOnError

`func (o *RecoverTaskRequest) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverTaskRequest) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverTaskRequest) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverTaskRequest) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverTaskRequest) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverTaskRequest) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDeployVmsToCloud

`func (o *RecoverTaskRequest) GetDeployVmsToCloud() DeployVmsToCloud`

GetDeployVmsToCloud returns the DeployVmsToCloud field if non-nil, zero value otherwise.

### GetDeployVmsToCloudOk

`func (o *RecoverTaskRequest) GetDeployVmsToCloudOk() (*DeployVmsToCloud, bool)`

GetDeployVmsToCloudOk returns a tuple with the DeployVmsToCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloud

`func (o *RecoverTaskRequest) SetDeployVmsToCloud(v DeployVmsToCloud)`

SetDeployVmsToCloud sets DeployVmsToCloud field to given value.

### HasDeployVmsToCloud

`func (o *RecoverTaskRequest) HasDeployVmsToCloud() bool`

HasDeployVmsToCloud returns a boolean if a field has been set.

### GetGlacierRetrievalType

`func (o *RecoverTaskRequest) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *RecoverTaskRequest) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *RecoverTaskRequest) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *RecoverTaskRequest) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *RecoverTaskRequest) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *RecoverTaskRequest) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
### GetHypervParameters

`func (o *RecoverTaskRequest) GetHypervParameters() HypervRestoreParameters`

GetHypervParameters returns the HypervParameters field if non-nil, zero value otherwise.

### GetHypervParametersOk

`func (o *RecoverTaskRequest) GetHypervParametersOk() (*HypervRestoreParameters, bool)`

GetHypervParametersOk returns a tuple with the HypervParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParameters

`func (o *RecoverTaskRequest) SetHypervParameters(v HypervRestoreParameters)`

SetHypervParameters sets HypervParameters field to given value.

### HasHypervParameters

`func (o *RecoverTaskRequest) HasHypervParameters() bool`

HasHypervParameters returns a boolean if a field has been set.

### GetKubernetesParameters

`func (o *RecoverTaskRequest) GetKubernetesParameters() KubernetesRestoreParameters`

GetKubernetesParameters returns the KubernetesParameters field if non-nil, zero value otherwise.

### GetKubernetesParametersOk

`func (o *RecoverTaskRequest) GetKubernetesParametersOk() (*KubernetesRestoreParameters, bool)`

GetKubernetesParametersOk returns a tuple with the KubernetesParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParameters

`func (o *RecoverTaskRequest) SetKubernetesParameters(v KubernetesRestoreParameters)`

SetKubernetesParameters sets KubernetesParameters field to given value.

### HasKubernetesParameters

`func (o *RecoverTaskRequest) HasKubernetesParameters() bool`

HasKubernetesParameters returns a boolean if a field has been set.

### GetMountParameters

`func (o *RecoverTaskRequest) GetMountParameters() MountVolumesParameters`

GetMountParameters returns the MountParameters field if non-nil, zero value otherwise.

### GetMountParametersOk

`func (o *RecoverTaskRequest) GetMountParametersOk() (*MountVolumesParameters, bool)`

GetMountParametersOk returns a tuple with the MountParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountParameters

`func (o *RecoverTaskRequest) SetMountParameters(v MountVolumesParameters)`

SetMountParameters sets MountParameters field to given value.

### HasMountParameters

`func (o *RecoverTaskRequest) HasMountParameters() bool`

HasMountParameters returns a boolean if a field has been set.

### GetName

`func (o *RecoverTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoverTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoverTaskRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *RecoverTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecoverTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewParentId

`func (o *RecoverTaskRequest) GetNewParentId() int64`

GetNewParentId returns the NewParentId field if non-nil, zero value otherwise.

### GetNewParentIdOk

`func (o *RecoverTaskRequest) GetNewParentIdOk() (*int64, bool)`

GetNewParentIdOk returns a tuple with the NewParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParentId

`func (o *RecoverTaskRequest) SetNewParentId(v int64)`

SetNewParentId sets NewParentId field to given value.

### HasNewParentId

`func (o *RecoverTaskRequest) HasNewParentId() bool`

HasNewParentId returns a boolean if a field has been set.

### SetNewParentIdNil

`func (o *RecoverTaskRequest) SetNewParentIdNil(b bool)`

 SetNewParentIdNil sets the value for NewParentId to be an explicit nil

### UnsetNewParentId
`func (o *RecoverTaskRequest) UnsetNewParentId()`

UnsetNewParentId ensures that no value is present for NewParentId, not even an explicit nil
### GetObjects

`func (o *RecoverTaskRequest) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverTaskRequest) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverTaskRequest) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverTaskRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverTaskRequest) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverTaskRequest) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOneDriveParameters

`func (o *RecoverTaskRequest) GetOneDriveParameters() OneDriveRestoreParameters`

GetOneDriveParameters returns the OneDriveParameters field if non-nil, zero value otherwise.

### GetOneDriveParametersOk

`func (o *RecoverTaskRequest) GetOneDriveParametersOk() (*OneDriveRestoreParameters, bool)`

GetOneDriveParametersOk returns a tuple with the OneDriveParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveParameters

`func (o *RecoverTaskRequest) SetOneDriveParameters(v OneDriveRestoreParameters)`

SetOneDriveParameters sets OneDriveParameters field to given value.

### HasOneDriveParameters

`func (o *RecoverTaskRequest) HasOneDriveParameters() bool`

HasOneDriveParameters returns a boolean if a field has been set.

### GetOutlookParameters

`func (o *RecoverTaskRequest) GetOutlookParameters() OutlookRestoreParameters`

GetOutlookParameters returns the OutlookParameters field if non-nil, zero value otherwise.

### GetOutlookParametersOk

`func (o *RecoverTaskRequest) GetOutlookParametersOk() (*OutlookRestoreParameters, bool)`

GetOutlookParametersOk returns a tuple with the OutlookParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookParameters

`func (o *RecoverTaskRequest) SetOutlookParameters(v OutlookRestoreParameters)`

SetOutlookParameters sets OutlookParameters field to given value.

### HasOutlookParameters

`func (o *RecoverTaskRequest) HasOutlookParameters() bool`

HasOutlookParameters returns a boolean if a field has been set.

### GetRestoreViewParameters

`func (o *RecoverTaskRequest) GetRestoreViewParameters() UpdateViewParam`

GetRestoreViewParameters returns the RestoreViewParameters field if non-nil, zero value otherwise.

### GetRestoreViewParametersOk

`func (o *RecoverTaskRequest) GetRestoreViewParametersOk() (*UpdateViewParam, bool)`

GetRestoreViewParametersOk returns a tuple with the RestoreViewParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreViewParameters

`func (o *RecoverTaskRequest) SetRestoreViewParameters(v UpdateViewParam)`

SetRestoreViewParameters sets RestoreViewParameters field to given value.

### HasRestoreViewParameters

`func (o *RecoverTaskRequest) HasRestoreViewParameters() bool`

HasRestoreViewParameters returns a boolean if a field has been set.

### GetSharePointParameters

`func (o *RecoverTaskRequest) GetSharePointParameters() SharePointRestoreParameters`

GetSharePointParameters returns the SharePointParameters field if non-nil, zero value otherwise.

### GetSharePointParametersOk

`func (o *RecoverTaskRequest) GetSharePointParametersOk() (*SharePointRestoreParameters, bool)`

GetSharePointParametersOk returns a tuple with the SharePointParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePointParameters

`func (o *RecoverTaskRequest) SetSharePointParameters(v SharePointRestoreParameters)`

SetSharePointParameters sets SharePointParameters field to given value.

### HasSharePointParameters

`func (o *RecoverTaskRequest) HasSharePointParameters() bool`

HasSharePointParameters returns a boolean if a field has been set.

### GetType

`func (o *RecoverTaskRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecoverTaskRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecoverTaskRequest) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *RecoverTaskRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RecoverTaskRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewName

`func (o *RecoverTaskRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RecoverTaskRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RecoverTaskRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RecoverTaskRequest) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RecoverTaskRequest) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RecoverTaskRequest) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVirtualDiskRestoreParameters

`func (o *RecoverTaskRequest) GetVirtualDiskRestoreParameters() VirtualDiskRestoreParameters`

GetVirtualDiskRestoreParameters returns the VirtualDiskRestoreParameters field if non-nil, zero value otherwise.

### GetVirtualDiskRestoreParametersOk

`func (o *RecoverTaskRequest) GetVirtualDiskRestoreParametersOk() (*VirtualDiskRestoreParameters, bool)`

GetVirtualDiskRestoreParametersOk returns a tuple with the VirtualDiskRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskRestoreParameters

`func (o *RecoverTaskRequest) SetVirtualDiskRestoreParameters(v VirtualDiskRestoreParameters)`

SetVirtualDiskRestoreParameters sets VirtualDiskRestoreParameters field to given value.

### HasVirtualDiskRestoreParameters

`func (o *RecoverTaskRequest) HasVirtualDiskRestoreParameters() bool`

HasVirtualDiskRestoreParameters returns a boolean if a field has been set.

### GetVlanParameters

`func (o *RecoverTaskRequest) GetVlanParameters() VlanParameters`

GetVlanParameters returns the VlanParameters field if non-nil, zero value otherwise.

### GetVlanParametersOk

`func (o *RecoverTaskRequest) GetVlanParametersOk() (*VlanParameters, bool)`

GetVlanParametersOk returns a tuple with the VlanParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParameters

`func (o *RecoverTaskRequest) SetVlanParameters(v VlanParameters)`

SetVlanParameters sets VlanParameters field to given value.

### HasVlanParameters

`func (o *RecoverTaskRequest) HasVlanParameters() bool`

HasVlanParameters returns a boolean if a field has been set.

### GetVmwareParameters

`func (o *RecoverTaskRequest) GetVmwareParameters() VmwareRestoreParameters`

GetVmwareParameters returns the VmwareParameters field if non-nil, zero value otherwise.

### GetVmwareParametersOk

`func (o *RecoverTaskRequest) GetVmwareParametersOk() (*VmwareRestoreParameters, bool)`

GetVmwareParametersOk returns a tuple with the VmwareParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParameters

`func (o *RecoverTaskRequest) SetVmwareParameters(v VmwareRestoreParameters)`

SetVmwareParameters sets VmwareParameters field to given value.

### HasVmwareParameters

`func (o *RecoverTaskRequest) HasVmwareParameters() bool`

HasVmwareParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


