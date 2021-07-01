# ProtectionObjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestArchivalSnapshotTimeUsecs** | Pointer to **NullableInt64** | Specifies the Unix epoch Timestamp (in microseconds) of the latest Archival Snapshot. | [optional] 
**LatestLocalSnapshotTimeUsecs** | Pointer to **NullableInt64** | Specifies the Unix epoch Timestamp (in microseconds) of the latest Local Snapshot. | [optional] 
**LatestReplicationSnapshotTimeUsecs** | Pointer to **NullableInt64** | Specifies the Unix epoch Timestamp (in microseconds) of the latest Replication Snapshot. | [optional] 
**ParentProtectionSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**ProtectionJobs** | Pointer to [**[]ProtectionRunInstance**](ProtectionRunInstance.md) | Returns the list of Protection Jobs with summary Information. | [optional] 
**ProtectionSource** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the leaf Protection Source Object such as a VM. | [optional] 
**RpoPolicies** | Pointer to [**[]ProtectionPolicy**](ProtectionPolicy.md) | Specifies the id of the RPO policy protecting this object. | [optional] 
**TotalArchivalSnapshots** | Pointer to **NullableInt32** | Specifies the total number of Archival Snapshots. | [optional] 
**TotalLocalSnapshots** | Pointer to **NullableInt32** | Specifies the total number of Local Snapshots. | [optional] 
**TotalReplicationSnapshots** | Pointer to **NullableInt32** | Specifies the total number of Replication Snapshots. | [optional] 

## Methods

### NewProtectionObjectSummary

`func NewProtectionObjectSummary() *ProtectionObjectSummary`

NewProtectionObjectSummary instantiates a new ProtectionObjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionObjectSummaryWithDefaults

`func NewProtectionObjectSummaryWithDefaults() *ProtectionObjectSummary`

NewProtectionObjectSummaryWithDefaults instantiates a new ProtectionObjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestArchivalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) GetLatestArchivalSnapshotTimeUsecs() int64`

GetLatestArchivalSnapshotTimeUsecs returns the LatestArchivalSnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetLatestArchivalSnapshotTimeUsecsOk

`func (o *ProtectionObjectSummary) GetLatestArchivalSnapshotTimeUsecsOk() (*int64, bool)`

GetLatestArchivalSnapshotTimeUsecsOk returns a tuple with the LatestArchivalSnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestArchivalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) SetLatestArchivalSnapshotTimeUsecs(v int64)`

SetLatestArchivalSnapshotTimeUsecs sets LatestArchivalSnapshotTimeUsecs field to given value.

### HasLatestArchivalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) HasLatestArchivalSnapshotTimeUsecs() bool`

HasLatestArchivalSnapshotTimeUsecs returns a boolean if a field has been set.

### SetLatestArchivalSnapshotTimeUsecsNil

`func (o *ProtectionObjectSummary) SetLatestArchivalSnapshotTimeUsecsNil(b bool)`

 SetLatestArchivalSnapshotTimeUsecsNil sets the value for LatestArchivalSnapshotTimeUsecs to be an explicit nil

### UnsetLatestArchivalSnapshotTimeUsecs
`func (o *ProtectionObjectSummary) UnsetLatestArchivalSnapshotTimeUsecs()`

UnsetLatestArchivalSnapshotTimeUsecs ensures that no value is present for LatestArchivalSnapshotTimeUsecs, not even an explicit nil
### GetLatestLocalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) GetLatestLocalSnapshotTimeUsecs() int64`

GetLatestLocalSnapshotTimeUsecs returns the LatestLocalSnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetLatestLocalSnapshotTimeUsecsOk

`func (o *ProtectionObjectSummary) GetLatestLocalSnapshotTimeUsecsOk() (*int64, bool)`

GetLatestLocalSnapshotTimeUsecsOk returns a tuple with the LatestLocalSnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestLocalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) SetLatestLocalSnapshotTimeUsecs(v int64)`

SetLatestLocalSnapshotTimeUsecs sets LatestLocalSnapshotTimeUsecs field to given value.

### HasLatestLocalSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) HasLatestLocalSnapshotTimeUsecs() bool`

HasLatestLocalSnapshotTimeUsecs returns a boolean if a field has been set.

### SetLatestLocalSnapshotTimeUsecsNil

`func (o *ProtectionObjectSummary) SetLatestLocalSnapshotTimeUsecsNil(b bool)`

 SetLatestLocalSnapshotTimeUsecsNil sets the value for LatestLocalSnapshotTimeUsecs to be an explicit nil

