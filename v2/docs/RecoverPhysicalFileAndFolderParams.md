# RecoverPhysicalFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the information about the files and folders to be recovered. | 
**PhysicalTargetParams** | Pointer to [**NullableRecoverPhysicalFileAndFolderParamsPhysicalTargetParams**](RecoverPhysicalFileAndFolderParamsPhysicalTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverPhysicalFileAndFolderParams

`func NewRecoverPhysicalFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverPhysicalFileAndFolderParams`

NewRecoverPhysicalFileAndFolderParams instantiates a new RecoverPhysicalFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalFileAndFolderParamsWithDefaults

`func NewRecoverPhysicalFileAndFolderParamsWithDefaults() *RecoverPhysicalFileAndFolderParams`

NewRecoverPhysicalFileAndFolderParamsWithDefaults instantiates a new RecoverPhysicalFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverPhysicalFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverPhysicalFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverPhysicalFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverPhysicalFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverPhysicalFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetPhysicalTargetParams

`func (o *RecoverPhysicalFileAndFolderParams) GetPhysicalTargetParams() RecoverPhysicalFileAndFolderParamsPhysicalTargetParams`

GetPhysicalTargetParams returns the PhysicalTargetParams field if non-nil, zero value otherwise.

### GetPhysicalTargetParamsOk

`func (o *RecoverPhysicalFileAndFolderParams) GetPhysicalTargetParamsOk() (*RecoverPhysicalFileAndFolderParamsPhysicalTargetParams, bool)`

GetPhysicalTargetParamsOk returns a tuple with the PhysicalTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalTargetParams

`func (o *RecoverPhysicalFileAndFolderParams) SetPhysicalTargetParams(v RecoverPhysicalFileAndFolderParamsPhysicalTargetParams)`

SetPhysicalTargetParams sets PhysicalTargetParams field to given value.

### HasPhysicalTargetParams

`func (o *RecoverPhysicalFileAndFolderParams) HasPhysicalTargetParams() bool`

HasPhysicalTargetParams returns a boolean if a field has been set.

### SetPhysicalTargetParamsNil

`func (o *RecoverPhysicalFileAndFolderParams) SetPhysicalTargetParamsNil(b bool)`

 SetPhysicalTargetParamsNil sets the value for PhysicalTargetParams to be an explicit nil

### UnsetPhysicalTargetParams
`func (o *RecoverPhysicalFileAndFolderParams) UnsetPhysicalTargetParams()`

UnsetPhysicalTargetParams ensures that no value is present for PhysicalTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverPhysicalFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverPhysicalFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverPhysicalFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


