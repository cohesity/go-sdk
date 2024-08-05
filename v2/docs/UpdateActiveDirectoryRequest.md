# UpdateActiveDirectoryRequest

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
**ActiveDirectoryAdminParams** | Pointer to **map[string]interface{}** | Specifies the params of a user with administrative privilege of this Active Directory. This field is mandatory if machine accounts are updated. | [optional] 
**IdMappingParams** | Pointer to **map[string]interface{}** | Specifies the params of the user id mapping info of an Active Directory. | [optional] 
**OverwriteMachineAccounts** | Pointer to **NullableBool** | Specifies if specified machine accounts should overwrite existing machine accounts. | [optional] 
**TransitiveAdTrustLevelLimit** | Pointer to **NullableInt32** | Specifies level of transitive Active Directory trust domains to be used. | [optional] 

## Methods

### NewUpdateActiveDirectoryRequest

`func NewUpdateActiveDirectoryRequest(machineAccounts []MachineAccount, ) *UpdateActiveDirectoryRequest`

NewUpdateActiveDirectoryRequest instantiates a new UpdateActiveDirectoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActiveDirectoryRequestWithDefaults

`func NewUpdateActiveDirectoryRequestWithDefaults() *UpdateActiveDirectoryRequest`

NewUpdateActiveDirectoryRequestWithDefaults instantiates a new UpdateActiveDirectoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *UpdateActiveDirectoryRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UpdateActiveDirectoryRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UpdateActiveDirectoryRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *UpdateActiveDirectoryRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *UpdateActiveDirectoryRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *UpdateActiveDirectoryRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetDomainControllersDenyList

`func (o *UpdateActiveDirectoryRequest) GetDomainControllersDenyList() []*string`

GetDomainControllersDenyList returns the DomainControllersDenyList field if non-nil, zero value otherwise.

### GetDomainControllersDenyListOk

`func (o *UpdateActiveDirectoryRequest) GetDomainControllersDenyListOk() (*[]*string, bool)`

GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllersDenyList

`func (o *UpdateActiveDirectoryRequest) SetDomainControllersDenyList(v []*string)`

SetDomainControllersDenyList sets DomainControllersDenyList field to given value.

### HasDomainControllersDenyList

`func (o *UpdateActiveDirectoryRequest) HasDomainControllersDenyList() bool`

HasDomainControllersDenyList returns a boolean if a field has been set.

### GetId

`func (o *UpdateActiveDirectoryRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateActiveDirectoryRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateActiveDirectoryRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateActiveDirectoryRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateActiveDirectoryRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateActiveDirectoryRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLdapProviderId

`func (o *UpdateActiveDirectoryRequest) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *UpdateActiveDirectoryRequest) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *UpdateActiveDirectoryRequest) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *UpdateActiveDirectoryRequest) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *UpdateActiveDirectoryRequest) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *UpdateActiveDirectoryRequest) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *UpdateActiveDirectoryRequest) GetMachineAccounts() []MachineAccount`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *UpdateActiveDirectoryRequest) GetMachineAccountsOk() (*[]MachineAccount, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *UpdateActiveDirectoryRequest) SetMachineAccounts(v []MachineAccount)`

SetMachineAccounts sets MachineAccounts field to given value.


### SetMachineAccountsNil

