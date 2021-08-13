/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// UserParams Specifies a User.
type UserParams struct {
	// Specifies the username.
	Username NullableString `json:"username"`
	// Specifies the domain of the User.
	Domain NullableString `json:"domain"`
	// Specifies the password of the User.
	Password NullableString `json:"password"`
	// Specifies the description of the User.
	Description NullableString `json:"description,omitempty"`
	// Specifies the email address of the User.
	Email NullableString `json:"email,omitempty"`
	// Specifies the Roles of the User.
	Roles []string `json:"roles,omitempty"`
	// Specifies the primary group of the User. Primary group is used for file access.
	PrimaryGroup NullableString `json:"primaryGroup,omitempty"`
	// Specifies other groups of the User.
	OtherGroups []string `json:"otherGroups,omitempty"`
	// Specifies whether the User is restricted.
	Restricted NullableBool `json:"restricted,omitempty"`
	// Specifies the epoch time in milliseconds when the user is effective.
	EffectiveTimeMsecs NullableInt64 `json:"effectiveTimeMsecs,omitempty"`
	// Specifies the epoch time in milliseconds when the user is expired.
	ExpiredTimeMsecs NullableInt64 `json:"expiredTimeMsecs,omitempty"`
	// Specifies whether the User is locked.
	Locked NullableBool `json:"locked,omitempty"`
	// Specifies the sid of the User.
	Sid NullableString `json:"sid,omitempty"`
	// Specifies the epoch time in milliseconds when the user account was created.
	CreatedTimeMsecs NullableInt64 `json:"createdTimeMsecs,omitempty"`
	// Specifies the epoch time in milliseconds when the user account was last modified.
	LastUpdatedTimeMsecs NullableInt64 `json:"lastUpdatedTimeMsecs,omitempty"`
	// Specifies the tenant id of the User.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the S3 Account parameters of the User.
	S3AccountParams *S3AccountParams `json:"s3AccountParams,omitempty"`
	// Specifies the reason for locking the User.
	LockedReason NullableString `json:"lockedReason,omitempty"`
}

// NewUserParams instantiates a new UserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserParams(username NullableString, domain NullableString, password NullableString) *UserParams {
	this := UserParams{}
	this.Username = username
	this.Domain = domain
	this.Password = password
	return &this
}

// NewUserParamsWithDefaults instantiates a new UserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserParamsWithDefaults() *UserParams {
	this := UserParams{}
	return &this
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserParams) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}

	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// SetUsername sets field value
func (o *UserParams) SetUsername(v string) {
	o.Username.Set(&v)
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserParams) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *UserParams) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserParams) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}

	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// SetPassword sets field value
func (o *UserParams) SetPassword(v string) {
	o.Password.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UserParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UserParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UserParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UserParams) UnsetDescription() {
	o.Description.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserParams) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserParams) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserParams) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserParams) UnsetEmail() {
	o.Email.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserParams) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserParams) SetRoles(v []string) {
	o.Roles = v
}

// GetPrimaryGroup returns the PrimaryGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetPrimaryGroup() string {
	if o == nil || o.PrimaryGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryGroup.Get()
}

// GetPrimaryGroupOk returns a tuple with the PrimaryGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetPrimaryGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrimaryGroup.Get(), o.PrimaryGroup.IsSet()
}

// HasPrimaryGroup returns a boolean if a field has been set.
func (o *UserParams) HasPrimaryGroup() bool {
	if o != nil && o.PrimaryGroup.IsSet() {
		return true
	}

	return false
}

// SetPrimaryGroup gets a reference to the given NullableString and assigns it to the PrimaryGroup field.
func (o *UserParams) SetPrimaryGroup(v string) {
	o.PrimaryGroup.Set(&v)
}
// SetPrimaryGroupNil sets the value for PrimaryGroup to be an explicit nil
func (o *UserParams) SetPrimaryGroupNil() {
	o.PrimaryGroup.Set(nil)
}

// UnsetPrimaryGroup ensures that no value is present for PrimaryGroup, not even an explicit nil
func (o *UserParams) UnsetPrimaryGroup() {
	o.PrimaryGroup.Unset()
}

// GetOtherGroups returns the OtherGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetOtherGroups() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.OtherGroups
}

// GetOtherGroupsOk returns a tuple with the OtherGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetOtherGroupsOk() (*[]string, bool) {
	if o == nil || o.OtherGroups == nil {
		return nil, false
	}
	return &o.OtherGroups, true
}

// HasOtherGroups returns a boolean if a field has been set.
func (o *UserParams) HasOtherGroups() bool {
	if o != nil && o.OtherGroups != nil {
		return true
	}

	return false
}

