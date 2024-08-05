# CommonTdmCloneTaskRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment of the TDM Clone task. | 
**PolicyId** | Pointer to **NullableString** | Specifies the ID of the policy, which should be used to protect this clone. This is useful for automatic snapshots. This must be specified if either of protectionGroupId and protectionGroupName is specified. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the ID of an existing protection group, which should start protecting this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupName should not be specified. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the name of a new protection group, which should be created to protect this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupId should not be specified. | [optional] 
**PointInTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp (in usecs from epoch) for creating the clone at a point in time in the past. | [optional] 
**SnapshotId** | **NullableString** | Specifies the snapshot ID, from which the clone is to be created. | 
**TargetHostId** | **NullableInt64** | Specifies the ID of the host, where the clone needs to be created. | 

## Methods

### NewCommonTdmCloneTaskRequestParams

`func NewCommonTdmCloneTaskRequestParams(environment NullableString, snapshotId NullableString, targetHostId NullableInt64, ) *CommonTdmCloneTaskRequestParams`

NewCommonTdmCloneTaskRequestParams instantiates a new CommonTdmCloneTaskRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmCloneTaskRequestParamsWithDefaults

`func NewCommonTdmCloneTaskRequestParamsWithDefaults() *CommonTdmCloneTaskRequestParams`

NewCommonTdmCloneTaskRequestParamsWithDefaults instantiates a new CommonTdmCloneTaskRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonTdmCloneTaskRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonTdmCloneTaskRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonTdmCloneTaskRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonTdmCloneTaskRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonTdmCloneTaskRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetPolicyId

`func (o *CommonTdmCloneTaskRequestParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonTdmCloneTaskRequestParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonTdmCloneTaskRequestParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CommonTdmCloneTaskRequestParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CommonTdmCloneTaskRequestParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonTdmCloneTaskRequestParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetProtectionGroupId

`func (o *CommonTdmCloneTaskRequestParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CommonTdmCloneTaskRequestParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CommonTdmCloneTaskRequestParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CommonTdmCloneTaskRequestParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CommonTdmCloneTaskRequestParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CommonTdmCloneTaskRequestParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *CommonTdmCloneTaskRequestParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *CommonTdmCloneTaskRequestParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *CommonTdmCloneTaskRequestParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *CommonTdmCloneTaskRequestParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *CommonTdmCloneTaskRequestParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *CommonTdmCloneTaskRequestParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParams) GetPointInTimeUsecs() int64`

GetPointInTimeUsecs returns the PointInTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeUsecsOk

`func (o *CommonTdmCloneTaskRequestParams) GetPointInTimeUsecsOk() (*int64, bool)`

GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParams) SetPointInTimeUsecs(v int64)`

SetPointInTimeUsecs sets PointInTimeUsecs field to given value.

### HasPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParams) HasPointInTimeUsecs() bool`

HasPointInTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeUsecsNil

`func (o *CommonTdmCloneTaskRequestParams) SetPointInTimeUsecsNil(b bool)`

 SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil

### UnsetPointInTimeUsecs
`func (o *CommonTdmCloneTaskRequestParams) UnsetPointInTimeUsecs()`

UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil
### GetSnapshotId

`func (o *CommonTdmCloneTaskRequestParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CommonTdmCloneTaskRequestParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CommonTdmCloneTaskRequestParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### SetSnapshotIdNil

`func (o *CommonTdmCloneTaskRequestParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *CommonTdmCloneTaskRequestParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetTargetHostId

`func (o *CommonTdmCloneTaskRequestParams) GetTargetHostId() int64`

GetTargetHostId returns the TargetHostId field if non-nil, zero value otherwise.

### GetTargetHostIdOk

`func (o *CommonTdmCloneTaskRequestParams) GetTargetHostIdOk() (*int64, bool)`

GetTargetHostIdOk returns a tuple with the TargetHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostId

`func (o *CommonTdmCloneTaskRequestParams) SetTargetHostId(v int64)`

SetTargetHostId sets TargetHostId field to given value.


### SetTargetHostIdNil

`func (o *CommonTdmCloneTaskRequestParams) SetTargetHostIdNil(b bool)`

 SetTargetHostIdNil sets the value for TargetHostId to be an explicit nil

### UnsetTargetHostId
`func (o *CommonTdmCloneTaskRequestParams) UnsetTargetHostId()`

UnsetTargetHostId ensures that no value is present for TargetHostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


