# UpdateArchivalSnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSnapshotConfig** | Pointer to [**[]RunArchivalConfig**](RunArchivalConfig.md) | Specifies the new configuration about adding Archival Snapshot to existing Protection Group Run. | [optional] 
**UpdateExistingSnapshotConfig** | Pointer to [**[]UpdateExistingArchivalSnapshotConfig**](UpdateExistingArchivalSnapshotConfig.md) | Specifies the configuration about updating an existing Archival Snapshot Run. | [optional] 

## Methods

### NewUpdateArchivalSnapshotConfig

`func NewUpdateArchivalSnapshotConfig() *UpdateArchivalSnapshotConfig`

NewUpdateArchivalSnapshotConfig instantiates a new UpdateArchivalSnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateArchivalSnapshotConfigWithDefaults

`func NewUpdateArchivalSnapshotConfigWithDefaults() *UpdateArchivalSnapshotConfig`

NewUpdateArchivalSnapshotConfigWithDefaults instantiates a new UpdateArchivalSnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) GetNewSnapshotConfig() []RunArchivalConfig`

GetNewSnapshotConfig returns the NewSnapshotConfig field if non-nil, zero value otherwise.

### GetNewSnapshotConfigOk

`func (o *UpdateArchivalSnapshotConfig) GetNewSnapshotConfigOk() (*[]RunArchivalConfig, bool)`

GetNewSnapshotConfigOk returns a tuple with the NewSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) SetNewSnapshotConfig(v []RunArchivalConfig)`

SetNewSnapshotConfig sets NewSnapshotConfig field to given value.

### HasNewSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) HasNewSnapshotConfig() bool`

HasNewSnapshotConfig returns a boolean if a field has been set.

### SetNewSnapshotConfigNil

`func (o *UpdateArchivalSnapshotConfig) SetNewSnapshotConfigNil(b bool)`

 SetNewSnapshotConfigNil sets the value for NewSnapshotConfig to be an explicit nil

### UnsetNewSnapshotConfig
`func (o *UpdateArchivalSnapshotConfig) UnsetNewSnapshotConfig()`

UnsetNewSnapshotConfig ensures that no value is present for NewSnapshotConfig, not even an explicit nil
### GetUpdateExistingSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) GetUpdateExistingSnapshotConfig() []UpdateExistingArchivalSnapshotConfig`

GetUpdateExistingSnapshotConfig returns the UpdateExistingSnapshotConfig field if non-nil, zero value otherwise.

### GetUpdateExistingSnapshotConfigOk

`func (o *UpdateArchivalSnapshotConfig) GetUpdateExistingSnapshotConfigOk() (*[]UpdateExistingArchivalSnapshotConfig, bool)`

GetUpdateExistingSnapshotConfigOk returns a tuple with the UpdateExistingSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) SetUpdateExistingSnapshotConfig(v []UpdateExistingArchivalSnapshotConfig)`

SetUpdateExistingSnapshotConfig sets UpdateExistingSnapshotConfig field to given value.

### HasUpdateExistingSnapshotConfig

`func (o *UpdateArchivalSnapshotConfig) HasUpdateExistingSnapshotConfig() bool`

HasUpdateExistingSnapshotConfig returns a boolean if a field has been set.

### SetUpdateExistingSnapshotConfigNil

`func (o *UpdateArchivalSnapshotConfig) SetUpdateExistingSnapshotConfigNil(b bool)`

 SetUpdateExistingSnapshotConfigNil sets the value for UpdateExistingSnapshotConfig to be an explicit nil

### UnsetUpdateExistingSnapshotConfig
`func (o *UpdateArchivalSnapshotConfig) UnsetUpdateExistingSnapshotConfig()`

UnsetUpdateExistingSnapshotConfig ensures that no value is present for UpdateExistingSnapshotConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


