# ViewIdMappingProtoFileLevelDataLockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLockDurationUsecs** | Pointer to **NullableInt64** | Auto-lock automatically commit files to WORM state in the filesystem if they have not been modified for an administrator-specified period of time. When the auto-lock is enabled, this field must be set to idle time duration after which file would be automatically locked. Auto locking will be disabled when configured with default value of -1. | [optional] 
**DefaultRetentionDurationUsecs** | Pointer to **NullableInt64** | Default retention duration is used when an explicit retention timestamp is not set by user/application when locking a file. If the administrator does not want to enforce this, this field must not be set. If file requires being retained forever by default, this must be set to INT64_MAX. If minimum and maximum retention are enforced, then this must be always between these two durations. | [optional] 
**HoldTimestampUsecs** | Pointer to **NullableInt64** | Specifies timestamp to protect locked files until a specific date. This would override retention periods and deny any mutable or remove operations on locked files until a specific date. | [optional] 
**MaxRetentionDurationUsecs** | Pointer to **NullableInt64** | Specifies maximum retention duration of worm locked file. If the administrator does not want to enforce this, this must not be set. If default and max retention duration are enforced, max retention duration must be greater than or equal to default retention duration. If min and max retention duration are enforced, max retention duration must be greater than and equal to min retention duration. | [optional] 
**MinRetentionDurationUsecs** | Pointer to **NullableInt64** | Minimum and maximum retention duration allow the administrator to enforce retention duration that falls within a specified range. If the administrator does not want to enforce this, this must not be set. If the file requires being retained forever, this must be set to INT64_MAX. If default retention is enforced, this must be less than or equal to default retention. If max retention are enforced, default retention duration must be less than and equal to max retention duration. | [optional] 
**Mode** | Pointer to **NullableInt32** | Explicit locking mode. | [optional] 
**Protocol** | Pointer to **NullableInt32** | Explicit locking protocol. | [optional] 

## Methods

### NewViewIdMappingProtoFileLevelDataLockConfig

`func NewViewIdMappingProtoFileLevelDataLockConfig() *ViewIdMappingProtoFileLevelDataLockConfig`

NewViewIdMappingProtoFileLevelDataLockConfig instantiates a new ViewIdMappingProtoFileLevelDataLockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIdMappingProtoFileLevelDataLockConfigWithDefaults

`func NewViewIdMappingProtoFileLevelDataLockConfigWithDefaults() *ViewIdMappingProtoFileLevelDataLockConfig`

NewViewIdMappingProtoFileLevelDataLockConfigWithDefaults instantiates a new ViewIdMappingProtoFileLevelDataLockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLockDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetAutoLockDurationUsecs() int64`

GetAutoLockDurationUsecs returns the AutoLockDurationUsecs field if non-nil, zero value otherwise.

### GetAutoLockDurationUsecsOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetAutoLockDurationUsecsOk() (*int64, bool)`

GetAutoLockDurationUsecsOk returns a tuple with the AutoLockDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLockDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetAutoLockDurationUsecs(v int64)`

SetAutoLockDurationUsecs sets AutoLockDurationUsecs field to given value.

### HasAutoLockDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasAutoLockDurationUsecs() bool`

HasAutoLockDurationUsecs returns a boolean if a field has been set.

### SetAutoLockDurationUsecsNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetAutoLockDurationUsecsNil(b bool)`

 SetAutoLockDurationUsecsNil sets the value for AutoLockDurationUsecs to be an explicit nil

### UnsetAutoLockDurationUsecs
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetAutoLockDurationUsecs()`

UnsetAutoLockDurationUsecs ensures that no value is present for AutoLockDurationUsecs, not even an explicit nil
### GetDefaultRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetDefaultRetentionDurationUsecs() int64`

GetDefaultRetentionDurationUsecs returns the DefaultRetentionDurationUsecs field if non-nil, zero value otherwise.

### GetDefaultRetentionDurationUsecsOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetDefaultRetentionDurationUsecsOk() (*int64, bool)`

GetDefaultRetentionDurationUsecsOk returns a tuple with the DefaultRetentionDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetDefaultRetentionDurationUsecs(v int64)`

SetDefaultRetentionDurationUsecs sets DefaultRetentionDurationUsecs field to given value.

### HasDefaultRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasDefaultRetentionDurationUsecs() bool`

HasDefaultRetentionDurationUsecs returns a boolean if a field has been set.

### SetDefaultRetentionDurationUsecsNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetDefaultRetentionDurationUsecsNil(b bool)`

 SetDefaultRetentionDurationUsecsNil sets the value for DefaultRetentionDurationUsecs to be an explicit nil

### UnsetDefaultRetentionDurationUsecs
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetDefaultRetentionDurationUsecs()`

UnsetDefaultRetentionDurationUsecs ensures that no value is present for DefaultRetentionDurationUsecs, not even an explicit nil
### GetHoldTimestampUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetHoldTimestampUsecs() int64`

GetHoldTimestampUsecs returns the HoldTimestampUsecs field if non-nil, zero value otherwise.

### GetHoldTimestampUsecsOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetHoldTimestampUsecsOk() (*int64, bool)`

GetHoldTimestampUsecsOk returns a tuple with the HoldTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimestampUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetHoldTimestampUsecs(v int64)`

SetHoldTimestampUsecs sets HoldTimestampUsecs field to given value.

### HasHoldTimestampUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasHoldTimestampUsecs() bool`

HasHoldTimestampUsecs returns a boolean if a field has been set.

### SetHoldTimestampUsecsNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetHoldTimestampUsecsNil(b bool)`

 SetHoldTimestampUsecsNil sets the value for HoldTimestampUsecs to be an explicit nil

### UnsetHoldTimestampUsecs
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetHoldTimestampUsecs()`

UnsetHoldTimestampUsecs ensures that no value is present for HoldTimestampUsecs, not even an explicit nil
### GetMaxRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetMaxRetentionDurationUsecs() int64`

GetMaxRetentionDurationUsecs returns the MaxRetentionDurationUsecs field if non-nil, zero value otherwise.

### GetMaxRetentionDurationUsecsOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetMaxRetentionDurationUsecsOk() (*int64, bool)`

GetMaxRetentionDurationUsecsOk returns a tuple with the MaxRetentionDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetMaxRetentionDurationUsecs(v int64)`

SetMaxRetentionDurationUsecs sets MaxRetentionDurationUsecs field to given value.

### HasMaxRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasMaxRetentionDurationUsecs() bool`

HasMaxRetentionDurationUsecs returns a boolean if a field has been set.

### SetMaxRetentionDurationUsecsNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetMaxRetentionDurationUsecsNil(b bool)`

 SetMaxRetentionDurationUsecsNil sets the value for MaxRetentionDurationUsecs to be an explicit nil

### UnsetMaxRetentionDurationUsecs
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetMaxRetentionDurationUsecs()`

UnsetMaxRetentionDurationUsecs ensures that no value is present for MaxRetentionDurationUsecs, not even an explicit nil
### GetMinRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetMinRetentionDurationUsecs() int64`

GetMinRetentionDurationUsecs returns the MinRetentionDurationUsecs field if non-nil, zero value otherwise.

### GetMinRetentionDurationUsecsOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetMinRetentionDurationUsecsOk() (*int64, bool)`

GetMinRetentionDurationUsecsOk returns a tuple with the MinRetentionDurationUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetMinRetentionDurationUsecs(v int64)`

SetMinRetentionDurationUsecs sets MinRetentionDurationUsecs field to given value.

### HasMinRetentionDurationUsecs

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasMinRetentionDurationUsecs() bool`

HasMinRetentionDurationUsecs returns a boolean if a field has been set.

### SetMinRetentionDurationUsecsNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetMinRetentionDurationUsecsNil(b bool)`

 SetMinRetentionDurationUsecsNil sets the value for MinRetentionDurationUsecs to be an explicit nil

### UnsetMinRetentionDurationUsecs
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetMinRetentionDurationUsecs()`

UnsetMinRetentionDurationUsecs ensures that no value is present for MinRetentionDurationUsecs, not even an explicit nil
### GetMode

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetProtocol

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *ViewIdMappingProtoFileLevelDataLockConfig) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *ViewIdMappingProtoFileLevelDataLockConfig) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


