# UpdateReplicationSnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSnapshotConfig** | Pointer to [**[]RunReplicationConfig**](RunReplicationConfig.md) | Specifies the new configuration about adding Replication Snapshot to existing Protection Group Run. | [optional] 
**UpdateExistingSnapshotConfig** | Pointer to [**[]UpdateExistingReplicationSnapshotConfig**](UpdateExistingReplicationSnapshotConfig.md) | Specifies the configuration about updating an existing Replication Snapshot Run. | [optional] 

## Methods

### NewUpdateReplicationSnapshotConfig

`func NewUpdateReplicationSnapshotConfig() *UpdateReplicationSnapshotConfig`

NewUpdateReplicationSnapshotConfig instantiates a new UpdateReplicationSnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReplicationSnapshotConfigWithDefaults

`func NewUpdateReplicationSnapshotConfigWithDefaults() *UpdateReplicationSnapshotConfig`

NewUpdateReplicationSnapshotConfigWithDefaults instantiates a new UpdateReplicationSnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) GetNewSnapshotConfig() []RunReplicationConfig`

GetNewSnapshotConfig returns the NewSnapshotConfig field if non-nil, zero value otherwise.

### GetNewSnapshotConfigOk

`func (o *UpdateReplicationSnapshotConfig) GetNewSnapshotConfigOk() (*[]RunReplicationConfig, bool)`

GetNewSnapshotConfigOk returns a tuple with the NewSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) SetNewSnapshotConfig(v []RunReplicationConfig)`

SetNewSnapshotConfig sets NewSnapshotConfig field to given value.

### HasNewSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) HasNewSnapshotConfig() bool`

HasNewSnapshotConfig returns a boolean if a field has been set.

### SetNewSnapshotConfigNil

`func (o *UpdateReplicationSnapshotConfig) SetNewSnapshotConfigNil(b bool)`

 SetNewSnapshotConfigNil sets the value for NewSnapshotConfig to be an explicit nil

### UnsetNewSnapshotConfig
`func (o *UpdateReplicationSnapshotConfig) UnsetNewSnapshotConfig()`

UnsetNewSnapshotConfig ensures that no value is present for NewSnapshotConfig, not even an explicit nil
### GetUpdateExistingSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) GetUpdateExistingSnapshotConfig() []UpdateExistingReplicationSnapshotConfig`

GetUpdateExistingSnapshotConfig returns the UpdateExistingSnapshotConfig field if non-nil, zero value otherwise.

### GetUpdateExistingSnapshotConfigOk

`func (o *UpdateReplicationSnapshotConfig) GetUpdateExistingSnapshotConfigOk() (*[]UpdateExistingReplicationSnapshotConfig, bool)`

GetUpdateExistingSnapshotConfigOk returns a tuple with the UpdateExistingSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) SetUpdateExistingSnapshotConfig(v []UpdateExistingReplicationSnapshotConfig)`

SetUpdateExistingSnapshotConfig sets UpdateExistingSnapshotConfig field to given value.

### HasUpdateExistingSnapshotConfig

`func (o *UpdateReplicationSnapshotConfig) HasUpdateExistingSnapshotConfig() bool`

HasUpdateExistingSnapshotConfig returns a boolean if a field has been set.

### SetUpdateExistingSnapshotConfigNil

`func (o *UpdateReplicationSnapshotConfig) SetUpdateExistingSnapshotConfigNil(b bool)`

 SetUpdateExistingSnapshotConfigNil sets the value for UpdateExistingSnapshotConfig to be an explicit nil

### UnsetUpdateExistingSnapshotConfig
`func (o *UpdateReplicationSnapshotConfig) UnsetUpdateExistingSnapshotConfig()`

UnsetUpdateExistingSnapshotConfig ensures that no value is present for UpdateExistingSnapshotConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


