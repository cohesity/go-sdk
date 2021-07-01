# ProtectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time for object retention. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the object. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the id of the policy. | [optional] 
**ProtectionJobId** | Pointer to **NullableInt64** | Specifies the id of the protection job. | [optional] 
**ProtectionJobName** | Pointer to **NullableString** | Specifies the protection job name which protects this object. | [optional] 
**RetentionPeriod** | Pointer to **NullableInt64** | Specifies the retention period. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time for object retention. | [optional] 
**StorageDomain** | Pointer to **NullableString** | Specifies the storage domain name. | [optional] 
**TotalSnapshots** | Pointer to **NullableInt64** | Specifies the total number of snapshots. | [optional] 

## Methods

### NewProtectionInfo

`func NewProtectionInfo() *ProtectionInfo`

NewProtectionInfo instantiates a new ProtectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionInfoWithDefaults

`func NewProtectionInfoWithDefaults() *ProtectionInfo`

NewProtectionInfoWithDefaults instantiates a new ProtectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *ProtectionInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetLocation

`func (o *ProtectionInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProtectionInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProtectionInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProtectionInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ProtectionInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ProtectionInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetPolicyId

`func (o *ProtectionInfo) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionInfo) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionInfo) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ProtectionInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ProtectionInfo) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionInfo) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetProtectionJobId

`func (o *ProtectionInfo) GetProtectionJobId() int64`

GetProtectionJobId returns the ProtectionJobId field if non-nil, zero value otherwise.

### GetProtectionJobIdOk

`func (o *ProtectionInfo) GetProtectionJobIdOk() (*int64, bool)`

GetProtectionJobIdOk returns a tuple with the ProtectionJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobId

`func (o *ProtectionInfo) SetProtectionJobId(v int64)`

SetProtectionJobId sets ProtectionJobId field to given value.

### HasProtectionJobId

`func (o *ProtectionInfo) HasProtectionJobId() bool`

HasProtectionJobId returns a boolean if a field has been set.

### SetProtectionJobIdNil

`func (o *ProtectionInfo) SetProtectionJobIdNil(b bool)`

 SetProtectionJobIdNil sets the value for ProtectionJobId to be an explicit nil

### UnsetProtectionJobId
`func (o *ProtectionInfo) UnsetProtectionJobId()`

UnsetProtectionJobId ensures that no value is present for ProtectionJobId, not even an explicit nil
### GetProtectionJobName

`func (o *ProtectionInfo) GetProtectionJobName() string`

GetProtectionJobName returns the ProtectionJobName field if non-nil, zero value otherwise.

### GetProtectionJobNameOk

`func (o *ProtectionInfo) GetProtectionJobNameOk() (*string, bool)`

GetProtectionJobNameOk returns a tuple with the ProtectionJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobName

`func (o *ProtectionInfo) SetProtectionJobName(v string)`

SetProtectionJobName sets ProtectionJobName field to given value.

### HasProtectionJobName

`func (o *ProtectionInfo) HasProtectionJobName() bool`

HasProtectionJobName returns a boolean if a field has been set.

### SetProtectionJobNameNil

`func (o *ProtectionInfo) SetProtectionJobNameNil(b bool)`

 SetProtectionJobNameNil sets the value for ProtectionJobName to be an explicit nil

### UnsetProtectionJobName
`func (o *ProtectionInfo) UnsetProtectionJobName()`

UnsetProtectionJobName ensures that no value is present for ProtectionJobName, not even an explicit nil
### GetRetentionPeriod

`func (o *ProtectionInfo) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *ProtectionInfo) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *ProtectionInfo) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *ProtectionInfo) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### SetRetentionPeriodNil

`func (o *ProtectionInfo) SetRetentionPeriodNil(b bool)`

 SetRetentionPeriodNil sets the value for RetentionPeriod to be an explicit nil

### UnsetRetentionPeriod
`func (o *ProtectionInfo) UnsetRetentionPeriod()`

UnsetRetentionPeriod ensures that no value is present for RetentionPeriod, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProtectionInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProtectionInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProtectionInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProtectionInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProtectionInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProtectionInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStorageDomain

`func (o *ProtectionInfo) GetStorageDomain() string`

GetStorageDomain returns the StorageDomain field if non-nil, zero value otherwise.

### GetStorageDomainOk

`func (o *ProtectionInfo) GetStorageDomainOk() (*string, bool)`

GetStorageDomainOk returns a tuple with the StorageDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomain

`func (o *ProtectionInfo) SetStorageDomain(v string)`

SetStorageDomain sets StorageDomain field to given value.

### HasStorageDomain

`func (o *ProtectionInfo) HasStorageDomain() bool`

HasStorageDomain returns a boolean if a field has been set.

### SetStorageDomainNil

`func (o *ProtectionInfo) SetStorageDomainNil(b bool)`

 SetStorageDomainNil sets the value for StorageDomain to be an explicit nil

### UnsetStorageDomain
`func (o *ProtectionInfo) UnsetStorageDomain()`

UnsetStorageDomain ensures that no value is present for StorageDomain, not even an explicit nil
### GetTotalSnapshots

`func (o *ProtectionInfo) GetTotalSnapshots() int64`

GetTotalSnapshots returns the TotalSnapshots field if non-nil, zero value otherwise.

### GetTotalSnapshotsOk

`func (o *ProtectionInfo) GetTotalSnapshotsOk() (*int64, bool)`

GetTotalSnapshotsOk returns a tuple with the TotalSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshots

`func (o *ProtectionInfo) SetTotalSnapshots(v int64)`

SetTotalSnapshots sets TotalSnapshots field to given value.

### HasTotalSnapshots

`func (o *ProtectionInfo) HasTotalSnapshots() bool`

HasTotalSnapshots returns a boolean if a field has been set.

### SetTotalSnapshotsNil

`func (o *ProtectionInfo) SetTotalSnapshotsNil(b bool)`

 SetTotalSnapshotsNil sets the value for TotalSnapshots to be an explicit nil

### UnsetTotalSnapshots
`func (o *ProtectionInfo) UnsetTotalSnapshots()`

UnsetTotalSnapshots ensures that no value is present for TotalSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


