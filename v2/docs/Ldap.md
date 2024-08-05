# Ldap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryId** | Pointer to **int64** | Specifies the Active Directory id which is mapped to this LDAP. | [optional] 
**AdDomainName** | Pointer to **NullableString** | Specifies the domain name of an Active Directory which is mapped to this LDAP provider | [optional] 
**AttributeCommonName** | Pointer to **NullableString** | Specifies name of the LDAP attribute used for common name of an object. | [optional] 
**AttributeGid** | Pointer to **NullableString** | Specifies name of the attribute used to lookup unix GID of an LDAP user. | [optional] 
**AttributeMemberOf** | Pointer to **NullableString** | Specifies name of the LDAP attribute used to lookup members of a group. | [optional] 
**AttributeUid** | Pointer to **NullableString** | Specifies name of the attribute used to lookup unix UID of an LDAP user. | [optional] 
**AttributeUsername** | Pointer to **NullableString** | Specifies name of the LDAP attribute used to lookup a user by user ID. | [optional] 
**AuthType** | **string** | Specifies the LDAP authentication type. | 
**BaseDistinguishedName** | **NullableString** | Specifies the base distinguished name used as the base for LDAP search requests. | 
**DomainName** | Pointer to **NullableString** | Specifies the name of the domain name to be used for querying LDAP servers from DNS. | [optional] 
**Name** | **NullableString** | Specifies the LDAP name. | 
**ObjectClassGroup** | Pointer to **NullableString** | Specifies name of the LDAP group object class for user accounts. | [optional] 
**ObjectClassUser** | Pointer to **NullableString** | Specifies name of the LDAP user object class for user accounts. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the LDAP server port. | [optional] 
**PreferredLdapServers** | Pointer to **[]string** | Specifies a list of preferred LDAP servers. Servers should either be FQDNs or IP addresses. | [optional] 
**SimpleAuthParams** | Pointer to [**CreateLdapParamsSimpleAuthParams**](CreateLdapParamsSimpleAuthParams.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the LDAP id. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the LDAP tenant id. | [optional] 

## Methods

### NewLdap

`func NewLdap(authType string, baseDistinguishedName NullableString, name NullableString, ) *Ldap`

NewLdap instantiates a new Ldap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapWithDefaults

`func NewLdapWithDefaults() *Ldap`

NewLdapWithDefaults instantiates a new Ldap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryId

`func (o *Ldap) GetActiveDirectoryId() int64`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *Ldap) GetActiveDirectoryIdOk() (*int64, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *Ldap) SetActiveDirectoryId(v int64)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *Ldap) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetAdDomainName

`func (o *Ldap) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *Ldap) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *Ldap) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *Ldap) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *Ldap) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *Ldap) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetAttributeCommonName

`func (o *Ldap) GetAttributeCommonName() string`

GetAttributeCommonName returns the AttributeCommonName field if non-nil, zero value otherwise.

### GetAttributeCommonNameOk

`func (o *Ldap) GetAttributeCommonNameOk() (*string, bool)`

GetAttributeCommonNameOk returns a tuple with the AttributeCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCommonName

`func (o *Ldap) SetAttributeCommonName(v string)`

SetAttributeCommonName sets AttributeCommonName field to given value.

### HasAttributeCommonName

`func (o *Ldap) HasAttributeCommonName() bool`

HasAttributeCommonName returns a boolean if a field has been set.

### SetAttributeCommonNameNil

`func (o *Ldap) SetAttributeCommonNameNil(b bool)`

 SetAttributeCommonNameNil sets the value for AttributeCommonName to be an explicit nil

### UnsetAttributeCommonName
`func (o *Ldap) UnsetAttributeCommonName()`

UnsetAttributeCommonName ensures that no value is present for AttributeCommonName, not even an explicit nil
### GetAttributeGid

`func (o *Ldap) GetAttributeGid() string`

GetAttributeGid returns the AttributeGid field if non-nil, zero value otherwise.

### GetAttributeGidOk

`func (o *Ldap) GetAttributeGidOk() (*string, bool)`

GetAttributeGidOk returns a tuple with the AttributeGid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeGid

`func (o *Ldap) SetAttributeGid(v string)`

SetAttributeGid sets AttributeGid field to given value.

### HasAttributeGid

`func (o *Ldap) HasAttributeGid() bool`

HasAttributeGid returns a boolean if a field has been set.

### SetAttributeGidNil

`func (o *Ldap) SetAttributeGidNil(b bool)`

 SetAttributeGidNil sets the value for AttributeGid to be an explicit nil

### UnsetAttributeGid
`func (o *Ldap) UnsetAttributeGid()`

UnsetAttributeGid ensures that no value is present for AttributeGid, not even an explicit nil
### GetAttributeMemberOf

`func (o *Ldap) GetAttributeMemberOf() string`

GetAttributeMemberOf returns the AttributeMemberOf field if non-nil, zero value otherwise.

### GetAttributeMemberOfOk

`func (o *Ldap) GetAttributeMemberOfOk() (*string, bool)`

GetAttributeMemberOfOk returns a tuple with the AttributeMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMemberOf

`func (o *Ldap) SetAttributeMemberOf(v string)`

SetAttributeMemberOf sets AttributeMemberOf field to given value.

### HasAttributeMemberOf

`func (o *Ldap) HasAttributeMemberOf() bool`

HasAttributeMemberOf returns a boolean if a field has been set.

### SetAttributeMemberOfNil

`func (o *Ldap) SetAttributeMemberOfNil(b bool)`

 SetAttributeMemberOfNil sets the value for AttributeMemberOf to be an explicit nil

### UnsetAttributeMemberOf
`func (o *Ldap) UnsetAttributeMemberOf()`

UnsetAttributeMemberOf ensures that no value is present for AttributeMemberOf, not even an explicit nil
### GetAttributeUid

`func (o *Ldap) GetAttributeUid() string`

GetAttributeUid returns the AttributeUid field if non-nil, zero value otherwise.

### GetAttributeUidOk

`func (o *Ldap) GetAttributeUidOk() (*string, bool)`

GetAttributeUidOk returns a tuple with the AttributeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeUid

`func (o *Ldap) SetAttributeUid(v string)`

SetAttributeUid sets AttributeUid field to given value.

### HasAttributeUid

`func (o *Ldap) HasAttributeUid() bool`

HasAttributeUid returns a boolean if a field has been set.

### SetAttributeUidNil

`func (o *Ldap) SetAttributeUidNil(b bool)`

 SetAttributeUidNil sets the value for AttributeUid to be an explicit nil

### UnsetAttributeUid
`func (o *Ldap) UnsetAttributeUid()`

UnsetAttributeUid ensures that no value is present for AttributeUid, not even an explicit nil
### GetAttributeUsername

`func (o *Ldap) GetAttributeUsername() string`

GetAttributeUsername returns the AttributeUsername field if non-nil, zero value otherwise.

### GetAttributeUsernameOk

`func (o *Ldap) GetAttributeUsernameOk() (*string, bool)`

GetAttributeUsernameOk returns a tuple with the AttributeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeUsername

`func (o *Ldap) SetAttributeUsername(v string)`

SetAttributeUsername sets AttributeUsername field to given value.

### HasAttributeUsername

`func (o *Ldap) HasAttributeUsername() bool`

HasAttributeUsername returns a boolean if a field has been set.

### SetAttributeUsernameNil

`func (o *Ldap) SetAttributeUsernameNil(b bool)`

 SetAttributeUsernameNil sets the value for AttributeUsername to be an explicit nil

### UnsetAttributeUsername
`func (o *Ldap) UnsetAttributeUsername()`

UnsetAttributeUsername ensures that no value is present for AttributeUsername, not even an explicit nil
### GetAuthType

`func (o *Ldap) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Ldap) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Ldap) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetBaseDistinguishedName

`func (o *Ldap) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *Ldap) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *Ldap) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.


### SetBaseDistinguishedNameNil

`func (o *Ldap) SetBaseDistinguishedNameNil(b bool)`

 SetBaseDistinguishedNameNil sets the value for BaseDistinguishedName to be an explicit nil

### UnsetBaseDistinguishedName
`func (o *Ldap) UnsetBaseDistinguishedName()`

UnsetBaseDistinguishedName ensures that no value is present for BaseDistinguishedName, not even an explicit nil
### GetDomainName

`func (o *Ldap) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Ldap) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Ldap) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Ldap) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *Ldap) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *Ldap) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetName

