# RecoverVmwareChildSnapshotParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantRecoveryInfo** | Pointer to [**NullableRecoveryTaskInfo**](RecoveryTaskInfo.md) | Specifies the info about instant recovery. This is only applicable for RecoverVm. | [optional] 
**DatastoreMigrationInfo** | Pointer to [**NullableRecoveryTaskInfo**](RecoveryTaskInfo.md) | Specifies the info about datastore migration. This is only applicable for RecoverVm. | [optional] 

## Methods

### NewRecoverVmwareChildSnapshotParamsAllOf

`func NewRecoverVmwareChildSnapshotParamsAllOf() *RecoverVmwareChildSnapshotParamsAllOf`

NewRecoverVmwareChildSnapshotParamsAllOf instantiates a new RecoverVmwareChildSnapshotParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareChildSnapshotParamsAllOfWithDefaults

`func NewRecoverVmwareChildSnapshotParamsAllOfWithDefaults() *RecoverVmwareChildSnapshotParamsAllOf`

NewRecoverVmwareChildSnapshotParamsAllOfWithDefaults instantiates a new RecoverVmwareChildSnapshotParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantRecoveryInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) GetInstantRecoveryInfo() RecoveryTaskInfo`

GetInstantRecoveryInfo returns the InstantRecoveryInfo field if non-nil, zero value otherwise.

### GetInstantRecoveryInfoOk

`func (o *RecoverVmwareChildSnapshotParamsAllOf) GetInstantRecoveryInfoOk() (*RecoveryTaskInfo, bool)`

GetInstantRecoveryInfoOk returns a tuple with the InstantRecoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRecoveryInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) SetInstantRecoveryInfo(v RecoveryTaskInfo)`

SetInstantRecoveryInfo sets InstantRecoveryInfo field to given value.

### HasInstantRecoveryInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) HasInstantRecoveryInfo() bool`

HasInstantRecoveryInfo returns a boolean if a field has been set.

### SetInstantRecoveryInfoNil

`func (o *RecoverVmwareChildSnapshotParamsAllOf) SetInstantRecoveryInfoNil(b bool)`

 SetInstantRecoveryInfoNil sets the value for InstantRecoveryInfo to be an explicit nil

### UnsetInstantRecoveryInfo
`func (o *RecoverVmwareChildSnapshotParamsAllOf) UnsetInstantRecoveryInfo()`

UnsetInstantRecoveryInfo ensures that no value is present for InstantRecoveryInfo, not even an explicit nil
### GetDatastoreMigrationInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) GetDatastoreMigrationInfo() RecoveryTaskInfo`

GetDatastoreMigrationInfo returns the DatastoreMigrationInfo field if non-nil, zero value otherwise.

### GetDatastoreMigrationInfoOk

`func (o *RecoverVmwareChildSnapshotParamsAllOf) GetDatastoreMigrationInfoOk() (*RecoveryTaskInfo, bool)`

GetDatastoreMigrationInfoOk returns a tuple with the DatastoreMigrationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreMigrationInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) SetDatastoreMigrationInfo(v RecoveryTaskInfo)`

SetDatastoreMigrationInfo sets DatastoreMigrationInfo field to given value.

### HasDatastoreMigrationInfo

`func (o *RecoverVmwareChildSnapshotParamsAllOf) HasDatastoreMigrationInfo() bool`

HasDatastoreMigrationInfo returns a boolean if a field has been set.

### SetDatastoreMigrationInfoNil

`func (o *RecoverVmwareChildSnapshotParamsAllOf) SetDatastoreMigrationInfoNil(b bool)`

 SetDatastoreMigrationInfoNil sets the value for DatastoreMigrationInfo to be an explicit nil

### UnsetDatastoreMigrationInfo
`func (o *RecoverVmwareChildSnapshotParamsAllOf) UnsetDatastoreMigrationInfo()`

UnsetDatastoreMigrationInfo ensures that no value is present for DatastoreMigrationInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


