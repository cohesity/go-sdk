# RecoverGcpFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**GcpTargetParams** | Pointer to [**NullableGcpTargetParamsForRecoverFileAndFolder**](GcpTargetParamsForRecoverFileAndFolder.md) | Specifies the parameters to recover to a GCP target. | [optional] 

## Methods

### NewRecoverGcpFileAndFolderParams

`func NewRecoverGcpFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverGcpFileAndFolderParams`

NewRecoverGcpFileAndFolderParams instantiates a new RecoverGcpFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpFileAndFolderParamsWithDefaults

`func NewRecoverGcpFileAndFolderParamsWithDefaults() *RecoverGcpFileAndFolderParams`

NewRecoverGcpFileAndFolderParamsWithDefaults instantiates a new RecoverGcpFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverGcpFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverGcpFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverGcpFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverGcpFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverGcpFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGcpFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGcpFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGcpFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetGcpTargetParams

`func (o *RecoverGcpFileAndFolderParams) GetGcpTargetParams() GcpTargetParamsForRecoverFileAndFolder`

GetGcpTargetParams returns the GcpTargetParams field if non-nil, zero value otherwise.

### GetGcpTargetParamsOk

`func (o *RecoverGcpFileAndFolderParams) GetGcpTargetParamsOk() (*GcpTargetParamsForRecoverFileAndFolder, bool)`

GetGcpTargetParamsOk returns a tuple with the GcpTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTargetParams

`func (o *RecoverGcpFileAndFolderParams) SetGcpTargetParams(v GcpTargetParamsForRecoverFileAndFolder)`

SetGcpTargetParams sets GcpTargetParams field to given value.

### HasGcpTargetParams

`func (o *RecoverGcpFileAndFolderParams) HasGcpTargetParams() bool`

HasGcpTargetParams returns a boolean if a field has been set.

### SetGcpTargetParamsNil

`func (o *RecoverGcpFileAndFolderParams) SetGcpTargetParamsNil(b bool)`

 SetGcpTargetParamsNil sets the value for GcpTargetParams to be an explicit nil

### UnsetGcpTargetParams
`func (o *RecoverGcpFileAndFolderParams) UnsetGcpTargetParams()`

UnsetGcpTargetParams ensures that no value is present for GcpTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


