# RecoverAzureParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureTargetParams** | Pointer to [**NullableRecoverAzureFileAndFolderParamsAzureTargetParams**](RecoverAzureFileAndFolderParamsAzureTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAzureParamsRecoverFileAndFolderParams

`func NewRecoverAzureParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverAzureParamsRecoverFileAndFolderParams`

NewRecoverAzureParamsRecoverFileAndFolderParams instantiates a new RecoverAzureParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverAzureParamsRecoverFileAndFolderParamsWithDefaults() *RecoverAzureParamsRecoverFileAndFolderParams`

NewRecoverAzureParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverAzureParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureTargetParams

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetAzureTargetParams() RecoverAzureFileAndFolderParamsAzureTargetParams`

GetAzureTargetParams returns the AzureTargetParams field if non-nil, zero value otherwise.

### GetAzureTargetParamsOk

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetAzureTargetParamsOk() (*RecoverAzureFileAndFolderParamsAzureTargetParams, bool)`

GetAzureTargetParamsOk returns a tuple with the AzureTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetParams

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) SetAzureTargetParams(v RecoverAzureFileAndFolderParamsAzureTargetParams)`

SetAzureTargetParams sets AzureTargetParams field to given value.

### HasAzureTargetParams

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) HasAzureTargetParams() bool`

HasAzureTargetParams returns a boolean if a field has been set.

### SetAzureTargetParamsNil

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) SetAzureTargetParamsNil(b bool)`

 SetAzureTargetParamsNil sets the value for AzureTargetParams to be an explicit nil

### UnsetAzureTargetParams
`func (o *RecoverAzureParamsRecoverFileAndFolderParams) UnsetAzureTargetParams()`

UnsetAzureTargetParams ensures that no value is present for AzureTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverAzureParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAzureParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


