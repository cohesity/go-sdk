# ClusterSnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotRetentionPolicy** | Pointer to [**SnapshotRetentionPolicy**](SnapshotRetentionPolicy.md) |  | [optional] 
**SnapshotSchedulePolicy** | Pointer to [**SnapshotSchedulePolicy**](SnapshotSchedulePolicy.md) |  | [optional] 

## Methods

### NewClusterSnapshotPolicy

`func NewClusterSnapshotPolicy() *ClusterSnapshotPolicy`

NewClusterSnapshotPolicy instantiates a new ClusterSnapshotPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSnapshotPolicyWithDefaults

`func NewClusterSnapshotPolicyWithDefaults() *ClusterSnapshotPolicy`

NewClusterSnapshotPolicyWithDefaults instantiates a new ClusterSnapshotPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionPolicy

`func (o *ClusterSnapshotPolicy) GetSnapshotRetentionPolicy() SnapshotRetentionPolicy`

GetSnapshotRetentionPolicy returns the SnapshotRetentionPolicy field if non-nil, zero value otherwise.

### GetSnapshotRetentionPolicyOk

`func (o *ClusterSnapshotPolicy) GetSnapshotRetentionPolicyOk() (*SnapshotRetentionPolicy, bool)`

GetSnapshotRetentionPolicyOk returns a tuple with the SnapshotRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionPolicy

`func (o *ClusterSnapshotPolicy) SetSnapshotRetentionPolicy(v SnapshotRetentionPolicy)`

SetSnapshotRetentionPolicy sets SnapshotRetentionPolicy field to given value.

### HasSnapshotRetentionPolicy

`func (o *ClusterSnapshotPolicy) HasSnapshotRetentionPolicy() bool`

HasSnapshotRetentionPolicy returns a boolean if a field has been set.

### GetSnapshotSchedulePolicy

`func (o *ClusterSnapshotPolicy) GetSnapshotSchedulePolicy() SnapshotSchedulePolicy`

GetSnapshotSchedulePolicy returns the SnapshotSchedulePolicy field if non-nil, zero value otherwise.

### GetSnapshotSchedulePolicyOk

`func (o *ClusterSnapshotPolicy) GetSnapshotSchedulePolicyOk() (*SnapshotSchedulePolicy, bool)`

GetSnapshotSchedulePolicyOk returns a tuple with the SnapshotSchedulePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedulePolicy

`func (o *ClusterSnapshotPolicy) SetSnapshotSchedulePolicy(v SnapshotSchedulePolicy)`

SetSnapshotSchedulePolicy sets SnapshotSchedulePolicy field to given value.

### HasSnapshotSchedulePolicy

`func (o *ClusterSnapshotPolicy) HasSnapshotSchedulePolicy() bool`

HasSnapshotSchedulePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


