# ReplicationBackupActivationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id that will be returned upon creation of new group or existing group for backing up failover entities. | [optional] 
**ReverseReplicationResult** | Pointer to [**ReverseReplicationResult**](ReverseReplicationResult.md) |  | [optional] 
**Objects** | Pointer to [**[]FailoverObject**](FailoverObject.md) | Specifies the list of failover object that are going to be protected on replication cluster. | [optional] 

## Methods

### NewReplicationBackupActivationResult

`func NewReplicationBackupActivationResult() *ReplicationBackupActivationResult`

NewReplicationBackupActivationResult instantiates a new ReplicationBackupActivationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationBackupActivationResultWithDefaults

`func NewReplicationBackupActivationResultWithDefaults() *ReplicationBackupActivationResult`

NewReplicationBackupActivationResultWithDefaults instantiates a new ReplicationBackupActivationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupId

`func (o *ReplicationBackupActivationResult) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ReplicationBackupActivationResult) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ReplicationBackupActivationResult) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ReplicationBackupActivationResult) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ReplicationBackupActivationResult) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ReplicationBackupActivationResult) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetReverseReplicationResult

`func (o *ReplicationBackupActivationResult) GetReverseReplicationResult() ReverseReplicationResult`

GetReverseReplicationResult returns the ReverseReplicationResult field if non-nil, zero value otherwise.

### GetReverseReplicationResultOk

`func (o *ReplicationBackupActivationResult) GetReverseReplicationResultOk() (*ReverseReplicationResult, bool)`

GetReverseReplicationResultOk returns a tuple with the ReverseReplicationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseReplicationResult

`func (o *ReplicationBackupActivationResult) SetReverseReplicationResult(v ReverseReplicationResult)`

SetReverseReplicationResult sets ReverseReplicationResult field to given value.

### HasReverseReplicationResult

`func (o *ReplicationBackupActivationResult) HasReverseReplicationResult() bool`

HasReverseReplicationResult returns a boolean if a field has been set.

### GetObjects

`func (o *ReplicationBackupActivationResult) GetObjects() []FailoverObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ReplicationBackupActivationResult) GetObjectsOk() (*[]FailoverObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ReplicationBackupActivationResult) SetObjects(v []FailoverObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ReplicationBackupActivationResult) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ReplicationBackupActivationResult) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ReplicationBackupActivationResult) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


