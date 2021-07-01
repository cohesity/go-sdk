/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// ActiveDirectoryEntry Specifies the join settings for a Microsoft Active Directory domain.
type ActiveDirectoryEntry struct {
	// Specifies the fully qualified domain name (FQDN) of an Active Directory.
	DomainName NullableString `json:"domainName,omitempty"`
	FallbackUserIdMappingInfo *UserIdMapping `json:"fallbackUserIdMappingInfo,omitempty"`
	// Specifies the list of trusted domains that were set by the user to be ignored during trusted domain discovery.
	IgnoredTrustedDomains []string `json:"ignoredTrustedDomains,omitempty"`
	// Specifies the LDAP provider id which is map to this Active Directory
	LdapProviderId NullableInt64 `json:"ldapProviderId,omitempty"`
	// Array of Machine Accounts.  Specifies an array of computer names used to identify the Cohesity Cluster on the domain.
	MachineAccounts []string `json:"machineAccounts,omitempty"`
	// Specifies an optional Organizational Unit name.
	OuName NullableString `json:"ouName,omitempty"`
	// Specifies the password for the specified userName.
	Password NullableString `json:"password,omitempty"`
	// Specifies Map of Active Directory domain names to its preferred domain controllers.
	PreferredDomainControllers []PreferredDomainController `json:"preferredDomainControllers,omitempty"`
	// Specifies the task path for AD joining task.
	TaskPath NullableString `json:"taskPath,omitempty"`
	// Specifies the unique id of the tenant.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the trusted domains of the Active Directory domain.
	TrustedDomains []string `json:"trustedDomains,omitempty"`
	// Specifies whether Trusted Domain discovery is disabled.
	TrustedDomainsEnabled NullableBool `json:"trustedDomainsEnabled,omitempty"`
	// Specifies the SID of the Active Directory domain user to be mapped to Unix root user.
	UnixRootSid NullableString `json:"unixRootSid,omitempty"`
	UserIdMappingInfo *UserIdMapping `json:"userIdMappingInfo,omitempty"`
	// Specifies a userName that has administrative privileges in the domain.
	UserName NullableString `json:"userName,omitempty"`
	// Specifies an optional Workgroup name.
	Workgroup NullableString `json:"workgroup,omitempty"`
}

// NewActiveDirectoryEntry instantiates a new ActiveDirectoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectoryEntry() *ActiveDirectoryEntry {
	this := ActiveDirectoryEntry{}
	return &this
}

// NewActiveDirectoryEntryWithDefaults instantiates a new ActiveDirectoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryEntryWithDefaults() *ActiveDirectoryEntry {
	this := ActiveDirectoryEntry{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetDomainName() string {
	if o == nil || o.DomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasDomainName() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *ActiveDirectoryEntry) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *ActiveDirectoryEntry) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetFallbackUserIdMappingInfo returns the FallbackUserIdMappingInfo field value if set, zero value otherwise.
func (o *ActiveDirectoryEntry) GetFallbackUserIdMappingInfo() UserIdMapping {
	if o == nil || o.FallbackUserIdMappingInfo == nil {
		var ret UserIdMapping
		return ret
	}
	return *o.FallbackUserIdMappingInfo
}

// GetFallbackUserIdMappingInfoOk returns a tuple with the FallbackUserIdMappingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryEntry) GetFallbackUserIdMappingInfoOk() (*UserIdMapping, bool) {
	if o == nil || o.FallbackUserIdMappingInfo == nil {
		return nil, false
	}
	return o.FallbackUserIdMappingInfo, true
}

// HasFallbackUserIdMappingInfo returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasFallbackUserIdMappingInfo() bool {
	if o != nil && o.FallbackUserIdMappingInfo != nil {
		return true
	}

	return false
}

// SetFallbackUserIdMappingInfo gets a reference to the given UserIdMapping and assigns it to the FallbackUserIdMappingInfo field.
func (o *ActiveDirectoryEntry) SetFallbackUserIdMappingInfo(v UserIdMapping) {
	o.FallbackUserIdMappingInfo = &v
}

// GetIgnoredTrustedDomains returns the IgnoredTrustedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetIgnoredTrustedDomains() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IgnoredTrustedDomains
}

