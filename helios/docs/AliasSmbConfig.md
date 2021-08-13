# AliasSmbConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryEnabled** | Pointer to **NullableBool** | Whether the share is discoverable. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether SMB encryption is enabled for this share. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EncryptionRequired** | Pointer to **NullableBool** | Whether to enforce encryption for all the sessions for this view. When enabled all unencrypted sessions are disallowed. | [optional] 
**Permissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Share level permissions. | [optional] 
**SuperUserSids** | Pointer to **[]string** | Specifies a list of super user sids. | [optional] 
**CachingEnabled** | Pointer to **NullableBool** | Indicate if offline file caching is supported. | [optional] 
**IsShareLevelPermissionEmpty** | Pointer to **NullableBool** | Indicate if share level permission is cleared by user. | [optional] 
**OplockEnabled** | Pointer to **NullableBool** | Indicate the operation lock is enabled by this view. | [optional] 
**ContinuousAvailability** | Pointer to **NullableBool** | Whether file open handles are persited to scribe to survive bridge process crash. When set to false, open handles will be kept in memory untill the current node has exclusive ticket for the entity handle. When the entity is opened from another node, the exclusive ticket would be revoked from the node. In revoke control flow, the current node would persist the state to scribe. On acquiring the exclusive ticket,another node would read the file open handles from scribe and resume the handling of operation. | [optional] 

## Methods

### NewAliasSmbConfig

`func NewAliasSmbConfig() *AliasSmbConfig`

NewAliasSmbConfig instantiates a new AliasSmbConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasSmbConfigWithDefaults

`func NewAliasSmbConfigWithDefaults() *AliasSmbConfig`

NewAliasSmbConfigWithDefaults instantiates a new AliasSmbConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryEnabled

`func (o *AliasSmbConfig) GetDiscoveryEnabled() bool`

GetDiscoveryEnabled returns the DiscoveryEnabled field if non-nil, zero value otherwise.

### GetDiscoveryEnabledOk

`func (o *AliasSmbConfig) GetDiscoveryEnabledOk() (*bool, bool)`

GetDiscoveryEnabledOk returns a tuple with the DiscoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEnabled

`func (o *AliasSmbConfig) SetDiscoveryEnabled(v bool)`

SetDiscoveryEnabled sets DiscoveryEnabled field to given value.

### HasDiscoveryEnabled

`func (o *AliasSmbConfig) HasDiscoveryEnabled() bool`

HasDiscoveryEnabled returns a boolean if a field has been set.

### SetDiscoveryEnabledNil

`func (o *AliasSmbConfig) SetDiscoveryEnabledNil(b bool)`

 SetDiscoveryEnabledNil sets the value for DiscoveryEnabled to be an explicit nil

### UnsetDiscoveryEnabled
`func (o *AliasSmbConfig) UnsetDiscoveryEnabled()`

UnsetDiscoveryEnabled ensures that no value is present for DiscoveryEnabled, not even an explicit nil
### GetEncryptionEnabled

`func (o *AliasSmbConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *AliasSmbConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *AliasSmbConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *AliasSmbConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *AliasSmbConfig) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *AliasSmbConfig) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetEncryptionRequired

`func (o *AliasSmbConfig) GetEncryptionRequired() bool`

GetEncryptionRequired returns the EncryptionRequired field if non-nil, zero value otherwise.

### GetEncryptionRequiredOk

`func (o *AliasSmbConfig) GetEncryptionRequiredOk() (*bool, bool)`

GetEncryptionRequiredOk returns a tuple with the EncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionRequired

`func (o *AliasSmbConfig) SetEncryptionRequired(v bool)`

SetEncryptionRequired sets EncryptionRequired field to given value.

### HasEncryptionRequired

`func (o *AliasSmbConfig) HasEncryptionRequired() bool`

HasEncryptionRequired returns a boolean if a field has been set.

### SetEncryptionRequiredNil

`func (o *AliasSmbConfig) SetEncryptionRequiredNil(b bool)`

 SetEncryptionRequiredNil sets the value for EncryptionRequired to be an explicit nil

### UnsetEncryptionRequired
`func (o *AliasSmbConfig) UnsetEncryptionRequired()`

UnsetEncryptionRequired ensures that no value is present for EncryptionRequired, not even an explicit nil
### GetPermissions

`func (o *AliasSmbConfig) GetPermissions() []SmbPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AliasSmbConfig) GetPermissionsOk() (*[]SmbPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AliasSmbConfig) SetPermissions(v []SmbPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AliasSmbConfig) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AliasSmbConfig) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AliasSmbConfig) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSuperUserSids

`func (o *AliasSmbConfig) GetSuperUserSids() []string`

GetSuperUserSids returns the SuperUserSids field if non-nil, zero value otherwise.

### GetSuperUserSidsOk

`func (o *AliasSmbConfig) GetSuperUserSidsOk() (*[]string, bool)`

GetSuperUserSidsOk returns a tuple with the SuperUserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUserSids

`func (o *AliasSmbConfig) SetSuperUserSids(v []string)`

SetSuperUserSids sets SuperUserSids field to given value.

### HasSuperUserSids

`func (o *AliasSmbConfig) HasSuperUserSids() bool`

HasSuperUserSids returns a boolean if a field has been set.

