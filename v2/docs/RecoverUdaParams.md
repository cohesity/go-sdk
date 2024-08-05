# RecoverUdaParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 1]
**Mounts** | Pointer to **NullableInt32** | Specifies the maximum number of view mounts per host. If not specified, the default value is taken as 1. | [optional] [default to 1]
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**RecoveryArgs** | Pointer to **NullableString** | Specifies the custom arguments to be supplied to the restore job script. This field is deprecated. Use recoveryJobArguments instead. | [optional] 
**RecoveryJobArguments** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the map of custom arguments to be supplied to the restore job script. | [optional] 
**Snapshots** | [**[]RecoverUdaSnapshotParams**](RecoverUdaSnapshotParams.md) | Specifies the local snapshot ids and other details of the objects to be recovered. | 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 

## Methods

### NewRecoverUdaParams

`func NewRecoverUdaParams(snapshots []RecoverUdaSnapshotParams, ) *RecoverUdaParams`

NewRecoverUdaParams instantiates a new RecoverUdaParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverUdaParamsWithDefaults

`func NewRecoverUdaParamsWithDefaults() *RecoverUdaParams`

NewRecoverUdaParamsWithDefaults instantiates a new RecoverUdaParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *RecoverUdaParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverUdaParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverUdaParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverUdaParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverUdaParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverUdaParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetMounts

`func (o *RecoverUdaParams) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *RecoverUdaParams) GetMountsOk() (*int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *RecoverUdaParams) SetMounts(v int32)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *RecoverUdaParams) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *RecoverUdaParams) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *RecoverUdaParams) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetRecoverTo

`func (o *RecoverUdaParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverUdaParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverUdaParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverUdaParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverUdaParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverUdaParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetRecoveryArgs

`func (o *RecoverUdaParams) GetRecoveryArgs() string`

GetRecoveryArgs returns the RecoveryArgs field if non-nil, zero value otherwise.

### GetRecoveryArgsOk

`func (o *RecoverUdaParams) GetRecoveryArgsOk() (*string, bool)`

GetRecoveryArgsOk returns a tuple with the RecoveryArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryArgs

`func (o *RecoverUdaParams) SetRecoveryArgs(v string)`

SetRecoveryArgs sets RecoveryArgs field to given value.

### HasRecoveryArgs

`func (o *RecoverUdaParams) HasRecoveryArgs() bool`

HasRecoveryArgs returns a boolean if a field has been set.

### SetRecoveryArgsNil

`func (o *RecoverUdaParams) SetRecoveryArgsNil(b bool)`

 SetRecoveryArgsNil sets the value for RecoveryArgs to be an explicit nil

### UnsetRecoveryArgs
`func (o *RecoverUdaParams) UnsetRecoveryArgs()`

UnsetRecoveryArgs ensures that no value is present for RecoveryArgs, not even an explicit nil
### GetRecoveryJobArguments

`func (o *RecoverUdaParams) GetRecoveryJobArguments() []KeyValuePair`

GetRecoveryJobArguments returns the RecoveryJobArguments field if non-nil, zero value otherwise.

### GetRecoveryJobArgumentsOk

`func (o *RecoverUdaParams) GetRecoveryJobArgumentsOk() (*[]KeyValuePair, bool)`

GetRecoveryJobArgumentsOk returns a tuple with the RecoveryJobArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryJobArguments

`func (o *RecoverUdaParams) SetRecoveryJobArguments(v []KeyValuePair)`

SetRecoveryJobArguments sets RecoveryJobArguments field to given value.

### HasRecoveryJobArguments

`func (o *RecoverUdaParams) HasRecoveryJobArguments() bool`

HasRecoveryJobArguments returns a boolean if a field has been set.

### SetRecoveryJobArgumentsNil

`func (o *RecoverUdaParams) SetRecoveryJobArgumentsNil(b bool)`

 SetRecoveryJobArgumentsNil sets the value for RecoveryJobArguments to be an explicit nil

### UnsetRecoveryJobArguments
`func (o *RecoverUdaParams) UnsetRecoveryJobArguments()`

UnsetRecoveryJobArguments ensures that no value is present for RecoveryJobArguments, not even an explicit nil
### GetSnapshots

`func (o *RecoverUdaParams) GetSnapshots() []RecoverUdaSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverUdaParams) GetSnapshotsOk() (*[]RecoverUdaSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverUdaParams) SetSnapshots(v []RecoverUdaSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverUdaParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverUdaParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetWarnings

`func (o *RecoverUdaParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverUdaParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverUdaParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverUdaParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverUdaParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverUdaParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


