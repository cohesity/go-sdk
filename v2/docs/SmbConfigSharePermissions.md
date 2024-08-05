# SmbConfigSharePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share permissions. | [optional] 
**SuperUserSids** | Pointer to **[]string** | Specifies a list of super user sids. | [optional] 

## Methods

### NewSmbConfigSharePermissions

`func NewSmbConfigSharePermissions() *SmbConfigSharePermissions`

NewSmbConfigSharePermissions instantiates a new SmbConfigSharePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbConfigSharePermissionsWithDefaults

`func NewSmbConfigSharePermissionsWithDefaults() *SmbConfigSharePermissions`

NewSmbConfigSharePermissionsWithDefaults instantiates a new SmbConfigSharePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *SmbConfigSharePermissions) GetPermissions() []SmbPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SmbConfigSharePermissions) GetPermissionsOk() (*[]SmbPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SmbConfigSharePermissions) SetPermissions(v []SmbPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SmbConfigSharePermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SmbConfigSharePermissions) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SmbConfigSharePermissions) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSuperUserSids

`func (o *SmbConfigSharePermissions) GetSuperUserSids() []string`

GetSuperUserSids returns the SuperUserSids field if non-nil, zero value otherwise.

### GetSuperUserSidsOk

`func (o *SmbConfigSharePermissions) GetSuperUserSidsOk() (*[]string, bool)`

GetSuperUserSidsOk returns a tuple with the SuperUserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUserSids

`func (o *SmbConfigSharePermissions) SetSuperUserSids(v []string)`

SetSuperUserSids sets SuperUserSids field to given value.

### HasSuperUserSids

`func (o *SmbConfigSharePermissions) HasSuperUserSids() bool`

HasSuperUserSids returns a boolean if a field has been set.

### SetSuperUserSidsNil

`func (o *SmbConfigSharePermissions) SetSuperUserSidsNil(b bool)`

 SetSuperUserSidsNil sets the value for SuperUserSids to be an explicit nil

### UnsetSuperUserSids
`func (o *SmbConfigSharePermissions) UnsetSuperUserSids()`

UnsetSuperUserSids ensures that no value is present for SuperUserSids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


