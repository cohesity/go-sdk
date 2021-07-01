# LatestProtectionJobRunInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestSnapshotInfo** | Pointer to [**LatestProtectionRun**](LatestProtectionRun.md) |  | [optional] 
**LocationName** | Pointer to **NullableString** | Specifies the name of location that the object is backedup to. | [optional] 
**NumSnapshots** | Pointer to **NullableInt64** | Specifies of number of successful snapshots. | [optional] 

## Methods

### NewLatestProtectionJobRunInfo

`func NewLatestProtectionJobRunInfo() *LatestProtectionJobRunInfo`

NewLatestProtectionJobRunInfo instantiates a new LatestProtectionJobRunInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestProtectionJobRunInfoWithDefaults

`func NewLatestProtectionJobRunInfoWithDefaults() *LatestProtectionJobRunInfo`

NewLatestProtectionJobRunInfoWithDefaults instantiates a new LatestProtectionJobRunInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestSnapshotInfo

`func (o *LatestProtectionJobRunInfo) GetLatestSnapshotInfo() LatestProtectionRun`

GetLatestSnapshotInfo returns the LatestSnapshotInfo field if non-nil, zero value otherwise.

### GetLatestSnapshotInfoOk

`func (o *LatestProtectionJobRunInfo) GetLatestSnapshotInfoOk() (*LatestProtectionRun, bool)`

GetLatestSnapshotInfoOk returns a tuple with the LatestSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotInfo

`func (o *LatestProtectionJobRunInfo) SetLatestSnapshotInfo(v LatestProtectionRun)`

SetLatestSnapshotInfo sets LatestSnapshotInfo field to given value.

### HasLatestSnapshotInfo

`func (o *LatestProtectionJobRunInfo) HasLatestSnapshotInfo() bool`

HasLatestSnapshotInfo returns a boolean if a field has been set.

### GetLocationName

`func (o *LatestProtectionJobRunInfo) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *LatestProtectionJobRunInfo) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *LatestProtectionJobRunInfo) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *LatestProtectionJobRunInfo) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *LatestProtectionJobRunInfo) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *LatestProtectionJobRunInfo) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetNumSnapshots

`func (o *LatestProtectionJobRunInfo) GetNumSnapshots() int64`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *LatestProtectionJobRunInfo) GetNumSnapshotsOk() (*int64, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *LatestProtectionJobRunInfo) SetNumSnapshots(v int64)`

SetNumSnapshots sets NumSnapshots field to given value.

### HasNumSnapshots

`func (o *LatestProtectionJobRunInfo) HasNumSnapshots() bool`

HasNumSnapshots returns a boolean if a field has been set.

### SetNumSnapshotsNil

`func (o *LatestProtectionJobRunInfo) SetNumSnapshotsNil(b bool)`

 SetNumSnapshotsNil sets the value for NumSnapshots to be an explicit nil

### UnsetNumSnapshots
`func (o *LatestProtectionJobRunInfo) UnsetNumSnapshots()`

UnsetNumSnapshots ensures that no value is present for NumSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


