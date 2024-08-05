# RecoverAwsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAwsParamsDownloadFileAndFolderParams**](RecoverAwsParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. For recovering files, specifies the object contains the file to recover. | [optional] 
**RecoverAuroraParams** | Pointer to [**NullableRecoverAwsParamsRecoverAuroraParams**](RecoverAwsParamsRecoverAuroraParams.md) |  | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverAwsParamsRecoverFileAndFolderParams**](RecoverAwsParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverRdsIngestParams** | Pointer to [**NullableRecoverAwsParamsRecoverRdsIngestParams**](RecoverAwsParamsRecoverRdsIngestParams.md) |  | [optional] 
**RecoverRdsParams** | Pointer to [**NullableRecoverAwsParamsRecoverRdsParams**](RecoverAwsParamsRecoverRdsParams.md) |  | [optional] 
**RecoverS3BucketParams** | Pointer to [**NullableRecoverAwsParamsRecoverS3BucketParams**](RecoverAwsParamsRecoverS3BucketParams.md) |  | [optional] 
**RecoverVmParams** | Pointer to [**NullableRecoverAwsParamsRecoverVmParams**](RecoverAwsParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverAwsParams

`func NewRecoverAwsParams(recoveryAction string, ) *RecoverAwsParams`

NewRecoverAwsParams instantiates a new RecoverAwsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsParamsWithDefaults

`func NewRecoverAwsParamsWithDefaults() *RecoverAwsParams`

NewRecoverAwsParamsWithDefaults instantiates a new RecoverAwsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverAwsParams) GetDownloadFileAndFolderParams() RecoverAwsParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverAwsParams) GetDownloadFileAndFolderParamsOk() (*RecoverAwsParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverAwsParams) SetDownloadFileAndFolderParams(v RecoverAwsParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverAwsParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverAwsParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverAwsParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverAwsParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverAwsParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverAwsParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverAwsParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverAwsParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverAwsParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverAuroraParams

`func (o *RecoverAwsParams) GetRecoverAuroraParams() RecoverAwsParamsRecoverAuroraParams`

GetRecoverAuroraParams returns the RecoverAuroraParams field if non-nil, zero value otherwise.

### GetRecoverAuroraParamsOk

`func (o *RecoverAwsParams) GetRecoverAuroraParamsOk() (*RecoverAwsParamsRecoverAuroraParams, bool)`

GetRecoverAuroraParamsOk returns a tuple with the RecoverAuroraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverAuroraParams

`func (o *RecoverAwsParams) SetRecoverAuroraParams(v RecoverAwsParamsRecoverAuroraParams)`

SetRecoverAuroraParams sets RecoverAuroraParams field to given value.

### HasRecoverAuroraParams

`func (o *RecoverAwsParams) HasRecoverAuroraParams() bool`

HasRecoverAuroraParams returns a boolean if a field has been set.

### SetRecoverAuroraParamsNil

`func (o *RecoverAwsParams) SetRecoverAuroraParamsNil(b bool)`

 SetRecoverAuroraParamsNil sets the value for RecoverAuroraParams to be an explicit nil

### UnsetRecoverAuroraParams
`func (o *RecoverAwsParams) UnsetRecoverAuroraParams()`

UnsetRecoverAuroraParams ensures that no value is present for RecoverAuroraParams, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverAwsParams) GetRecoverFileAndFolderParams() RecoverAwsParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverAwsParams) GetRecoverFileAndFolderParamsOk() (*RecoverAwsParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverAwsParams) SetRecoverFileAndFolderParams(v RecoverAwsParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverAwsParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverAwsParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverAwsParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverRdsIngestParams

`func (o *RecoverAwsParams) GetRecoverRdsIngestParams() RecoverAwsParamsRecoverRdsIngestParams`

GetRecoverRdsIngestParams returns the RecoverRdsIngestParams field if non-nil, zero value otherwise.

### GetRecoverRdsIngestParamsOk

`func (o *RecoverAwsParams) GetRecoverRdsIngestParamsOk() (*RecoverAwsParamsRecoverRdsIngestParams, bool)`

GetRecoverRdsIngestParamsOk returns a tuple with the RecoverRdsIngestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverRdsIngestParams

`func (o *RecoverAwsParams) SetRecoverRdsIngestParams(v RecoverAwsParamsRecoverRdsIngestParams)`

SetRecoverRdsIngestParams sets RecoverRdsIngestParams field to given value.

### HasRecoverRdsIngestParams

`func (o *RecoverAwsParams) HasRecoverRdsIngestParams() bool`

HasRecoverRdsIngestParams returns a boolean if a field has been set.

### SetRecoverRdsIngestParamsNil

`func (o *RecoverAwsParams) SetRecoverRdsIngestParamsNil(b bool)`

 SetRecoverRdsIngestParamsNil sets the value for RecoverRdsIngestParams to be an explicit nil

### UnsetRecoverRdsIngestParams
`func (o *RecoverAwsParams) UnsetRecoverRdsIngestParams()`

UnsetRecoverRdsIngestParams ensures that no value is present for RecoverRdsIngestParams, not even an explicit nil
### GetRecoverRdsParams

`func (o *RecoverAwsParams) GetRecoverRdsParams() RecoverAwsParamsRecoverRdsParams`

GetRecoverRdsParams returns the RecoverRdsParams field if non-nil, zero value otherwise.

### GetRecoverRdsParamsOk

`func (o *RecoverAwsParams) GetRecoverRdsParamsOk() (*RecoverAwsParamsRecoverRdsParams, bool)`

GetRecoverRdsParamsOk returns a tuple with the RecoverRdsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverRdsParams

`func (o *RecoverAwsParams) SetRecoverRdsParams(v RecoverAwsParamsRecoverRdsParams)`

SetRecoverRdsParams sets RecoverRdsParams field to given value.

### HasRecoverRdsParams

`func (o *RecoverAwsParams) HasRecoverRdsParams() bool`

HasRecoverRdsParams returns a boolean if a field has been set.

### SetRecoverRdsParamsNil

`func (o *RecoverAwsParams) SetRecoverRdsParamsNil(b bool)`

 SetRecoverRdsParamsNil sets the value for RecoverRdsParams to be an explicit nil

### UnsetRecoverRdsParams
`func (o *RecoverAwsParams) UnsetRecoverRdsParams()`

UnsetRecoverRdsParams ensures that no value is present for RecoverRdsParams, not even an explicit nil
### GetRecoverS3BucketParams

`func (o *RecoverAwsParams) GetRecoverS3BucketParams() RecoverAwsParamsRecoverS3BucketParams`

GetRecoverS3BucketParams returns the RecoverS3BucketParams field if non-nil, zero value otherwise.

### GetRecoverS3BucketParamsOk

`func (o *RecoverAwsParams) GetRecoverS3BucketParamsOk() (*RecoverAwsParamsRecoverS3BucketParams, bool)`

GetRecoverS3BucketParamsOk returns a tuple with the RecoverS3BucketParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverS3BucketParams

`func (o *RecoverAwsParams) SetRecoverS3BucketParams(v RecoverAwsParamsRecoverS3BucketParams)`

SetRecoverS3BucketParams sets RecoverS3BucketParams field to given value.

### HasRecoverS3BucketParams

`func (o *RecoverAwsParams) HasRecoverS3BucketParams() bool`

HasRecoverS3BucketParams returns a boolean if a field has been set.

### SetRecoverS3BucketParamsNil

`func (o *RecoverAwsParams) SetRecoverS3BucketParamsNil(b bool)`

 SetRecoverS3BucketParamsNil sets the value for RecoverS3BucketParams to be an explicit nil

### UnsetRecoverS3BucketParams
`func (o *RecoverAwsParams) UnsetRecoverS3BucketParams()`

UnsetRecoverS3BucketParams ensures that no value is present for RecoverS3BucketParams, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverAwsParams) GetRecoverVmParams() RecoverAwsParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverAwsParams) GetRecoverVmParamsOk() (*RecoverAwsParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverAwsParams) SetRecoverVmParams(v RecoverAwsParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverAwsParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverAwsParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverAwsParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverAwsParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverAwsParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverAwsParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


