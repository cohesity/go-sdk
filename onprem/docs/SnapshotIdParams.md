# SnapshotIdParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupId** | **NullableString** | Specifies the protection group id. | 
**SnapshotJobInstanceId** | Pointer to **NullableInt64** | Specifies the instance id of the snapshot. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the snapshot in micro seconds. | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the vault id. | [optional] 

## Methods

### NewSnapshotIdParams

`func NewSnapshotIdParams(protectionGroupId NullableString, ) *SnapshotIdParams`

NewSnapshotIdParams instantiates a new SnapshotIdParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotIdParamsWithDefaults

`func NewSnapshotIdParamsWithDefaults() *SnapshotIdParams`

NewSnapshotIdParamsWithDefaults instantiates a new SnapshotIdParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupId

`func (o *SnapshotIdParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *SnapshotIdParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *SnapshotIdParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.


### SetProtectionGroupIdNil

`func (o *SnapshotIdParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *SnapshotIdParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetSnapshotJobInstanceId

`func (o *SnapshotIdParams) GetSnapshotJobInstanceId() int64`

GetSnapshotJobInstanceId returns the SnapshotJobInstanceId field if non-nil, zero value otherwise.

### GetSnapshotJobInstanceIdOk

`func (o *SnapshotIdParams) GetSnapshotJobInstanceIdOk() (*int64, bool)`

GetSnapshotJobInstanceIdOk returns a tuple with the SnapshotJobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotJobInstanceId

`func (o *SnapshotIdParams) SetSnapshotJobInstanceId(v int64)`

SetSnapshotJobInstanceId sets SnapshotJobInstanceId field to given value.

### HasSnapshotJobInstanceId

`func (o *SnapshotIdParams) HasSnapshotJobInstanceId() bool`

HasSnapshotJobInstanceId returns a boolean if a field has been set.

### SetSnapshotJobInstanceIdNil

`func (o *SnapshotIdParams) SetSnapshotJobInstanceIdNil(b bool)`

 SetSnapshotJobInstanceIdNil sets the value for SnapshotJobInstanceId to be an explicit nil

### UnsetSnapshotJobInstanceId
`func (o *SnapshotIdParams) UnsetSnapshotJobInstanceId()`

UnsetSnapshotJobInstanceId ensures that no value is present for SnapshotJobInstanceId, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *SnapshotIdParams) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *SnapshotIdParams) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *SnapshotIdParams) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *SnapshotIdParams) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *SnapshotIdParams) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *SnapshotIdParams) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetVaultId

`func (o *SnapshotIdParams) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *SnapshotIdParams) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *SnapshotIdParams) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *SnapshotIdParams) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *SnapshotIdParams) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *SnapshotIdParams) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


