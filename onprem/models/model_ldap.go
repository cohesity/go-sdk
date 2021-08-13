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

// Ldap Specifies an LDAP.
type Ldap struct {
	// Specifies the LDAP name.
	Name NullableString `json:"name"`
	// Specifies a list of preferred LDAP servers. Servers should either be FQDNs or IP addresses.
	PreferredLdapServers []string `json:"preferredLdapServers"`
	// Specifies the base distinguished name used as the base for LDAP search requests.
	BaseDistinguishedName NullableString `json:"baseDistinguishedName"`
	// Specifies the LDAP authentication type.
	AuthType string `json:"authType"`
	// Specifies the LDAP server port.
	Port NullableInt32 `json:"port,omitempty"`
	// Specifies the parameters for LDAP with 'Simple' authentication type.
	SimpleAuthParams *SimpleAuthParams `json:"simpleAuthParams,omitempty"`
	// Specifies the Active Directory id which is mapped to this LDAP.
	ActiveDirectoryId NullableInt64 `json:"activeDirectoryId,omitempty"`
	// Specifies name of the LDAP attribute used for common name of an object.
	AttributeCommonName NullableString `json:"attributeCommonName,omitempty"`
	// Specifies name of the attribute used to lookup unix GID of an LDAP user.
	AttributeGid NullableString `json:"attributeGid,omitempty"`
	// Specifies name of the attribute used to lookup unix UID of an LDAP user.
	AttributeUid NullableString `json:"attributeUid,omitempty"`
	// Specifies name of the LDAP attribute used to lookup members of a group.
	AttributeMemberOf NullableString `json:"attributeMemberOf,omitempty"`
	// Specifies name of the LDAP attribute used to lookup a user by user ID.
	AttributeUsername NullableString `json:"attributeUsername,omitempty"`
	// Specifies name of the LDAP group object class for user accounts.
	ObjectClassGroup NullableString `json:"objectClassGroup,omitempty"`
	// Specifies name of the LDAP user object class for user accounts.
	ObjectClassUser NullableString `json:"objectClassUser,omitempty"`
	// Specifies the LDAP id.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the LDAP tenant id.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewLdap instantiates a new Ldap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdap(name NullableString, preferredLdapServers []string, baseDistinguishedName NullableString, authType string) *Ldap {
	this := Ldap{}
	this.Name = name
	this.PreferredLdapServers = preferredLdapServers
	this.BaseDistinguishedName = baseDistinguishedName
	this.AuthType = authType
	return &this
}

// NewLdapWithDefaults instantiates a new Ldap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapWithDefaults() *Ldap {
	this := Ldap{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Ldap) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Ldap) SetName(v string) {
	o.Name.Set(&v)
}

// GetPreferredLdapServers returns the PreferredLdapServers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Ldap) GetPreferredLdapServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreferredLdapServers
}

// GetPreferredLdapServersOk returns a tuple with the PreferredLdapServers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetPreferredLdapServersOk() (*[]string, bool) {
	if o == nil || o.PreferredLdapServers == nil {
		return nil, false
	}
	return &o.PreferredLdapServers, true
}

// SetPreferredLdapServers sets field value
func (o *Ldap) SetPreferredLdapServers(v []string) {
	o.PreferredLdapServers = v
}

// GetBaseDistinguishedName returns the BaseDistinguishedName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Ldap) GetBaseDistinguishedName() string {
	if o == nil || o.BaseDistinguishedName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BaseDistinguishedName.Get()
}

// GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetBaseDistinguishedNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseDistinguishedName.Get(), o.BaseDistinguishedName.IsSet()
}

// SetBaseDistinguishedName sets field value
func (o *Ldap) SetBaseDistinguishedName(v string) {
	o.BaseDistinguishedName.Set(&v)
}

// GetAuthType returns the AuthType field value
func (o *Ldap) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *Ldap) GetAuthTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *Ldap) SetAuthType(v string) {
	o.AuthType = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *Ldap) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *Ldap) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *Ldap) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *Ldap) UnsetPort() {
	o.Port.Unset()
}

