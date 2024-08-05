# StorageArraySnapshotThrottlingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the volume ID of the Storage Snapshot Mgmt throttling Policy. | [optional] 
**MaxSnapshotsConfigEnabled** | Pointer to **NullableBool** | Specifies whether we will use storage snapshot managmement max snapshots config to all volumes/luns that are part of the registered entity. | [optional] 
**MaxSnapshotsMgmtSnapshotConfig** | Pointer to [**StorageArraySnapshotMaxSnapshotConfig**](StorageArraySnapshotMaxSnapshotConfig.md) |  | [optional] 
**MaxSnapshotsMgmtSpaceConfig** | Pointer to [**StorageArraySnapshotMaxSpaceConfig**](StorageArraySnapshotMaxSpaceConfig.md) |  | [optional] 
**MaxSpaceConfigEnabled** | Pointer to **NullableBool** | Specifies whether we will use storage snapshot managmement max space config to all volumes/luns that are part of the registered entity. | [optional] 

## Methods

### NewStorageArraySnapshotThrottlingPolicy

`func NewStorageArraySnapshotThrottlingPolicy() *StorageArraySnapshotThrottlingPolicy`

NewStorageArraySnapshotThrottlingPolicy instantiates a new StorageArraySnapshotThrottlingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageArraySnapshotThrottlingPolicyWithDefaults

`func NewStorageArraySnapshotThrottlingPolicyWithDefaults() *StorageArraySnapshotThrottlingPolicy`

NewStorageArraySnapshotThrottlingPolicyWithDefaults instantiates a new StorageArraySnapshotThrottlingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageArraySnapshotThrottlingPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageArraySnapshotThrottlingPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageArraySnapshotThrottlingPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageArraySnapshotThrottlingPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StorageArraySnapshotThrottlingPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StorageArraySnapshotThrottlingPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsConfigEnabled() bool`

GetMaxSnapshotsConfigEnabled returns the MaxSnapshotsConfigEnabled field if non-nil, zero value otherwise.

### GetMaxSnapshotsConfigEnabledOk

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsConfigEnabledOk() (*bool, bool)`

GetMaxSnapshotsConfigEnabledOk returns a tuple with the MaxSnapshotsConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSnapshotsConfigEnabled(v bool)`

SetMaxSnapshotsConfigEnabled sets MaxSnapshotsConfigEnabled field to given value.

### HasMaxSnapshotsConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) HasMaxSnapshotsConfigEnabled() bool`

HasMaxSnapshotsConfigEnabled returns a boolean if a field has been set.

### SetMaxSnapshotsConfigEnabledNil

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSnapshotsConfigEnabledNil(b bool)`

 SetMaxSnapshotsConfigEnabledNil sets the value for MaxSnapshotsConfigEnabled to be an explicit nil

### UnsetMaxSnapshotsConfigEnabled
`func (o *StorageArraySnapshotThrottlingPolicy) UnsetMaxSnapshotsConfigEnabled()`

UnsetMaxSnapshotsConfigEnabled ensures that no value is present for MaxSnapshotsConfigEnabled, not even an explicit nil
### GetMaxSnapshotsMgmtSnapshotConfig

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsMgmtSnapshotConfig() StorageArraySnapshotMaxSnapshotConfig`

GetMaxSnapshotsMgmtSnapshotConfig returns the MaxSnapshotsMgmtSnapshotConfig field if non-nil, zero value otherwise.

### GetMaxSnapshotsMgmtSnapshotConfigOk

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsMgmtSnapshotConfigOk() (*StorageArraySnapshotMaxSnapshotConfig, bool)`

GetMaxSnapshotsMgmtSnapshotConfigOk returns a tuple with the MaxSnapshotsMgmtSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsMgmtSnapshotConfig

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSnapshotsMgmtSnapshotConfig(v StorageArraySnapshotMaxSnapshotConfig)`

SetMaxSnapshotsMgmtSnapshotConfig sets MaxSnapshotsMgmtSnapshotConfig field to given value.

### HasMaxSnapshotsMgmtSnapshotConfig

`func (o *StorageArraySnapshotThrottlingPolicy) HasMaxSnapshotsMgmtSnapshotConfig() bool`

HasMaxSnapshotsMgmtSnapshotConfig returns a boolean if a field has been set.

### GetMaxSnapshotsMgmtSpaceConfig

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsMgmtSpaceConfig() StorageArraySnapshotMaxSpaceConfig`

GetMaxSnapshotsMgmtSpaceConfig returns the MaxSnapshotsMgmtSpaceConfig field if non-nil, zero value otherwise.

### GetMaxSnapshotsMgmtSpaceConfigOk

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSnapshotsMgmtSpaceConfigOk() (*StorageArraySnapshotMaxSpaceConfig, bool)`

GetMaxSnapshotsMgmtSpaceConfigOk returns a tuple with the MaxSnapshotsMgmtSpaceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsMgmtSpaceConfig

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSnapshotsMgmtSpaceConfig(v StorageArraySnapshotMaxSpaceConfig)`

SetMaxSnapshotsMgmtSpaceConfig sets MaxSnapshotsMgmtSpaceConfig field to given value.

### HasMaxSnapshotsMgmtSpaceConfig

`func (o *StorageArraySnapshotThrottlingPolicy) HasMaxSnapshotsMgmtSpaceConfig() bool`

HasMaxSnapshotsMgmtSpaceConfig returns a boolean if a field has been set.

### GetMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSpaceConfigEnabled() bool`

GetMaxSpaceConfigEnabled returns the MaxSpaceConfigEnabled field if non-nil, zero value otherwise.

### GetMaxSpaceConfigEnabledOk

`func (o *StorageArraySnapshotThrottlingPolicy) GetMaxSpaceConfigEnabledOk() (*bool, bool)`

GetMaxSpaceConfigEnabledOk returns a tuple with the MaxSpaceConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSpaceConfigEnabled(v bool)`

SetMaxSpaceConfigEnabled sets MaxSpaceConfigEnabled field to given value.

### HasMaxSpaceConfigEnabled

`func (o *StorageArraySnapshotThrottlingPolicy) HasMaxSpaceConfigEnabled() bool`

HasMaxSpaceConfigEnabled returns a boolean if a field has been set.

### SetMaxSpaceConfigEnabledNil

`func (o *StorageArraySnapshotThrottlingPolicy) SetMaxSpaceConfigEnabledNil(b bool)`

 SetMaxSpaceConfigEnabledNil sets the value for MaxSpaceConfigEnabled to be an explicit nil

### UnsetMaxSpaceConfigEnabled
`func (o *StorageArraySnapshotThrottlingPolicy) UnsetMaxSpaceConfigEnabled()`

UnsetMaxSpaceConfigEnabled ensures that no value is present for MaxSpaceConfigEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


