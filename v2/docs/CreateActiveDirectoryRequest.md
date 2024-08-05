# CreateActiveDirectoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection. | [optional] 
**DomainControllersDenyList** | Pointer to **[]string** | Specifies a list of denied domain controllers of this Active Directory Domain. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the Active Directory. | [optional] [readonly] 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id which is mapped to this Active Directory | [optional] 
**MachineAccounts** | [**[]MachineAccount**](MachineAccount.md) | Specifies a list of computer names used to identify the Cohesity Cluster on the Active Directory domain. The first machine account is used as primary machine account and it can not be modified. | 
**NisProviderDomainName** | Pointer to **NullableString** | Specifies the name of the NIS Provider which is mapped to this Active Directory. | [optional] 
**OrganizationalUnitName** | Pointer to **NullableString** | Specifies an optional organizational unit name. | [optional] 
**PreferredDomainControllers** | Pointer to [**[]DomainController**](DomainController.md) | Specifies a list of preferred domain controllers of this Active Directory. | [optional] 
**TrustedDomainParams** | Pointer to [**NullableCommonActiveDirectoryParamsTrustedDomainParams**](CommonActiveDirectoryParamsTrustedDomainParams.md) |  | [optional] 
**WorkGroupName** | Pointer to **NullableString** | Specifies a work group name. | [optional] 
**ActiveDirectoryAdminParams** | **map[string]interface{}** | Specifies the params of a user with administrative privilege of this Active Directory. | 
**DomainName** | **NullableString** | Specifies the domain name of the Active Directory. | 
**OverwriteMachineAccounts** | Pointer to **NullableBool** | Specifies if specified machine accounts should overwrite existing machine accounts. | [optional] 

## Methods

### NewCreateActiveDirectoryRequest

`func NewCreateActiveDirectoryRequest(machineAccounts []MachineAccount, activeDirectoryAdminParams map[string]interface{}, domainName NullableString, ) *CreateActiveDirectoryRequest`

NewCreateActiveDirectoryRequest instantiates a new CreateActiveDirectoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActiveDirectoryRequestWithDefaults

`func NewCreateActiveDirectoryRequestWithDefaults() *CreateActiveDirectoryRequest`

NewCreateActiveDirectoryRequestWithDefaults instantiates a new CreateActiveDirectoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *CreateActiveDirectoryRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateActiveDirectoryRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateActiveDirectoryRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CreateActiveDirectoryRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *CreateActiveDirectoryRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CreateActiveDirectoryRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetDomainControllersDenyList

`func (o *CreateActiveDirectoryRequest) GetDomainControllersDenyList() []*string`

GetDomainControllersDenyList returns the DomainControllersDenyList field if non-nil, zero value otherwise.

### GetDomainControllersDenyListOk

`func (o *CreateActiveDirectoryRequest) GetDomainControllersDenyListOk() (*[]*string, bool)`

GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllersDenyList

`func (o *CreateActiveDirectoryRequest) SetDomainControllersDenyList(v []*string)`

SetDomainControllersDenyList sets DomainControllersDenyList field to given value.

### HasDomainControllersDenyList

`func (o *CreateActiveDirectoryRequest) HasDomainControllersDenyList() bool`

HasDomainControllersDenyList returns a boolean if a field has been set.

### GetId

`func (o *CreateActiveDirectoryRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateActiveDirectoryRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateActiveDirectoryRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateActiveDirectoryRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateActiveDirectoryRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateActiveDirectoryRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLdapProviderId

`func (o *CreateActiveDirectoryRequest) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *CreateActiveDirectoryRequest) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *CreateActiveDirectoryRequest) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *CreateActiveDirectoryRequest) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *CreateActiveDirectoryRequest) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *CreateActiveDirectoryRequest) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *CreateActiveDirectoryRequest) GetMachineAccounts() []MachineAccount`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *CreateActiveDirectoryRequest) GetMachineAccountsOk() (*[]MachineAccount, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *CreateActiveDirectoryRequest) SetMachineAccounts(v []MachineAccount)`

SetMachineAccounts sets MachineAccounts field to given value.


### SetMachineAccountsNil

`func (o *CreateActiveDirectoryRequest) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *CreateActiveDirectoryRequest) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetNisProviderDomainName

`func (o *CreateActiveDirectoryRequest) GetNisProviderDomainName() string`

GetNisProviderDomainName returns the NisProviderDomainName field if non-nil, zero value otherwise.

### GetNisProviderDomainNameOk

`func (o *CreateActiveDirectoryRequest) GetNisProviderDomainNameOk() (*string, bool)`

GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderDomainName

`func (o *CreateActiveDirectoryRequest) SetNisProviderDomainName(v string)`

SetNisProviderDomainName sets NisProviderDomainName field to given value.

### HasNisProviderDomainName

`func (o *CreateActiveDirectoryRequest) HasNisProviderDomainName() bool`

HasNisProviderDomainName returns a boolean if a field has been set.

### SetNisProviderDomainNameNil

`func (o *CreateActiveDirectoryRequest) SetNisProviderDomainNameNil(b bool)`

 SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil

### UnsetNisProviderDomainName
`func (o *CreateActiveDirectoryRequest) UnsetNisProviderDomainName()`

UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
### GetOrganizationalUnitName

`func (o *CreateActiveDirectoryRequest) GetOrganizationalUnitName() string`

GetOrganizationalUnitName returns the OrganizationalUnitName field if non-nil, zero value otherwise.

### GetOrganizationalUnitNameOk

`func (o *CreateActiveDirectoryRequest) GetOrganizationalUnitNameOk() (*string, bool)`

GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnitName

`func (o *CreateActiveDirectoryRequest) SetOrganizationalUnitName(v string)`

SetOrganizationalUnitName sets OrganizationalUnitName field to given value.

### HasOrganizationalUnitName

`func (o *CreateActiveDirectoryRequest) HasOrganizationalUnitName() bool`

HasOrganizationalUnitName returns a boolean if a field has been set.

### SetOrganizationalUnitNameNil

`func (o *CreateActiveDirectoryRequest) SetOrganizationalUnitNameNil(b bool)`

 SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil

### UnsetOrganizationalUnitName
`func (o *CreateActiveDirectoryRequest) UnsetOrganizationalUnitName()`

UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
### GetPreferredDomainControllers

`func (o *CreateActiveDirectoryRequest) GetPreferredDomainControllers() []DomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *CreateActiveDirectoryRequest) GetPreferredDomainControllersOk() (*[]DomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *CreateActiveDirectoryRequest) SetPreferredDomainControllers(v []DomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *CreateActiveDirectoryRequest) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *CreateActiveDirectoryRequest) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *CreateActiveDirectoryRequest) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTrustedDomainParams

