# ACLProtoGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grantee** | Pointer to [**GranteeProto**](GranteeProto.md) |  | [optional] 
**PermissionVec** | Pointer to **[]int32** | Vector of permission granted to this grantee. | [optional] 

## Methods

### NewACLProtoGrant

`func NewACLProtoGrant() *ACLProtoGrant`

NewACLProtoGrant instantiates a new ACLProtoGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLProtoGrantWithDefaults

`func NewACLProtoGrantWithDefaults() *ACLProtoGrant`

NewACLProtoGrantWithDefaults instantiates a new ACLProtoGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantee

`func (o *ACLProtoGrant) GetGrantee() GranteeProto`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *ACLProtoGrant) GetGranteeOk() (*GranteeProto, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *ACLProtoGrant) SetGrantee(v GranteeProto)`

SetGrantee sets Grantee field to given value.

### HasGrantee

`func (o *ACLProtoGrant) HasGrantee() bool`

HasGrantee returns a boolean if a field has been set.

### GetPermissionVec

`func (o *ACLProtoGrant) GetPermissionVec() []int32`

GetPermissionVec returns the PermissionVec field if non-nil, zero value otherwise.

### GetPermissionVecOk

`func (o *ACLProtoGrant) GetPermissionVecOk() (*[]int32, bool)`

GetPermissionVecOk returns a tuple with the PermissionVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionVec

`func (o *ACLProtoGrant) SetPermissionVec(v []int32)`

SetPermissionVec sets PermissionVec field to given value.

### HasPermissionVec

`func (o *ACLProtoGrant) HasPermissionVec() bool`

HasPermissionVec returns a boolean if a field has been set.

### SetPermissionVecNil

`func (o *ACLProtoGrant) SetPermissionVecNil(b bool)`

 SetPermissionVecNil sets the value for PermissionVec to be an explicit nil

### UnsetPermissionVec
`func (o *ACLProtoGrant) UnsetPermissionVec()`

UnsetPermissionVec ensures that no value is present for PermissionVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


