# UptieringPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileAge** | Pointer to [**UptieringFileAgePolicy**](UptieringFileAgePolicy.md) |  | [optional] 
**IncludeAllFiles** | Pointer to **NullableBool** | If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints. | [optional] [default to false]
**Target** | Pointer to [**NullableUptieringTarget**](UptieringTarget.md) |  | [optional] 
**EnableAuditLogging** | Pointer to **NullableBool** | Specifies whether to audit log the file tiering activity. | [optional] [default to false]
**FileSize** | Pointer to [**FileSizePolicy**](FileSizePolicy.md) |  | [optional] 
**FilePath** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 

## Methods

### NewUptieringPolicy

`func NewUptieringPolicy() *UptieringPolicy`

NewUptieringPolicy instantiates a new UptieringPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptieringPolicyWithDefaults

`func NewUptieringPolicyWithDefaults() *UptieringPolicy`

NewUptieringPolicyWithDefaults instantiates a new UptieringPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileAge

`func (o *UptieringPolicy) GetFileAge() UptieringFileAgePolicy`

GetFileAge returns the FileAge field if non-nil, zero value otherwise.

### GetFileAgeOk

`func (o *UptieringPolicy) GetFileAgeOk() (*UptieringFileAgePolicy, bool)`

GetFileAgeOk returns a tuple with the FileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAge

`func (o *UptieringPolicy) SetFileAge(v UptieringFileAgePolicy)`

SetFileAge sets FileAge field to given value.

### HasFileAge

`func (o *UptieringPolicy) HasFileAge() bool`

HasFileAge returns a boolean if a field has been set.

### GetIncludeAllFiles

`func (o *UptieringPolicy) GetIncludeAllFiles() bool`

GetIncludeAllFiles returns the IncludeAllFiles field if non-nil, zero value otherwise.

### GetIncludeAllFilesOk

`func (o *UptieringPolicy) GetIncludeAllFilesOk() (*bool, bool)`

GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllFiles

`func (o *UptieringPolicy) SetIncludeAllFiles(v bool)`

SetIncludeAllFiles sets IncludeAllFiles field to given value.

### HasIncludeAllFiles

`func (o *UptieringPolicy) HasIncludeAllFiles() bool`

HasIncludeAllFiles returns a boolean if a field has been set.

### SetIncludeAllFilesNil

`func (o *UptieringPolicy) SetIncludeAllFilesNil(b bool)`

 SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil

### UnsetIncludeAllFiles
`func (o *UptieringPolicy) UnsetIncludeAllFiles()`

UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
### GetTarget

`func (o *UptieringPolicy) GetTarget() UptieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UptieringPolicy) GetTargetOk() (*UptieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UptieringPolicy) SetTarget(v UptieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *UptieringPolicy) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *UptieringPolicy) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *UptieringPolicy) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetEnableAuditLogging

`func (o *UptieringPolicy) GetEnableAuditLogging() bool`

GetEnableAuditLogging returns the EnableAuditLogging field if non-nil, zero value otherwise.

### GetEnableAuditLoggingOk

`func (o *UptieringPolicy) GetEnableAuditLoggingOk() (*bool, bool)`

GetEnableAuditLoggingOk returns a tuple with the EnableAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuditLogging

`func (o *UptieringPolicy) SetEnableAuditLogging(v bool)`

SetEnableAuditLogging sets EnableAuditLogging field to given value.

### HasEnableAuditLogging

`func (o *UptieringPolicy) HasEnableAuditLogging() bool`

HasEnableAuditLogging returns a boolean if a field has been set.

### SetEnableAuditLoggingNil

`func (o *UptieringPolicy) SetEnableAuditLoggingNil(b bool)`

 SetEnableAuditLoggingNil sets the value for EnableAuditLogging to be an explicit nil

### UnsetEnableAuditLogging
`func (o *UptieringPolicy) UnsetEnableAuditLogging()`

UnsetEnableAuditLogging ensures that no value is present for EnableAuditLogging, not even an explicit nil
### GetFileSize

`func (o *UptieringPolicy) GetFileSize() FileSizePolicy`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *UptieringPolicy) GetFileSizeOk() (*FileSizePolicy, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *UptieringPolicy) SetFileSize(v FileSizePolicy)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *UptieringPolicy) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFilePath

`func (o *UptieringPolicy) GetFilePath() FileFilteringPolicy`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *UptieringPolicy) GetFilePathOk() (*FileFilteringPolicy, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *UptieringPolicy) SetFilePath(v FileFilteringPolicy)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *UptieringPolicy) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