// GetIgnoredTrustedDomainsOk returns a tuple with the IgnoredTrustedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetIgnoredTrustedDomainsOk() (*[]string, bool) {
	if o == nil || o.IgnoredTrustedDomains == nil {
		return nil, false
	}
	return &o.IgnoredTrustedDomains, true
}

// HasIgnoredTrustedDomains returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasIgnoredTrustedDomains() bool {
	if o != nil && o.IgnoredTrustedDomains != nil {
		return true
	}

	return false
}

// SetIgnoredTrustedDomains gets a reference to the given []string and assigns it to the IgnoredTrustedDomains field.
func (o *ActiveDirectoryEntry) SetIgnoredTrustedDomains(v []string) {
	o.IgnoredTrustedDomains = v
}

// GetLdapProviderId returns the LdapProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetLdapProviderId() int64 {
	if o == nil || o.LdapProviderId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LdapProviderId.Get()
}

// GetLdapProviderIdOk returns a tuple with the LdapProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetLdapProviderIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LdapProviderId.Get(), o.LdapProviderId.IsSet()
}

// HasLdapProviderId returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasLdapProviderId() bool {
	if o != nil && o.LdapProviderId.IsSet() {
		return true
	}

	return false
}

// SetLdapProviderId gets a reference to the given NullableInt64 and assigns it to the LdapProviderId field.
func (o *ActiveDirectoryEntry) SetLdapProviderId(v int64) {
	o.LdapProviderId.Set(&v)
}
// SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil
func (o *ActiveDirectoryEntry) SetLdapProviderIdNil() {
	o.LdapProviderId.Set(nil)
}

// UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetLdapProviderId() {
	o.LdapProviderId.Unset()
}

// GetMachineAccounts returns the MachineAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetMachineAccounts() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.MachineAccounts
}

// GetMachineAccountsOk returns a tuple with the MachineAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetMachineAccountsOk() (*[]string, bool) {
	if o == nil || o.MachineAccounts == nil {
		return nil, false
	}
	return &o.MachineAccounts, true
}

// HasMachineAccounts returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasMachineAccounts() bool {
	if o != nil && o.MachineAccounts != nil {
		return true
	}

	return false
}

// SetMachineAccounts gets a reference to the given []string and assigns it to the MachineAccounts field.
func (o *ActiveDirectoryEntry) SetMachineAccounts(v []string) {
	o.MachineAccounts = v
}

// GetOuName returns the OuName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetOuName() string {
	if o == nil || o.OuName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OuName.Get()
}

// GetOuNameOk returns a tuple with the OuName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetOuNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OuName.Get(), o.OuName.IsSet()
}

// HasOuName returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasOuName() bool {
	if o != nil && o.OuName.IsSet() {
		return true
	}

	return false
}

// SetOuName gets a reference to the given NullableString and assigns it to the OuName field.
func (o *ActiveDirectoryEntry) SetOuName(v string) {
	o.OuName.Set(&v)
}
// SetOuNameNil sets the value for OuName to be an explicit nil
func (o *ActiveDirectoryEntry) SetOuNameNil() {
	o.OuName.Set(nil)
}

// UnsetOuName ensures that no value is present for OuName, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetOuName() {
	o.OuName.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ActiveDirectoryEntry) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ActiveDirectoryEntry) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetPassword() {
	o.Password.Unset()
}

// GetPreferredDomainControllers returns the PreferredDomainControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetPreferredDomainControllers() []PreferredDomainController {
	if o == nil  {
		var ret []PreferredDomainController
		return ret
	}
	return o.PreferredDomainControllers
}

// GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetPreferredDomainControllersOk() (*[]PreferredDomainController, bool) {
	if o == nil || o.PreferredDomainControllers == nil {
		return nil, false
	}
	return &o.PreferredDomainControllers, true
}

// HasPreferredDomainControllers returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasPreferredDomainControllers() bool {
	if o != nil && o.PreferredDomainControllers != nil {
		return true
	}

	return false
}

// SetPreferredDomainControllers gets a reference to the given []PreferredDomainController and assigns it to the PreferredDomainControllers field.
func (o *ActiveDirectoryEntry) SetPreferredDomainControllers(v []PreferredDomainController) {
	o.PreferredDomainControllers = v
}

