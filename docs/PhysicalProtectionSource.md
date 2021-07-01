# PhysicalProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]AgentInformation**](AgentInformation.md) | Array of Agents on the Physical Protection Source.  Specifiles the agents running on the Physical Protection Source and the status information. | [optional] 
**HostName** | Pointer to **NullableString** | Specifies the hostname. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the environment type for the host. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**Id** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a unique id of a Physical Protection Source. The id is unique across Cohesity Clusters. | [optional] 
**MemorySizeBytes** | Pointer to **NullableInt64** | Specifies the total memory on the host in bytes. | [optional] 
**Name** | Pointer to **NullableString** | Specifies a human readable name of the Protection Source. | [optional] 
**NetworkingInfo** | Pointer to [**NetworkingInformation**](NetworkingInformation.md) |  | [optional] 
**NumProcessors** | Pointer to **NullableInt64** | Specifies the number of processors on the host. | [optional] 
**OsName** | Pointer to **NullableString** | Specifies a human readable name of the OS of the Protection Source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed Object in a Physical Protection Source. &#39;kGroup&#39; indicates the EH container. &#39;kHost&#39; indicates a single physical server. &#39;kWindowsCluster&#39; indicates a Microsoft Windows cluster. &#39;kOracleRACCluster&#39; indicates an Oracle Real Application Cluster(RAC). &#39;kOracleAPCluster&#39; indicates an Oracle Active-Passive Cluster. | [optional] 
**VcsVersion** | Pointer to **NullableString** | Specifies cluster version for VCS host. | [optional] 
**Volumes** | Pointer to [**[]PhysicalVolume**](PhysicalVolume.md) | Array of Physical Volumes.  Specifies the volumes available on the physical host. These fields are populated only for the kPhysicalHost type. | [optional] 

## Methods

### NewPhysicalProtectionSource

`func NewPhysicalProtectionSource() *PhysicalProtectionSource`

NewPhysicalProtectionSource instantiates a new PhysicalProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalProtectionSourceWithDefaults

`func NewPhysicalProtectionSourceWithDefaults() *PhysicalProtectionSource`

NewPhysicalProtectionSourceWithDefaults instantiates a new PhysicalProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *PhysicalProtectionSource) GetAgents() []AgentInformation`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *PhysicalProtectionSource) GetAgentsOk() (*[]AgentInformation, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *PhysicalProtectionSource) SetAgents(v []AgentInformation)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *PhysicalProtectionSource) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *PhysicalProtectionSource) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *PhysicalProtectionSource) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil
### GetHostName

`func (o *PhysicalProtectionSource) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *PhysicalProtectionSource) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *PhysicalProtectionSource) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *PhysicalProtectionSource) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *PhysicalProtectionSource) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *PhysicalProtectionSource) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetHostType

