# RecoverAwsParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTargetParams** | Pointer to [**NullableRecoverAwsFileAndFolderParamsAwsTargetParams**](RecoverAwsFileAndFolderParamsAwsTargetParams.md) |  | [optional] 
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAwsParamsRecoverFileAndFolderParams

`func NewRecoverAwsParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverAwsParamsRecoverFileAndFolderParams`

NewRecoverAwsParamsRecoverFileAndFolderParams instantiates a new RecoverAwsParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverAwsParamsRecoverFileAndFolderParamsWithDefaults() *RecoverAwsParamsRecoverFileAndFolderParams`

NewRecoverAwsParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverAwsParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTargetParams

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetAwsTargetParams() RecoverAwsFileAndFolderParamsAwsTargetParams`

GetAwsTargetParams returns the AwsTargetParams field if non-nil, zero value otherwise.

### GetAwsTargetParamsOk

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetAwsTargetParamsOk() (*RecoverAwsFileAndFolderParamsAwsTargetParams, bool)`

GetAwsTargetParamsOk returns a tuple with the AwsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetParams

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetAwsTargetParams(v RecoverAwsFileAndFolderParamsAwsTargetParams)`

SetAwsTargetParams sets AwsTargetParams field to given value.

### HasAwsTargetParams

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) HasAwsTargetParams() bool`

HasAwsTargetParams returns a boolean if a field has been set.

### SetAwsTargetParamsNil

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetAwsTargetParamsNil(b bool)`

 SetAwsTargetParamsNil sets the value for AwsTargetParams to be an explicit nil

### UnsetAwsTargetParams
`func (o *RecoverAwsParamsRecoverFileAndFolderParams) UnsetAwsTargetParams()`

UnsetAwsTargetParams ensures that no value is present for AwsTargetParams, not even an explicit nil
### GetFilesAndFolders

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverAwsParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAwsParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


