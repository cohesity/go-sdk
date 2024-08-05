# ProtectedObjectGroupBackupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id, if given object is also protected by a protection group. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the protection group name, if given object is also protected by a protection group. | [optional] [readonly] 

## Methods

### NewProtectedObjectGroupBackupConfig

`func NewProtectedObjectGroupBackupConfig() *ProtectedObjectGroupBackupConfig`

NewProtectedObjectGroupBackupConfig instantiates a new ProtectedObjectGroupBackupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectGroupBackupConfigWithDefaults

`func NewProtectedObjectGroupBackupConfigWithDefaults() *ProtectedObjectGroupBackupConfig`

NewProtectedObjectGroupBackupConfigWithDefaults instantiates a new ProtectedObjectGroupBackupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupId

`func (o *ProtectedObjectGroupBackupConfig) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ProtectedObjectGroupBackupConfig) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ProtectedObjectGroupBackupConfig) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ProtectedObjectGroupBackupConfig) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ProtectedObjectGroupBackupConfig) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ProtectedObjectGroupBackupConfig) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ProtectedObjectGroupBackupConfig) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ProtectedObjectGroupBackupConfig) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ProtectedObjectGroupBackupConfig) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ProtectedObjectGroupBackupConfig) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ProtectedObjectGroupBackupConfig) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ProtectedObjectGroupBackupConfig) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