`func (o *Ldap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ldap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ldap) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Ldap) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Ldap) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectClassGroup

`func (o *Ldap) GetObjectClassGroup() string`

GetObjectClassGroup returns the ObjectClassGroup field if non-nil, zero value otherwise.

### GetObjectClassGroupOk

`func (o *Ldap) GetObjectClassGroupOk() (*string, bool)`

GetObjectClassGroupOk returns a tuple with the ObjectClassGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassGroup

`func (o *Ldap) SetObjectClassGroup(v string)`

SetObjectClassGroup sets ObjectClassGroup field to given value.

### HasObjectClassGroup

`func (o *Ldap) HasObjectClassGroup() bool`

HasObjectClassGroup returns a boolean if a field has been set.

### SetObjectClassGroupNil

`func (o *Ldap) SetObjectClassGroupNil(b bool)`

 SetObjectClassGroupNil sets the value for ObjectClassGroup to be an explicit nil

### UnsetObjectClassGroup
`func (o *Ldap) UnsetObjectClassGroup()`

UnsetObjectClassGroup ensures that no value is present for ObjectClassGroup, not even an explicit nil
### GetObjectClassUser

`func (o *Ldap) GetObjectClassUser() string`

GetObjectClassUser returns the ObjectClassUser field if non-nil, zero value otherwise.

### GetObjectClassUserOk

`func (o *Ldap) GetObjectClassUserOk() (*string, bool)`

GetObjectClassUserOk returns a tuple with the ObjectClassUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassUser

`func (o *Ldap) SetObjectClassUser(v string)`

SetObjectClassUser sets ObjectClassUser field to given value.

### HasObjectClassUser

`func (o *Ldap) HasObjectClassUser() bool`

HasObjectClassUser returns a boolean if a field has been set.

### SetObjectClassUserNil

`func (o *Ldap) SetObjectClassUserNil(b bool)`

 SetObjectClassUserNil sets the value for ObjectClassUser to be an explicit nil

### UnsetObjectClassUser
`func (o *Ldap) UnsetObjectClassUser()`

UnsetObjectClassUser ensures that no value is present for ObjectClassUser, not even an explicit nil
### GetPort

`func (o *Ldap) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Ldap) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Ldap) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Ldap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *Ldap) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Ldap) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPreferredLdapServers

`func (o *Ldap) GetPreferredLdapServers() []string`

GetPreferredLdapServers returns the PreferredLdapServers field if non-nil, zero value otherwise.

### GetPreferredLdapServersOk

`func (o *Ldap) GetPreferredLdapServersOk() (*[]string, bool)`

GetPreferredLdapServersOk returns a tuple with the PreferredLdapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLdapServers

`func (o *Ldap) SetPreferredLdapServers(v []string)`

SetPreferredLdapServers sets PreferredLdapServers field to given value.

### HasPreferredLdapServers

`func (o *Ldap) HasPreferredLdapServers() bool`

HasPreferredLdapServers returns a boolean if a field has been set.

### SetPreferredLdapServersNil

`func (o *Ldap) SetPreferredLdapServersNil(b bool)`

 SetPreferredLdapServersNil sets the value for PreferredLdapServers to be an explicit nil

### UnsetPreferredLdapServers
`func (o *Ldap) UnsetPreferredLdapServers()`

UnsetPreferredLdapServers ensures that no value is present for PreferredLdapServers, not even an explicit nil
### GetSimpleAuthParams

`func (o *Ldap) GetSimpleAuthParams() CreateLdapParamsSimpleAuthParams`

GetSimpleAuthParams returns the SimpleAuthParams field if non-nil, zero value otherwise.

### GetSimpleAuthParamsOk

`func (o *Ldap) GetSimpleAuthParamsOk() (*CreateLdapParamsSimpleAuthParams, bool)`

GetSimpleAuthParamsOk returns a tuple with the SimpleAuthParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAuthParams

`func (o *Ldap) SetSimpleAuthParams(v CreateLdapParamsSimpleAuthParams)`

SetSimpleAuthParams sets SimpleAuthParams field to given value.

### HasSimpleAuthParams

`func (o *Ldap) HasSimpleAuthParams() bool`

HasSimpleAuthParams returns a boolean if a field has been set.

### GetId

`func (o *Ldap) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ldap) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ldap) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Ldap) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Ldap) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Ldap) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *Ldap) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Ldap) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Ldap) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Ldap) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Ldap) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Ldap) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


