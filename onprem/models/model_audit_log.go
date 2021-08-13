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

// AuditLog Specifies an audit log message.
type AuditLog struct {
	// Specifies the change details of this audit log.
	Details NullableString `json:"details,omitempty"`
	// Specifies the username who made this audit log.
	Username NullableString `json:"username,omitempty"`
	// Specifies the domain of user who made this audit log.
	Domain NullableString `json:"domain,omitempty"`
	// Specifies the entity name.
	EntityName NullableString `json:"entityName,omitempty"`
	// Specifies the entity type.
	EntityType NullableString `json:"entityType,omitempty"`
	// Specifies the action type of this audit log.
	Action NullableString `json:"action,omitempty"`
	// Specifies a unix timestamp in micro seconds when the audit log was taken.
	TimestampUsecs NullableInt64 `json:"timestampUsecs,omitempty"`
	// Specifies the ip of user who made this audit log.
	Ip NullableString `json:"ip,omitempty"`
	// Specifies if the action is made through impersonation.
	IsImpersonation NullableBool `json:"isImpersonation,omitempty"`
	// Specifies the tenant id who made this audit log.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the tenant name who made this audit log.
	TenantName NullableString `json:"tenantName,omitempty"`
	// Specifies the original tenant id who made this audit log.
	OriginalTenantId NullableString `json:"originalTenantId,omitempty"`
	// Specifies the original tenant name who made this audit log.
	OriginalTenantName NullableString `json:"originalTenantName,omitempty"`
	// Specifies the record before the action is invoked. This will be returned only if verbose audit is enabled. 
	PreviousRecord NullableString `json:"previousRecord,omitempty"`
	// Specifies the record after the action is invoked. This will be returned only if verbose audit is enabled. 
	NewRecord NullableString `json:"newRecord,omitempty"`
}

// NewAuditLog instantiates a new AuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLog() *AuditLog {
	this := AuditLog{}
	return &this
}

// NewAuditLogWithDefaults instantiates a new AuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogWithDefaults() *AuditLog {
	this := AuditLog{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetDetails() string {
	if o == nil || o.Details.Get() == nil {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetDetailsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *AuditLog) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *AuditLog) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *AuditLog) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *AuditLog) UnsetDetails() {
	o.Details.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *AuditLog) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *AuditLog) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *AuditLog) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *AuditLog) UnsetUsername() {
	o.Username.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *AuditLog) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *AuditLog) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *AuditLog) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *AuditLog) UnsetDomain() {
	o.Domain.Unset()
}

// GetEntityName returns the EntityName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetEntityName() string {
	if o == nil || o.EntityName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityName.Get()
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetEntityNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityName.Get(), o.EntityName.IsSet()
}

// HasEntityName returns a boolean if a field has been set.
func (o *AuditLog) HasEntityName() bool {
	if o != nil && o.EntityName.IsSet() {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given NullableString and assigns it to the EntityName field.
func (o *AuditLog) SetEntityName(v string) {
	o.EntityName.Set(&v)
}
// SetEntityNameNil sets the value for EntityName to be an explicit nil
func (o *AuditLog) SetEntityNameNil() {
	o.EntityName.Set(nil)
}

// UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
func (o *AuditLog) UnsetEntityName() {
	o.EntityName.Unset()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetEntityType() string {
	if o == nil || o.EntityType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetEntityTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *AuditLog) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableString and assigns it to the EntityType field.
func (o *AuditLog) SetEntityType(v string) {
	o.EntityType.Set(&v)
}
// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *AuditLog) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *AuditLog) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *AuditLog) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *AuditLog) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *AuditLog) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *AuditLog) UnsetAction() {
	o.Action.Unset()
}

// GetTimestampUsecs returns the TimestampUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetTimestampUsecs() int64 {
	if o == nil || o.TimestampUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TimestampUsecs.Get()
}

// GetTimestampUsecsOk returns a tuple with the TimestampUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetTimestampUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimestampUsecs.Get(), o.TimestampUsecs.IsSet()
}

// HasTimestampUsecs returns a boolean if a field has been set.
func (o *AuditLog) HasTimestampUsecs() bool {
	if o != nil && o.TimestampUsecs.IsSet() {
		return true
	}

	return false
}

