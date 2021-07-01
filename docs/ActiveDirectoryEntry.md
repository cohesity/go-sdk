# ActiveDirectoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Specifies the fully qualified domain name (FQDN) of an Active Directory. | [optional] 
**FallbackUserIdMappingInfo** | Pointer to [**UserIdMapping**](UserIdMapping.md) |  | [optional] 
**IgnoredTrustedDomains** | Pointer to **[]string** | Specifies the list of trusted domains that were set by the user to be ignored during trusted domain discovery. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id which is map to this Active Directory | [optional] 
**MachineAccounts** | Pointer to **[]string** | Array of Machine Accounts.  Specifies an array of computer names used to identify the Cohesity Cluster on the domain. | [optional] 
**OuName** | Pointer to **NullableString** | Specifies an optional Organizational Unit name. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for the specified userName. | [optional] 
**PreferredDomainControllers** | Pointer to [**[]PreferredDomainController**](PreferredDomainController.md) | Specifies Map of Active Directory domain names to its preferred domain controllers. | [optional] 
**TaskPath** | Pointer to **NullableString** | Specifies the task path for AD joining task. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**TrustedDomains** | Pointer to **[]string** | Specifies the trusted domains of the Active Directory domain. | [optional] [readonly] 
**TrustedDomainsEnabled** | Pointer to **NullableBool** | Specifies whether Trusted Domain discovery is disabled. | [optional] 
**UnixRootSid** | Pointer to **NullableString** | Specifies the SID of the Active Directory domain user to be mapped to Unix root user. | [optional] 
**UserIdMappingInfo** | Pointer to [**UserIdMapping**](UserIdMapping.md) |  | [optional] 
**UserName** | Pointer to **NullableString** | Specifies a userName that has administrative privileges in the domain. | [optional] 
**Workgroup** | Pointer to **NullableString** | Specifies an optional Workgroup name. | [optional] 

## Methods

### NewActiveDirectoryEntry

`func NewActiveDirectoryEntry() *ActiveDirectoryEntry`

NewActiveDirectoryEntry instantiates a new ActiveDirectoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryEntryWithDefaults

`func NewActiveDirectoryEntryWithDefaults() *ActiveDirectoryEntry`

NewActiveDirectoryEntryWithDefaults instantiates a new ActiveDirectoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *ActiveDirectoryEntry) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ActiveDirectoryEntry) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ActiveDirectoryEntry) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ActiveDirectoryEntry) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ActiveDirectoryEntry) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ActiveDirectoryEntry) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetFallbackUserIdMappingInfo

`func (o *ActiveDirectoryEntry) GetFallbackUserIdMappingInfo() UserIdMapping`

GetFallbackUserIdMappingInfo returns the FallbackUserIdMappingInfo field if non-nil, zero value otherwise.

### GetFallbackUserIdMappingInfoOk

`func (o *ActiveDirectoryEntry) GetFallbackUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetFallbackUserIdMappingInfoOk returns a tuple with the FallbackUserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUserIdMappingInfo

`func (o *ActiveDirectoryEntry) SetFallbackUserIdMappingInfo(v UserIdMapping)`

SetFallbackUserIdMappingInfo sets FallbackUserIdMappingInfo field to given value.

### HasFallbackUserIdMappingInfo

`func (o *ActiveDirectoryEntry) HasFallbackUserIdMappingInfo() bool`

HasFallbackUserIdMappingInfo returns a boolean if a field has been set.

### GetIgnoredTrustedDomains

`func (o *ActiveDirectoryEntry) GetIgnoredTrustedDomains() []string`

GetIgnoredTrustedDomains returns the IgnoredTrustedDomains field if non-nil, zero value otherwise.

### GetIgnoredTrustedDomainsOk

`func (o *ActiveDirectoryEntry) GetIgnoredTrustedDomainsOk() (*[]string, bool)`

GetIgnoredTrustedDomainsOk returns a tuple with the IgnoredTrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredTrustedDomains

`func (o *ActiveDirectoryEntry) SetIgnoredTrustedDomains(v []string)`

SetIgnoredTrustedDomains sets IgnoredTrustedDomains field to given value.

### HasIgnoredTrustedDomains

`func (o *ActiveDirectoryEntry) HasIgnoredTrustedDomains() bool`

HasIgnoredTrustedDomains returns a boolean if a field has been set.

### SetIgnoredTrustedDomainsNil

`func (o *ActiveDirectoryEntry) SetIgnoredTrustedDomainsNil(b bool)`

 SetIgnoredTrustedDomainsNil sets the value for IgnoredTrustedDomains to be an explicit nil

### UnsetIgnoredTrustedDomains
`func (o *ActiveDirectoryEntry) UnsetIgnoredTrustedDomains()`

UnsetIgnoredTrustedDomains ensures that no value is present for IgnoredTrustedDomains, not even an explicit nil
### GetLdapProviderId

`func (o *ActiveDirectoryEntry) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *ActiveDirectoryEntry) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *ActiveDirectoryEntry) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *ActiveDirectoryEntry) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *ActiveDirectoryEntry) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *ActiveDirectoryEntry) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *ActiveDirectoryEntry) GetMachineAccounts() []string`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *ActiveDirectoryEntry) GetMachineAccountsOk() (*[]string, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *ActiveDirectoryEntry) SetMachineAccounts(v []string)`

