# LocalUserUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Specifies the email address of the User. | [optional] 
**Groups** | Pointer to **[]string** | Specifies additional groups the User may belong to. | [optional] [readonly] 
**Password** | Pointer to **NullableString** | Specifies the password of the User. | [optional] 
**PrimaryGroup** | Pointer to **NullableString** | Specifies the primary group of the User. Primary group is used for file access. | [optional] [readonly] 
**CurrentPassword** | Pointer to **NullableString** | Specifies the current password of the user. This is required when a session user tries to update his own password. | [optional] 

## Methods

### NewLocalUserUpdateParams

`func NewLocalUserUpdateParams() *LocalUserUpdateParams`

NewLocalUserUpdateParams instantiates a new LocalUserUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalUserUpdateParamsWithDefaults

`func NewLocalUserUpdateParamsWithDefaults() *LocalUserUpdateParams`

NewLocalUserUpdateParamsWithDefaults instantiates a new LocalUserUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LocalUserUpdateParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocalUserUpdateParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocalUserUpdateParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocalUserUpdateParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocalUserUpdateParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocalUserUpdateParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroups

`func (o *LocalUserUpdateParams) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LocalUserUpdateParams) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LocalUserUpdateParams) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *LocalUserUpdateParams) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPassword

`func (o *LocalUserUpdateParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LocalUserUpdateParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LocalUserUpdateParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LocalUserUpdateParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *LocalUserUpdateParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LocalUserUpdateParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPrimaryGroup

`func (o *LocalUserUpdateParams) GetPrimaryGroup() string`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *LocalUserUpdateParams) GetPrimaryGroupOk() (*string, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *LocalUserUpdateParams) SetPrimaryGroup(v string)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *LocalUserUpdateParams) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### SetPrimaryGroupNil

`func (o *LocalUserUpdateParams) SetPrimaryGroupNil(b bool)`

 SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil

### UnsetPrimaryGroup
`func (o *LocalUserUpdateParams) UnsetPrimaryGroup()`

UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil
### GetCurrentPassword

`func (o *LocalUserUpdateParams) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *LocalUserUpdateParams) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *LocalUserUpdateParams) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *LocalUserUpdateParams) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *LocalUserUpdateParams) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *LocalUserUpdateParams) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


