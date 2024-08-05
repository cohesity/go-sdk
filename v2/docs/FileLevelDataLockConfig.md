# FileLevelDataLockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockAfterDurationIdleMsecs** | Pointer to **NullableInt32** | Specifies the duration to lock a file that has not been accessed or modified (ie. has been idle) for a certain duration of time in milliseconds. Do not set if it is required to disable auto lock. | [optional] 
**CoexistingLockMode** | Pointer to **NullableBool** | Specified if files in the View can be locked in different modes. This property is immutable and can only be set when enabling File level datalock. If this property is set for an S3 View, S3 bucket Versioning should also be enabled. | [optional] 
**DefaultRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a global default retention duration for files in this view, if file lock is enabled for this view. Also, it is a required field if file lock is enabled. Set to -1 if the required default retention period is forever. | [optional] 
**DefaultRetentionDurationYears** | Pointer to **NullableInt64** | Specifies a global default retention duration in years for files in this view, if file/object lock is enabled for this view. | [optional] 
**ExpiryTimestampMsecs** | Pointer to **NullableInt32** | Specifies a definite timestamp in milliseconds for retaining the file. | [optional] 
**LockingProtocol** | Pointer to **NullableString** | Specifies the supported mechanisms to explicity lock a file from NFS/SMB interface. Supported locking protocols: SetReadOnly, SetAtime. &#39;SetReadOnly&#39; is compatible with Isilon/Netapp behaviour. This locks the file and the retention duration is determined in this order: 1) atime, if set by user/application and within min and max retention duration. 2) Min retention duration, if set. 3) Otherwise, file is switched to expired data automatically. &#39;SetAtime&#39; is compatible with Data Domain behaviour. | [optional] 
**MaxRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a maximum duration in milliseconds for which any file in this view can be retained for. Set to -1 if the required retention duration is forever. If set, it should be greater than or equal to the default retention period as well as the min retention period. | [optional] 
**MinRetentionDurationMsecs** | Pointer to **NullableInt64** | Specifies a minimum retention duration in milliseconds after a file gets locked. The file cannot be modified or deleted during this timeframe. Set to -1 if the required retention duration is forever. This should be set less than or equal to the default retention duration. | [optional] 
**Mode** | Pointer to **NullableString** | Specifies the mode of file level datalock. Enterprise mode can be upgraded to Compliance mode, but Compliance mode cannot be downgraded to Enterprise mode. Compliance: This mode would disallow all user to delete/modify file or view under any condition when it &#39;s in locked status except for deleting view when the view is empty. Enterprise: This mode would follow the rules as compliance mode for normal users. But it would allow the storage admin (1) to delete view or file anytime no matter it is in locked status or expired. (2) to rename the view (3) to bring back the retention period when it&#39;s in locked mode A lock mode of a file in a view can be in one of the following: &#39;Compliance&#39;: Default mode of datalock, in this mode, Data Security Admin cannot modify/delete this view when datalock is in effect. Data Security Admin can delete this view when datalock is expired. &#39;kEnterprise&#39; : In this mode, Data Security Admin can change view name or delete view when datalock is in effect. Datalock in this mode can be upgraded to &#39;Compliance&#39; mode. | [optional] 

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

### GetAutoLockAfterDurationIdleMsecs

`func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdleMsecs() int32`

GetAutoLockAfterDurationIdleMsecs returns the AutoLockAfterDurationIdleMsecs field if non-nil, zero value otherwise.

### GetAutoLockAfterDurationIdleMsecsOk

`func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdleMsecsOk() (*int32, bool)`

GetAutoLockAfterDurationIdleMsecsOk returns a tuple with the AutoLockAfterDurationIdleMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockAfterDurationIdleMsecs

`func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdleMsecs(v int32)`

SetAutoLockAfterDurationIdleMsecs sets AutoLockAfterDurationIdleMsecs field to given value.

### HasAutoLockAfterDurationIdleMsecs

`func (o *FileLevelDataLockConfig) HasAutoLockAfterDurationIdleMsecs() bool`

HasAutoLockAfterDurationIdleMsecs returns a boolean if a field has been set.

### SetAutoLockAfterDurationIdleMsecsNil

