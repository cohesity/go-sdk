# RestoreVMwareVMParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogUuid** | Pointer to **NullableString** | Specifies the name of the catalog for vapp template recovery. This is applicable for recovery to a VCD. | [optional] 
**CopyRecovery** | Pointer to **NullableBool** | Whether to perform copy recovery instead of instant recovery. | [optional] 
**DatastoreEntityVec** | Pointer to [**[]EntityProto**](EntityProto.md) | Datastore entities if the restore is to alternate location. | [optional] 
**PreserveCustomAttributesDuringClone** | Pointer to **NullableBool** | Whether to preserve custom attributes for the clone op. | [optional] 
**PreserveTagsDuringClone** | Pointer to **NullableBool** | Whether to preserve tags for the clone op. | [optional] 
**ResourcePoolEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**StorageProfileName** | Pointer to **NullableString** | This is only populated for VCD restore to alternate location. It contains the name of the destination storage profile. | [optional] 
**StorageProfileVcdUuid** | Pointer to **NullableString** | This is only populated for VCD restore to alternate location. It contains the vcd uuid of the destination storage profile. | [optional] 
**TargetDatastoreFolder** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetVmFolder** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreVMwareVMParams

`func NewRestoreVMwareVMParams() *RestoreVMwareVMParams`

NewRestoreVMwareVMParams instantiates a new RestoreVMwareVMParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreVMwareVMParamsWithDefaults

`func NewRestoreVMwareVMParamsWithDefaults() *RestoreVMwareVMParams`

NewRestoreVMwareVMParamsWithDefaults instantiates a new RestoreVMwareVMParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogUuid

`func (o *RestoreVMwareVMParams) GetCatalogUuid() string`

GetCatalogUuid returns the CatalogUuid field if non-nil, zero value otherwise.

### GetCatalogUuidOk

`func (o *RestoreVMwareVMParams) GetCatalogUuidOk() (*string, bool)`

GetCatalogUuidOk returns a tuple with the CatalogUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUuid

`func (o *RestoreVMwareVMParams) SetCatalogUuid(v string)`

SetCatalogUuid sets CatalogUuid field to given value.

### HasCatalogUuid

`func (o *RestoreVMwareVMParams) HasCatalogUuid() bool`

HasCatalogUuid returns a boolean if a field has been set.

### SetCatalogUuidNil

`func (o *RestoreVMwareVMParams) SetCatalogUuidNil(b bool)`

 SetCatalogUuidNil sets the value for CatalogUuid to be an explicit nil

### UnsetCatalogUuid
`func (o *RestoreVMwareVMParams) UnsetCatalogUuid()`

UnsetCatalogUuid ensures that no value is present for CatalogUuid, not even an explicit nil
### GetCopyRecovery

`func (o *RestoreVMwareVMParams) GetCopyRecovery() bool`

GetCopyRecovery returns the CopyRecovery field if non-nil, zero value otherwise.

### GetCopyRecoveryOk

`func (o *RestoreVMwareVMParams) GetCopyRecoveryOk() (*bool, bool)`

GetCopyRecoveryOk returns a tuple with the CopyRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRecovery

`func (o *RestoreVMwareVMParams) SetCopyRecovery(v bool)`

SetCopyRecovery sets CopyRecovery field to given value.

### HasCopyRecovery

`func (o *RestoreVMwareVMParams) HasCopyRecovery() bool`

HasCopyRecovery returns a boolean if a field has been set.

### SetCopyRecoveryNil

`func (o *RestoreVMwareVMParams) SetCopyRecoveryNil(b bool)`

 SetCopyRecoveryNil sets the value for CopyRecovery to be an explicit nil

### UnsetCopyRecovery
`func (o *RestoreVMwareVMParams) UnsetCopyRecovery()`

UnsetCopyRecovery ensures that no value is present for CopyRecovery, not even an explicit nil
### GetDatastoreEntityVec

`func (o *RestoreVMwareVMParams) GetDatastoreEntityVec() []EntityProto`

GetDatastoreEntityVec returns the DatastoreEntityVec field if non-nil, zero value otherwise.

### GetDatastoreEntityVecOk

`func (o *RestoreVMwareVMParams) GetDatastoreEntityVecOk() (*[]EntityProto, bool)`

GetDatastoreEntityVecOk returns a tuple with the DatastoreEntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEntityVec

`func (o *RestoreVMwareVMParams) SetDatastoreEntityVec(v []EntityProto)`

