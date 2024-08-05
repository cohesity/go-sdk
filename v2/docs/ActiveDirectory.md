# ActiveDirectory

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
**CentrifyZones** | Pointer to [**[]CentrifyZones**](CentrifyZones.md) | Specifies a list of Centrify zones. | [optional] 
**DomainControllers** | Pointer to [**[]DomainControllers**](DomainControllers.md) | A list of domain names with a list of it&#39;s domain controllers. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the domain name of the Active Directory. | [optional] 
**Error** | Pointer to [**ActiveDirectoryError**](ActiveDirectoryError.md) |  | [optional] 
**IdMappingParams** | Pointer to **map[string]interface{}** | Specifies the params of the user id mapping info of an Active Directory. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this Active Directory. | [optional] 
**SecurityPrincipals** | Pointer to [**[]SecurityPrincipal**](SecurityPrincipal.md) | Specifies a list of security principals. | [optional] 
**TaskLogs** | Pointer to [**TaskLogs**](TaskLogs.md) |  | [optional] 
**TransitiveAdTrustLevelLimit** | Pointer to **NullableInt32** | Specifies level of transitive Active Directory trust domains to be used. | [optional] 

## Methods

### NewActiveDirectory

`func NewActiveDirectory(machineAccounts []MachineAccount, ) *ActiveDirectory`

NewActiveDirectory instantiates a new ActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryWithDefaults

`func NewActiveDirectoryWithDefaults() *ActiveDirectory`

NewActiveDirectoryWithDefaults instantiates a new ActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ActiveDirectory) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ActiveDirectory) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ActiveDirectory) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ActiveDirectory) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *ActiveDirectory) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *ActiveDirectory) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetDomainControllersDenyList

`func (o *ActiveDirectory) GetDomainControllersDenyList() []*string`

GetDomainControllersDenyList returns the DomainControllersDenyList field if non-nil, zero value otherwise.

### GetDomainControllersDenyListOk

`func (o *ActiveDirectory) GetDomainControllersDenyListOk() (*[]*string, bool)`

GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllersDenyList

`func (o *ActiveDirectory) SetDomainControllersDenyList(v []*string)`

SetDomainControllersDenyList sets DomainControllersDenyList field to given value.

### HasDomainControllersDenyList

`func (o *ActiveDirectory) HasDomainControllersDenyList() bool`

HasDomainControllersDenyList returns a boolean if a field has been set.

### GetId

`func (o *ActiveDirectory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveDirectory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveDirectory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActiveDirectory) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActiveDirectory) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActiveDirectory) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLdapProviderId

`func (o *ActiveDirectory) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *ActiveDirectory) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *ActiveDirectory) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *ActiveDirectory) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *ActiveDirectory) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *ActiveDirectory) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *ActiveDirectory) GetMachineAccounts() []MachineAccount`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *ActiveDirectory) GetMachineAccountsOk() (*[]MachineAccount, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *ActiveDirectory) SetMachineAccounts(v []MachineAccount)`

SetMachineAccounts sets MachineAccounts field to given value.


### SetMachineAccountsNil

`func (o *ActiveDirectory) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *ActiveDirectory) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetNisProviderDomainName

`func (o *ActiveDirectory) GetNisProviderDomainName() string`

GetNisProviderDomainName returns the NisProviderDomainName field if non-nil, zero value otherwise.

### GetNisProviderDomainNameOk

`func (o *ActiveDirectory) GetNisProviderDomainNameOk() (*string, bool)`

GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderDomainName

`func (o *ActiveDirectory) SetNisProviderDomainName(v string)`

SetNisProviderDomainName sets NisProviderDomainName field to given value.

### HasNisProviderDomainName

`func (o *ActiveDirectory) HasNisProviderDomainName() bool`

HasNisProviderDomainName returns a boolean if a field has been set.

### SetNisProviderDomainNameNil

`func (o *ActiveDirectory) SetNisProviderDomainNameNil(b bool)`

 SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil

### UnsetNisProviderDomainName
`func (o *ActiveDirectory) UnsetNisProviderDomainName()`

UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
### GetOrganizationalUnitName

`func (o *ActiveDirectory) GetOrganizationalUnitName() string`

GetOrganizationalUnitName returns the OrganizationalUnitName field if non-nil, zero value otherwise.

### GetOrganizationalUnitNameOk

`func (o *ActiveDirectory) GetOrganizationalUnitNameOk() (*string, bool)`

GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnitName

`func (o *ActiveDirectory) SetOrganizationalUnitName(v string)`

SetOrganizationalUnitName sets OrganizationalUnitName field to given value.

### HasOrganizationalUnitName

`func (o *ActiveDirectory) HasOrganizationalUnitName() bool`

HasOrganizationalUnitName returns a boolean if a field has been set.

### SetOrganizationalUnitNameNil

`func (o *ActiveDirectory) SetOrganizationalUnitNameNil(b bool)`

 SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil

### UnsetOrganizationalUnitName
`func (o *ActiveDirectory) UnsetOrganizationalUnitName()`

UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
### GetPreferredDomainControllers

`func (o *ActiveDirectory) GetPreferredDomainControllers() []DomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *ActiveDirectory) GetPreferredDomainControllersOk() (*[]DomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *ActiveDirectory) SetPreferredDomainControllers(v []DomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *ActiveDirectory) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *ActiveDirectory) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *ActiveDirectory) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTrustedDomainParams

