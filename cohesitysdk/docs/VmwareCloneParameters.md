# VmwareCloneParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreFolderId** | Pointer to **NullableInt64** | Specifies the folder where the restore datastore should be created. This is applicable only when the VMs are being cloned. | [optional] 
**DetachNetwork** | Pointer to **NullableBool** | Specifies whether the network should be detached from the recovered or cloned VMs. | [optional] 
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the network should be left in disabled state. Attached network is enabled by default. Set this flag to true to disable it. | [optional] 
**NetworkId** | Pointer to **NullableInt64** | Specifies a network configuration to be attached to the cloned or recovered object. For kCloneVMs and kRecoverVMs tasks, original network configuration is detached if the cloned or recovered object is kept under a different parent Protection Source or a different Resource Pool. By default, for kRecoverVMs task, original network configuration is preserved if the recovered object is kept under the same parent Protection Source and the same Resource Pool. Specify this field to override the preserved network configuration or to attach a new network configuration to the cloned or recovered objects. You can get the networkId of the kNetwork object by setting includeNetworks to &#39;true&#39; in the GET /public/protectionSources operation. In the response, get the id of the desired kNetwork object, the resource pool, and the registered parent Protection Source. | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapping**](NetworkMapping.md) | Specifies the parameters for mapping the source and target networks. This field can be used if restoring to a different parent source. This will replace the NetworkId and DisableNetwork that are used to provide configuration for a single network. Unless the support for mapping is available for all the entities old keys can be used to attach a new network. Supports &#39;kVMware&#39; for now. | [optional] 
**PoweredOn** | Pointer to **NullableBool** | Specifies the power state of the cloned or recovered objects. By default, the cloned or recovered objects are powered off. | [optional] 
**Prefix** | Pointer to **NullableString** | Specifies a prefix to prepended to the source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 
**PreserveCustomAttributesDuringClone** | Pointer to **NullableBool** | Specifies whether or not to preserve the custom attributes during the clone operation. The default behavior is &#39;true&#39;. | [optional] 
**PreserveTags** | Pointer to **NullableBool** | Specifies whether or not to preserve tags during the clone operation. The default behavior is &#39;true&#39;. | [optional] 
**RecoveryProcessType** | Pointer to **NullableString** | Specifies the type of recovery process to be performed. If unspecified, then an instant recovery will be performed. Specifies the recovery process type to be used.. &#39;kInstantRecovery&#39; indicates that an instant recovery should be performed. &#39;kCopyRecovery&#39; indicates that a copy recovery should be performed. | [optional] 
**ResourcePoolId** | Pointer to **NullableInt64** | Specifies the resource pool where the cloned or recovered objects are attached. This field is mandatory for kCloneVMs Restore Tasks always. For kRecoverVMs Restore Tasks, this field is mandatory only if newParentId field is specified. If this field is not specified, recovered objects are attached to the original resource pool under the original parent. | [optional] 
**StorageProfileName** | Pointer to **NullableString** | Specifies the name of the destination storage profile while restoring to an alternate VCD location. | [optional] 
**StorageProfileVcdUuid** | Pointer to **NullableString** | Specifies the UUID of the storage profile while restoring to an alternate VCD location. | [optional] 
**Suffix** | Pointer to **NullableString** | Specifies a suffix to appended to the original source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 
**VAppId** | Pointer to **NullableInt64** | Specifies the ID of the vApp to which a VM should be restored. | [optional] 
**VdcId** | Pointer to **NullableInt64** | Specifies the ID of the VDC to which a VM should be restored. | [optional] 
**VmFolderId** | Pointer to **NullableInt64** | Specifies a folder where the VMs should be restored. This is applicable only when the VMs are being restored to an alternate location or if clone is being performed. | [optional] 

## Methods

### NewVmwareCloneParameters

`func NewVmwareCloneParameters() *VmwareCloneParameters`

NewVmwareCloneParameters instantiates a new VmwareCloneParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareCloneParametersWithDefaults

