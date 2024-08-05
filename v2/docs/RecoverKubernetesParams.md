# RecoverKubernetesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverGcpParamsDownloadFileAndFolderParams**](RecoverGcpParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of objects which need to be recovered. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverKubernetesParamsRecoverFileAndFolderParams**](RecoverKubernetesParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverNamespaceParams** | Pointer to [**NullableRecoverKubernetesParamsRecoverNamespaceParams**](RecoverKubernetesParamsRecoverNamespaceParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverKubernetesParams

`func NewRecoverKubernetesParams(recoveryAction string, ) *RecoverKubernetesParams`

NewRecoverKubernetesParams instantiates a new RecoverKubernetesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKubernetesParamsWithDefaults

`func NewRecoverKubernetesParamsWithDefaults() *RecoverKubernetesParams`

NewRecoverKubernetesParamsWithDefaults instantiates a new RecoverKubernetesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverKubernetesParams) GetDownloadFileAndFolderParams() RecoverGcpParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverKubernetesParams) GetDownloadFileAndFolderParamsOk() (*RecoverGcpParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverKubernetesParams) SetDownloadFileAndFolderParams(v RecoverGcpParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverKubernetesParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverKubernetesParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverKubernetesParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverKubernetesParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverKubernetesParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverKubernetesParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverKubernetesParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverKubernetesParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverKubernetesParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverKubernetesParams) GetRecoverFileAndFolderParams() RecoverKubernetesParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverKubernetesParams) GetRecoverFileAndFolderParamsOk() (*RecoverKubernetesParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverKubernetesParams) SetRecoverFileAndFolderParams(v RecoverKubernetesParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverKubernetesParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverKubernetesParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverKubernetesParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverNamespaceParams

`func (o *RecoverKubernetesParams) GetRecoverNamespaceParams() RecoverKubernetesParamsRecoverNamespaceParams`

GetRecoverNamespaceParams returns the RecoverNamespaceParams field if non-nil, zero value otherwise.

### GetRecoverNamespaceParamsOk

`func (o *RecoverKubernetesParams) GetRecoverNamespaceParamsOk() (*RecoverKubernetesParamsRecoverNamespaceParams, bool)`

GetRecoverNamespaceParamsOk returns a tuple with the RecoverNamespaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverNamespaceParams

`func (o *RecoverKubernetesParams) SetRecoverNamespaceParams(v RecoverKubernetesParamsRecoverNamespaceParams)`

SetRecoverNamespaceParams sets RecoverNamespaceParams field to given value.

### HasRecoverNamespaceParams

`func (o *RecoverKubernetesParams) HasRecoverNamespaceParams() bool`

HasRecoverNamespaceParams returns a boolean if a field has been set.

### SetRecoverNamespaceParamsNil

`func (o *RecoverKubernetesParams) SetRecoverNamespaceParamsNil(b bool)`

 SetRecoverNamespaceParamsNil sets the value for RecoverNamespaceParams to be an explicit nil

### UnsetRecoverNamespaceParams
`func (o *RecoverKubernetesParams) UnsetRecoverNamespaceParams()`

UnsetRecoverNamespaceParams ensures that no value is present for RecoverNamespaceParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverKubernetesParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverKubernetesParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverKubernetesParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