### SetSuperUserSidsNil

`func (o *AliasSmbConfig) SetSuperUserSidsNil(b bool)`

 SetSuperUserSidsNil sets the value for SuperUserSids to be an explicit nil

### UnsetSuperUserSids
`func (o *AliasSmbConfig) UnsetSuperUserSids()`

UnsetSuperUserSids ensures that no value is present for SuperUserSids, not even an explicit nil
### GetCachingEnabled

`func (o *AliasSmbConfig) GetCachingEnabled() bool`

GetCachingEnabled returns the CachingEnabled field if non-nil, zero value otherwise.

### GetCachingEnabledOk

`func (o *AliasSmbConfig) GetCachingEnabledOk() (*bool, bool)`

GetCachingEnabledOk returns a tuple with the CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingEnabled

`func (o *AliasSmbConfig) SetCachingEnabled(v bool)`

SetCachingEnabled sets CachingEnabled field to given value.

### HasCachingEnabled

`func (o *AliasSmbConfig) HasCachingEnabled() bool`

HasCachingEnabled returns a boolean if a field has been set.

### SetCachingEnabledNil

`func (o *AliasSmbConfig) SetCachingEnabledNil(b bool)`

 SetCachingEnabledNil sets the value for CachingEnabled to be an explicit nil

### UnsetCachingEnabled
`func (o *AliasSmbConfig) UnsetCachingEnabled()`

UnsetCachingEnabled ensures that no value is present for CachingEnabled, not even an explicit nil
### GetIsShareLevelPermissionEmpty

`func (o *AliasSmbConfig) GetIsShareLevelPermissionEmpty() bool`

GetIsShareLevelPermissionEmpty returns the IsShareLevelPermissionEmpty field if non-nil, zero value otherwise.

### GetIsShareLevelPermissionEmptyOk

`func (o *AliasSmbConfig) GetIsShareLevelPermissionEmptyOk() (*bool, bool)`

GetIsShareLevelPermissionEmptyOk returns a tuple with the IsShareLevelPermissionEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareLevelPermissionEmpty

`func (o *AliasSmbConfig) SetIsShareLevelPermissionEmpty(v bool)`

SetIsShareLevelPermissionEmpty sets IsShareLevelPermissionEmpty field to given value.

### HasIsShareLevelPermissionEmpty

`func (o *AliasSmbConfig) HasIsShareLevelPermissionEmpty() bool`

HasIsShareLevelPermissionEmpty returns a boolean if a field has been set.

### SetIsShareLevelPermissionEmptyNil

`func (o *AliasSmbConfig) SetIsShareLevelPermissionEmptyNil(b bool)`

 SetIsShareLevelPermissionEmptyNil sets the value for IsShareLevelPermissionEmpty to be an explicit nil

### UnsetIsShareLevelPermissionEmpty
`func (o *AliasSmbConfig) UnsetIsShareLevelPermissionEmpty()`

UnsetIsShareLevelPermissionEmpty ensures that no value is present for IsShareLevelPermissionEmpty, not even an explicit nil
### GetOplockEnabled

`func (o *AliasSmbConfig) GetOplockEnabled() bool`

GetOplockEnabled returns the OplockEnabled field if non-nil, zero value otherwise.

### GetOplockEnabledOk

`func (o *AliasSmbConfig) GetOplockEnabledOk() (*bool, bool)`

GetOplockEnabledOk returns a tuple with the OplockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplockEnabled

`func (o *AliasSmbConfig) SetOplockEnabled(v bool)`

SetOplockEnabled sets OplockEnabled field to given value.

### HasOplockEnabled

`func (o *AliasSmbConfig) HasOplockEnabled() bool`

HasOplockEnabled returns a boolean if a field has been set.

### SetOplockEnabledNil

`func (o *AliasSmbConfig) SetOplockEnabledNil(b bool)`

 SetOplockEnabledNil sets the value for OplockEnabled to be an explicit nil

### UnsetOplockEnabled
`func (o *AliasSmbConfig) UnsetOplockEnabled()`

UnsetOplockEnabled ensures that no value is present for OplockEnabled, not even an explicit nil
### GetContinuousAvailability

`func (o *AliasSmbConfig) GetContinuousAvailability() bool`

GetContinuousAvailability returns the ContinuousAvailability field if non-nil, zero value otherwise.

### GetContinuousAvailabilityOk

`func (o *AliasSmbConfig) GetContinuousAvailabilityOk() (*bool, bool)`

GetContinuousAvailabilityOk returns a tuple with the ContinuousAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousAvailability

`func (o *AliasSmbConfig) SetContinuousAvailability(v bool)`

SetContinuousAvailability sets ContinuousAvailability field to given value.

### HasContinuousAvailability

`func (o *AliasSmbConfig) HasContinuousAvailability() bool`

HasContinuousAvailability returns a boolean if a field has been set.

### SetContinuousAvailabilityNil

`func (o *AliasSmbConfig) SetContinuousAvailabilityNil(b bool)`

 SetContinuousAvailabilityNil sets the value for ContinuousAvailability to be an explicit nil

### UnsetContinuousAvailability
`func (o *AliasSmbConfig) UnsetContinuousAvailability()`

UnsetContinuousAvailability ensures that no value is present for ContinuousAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


