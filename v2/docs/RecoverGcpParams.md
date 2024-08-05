# RecoverGcpParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverGcpParamsDownloadFileAndFolderParams**](RecoverGcpParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverGcpParamsRecoverFileAndFolderParams**](RecoverGcpParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverVmParams** | Pointer to [**NullableRecoverGcpParamsRecoverVmParams**](RecoverGcpParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverGcpParams

`func NewRecoverGcpParams(recoveryAction string, ) *RecoverGcpParams`

NewRecoverGcpParams instantiates a new RecoverGcpParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpParamsWithDefaults

`func NewRecoverGcpParamsWithDefaults() *RecoverGcpParams`

NewRecoverGcpParamsWithDefaults instantiates a new RecoverGcpParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverGcpParams) GetDownloadFileAndFolderParams() RecoverGcpParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverGcpParams) GetDownloadFileAndFolderParamsOk() (*RecoverGcpParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverGcpParams) SetDownloadFileAndFolderParams(v RecoverGcpParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverGcpParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverGcpParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverGcpParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverGcpParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverGcpParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverGcpParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverGcpParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverGcpParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverGcpParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverGcpParams) GetRecoverFileAndFolderParams() RecoverGcpParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverGcpParams) GetRecoverFileAndFolderParamsOk() (*RecoverGcpParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverGcpParams) SetRecoverFileAndFolderParams(v RecoverGcpParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverGcpParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverGcpParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverGcpParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverGcpParams) GetRecoverVmParams() RecoverGcpParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverGcpParams) GetRecoverVmParamsOk() (*RecoverGcpParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverGcpParams) SetRecoverVmParams(v RecoverGcpParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverGcpParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverGcpParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverGcpParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverGcpParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverGcpParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverGcpParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


