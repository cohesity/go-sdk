# IdentityAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderType** | **string** | Specifies the type of identity provider the action will be performed on. | 
**OAuth2Params** | Pointer to [**OAuth2Action**](OAuth2Action.md) |  | [optional] 
**OpenIdConnectParams** | Pointer to [**OpenIdConnectAction**](OpenIdConnectAction.md) |  | [optional] 

## Methods

### NewIdentityAction

`func NewIdentityAction(identityProviderType string, ) *IdentityAction`

NewIdentityAction instantiates a new IdentityAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityActionWithDefaults

`func NewIdentityActionWithDefaults() *IdentityAction`

NewIdentityActionWithDefaults instantiates a new IdentityAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviderType

`func (o *IdentityAction) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *IdentityAction) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *IdentityAction) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.


### GetOAuth2Params

`func (o *IdentityAction) GetOAuth2Params() OAuth2Action`

GetOAuth2Params returns the OAuth2Params field if non-nil, zero value otherwise.

### GetOAuth2ParamsOk

`func (o *IdentityAction) GetOAuth2ParamsOk() (*OAuth2Action, bool)`

GetOAuth2ParamsOk returns a tuple with the OAuth2Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2Params

`func (o *IdentityAction) SetOAuth2Params(v OAuth2Action)`

SetOAuth2Params sets OAuth2Params field to given value.

### HasOAuth2Params

`func (o *IdentityAction) HasOAuth2Params() bool`

HasOAuth2Params returns a boolean if a field has been set.

### GetOpenIdConnectParams

`func (o *IdentityAction) GetOpenIdConnectParams() OpenIdConnectAction`

GetOpenIdConnectParams returns the OpenIdConnectParams field if non-nil, zero value otherwise.

### GetOpenIdConnectParamsOk

`func (o *IdentityAction) GetOpenIdConnectParamsOk() (*OpenIdConnectAction, bool)`

GetOpenIdConnectParamsOk returns a tuple with the OpenIdConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIdConnectParams

`func (o *IdentityAction) SetOpenIdConnectParams(v OpenIdConnectAction)`

SetOpenIdConnectParams sets OpenIdConnectParams field to given value.

### HasOpenIdConnectParams

`func (o *IdentityAction) HasOpenIdConnectParams() bool`

HasOpenIdConnectParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


