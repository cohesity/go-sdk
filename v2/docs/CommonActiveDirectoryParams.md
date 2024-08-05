# CommonActiveDirectoryParams

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

## Methods

### NewCommonActiveDirectoryParams

`func NewCommonActiveDirectoryParams(machineAccounts []MachineAccount, ) *CommonActiveDirectoryParams`

NewCommonActiveDirectoryParams instantiates a new CommonActiveDirectoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonActiveDirectoryParamsWithDefaults

`func NewCommonActiveDirectoryParamsWithDefaults() *CommonActiveDirectoryParams`

NewCommonActiveDirectoryParamsWithDefaults instantiates a new CommonActiveDirectoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *CommonActiveDirectoryParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CommonActiveDirectoryParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CommonActiveDirectoryParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CommonActiveDirectoryParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *CommonActiveDirectoryParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CommonActiveDirectoryParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetDomainControllersDenyList

`func (o *CommonActiveDirectoryParams) GetDomainControllersDenyList() []*string`

GetDomainControllersDenyList returns the DomainControllersDenyList field if non-nil, zero value otherwise.

### GetDomainControllersDenyListOk

`func (o *CommonActiveDirectoryParams) GetDomainControllersDenyListOk() (*[]*string, bool)`

GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllersDenyList

`func (o *CommonActiveDirectoryParams) SetDomainControllersDenyList(v []*string)`

SetDomainControllersDenyList sets DomainControllersDenyList field to given value.

### HasDomainControllersDenyList

`func (o *CommonActiveDirectoryParams) HasDomainControllersDenyList() bool`

HasDomainControllersDenyList returns a boolean if a field has been set.

### GetId

`func (o *CommonActiveDirectoryParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonActiveDirectoryParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonActiveDirectoryParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonActiveDirectoryParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonActiveDirectoryParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonActiveDirectoryParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLdapProviderId

`func (o *CommonActiveDirectoryParams) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *CommonActiveDirectoryParams) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *CommonActiveDirectoryParams) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *CommonActiveDirectoryParams) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *CommonActiveDirectoryParams) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *CommonActiveDirectoryParams) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *CommonActiveDirectoryParams) GetMachineAccounts() []MachineAccount`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *CommonActiveDirectoryParams) GetMachineAccountsOk() (*[]MachineAccount, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *CommonActiveDirectoryParams) SetMachineAccounts(v []MachineAccount)`

SetMachineAccounts sets MachineAccounts field to given value.


### SetMachineAccountsNil

`func (o *CommonActiveDirectoryParams) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *CommonActiveDirectoryParams) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetNisProviderDomainName

`func (o *CommonActiveDirectoryParams) GetNisProviderDomainName() string`

GetNisProviderDomainName returns the NisProviderDomainName field if non-nil, zero value otherwise.

### GetNisProviderDomainNameOk

`func (o *CommonActiveDirectoryParams) GetNisProviderDomainNameOk() (*string, bool)`

GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderDomainName

`func (o *CommonActiveDirectoryParams) SetNisProviderDomainName(v string)`

SetNisProviderDomainName sets NisProviderDomainName field to given value.

### HasNisProviderDomainName

`func (o *CommonActiveDirectoryParams) HasNisProviderDomainName() bool`

HasNisProviderDomainName returns a boolean if a field has been set.

### SetNisProviderDomainNameNil

`func (o *CommonActiveDirectoryParams) SetNisProviderDomainNameNil(b bool)`

 SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil

### UnsetNisProviderDomainName
`func (o *CommonActiveDirectoryParams) UnsetNisProviderDomainName()`

UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
### GetOrganizationalUnitName

`func (o *CommonActiveDirectoryParams) GetOrganizationalUnitName() string`

GetOrganizationalUnitName returns the OrganizationalUnitName field if non-nil, zero value otherwise.

### GetOrganizationalUnitNameOk

`func (o *CommonActiveDirectoryParams) GetOrganizationalUnitNameOk() (*string, bool)`

GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnitName

`func (o *CommonActiveDirectoryParams) SetOrganizationalUnitName(v string)`

SetOrganizationalUnitName sets OrganizationalUnitName field to given value.

### HasOrganizationalUnitName

`func (o *CommonActiveDirectoryParams) HasOrganizationalUnitName() bool`

HasOrganizationalUnitName returns a boolean if a field has been set.

### SetOrganizationalUnitNameNil

`func (o *CommonActiveDirectoryParams) SetOrganizationalUnitNameNil(b bool)`

 SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil

### UnsetOrganizationalUnitName
`func (o *CommonActiveDirectoryParams) UnsetOrganizationalUnitName()`

UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
### GetPreferredDomainControllers

`func (o *CommonActiveDirectoryParams) GetPreferredDomainControllers() []DomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *CommonActiveDirectoryParams) GetPreferredDomainControllersOk() (*[]DomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *CommonActiveDirectoryParams) SetPreferredDomainControllers(v []DomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *CommonActiveDirectoryParams) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *CommonActiveDirectoryParams) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *CommonActiveDirectoryParams) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTrustedDomainParams

`func (o *CommonActiveDirectoryParams) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams`

GetTrustedDomainParams returns the TrustedDomainParams field if non-nil, zero value otherwise.

### GetTrustedDomainParamsOk

`func (o *CommonActiveDirectoryParams) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool)`

GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainParams

`func (o *CommonActiveDirectoryParams) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams)`

SetTrustedDomainParams sets TrustedDomainParams field to given value.

### HasTrustedDomainParams

`func (o *CommonActiveDirectoryParams) HasTrustedDomainParams() bool`

HasTrustedDomainParams returns a boolean if a field has been set.

### SetTrustedDomainParamsNil

`func (o *CommonActiveDirectoryParams) SetTrustedDomainParamsNil(b bool)`

 SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil

### UnsetTrustedDomainParams
`func (o *CommonActiveDirectoryParams) UnsetTrustedDomainParams()`

UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
### GetWorkGroupName

`func (o *CommonActiveDirectoryParams) GetWorkGroupName() string`

GetWorkGroupName returns the WorkGroupName field if non-nil, zero value otherwise.

### GetWorkGroupNameOk

`func (o *CommonActiveDirectoryParams) GetWorkGroupNameOk() (*string, bool)`

GetWorkGroupNameOk returns a tuple with the WorkGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupName

`func (o *CommonActiveDirectoryParams) SetWorkGroupName(v string)`

SetWorkGroupName sets WorkGroupName field to given value.

### HasWorkGroupName

`func (o *CommonActiveDirectoryParams) HasWorkGroupName() bool`

HasWorkGroupName returns a boolean if a field has been set.

### SetWorkGroupNameNil

`func (o *CommonActiveDirectoryParams) SetWorkGroupNameNil(b bool)`

 SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil

### UnsetWorkGroupName
`func (o *CommonActiveDirectoryParams) UnsetWorkGroupName()`

UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


