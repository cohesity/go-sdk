# CreateActiveDirectoryEntryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Specifies the fully qualified domain name (FQDN) of an Active Directory. | [optional] 
**FallbackUserIdMappingInfo** | Pointer to [**UserIdMapping**](UserIdMapping.md) |  | [optional] 
**IgnoredTrustedDomains** | Pointer to **[]string** | Specifies the list of trusted domains that were set by the user to be ignored during trusted domain discovery. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id which is map to this Active Directory | [optional] 
**MachineAccounts** | Pointer to **[]string** | Array of Machine Accounts.  Specifies an array of computer names used to identify the Cohesity Cluster on the domain. | [optional] 
**OuName** | Pointer to **NullableString** | Specifies an optional Organizational Unit name. | [optional] 
**OverwriteExistingAccounts** | Pointer to **NullableBool** | Specifies whether the specified machine accounts should overwrite the existing machine accounts in this domain. | [optional] 
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

### NewCreateActiveDirectoryEntryParams

`func NewCreateActiveDirectoryEntryParams() *CreateActiveDirectoryEntryParams`

NewCreateActiveDirectoryEntryParams instantiates a new CreateActiveDirectoryEntryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActiveDirectoryEntryParamsWithDefaults

`func NewCreateActiveDirectoryEntryParamsWithDefaults() *CreateActiveDirectoryEntryParams`

NewCreateActiveDirectoryEntryParamsWithDefaults instantiates a new CreateActiveDirectoryEntryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CreateActiveDirectoryEntryParams) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateActiveDirectoryEntryParams) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateActiveDirectoryEntryParams) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateActiveDirectoryEntryParams) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *CreateActiveDirectoryEntryParams) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CreateActiveDirectoryEntryParams) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetFallbackUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) GetFallbackUserIdMappingInfo() UserIdMapping`

GetFallbackUserIdMappingInfo returns the FallbackUserIdMappingInfo field if non-nil, zero value otherwise.

### GetFallbackUserIdMappingInfoOk

`func (o *CreateActiveDirectoryEntryParams) GetFallbackUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetFallbackUserIdMappingInfoOk returns a tuple with the FallbackUserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) SetFallbackUserIdMappingInfo(v UserIdMapping)`

SetFallbackUserIdMappingInfo sets FallbackUserIdMappingInfo field to given value.

### HasFallbackUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) HasFallbackUserIdMappingInfo() bool`

HasFallbackUserIdMappingInfo returns a boolean if a field has been set.

### GetIgnoredTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) GetIgnoredTrustedDomains() []string`

GetIgnoredTrustedDomains returns the IgnoredTrustedDomains field if non-nil, zero value otherwise.

### GetIgnoredTrustedDomainsOk

`func (o *CreateActiveDirectoryEntryParams) GetIgnoredTrustedDomainsOk() (*[]string, bool)`

GetIgnoredTrustedDomainsOk returns a tuple with the IgnoredTrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) SetIgnoredTrustedDomains(v []string)`

SetIgnoredTrustedDomains sets IgnoredTrustedDomains field to given value.

### HasIgnoredTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) HasIgnoredTrustedDomains() bool`

HasIgnoredTrustedDomains returns a boolean if a field has been set.

### SetIgnoredTrustedDomainsNil

`func (o *CreateActiveDirectoryEntryParams) SetIgnoredTrustedDomainsNil(b bool)`

 SetIgnoredTrustedDomainsNil sets the value for IgnoredTrustedDomains to be an explicit nil

### UnsetIgnoredTrustedDomains
`func (o *CreateActiveDirectoryEntryParams) UnsetIgnoredTrustedDomains()`

UnsetIgnoredTrustedDomains ensures that no value is present for IgnoredTrustedDomains, not even an explicit nil
### GetLdapProviderId

`func (o *CreateActiveDirectoryEntryParams) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *CreateActiveDirectoryEntryParams) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *CreateActiveDirectoryEntryParams) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *CreateActiveDirectoryEntryParams) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *CreateActiveDirectoryEntryParams) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *CreateActiveDirectoryEntryParams) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetMachineAccounts

`func (o *CreateActiveDirectoryEntryParams) GetMachineAccounts() []string`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *CreateActiveDirectoryEntryParams) GetMachineAccountsOk() (*[]string, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *CreateActiveDirectoryEntryParams) SetMachineAccounts(v []string)`

SetMachineAccounts sets MachineAccounts field to given value.

### HasMachineAccounts

`func (o *CreateActiveDirectoryEntryParams) HasMachineAccounts() bool`

HasMachineAccounts returns a boolean if a field has been set.

### SetMachineAccountsNil

