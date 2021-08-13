# RecoverVmwareVmVCDSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Vdc** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the VDC object where the recovered objects will be attached. | 
**VApp** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the vApp object where the recovered objects will be attached. | [optional] 
**Datastores** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore objects where the object&#39;s files should be recovered to. This should only be specified if storageProfile is not specified. | [optional] 
**StorageProfile** | Pointer to [**VcdStorageProfileParams**](VcdStorageProfileParams.md) | Specifies the storage profile to which the objects should be recovered. This should only be specified if datastores are not specified. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVmNewSourceNetworkConfig**](RecoverVmwareVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 
**OrgVdcNetwork** | Pointer to [**OrgVDCNetwork**](OrgVDCNetwork.md) |  | [optional] 

## Methods

### NewRecoverVmwareVmVCDSourceConfig

`func NewRecoverVmwareVmVCDSourceConfig(source NullableRecoveryObjectIdentifier, vdc NullableRecoveryObjectIdentifier, ) *RecoverVmwareVmVCDSourceConfig`

NewRecoverVmwareVmVCDSourceConfig instantiates a new RecoverVmwareVmVCDSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmVCDSourceConfigWithDefaults

`func NewRecoverVmwareVmVCDSourceConfigWithDefaults() *RecoverVmwareVmVCDSourceConfig`

NewRecoverVmwareVmVCDSourceConfigWithDefaults instantiates a new RecoverVmwareVmVCDSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverVmwareVmVCDSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverVmwareVmVCDSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverVmwareVmVCDSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverVmwareVmVCDSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetVdc

`func (o *RecoverVmwareVmVCDSourceConfig) GetVdc() RecoveryObjectIdentifier`

GetVdc returns the Vdc field if non-nil, zero value otherwise.

### GetVdcOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetVdcOk() (*RecoveryObjectIdentifier, bool)`

GetVdcOk returns a tuple with the Vdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdc

`func (o *RecoverVmwareVmVCDSourceConfig) SetVdc(v RecoveryObjectIdentifier)`

SetVdc sets Vdc field to given value.


### SetVdcNil

`func (o *RecoverVmwareVmVCDSourceConfig) SetVdcNil(b bool)`

 SetVdcNil sets the value for Vdc to be an explicit nil

### UnsetVdc
`func (o *RecoverVmwareVmVCDSourceConfig) UnsetVdc()`

UnsetVdc ensures that no value is present for Vdc, not even an explicit nil
### GetVApp

`func (o *RecoverVmwareVmVCDSourceConfig) GetVApp() RecoveryObjectIdentifier`

GetVApp returns the VApp field if non-nil, zero value otherwise.

### GetVAppOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetVAppOk() (*RecoveryObjectIdentifier, bool)`

GetVAppOk returns a tuple with the VApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVApp

`func (o *RecoverVmwareVmVCDSourceConfig) SetVApp(v RecoveryObjectIdentifier)`

SetVApp sets VApp field to given value.

### HasVApp

`func (o *RecoverVmwareVmVCDSourceConfig) HasVApp() bool`

HasVApp returns a boolean if a field has been set.

### SetVAppNil

`func (o *RecoverVmwareVmVCDSourceConfig) SetVAppNil(b bool)`

 SetVAppNil sets the value for VApp to be an explicit nil

### UnsetVApp
`func (o *RecoverVmwareVmVCDSourceConfig) UnsetVApp()`

UnsetVApp ensures that no value is present for VApp, not even an explicit nil
### GetDatastores

`func (o *RecoverVmwareVmVCDSourceConfig) GetDatastores() []RecoveryObjectIdentifier`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetDatastoresOk() (*[]RecoveryObjectIdentifier, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *RecoverVmwareVmVCDSourceConfig) SetDatastores(v []RecoveryObjectIdentifier)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *RecoverVmwareVmVCDSourceConfig) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *RecoverVmwareVmVCDSourceConfig) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *RecoverVmwareVmVCDSourceConfig) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetStorageProfile

`func (o *RecoverVmwareVmVCDSourceConfig) GetStorageProfile() VcdStorageProfileParams`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetStorageProfileOk() (*VcdStorageProfileParams, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *RecoverVmwareVmVCDSourceConfig) SetStorageProfile(v VcdStorageProfileParams)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *RecoverVmwareVmVCDSourceConfig) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### GetNetworkConfig

`func (o *RecoverVmwareVmVCDSourceConfig) GetNetworkConfig() RecoverVmwareVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverVmwareVmVCDSourceConfig) SetNetworkConfig(v RecoverVmwareVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverVmwareVmVCDSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverVmwareVmVCDSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverVmwareVmVCDSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetOrgVdcNetwork

`func (o *RecoverVmwareVmVCDSourceConfig) GetOrgVdcNetwork() OrgVDCNetwork`

GetOrgVdcNetwork returns the OrgVdcNetwork field if non-nil, zero value otherwise.

### GetOrgVdcNetworkOk

`func (o *RecoverVmwareVmVCDSourceConfig) GetOrgVdcNetworkOk() (*OrgVDCNetwork, bool)`

GetOrgVdcNetworkOk returns a tuple with the OrgVdcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgVdcNetwork

`func (o *RecoverVmwareVmVCDSourceConfig) SetOrgVdcNetwork(v OrgVDCNetwork)`

SetOrgVdcNetwork sets OrgVdcNetwork field to given value.

### HasOrgVdcNetwork

`func (o *RecoverVmwareVmVCDSourceConfig) HasOrgVdcNetwork() bool`

HasOrgVdcNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