// GetSimpleAuthParams returns the SimpleAuthParams field value if set, zero value otherwise.
func (o *Ldap) GetSimpleAuthParams() SimpleAuthParams {
	if o == nil || o.SimpleAuthParams == nil {
		var ret SimpleAuthParams
		return ret
	}
	return *o.SimpleAuthParams
}

// GetSimpleAuthParamsOk returns a tuple with the SimpleAuthParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ldap) GetSimpleAuthParamsOk() (*SimpleAuthParams, bool) {
	if o == nil || o.SimpleAuthParams == nil {
		return nil, false
	}
	return o.SimpleAuthParams, true
}

// HasSimpleAuthParams returns a boolean if a field has been set.
func (o *Ldap) HasSimpleAuthParams() bool {
	if o != nil && o.SimpleAuthParams != nil {
		return true
	}

	return false
}

// SetSimpleAuthParams gets a reference to the given SimpleAuthParams and assigns it to the SimpleAuthParams field.
func (o *Ldap) SetSimpleAuthParams(v SimpleAuthParams) {
	o.SimpleAuthParams = &v
}

// GetActiveDirectoryId returns the ActiveDirectoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetActiveDirectoryId() int64 {
	if o == nil || o.ActiveDirectoryId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ActiveDirectoryId.Get()
}

// GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetActiveDirectoryIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveDirectoryId.Get(), o.ActiveDirectoryId.IsSet()
}

// HasActiveDirectoryId returns a boolean if a field has been set.
func (o *Ldap) HasActiveDirectoryId() bool {
	if o != nil && o.ActiveDirectoryId.IsSet() {
		return true
	}

	return false
}

// SetActiveDirectoryId gets a reference to the given NullableInt64 and assigns it to the ActiveDirectoryId field.
func (o *Ldap) SetActiveDirectoryId(v int64) {
	o.ActiveDirectoryId.Set(&v)
}
// SetActiveDirectoryIdNil sets the value for ActiveDirectoryId to be an explicit nil
func (o *Ldap) SetActiveDirectoryIdNil() {
	o.ActiveDirectoryId.Set(nil)
}

// UnsetActiveDirectoryId ensures that no value is present for ActiveDirectoryId, not even an explicit nil
func (o *Ldap) UnsetActiveDirectoryId() {
	o.ActiveDirectoryId.Unset()
}

// GetAttributeCommonName returns the AttributeCommonName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetAttributeCommonName() string {
	if o == nil || o.AttributeCommonName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributeCommonName.Get()
}

// GetAttributeCommonNameOk returns a tuple with the AttributeCommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetAttributeCommonNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttributeCommonName.Get(), o.AttributeCommonName.IsSet()
}

// HasAttributeCommonName returns a boolean if a field has been set.
func (o *Ldap) HasAttributeCommonName() bool {
	if o != nil && o.AttributeCommonName.IsSet() {
		return true
	}

	return false
}

// SetAttributeCommonName gets a reference to the given NullableString and assigns it to the AttributeCommonName field.
func (o *Ldap) SetAttributeCommonName(v string) {
	o.AttributeCommonName.Set(&v)
}
// SetAttributeCommonNameNil sets the value for AttributeCommonName to be an explicit nil
func (o *Ldap) SetAttributeCommonNameNil() {
	o.AttributeCommonName.Set(nil)
}

// UnsetAttributeCommonName ensures that no value is present for AttributeCommonName, not even an explicit nil
func (o *Ldap) UnsetAttributeCommonName() {
	o.AttributeCommonName.Unset()
}

// GetAttributeGid returns the AttributeGid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetAttributeGid() string {
	if o == nil || o.AttributeGid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributeGid.Get()
}

// GetAttributeGidOk returns a tuple with the AttributeGid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetAttributeGidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttributeGid.Get(), o.AttributeGid.IsSet()
}

// HasAttributeGid returns a boolean if a field has been set.
func (o *Ldap) HasAttributeGid() bool {
	if o != nil && o.AttributeGid.IsSet() {
		return true
	}

	return false
}