// SetTimestampUsecs gets a reference to the given NullableInt64 and assigns it to the TimestampUsecs field.
func (o *AuditLog) SetTimestampUsecs(v int64) {
	o.TimestampUsecs.Set(&v)
}
// SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil
func (o *AuditLog) SetTimestampUsecsNil() {
	o.TimestampUsecs.Set(nil)
}

// UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
func (o *AuditLog) UnsetTimestampUsecs() {
	o.TimestampUsecs.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *AuditLog) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *AuditLog) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *AuditLog) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *AuditLog) UnsetIp() {
	o.Ip.Unset()
}

// GetIsImpersonation returns the IsImpersonation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetIsImpersonation() bool {
	if o == nil || o.IsImpersonation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsImpersonation.Get()
}

// GetIsImpersonationOk returns a tuple with the IsImpersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetIsImpersonationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsImpersonation.Get(), o.IsImpersonation.IsSet()
}

// HasIsImpersonation returns a boolean if a field has been set.
func (o *AuditLog) HasIsImpersonation() bool {
	if o != nil && o.IsImpersonation.IsSet() {
		return true
	}

	return false
}

// SetIsImpersonation gets a reference to the given NullableBool and assigns it to the IsImpersonation field.
func (o *AuditLog) SetIsImpersonation(v bool) {
	o.IsImpersonation.Set(&v)
}
// SetIsImpersonationNil sets the value for IsImpersonation to be an explicit nil
func (o *AuditLog) SetIsImpersonationNil() {
	o.IsImpersonation.Set(nil)
}

// UnsetIsImpersonation ensures that no value is present for IsImpersonation, not even an explicit nil
func (o *AuditLog) UnsetIsImpersonation() {
	o.IsImpersonation.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *AuditLog) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *AuditLog) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *AuditLog) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *AuditLog) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTenantName returns the TenantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetTenantName() string {
	if o == nil || o.TenantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantName.Get()
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetTenantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantName.Get(), o.TenantName.IsSet()
}

// HasTenantName returns a boolean if a field has been set.
func (o *AuditLog) HasTenantName() bool {
	if o != nil && o.TenantName.IsSet() {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given NullableString and assigns it to the TenantName field.
func (o *AuditLog) SetTenantName(v string) {
	o.TenantName.Set(&v)
}
// SetTenantNameNil sets the value for TenantName to be an explicit nil
func (o *AuditLog) SetTenantNameNil() {
	o.TenantName.Set(nil)
}

// UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
func (o *AuditLog) UnsetTenantName() {
	o.TenantName.Unset()
}

// GetOriginalTenantId returns the OriginalTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetOriginalTenantId() string {
	if o == nil || o.OriginalTenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginalTenantId.Get()
}

// GetOriginalTenantIdOk returns a tuple with the OriginalTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetOriginalTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalTenantId.Get(), o.OriginalTenantId.IsSet()
}

// HasOriginalTenantId returns a boolean if a field has been set.
func (o *AuditLog) HasOriginalTenantId() bool {
	if o != nil && o.OriginalTenantId.IsSet() {
		return true
	}

	return false
}

// SetOriginalTenantId gets a reference to the given NullableString and assigns it to the OriginalTenantId field.
func (o *AuditLog) SetOriginalTenantId(v string) {
	o.OriginalTenantId.Set(&v)
}
// SetOriginalTenantIdNil sets the value for OriginalTenantId to be an explicit nil
func (o *AuditLog) SetOriginalTenantIdNil() {
	o.OriginalTenantId.Set(nil)
}

// UnsetOriginalTenantId ensures that no value is present for OriginalTenantId, not even an explicit nil
func (o *AuditLog) UnsetOriginalTenantId() {
	o.OriginalTenantId.Unset()
}

// GetOriginalTenantName returns the OriginalTenantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetOriginalTenantName() string {
	if o == nil || o.OriginalTenantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginalTenantName.Get()
}

// GetOriginalTenantNameOk returns a tuple with the OriginalTenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetOriginalTenantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalTenantName.Get(), o.OriginalTenantName.IsSet()
}

// HasOriginalTenantName returns a boolean if a field has been set.
func (o *AuditLog) HasOriginalTenantName() bool {
	if o != nil && o.OriginalTenantName.IsSet() {
		return true
	}

	return false
}

