# KubernetesTargetParamsForRecoverKubernetesNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the objects to be recovered. | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**RenameRecoveredNamespacesParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the Namespaces that are recovered. If not specified, the original names of the Namespaces are preserved. If a name collision occurs then the Namespace being recovered will overwrite the Namespace already present on the source. | [optional] 
**RecoveryTargetConfig** | [**NullableKubernetesNamespaceRecoveryTargetConfig**](KubernetesNamespaceRecoveryTargetConfig.md) | Specifies the recovery target configuration of the Namespace recovery. | 

## Methods

### NewKubernetesTargetParamsForRecoverKubernetesNamespace

`func NewKubernetesTargetParamsForRecoverKubernetesNamespace(recoveryTargetConfig NullableKubernetesNamespaceRecoveryTargetConfig, ) *KubernetesTargetParamsForRecoverKubernetesNamespace`

NewKubernetesTargetParamsForRecoverKubernetesNamespace instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults

`func NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults() *KubernetesTargetParamsForRecoverKubernetesNamespace`

NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetRenameRecoveredNamespacesParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredNamespacesParams returns the RenameRecoveredNamespacesParams field if non-nil, zero value otherwise.

### GetRenameRecoveredNamespacesParamsOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredNamespacesParamsOk returns a tuple with the RenameRecoveredNamespacesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredNamespacesParams

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRenameRecoveredNamespacesParams(v RecoveredOrClonedVmsRenameConfig)`

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
### GetRecoveryTargetConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfig() KubernetesNamespaceRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfigOk() (*KubernetesNamespaceRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoveryTargetConfig(v KubernetesNamespaceRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.


### SetRecoveryTargetConfigNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