`func (o *ActiveDirectory) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams`

GetTrustedDomainParams returns the TrustedDomainParams field if non-nil, zero value otherwise.

### GetTrustedDomainParamsOk

`func (o *ActiveDirectory) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool)`

GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainParams

`func (o *ActiveDirectory) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams)`

SetTrustedDomainParams sets TrustedDomainParams field to given value.

### HasTrustedDomainParams

`func (o *ActiveDirectory) HasTrustedDomainParams() bool`

HasTrustedDomainParams returns a boolean if a field has been set.

### SetTrustedDomainParamsNil

`func (o *ActiveDirectory) SetTrustedDomainParamsNil(b bool)`

 SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil

### UnsetTrustedDomainParams
`func (o *ActiveDirectory) UnsetTrustedDomainParams()`

UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
### GetWorkGroupName

`func (o *ActiveDirectory) GetWorkGroupName() string`

GetWorkGroupName returns the WorkGroupName field if non-nil, zero value otherwise.

### GetWorkGroupNameOk

`func (o *ActiveDirectory) GetWorkGroupNameOk() (*string, bool)`

GetWorkGroupNameOk returns a tuple with the WorkGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupName

`func (o *ActiveDirectory) SetWorkGroupName(v string)`

SetWorkGroupName sets WorkGroupName field to given value.

### HasWorkGroupName

`func (o *ActiveDirectory) HasWorkGroupName() bool`

HasWorkGroupName returns a boolean if a field has been set.

### SetWorkGroupNameNil

`func (o *ActiveDirectory) SetWorkGroupNameNil(b bool)`

 SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil

### UnsetWorkGroupName
`func (o *ActiveDirectory) UnsetWorkGroupName()`

UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil
### GetCentrifyZones

`func (o *ActiveDirectory) GetCentrifyZones() []CentrifyZones`

GetCentrifyZones returns the CentrifyZones field if non-nil, zero value otherwise.

### GetCentrifyZonesOk

`func (o *ActiveDirectory) GetCentrifyZonesOk() (*[]CentrifyZones, bool)`

GetCentrifyZonesOk returns a tuple with the CentrifyZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifyZones

`func (o *ActiveDirectory) SetCentrifyZones(v []CentrifyZones)`

SetCentrifyZones sets CentrifyZones field to given value.

### HasCentrifyZones

`func (o *ActiveDirectory) HasCentrifyZones() bool`

HasCentrifyZones returns a boolean if a field has been set.

### SetCentrifyZonesNil

`func (o *ActiveDirectory) SetCentrifyZonesNil(b bool)`

 SetCentrifyZonesNil sets the value for CentrifyZones to be an explicit nil

### UnsetCentrifyZones
`func (o *ActiveDirectory) UnsetCentrifyZones()`

UnsetCentrifyZones ensures that no value is present for CentrifyZones, not even an explicit nil
### GetDomainControllers

`func (o *ActiveDirectory) GetDomainControllers() []DomainControllers`

GetDomainControllers returns the DomainControllers field if non-nil, zero value otherwise.

### GetDomainControllersOk

`func (o *ActiveDirectory) GetDomainControllersOk() (*[]DomainControllers, bool)`

GetDomainControllersOk returns a tuple with the DomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllers

`func (o *ActiveDirectory) SetDomainControllers(v []DomainControllers)`

SetDomainControllers sets DomainControllers field to given value.

### HasDomainControllers

`func (o *ActiveDirectory) HasDomainControllers() bool`

HasDomainControllers returns a boolean if a field has been set.

### SetDomainControllersNil

`func (o *ActiveDirectory) SetDomainControllersNil(b bool)`

 SetDomainControllersNil sets the value for DomainControllers to be an explicit nil

### UnsetDomainControllers
`func (o *ActiveDirectory) UnsetDomainControllers()`

UnsetDomainControllers ensures that no value is present for DomainControllers, not even an explicit nil
### GetDomainName

