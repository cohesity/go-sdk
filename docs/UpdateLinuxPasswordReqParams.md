# UpdateLinuxPasswordReqParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinuxCurrentPassword** | Pointer to **NullableString** | Specifies the current password. | [optional] 
**LinuxPassword** | **NullableString** | Specifies the new linux password. | 
**LinuxUsername** | **NullableString** | Specifies the linux username for which the password will be updated. | 

## Methods

### NewUpdateLinuxPasswordReqParams

`func NewUpdateLinuxPasswordReqParams(linuxPassword NullableString, linuxUsername NullableString, ) *UpdateLinuxPasswordReqParams`

NewUpdateLinuxPasswordReqParams instantiates a new UpdateLinuxPasswordReqParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLinuxPasswordReqParamsWithDefaults

`func NewUpdateLinuxPasswordReqParamsWithDefaults() *UpdateLinuxPasswordReqParams`

NewUpdateLinuxPasswordReqParamsWithDefaults instantiates a new UpdateLinuxPasswordReqParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinuxCurrentPassword

`func (o *UpdateLinuxPasswordReqParams) GetLinuxCurrentPassword() string`

GetLinuxCurrentPassword returns the LinuxCurrentPassword field if non-nil, zero value otherwise.

### GetLinuxCurrentPasswordOk

`func (o *UpdateLinuxPasswordReqParams) GetLinuxCurrentPasswordOk() (*string, bool)`

GetLinuxCurrentPasswordOk returns a tuple with the LinuxCurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxCurrentPassword

`func (o *UpdateLinuxPasswordReqParams) SetLinuxCurrentPassword(v string)`

SetLinuxCurrentPassword sets LinuxCurrentPassword field to given value.

### HasLinuxCurrentPassword

`func (o *UpdateLinuxPasswordReqParams) HasLinuxCurrentPassword() bool`

HasLinuxCurrentPassword returns a boolean if a field has been set.

### SetLinuxCurrentPasswordNil

`func (o *UpdateLinuxPasswordReqParams) SetLinuxCurrentPasswordNil(b bool)`

 SetLinuxCurrentPasswordNil sets the value for LinuxCurrentPassword to be an explicit nil

### UnsetLinuxCurrentPassword
`func (o *UpdateLinuxPasswordReqParams) UnsetLinuxCurrentPassword()`

UnsetLinuxCurrentPassword ensures that no value is present for LinuxCurrentPassword, not even an explicit nil
### GetLinuxPassword

`func (o *UpdateLinuxPasswordReqParams) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UpdateLinuxPasswordReqParams) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UpdateLinuxPasswordReqParams) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.


### SetLinuxPasswordNil

`func (o *UpdateLinuxPasswordReqParams) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *UpdateLinuxPasswordReqParams) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxUsername

`func (o *UpdateLinuxPasswordReqParams) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UpdateLinuxPasswordReqParams) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UpdateLinuxPasswordReqParams) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.


### SetLinuxUsernameNil

`func (o *UpdateLinuxPasswordReqParams) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *UpdateLinuxPasswordReqParams) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


