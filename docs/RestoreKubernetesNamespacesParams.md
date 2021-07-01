# RestoreKubernetesNamespacesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupJobName** | Pointer to **NullableString** | Backup job that needs to be used for recovering the namespace. | [optional] 
**ClusterEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ManagementNamespace** | Pointer to **NullableString** | Namespace in which restore job will be created in K8s cluster. | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 

## Methods

### NewRestoreKubernetesNamespacesParams

`func NewRestoreKubernetesNamespacesParams() *RestoreKubernetesNamespacesParams`

NewRestoreKubernetesNamespacesParams instantiates a new RestoreKubernetesNamespacesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreKubernetesNamespacesParamsWithDefaults

`func NewRestoreKubernetesNamespacesParamsWithDefaults() *RestoreKubernetesNamespacesParams`

NewRestoreKubernetesNamespacesParamsWithDefaults instantiates a new RestoreKubernetesNamespacesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupJobName

`func (o *RestoreKubernetesNamespacesParams) GetBackupJobName() string`

GetBackupJobName returns the BackupJobName field if non-nil, zero value otherwise.

### GetBackupJobNameOk

`func (o *RestoreKubernetesNamespacesParams) GetBackupJobNameOk() (*string, bool)`

GetBackupJobNameOk returns a tuple with the BackupJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobName

`func (o *RestoreKubernetesNamespacesParams) SetBackupJobName(v string)`

SetBackupJobName sets BackupJobName field to given value.

### HasBackupJobName

`func (o *RestoreKubernetesNamespacesParams) HasBackupJobName() bool`

HasBackupJobName returns a boolean if a field has been set.

### SetBackupJobNameNil

`func (o *RestoreKubernetesNamespacesParams) SetBackupJobNameNil(b bool)`

 SetBackupJobNameNil sets the value for BackupJobName to be an explicit nil

### UnsetBackupJobName
`func (o *RestoreKubernetesNamespacesParams) UnsetBackupJobName()`

UnsetBackupJobName ensures that no value is present for BackupJobName, not even an explicit nil
### GetClusterEntity

`func (o *RestoreKubernetesNamespacesParams) GetClusterEntity() EntityProto`

GetClusterEntity returns the ClusterEntity field if non-nil, zero value otherwise.

### GetClusterEntityOk

`func (o *RestoreKubernetesNamespacesParams) GetClusterEntityOk() (*EntityProto, bool)`

GetClusterEntityOk returns a tuple with the ClusterEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntity

`func (o *RestoreKubernetesNamespacesParams) SetClusterEntity(v EntityProto)`

SetClusterEntity sets ClusterEntity field to given value.

### HasClusterEntity

`func (o *RestoreKubernetesNamespacesParams) HasClusterEntity() bool`

HasClusterEntity returns a boolean if a field has been set.

### GetManagementNamespace

`func (o *RestoreKubernetesNamespacesParams) GetManagementNamespace() string`

GetManagementNamespace returns the ManagementNamespace field if non-nil, zero value otherwise.

### GetManagementNamespaceOk

`func (o *RestoreKubernetesNamespacesParams) GetManagementNamespaceOk() (*string, bool)`

GetManagementNamespaceOk returns a tuple with the ManagementNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNamespace

`func (o *RestoreKubernetesNamespacesParams) SetManagementNamespace(v string)`

SetManagementNamespace sets ManagementNamespace field to given value.

### HasManagementNamespace

`func (o *RestoreKubernetesNamespacesParams) HasManagementNamespace() bool`

HasManagementNamespace returns a boolean if a field has been set.

### SetManagementNamespaceNil

`func (o *RestoreKubernetesNamespacesParams) SetManagementNamespaceNil(b bool)`

 SetManagementNamespaceNil sets the value for ManagementNamespace to be an explicit nil

### UnsetManagementNamespace
`func (o *RestoreKubernetesNamespacesParams) UnsetManagementNamespace()`

UnsetManagementNamespace ensures that no value is present for ManagementNamespace, not even an explicit nil
### GetRenameRestoredObjectParam

`func (o *RestoreKubernetesNamespacesParams) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *RestoreKubernetesNamespacesParams) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *RestoreKubernetesNamespacesParams) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *RestoreKubernetesNamespacesParams) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


