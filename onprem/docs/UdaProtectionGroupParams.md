# UdaProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **NullableInt64** | Specifies the source Id of the objects to be protected. | 
**Objects** | [**[]UdaProtectionGroupObjectParams**](UdaProtectionGroupObjectParams.md) | Specifies a list of fully qualified names of the objects to be protected. | 
**FullBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the full backup script when a full backup is enabled in the policy. | [optional] 
**IncrBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the incremental backup script when an incremental backup is enabled in the policy. | [optional] 
**LogBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the log backup script when a log backup is enabled in the policy. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 1]
**Mounts** | Pointer to **NullableInt32** | Specifies the maximum number of view mounts per host. If not specified, the default value is taken as 1. | [optional] [default to 1]

## Methods

### NewUdaProtectionGroupParams

`func NewUdaProtectionGroupParams(sourceId NullableInt64, objects []UdaProtectionGroupObjectParams, ) *UdaProtectionGroupParams`

NewUdaProtectionGroupParams instantiates a new UdaProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaProtectionGroupParamsWithDefaults

`func NewUdaProtectionGroupParamsWithDefaults() *UdaProtectionGroupParams`

NewUdaProtectionGroupParamsWithDefaults instantiates a new UdaProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *UdaProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UdaProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UdaProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *UdaProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *UdaProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetObjects

`func (o *UdaProtectionGroupParams) GetObjects() []UdaProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *UdaProtectionGroupParams) GetObjectsOk() (*[]UdaProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *UdaProtectionGroupParams) SetObjects(v []UdaProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetFullBackupArgs

`func (o *UdaProtectionGroupParams) GetFullBackupArgs() string`

GetFullBackupArgs returns the FullBackupArgs field if non-nil, zero value otherwise.

### GetFullBackupArgsOk

`func (o *UdaProtectionGroupParams) GetFullBackupArgsOk() (*string, bool)`

GetFullBackupArgsOk returns a tuple with the FullBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupArgs

`func (o *UdaProtectionGroupParams) SetFullBackupArgs(v string)`

SetFullBackupArgs sets FullBackupArgs field to given value.

### HasFullBackupArgs

`func (o *UdaProtectionGroupParams) HasFullBackupArgs() bool`

HasFullBackupArgs returns a boolean if a field has been set.

### SetFullBackupArgsNil

`func (o *UdaProtectionGroupParams) SetFullBackupArgsNil(b bool)`

 SetFullBackupArgsNil sets the value for FullBackupArgs to be an explicit nil

### UnsetFullBackupArgs
`func (o *UdaProtectionGroupParams) UnsetFullBackupArgs()`

UnsetFullBackupArgs ensures that no value is present for FullBackupArgs, not even an explicit nil
### GetIncrBackupArgs

`func (o *UdaProtectionGroupParams) GetIncrBackupArgs() string`

GetIncrBackupArgs returns the IncrBackupArgs field if non-nil, zero value otherwise.

### GetIncrBackupArgsOk

`func (o *UdaProtectionGroupParams) GetIncrBackupArgsOk() (*string, bool)`

GetIncrBackupArgsOk returns a tuple with the IncrBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrBackupArgs

`func (o *UdaProtectionGroupParams) SetIncrBackupArgs(v string)`

SetIncrBackupArgs sets IncrBackupArgs field to given value.

### HasIncrBackupArgs

`func (o *UdaProtectionGroupParams) HasIncrBackupArgs() bool`

HasIncrBackupArgs returns a boolean if a field has been set.

### SetIncrBackupArgsNil

`func (o *UdaProtectionGroupParams) SetIncrBackupArgsNil(b bool)`

 SetIncrBackupArgsNil sets the value for IncrBackupArgs to be an explicit nil

### UnsetIncrBackupArgs
`func (o *UdaProtectionGroupParams) UnsetIncrBackupArgs()`

UnsetIncrBackupArgs ensures that no value is present for IncrBackupArgs, not even an explicit nil
### GetLogBackupArgs

`func (o *UdaProtectionGroupParams) GetLogBackupArgs() string`

GetLogBackupArgs returns the LogBackupArgs field if non-nil, zero value otherwise.

### GetLogBackupArgsOk

`func (o *UdaProtectionGroupParams) GetLogBackupArgsOk() (*string, bool)`

GetLogBackupArgsOk returns a tuple with the LogBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupArgs

`func (o *UdaProtectionGroupParams) SetLogBackupArgs(v string)`

SetLogBackupArgs sets LogBackupArgs field to given value.

### HasLogBackupArgs

`func (o *UdaProtectionGroupParams) HasLogBackupArgs() bool`

HasLogBackupArgs returns a boolean if a field has been set.

### SetLogBackupArgsNil

`func (o *UdaProtectionGroupParams) SetLogBackupArgsNil(b bool)`

 SetLogBackupArgsNil sets the value for LogBackupArgs to be an explicit nil

### UnsetLogBackupArgs
`func (o *UdaProtectionGroupParams) UnsetLogBackupArgs()`

UnsetLogBackupArgs ensures that no value is present for LogBackupArgs, not even an explicit nil
### GetConcurrency

`func (o *UdaProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UdaProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UdaProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *UdaProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *UdaProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *UdaProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetMounts

`func (o *UdaProtectionGroupParams) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *UdaProtectionGroupParams) GetMountsOk() (*int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *UdaProtectionGroupParams) SetMounts(v int32)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *UdaProtectionGroupParams) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *UdaProtectionGroupParams) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *UdaProtectionGroupParams) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


