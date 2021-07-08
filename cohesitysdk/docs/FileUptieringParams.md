# FileUptieringParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSelectPolicy** | Pointer to **NullableInt32** | File uptier policy based on file access/modify time. | [optional] 
**FileSize** | Pointer to **NullableInt64** | Gives the size criteria to be used for selecting the files to be uptiered. The hot files, which are greater or smaller than file_size, are uptiered. | [optional] 
**FileSizePolicy** | Pointer to **NullableInt32** | File size policy for selecting files to uptier. | [optional] 
**HotFileWindow** | Pointer to **NullableInt64** | Identifies the hot files in the view. Files which are accessed num_file_access times in hot_file_window msecs, are uptiered. It is only applicable when file_select_policy is kLastAccessed and num_file_access is greater than 1. | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Mount path where the Cohesity target view is mounted on NFS clients while migrating the data. | [optional] 
**NumFileAccess** | Pointer to **NullableInt32** | Number of times file must be accessed within hot_file_window in order to qualify for uptiering. Applicable only when file_select_policy is kLastAccessed. | [optional] 
**SourceViewName** | Pointer to **NullableString** | The source view name from which the data will be uptiered. | [optional] 
**UptierAllFiles** | Pointer to **NullableBool** | If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints. | [optional] 

## Methods

### NewFileUptieringParams

`func NewFileUptieringParams() *FileUptieringParams`

NewFileUptieringParams instantiates a new FileUptieringParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUptieringParamsWithDefaults

`func NewFileUptieringParamsWithDefaults() *FileUptieringParams`

NewFileUptieringParamsWithDefaults instantiates a new FileUptieringParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSelectPolicy

`func (o *FileUptieringParams) GetFileSelectPolicy() int32`

GetFileSelectPolicy returns the FileSelectPolicy field if non-nil, zero value otherwise.

### GetFileSelectPolicyOk

`func (o *FileUptieringParams) GetFileSelectPolicyOk() (*int32, bool)`

GetFileSelectPolicyOk returns a tuple with the FileSelectPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSelectPolicy

`func (o *FileUptieringParams) SetFileSelectPolicy(v int32)`

SetFileSelectPolicy sets FileSelectPolicy field to given value.

### HasFileSelectPolicy

`func (o *FileUptieringParams) HasFileSelectPolicy() bool`

HasFileSelectPolicy returns a boolean if a field has been set.

### SetFileSelectPolicyNil

`func (o *FileUptieringParams) SetFileSelectPolicyNil(b bool)`

 SetFileSelectPolicyNil sets the value for FileSelectPolicy to be an explicit nil

### UnsetFileSelectPolicy
`func (o *FileUptieringParams) UnsetFileSelectPolicy()`

UnsetFileSelectPolicy ensures that no value is present for FileSelectPolicy, not even an explicit nil
### GetFileSize

`func (o *FileUptieringParams) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileUptieringParams) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileUptieringParams) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileUptieringParams) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *FileUptieringParams) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *FileUptieringParams) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetFileSizePolicy

`func (o *FileUptieringParams) GetFileSizePolicy() int32`

GetFileSizePolicy returns the FileSizePolicy field if non-nil, zero value otherwise.

### GetFileSizePolicyOk

`func (o *FileUptieringParams) GetFileSizePolicyOk() (*int32, bool)`

GetFileSizePolicyOk returns a tuple with the FileSizePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizePolicy

`func (o *FileUptieringParams) SetFileSizePolicy(v int32)`

SetFileSizePolicy sets FileSizePolicy field to given value.

### HasFileSizePolicy

`func (o *FileUptieringParams) HasFileSizePolicy() bool`

HasFileSizePolicy returns a boolean if a field has been set.

### SetFileSizePolicyNil

`func (o *FileUptieringParams) SetFileSizePolicyNil(b bool)`

 SetFileSizePolicyNil sets the value for FileSizePolicy to be an explicit nil

### UnsetFileSizePolicy
`func (o *FileUptieringParams) UnsetFileSizePolicy()`

UnsetFileSizePolicy ensures that no value is present for FileSizePolicy, not even an explicit nil
### GetHotFileWindow

`func (o *FileUptieringParams) GetHotFileWindow() int64`

GetHotFileWindow returns the HotFileWindow field if non-nil, zero value otherwise.

### GetHotFileWindowOk

`func (o *FileUptieringParams) GetHotFileWindowOk() (*int64, bool)`

GetHotFileWindowOk returns a tuple with the HotFileWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotFileWindow

`func (o *FileUptieringParams) SetHotFileWindow(v int64)`

SetHotFileWindow sets HotFileWindow field to given value.

### HasHotFileWindow

`func (o *FileUptieringParams) HasHotFileWindow() bool`

HasHotFileWindow returns a boolean if a field has been set.

### SetHotFileWindowNil

`func (o *FileUptieringParams) SetHotFileWindowNil(b bool)`

 SetHotFileWindowNil sets the value for HotFileWindow to be an explicit nil

### UnsetHotFileWindow
`func (o *FileUptieringParams) UnsetHotFileWindow()`

UnsetHotFileWindow ensures that no value is present for HotFileWindow, not even an explicit nil
### GetNfsMountPath

`func (o *FileUptieringParams) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *FileUptieringParams) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *FileUptieringParams) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *FileUptieringParams) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *FileUptieringParams) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *FileUptieringParams) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetNumFileAccess

