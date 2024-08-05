# KubernetesTargetParamsForRecoverKubernetesNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeParams** | Pointer to [**NullableKubernetesFilterParams**](KubernetesFilterParams.md) |  | [optional] 
**ExcludedPvcs** | Pointer to [**[]KubernetesPvcInfo**](KubernetesPvcInfo.md) | Specifies the list of pvc to be excluded from recovery. | [optional] 
**IncludeParams** | Pointer to [**NullableKubernetesFilterParams**](KubernetesFilterParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the objects to be recovered. | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**RecoverPvcsOnly** | Pointer to **NullableBool** | Specifies whether to recover PVCs only during recovery. | [optional] 
**RecoveryTargetConfig** | [**NullableKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig**](KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig.md) |  | 
**RenameRecoveredNamespacesParams** | Pointer to [**NullableKubernetesTargetParamsForRecoverKubernetesNamespaceRenameRecoveredNamespacesParams**](KubernetesTargetParamsForRecoverKubernetesNamespaceRenameRecoveredNamespacesParams.md) |  | [optional] 
**StorageClass** | Pointer to [**KubernetesStorageClassParams**](KubernetesStorageClassParams.md) |  | [optional] 

## Methods

### NewKubernetesTargetParamsForRecoverKubernetesNamespace

`func NewKubernetesTargetParamsForRecoverKubernetesNamespace(recoveryTargetConfig NullableKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig, ) *KubernetesTargetParamsForRecoverKubernetesNamespace`

NewKubernetesTargetParamsForRecoverKubernetesNamespace instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults

`func NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults() *KubernetesTargetParamsForRecoverKubernetesNamespace`

NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetExcludeParams() KubernetesFilterParams`

GetExcludeParams returns the ExcludeParams field if non-nil, zero value otherwise.

### GetExcludeParamsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetExcludeParamsOk() (*KubernetesFilterParams, bool)`

GetExcludeParamsOk returns a tuple with the ExcludeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetExcludeParams(v KubernetesFilterParams)`

SetExcludeParams sets ExcludeParams field to given value.

### HasExcludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasExcludeParams() bool`

HasExcludeParams returns a boolean if a field has been set.

### SetExcludeParamsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetExcludeParamsNil(b bool)`

 SetExcludeParamsNil sets the value for ExcludeParams to be an explicit nil

### UnsetExcludeParams
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetExcludeParams()`

UnsetExcludeParams ensures that no value is present for ExcludeParams, not even an explicit nil
### GetExcludedPvcs

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetExcludedPvcs() []KubernetesPvcInfo`

GetExcludedPvcs returns the ExcludedPvcs field if non-nil, zero value otherwise.

### GetExcludedPvcsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetExcludedPvcsOk() (*[]KubernetesPvcInfo, bool)`

GetExcludedPvcsOk returns a tuple with the ExcludedPvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPvcs

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetExcludedPvcs(v []KubernetesPvcInfo)`

SetExcludedPvcs sets ExcludedPvcs field to given value.

### HasExcludedPvcs

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasExcludedPvcs() bool`

HasExcludedPvcs returns a boolean if a field has been set.

### SetExcludedPvcsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetExcludedPvcsNil(b bool)`

 SetExcludedPvcsNil sets the value for ExcludedPvcs to be an explicit nil

### UnsetExcludedPvcs
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetExcludedPvcs()`

UnsetExcludedPvcs ensures that no value is present for ExcludedPvcs, not even an explicit nil
### GetIncludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetIncludeParams() KubernetesFilterParams`

GetIncludeParams returns the IncludeParams field if non-nil, zero value otherwise.

### GetIncludeParamsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetIncludeParamsOk() (*KubernetesFilterParams, bool)`

GetIncludeParamsOk returns a tuple with the IncludeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetIncludeParams(v KubernetesFilterParams)`

SetIncludeParams sets IncludeParams field to given value.

### HasIncludeParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasIncludeParams() bool`

HasIncludeParams returns a boolean if a field has been set.

### SetIncludeParamsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetIncludeParamsNil(b bool)`

 SetIncludeParamsNil sets the value for IncludeParams to be an explicit nil

### UnsetIncludeParams
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetIncludeParams()`

UnsetIncludeParams ensures that no value is present for IncludeParams, not even an explicit nil
### GetObjects

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetRecoverPvcsOnly

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverPvcsOnly() bool`

GetRecoverPvcsOnly returns the RecoverPvcsOnly field if non-nil, zero value otherwise.

### GetRecoverPvcsOnlyOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverPvcsOnlyOk() (*bool, bool)`

GetRecoverPvcsOnlyOk returns a tuple with the RecoverPvcsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPvcsOnly

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoverPvcsOnly(v bool)`

SetRecoverPvcsOnly sets RecoverPvcsOnly field to given value.

### HasRecoverPvcsOnly

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasRecoverPvcsOnly() bool`

HasRecoverPvcsOnly returns a boolean if a field has been set.

### SetRecoverPvcsOnlyNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoverPvcsOnlyNil(b bool)`

 SetRecoverPvcsOnlyNil sets the value for RecoverPvcsOnly to be an explicit nil

### UnsetRecoverPvcsOnly
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRecoverPvcsOnly()`

UnsetRecoverPvcsOnly ensures that no value is present for RecoverPvcsOnly, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfig() KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfigOk() (*KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoveryTargetConfig(v KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.


### SetRecoveryTargetConfigNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredNamespacesParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParams() KubernetesTargetParamsForRecoverKubernetesNamespaceRenameRecoveredNamespacesParams`

GetRenameRecoveredNamespacesParams returns the RenameRecoveredNamespacesParams field if non-nil, zero value otherwise.

### GetRenameRecoveredNamespacesParamsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParamsOk() (*KubernetesTargetParamsForRecoverKubernetesNamespaceRenameRecoveredNamespacesParams, bool)`

GetRenameRecoveredNamespacesParamsOk returns a tuple with the RenameRecoveredNamespacesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredNamespacesParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRenameRecoveredNamespacesParams(v KubernetesTargetParamsForRecoverKubernetesNamespaceRenameRecoveredNamespacesParams)`

SetRenameRecoveredNamespacesParams sets RenameRecoveredNamespacesParams field to given value.

### HasRenameRecoveredNamespacesParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasRenameRecoveredNamespacesParams() bool`

HasRenameRecoveredNamespacesParams returns a boolean if a field has been set.

### SetRenameRecoveredNamespacesParamsNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRenameRecoveredNamespacesParamsNil(b bool)`

 SetRenameRecoveredNamespacesParamsNil sets the value for RenameRecoveredNamespacesParams to be an explicit nil

### UnsetRenameRecoveredNamespacesParams
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRenameRecoveredNamespacesParams()`

UnsetRenameRecoveredNamespacesParams ensures that no value is present for RenameRecoveredNamespacesParams, not even an explicit nil
### GetStorageClass

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetStorageClass() KubernetesStorageClassParams`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetStorageClassOk() (*KubernetesStorageClassParams, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetStorageClass(v KubernetesStorageClassParams)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