// SetAttributeGid gets a reference to the given NullableString and assigns it to the AttributeGid field.
func (o *Ldap) SetAttributeGid(v string) {
	o.AttributeGid.Set(&v)
}
// SetAttributeGidNil sets the value for AttributeGid to be an explicit nil
func (o *Ldap) SetAttributeGidNil() {
	o.AttributeGid.Set(nil)
}

// UnsetAttributeGid ensures that no value is present for AttributeGid, not even an explicit nil
func (o *Ldap) UnsetAttributeGid() {
	o.AttributeGid.Unset()
}

// GetAttributeUid returns the AttributeUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetAttributeUid() string {
	if o == nil || o.AttributeUid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributeUid.Get()
}

// GetAttributeUidOk returns a tuple with the AttributeUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetAttributeUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttributeUid.Get(), o.AttributeUid.IsSet()
}

// HasAttributeUid returns a boolean if a field has been set.
func (o *Ldap) HasAttributeUid() bool {
	if o != nil && o.AttributeUid.IsSet() {
		return true
	}

	return false
}

// SetAttributeUid gets a reference to the given NullableString and assigns it to the AttributeUid field.
func (o *Ldap) SetAttributeUid(v string) {
	o.AttributeUid.Set(&v)
}
// SetAttributeUidNil sets the value for AttributeUid to be an explicit nil
func (o *Ldap) SetAttributeUidNil() {
	o.AttributeUid.Set(nil)
}

// UnsetAttributeUid ensures that no value is present for AttributeUid, not even an explicit nil
func (o *Ldap) UnsetAttributeUid() {
	o.AttributeUid.Unset()
}

// GetAttributeMemberOf returns the AttributeMemberOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetAttributeMemberOf() string {
	if o == nil || o.AttributeMemberOf.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributeMemberOf.Get()
}

// GetAttributeMemberOfOk returns a tuple with the AttributeMemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetAttributeMemberOfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttributeMemberOf.Get(), o.AttributeMemberOf.IsSet()
}

// HasAttributeMemberOf returns a boolean if a field has been set.
func (o *Ldap) HasAttributeMemberOf() bool {
	if o != nil && o.AttributeMemberOf.IsSet() {
		return true
	}

	return false
}

// SetAttributeMemberOf gets a reference to the given NullableString and assigns it to the AttributeMemberOf field.
func (o *Ldap) SetAttributeMemberOf(v string) {
	o.AttributeMemberOf.Set(&v)
}
// SetAttributeMemberOfNil sets the value for AttributeMemberOf to be an explicit nil
func (o *Ldap) SetAttributeMemberOfNil() {
	o.AttributeMemberOf.Set(nil)
}

// UnsetAttributeMemberOf ensures that no value is present for AttributeMemberOf, not even an explicit nil
func (o *Ldap) UnsetAttributeMemberOf() {
	o.AttributeMemberOf.Unset()
}

// GetAttributeUsername returns the AttributeUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetAttributeUsername() string {
	if o == nil || o.AttributeUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributeUsername.Get()
}

// GetAttributeUsernameOk returns a tuple with the AttributeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetAttributeUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AttributeUsername.Get(), o.AttributeUsername.IsSet()
}

// HasAttributeUsername returns a boolean if a field has been set.
func (o *Ldap) HasAttributeUsername() bool {
	if o != nil && o.AttributeUsername.IsSet() {
		return true
	}

	return false
}

// SetAttributeUsername gets a reference to the given NullableString and assigns it to the AttributeUsername field.
func (o *Ldap) SetAttributeUsername(v string) {
	o.AttributeUsername.Set(&v)
}
// SetAttributeUsernameNil sets the value for AttributeUsername to be an explicit nil
func (o *Ldap) SetAttributeUsernameNil() {
	o.AttributeUsername.Set(nil)
}

// UnsetAttributeUsername ensures that no value is present for AttributeUsername, not even an explicit nil
func (o *Ldap) UnsetAttributeUsername() {
	o.AttributeUsername.Unset()
}

// GetObjectClassGroup returns the ObjectClassGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetObjectClassGroup() string {
	if o == nil || o.ObjectClassGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectClassGroup.Get()
}

// GetObjectClassGroupOk returns a tuple with the ObjectClassGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetObjectClassGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectClassGroup.Get(), o.ObjectClassGroup.IsSet()
}