`func NewVmwareCloneParametersWithDefaults() *VmwareCloneParameters`

NewVmwareCloneParametersWithDefaults instantiates a new VmwareCloneParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreFolderId

`func (o *VmwareCloneParameters) GetDatastoreFolderId() int64`

GetDatastoreFolderId returns the DatastoreFolderId field if non-nil, zero value otherwise.

### GetDatastoreFolderIdOk

`func (o *VmwareCloneParameters) GetDatastoreFolderIdOk() (*int64, bool)`

GetDatastoreFolderIdOk returns a tuple with the DatastoreFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreFolderId

`func (o *VmwareCloneParameters) SetDatastoreFolderId(v int64)`

SetDatastoreFolderId sets DatastoreFolderId field to given value.

### HasDatastoreFolderId

`func (o *VmwareCloneParameters) HasDatastoreFolderId() bool`

HasDatastoreFolderId returns a boolean if a field has been set.

### SetDatastoreFolderIdNil

`func (o *VmwareCloneParameters) SetDatastoreFolderIdNil(b bool)`

 SetDatastoreFolderIdNil sets the value for DatastoreFolderId to be an explicit nil

### UnsetDatastoreFolderId
`func (o *VmwareCloneParameters) UnsetDatastoreFolderId()`

UnsetDatastoreFolderId ensures that no value is present for DatastoreFolderId, not even an explicit nil
### GetDetachNetwork

`func (o *VmwareCloneParameters) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *VmwareCloneParameters) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *VmwareCloneParameters) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *VmwareCloneParameters) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *VmwareCloneParameters) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *VmwareCloneParameters) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetDisableNetwork

`func (o *VmwareCloneParameters) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *VmwareCloneParameters) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *VmwareCloneParameters) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *VmwareCloneParameters) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *VmwareCloneParameters) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *VmwareCloneParameters) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetNetworkId

`func (o *VmwareCloneParameters) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VmwareCloneParameters) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VmwareCloneParameters) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VmwareCloneParameters) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *VmwareCloneParameters) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *VmwareCloneParameters) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetNetworkMappings

`func (o *VmwareCloneParameters) GetNetworkMappings() []NetworkMapping`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *VmwareCloneParameters) GetNetworkMappingsOk() (*[]NetworkMapping, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *VmwareCloneParameters) SetNetworkMappings(v []NetworkMapping)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *VmwareCloneParameters) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *VmwareCloneParameters) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *VmwareCloneParameters) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetPoweredOn

`func (o *VmwareCloneParameters) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *VmwareCloneParameters) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *VmwareCloneParameters) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *VmwareCloneParameters) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.

### SetPoweredOnNil

`func (o *VmwareCloneParameters) SetPoweredOnNil(b bool)`

 SetPoweredOnNil sets the value for PoweredOn to be an explicit nil

### UnsetPoweredOn
`func (o *VmwareCloneParameters) UnsetPoweredOn()`

UnsetPoweredOn ensures that no value is present for PoweredOn, not even an explicit nil
### GetPrefix

