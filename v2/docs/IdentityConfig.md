# IdentityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **NullableString** | Specifies domain of idp configuration | 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the IDP. | [optional] [readonly] 
**IdentityProviderType** | **NullableString** | Specifies the type of identity provider. | 
**IsEnabled** | Pointer to **NullableBool** | Specifies a flag to enable or disable this idp service. When it is set to true, idp service is enabled. When it is set to false, idp service is disabled. By defaut idp is enabled i.e the value is true. | [optional] [default to true]
**LastModifiedTimestampUsecs** | Pointer to **NullableInt64** | Specifies the last time this configuration was modified in microseconds since the epoch. This is may be specified for PUT operations to prevent stale requests from being written. If it is specified during a PUT operation then the request will be rejected if the specified time does not match the actual last modified time. | [optional] 
**OAuth2Params** | Pointer to [**OAuth2Provider**](OAuth2Provider.md) |  | [optional] 
**OpenIdConnectParams** | Pointer to [**OpenIdProvider**](OpenIdProvider.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id if the idp is configured for a tenant. If this is not set, this idp configuration is used for the cluster level users and for all users of tenants not having an idp configuration. | [optional] 

## Methods

### NewIdentityConfig

`func NewIdentityConfig(domain NullableString, identityProviderType NullableString, ) *IdentityConfig`

NewIdentityConfig instantiates a new IdentityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityConfigWithDefaults

`func NewIdentityConfigWithDefaults() *IdentityConfig`

NewIdentityConfigWithDefaults instantiates a new IdentityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdentityConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *IdentityConfig) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityConfig) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetId

`func (o *IdentityConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IdentityConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IdentityConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentityProviderType

`func (o *IdentityConfig) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *IdentityConfig) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *IdentityConfig) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.


### SetIdentityProviderTypeNil

`func (o *IdentityConfig) SetIdentityProviderTypeNil(b bool)`

 SetIdentityProviderTypeNil sets the value for IdentityProviderType to be an explicit nil

### UnsetIdentityProviderType
`func (o *IdentityConfig) UnsetIdentityProviderType()`

UnsetIdentityProviderType ensures that no value is present for IdentityProviderType, not even an explicit nil
### GetIsEnabled

`func (o *IdentityConfig) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IdentityConfig) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IdentityConfig) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IdentityConfig) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *IdentityConfig) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *IdentityConfig) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetLastModifiedTimestampUsecs

`func (o *IdentityConfig) GetLastModifiedTimestampUsecs() int64`

GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimestampUsecsOk

`func (o *IdentityConfig) GetLastModifiedTimestampUsecsOk() (*int64, bool)`

GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestampUsecs

`func (o *IdentityConfig) SetLastModifiedTimestampUsecs(v int64)`

SetLastModifiedTimestampUsecs sets LastModifiedTimestampUsecs field to given value.

### HasLastModifiedTimestampUsecs

`func (o *IdentityConfig) HasLastModifiedTimestampUsecs() bool`

HasLastModifiedTimestampUsecs returns a boolean if a field has been set.

### SetLastModifiedTimestampUsecsNil

`func (o *IdentityConfig) SetLastModifiedTimestampUsecsNil(b bool)`

 SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil

### UnsetLastModifiedTimestampUsecs
`func (o *IdentityConfig) UnsetLastModifiedTimestampUsecs()`

UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
### GetOAuth2Params

`func (o *IdentityConfig) GetOAuth2Params() OAuth2Provider`

GetOAuth2Params returns the OAuth2Params field if non-nil, zero value otherwise.

### GetOAuth2ParamsOk

`func (o *IdentityConfig) GetOAuth2ParamsOk() (*OAuth2Provider, bool)`

GetOAuth2ParamsOk returns a tuple with the OAuth2Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2Params

`func (o *IdentityConfig) SetOAuth2Params(v OAuth2Provider)`

SetOAuth2Params sets OAuth2Params field to given value.

### HasOAuth2Params

`func (o *IdentityConfig) HasOAuth2Params() bool`

HasOAuth2Params returns a boolean if a field has been set.

### GetOpenIdConnectParams

`func (o *IdentityConfig) GetOpenIdConnectParams() OpenIdProvider`

GetOpenIdConnectParams returns the OpenIdConnectParams field if non-nil, zero value otherwise.

### GetOpenIdConnectParamsOk

`func (o *IdentityConfig) GetOpenIdConnectParamsOk() (*OpenIdProvider, bool)`

GetOpenIdConnectParamsOk returns a tuple with the OpenIdConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIdConnectParams

`func (o *IdentityConfig) SetOpenIdConnectParams(v OpenIdProvider)`

SetOpenIdConnectParams sets OpenIdConnectParams field to given value.

### HasOpenIdConnectParams

`func (o *IdentityConfig) HasOpenIdConnectParams() bool`

HasOpenIdConnectParams returns a boolean if a field has been set.

### GetTenantId

`func (o *IdentityConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IdentityConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IdentityConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IdentityConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *IdentityConfig) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *IdentityConfig) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


