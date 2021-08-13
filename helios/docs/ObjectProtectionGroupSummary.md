# ObjectProtectionGroupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the protection group name. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the protection group id. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the policy name for this group. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the policy id for this group. | [optional] 
**StorageDomainId** | Pointer to **NullableString** | Specifies the storage domain id of this group. Format is clusterId:clusterIncarnationId:storageDomainId. | [optional] 
**LastBackupRunStatus** | Pointer to **NullableString** | Specifies the status of last local back up run. | [optional] 
**LastArchivalRunStatus** | Pointer to **NullableString** | Specifies the status of last archival run. | [optional] 
**LastReplicationRunStatus** | Pointer to **NullableString** | Specifies the status of last replication run. | [optional] 
**LastRunSlaViolated** | Pointer to **NullableBool** | Specifies if the sla is violated in last run. | [optional] 

## Methods

### NewObjectProtectionGroupSummary

`func NewObjectProtectionGroupSummary() *ObjectProtectionGroupSummary`

NewObjectProtectionGroupSummary instantiates a new ObjectProtectionGroupSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProtectionGroupSummaryWithDefaults

`func NewObjectProtectionGroupSummaryWithDefaults() *ObjectProtectionGroupSummary`

NewObjectProtectionGroupSummaryWithDefaults instantiates a new ObjectProtectionGroupSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectProtectionGroupSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectProtectionGroupSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectProtectionGroupSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectProtectionGroupSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectProtectionGroupSummary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectProtectionGroupSummary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ObjectProtectionGroupSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectProtectionGroupSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectProtectionGroupSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectProtectionGroupSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectProtectionGroupSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectProtectionGroupSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPolicyName

