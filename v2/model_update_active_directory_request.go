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

// checks if the UpdateActiveDirectoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateActiveDirectoryRequest{}

// UpdateActiveDirectoryRequest Specifies the request to create an Active Directory.
type UpdateActiveDirectoryRequest struct {
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
	// Specifies the params of a user with administrative privilege of this Active Directory. This field is mandatory if machine accounts are updated.
	ActiveDirectoryAdminParams map[string]interface{} `json:"activeDirectoryAdminParams,omitempty"`
	// Specifies the params of the user id mapping info of an Active Directory.
	IdMappingParams map[string]interface{} `json:"idMappingParams,omitempty"`
	// Specifies if specified machine accounts should overwrite existing machine accounts.
	OverwriteMachineAccounts NullableBool `json:"overwriteMachineAccounts,omitempty"`
	// Specifies level of transitive Active Directory trust domains to be used.
	TransitiveAdTrustLevelLimit NullableInt32 `json:"transitiveAdTrustLevelLimit,omitempty"`
}

type _UpdateActiveDirectoryRequest UpdateActiveDirectoryRequest

// NewUpdateActiveDirectoryRequest instantiates a new UpdateActiveDirectoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateActiveDirectoryRequest(machineAccounts []MachineAccount) *UpdateActiveDirectoryRequest {
	this := UpdateActiveDirectoryRequest{}
	this.MachineAccounts = machineAccounts
	return &this
}

// NewUpdateActiveDirectoryRequestWithDefaults instantiates a new UpdateActiveDirectoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateActiveDirectoryRequestWithDefaults() *UpdateActiveDirectoryRequest {
	this := UpdateActiveDirectoryRequest{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetConnectionId() int64 {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret int64
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetConnectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt64 and assigns it to the ConnectionId field.
func (o *UpdateActiveDirectoryRequest) SetConnectionId(v int64) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetDomainControllersDenyList returns the DomainControllersDenyList field value if set, zero value otherwise.
func (o *UpdateActiveDirectoryRequest) GetDomainControllersDenyList() []*string {
	if o == nil || IsNil(o.DomainControllersDenyList) {
		var ret []*string
		return ret
	}
	return o.DomainControllersDenyList
}

// GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActiveDirectoryRequest) GetDomainControllersDenyListOk() ([]*string, bool) {
	if o == nil || IsNil(o.DomainControllersDenyList) {
		return nil, false
	}
	return o.DomainControllersDenyList, true
}

// HasDomainControllersDenyList returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasDomainControllersDenyList() bool {
	if o != nil && !IsNil(o.DomainControllersDenyList) {
		return true
	}

	return false
}

// SetDomainControllersDenyList gets a reference to the given []*string and assigns it to the DomainControllersDenyList field.
func (o *UpdateActiveDirectoryRequest) SetDomainControllersDenyList(v []*string) {
	o.DomainControllersDenyList = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *UpdateActiveDirectoryRequest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetId() {
	o.Id.Unset()
}

// GetLdapProviderId returns the LdapProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetLdapProviderId() int64 {
	if o == nil || IsNil(o.LdapProviderId.Get()) {
		var ret int64
		return ret
	}
	return *o.LdapProviderId.Get()
}

// GetLdapProviderIdOk returns a tuple with the LdapProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetLdapProviderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LdapProviderId.Get(), o.LdapProviderId.IsSet()
}

// HasLdapProviderId returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasLdapProviderId() bool {
	if o != nil && o.LdapProviderId.IsSet() {
		return true
	}

	return false
}

// SetLdapProviderId gets a reference to the given NullableInt64 and assigns it to the LdapProviderId field.
func (o *UpdateActiveDirectoryRequest) SetLdapProviderId(v int64) {
	o.LdapProviderId.Set(&v)
}
// SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetLdapProviderIdNil() {
	o.LdapProviderId.Set(nil)
}

// UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetLdapProviderId() {
	o.LdapProviderId.Unset()
}

// GetMachineAccounts returns the MachineAccounts field value
// If the value is explicit nil, the zero value for []MachineAccount will be returned
func (o *UpdateActiveDirectoryRequest) GetMachineAccounts() []MachineAccount {
	if o == nil {
		var ret []MachineAccount
		return ret
	}

	return o.MachineAccounts
}