SetDatastoreEntityVec sets DatastoreEntityVec field to given value.

### HasDatastoreEntityVec

`func (o *RestoreVMwareVMParams) HasDatastoreEntityVec() bool`

HasDatastoreEntityVec returns a boolean if a field has been set.

### SetDatastoreEntityVecNil

`func (o *RestoreVMwareVMParams) SetDatastoreEntityVecNil(b bool)`

 SetDatastoreEntityVecNil sets the value for DatastoreEntityVec to be an explicit nil

### UnsetDatastoreEntityVec
`func (o *RestoreVMwareVMParams) UnsetDatastoreEntityVec()`

UnsetDatastoreEntityVec ensures that no value is present for DatastoreEntityVec, not even an explicit nil
### GetPreserveCustomAttributesDuringClone

`func (o *RestoreVMwareVMParams) GetPreserveCustomAttributesDuringClone() bool`

GetPreserveCustomAttributesDuringClone returns the PreserveCustomAttributesDuringClone field if non-nil, zero value otherwise.

### GetPreserveCustomAttributesDuringCloneOk

`func (o *RestoreVMwareVMParams) GetPreserveCustomAttributesDuringCloneOk() (*bool, bool)`

GetPreserveCustomAttributesDuringCloneOk returns a tuple with the PreserveCustomAttributesDuringClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveCustomAttributesDuringClone

`func (o *RestoreVMwareVMParams) SetPreserveCustomAttributesDuringClone(v bool)`

SetPreserveCustomAttributesDuringClone sets PreserveCustomAttributesDuringClone field to given value.

### HasPreserveCustomAttributesDuringClone

`func (o *RestoreVMwareVMParams) HasPreserveCustomAttributesDuringClone() bool`

HasPreserveCustomAttributesDuringClone returns a boolean if a field has been set.

### SetPreserveCustomAttributesDuringCloneNil

`func (o *RestoreVMwareVMParams) SetPreserveCustomAttributesDuringCloneNil(b bool)`

 SetPreserveCustomAttributesDuringCloneNil sets the value for PreserveCustomAttributesDuringClone to be an explicit nil

### UnsetPreserveCustomAttributesDuringClone
`func (o *RestoreVMwareVMParams) UnsetPreserveCustomAttributesDuringClone()`

UnsetPreserveCustomAttributesDuringClone ensures that no value is present for PreserveCustomAttributesDuringClone, not even an explicit nil
### GetPreserveTagsDuringClone

`func (o *RestoreVMwareVMParams) GetPreserveTagsDuringClone() bool`

GetPreserveTagsDuringClone returns the PreserveTagsDuringClone field if non-nil, zero value otherwise.

### GetPreserveTagsDuringCloneOk

`func (o *RestoreVMwareVMParams) GetPreserveTagsDuringCloneOk() (*bool, bool)`

GetPreserveTagsDuringCloneOk returns a tuple with the PreserveTagsDuringClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTagsDuringClone

`func (o *RestoreVMwareVMParams) SetPreserveTagsDuringClone(v bool)`

SetPreserveTagsDuringClone sets PreserveTagsDuringClone field to given value.

### HasPreserveTagsDuringClone

`func (o *RestoreVMwareVMParams) HasPreserveTagsDuringClone() bool`

HasPreserveTagsDuringClone returns a boolean if a field has been set.

### SetPreserveTagsDuringCloneNil

`func (o *RestoreVMwareVMParams) SetPreserveTagsDuringCloneNil(b bool)`

 SetPreserveTagsDuringCloneNil sets the value for PreserveTagsDuringClone to be an explicit nil

### UnsetPreserveTagsDuringClone
`func (o *RestoreVMwareVMParams) UnsetPreserveTagsDuringClone()`

UnsetPreserveTagsDuringClone ensures that no value is present for PreserveTagsDuringClone, not even an explicit nil
### GetResourcePoolEntity

`func (o *RestoreVMwareVMParams) GetResourcePoolEntity() EntityProto`

GetResourcePoolEntity returns the ResourcePoolEntity field if non-nil, zero value otherwise.

### GetResourcePoolEntityOk

`func (o *RestoreVMwareVMParams) GetResourcePoolEntityOk() (*EntityProto, bool)`

GetResourcePoolEntityOk returns a tuple with the ResourcePoolEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolEntity

`func (o *RestoreVMwareVMParams) SetResourcePoolEntity(v EntityProto)`

