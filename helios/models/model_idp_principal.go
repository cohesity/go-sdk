/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// IdpPrincipal Specifies the parameters of an IDP Principal.
type IdpPrincipal struct {
	// Specifies the name of the Principal.
	Name NullableString `json:"name"`
	// Specifies the IDP of the IDP with which this principal is associated.
	IdpId NullableInt64 `json:"idpId"`
	// Specifies the type of this principal. It can be a 'User' or a 'Group'.
	PrincipalType NullableString `json:"principalType"`
	// Specifies a list of roles associated with this Principal.
	Roles *[]string `json:"roles,omitempty"`
	// Specifies a list of clusters associated with this Principal. They should be in a string with the format '{cluster ID}:{cluster incarnation ID}'.
	Clusters *[]string `json:"clusters,omitempty"`
	// Specifies whether or not this principal is currently active.
	IsActive NullableBool `json:"isActive,omitempty"`
	// Specifies the timestamp in microseconds since the epoch when this Principal was created.
	CreatedTimeUsecs NullableInt64 `json:"createdTimeUsecs,omitempty"`
	// Specifies the timestamp in microseconds since the epoch when this Principal was updated.
	LastUpdatedTimeUsecs NullableInt64 `json:"lastUpdatedTimeUsecs,omitempty"`
	// Specifies the starting timestamp in microseconds since the epoch when this principal will be able to log in.
	EffectiveTimeUsecs NullableInt64 `json:"effectiveTimeUsecs,omitempty"`
	// Specifies the timestamp in microseconds since the epoch when this principal will no longer be able to log in.
	ExpiredTimeUsecs NullableInt64 `json:"expiredTimeUsecs,omitempty"`
	// Specifies the unique SID of the principal.
	Sid NullableString `json:"sid,omitempty"`
	// Specifies the list of tenant profiles associated to this principal if any.
	Profiles *[]McmTenantProfile `json:"profiles,omitempty"`
	// Specifies the list of tenant access associated to this principal.
	TenantAccesses *[]TenantAccess `json:"tenantAccesses,omitempty"`
}

// NewIdpPrincipal instantiates a new IdpPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPrincipal(name NullableString, idpId NullableInt64, principalType NullableString) *IdpPrincipal {
	this := IdpPrincipal{}
	this.Name = name
	this.IdpId = idpId
	this.PrincipalType = principalType
	return &this
}

// NewIdpPrincipalWithDefaults instantiates a new IdpPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPrincipalWithDefaults() *IdpPrincipal {
	this := IdpPrincipal{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdpPrincipal) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *IdpPrincipal) SetName(v string) {
	o.Name.Set(&v)
}

// GetIdpId returns the IdpId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *IdpPrincipal) GetIdpId() int64 {
	if o == nil || o.IdpId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.IdpId.Get()
}

// GetIdpIdOk returns a tuple with the IdpId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetIdpIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdpId.Get(), o.IdpId.IsSet()
}

// SetIdpId sets field value
func (o *IdpPrincipal) SetIdpId(v int64) {
	o.IdpId.Set(&v)
}

// GetPrincipalType returns the PrincipalType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdpPrincipal) GetPrincipalType() string {
	if o == nil || o.PrincipalType.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrincipalType.Get()
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetPrincipalTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalType.Get(), o.PrincipalType.IsSet()
}

// SetPrincipalType sets field value
func (o *IdpPrincipal) SetPrincipalType(v string) {
	o.PrincipalType.Set(&v)
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *IdpPrincipal) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPrincipal) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IdpPrincipal) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *IdpPrincipal) SetRoles(v []string) {
	o.Roles = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *IdpPrincipal) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPrincipal) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *IdpPrincipal) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *IdpPrincipal) SetClusters(v []string) {
	o.Clusters = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *IdpPrincipal) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *IdpPrincipal) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *IdpPrincipal) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *IdpPrincipal) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCreatedTimeUsecs returns the CreatedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetCreatedTimeUsecs() int64 {
	if o == nil || o.CreatedTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTimeUsecs.Get()
}

