# KeystoneCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminCreds** | [**KeystoneAdminParams**](KeystoneAdminParams.md) | Specifies parameters related to Keystone administrator. | 
**Scope** | [**KeystoneScopeParams**](KeystoneScopeParams.md) | Specifies parameters related to Keystone scope. | 

## Methods

### NewKeystoneCredentials

`func NewKeystoneCredentials(adminCreds KeystoneAdminParams, scope KeystoneScopeParams, ) *KeystoneCredentials`

NewKeystoneCredentials instantiates a new KeystoneCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneCredentialsWithDefaults

`func NewKeystoneCredentialsWithDefaults() *KeystoneCredentials`

NewKeystoneCredentialsWithDefaults instantiates a new KeystoneCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminCreds

`func (o *KeystoneCredentials) GetAdminCreds() KeystoneAdminParams`

GetAdminCreds returns the AdminCreds field if non-nil, zero value otherwise.

### GetAdminCredsOk

`func (o *KeystoneCredentials) GetAdminCredsOk() (*KeystoneAdminParams, bool)`

GetAdminCredsOk returns a tuple with the AdminCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCreds

`func (o *KeystoneCredentials) SetAdminCreds(v KeystoneAdminParams)`

SetAdminCreds sets AdminCreds field to given value.


### GetScope

`func (o *KeystoneCredentials) GetScope() KeystoneScopeParams`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *KeystoneCredentials) GetScopeOk() (*KeystoneScopeParams, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *KeystoneCredentials) SetScope(v KeystoneScopeParams)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