SetResourcePoolEntity sets ResourcePoolEntity field to given value.

### HasResourcePoolEntity

`func (o *RestoreVMwareVMParams) HasResourcePoolEntity() bool`

HasResourcePoolEntity returns a boolean if a field has been set.

### GetStorageProfileName

`func (o *RestoreVMwareVMParams) GetStorageProfileName() string`

GetStorageProfileName returns the StorageProfileName field if non-nil, zero value otherwise.

### GetStorageProfileNameOk

`func (o *RestoreVMwareVMParams) GetStorageProfileNameOk() (*string, bool)`

GetStorageProfileNameOk returns a tuple with the StorageProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileName

`func (o *RestoreVMwareVMParams) SetStorageProfileName(v string)`

SetStorageProfileName sets StorageProfileName field to given value.

### HasStorageProfileName

`func (o *RestoreVMwareVMParams) HasStorageProfileName() bool`

HasStorageProfileName returns a boolean if a field has been set.

### SetStorageProfileNameNil

`func (o *RestoreVMwareVMParams) SetStorageProfileNameNil(b bool)`

 SetStorageProfileNameNil sets the value for StorageProfileName to be an explicit nil

### UnsetStorageProfileName
`func (o *RestoreVMwareVMParams) UnsetStorageProfileName()`

UnsetStorageProfileName ensures that no value is present for StorageProfileName, not even an explicit nil
### GetStorageProfileVcdUuid

`func (o *RestoreVMwareVMParams) GetStorageProfileVcdUuid() string`

GetStorageProfileVcdUuid returns the StorageProfileVcdUuid field if non-nil, zero value otherwise.

### GetStorageProfileVcdUuidOk

`func (o *RestoreVMwareVMParams) GetStorageProfileVcdUuidOk() (*string, bool)`

GetStorageProfileVcdUuidOk returns a tuple with the StorageProfileVcdUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileVcdUuid

`func (o *RestoreVMwareVMParams) SetStorageProfileVcdUuid(v string)`

SetStorageProfileVcdUuid sets StorageProfileVcdUuid field to given value.

### HasStorageProfileVcdUuid

`func (o *RestoreVMwareVMParams) HasStorageProfileVcdUuid() bool`

HasStorageProfileVcdUuid returns a boolean if a field has been set.

### SetStorageProfileVcdUuidNil

`func (o *RestoreVMwareVMParams) SetStorageProfileVcdUuidNil(b bool)`

 SetStorageProfileVcdUuidNil sets the value for StorageProfileVcdUuid to be an explicit nil

### UnsetStorageProfileVcdUuid
`func (o *RestoreVMwareVMParams) UnsetStorageProfileVcdUuid()`

UnsetStorageProfileVcdUuid ensures that no value is present for StorageProfileVcdUuid, not even an explicit nil
### GetTargetDatastoreFolder

`func (o *RestoreVMwareVMParams) GetTargetDatastoreFolder() EntityProto`

GetTargetDatastoreFolder returns the TargetDatastoreFolder field if non-nil, zero value otherwise.

### GetTargetDatastoreFolderOk

`func (o *RestoreVMwareVMParams) GetTargetDatastoreFolderOk() (*EntityProto, bool)`

GetTargetDatastoreFolderOk returns a tuple with the TargetDatastoreFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastoreFolder

`func (o *RestoreVMwareVMParams) SetTargetDatastoreFolder(v EntityProto)`

SetTargetDatastoreFolder sets TargetDatastoreFolder field to given value.

### HasTargetDatastoreFolder

`func (o *RestoreVMwareVMParams) HasTargetDatastoreFolder() bool`

HasTargetDatastoreFolder returns a boolean if a field has been set.

### GetTargetVmFolder

`func (o *RestoreVMwareVMParams) GetTargetVmFolder() EntityProto`

GetTargetVmFolder returns the TargetVmFolder field if non-nil, zero value otherwise.

### GetTargetVmFolderOk

`func (o *RestoreVMwareVMParams) GetTargetVmFolderOk() (*EntityProto, bool)`

GetTargetVmFolderOk returns a tuple with the TargetVmFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmFolder

`func (o *RestoreVMwareVMParams) SetTargetVmFolder(v EntityProto)`

SetTargetVmFolder sets TargetVmFolder field to given value.

### HasTargetVmFolder

`func (o *RestoreVMwareVMParams) HasTargetVmFolder() bool`

HasTargetVmFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


