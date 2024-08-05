# S3ConfigAclConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]AclGrant**](AclGrant.md) | Specifies a list of ACL grants. | [optional] 

## Methods

### NewS3ConfigAclConfig

`func NewS3ConfigAclConfig() *S3ConfigAclConfig`

NewS3ConfigAclConfig instantiates a new S3ConfigAclConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigAclConfigWithDefaults

`func NewS3ConfigAclConfigWithDefaults() *S3ConfigAclConfig`

NewS3ConfigAclConfigWithDefaults instantiates a new S3ConfigAclConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *S3ConfigAclConfig) GetGrants() []AclGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *S3ConfigAclConfig) GetGrantsOk() (*[]AclGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *S3ConfigAclConfig) SetGrants(v []AclGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *S3ConfigAclConfig) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### SetGrantsNil

`func (o *S3ConfigAclConfig) SetGrantsNil(b bool)`

 SetGrantsNil sets the value for Grants to be an explicit nil

### UnsetGrants
`func (o *S3ConfigAclConfig) UnsetGrants()`

UnsetGrants ensures that no value is present for Grants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


