# RecoverKubernetesFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the information about the files and folders to be recovered. | 
**KubernetesTargetParams** | Pointer to [**NullableRecoverKubernetesFileAndFolderParamsKubernetesTargetParams**](RecoverKubernetesFileAndFolderParamsKubernetesTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverKubernetesFileAndFolderParams

`func NewRecoverKubernetesFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverKubernetesFileAndFolderParams`

NewRecoverKubernetesFileAndFolderParams instantiates a new RecoverKubernetesFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKubernetesFileAndFolderParamsWithDefaults

`func NewRecoverKubernetesFileAndFolderParamsWithDefaults() *RecoverKubernetesFileAndFolderParams`

NewRecoverKubernetesFileAndFolderParamsWithDefaults instantiates a new RecoverKubernetesFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverKubernetesFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverKubernetesFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverKubernetesFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverKubernetesFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverKubernetesFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetKubernetesTargetParams

`func (o *RecoverKubernetesFileAndFolderParams) GetKubernetesTargetParams() RecoverKubernetesFileAndFolderParamsKubernetesTargetParams`

GetKubernetesTargetParams returns the KubernetesTargetParams field if non-nil, zero value otherwise.

### GetKubernetesTargetParamsOk

`func (o *RecoverKubernetesFileAndFolderParams) GetKubernetesTargetParamsOk() (*RecoverKubernetesFileAndFolderParamsKubernetesTargetParams, bool)`

GetKubernetesTargetParamsOk returns a tuple with the KubernetesTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetParams

`func (o *RecoverKubernetesFileAndFolderParams) SetKubernetesTargetParams(v RecoverKubernetesFileAndFolderParamsKubernetesTargetParams)`

SetKubernetesTargetParams sets KubernetesTargetParams field to given value.

### HasKubernetesTargetParams

`func (o *RecoverKubernetesFileAndFolderParams) HasKubernetesTargetParams() bool`

HasKubernetesTargetParams returns a boolean if a field has been set.

### SetKubernetesTargetParamsNil

`func (o *RecoverKubernetesFileAndFolderParams) SetKubernetesTargetParamsNil(b bool)`

 SetKubernetesTargetParamsNil sets the value for KubernetesTargetParams to be an explicit nil

### UnsetKubernetesTargetParams
`func (o *RecoverKubernetesFileAndFolderParams) UnsetKubernetesTargetParams()`

UnsetKubernetesTargetParams ensures that no value is present for KubernetesTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverKubernetesFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverKubernetesFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverKubernetesFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


