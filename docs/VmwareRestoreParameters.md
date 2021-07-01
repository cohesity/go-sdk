# VmwareRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDatastoreIds** | Pointer to **[]int64** | Specifies additional datastores where the object should be recovered to. | [optional] 
**DatastoreFolderId** | Pointer to **NullableInt64** | Specifies the folder where the restore datastore should be created. This is applicable only when the VMs are being cloned. | [optional] 
**DatastoreId** | Pointer to **NullableInt64** | Specifies the datastore where the object&#39;s files should be recovered to. This field is mandatory to recover objects to a different resource pool or to a different parent source. If not specified, objects are recovered to their original datastore locations in the parent source. | [optional] 
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

### NewVmwareRestoreParameters

`func NewVmwareRestoreParameters() *VmwareRestoreParameters`

NewVmwareRestoreParameters instantiates a new VmwareRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRestoreParametersWithDefaults

`func NewVmwareRestoreParametersWithDefaults() *VmwareRestoreParameters`

NewVmwareRestoreParametersWithDefaults instantiates a new VmwareRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDatastoreIds

`func (o *VmwareRestoreParameters) GetAdditionalDatastoreIds() []int64`

GetAdditionalDatastoreIds returns the AdditionalDatastoreIds field if non-nil, zero value otherwise.

### GetAdditionalDatastoreIdsOk

`func (o *VmwareRestoreParameters) GetAdditionalDatastoreIdsOk() (*[]int64, bool)`

GetAdditionalDatastoreIdsOk returns a tuple with the AdditionalDatastoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDatastoreIds

`func (o *VmwareRestoreParameters) SetAdditionalDatastoreIds(v []int64)`

SetAdditionalDatastoreIds sets AdditionalDatastoreIds field to given value.

### HasAdditionalDatastoreIds

`func (o *VmwareRestoreParameters) HasAdditionalDatastoreIds() bool`

HasAdditionalDatastoreIds returns a boolean if a field has been set.

### SetAdditionalDatastoreIdsNil

`func (o *VmwareRestoreParameters) SetAdditionalDatastoreIdsNil(b bool)`

 SetAdditionalDatastoreIdsNil sets the value for AdditionalDatastoreIds to be an explicit nil

### UnsetAdditionalDatastoreIds
`func (o *VmwareRestoreParameters) UnsetAdditionalDatastoreIds()`

UnsetAdditionalDatastoreIds ensures that no value is present for AdditionalDatastoreIds, not even an explicit nil
### GetDatastoreFolderId

`func (o *VmwareRestoreParameters) GetDatastoreFolderId() int64`

GetDatastoreFolderId returns the DatastoreFolderId field if non-nil, zero value otherwise.

### GetDatastoreFolderIdOk

`func (o *VmwareRestoreParameters) GetDatastoreFolderIdOk() (*int64, bool)`

GetDatastoreFolderIdOk returns a tuple with the DatastoreFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreFolderId

`func (o *VmwareRestoreParameters) SetDatastoreFolderId(v int64)`

SetDatastoreFolderId sets DatastoreFolderId field to given value.

### HasDatastoreFolderId

`func (o *VmwareRestoreParameters) HasDatastoreFolderId() bool`

HasDatastoreFolderId returns a boolean if a field has been set.

### SetDatastoreFolderIdNil

`func (o *VmwareRestoreParameters) SetDatastoreFolderIdNil(b bool)`

 SetDatastoreFolderIdNil sets the value for DatastoreFolderId to be an explicit nil

### UnsetDatastoreFolderId
`func (o *VmwareRestoreParameters) UnsetDatastoreFolderId()`

UnsetDatastoreFolderId ensures that no value is present for DatastoreFolderId, not even an explicit nil
### GetDatastoreId

`func (o *VmwareRestoreParameters) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *VmwareRestoreParameters) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *VmwareRestoreParameters) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *VmwareRestoreParameters) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *VmwareRestoreParameters) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *VmwareRestoreParameters) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDetachNetwork

`func (o *VmwareRestoreParameters) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *VmwareRestoreParameters) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *VmwareRestoreParameters) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *VmwareRestoreParameters) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *VmwareRestoreParameters) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *VmwareRestoreParameters) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetDisableNetwork

`func (o *VmwareRestoreParameters) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *VmwareRestoreParameters) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *VmwareRestoreParameters) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *VmwareRestoreParameters) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *VmwareRestoreParameters) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *VmwareRestoreParameters) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetNetworkId

`func (o *VmwareRestoreParameters) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VmwareRestoreParameters) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VmwareRestoreParameters) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VmwareRestoreParameters) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *VmwareRestoreParameters) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *VmwareRestoreParameters) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetNetworkMappings

`func (o *VmwareRestoreParameters) GetNetworkMappings() []NetworkMapping`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *VmwareRestoreParameters) GetNetworkMappingsOk() (*[]NetworkMapping, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *VmwareRestoreParameters) SetNetworkMappings(v []NetworkMapping)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *VmwareRestoreParameters) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *VmwareRestoreParameters) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *VmwareRestoreParameters) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetPoweredOn

