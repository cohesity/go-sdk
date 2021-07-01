# TenantLdapProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapProviderIds** | Pointer to **[]int64** | Specifies the ids of ldap providers for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantLdapProviderUpdate

`func NewTenantLdapProviderUpdate() *TenantLdapProviderUpdate`

NewTenantLdapProviderUpdate instantiates a new TenantLdapProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantLdapProviderUpdateWithDefaults

`func NewTenantLdapProviderUpdateWithDefaults() *TenantLdapProviderUpdate`

NewTenantLdapProviderUpdateWithDefaults instantiates a new TenantLdapProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapProviderIds

`func (o *TenantLdapProviderUpdate) GetLdapProviderIds() []int64`

GetLdapProviderIds returns the LdapProviderIds field if non-nil, zero value otherwise.

### GetLdapProviderIdsOk

`func (o *TenantLdapProviderUpdate) GetLdapProviderIdsOk() (*[]int64, bool)`

GetLdapProviderIdsOk returns a tuple with the LdapProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderIds

`func (o *TenantLdapProviderUpdate) SetLdapProviderIds(v []int64)`

SetLdapProviderIds sets LdapProviderIds field to given value.

### HasLdapProviderIds

`func (o *TenantLdapProviderUpdate) HasLdapProviderIds() bool`

HasLdapProviderIds returns a boolean if a field has been set.

### SetLdapProviderIdsNil

`func (o *TenantLdapProviderUpdate) SetLdapProviderIdsNil(b bool)`

 SetLdapProviderIdsNil sets the value for LdapProviderIds to be an explicit nil

### UnsetLdapProviderIds
`func (o *TenantLdapProviderUpdate) UnsetLdapProviderIds()`

UnsetLdapProviderIds ensures that no value is present for LdapProviderIds, not even an explicit nil
### GetTenantId

`func (o *TenantLdapProviderUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantLdapProviderUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantLdapProviderUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantLdapProviderUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantLdapProviderUpdate) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantLdapProviderUpdate) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


