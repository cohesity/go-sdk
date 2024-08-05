# UpdateLinuxPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **NullableString** | Specifies the current password of the user. This is required when trying to update the current user&#39;s password. | [optional] 
**NewPassword** | Pointer to **NullableString** | Specifies the new linux password. | [optional] 
**Username** | **NullableString** | Specifies the linux username for which the password will be updated. | 
**VerifyPassword** | Pointer to **NullableBool** | True if request is only to verify if current password matches with set password. | [optional] 

## Methods

### NewUpdateLinuxPasswordRequest

`func NewUpdateLinuxPasswordRequest(username NullableString, ) *UpdateLinuxPasswordRequest`

NewUpdateLinuxPasswordRequest instantiates a new UpdateLinuxPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLinuxPasswordRequestWithDefaults

`func NewUpdateLinuxPasswordRequestWithDefaults() *UpdateLinuxPasswordRequest`

NewUpdateLinuxPasswordRequestWithDefaults instantiates a new UpdateLinuxPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdateLinuxPasswordRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateLinuxPasswordRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateLinuxPasswordRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateLinuxPasswordRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *UpdateLinuxPasswordRequest) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *UpdateLinuxPasswordRequest) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetNewPassword

`func (o *UpdateLinuxPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateLinuxPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateLinuxPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateLinuxPasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *UpdateLinuxPasswordRequest) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *UpdateLinuxPasswordRequest) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
### GetUsername

`func (o *UpdateLinuxPasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateLinuxPasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateLinuxPasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *UpdateLinuxPasswordRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateLinuxPasswordRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetVerifyPassword

`func (o *UpdateLinuxPasswordRequest) GetVerifyPassword() bool`

GetVerifyPassword returns the VerifyPassword field if non-nil, zero value otherwise.

### GetVerifyPasswordOk

`func (o *UpdateLinuxPasswordRequest) GetVerifyPasswordOk() (*bool, bool)`

GetVerifyPasswordOk returns a tuple with the VerifyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPassword

`func (o *UpdateLinuxPasswordRequest) SetVerifyPassword(v bool)`

SetVerifyPassword sets VerifyPassword field to given value.

### HasVerifyPassword

`func (o *UpdateLinuxPasswordRequest) HasVerifyPassword() bool`

HasVerifyPassword returns a boolean if a field has been set.

### SetVerifyPasswordNil

`func (o *UpdateLinuxPasswordRequest) SetVerifyPasswordNil(b bool)`

 SetVerifyPasswordNil sets the value for VerifyPassword to be an explicit nil

### UnsetVerifyPassword
`func (o *UpdateLinuxPasswordRequest) UnsetVerifyPassword()`

UnsetVerifyPassword ensures that no value is present for VerifyPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