`func (o *VmwareRestoreParameters) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *VmwareRestoreParameters) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *VmwareRestoreParameters) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *VmwareRestoreParameters) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.

### SetPoweredOnNil

`func (o *VmwareRestoreParameters) SetPoweredOnNil(b bool)`

 SetPoweredOnNil sets the value for PoweredOn to be an explicit nil

### UnsetPoweredOn
`func (o *VmwareRestoreParameters) UnsetPoweredOn()`

UnsetPoweredOn ensures that no value is present for PoweredOn, not even an explicit nil
### GetPrefix

`func (o *VmwareRestoreParameters) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VmwareRestoreParameters) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VmwareRestoreParameters) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VmwareRestoreParameters) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *VmwareRestoreParameters) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *VmwareRestoreParameters) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetPreserveCustomAttributesDuringClone

`func (o *VmwareRestoreParameters) GetPreserveCustomAttributesDuringClone() bool`

GetPreserveCustomAttributesDuringClone returns the PreserveCustomAttributesDuringClone field if non-nil, zero value otherwise.

### GetPreserveCustomAttributesDuringCloneOk

`func (o *VmwareRestoreParameters) GetPreserveCustomAttributesDuringCloneOk() (*bool, bool)`

GetPreserveCustomAttributesDuringCloneOk returns a tuple with the PreserveCustomAttributesDuringClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveCustomAttributesDuringClone

`func (o *VmwareRestoreParameters) SetPreserveCustomAttributesDuringClone(v bool)`

SetPreserveCustomAttributesDuringClone sets PreserveCustomAttributesDuringClone field to given value.

### HasPreserveCustomAttributesDuringClone

`func (o *VmwareRestoreParameters) HasPreserveCustomAttributesDuringClone() bool`

HasPreserveCustomAttributesDuringClone returns a boolean if a field has been set.

### SetPreserveCustomAttributesDuringCloneNil

`func (o *VmwareRestoreParameters) SetPreserveCustomAttributesDuringCloneNil(b bool)`

 SetPreserveCustomAttributesDuringCloneNil sets the value for PreserveCustomAttributesDuringClone to be an explicit nil

### UnsetPreserveCustomAttributesDuringClone
`func (o *VmwareRestoreParameters) UnsetPreserveCustomAttributesDuringClone()`

UnsetPreserveCustomAttributesDuringClone ensures that no value is present for PreserveCustomAttributesDuringClone, not even an explicit nil
### GetPreserveTags

`func (o *VmwareRestoreParameters) GetPreserveTags() bool`

GetPreserveTags returns the PreserveTags field if non-nil, zero value otherwise.

### GetPreserveTagsOk

`func (o *VmwareRestoreParameters) GetPreserveTagsOk() (*bool, bool)`

GetPreserveTagsOk returns a tuple with the PreserveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTags

`func (o *VmwareRestoreParameters) SetPreserveTags(v bool)`

SetPreserveTags sets PreserveTags field to given value.

### HasPreserveTags

`func (o *VmwareRestoreParameters) HasPreserveTags() bool`

HasPreserveTags returns a boolean if a field has been set.

### SetPreserveTagsNil

`func (o *VmwareRestoreParameters) SetPreserveTagsNil(b bool)`

 SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil

### UnsetPreserveTags
`func (o *VmwareRestoreParameters) UnsetPreserveTags()`

UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
### GetRecoveryProcessType

`func (o *VmwareRestoreParameters) GetRecoveryProcessType() string`

GetRecoveryProcessType returns the RecoveryProcessType field if non-nil, zero value otherwise.

### GetRecoveryProcessTypeOk

`func (o *VmwareRestoreParameters) GetRecoveryProcessTypeOk() (*string, bool)`

GetRecoveryProcessTypeOk returns a tuple with the RecoveryProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryProcessType

`func (o *VmwareRestoreParameters) SetRecoveryProcessType(v string)`

SetRecoveryProcessType sets RecoveryProcessType field to given value.

### HasRecoveryProcessType

`func (o *VmwareRestoreParameters) HasRecoveryProcessType() bool`

HasRecoveryProcessType returns a boolean if a field has been set.

### SetRecoveryProcessTypeNil

`func (o *VmwareRestoreParameters) SetRecoveryProcessTypeNil(b bool)`

 SetRecoveryProcessTypeNil sets the value for RecoveryProcessType to be an explicit nil

### UnsetRecoveryProcessType
`func (o *VmwareRestoreParameters) UnsetRecoveryProcessType()`

UnsetRecoveryProcessType ensures that no value is present for RecoveryProcessType, not even an explicit nil
### GetResourcePoolId

