# AccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **NullableString** | Created access token. | [optional] 
**Privileges** | Pointer to **[]string** | Privileges for the user. | [optional] 
**TokenType** | Pointer to **NullableString** | Access token type. | [optional] 

## Methods

### NewAccessTokenResponse

`func NewAccessTokenResponse() *AccessTokenResponse`

NewAccessTokenResponse instantiates a new AccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenResponseWithDefaults

`func NewAccessTokenResponseWithDefaults() *AccessTokenResponse`

NewAccessTokenResponseWithDefaults instantiates a new AccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccessTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AccessTokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *AccessTokenResponse) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AccessTokenResponse) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetPrivileges

`func (o *AccessTokenResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *AccessTokenResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *AccessTokenResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *AccessTokenResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetTokenType

`func (o *AccessTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AccessTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *AccessTokenResponse) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *AccessTokenResponse) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


