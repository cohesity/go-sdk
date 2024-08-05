# UpdateShareParamSmbConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachingEnabled** | Pointer to **NullableBool** | Indicate if offline file caching is supported. | [optional] 
**ContinuousAvailability** | Pointer to **NullableBool** | Whether file open handles are persisted to scribe to survive bridge process crash. When set to false, open handles will be kept in memory until the current node has exclusive ticket for the entity handle. When the entity is opened from another node, the exclusive ticket would be revoked from the node. In revoke control flow, the current node would persist the state to scribe. On acquiring the exclusive ticket,another node would read the file open handles from scribe and resume the handling of operation. | [optional] 
**DiscoveryEnabled** | Pointer to **NullableBool** | Whether the share is discoverable. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether SMB encryption is enabled for this share. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EncryptionRequired** | Pointer to **NullableBool** | Whether to enforce encryption for all the sessions for this view. When enabled all unencrypted sessions are disallowed. | [optional] 
**IsShareLevelPermissionEmpty** | Pointer to **NullableBool** | Indicate if share level permission is cleared by user. | [optional] 
**OplockEnabled** | Pointer to **NullableBool** | Indicate the operation lock is enabled by this view. | [optional] 
**Permissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Share level permissions. Note: Supported Access: FullControl, Modify, ReadOnly. Supported type: Allow, Deny. | [optional] 
**SuperUserSids** | Pointer to **[]string** | Specifies a list of super user sids. Duplicate SIDs are not allowed. | [optional] 

## Methods

### NewUpdateShareParamSmbConfig

`func NewUpdateShareParamSmbConfig() *UpdateShareParamSmbConfig`

NewUpdateShareParamSmbConfig instantiates a new UpdateShareParamSmbConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShareParamSmbConfigWithDefaults

`func NewUpdateShareParamSmbConfigWithDefaults() *UpdateShareParamSmbConfig`

NewUpdateShareParamSmbConfigWithDefaults instantiates a new UpdateShareParamSmbConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachingEnabled

`func (o *UpdateShareParamSmbConfig) GetCachingEnabled() bool`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *UpdateShareParamSmbConfig) GetCachingEnabledOk() (*bool, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *UpdateShareParamSmbConfig) SetCachingEnabled(v bool)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *UpdateShareParamSmbConfig) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### SetCachingEnabledNil

`func (o *UpdateShareParamSmbConfig) SetCachingEnabledNil(b bool)`

 SetCachingEnabledNil sets the value for CachingEnabled to be an explicit nil

### UnsetCachingEnabled
`func (o *UpdateShareParamSmbConfig) UnsetCachingEnabled()`

UnsetCachingEnabled ensures that no value is present for CachingEnabled, not even an explicit nil
### GetContinuousAvailability

`func (o *UpdateShareParamSmbConfig) GetContinuousAvailability() bool`

GetContinuousAvailability returns the ContinuousAvailability field if non-nil, zero value otherwise.

### GetContinuousAvailabilityOk

`func (o *UpdateShareParamSmbConfig) GetContinuousAvailabilityOk() (*bool, bool)`

GetContinuousAvailabilityOk returns a tuple with the ContinuousAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousAvailability

`func (o *UpdateShareParamSmbConfig) SetContinuousAvailability(v bool)`

SetContinuousAvailability sets ContinuousAvailability field to given value.

### HasContinuousAvailability

`func (o *UpdateShareParamSmbConfig) HasContinuousAvailability() bool`

HasContinuousAvailability returns a boolean if a field has been set.

### SetContinuousAvailabilityNil

`func (o *UpdateShareParamSmbConfig) SetContinuousAvailabilityNil(b bool)`

 SetContinuousAvailabilityNil sets the value for ContinuousAvailability to be an explicit nil

### UnsetContinuousAvailability
`func (o *UpdateShareParamSmbConfig) UnsetContinuousAvailability()`

UnsetContinuousAvailability ensures that no value is present for ContinuousAvailability, not even an explicit nil
### GetDiscoveryEnabled

`func (o *UpdateShareParamSmbConfig) GetDiscoveryEnabled() bool`

GetDiscoveryEnabled returns the DiscoveryEnabled field if non-nil, zero value otherwise.

### GetDiscoveryEnabledOk

`func (o *UpdateShareParamSmbConfig) GetDiscoveryEnabledOk() (*bool, bool)`

GetDiscoveryEnabledOk returns a tuple with the DiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEnabled

`func (o *UpdateShareParamSmbConfig) SetDiscoveryEnabled(v bool)`

SetDiscoveryEnabled sets DiscoveryEnabled field to given value.

### HasDiscoveryEnabled

`func (o *UpdateShareParamSmbConfig) HasDiscoveryEnabled() bool`

HasDiscoveryEnabled returns a boolean if a field has been set.

### SetDiscoveryEnabledNil

`func (o *UpdateShareParamSmbConfig) SetDiscoveryEnabledNil(b bool)`

 SetDiscoveryEnabledNil sets the value for DiscoveryEnabled to be an explicit nil

### UnsetDiscoveryEnabled
`func (o *UpdateShareParamSmbConfig) UnsetDiscoveryEnabled()`

UnsetDiscoveryEnabled ensures that no value is present for DiscoveryEnabled, not even an explicit nil
### GetEncryptionEnabled

