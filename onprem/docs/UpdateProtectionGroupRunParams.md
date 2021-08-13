# UpdateProtectionGroupRunParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **NullableString** | Specifies a unique Protection Group Run id. | 
**LocalSnapshotConfig** | Pointer to [**UpdateLocalSnapshotConfig**](UpdateLocalSnapshotConfig.md) |  | [optional] 
**ReplicationSnapshotConfig** | Pointer to [**UpdateReplicationSnapshotConfig**](UpdateReplicationSnapshotConfig.md) |  | [optional] 
**ArchivalSnapshotConfig** | Pointer to [**UpdateArchivalSnapshotConfig**](UpdateArchivalSnapshotConfig.md) |  | [optional] 

## Methods

### NewUpdateProtectionGroupRunParams

`func NewUpdateProtectionGroupRunParams(runId NullableString, ) *UpdateProtectionGroupRunParams`

NewUpdateProtectionGroupRunParams instantiates a new UpdateProtectionGroupRunParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionGroupRunParamsWithDefaults

`func NewUpdateProtectionGroupRunParamsWithDefaults() *UpdateProtectionGroupRunParams`

NewUpdateProtectionGroupRunParamsWithDefaults instantiates a new UpdateProtectionGroupRunParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *UpdateProtectionGroupRunParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *UpdateProtectionGroupRunParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *UpdateProtectionGroupRunParams) SetRunId(v string)`

SetRunId sets RunId field to given value.


### SetRunIdNil

`func (o *UpdateProtectionGroupRunParams) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *UpdateProtectionGroupRunParams) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetLocalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) GetLocalSnapshotConfig() UpdateLocalSnapshotConfig`

GetLocalSnapshotConfig returns the LocalSnapshotConfig field if non-nil, zero value otherwise.

### GetLocalSnapshotConfigOk

`func (o *UpdateProtectionGroupRunParams) GetLocalSnapshotConfigOk() (*UpdateLocalSnapshotConfig, bool)`

GetLocalSnapshotConfigOk returns a tuple with the LocalSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) SetLocalSnapshotConfig(v UpdateLocalSnapshotConfig)`

SetLocalSnapshotConfig sets LocalSnapshotConfig field to given value.

### HasLocalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) HasLocalSnapshotConfig() bool`

HasLocalSnapshotConfig returns a boolean if a field has been set.

### GetReplicationSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) GetReplicationSnapshotConfig() UpdateReplicationSnapshotConfig`

GetReplicationSnapshotConfig returns the ReplicationSnapshotConfig field if non-nil, zero value otherwise.

### GetReplicationSnapshotConfigOk

`func (o *UpdateProtectionGroupRunParams) GetReplicationSnapshotConfigOk() (*UpdateReplicationSnapshotConfig, bool)`

GetReplicationSnapshotConfigOk returns a tuple with the ReplicationSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) SetReplicationSnapshotConfig(v UpdateReplicationSnapshotConfig)`

SetReplicationSnapshotConfig sets ReplicationSnapshotConfig field to given value.

### HasReplicationSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) HasReplicationSnapshotConfig() bool`

HasReplicationSnapshotConfig returns a boolean if a field has been set.

### GetArchivalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) GetArchivalSnapshotConfig() UpdateArchivalSnapshotConfig`

GetArchivalSnapshotConfig returns the ArchivalSnapshotConfig field if non-nil, zero value otherwise.

### GetArchivalSnapshotConfigOk

`func (o *UpdateProtectionGroupRunParams) GetArchivalSnapshotConfigOk() (*UpdateArchivalSnapshotConfig, bool)`

GetArchivalSnapshotConfigOk returns a tuple with the ArchivalSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) SetArchivalSnapshotConfig(v UpdateArchivalSnapshotConfig)`

SetArchivalSnapshotConfig sets ArchivalSnapshotConfig field to given value.

### HasArchivalSnapshotConfig

`func (o *UpdateProtectionGroupRunParams) HasArchivalSnapshotConfig() bool`

HasArchivalSnapshotConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


