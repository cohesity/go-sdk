/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActiveDirectory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveDirectory{}

// ActiveDirectory Specifies an Active Directory.
type ActiveDirectory struct {
	// Specifies the id of the connection.
	ConnectionId NullableInt64 `json:"connectionId,omitempty"`
	// Specifies a list of denied domain controllers of this Active Directory Domain.
	DomainControllersDenyList []*string `json:"domainControllersDenyList,omitempty"`
	// Specifies the id of the Active Directory.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the LDAP provider id which is mapped to this Active Directory
	LdapProviderId NullableInt64 `json:"ldapProviderId,omitempty"`
	// Specifies a list of computer names used to identify the Cohesity Cluster on the Active Directory domain. The first machine account is used as primary machine account and it can not be modified.
	MachineAccounts []MachineAccount `json:"machineAccounts"`
	// Specifies the name of the NIS Provider which is mapped to this Active Directory.
	NisProviderDomainName NullableString `json:"nisProviderDomainName,omitempty"`
	// Specifies an optional organizational unit name.
	OrganizationalUnitName NullableString `json:"organizationalUnitName,omitempty"`
	// Specifies a list of preferred domain controllers of this Active Directory.
	PreferredDomainControllers []DomainController `json:"preferredDomainControllers,omitempty"`
	TrustedDomainParams NullableCommonActiveDirectoryParamsTrustedDomainParams `json:"trustedDomainParams,omitempty"`
	// Specifies a work group name.
	WorkGroupName NullableString `json:"workGroupName,omitempty"`
	// Specifies a list of Centrify zones.
	CentrifyZones []CentrifyZones `json:"centrifyZones,omitempty"`
	// A list of domain names with a list of it's domain controllers.
	DomainControllers []DomainControllers `json:"domainControllers,omitempty"`
	// Specifies the domain name of the Active Directory.
	DomainName NullableString `json:"domainName,omitempty"`
	Error *ActiveDirectoryError `json:"error,omitempty"`
	// Specifies the params of the user id mapping info of an Active Directory.
	IdMappingParams map[string]interface{} `json:"idMappingParams,omitempty"`
	// Specifies the list of tenants that have permissions for this Active Directory.
	Permissions []Tenant `json:"permissions,omitempty"`
	// Specifies a list of security principals.
	SecurityPrincipals []SecurityPrincipal `json:"securityPrincipals,omitempty"`
	TaskLogs *TaskLogs `json:"taskLogs,omitempty"`
	// Specifies level of transitive Active Directory trust domains to be used.
	TransitiveAdTrustLevelLimit NullableInt32 `json:"transitiveAdTrustLevelLimit,omitempty"`
}

type _ActiveDirectory ActiveDirectory

// NewActiveDirectory instantiates a new ActiveDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveDirectory(machineAccounts []MachineAccount) *ActiveDirectory {
	this := ActiveDirectory{}
	this.MachineAccounts = machineAccounts
	return &this
}

// NewActiveDirectoryWithDefaults instantiates a new ActiveDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryWithDefaults() *ActiveDirectory {
	this := ActiveDirectory{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetConnectionId() int64 {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret int64
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetConnectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *ActiveDirectory) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt64 and assigns it to the ConnectionId field.
func (o *ActiveDirectory) SetConnectionId(v int64) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *ActiveDirectory) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *ActiveDirectory) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetDomainControllersDenyList returns the DomainControllersDenyList field value if set, zero value otherwise.
func (o *ActiveDirectory) GetDomainControllersDenyList() []*string {
	if o == nil || IsNil(o.DomainControllersDenyList) {
		var ret []*string
		return ret
	}
	return o.DomainControllersDenyList
}

// GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectory) GetDomainControllersDenyListOk() ([]*string, bool) {
	if o == nil || IsNil(o.DomainControllersDenyList) {
		return nil, false
	}
	return o.DomainControllersDenyList, true
}

// HasDomainControllersDenyList returns a boolean if a field has been set.
func (o *ActiveDirectory) HasDomainControllersDenyList() bool {
	if o != nil && !IsNil(o.DomainControllersDenyList) {
		return true
	}

	return false
}

