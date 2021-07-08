# RetrieveArchiveTaskStateProtoDownloadFilesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathVec** | Pointer to **[]string** | The file paths to be retrieved from the archive. For example, if the file paths are /foo/bar and /foo2/bar and target_dir_path is \&quot;downloaded_files\&quot;. The retrieved files will have the full path structure maintained in target_dir_path such as downloaded_files/foo/bar and /downloaded_files/foo2/bar. | [optional] 
**MagnetoInstanceId** | Pointer to [**MagnetoInstanceId**](MagnetoInstanceId.md) |  | [optional] 
**ObjectId** | Pointer to [**MagnetoObjectId**](MagnetoObjectId.md) |  | [optional] 
**SkipRestoreInStubView** | Pointer to **NullableBool** | Ask Icebox to only create the stub view and skip restoring files in the stub view. | [optional] 
**TargetDirPath** | Pointer to **NullableString** | Path to the directory under which the downloaded files will be placed. | [optional] 
**TargetViewName** | Pointer to **NullableString** | Target view name where the downloaded files will be placed. | [optional] 

## Methods

### NewRetrieveArchiveTaskStateProtoDownloadFilesInfo

`func NewRetrieveArchiveTaskStateProtoDownloadFilesInfo() *RetrieveArchiveTaskStateProtoDownloadFilesInfo`

NewRetrieveArchiveTaskStateProtoDownloadFilesInfo instantiates a new RetrieveArchiveTaskStateProtoDownloadFilesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveArchiveTaskStateProtoDownloadFilesInfoWithDefaults

`func NewRetrieveArchiveTaskStateProtoDownloadFilesInfoWithDefaults() *RetrieveArchiveTaskStateProtoDownloadFilesInfo`

NewRetrieveArchiveTaskStateProtoDownloadFilesInfoWithDefaults instantiates a new RetrieveArchiveTaskStateProtoDownloadFilesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathVec

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetFilePathVec() []string`

GetFilePathVec returns the FilePathVec field if non-nil, zero value otherwise.

### GetFilePathVecOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetFilePathVecOk() (*[]string, bool)`

GetFilePathVecOk returns a tuple with the FilePathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathVec

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetFilePathVec(v []string)`

SetFilePathVec sets FilePathVec field to given value.

### HasFilePathVec

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasFilePathVec() bool`

HasFilePathVec returns a boolean if a field has been set.

### SetFilePathVecNil

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetFilePathVecNil(b bool)`

 SetFilePathVecNil sets the value for FilePathVec to be an explicit nil

### UnsetFilePathVec
`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) UnsetFilePathVec()`

UnsetFilePathVec ensures that no value is present for FilePathVec, not even an explicit nil
### GetMagnetoInstanceId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetMagnetoInstanceId() MagnetoInstanceId`

GetMagnetoInstanceId returns the MagnetoInstanceId field if non-nil, zero value otherwise.

### GetMagnetoInstanceIdOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetMagnetoInstanceIdOk() (*MagnetoInstanceId, bool)`

GetMagnetoInstanceIdOk returns a tuple with the MagnetoInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetoInstanceId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetMagnetoInstanceId(v MagnetoInstanceId)`

SetMagnetoInstanceId sets MagnetoInstanceId field to given value.

### HasMagnetoInstanceId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasMagnetoInstanceId() bool`

HasMagnetoInstanceId returns a boolean if a field has been set.

### GetObjectId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetObjectId() MagnetoObjectId`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetObjectIdOk() (*MagnetoObjectId, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetObjectId(v MagnetoObjectId)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSkipRestoreInStubView

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetSkipRestoreInStubView() bool`

GetSkipRestoreInStubView returns the SkipRestoreInStubView field if non-nil, zero value otherwise.

### GetSkipRestoreInStubViewOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetSkipRestoreInStubViewOk() (*bool, bool)`

GetSkipRestoreInStubViewOk returns a tuple with the SkipRestoreInStubView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRestoreInStubView

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetSkipRestoreInStubView(v bool)`

SetSkipRestoreInStubView sets SkipRestoreInStubView field to given value.

### HasSkipRestoreInStubView

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasSkipRestoreInStubView() bool`

HasSkipRestoreInStubView returns a boolean if a field has been set.

### SetSkipRestoreInStubViewNil

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetSkipRestoreInStubViewNil(b bool)`

 SetSkipRestoreInStubViewNil sets the value for SkipRestoreInStubView to be an explicit nil

### UnsetSkipRestoreInStubView
`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) UnsetSkipRestoreInStubView()`

UnsetSkipRestoreInStubView ensures that no value is present for SkipRestoreInStubView, not even an explicit nil
### GetTargetDirPath

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetTargetDirPath() string`

GetTargetDirPath returns the TargetDirPath field if non-nil, zero value otherwise.

### GetTargetDirPathOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetTargetDirPathOk() (*string, bool)`

GetTargetDirPathOk returns a tuple with the TargetDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirPath

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetTargetDirPath(v string)`

SetTargetDirPath sets TargetDirPath field to given value.

### HasTargetDirPath

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasTargetDirPath() bool`

HasTargetDirPath returns a boolean if a field has been set.

### SetTargetDirPathNil

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetTargetDirPathNil(b bool)`

 SetTargetDirPathNil sets the value for TargetDirPath to be an explicit nil

### UnsetTargetDirPath
`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) UnsetTargetDirPath()`

UnsetTargetDirPath ensures that no value is present for TargetDirPath, not even an explicit nil
### GetTargetViewName

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.

### SetTargetViewNameNil

`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) SetTargetViewNameNil(b bool)`

 SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil

### UnsetTargetViewName
`func (o *RetrieveArchiveTaskStateProtoDownloadFilesInfo) UnsetTargetViewName()`

UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