// GetCreatedTimeUsecsOk returns a tuple with the CreatedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetCreatedTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimeUsecs.Get(), o.CreatedTimeUsecs.IsSet()
}

// HasCreatedTimeUsecs returns a boolean if a field has been set.
func (o *IdpPrincipal) HasCreatedTimeUsecs() bool {
	if o != nil && o.CreatedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimeUsecs gets a reference to the given NullableInt64 and assigns it to the CreatedTimeUsecs field.
func (o *IdpPrincipal) SetCreatedTimeUsecs(v int64) {
	o.CreatedTimeUsecs.Set(&v)
}
// SetCreatedTimeUsecsNil sets the value for CreatedTimeUsecs to be an explicit nil
func (o *IdpPrincipal) SetCreatedTimeUsecsNil() {
	o.CreatedTimeUsecs.Set(nil)
}

// UnsetCreatedTimeUsecs ensures that no value is present for CreatedTimeUsecs, not even an explicit nil
func (o *IdpPrincipal) UnsetCreatedTimeUsecs() {
	o.CreatedTimeUsecs.Unset()
}

// GetLastUpdatedTimeUsecs returns the LastUpdatedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetLastUpdatedTimeUsecs() int64 {
	if o == nil || o.LastUpdatedTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimeUsecs.Get()
}

// GetLastUpdatedTimeUsecsOk returns a tuple with the LastUpdatedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetLastUpdatedTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedTimeUsecs.Get(), o.LastUpdatedTimeUsecs.IsSet()
}

// HasLastUpdatedTimeUsecs returns a boolean if a field has been set.
func (o *IdpPrincipal) HasLastUpdatedTimeUsecs() bool {
	if o != nil && o.LastUpdatedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedTimeUsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedTimeUsecs field.
func (o *IdpPrincipal) SetLastUpdatedTimeUsecs(v int64) {
	o.LastUpdatedTimeUsecs.Set(&v)
}
// SetLastUpdatedTimeUsecsNil sets the value for LastUpdatedTimeUsecs to be an explicit nil
func (o *IdpPrincipal) SetLastUpdatedTimeUsecsNil() {
	o.LastUpdatedTimeUsecs.Set(nil)
}

// UnsetLastUpdatedTimeUsecs ensures that no value is present for LastUpdatedTimeUsecs, not even an explicit nil
func (o *IdpPrincipal) UnsetLastUpdatedTimeUsecs() {
	o.LastUpdatedTimeUsecs.Unset()
}

// GetEffectiveTimeUsecs returns the EffectiveTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetEffectiveTimeUsecs() int64 {
	if o == nil || o.EffectiveTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EffectiveTimeUsecs.Get()
}

// GetEffectiveTimeUsecsOk returns a tuple with the EffectiveTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetEffectiveTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EffectiveTimeUsecs.Get(), o.EffectiveTimeUsecs.IsSet()
}

// HasEffectiveTimeUsecs returns a boolean if a field has been set.
func (o *IdpPrincipal) HasEffectiveTimeUsecs() bool {
	if o != nil && o.EffectiveTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEffectiveTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EffectiveTimeUsecs field.
func (o *IdpPrincipal) SetEffectiveTimeUsecs(v int64) {
	o.EffectiveTimeUsecs.Set(&v)
}
// SetEffectiveTimeUsecsNil sets the value for EffectiveTimeUsecs to be an explicit nil
func (o *IdpPrincipal) SetEffectiveTimeUsecsNil() {
	o.EffectiveTimeUsecs.Set(nil)
}

// UnsetEffectiveTimeUsecs ensures that no value is present for EffectiveTimeUsecs, not even an explicit nil
func (o *IdpPrincipal) UnsetEffectiveTimeUsecs() {
	o.EffectiveTimeUsecs.Unset()
}

// GetExpiredTimeUsecs returns the ExpiredTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetExpiredTimeUsecs() int64 {
	if o == nil || o.ExpiredTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpiredTimeUsecs.Get()
}