// GetMachineAccountsOk returns a tuple with the MachineAccounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetMachineAccountsOk() ([]MachineAccount, bool) {
	if o == nil || IsNil(o.MachineAccounts) {
		return nil, false
	}
	return o.MachineAccounts, true
}

// SetMachineAccounts sets field value
func (o *UpdateActiveDirectoryRequest) SetMachineAccounts(v []MachineAccount) {
	o.MachineAccounts = v
}

// GetNisProviderDomainName returns the NisProviderDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetNisProviderDomainName() string {
	if o == nil || IsNil(o.NisProviderDomainName.Get()) {
		var ret string
		return ret
	}
	return *o.NisProviderDomainName.Get()
}

// GetNisProviderDomainNameOk returns a tuple with the NisProviderDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetNisProviderDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NisProviderDomainName.Get(), o.NisProviderDomainName.IsSet()
}

// HasNisProviderDomainName returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasNisProviderDomainName() bool {
	if o != nil && o.NisProviderDomainName.IsSet() {
		return true
	}

	return false
}

// SetNisProviderDomainName gets a reference to the given NullableString and assigns it to the NisProviderDomainName field.
func (o *UpdateActiveDirectoryRequest) SetNisProviderDomainName(v string) {
	o.NisProviderDomainName.Set(&v)
}
// SetNisProviderDomainNameNil sets the value for NisProviderDomainName to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetNisProviderDomainNameNil() {
	o.NisProviderDomainName.Set(nil)
}

// UnsetNisProviderDomainName ensures that no value is present for NisProviderDomainName, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetNisProviderDomainName() {
	o.NisProviderDomainName.Unset()
}

// GetOrganizationalUnitName returns the OrganizationalUnitName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetOrganizationalUnitName() string {
	if o == nil || IsNil(o.OrganizationalUnitName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationalUnitName.Get()
}

// GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetOrganizationalUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationalUnitName.Get(), o.OrganizationalUnitName.IsSet()
}

// HasOrganizationalUnitName returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasOrganizationalUnitName() bool {
	if o != nil && o.OrganizationalUnitName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationalUnitName gets a reference to the given NullableString and assigns it to the OrganizationalUnitName field.
func (o *UpdateActiveDirectoryRequest) SetOrganizationalUnitName(v string) {
	o.OrganizationalUnitName.Set(&v)
}
// SetOrganizationalUnitNameNil sets the value for OrganizationalUnitName to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetOrganizationalUnitNameNil() {
	o.OrganizationalUnitName.Set(nil)
}

// UnsetOrganizationalUnitName ensures that no value is present for OrganizationalUnitName, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetOrganizationalUnitName() {
	o.OrganizationalUnitName.Unset()
}

// GetPreferredDomainControllers returns the PreferredDomainControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetPreferredDomainControllers() []DomainController {
	if o == nil {
		var ret []DomainController
		return ret
	}
	return o.PreferredDomainControllers
}

// GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetPreferredDomainControllersOk() ([]DomainController, bool) {
	if o == nil || IsNil(o.PreferredDomainControllers) {
		return nil, false
	}
	return o.PreferredDomainControllers, true
}

// HasPreferredDomainControllers returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasPreferredDomainControllers() bool {
	if o != nil && !IsNil(o.PreferredDomainControllers) {
		return true
	}

	return false
}

// SetPreferredDomainControllers gets a reference to the given []DomainController and assigns it to the PreferredDomainControllers field.
func (o *UpdateActiveDirectoryRequest) SetPreferredDomainControllers(v []DomainController) {
	o.PreferredDomainControllers = v
}

// GetTrustedDomainParams returns the TrustedDomainParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetTrustedDomainParams() CommonActiveDirectoryParamsTrustedDomainParams {
	if o == nil || IsNil(o.TrustedDomainParams.Get()) {
		var ret CommonActiveDirectoryParamsTrustedDomainParams
		return ret
	}
	return *o.TrustedDomainParams.Get()
}

