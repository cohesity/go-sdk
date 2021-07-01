# Office365Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the application ID that the registration portal (apps.dev.microsoft.com) assigned. | [optional] 
**ClientSecret** | Pointer to **NullableString** | Specifies the application secret that was created in app registration portal. | [optional] 
**GrantType** | Pointer to **NullableString** | Specifies the application grant type. eg: For client credentials flow, set this to \&quot;client_credentials\&quot;; For refreshing access-token, set this to \&quot;refresh_token\&quot;. | [optional] 
**Scope** | Pointer to **NullableString** | Specifies a space separated list of scopes/permissions for the user. eg: Incase of MS Graph APIs for Office365, scope is set to default: https://graph.microsoft.com/.default | [optional] 
**UseOAuthForExchangeOnline** | Pointer to **NullableBool** | This field is deprecated from here and placed in RegisteredSourceInfo and ProtectionSourceParameters. deprecated: true | [optional] 

## Methods

### NewOffice365Credentials

`func NewOffice365Credentials() *Office365Credentials`

NewOffice365Credentials instantiates a new Office365Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365CredentialsWithDefaults

`func NewOffice365CredentialsWithDefaults() *Office365Credentials`

NewOffice365CredentialsWithDefaults instantiates a new Office365Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Office365Credentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Office365Credentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Office365Credentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Office365Credentials) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *Office365Credentials) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *Office365Credentials) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *Office365Credentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Office365Credentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Office365Credentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Office365Credentials) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *Office365Credentials) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *Office365Credentials) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetGrantType

`func (o *Office365Credentials) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Office365Credentials) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Office365Credentials) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Office365Credentials) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### SetGrantTypeNil

`func (o *Office365Credentials) SetGrantTypeNil(b bool)`

 SetGrantTypeNil sets the value for GrantType to be an explicit nil

### UnsetGrantType
`func (o *Office365Credentials) UnsetGrantType()`

UnsetGrantType ensures that no value is present for GrantType, not even an explicit nil
### GetScope

`func (o *Office365Credentials) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Office365Credentials) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Office365Credentials) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Office365Credentials) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Office365Credentials) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Office365Credentials) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetUseOAuthForExchangeOnline

`func (o *Office365Credentials) GetUseOAuthForExchangeOnline() bool`

GetUseOAuthForExchangeOnline returns the UseOAuthForExchangeOnline field if non-nil, zero value otherwise.

### GetUseOAuthForExchangeOnlineOk

`func (o *Office365Credentials) GetUseOAuthForExchangeOnlineOk() (*bool, bool)`

GetUseOAuthForExchangeOnlineOk returns a tuple with the UseOAuthForExchangeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOAuthForExchangeOnline

`func (o *Office365Credentials) SetUseOAuthForExchangeOnline(v bool)`

SetUseOAuthForExchangeOnline sets UseOAuthForExchangeOnline field to given value.

### HasUseOAuthForExchangeOnline

`func (o *Office365Credentials) HasUseOAuthForExchangeOnline() bool`

HasUseOAuthForExchangeOnline returns a boolean if a field has been set.

### SetUseOAuthForExchangeOnlineNil

`func (o *Office365Credentials) SetUseOAuthForExchangeOnlineNil(b bool)`

 SetUseOAuthForExchangeOnlineNil sets the value for UseOAuthForExchangeOnline to be an explicit nil

### UnsetUseOAuthForExchangeOnline
`func (o *Office365Credentials) UnsetUseOAuthForExchangeOnline()`

UnsetUseOAuthForExchangeOnline ensures that no value is present for UseOAuthForExchangeOnline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


