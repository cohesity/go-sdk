# GaiaIndexingStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumConvertedDocs** | Pointer to **NullableInt64** |  | [optional] 
**NumErrorsSeen** | Pointer to **NullableInt64** |  | [optional] 
**NumFinishedDocs** | Pointer to **NullableInt64** |  | [optional] 
**NumFinishedObjects** | Pointer to **NullableInt64** |  | [optional] 
**NumFinishedSnapshots** | Pointer to **NullableInt64** |  | [optional] 
**NumFinishedSubObjects** | Pointer to **NullableInt64** |  | [optional] 
**NumIndexedBytes** | Pointer to **NullableInt64** |  | [optional] 
**NumIndexedChunks** | Pointer to **NullableInt64** | Embeddings are created over chunks of a doc | [optional] 
**NumIndexedDocs** | Pointer to **NullableInt64** |  | [optional] 
**NumObjects** | Pointer to **NullableInt64** |  | [optional] 
**NumSnapshots** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewGaiaIndexingStats

`func NewGaiaIndexingStats() *GaiaIndexingStats`

NewGaiaIndexingStats instantiates a new GaiaIndexingStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaiaIndexingStatsWithDefaults

`func NewGaiaIndexingStatsWithDefaults() *GaiaIndexingStats`

NewGaiaIndexingStatsWithDefaults instantiates a new GaiaIndexingStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumConvertedDocs

`func (o *GaiaIndexingStats) GetNumConvertedDocs() int64`

GetNumConvertedDocs returns the NumConvertedDocs field if non-nil, zero value otherwise.

### GetNumConvertedDocsOk

`func (o *GaiaIndexingStats) GetNumConvertedDocsOk() (*int64, bool)`

GetNumConvertedDocsOk returns a tuple with the NumConvertedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConvertedDocs

`func (o *GaiaIndexingStats) SetNumConvertedDocs(v int64)`

SetNumConvertedDocs sets NumConvertedDocs field to given value.

### HasNumConvertedDocs

`func (o *GaiaIndexingStats) HasNumConvertedDocs() bool`

HasNumConvertedDocs returns a boolean if a field has been set.

### SetNumConvertedDocsNil

`func (o *GaiaIndexingStats) SetNumConvertedDocsNil(b bool)`

 SetNumConvertedDocsNil sets the value for NumConvertedDocs to be an explicit nil

### UnsetNumConvertedDocs
`func (o *GaiaIndexingStats) UnsetNumConvertedDocs()`

UnsetNumConvertedDocs ensures that no value is present for NumConvertedDocs, not even an explicit nil
### GetNumErrorsSeen

`func (o *GaiaIndexingStats) GetNumErrorsSeen() int64`

GetNumErrorsSeen returns the NumErrorsSeen field if non-nil, zero value otherwise.

### GetNumErrorsSeenOk

`func (o *GaiaIndexingStats) GetNumErrorsSeenOk() (*int64, bool)`

GetNumErrorsSeenOk returns a tuple with the NumErrorsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrorsSeen

`func (o *GaiaIndexingStats) SetNumErrorsSeen(v int64)`

SetNumErrorsSeen sets NumErrorsSeen field to given value.

### HasNumErrorsSeen

`func (o *GaiaIndexingStats) HasNumErrorsSeen() bool`

HasNumErrorsSeen returns a boolean if a field has been set.

### SetNumErrorsSeenNil

`func (o *GaiaIndexingStats) SetNumErrorsSeenNil(b bool)`

 SetNumErrorsSeenNil sets the value for NumErrorsSeen to be an explicit nil

### UnsetNumErrorsSeen
`func (o *GaiaIndexingStats) UnsetNumErrorsSeen()`

UnsetNumErrorsSeen ensures that no value is present for NumErrorsSeen, not even an explicit nil
### GetNumFinishedDocs

`func (o *GaiaIndexingStats) GetNumFinishedDocs() int64`

GetNumFinishedDocs returns the NumFinishedDocs field if non-nil, zero value otherwise.

### GetNumFinishedDocsOk

`func (o *GaiaIndexingStats) GetNumFinishedDocsOk() (*int64, bool)`

GetNumFinishedDocsOk returns a tuple with the NumFinishedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFinishedDocs

`func (o *GaiaIndexingStats) SetNumFinishedDocs(v int64)`

SetNumFinishedDocs sets NumFinishedDocs field to given value.

### HasNumFinishedDocs

`func (o *GaiaIndexingStats) HasNumFinishedDocs() bool`

HasNumFinishedDocs returns a boolean if a field has been set.

### SetNumFinishedDocsNil

`func (o *GaiaIndexingStats) SetNumFinishedDocsNil(b bool)`

 SetNumFinishedDocsNil sets the value for NumFinishedDocs to be an explicit nil

