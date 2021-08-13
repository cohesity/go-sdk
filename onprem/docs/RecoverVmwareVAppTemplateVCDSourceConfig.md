# RecoverVmwareVAppTemplateVCDSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the vApp templates. | 
**Vdc** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the VDC object where the recovered objects will be attached. | 
**Catalog** | [**NullableVdcCatalog**](VdcCatalog.md) | Specifies the catalog where the vApp template should reside. | 
**Datastores** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore objects where the object&#39;s files should be recovered to. | [optional] 
**StorageProfile** | Pointer to [**VcdStorageProfileParams**](VcdStorageProfileParams.md) | Specifies the storage profile to which the objects should be recovered. This should only be specified if datastores are not specified. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmNewSourceNetworkConfig**](RecoverVmwareVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered vApp templates. | [optional] 
**OrgVdcNetwork** | Pointer to [**OrgVDCNetwork**](OrgVDCNetwork.md) |  | [optional] 

## Methods

### NewRecoverVmwareVAppTemplateVCDSourceConfig

`func NewRecoverVmwareVAppTemplateVCDSourceConfig(source NullableRecoveryObjectIdentifier, vdc NullableRecoveryObjectIdentifier, catalog NullableVdcCatalog, ) *RecoverVmwareVAppTemplateVCDSourceConfig`

NewRecoverVmwareVAppTemplateVCDSourceConfig instantiates a new RecoverVmwareVAppTemplateVCDSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppTemplateVCDSourceConfigWithDefaults

`func NewRecoverVmwareVAppTemplateVCDSourceConfigWithDefaults() *RecoverVmwareVAppTemplateVCDSourceConfig`

NewRecoverVmwareVAppTemplateVCDSourceConfigWithDefaults instantiates a new RecoverVmwareVAppTemplateVCDSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetVdc

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetVdc() RecoveryObjectIdentifier`

GetVdc returns the Vdc field if non-nil, zero value otherwise.

### GetVdcOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetVdcOk() (*RecoveryObjectIdentifier, bool)`

GetVdcOk returns a tuple with the Vdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdc

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetVdc(v RecoveryObjectIdentifier)`

SetVdc sets Vdc field to given value.


### SetVdcNil

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetVdcNil(b bool)`

 SetVdcNil sets the value for Vdc to be an explicit nil

### UnsetVdc
`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) UnsetVdc()`

UnsetVdc ensures that no value is present for Vdc, not even an explicit nil
### GetCatalog

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetCatalog() VdcCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetCatalogOk() (*VdcCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetCatalog(v VdcCatalog)`

SetCatalog sets Catalog field to given value.


### SetCatalogNil

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetDatastores

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetDatastores() []RecoveryObjectIdentifier`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetDatastoresOk() (*[]RecoveryObjectIdentifier, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetDatastores(v []RecoveryObjectIdentifier)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetStorageProfile

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetStorageProfile() VcdStorageProfileParams`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetStorageProfileOk() (*VcdStorageProfileParams, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetStorageProfile(v VcdStorageProfileParams)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### GetNetworkConfig

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetNetworkConfig() RecoverVmwareVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetOrgVdcNetwork

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetOrgVdcNetwork() OrgVDCNetwork`

GetOrgVdcNetwork returns the OrgVdcNetwork field if non-nil, zero value otherwise.

### GetOrgVdcNetworkOk

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) GetOrgVdcNetworkOk() (*OrgVDCNetwork, bool)`

GetOrgVdcNetworkOk returns a tuple with the OrgVdcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgVdcNetwork

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) SetOrgVdcNetwork(v OrgVDCNetwork)`

SetOrgVdcNetwork sets OrgVdcNetwork field to given value.

### HasOrgVdcNetwork

`func (o *RecoverVmwareVAppTemplateVCDSourceConfig) HasOrgVdcNetwork() bool`

HasOrgVdcNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


