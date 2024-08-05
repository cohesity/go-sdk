# UdaObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupJobArguments** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the map of custom arguments to be supplied to the various backup scripts. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams thatwill be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 1]
**FullBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the full backup script when a full backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead. | [optional] 
**HasEntitySupport** | Pointer to **NullableBool** | Specifies whether this Protection Group is created from a source having entity support. | [optional] [readonly] 
**IncrBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the incremental backup script when an incremental backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead. | [optional] 
**LogBackupArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the log backup script when a log backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead. | [optional] 
**Mounts** | Pointer to **NullableInt32** | Specifies the maximum number of view mounts per host. If not specified, the default value is taken as 1. | [optional] [default to 1]
**Objects** | [**[]UdaObjectProtectionObjectParams**](UdaObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 

## Methods

### NewUdaObjectProtectionParams

`func NewUdaObjectProtectionParams(objects []UdaObjectProtectionObjectParams, ) *UdaObjectProtectionParams`

NewUdaObjectProtectionParams instantiates a new UdaObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaObjectProtectionParamsWithDefaults

`func NewUdaObjectProtectionParamsWithDefaults() *UdaObjectProtectionParams`

NewUdaObjectProtectionParamsWithDefaults instantiates a new UdaObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupJobArguments

`func (o *UdaObjectProtectionParams) GetBackupJobArguments() []KeyValuePair`

GetBackupJobArguments returns the BackupJobArguments field if non-nil, zero value otherwise.

### GetBackupJobArgumentsOk

`func (o *UdaObjectProtectionParams) GetBackupJobArgumentsOk() (*[]KeyValuePair, bool)`

GetBackupJobArgumentsOk returns a tuple with the BackupJobArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobArguments

`func (o *UdaObjectProtectionParams) SetBackupJobArguments(v []KeyValuePair)`

SetBackupJobArguments sets BackupJobArguments field to given value.

### HasBackupJobArguments

`func (o *UdaObjectProtectionParams) HasBackupJobArguments() bool`

HasBackupJobArguments returns a boolean if a field has been set.

### SetBackupJobArgumentsNil

`func (o *UdaObjectProtectionParams) SetBackupJobArgumentsNil(b bool)`

 SetBackupJobArgumentsNil sets the value for BackupJobArguments to be an explicit nil

### UnsetBackupJobArguments
`func (o *UdaObjectProtectionParams) UnsetBackupJobArguments()`

UnsetBackupJobArguments ensures that no value is present for BackupJobArguments, not even an explicit nil
### GetConcurrency

`func (o *UdaObjectProtectionParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UdaObjectProtectionParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UdaObjectProtectionParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *UdaObjectProtectionParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *UdaObjectProtectionParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *UdaObjectProtectionParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetFullBackupArgs

`func (o *UdaObjectProtectionParams) GetFullBackupArgs() string`

GetFullBackupArgs returns the FullBackupArgs field if non-nil, zero value otherwise.

### GetFullBackupArgsOk

`func (o *UdaObjectProtectionParams) GetFullBackupArgsOk() (*string, bool)`

GetFullBackupArgsOk returns a tuple with the FullBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupArgs

`func (o *UdaObjectProtectionParams) SetFullBackupArgs(v string)`

SetFullBackupArgs sets FullBackupArgs field to given value.

### HasFullBackupArgs

`func (o *UdaObjectProtectionParams) HasFullBackupArgs() bool`

HasFullBackupArgs returns a boolean if a field has been set.

### SetFullBackupArgsNil

`func (o *UdaObjectProtectionParams) SetFullBackupArgsNil(b bool)`

 SetFullBackupArgsNil sets the value for FullBackupArgs to be an explicit nil

### UnsetFullBackupArgs
`func (o *UdaObjectProtectionParams) UnsetFullBackupArgs()`

UnsetFullBackupArgs ensures that no value is present for FullBackupArgs, not even an explicit nil
### GetHasEntitySupport

`func (o *UdaObjectProtectionParams) GetHasEntitySupport() bool`

GetHasEntitySupport returns the HasEntitySupport field if non-nil, zero value otherwise.

### GetHasEntitySupportOk

`func (o *UdaObjectProtectionParams) GetHasEntitySupportOk() (*bool, bool)`

GetHasEntitySupportOk returns a tuple with the HasEntitySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitySupport

`func (o *UdaObjectProtectionParams) SetHasEntitySupport(v bool)`

SetHasEntitySupport sets HasEntitySupport field to given value.

### HasHasEntitySupport

`func (o *UdaObjectProtectionParams) HasHasEntitySupport() bool`

HasHasEntitySupport returns a boolean if a field has been set.

### SetHasEntitySupportNil

`func (o *UdaObjectProtectionParams) SetHasEntitySupportNil(b bool)`

 SetHasEntitySupportNil sets the value for HasEntitySupport to be an explicit nil

### UnsetHasEntitySupport
`func (o *UdaObjectProtectionParams) UnsetHasEntitySupport()`

UnsetHasEntitySupport ensures that no value is present for HasEntitySupport, not even an explicit nil
### GetIncrBackupArgs

`func (o *UdaObjectProtectionParams) GetIncrBackupArgs() string`

GetIncrBackupArgs returns the IncrBackupArgs field if non-nil, zero value otherwise.

### GetIncrBackupArgsOk

`func (o *UdaObjectProtectionParams) GetIncrBackupArgsOk() (*string, bool)`

GetIncrBackupArgsOk returns a tuple with the IncrBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrBackupArgs

`func (o *UdaObjectProtectionParams) SetIncrBackupArgs(v string)`

SetIncrBackupArgs sets IncrBackupArgs field to given value.

### HasIncrBackupArgs

`func (o *UdaObjectProtectionParams) HasIncrBackupArgs() bool`

HasIncrBackupArgs returns a boolean if a field has been set.

### SetIncrBackupArgsNil

`func (o *UdaObjectProtectionParams) SetIncrBackupArgsNil(b bool)`

 SetIncrBackupArgsNil sets the value for IncrBackupArgs to be an explicit nil

### UnsetIncrBackupArgs
`func (o *UdaObjectProtectionParams) UnsetIncrBackupArgs()`

UnsetIncrBackupArgs ensures that no value is present for IncrBackupArgs, not even an explicit nil
### GetLogBackupArgs

`func (o *UdaObjectProtectionParams) GetLogBackupArgs() string`

GetLogBackupArgs returns the LogBackupArgs field if non-nil, zero value otherwise.

### GetLogBackupArgsOk

`func (o *UdaObjectProtectionParams) GetLogBackupArgsOk() (*string, bool)`

GetLogBackupArgsOk returns a tuple with the LogBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupArgs

`func (o *UdaObjectProtectionParams) SetLogBackupArgs(v string)`

SetLogBackupArgs sets LogBackupArgs field to given value.

### HasLogBackupArgs

`func (o *UdaObjectProtectionParams) HasLogBackupArgs() bool`

HasLogBackupArgs returns a boolean if a field has been set.

### SetLogBackupArgsNil

`func (o *UdaObjectProtectionParams) SetLogBackupArgsNil(b bool)`

 SetLogBackupArgsNil sets the value for LogBackupArgs to be an explicit nil

### UnsetLogBackupArgs
`func (o *UdaObjectProtectionParams) UnsetLogBackupArgs()`

UnsetLogBackupArgs ensures that no value is present for LogBackupArgs, not even an explicit nil
### GetMounts

`func (o *UdaObjectProtectionParams) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *UdaObjectProtectionParams) GetMountsOk() (*int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *UdaObjectProtectionParams) SetMounts(v int32)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *UdaObjectProtectionParams) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *UdaObjectProtectionParams) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *UdaObjectProtectionParams) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetObjects

`func (o *UdaObjectProtectionParams) GetObjects() []UdaObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *UdaObjectProtectionParams) GetObjectsOk() (*[]UdaObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *UdaObjectProtectionParams) SetObjects(v []UdaObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