`func (o *UpdateShareParamSmbConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *UpdateShareParamSmbConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *UpdateShareParamSmbConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *UpdateShareParamSmbConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *UpdateShareParamSmbConfig) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *UpdateShareParamSmbConfig) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetEncryptionRequired

`func (o *UpdateShareParamSmbConfig) GetEncryptionRequired() bool`

GetEncryptionRequired returns the EncryptionRequired field if non-nil, zero value otherwise.

### GetEncryptionRequiredOk

`func (o *UpdateShareParamSmbConfig) GetEncryptionRequiredOk() (*bool, bool)`

GetEncryptionRequiredOk returns a tuple with the EncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionRequired

`func (o *UpdateShareParamSmbConfig) SetEncryptionRequired(v bool)`

SetEncryptionRequired sets EncryptionRequired field to given value.

### HasEncryptionRequired

`func (o *UpdateShareParamSmbConfig) HasEncryptionRequired() bool`

HasEncryptionRequired returns a boolean if a field has been set.

### SetEncryptionRequiredNil

`func (o *UpdateShareParamSmbConfig) SetEncryptionRequiredNil(b bool)`

 SetEncryptionRequiredNil sets the value for EncryptionRequired to be an explicit nil

### UnsetEncryptionRequired
`func (o *UpdateShareParamSmbConfig) UnsetEncryptionRequired()`

UnsetEncryptionRequired ensures that no value is present for EncryptionRequired, not even an explicit nil
### GetIsShareLevelPermissionEmpty

`func (o *UpdateShareParamSmbConfig) GetIsShareLevelPermissionEmpty() bool`

GetIsShareLevelPermissionEmpty returns the IsShareLevelPermissionEmpty field if non-nil, zero value otherwise.

### GetIsShareLevelPermissionEmptyOk

`func (o *UpdateShareParamSmbConfig) GetIsShareLevelPermissionEmptyOk() (*bool, bool)`

GetIsShareLevelPermissionEmptyOk returns a tuple with the IsShareLevelPermissionEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareLevelPermissionEmpty

`func (o *UpdateShareParamSmbConfig) SetIsShareLevelPermissionEmpty(v bool)`

SetIsShareLevelPermissionEmpty sets IsShareLevelPermissionEmpty field to given value.

### HasIsShareLevelPermissionEmpty

`func (o *UpdateShareParamSmbConfig) HasIsShareLevelPermissionEmpty() bool`

HasIsShareLevelPermissionEmpty returns a boolean if a field has been set.

### SetIsShareLevelPermissionEmptyNil

`func (o *UpdateShareParamSmbConfig) SetIsShareLevelPermissionEmptyNil(b bool)`

 SetIsShareLevelPermissionEmptyNil sets the value for IsShareLevelPermissionEmpty to be an explicit nil

### UnsetIsShareLevelPermissionEmpty
`func (o *UpdateShareParamSmbConfig) UnsetIsShareLevelPermissionEmpty()`

UnsetIsShareLevelPermissionEmpty ensures that no value is present for IsShareLevelPermissionEmpty, not even an explicit nil
### GetOplockEnabled

`func (o *UpdateShareParamSmbConfig) GetOplockEnabled() bool`

GetOplockEnabled returns the OplockEnabled field if non-nil, zero value otherwise.

### GetOplockEnabledOk

`func (o *UpdateShareParamSmbConfig) GetOplockEnabledOk() (*bool, bool)`

GetOplockEnabledOk returns a tuple with the OplockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplockEnabled

`func (o *UpdateShareParamSmbConfig) SetOplockEnabled(v bool)`

SetOplockEnabled sets OplockEnabled field to given value.

### HasOplockEnabled

`func (o *UpdateShareParamSmbConfig) HasOplockEnabled() bool`

HasOplockEnabled returns a boolean if a field has been set.

### SetOplockEnabledNil

`func (o *UpdateShareParamSmbConfig) SetOplockEnabledNil(b bool)`

 SetOplockEnabledNil sets the value for OplockEnabled to be an explicit nil

### UnsetOplockEnabled
`func (o *UpdateShareParamSmbConfig) UnsetOplockEnabled()`

UnsetOplockEnabled ensures that no value is present for OplockEnabled, not even an explicit nil
### GetPermissions

`func (o *UpdateShareParamSmbConfig) GetPermissions() []SmbPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateShareParamSmbConfig) GetPermissionsOk() (*[]SmbPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateShareParamSmbConfig) SetPermissions(v []SmbPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateShareParamSmbConfig) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *UpdateShareParamSmbConfig) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UpdateShareParamSmbConfig) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSuperUserSids

`func (o *UpdateShareParamSmbConfig) GetSuperUserSids() []string`

GetSuperUserSids returns the SuperUserSids field if non-nil, zero value otherwise.

### GetSuperUserSidsOk

`func (o *UpdateShareParamSmbConfig) GetSuperUserSidsOk() (*[]string, bool)`

GetSuperUserSidsOk returns a tuple with the SuperUserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUserSids

`func (o *UpdateShareParamSmbConfig) SetSuperUserSids(v []string)`

SetSuperUserSids sets SuperUserSids field to given value.

### HasSuperUserSids

`func (o *UpdateShareParamSmbConfig) HasSuperUserSids() bool`

HasSuperUserSids returns a boolean if a field has been set.

### SetSuperUserSidsNil

`func (o *UpdateShareParamSmbConfig) SetSuperUserSidsNil(b bool)`

 SetSuperUserSidsNil sets the value for SuperUserSids to be an explicit nil

### UnsetSuperUserSids
`func (o *UpdateShareParamSmbConfig) UnsetSuperUserSids()`

UnsetSuperUserSids ensures that no value is present for SuperUserSids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


