# CreateLdapParams

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

## Methods

### NewCreateLdapParams

`func NewCreateLdapParams(authType string, baseDistinguishedName NullableString, name NullableString, ) *CreateLdapParams`

NewCreateLdapParams instantiates a new CreateLdapParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapParamsWithDefaults

`func NewCreateLdapParamsWithDefaults() *CreateLdapParams`

NewCreateLdapParamsWithDefaults instantiates a new CreateLdapParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryId

`func (o *CreateLdapParams) GetActiveDirectoryId() int64`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *CreateLdapParams) GetActiveDirectoryIdOk() (*int64, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *CreateLdapParams) SetActiveDirectoryId(v int64)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *CreateLdapParams) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetAdDomainName

`func (o *CreateLdapParams) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *CreateLdapParams) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *CreateLdapParams) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *CreateLdapParams) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *CreateLdapParams) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *CreateLdapParams) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetAttributeCommonName

`func (o *CreateLdapParams) GetAttributeCommonName() string`

GetAttributeCommonName returns the AttributeCommonName field if non-nil, zero value otherwise.

### GetAttributeCommonNameOk

`func (o *CreateLdapParams) GetAttributeCommonNameOk() (*string, bool)`

GetAttributeCommonNameOk returns a tuple with the AttributeCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCommonName

`func (o *CreateLdapParams) SetAttributeCommonName(v string)`

SetAttributeCommonName sets AttributeCommonName field to given value.

### HasAttributeCommonName

`func (o *CreateLdapParams) HasAttributeCommonName() bool`

HasAttributeCommonName returns a boolean if a field has been set.

### SetAttributeCommonNameNil

`func (o *CreateLdapParams) SetAttributeCommonNameNil(b bool)`

 SetAttributeCommonNameNil sets the value for AttributeCommonName to be an explicit nil

### UnsetAttributeCommonName
`func (o *CreateLdapParams) UnsetAttributeCommonName()`

UnsetAttributeCommonName ensures that no value is present for AttributeCommonName, not even an explicit nil
### GetAttributeGid

`func (o *CreateLdapParams) GetAttributeGid() string`

GetAttributeGid returns the AttributeGid field if non-nil, zero value otherwise.

### GetAttributeGidOk

`func (o *CreateLdapParams) GetAttributeGidOk() (*string, bool)`

GetAttributeGidOk returns a tuple with the AttributeGid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeGid

`func (o *CreateLdapParams) SetAttributeGid(v string)`

SetAttributeGid sets AttributeGid field to given value.

### HasAttributeGid

`func (o *CreateLdapParams) HasAttributeGid() bool`

HasAttributeGid returns a boolean if a field has been set.

### SetAttributeGidNil

`func (o *CreateLdapParams) SetAttributeGidNil(b bool)`

 SetAttributeGidNil sets the value for AttributeGid to be an explicit nil

### UnsetAttributeGid
`func (o *CreateLdapParams) UnsetAttributeGid()`

UnsetAttributeGid ensures that no value is present for AttributeGid, not even an explicit nil
### GetAttributeMemberOf

`func (o *CreateLdapParams) GetAttributeMemberOf() string`

GetAttributeMemberOf returns the AttributeMemberOf field if non-nil, zero value otherwise.

### GetAttributeMemberOfOk

`func (o *CreateLdapParams) GetAttributeMemberOfOk() (*string, bool)`

GetAttributeMemberOfOk returns a tuple with the AttributeMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMemberOf

`func (o *CreateLdapParams) SetAttributeMemberOf(v string)`

SetAttributeMemberOf sets AttributeMemberOf field to given value.

### HasAttributeMemberOf

`func (o *CreateLdapParams) HasAttributeMemberOf() bool`

HasAttributeMemberOf returns a boolean if a field has been set.

### SetAttributeMemberOfNil

`func (o *CreateLdapParams) SetAttributeMemberOfNil(b bool)`

 SetAttributeMemberOfNil sets the value for AttributeMemberOf to be an explicit nil

### UnsetAttributeMemberOf
`func (o *CreateLdapParams) UnsetAttributeMemberOf()`

UnsetAttributeMemberOf ensures that no value is present for AttributeMemberOf, not even an explicit nil
### GetAttributeUid

`func (o *CreateLdapParams) GetAttributeUid() string`

GetAttributeUid returns the AttributeUid field if non-nil, zero value otherwise.

### GetAttributeUidOk

`func (o *CreateLdapParams) GetAttributeUidOk() (*string, bool)`

GetAttributeUidOk returns a tuple with the AttributeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeUid

`func (o *CreateLdapParams) SetAttributeUid(v string)`

SetAttributeUid sets AttributeUid field to given value.

### HasAttributeUid

`func (o *CreateLdapParams) HasAttributeUid() bool`

HasAttributeUid returns a boolean if a field has been set.

### SetAttributeUidNil

`func (o *CreateLdapParams) SetAttributeUidNil(b bool)`

 SetAttributeUidNil sets the value for AttributeUid to be an explicit nil

### UnsetAttributeUid
`func (o *CreateLdapParams) UnsetAttributeUid()`

UnsetAttributeUid ensures that no value is present for AttributeUid, not even an explicit nil
### GetAttributeUsername

`func (o *CreateLdapParams) GetAttributeUsername() string`

GetAttributeUsername returns the AttributeUsername field if non-nil, zero value otherwise.

### GetAttributeUsernameOk

`func (o *CreateLdapParams) GetAttributeUsernameOk() (*string, bool)`

GetAttributeUsernameOk returns a tuple with the AttributeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeUsername

`func (o *CreateLdapParams) SetAttributeUsername(v string)`

SetAttributeUsername sets AttributeUsername field to given value.

### HasAttributeUsername

`func (o *CreateLdapParams) HasAttributeUsername() bool`

HasAttributeUsername returns a boolean if a field has been set.

### SetAttributeUsernameNil

`func (o *CreateLdapParams) SetAttributeUsernameNil(b bool)`

 SetAttributeUsernameNil sets the value for AttributeUsername to be an explicit nil

### UnsetAttributeUsername
`func (o *CreateLdapParams) UnsetAttributeUsername()`

UnsetAttributeUsername ensures that no value is present for AttributeUsername, not even an explicit nil
### GetAuthType

`func (o *CreateLdapParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateLdapParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateLdapParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetBaseDistinguishedName

`func (o *CreateLdapParams) GetBaseDistinguishedName() string`

GetBaseDistinguishedName returns the BaseDistinguishedName field if non-nil, zero value otherwise.

### GetBaseDistinguishedNameOk

`func (o *CreateLdapParams) GetBaseDistinguishedNameOk() (*string, bool)`

GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDistinguishedName

`func (o *CreateLdapParams) SetBaseDistinguishedName(v string)`

SetBaseDistinguishedName sets BaseDistinguishedName field to given value.


### SetBaseDistinguishedNameNil

`func (o *CreateLdapParams) SetBaseDistinguishedNameNil(b bool)`

 SetBaseDistinguishedNameNil sets the value for BaseDistinguishedName to be an explicit nil

### UnsetBaseDistinguishedName
`func (o *CreateLdapParams) UnsetBaseDistinguishedName()`

UnsetBaseDistinguishedName ensures that no value is present for BaseDistinguishedName, not even an explicit nil
### GetDomainName

`func (o *CreateLdapParams) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateLdapParams) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateLdapParams) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateLdapParams) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *CreateLdapParams) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CreateLdapParams) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetName

`func (o *CreateLdapParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLdapParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLdapParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateLdapParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateLdapParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectClassGroup

`func (o *CreateLdapParams) GetObjectClassGroup() string`

GetObjectClassGroup returns the ObjectClassGroup field if non-nil, zero value otherwise.

### GetObjectClassGroupOk

`func (o *CreateLdapParams) GetObjectClassGroupOk() (*string, bool)`

GetObjectClassGroupOk returns a tuple with the ObjectClassGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassGroup

`func (o *CreateLdapParams) SetObjectClassGroup(v string)`

SetObjectClassGroup sets ObjectClassGroup field to given value.

### HasObjectClassGroup

`func (o *CreateLdapParams) HasObjectClassGroup() bool`

HasObjectClassGroup returns a boolean if a field has been set.

### SetObjectClassGroupNil

`func (o *CreateLdapParams) SetObjectClassGroupNil(b bool)`

 SetObjectClassGroupNil sets the value for ObjectClassGroup to be an explicit nil

### UnsetObjectClassGroup
`func (o *CreateLdapParams) UnsetObjectClassGroup()`

UnsetObjectClassGroup ensures that no value is present for ObjectClassGroup, not even an explicit nil
### GetObjectClassUser

`func (o *CreateLdapParams) GetObjectClassUser() string`

GetObjectClassUser returns the ObjectClassUser field if non-nil, zero value otherwise.

### GetObjectClassUserOk

`func (o *CreateLdapParams) GetObjectClassUserOk() (*string, bool)`

GetObjectClassUserOk returns a tuple with the ObjectClassUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassUser

`func (o *CreateLdapParams) SetObjectClassUser(v string)`

SetObjectClassUser sets ObjectClassUser field to given value.

### HasObjectClassUser

`func (o *CreateLdapParams) HasObjectClassUser() bool`

HasObjectClassUser returns a boolean if a field has been set.

### SetObjectClassUserNil

`func (o *CreateLdapParams) SetObjectClassUserNil(b bool)`

 SetObjectClassUserNil sets the value for ObjectClassUser to be an explicit nil

### UnsetObjectClassUser
`func (o *CreateLdapParams) UnsetObjectClassUser()`

UnsetObjectClassUser ensures that no value is present for ObjectClassUser, not even an explicit nil
### GetPort

`func (o *CreateLdapParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateLdapParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateLdapParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateLdapParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CreateLdapParams) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CreateLdapParams) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPreferredLdapServers

`func (o *CreateLdapParams) GetPreferredLdapServers() []string`

GetPreferredLdapServers returns the PreferredLdapServers field if non-nil, zero value otherwise.

### GetPreferredLdapServersOk

`func (o *CreateLdapParams) GetPreferredLdapServersOk() (*[]string, bool)`

GetPreferredLdapServersOk returns a tuple with the PreferredLdapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLdapServers

`func (o *CreateLdapParams) SetPreferredLdapServers(v []string)`

SetPreferredLdapServers sets PreferredLdapServers field to given value.

### HasPreferredLdapServers

`func (o *CreateLdapParams) HasPreferredLdapServers() bool`

HasPreferredLdapServers returns a boolean if a field has been set.

### SetPreferredLdapServersNil

`func (o *CreateLdapParams) SetPreferredLdapServersNil(b bool)`

 SetPreferredLdapServersNil sets the value for PreferredLdapServers to be an explicit nil

### UnsetPreferredLdapServers
`func (o *CreateLdapParams) UnsetPreferredLdapServers()`

UnsetPreferredLdapServers ensures that no value is present for PreferredLdapServers, not even an explicit nil
### GetSimpleAuthParams

`func (o *CreateLdapParams) GetSimpleAuthParams() CreateLdapParamsSimpleAuthParams`

GetSimpleAuthParams returns the SimpleAuthParams field if non-nil, zero value otherwise.

### GetSimpleAuthParamsOk

`func (o *CreateLdapParams) GetSimpleAuthParamsOk() (*CreateLdapParamsSimpleAuthParams, bool)`

GetSimpleAuthParamsOk returns a tuple with the SimpleAuthParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAuthParams

`func (o *CreateLdapParams) SetSimpleAuthParams(v CreateLdapParamsSimpleAuthParams)`

SetSimpleAuthParams sets SimpleAuthParams field to given value.

### HasSimpleAuthParams

`func (o *CreateLdapParams) HasSimpleAuthParams() bool`

HasSimpleAuthParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