`func (o *PhysicalProtectionSource) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *PhysicalProtectionSource) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *PhysicalProtectionSource) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *PhysicalProtectionSource) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *PhysicalProtectionSource) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *PhysicalProtectionSource) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetId

`func (o *PhysicalProtectionSource) GetId() UniversalId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalProtectionSource) GetIdOk() (*UniversalId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalProtectionSource) SetId(v UniversalId)`

SetId sets Id field to given value.

### HasId

`func (o *PhysicalProtectionSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PhysicalProtectionSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PhysicalProtectionSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMemorySizeBytes

`func (o *PhysicalProtectionSource) GetMemorySizeBytes() int64`

GetMemorySizeBytes returns the MemorySizeBytes field if non-nil, zero value otherwise.

### GetMemorySizeBytesOk

`func (o *PhysicalProtectionSource) GetMemorySizeBytesOk() (*int64, bool)`

GetMemorySizeBytesOk returns a tuple with the MemorySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeBytes

`func (o *PhysicalProtectionSource) SetMemorySizeBytes(v int64)`

SetMemorySizeBytes sets MemorySizeBytes field to given value.

### HasMemorySizeBytes

`func (o *PhysicalProtectionSource) HasMemorySizeBytes() bool`

HasMemorySizeBytes returns a boolean if a field has been set.

### SetMemorySizeBytesNil

`func (o *PhysicalProtectionSource) SetMemorySizeBytesNil(b bool)`

 SetMemorySizeBytesNil sets the value for MemorySizeBytes to be an explicit nil

### UnsetMemorySizeBytes
`func (o *PhysicalProtectionSource) UnsetMemorySizeBytes()`

UnsetMemorySizeBytes ensures that no value is present for MemorySizeBytes, not even an explicit nil
### GetName

`func (o *PhysicalProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhysicalProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PhysicalProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PhysicalProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkingInfo

`func (o *PhysicalProtectionSource) GetNetworkingInfo() NetworkingInformation`

GetNetworkingInfo returns the NetworkingInfo field if non-nil, zero value otherwise.

### GetNetworkingInfoOk

`func (o *PhysicalProtectionSource) GetNetworkingInfoOk() (*NetworkingInformation, bool)`

GetNetworkingInfoOk returns a tuple with the NetworkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingInfo

`func (o *PhysicalProtectionSource) SetNetworkingInfo(v NetworkingInformation)`

SetNetworkingInfo sets NetworkingInfo field to given value.

### HasNetworkingInfo

`func (o *PhysicalProtectionSource) HasNetworkingInfo() bool`

HasNetworkingInfo returns a boolean if a field has been set.

### GetNumProcessors

`func (o *PhysicalProtectionSource) GetNumProcessors() int64`

GetNumProcessors returns the NumProcessors field if non-nil, zero value otherwise.

### GetNumProcessorsOk

`func (o *PhysicalProtectionSource) GetNumProcessorsOk() (*int64, bool)`

GetNumProcessorsOk returns a tuple with the NumProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProcessors

`func (o *PhysicalProtectionSource) SetNumProcessors(v int64)`

SetNumProcessors sets NumProcessors field to given value.

### HasNumProcessors

`func (o *PhysicalProtectionSource) HasNumProcessors() bool`

HasNumProcessors returns a boolean if a field has been set.

### SetNumProcessorsNil

`func (o *PhysicalProtectionSource) SetNumProcessorsNil(b bool)`

 SetNumProcessorsNil sets the value for NumProcessors to be an explicit nil

### UnsetNumProcessors
`func (o *PhysicalProtectionSource) UnsetNumProcessors()`

UnsetNumProcessors ensures that no value is present for NumProcessors, not even an explicit nil
### GetOsName

`func (o *PhysicalProtectionSource) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *PhysicalProtectionSource) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *PhysicalProtectionSource) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *PhysicalProtectionSource) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### SetOsNameNil

`func (o *PhysicalProtectionSource) SetOsNameNil(b bool)`

 SetOsNameNil sets the value for OsName to be an explicit nil

### UnsetOsName
`func (o *PhysicalProtectionSource) UnsetOsName()`

UnsetOsName ensures that no value is present for OsName, not even an explicit nil
### GetType

`func (o *PhysicalProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhysicalProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhysicalProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhysicalProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PhysicalProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PhysicalProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVcsVersion

`func (o *PhysicalProtectionSource) GetVcsVersion() string`

GetVcsVersion returns the VcsVersion field if non-nil, zero value otherwise.

### GetVcsVersionOk

`func (o *PhysicalProtectionSource) GetVcsVersionOk() (*string, bool)`

GetVcsVersionOk returns a tuple with the VcsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsVersion

`func (o *PhysicalProtectionSource) SetVcsVersion(v string)`

SetVcsVersion sets VcsVersion field to given value.

### HasVcsVersion

`func (o *PhysicalProtectionSource) HasVcsVersion() bool`

HasVcsVersion returns a boolean if a field has been set.

### SetVcsVersionNil

`func (o *PhysicalProtectionSource) SetVcsVersionNil(b bool)`

 SetVcsVersionNil sets the value for VcsVersion to be an explicit nil

### UnsetVcsVersion
`func (o *PhysicalProtectionSource) UnsetVcsVersion()`

UnsetVcsVersion ensures that no value is present for VcsVersion, not even an explicit nil
### GetVolumes

`func (o *PhysicalProtectionSource) GetVolumes() []PhysicalVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *PhysicalProtectionSource) GetVolumesOk() (*[]PhysicalVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *PhysicalProtectionSource) SetVolumes(v []PhysicalVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *PhysicalProtectionSource) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *PhysicalProtectionSource) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *PhysicalProtectionSource) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


