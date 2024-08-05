# ActiveDirectoryAdminParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **NullableString** | Specifies the password of the user. | 
**Username** | **NullableString** | Specifies the user name. | 

## Methods

### NewActiveDirectoryAdminParams

`func NewActiveDirectoryAdminParams(password NullableString, username NullableString, ) *ActiveDirectoryAdminParams`

NewActiveDirectoryAdminParams instantiates a new ActiveDirectoryAdminParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryAdminParamsWithDefaults

`func NewActiveDirectoryAdminParamsWithDefaults() *ActiveDirectoryAdminParams`

NewActiveDirectoryAdminParamsWithDefaults instantiates a new ActiveDirectoryAdminParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ActiveDirectoryAdminParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ActiveDirectoryAdminParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ActiveDirectoryAdminParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *ActiveDirectoryAdminParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ActiveDirectoryAdminParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUsername

`func (o *ActiveDirectoryAdminParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActiveDirectoryAdminParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActiveDirectoryAdminParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *ActiveDirectoryAdminParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ActiveDirectoryAdminParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


