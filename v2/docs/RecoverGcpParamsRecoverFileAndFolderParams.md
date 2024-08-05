# RecoverGcpParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**GcpTargetParams** | Pointer to [**NullableRecoverGcpFileAndFolderParamsGcpTargetParams**](RecoverGcpFileAndFolderParamsGcpTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverGcpParamsRecoverFileAndFolderParams

`func NewRecoverGcpParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverGcpParamsRecoverFileAndFolderParams`

NewRecoverGcpParamsRecoverFileAndFolderParams instantiates a new RecoverGcpParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverGcpParamsRecoverFileAndFolderParamsWithDefaults() *RecoverGcpParamsRecoverFileAndFolderParams`

NewRecoverGcpParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverGcpParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverGcpParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetGcpTargetParams

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetGcpTargetParams() RecoverGcpFileAndFolderParamsGcpTargetParams`

GetGcpTargetParams returns the GcpTargetParams field if non-nil, zero value otherwise.

### GetGcpTargetParamsOk

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetGcpTargetParamsOk() (*RecoverGcpFileAndFolderParamsGcpTargetParams, bool)`

GetGcpTargetParamsOk returns a tuple with the GcpTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTargetParams

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) SetGcpTargetParams(v RecoverGcpFileAndFolderParamsGcpTargetParams)`

SetGcpTargetParams sets GcpTargetParams field to given value.

### HasGcpTargetParams

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) HasGcpTargetParams() bool`

HasGcpTargetParams returns a boolean if a field has been set.

### SetGcpTargetParamsNil

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) SetGcpTargetParamsNil(b bool)`

 SetGcpTargetParamsNil sets the value for GcpTargetParams to be an explicit nil

### UnsetGcpTargetParams
`func (o *RecoverGcpParamsRecoverFileAndFolderParams) UnsetGcpTargetParams()`

UnsetGcpTargetParams ensures that no value is present for GcpTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGcpParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


