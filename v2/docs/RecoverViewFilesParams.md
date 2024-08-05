# RecoverViewFilesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAndFolders** | [**[]ViewRecoverFileAndFolderInfo**](ViewRecoverFileAndFolderInfo.md) | Specifies the list of info about the view files and folders to be recovered. | 
**ViewTargetParams** | Pointer to [**NullableRecoverViewFilesParamsViewTargetParams**](RecoverViewFilesParamsViewTargetParams.md) |  | [optional] 

## Methods

### NewRecoverViewFilesParams

`func NewRecoverViewFilesParams(filesAndFolders []ViewRecoverFileAndFolderInfo, ) *RecoverViewFilesParams`

NewRecoverViewFilesParams instantiates a new RecoverViewFilesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverViewFilesParamsWithDefaults

`func NewRecoverViewFilesParamsWithDefaults() *RecoverViewFilesParams`

NewRecoverViewFilesParamsWithDefaults instantiates a new RecoverViewFilesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAndFolders

`func (o *RecoverViewFilesParams) GetFilesAndFolders() []ViewRecoverFileAndFolderInfo`

GetFilesAndFolders returns the FilesAndFolders field if non-nil, zero value otherwise.

### GetFilesAndFoldersOk

`func (o *RecoverViewFilesParams) GetFilesAndFoldersOk() (*[]ViewRecoverFileAndFolderInfo, bool)`

GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAndFolders

`func (o *RecoverViewFilesParams) SetFilesAndFolders(v []ViewRecoverFileAndFolderInfo)`

SetFilesAndFolders sets FilesAndFolders field to given value.


### SetFilesAndFoldersNil

`func (o *RecoverViewFilesParams) SetFilesAndFoldersNil(b bool)`

 SetFilesAndFoldersNil sets the value for FilesAndFolders to be an explicit nil

### UnsetFilesAndFolders
`func (o *RecoverViewFilesParams) UnsetFilesAndFolders()`

UnsetFilesAndFolders ensures that no value is present for FilesAndFolders, not even an explicit nil
### GetViewTargetParams

`func (o *RecoverViewFilesParams) GetViewTargetParams() RecoverViewFilesParamsViewTargetParams`

GetViewTargetParams returns the ViewTargetParams field if non-nil, zero value otherwise.

### GetViewTargetParamsOk

`func (o *RecoverViewFilesParams) GetViewTargetParamsOk() (*RecoverViewFilesParamsViewTargetParams, bool)`

GetViewTargetParamsOk returns a tuple with the ViewTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewTargetParams

`func (o *RecoverViewFilesParams) SetViewTargetParams(v RecoverViewFilesParamsViewTargetParams)`

SetViewTargetParams sets ViewTargetParams field to given value.

### HasViewTargetParams

`func (o *RecoverViewFilesParams) HasViewTargetParams() bool`

HasViewTargetParams returns a boolean if a field has been set.

### SetViewTargetParamsNil

`func (o *RecoverViewFilesParams) SetViewTargetParamsNil(b bool)`

 SetViewTargetParamsNil sets the value for ViewTargetParams to be an explicit nil

### UnsetViewTargetParams
`func (o *RecoverViewFilesParams) UnsetViewTargetParams()`

UnsetViewTargetParams ensures that no value is present for ViewTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


