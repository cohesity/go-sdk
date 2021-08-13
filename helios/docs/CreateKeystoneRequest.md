# CreateKeystoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminCreds** | [**KeystoneAdminParams**](KeystoneAdminParams.md) | Specifies parameters related to Keystone administrator. | 
**Scope** | [**KeystoneScopeParams**](KeystoneScopeParams.md) | Specifies parameters related to Keystone scope. | 
**Name** | **NullableString** | Specifies the Keystone configuration name. | 
**Id** | Pointer to **NullableInt64** | Specifies the Keystone configuration id. | [optional] [readonly] 
**AuthUrl** | **NullableString** | Specifies the url points to the Keystone service. | 

## Methods

### NewCreateKeystoneRequest

`func NewCreateKeystoneRequest(adminCreds KeystoneAdminParams, scope KeystoneScopeParams, name NullableString, authUrl NullableString, ) *CreateKeystoneRequest`

NewCreateKeystoneRequest instantiates a new CreateKeystoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeystoneRequestWithDefaults

`func NewCreateKeystoneRequestWithDefaults() *CreateKeystoneRequest`

NewCreateKeystoneRequestWithDefaults instantiates a new CreateKeystoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminCreds

`func (o *CreateKeystoneRequest) GetAdminCreds() KeystoneAdminParams`

GetAdminCreds returns the AdminCreds field if non-nil, zero value otherwise.

### GetAdminCredsOk

`func (o *CreateKeystoneRequest) GetAdminCredsOk() (*KeystoneAdminParams, bool)`

GetAdminCredsOk returns a tuple with the AdminCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminCreds

`func (o *CreateKeystoneRequest) SetAdminCreds(v KeystoneAdminParams)`

SetAdminCreds sets AdminCreds field to given value.


### GetScope

`func (o *CreateKeystoneRequest) GetScope() KeystoneScopeParams`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateKeystoneRequest) GetScopeOk() (*KeystoneScopeParams, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateKeystoneRequest) SetScope(v KeystoneScopeParams)`

SetScope sets Scope field to given value.


### GetName

`func (o *CreateKeystoneRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeystoneRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeystoneRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateKeystoneRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateKeystoneRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *CreateKeystoneRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateKeystoneRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateKeystoneRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateKeystoneRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateKeystoneRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateKeystoneRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAuthUrl

`func (o *CreateKeystoneRequest) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *CreateKeystoneRequest) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *CreateKeystoneRequest) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### SetAuthUrlNil

`func (o *CreateKeystoneRequest) SetAuthUrlNil(b bool)`

 SetAuthUrlNil sets the value for AuthUrl to be an explicit nil

### UnsetAuthUrl
`func (o *CreateKeystoneRequest) UnsetAuthUrl()`

UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