`func (o *CreateActiveDirectoryEntryParams) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *CreateActiveDirectoryEntryParams) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetOuName

`func (o *CreateActiveDirectoryEntryParams) GetOuName() string`

GetOuName returns the OuName field if non-nil, zero value otherwise.

### GetOuNameOk

`func (o *CreateActiveDirectoryEntryParams) GetOuNameOk() (*string, bool)`

GetOuNameOk returns a tuple with the OuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuName

`func (o *CreateActiveDirectoryEntryParams) SetOuName(v string)`

SetOuName sets OuName field to given value.

### HasOuName

`func (o *CreateActiveDirectoryEntryParams) HasOuName() bool`

HasOuName returns a boolean if a field has been set.

### SetOuNameNil

`func (o *CreateActiveDirectoryEntryParams) SetOuNameNil(b bool)`

 SetOuNameNil sets the value for OuName to be an explicit nil

### UnsetOuName
`func (o *CreateActiveDirectoryEntryParams) UnsetOuName()`

UnsetOuName ensures that no value is present for OuName, not even an explicit nil
### GetOverwriteExistingAccounts

`func (o *CreateActiveDirectoryEntryParams) GetOverwriteExistingAccounts() bool`

GetOverwriteExistingAccounts returns the OverwriteExistingAccounts field if non-nil, zero value otherwise.

### GetOverwriteExistingAccountsOk

`func (o *CreateActiveDirectoryEntryParams) GetOverwriteExistingAccountsOk() (*bool, bool)`

GetOverwriteExistingAccountsOk returns a tuple with the OverwriteExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingAccounts

`func (o *CreateActiveDirectoryEntryParams) SetOverwriteExistingAccounts(v bool)`

SetOverwriteExistingAccounts sets OverwriteExistingAccounts field to given value.

### HasOverwriteExistingAccounts

`func (o *CreateActiveDirectoryEntryParams) HasOverwriteExistingAccounts() bool`

HasOverwriteExistingAccounts returns a boolean if a field has been set.

### SetOverwriteExistingAccountsNil

`func (o *CreateActiveDirectoryEntryParams) SetOverwriteExistingAccountsNil(b bool)`

 SetOverwriteExistingAccountsNil sets the value for OverwriteExistingAccounts to be an explicit nil

### UnsetOverwriteExistingAccounts
`func (o *CreateActiveDirectoryEntryParams) UnsetOverwriteExistingAccounts()`

UnsetOverwriteExistingAccounts ensures that no value is present for OverwriteExistingAccounts, not even an explicit nil
### GetPassword

`func (o *CreateActiveDirectoryEntryParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateActiveDirectoryEntryParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateActiveDirectoryEntryParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateActiveDirectoryEntryParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateActiveDirectoryEntryParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateActiveDirectoryEntryParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPreferredDomainControllers

`func (o *CreateActiveDirectoryEntryParams) GetPreferredDomainControllers() []PreferredDomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *CreateActiveDirectoryEntryParams) GetPreferredDomainControllersOk() (*[]PreferredDomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *CreateActiveDirectoryEntryParams) SetPreferredDomainControllers(v []PreferredDomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *CreateActiveDirectoryEntryParams) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.

### SetPreferredDomainControllersNil

`func (o *CreateActiveDirectoryEntryParams) SetPreferredDomainControllersNil(b bool)`

 SetPreferredDomainControllersNil sets the value for PreferredDomainControllers to be an explicit nil

### UnsetPreferredDomainControllers
`func (o *CreateActiveDirectoryEntryParams) UnsetPreferredDomainControllers()`

UnsetPreferredDomainControllers ensures that no value is present for PreferredDomainControllers, not even an explicit nil
### GetTaskPath

`func (o *CreateActiveDirectoryEntryParams) GetTaskPath() string`

GetTaskPath returns the TaskPath field if non-nil, zero value otherwise.

### GetTaskPathOk

`func (o *CreateActiveDirectoryEntryParams) GetTaskPathOk() (*string, bool)`

GetTaskPathOk returns a tuple with the TaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPath

`func (o *CreateActiveDirectoryEntryParams) SetTaskPath(v string)`

SetTaskPath sets TaskPath field to given value.

### HasTaskPath

`func (o *CreateActiveDirectoryEntryParams) HasTaskPath() bool`

HasTaskPath returns a boolean if a field has been set.

### SetTaskPathNil

`func (o *CreateActiveDirectoryEntryParams) SetTaskPathNil(b bool)`

 SetTaskPathNil sets the value for TaskPath to be an explicit nil

### UnsetTaskPath
`func (o *CreateActiveDirectoryEntryParams) UnsetTaskPath()`

UnsetTaskPath ensures that no value is present for TaskPath, not even an explicit nil
### GetTenantId

`func (o *CreateActiveDirectoryEntryParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateActiveDirectoryEntryParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateActiveDirectoryEntryParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateActiveDirectoryEntryParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateActiveDirectoryEntryParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateActiveDirectoryEntryParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *CreateActiveDirectoryEntryParams) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *CreateActiveDirectoryEntryParams) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### SetTrustedDomainsNil

