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

// UserDeleteParameters Specifies the users to delete on the Cohesity Cluster.
type UserDeleteParameters struct {
	// Specifies the domain associated with the users to delete. Only users associated with the same domain can be deleted by a single request. If no domain is specified, the specified users are deleted from the LOCAL domain on the Cohesity Cluster. If a non-LOCAL domain is specified, the specified users are deleted on the Cohesity Cluster. However, the referenced user principals on the Active Directory are not deleted.
	Domain NullableString `json:"domain,omitempty"`
	// Specifies the tenant for which the users are to be deleted.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Array of Users.  Specifies the list of users to delete on Cohesity Cluster. Only users from the same domain can be deleted by a single request.
	Users []string `json:"users,omitempty"`
}

// NewUserDeleteParameters instantiates a new UserDeleteParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDeleteParameters() *UserDeleteParameters {
	this := UserDeleteParameters{}
	return &this
}

// NewUserDeleteParametersWithDefaults instantiates a new UserDeleteParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDeleteParametersWithDefaults() *UserDeleteParameters {
	this := UserDeleteParameters{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDeleteParameters) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDeleteParameters) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *UserDeleteParameters) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *UserDeleteParameters) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *UserDeleteParameters) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *UserDeleteParameters) UnsetDomain() {
	o.Domain.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDeleteParameters) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDeleteParameters) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *UserDeleteParameters) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *UserDeleteParameters) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *UserDeleteParameters) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *UserDeleteParameters) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDeleteParameters) GetUsers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDeleteParameters) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserDeleteParameters) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *UserDeleteParameters) SetUsers(v []string) {
	o.Users = v
}

func (o UserDeleteParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUserDeleteParameters struct {
	value *UserDeleteParameters
	isSet bool
}

func (v NullableUserDeleteParameters) Get() *UserDeleteParameters {
	return v.value
}

func (v *NullableUserDeleteParameters) Set(val *UserDeleteParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDeleteParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDeleteParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDeleteParameters(val *UserDeleteParameters) *NullableUserDeleteParameters {
	return &NullableUserDeleteParameters{value: val, isSet: true}
}

func (v NullableUserDeleteParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDeleteParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


