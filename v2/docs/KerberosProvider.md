# KerberosProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminServer** | **NullableString** | Specifies the admin server used for registration from the list of KDC servers. | 
**HostAlias** | **[]string** | Specifies the DNS routable host alias names. | 
**Id** | Pointer to **NullableString** | Specifies the id. | [optional] [readonly] 
**KdcServers** | **[]string** | Specifies a list of Key distribution Centre(KDC) Severs. | 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id to be mapped | [optional] 
**OverwritehostAlias** | Pointer to **NullableBool** | Specifies if specified host alias should overwrite existing host alias. | [optional] 
**RealmName** | **NullableString** | Specifies the realm name. | 

## Methods

### NewKerberosProvider

`func NewKerberosProvider(adminServer NullableString, hostAlias []string, kdcServers []string, realmName NullableString, ) *KerberosProvider`

NewKerberosProvider instantiates a new KerberosProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosProviderWithDefaults

`func NewKerberosProviderWithDefaults() *KerberosProvider`

NewKerberosProviderWithDefaults instantiates a new KerberosProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminServer

`func (o *KerberosProvider) GetAdminServer() string`

GetAdminServer returns the AdminServer field if non-nil, zero value otherwise.

### GetAdminServerOk

`func (o *KerberosProvider) GetAdminServerOk() (*string, bool)`

GetAdminServerOk returns a tuple with the AdminServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminServer

`func (o *KerberosProvider) SetAdminServer(v string)`

SetAdminServer sets AdminServer field to given value.


### SetAdminServerNil

`func (o *KerberosProvider) SetAdminServerNil(b bool)`

 SetAdminServerNil sets the value for AdminServer to be an explicit nil

### UnsetAdminServer
`func (o *KerberosProvider) UnsetAdminServer()`

UnsetAdminServer ensures that no value is present for AdminServer, not even an explicit nil
### GetHostAlias

`func (o *KerberosProvider) GetHostAlias() []string`

GetHostAlias returns the HostAlias field if non-nil, zero value otherwise.

### GetHostAliasOk

`func (o *KerberosProvider) GetHostAliasOk() (*[]string, bool)`

GetHostAliasOk returns a tuple with the HostAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAlias

`func (o *KerberosProvider) SetHostAlias(v []string)`

SetHostAlias sets HostAlias field to given value.


### GetId

`func (o *KerberosProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KerberosProvider) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KerberosProvider) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKdcServers

`func (o *KerberosProvider) GetKdcServers() []string`

GetKdcServers returns the KdcServers field if non-nil, zero value otherwise.

### GetKdcServersOk

`func (o *KerberosProvider) GetKdcServersOk() (*[]string, bool)`

GetKdcServersOk returns a tuple with the KdcServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcServers

`func (o *KerberosProvider) SetKdcServers(v []string)`

SetKdcServers sets KdcServers field to given value.


### GetLdapProviderId

`func (o *KerberosProvider) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *KerberosProvider) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *KerberosProvider) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *KerberosProvider) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *KerberosProvider) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *KerberosProvider) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetOverwritehostAlias

`func (o *KerberosProvider) GetOverwritehostAlias() bool`

GetOverwritehostAlias returns the OverwritehostAlias field if non-nil, zero value otherwise.

### GetOverwritehostAliasOk

`func (o *KerberosProvider) GetOverwritehostAliasOk() (*bool, bool)`

GetOverwritehostAliasOk returns a tuple with the OverwritehostAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwritehostAlias

`func (o *KerberosProvider) SetOverwritehostAlias(v bool)`

SetOverwritehostAlias sets OverwritehostAlias field to given value.

### HasOverwritehostAlias

`func (o *KerberosProvider) HasOverwritehostAlias() bool`

HasOverwritehostAlias returns a boolean if a field has been set.

### SetOverwritehostAliasNil

`func (o *KerberosProvider) SetOverwritehostAliasNil(b bool)`

 SetOverwritehostAliasNil sets the value for OverwritehostAlias to be an explicit nil

### UnsetOverwritehostAlias
`func (o *KerberosProvider) UnsetOverwritehostAlias()`

UnsetOverwritehostAlias ensures that no value is present for OverwritehostAlias, not even an explicit nil
### GetRealmName

`func (o *KerberosProvider) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *KerberosProvider) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *KerberosProvider) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.


### SetRealmNameNil

`func (o *KerberosProvider) SetRealmNameNil(b bool)`

 SetRealmNameNil sets the value for RealmName to be an explicit nil

### UnsetRealmName
`func (o *KerberosProvider) UnsetRealmName()`

UnsetRealmName ensures that no value is present for RealmName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