// GetExpiredTimeUsecsOk returns a tuple with the ExpiredTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetExpiredTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiredTimeUsecs.Get(), o.ExpiredTimeUsecs.IsSet()
}

// HasExpiredTimeUsecs returns a boolean if a field has been set.
func (o *IdpPrincipal) HasExpiredTimeUsecs() bool {
	if o != nil && o.ExpiredTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiredTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ExpiredTimeUsecs field.
func (o *IdpPrincipal) SetExpiredTimeUsecs(v int64) {
	o.ExpiredTimeUsecs.Set(&v)
}
// SetExpiredTimeUsecsNil sets the value for ExpiredTimeUsecs to be an explicit nil
func (o *IdpPrincipal) SetExpiredTimeUsecsNil() {
	o.ExpiredTimeUsecs.Set(nil)
}

// UnsetExpiredTimeUsecs ensures that no value is present for ExpiredTimeUsecs, not even an explicit nil
func (o *IdpPrincipal) UnsetExpiredTimeUsecs() {
	o.ExpiredTimeUsecs.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpPrincipal) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpPrincipal) GetSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *IdpPrincipal) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *IdpPrincipal) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *IdpPrincipal) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *IdpPrincipal) UnsetSid() {
	o.Sid.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *IdpPrincipal) GetProfiles() []McmTenantProfile {
	if o == nil || o.Profiles == nil {
		var ret []McmTenantProfile
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPrincipal) GetProfilesOk() (*[]McmTenantProfile, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IdpPrincipal) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []McmTenantProfile and assigns it to the Profiles field.
func (o *IdpPrincipal) SetProfiles(v []McmTenantProfile) {
	o.Profiles = &v
}

// GetTenantAccesses returns the TenantAccesses field value if set, zero value otherwise.
func (o *IdpPrincipal) GetTenantAccesses() []TenantAccess {
	if o == nil || o.TenantAccesses == nil {
		var ret []TenantAccess
		return ret
	}
	return *o.TenantAccesses
}

// GetTenantAccessesOk returns a tuple with the TenantAccesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPrincipal) GetTenantAccessesOk() (*[]TenantAccess, bool) {
	if o == nil || o.TenantAccesses == nil {
		return nil, false
	}
	return o.TenantAccesses, true
}

// HasTenantAccesses returns a boolean if a field has been set.
func (o *IdpPrincipal) HasTenantAccesses() bool {
	if o != nil && o.TenantAccesses != nil {
		return true
	}

	return false
}

// SetTenantAccesses gets a reference to the given []TenantAccess and assigns it to the TenantAccesses field.
func (o *IdpPrincipal) SetTenantAccesses(v []TenantAccess) {
	o.TenantAccesses = &v
}

func (o IdpPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["idpId"] = o.IdpId.Get()
	}
	if true {
		toSerialize["principalType"] = o.PrincipalType.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.CreatedTimeUsecs.IsSet() {
		toSerialize["createdTimeUsecs"] = o.CreatedTimeUsecs.Get()
	}
	if o.LastUpdatedTimeUsecs.IsSet() {
		toSerialize["lastUpdatedTimeUsecs"] = o.LastUpdatedTimeUsecs.Get()
	}
	if o.EffectiveTimeUsecs.IsSet() {
		toSerialize["effectiveTimeUsecs"] = o.EffectiveTimeUsecs.Get()
	}
	if o.ExpiredTimeUsecs.IsSet() {
		toSerialize["expiredTimeUsecs"] = o.ExpiredTimeUsecs.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.TenantAccesses != nil {
		toSerialize["tenantAccesses"] = o.TenantAccesses
	}
	return json.Marshal(toSerialize)
}

type NullableIdpPrincipal struct {
	value *IdpPrincipal
	isSet bool
}

func (v NullableIdpPrincipal) Get() *IdpPrincipal {
	return v.value
}

func (v *NullableIdpPrincipal) Set(val *IdpPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPrincipal(val *IdpPrincipal) *NullableIdpPrincipal {
	return &NullableIdpPrincipal{value: val, isSet: true}
}

func (v NullableIdpPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


