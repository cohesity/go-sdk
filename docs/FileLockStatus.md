# FileLockStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimestampMsecs** | Pointer to **NullableInt64** | Specifies a expiry timestamp in milliseconds until the file is locked. | [optional] 
**HoldTimestampMsecs** | Pointer to **NullableInt64** | Specifies a override timestamp in milliseconds when an expired file is kept on hold. | [optional] 
**LockTimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which the file was locked. | [optional] 
**Mode** | Pointer to **NullableString** | Specifies the mode of the file lock. &#39;kCompliance&#39;, &#39;kEnterprise&#39;. A lock mode of a file in a view can be in one of the following:  &#39;kCompliance&#39;: Default mode of datalock, in this mode, Data Security Admin cannot modify/delete this view when datalock is in effect. Data Security Admin can delete this view when datalock is expired. &#39;kEnterprise&#39; : In this mode, Data Security Admin can change view name or delete view when datalock is in effect. Datalock in this mode can be upgraded to &#39;kCompliance&#39; mode. | [optional] 
**State** | Pointer to **NullableInt32** | Specifies the lock state of the file. | [optional] 

## Methods

### NewFileLockStatus

`func NewFileLockStatus() *FileLockStatus`

NewFileLockStatus instantiates a new FileLockStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLockStatusWithDefaults

`func NewFileLockStatusWithDefaults() *FileLockStatus`

NewFileLockStatusWithDefaults instantiates a new FileLockStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTimestampMsecs

`func (o *FileLockStatus) GetExpiryTimestampMsecs() int64`

GetExpiryTimestampMsecs returns the ExpiryTimestampMsecs field if non-nil, zero value otherwise.

### GetExpiryTimestampMsecsOk

`func (o *FileLockStatus) GetExpiryTimestampMsecsOk() (*int64, bool)`

GetExpiryTimestampMsecsOk returns a tuple with the ExpiryTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestampMsecs

`func (o *FileLockStatus) SetExpiryTimestampMsecs(v int64)`

SetExpiryTimestampMsecs sets ExpiryTimestampMsecs field to given value.

### HasExpiryTimestampMsecs

`func (o *FileLockStatus) HasExpiryTimestampMsecs() bool`

HasExpiryTimestampMsecs returns a boolean if a field has been set.

### SetExpiryTimestampMsecsNil

`func (o *FileLockStatus) SetExpiryTimestampMsecsNil(b bool)`

 SetExpiryTimestampMsecsNil sets the value for ExpiryTimestampMsecs to be an explicit nil

### UnsetExpiryTimestampMsecs
`func (o *FileLockStatus) UnsetExpiryTimestampMsecs()`

UnsetExpiryTimestampMsecs ensures that no value is present for ExpiryTimestampMsecs, not even an explicit nil
### GetHoldTimestampMsecs

`func (o *FileLockStatus) GetHoldTimestampMsecs() int64`

GetHoldTimestampMsecs returns the HoldTimestampMsecs field if non-nil, zero value otherwise.

### GetHoldTimestampMsecsOk

`func (o *FileLockStatus) GetHoldTimestampMsecsOk() (*int64, bool)`

GetHoldTimestampMsecsOk returns a tuple with the HoldTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimestampMsecs

`func (o *FileLockStatus) SetHoldTimestampMsecs(v int64)`

SetHoldTimestampMsecs sets HoldTimestampMsecs field to given value.

### HasHoldTimestampMsecs

`func (o *FileLockStatus) HasHoldTimestampMsecs() bool`

HasHoldTimestampMsecs returns a boolean if a field has been set.

### SetHoldTimestampMsecsNil

`func (o *FileLockStatus) SetHoldTimestampMsecsNil(b bool)`

 SetHoldTimestampMsecsNil sets the value for HoldTimestampMsecs to be an explicit nil

### UnsetHoldTimestampMsecs
`func (o *FileLockStatus) UnsetHoldTimestampMsecs()`

UnsetHoldTimestampMsecs ensures that no value is present for HoldTimestampMsecs, not even an explicit nil
### GetLockTimestampMsecs

`func (o *FileLockStatus) GetLockTimestampMsecs() int64`

GetLockTimestampMsecs returns the LockTimestampMsecs field if non-nil, zero value otherwise.

### GetLockTimestampMsecsOk

`func (o *FileLockStatus) GetLockTimestampMsecsOk() (*int64, bool)`

GetLockTimestampMsecsOk returns a tuple with the LockTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTimestampMsecs

`func (o *FileLockStatus) SetLockTimestampMsecs(v int64)`

SetLockTimestampMsecs sets LockTimestampMsecs field to given value.

### HasLockTimestampMsecs

`func (o *FileLockStatus) HasLockTimestampMsecs() bool`

HasLockTimestampMsecs returns a boolean if a field has been set.

### SetLockTimestampMsecsNil

`func (o *FileLockStatus) SetLockTimestampMsecsNil(b bool)`

 SetLockTimestampMsecsNil sets the value for LockTimestampMsecs to be an explicit nil

### UnsetLockTimestampMsecs
`func (o *FileLockStatus) UnsetLockTimestampMsecs()`

UnsetLockTimestampMsecs ensures that no value is present for LockTimestampMsecs, not even an explicit nil
### GetMode

`func (o *FileLockStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileLockStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileLockStatus) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileLockStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *FileLockStatus) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *FileLockStatus) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetState

`func (o *FileLockStatus) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FileLockStatus) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FileLockStatus) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *FileLockStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *FileLockStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *FileLockStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


