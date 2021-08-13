# RecoverVmwareSnapshotParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildSnapshots** | Pointer to [**[]RecoverVmwareChildSnapshotParams**](RecoverVmwareChildSnapshotParams.md) | Specifies optional information about any child snapshots of this object. For example a VCD snapshot may have child VM information populated here. | [optional] 

## Methods

### NewRecoverVmwareSnapshotParamsAllOf

`func NewRecoverVmwareSnapshotParamsAllOf() *RecoverVmwareSnapshotParamsAllOf`

NewRecoverVmwareSnapshotParamsAllOf instantiates a new RecoverVmwareSnapshotParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareSnapshotParamsAllOfWithDefaults

`func NewRecoverVmwareSnapshotParamsAllOfWithDefaults() *RecoverVmwareSnapshotParamsAllOf`

NewRecoverVmwareSnapshotParamsAllOfWithDefaults instantiates a new RecoverVmwareSnapshotParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildSnapshots

`func (o *RecoverVmwareSnapshotParamsAllOf) GetChildSnapshots() []RecoverVmwareChildSnapshotParams`

GetChildSnapshots returns the ChildSnapshots field if non-nil, zero value otherwise.

### GetChildSnapshotsOk

`func (o *RecoverVmwareSnapshotParamsAllOf) GetChildSnapshotsOk() (*[]RecoverVmwareChildSnapshotParams, bool)`

GetChildSnapshotsOk returns a tuple with the ChildSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSnapshots

`func (o *RecoverVmwareSnapshotParamsAllOf) SetChildSnapshots(v []RecoverVmwareChildSnapshotParams)`

SetChildSnapshots sets ChildSnapshots field to given value.

### HasChildSnapshots

`func (o *RecoverVmwareSnapshotParamsAllOf) HasChildSnapshots() bool`

HasChildSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


