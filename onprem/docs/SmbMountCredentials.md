# SmbMountCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Specifies the username to access target entity. | 
**Password** | **string** | Specifies the password to access target entity. | 

## Methods

### NewSmbMountCredentials

`func NewSmbMountCredentials(username string, password string, ) *SmbMountCredentials`

NewSmbMountCredentials instantiates a new SmbMountCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbMountCredentialsWithDefaults

`func NewSmbMountCredentialsWithDefaults() *SmbMountCredentials`

NewSmbMountCredentialsWithDefaults instantiates a new SmbMountCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SmbMountCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SmbMountCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SmbMountCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *SmbMountCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmbMountCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmbMountCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


