# CommonTieringPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAuditLogging** | Pointer to **NullableBool** | Specifies whether to audit log the file tiering activity. | [optional] [default to false]
**FilePath** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileSize** | Pointer to [**FileSizePolicy**](FileSizePolicy.md) |  | [optional] 

## Methods

### NewCommonTieringPolicy

`func NewCommonTieringPolicy() *CommonTieringPolicy`

NewCommonTieringPolicy instantiates a new CommonTieringPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTieringPolicyWithDefaults

`func NewCommonTieringPolicyWithDefaults() *CommonTieringPolicy`

NewCommonTieringPolicyWithDefaults instantiates a new CommonTieringPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAuditLogging

`func (o *CommonTieringPolicy) GetEnableAuditLogging() bool`

GetEnableAuditLogging returns the EnableAuditLogging field if non-nil, zero value otherwise.

### GetEnableAuditLoggingOk

`func (o *CommonTieringPolicy) GetEnableAuditLoggingOk() (*bool, bool)`

GetEnableAuditLoggingOk returns a tuple with the EnableAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuditLogging

`func (o *CommonTieringPolicy) SetEnableAuditLogging(v bool)`

SetEnableAuditLogging sets EnableAuditLogging field to given value.

### HasEnableAuditLogging

`func (o *CommonTieringPolicy) HasEnableAuditLogging() bool`

HasEnableAuditLogging returns a boolean if a field has been set.

### SetEnableAuditLoggingNil

`func (o *CommonTieringPolicy) SetEnableAuditLoggingNil(b bool)`

 SetEnableAuditLoggingNil sets the value for EnableAuditLogging to be an explicit nil

### UnsetEnableAuditLogging
`func (o *CommonTieringPolicy) UnsetEnableAuditLogging()`

UnsetEnableAuditLogging ensures that no value is present for EnableAuditLogging, not even an explicit nil
### GetFilePath

`func (o *CommonTieringPolicy) GetFilePath() FileFilteringPolicy`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CommonTieringPolicy) GetFilePathOk() (*FileFilteringPolicy, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CommonTieringPolicy) SetFilePath(v FileFilteringPolicy)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *CommonTieringPolicy) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileSize

`func (o *CommonTieringPolicy) GetFileSize() FileSizePolicy`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *CommonTieringPolicy) GetFileSizeOk() (*FileSizePolicy, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *CommonTieringPolicy) SetFileSize(v FileSizePolicy)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *CommonTieringPolicy) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


