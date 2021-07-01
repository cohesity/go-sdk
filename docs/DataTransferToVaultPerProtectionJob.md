# DataTransferToVaultPerProtectionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumLogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of logical bytes that are transferred from this Cohesity Cluster to this Vault for this Protection Job. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumPhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of physical bytes that are transferred from this Cohesity Cluster to this Vault for this Protection Job. | [optional] 
**ProtectionJobName** | Pointer to **NullableString** | Specifies the name of the Protection Job. | [optional] 
**StorageConsumed** | Pointer to **NullableInt64** | Specifies the total number of storage bytes consumed that are transferred from this Cohesity Cluster to this vault for this Protection Job. | [optional] 

## Methods

### NewDataTransferToVaultPerProtectionJob

`func NewDataTransferToVaultPerProtectionJob() *DataTransferToVaultPerProtectionJob`

NewDataTransferToVaultPerProtectionJob instantiates a new DataTransferToVaultPerProtectionJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferToVaultPerProtectionJobWithDefaults

`func NewDataTransferToVaultPerProtectionJobWithDefaults() *DataTransferToVaultPerProtectionJob`

NewDataTransferToVaultPerProtectionJobWithDefaults instantiates a new DataTransferToVaultPerProtectionJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumLogicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) GetNumLogicalBytesTransferred() int64`

GetNumLogicalBytesTransferred returns the NumLogicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumLogicalBytesTransferredOk

`func (o *DataTransferToVaultPerProtectionJob) GetNumLogicalBytesTransferredOk() (*int64, bool)`

GetNumLogicalBytesTransferredOk returns a tuple with the NumLogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) SetNumLogicalBytesTransferred(v int64)`

SetNumLogicalBytesTransferred sets NumLogicalBytesTransferred field to given value.

### HasNumLogicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) HasNumLogicalBytesTransferred() bool`

HasNumLogicalBytesTransferred returns a boolean if a field has been set.

### SetNumLogicalBytesTransferredNil

`func (o *DataTransferToVaultPerProtectionJob) SetNumLogicalBytesTransferredNil(b bool)`

 SetNumLogicalBytesTransferredNil sets the value for NumLogicalBytesTransferred to be an explicit nil

### UnsetNumLogicalBytesTransferred
`func (o *DataTransferToVaultPerProtectionJob) UnsetNumLogicalBytesTransferred()`

UnsetNumLogicalBytesTransferred ensures that no value is present for NumLogicalBytesTransferred, not even an explicit nil
### GetNumPhysicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) GetNumPhysicalBytesTransferred() int64`

GetNumPhysicalBytesTransferred returns the NumPhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumPhysicalBytesTransferredOk

`func (o *DataTransferToVaultPerProtectionJob) GetNumPhysicalBytesTransferredOk() (*int64, bool)`

GetNumPhysicalBytesTransferredOk returns a tuple with the NumPhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) SetNumPhysicalBytesTransferred(v int64)`

SetNumPhysicalBytesTransferred sets NumPhysicalBytesTransferred field to given value.

### HasNumPhysicalBytesTransferred

`func (o *DataTransferToVaultPerProtectionJob) HasNumPhysicalBytesTransferred() bool`

HasNumPhysicalBytesTransferred returns a boolean if a field has been set.

### SetNumPhysicalBytesTransferredNil

`func (o *DataTransferToVaultPerProtectionJob) SetNumPhysicalBytesTransferredNil(b bool)`

 SetNumPhysicalBytesTransferredNil sets the value for NumPhysicalBytesTransferred to be an explicit nil

### UnsetNumPhysicalBytesTransferred
`func (o *DataTransferToVaultPerProtectionJob) UnsetNumPhysicalBytesTransferred()`

UnsetNumPhysicalBytesTransferred ensures that no value is present for NumPhysicalBytesTransferred, not even an explicit nil
### GetProtectionJobName

`func (o *DataTransferToVaultPerProtectionJob) GetProtectionJobName() string`

GetProtectionJobName returns the ProtectionJobName field if non-nil, zero value otherwise.

### GetProtectionJobNameOk

`func (o *DataTransferToVaultPerProtectionJob) GetProtectionJobNameOk() (*string, bool)`

GetProtectionJobNameOk returns a tuple with the ProtectionJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobName

`func (o *DataTransferToVaultPerProtectionJob) SetProtectionJobName(v string)`

SetProtectionJobName sets ProtectionJobName field to given value.

### HasProtectionJobName

`func (o *DataTransferToVaultPerProtectionJob) HasProtectionJobName() bool`

HasProtectionJobName returns a boolean if a field has been set.

### SetProtectionJobNameNil

`func (o *DataTransferToVaultPerProtectionJob) SetProtectionJobNameNil(b bool)`

 SetProtectionJobNameNil sets the value for ProtectionJobName to be an explicit nil

### UnsetProtectionJobName
`func (o *DataTransferToVaultPerProtectionJob) UnsetProtectionJobName()`

UnsetProtectionJobName ensures that no value is present for ProtectionJobName, not even an explicit nil
### GetStorageConsumed

`func (o *DataTransferToVaultPerProtectionJob) GetStorageConsumed() int64`

GetStorageConsumed returns the StorageConsumed field if non-nil, zero value otherwise.

### GetStorageConsumedOk

`func (o *DataTransferToVaultPerProtectionJob) GetStorageConsumedOk() (*int64, bool)`

GetStorageConsumedOk returns a tuple with the StorageConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumed

`func (o *DataTransferToVaultPerProtectionJob) SetStorageConsumed(v int64)`

SetStorageConsumed sets StorageConsumed field to given value.

### HasStorageConsumed

`func (o *DataTransferToVaultPerProtectionJob) HasStorageConsumed() bool`

HasStorageConsumed returns a boolean if a field has been set.

### SetStorageConsumedNil

`func (o *DataTransferToVaultPerProtectionJob) SetStorageConsumedNil(b bool)`

 SetStorageConsumedNil sets the value for StorageConsumed to be an explicit nil

### UnsetStorageConsumed
`func (o *DataTransferToVaultPerProtectionJob) UnsetStorageConsumed()`

UnsetStorageConsumed ensures that no value is present for StorageConsumed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


