# RegisterOrUpdateKerberosProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id. | [optional] [readonly] 
**RealmName** | **NullableString** | Specifies the realm name. | 
**KdcServers** | **[]string** | Specifies a list of Key distribution Centre(KDC) Severs. | 
**AdminServer** | **NullableString** | Specifies the admin server used for registration from the list of KDC servers. | 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id to be mapped | [optional] 
**OverwritehostAlias** | Pointer to **NullableBool** | Specifies if specified host alias should overwrite existing host alias. | [optional] 
**HostAlias** | **[]string** | Specifies the DNS routable host alias names. | 
**AdminPassword** | **NullableString** | Specifies the password | 

## Methods

### NewRegisterOrUpdateKerberosProviderRequest

`func NewRegisterOrUpdateKerberosProviderRequest(realmName NullableString, kdcServers []string, adminServer NullableString, hostAlias []string, adminPassword NullableString, ) *RegisterOrUpdateKerberosProviderRequest`

NewRegisterOrUpdateKerberosProviderRequest instantiates a new RegisterOrUpdateKerberosProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterOrUpdateKerberosProviderRequestWithDefaults

`func NewRegisterOrUpdateKerberosProviderRequestWithDefaults() *RegisterOrUpdateKerberosProviderRequest`

NewRegisterOrUpdateKerberosProviderRequestWithDefaults instantiates a new RegisterOrUpdateKerberosProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterOrUpdateKerberosProviderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterOrUpdateKerberosProviderRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterOrUpdateKerberosProviderRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRealmName

`func (o *RegisterOrUpdateKerberosProviderRequest) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *RegisterOrUpdateKerberosProviderRequest) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.


### SetRealmNameNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetRealmNameNil(b bool)`

 SetRealmNameNil sets the value for RealmName to be an explicit nil

### UnsetRealmName
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetRealmName()`

UnsetRealmName ensures that no value is present for RealmName, not even an explicit nil
### GetKdcServers

`func (o *RegisterOrUpdateKerberosProviderRequest) GetKdcServers() []string`

GetKdcServers returns the KdcServers field if non-nil, zero value otherwise.

### GetKdcServersOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetKdcServersOk() (*[]string, bool)`

GetKdcServersOk returns a tuple with the KdcServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcServers

`func (o *RegisterOrUpdateKerberosProviderRequest) SetKdcServers(v []string)`

SetKdcServers sets KdcServers field to given value.


### GetAdminServer

`func (o *RegisterOrUpdateKerberosProviderRequest) GetAdminServer() string`

GetAdminServer returns the AdminServer field if non-nil, zero value otherwise.

### GetAdminServerOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetAdminServerOk() (*string, bool)`

GetAdminServerOk returns a tuple with the AdminServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminServer

`func (o *RegisterOrUpdateKerberosProviderRequest) SetAdminServer(v string)`

SetAdminServer sets AdminServer field to given value.


### SetAdminServerNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetAdminServerNil(b bool)`

 SetAdminServerNil sets the value for AdminServer to be an explicit nil

### UnsetAdminServer
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetAdminServer()`

UnsetAdminServer ensures that no value is present for AdminServer, not even an explicit nil
### GetLdapProviderId

`func (o *RegisterOrUpdateKerberosProviderRequest) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *RegisterOrUpdateKerberosProviderRequest) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *RegisterOrUpdateKerberosProviderRequest) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetOverwritehostAlias

`func (o *RegisterOrUpdateKerberosProviderRequest) GetOverwritehostAlias() bool`

GetOverwritehostAlias returns the OverwritehostAlias field if non-nil, zero value otherwise.

### GetOverwritehostAliasOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetOverwritehostAliasOk() (*bool, bool)`

GetOverwritehostAliasOk returns a tuple with the OverwritehostAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwritehostAlias

`func (o *RegisterOrUpdateKerberosProviderRequest) SetOverwritehostAlias(v bool)`

SetOverwritehostAlias sets OverwritehostAlias field to given value.

### HasOverwritehostAlias

`func (o *RegisterOrUpdateKerberosProviderRequest) HasOverwritehostAlias() bool`

HasOverwritehostAlias returns a boolean if a field has been set.

### SetOverwritehostAliasNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetOverwritehostAliasNil(b bool)`

 SetOverwritehostAliasNil sets the value for OverwritehostAlias to be an explicit nil

### UnsetOverwritehostAlias
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetOverwritehostAlias()`

UnsetOverwritehostAlias ensures that no value is present for OverwritehostAlias, not even an explicit nil
### GetHostAlias

`func (o *RegisterOrUpdateKerberosProviderRequest) GetHostAlias() []string`

GetHostAlias returns the HostAlias field if non-nil, zero value otherwise.

### GetHostAliasOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetHostAliasOk() (*[]string, bool)`

GetHostAliasOk returns a tuple with the HostAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAlias

`func (o *RegisterOrUpdateKerberosProviderRequest) SetHostAlias(v []string)`

SetHostAlias sets HostAlias field to given value.


### GetAdminPassword

`func (o *RegisterOrUpdateKerberosProviderRequest) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *RegisterOrUpdateKerberosProviderRequest) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *RegisterOrUpdateKerberosProviderRequest) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.


### SetAdminPasswordNil

`func (o *RegisterOrUpdateKerberosProviderRequest) SetAdminPasswordNil(b bool)`

 SetAdminPasswordNil sets the value for AdminPassword to be an explicit nil

### UnsetAdminPassword
`func (o *RegisterOrUpdateKerberosProviderRequest) UnsetAdminPassword()`

UnsetAdminPassword ensures that no value is present for AdminPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


