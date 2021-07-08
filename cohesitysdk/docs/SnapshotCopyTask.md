# SnapshotCopyTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyStatus** | Pointer to **NullableString** | Specifies the status of the copy task. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies when the Snapshot expires on the target. | [optional] 
**Message** | Pointer to **NullableString** | Specifies warning or error information when the copy task is not successful. | [optional] 
**SnapshotTarget** | Pointer to [**SnapshotTargetSettings**](SnapshotTargetSettings.md) |  | [optional] 

## Methods

### NewSnapshotCopyTask

`func NewSnapshotCopyTask() *SnapshotCopyTask`

NewSnapshotCopyTask instantiates a new SnapshotCopyTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCopyTaskWithDefaults

`func NewSnapshotCopyTaskWithDefaults() *SnapshotCopyTask`

NewSnapshotCopyTaskWithDefaults instantiates a new SnapshotCopyTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyStatus

`func (o *SnapshotCopyTask) GetCopyStatus() string`

GetCopyStatus returns the CopyStatus field if non-nil, zero value otherwise.

### GetCopyStatusOk

`func (o *SnapshotCopyTask) GetCopyStatusOk() (*string, bool)`

GetCopyStatusOk returns a tuple with the CopyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyStatus

`func (o *SnapshotCopyTask) SetCopyStatus(v string)`

SetCopyStatus sets CopyStatus field to given value.

### HasCopyStatus

`func (o *SnapshotCopyTask) HasCopyStatus() bool`

HasCopyStatus returns a boolean if a field has been set.

### SetCopyStatusNil

`func (o *SnapshotCopyTask) SetCopyStatusNil(b bool)`

 SetCopyStatusNil sets the value for CopyStatus to be an explicit nil

### UnsetCopyStatus
`func (o *SnapshotCopyTask) UnsetCopyStatus()`

UnsetCopyStatus ensures that no value is present for CopyStatus, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *SnapshotCopyTask) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *SnapshotCopyTask) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *SnapshotCopyTask) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *SnapshotCopyTask) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *SnapshotCopyTask) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *SnapshotCopyTask) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetMessage

`func (o *SnapshotCopyTask) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotCopyTask) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotCopyTask) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SnapshotCopyTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SnapshotCopyTask) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SnapshotCopyTask) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSnapshotTarget

`func (o *SnapshotCopyTask) GetSnapshotTarget() SnapshotTargetSettings`

GetSnapshotTarget returns the SnapshotTarget field if non-nil, zero value otherwise.

### GetSnapshotTargetOk

`func (o *SnapshotCopyTask) GetSnapshotTargetOk() (*SnapshotTargetSettings, bool)`

GetSnapshotTargetOk returns a tuple with the SnapshotTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTarget

`func (o *SnapshotCopyTask) SetSnapshotTarget(v SnapshotTargetSettings)`

SetSnapshotTarget sets SnapshotTarget field to given value.

### HasSnapshotTarget

`func (o *SnapshotCopyTask) HasSnapshotTarget() bool`

HasSnapshotTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