`func (o *VmwareRestoreParameters) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *VmwareRestoreParameters) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *VmwareRestoreParameters) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *VmwareRestoreParameters) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *VmwareRestoreParameters) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *VmwareRestoreParameters) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetStorageProfileName

`func (o *VmwareRestoreParameters) GetStorageProfileName() string`

GetStorageProfileName returns the StorageProfileName field if non-nil, zero value otherwise.

### GetStorageProfileNameOk

`func (o *VmwareRestoreParameters) GetStorageProfileNameOk() (*string, bool)`

GetStorageProfileNameOk returns a tuple with the StorageProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileName

`func (o *VmwareRestoreParameters) SetStorageProfileName(v string)`

SetStorageProfileName sets StorageProfileName field to given value.

### HasStorageProfileName

`func (o *VmwareRestoreParameters) HasStorageProfileName() bool`

HasStorageProfileName returns a boolean if a field has been set.

### SetStorageProfileNameNil

`func (o *VmwareRestoreParameters) SetStorageProfileNameNil(b bool)`

 SetStorageProfileNameNil sets the value for StorageProfileName to be an explicit nil

### UnsetStorageProfileName
`func (o *VmwareRestoreParameters) UnsetStorageProfileName()`

UnsetStorageProfileName ensures that no value is present for StorageProfileName, not even an explicit nil
### GetStorageProfileVcdUuid

`func (o *VmwareRestoreParameters) GetStorageProfileVcdUuid() string`

GetStorageProfileVcdUuid returns the StorageProfileVcdUuid field if non-nil, zero value otherwise.

### GetStorageProfileVcdUuidOk

`func (o *VmwareRestoreParameters) GetStorageProfileVcdUuidOk() (*string, bool)`

GetStorageProfileVcdUuidOk returns a tuple with the StorageProfileVcdUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileVcdUuid

`func (o *VmwareRestoreParameters) SetStorageProfileVcdUuid(v string)`

SetStorageProfileVcdUuid sets StorageProfileVcdUuid field to given value.

### HasStorageProfileVcdUuid

`func (o *VmwareRestoreParameters) HasStorageProfileVcdUuid() bool`

HasStorageProfileVcdUuid returns a boolean if a field has been set.

### SetStorageProfileVcdUuidNil

`func (o *VmwareRestoreParameters) SetStorageProfileVcdUuidNil(b bool)`

 SetStorageProfileVcdUuidNil sets the value for StorageProfileVcdUuid to be an explicit nil

### UnsetStorageProfileVcdUuid
`func (o *VmwareRestoreParameters) UnsetStorageProfileVcdUuid()`

UnsetStorageProfileVcdUuid ensures that no value is present for StorageProfileVcdUuid, not even an explicit nil
### GetSuffix

`func (o *VmwareRestoreParameters) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *VmwareRestoreParameters) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *VmwareRestoreParameters) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *VmwareRestoreParameters) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *VmwareRestoreParameters) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *VmwareRestoreParameters) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetVAppId

`func (o *VmwareRestoreParameters) GetVAppId() int64`

GetVAppId returns the VAppId field if non-nil, zero value otherwise.

### GetVAppIdOk

`func (o *VmwareRestoreParameters) GetVAppIdOk() (*int64, bool)`

GetVAppIdOk returns a tuple with the VAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAppId

`func (o *VmwareRestoreParameters) SetVAppId(v int64)`

SetVAppId sets VAppId field to given value.

### HasVAppId

`func (o *VmwareRestoreParameters) HasVAppId() bool`

HasVAppId returns a boolean if a field has been set.

### SetVAppIdNil

`func (o *VmwareRestoreParameters) SetVAppIdNil(b bool)`

 SetVAppIdNil sets the value for VAppId to be an explicit nil

### UnsetVAppId
`func (o *VmwareRestoreParameters) UnsetVAppId()`

UnsetVAppId ensures that no value is present for VAppId, not even an explicit nil
### GetVdcId

`func (o *VmwareRestoreParameters) GetVdcId() int64`

GetVdcId returns the VdcId field if non-nil, zero value otherwise.

### GetVdcIdOk

`func (o *VmwareRestoreParameters) GetVdcIdOk() (*int64, bool)`

GetVdcIdOk returns a tuple with the VdcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcId

`func (o *VmwareRestoreParameters) SetVdcId(v int64)`

SetVdcId sets VdcId field to given value.

### HasVdcId

`func (o *VmwareRestoreParameters) HasVdcId() bool`

HasVdcId returns a boolean if a field has been set.

### SetVdcIdNil

`func (o *VmwareRestoreParameters) SetVdcIdNil(b bool)`

 SetVdcIdNil sets the value for VdcId to be an explicit nil

### UnsetVdcId
`func (o *VmwareRestoreParameters) UnsetVdcId()`

UnsetVdcId ensures that no value is present for VdcId, not even an explicit nil
### GetVmFolderId

`func (o *VmwareRestoreParameters) GetVmFolderId() int64`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *VmwareRestoreParameters) GetVmFolderIdOk() (*int64, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *VmwareRestoreParameters) SetVmFolderId(v int64)`

SetVmFolderId sets VmFolderId field to given value.

### HasVmFolderId

`func (o *VmwareRestoreParameters) HasVmFolderId() bool`

HasVmFolderId returns a boolean if a field has been set.

### SetVmFolderIdNil

`func (o *VmwareRestoreParameters) SetVmFolderIdNil(b bool)`

 SetVmFolderIdNil sets the value for VmFolderId to be an explicit nil

### UnsetVmFolderId
`func (o *VmwareRestoreParameters) UnsetVmFolderId()`

UnsetVmFolderId ensures that no value is present for VmFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