// GetTrustedDomainParamsOk returns a tuple with the TrustedDomainParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetTrustedDomainParamsOk() (*CommonActiveDirectoryParamsTrustedDomainParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustedDomainParams.Get(), o.TrustedDomainParams.IsSet()
}

// HasTrustedDomainParams returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasTrustedDomainParams() bool {
	if o != nil && o.TrustedDomainParams.IsSet() {
		return true
	}

	return false
}

// SetTrustedDomainParams gets a reference to the given NullableCommonActiveDirectoryParamsTrustedDomainParams and assigns it to the TrustedDomainParams field.
func (o *UpdateActiveDirectoryRequest) SetTrustedDomainParams(v CommonActiveDirectoryParamsTrustedDomainParams) {
	o.TrustedDomainParams.Set(&v)
}
// SetTrustedDomainParamsNil sets the value for TrustedDomainParams to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetTrustedDomainParamsNil() {
	o.TrustedDomainParams.Set(nil)
}

// UnsetTrustedDomainParams ensures that no value is present for TrustedDomainParams, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetTrustedDomainParams() {
	o.TrustedDomainParams.Unset()
}

// GetWorkGroupName returns the WorkGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetWorkGroupName() string {
	if o == nil || IsNil(o.WorkGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.WorkGroupName.Get()
}

// GetWorkGroupNameOk returns a tuple with the WorkGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetWorkGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkGroupName.Get(), o.WorkGroupName.IsSet()
}

// HasWorkGroupName returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasWorkGroupName() bool {
	if o != nil && o.WorkGroupName.IsSet() {
		return true
	}

	return false
}

// SetWorkGroupName gets a reference to the given NullableString and assigns it to the WorkGroupName field.
func (o *UpdateActiveDirectoryRequest) SetWorkGroupName(v string) {
	o.WorkGroupName.Set(&v)
}
// SetWorkGroupNameNil sets the value for WorkGroupName to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetWorkGroupNameNil() {
	o.WorkGroupName.Set(nil)
}

// UnsetWorkGroupName ensures that no value is present for WorkGroupName, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetWorkGroupName() {
	o.WorkGroupName.Unset()
}

// GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetActiveDirectoryAdminParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ActiveDirectoryAdminParams
}

// GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetActiveDirectoryAdminParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ActiveDirectoryAdminParams) {
		return map[string]interface{}{}, false
	}
	return o.ActiveDirectoryAdminParams, true
}

// HasActiveDirectoryAdminParams returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasActiveDirectoryAdminParams() bool {
	if o != nil && !IsNil(o.ActiveDirectoryAdminParams) {
		return true
	}

	return false
}

// SetActiveDirectoryAdminParams gets a reference to the given map[string]interface{} and assigns it to the ActiveDirectoryAdminParams field.
func (o *UpdateActiveDirectoryRequest) SetActiveDirectoryAdminParams(v map[string]interface{}) {
	o.ActiveDirectoryAdminParams = v
}

// GetIdMappingParams returns the IdMappingParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetIdMappingParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IdMappingParams
}

// GetIdMappingParamsOk returns a tuple with the IdMappingParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetIdMappingParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IdMappingParams) {
		return map[string]interface{}{}, false
	}
	return o.IdMappingParams, true
}

// HasIdMappingParams returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasIdMappingParams() bool {
	if o != nil && !IsNil(o.IdMappingParams) {
		return true
	}

	return false
}

// SetIdMappingParams gets a reference to the given map[string]interface{} and assigns it to the IdMappingParams field.
func (o *UpdateActiveDirectoryRequest) SetIdMappingParams(v map[string]interface{}) {
	o.IdMappingParams = v
}

// GetOverwriteMachineAccounts returns the OverwriteMachineAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetOverwriteMachineAccounts() bool {
	if o == nil || IsNil(o.OverwriteMachineAccounts.Get()) {
		var ret bool
		return ret
	}
	return *o.OverwriteMachineAccounts.Get()
}

// GetOverwriteMachineAccountsOk returns a tuple with the OverwriteMachineAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetOverwriteMachineAccountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwriteMachineAccounts.Get(), o.OverwriteMachineAccounts.IsSet()
}

