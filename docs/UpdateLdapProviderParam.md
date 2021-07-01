# UpdateLdapProviderParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainName** | Pointer to **NullableString** | Specifies the domain name of an Active Directory which is mapped to this LDAP provider | [optional] 
**AuthType** | Pointer to **NullableString** | Specifies the authentication type used while connecting to LDAP servers. Authentication level. &#39;kAnonymous&#39; indicates LDAP authentication type &#39;Anonymous&#39; &#39;kSimple&#39; indicates LDAP authentication type &#39;Simple&#39; | [optional] 
**BaseDistinguishedName** | Pointer to **NullableString** | Specifies the base distinguished name used as the base for LDAP search requests. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the name of the domain name to be used for querying LDAP servers from DNS. If PreferredLdapServerList is set, then DomainName field is ignored. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the LDAP provider. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the LDAP provider. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies LDAP server port. | [optional] 
**PreferredLdapServerList** | Pointer to **[]string** | Specifies the preferred LDAP servers. Server names should be either in fully qualified domain name (FQDN) format or IP addresses. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**UseSsl** | Pointer to **NullableBool** | Specifies whether to use SSL for LDAP connections. | [optional] 
**UserDistinguishedName** | Pointer to **NullableString** | Specifies the user distinguished name that is used for LDAP authentication. It should be provided if the AuthType is set to either kSimple or kSasl. | [optional] 
**UserPassword** | Pointer to **NullableString** | Specifies the user password that is used for LDAP authentication. | [optional] 

## Methods

### NewUpdateLdapProviderParam

`func NewUpdateLdapProviderParam() *UpdateLdapProviderParam`

NewUpdateLdapProviderParam instantiates a new UpdateLdapProviderParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLdapProviderParamWithDefaults

`func NewUpdateLdapProviderParamWithDefaults() *UpdateLdapProviderParam`

NewUpdateLdapProviderParamWithDefaults instantiates a new UpdateLdapProviderParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *UpdateLdapProviderParam) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *UpdateLdapProviderParam) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *UpdateLdapProviderParam) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *UpdateLdapProviderParam) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *UpdateLdapProviderParam) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *UpdateLdapProviderParam) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetAuthType

`func (o *UpdateLdapProviderParam) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateLdapProviderParam) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateLdapProviderParam) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateLdapProviderParam) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *UpdateLdapProviderParam) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *UpdateLdapProviderParam) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetBaseDistinguishedName

`func (o *UpdateLdapProviderParam) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *UpdateLdapProviderParam) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *UpdateLdapProviderParam) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.

### HasBaseDistinguishedName

`func (o *UpdateLdapProviderParam) HasBaseDistinguishedName() bool`

HasBaseDistinguishedName returns a boolean if a field has been set.

### SetBaseDistinguishedNameNil

`func (o *UpdateLdapProviderParam) SetBaseDistinguishedNameNil(b bool)`

 SetBaseDistinguishedNameNil sets the value for BaseDistinguishedName to be an explicit nil

### UnsetBaseDistinguishedName
`func (o *UpdateLdapProviderParam) UnsetBaseDistinguishedName()`

UnsetBaseDistinguishedName ensures that no value is present for BaseDistinguishedName, not even an explicit nil
### GetDomainName

`func (o *UpdateLdapProviderParam) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateLdapProviderParam) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateLdapProviderParam) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UpdateLdapProviderParam) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *UpdateLdapProviderParam) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *UpdateLdapProviderParam) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetId

`func (o *UpdateLdapProviderParam) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLdapProviderParam) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLdapProviderParam) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLdapProviderParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateLdapProviderParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateLdapProviderParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UpdateLdapProviderParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLdapProviderParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLdapProviderParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLdapProviderParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateLdapProviderParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateLdapProviderParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPort

`func (o *UpdateLdapProviderParam) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateLdapProviderParam) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateLdapProviderParam) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateLdapProviderParam) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *UpdateLdapProviderParam) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *UpdateLdapProviderParam) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPreferredLdapServerList

`func (o *UpdateLdapProviderParam) GetPreferredLdapServerList() []string`

GetPreferredLdapServerList returns the PreferredLdapServerList field if non-nil, zero value otherwise.

### GetPreferredLdapServerListOk

`func (o *UpdateLdapProviderParam) GetPreferredLdapServerListOk() (*[]string, bool)`

GetPreferredLdapServerListOk returns a tuple with the PreferredLdapServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLdapServerList

`func (o *UpdateLdapProviderParam) SetPreferredLdapServerList(v []string)`

SetPreferredLdapServerList sets PreferredLdapServerList field to given value.

### HasPreferredLdapServerList

`func (o *UpdateLdapProviderParam) HasPreferredLdapServerList() bool`

HasPreferredLdapServerList returns a boolean if a field has been set.

### SetPreferredLdapServerListNil

`func (o *UpdateLdapProviderParam) SetPreferredLdapServerListNil(b bool)`

 SetPreferredLdapServerListNil sets the value for PreferredLdapServerList to be an explicit nil

### UnsetPreferredLdapServerList
`func (o *UpdateLdapProviderParam) UnsetPreferredLdapServerList()`

UnsetPreferredLdapServerList ensures that no value is present for PreferredLdapServerList, not even an explicit nil
### GetTenantId

`func (o *UpdateLdapProviderParam) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateLdapProviderParam) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateLdapProviderParam) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateLdapProviderParam) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateLdapProviderParam) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateLdapProviderParam) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseSsl

`func (o *UpdateLdapProviderParam) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *UpdateLdapProviderParam) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *UpdateLdapProviderParam) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *UpdateLdapProviderParam) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### SetUseSslNil

`func (o *UpdateLdapProviderParam) SetUseSslNil(b bool)`

 SetUseSslNil sets the value for UseSsl to be an explicit nil

### UnsetUseSsl
`func (o *UpdateLdapProviderParam) UnsetUseSsl()`

UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
### GetUserDistinguishedName

`func (o *UpdateLdapProviderParam) GetUserDistinguishedName() string`

GetUserDistinguishedName returns the UserDistinguishedName field if non-nil, zero value otherwise.

### GetUserDistinguishedNameOk

`func (o *UpdateLdapProviderParam) GetUserDistinguishedNameOk() (*string, bool)`

GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDistinguishedName

`func (o *UpdateLdapProviderParam) SetUserDistinguishedName(v string)`

SetUserDistinguishedName sets UserDistinguishedName field to given value.

### HasUserDistinguishedName

`func (o *UpdateLdapProviderParam) HasUserDistinguishedName() bool`

HasUserDistinguishedName returns a boolean if a field has been set.

### SetUserDistinguishedNameNil

`func (o *UpdateLdapProviderParam) SetUserDistinguishedNameNil(b bool)`

 SetUserDistinguishedNameNil sets the value for UserDistinguishedName to be an explicit nil

### UnsetUserDistinguishedName
`func (o *UpdateLdapProviderParam) UnsetUserDistinguishedName()`

UnsetUserDistinguishedName ensures that no value is present for UserDistinguishedName, not even an explicit nil
### GetUserPassword

`func (o *UpdateLdapProviderParam) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *UpdateLdapProviderParam) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *UpdateLdapProviderParam) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *UpdateLdapProviderParam) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### SetUserPasswordNil

`func (o *UpdateLdapProviderParam) SetUserPasswordNil(b bool)`

 SetUserPasswordNil sets the value for UserPassword to be an explicit nil

### UnsetUserPassword
`func (o *UpdateLdapProviderParam) UnsetUserPassword()`

UnsetUserPassword ensures that no value is present for UserPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