### UnsetNumFinishedDocs
`func (o *GaiaIndexingStats) UnsetNumFinishedDocs()`

UnsetNumFinishedDocs ensures that no value is present for NumFinishedDocs, not even an explicit nil
### GetNumFinishedObjects

`func (o *GaiaIndexingStats) GetNumFinishedObjects() int64`

GetNumFinishedObjects returns the NumFinishedObjects field if non-nil, zero value otherwise.

### GetNumFinishedObjectsOk

`func (o *GaiaIndexingStats) GetNumFinishedObjectsOk() (*int64, bool)`

GetNumFinishedObjectsOk returns a tuple with the NumFinishedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFinishedObjects

`func (o *GaiaIndexingStats) SetNumFinishedObjects(v int64)`

SetNumFinishedObjects sets NumFinishedObjects field to given value.

### HasNumFinishedObjects

`func (o *GaiaIndexingStats) HasNumFinishedObjects() bool`

HasNumFinishedObjects returns a boolean if a field has been set.

### SetNumFinishedObjectsNil

`func (o *GaiaIndexingStats) SetNumFinishedObjectsNil(b bool)`

 SetNumFinishedObjectsNil sets the value for NumFinishedObjects to be an explicit nil

### UnsetNumFinishedObjects
`func (o *GaiaIndexingStats) UnsetNumFinishedObjects()`

UnsetNumFinishedObjects ensures that no value is present for NumFinishedObjects, not even an explicit nil
### GetNumFinishedSnapshots

`func (o *GaiaIndexingStats) GetNumFinishedSnapshots() int64`

GetNumFinishedSnapshots returns the NumFinishedSnapshots field if non-nil, zero value otherwise.

### GetNumFinishedSnapshotsOk

`func (o *GaiaIndexingStats) GetNumFinishedSnapshotsOk() (*int64, bool)`

GetNumFinishedSnapshotsOk returns a tuple with the NumFinishedSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFinishedSnapshots

`func (o *GaiaIndexingStats) SetNumFinishedSnapshots(v int64)`

SetNumFinishedSnapshots sets NumFinishedSnapshots field to given value.

### HasNumFinishedSnapshots

`func (o *GaiaIndexingStats) HasNumFinishedSnapshots() bool`

HasNumFinishedSnapshots returns a boolean if a field has been set.

### SetNumFinishedSnapshotsNil

`func (o *GaiaIndexingStats) SetNumFinishedSnapshotsNil(b bool)`

 SetNumFinishedSnapshotsNil sets the value for NumFinishedSnapshots to be an explicit nil

### UnsetNumFinishedSnapshots
`func (o *GaiaIndexingStats) UnsetNumFinishedSnapshots()`

UnsetNumFinishedSnapshots ensures that no value is present for NumFinishedSnapshots, not even an explicit nil
### GetNumFinishedSubObjects

`func (o *GaiaIndexingStats) GetNumFinishedSubObjects() int64`

GetNumFinishedSubObjects returns the NumFinishedSubObjects field if non-nil, zero value otherwise.

### GetNumFinishedSubObjectsOk

`func (o *GaiaIndexingStats) GetNumFinishedSubObjectsOk() (*int64, bool)`

GetNumFinishedSubObjectsOk returns a tuple with the NumFinishedSubObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFinishedSubObjects

`func (o *GaiaIndexingStats) SetNumFinishedSubObjects(v int64)`

SetNumFinishedSubObjects sets NumFinishedSubObjects field to given value.

### HasNumFinishedSubObjects

`func (o *GaiaIndexingStats) HasNumFinishedSubObjects() bool`

HasNumFinishedSubObjects returns a boolean if a field has been set.

### SetNumFinishedSubObjectsNil

`func (o *GaiaIndexingStats) SetNumFinishedSubObjectsNil(b bool)`

 SetNumFinishedSubObjectsNil sets the value for NumFinishedSubObjects to be an explicit nil

### UnsetNumFinishedSubObjects
`func (o *GaiaIndexingStats) UnsetNumFinishedSubObjects()`

UnsetNumFinishedSubObjects ensures that no value is present for NumFinishedSubObjects, not even an explicit nil
### GetNumIndexedBytes

`func (o *GaiaIndexingStats) GetNumIndexedBytes() int64`

GetNumIndexedBytes returns the NumIndexedBytes field if non-nil, zero value otherwise.

### GetNumIndexedBytesOk

`func (o *GaiaIndexingStats) GetNumIndexedBytesOk() (*int64, bool)`

GetNumIndexedBytesOk returns a tuple with the NumIndexedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIndexedBytes

`func (o *GaiaIndexingStats) SetNumIndexedBytes(v int64)`

