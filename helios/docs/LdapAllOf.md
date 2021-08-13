# LdapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the LDAP id. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the LDAP tenant id. | [optional] 

## Methods

### NewLdapAllOf

`func NewLdapAllOf() *LdapAllOf`

NewLdapAllOf instantiates a new LdapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAllOfWithDefaults

`func NewLdapAllOfWithDefaults() *LdapAllOf`

NewLdapAllOfWithDefaults instantiates a new LdapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapAllOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LdapAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LdapAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LdapAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *LdapAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LdapAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LdapAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LdapAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LdapAllOf) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LdapAllOf) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


