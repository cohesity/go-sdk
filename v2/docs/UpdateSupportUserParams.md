# UpdateSupportUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **NullableString** | Specifies the current password of the user. This is required when trying to update the current user&#39;s password. | [optional] 
**EnableSudoAccess** | Pointer to **NullableBool** | If set to true, sudo access will be enabled for the user. If null, the endpoint will not attempt to alter sudo access privilege for the support user. | [optional] 
**NewPassword** | Pointer to **NullableString** | Specifies the new password for the support user. | [optional] 

## Methods

### NewUpdateSupportUserParams

`func NewUpdateSupportUserParams() *UpdateSupportUserParams`

NewUpdateSupportUserParams instantiates a new UpdateSupportUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSupportUserParamsWithDefaults

`func NewUpdateSupportUserParamsWithDefaults() *UpdateSupportUserParams`

NewUpdateSupportUserParamsWithDefaults instantiates a new UpdateSupportUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdateSupportUserParams) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateSupportUserParams) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateSupportUserParams) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateSupportUserParams) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *UpdateSupportUserParams) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *UpdateSupportUserParams) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetEnableSudoAccess

`func (o *UpdateSupportUserParams) GetEnableSudoAccess() bool`

GetEnableSudoAccess returns the EnableSudoAccess field if non-nil, zero value otherwise.

### GetEnableSudoAccessOk

`func (o *UpdateSupportUserParams) GetEnableSudoAccessOk() (*bool, bool)`

GetEnableSudoAccessOk returns a tuple with the EnableSudoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSudoAccess

`func (o *UpdateSupportUserParams) SetEnableSudoAccess(v bool)`

SetEnableSudoAccess sets EnableSudoAccess field to given value.

### HasEnableSudoAccess

`func (o *UpdateSupportUserParams) HasEnableSudoAccess() bool`

HasEnableSudoAccess returns a boolean if a field has been set.

### SetEnableSudoAccessNil

`func (o *UpdateSupportUserParams) SetEnableSudoAccessNil(b bool)`

 SetEnableSudoAccessNil sets the value for EnableSudoAccess to be an explicit nil

### UnsetEnableSudoAccess
`func (o *UpdateSupportUserParams) UnsetEnableSudoAccess()`

UnsetEnableSudoAccess ensures that no value is present for EnableSudoAccess, not even an explicit nil
### GetNewPassword

`func (o *UpdateSupportUserParams) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateSupportUserParams) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateSupportUserParams) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateSupportUserParams) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *UpdateSupportUserParams) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *UpdateSupportUserParams) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


