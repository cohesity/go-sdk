# AclGrantGrantee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **NullableString** | Specifies the group to which permissions are granted if the &#x60;type&#x60; is Group. | [optional] 
**Type** | **NullableString** | Specifies the grantee type. | 
**UserId** | Pointer to **NullableString** | Specifies the user id of the grantee if the &#x60;type&#x60; is Registered User. | [optional] 

## Methods

### NewAclGrantGrantee

`func NewAclGrantGrantee(type_ NullableString, ) *AclGrantGrantee`

NewAclGrantGrantee instantiates a new AclGrantGrantee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclGrantGranteeWithDefaults

`func NewAclGrantGranteeWithDefaults() *AclGrantGrantee`

NewAclGrantGranteeWithDefaults instantiates a new AclGrantGrantee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AclGrantGrantee) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AclGrantGrantee) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AclGrantGrantee) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AclGrantGrantee) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *AclGrantGrantee) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *AclGrantGrantee) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetType

`func (o *AclGrantGrantee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AclGrantGrantee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AclGrantGrantee) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AclGrantGrantee) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AclGrantGrantee) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserId

`func (o *AclGrantGrantee) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AclGrantGrantee) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AclGrantGrantee) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AclGrantGrantee) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AclGrantGrantee) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AclGrantGrantee) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