// GetTaskPath returns the TaskPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetTaskPath() string {
	if o == nil || o.TaskPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskPath.Get()
}

// GetTaskPathOk returns a tuple with the TaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetTaskPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaskPath.Get(), o.TaskPath.IsSet()
}

// HasTaskPath returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasTaskPath() bool {
	if o != nil && o.TaskPath.IsSet() {
		return true
	}

	return false
}

// SetTaskPath gets a reference to the given NullableString and assigns it to the TaskPath field.
func (o *ActiveDirectoryEntry) SetTaskPath(v string) {
	o.TaskPath.Set(&v)
}
// SetTaskPathNil sets the value for TaskPath to be an explicit nil
func (o *ActiveDirectoryEntry) SetTaskPathNil() {
	o.TaskPath.Set(nil)
}

// UnsetTaskPath ensures that no value is present for TaskPath, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetTaskPath() {
	o.TaskPath.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ActiveDirectoryEntry) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ActiveDirectoryEntry) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTrustedDomains returns the TrustedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetTrustedDomains() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TrustedDomains
}

// GetTrustedDomainsOk returns a tuple with the TrustedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetTrustedDomainsOk() (*[]string, bool) {
	if o == nil || o.TrustedDomains == nil {
		return nil, false
	}
	return &o.TrustedDomains, true
}

// HasTrustedDomains returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasTrustedDomains() bool {
	if o != nil && o.TrustedDomains != nil {
		return true
	}

	return false
}

// SetTrustedDomains gets a reference to the given []string and assigns it to the TrustedDomains field.
func (o *ActiveDirectoryEntry) SetTrustedDomains(v []string) {
	o.TrustedDomains = v
}

// GetTrustedDomainsEnabled returns the TrustedDomainsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetTrustedDomainsEnabled() bool {
	if o == nil || o.TrustedDomainsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TrustedDomainsEnabled.Get()
}

// GetTrustedDomainsEnabledOk returns a tuple with the TrustedDomainsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetTrustedDomainsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrustedDomainsEnabled.Get(), o.TrustedDomainsEnabled.IsSet()
}

// HasTrustedDomainsEnabled returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasTrustedDomainsEnabled() bool {
	if o != nil && o.TrustedDomainsEnabled.IsSet() {
		return true
	}

	return false
}

// SetTrustedDomainsEnabled gets a reference to the given NullableBool and assigns it to the TrustedDomainsEnabled field.
func (o *ActiveDirectoryEntry) SetTrustedDomainsEnabled(v bool) {
	o.TrustedDomainsEnabled.Set(&v)
}
// SetTrustedDomainsEnabledNil sets the value for TrustedDomainsEnabled to be an explicit nil
func (o *ActiveDirectoryEntry) SetTrustedDomainsEnabledNil() {
	o.TrustedDomainsEnabled.Set(nil)
}

// UnsetTrustedDomainsEnabled ensures that no value is present for TrustedDomainsEnabled, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetTrustedDomainsEnabled() {
	o.TrustedDomainsEnabled.Unset()
}

// GetUnixRootSid returns the UnixRootSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetUnixRootSid() string {
	if o == nil || o.UnixRootSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnixRootSid.Get()
}

// GetUnixRootSidOk returns a tuple with the UnixRootSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetUnixRootSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnixRootSid.Get(), o.UnixRootSid.IsSet()
}

// HasUnixRootSid returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasUnixRootSid() bool {
	if o != nil && o.UnixRootSid.IsSet() {
		return true
	}

	return false
}

// SetUnixRootSid gets a reference to the given NullableString and assigns it to the UnixRootSid field.
func (o *ActiveDirectoryEntry) SetUnixRootSid(v string) {
	o.UnixRootSid.Set(&v)
}
// SetUnixRootSidNil sets the value for UnixRootSid to be an explicit nil
func (o *ActiveDirectoryEntry) SetUnixRootSidNil() {
	o.UnixRootSid.Set(nil)
}

// UnsetUnixRootSid ensures that no value is present for UnixRootSid, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetUnixRootSid() {
	o.UnixRootSid.Unset()
}

// GetUserIdMappingInfo returns the UserIdMappingInfo field value if set, zero value otherwise.
func (o *ActiveDirectoryEntry) GetUserIdMappingInfo() UserIdMapping {
	if o == nil || o.UserIdMappingInfo == nil {
		var ret UserIdMapping
		return ret
	}
	return *o.UserIdMappingInfo
}

