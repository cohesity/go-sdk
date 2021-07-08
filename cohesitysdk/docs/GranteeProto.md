# GranteeProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **NullableString** | If grantee is of type &#39;kEmailUser&#39;, this field will contain the email address of the user. | [optional] 
**Group** | Pointer to **NullableInt32** | If grantee is of type &#39;kGroup&#39;, this field will contain the group to which permission is granted. | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**UserId** | Pointer to **NullableString** | If grantee is of type &#39;kRegisteredUser&#39;, this field will contain the canonical id of the user. | [optional] 

## Methods

### NewGranteeProto

`func NewGranteeProto() *GranteeProto`

NewGranteeProto instantiates a new GranteeProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGranteeProtoWithDefaults

`func NewGranteeProtoWithDefaults() *GranteeProto`

NewGranteeProtoWithDefaults instantiates a new GranteeProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *GranteeProto) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *GranteeProto) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *GranteeProto) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *GranteeProto) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *GranteeProto) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *GranteeProto) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetGroup

`func (o *GranteeProto) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GranteeProto) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GranteeProto) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GranteeProto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *GranteeProto) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *GranteeProto) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetType

`func (o *GranteeProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GranteeProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GranteeProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GranteeProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GranteeProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GranteeProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserId

`func (o *GranteeProto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GranteeProto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GranteeProto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GranteeProto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GranteeProto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GranteeProto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


