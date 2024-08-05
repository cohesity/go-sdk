# AclGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grantee** | [**AclGrantGrantee**](AclGrantGrantee.md) |  | 
**Permissions** | **[]string** | Specifies a list of permissions granted to the grantees. | 

## Methods

### NewAclGrant

`func NewAclGrant(grantee AclGrantGrantee, permissions []string, ) *AclGrant`

NewAclGrant instantiates a new AclGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclGrantWithDefaults

`func NewAclGrantWithDefaults() *AclGrant`

NewAclGrantWithDefaults instantiates a new AclGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantee

`func (o *AclGrant) GetGrantee() AclGrantGrantee`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *AclGrant) GetGranteeOk() (*AclGrantGrantee, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *AclGrant) SetGrantee(v AclGrantGrantee)`

SetGrantee sets Grantee field to given value.


### GetPermissions

`func (o *AclGrant) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AclGrant) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AclGrant) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### SetPermissionsNil

`func (o *AclGrant) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AclGrant) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


