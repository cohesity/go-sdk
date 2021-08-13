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

// CreateGroupParams Specifies the parameters to create a Group.
type CreateGroupParams struct {
	// Specifies the Group name.
	Name NullableString `json:"name"`
	// Specifies the Group domain.
	Domain NullableString `json:"domain"`
	// Specifies the description of the Group.
	Description NullableString `json:"description,omitempty"`
	// Specifies the Roles of the Group.
	Roles []string `json:"roles,omitempty"`
	// Specifies a list of Users who are member of this Group.
	Users []string `json:"users,omitempty"`
	// Specifies whether the Group is restricted.
	Restricted NullableBool `json:"restricted,omitempty"`
	// Specifies a list of tenant ids who can access this group.
	TenantIds []string `json:"tenantIds,omitempty"`
}

// NewCreateGroupParams instantiates a new CreateGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupParams(name NullableString, domain NullableString) *CreateGroupParams {
	this := CreateGroupParams{}
	this.Name = name
	this.Domain = domain
	return &this
}

// NewCreateGroupParamsWithDefaults instantiates a new CreateGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupParamsWithDefaults() *CreateGroupParams {
	this := CreateGroupParams{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateGroupParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CreateGroupParams) SetName(v string) {
	o.Name.Set(&v)
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateGroupParams) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *CreateGroupParams) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupParams) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGroupParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateGroupParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateGroupParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateGroupParams) UnsetDescription() {
	o.Description.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupParams) GetRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateGroupParams) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateGroupParams) SetRoles(v []string) {
	o.Roles = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupParams) GetUsers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CreateGroupParams) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *CreateGroupParams) SetUsers(v []string) {
	o.Users = v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupParams) GetRestricted() bool {
	if o == nil || o.Restricted.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Restricted.Get()
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetRestrictedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Restricted.Get(), o.Restricted.IsSet()
}

// HasRestricted returns a boolean if a field has been set.
func (o *CreateGroupParams) HasRestricted() bool {
	if o != nil && o.Restricted.IsSet() {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given NullableBool and assigns it to the Restricted field.
func (o *CreateGroupParams) SetRestricted(v bool) {
	o.Restricted.Set(&v)
}
// SetRestrictedNil sets the value for Restricted to be an explicit nil
func (o *CreateGroupParams) SetRestrictedNil() {
	o.Restricted.Set(nil)
}

// UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
func (o *CreateGroupParams) UnsetRestricted() {
	o.Restricted.Unset()
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGroupParams) GetTenantIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGroupParams) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return &o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *CreateGroupParams) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *CreateGroupParams) SetTenantIds(v []string) {
	o.TenantIds = v
}

func (o CreateGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Restricted.IsSet() {
		toSerialize["restricted"] = o.Restricted.Get()
	}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupParams struct {
	value *CreateGroupParams
	isSet bool
}

func (v NullableCreateGroupParams) Get() *CreateGroupParams {
	return v.value
}

func (v *NullableCreateGroupParams) Set(val *CreateGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupParams(val *CreateGroupParams) *NullableCreateGroupParams {
	return &NullableCreateGroupParams{value: val, isSet: true}
}

func (v NullableCreateGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CreateGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}