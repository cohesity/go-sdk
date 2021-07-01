# TestIdpReachability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerId** | Pointer to **NullableString** | Specifies the IdP provided Issuer ID for the app. | [optional] 
**SsoUrl** | Pointer to **NullableString** | Specifies the SSO URL of the IdP service for the customer. This is the URL given by IdP when the customer created an account. Customers may use this for several clusters that are registered with on IdP site. | [optional] 

## Methods

### NewTestIdpReachability

`func NewTestIdpReachability() *TestIdpReachability`

NewTestIdpReachability instantiates a new TestIdpReachability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestIdpReachabilityWithDefaults

`func NewTestIdpReachabilityWithDefaults() *TestIdpReachability`

NewTestIdpReachabilityWithDefaults instantiates a new TestIdpReachability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerId

`func (o *TestIdpReachability) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *TestIdpReachability) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *TestIdpReachability) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *TestIdpReachability) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *TestIdpReachability) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *TestIdpReachability) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetSsoUrl

`func (o *TestIdpReachability) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *TestIdpReachability) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *TestIdpReachability) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *TestIdpReachability) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### SetSsoUrlNil

`func (o *TestIdpReachability) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *TestIdpReachability) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