// SetDomainControllersDenyList gets a reference to the given []*string and assigns it to the DomainControllersDenyList field.
func (o *ActiveDirectory) SetDomainControllersDenyList(v []*string) {
	o.DomainControllersDenyList = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ActiveDirectory) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ActiveDirectory) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ActiveDirectory) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ActiveDirectory) UnsetId() {
	o.Id.Unset()
}

// GetLdapProviderId returns the LdapProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetLdapProviderId() int64 {
	if o == nil || IsNil(o.LdapProviderId.Get()) {
		var ret int64
		return ret
	}
	return *o.LdapProviderId.Get()
}

// GetLdapProviderIdOk returns a tuple with the LdapProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetLdapProviderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LdapProviderId.Get(), o.LdapProviderId.IsSet()
}

// HasLdapProviderId returns a boolean if a field has been set.
func (o *ActiveDirectory) HasLdapProviderId() bool {
	if o != nil && o.LdapProviderId.IsSet() {
		return true
	}

	return false
}

// SetLdapProviderId gets a reference to the given NullableInt64 and assigns it to the LdapProviderId field.
func (o *ActiveDirectory) SetLdapProviderId(v int64) {
	o.LdapProviderId.Set(&v)
}
// SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil
func (o *ActiveDirectory) SetLdapProviderIdNil() {
	o.LdapProviderId.Set(nil)
}

// UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
func (o *ActiveDirectory) UnsetLdapProviderId() {
	o.LdapProviderId.Unset()
}

// GetMachineAccounts returns the MachineAccounts field value
// If the value is explicit nil, the zero value for []MachineAccount will be returned
func (o *ActiveDirectory) GetMachineAccounts() []MachineAccount {
	if o == nil {
		var ret []MachineAccount
		return ret
	}

	return o.MachineAccounts
}

// GetMachineAccountsOk returns a tuple with the MachineAccounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetMachineAccountsOk() ([]MachineAccount, bool) {
	if o == nil || IsNil(o.MachineAccounts) {
		return nil, false
	}
	return o.MachineAccounts, true
}

// SetMachineAccounts sets field value
func (o *ActiveDirectory) SetMachineAccounts(v []MachineAccount) {
	o.MachineAccounts = v
}

// GetNisProviderDomainName returns the NisProviderDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetNisProviderDomainName() string {
	if o == nil || IsNil(o.NisProviderDomainName.Get()) {
		var ret string
		return ret
	}
	return *o.NisProviderDomainName.Get()
}

// GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetNisProviderDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NisProviderDomainName.Get(), o.NisProviderDomainName.IsSet()
}

// HasNisProviderDomainName returns a boolean if a field has been set.
func (o *ActiveDirectory) HasNisProviderDomainName() bool {
	if o != nil && o.NisProviderDomainName.IsSet() {
		return true
	}

	return false
}

// SetNisProviderDomainName gets a reference to the given NullableString and assigns it to the NisProviderDomainName field.
func (o *ActiveDirectory) SetNisProviderDomainName(v string) {
	o.NisProviderDomainName.Set(&v)
}
// SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil
func (o *ActiveDirectory) SetNisProviderDomainNameNil() {
	o.NisProviderDomainName.Set(nil)
}

// UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
func (o *ActiveDirectory) UnsetNisProviderDomainName() {
	o.NisProviderDomainName.Unset()
}

// GetOrganizationalUnitName returns the OrganizationalUnitName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetOrganizationalUnitName() string {
	if o == nil || IsNil(o.OrganizationalUnitName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationalUnitName.Get()
}

// GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetOrganizationalUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationalUnitName.Get(), o.OrganizationalUnitName.IsSet()
}

// HasOrganizationalUnitName returns a boolean if a field has been set.
func (o *ActiveDirectory) HasOrganizationalUnitName() bool {
	if o != nil && o.OrganizationalUnitName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationalUnitName gets a reference to the given NullableString and assigns it to the OrganizationalUnitName field.
func (o *ActiveDirectory) SetOrganizationalUnitName(v string) {
	o.OrganizationalUnitName.Set(&v)
}
// SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil
func (o *ActiveDirectory) SetOrganizationalUnitNameNil() {
	o.OrganizationalUnitName.Set(nil)
}

// UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
func (o *ActiveDirectory) UnsetOrganizationalUnitName() {
	o.OrganizationalUnitName.Unset()
}

// GetPreferredDomainControllers returns the PreferredDomainControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetPreferredDomainControllers() []DomainController {
	if o == nil {
		var ret []DomainController
		return ret
	}
	return o.PreferredDomainControllers
}

// GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetPreferredDomainControllersOk() ([]DomainController, bool) {
	if o == nil || IsNil(o.PreferredDomainControllers) {
		return nil, false
	}
	return o.PreferredDomainControllers, true
}

// HasPreferredDomainControllers returns a boolean if a field has been set.
func (o *ActiveDirectory) HasPreferredDomainControllers() bool {
	if o != nil && !IsNil(o.PreferredDomainControllers) {
		return true
	}

	return false
}

// SetPreferredDomainControllers gets a reference to the given []DomainController and assigns it to the PreferredDomainControllers field.
func (o *ActiveDirectory) SetPreferredDomainControllers(v []DomainController) {
	o.PreferredDomainControllers = v
}

// GetTrustedDomainParams returns the TrustedDomainParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams {
	if o == nil || IsNil(o.TrustedDomainParams.Get()) {
		var ret CommonActiveDirectoryParamsTrustedDomainParams
		return ret
	}
	return *o.TrustedDomainParams.Get()
}

// GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustedDomainParams.Get(), o.TrustedDomainParams.IsSet()
}

// HasTrustedDomainParams returns a boolean if a field has been set.
func (o *ActiveDirectory) HasTrustedDomainParams() bool {
	if o != nil && o.TrustedDomainParams.IsSet() {
		return true
	}

	return false
}

// SetTrustedDomainParams gets a reference to the given NullableCommonActiveDirectoryParamsTrustedDomainParams and assigns it to the TrustedDomainParams field.
func (o *ActiveDirectory) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams) {
	o.TrustedDomainParams.Set(&v)
}
// SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil
func (o *ActiveDirectory) SetTrustedDomainParamsNil() {
	o.TrustedDomainParams.Set(nil)
}

// UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
func (o *ActiveDirectory) UnsetTrustedDomainParams() {
	o.TrustedDomainParams.Unset()
}

// GetWorkGroupName returns the WorkGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetWorkGroupName() string {
	if o == nil || IsNil(o.WorkGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.WorkGroupName.Get()
}

// GetWorkGroupNameOk returns a tuple with the WorkGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetWorkGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkGroupName.Get(), o.WorkGroupName.IsSet()
}

// HasWorkGroupName returns a boolean if a field has been set.
func (o *ActiveDirectory) HasWorkGroupName() bool {
	if o != nil && o.WorkGroupName.IsSet() {
		return true
	}

	return false
}

// SetWorkGroupName gets a reference to the given NullableString and assigns it to the WorkGroupName field.
func (o *ActiveDirectory) SetWorkGroupName(v string) {
	o.WorkGroupName.Set(&v)
}
// SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil
func (o *ActiveDirectory) SetWorkGroupNameNil() {
	o.WorkGroupName.Set(nil)
}

// UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil
func (o *ActiveDirectory) UnsetWorkGroupName() {
	o.WorkGroupName.Unset()
}

// GetCentrifyZones returns the CentrifyZones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetCentrifyZones() []CentrifyZones {
	if o == nil {
		var ret []CentrifyZones
		return ret
	}
	return o.CentrifyZones
}

// GetCentrifyZonesOk returns a tuple with the CentrifyZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetCentrifyZonesOk() ([]CentrifyZones, bool) {
	if o == nil || IsNil(o.CentrifyZones) {
		return nil, false
	}
	return o.CentrifyZones, true
}

// HasCentrifyZones returns a boolean if a field has been set.
func (o *ActiveDirectory) HasCentrifyZones() bool {
	if o != nil && !IsNil(o.CentrifyZones) {
		return true
	}

	return false
}

// SetCentrifyZones gets a reference to the given []CentrifyZones and assigns it to the CentrifyZones field.
func (o *ActiveDirectory) SetCentrifyZones(v []CentrifyZones) {
	o.CentrifyZones = v
}