`func (o *VmwareCloneParameters) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VmwareCloneParameters) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VmwareCloneParameters) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VmwareCloneParameters) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *VmwareCloneParameters) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *VmwareCloneParameters) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetPreserveCustomAttributesDuringClone

`func (o *VmwareCloneParameters) GetPreserveCustomAttributesDuringClone() bool`

GetPreserveCustomAttributesDuringClone returns the PreserveCustomAttributesDuringClone field if non-nil, zero value otherwise.

### GetPreserveCustomAttributesDuringCloneOk

`func (o *VmwareCloneParameters) GetPreserveCustomAttributesDuringCloneOk() (*bool, bool)`

GetPreserveCustomAttributesDuringCloneOk returns a tuple with the PreserveCustomAttributesDuringClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveCustomAttributesDuringClone

`func (o *VmwareCloneParameters) SetPreserveCustomAttributesDuringClone(v bool)`

SetPreserveCustomAttributesDuringClone sets PreserveCustomAttributesDuringClone field to given value.

### HasPreserveCustomAttributesDuringClone

`func (o *VmwareCloneParameters) HasPreserveCustomAttributesDuringClone() bool`

HasPreserveCustomAttributesDuringClone returns a boolean if a field has been set.

### SetPreserveCustomAttributesDuringCloneNil

`func (o *VmwareCloneParameters) SetPreserveCustomAttributesDuringCloneNil(b bool)`

 SetPreserveCustomAttributesDuringCloneNil sets the value for PreserveCustomAttributesDuringClone to be an explicit nil

### UnsetPreserveCustomAttributesDuringClone
`func (o *VmwareCloneParameters) UnsetPreserveCustomAttributesDuringClone()`

UnsetPreserveCustomAttributesDuringClone ensures that no value is present for PreserveCustomAttributesDuringClone, not even an explicit nil
### GetPreserveTags

`func (o *VmwareCloneParameters) GetPreserveTags() bool`

GetPreserveTags returns the PreserveTags field if non-nil, zero value otherwise.

### GetPreserveTagsOk

`func (o *VmwareCloneParameters) GetPreserveTagsOk() (*bool, bool)`

GetPreserveTagsOk returns a tuple with the PreserveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTags

`func (o *VmwareCloneParameters) SetPreserveTags(v bool)`

SetPreserveTags sets PreserveTags field to given value.

### HasPreserveTags

`func (o *VmwareCloneParameters) HasPreserveTags() bool`

HasPreserveTags returns a boolean if a field has been set.

### SetPreserveTagsNil

`func (o *VmwareCloneParameters) SetPreserveTagsNil(b bool)`

 SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil

### UnsetPreserveTags
`func (o *VmwareCloneParameters) UnsetPreserveTags()`

UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
### GetRecoveryProcessType

`func (o *VmwareCloneParameters) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *VmwareCloneParameters) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *VmwareCloneParameters) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *VmwareCloneParameters) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### SetRecoveryProcessTypeNil

`func (o *VmwareCloneParameters) SetRecoveryProcessTypeNil(b bool)`

 SetRecoveryProcessTypeNil sets the value for RecoveryProcessType to be an explicit nil

### UnsetRecoveryProcessType
`func (o *VmwareCloneParameters) UnsetRecoveryProcessType()`

UnsetRecoveryProcessType ensures that no value is present for RecoveryProcessType, not even an explicit nil
### GetResourcePoolId

`func (o *VmwareCloneParameters) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *VmwareCloneParameters) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *VmwareCloneParameters) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *VmwareCloneParameters) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *VmwareCloneParameters) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *VmwareCloneParameters) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetStorageProfileName

`func (o *VmwareCloneParameters) GetStorageProfileName() string`

GetStorageProfileName returns the StorageProfileName field if non-nil, zero value otherwise.

### GetStorageProfileNameOk

`func (o *VmwareCloneParameters) GetStorageProfileNameOk() (*string, bool)`

GetStorageProfileNameOk returns a tuple with the StorageProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileName

`func (o *VmwareCloneParameters) SetStorageProfileName(v string)`

SetStorageProfileName sets StorageProfileName field to given value.

### HasStorageProfileName

`func (o *VmwareCloneParameters) HasStorageProfileName() bool`

HasStorageProfileName returns a boolean if a field has been set.

### SetStorageProfileNameNil

`func (o *VmwareCloneParameters) SetStorageProfileNameNil(b bool)`

 SetStorageProfileNameNil sets the value for StorageProfileName to be an explicit nil

### UnsetStorageProfileName
`func (o *VmwareCloneParameters) UnsetStorageProfileName()`

UnsetStorageProfileName ensures that no value is present for StorageProfileName, not even an explicit nil
### GetStorageProfileVcdUuid

`func (o *VmwareCloneParameters) GetStorageProfileVcdUuid() string`

GetStorageProfileVcdUuid returns the StorageProfileVcdUuid field if non-nil, zero value otherwise.

### GetStorageProfileVcdUuidOk

`func (o *VmwareCloneParameters) GetStorageProfileVcdUuidOk() (*string, bool)`

GetStorageProfileVcdUuidOk returns a tuple with the StorageProfileVcdUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileVcdUuid

`func (o *VmwareCloneParameters) SetStorageProfileVcdUuid(v string)`

SetStorageProfileVcdUuid sets StorageProfileVcdUuid field to given value.

### HasStorageProfileVcdUuid

`func (o *VmwareCloneParameters) HasStorageProfileVcdUuid() bool`

HasStorageProfileVcdUuid returns a boolean if a field has been set.

### SetStorageProfileVcdUuidNil

`func (o *VmwareCloneParameters) SetStorageProfileVcdUuidNil(b bool)`

 SetStorageProfileVcdUuidNil sets the value for StorageProfileVcdUuid to be an explicit nil

### UnsetStorageProfileVcdUuid
`func (o *VmwareCloneParameters) UnsetStorageProfileVcdUuid()`

UnsetStorageProfileVcdUuid ensures that no value is present for StorageProfileVcdUuid, not even an explicit nil
### GetSuffix

`func (o *VmwareCloneParameters) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *VmwareCloneParameters) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *VmwareCloneParameters) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *VmwareCloneParameters) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *VmwareCloneParameters) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *VmwareCloneParameters) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetVAppId

