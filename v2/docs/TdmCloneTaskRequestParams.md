# TdmCloneTaskRequestParams

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
**OracleParams** | Pointer to [**OracleCloneTask**](OracleCloneTask.md) |  | [optional] 

## Methods

### NewTdmCloneTaskRequestParams

`func NewTdmCloneTaskRequestParams(environment NullableString, snapshotId NullableString, targetHostId NullableInt64, ) *TdmCloneTaskRequestParams`

NewTdmCloneTaskRequestParams instantiates a new TdmCloneTaskRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmCloneTaskRequestParamsWithDefaults

`func NewTdmCloneTaskRequestParamsWithDefaults() *TdmCloneTaskRequestParams`

NewTdmCloneTaskRequestParamsWithDefaults instantiates a new TdmCloneTaskRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *TdmCloneTaskRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TdmCloneTaskRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TdmCloneTaskRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *TdmCloneTaskRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TdmCloneTaskRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetPolicyId

`func (o *TdmCloneTaskRequestParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *TdmCloneTaskRequestParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *TdmCloneTaskRequestParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *TdmCloneTaskRequestParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *TdmCloneTaskRequestParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *TdmCloneTaskRequestParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetProtectionGroupId

`func (o *TdmCloneTaskRequestParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *TdmCloneTaskRequestParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *TdmCloneTaskRequestParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *TdmCloneTaskRequestParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *TdmCloneTaskRequestParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *TdmCloneTaskRequestParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *TdmCloneTaskRequestParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *TdmCloneTaskRequestParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *TdmCloneTaskRequestParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *TdmCloneTaskRequestParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *TdmCloneTaskRequestParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *TdmCloneTaskRequestParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetPointInTimeUsecs

`func (o *TdmCloneTaskRequestParams) GetPointInTimeUsecs() int64`

GetPointInTimeUsecs returns the PointInTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeUsecsOk

`func (o *TdmCloneTaskRequestParams) GetPointInTimeUsecsOk() (*int64, bool)`

GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUsecs

`func (o *TdmCloneTaskRequestParams) SetPointInTimeUsecs(v int64)`

SetPointInTimeUsecs sets PointInTimeUsecs field to given value.

### HasPointInTimeUsecs

`func (o *TdmCloneTaskRequestParams) HasPointInTimeUsecs() bool`

HasPointInTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeUsecsNil

`func (o *TdmCloneTaskRequestParams) SetPointInTimeUsecsNil(b bool)`

 SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil

### UnsetPointInTimeUsecs
`func (o *TdmCloneTaskRequestParams) UnsetPointInTimeUsecs()`

UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil
### GetSnapshotId

`func (o *TdmCloneTaskRequestParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TdmCloneTaskRequestParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TdmCloneTaskRequestParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### SetSnapshotIdNil

`func (o *TdmCloneTaskRequestParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *TdmCloneTaskRequestParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetTargetHostId

`func (o *TdmCloneTaskRequestParams) GetTargetHostId() int64`

GetTargetHostId returns the TargetHostId field if non-nil, zero value otherwise.

### GetTargetHostIdOk

`func (o *TdmCloneTaskRequestParams) GetTargetHostIdOk() (*int64, bool)`

GetTargetHostIdOk returns a tuple with the TargetHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostId

`func (o *TdmCloneTaskRequestParams) SetTargetHostId(v int64)`

SetTargetHostId sets TargetHostId field to given value.


### SetTargetHostIdNil

`func (o *TdmCloneTaskRequestParams) SetTargetHostIdNil(b bool)`

 SetTargetHostIdNil sets the value for TargetHostId to be an explicit nil

### UnsetTargetHostId
`func (o *TdmCloneTaskRequestParams) UnsetTargetHostId()`

UnsetTargetHostId ensures that no value is present for TargetHostId, not even an explicit nil
### GetOracleParams

`func (o *TdmCloneTaskRequestParams) GetOracleParams() OracleCloneTask`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *TdmCloneTaskRequestParams) GetOracleParamsOk() (*OracleCloneTask, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *TdmCloneTaskRequestParams) SetOracleParams(v OracleCloneTask)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *TdmCloneTaskRequestParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


