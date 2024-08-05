/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the TenantAssignmentsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAssignmentsResult{}

// TenantAssignmentsResult Result of attempted tenant assignments
type TenantAssignmentsResult struct {
	// A list of Ids of properties assigned to the tenant.
	Objects []int64 `json:"objects,omitempty"`
	// A list of Ids of properties assigned to the tenant.
	Policies []string `json:"policies,omitempty"`
	// A list of Ids of properties assigned to the tenant.
	StorageDomains []int64 `json:"storageDomains,omitempty"`
	// A list of Ids of properties assigned to the tenant.
	Views []int64 `json:"views,omitempty"`
	// A list of Ids of properties assigned to the tenant.
	Vlans []string `json:"vlans,omitempty"`
}

// NewTenantAssignmentsResult instantiates a new TenantAssignmentsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAssignmentsResult() *TenantAssignmentsResult {
	this := TenantAssignmentsResult{}
	return &this
}

// NewTenantAssignmentsResultWithDefaults instantiates a new TenantAssignmentsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAssignmentsResultWithDefaults() *TenantAssignmentsResult {
	this := TenantAssignmentsResult{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *TenantAssignmentsResult) GetObjects() []int64 {
	if o == nil || IsNil(o.Objects) {
		var ret []int64
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAssignmentsResult) GetObjectsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *TenantAssignmentsResult) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []int64 and assigns it to the Objects field.
func (o *TenantAssignmentsResult) SetObjects(v []int64) {
	o.Objects = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *TenantAssignmentsResult) GetPolicies() []string {
	if o == nil || IsNil(o.Policies) {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAssignmentsResult) GetPoliciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *TenantAssignmentsResult) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *TenantAssignmentsResult) SetPolicies(v []string) {
	o.Policies = v
}

// GetStorageDomains returns the StorageDomains field value if set, zero value otherwise.
func (o *TenantAssignmentsResult) GetStorageDomains() []int64 {
	if o == nil || IsNil(o.StorageDomains) {
		var ret []int64
		return ret
	}
	return o.StorageDomains
}

// GetStorageDomainsOk returns a tuple with the StorageDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAssignmentsResult) GetStorageDomainsOk() ([]int64, bool) {
	if o == nil || IsNil(o.StorageDomains) {
		return nil, false
	}
	return o.StorageDomains, true
}

// HasStorageDomains returns a boolean if a field has been set.
func (o *TenantAssignmentsResult) HasStorageDomains() bool {
	if o != nil && !IsNil(o.StorageDomains) {
		return true
	}

	return false
}

// SetStorageDomains gets a reference to the given []int64 and assigns it to the StorageDomains field.
func (o *TenantAssignmentsResult) SetStorageDomains(v []int64) {
	o.StorageDomains = v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *TenantAssignmentsResult) GetViews() []int64 {
	if o == nil || IsNil(o.Views) {
		var ret []int64
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAssignmentsResult) GetViewsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *TenantAssignmentsResult) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []int64 and assigns it to the Views field.
func (o *TenantAssignmentsResult) SetViews(v []int64) {
	o.Views = v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *TenantAssignmentsResult) GetVlans() []string {
	if o == nil || IsNil(o.Vlans) {
		var ret []string
		return ret
	}
	return o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAssignmentsResult) GetVlansOk() ([]string, bool) {
	if o == nil || IsNil(o.Vlans) {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *TenantAssignmentsResult) HasVlans() bool {
	if o != nil && !IsNil(o.Vlans) {
		return true
	}

	return false
}

// SetVlans gets a reference to the given []string and assigns it to the Vlans field.
func (o *TenantAssignmentsResult) SetVlans(v []string) {
	o.Vlans = v
}

func (o TenantAssignmentsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAssignmentsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.StorageDomains) {
		toSerialize["storageDomains"] = o.StorageDomains
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !IsNil(o.Vlans) {
		toSerialize["vlans"] = o.Vlans
	}
	return toSerialize, nil
}

type NullableTenantAssignmentsResult struct {
	value *TenantAssignmentsResult
	isSet bool
}

func (v NullableTenantAssignmentsResult) Get() *TenantAssignmentsResult {
	return v.value
}

func (v *NullableTenantAssignmentsResult) Set(val *TenantAssignmentsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAssignmentsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAssignmentsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAssignmentsResult(val *TenantAssignmentsResult) *NullableTenantAssignmentsResult {
	return &NullableTenantAssignmentsResult{value: val, isSet: true}
}

func (v NullableTenantAssignmentsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAssignmentsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


