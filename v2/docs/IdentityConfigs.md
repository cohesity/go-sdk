# IdentityConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idps** | Pointer to [**[]IdentityConfig**](IdentityConfig.md) | Specifies a list of Identity Providers. | [optional] 

## Methods

### NewIdentityConfigs

`func NewIdentityConfigs() *IdentityConfigs`

NewIdentityConfigs instantiates a new IdentityConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityConfigsWithDefaults

`func NewIdentityConfigsWithDefaults() *IdentityConfigs`

NewIdentityConfigsWithDefaults instantiates a new IdentityConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdps

`func (o *IdentityConfigs) GetIdps() []IdentityConfig`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *IdentityConfigs) GetIdpsOk() (*[]IdentityConfig, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *IdentityConfigs) SetIdps(v []IdentityConfig)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *IdentityConfigs) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### SetIdpsNil

`func (o *IdentityConfigs) SetIdpsNil(b bool)`

 SetIdpsNil sets the value for Idps to be an explicit nil

### UnsetIdps
`func (o *IdentityConfigs) UnsetIdps()`

UnsetIdps ensures that no value is present for Idps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


