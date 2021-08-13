# PrivateKeyCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Specifies the userId to access target entity. | 
**PrivateKey** | **string** | Specifies the private key to access target entity. | 
**Passphrase** | Pointer to **string** | Specifies the passphrase to access target entity. | [optional] 

## Methods

### NewPrivateKeyCredentials

`func NewPrivateKeyCredentials(userId string, privateKey string, ) *PrivateKeyCredentials`

NewPrivateKeyCredentials instantiates a new PrivateKeyCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateKeyCredentialsWithDefaults

`func NewPrivateKeyCredentialsWithDefaults() *PrivateKeyCredentials`

NewPrivateKeyCredentialsWithDefaults instantiates a new PrivateKeyCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PrivateKeyCredentials) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PrivateKeyCredentials) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PrivateKeyCredentials) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPrivateKey

`func (o *PrivateKeyCredentials) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PrivateKeyCredentials) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PrivateKeyCredentials) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPassphrase

`func (o *PrivateKeyCredentials) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *PrivateKeyCredentials) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *PrivateKeyCredentials) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *PrivateKeyCredentials) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