`func (o *CreateActiveDirectoryRequest) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams`

GetTrustedDomainParams returns the TrustedDomainParams field if non-nil, zero value otherwise.

### GetTrustedDomainParamsOk

`func (o *CreateActiveDirectoryRequest) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool)`

GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainParams

`func (o *CreateActiveDirectoryRequest) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams)`

SetTrustedDomainParams sets TrustedDomainParams field to given value.

### HasTrustedDomainParams

`func (o *CreateActiveDirectoryRequest) HasTrustedDomainParams() bool`

HasTrustedDomainParams returns a boolean if a field has been set.

### SetTrustedDomainParamsNil

`func (o *CreateActiveDirectoryRequest) SetTrustedDomainParamsNil(b bool)`

 SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil

### UnsetTrustedDomainParams
`func (o *CreateActiveDirectoryRequest) UnsetTrustedDomainParams()`

UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
### GetWorkGroupName

`func (o *CreateActiveDirectoryRequest) GetWorkGroupName() string`

GetWorkGroupName returns the WorkGroupName field if non-nil, zero value otherwise.

### GetWorkGroupNameOk

`func (o *CreateActiveDirectoryRequest) GetWorkGroupNameOk() (*string, bool)`

GetWorkGroupNameOk returns a tuple with the WorkGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupName

`func (o *CreateActiveDirectoryRequest) SetWorkGroupName(v string)`

SetWorkGroupName sets WorkGroupName field to given value.

### HasWorkGroupName

`func (o *CreateActiveDirectoryRequest) HasWorkGroupName() bool`

HasWorkGroupName returns a boolean if a field has been set.

### SetWorkGroupNameNil

`func (o *CreateActiveDirectoryRequest) SetWorkGroupNameNil(b bool)`

 SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil

### UnsetWorkGroupName
`func (o *CreateActiveDirectoryRequest) UnsetWorkGroupName()`

UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil
### GetActiveDirectoryAdminParams

`func (o *CreateActiveDirectoryRequest) GetActiveDirectoryAdminParams() map[string]interface{}`

GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field if non-nil, zero value otherwise.

### GetActiveDirectoryAdminParamsOk

`func (o *CreateActiveDirectoryRequest) GetActiveDirectoryAdminParamsOk() (*map[string]interface{}, bool)`

GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryAdminParams

`func (o *CreateActiveDirectoryRequest) SetActiveDirectoryAdminParams(v map[string]interface{})`

SetActiveDirectoryAdminParams sets ActiveDirectoryAdminParams field to given value.


### SetActiveDirectoryAdminParamsNil

`func (o *CreateActiveDirectoryRequest) SetActiveDirectoryAdminParamsNil(b bool)`

 SetActiveDirectoryAdminParamsNil sets the value for ActiveDirectoryAdminParams to be an explicit nil

### UnsetActiveDirectoryAdminParams
`func (o *CreateActiveDirectoryRequest) UnsetActiveDirectoryAdminParams()`

UnsetActiveDirectoryAdminParams ensures that no value is present for ActiveDirectoryAdminParams, not even an explicit nil
### GetDomainName

`func (o *CreateActiveDirectoryRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateActiveDirectoryRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateActiveDirectoryRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *CreateActiveDirectoryRequest) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CreateActiveDirectoryRequest) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequest) GetOverwriteMachineAccounts() bool`

GetOverwriteMachineAccounts returns the OverwriteMachineAccounts field if non-nil, zero value otherwise.

### GetOverwriteMachineAccountsOk

`func (o *CreateActiveDirectoryRequest) GetOverwriteMachineAccountsOk() (*bool, bool)`

GetOverwriteMachineAccountsOk returns a tuple with the OverwriteMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequest) SetOverwriteMachineAccounts(v bool)`

SetOverwriteMachineAccounts sets OverwriteMachineAccounts field to given value.

### HasOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequest) HasOverwriteMachineAccounts() bool`

HasOverwriteMachineAccounts returns a boolean if a field has been set.

### SetOverwriteMachineAccountsNil

`func (o *CreateActiveDirectoryRequest) SetOverwriteMachineAccountsNil(b bool)`

 SetOverwriteMachineAccountsNil sets the value for OverwriteMachineAccounts to be an explicit nil

### UnsetOverwriteMachineAccounts
`func (o *CreateActiveDirectoryRequest) UnsetOverwriteMachineAccounts()`

UnsetOverwriteMachineAccounts ensures that no value is present for OverwriteMachineAccounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


