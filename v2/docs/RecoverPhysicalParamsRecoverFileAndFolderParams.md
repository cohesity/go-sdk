# RecoverPhysicalParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the information about the files and folders to be recovered. | 
**PhysicalTargetParams** | Pointer to [**NullableRecoverPhysicalFileAndFolderParamsPhysicalTargetParams**](RecoverPhysicalFileAndFolderParamsPhysicalTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverPhysicalParamsRecoverFileAndFolderParams

`func NewRecoverPhysicalParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverPhysicalParamsRecoverFileAndFolderParams`

NewRecoverPhysicalParamsRecoverFileAndFolderParams instantiates a new RecoverPhysicalParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverPhysicalParamsRecoverFileAndFolderParamsWithDefaults() *RecoverPhysicalParamsRecoverFileAndFolderParams`

NewRecoverPhysicalParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverPhysicalParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetPhysicalTargetParams

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetPhysicalTargetParams() RecoverPhysicalFileAndFolderParamsPhysicalTargetParams`

GetPhysicalTargetParams returns the PhysicalTargetParams field if non-nil, zero value otherwise.

### GetPhysicalTargetParamsOk

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetPhysicalTargetParamsOk() (*RecoverPhysicalFileAndFolderParamsPhysicalTargetParams, bool)`

GetPhysicalTargetParamsOk returns a tuple with the PhysicalTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalTargetParams

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) SetPhysicalTargetParams(v RecoverPhysicalFileAndFolderParamsPhysicalTargetParams)`

SetPhysicalTargetParams sets PhysicalTargetParams field to given value.

### HasPhysicalTargetParams

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) HasPhysicalTargetParams() bool`

HasPhysicalTargetParams returns a boolean if a field has been set.

### SetPhysicalTargetParamsNil

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) SetPhysicalTargetParamsNil(b bool)`

 SetPhysicalTargetParamsNil sets the value for PhysicalTargetParams to be an explicit nil

### UnsetPhysicalTargetParams
`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) UnsetPhysicalTargetParams()`

UnsetPhysicalTargetParams ensures that no value is present for PhysicalTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverPhysicalParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