// GetUserIdMappingInfoOk returns a tuple with the UserIdMappingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectoryEntry) GetUserIdMappingInfoOk() (*UserIdMapping, bool) {
	if o == nil || o.UserIdMappingInfo == nil {
		return nil, false
	}
	return o.UserIdMappingInfo, true
}

// HasUserIdMappingInfo returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasUserIdMappingInfo() bool {
	if o != nil && o.UserIdMappingInfo != nil {
		return true
	}

	return false
}

// SetUserIdMappingInfo gets a reference to the given UserIdMapping and assigns it to the UserIdMappingInfo field.
func (o *ActiveDirectoryEntry) SetUserIdMappingInfo(v UserIdMapping) {
	o.UserIdMappingInfo = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *ActiveDirectoryEntry) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *ActiveDirectoryEntry) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetUserName() {
	o.UserName.Unset()
}

// GetWorkgroup returns the Workgroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectoryEntry) GetWorkgroup() string {
	if o == nil || o.Workgroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.Workgroup.Get()
}

// GetWorkgroupOk returns a tuple with the Workgroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectoryEntry) GetWorkgroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Workgroup.Get(), o.Workgroup.IsSet()
}

// HasWorkgroup returns a boolean if a field has been set.
func (o *ActiveDirectoryEntry) HasWorkgroup() bool {
	if o != nil && o.Workgroup.IsSet() {
		return true
	}

	return false
}

// SetWorkgroup gets a reference to the given NullableString and assigns it to the Workgroup field.
func (o *ActiveDirectoryEntry) SetWorkgroup(v string) {
	o.Workgroup.Set(&v)
}
// SetWorkgroupNil sets the value for Workgroup to be an explicit nil
func (o *ActiveDirectoryEntry) SetWorkgroupNil() {
	o.Workgroup.Set(nil)
}

// UnsetWorkgroup ensures that no value is present for Workgroup, not even an explicit nil
func (o *ActiveDirectoryEntry) UnsetWorkgroup() {
	o.Workgroup.Unset()
}

func (o ActiveDirectoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if o.FallbackUserIdMappingInfo != nil {
		toSerialize["fallbackUserIdMappingInfo"] = o.FallbackUserIdMappingInfo
	}
	if o.IgnoredTrustedDomains != nil {
		toSerialize["ignoredTrustedDomains"] = o.IgnoredTrustedDomains
	}
	if o.LdapProviderId.IsSet() {
		toSerialize["ldapProviderId"] = o.LdapProviderId.Get()
	}
	if o.MachineAccounts != nil {
		toSerialize["machineAccounts"] = o.MachineAccounts
	}
	if o.OuName.IsSet() {
		toSerialize["ouName"] = o.OuName.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PreferredDomainControllers != nil {
		toSerialize["preferredDomainControllers"] = o.PreferredDomainControllers
	}
	if o.TaskPath.IsSet() {
		toSerialize["taskPath"] = o.TaskPath.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.TrustedDomains != nil {
		toSerialize["trustedDomains"] = o.TrustedDomains
	}
	if o.TrustedDomainsEnabled.IsSet() {
		toSerialize["trustedDomainsEnabled"] = o.TrustedDomainsEnabled.Get()
	}
	if o.UnixRootSid.IsSet() {
		toSerialize["unixRootSid"] = o.UnixRootSid.Get()
	}
	if o.UserIdMappingInfo != nil {
		toSerialize["userIdMappingInfo"] = o.UserIdMappingInfo
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Workgroup.IsSet() {
		toSerialize["workgroup"] = o.Workgroup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableActiveDirectoryEntry struct {
	value *ActiveDirectoryEntry
	isSet bool
}

func (v NullableActiveDirectoryEntry) Get() *ActiveDirectoryEntry {
	return v.value
}

func (v *NullableActiveDirectoryEntry) Set(val *ActiveDirectoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectoryEntry(val *ActiveDirectoryEntry) *NullableActiveDirectoryEntry {
	return &NullableActiveDirectoryEntry{value: val, isSet: true}
}

func (v NullableActiveDirectoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


