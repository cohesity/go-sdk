# RecoverKvmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverGcpParamsDownloadFileAndFolderParams**](RecoverGcpParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. | [optional] 
**RecoverVmParams** | Pointer to [**RecoverKvmParamsRecoverVmParams**](RecoverKvmParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 

## Methods

### NewRecoverKvmParams

`func NewRecoverKvmParams(recoveryAction string, ) *RecoverKvmParams`

NewRecoverKvmParams instantiates a new RecoverKvmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmParamsWithDefaults

`func NewRecoverKvmParamsWithDefaults() *RecoverKvmParams`

NewRecoverKvmParamsWithDefaults instantiates a new RecoverKvmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverKvmParams) GetDownloadFileAndFolderParams() RecoverGcpParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverKvmParams) GetDownloadFileAndFolderParamsOk() (*RecoverGcpParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverKvmParams) SetDownloadFileAndFolderParams(v RecoverGcpParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverKvmParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverKvmParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverKvmParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverKvmParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverKvmParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverKvmParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverKvmParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverKvmParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverKvmParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverKvmParams) GetRecoverVmParams() RecoverKvmParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverKvmParams) GetRecoverVmParamsOk() (*RecoverKvmParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverKvmParams) SetRecoverVmParams(v RecoverKvmParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverKvmParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### GetRecoveryAction

`func (o *RecoverKvmParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverKvmParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverKvmParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


