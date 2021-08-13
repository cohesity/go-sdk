# ObjectLastRunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** | Specifies the last run id. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the protection group name of last run. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id of last run. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the policy name of last run. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the policy id of last run. | [optional] 
**BackupRunStatus** | Pointer to **NullableString** | Specifies the status of last local back up run. | [optional] 
**ArchivalRunStatus** | Pointer to **NullableString** | Specifies the status of last archival run. | [optional] 
**ReplicationRunStatus** | Pointer to **NullableString** | Specifies the status of last replication run. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Specifies if the sla is violated in last run. | [optional] 

## Methods

### NewObjectLastRunAllOf

`func NewObjectLastRunAllOf() *ObjectLastRunAllOf`

NewObjectLastRunAllOf instantiates a new ObjectLastRunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectLastRunAllOfWithDefaults

`func NewObjectLastRunAllOfWithDefaults() *ObjectLastRunAllOf`

NewObjectLastRunAllOfWithDefaults instantiates a new ObjectLastRunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *ObjectLastRunAllOf) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ObjectLastRunAllOf) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ObjectLastRunAllOf) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ObjectLastRunAllOf) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetProtectionGroupName

`func (o *ObjectLastRunAllOf) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectLastRunAllOf) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectLastRunAllOf) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectLastRunAllOf) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectLastRunAllOf) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectLastRunAllOf) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectLastRunAllOf) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectLastRunAllOf) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectLastRunAllOf) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectLastRunAllOf) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectLastRunAllOf) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectLastRunAllOf) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetPolicyName

`func (o *ObjectLastRunAllOf) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectLastRunAllOf) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectLastRunAllOf) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectLastRunAllOf) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectLastRunAllOf) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectLastRunAllOf) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyId

`func (o *ObjectLastRunAllOf) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectLastRunAllOf) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectLastRunAllOf) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectLastRunAllOf) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectLastRunAllOf) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectLastRunAllOf) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetBackupRunStatus

`func (o *ObjectLastRunAllOf) GetBackupRunStatus() string`

GetBackupRunStatus returns the BackupRunStatus field if non-nil, zero value otherwise.

### GetBackupRunStatusOk

`func (o *ObjectLastRunAllOf) GetBackupRunStatusOk() (*string, bool)`

GetBackupRunStatusOk returns a tuple with the BackupRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRunStatus

`func (o *ObjectLastRunAllOf) SetBackupRunStatus(v string)`

SetBackupRunStatus sets BackupRunStatus field to given value.

### HasBackupRunStatus

`func (o *ObjectLastRunAllOf) HasBackupRunStatus() bool`

HasBackupRunStatus returns a boolean if a field has been set.

### SetBackupRunStatusNil

`func (o *ObjectLastRunAllOf) SetBackupRunStatusNil(b bool)`

 SetBackupRunStatusNil sets the value for BackupRunStatus to be an explicit nil

### UnsetBackupRunStatus
`func (o *ObjectLastRunAllOf) UnsetBackupRunStatus()`

UnsetBackupRunStatus ensures that no value is present for BackupRunStatus, not even an explicit nil
### GetArchivalRunStatus

`func (o *ObjectLastRunAllOf) GetArchivalRunStatus() string`

GetArchivalRunStatus returns the ArchivalRunStatus field if non-nil, zero value otherwise.

### GetArchivalRunStatusOk

`func (o *ObjectLastRunAllOf) GetArchivalRunStatusOk() (*string, bool)`

GetArchivalRunStatusOk returns a tuple with the ArchivalRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRunStatus

`func (o *ObjectLastRunAllOf) SetArchivalRunStatus(v string)`

SetArchivalRunStatus sets ArchivalRunStatus field to given value.

### HasArchivalRunStatus

`func (o *ObjectLastRunAllOf) HasArchivalRunStatus() bool`

HasArchivalRunStatus returns a boolean if a field has been set.

### SetArchivalRunStatusNil

`func (o *ObjectLastRunAllOf) SetArchivalRunStatusNil(b bool)`

 SetArchivalRunStatusNil sets the value for ArchivalRunStatus to be an explicit nil

### UnsetArchivalRunStatus
`func (o *ObjectLastRunAllOf) UnsetArchivalRunStatus()`

UnsetArchivalRunStatus ensures that no value is present for ArchivalRunStatus, not even an explicit nil
### GetReplicationRunStatus

`func (o *ObjectLastRunAllOf) GetReplicationRunStatus() string`

GetReplicationRunStatus returns the ReplicationRunStatus field if non-nil, zero value otherwise.

### GetReplicationRunStatusOk

`func (o *ObjectLastRunAllOf) GetReplicationRunStatusOk() (*string, bool)`

GetReplicationRunStatusOk returns a tuple with the ReplicationRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRunStatus

`func (o *ObjectLastRunAllOf) SetReplicationRunStatus(v string)`

SetReplicationRunStatus sets ReplicationRunStatus field to given value.

### HasReplicationRunStatus

`func (o *ObjectLastRunAllOf) HasReplicationRunStatus() bool`

HasReplicationRunStatus returns a boolean if a field has been set.

### SetReplicationRunStatusNil

`func (o *ObjectLastRunAllOf) SetReplicationRunStatusNil(b bool)`

 SetReplicationRunStatusNil sets the value for ReplicationRunStatus to be an explicit nil

### UnsetReplicationRunStatus
`func (o *ObjectLastRunAllOf) UnsetReplicationRunStatus()`

UnsetReplicationRunStatus ensures that no value is present for ReplicationRunStatus, not even an explicit nil
### GetIsSlaViolated

`func (o *ObjectLastRunAllOf) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ObjectLastRunAllOf) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ObjectLastRunAllOf) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ObjectLastRunAllOf) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ObjectLastRunAllOf) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ObjectLastRunAllOf) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