// HasObjectClassGroup returns a boolean if a field has been set.
func (o *Ldap) HasObjectClassGroup() bool {
	if o != nil && o.ObjectClassGroup.IsSet() {
		return true
	}

	return false
}

// SetObjectClassGroup gets a reference to the given NullableString and assigns it to the ObjectClassGroup field.
func (o *Ldap) SetObjectClassGroup(v string) {
	o.ObjectClassGroup.Set(&v)
}
// SetObjectClassGroupNil sets the value for ObjectClassGroup to be an explicit nil
func (o *Ldap) SetObjectClassGroupNil() {
	o.ObjectClassGroup.Set(nil)
}

// UnsetObjectClassGroup ensures that no value is present for ObjectClassGroup, not even an explicit nil
func (o *Ldap) UnsetObjectClassGroup() {
	o.ObjectClassGroup.Unset()
}

// GetObjectClassUser returns the ObjectClassUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetObjectClassUser() string {
	if o == nil || o.ObjectClassUser.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectClassUser.Get()
}

// GetObjectClassUserOk returns a tuple with the ObjectClassUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetObjectClassUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectClassUser.Get(), o.ObjectClassUser.IsSet()
}

// HasObjectClassUser returns a boolean if a field has been set.
func (o *Ldap) HasObjectClassUser() bool {
	if o != nil && o.ObjectClassUser.IsSet() {
		return true
	}

	return false
}

// SetObjectClassUser gets a reference to the given NullableString and assigns it to the ObjectClassUser field.
func (o *Ldap) SetObjectClassUser(v string) {
	o.ObjectClassUser.Set(&v)
}
// SetObjectClassUserNil sets the value for ObjectClassUser to be an explicit nil
func (o *Ldap) SetObjectClassUserNil() {
	o.ObjectClassUser.Set(nil)
}

// UnsetObjectClassUser ensures that no value is present for ObjectClassUser, not even an explicit nil
func (o *Ldap) UnsetObjectClassUser() {
	o.ObjectClassUser.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Ldap) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Ldap) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Ldap) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Ldap) UnsetId() {
	o.Id.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ldap) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ldap) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Ldap) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Ldap) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Ldap) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Ldap) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o Ldap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PreferredLdapServers != nil {
		toSerialize["preferredLdapServers"] = o.PreferredLdapServers
	}
	if true {
		toSerialize["baseDistinguishedName"] = o.BaseDistinguishedName.Get()
	}
	if true {
		toSerialize["authType"] = o.AuthType
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.SimpleAuthParams != nil {
		toSerialize["simpleAuthParams"] = o.SimpleAuthParams
	}
	if o.ActiveDirectoryId.IsSet() {
		toSerialize["activeDirectoryId"] = o.ActiveDirectoryId.Get()
	}
	if o.AttributeCommonName.IsSet() {
		toSerialize["attributeCommonName"] = o.AttributeCommonName.Get()
	}
	if o.AttributeGid.IsSet() {
		toSerialize["attributeGid"] = o.AttributeGid.Get()
	}
	if o.AttributeUid.IsSet() {
		toSerialize["attributeUid"] = o.AttributeUid.Get()
	}
	if o.AttributeMemberOf.IsSet() {
		toSerialize["attributeMemberOf"] = o.AttributeMemberOf.Get()
	}
	if o.AttributeUsername.IsSet() {
		toSerialize["attributeUsername"] = o.AttributeUsername.Get()
	}
	if o.ObjectClassGroup.IsSet() {
		toSerialize["objectClassGroup"] = o.ObjectClassGroup.Get()
	}
	if o.ObjectClassUser.IsSet() {
		toSerialize["objectClassUser"] = o.ObjectClassUser.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLdap struct {
	value *Ldap
	isSet bool
}

func (v NullableLdap) Get() *Ldap {
	return v.value
}

func (v *NullableLdap) Set(val *Ldap) {
	v.value = val
	v.isSet = true
}

func (v NullableLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdap(val *Ldap) *NullableLdap {
	return &NullableLdap{value: val, isSet: true}
}

func (v NullableLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Ldap) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}