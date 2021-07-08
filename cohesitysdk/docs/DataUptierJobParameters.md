# DataUptierJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSelectionPolicy** | Pointer to **NullableString** | Specifies policy to select a file to uptier based on file access or modification time. eg. A file can be selected to uptier if it has been accessed in the HotFileWindow or it is modified. enum: kLastAccessed, kLastModified. Specifies policy for file selection in data uptier jobs. &#39;kLastAccessed&#39;: Uptier the files which are accessed for at least num_file_access in hot_file_window. &#39;kLastModified&#39;: Uptier the files which are modified. | [optional] 
**FileSizeBytes** | Pointer to **NullableInt64** | Gives the size criteria to be used for selecting the files to be uptiered in bytes. The hot files that are smaller or greater than this size are uptiered. | [optional] 
**FileSizePolicy** | Pointer to **NullableString** | Specifies policy to select a file to uptier based on its size. eg. A file can be selected to uptier if its size is greater than or smaller than the FileSizeBytes. enum: kGreaterThan, kSmallerThan. Specifies policy for file selection in data uptier jobs based on file size. &#39;kGreaterThan&#39;: Uptier the files having size greater than file_size. &#39;kSmallerThan&#39;: Uptier the files having size smaller than file_size. | [optional] 
**HotFileWindow** | Pointer to **NullableInt64** | Identifies the hot files in the NAS source. Files that have been modified in the last hot_file_window are uptiered. Applicable only when file_select_policy is kLastAccessed. | [optional] 
**IncludeAllFiles** | Pointer to **NullableBool** | Specifies whether uptier all files found in the view by overriding the FileUptierSelectionPolicy &amp; FileUptierSizePolicy constraints. Default value false. | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Mount path where the Cohesity target view is mounted on NFS clients while migrating the data. | [optional] 
**NumFileAccess** | Pointer to **NullableInt32** | Number of times file must be accessed within hot_file_window in order to qualify for uptiering. Applicable only when file_select_policy is kLastAccessed. | [optional] 
**SourceViewName** | Pointer to **NullableString** | The source view name from which the data will be uptiered. | [optional] 

## Methods

### NewDataUptierJobParameters

`func NewDataUptierJobParameters() *DataUptierJobParameters`

NewDataUptierJobParameters instantiates a new DataUptierJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUptierJobParametersWithDefaults

`func NewDataUptierJobParametersWithDefaults() *DataUptierJobParameters`

NewDataUptierJobParametersWithDefaults instantiates a new DataUptierJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSelectionPolicy

`func (o *DataUptierJobParameters) GetFileSelectionPolicy() string`

GetFileSelectionPolicy returns the FileSelectionPolicy field if non-nil, zero value otherwise.

### GetFileSelectionPolicyOk

`func (o *DataUptierJobParameters) GetFileSelectionPolicyOk() (*string, bool)`

GetFileSelectionPolicyOk returns a tuple with the FileSelectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSelectionPolicy

`func (o *DataUptierJobParameters) SetFileSelectionPolicy(v string)`

SetFileSelectionPolicy sets FileSelectionPolicy field to given value.

### HasFileSelectionPolicy

`func (o *DataUptierJobParameters) HasFileSelectionPolicy() bool`

HasFileSelectionPolicy returns a boolean if a field has been set.

### SetFileSelectionPolicyNil

`func (o *DataUptierJobParameters) SetFileSelectionPolicyNil(b bool)`

 SetFileSelectionPolicyNil sets the value for FileSelectionPolicy to be an explicit nil

### UnsetFileSelectionPolicy
`func (o *DataUptierJobParameters) UnsetFileSelectionPolicy()`

UnsetFileSelectionPolicy ensures that no value is present for FileSelectionPolicy, not even an explicit nil
### GetFileSizeBytes

