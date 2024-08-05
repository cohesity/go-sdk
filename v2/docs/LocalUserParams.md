# LocalUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Specifies the email address of the User. | [optional] 
**Groups** | Pointer to **[]string** | Specifies additional groups the User may belong to. | [optional] [readonly] 
**Password** | Pointer to **NullableString** | Specifies the password of the User. | [optional] 
**PrimaryGroup** | Pointer to **NullableString** | Specifies the primary group of the User. Primary group is used for file access. | [optional] [readonly] 

## Methods

### NewLocalUserParams

`func NewLocalUserParams() *LocalUserParams`

NewLocalUserParams instantiates a new LocalUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalUserParamsWithDefaults

`func NewLocalUserParamsWithDefaults() *LocalUserParams`

NewLocalUserParamsWithDefaults instantiates a new LocalUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LocalUserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocalUserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocalUserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocalUserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocalUserParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocalUserParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroups

`func (o *LocalUserParams) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LocalUserParams) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LocalUserParams) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *LocalUserParams) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPassword

`func (o *LocalUserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LocalUserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LocalUserParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LocalUserParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *LocalUserParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LocalUserParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPrimaryGroup

`func (o *LocalUserParams) GetPrimaryGroup() string`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *LocalUserParams) GetPrimaryGroupOk() (*string, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *LocalUserParams) SetPrimaryGroup(v string)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *LocalUserParams) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### SetPrimaryGroupNil

`func (o *LocalUserParams) SetPrimaryGroupNil(b bool)`

 SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil

### UnsetPrimaryGroup
`func (o *LocalUserParams) UnsetPrimaryGroup()`

UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