// GetDomainControllers returns the DomainControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetDomainControllers() []DomainControllers {
	if o == nil {
		var ret []DomainControllers
		return ret
	}
	return o.DomainControllers
}

// GetDomainControllersOk returns a tuple with the DomainControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetDomainControllersOk() ([]DomainControllers, bool) {
	if o == nil || IsNil(o.DomainControllers) {
		return nil, false
	}
	return o.DomainControllers, true
}

// HasDomainControllers returns a boolean if a field has been set.
func (o *ActiveDirectory) HasDomainControllers() bool {
	if o != nil && !IsNil(o.DomainControllers) {
		return true
	}

	return false
}

// SetDomainControllers gets a reference to the given []DomainControllers and assigns it to the DomainControllers field.
func (o *ActiveDirectory) SetDomainControllers(v []DomainControllers) {
	o.DomainControllers = v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetDomainName() string {
	if o == nil || IsNil(o.DomainName.Get()) {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *ActiveDirectory) HasDomainName() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *ActiveDirectory) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *ActiveDirectory) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *ActiveDirectory) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ActiveDirectory) GetError() ActiveDirectoryError {
	if o == nil || IsNil(o.Error) {
		var ret ActiveDirectoryError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectory) GetErrorOk() (*ActiveDirectoryError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ActiveDirectory) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ActiveDirectoryError and assigns it to the Error field.
func (o *ActiveDirectory) SetError(v ActiveDirectoryError) {
	o.Error = &v
}

// GetIdMappingParams returns the IdMappingParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetIdMappingParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IdMappingParams
}

// GetIdMappingParamsOk returns a tuple with the IdMappingParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetIdMappingParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IdMappingParams) {
		return map[string]interface{}{}, false
	}
	return o.IdMappingParams, true
}

// HasIdMappingParams returns a boolean if a field has been set.
func (o *ActiveDirectory) HasIdMappingParams() bool {
	if o != nil && !IsNil(o.IdMappingParams) {
		return true
	}

	return false
}

// SetIdMappingParams gets a reference to the given map[string]interface{} and assigns it to the IdMappingParams field.
func (o *ActiveDirectory) SetIdMappingParams(v map[string]interface{}) {
	o.IdMappingParams = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetPermissions() []Tenant {
	if o == nil {
		var ret []Tenant
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetPermissionsOk() ([]Tenant, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ActiveDirectory) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Tenant and assigns it to the Permissions field.
func (o *ActiveDirectory) SetPermissions(v []Tenant) {
	o.Permissions = v
}

// GetSecurityPrincipals returns the SecurityPrincipals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetSecurityPrincipals() []SecurityPrincipal {
	if o == nil {
		var ret []SecurityPrincipal
		return ret
	}
	return o.SecurityPrincipals
}

// GetSecurityPrincipalsOk returns a tuple with the SecurityPrincipals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetSecurityPrincipalsOk() ([]SecurityPrincipal, bool) {
	if o == nil || IsNil(o.SecurityPrincipals) {
		return nil, false
	}
	return o.SecurityPrincipals, true
}

// HasSecurityPrincipals returns a boolean if a field has been set.
func (o *ActiveDirectory) HasSecurityPrincipals() bool {
	if o != nil && !IsNil(o.SecurityPrincipals) {
		return true
	}

	return false
}

// SetSecurityPrincipals gets a reference to the given []SecurityPrincipal and assigns it to the SecurityPrincipals field.
func (o *ActiveDirectory) SetSecurityPrincipals(v []SecurityPrincipal) {
	o.SecurityPrincipals = v
}

// GetTaskLogs returns the TaskLogs field value if set, zero value otherwise.
func (o *ActiveDirectory) GetTaskLogs() TaskLogs {
	if o == nil || IsNil(o.TaskLogs) {
		var ret TaskLogs
		return ret
	}
	return *o.TaskLogs
}

// GetTaskLogsOk returns a tuple with the TaskLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveDirectory) GetTaskLogsOk() (*TaskLogs, bool) {
	if o == nil || IsNil(o.TaskLogs) {
		return nil, false
	}
	return o.TaskLogs, true
}

