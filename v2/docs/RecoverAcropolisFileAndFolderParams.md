# RecoverAcropolisFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisTargetParams** | Pointer to [**NullableRecoverAcropolisFileAndFolderParamsAcropolisTargetParams**](RecoverAcropolisFileAndFolderParamsAcropolisTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAcropolisFileAndFolderParams

`func NewRecoverAcropolisFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverAcropolisFileAndFolderParams`

NewRecoverAcropolisFileAndFolderParams instantiates a new RecoverAcropolisFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisFileAndFolderParamsWithDefaults

`func NewRecoverAcropolisFileAndFolderParamsWithDefaults() *RecoverAcropolisFileAndFolderParams`

NewRecoverAcropolisFileAndFolderParamsWithDefaults instantiates a new RecoverAcropolisFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisTargetParams

`func (o *RecoverAcropolisFileAndFolderParams) GetAcropolisTargetParams() RecoverAcropolisFileAndFolderParamsAcropolisTargetParams`

GetAcropolisTargetParams returns the AcropolisTargetParams field if non-nil, zero value otherwise.

### GetAcropolisTargetParamsOk

`func (o *RecoverAcropolisFileAndFolderParams) GetAcropolisTargetParamsOk() (*RecoverAcropolisFileAndFolderParamsAcropolisTargetParams, bool)`

GetAcropolisTargetParamsOk returns a tuple with the AcropolisTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisTargetParams

`func (o *RecoverAcropolisFileAndFolderParams) SetAcropolisTargetParams(v RecoverAcropolisFileAndFolderParamsAcropolisTargetParams)`

SetAcropolisTargetParams sets AcropolisTargetParams field to given value.

### HasAcropolisTargetParams

`func (o *RecoverAcropolisFileAndFolderParams) HasAcropolisTargetParams() bool`

HasAcropolisTargetParams returns a boolean if a field has been set.

### SetAcropolisTargetParamsNil

`func (o *RecoverAcropolisFileAndFolderParams) SetAcropolisTargetParamsNil(b bool)`

 SetAcropolisTargetParamsNil sets the value for AcropolisTargetParams to be an explicit nil

### UnsetAcropolisTargetParams
`func (o *RecoverAcropolisFileAndFolderParams) UnsetAcropolisTargetParams()`

UnsetAcropolisTargetParams ensures that no value is present for AcropolisTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverAcropolisFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverAcropolisFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverAcropolisFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverAcropolisFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverAcropolisFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAcropolisFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAcropolisFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAcropolisFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