// SetOriginalTenantName gets a reference to the given NullableString and assigns it to the OriginalTenantName field.
func (o *AuditLog) SetOriginalTenantName(v string) {
	o.OriginalTenantName.Set(&v)
}
// SetOriginalTenantNameNil sets the value for OriginalTenantName to be an explicit nil
func (o *AuditLog) SetOriginalTenantNameNil() {
	o.OriginalTenantName.Set(nil)
}

// UnsetOriginalTenantName ensures that no value is present for OriginalTenantName, not even an explicit nil
func (o *AuditLog) UnsetOriginalTenantName() {
	o.OriginalTenantName.Unset()
}

// GetPreviousRecord returns the PreviousRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetPreviousRecord() string {
	if o == nil || o.PreviousRecord.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousRecord.Get()
}

// GetPreviousRecordOk returns a tuple with the PreviousRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetPreviousRecordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousRecord.Get(), o.PreviousRecord.IsSet()
}

// HasPreviousRecord returns a boolean if a field has been set.
func (o *AuditLog) HasPreviousRecord() bool {
	if o != nil && o.PreviousRecord.IsSet() {
		return true
	}

	return false
}

// SetPreviousRecord gets a reference to the given NullableString and assigns it to the PreviousRecord field.
func (o *AuditLog) SetPreviousRecord(v string) {
	o.PreviousRecord.Set(&v)
}
// SetPreviousRecordNil sets the value for PreviousRecord to be an explicit nil
func (o *AuditLog) SetPreviousRecordNil() {
	o.PreviousRecord.Set(nil)
}

// UnsetPreviousRecord ensures that no value is present for PreviousRecord, not even an explicit nil
func (o *AuditLog) UnsetPreviousRecord() {
	o.PreviousRecord.Unset()
}

// GetNewRecord returns the NewRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLog) GetNewRecord() string {
	if o == nil || o.NewRecord.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewRecord.Get()
}

// GetNewRecordOk returns a tuple with the NewRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLog) GetNewRecordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewRecord.Get(), o.NewRecord.IsSet()
}

// HasNewRecord returns a boolean if a field has been set.
func (o *AuditLog) HasNewRecord() bool {
	if o != nil && o.NewRecord.IsSet() {
		return true
	}

	return false
}

// SetNewRecord gets a reference to the given NullableString and assigns it to the NewRecord field.
func (o *AuditLog) SetNewRecord(v string) {
	o.NewRecord.Set(&v)
}
// SetNewRecordNil sets the value for NewRecord to be an explicit nil
func (o *AuditLog) SetNewRecordNil() {
	o.NewRecord.Set(nil)
}

// UnsetNewRecord ensures that no value is present for NewRecord, not even an explicit nil
func (o *AuditLog) UnsetNewRecord() {
	o.NewRecord.Unset()
}

func (o AuditLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.EntityName.IsSet() {
		toSerialize["entityName"] = o.EntityName.Get()
	}
	if o.EntityType.IsSet() {
		toSerialize["entityType"] = o.EntityType.Get()
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.TimestampUsecs.IsSet() {
		toSerialize["timestampUsecs"] = o.TimestampUsecs.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.IsImpersonation.IsSet() {
		toSerialize["isImpersonation"] = o.IsImpersonation.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.TenantName.IsSet() {
		toSerialize["tenantName"] = o.TenantName.Get()
	}
	if o.OriginalTenantId.IsSet() {
		toSerialize["originalTenantId"] = o.OriginalTenantId.Get()
	}
	if o.OriginalTenantName.IsSet() {
		toSerialize["originalTenantName"] = o.OriginalTenantName.Get()
	}
	if o.PreviousRecord.IsSet() {
		toSerialize["previousRecord"] = o.PreviousRecord.Get()
	}
	if o.NewRecord.IsSet() {
		toSerialize["newRecord"] = o.NewRecord.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLog struct {
	value *AuditLog
	isSet bool
}

func (v NullableAuditLog) Get() *AuditLog {
	return v.value
}

func (v *NullableAuditLog) Set(val *AuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLog(val *AuditLog) *NullableAuditLog {
	return &NullableAuditLog{value: val, isSet: true}
}

func (v NullableAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AuditLog) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}