`func (o *DataUptierJobParameters) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *DataUptierJobParameters) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *DataUptierJobParameters) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *DataUptierJobParameters) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### SetFileSizeBytesNil

`func (o *DataUptierJobParameters) SetFileSizeBytesNil(b bool)`

 SetFileSizeBytesNil sets the value for FileSizeBytes to be an explicit nil

### UnsetFileSizeBytes
`func (o *DataUptierJobParameters) UnsetFileSizeBytes()`

UnsetFileSizeBytes ensures that no value is present for FileSizeBytes, not even an explicit nil
### GetFileSizePolicy

`func (o *DataUptierJobParameters) GetFileSizePolicy() string`

GetFileSizePolicy returns the FileSizePolicy field if non-nil, zero value otherwise.

### GetFileSizePolicyOk

`func (o *DataUptierJobParameters) GetFileSizePolicyOk() (*string, bool)`

GetFileSizePolicyOk returns a tuple with the FileSizePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizePolicy

`func (o *DataUptierJobParameters) SetFileSizePolicy(v string)`

SetFileSizePolicy sets FileSizePolicy field to given value.

### HasFileSizePolicy

`func (o *DataUptierJobParameters) HasFileSizePolicy() bool`

HasFileSizePolicy returns a boolean if a field has been set.

### SetFileSizePolicyNil

`func (o *DataUptierJobParameters) SetFileSizePolicyNil(b bool)`

 SetFileSizePolicyNil sets the value for FileSizePolicy to be an explicit nil

### UnsetFileSizePolicy
`func (o *DataUptierJobParameters) UnsetFileSizePolicy()`

UnsetFileSizePolicy ensures that no value is present for FileSizePolicy, not even an explicit nil
### GetHotFileWindow

`func (o *DataUptierJobParameters) GetHotFileWindow() int64`

GetHotFileWindow returns the HotFileWindow field if non-nil, zero value otherwise.

### GetHotFileWindowOk

`func (o *DataUptierJobParameters) GetHotFileWindowOk() (*int64, bool)`

GetHotFileWindowOk returns a tuple with the HotFileWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotFileWindow

`func (o *DataUptierJobParameters) SetHotFileWindow(v int64)`

SetHotFileWindow sets HotFileWindow field to given value.

### HasHotFileWindow

`func (o *DataUptierJobParameters) HasHotFileWindow() bool`

HasHotFileWindow returns a boolean if a field has been set.

### SetHotFileWindowNil

`func (o *DataUptierJobParameters) SetHotFileWindowNil(b bool)`

 SetHotFileWindowNil sets the value for HotFileWindow to be an explicit nil

### UnsetHotFileWindow
`func (o *DataUptierJobParameters) UnsetHotFileWindow()`

UnsetHotFileWindow ensures that no value is present for HotFileWindow, not even an explicit nil
### GetIncludeAllFiles

`func (o *DataUptierJobParameters) GetIncludeAllFiles() bool`

GetIncludeAllFiles returns the IncludeAllFiles field if non-nil, zero value otherwise.

### GetIncludeAllFilesOk

`func (o *DataUptierJobParameters) GetIncludeAllFilesOk() (*bool, bool)`

GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllFiles

`func (o *DataUptierJobParameters) SetIncludeAllFiles(v bool)`

SetIncludeAllFiles sets IncludeAllFiles field to given value.

### HasIncludeAllFiles

`func (o *DataUptierJobParameters) HasIncludeAllFiles() bool`

HasIncludeAllFiles returns a boolean if a field has been set.

### SetIncludeAllFilesNil

`func (o *DataUptierJobParameters) SetIncludeAllFilesNil(b bool)`

 SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil

### UnsetIncludeAllFiles
`func (o *DataUptierJobParameters) UnsetIncludeAllFiles()`

UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
### GetNfsMountPath

`func (o *DataUptierJobParameters) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *DataUptierJobParameters) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *DataUptierJobParameters) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *DataUptierJobParameters) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *DataUptierJobParameters) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *DataUptierJobParameters) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetNumFileAccess

`func (o *DataUptierJobParameters) GetNumFileAccess() int32`

GetNumFileAccess returns the NumFileAccess field if non-nil, zero value otherwise.

### GetNumFileAccessOk

`func (o *DataUptierJobParameters) GetNumFileAccessOk() (*int32, bool)`

GetNumFileAccessOk returns a tuple with the NumFileAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFileAccess

`func (o *DataUptierJobParameters) SetNumFileAccess(v int32)`

SetNumFileAccess sets NumFileAccess field to given value.

### HasNumFileAccess

`func (o *DataUptierJobParameters) HasNumFileAccess() bool`

HasNumFileAccess returns a boolean if a field has been set.

### SetNumFileAccessNil

`func (o *DataUptierJobParameters) SetNumFileAccessNil(b bool)`

 SetNumFileAccessNil sets the value for NumFileAccess to be an explicit nil

### UnsetNumFileAccess
`func (o *DataUptierJobParameters) UnsetNumFileAccess()`

UnsetNumFileAccess ensures that no value is present for NumFileAccess, not even an explicit nil
### GetSourceViewName

`func (o *DataUptierJobParameters) GetSourceViewName() string`

GetSourceViewName returns the SourceViewName field if non-nil, zero value otherwise.

### GetSourceViewNameOk

`func (o *DataUptierJobParameters) GetSourceViewNameOk() (*string, bool)`

GetSourceViewNameOk returns a tuple with the SourceViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceViewName

`func (o *DataUptierJobParameters) SetSourceViewName(v string)`

SetSourceViewName sets SourceViewName field to given value.

### HasSourceViewName

`func (o *DataUptierJobParameters) HasSourceViewName() bool`

HasSourceViewName returns a boolean if a field has been set.

### SetSourceViewNameNil

`func (o *DataUptierJobParameters) SetSourceViewNameNil(b bool)`

 SetSourceViewNameNil sets the value for SourceViewName to be an explicit nil

### UnsetSourceViewName
`func (o *DataUptierJobParameters) UnsetSourceViewName()`

UnsetSourceViewName ensures that no value is present for SourceViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