`func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdleMsecsNil(b bool)`

 SetAutoLockAfterDurationIdleMsecsNil sets the value for AutoLockAfterDurationIdleMsecs to be an explicit nil

### UnsetAutoLockAfterDurationIdleMsecs
`func (o *FileLevelDataLockConfig) UnsetAutoLockAfterDurationIdleMsecs()`

UnsetAutoLockAfterDurationIdleMsecs ensures that no value is present for AutoLockAfterDurationIdleMsecs, not even an explicit nil
### GetCoexistingLockMode

`func (o *FileLevelDataLockConfig) GetCoexistingLockMode() bool`

GetCoexistingLockMode returns the CoexistingLockMode field if non-nil, zero value otherwise.

### GetCoexistingLockModeOk

`func (o *FileLevelDataLockConfig) GetCoexistingLockModeOk() (*bool, bool)`

GetCoexistingLockModeOk returns a tuple with the CoexistingLockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoexistingLockMode

`func (o *FileLevelDataLockConfig) SetCoexistingLockMode(v bool)`

SetCoexistingLockMode sets CoexistingLockMode field to given value.

### HasCoexistingLockMode

`func (o *FileLevelDataLockConfig) HasCoexistingLockMode() bool`

HasCoexistingLockMode returns a boolean if a field has been set.

### SetCoexistingLockModeNil

`func (o *FileLevelDataLockConfig) SetCoexistingLockModeNil(b bool)`

 SetCoexistingLockModeNil sets the value for CoexistingLockMode to be an explicit nil

### UnsetCoexistingLockMode
`func (o *FileLevelDataLockConfig) UnsetCoexistingLockMode()`

UnsetCoexistingLockMode ensures that no value is present for CoexistingLockMode, not even an explicit nil
### GetDefaultRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationMsecs() int64`

GetDefaultRetentionDurationMsecs returns the DefaultRetentionDurationMsecs field if non-nil, zero value otherwise.

### GetDefaultRetentionDurationMsecsOk

`func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationMsecsOk() (*int64, bool)`

GetDefaultRetentionDurationMsecsOk returns a tuple with the DefaultRetentionDurationMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationMsecs(v int64)`

SetDefaultRetentionDurationMsecs sets DefaultRetentionDurationMsecs field to given value.

### HasDefaultRetentionDurationMsecs

`func (o *FileLevelDataLockConfig) HasDefaultRetentionDurationMsecs() bool`

HasDefaultRetentionDurationMsecs returns a boolean if a field has been set.

### SetDefaultRetentionDurationMsecsNil

`func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationMsecsNil(b bool)`

 SetDefaultRetentionDurationMsecsNil sets the value for DefaultRetentionDurationMsecs to be an explicit nil

### UnsetDefaultRetentionDurationMsecs
`func (o *FileLevelDataLockConfig) UnsetDefaultRetentionDurationMsecs()`

UnsetDefaultRetentionDurationMsecs ensures that no value is present for DefaultRetentionDurationMsecs, not even an explicit nil
### GetDefaultRetentionDurationYears

`func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationYears() int64`

GetDefaultRetentionDurationYears returns the DefaultRetentionDurationYears field if non-nil, zero value otherwise.

### GetDefaultRetentionDurationYearsOk

`func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationYearsOk() (*int64, bool)`

GetDefaultRetentionDurationYearsOk returns a tuple with the DefaultRetentionDurationYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRetentionDurationYears

`func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationYears(v int64)`

SetDefaultRetentionDurationYears sets DefaultRetentionDurationYears field to given value.

### HasDefaultRetentionDurationYears

`func (o *FileLevelDataLockConfig) HasDefaultRetentionDurationYears() bool`

HasDefaultRetentionDurationYears returns a boolean if a field has been set.

### SetDefaultRetentionDurationYearsNil

`func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationYearsNil(b bool)`

 SetDefaultRetentionDurationYearsNil sets the value for DefaultRetentionDurationYears to be an explicit nil

### UnsetDefaultRetentionDurationYears
`func (o *FileLevelDataLockConfig) UnsetDefaultRetentionDurationYears()`

UnsetDefaultRetentionDurationYears ensures that no value is present for DefaultRetentionDurationYears, not even an explicit nil
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


