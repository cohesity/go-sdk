# SshPrivateKeyCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passphrase** | Pointer to **string** | Passphrase for the private key. | [optional] 
**PrivateKey** | **string** | The private key. | 
**UserId** | **string** | userId for PrivateKey credentials. | 

## Methods

### NewSshPrivateKeyCredentials

`func NewSshPrivateKeyCredentials(privateKey string, userId string, ) *SshPrivateKeyCredentials`

NewSshPrivateKeyCredentials instantiates a new SshPrivateKeyCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshPrivateKeyCredentialsWithDefaults

`func NewSshPrivateKeyCredentialsWithDefaults() *SshPrivateKeyCredentials`

NewSshPrivateKeyCredentialsWithDefaults instantiates a new SshPrivateKeyCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassphrase

`func (o *SshPrivateKeyCredentials) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *SshPrivateKeyCredentials) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *SshPrivateKeyCredentials) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *SshPrivateKeyCredentials) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SshPrivateKeyCredentials) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SshPrivateKeyCredentials) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SshPrivateKeyCredentials) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetUserId

`func (o *SshPrivateKeyCredentials) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SshPrivateKeyCredentials) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SshPrivateKeyCredentials) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


