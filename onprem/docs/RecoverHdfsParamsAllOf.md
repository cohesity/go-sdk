# RecoverHdfsParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | [**[]RecoverHdfsSnapshotParams**](RecoverHdfsSnapshotParams.md) | Specifies the local snapshot ids of the Objects to be recovered. | 
**TargetDirectory** | Pointer to **NullableString** | Specifies the target directory where files and folders are to be recovered. If not set, this will be taken as &#39;/&#39; | [optional] 

## Methods

### NewRecoverHdfsParamsAllOf

`func NewRecoverHdfsParamsAllOf(snapshots []RecoverHdfsSnapshotParams, ) *RecoverHdfsParamsAllOf`

NewRecoverHdfsParamsAllOf instantiates a new RecoverHdfsParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHdfsParamsAllOfWithDefaults

`func NewRecoverHdfsParamsAllOfWithDefaults() *RecoverHdfsParamsAllOf`

NewRecoverHdfsParamsAllOfWithDefaults instantiates a new RecoverHdfsParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *RecoverHdfsParamsAllOf) GetSnapshots() []RecoverHdfsSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverHdfsParamsAllOf) GetSnapshotsOk() (*[]RecoverHdfsSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverHdfsParamsAllOf) SetSnapshots(v []RecoverHdfsSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverHdfsParamsAllOf) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverHdfsParamsAllOf) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetTargetDirectory

`func (o *RecoverHdfsParamsAllOf) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *RecoverHdfsParamsAllOf) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *RecoverHdfsParamsAllOf) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.

### HasTargetDirectory

`func (o *RecoverHdfsParamsAllOf) HasTargetDirectory() bool`

HasTargetDirectory returns a boolean if a field has been set.

### SetTargetDirectoryNil

`func (o *RecoverHdfsParamsAllOf) SetTargetDirectoryNil(b bool)`

 SetTargetDirectoryNil sets the value for TargetDirectory to be an explicit nil

### UnsetTargetDirectory
`func (o *RecoverHdfsParamsAllOf) UnsetTargetDirectory()`

UnsetTargetDirectory ensures that no value is present for TargetDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


