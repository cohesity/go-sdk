# NisProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **NullableString** | Specifies the Domain Name of NIS Provider. | 
**MasterServerHostname** | **NullableString** | Specifies the hostname of Master Server. | 
**SlaveServers** | Pointer to **[]string** | Specifies a list of slave servers in the NIS Domain. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the list of tenant Ids for NIS Provider. | [optional] 

## Methods

### NewNisProvider

`func NewNisProvider(domain NullableString, masterServerHostname NullableString, ) *NisProvider`

NewNisProvider instantiates a new NisProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNisProviderWithDefaults

`func NewNisProviderWithDefaults() *NisProvider`

NewNisProviderWithDefaults instantiates a new NisProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *NisProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NisProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NisProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *NisProvider) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *NisProvider) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetMasterServerHostname

`func (o *NisProvider) GetMasterServerHostname() string`

GetMasterServerHostname returns the MasterServerHostname field if non-nil, zero value otherwise.

### GetMasterServerHostnameOk

`func (o *NisProvider) GetMasterServerHostnameOk() (*string, bool)`

GetMasterServerHostnameOk returns a tuple with the MasterServerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterServerHostname

`func (o *NisProvider) SetMasterServerHostname(v string)`

SetMasterServerHostname sets MasterServerHostname field to given value.


### SetMasterServerHostnameNil

`func (o *NisProvider) SetMasterServerHostnameNil(b bool)`

 SetMasterServerHostnameNil sets the value for MasterServerHostname to be an explicit nil

### UnsetMasterServerHostname
`func (o *NisProvider) UnsetMasterServerHostname()`

UnsetMasterServerHostname ensures that no value is present for MasterServerHostname, not even an explicit nil
### GetSlaveServers

`func (o *NisProvider) GetSlaveServers() []string`

GetSlaveServers returns the SlaveServers field if non-nil, zero value otherwise.

### GetSlaveServersOk

`func (o *NisProvider) GetSlaveServersOk() (*[]string, bool)`

GetSlaveServersOk returns a tuple with the SlaveServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveServers

`func (o *NisProvider) SetSlaveServers(v []string)`

SetSlaveServers sets SlaveServers field to given value.

### HasSlaveServers

`func (o *NisProvider) HasSlaveServers() bool`

HasSlaveServers returns a boolean if a field has been set.

### SetSlaveServersNil

`func (o *NisProvider) SetSlaveServersNil(b bool)`

 SetSlaveServersNil sets the value for SlaveServers to be an explicit nil

### UnsetSlaveServers
`func (o *NisProvider) UnsetSlaveServers()`

UnsetSlaveServers ensures that no value is present for SlaveServers, not even an explicit nil
### GetTenantIds

`func (o *NisProvider) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *NisProvider) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *NisProvider) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *NisProvider) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *NisProvider) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *NisProvider) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


