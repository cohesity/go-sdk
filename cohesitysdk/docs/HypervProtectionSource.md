# HypervProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]AgentInformation**](AgentInformation.md) | Array of Agents on the Physical Protection Source.  Specifiles the agents running on the HyperV Protection Source and the status information. | [optional] 
**BackupType** | Pointer to **NullableString** | Specifies the type of backup supported by the VM. &#39;kRctBackup&#39;, &#39;kVssBackup&#39; Specifies the type of an HyperV datastore object. &#39;kRctBackup&#39; indicates backup is done using RCT/checkpoints. &#39;kVssBackup&#39; indicates backup is done using VSS. | [optional] 
**BuildNumber** | Pointer to **NullableString** | Specifies the build number for HyperV SCVMM Servers. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name for &#39;kHostCluster&#39; objects. | [optional] 
**DatastoreInfo** | Pointer to [**HypervDatastore**](HypervDatastore.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the Protection Source. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies host OS type for &#39;kVirtualMachine&#39; objects. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**HypervUuid** | Pointer to **NullableString** | Specifies the UUID for &#39;kVirtualMachine&#39; HyperV objects. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the HyperV Object. | [optional] 
**TagAttributes** | Pointer to [**[]TagAttribute**](TagAttribute.md) | Specifies the list of VM Tag attributes associated with this Object. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of an HyperV Protection Source Object such as &#39;kSCVMMServer&#39;, &#39;kStandaloneHost&#39;, &#39;kNetwork&#39;, etc. overrideDescription: true Specifies the type of an HyperV Protection Source. &#39;kSCVMMServer&#39; indicates a collection of root folders clusters. &#39;kStandaloneHost&#39; indicates a single Nutanix cluster. &#39;kStandaloneCluster&#39; indicates a single Nutanix cluster. &#39;kHostGroup&#39; indicates a Nutanix cluster managed by a Prism Central. &#39;kHypervHost&#39; indicates an HyperV host. &#39;kHostCluster&#39; indicates a Nutanix cluster managed by a Prism Central. &#39;kVirtualMachine&#39; indicates a Virtual Machine. &#39;kNetwork&#39; indicates a Virtual Machine network object. &#39;kDatastore&#39; represents a storage container object. &#39;kTag&#39; indicates a tag type object. &#39;kCustomProperty&#39; indicates a custom property including tag type. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the Object. This is unique within the HyperV environment. | [optional] 
**VmInfo** | Pointer to [**HypervVirtualMachine**](HypervVirtualMachine.md) |  | [optional] 
**WindowsVersion** | Pointer to **NullableString** | Specifies the windows version for HyperV hosts. | [optional] 

## Methods

### NewHypervProtectionSource

`func NewHypervProtectionSource() *HypervProtectionSource`

NewHypervProtectionSource instantiates a new HypervProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervProtectionSourceWithDefaults

`func NewHypervProtectionSourceWithDefaults() *HypervProtectionSource`

NewHypervProtectionSourceWithDefaults instantiates a new HypervProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *HypervProtectionSource) GetAgents() []AgentInformation`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *HypervProtectionSource) GetAgentsOk() (*[]AgentInformation, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *HypervProtectionSource) SetAgents(v []AgentInformation)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *HypervProtectionSource) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *HypervProtectionSource) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *HypervProtectionSource) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil
### GetBackupType

`func (o *HypervProtectionSource) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *HypervProtectionSource) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *HypervProtectionSource) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *HypervProtectionSource) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### SetBackupTypeNil

`func (o *HypervProtectionSource) SetBackupTypeNil(b bool)`

 SetBackupTypeNil sets the value for BackupType to be an explicit nil

### UnsetBackupType
`func (o *HypervProtectionSource) UnsetBackupType()`

UnsetBackupType ensures that no value is present for BackupType, not even an explicit nil
### GetBuildNumber

`func (o *HypervProtectionSource) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *HypervProtectionSource) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *HypervProtectionSource) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *HypervProtectionSource) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### SetBuildNumberNil

`func (o *HypervProtectionSource) SetBuildNumberNil(b bool)`

 SetBuildNumberNil sets the value for BuildNumber to be an explicit nil

### UnsetBuildNumber
`func (o *HypervProtectionSource) UnsetBuildNumber()`

UnsetBuildNumber ensures that no value is present for BuildNumber, not even an explicit nil
### GetClusterName

`func (o *HypervProtectionSource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HypervProtectionSource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HypervProtectionSource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HypervProtectionSource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *HypervProtectionSource) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *HypervProtectionSource) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetDatastoreInfo

`func (o *HypervProtectionSource) GetDatastoreInfo() HypervDatastore`

GetDatastoreInfo returns the DatastoreInfo field if non-nil, zero value otherwise.

### GetDatastoreInfoOk

`func (o *HypervProtectionSource) GetDatastoreInfoOk() (*HypervDatastore, bool)`

