# RecoverSapHanaParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 1]
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Snapshots** | [**[]RecoverUdaSnapshotParams**](RecoverUdaSnapshotParams.md) | Specifies the local snapshot ids and other details of the objects to be recovered. | 
**StartDatabase** | Pointer to **NullableBool** | Start the database after the recovery is complete. | [optional] [default to true]
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 

## Methods

### NewRecoverSapHanaParams

`func NewRecoverSapHanaParams(snapshots []RecoverUdaSnapshotParams, ) *RecoverSapHanaParams`

NewRecoverSapHanaParams instantiates a new RecoverSapHanaParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSapHanaParamsWithDefaults

`func NewRecoverSapHanaParamsWithDefaults() *RecoverSapHanaParams`

NewRecoverSapHanaParamsWithDefaults instantiates a new RecoverSapHanaParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *RecoverSapHanaParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverSapHanaParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverSapHanaParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverSapHanaParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverSapHanaParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverSapHanaParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetRecoverTo

`func (o *RecoverSapHanaParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverSapHanaParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverSapHanaParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverSapHanaParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverSapHanaParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverSapHanaParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetSnapshots

`func (o *RecoverSapHanaParams) GetSnapshots() []RecoverUdaSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverSapHanaParams) GetSnapshotsOk() (*[]RecoverUdaSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverSapHanaParams) SetSnapshots(v []RecoverUdaSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverSapHanaParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverSapHanaParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetStartDatabase

`func (o *RecoverSapHanaParams) GetStartDatabase() bool`

GetStartDatabase returns the StartDatabase field if non-nil, zero value otherwise.

### GetStartDatabaseOk

`func (o *RecoverSapHanaParams) GetStartDatabaseOk() (*bool, bool)`

GetStartDatabaseOk returns a tuple with the StartDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatabase

`func (o *RecoverSapHanaParams) SetStartDatabase(v bool)`

SetStartDatabase sets StartDatabase field to given value.

### HasStartDatabase

`func (o *RecoverSapHanaParams) HasStartDatabase() bool`

HasStartDatabase returns a boolean if a field has been set.

### SetStartDatabaseNil

`func (o *RecoverSapHanaParams) SetStartDatabaseNil(b bool)`

 SetStartDatabaseNil sets the value for StartDatabase to be an explicit nil

### UnsetStartDatabase
`func (o *RecoverSapHanaParams) UnsetStartDatabase()`

UnsetStartDatabase ensures that no value is present for StartDatabase, not even an explicit nil
### GetWarnings

`func (o *RecoverSapHanaParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverSapHanaParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverSapHanaParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverSapHanaParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverSapHanaParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverSapHanaParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


