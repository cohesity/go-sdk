# StorageArraySnapshotMaxSnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSnapshots** | Pointer to **NullableInt32** | Specifies max number of storage snapshots allowed per volume/lun. | [optional] 

## Methods

### NewStorageArraySnapshotMaxSnapshotConfig

`func NewStorageArraySnapshotMaxSnapshotConfig() *StorageArraySnapshotMaxSnapshotConfig`

NewStorageArraySnapshotMaxSnapshotConfig instantiates a new StorageArraySnapshotMaxSnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageArraySnapshotMaxSnapshotConfigWithDefaults

`func NewStorageArraySnapshotMaxSnapshotConfigWithDefaults() *StorageArraySnapshotMaxSnapshotConfig`

NewStorageArraySnapshotMaxSnapshotConfigWithDefaults instantiates a new StorageArraySnapshotMaxSnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSnapshots

`func (o *StorageArraySnapshotMaxSnapshotConfig) GetMaxSnapshots() int32`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *StorageArraySnapshotMaxSnapshotConfig) GetMaxSnapshotsOk() (*int32, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *StorageArraySnapshotMaxSnapshotConfig) SetMaxSnapshots(v int32)`

SetMaxSnapshots sets MaxSnapshots field to given value.

### HasMaxSnapshots

`func (o *StorageArraySnapshotMaxSnapshotConfig) HasMaxSnapshots() bool`

HasMaxSnapshots returns a boolean if a field has been set.

### SetMaxSnapshotsNil

`func (o *StorageArraySnapshotMaxSnapshotConfig) SetMaxSnapshotsNil(b bool)`

 SetMaxSnapshotsNil sets the value for MaxSnapshots to be an explicit nil

### UnsetMaxSnapshots
`func (o *StorageArraySnapshotMaxSnapshotConfig) UnsetMaxSnapshots()`

UnsetMaxSnapshots ensures that no value is present for MaxSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


