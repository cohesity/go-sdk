# StorageArraySnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSnapshotConfig** | Pointer to [**StorageArraySnapshotMaxSnapshotConfig**](StorageArraySnapshotMaxSnapshotConfig.md) |  | [optional] 
**MaxSnapshotsConfigEnabled** | Pointer to **NullableBool** | Specifies whether we will use storage snapshot managmement max snapshots config to all volumes/luns that are part of the registered entity. | [optional] 
**MaxSpaceConfig** | Pointer to [**StorageArraySnapshotMaxSpaceConfig**](StorageArraySnapshotMaxSpaceConfig.md) |  | [optional] 
**MaxSpaceConfigEnabled** | Pointer to **NullableBool** | Specifies whether we will use storage snapshot managmement max space config to all volumes/luns that are part of the registered entity. | [optional] 
**StorageArraySnapshotThrottlingPolicies** | Pointer to [**[]StorageArraySnapshotThrottlingPolicy**](StorageArraySnapshotThrottlingPolicy.md) | Specifies the list of storage array snapshot management throttling policies for individual volume/lun. | [optional] 

## Methods

### NewStorageArraySnapshotConfig

`func NewStorageArraySnapshotConfig() *StorageArraySnapshotConfig`

NewStorageArraySnapshotConfig instantiates a new StorageArraySnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageArraySnapshotConfigWithDefaults

`func NewStorageArraySnapshotConfigWithDefaults() *StorageArraySnapshotConfig`

NewStorageArraySnapshotConfigWithDefaults instantiates a new StorageArraySnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSnapshotConfig

`func (o *StorageArraySnapshotConfig) GetMaxSnapshotConfig() StorageArraySnapshotMaxSnapshotConfig`

GetMaxSnapshotConfig returns the MaxSnapshotConfig field if non-nil, zero value otherwise.

### GetMaxSnapshotConfigOk

`func (o *StorageArraySnapshotConfig) GetMaxSnapshotConfigOk() (*StorageArraySnapshotMaxSnapshotConfig, bool)`

GetMaxSnapshotConfigOk returns a tuple with the MaxSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotConfig

`func (o *StorageArraySnapshotConfig) SetMaxSnapshotConfig(v StorageArraySnapshotMaxSnapshotConfig)`

SetMaxSnapshotConfig sets MaxSnapshotConfig field to given value.

### HasMaxSnapshotConfig

`func (o *StorageArraySnapshotConfig) HasMaxSnapshotConfig() bool`

HasMaxSnapshotConfig returns a boolean if a field has been set.

### GetMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotConfig) GetMaxSnapshotsConfigEnabled() bool`

GetMaxSnapshotsConfigEnabled returns the MaxSnapshotsConfigEnabled field if non-nil, zero value otherwise.

### GetMaxSnapshotsConfigEnabledOk

`func (o *StorageArraySnapshotConfig) GetMaxSnapshotsConfigEnabledOk() (*bool, bool)`

GetMaxSnapshotsConfigEnabledOk returns a tuple with the MaxSnapshotsConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotConfig) SetMaxSnapshotsConfigEnabled(v bool)`

SetMaxSnapshotsConfigEnabled sets MaxSnapshotsConfigEnabled field to given value.

### HasMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotConfig) HasMaxSnapshotsConfigEnabled() bool`

HasMaxSnapshotsConfigEnabled returns a boolean if a field has been set.

### SetMaxSnapshotsConfigEnabledNil

`func (o *StorageArraySnapshotConfig) SetMaxSnapshotsConfigEnabledNil(b bool)`

 SetMaxSnapshotsConfigEnabledNil sets the value for MaxSnapshotsConfigEnabled to be an explicit nil

### UnsetMaxSnapshotsConfigEnabled
`func (o *StorageArraySnapshotConfig) UnsetMaxSnapshotsConfigEnabled()`

UnsetMaxSnapshotsConfigEnabled ensures that no value is present for MaxSnapshotsConfigEnabled, not even an explicit nil
### GetMaxSpaceConfig

`func (o *StorageArraySnapshotConfig) GetMaxSpaceConfig() StorageArraySnapshotMaxSpaceConfig`