`func (o *ActiveDirectory) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ActiveDirectory) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ActiveDirectory) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ActiveDirectory) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ActiveDirectory) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ActiveDirectory) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetError

`func (o *ActiveDirectory) GetError() ActiveDirectoryError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ActiveDirectory) GetErrorOk() (*ActiveDirectoryError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ActiveDirectory) SetError(v ActiveDirectoryError)`

SetError sets Error field to given value.

### HasError

`func (o *ActiveDirectory) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdMappingParams

`func (o *ActiveDirectory) GetIdMappingParams() map[string]interface{}`

GetIdMappingParams returns the IdMappingParams field if non-nil, zero value otherwise.

### GetIdMappingParamsOk

`func (o *ActiveDirectory) GetIdMappingParamsOk() (*map[string]interface{}, bool)`

GetIdMappingParamsOk returns a tuple with the IdMappingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingParams

`func (o *ActiveDirectory) SetIdMappingParams(v map[string]interface{})`

SetIdMappingParams sets IdMappingParams field to given value.

### HasIdMappingParams

`func (o *ActiveDirectory) HasIdMappingParams() bool`

HasIdMappingParams returns a boolean if a field has been set.

### SetIdMappingParamsNil

`func (o *ActiveDirectory) SetIdMappingParamsNil(b bool)`

 SetIdMappingParamsNil sets the value for IdMappingParams to be an explicit nil

### UnsetIdMappingParams
`func (o *ActiveDirectory) UnsetIdMappingParams()`

UnsetIdMappingParams ensures that no value is present for IdMappingParams, not even an explicit nil
### GetPermissions

`func (o *ActiveDirectory) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ActiveDirectory) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ActiveDirectory) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ActiveDirectory) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ActiveDirectory) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ActiveDirectory) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSecurityPrincipals

`func (o *ActiveDirectory) GetSecurityPrincipals() []SecurityPrincipal`

GetSecurityPrincipals returns the SecurityPrincipals field if non-nil, zero value otherwise.

### GetSecurityPrincipalsOk

`func (o *ActiveDirectory) GetSecurityPrincipalsOk() (*[]SecurityPrincipal, bool)`

GetSecurityPrincipalsOk returns a tuple with the SecurityPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrincipals

`func (o *ActiveDirectory) SetSecurityPrincipals(v []SecurityPrincipal)`

SetSecurityPrincipals sets SecurityPrincipals field to given value.

### HasSecurityPrincipals

`func (o *ActiveDirectory) HasSecurityPrincipals() bool`

HasSecurityPrincipals returns a boolean if a field has been set.

### SetSecurityPrincipalsNil

`func (o *ActiveDirectory) SetSecurityPrincipalsNil(b bool)`

 SetSecurityPrincipalsNil sets the value for SecurityPrincipals to be an explicit nil

### UnsetSecurityPrincipals
`func (o *ActiveDirectory) UnsetSecurityPrincipals()`

UnsetSecurityPrincipals ensures that no value is present for SecurityPrincipals, not even an explicit nil
### GetTaskLogs

`func (o *ActiveDirectory) GetTaskLogs() TaskLogs`

GetTaskLogs returns the TaskLogs field if non-nil, zero value otherwise.

### GetTaskLogsOk

`func (o *ActiveDirectory) GetTaskLogsOk() (*TaskLogs, bool)`

GetTaskLogsOk returns a tuple with the TaskLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskLogs

`func (o *ActiveDirectory) SetTaskLogs(v TaskLogs)`

SetTaskLogs sets TaskLogs field to given value.

### HasTaskLogs

`func (o *ActiveDirectory) HasTaskLogs() bool`

HasTaskLogs returns a boolean if a field has been set.

### GetTransitiveAdTrustLevelLimit

`func (o *ActiveDirectory) GetTransitiveAdTrustLevelLimit() int32`

GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field if non-nil, zero value otherwise.

### GetTransitiveAdTrustLevelLimitOk

`func (o *ActiveDirectory) GetTransitiveAdTrustLevelLimitOk() (*int32, bool)`

GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveAdTrustLevelLimit

`func (o *ActiveDirectory) SetTransitiveAdTrustLevelLimit(v int32)`

SetTransitiveAdTrustLevelLimit sets TransitiveAdTrustLevelLimit field to given value.

### HasTransitiveAdTrustLevelLimit

`func (o *ActiveDirectory) HasTransitiveAdTrustLevelLimit() bool`

HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.

### SetTransitiveAdTrustLevelLimitNil

`func (o *ActiveDirectory) SetTransitiveAdTrustLevelLimitNil(b bool)`

 SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil

### UnsetTransitiveAdTrustLevelLimit
`func (o *ActiveDirectory) UnsetTransitiveAdTrustLevelLimit()`

UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


