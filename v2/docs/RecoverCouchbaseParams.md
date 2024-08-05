# RecoverCouchbaseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a recovery job. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 
**AppendDocuments** | Pointer to **NullableBool** | If set to true, docuements from the bucket being recovered will be appended into the bucket at the destination. | [optional] 
**DdlOnlyRecovery** | Pointer to **NullableBool** | Set to true to recover only the bucket configurations. No documents will be recovered. | [optional] 
**FilterDocumentsParams** | [**FilterDocumentsParams**](FilterDocumentsParams.md) |  | 
**OverwriteUsers** | Pointer to **NullableBool** | If set to true existing users will be replaced with users from the bucket being recovered. | [optional] 
**Snapshots** | [**[]RecoverCouchbaseSnapshotParams**](RecoverCouchbaseSnapshotParams.md) | Specifies the local snapshot ids of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 

## Methods

### NewRecoverCouchbaseParams

`func NewRecoverCouchbaseParams(filterDocumentsParams FilterDocumentsParams, snapshots []RecoverCouchbaseSnapshotParams, ) *RecoverCouchbaseParams`

NewRecoverCouchbaseParams instantiates a new RecoverCouchbaseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCouchbaseParamsWithDefaults

`func NewRecoverCouchbaseParamsWithDefaults() *RecoverCouchbaseParams`

NewRecoverCouchbaseParamsWithDefaults instantiates a new RecoverCouchbaseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfigs

`func (o *RecoverCouchbaseParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *RecoverCouchbaseParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *RecoverCouchbaseParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *RecoverCouchbaseParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *RecoverCouchbaseParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *RecoverCouchbaseParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetBandwidthMBPS

`func (o *RecoverCouchbaseParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *RecoverCouchbaseParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *RecoverCouchbaseParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *RecoverCouchbaseParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *RecoverCouchbaseParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *RecoverCouchbaseParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *RecoverCouchbaseParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverCouchbaseParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverCouchbaseParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverCouchbaseParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverCouchbaseParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverCouchbaseParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetOverwrite

`func (o *RecoverCouchbaseParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverCouchbaseParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverCouchbaseParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverCouchbaseParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverCouchbaseParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverCouchbaseParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRecoverTo

`func (o *RecoverCouchbaseParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverCouchbaseParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverCouchbaseParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverCouchbaseParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverCouchbaseParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverCouchbaseParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetWarnings

`func (o *RecoverCouchbaseParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverCouchbaseParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverCouchbaseParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverCouchbaseParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverCouchbaseParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverCouchbaseParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetAppendDocuments

`func (o *RecoverCouchbaseParams) GetAppendDocuments() bool`

GetAppendDocuments returns the AppendDocuments field if non-nil, zero value otherwise.

### GetAppendDocumentsOk

`func (o *RecoverCouchbaseParams) GetAppendDocumentsOk() (*bool, bool)`

GetAppendDocumentsOk returns a tuple with the AppendDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendDocuments

`func (o *RecoverCouchbaseParams) SetAppendDocuments(v bool)`

SetAppendDocuments sets AppendDocuments field to given value.

### HasAppendDocuments

`func (o *RecoverCouchbaseParams) HasAppendDocuments() bool`

HasAppendDocuments returns a boolean if a field has been set.

### SetAppendDocumentsNil

`func (o *RecoverCouchbaseParams) SetAppendDocumentsNil(b bool)`

 SetAppendDocumentsNil sets the value for AppendDocuments to be an explicit nil

### UnsetAppendDocuments
`func (o *RecoverCouchbaseParams) UnsetAppendDocuments()`

UnsetAppendDocuments ensures that no value is present for AppendDocuments, not even an explicit nil
### GetDdlOnlyRecovery

`func (o *RecoverCouchbaseParams) GetDdlOnlyRecovery() bool`

GetDdlOnlyRecovery returns the DdlOnlyRecovery field if non-nil, zero value otherwise.

### GetDdlOnlyRecoveryOk

`func (o *RecoverCouchbaseParams) GetDdlOnlyRecoveryOk() (*bool, bool)`

GetDdlOnlyRecoveryOk returns a tuple with the DdlOnlyRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdlOnlyRecovery

`func (o *RecoverCouchbaseParams) SetDdlOnlyRecovery(v bool)`

SetDdlOnlyRecovery sets DdlOnlyRecovery field to given value.

### HasDdlOnlyRecovery

`func (o *RecoverCouchbaseParams) HasDdlOnlyRecovery() bool`

HasDdlOnlyRecovery returns a boolean if a field has been set.

### SetDdlOnlyRecoveryNil

`func (o *RecoverCouchbaseParams) SetDdlOnlyRecoveryNil(b bool)`

 SetDdlOnlyRecoveryNil sets the value for DdlOnlyRecovery to be an explicit nil

### UnsetDdlOnlyRecovery
`func (o *RecoverCouchbaseParams) UnsetDdlOnlyRecovery()`

UnsetDdlOnlyRecovery ensures that no value is present for DdlOnlyRecovery, not even an explicit nil
### GetFilterDocumentsParams

`func (o *RecoverCouchbaseParams) GetFilterDocumentsParams() FilterDocumentsParams`

GetFilterDocumentsParams returns the FilterDocumentsParams field if non-nil, zero value otherwise.

### GetFilterDocumentsParamsOk

`func (o *RecoverCouchbaseParams) GetFilterDocumentsParamsOk() (*FilterDocumentsParams, bool)`

GetFilterDocumentsParamsOk returns a tuple with the FilterDocumentsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDocumentsParams

`func (o *RecoverCouchbaseParams) SetFilterDocumentsParams(v FilterDocumentsParams)`

SetFilterDocumentsParams sets FilterDocumentsParams field to given value.


### GetOverwriteUsers

`func (o *RecoverCouchbaseParams) GetOverwriteUsers() bool`

GetOverwriteUsers returns the OverwriteUsers field if non-nil, zero value otherwise.

### GetOverwriteUsersOk

`func (o *RecoverCouchbaseParams) GetOverwriteUsersOk() (*bool, bool)`

GetOverwriteUsersOk returns a tuple with the OverwriteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteUsers

`func (o *RecoverCouchbaseParams) SetOverwriteUsers(v bool)`

SetOverwriteUsers sets OverwriteUsers field to given value.

### HasOverwriteUsers

`func (o *RecoverCouchbaseParams) HasOverwriteUsers() bool`

HasOverwriteUsers returns a boolean if a field has been set.

### SetOverwriteUsersNil

`func (o *RecoverCouchbaseParams) SetOverwriteUsersNil(b bool)`

 SetOverwriteUsersNil sets the value for OverwriteUsers to be an explicit nil

### UnsetOverwriteUsers
`func (o *RecoverCouchbaseParams) UnsetOverwriteUsers()`

UnsetOverwriteUsers ensures that no value is present for OverwriteUsers, not even an explicit nil
### GetSnapshots

`func (o *RecoverCouchbaseParams) GetSnapshots() []RecoverCouchbaseSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverCouchbaseParams) GetSnapshotsOk() (*[]RecoverCouchbaseSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverCouchbaseParams) SetSnapshots(v []RecoverCouchbaseSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverCouchbaseParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverCouchbaseParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverCouchbaseParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverCouchbaseParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverCouchbaseParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverCouchbaseParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverCouchbaseParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverCouchbaseParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


