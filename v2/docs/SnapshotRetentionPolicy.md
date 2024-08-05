# SnapshotRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumOfDaysToKeep** | Pointer to **NullableInt64** | Number of days to keep the snapshot. | [optional] 
**NumOfVersionsToKeep** | Pointer to **NullableInt64** | Number of snapshot versions to keep. | [optional] 
**Suspended** | Pointer to **NullableBool** | Bool to denote if the policy is suspended. | [optional] 

## Methods

### NewSnapshotRetentionPolicy

`func NewSnapshotRetentionPolicy() *SnapshotRetentionPolicy`

NewSnapshotRetentionPolicy instantiates a new SnapshotRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetentionPolicyWithDefaults

`func NewSnapshotRetentionPolicyWithDefaults() *SnapshotRetentionPolicy`

NewSnapshotRetentionPolicyWithDefaults instantiates a new SnapshotRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumOfDaysToKeep

`func (o *SnapshotRetentionPolicy) GetNumOfDaysToKeep() int64`

GetNumOfDaysToKeep returns the NumOfDaysToKeep field if non-nil, zero value otherwise.

### GetNumOfDaysToKeepOk

`func (o *SnapshotRetentionPolicy) GetNumOfDaysToKeepOk() (*int64, bool)`

GetNumOfDaysToKeepOk returns a tuple with the NumOfDaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfDaysToKeep

`func (o *SnapshotRetentionPolicy) SetNumOfDaysToKeep(v int64)`

SetNumOfDaysToKeep sets NumOfDaysToKeep field to given value.

### HasNumOfDaysToKeep

`func (o *SnapshotRetentionPolicy) HasNumOfDaysToKeep() bool`

HasNumOfDaysToKeep returns a boolean if a field has been set.

### SetNumOfDaysToKeepNil

`func (o *SnapshotRetentionPolicy) SetNumOfDaysToKeepNil(b bool)`

 SetNumOfDaysToKeepNil sets the value for NumOfDaysToKeep to be an explicit nil

### UnsetNumOfDaysToKeep
`func (o *SnapshotRetentionPolicy) UnsetNumOfDaysToKeep()`

UnsetNumOfDaysToKeep ensures that no value is present for NumOfDaysToKeep, not even an explicit nil
### GetNumOfVersionsToKeep

`func (o *SnapshotRetentionPolicy) GetNumOfVersionsToKeep() int64`

GetNumOfVersionsToKeep returns the NumOfVersionsToKeep field if non-nil, zero value otherwise.

### GetNumOfVersionsToKeepOk

`func (o *SnapshotRetentionPolicy) GetNumOfVersionsToKeepOk() (*int64, bool)`

GetNumOfVersionsToKeepOk returns a tuple with the NumOfVersionsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfVersionsToKeep

`func (o *SnapshotRetentionPolicy) SetNumOfVersionsToKeep(v int64)`

SetNumOfVersionsToKeep sets NumOfVersionsToKeep field to given value.

### HasNumOfVersionsToKeep

`func (o *SnapshotRetentionPolicy) HasNumOfVersionsToKeep() bool`

HasNumOfVersionsToKeep returns a boolean if a field has been set.

### SetNumOfVersionsToKeepNil

`func (o *SnapshotRetentionPolicy) SetNumOfVersionsToKeepNil(b bool)`

 SetNumOfVersionsToKeepNil sets the value for NumOfVersionsToKeep to be an explicit nil

### UnsetNumOfVersionsToKeep
`func (o *SnapshotRetentionPolicy) UnsetNumOfVersionsToKeep()`

UnsetNumOfVersionsToKeep ensures that no value is present for NumOfVersionsToKeep, not even an explicit nil
### GetSuspended

`func (o *SnapshotRetentionPolicy) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *SnapshotRetentionPolicy) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *SnapshotRetentionPolicy) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *SnapshotRetentionPolicy) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *SnapshotRetentionPolicy) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *SnapshotRetentionPolicy) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


