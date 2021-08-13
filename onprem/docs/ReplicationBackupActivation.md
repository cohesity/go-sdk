# ReplicationBackupActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]FailoverObject**](FailoverObject.md) | Specifies the list of failover object that need to be protected on replication cluster. If the object set that was sent earlier is provided again then API will return an error. If this objects list is not specified then internally it will be inferred if &#39;/objectLinkage&#39; API has been called previously. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id that will be used for backing up the failover entities on replication cluster. This is a optional argument and only need to be passed if user wants to use the existing job for the backup. If specified then Orchastrator should enusre that protection group is compatible to handle all provided failover objects. | [optional] 
**EnableReverseReplication** | Pointer to **NullableBool** | If this is specifed as true, then reverse replication of failover objects will be enabled from replication cluster to source cluster. If source cluster is not reachable, then replications will fail until source cluster comes up again. Here orchastrator should also ensure that storage domain on replication cluster is correctly mapped to the same storage domain on the source cluster. | [optional] 

## Methods

### NewReplicationBackupActivation

`func NewReplicationBackupActivation() *ReplicationBackupActivation`

NewReplicationBackupActivation instantiates a new ReplicationBackupActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationBackupActivationWithDefaults

`func NewReplicationBackupActivationWithDefaults() *ReplicationBackupActivation`

NewReplicationBackupActivationWithDefaults instantiates a new ReplicationBackupActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *ReplicationBackupActivation) GetObjects() []FailoverObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ReplicationBackupActivation) GetObjectsOk() (*[]FailoverObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ReplicationBackupActivation) SetObjects(v []FailoverObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ReplicationBackupActivation) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ReplicationBackupActivation) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ReplicationBackupActivation) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetProtectionGroupId

`func (o *ReplicationBackupActivation) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ReplicationBackupActivation) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ReplicationBackupActivation) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ReplicationBackupActivation) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ReplicationBackupActivation) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ReplicationBackupActivation) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetEnableReverseReplication

`func (o *ReplicationBackupActivation) GetEnableReverseReplication() bool`

GetEnableReverseReplication returns the EnableReverseReplication field if non-nil, zero value otherwise.

### GetEnableReverseReplicationOk

`func (o *ReplicationBackupActivation) GetEnableReverseReplicationOk() (*bool, bool)`

GetEnableReverseReplicationOk returns a tuple with the EnableReverseReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReverseReplication

`func (o *ReplicationBackupActivation) SetEnableReverseReplication(v bool)`

SetEnableReverseReplication sets EnableReverseReplication field to given value.

### HasEnableReverseReplication

`func (o *ReplicationBackupActivation) HasEnableReverseReplication() bool`

HasEnableReverseReplication returns a boolean if a field has been set.

### SetEnableReverseReplicationNil

`func (o *ReplicationBackupActivation) SetEnableReverseReplicationNil(b bool)`

 SetEnableReverseReplicationNil sets the value for EnableReverseReplication to be an explicit nil

### UnsetEnableReverseReplication
`func (o *ReplicationBackupActivation) UnsetEnableReverseReplication()`

UnsetEnableReverseReplication ensures that no value is present for EnableReverseReplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


