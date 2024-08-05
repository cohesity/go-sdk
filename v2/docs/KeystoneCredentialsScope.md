# KeystoneCredentialsScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainScopeParams** | Pointer to [**KeystoneScopeParamsDomainScopeParams**](KeystoneScopeParamsDomainScopeParams.md) |  | [optional] 
**ProjectScopeParams** | Pointer to [**KeystoneScopeParamsProjectScopeParams**](KeystoneScopeParamsProjectScopeParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the scope type. | 

## Methods

### NewKeystoneCredentialsScope

`func NewKeystoneCredentialsScope(type_ NullableString, ) *KeystoneCredentialsScope`

NewKeystoneCredentialsScope instantiates a new KeystoneCredentialsScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneCredentialsScopeWithDefaults

`func NewKeystoneCredentialsScopeWithDefaults() *KeystoneCredentialsScope`

NewKeystoneCredentialsScopeWithDefaults instantiates a new KeystoneCredentialsScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainScopeParams

`func (o *KeystoneCredentialsScope) GetDomainScopeParams() KeystoneScopeParamsDomainScopeParams`

GetDomainScopeParams returns the DomainScopeParams field if non-nil, zero value otherwise.

### GetDomainScopeParamsOk

`func (o *KeystoneCredentialsScope) GetDomainScopeParamsOk() (*KeystoneScopeParamsDomainScopeParams, bool)`

GetDomainScopeParamsOk returns a tuple with the DomainScopeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainScopeParams

`func (o *KeystoneCredentialsScope) SetDomainScopeParams(v KeystoneScopeParamsDomainScopeParams)`

SetDomainScopeParams sets DomainScopeParams field to given value.

### HasDomainScopeParams

`func (o *KeystoneCredentialsScope) HasDomainScopeParams() bool`

HasDomainScopeParams returns a boolean if a field has been set.

### GetProjectScopeParams

`func (o *KeystoneCredentialsScope) GetProjectScopeParams() KeystoneScopeParamsProjectScopeParams`

GetProjectScopeParams returns the ProjectScopeParams field if non-nil, zero value otherwise.

### GetProjectScopeParamsOk

`func (o *KeystoneCredentialsScope) GetProjectScopeParamsOk() (*KeystoneScopeParamsProjectScopeParams, bool)`

GetProjectScopeParamsOk returns a tuple with the ProjectScopeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectScopeParams

`func (o *KeystoneCredentialsScope) SetProjectScopeParams(v KeystoneScopeParamsProjectScopeParams)`

SetProjectScopeParams sets ProjectScopeParams field to given value.

### HasProjectScopeParams

`func (o *KeystoneCredentialsScope) HasProjectScopeParams() bool`

HasProjectScopeParams returns a boolean if a field has been set.

### GetType

`func (o *KeystoneCredentialsScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeystoneCredentialsScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeystoneCredentialsScope) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *KeystoneCredentialsScope) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KeystoneCredentialsScope) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