// HasOverwriteMachineAccounts returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasOverwriteMachineAccounts() bool {
	if o != nil && o.OverwriteMachineAccounts.IsSet() {
		return true
	}

	return false
}

// SetOverwriteMachineAccounts gets a reference to the given NullableBool and assigns it to the OverwriteMachineAccounts field.
func (o *UpdateActiveDirectoryRequest) SetOverwriteMachineAccounts(v bool) {
	o.OverwriteMachineAccounts.Set(&v)
}
// SetOverwriteMachineAccountsNil sets the value for OverwriteMachineAccounts to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetOverwriteMachineAccountsNil() {
	o.OverwriteMachineAccounts.Set(nil)
}

// UnsetOverwriteMachineAccounts ensures that no value is present for OverwriteMachineAccounts, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetOverwriteMachineAccounts() {
	o.OverwriteMachineAccounts.Unset()
}

// GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActiveDirectoryRequest) GetTransitiveAdTrustLevelLimit() int32 {
	if o == nil || IsNil(o.TransitiveAdTrustLevelLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.TransitiveAdTrustLevelLimit.Get()
}

// GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActiveDirectoryRequest) GetTransitiveAdTrustLevelLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransitiveAdTrustLevelLimit.Get(), o.TransitiveAdTrustLevelLimit.IsSet()
}

// HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.
func (o *UpdateActiveDirectoryRequest) HasTransitiveAdTrustLevelLimit() bool {
	if o != nil && o.TransitiveAdTrustLevelLimit.IsSet() {
		return true
	}

	return false
}

// SetTransitiveAdTrustLevelLimit gets a reference to the given NullableInt32 and assigns it to the TransitiveAdTrustLevelLimit field.
func (o *UpdateActiveDirectoryRequest) SetTransitiveAdTrustLevelLimit(v int32) {
	o.TransitiveAdTrustLevelLimit.Set(&v)
}
// SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil
func (o *UpdateActiveDirectoryRequest) SetTransitiveAdTrustLevelLimitNil() {
	o.TransitiveAdTrustLevelLimit.Set(nil)
}

// UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil
func (o *UpdateActiveDirectoryRequest) UnsetTransitiveAdTrustLevelLimit() {
	o.TransitiveAdTrustLevelLimit.Unset()
}

func (o UpdateActiveDirectoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateActiveDirectoryRequest) ToMap() (map[string]interface{}, error) {
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
	if o.ActiveDirectoryAdminParams != nil {
		toSerialize["activeDirectoryAdminParams"] = o.ActiveDirectoryAdminParams
	}
	if o.IdMappingParams != nil {
		toSerialize["idMappingParams"] = o.IdMappingParams
	}
	if o.OverwriteMachineAccounts.IsSet() {
		toSerialize["overwriteMachineAccounts"] = o.OverwriteMachineAccounts.Get()
	}
	if o.TransitiveAdTrustLevelLimit.IsSet() {
		toSerialize["transitiveAdTrustLevelLimit"] = o.TransitiveAdTrustLevelLimit.Get()
	}
	return toSerialize, nil
}

func (o *UpdateActiveDirectoryRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateActiveDirectoryRequest := _UpdateActiveDirectoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateActiveDirectoryRequest)

	if err != nil {
		return err
	}

	*o = UpdateActiveDirectoryRequest(varUpdateActiveDirectoryRequest)

	return err
}

type NullableUpdateActiveDirectoryRequest struct {
	value *UpdateActiveDirectoryRequest
	isSet bool
}

func (v NullableUpdateActiveDirectoryRequest) Get() *UpdateActiveDirectoryRequest {
	return v.value
}

func (v *NullableUpdateActiveDirectoryRequest) Set(val *UpdateActiveDirectoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateActiveDirectoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateActiveDirectoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateActiveDirectoryRequest(val *UpdateActiveDirectoryRequest) *NullableUpdateActiveDirectoryRequest {
	return &NullableUpdateActiveDirectoryRequest{value: val, isSet: true}
}

func (v NullableUpdateActiveDirectoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateActiveDirectoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