SetMachineAccounts sets MachineAccounts field to given value.

### HasMachineAccounts

`func (o *ActiveDirectoryEntry) HasMachineAccounts() bool`

HasMachineAccounts returns a boolean if a field has been set.

### SetMachineAccountsNil

`func (o *ActiveDirectoryEntry) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *ActiveDirectoryEntry) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetOuName

`func (o *ActiveDirectoryEntry) GetOuName() string`

GetOuName returns the OuName field if non-nil, zero value otherwise.

### GetOuNameOk

`func (o *ActiveDirectoryEntry) GetOuNameOk() (*string, bool)`

GetOuNameOk returns a tuple with the OuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuName

`func (o *ActiveDirectoryEntry) SetOuName(v string)`

SetOuName sets OuName field to given value.

### HasOuName

`func (o *ActiveDirectoryEntry) HasOuName() bool`

HasOuName returns a boolean if a field has been set.

### SetOuNameNil

`func (o *ActiveDirectoryEntry) SetOuNameNil(b bool)`

 SetOuNameNil sets the value for OuName to be an explicit nil

### UnsetOuName
`func (o *ActiveDirectoryEntry) UnsetOuName()`

UnsetOuName ensures that no value is present for OuName, not even an explicit nil
### GetPassword

`func (o *ActiveDirectoryEntry) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ActiveDirectoryEntry) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ActiveDirectoryEntry) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ActiveDirectoryEntry) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ActiveDirectoryEntry) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ActiveDirectoryEntry) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPreferredDomainControllers

`func (o *ActiveDirectoryEntry) GetPreferredDomainControllers() []PreferredDomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *ActiveDirectoryEntry) GetPreferredDomainControllersOk() (*[]PreferredDomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *ActiveDirectoryEntry) SetPreferredDomainControllers(v []PreferredDomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *ActiveDirectoryEntry) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *ActiveDirectoryEntry) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *ActiveDirectoryEntry) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTaskPath

`func (o *ActiveDirectoryEntry) GetTaskPath() string`

GetTaskPath returns the TaskPath field if non-nil, zero value otherwise.

### GetTaskPathOk

`func (o *ActiveDirectoryEntry) GetTaskPathOk() (*string, bool)`

GetTaskPathOk returns a tuple with the TaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPath

`func (o *ActiveDirectoryEntry) SetTaskPath(v string)`

SetTaskPath sets TaskPath field to given value.

### HasTaskPath

`func (o *ActiveDirectoryEntry) HasTaskPath() bool`

HasTaskPath returns a boolean if a field has been set.

### SetTaskPathNil

`func (o *ActiveDirectoryEntry) SetTaskPathNil(b bool)`

 SetTaskPathNil sets the value for TaskPath to be an explicit nil

### UnsetTaskPath
`func (o *ActiveDirectoryEntry) UnsetTaskPath()`

UnsetTaskPath ensures that no value is present for TaskPath, not even an explicit nil
### GetTenantId

`func (o *ActiveDirectoryEntry) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ActiveDirectoryEntry) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ActiveDirectoryEntry) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ActiveDirectoryEntry) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ActiveDirectoryEntry) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ActiveDirectoryEntry) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTrustedDomains

`func (o *ActiveDirectoryEntry) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *ActiveDirectoryEntry) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *ActiveDirectoryEntry) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *ActiveDirectoryEntry) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### SetTrustedDomainsNil

`func (o *ActiveDirectoryEntry) SetTrustedDomainsNil(b bool)`

 SetTrustedDomainsNil sets the value for TrustedDomains to be an explicit nil

### UnsetTrustedDomains
`func (o *ActiveDirectoryEntry) UnsetTrustedDomains()`

UnsetTrustedDomains ensures that no value is present for TrustedDomains, not even an explicit nil
### GetTrustedDomainsEnabled

`func (o *ActiveDirectoryEntry) GetTrustedDomainsEnabled() bool`

GetTrustedDomainsEnabled returns the TrustedDomainsEnabled field if non-nil, zero value otherwise.

### GetTrustedDomainsEnabledOk

`func (o *ActiveDirectoryEntry) GetTrustedDomainsEnabledOk() (*bool, bool)`

GetTrustedDomainsEnabledOk returns a tuple with the TrustedDomainsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsEnabled

`func (o *ActiveDirectoryEntry) SetTrustedDomainsEnabled(v bool)`

SetTrustedDomainsEnabled sets TrustedDomainsEnabled field to given value.

### HasTrustedDomainsEnabled

`func (o *ActiveDirectoryEntry) HasTrustedDomainsEnabled() bool`

HasTrustedDomainsEnabled returns a boolean if a field has been set.

### SetTrustedDomainsEnabledNil

`func (o *ActiveDirectoryEntry) SetTrustedDomainsEnabledNil(b bool)`

 SetTrustedDomainsEnabledNil sets the value for TrustedDomainsEnabled to be an explicit nil

### UnsetTrustedDomainsEnabled
`func (o *ActiveDirectoryEntry) UnsetTrustedDomainsEnabled()`

UnsetTrustedDomainsEnabled ensures that no value is present for TrustedDomainsEnabled, not even an explicit nil
### GetUnixRootSid

`func (o *ActiveDirectoryEntry) GetUnixRootSid() string`

GetUnixRootSid returns the UnixRootSid field if non-nil, zero value otherwise.

### GetUnixRootSidOk

`func (o *ActiveDirectoryEntry) GetUnixRootSidOk() (*string, bool)`

GetUnixRootSidOk returns a tuple with the UnixRootSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixRootSid

`func (o *ActiveDirectoryEntry) SetUnixRootSid(v string)`

SetUnixRootSid sets UnixRootSid field to given value.

### HasUnixRootSid

`func (o *ActiveDirectoryEntry) HasUnixRootSid() bool`

HasUnixRootSid returns a boolean if a field has been set.

### SetUnixRootSidNil

`func (o *ActiveDirectoryEntry) SetUnixRootSidNil(b bool)`

 SetUnixRootSidNil sets the value for UnixRootSid to be an explicit nil

### UnsetUnixRootSid
`func (o *ActiveDirectoryEntry) UnsetUnixRootSid()`

UnsetUnixRootSid ensures that no value is present for UnixRootSid, not even an explicit nil
### GetUserIdMappingInfo

`func (o *ActiveDirectoryEntry) GetUserIdMappingInfo() UserIdMapping`

GetUserIdMappingInfo returns the UserIdMappingInfo field if non-nil, zero value otherwise.

### GetUserIdMappingInfoOk

`func (o *ActiveDirectoryEntry) GetUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetUserIdMappingInfoOk returns a tuple with the UserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdMappingInfo

`func (o *ActiveDirectoryEntry) SetUserIdMappingInfo(v UserIdMapping)`

SetUserIdMappingInfo sets UserIdMappingInfo field to given value.

### HasUserIdMappingInfo

`func (o *ActiveDirectoryEntry) HasUserIdMappingInfo() bool`

HasUserIdMappingInfo returns a boolean if a field has been set.

### GetUserName

`func (o *ActiveDirectoryEntry) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ActiveDirectoryEntry) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ActiveDirectoryEntry) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ActiveDirectoryEntry) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ActiveDirectoryEntry) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ActiveDirectoryEntry) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetWorkgroup

`func (o *ActiveDirectoryEntry) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *ActiveDirectoryEntry) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *ActiveDirectoryEntry) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *ActiveDirectoryEntry) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### SetWorkgroupNil

`func (o *ActiveDirectoryEntry) SetWorkgroupNil(b bool)`

 SetWorkgroupNil sets the value for Workgroup to be an explicit nil

### UnsetWorkgroup
`func (o *ActiveDirectoryEntry) UnsetWorkgroup()`

UnsetWorkgroup ensures that no value is present for Workgroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


