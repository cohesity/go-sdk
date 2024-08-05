# UdaObjectProtectionRequestParams

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

### NewUdaObjectProtectionRequestParams

`func NewUdaObjectProtectionRequestParams(objects []UdaObjectProtectionObjectParams, ) *UdaObjectProtectionRequestParams`

NewUdaObjectProtectionRequestParams instantiates a new UdaObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaObjectProtectionRequestParamsWithDefaults

`func NewUdaObjectProtectionRequestParamsWithDefaults() *UdaObjectProtectionRequestParams`

NewUdaObjectProtectionRequestParamsWithDefaults instantiates a new UdaObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupJobArguments

`func (o *UdaObjectProtectionRequestParams) GetBackupJobArguments() []KeyValuePair`

GetBackupJobArguments returns the BackupJobArguments field if non-nil, zero value otherwise.

### GetBackupJobArgumentsOk

`func (o *UdaObjectProtectionRequestParams) GetBackupJobArgumentsOk() (*[]KeyValuePair, bool)`

GetBackupJobArgumentsOk returns a tuple with the BackupJobArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobArguments

`func (o *UdaObjectProtectionRequestParams) SetBackupJobArguments(v []KeyValuePair)`

SetBackupJobArguments sets BackupJobArguments field to given value.

### HasBackupJobArguments

`func (o *UdaObjectProtectionRequestParams) HasBackupJobArguments() bool`

HasBackupJobArguments returns a boolean if a field has been set.

### SetBackupJobArgumentsNil

`func (o *UdaObjectProtectionRequestParams) SetBackupJobArgumentsNil(b bool)`

 SetBackupJobArgumentsNil sets the value for BackupJobArguments to be an explicit nil

### UnsetBackupJobArguments
`func (o *UdaObjectProtectionRequestParams) UnsetBackupJobArguments()`

UnsetBackupJobArguments ensures that no value is present for BackupJobArguments, not even an explicit nil
### GetConcurrency

`func (o *UdaObjectProtectionRequestParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UdaObjectProtectionRequestParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UdaObjectProtectionRequestParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *UdaObjectProtectionRequestParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *UdaObjectProtectionRequestParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *UdaObjectProtectionRequestParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetFullBackupArgs

`func (o *UdaObjectProtectionRequestParams) GetFullBackupArgs() string`

GetFullBackupArgs returns the FullBackupArgs field if non-nil, zero value otherwise.

### GetFullBackupArgsOk

`func (o *UdaObjectProtectionRequestParams) GetFullBackupArgsOk() (*string, bool)`

GetFullBackupArgsOk returns a tuple with the FullBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupArgs

`func (o *UdaObjectProtectionRequestParams) SetFullBackupArgs(v string)`

SetFullBackupArgs sets FullBackupArgs field to given value.

### HasFullBackupArgs

`func (o *UdaObjectProtectionRequestParams) HasFullBackupArgs() bool`

HasFullBackupArgs returns a boolean if a field has been set.

### SetFullBackupArgsNil

`func (o *UdaObjectProtectionRequestParams) SetFullBackupArgsNil(b bool)`

 SetFullBackupArgsNil sets the value for FullBackupArgs to be an explicit nil

### UnsetFullBackupArgs
`func (o *UdaObjectProtectionRequestParams) UnsetFullBackupArgs()`

UnsetFullBackupArgs ensures that no value is present for FullBackupArgs, not even an explicit nil
### GetHasEntitySupport

`func (o *UdaObjectProtectionRequestParams) GetHasEntitySupport() bool`

GetHasEntitySupport returns the HasEntitySupport field if non-nil, zero value otherwise.

### GetHasEntitySupportOk

`func (o *UdaObjectProtectionRequestParams) GetHasEntitySupportOk() (*bool, bool)`

GetHasEntitySupportOk returns a tuple with the HasEntitySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitySupport

`func (o *UdaObjectProtectionRequestParams) SetHasEntitySupport(v bool)`

SetHasEntitySupport sets HasEntitySupport field to given value.

### HasHasEntitySupport

`func (o *UdaObjectProtectionRequestParams) HasHasEntitySupport() bool`

HasHasEntitySupport returns a boolean if a field has been set.

### SetHasEntitySupportNil

`func (o *UdaObjectProtectionRequestParams) SetHasEntitySupportNil(b bool)`

 SetHasEntitySupportNil sets the value for HasEntitySupport to be an explicit nil

### UnsetHasEntitySupport
`func (o *UdaObjectProtectionRequestParams) UnsetHasEntitySupport()`

UnsetHasEntitySupport ensures that no value is present for HasEntitySupport, not even an explicit nil
### GetIncrBackupArgs

`func (o *UdaObjectProtectionRequestParams) GetIncrBackupArgs() string`

GetIncrBackupArgs returns the IncrBackupArgs field if non-nil, zero value otherwise.

### GetIncrBackupArgsOk

`func (o *UdaObjectProtectionRequestParams) GetIncrBackupArgsOk() (*string, bool)`

GetIncrBackupArgsOk returns a tuple with the IncrBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrBackupArgs

`func (o *UdaObjectProtectionRequestParams) SetIncrBackupArgs(v string)`

SetIncrBackupArgs sets IncrBackupArgs field to given value.

### HasIncrBackupArgs

`func (o *UdaObjectProtectionRequestParams) HasIncrBackupArgs() bool`

HasIncrBackupArgs returns a boolean if a field has been set.

### SetIncrBackupArgsNil

`func (o *UdaObjectProtectionRequestParams) SetIncrBackupArgsNil(b bool)`

 SetIncrBackupArgsNil sets the value for IncrBackupArgs to be an explicit nil

### UnsetIncrBackupArgs
`func (o *UdaObjectProtectionRequestParams) UnsetIncrBackupArgs()`

UnsetIncrBackupArgs ensures that no value is present for IncrBackupArgs, not even an explicit nil
### GetLogBackupArgs

`func (o *UdaObjectProtectionRequestParams) GetLogBackupArgs() string`

GetLogBackupArgs returns the LogBackupArgs field if non-nil, zero value otherwise.

### GetLogBackupArgsOk

`func (o *UdaObjectProtectionRequestParams) GetLogBackupArgsOk() (*string, bool)`

GetLogBackupArgsOk returns a tuple with the LogBackupArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupArgs

`func (o *UdaObjectProtectionRequestParams) SetLogBackupArgs(v string)`

SetLogBackupArgs sets LogBackupArgs field to given value.

### HasLogBackupArgs

`func (o *UdaObjectProtectionRequestParams) HasLogBackupArgs() bool`

HasLogBackupArgs returns a boolean if a field has been set.

### SetLogBackupArgsNil

`func (o *UdaObjectProtectionRequestParams) SetLogBackupArgsNil(b bool)`

 SetLogBackupArgsNil sets the value for LogBackupArgs to be an explicit nil

### UnsetLogBackupArgs
`func (o *UdaObjectProtectionRequestParams) UnsetLogBackupArgs()`

UnsetLogBackupArgs ensures that no value is present for LogBackupArgs, not even an explicit nil
### GetMounts

`func (o *UdaObjectProtectionRequestParams) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *UdaObjectProtectionRequestParams) GetMountsOk() (*int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *UdaObjectProtectionRequestParams) SetMounts(v int32)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *UdaObjectProtectionRequestParams) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *UdaObjectProtectionRequestParams) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *UdaObjectProtectionRequestParams) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetObjects

`func (o *UdaObjectProtectionRequestParams) GetObjects() []UdaObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *UdaObjectProtectionRequestParams) GetObjectsOk() (*[]UdaObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *UdaObjectProtectionRequestParams) SetObjects(v []UdaObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