GetDatastoreInfoOk returns a tuple with the DatastoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreInfo

`func (o *HypervProtectionSource) SetDatastoreInfo(v HypervDatastore)`

SetDatastoreInfo sets DatastoreInfo field to given value.

### HasDatastoreInfo

`func (o *HypervProtectionSource) HasDatastoreInfo() bool`

HasDatastoreInfo returns a boolean if a field has been set.

### GetDescription

`func (o *HypervProtectionSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervProtectionSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervProtectionSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervProtectionSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HypervProtectionSource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HypervProtectionSource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHostType

`func (o *HypervProtectionSource) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HypervProtectionSource) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HypervProtectionSource) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *HypervProtectionSource) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *HypervProtectionSource) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *HypervProtectionSource) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetHypervUuid

`func (o *HypervProtectionSource) GetHypervUuid() string`

GetHypervUuid returns the HypervUuid field if non-nil, zero value otherwise.

### GetHypervUuidOk

`func (o *HypervProtectionSource) GetHypervUuidOk() (*string, bool)`

GetHypervUuidOk returns a tuple with the HypervUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervUuid

`func (o *HypervProtectionSource) SetHypervUuid(v string)`

SetHypervUuid sets HypervUuid field to given value.

### HasHypervUuid

`func (o *HypervProtectionSource) HasHypervUuid() bool`

HasHypervUuid returns a boolean if a field has been set.

### SetHypervUuidNil

`func (o *HypervProtectionSource) SetHypervUuidNil(b bool)`

 SetHypervUuidNil sets the value for HypervUuid to be an explicit nil

### UnsetHypervUuid
`func (o *HypervProtectionSource) UnsetHypervUuid()`

UnsetHypervUuid ensures that no value is present for HypervUuid, not even an explicit nil
### GetName

`func (o *HypervProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTagAttributes

`func (o *HypervProtectionSource) GetTagAttributes() []TagAttribute`

GetTagAttributes returns the TagAttributes field if non-nil, zero value otherwise.

### GetTagAttributesOk

`func (o *HypervProtectionSource) GetTagAttributesOk() (*[]TagAttribute, bool)`

GetTagAttributesOk returns a tuple with the TagAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAttributes

`func (o *HypervProtectionSource) SetTagAttributes(v []TagAttribute)`

SetTagAttributes sets TagAttributes field to given value.

### HasTagAttributes

`func (o *HypervProtectionSource) HasTagAttributes() bool`

HasTagAttributes returns a boolean if a field has been set.

### SetTagAttributesNil

`func (o *HypervProtectionSource) SetTagAttributesNil(b bool)`

 SetTagAttributesNil sets the value for TagAttributes to be an explicit nil

### UnsetTagAttributes
`func (o *HypervProtectionSource) UnsetTagAttributes()`

UnsetTagAttributes ensures that no value is present for TagAttributes, not even an explicit nil
### GetType

`func (o *HypervProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HypervProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HypervProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HypervProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HypervProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HypervProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *HypervProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HypervProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HypervProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HypervProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *HypervProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *HypervProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVmInfo

`func (o *HypervProtectionSource) GetVmInfo() HypervVirtualMachine`

GetVmInfo returns the VmInfo field if non-nil, zero value otherwise.

### GetVmInfoOk

`func (o *HypervProtectionSource) GetVmInfoOk() (*HypervVirtualMachine, bool)`

GetVmInfoOk returns a tuple with the VmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInfo

`func (o *HypervProtectionSource) SetVmInfo(v HypervVirtualMachine)`

SetVmInfo sets VmInfo field to given value.

### HasVmInfo

`func (o *HypervProtectionSource) HasVmInfo() bool`

HasVmInfo returns a boolean if a field has been set.

### GetWindowsVersion

`func (o *HypervProtectionSource) GetWindowsVersion() string`

GetWindowsVersion returns the WindowsVersion field if non-nil, zero value otherwise.

### GetWindowsVersionOk

`func (o *HypervProtectionSource) GetWindowsVersionOk() (*string, bool)`

GetWindowsVersionOk returns a tuple with the WindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsVersion

`func (o *HypervProtectionSource) SetWindowsVersion(v string)`

SetWindowsVersion sets WindowsVersion field to given value.

### HasWindowsVersion

`func (o *HypervProtectionSource) HasWindowsVersion() bool`

HasWindowsVersion returns a boolean if a field has been set.

### SetWindowsVersionNil

`func (o *HypervProtectionSource) SetWindowsVersionNil(b bool)`

 SetWindowsVersionNil sets the value for WindowsVersion to be an explicit nil

### UnsetWindowsVersion
`func (o *HypervProtectionSource) UnsetWindowsVersion()`

UnsetWindowsVersion ensures that no value is present for WindowsVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


