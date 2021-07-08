# UpdateLdapProviderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id which is mapped to an Active Directory | [optional] 

## Methods

### NewUpdateLdapProviderParams

`func NewUpdateLdapProviderParams() *UpdateLdapProviderParams`

NewUpdateLdapProviderParams instantiates a new UpdateLdapProviderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLdapProviderParamsWithDefaults

`func NewUpdateLdapProviderParamsWithDefaults() *UpdateLdapProviderParams`

NewUpdateLdapProviderParamsWithDefaults instantiates a new UpdateLdapProviderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapProviderId

`func (o *UpdateLdapProviderParams) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *UpdateLdapProviderParams) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *UpdateLdapProviderParams) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *UpdateLdapProviderParams) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *UpdateLdapProviderParams) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *UpdateLdapProviderParams) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


