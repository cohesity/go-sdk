# ProtectionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **NullableString** | Specifies the policy name for this group. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the policy id for this protection. | [optional] 
**StorageDomainId** | Pointer to **NullableString** | Specifies the storage domain id of this protection. Format is clusterId:clusterIncarnationId:storageDomainId. | [optional] 
**LastBackupRunStatus** | Pointer to **NullableString** | Specifies the status of last local back up run. | [optional] 
**LastArchivalRunStatus** | Pointer to **NullableString** | Specifies the status of last archival run. | [optional] 
**LastReplicationRunStatus** | Pointer to **NullableString** | Specifies the status of last replication run. | [optional] 
**LastRunSlaViolated** | Pointer to **NullableBool** | Specifies if the sla is violated in last run. | [optional] 

## Methods

### NewProtectionSummary

`func NewProtectionSummary() *ProtectionSummary`

NewProtectionSummary instantiates a new ProtectionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSummaryWithDefaults

`func NewProtectionSummaryWithDefaults() *ProtectionSummary`

NewProtectionSummaryWithDefaults instantiates a new ProtectionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *ProtectionSummary) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ProtectionSummary) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ProtectionSummary) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ProtectionSummary) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ProtectionSummary) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ProtectionSummary) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyId

`func (o *ProtectionSummary) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionSummary) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionSummary) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ProtectionSummary) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ProtectionSummary) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionSummary) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetStorageDomainId

`func (o *ProtectionSummary) GetStorageDomainId() string`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ProtectionSummary) GetStorageDomainIdOk() (*string, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ProtectionSummary) SetStorageDomainId(v string)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ProtectionSummary) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ProtectionSummary) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ProtectionSummary) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetLastBackupRunStatus

`func (o *ProtectionSummary) GetLastBackupRunStatus() string`

GetLastBackupRunStatus returns the LastBackupRunStatus field if non-nil, zero value otherwise.

### GetLastBackupRunStatusOk

`func (o *ProtectionSummary) GetLastBackupRunStatusOk() (*string, bool)`

GetLastBackupRunStatusOk returns a tuple with the LastBackupRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupRunStatus

`func (o *ProtectionSummary) SetLastBackupRunStatus(v string)`

SetLastBackupRunStatus sets LastBackupRunStatus field to given value.

### HasLastBackupRunStatus

`func (o *ProtectionSummary) HasLastBackupRunStatus() bool`

HasLastBackupRunStatus returns a boolean if a field has been set.

### SetLastBackupRunStatusNil

`func (o *ProtectionSummary) SetLastBackupRunStatusNil(b bool)`

 SetLastBackupRunStatusNil sets the value for LastBackupRunStatus to be an explicit nil

### UnsetLastBackupRunStatus
`func (o *ProtectionSummary) UnsetLastBackupRunStatus()`

UnsetLastBackupRunStatus ensures that no value is present for LastBackupRunStatus, not even an explicit nil
### GetLastArchivalRunStatus

`func (o *ProtectionSummary) GetLastArchivalRunStatus() string`

GetLastArchivalRunStatus returns the LastArchivalRunStatus field if non-nil, zero value otherwise.

### GetLastArchivalRunStatusOk

`func (o *ProtectionSummary) GetLastArchivalRunStatusOk() (*string, bool)`

GetLastArchivalRunStatusOk returns a tuple with the LastArchivalRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastArchivalRunStatus

`func (o *ProtectionSummary) SetLastArchivalRunStatus(v string)`

SetLastArchivalRunStatus sets LastArchivalRunStatus field to given value.

### HasLastArchivalRunStatus

`func (o *ProtectionSummary) HasLastArchivalRunStatus() bool`

HasLastArchivalRunStatus returns a boolean if a field has been set.

### SetLastArchivalRunStatusNil

`func (o *ProtectionSummary) SetLastArchivalRunStatusNil(b bool)`

 SetLastArchivalRunStatusNil sets the value for LastArchivalRunStatus to be an explicit nil

### UnsetLastArchivalRunStatus
`func (o *ProtectionSummary) UnsetLastArchivalRunStatus()`

UnsetLastArchivalRunStatus ensures that no value is present for LastArchivalRunStatus, not even an explicit nil
### GetLastReplicationRunStatus

`func (o *ProtectionSummary) GetLastReplicationRunStatus() string`

GetLastReplicationRunStatus returns the LastReplicationRunStatus field if non-nil, zero value otherwise.

### GetLastReplicationRunStatusOk

`func (o *ProtectionSummary) GetLastReplicationRunStatusOk() (*string, bool)`

GetLastReplicationRunStatusOk returns a tuple with the LastReplicationRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplicationRunStatus

`func (o *ProtectionSummary) SetLastReplicationRunStatus(v string)`

SetLastReplicationRunStatus sets LastReplicationRunStatus field to given value.

### HasLastReplicationRunStatus

`func (o *ProtectionSummary) HasLastReplicationRunStatus() bool`

HasLastReplicationRunStatus returns a boolean if a field has been set.

### SetLastReplicationRunStatusNil

`func (o *ProtectionSummary) SetLastReplicationRunStatusNil(b bool)`

 SetLastReplicationRunStatusNil sets the value for LastReplicationRunStatus to be an explicit nil

### UnsetLastReplicationRunStatus
`func (o *ProtectionSummary) UnsetLastReplicationRunStatus()`

UnsetLastReplicationRunStatus ensures that no value is present for LastReplicationRunStatus, not even an explicit nil
### GetLastRunSlaViolated

`func (o *ProtectionSummary) GetLastRunSlaViolated() bool`

GetLastRunSlaViolated returns the LastRunSlaViolated field if non-nil, zero value otherwise.

### GetLastRunSlaViolatedOk

`func (o *ProtectionSummary) GetLastRunSlaViolatedOk() (*bool, bool)`

GetLastRunSlaViolatedOk returns a tuple with the LastRunSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunSlaViolated

`func (o *ProtectionSummary) SetLastRunSlaViolated(v bool)`

SetLastRunSlaViolated sets LastRunSlaViolated field to given value.

### HasLastRunSlaViolated

`func (o *ProtectionSummary) HasLastRunSlaViolated() bool`

HasLastRunSlaViolated returns a boolean if a field has been set.

### SetLastRunSlaViolatedNil

`func (o *ProtectionSummary) SetLastRunSlaViolatedNil(b bool)`

 SetLastRunSlaViolatedNil sets the value for LastRunSlaViolated to be an explicit nil

### UnsetLastRunSlaViolated
`func (o *ProtectionSummary) UnsetLastRunSlaViolated()`

UnsetLastRunSlaViolated ensures that no value is present for LastRunSlaViolated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