`func (o *UpdateActiveDirectoryRequest) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *UpdateActiveDirectoryRequest) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetNisProviderDomainName

`func (o *UpdateActiveDirectoryRequest) GetNisProviderDomainName() string`

GetNisProviderDomainName returns the NisProviderDomainName field if non-nil, zero value otherwise.

### GetNisProviderDomainNameOk

`func (o *UpdateActiveDirectoryRequest) GetNisProviderDomainNameOk() (*string, bool)`

GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderDomainName

`func (o *UpdateActiveDirectoryRequest) SetNisProviderDomainName(v string)`

SetNisProviderDomainName sets NisProviderDomainName field to given value.

### HasNisProviderDomainName

`func (o *UpdateActiveDirectoryRequest) HasNisProviderDomainName() bool`

HasNisProviderDomainName returns a boolean if a field has been set.

### SetNisProviderDomainNameNil

`func (o *UpdateActiveDirectoryRequest) SetNisProviderDomainNameNil(b bool)`

 SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil

### UnsetNisProviderDomainName
`func (o *UpdateActiveDirectoryRequest) UnsetNisProviderDomainName()`

UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
### GetOrganizationalUnitName

`func (o *UpdateActiveDirectoryRequest) GetOrganizationalUnitName() string`

GetOrganizationalUnitName returns the OrganizationalUnitName field if non-nil, zero value otherwise.

### GetOrganizationalUnitNameOk

`func (o *UpdateActiveDirectoryRequest) GetOrganizationalUnitNameOk() (*string, bool)`

GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnitName

`func (o *UpdateActiveDirectoryRequest) SetOrganizationalUnitName(v string)`

SetOrganizationalUnitName sets OrganizationalUnitName field to given value.

### HasOrganizationalUnitName

`func (o *UpdateActiveDirectoryRequest) HasOrganizationalUnitName() bool`

HasOrganizationalUnitName returns a boolean if a field has been set.

### SetOrganizationalUnitNameNil

`func (o *UpdateActiveDirectoryRequest) SetOrganizationalUnitNameNil(b bool)`

 SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil

### UnsetOrganizationalUnitName
`func (o *UpdateActiveDirectoryRequest) UnsetOrganizationalUnitName()`

UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
### GetPreferredDomainControllers

`func (o *UpdateActiveDirectoryRequest) GetPreferredDomainControllers() []DomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *UpdateActiveDirectoryRequest) GetPreferredDomainControllersOk() (*[]DomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *UpdateActiveDirectoryRequest) SetPreferredDomainControllers(v []DomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *UpdateActiveDirectoryRequest) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *UpdateActiveDirectoryRequest) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *UpdateActiveDirectoryRequest) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTrustedDomainParams

`func (o *UpdateActiveDirectoryRequest) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams`

GetTrustedDomainParams returns the TrustedDomainParams field if non-nil, zero value otherwise.

### GetTrustedDomainParamsOk

`func (o *UpdateActiveDirectoryRequest) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool)`

GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainParams

`func (o *UpdateActiveDirectoryRequest) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams)`

SetTrustedDomainParams sets TrustedDomainParams field to given value.

### HasTrustedDomainParams

`func (o *UpdateActiveDirectoryRequest) HasTrustedDomainParams() bool`

HasTrustedDomainParams returns a boolean if a field has been set.

### SetTrustedDomainParamsNil

`func (o *UpdateActiveDirectoryRequest) SetTrustedDomainParamsNil(b bool)`

 SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil

### UnsetTrustedDomainParams
`func (o *UpdateActiveDirectoryRequest) UnsetTrustedDomainParams()`

UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
### GetWorkGroupName

`func (o *UpdateActiveDirectoryRequest) GetWorkGroupName() string`

GetWorkGroupName returns the WorkGroupName field if non-nil, zero value otherwise.

### GetWorkGroupNameOk

`func (o *UpdateActiveDirectoryRequest) GetWorkGroupNameOk() (*string, bool)`

GetWorkGroupNameOk returns a tuple with the WorkGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupName

`func (o *UpdateActiveDirectoryRequest) SetWorkGroupName(v string)`

SetWorkGroupName sets WorkGroupName field to given value.

### HasWorkGroupName

`func (o *UpdateActiveDirectoryRequest) HasWorkGroupName() bool`

HasWorkGroupName returns a boolean if a field has been set.

### SetWorkGroupNameNil

`func (o *UpdateActiveDirectoryRequest) SetWorkGroupNameNil(b bool)`

 SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil

### UnsetWorkGroupName
`func (o *UpdateActiveDirectoryRequest) UnsetWorkGroupName()`

UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil
### GetActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequest) GetActiveDirectoryAdminParams() map[string]interface{}`

GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field if non-nil, zero value otherwise.

### GetActiveDirectoryAdminParamsOk

`func (o *UpdateActiveDirectoryRequest) GetActiveDirectoryAdminParamsOk() (*map[string]interface{}, bool)`

GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequest) SetActiveDirectoryAdminParams(v map[string]interface{})`

SetActiveDirectoryAdminParams sets ActiveDirectoryAdminParams field to given value.

### HasActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequest) HasActiveDirectoryAdminParams() bool`

HasActiveDirectoryAdminParams returns a boolean if a field has been set.

### SetActiveDirectoryAdminParamsNil

`func (o *UpdateActiveDirectoryRequest) SetActiveDirectoryAdminParamsNil(b bool)`

 SetActiveDirectoryAdminParamsNil sets the value for ActiveDirectoryAdminParams to be an explicit nil

### UnsetActiveDirectoryAdminParams
`func (o *UpdateActiveDirectoryRequest) UnsetActiveDirectoryAdminParams()`

UnsetActiveDirectoryAdminParams ensures that no value is present for ActiveDirectoryAdminParams, not even an explicit nil
### GetIdMappingParams

`func (o *UpdateActiveDirectoryRequest) GetIdMappingParams() map[string]interface{}`

GetIdMappingParams returns the IdMappingParams field if non-nil, zero value otherwise.

### GetIdMappingParamsOk

`func (o *UpdateActiveDirectoryRequest) GetIdMappingParamsOk() (*map[string]interface{}, bool)`

GetIdMappingParamsOk returns a tuple with the IdMappingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingParams

`func (o *UpdateActiveDirectoryRequest) SetIdMappingParams(v map[string]interface{})`

SetIdMappingParams sets IdMappingParams field to given value.

### HasIdMappingParams

`func (o *UpdateActiveDirectoryRequest) HasIdMappingParams() bool`

HasIdMappingParams returns a boolean if a field has been set.

### SetIdMappingParamsNil

`func (o *UpdateActiveDirectoryRequest) SetIdMappingParamsNil(b bool)`

 SetIdMappingParamsNil sets the value for IdMappingParams to be an explicit nil

### UnsetIdMappingParams
`func (o *UpdateActiveDirectoryRequest) UnsetIdMappingParams()`

UnsetIdMappingParams ensures that no value is present for IdMappingParams, not even an explicit nil
### GetOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequest) GetOverwriteMachineAccounts() bool`

GetOverwriteMachineAccounts returns the OverwriteMachineAccounts field if non-nil, zero value otherwise.

### GetOverwriteMachineAccountsOk

`func (o *UpdateActiveDirectoryRequest) GetOverwriteMachineAccountsOk() (*bool, bool)`

GetOverwriteMachineAccountsOk returns a tuple with the OverwriteMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequest) SetOverwriteMachineAccounts(v bool)`

SetOverwriteMachineAccounts sets OverwriteMachineAccounts field to given value.

### HasOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequest) HasOverwriteMachineAccounts() bool`

HasOverwriteMachineAccounts returns a boolean if a field has been set.

### SetOverwriteMachineAccountsNil

`func (o *UpdateActiveDirectoryRequest) SetOverwriteMachineAccountsNil(b bool)`

 SetOverwriteMachineAccountsNil sets the value for OverwriteMachineAccounts to be an explicit nil

### UnsetOverwriteMachineAccounts
`func (o *UpdateActiveDirectoryRequest) UnsetOverwriteMachineAccounts()`

UnsetOverwriteMachineAccounts ensures that no value is present for OverwriteMachineAccounts, not even an explicit nil
### GetTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequest) GetTransitiveAdTrustLevelLimit() int32`

GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field if non-nil, zero value otherwise.

### GetTransitiveAdTrustLevelLimitOk

`func (o *UpdateActiveDirectoryRequest) GetTransitiveAdTrustLevelLimitOk() (*int32, bool)`

GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequest) SetTransitiveAdTrustLevelLimit(v int32)`

SetTransitiveAdTrustLevelLimit sets TransitiveAdTrustLevelLimit field to given value.

### HasTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequest) HasTransitiveAdTrustLevelLimit() bool`

HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.

### SetTransitiveAdTrustLevelLimitNil

`func (o *UpdateActiveDirectoryRequest) SetTransitiveAdTrustLevelLimitNil(b bool)`

 SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil

### UnsetTransitiveAdTrustLevelLimit
`func (o *UpdateActiveDirectoryRequest) UnsetTransitiveAdTrustLevelLimit()`

UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


