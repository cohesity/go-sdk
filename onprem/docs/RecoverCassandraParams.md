# RecoverCassandraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 
**Snapshots** | [**[]RecoverCassandraSnapshotParams**](RecoverCassandraSnapshotParams.md) | Specifies the local snapshot ids and other details of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 
**SelectedDataCenters** | Pointer to **[]string** | Selected Data centers for this cluster. | [optional] 
**StagingDirectoryList** | Pointer to **[]string** | Specifies the directory on the primary to copy the files which are to be uploaded using destination sstableloader. | [optional] 

## Methods

### NewRecoverCassandraParams

`func NewRecoverCassandraParams(snapshots []RecoverCassandraSnapshotParams, ) *RecoverCassandraParams`

NewRecoverCassandraParams instantiates a new RecoverCassandraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCassandraParamsWithDefaults

`func NewRecoverCassandraParamsWithDefaults() *RecoverCassandraParams`

NewRecoverCassandraParamsWithDefaults instantiates a new RecoverCassandraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverTo

`func (o *RecoverCassandraParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverCassandraParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverCassandraParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverCassandraParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverCassandraParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverCassandraParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetOverwrite

`func (o *RecoverCassandraParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverCassandraParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverCassandraParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverCassandraParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverCassandraParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverCassandraParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetConcurrency

`func (o *RecoverCassandraParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverCassandraParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverCassandraParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverCassandraParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverCassandraParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverCassandraParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetBandwidthMBPS

`func (o *RecoverCassandraParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *RecoverCassandraParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *RecoverCassandraParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *RecoverCassandraParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *RecoverCassandraParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *RecoverCassandraParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetWarnings

`func (o *RecoverCassandraParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverCassandraParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverCassandraParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverCassandraParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverCassandraParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverCassandraParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetSnapshots

`func (o *RecoverCassandraParams) GetSnapshots() []RecoverCassandraSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverCassandraParams) GetSnapshotsOk() (*[]RecoverCassandraSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverCassandraParams) SetSnapshots(v []RecoverCassandraSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverCassandraParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverCassandraParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverCassandraParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverCassandraParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverCassandraParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverCassandraParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverCassandraParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverCassandraParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetSelectedDataCenters

`func (o *RecoverCassandraParams) GetSelectedDataCenters() []string`

GetSelectedDataCenters returns the SelectedDataCenters field if non-nil, zero value otherwise.

### GetSelectedDataCentersOk

`func (o *RecoverCassandraParams) GetSelectedDataCentersOk() (*[]string, bool)`

GetSelectedDataCentersOk returns a tuple with the SelectedDataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDataCenters

`func (o *RecoverCassandraParams) SetSelectedDataCenters(v []string)`

SetSelectedDataCenters sets SelectedDataCenters field to given value.

### HasSelectedDataCenters

`func (o *RecoverCassandraParams) HasSelectedDataCenters() bool`

HasSelectedDataCenters returns a boolean if a field has been set.

### GetStagingDirectoryList

`func (o *RecoverCassandraParams) GetStagingDirectoryList() []string`

GetStagingDirectoryList returns the StagingDirectoryList field if non-nil, zero value otherwise.

### GetStagingDirectoryListOk

`func (o *RecoverCassandraParams) GetStagingDirectoryListOk() (*[]string, bool)`

GetStagingDirectoryListOk returns a tuple with the StagingDirectoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDirectoryList

`func (o *RecoverCassandraParams) SetStagingDirectoryList(v []string)`

SetStagingDirectoryList sets StagingDirectoryList field to given value.

### HasStagingDirectoryList

`func (o *RecoverCassandraParams) HasStagingDirectoryList() bool`

HasStagingDirectoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


