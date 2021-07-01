# LdapProviderResponse

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

### NewLdapProviderResponse

`func NewLdapProviderResponse() *LdapProviderResponse`

NewLdapProviderResponse instantiates a new LdapProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderResponseWithDefaults

`func NewLdapProviderResponseWithDefaults() *LdapProviderResponse`

NewLdapProviderResponseWithDefaults instantiates a new LdapProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *LdapProviderResponse) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *LdapProviderResponse) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *LdapProviderResponse) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *LdapProviderResponse) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *LdapProviderResponse) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *LdapProviderResponse) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetAuthType

`func (o *LdapProviderResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *LdapProviderResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *LdapProviderResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *LdapProviderResponse) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *LdapProviderResponse) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *LdapProviderResponse) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetBaseDistinguishedName

`func (o *LdapProviderResponse) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *LdapProviderResponse) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *LdapProviderResponse) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.

### HasBaseDistinguishedName

`func (o *LdapProviderResponse) HasBaseDistinguishedName() bool`

HasBaseDistinguishedName returns a boolean if a field has been set.

### SetBaseDistinguishedNameNil

`func (o *LdapProviderResponse) SetBaseDistinguishedNameNil(b bool)`

 SetBaseDistinguishedNameNil sets the value for BaseDistinguishedName to be an explicit nil

### UnsetBaseDistinguishedName
`func (o *LdapProviderResponse) UnsetBaseDistinguishedName()`

UnsetBaseDistinguishedName ensures that no value is present for BaseDistinguishedName, not even an explicit nil
### GetDomainName

`func (o *LdapProviderResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *LdapProviderResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *LdapProviderResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *LdapProviderResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *LdapProviderResponse) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *LdapProviderResponse) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetId

`func (o *LdapProviderResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapProviderResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapProviderResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LdapProviderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LdapProviderResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LdapProviderResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *LdapProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapProviderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapProviderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LdapProviderResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LdapProviderResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPort

`func (o *LdapProviderResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LdapProviderResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LdapProviderResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LdapProviderResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LdapProviderResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LdapProviderResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPreferredLdapServerList

`func (o *LdapProviderResponse) GetPreferredLdapServerList() []string`

GetPreferredLdapServerList returns the PreferredLdapServerList field if non-nil, zero value otherwise.

### GetPreferredLdapServerListOk

`func (o *LdapProviderResponse) GetPreferredLdapServerListOk() (*[]string, bool)`

GetPreferredLdapServerListOk returns a tuple with the PreferredLdapServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLdapServerList

`func (o *LdapProviderResponse) SetPreferredLdapServerList(v []string)`

SetPreferredLdapServerList sets PreferredLdapServerList field to given value.

### HasPreferredLdapServerList

`func (o *LdapProviderResponse) HasPreferredLdapServerList() bool`

HasPreferredLdapServerList returns a boolean if a field has been set.

### SetPreferredLdapServerListNil

`func (o *LdapProviderResponse) SetPreferredLdapServerListNil(b bool)`

 SetPreferredLdapServerListNil sets the value for PreferredLdapServerList to be an explicit nil

### UnsetPreferredLdapServerList
`func (o *LdapProviderResponse) UnsetPreferredLdapServerList()`

UnsetPreferredLdapServerList ensures that no value is present for PreferredLdapServerList, not even an explicit nil
### GetTenantId

`func (o *LdapProviderResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LdapProviderResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LdapProviderResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LdapProviderResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LdapProviderResponse) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LdapProviderResponse) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseSsl

`func (o *LdapProviderResponse) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *LdapProviderResponse) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *LdapProviderResponse) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *LdapProviderResponse) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### SetUseSslNil

`func (o *LdapProviderResponse) SetUseSslNil(b bool)`

 SetUseSslNil sets the value for UseSsl to be an explicit nil

### UnsetUseSsl
`func (o *LdapProviderResponse) UnsetUseSsl()`

UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
### GetUserDistinguishedName

`func (o *LdapProviderResponse) GetUserDistinguishedName() string`

GetUserDistinguishedName returns the UserDistinguishedName field if non-nil, zero value otherwise.

### GetUserDistinguishedNameOk

`func (o *LdapProviderResponse) GetUserDistinguishedNameOk() (*string, bool)`

GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDistinguishedName

`func (o *LdapProviderResponse) SetUserDistinguishedName(v string)`

SetUserDistinguishedName sets UserDistinguishedName field to given value.

### HasUserDistinguishedName

`func (o *LdapProviderResponse) HasUserDistinguishedName() bool`

HasUserDistinguishedName returns a boolean if a field has been set.

### SetUserDistinguishedNameNil

`func (o *LdapProviderResponse) SetUserDistinguishedNameNil(b bool)`

 SetUserDistinguishedNameNil sets the value for UserDistinguishedName to be an explicit nil

### UnsetUserDistinguishedName
`func (o *LdapProviderResponse) UnsetUserDistinguishedName()`

UnsetUserDistinguishedName ensures that no value is present for UserDistinguishedName, not even an explicit nil
### GetUserPassword

`func (o *LdapProviderResponse) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *LdapProviderResponse) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *LdapProviderResponse) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *LdapProviderResponse) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### SetUserPasswordNil

`func (o *LdapProviderResponse) SetUserPasswordNil(b bool)`

 SetUserPasswordNil sets the value for UserPassword to be an explicit nil

### UnsetUserPassword
`func (o *LdapProviderResponse) UnsetUserPassword()`

UnsetUserPassword ensures that no value is present for UserPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


