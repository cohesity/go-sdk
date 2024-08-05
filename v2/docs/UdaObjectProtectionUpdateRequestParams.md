# UdaObjectProtectionUpdateRequestParams

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

### NewUdaObjectProtectionUpdateRequestParams

`func NewUdaObjectProtectionUpdateRequestParams(objects []UdaObjectProtectionObjectParams, ) *UdaObjectProtectionUpdateRequestParams`

NewUdaObjectProtectionUpdateRequestParams instantiates a new UdaObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaObjectProtectionUpdateRequestParamsWithDefaults

`func NewUdaObjectProtectionUpdateRequestParamsWithDefaults() *UdaObjectProtectionUpdateRequestParams`

NewUdaObjectProtectionUpdateRequestParamsWithDefaults instantiates a new UdaObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupJobArguments

`func (o *UdaObjectProtectionUpdateRequestParams) GetBackupJobArguments() []KeyValuePair`

GetBackupJobArguments returns the BackupJobArguments field if non-nil, zero value otherwise.

### GetBackupJobArgumentsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetBackupJobArgumentsOk() (*[]KeyValuePair, bool)`

GetBackupJobArgumentsOk returns a tuple with the BackupJobArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobArguments

`func (o *UdaObjectProtectionUpdateRequestParams) SetBackupJobArguments(v []KeyValuePair)`

SetBackupJobArguments sets BackupJobArguments field to given value.

### HasBackupJobArguments

`func (o *UdaObjectProtectionUpdateRequestParams) HasBackupJobArguments() bool`

HasBackupJobArguments returns a boolean if a field has been set.

### SetBackupJobArgumentsNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetBackupJobArgumentsNil(b bool)`

 SetBackupJobArgumentsNil sets the value for BackupJobArguments to be an explicit nil

### UnsetBackupJobArguments
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetBackupJobArguments()`

UnsetBackupJobArguments ensures that no value is present for BackupJobArguments, not even an explicit nil
### GetConcurrency

`func (o *UdaObjectProtectionUpdateRequestParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UdaObjectProtectionUpdateRequestParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *UdaObjectProtectionUpdateRequestParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetFullBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) GetFullBackupArgs() string`

GetFullBackupArgs returns the FullBackupArgs field if non-nil, zero value otherwise.

### GetFullBackupArgsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetFullBackupArgsOk() (*string, bool)`

GetFullBackupArgsOk returns a tuple with the FullBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) SetFullBackupArgs(v string)`

SetFullBackupArgs sets FullBackupArgs field to given value.

### HasFullBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) HasFullBackupArgs() bool`

HasFullBackupArgs returns a boolean if a field has been set.

### SetFullBackupArgsNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetFullBackupArgsNil(b bool)`

 SetFullBackupArgsNil sets the value for FullBackupArgs to be an explicit nil

### UnsetFullBackupArgs
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetFullBackupArgs()`

UnsetFullBackupArgs ensures that no value is present for FullBackupArgs, not even an explicit nil
### GetHasEntitySupport

`func (o *UdaObjectProtectionUpdateRequestParams) GetHasEntitySupport() bool`

GetHasEntitySupport returns the HasEntitySupport field if non-nil, zero value otherwise.

### GetHasEntitySupportOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetHasEntitySupportOk() (*bool, bool)`

GetHasEntitySupportOk returns a tuple with the HasEntitySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitySupport

`func (o *UdaObjectProtectionUpdateRequestParams) SetHasEntitySupport(v bool)`

SetHasEntitySupport sets HasEntitySupport field to given value.

### HasHasEntitySupport

`func (o *UdaObjectProtectionUpdateRequestParams) HasHasEntitySupport() bool`

HasHasEntitySupport returns a boolean if a field has been set.

### SetHasEntitySupportNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetHasEntitySupportNil(b bool)`

 SetHasEntitySupportNil sets the value for HasEntitySupport to be an explicit nil

### UnsetHasEntitySupport
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetHasEntitySupport()`

UnsetHasEntitySupport ensures that no value is present for HasEntitySupport, not even an explicit nil
### GetIncrBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) GetIncrBackupArgs() string`

GetIncrBackupArgs returns the IncrBackupArgs field if non-nil, zero value otherwise.

### GetIncrBackupArgsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetIncrBackupArgsOk() (*string, bool)`

GetIncrBackupArgsOk returns a tuple with the IncrBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) SetIncrBackupArgs(v string)`

SetIncrBackupArgs sets IncrBackupArgs field to given value.

### HasIncrBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) HasIncrBackupArgs() bool`

HasIncrBackupArgs returns a boolean if a field has been set.

### SetIncrBackupArgsNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetIncrBackupArgsNil(b bool)`

 SetIncrBackupArgsNil sets the value for IncrBackupArgs to be an explicit nil

### UnsetIncrBackupArgs
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetIncrBackupArgs()`

UnsetIncrBackupArgs ensures that no value is present for IncrBackupArgs, not even an explicit nil
### GetLogBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) GetLogBackupArgs() string`

GetLogBackupArgs returns the LogBackupArgs field if non-nil, zero value otherwise.

### GetLogBackupArgsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetLogBackupArgsOk() (*string, bool)`

GetLogBackupArgsOk returns a tuple with the LogBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) SetLogBackupArgs(v string)`

SetLogBackupArgs sets LogBackupArgs field to given value.

### HasLogBackupArgs

`func (o *UdaObjectProtectionUpdateRequestParams) HasLogBackupArgs() bool`

HasLogBackupArgs returns a boolean if a field has been set.

### SetLogBackupArgsNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetLogBackupArgsNil(b bool)`

 SetLogBackupArgsNil sets the value for LogBackupArgs to be an explicit nil

### UnsetLogBackupArgs
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetLogBackupArgs()`

UnsetLogBackupArgs ensures that no value is present for LogBackupArgs, not even an explicit nil
### GetMounts

`func (o *UdaObjectProtectionUpdateRequestParams) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetMountsOk() (*int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *UdaObjectProtectionUpdateRequestParams) SetMounts(v int32)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *UdaObjectProtectionUpdateRequestParams) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *UdaObjectProtectionUpdateRequestParams) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *UdaObjectProtectionUpdateRequestParams) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetObjects

`func (o *UdaObjectProtectionUpdateRequestParams) GetObjects() []UdaObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *UdaObjectProtectionUpdateRequestParams) GetObjectsOk() (*[]UdaObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *UdaObjectProtectionUpdateRequestParams) SetObjects(v []UdaObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


