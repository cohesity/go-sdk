# ProtectedObjectBackupConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoProtectConfig** | Pointer to **NullableBool** | Specifies whether or not this configuration is applied to an autoprotected object rather than this specific object. | [optional] 
**AutoProtectParentId** | Pointer to **NullableInt64** | Specifies the parent ID of the object which the backup configuration is applied to if this is an auto protect config. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies whether or not protection has been paused on this object. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies whether or not protection has been deactivated on this object. | [optional] 

## Methods

### NewProtectedObjectBackupConfigAllOf

`func NewProtectedObjectBackupConfigAllOf() *ProtectedObjectBackupConfigAllOf`

NewProtectedObjectBackupConfigAllOf instantiates a new ProtectedObjectBackupConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectBackupConfigAllOfWithDefaults

`func NewProtectedObjectBackupConfigAllOfWithDefaults() *ProtectedObjectBackupConfigAllOf`

NewProtectedObjectBackupConfigAllOfWithDefaults instantiates a new ProtectedObjectBackupConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfigAllOf) GetIsAutoProtectConfig() bool`

GetIsAutoProtectConfig returns the IsAutoProtectConfig field if non-nil, zero value otherwise.

### GetIsAutoProtectConfigOk

`func (o *ProtectedObjectBackupConfigAllOf) GetIsAutoProtectConfigOk() (*bool, bool)`

GetIsAutoProtectConfigOk returns a tuple with the IsAutoProtectConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfigAllOf) SetIsAutoProtectConfig(v bool)`

SetIsAutoProtectConfig sets IsAutoProtectConfig field to given value.

### HasIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfigAllOf) HasIsAutoProtectConfig() bool`

HasIsAutoProtectConfig returns a boolean if a field has been set.

### SetIsAutoProtectConfigNil

`func (o *ProtectedObjectBackupConfigAllOf) SetIsAutoProtectConfigNil(b bool)`

 SetIsAutoProtectConfigNil sets the value for IsAutoProtectConfig to be an explicit nil

### UnsetIsAutoProtectConfig
`func (o *ProtectedObjectBackupConfigAllOf) UnsetIsAutoProtectConfig()`

UnsetIsAutoProtectConfig ensures that no value is present for IsAutoProtectConfig, not even an explicit nil
### GetAutoProtectParentId

`func (o *ProtectedObjectBackupConfigAllOf) GetAutoProtectParentId() int64`

GetAutoProtectParentId returns the AutoProtectParentId field if non-nil, zero value otherwise.

### GetAutoProtectParentIdOk

`func (o *ProtectedObjectBackupConfigAllOf) GetAutoProtectParentIdOk() (*int64, bool)`

GetAutoProtectParentIdOk returns a tuple with the AutoProtectParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoProtectParentId

`func (o *ProtectedObjectBackupConfigAllOf) SetAutoProtectParentId(v int64)`

SetAutoProtectParentId sets AutoProtectParentId field to given value.

### HasAutoProtectParentId

`func (o *ProtectedObjectBackupConfigAllOf) HasAutoProtectParentId() bool`

HasAutoProtectParentId returns a boolean if a field has been set.

### SetAutoProtectParentIdNil

`func (o *ProtectedObjectBackupConfigAllOf) SetAutoProtectParentIdNil(b bool)`

 SetAutoProtectParentIdNil sets the value for AutoProtectParentId to be an explicit nil

### UnsetAutoProtectParentId
`func (o *ProtectedObjectBackupConfigAllOf) UnsetAutoProtectParentId()`

UnsetAutoProtectParentId ensures that no value is present for AutoProtectParentId, not even an explicit nil
### GetIsPaused

`func (o *ProtectedObjectBackupConfigAllOf) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectedObjectBackupConfigAllOf) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectedObjectBackupConfigAllOf) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectedObjectBackupConfigAllOf) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectedObjectBackupConfigAllOf) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectedObjectBackupConfigAllOf) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsActive

`func (o *ProtectedObjectBackupConfigAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProtectedObjectBackupConfigAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProtectedObjectBackupConfigAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProtectedObjectBackupConfigAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProtectedObjectBackupConfigAllOf) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProtectedObjectBackupConfigAllOf) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


