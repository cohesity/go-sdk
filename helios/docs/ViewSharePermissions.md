# ViewSharePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuperUserSids** | Pointer to **[]string** | Specifies a list of super user sids. | [optional] 
**Permissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share permissions. | [optional] 

## Methods

### NewViewSharePermissions

`func NewViewSharePermissions() *ViewSharePermissions`

NewViewSharePermissions instantiates a new ViewSharePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewSharePermissionsWithDefaults

`func NewViewSharePermissionsWithDefaults() *ViewSharePermissions`

NewViewSharePermissionsWithDefaults instantiates a new ViewSharePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuperUserSids

`func (o *ViewSharePermissions) GetSuperUserSids() []string`

GetSuperUserSids returns the SuperUserSids field if non-nil, zero value otherwise.

### GetSuperUserSidsOk

`func (o *ViewSharePermissions) GetSuperUserSidsOk() (*[]string, bool)`

GetSuperUserSidsOk returns a tuple with the SuperUserSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUserSids

`func (o *ViewSharePermissions) SetSuperUserSids(v []string)`

SetSuperUserSids sets SuperUserSids field to given value.

### HasSuperUserSids

`func (o *ViewSharePermissions) HasSuperUserSids() bool`

HasSuperUserSids returns a boolean if a field has been set.

### SetSuperUserSidsNil

`func (o *ViewSharePermissions) SetSuperUserSidsNil(b bool)`

 SetSuperUserSidsNil sets the value for SuperUserSids to be an explicit nil

### UnsetSuperUserSids
`func (o *ViewSharePermissions) UnsetSuperUserSids()`

UnsetSuperUserSids ensures that no value is present for SuperUserSids, not even an explicit nil
### GetPermissions

`func (o *ViewSharePermissions) GetPermissions() []SmbPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ViewSharePermissions) GetPermissionsOk() (*[]SmbPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ViewSharePermissions) SetPermissions(v []SmbPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ViewSharePermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ViewSharePermissions) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ViewSharePermissions) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


