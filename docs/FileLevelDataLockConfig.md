# FileLevelDataLockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockAfterDurationIdle** | Pointer to **NullableInt32** | Specifies the duration to lock a file that has not been accessed or modified (ie. has been idle) for a certain duration of time in milliseconds. Do not set if it is required to disable auto lock. | [optional] 
**DefaultFileRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a global default retention duration for files in this view, if file lock is enabled for this view. Also, it is a required field if file lock is enabled. Set to -1 if the required default retention period is forever. | [optional] 
**ExpiryTimestampMsecs** | Pointer to **NullableInt32** | Specifies a definite timestamp in milliseconds for retaining the file. | [optional] 
**LockingProtocol** | Pointer to **NullableString** | Specifies the supported mechanisms to explicity lock a file from NFS/SMB interface. Supported locking protocols: kSetReadOnly, kSetAtime. &#39;kSetReadOnly&#39; is compatible with Isilon/Netapp behaviour. This locks the file and the retention duration is determined in this order: 1) atime, if set by user/application and within min and max retention duration. 2) Min retention duration, if set. 3) Otherwise, file is switched to expired data automatically. &#39;kSetAtime&#39; is compatible with Data Domain behaviour. | [optional] 
**MaxRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a maximum duration in milliseconds for which any file in this view can be retained for. Set to -1 if the required retention duration is forever. If set, it should be greater than or equal to the default retention period as well as the min retention period. | [optional] 
**MinRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a minimum retention duration in milliseconds after a file gets locked. The file cannot be modified or deleted during this timeframe. Set to -1 if the required retention duration is forever. This should be set less than or equal to the default retention duration. | [optional] 
**Mode** | Pointer to **NullableString** | Specifies the mode of file level datalock. Enterprise mode can be upgraded to Compliance mode, but Compliance mode cannot be downgraded to Enterprise mode. kCompliance: This mode would disallow all user to delete/modify file or view under any condition when it &#39;s in locked status except for deleting view when the view is empty. kEnterprise: This mode would follow the rules as compliance mode for normal users. But it would allow the storage admin (1) to delete view or file anytime no matter it is in locked status or expired. (2) to rename the view (3) to bring back the retention period when it&#39;s in locked mode A lock mode of a file in a view can be in one of the following:  &#39;kCompliance&#39;: Default mode of datalock, in this mode, Data Security Admin cannot modify/delete this view when datalock is in effect. Data Security Admin can delete this view when datalock is expired. &#39;kEnterprise&#39; : In this mode, Data Security Admin can change view name or delete view when datalock is in effect. Datalock in this mode can be upgraded to &#39;kCompliance&#39; mode. | [optional] 

## Methods

### NewFileLevelDataLockConfig

`func NewFileLevelDataLockConfig() *FileLevelDataLockConfig`

NewFileLevelDataLockConfig instantiates a new FileLevelDataLockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLevelDataLockConfigWithDefaults

`func NewFileLevelDataLockConfigWithDefaults() *FileLevelDataLockConfig`

NewFileLevelDataLockConfigWithDefaults instantiates a new FileLevelDataLockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLockAfterDurationIdle

`func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdle() int32`

GetAutoLockAfterDurationIdle returns the AutoLockAfterDurationIdle field if non-nil, zero value otherwise.

### GetAutoLockAfterDurationIdleOk

`func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdleOk() (*int32, bool)`

GetAutoLockAfterDurationIdleOk returns a tuple with the AutoLockAfterDurationIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockAfterDurationIdle

`func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdle(v int32)`

SetAutoLockAfterDurationIdle sets AutoLockAfterDurationIdle field to given value.

### HasAutoLockAfterDurationIdle

`func (o *FileLevelDataLockConfig) HasAutoLockAfterDurationIdle() bool`

HasAutoLockAfterDurationIdle returns a boolean if a field has been set.

### SetAutoLockAfterDurationIdleNil

`func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdleNil(b bool)`

 SetAutoLockAfterDurationIdleNil sets the value for AutoLockAfterDurationIdle to be an explicit nil

### UnsetAutoLockAfterDurationIdle
`func (o *FileLevelDataLockConfig) UnsetAutoLockAfterDurationIdle()`

UnsetAutoLockAfterDurationIdle ensures that no value is present for AutoLockAfterDurationIdle, not even an explicit nil
### GetDefaultFileRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) GetDefaultFileRetentionDurationMsecs() int64`

GetDefaultFileRetentionDurationMsecs returns the DefaultFileRetentionDurationMsecs field if non-nil, zero value otherwise.

### GetDefaultFileRetentionDurationMsecsOk

`func (o *FileLevelDataLockConfig) GetDefaultFileRetentionDurationMsecsOk() (*int64, bool)`

GetDefaultFileRetentionDurationMsecsOk returns a tuple with the DefaultFileRetentionDurationMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFileRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) SetDefaultFileRetentionDurationMsecs(v int64)`

SetDefaultFileRetentionDurationMsecs sets DefaultFileRetentionDurationMsecs field to given value.

### HasDefaultFileRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) HasDefaultFileRetentionDurationMsecs() bool`

HasDefaultFileRetentionDurationMsecs returns a boolean if a field has been set.

### SetDefaultFileRetentionDurationMsecsNil

`func (o *FileLevelDataLockConfig) SetDefaultFileRetentionDurationMsecsNil(b bool)`

 SetDefaultFileRetentionDurationMsecsNil sets the value for DefaultFileRetentionDurationMsecs to be an explicit nil

