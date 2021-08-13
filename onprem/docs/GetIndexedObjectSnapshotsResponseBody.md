# GetIndexedObjectSnapshotsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]IndexedObjectSnapshot**](IndexedObjectSnapshot.md) | Specifies a list of snapshots containing the indexed object. | [optional] 

## Methods

### NewGetIndexedObjectSnapshotsResponseBody

`func NewGetIndexedObjectSnapshotsResponseBody() *GetIndexedObjectSnapshotsResponseBody`

NewGetIndexedObjectSnapshotsResponseBody instantiates a new GetIndexedObjectSnapshotsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIndexedObjectSnapshotsResponseBodyWithDefaults

`func NewGetIndexedObjectSnapshotsResponseBodyWithDefaults() *GetIndexedObjectSnapshotsResponseBody`

NewGetIndexedObjectSnapshotsResponseBodyWithDefaults instantiates a new GetIndexedObjectSnapshotsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *GetIndexedObjectSnapshotsResponseBody) GetSnapshots() []IndexedObjectSnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *GetIndexedObjectSnapshotsResponseBody) GetSnapshotsOk() (*[]IndexedObjectSnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *GetIndexedObjectSnapshotsResponseBody) SetSnapshots(v []IndexedObjectSnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *GetIndexedObjectSnapshotsResponseBody) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### SetSnapshotsNil

`func (o *GetIndexedObjectSnapshotsResponseBody) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *GetIndexedObjectSnapshotsResponseBody) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