`func (o *CreateActiveDirectoryEntryParams) SetTrustedDomainsNil(b bool)`

 SetTrustedDomainsNil sets the value for TrustedDomains to be an explicit nil

### UnsetTrustedDomains
`func (o *CreateActiveDirectoryEntryParams) UnsetTrustedDomains()`

UnsetTrustedDomains ensures that no value is present for TrustedDomains, not even an explicit nil
### GetTrustedDomainsEnabled

`func (o *CreateActiveDirectoryEntryParams) GetTrustedDomainsEnabled() bool`

GetTrustedDomainsEnabled returns the TrustedDomainsEnabled field if non-nil, zero value otherwise.

### GetTrustedDomainsEnabledOk

`func (o *CreateActiveDirectoryEntryParams) GetTrustedDomainsEnabledOk() (*bool, bool)`

GetTrustedDomainsEnabledOk returns a tuple with the TrustedDomainsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomainsEnabled

`func (o *CreateActiveDirectoryEntryParams) SetTrustedDomainsEnabled(v bool)`

SetTrustedDomainsEnabled sets TrustedDomainsEnabled field to given value.

### HasTrustedDomainsEnabled

`func (o *CreateActiveDirectoryEntryParams) HasTrustedDomainsEnabled() bool`

HasTrustedDomainsEnabled returns a boolean if a field has been set.

### SetTrustedDomainsEnabledNil

`func (o *CreateActiveDirectoryEntryParams) SetTrustedDomainsEnabledNil(b bool)`

 SetTrustedDomainsEnabledNil sets the value for TrustedDomainsEnabled to be an explicit nil

### UnsetTrustedDomainsEnabled
`func (o *CreateActiveDirectoryEntryParams) UnsetTrustedDomainsEnabled()`

UnsetTrustedDomainsEnabled ensures that no value is present for TrustedDomainsEnabled, not even an explicit nil
### GetUnixRootSid

`func (o *CreateActiveDirectoryEntryParams) GetUnixRootSid() string`

GetUnixRootSid returns the UnixRootSid field if non-nil, zero value otherwise.

### GetUnixRootSidOk

`func (o *CreateActiveDirectoryEntryParams) GetUnixRootSidOk() (*string, bool)`

GetUnixRootSidOk returns a tuple with the UnixRootSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixRootSid

`func (o *CreateActiveDirectoryEntryParams) SetUnixRootSid(v string)`

SetUnixRootSid sets UnixRootSid field to given value.

### HasUnixRootSid

`func (o *CreateActiveDirectoryEntryParams) HasUnixRootSid() bool`

HasUnixRootSid returns a boolean if a field has been set.

### SetUnixRootSidNil

`func (o *CreateActiveDirectoryEntryParams) SetUnixRootSidNil(b bool)`

 SetUnixRootSidNil sets the value for UnixRootSid to be an explicit nil

### UnsetUnixRootSid
`func (o *CreateActiveDirectoryEntryParams) UnsetUnixRootSid()`

UnsetUnixRootSid ensures that no value is present for UnixRootSid, not even an explicit nil
### GetUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) GetUserIdMappingInfo() UserIdMapping`

GetUserIdMappingInfo returns the UserIdMappingInfo field if non-nil, zero value otherwise.

### GetUserIdMappingInfoOk

`func (o *CreateActiveDirectoryEntryParams) GetUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetUserIdMappingInfoOk returns a tuple with the UserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) SetUserIdMappingInfo(v UserIdMapping)`

SetUserIdMappingInfo sets UserIdMappingInfo field to given value.

### HasUserIdMappingInfo

`func (o *CreateActiveDirectoryEntryParams) HasUserIdMappingInfo() bool`

HasUserIdMappingInfo returns a boolean if a field has been set.

### GetUserName

`func (o *CreateActiveDirectoryEntryParams) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateActiveDirectoryEntryParams) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateActiveDirectoryEntryParams) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CreateActiveDirectoryEntryParams) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *CreateActiveDirectoryEntryParams) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *CreateActiveDirectoryEntryParams) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetWorkgroup

`func (o *CreateActiveDirectoryEntryParams) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *CreateActiveDirectoryEntryParams) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *CreateActiveDirectoryEntryParams) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *CreateActiveDirectoryEntryParams) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### SetWorkgroupNil

`func (o *CreateActiveDirectoryEntryParams) SetWorkgroupNil(b bool)`

 SetWorkgroupNil sets the value for Workgroup to be an explicit nil

### UnsetWorkgroup
`func (o *CreateActiveDirectoryEntryParams) UnsetWorkgroup()`

UnsetWorkgroup ensures that no value is present for Workgroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