// SetOtherGroups gets a reference to the given []string and assigns it to the OtherGroups field.
func (o *UserParams) SetOtherGroups(v []string) {
	o.OtherGroups = v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetRestricted() bool {
	if o == nil || o.Restricted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Restricted.Get()
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetRestrictedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Restricted.Get(), o.Restricted.IsSet()
}

// HasRestricted returns a boolean if a field has been set.
func (o *UserParams) HasRestricted() bool {
	if o != nil && o.Restricted.IsSet() {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given NullableBool and assigns it to the Restricted field.
func (o *UserParams) SetRestricted(v bool) {
	o.Restricted.Set(&v)
}
// SetRestrictedNil sets the value for Restricted to be an explicit nil
func (o *UserParams) SetRestrictedNil() {
	o.Restricted.Set(nil)
}

// UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
func (o *UserParams) UnsetRestricted() {
	o.Restricted.Unset()
}

// GetEffectiveTimeMsecs returns the EffectiveTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetEffectiveTimeMsecs() int64 {
	if o == nil || o.EffectiveTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EffectiveTimeMsecs.Get()
}

// GetEffectiveTimeMsecsOk returns a tuple with the EffectiveTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetEffectiveTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EffectiveTimeMsecs.Get(), o.EffectiveTimeMsecs.IsSet()
}

// HasEffectiveTimeMsecs returns a boolean if a field has been set.
func (o *UserParams) HasEffectiveTimeMsecs() bool {
	if o != nil && o.EffectiveTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetEffectiveTimeMsecs gets a reference to the given NullableInt64 and assigns it to the EffectiveTimeMsecs field.
func (o *UserParams) SetEffectiveTimeMsecs(v int64) {
	o.EffectiveTimeMsecs.Set(&v)
}
// SetEffectiveTimeMsecsNil sets the value for EffectiveTimeMsecs to be an explicit nil
func (o *UserParams) SetEffectiveTimeMsecsNil() {
	o.EffectiveTimeMsecs.Set(nil)
}

// UnsetEffectiveTimeMsecs ensures that no value is present for EffectiveTimeMsecs, not even an explicit nil
func (o *UserParams) UnsetEffectiveTimeMsecs() {
	o.EffectiveTimeMsecs.Unset()
}

// GetExpiredTimeMsecs returns the ExpiredTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetExpiredTimeMsecs() int64 {
	if o == nil || o.ExpiredTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpiredTimeMsecs.Get()
}

// GetExpiredTimeMsecsOk returns a tuple with the ExpiredTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetExpiredTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiredTimeMsecs.Get(), o.ExpiredTimeMsecs.IsSet()
}

// HasExpiredTimeMsecs returns a boolean if a field has been set.
func (o *UserParams) HasExpiredTimeMsecs() bool {
	if o != nil && o.ExpiredTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiredTimeMsecs gets a reference to the given NullableInt64 and assigns it to the ExpiredTimeMsecs field.
func (o *UserParams) SetExpiredTimeMsecs(v int64) {
	o.ExpiredTimeMsecs.Set(&v)
}
// SetExpiredTimeMsecsNil sets the value for ExpiredTimeMsecs to be an explicit nil
func (o *UserParams) SetExpiredTimeMsecsNil() {
	o.ExpiredTimeMsecs.Set(nil)
}

// UnsetExpiredTimeMsecs ensures that no value is present for ExpiredTimeMsecs, not even an explicit nil
func (o *UserParams) UnsetExpiredTimeMsecs() {
	o.ExpiredTimeMsecs.Unset()
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetLocked() bool {
	if o == nil || o.Locked.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetLockedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *UserParams) HasLocked() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *UserParams) SetLocked(v bool) {
	o.Locked.Set(&v)
}
// SetLockedNil sets the value for Locked to be an explicit nil
func (o *UserParams) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *UserParams) UnsetLocked() {
	o.Locked.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *UserParams) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *UserParams) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *UserParams) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *UserParams) UnsetSid() {
	o.Sid.Unset()
}

// GetCreatedTimeMsecs returns the CreatedTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetCreatedTimeMsecs() int64 {
	if o == nil || o.CreatedTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTimeMsecs.Get()
}

// GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetCreatedTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimeMsecs.Get(), o.CreatedTimeMsecs.IsSet()
}

// HasCreatedTimeMsecs returns a boolean if a field has been set.
func (o *UserParams) HasCreatedTimeMsecs() bool {
	if o != nil && o.CreatedTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedTimeMsecs field.
func (o *UserParams) SetCreatedTimeMsecs(v int64) {
	o.CreatedTimeMsecs.Set(&v)
}
// SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil
func (o *UserParams) SetCreatedTimeMsecsNil() {
	o.CreatedTimeMsecs.Set(nil)
}

// UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
func (o *UserParams) UnsetCreatedTimeMsecs() {
	o.CreatedTimeMsecs.Unset()
}

// GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetLastUpdatedTimeMsecs() int64 {
	if o == nil || o.LastUpdatedTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimeMsecs.Get()
}

// GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetLastUpdatedTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedTimeMsecs.Get(), o.LastUpdatedTimeMsecs.IsSet()
}

// HasLastUpdatedTimeMsecs returns a boolean if a field has been set.
func (o *UserParams) HasLastUpdatedTimeMsecs() bool {
	if o != nil && o.LastUpdatedTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedTimeMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedTimeMsecs field.
func (o *UserParams) SetLastUpdatedTimeMsecs(v int64) {
	o.LastUpdatedTimeMsecs.Set(&v)
}
// SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil
func (o *UserParams) SetLastUpdatedTimeMsecsNil() {
	o.LastUpdatedTimeMsecs.Set(nil)
}

// UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
func (o *UserParams) UnsetLastUpdatedTimeMsecs() {
	o.LastUpdatedTimeMsecs.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *UserParams) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *UserParams) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *UserParams) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *UserParams) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetS3AccountParams returns the S3AccountParams field value if set, zero value otherwise.
func (o *UserParams) GetS3AccountParams() S3AccountParams {
	if o == nil || o.S3AccountParams == nil {
		var ret S3AccountParams
		return ret
	}
	return *o.S3AccountParams
}

// GetS3AccountParamsOk returns a tuple with the S3AccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserParams) GetS3AccountParamsOk() (*S3AccountParams, bool) {
	if o == nil || o.S3AccountParams == nil {
		return nil, false
	}
	return o.S3AccountParams, true
}

// HasS3AccountParams returns a boolean if a field has been set.
func (o *UserParams) HasS3AccountParams() bool {
	if o != nil && o.S3AccountParams != nil {
		return true
	}

	return false
}

// SetS3AccountParams gets a reference to the given S3AccountParams and assigns it to the S3AccountParams field.
func (o *UserParams) SetS3AccountParams(v S3AccountParams) {
	o.S3AccountParams = &v
}

// GetLockedReason returns the LockedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserParams) GetLockedReason() string {
	if o == nil || o.LockedReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.LockedReason.Get()
}

// GetLockedReasonOk returns a tuple with the LockedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserParams) GetLockedReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LockedReason.Get(), o.LockedReason.IsSet()
}

// HasLockedReason returns a boolean if a field has been set.
func (o *UserParams) HasLockedReason() bool {
	if o != nil && o.LockedReason.IsSet() {
		return true
	}

	return false
}

// SetLockedReason gets a reference to the given NullableString and assigns it to the LockedReason field.
func (o *UserParams) SetLockedReason(v string) {
	o.LockedReason.Set(&v)
}
// SetLockedReasonNil sets the value for LockedReason to be an explicit nil
func (o *UserParams) SetLockedReasonNil() {
	o.LockedReason.Set(nil)
}

// UnsetLockedReason ensures that no value is present for LockedReason, not even an explicit nil
func (o *UserParams) UnsetLockedReason() {
	o.LockedReason.Unset()
}

func (o UserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username.Get()
	}
	if true {
		toSerialize["domain"] = o.Domain.Get()
	}
	if true {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.PrimaryGroup.IsSet() {
		toSerialize["primaryGroup"] = o.PrimaryGroup.Get()
	}
	if o.OtherGroups != nil {
		toSerialize["otherGroups"] = o.OtherGroups
	}
	if o.Restricted.IsSet() {
		toSerialize["restricted"] = o.Restricted.Get()
	}
	if o.EffectiveTimeMsecs.IsSet() {
		toSerialize["effectiveTimeMsecs"] = o.EffectiveTimeMsecs.Get()
	}
	if o.ExpiredTimeMsecs.IsSet() {
		toSerialize["expiredTimeMsecs"] = o.ExpiredTimeMsecs.Get()
	}
	if o.Locked.IsSet() {
		toSerialize["locked"] = o.Locked.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.CreatedTimeMsecs.IsSet() {
		toSerialize["createdTimeMsecs"] = o.CreatedTimeMsecs.Get()
	}
	if o.LastUpdatedTimeMsecs.IsSet() {
		toSerialize["lastUpdatedTimeMsecs"] = o.LastUpdatedTimeMsecs.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.S3AccountParams != nil {
		toSerialize["s3AccountParams"] = o.S3AccountParams
	}
	if o.LockedReason.IsSet() {
		toSerialize["lockedReason"] = o.LockedReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserParams struct {
	value *UserParams
	isSet bool
}

func (v NullableUserParams) Get() *UserParams {
	return v.value
}

func (v *NullableUserParams) Set(val *UserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserParams(val *UserParams) *NullableUserParams {
	return &NullableUserParams{value: val, isSet: true}
}

func (v NullableUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o UserParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}