### UnsetLatestLocalSnapshotTimeUsecs
`func (o *ProtectionObjectSummary) UnsetLatestLocalSnapshotTimeUsecs()`

UnsetLatestLocalSnapshotTimeUsecs ensures that no value is present for LatestLocalSnapshotTimeUsecs, not even an explicit nil
### GetLatestReplicationSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) GetLatestReplicationSnapshotTimeUsecs() int64`

GetLatestReplicationSnapshotTimeUsecs returns the LatestReplicationSnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetLatestReplicationSnapshotTimeUsecsOk

`func (o *ProtectionObjectSummary) GetLatestReplicationSnapshotTimeUsecsOk() (*int64, bool)`

GetLatestReplicationSnapshotTimeUsecsOk returns a tuple with the LatestReplicationSnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReplicationSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) SetLatestReplicationSnapshotTimeUsecs(v int64)`

SetLatestReplicationSnapshotTimeUsecs sets LatestReplicationSnapshotTimeUsecs field to given value.

### HasLatestReplicationSnapshotTimeUsecs

`func (o *ProtectionObjectSummary) HasLatestReplicationSnapshotTimeUsecs() bool`

HasLatestReplicationSnapshotTimeUsecs returns a boolean if a field has been set.

### SetLatestReplicationSnapshotTimeUsecsNil

`func (o *ProtectionObjectSummary) SetLatestReplicationSnapshotTimeUsecsNil(b bool)`

 SetLatestReplicationSnapshotTimeUsecsNil sets the value for LatestReplicationSnapshotTimeUsecs to be an explicit nil

### UnsetLatestReplicationSnapshotTimeUsecs
`func (o *ProtectionObjectSummary) UnsetLatestReplicationSnapshotTimeUsecs()`

UnsetLatestReplicationSnapshotTimeUsecs ensures that no value is present for LatestReplicationSnapshotTimeUsecs, not even an explicit nil
### GetParentProtectionSource

`func (o *ProtectionObjectSummary) GetParentProtectionSource() ProtectionSource`

GetParentProtectionSource returns the ParentProtectionSource field if non-nil, zero value otherwise.

### GetParentProtectionSourceOk

`func (o *ProtectionObjectSummary) GetParentProtectionSourceOk() (*ProtectionSource, bool)`

GetParentProtectionSourceOk returns a tuple with the ParentProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProtectionSource

`func (o *ProtectionObjectSummary) SetParentProtectionSource(v ProtectionSource)`

SetParentProtectionSource sets ParentProtectionSource field to given value.

### HasParentProtectionSource

`func (o *ProtectionObjectSummary) HasParentProtectionSource() bool`

HasParentProtectionSource returns a boolean if a field has been set.

### GetProtectionJobs

`func (o *ProtectionObjectSummary) GetProtectionJobs() []ProtectionRunInstance`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *ProtectionObjectSummary) GetProtectionJobsOk() (*[]ProtectionRunInstance, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *ProtectionObjectSummary) SetProtectionJobs(v []ProtectionRunInstance)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *ProtectionObjectSummary) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *ProtectionObjectSummary) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *ProtectionObjectSummary) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil
### GetProtectionSource

`func (o *ProtectionObjectSummary) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectionObjectSummary) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectionObjectSummary) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectionObjectSummary) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### SetProtectionSourceNil

`func (o *ProtectionObjectSummary) SetProtectionSourceNil(b bool)`

 SetProtectionSourceNil sets the value for ProtectionSource to be an explicit nil

### UnsetProtectionSource
`func (o *ProtectionObjectSummary) UnsetProtectionSource()`

UnsetProtectionSource ensures that no value is present for ProtectionSource, not even an explicit nil
### GetRpoPolicies

`func (o *ProtectionObjectSummary) GetRpoPolicies() []ProtectionPolicy`

GetRpoPolicies returns the RpoPolicies field if non-nil, zero value otherwise.

### GetRpoPoliciesOk

`func (o *ProtectionObjectSummary) GetRpoPoliciesOk() (*[]ProtectionPolicy, bool)`

GetRpoPoliciesOk returns a tuple with the RpoPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicies

`func (o *ProtectionObjectSummary) SetRpoPolicies(v []ProtectionPolicy)`

SetRpoPolicies sets RpoPolicies field to given value.

### HasRpoPolicies

`func (o *ProtectionObjectSummary) HasRpoPolicies() bool`

HasRpoPolicies returns a boolean if a field has been set.

### SetRpoPoliciesNil