SetNumIndexedBytes sets NumIndexedBytes field to given value.

### HasNumIndexedBytes

`func (o *GaiaIndexingStats) HasNumIndexedBytes() bool`

HasNumIndexedBytes returns a boolean if a field has been set.

### SetNumIndexedBytesNil

`func (o *GaiaIndexingStats) SetNumIndexedBytesNil(b bool)`

 SetNumIndexedBytesNil sets the value for NumIndexedBytes to be an explicit nil

### UnsetNumIndexedBytes
`func (o *GaiaIndexingStats) UnsetNumIndexedBytes()`

UnsetNumIndexedBytes ensures that no value is present for NumIndexedBytes, not even an explicit nil
### GetNumIndexedChunks

`func (o *GaiaIndexingStats) GetNumIndexedChunks() int64`

GetNumIndexedChunks returns the NumIndexedChunks field if non-nil, zero value otherwise.

### GetNumIndexedChunksOk

`func (o *GaiaIndexingStats) GetNumIndexedChunksOk() (*int64, bool)`

GetNumIndexedChunksOk returns a tuple with the NumIndexedChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIndexedChunks

`func (o *GaiaIndexingStats) SetNumIndexedChunks(v int64)`

SetNumIndexedChunks sets NumIndexedChunks field to given value.

### HasNumIndexedChunks

`func (o *GaiaIndexingStats) HasNumIndexedChunks() bool`

HasNumIndexedChunks returns a boolean if a field has been set.

### SetNumIndexedChunksNil

`func (o *GaiaIndexingStats) SetNumIndexedChunksNil(b bool)`

 SetNumIndexedChunksNil sets the value for NumIndexedChunks to be an explicit nil

### UnsetNumIndexedChunks
`func (o *GaiaIndexingStats) UnsetNumIndexedChunks()`

UnsetNumIndexedChunks ensures that no value is present for NumIndexedChunks, not even an explicit nil
### GetNumIndexedDocs

`func (o *GaiaIndexingStats) GetNumIndexedDocs() int64`

GetNumIndexedDocs returns the NumIndexedDocs field if non-nil, zero value otherwise.

### GetNumIndexedDocsOk

`func (o *GaiaIndexingStats) GetNumIndexedDocsOk() (*int64, bool)`

GetNumIndexedDocsOk returns a tuple with the NumIndexedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIndexedDocs

`func (o *GaiaIndexingStats) SetNumIndexedDocs(v int64)`

SetNumIndexedDocs sets NumIndexedDocs field to given value.

### HasNumIndexedDocs

`func (o *GaiaIndexingStats) HasNumIndexedDocs() bool`

HasNumIndexedDocs returns a boolean if a field has been set.

### SetNumIndexedDocsNil

`func (o *GaiaIndexingStats) SetNumIndexedDocsNil(b bool)`

 SetNumIndexedDocsNil sets the value for NumIndexedDocs to be an explicit nil

### UnsetNumIndexedDocs
`func (o *GaiaIndexingStats) UnsetNumIndexedDocs()`

UnsetNumIndexedDocs ensures that no value is present for NumIndexedDocs, not even an explicit nil
### GetNumObjects

`func (o *GaiaIndexingStats) GetNumObjects() int64`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *GaiaIndexingStats) GetNumObjectsOk() (*int64, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *GaiaIndexingStats) SetNumObjects(v int64)`

SetNumObjects sets NumObjects field to given value.

### HasNumObjects

`func (o *GaiaIndexingStats) HasNumObjects() bool`

HasNumObjects returns a boolean if a field has been set.

### SetNumObjectsNil

`func (o *GaiaIndexingStats) SetNumObjectsNil(b bool)`

 SetNumObjectsNil sets the value for NumObjects to be an explicit nil

### UnsetNumObjects
`func (o *GaiaIndexingStats) UnsetNumObjects()`

UnsetNumObjects ensures that no value is present for NumObjects, not even an explicit nil
### GetNumSnapshots

`func (o *GaiaIndexingStats) GetNumSnapshots() int64`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *GaiaIndexingStats) GetNumSnapshotsOk() (*int64, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *GaiaIndexingStats) SetNumSnapshots(v int64)`

SetNumSnapshots sets NumSnapshots field to given value.

### HasNumSnapshots

`func (o *GaiaIndexingStats) HasNumSnapshots() bool`

HasNumSnapshots returns a boolean if a field has been set.

### SetNumSnapshotsNil

`func (o *GaiaIndexingStats) SetNumSnapshotsNil(b bool)`

 SetNumSnapshotsNil sets the value for NumSnapshots to be an explicit nil

### UnsetNumSnapshots
`func (o *GaiaIndexingStats) UnsetNumSnapshots()`

UnsetNumSnapshots ensures that no value is present for NumSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


