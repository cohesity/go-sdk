# SshPasswordCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | SSH password. | 
**Username** | **string** | SSH username. | 

## Methods

### NewSshPasswordCredentials

`func NewSshPasswordCredentials(password string, username string, ) *SshPasswordCredentials`

NewSshPasswordCredentials instantiates a new SshPasswordCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshPasswordCredentialsWithDefaults

`func NewSshPasswordCredentialsWithDefaults() *SshPasswordCredentials`

NewSshPasswordCredentialsWithDefaults instantiates a new SshPasswordCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SshPasswordCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SshPasswordCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SshPasswordCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *SshPasswordCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SshPasswordCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SshPasswordCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