`func (o *ProtectionObjectSummary) SetRpoPoliciesNil(b bool)`

 SetRpoPoliciesNil sets the value for RpoPolicies to be an explicit nil

### UnsetRpoPolicies
`func (o *ProtectionObjectSummary) UnsetRpoPolicies()`

UnsetRpoPolicies ensures that no value is present for RpoPolicies, not even an explicit nil
### GetTotalArchivalSnapshots

`func (o *ProtectionObjectSummary) GetTotalArchivalSnapshots() int32`

GetTotalArchivalSnapshots returns the TotalArchivalSnapshots field if non-nil, zero value otherwise.

### GetTotalArchivalSnapshotsOk

`func (o *ProtectionObjectSummary) GetTotalArchivalSnapshotsOk() (*int32, bool)`

GetTotalArchivalSnapshotsOk returns a tuple with the TotalArchivalSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalArchivalSnapshots

`func (o *ProtectionObjectSummary) SetTotalArchivalSnapshots(v int32)`

SetTotalArchivalSnapshots sets TotalArchivalSnapshots field to given value.

### HasTotalArchivalSnapshots

`func (o *ProtectionObjectSummary) HasTotalArchivalSnapshots() bool`

HasTotalArchivalSnapshots returns a boolean if a field has been set.

### SetTotalArchivalSnapshotsNil

`func (o *ProtectionObjectSummary) SetTotalArchivalSnapshotsNil(b bool)`

 SetTotalArchivalSnapshotsNil sets the value for TotalArchivalSnapshots to be an explicit nil

### UnsetTotalArchivalSnapshots
`func (o *ProtectionObjectSummary) UnsetTotalArchivalSnapshots()`

UnsetTotalArchivalSnapshots ensures that no value is present for TotalArchivalSnapshots, not even an explicit nil
### GetTotalLocalSnapshots

`func (o *ProtectionObjectSummary) GetTotalLocalSnapshots() int32`

GetTotalLocalSnapshots returns the TotalLocalSnapshots field if non-nil, zero value otherwise.

### GetTotalLocalSnapshotsOk

`func (o *ProtectionObjectSummary) GetTotalLocalSnapshotsOk() (*int32, bool)`

GetTotalLocalSnapshotsOk returns a tuple with the TotalLocalSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLocalSnapshots

`func (o *ProtectionObjectSummary) SetTotalLocalSnapshots(v int32)`

SetTotalLocalSnapshots sets TotalLocalSnapshots field to given value.

### HasTotalLocalSnapshots

`func (o *ProtectionObjectSummary) HasTotalLocalSnapshots() bool`

HasTotalLocalSnapshots returns a boolean if a field has been set.

### SetTotalLocalSnapshotsNil

`func (o *ProtectionObjectSummary) SetTotalLocalSnapshotsNil(b bool)`

 SetTotalLocalSnapshotsNil sets the value for TotalLocalSnapshots to be an explicit nil

### UnsetTotalLocalSnapshots
`func (o *ProtectionObjectSummary) UnsetTotalLocalSnapshots()`

UnsetTotalLocalSnapshots ensures that no value is present for TotalLocalSnapshots, not even an explicit nil
### GetTotalReplicationSnapshots

`func (o *ProtectionObjectSummary) GetTotalReplicationSnapshots() int32`

GetTotalReplicationSnapshots returns the TotalReplicationSnapshots field if non-nil, zero value otherwise.

### GetTotalReplicationSnapshotsOk

`func (o *ProtectionObjectSummary) GetTotalReplicationSnapshotsOk() (*int32, bool)`

GetTotalReplicationSnapshotsOk returns a tuple with the TotalReplicationSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReplicationSnapshots

`func (o *ProtectionObjectSummary) SetTotalReplicationSnapshots(v int32)`

SetTotalReplicationSnapshots sets TotalReplicationSnapshots field to given value.

### HasTotalReplicationSnapshots

`func (o *ProtectionObjectSummary) HasTotalReplicationSnapshots() bool`

HasTotalReplicationSnapshots returns a boolean if a field has been set.

### SetTotalReplicationSnapshotsNil

`func (o *ProtectionObjectSummary) SetTotalReplicationSnapshotsNil(b bool)`

 SetTotalReplicationSnapshotsNil sets the value for TotalReplicationSnapshots to be an explicit nil

### UnsetTotalReplicationSnapshots
`func (o *ProtectionObjectSummary) UnsetTotalReplicationSnapshots()`

UnsetTotalReplicationSnapshots ensures that no value is present for TotalReplicationSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


