# RecoverHyperVParamsRecoverFileAndFolderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]CommonRecoverFileAndFolderInfo**](CommonRecoverFileAndFolderInfo.md) | Specifies the info about the files and folders to be recovered. | 
**HypervTargetParams** | Pointer to [**NullableRecoverHyperVFileAndFolderParamsHypervTargetParams**](RecoverHyperVFileAndFolderParamsHypervTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverHyperVParamsRecoverFileAndFolderParams

`func NewRecoverHyperVParamsRecoverFileAndFolderParams(filesAndFolders []CommonRecoverFileAndFolderInfo, targetEnvironment string, ) *RecoverHyperVParamsRecoverFileAndFolderParams`

NewRecoverHyperVParamsRecoverFileAndFolderParams instantiates a new RecoverHyperVParamsRecoverFileAndFolderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVParamsRecoverFileAndFolderParamsWithDefaults

`func NewRecoverHyperVParamsRecoverFileAndFolderParamsWithDefaults() *RecoverHyperVParamsRecoverFileAndFolderParams`

NewRecoverHyperVParamsRecoverFileAndFolderParamsWithDefaults instantiates a new RecoverHyperVParamsRecoverFileAndFolderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetFilesAndFoldersOk() (*[]CommonRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetHypervTargetParams

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetHypervTargetParams() RecoverHyperVFileAndFolderParamsHypervTargetParams`

GetHypervTargetParams returns the HypervTargetParams field if non-nil, zero value otherwise.

### GetHypervTargetParamsOk

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetHypervTargetParamsOk() (*RecoverHyperVFileAndFolderParamsHypervTargetParams, bool)`

GetHypervTargetParamsOk returns a tuple with the HypervTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervTargetParams

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) SetHypervTargetParams(v RecoverHyperVFileAndFolderParamsHypervTargetParams)`

SetHypervTargetParams sets HypervTargetParams field to given value.

### HasHypervTargetParams

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) HasHypervTargetParams() bool`

HasHypervTargetParams returns a boolean if a field has been set.

### SetHypervTargetParamsNil

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) SetHypervTargetParamsNil(b bool)`

 SetHypervTargetParamsNil sets the value for HypervTargetParams to be an explicit nil

### UnsetHypervTargetParams
`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) UnsetHypervTargetParams()`

UnsetHypervTargetParams ensures that no value is present for HypervTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverHyperVParamsRecoverFileAndFolderParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


