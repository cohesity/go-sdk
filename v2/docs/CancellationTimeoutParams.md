# CancellationTimeoutParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | Pointer to **NullableString** | The scheduled backup type(kFull, kRegular etc.) | [optional] 
**TimeoutMins** | Pointer to **NullableInt64** | Specifies the timeout in mins | [optional] 

## Methods

### NewCancellationTimeoutParams

`func NewCancellationTimeoutParams() *CancellationTimeoutParams`

NewCancellationTimeoutParams instantiates a new CancellationTimeoutParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationTimeoutParamsWithDefaults

`func NewCancellationTimeoutParamsWithDefaults() *CancellationTimeoutParams`

NewCancellationTimeoutParamsWithDefaults instantiates a new CancellationTimeoutParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupType

`func (o *CancellationTimeoutParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *CancellationTimeoutParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *CancellationTimeoutParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *CancellationTimeoutParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### SetBackupTypeNil

`func (o *CancellationTimeoutParams) SetBackupTypeNil(b bool)`

 SetBackupTypeNil sets the value for BackupType to be an explicit nil

### UnsetBackupType
`func (o *CancellationTimeoutParams) UnsetBackupType()`

UnsetBackupType ensures that no value is present for BackupType, not even an explicit nil
### GetTimeoutMins

`func (o *CancellationTimeoutParams) GetTimeoutMins() int64`

GetTimeoutMins returns the TimeoutMins field if non-nil, zero value otherwise.

### GetTimeoutMinsOk

`func (o *CancellationTimeoutParams) GetTimeoutMinsOk() (*int64, bool)`

GetTimeoutMinsOk returns a tuple with the TimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMins

`func (o *CancellationTimeoutParams) SetTimeoutMins(v int64)`

SetTimeoutMins sets TimeoutMins field to given value.

### HasTimeoutMins

`func (o *CancellationTimeoutParams) HasTimeoutMins() bool`

HasTimeoutMins returns a boolean if a field has been set.

### SetTimeoutMinsNil

`func (o *CancellationTimeoutParams) SetTimeoutMinsNil(b bool)`

 SetTimeoutMinsNil sets the value for TimeoutMins to be an explicit nil

### UnsetTimeoutMins
`func (o *CancellationTimeoutParams) UnsetTimeoutMins()`

UnsetTimeoutMins ensures that no value is present for TimeoutMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