// HasTaskLogs returns a boolean if a field has been set.
func (o *ActiveDirectory) HasTaskLogs() bool {
	if o != nil && !IsNil(o.TaskLogs) {
		return true
	}

	return false
}

// SetTaskLogs gets a reference to the given TaskLogs and assigns it to the TaskLogs field.
func (o *ActiveDirectory) SetTaskLogs(v TaskLogs) {
	o.TaskLogs = &v
}

// GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveDirectory) GetTransitiveAdTrustLevelLimit() int32 {
	if o == nil || IsNil(o.TransitiveAdTrustLevelLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.TransitiveAdTrustLevelLimit.Get()
}

// GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveDirectory) GetTransitiveAdTrustLevelLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransitiveAdTrustLevelLimit.Get(), o.TransitiveAdTrustLevelLimit.IsSet()
}

// HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.
func (o *ActiveDirectory) HasTransitiveAdTrustLevelLimit() bool {
	if o != nil && o.TransitiveAdTrustLevelLimit.IsSet() {
		return true
	}

	return false
}

// SetTransitiveAdTrustLevelLimit gets a reference to the given NullableInt32 and assigns it to the TransitiveAdTrustLevelLimit field.
func (o *ActiveDirectory) SetTransitiveAdTrustLevelLimit(v int32) {
	o.TransitiveAdTrustLevelLimit.Set(&v)
}
// SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil
func (o *ActiveDirectory) SetTransitiveAdTrustLevelLimitNil() {
	o.TransitiveAdTrustLevelLimit.Set(nil)
}

// UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil
func (o *ActiveDirectory) UnsetTransitiveAdTrustLevelLimit() {
	o.TransitiveAdTrustLevelLimit.Unset()
}

func (o ActiveDirectory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveDirectory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if !IsNil(o.DomainControllersDenyList) {
		toSerialize["domainControllersDenyList"] = o.DomainControllersDenyList
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.LdapProviderId.IsSet() {
		toSerialize["ldapProviderId"] = o.LdapProviderId.Get()
	}
	if o.MachineAccounts != nil {
		toSerialize["machineAccounts"] = o.MachineAccounts
	}
	if o.NisProviderDomainName.IsSet() {
		toSerialize["nisProviderDomainName"] = o.NisProviderDomainName.Get()
	}
	if o.OrganizationalUnitName.IsSet() {
		toSerialize["organizationalUnitName"] = o.OrganizationalUnitName.Get()
	}
	if o.PreferredDomainControllers != nil {
		toSerialize["preferredDomainControllers"] = o.PreferredDomainControllers
	}
	if o.TrustedDomainParams.IsSet() {
		toSerialize["trustedDomainParams"] = o.TrustedDomainParams.Get()
	}
	if o.WorkGroupName.IsSet() {
		toSerialize["workGroupName"] = o.WorkGroupName.Get()
	}
	if o.CentrifyZones != nil {
		toSerialize["centrifyZones"] = o.CentrifyZones
	}
	if o.DomainControllers != nil {
		toSerialize["domainControllers"] = o.DomainControllers
	}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if o.IdMappingParams != nil {
		toSerialize["idMappingParams"] = o.IdMappingParams
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.SecurityPrincipals != nil {
		toSerialize["securityPrincipals"] = o.SecurityPrincipals
	}
	if !IsNil(o.TaskLogs) {
		toSerialize["taskLogs"] = o.TaskLogs
	}
	if o.TransitiveAdTrustLevelLimit.IsSet() {
		toSerialize["transitiveAdTrustLevelLimit"] = o.TransitiveAdTrustLevelLimit.Get()
	}
	return toSerialize, nil
}

func (o *ActiveDirectory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"machineAccounts",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActiveDirectory := _ActiveDirectory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActiveDirectory)

	if err != nil {
		return err
	}

	*o = ActiveDirectory(varActiveDirectory)

	return err
}

type NullableActiveDirectory struct {
	value *ActiveDirectory
	isSet bool
}

func (v NullableActiveDirectory) Get() *ActiveDirectory {
	return v.value
}

func (v *NullableActiveDirectory) Set(val *ActiveDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveDirectory(val *ActiveDirectory) *NullableActiveDirectory {
	return &NullableActiveDirectory{value: val, isSet: true}
}

func (v NullableActiveDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


