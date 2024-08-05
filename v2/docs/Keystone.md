# Keystone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminCreds** | [**KeystoneCredentialsAdminCreds**](KeystoneCredentialsAdminCreds.md) |  | 
**Scope** | [**KeystoneCredentialsScope**](KeystoneCredentialsScope.md) |  | 
**AuthUrl** | **NullableString** | Specifies the url points to the Keystone service. | 
**Id** | Pointer to **NullableInt64** | Specifies the Keystone configuration id. | [optional] [readonly] 
**Name** | **NullableString** | Specifies the Keystone configuration name. | 

## Methods

### NewKeystone

`func NewKeystone(adminCreds KeystoneCredentialsAdminCreds, scope KeystoneCredentialsScope, authUrl NullableString, name NullableString, ) *Keystone`

NewKeystone instantiates a new Keystone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneWithDefaults

`func NewKeystoneWithDefaults() *Keystone`

NewKeystoneWithDefaults instantiates a new Keystone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminCreds

`func (o *Keystone) GetAdminCreds() KeystoneCredentialsAdminCreds`

GetAdminCreds returns the AdminCreds field if non-nil, zero value otherwise.

### GetAdminCredsOk

`func (o *Keystone) GetAdminCredsOk() (*KeystoneCredentialsAdminCreds, bool)`

GetAdminCredsOk returns a tuple with the AdminCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCreds

`func (o *Keystone) SetAdminCreds(v KeystoneCredentialsAdminCreds)`

SetAdminCreds sets AdminCreds field to given value.


### GetScope

`func (o *Keystone) GetScope() KeystoneCredentialsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Keystone) GetScopeOk() (*KeystoneCredentialsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Keystone) SetScope(v KeystoneCredentialsScope)`

SetScope sets Scope field to given value.


### GetAuthUrl

`func (o *Keystone) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *Keystone) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *Keystone) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### SetAuthUrlNil

`func (o *Keystone) SetAuthUrlNil(b bool)`

 SetAuthUrlNil sets the value for AuthUrl to be an explicit nil

### UnsetAuthUrl
`func (o *Keystone) UnsetAuthUrl()`

UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil
### GetId

`func (o *Keystone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Keystone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Keystone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Keystone) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Keystone) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Keystone) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Keystone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Keystone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Keystone) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Keystone) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Keystone) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