`func (o *VmwareCloneParameters) GetVAppId() int64`

GetVAppId returns the VAppId field if non-nil, zero value otherwise.

### GetVAppIdOk

`func (o *VmwareCloneParameters) GetVAppIdOk() (*int64, bool)`

GetVAppIdOk returns a tuple with the VAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAppId

`func (o *VmwareCloneParameters) SetVAppId(v int64)`

SetVAppId sets VAppId field to given value.

### HasVAppId

`func (o *VmwareCloneParameters) HasVAppId() bool`

HasVAppId returns a boolean if a field has been set.

### SetVAppIdNil

`func (o *VmwareCloneParameters) SetVAppIdNil(b bool)`

 SetVAppIdNil sets the value for VAppId to be an explicit nil

### UnsetVAppId
`func (o *VmwareCloneParameters) UnsetVAppId()`

UnsetVAppId ensures that no value is present for VAppId, not even an explicit nil
### GetVdcId

`func (o *VmwareCloneParameters) GetVdcId() int64`

GetVdcId returns the VdcId field if non-nil, zero value otherwise.

### GetVdcIdOk

`func (o *VmwareCloneParameters) GetVdcIdOk() (*int64, bool)`

GetVdcIdOk returns a tuple with the VdcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcId

`func (o *VmwareCloneParameters) SetVdcId(v int64)`

SetVdcId sets VdcId field to given value.

### HasVdcId

`func (o *VmwareCloneParameters) HasVdcId() bool`

HasVdcId returns a boolean if a field has been set.

### SetVdcIdNil

`func (o *VmwareCloneParameters) SetVdcIdNil(b bool)`

 SetVdcIdNil sets the value for VdcId to be an explicit nil

### UnsetVdcId
`func (o *VmwareCloneParameters) UnsetVdcId()`

UnsetVdcId ensures that no value is present for VdcId, not even an explicit nil
### GetVmFolderId

`func (o *VmwareCloneParameters) GetVmFolderId() int64`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *VmwareCloneParameters) GetVmFolderIdOk() (*int64, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *VmwareCloneParameters) SetVmFolderId(v int64)`

SetVmFolderId sets VmFolderId field to given value.

### HasVmFolderId

`func (o *VmwareCloneParameters) HasVmFolderId() bool`

HasVmFolderId returns a boolean if a field has been set.

### SetVmFolderIdNil

`func (o *VmwareCloneParameters) SetVmFolderIdNil(b bool)`

 SetVmFolderIdNil sets the value for VmFolderId to be an explicit nil

### UnsetVmFolderId
`func (o *VmwareCloneParameters) UnsetVmFolderId()`

UnsetVmFolderId ensures that no value is present for VmFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


