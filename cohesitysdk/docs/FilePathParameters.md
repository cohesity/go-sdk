# FilePathParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupFilePath** | Pointer to **NullableString** | Specifies absolute path to a file or a directory in a Physical Server to protect. | [optional] 
**ExcludedFilePaths** | Pointer to **[]string** | Array of Excluded File Paths.  Specifies absolute paths to descendant files or subdirectories of backupFilePath that should not be protected. | [optional] 
**SkipNestedVolumes** | Pointer to **NullableBool** | Specifies if any subdirectories under backupFilePath, where local or network volumes are mounted, should be excluded from being protected. If true, the volumes are not protected. deprecated: true | [optional] 

## Methods

### NewFilePathParameters

`func NewFilePathParameters() *FilePathParameters`

NewFilePathParameters instantiates a new FilePathParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePathParametersWithDefaults

`func NewFilePathParametersWithDefaults() *FilePathParameters`

NewFilePathParametersWithDefaults instantiates a new FilePathParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupFilePath

`func (o *FilePathParameters) GetBackupFilePath() string`

GetBackupFilePath returns the BackupFilePath field if non-nil, zero value otherwise.

### GetBackupFilePathOk

`func (o *FilePathParameters) GetBackupFilePathOk() (*string, bool)`

GetBackupFilePathOk returns a tuple with the BackupFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFilePath

`func (o *FilePathParameters) SetBackupFilePath(v string)`

SetBackupFilePath sets BackupFilePath field to given value.

### HasBackupFilePath

`func (o *FilePathParameters) HasBackupFilePath() bool`

HasBackupFilePath returns a boolean if a field has been set.

### SetBackupFilePathNil

`func (o *FilePathParameters) SetBackupFilePathNil(b bool)`

 SetBackupFilePathNil sets the value for BackupFilePath to be an explicit nil

### UnsetBackupFilePath
`func (o *FilePathParameters) UnsetBackupFilePath()`

UnsetBackupFilePath ensures that no value is present for BackupFilePath, not even an explicit nil
### GetExcludedFilePaths

`func (o *FilePathParameters) GetExcludedFilePaths() []string`

GetExcludedFilePaths returns the ExcludedFilePaths field if non-nil, zero value otherwise.

### GetExcludedFilePathsOk

`func (o *FilePathParameters) GetExcludedFilePathsOk() (*[]string, bool)`

GetExcludedFilePathsOk returns a tuple with the ExcludedFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFilePaths

`func (o *FilePathParameters) SetExcludedFilePaths(v []string)`

SetExcludedFilePaths sets ExcludedFilePaths field to given value.

### HasExcludedFilePaths

`func (o *FilePathParameters) HasExcludedFilePaths() bool`

HasExcludedFilePaths returns a boolean if a field has been set.

### SetExcludedFilePathsNil

`func (o *FilePathParameters) SetExcludedFilePathsNil(b bool)`

 SetExcludedFilePathsNil sets the value for ExcludedFilePaths to be an explicit nil

### UnsetExcludedFilePaths
`func (o *FilePathParameters) UnsetExcludedFilePaths()`

UnsetExcludedFilePaths ensures that no value is present for ExcludedFilePaths, not even an explicit nil
### GetSkipNestedVolumes

`func (o *FilePathParameters) GetSkipNestedVolumes() bool`

GetSkipNestedVolumes returns the SkipNestedVolumes field if non-nil, zero value otherwise.

### GetSkipNestedVolumesOk

`func (o *FilePathParameters) GetSkipNestedVolumesOk() (*bool, bool)`

GetSkipNestedVolumesOk returns a tuple with the SkipNestedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNestedVolumes

`func (o *FilePathParameters) SetSkipNestedVolumes(v bool)`

SetSkipNestedVolumes sets SkipNestedVolumes field to given value.

### HasSkipNestedVolumes

`func (o *FilePathParameters) HasSkipNestedVolumes() bool`

HasSkipNestedVolumes returns a boolean if a field has been set.

### SetSkipNestedVolumesNil

`func (o *FilePathParameters) SetSkipNestedVolumesNil(b bool)`

 SetSkipNestedVolumesNil sets the value for SkipNestedVolumes to be an explicit nil

### UnsetSkipNestedVolumes
`func (o *FilePathParameters) UnsetSkipNestedVolumes()`

UnsetSkipNestedVolumes ensures that no value is present for SkipNestedVolumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