### UnsetDefaultFileRetentionDurationMsecs
`func (o *FileLevelDataLockConfig) UnsetDefaultFileRetentionDurationMsecs()`

UnsetDefaultFileRetentionDurationMsecs ensures that no value is present for DefaultFileRetentionDurationMsecs, not even an explicit nil
### GetExpiryTimestampMsecs

`func (o *FileLevelDataLockConfig) GetExpiryTimestampMsecs() int32`

GetExpiryTimestampMsecs returns the ExpiryTimestampMsecs field if non-nil, zero value otherwise.

### GetExpiryTimestampMsecsOk

`func (o *FileLevelDataLockConfig) GetExpiryTimestampMsecsOk() (*int32, bool)`

GetExpiryTimestampMsecsOk returns a tuple with the ExpiryTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestampMsecs

`func (o *FileLevelDataLockConfig) SetExpiryTimestampMsecs(v int32)`

SetExpiryTimestampMsecs sets ExpiryTimestampMsecs field to given value.

### HasExpiryTimestampMsecs

`func (o *FileLevelDataLockConfig) HasExpiryTimestampMsecs() bool`

HasExpiryTimestampMsecs returns a boolean if a field has been set.

### SetExpiryTimestampMsecsNil

`func (o *FileLevelDataLockConfig) SetExpiryTimestampMsecsNil(b bool)`

 SetExpiryTimestampMsecsNil sets the value for ExpiryTimestampMsecs to be an explicit nil

### UnsetExpiryTimestampMsecs
`func (o *FileLevelDataLockConfig) UnsetExpiryTimestampMsecs()`

UnsetExpiryTimestampMsecs ensures that no value is present for ExpiryTimestampMsecs, not even an explicit nil
### GetLockingProtocol

`func (o *FileLevelDataLockConfig) GetLockingProtocol() string`

GetLockingProtocol returns the LockingProtocol field if non-nil, zero value otherwise.

### GetLockingProtocolOk

`func (o *FileLevelDataLockConfig) GetLockingProtocolOk() (*string, bool)`

GetLockingProtocolOk returns a tuple with the LockingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockingProtocol

`func (o *FileLevelDataLockConfig) SetLockingProtocol(v string)`

SetLockingProtocol sets LockingProtocol field to given value.

### HasLockingProtocol

`func (o *FileLevelDataLockConfig) HasLockingProtocol() bool`

HasLockingProtocol returns a boolean if a field has been set.

### SetLockingProtocolNil

`func (o *FileLevelDataLockConfig) SetLockingProtocolNil(b bool)`

 SetLockingProtocolNil sets the value for LockingProtocol to be an explicit nil

### UnsetLockingProtocol
`func (o *FileLevelDataLockConfig) UnsetLockingProtocol()`

UnsetLockingProtocol ensures that no value is present for LockingProtocol, not even an explicit nil
### GetMaxRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) GetMaxRetentionDurationMsecs() int64`

GetMaxRetentionDurationMsecs returns the MaxRetentionDurationMsecs field if non-nil, zero value otherwise.

### GetMaxRetentionDurationMsecsOk

`func (o *FileLevelDataLockConfig) GetMaxRetentionDurationMsecsOk() (*int64, bool)`

GetMaxRetentionDurationMsecsOk returns a tuple with the MaxRetentionDurationMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) SetMaxRetentionDurationMsecs(v int64)`

SetMaxRetentionDurationMsecs sets MaxRetentionDurationMsecs field to given value.

### HasMaxRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) HasMaxRetentionDurationMsecs() bool`

HasMaxRetentionDurationMsecs returns a boolean if a field has been set.

### SetMaxRetentionDurationMsecsNil

`func (o *FileLevelDataLockConfig) SetMaxRetentionDurationMsecsNil(b bool)`

 SetMaxRetentionDurationMsecsNil sets the value for MaxRetentionDurationMsecs to be an explicit nil

### UnsetMaxRetentionDurationMsecs
`func (o *FileLevelDataLockConfig) UnsetMaxRetentionDurationMsecs()`

UnsetMaxRetentionDurationMsecs ensures that no value is present for MaxRetentionDurationMsecs, not even an explicit nil
### GetMinRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) GetMinRetentionDurationMsecs() int64`

GetMinRetentionDurationMsecs returns the MinRetentionDurationMsecs field if non-nil, zero value otherwise.

### GetMinRetentionDurationMsecsOk

`func (o *FileLevelDataLockConfig) GetMinRetentionDurationMsecsOk() (*int64, bool)`

GetMinRetentionDurationMsecsOk returns a tuple with the MinRetentionDurationMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) SetMinRetentionDurationMsecs(v int64)`

SetMinRetentionDurationMsecs sets MinRetentionDurationMsecs field to given value.

### HasMinRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) HasMinRetentionDurationMsecs() bool`

HasMinRetentionDurationMsecs returns a boolean if a field has been set.

### SetMinRetentionDurationMsecsNil

`func (o *FileLevelDataLockConfig) SetMinRetentionDurationMsecsNil(b bool)`

 SetMinRetentionDurationMsecsNil sets the value for MinRetentionDurationMsecs to be an explicit nil

### UnsetMinRetentionDurationMsecs
`func (o *FileLevelDataLockConfig) UnsetMinRetentionDurationMsecs()`

UnsetMinRetentionDurationMsecs ensures that no value is present for MinRetentionDurationMsecs, not even an explicit nil
### GetMode

`func (o *FileLevelDataLockConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileLevelDataLockConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileLevelDataLockConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileLevelDataLockConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *FileLevelDataLockConfig) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *FileLevelDataLockConfig) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


