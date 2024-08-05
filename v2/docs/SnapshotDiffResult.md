# SnapshotDiffResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileOperations** | Pointer to [**[]FileOperation**](FileOperation.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewSnapshotDiffResult

`func NewSnapshotDiffResult(status string, ) *SnapshotDiffResult`

NewSnapshotDiffResult instantiates a new SnapshotDiffResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotDiffResultWithDefaults

`func NewSnapshotDiffResultWithDefaults() *SnapshotDiffResult`

NewSnapshotDiffResultWithDefaults instantiates a new SnapshotDiffResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileOperations

`func (o *SnapshotDiffResult) GetFileOperations() []FileOperation`

GetFileOperations returns the FileOperations field if non-nil, zero value otherwise.

### GetFileOperationsOk

`func (o *SnapshotDiffResult) GetFileOperationsOk() (*[]FileOperation, bool)`

GetFileOperationsOk returns a tuple with the FileOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOperations

`func (o *SnapshotDiffResult) SetFileOperations(v []FileOperation)`

SetFileOperations sets FileOperations field to given value.

### HasFileOperations

`func (o *SnapshotDiffResult) HasFileOperations() bool`

HasFileOperations returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotDiffResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotDiffResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotDiffResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


