# ObjectSnapshotIdParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupId** | **NullableString** | Specifies the protection group id. | 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the snapshot in micro seconds. | [optional] 
**SnapshotJobInstanceId** | Pointer to **NullableInt64** | Specifies the instance id of the snapshot. | [optional] 
**SourceGroupId** | Pointer to **NullableString** | Specifies the source protection group id. | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the vault id. | [optional] 

## Methods

### NewObjectSnapshotIdParams

`func NewObjectSnapshotIdParams(protectionGroupId NullableString, ) *ObjectSnapshotIdParams`

NewObjectSnapshotIdParams instantiates a new ObjectSnapshotIdParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotIdParamsWithDefaults

`func NewObjectSnapshotIdParamsWithDefaults() *ObjectSnapshotIdParams`

NewObjectSnapshotIdParamsWithDefaults instantiates a new ObjectSnapshotIdParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupId

`func (o *ObjectSnapshotIdParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectSnapshotIdParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectSnapshotIdParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.


### SetProtectionGroupIdNil

`func (o *ObjectSnapshotIdParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectSnapshotIdParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *ObjectSnapshotIdParams) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *ObjectSnapshotIdParams) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *ObjectSnapshotIdParams) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *ObjectSnapshotIdParams) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *ObjectSnapshotIdParams) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *ObjectSnapshotIdParams) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetSnapshotJobInstanceId

`func (o *ObjectSnapshotIdParams) GetSnapshotJobInstanceId() int64`

GetSnapshotJobInstanceId returns the SnapshotJobInstanceId field if non-nil, zero value otherwise.

### GetSnapshotJobInstanceIdOk

`func (o *ObjectSnapshotIdParams) GetSnapshotJobInstanceIdOk() (*int64, bool)`

GetSnapshotJobInstanceIdOk returns a tuple with the SnapshotJobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotJobInstanceId

`func (o *ObjectSnapshotIdParams) SetSnapshotJobInstanceId(v int64)`

SetSnapshotJobInstanceId sets SnapshotJobInstanceId field to given value.

### HasSnapshotJobInstanceId

`func (o *ObjectSnapshotIdParams) HasSnapshotJobInstanceId() bool`

HasSnapshotJobInstanceId returns a boolean if a field has been set.

### SetSnapshotJobInstanceIdNil

`func (o *ObjectSnapshotIdParams) SetSnapshotJobInstanceIdNil(b bool)`

 SetSnapshotJobInstanceIdNil sets the value for SnapshotJobInstanceId to be an explicit nil

### UnsetSnapshotJobInstanceId
`func (o *ObjectSnapshotIdParams) UnsetSnapshotJobInstanceId()`

UnsetSnapshotJobInstanceId ensures that no value is present for SnapshotJobInstanceId, not even an explicit nil
### GetSourceGroupId

`func (o *ObjectSnapshotIdParams) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *ObjectSnapshotIdParams) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *ObjectSnapshotIdParams) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *ObjectSnapshotIdParams) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### SetSourceGroupIdNil

`func (o *ObjectSnapshotIdParams) SetSourceGroupIdNil(b bool)`

 SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil

### UnsetSourceGroupId
`func (o *ObjectSnapshotIdParams) UnsetSourceGroupId()`

UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil
### GetVaultId

`func (o *ObjectSnapshotIdParams) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *ObjectSnapshotIdParams) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *ObjectSnapshotIdParams) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *ObjectSnapshotIdParams) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *ObjectSnapshotIdParams) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *ObjectSnapshotIdParams) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


