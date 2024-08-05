# RecoverAzureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureEntraIdParams** | Pointer to [**NullableRecoverAzureParamsAzureEntraIdParams**](RecoverAzureParamsAzureEntraIdParams.md) |  | [optional] 
**AzureSqlParams** | Pointer to [**NullableRecoverAzureParamsAzureSqlParams**](RecoverAzureParamsAzureSqlParams.md) |  | [optional] 
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsDownloadFileAndFolderParams**](RecoverAcropolisParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. For recovering files, specifies the object contains the file to recover. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverAzureParamsRecoverFileAndFolderParams**](RecoverAzureParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverVmParams** | Pointer to [**NullableRecoverAzureParamsRecoverVmParams**](RecoverAzureParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverAzureParams

`func NewRecoverAzureParams(recoveryAction string, ) *RecoverAzureParams`

NewRecoverAzureParams instantiates a new RecoverAzureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureParamsWithDefaults

`func NewRecoverAzureParamsWithDefaults() *RecoverAzureParams`

NewRecoverAzureParamsWithDefaults instantiates a new RecoverAzureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureEntraIdParams

`func (o *RecoverAzureParams) GetAzureEntraIdParams() RecoverAzureParamsAzureEntraIdParams`

GetAzureEntraIdParams returns the AzureEntraIdParams field if non-nil, zero value otherwise.

### GetAzureEntraIdParamsOk

`func (o *RecoverAzureParams) GetAzureEntraIdParamsOk() (*RecoverAzureParamsAzureEntraIdParams, bool)`

GetAzureEntraIdParamsOk returns a tuple with the AzureEntraIdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEntraIdParams

`func (o *RecoverAzureParams) SetAzureEntraIdParams(v RecoverAzureParamsAzureEntraIdParams)`

SetAzureEntraIdParams sets AzureEntraIdParams field to given value.

### HasAzureEntraIdParams

`func (o *RecoverAzureParams) HasAzureEntraIdParams() bool`

HasAzureEntraIdParams returns a boolean if a field has been set.

### SetAzureEntraIdParamsNil

`func (o *RecoverAzureParams) SetAzureEntraIdParamsNil(b bool)`

 SetAzureEntraIdParamsNil sets the value for AzureEntraIdParams to be an explicit nil

### UnsetAzureEntraIdParams
`func (o *RecoverAzureParams) UnsetAzureEntraIdParams()`

UnsetAzureEntraIdParams ensures that no value is present for AzureEntraIdParams, not even an explicit nil
### GetAzureSqlParams

`func (o *RecoverAzureParams) GetAzureSqlParams() RecoverAzureParamsAzureSqlParams`

GetAzureSqlParams returns the AzureSqlParams field if non-nil, zero value otherwise.

### GetAzureSqlParamsOk

`func (o *RecoverAzureParams) GetAzureSqlParamsOk() (*RecoverAzureParamsAzureSqlParams, bool)`

GetAzureSqlParamsOk returns a tuple with the AzureSqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSqlParams

`func (o *RecoverAzureParams) SetAzureSqlParams(v RecoverAzureParamsAzureSqlParams)`

SetAzureSqlParams sets AzureSqlParams field to given value.

### HasAzureSqlParams

`func (o *RecoverAzureParams) HasAzureSqlParams() bool`

HasAzureSqlParams returns a boolean if a field has been set.

### SetAzureSqlParamsNil

`func (o *RecoverAzureParams) SetAzureSqlParamsNil(b bool)`

 SetAzureSqlParamsNil sets the value for AzureSqlParams to be an explicit nil

### UnsetAzureSqlParams
`func (o *RecoverAzureParams) UnsetAzureSqlParams()`

UnsetAzureSqlParams ensures that no value is present for AzureSqlParams, not even an explicit nil
### GetDownloadFileAndFolderParams

`func (o *RecoverAzureParams) GetDownloadFileAndFolderParams() RecoverAcropolisParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverAzureParams) GetDownloadFileAndFolderParamsOk() (*RecoverAcropolisParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverAzureParams) SetDownloadFileAndFolderParams(v RecoverAcropolisParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverAzureParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverAzureParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverAzureParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverAzureParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverAzureParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverAzureParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverAzureParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverAzureParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverAzureParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverAzureParams) GetRecoverFileAndFolderParams() RecoverAzureParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverAzureParams) GetRecoverFileAndFolderParamsOk() (*RecoverAzureParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverAzureParams) SetRecoverFileAndFolderParams(v RecoverAzureParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverAzureParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverAzureParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverAzureParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverAzureParams) GetRecoverVmParams() RecoverAzureParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverAzureParams) GetRecoverVmParamsOk() (*RecoverAzureParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverAzureParams) SetRecoverVmParams(v RecoverAzureParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverAzureParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverAzureParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverAzureParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverAzureParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverAzureParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverAzureParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


