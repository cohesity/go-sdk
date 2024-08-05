# RecoverAcropolisParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisTargetParams** | Pointer to [**NullableRecoverAcropolisFileAndFolderParamsAcropolisTargetParams**](RecoverAcropolisFileAndFolderParamsAcropolisTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAcropolisParamsRecoverFileAndFolderParams

`func NewRecoverAcropolisParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverAcropolisParamsRecoverFileAndFolderParams`

NewRecoverAcropolisParamsRecoverFileAndFolderParams instantiates a new RecoverAcropolisParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverAcropolisParamsRecoverFileAndFolderParamsWithDefaults() *RecoverAcropolisParamsRecoverFileAndFolderParams`

NewRecoverAcropolisParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverAcropolisParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisTargetParams

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetAcropolisTargetParams() RecoverAcropolisFileAndFolderParamsAcropolisTargetParams`

GetAcropolisTargetParams returns the AcropolisTargetParams field if non-nil, zero value otherwise.

### GetAcropolisTargetParamsOk

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetAcropolisTargetParamsOk() (*RecoverAcropolisFileAndFolderParamsAcropolisTargetParams, bool)`

GetAcropolisTargetParamsOk returns a tuple with the AcropolisTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisTargetParams

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) SetAcropolisTargetParams(v RecoverAcropolisFileAndFolderParamsAcropolisTargetParams)`

SetAcropolisTargetParams sets AcropolisTargetParams field to given value.

### HasAcropolisTargetParams

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) HasAcropolisTargetParams() bool`

HasAcropolisTargetParams returns a boolean if a field has been set.

### SetAcropolisTargetParamsNil

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) SetAcropolisTargetParamsNil(b bool)`

 SetAcropolisTargetParamsNil sets the value for AcropolisTargetParams to be an explicit nil

### UnsetAcropolisTargetParams
`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) UnsetAcropolisTargetParams()`

UnsetAcropolisTargetParams ensures that no value is present for AcropolisTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAcropolisParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


