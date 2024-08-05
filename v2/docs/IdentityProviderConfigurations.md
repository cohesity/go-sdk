# IdentityProviderConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idps** | Pointer to [**[]IdentityProviderConfiguration**](IdentityProviderConfiguration.md) | Specifies a list of identity provider configurations | [optional] 

## Methods

### NewIdentityProviderConfigurations

`func NewIdentityProviderConfigurations() *IdentityProviderConfigurations`

NewIdentityProviderConfigurations instantiates a new IdentityProviderConfigurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderConfigurationsWithDefaults

`func NewIdentityProviderConfigurationsWithDefaults() *IdentityProviderConfigurations`

NewIdentityProviderConfigurationsWithDefaults instantiates a new IdentityProviderConfigurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdps

`func (o *IdentityProviderConfigurations) GetIdps() []IdentityProviderConfiguration`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *IdentityProviderConfigurations) GetIdpsOk() (*[]IdentityProviderConfiguration, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *IdentityProviderConfigurations) SetIdps(v []IdentityProviderConfiguration)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *IdentityProviderConfigurations) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### SetIdpsNil

`func (o *IdentityProviderConfigurations) SetIdpsNil(b bool)`

 SetIdpsNil sets the value for Idps to be an explicit nil

### UnsetIdps
`func (o *IdentityProviderConfigurations) UnsetIdps()`

UnsetIdps ensures that no value is present for Idps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


