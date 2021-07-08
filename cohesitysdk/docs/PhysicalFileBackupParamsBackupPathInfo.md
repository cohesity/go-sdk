# PhysicalFileBackupParamsBackupPathInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludePaths** | Pointer to **[]string** | A list of absolute paths on the Physical source that should not be backed up. Any path that is a descendant of these paths will also be skipped. | [optional] 
**IncludePath** | Pointer to **NullableString** | An absolute path on the Physical source that should be backed up. Any path that is a descendant of this path will also be backed up, unless (a) it is excluded from backup by exclude_paths below, OR (b) it belongs to a volume that is different from the volume include_path belongs to and there are no relevant paths in that volume being backed up. | [optional] 
**SkipNestedVolumes** | Pointer to **NullableBool** | Whether to skip any nested volumes (both local and network) that are mounted under &#39;include_path&#39;. Note that if a path to a nested volume is specified as an include_path in another BackupPathInfo entry, that path will still get backed up. | [optional] 

## Methods

### NewPhysicalFileBackupParamsBackupPathInfo

`func NewPhysicalFileBackupParamsBackupPathInfo() *PhysicalFileBackupParamsBackupPathInfo`

NewPhysicalFileBackupParamsBackupPathInfo instantiates a new PhysicalFileBackupParamsBackupPathInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalFileBackupParamsBackupPathInfoWithDefaults

`func NewPhysicalFileBackupParamsBackupPathInfoWithDefaults() *PhysicalFileBackupParamsBackupPathInfo`

NewPhysicalFileBackupParamsBackupPathInfoWithDefaults instantiates a new PhysicalFileBackupParamsBackupPathInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludePaths

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *PhysicalFileBackupParamsBackupPathInfo) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### SetExcludePathsNil

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetExcludePathsNil(b bool)`

 SetExcludePathsNil sets the value for ExcludePaths to be an explicit nil

### UnsetExcludePaths
`func (o *PhysicalFileBackupParamsBackupPathInfo) UnsetExcludePaths()`

UnsetExcludePaths ensures that no value is present for ExcludePaths, not even an explicit nil
### GetIncludePath

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetIncludePath() string`

GetIncludePath returns the IncludePath field if non-nil, zero value otherwise.

### GetIncludePathOk

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetIncludePathOk() (*string, bool)`

GetIncludePathOk returns a tuple with the IncludePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePath

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetIncludePath(v string)`

SetIncludePath sets IncludePath field to given value.

### HasIncludePath

`func (o *PhysicalFileBackupParamsBackupPathInfo) HasIncludePath() bool`

HasIncludePath returns a boolean if a field has been set.

### SetIncludePathNil

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetIncludePathNil(b bool)`

 SetIncludePathNil sets the value for IncludePath to be an explicit nil

### UnsetIncludePath
`func (o *PhysicalFileBackupParamsBackupPathInfo) UnsetIncludePath()`

UnsetIncludePath ensures that no value is present for IncludePath, not even an explicit nil
### GetSkipNestedVolumes

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetSkipNestedVolumes() bool`

GetSkipNestedVolumes returns the SkipNestedVolumes field if non-nil, zero value otherwise.

### GetSkipNestedVolumesOk

`func (o *PhysicalFileBackupParamsBackupPathInfo) GetSkipNestedVolumesOk() (*bool, bool)`

GetSkipNestedVolumesOk returns a tuple with the SkipNestedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNestedVolumes

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetSkipNestedVolumes(v bool)`

SetSkipNestedVolumes sets SkipNestedVolumes field to given value.

### HasSkipNestedVolumes

`func (o *PhysicalFileBackupParamsBackupPathInfo) HasSkipNestedVolumes() bool`

HasSkipNestedVolumes returns a boolean if a field has been set.

### SetSkipNestedVolumesNil

`func (o *PhysicalFileBackupParamsBackupPathInfo) SetSkipNestedVolumesNil(b bool)`

 SetSkipNestedVolumesNil sets the value for SkipNestedVolumes to be an explicit nil

### UnsetSkipNestedVolumes
`func (o *PhysicalFileBackupParamsBackupPathInfo) UnsetSkipNestedVolumes()`

UnsetSkipNestedVolumes ensures that no value is present for SkipNestedVolumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