GetMaxSpaceConfig returns the MaxSpaceConfig field if non-nil, zero value otherwise.

### GetMaxSpaceConfigOk

`func (o *StorageArraySnapshotConfig) GetMaxSpaceConfigOk() (*StorageArraySnapshotMaxSpaceConfig, bool)`

GetMaxSpaceConfigOk returns a tuple with the MaxSpaceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpaceConfig

`func (o *StorageArraySnapshotConfig) SetMaxSpaceConfig(v StorageArraySnapshotMaxSpaceConfig)`

SetMaxSpaceConfig sets MaxSpaceConfig field to given value.

### HasMaxSpaceConfig

`func (o *StorageArraySnapshotConfig) HasMaxSpaceConfig() bool`

HasMaxSpaceConfig returns a boolean if a field has been set.

### GetMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotConfig) GetMaxSpaceConfigEnabled() bool`

GetMaxSpaceConfigEnabled returns the MaxSpaceConfigEnabled field if non-nil, zero value otherwise.

### GetMaxSpaceConfigEnabledOk

`func (o *StorageArraySnapshotConfig) GetMaxSpaceConfigEnabledOk() (*bool, bool)`

GetMaxSpaceConfigEnabledOk returns a tuple with the MaxSpaceConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotConfig) SetMaxSpaceConfigEnabled(v bool)`

SetMaxSpaceConfigEnabled sets MaxSpaceConfigEnabled field to given value.

### HasMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotConfig) HasMaxSpaceConfigEnabled() bool`

HasMaxSpaceConfigEnabled returns a boolean if a field has been set.

### SetMaxSpaceConfigEnabledNil

`func (o *StorageArraySnapshotConfig) SetMaxSpaceConfigEnabledNil(b bool)`

 SetMaxSpaceConfigEnabledNil sets the value for MaxSpaceConfigEnabled to be an explicit nil

### UnsetMaxSpaceConfigEnabled
`func (o *StorageArraySnapshotConfig) UnsetMaxSpaceConfigEnabled()`

UnsetMaxSpaceConfigEnabled ensures that no value is present for MaxSpaceConfigEnabled, not even an explicit nil
### GetStorageArraySnapshotThrottlingPolicies

`func (o *StorageArraySnapshotConfig) GetStorageArraySnapshotThrottlingPolicies() []StorageArraySnapshotThrottlingPolicy`

GetStorageArraySnapshotThrottlingPolicies returns the StorageArraySnapshotThrottlingPolicies field if non-nil, zero value otherwise.

### GetStorageArraySnapshotThrottlingPoliciesOk

`func (o *StorageArraySnapshotConfig) GetStorageArraySnapshotThrottlingPoliciesOk() (*[]StorageArraySnapshotThrottlingPolicy, bool)`

GetStorageArraySnapshotThrottlingPoliciesOk returns a tuple with the StorageArraySnapshotThrottlingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArraySnapshotThrottlingPolicies

`func (o *StorageArraySnapshotConfig) SetStorageArraySnapshotThrottlingPolicies(v []StorageArraySnapshotThrottlingPolicy)`

SetStorageArraySnapshotThrottlingPolicies sets StorageArraySnapshotThrottlingPolicies field to given value.

### HasStorageArraySnapshotThrottlingPolicies

`func (o *StorageArraySnapshotConfig) HasStorageArraySnapshotThrottlingPolicies() bool`

HasStorageArraySnapshotThrottlingPolicies returns a boolean if a field has been set.

### SetStorageArraySnapshotThrottlingPoliciesNil

`func (o *StorageArraySnapshotConfig) SetStorageArraySnapshotThrottlingPoliciesNil(b bool)`

 SetStorageArraySnapshotThrottlingPoliciesNil sets the value for StorageArraySnapshotThrottlingPolicies to be an explicit nil

### UnsetStorageArraySnapshotThrottlingPolicies
`func (o *StorageArraySnapshotConfig) UnsetStorageArraySnapshotThrottlingPolicies()`

UnsetStorageArraySnapshotThrottlingPolicies ensures that no value is present for StorageArraySnapshotThrottlingPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