`func (o *ObjectProtectionGroupSummary) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectProtectionGroupSummary) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectProtectionGroupSummary) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectProtectionGroupSummary) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectProtectionGroupSummary) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectProtectionGroupSummary) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyId

`func (o *ObjectProtectionGroupSummary) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectProtectionGroupSummary) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectProtectionGroupSummary) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectProtectionGroupSummary) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectProtectionGroupSummary) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectProtectionGroupSummary) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectProtectionGroupSummary) GetStorageDomainId() string`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectProtectionGroupSummary) GetStorageDomainIdOk() (*string, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectProtectionGroupSummary) SetStorageDomainId(v string)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectProtectionGroupSummary) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectProtectionGroupSummary) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectProtectionGroupSummary) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetLastBackupRunStatus

`func (o *ObjectProtectionGroupSummary) GetLastBackupRunStatus() string`

GetLastBackupRunStatus returns the LastBackupRunStatus field if non-nil, zero value otherwise.

### GetLastBackupRunStatusOk

`func (o *ObjectProtectionGroupSummary) GetLastBackupRunStatusOk() (*string, bool)`

GetLastBackupRunStatusOk returns a tuple with the LastBackupRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupRunStatus

`func (o *ObjectProtectionGroupSummary) SetLastBackupRunStatus(v string)`

SetLastBackupRunStatus sets LastBackupRunStatus field to given value.

### HasLastBackupRunStatus

`func (o *ObjectProtectionGroupSummary) HasLastBackupRunStatus() bool`

HasLastBackupRunStatus returns a boolean if a field has been set.

### SetLastBackupRunStatusNil

`func (o *ObjectProtectionGroupSummary) SetLastBackupRunStatusNil(b bool)`

 SetLastBackupRunStatusNil sets the value for LastBackupRunStatus to be an explicit nil

### UnsetLastBackupRunStatus
`func (o *ObjectProtectionGroupSummary) UnsetLastBackupRunStatus()`

UnsetLastBackupRunStatus ensures that no value is present for LastBackupRunStatus, not even an explicit nil
### GetLastArchivalRunStatus

`func (o *ObjectProtectionGroupSummary) GetLastArchivalRunStatus() string`

GetLastArchivalRunStatus returns the LastArchivalRunStatus field if non-nil, zero value otherwise.

### GetLastArchivalRunStatusOk

`func (o *ObjectProtectionGroupSummary) GetLastArchivalRunStatusOk() (*string, bool)`

GetLastArchivalRunStatusOk returns a tuple with the LastArchivalRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastArchivalRunStatus

`func (o *ObjectProtectionGroupSummary) SetLastArchivalRunStatus(v string)`

SetLastArchivalRunStatus sets LastArchivalRunStatus field to given value.

### HasLastArchivalRunStatus

`func (o *ObjectProtectionGroupSummary) HasLastArchivalRunStatus() bool`

HasLastArchivalRunStatus returns a boolean if a field has been set.

### SetLastArchivalRunStatusNil

`func (o *ObjectProtectionGroupSummary) SetLastArchivalRunStatusNil(b bool)`

 SetLastArchivalRunStatusNil sets the value for LastArchivalRunStatus to be an explicit nil

### UnsetLastArchivalRunStatus
`func (o *ObjectProtectionGroupSummary) UnsetLastArchivalRunStatus()`

UnsetLastArchivalRunStatus ensures that no value is present for LastArchivalRunStatus, not even an explicit nil
### GetLastReplicationRunStatus

`func (o *ObjectProtectionGroupSummary) GetLastReplicationRunStatus() string`

GetLastReplicationRunStatus returns the LastReplicationRunStatus field if non-nil, zero value otherwise.

### GetLastReplicationRunStatusOk

`func (o *ObjectProtectionGroupSummary) GetLastReplicationRunStatusOk() (*string, bool)`

GetLastReplicationRunStatusOk returns a tuple with the LastReplicationRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplicationRunStatus

`func (o *ObjectProtectionGroupSummary) SetLastReplicationRunStatus(v string)`

SetLastReplicationRunStatus sets LastReplicationRunStatus field to given value.

### HasLastReplicationRunStatus

`func (o *ObjectProtectionGroupSummary) HasLastReplicationRunStatus() bool`

HasLastReplicationRunStatus returns a boolean if a field has been set.

### SetLastReplicationRunStatusNil

`func (o *ObjectProtectionGroupSummary) SetLastReplicationRunStatusNil(b bool)`

 SetLastReplicationRunStatusNil sets the value for LastReplicationRunStatus to be an explicit nil

### UnsetLastReplicationRunStatus
`func (o *ObjectProtectionGroupSummary) UnsetLastReplicationRunStatus()`

UnsetLastReplicationRunStatus ensures that no value is present for LastReplicationRunStatus, not even an explicit nil
### GetLastRunSlaViolated

`func (o *ObjectProtectionGroupSummary) GetLastRunSlaViolated() bool`

GetLastRunSlaViolated returns the LastRunSlaViolated field if non-nil, zero value otherwise.

### GetLastRunSlaViolatedOk

`func (o *ObjectProtectionGroupSummary) GetLastRunSlaViolatedOk() (*bool, bool)`

GetLastRunSlaViolatedOk returns a tuple with the LastRunSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunSlaViolated

`func (o *ObjectProtectionGroupSummary) SetLastRunSlaViolated(v bool)`

SetLastRunSlaViolated sets LastRunSlaViolated field to given value.

### HasLastRunSlaViolated

`func (o *ObjectProtectionGroupSummary) HasLastRunSlaViolated() bool`

HasLastRunSlaViolated returns a boolean if a field has been set.

### SetLastRunSlaViolatedNil

`func (o *ObjectProtectionGroupSummary) SetLastRunSlaViolatedNil(b bool)`

 SetLastRunSlaViolatedNil sets the value for LastRunSlaViolated to be an explicit nil

### UnsetLastRunSlaViolated
`func (o *ObjectProtectionGroupSummary) UnsetLastRunSlaViolated()`

UnsetLastRunSlaViolated ensures that no value is present for LastRunSlaViolated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