`func (o *FileUptieringParams) GetNumFileAccess() int32`

GetNumFileAccess returns the NumFileAccess field if non-nil, zero value otherwise.

### GetNumFileAccessOk

`func (o *FileUptieringParams) GetNumFileAccessOk() (*int32, bool)`

GetNumFileAccessOk returns a tuple with the NumFileAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFileAccess

`func (o *FileUptieringParams) SetNumFileAccess(v int32)`

SetNumFileAccess sets NumFileAccess field to given value.

### HasNumFileAccess

`func (o *FileUptieringParams) HasNumFileAccess() bool`

HasNumFileAccess returns a boolean if a field has been set.

### SetNumFileAccessNil

`func (o *FileUptieringParams) SetNumFileAccessNil(b bool)`

 SetNumFileAccessNil sets the value for NumFileAccess to be an explicit nil

### UnsetNumFileAccess
`func (o *FileUptieringParams) UnsetNumFileAccess()`

UnsetNumFileAccess ensures that no value is present for NumFileAccess, not even an explicit nil
### GetSourceViewName

`func (o *FileUptieringParams) GetSourceViewName() string`

GetSourceViewName returns the SourceViewName field if non-nil, zero value otherwise.

### GetSourceViewNameOk

`func (o *FileUptieringParams) GetSourceViewNameOk() (*string, bool)`

GetSourceViewNameOk returns a tuple with the SourceViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceViewName

`func (o *FileUptieringParams) SetSourceViewName(v string)`

SetSourceViewName sets SourceViewName field to given value.

### HasSourceViewName

`func (o *FileUptieringParams) HasSourceViewName() bool`

HasSourceViewName returns a boolean if a field has been set.

### SetSourceViewNameNil

`func (o *FileUptieringParams) SetSourceViewNameNil(b bool)`

 SetSourceViewNameNil sets the value for SourceViewName to be an explicit nil

### UnsetSourceViewName
`func (o *FileUptieringParams) UnsetSourceViewName()`

UnsetSourceViewName ensures that no value is present for SourceViewName, not even an explicit nil
### GetUptierAllFiles

`func (o *FileUptieringParams) GetUptierAllFiles() bool`

GetUptierAllFiles returns the UptierAllFiles field if non-nil, zero value otherwise.

### GetUptierAllFilesOk

`func (o *FileUptieringParams) GetUptierAllFilesOk() (*bool, bool)`

GetUptierAllFilesOk returns a tuple with the UptierAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptierAllFiles

`func (o *FileUptieringParams) SetUptierAllFiles(v bool)`

SetUptierAllFiles sets UptierAllFiles field to given value.

### HasUptierAllFiles

`func (o *FileUptieringParams) HasUptierAllFiles() bool`

HasUptierAllFiles returns a boolean if a field has been set.

### SetUptierAllFilesNil

`func (o *FileUptieringParams) SetUptierAllFilesNil(b bool)`

 SetUptierAllFilesNil sets the value for UptierAllFiles to be an explicit nil

### UnsetUptierAllFiles
`func (o *FileUptieringParams) UnsetUptierAllFiles()`

UnsetUptierAllFiles ensures that no value is present for UptierAllFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


