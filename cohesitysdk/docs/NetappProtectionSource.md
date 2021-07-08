# NetappProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterInfo** | Pointer to [**NetappClusterInfo**](NetappClusterInfo.md) |  | [optional] 
**IsTopLevel** | Pointer to **NullableBool** | Specifies if this Object is a top level Object. Because a top level Object can either be a NetApp cluster or a Vserver, this cannot be determined only by type. | [optional] 
**LicenseTypes** | Pointer to **[]string** | Specifies the type of license available on Netapp Cluster &#39;kSnapmirrorCloud&#39; indicates a SnapMirror license on Netapp. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the NetApp Object. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed NetApp Object in a NetApp Protection Source such as &#39;kCluster&#39;, &#39;kVserver&#39; or &#39;kVolume&#39;. &#39;kCluster&#39; indicates a Netapp cluster as a protection source. &#39;kVserver&#39; indicates a Netapp vserver in a cluster as a protection source. &#39;kVolume&#39; indicates  a volume in Netapp vserver as a protection source. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the globally unique ID of this Object assigned by the NetApp server. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the version of Netapp Cluster. | [optional] 
**VolumeInfo** | Pointer to [**NetappVolumeInfo**](NetappVolumeInfo.md) |  | [optional] 
**VserverInfo** | Pointer to [**NetappVserverInfo**](NetappVserverInfo.md) |  | [optional] 

## Methods

### NewNetappProtectionSource

`func NewNetappProtectionSource() *NetappProtectionSource`

NewNetappProtectionSource instantiates a new NetappProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappProtectionSourceWithDefaults

`func NewNetappProtectionSourceWithDefaults() *NetappProtectionSource`

NewNetappProtectionSourceWithDefaults instantiates a new NetappProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterInfo

`func (o *NetappProtectionSource) GetClusterInfo() NetappClusterInfo`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *NetappProtectionSource) GetClusterInfoOk() (*NetappClusterInfo, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *NetappProtectionSource) SetClusterInfo(v NetappClusterInfo)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *NetappProtectionSource) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### GetIsTopLevel

`func (o *NetappProtectionSource) GetIsTopLevel() bool`

GetIsTopLevel returns the IsTopLevel field if non-nil, zero value otherwise.

### GetIsTopLevelOk

`func (o *NetappProtectionSource) GetIsTopLevelOk() (*bool, bool)`

GetIsTopLevelOk returns a tuple with the IsTopLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTopLevel

`func (o *NetappProtectionSource) SetIsTopLevel(v bool)`

SetIsTopLevel sets IsTopLevel field to given value.

### HasIsTopLevel

`func (o *NetappProtectionSource) HasIsTopLevel() bool`

HasIsTopLevel returns a boolean if a field has been set.

### SetIsTopLevelNil

`func (o *NetappProtectionSource) SetIsTopLevelNil(b bool)`

 SetIsTopLevelNil sets the value for IsTopLevel to be an explicit nil

### UnsetIsTopLevel
`func (o *NetappProtectionSource) UnsetIsTopLevel()`

UnsetIsTopLevel ensures that no value is present for IsTopLevel, not even an explicit nil
### GetLicenseTypes

`func (o *NetappProtectionSource) GetLicenseTypes() []string`

GetLicenseTypes returns the LicenseTypes field if non-nil, zero value otherwise.

### GetLicenseTypesOk

`func (o *NetappProtectionSource) GetLicenseTypesOk() (*[]string, bool)`

GetLicenseTypesOk returns a tuple with the LicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypes

`func (o *NetappProtectionSource) SetLicenseTypes(v []string)`

SetLicenseTypes sets LicenseTypes field to given value.

### HasLicenseTypes

`func (o *NetappProtectionSource) HasLicenseTypes() bool`

HasLicenseTypes returns a boolean if a field has been set.

### SetLicenseTypesNil

`func (o *NetappProtectionSource) SetLicenseTypesNil(b bool)`

 SetLicenseTypesNil sets the value for LicenseTypes to be an explicit nil

### UnsetLicenseTypes
`func (o *NetappProtectionSource) UnsetLicenseTypes()`

UnsetLicenseTypes ensures that no value is present for LicenseTypes, not even an explicit nil
### GetName

`func (o *NetappProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetappProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetappProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetappProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NetappProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NetappProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *NetappProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetappProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetappProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetappProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NetappProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NetappProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *NetappProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NetappProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NetappProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NetappProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *NetappProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *NetappProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *NetappProtectionSource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetappProtectionSource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetappProtectionSource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetappProtectionSource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NetappProtectionSource) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NetappProtectionSource) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVolumeInfo

`func (o *NetappProtectionSource) GetVolumeInfo() NetappVolumeInfo`

GetVolumeInfo returns the VolumeInfo field if non-nil, zero value otherwise.

### GetVolumeInfoOk

`func (o *NetappProtectionSource) GetVolumeInfoOk() (*NetappVolumeInfo, bool)`

GetVolumeInfoOk returns a tuple with the VolumeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInfo

`func (o *NetappProtectionSource) SetVolumeInfo(v NetappVolumeInfo)`

SetVolumeInfo sets VolumeInfo field to given value.

### HasVolumeInfo

`func (o *NetappProtectionSource) HasVolumeInfo() bool`

HasVolumeInfo returns a boolean if a field has been set.

### GetVserverInfo

`func (o *NetappProtectionSource) GetVserverInfo() NetappVserverInfo`

GetVserverInfo returns the VserverInfo field if non-nil, zero value otherwise.

### GetVserverInfoOk

`func (o *NetappProtectionSource) GetVserverInfoOk() (*NetappVserverInfo, bool)`

GetVserverInfoOk returns a tuple with the VserverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVserverInfo

`func (o *NetappProtectionSource) SetVserverInfo(v NetappVserverInfo)`

SetVserverInfo sets VserverInfo field to given value.

### HasVserverInfo

`func (o *NetappProtectionSource) HasVserverInfo() bool`

HasVserverInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


