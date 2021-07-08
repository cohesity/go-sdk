# UserSshKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKey** | Pointer to **NullableString** | Specifies SSH key needed to be added to the username passed. | [optional] 
**Username** | Pointer to **NullableString** | Specifies name of the user to add. | [optional] 

## Methods

### NewUserSshKey

`func NewUserSshKey() *UserSshKey`

NewUserSshKey instantiates a new UserSshKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSshKeyWithDefaults

`func NewUserSshKeyWithDefaults() *UserSshKey`

NewUserSshKeyWithDefaults instantiates a new UserSshKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKey

`func (o *UserSshKey) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *UserSshKey) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *UserSshKey) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *UserSshKey) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *UserSshKey) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *UserSshKey) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetUsername

`func (o *UserSshKey) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSshKey) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSshKey) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSshKey) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserSshKey) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserSshKey) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


