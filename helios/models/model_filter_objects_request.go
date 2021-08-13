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

// FilterObjectsRequest Specifies the filter details.
type FilterObjectsRequest struct {
	// Specifies the type of filtering user wants to perform. Currently, we only support exclude type of filter.
	FilterType NullableString `json:"filterType"`
	// Specifies the list of filters that need to be applied on given list of discovered objects.
	Filters []Filter `json:"filters"`
	// Specifies a list of non leaf object ids to filter the leaf level objects. Non leaf object such host (physical or vm) or database instance can be specified.
	ObjectIds []int64 `json:"objectIds"`
	// Specifies the type of application enviornment needed for filtering to be applied on. This is needed because in case of applications like SQL, Oracle, a single source can contain multiple application enviornments.
	ApplicationEnvironment NullableString `json:"applicationEnvironment,omitempty"`
	// TenantIds contains list of the tenant for which objects are to be returned.
	TenantIds []string `json:"tenantIds,omitempty"`
	// If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false.
	IncludeTenants NullableBool `json:"includeTenants,omitempty"`
}

// NewFilterObjectsRequest instantiates a new FilterObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterObjectsRequest(filterType NullableString, filters []Filter, objectIds []int64) *FilterObjectsRequest {
	this := FilterObjectsRequest{}
	this.FilterType = filterType
	this.Filters = filters
	this.ObjectIds = objectIds
	var includeTenants bool = false
	this.IncludeTenants = *NewNullableBool(&includeTenants)
	return &this
}

// NewFilterObjectsRequestWithDefaults instantiates a new FilterObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterObjectsRequestWithDefaults() *FilterObjectsRequest {
	this := FilterObjectsRequest{}
	var includeTenants bool = false
	this.IncludeTenants = *NewNullableBool(&includeTenants) // model_simple false
	return &this
}

// GetFilterType returns the FilterType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilterObjectsRequest) GetFilterType() string {
	if o == nil || o.FilterType.Get() == nil {
		var ret string
		return ret
	}

	return *o.FilterType.Get()
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetFilterTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FilterType.Get(), o.FilterType.IsSet()
}

// SetFilterType sets field value
func (o *FilterObjectsRequest) SetFilterType(v string) {
	o.FilterType.Set(&v)
}

// GetFilters returns the Filters field value
// If the value is explicit nil, the zero value for []Filter will be returned
func (o *FilterObjectsRequest) GetFilters() []Filter {
	if o == nil {
		var ret []Filter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetFiltersOk() (*[]Filter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *FilterObjectsRequest) SetFilters(v []Filter) {
	o.Filters = v
}

// GetObjectIds returns the ObjectIds field value
// If the value is explicit nil, the zero value for []int64 will be returned
func (o *FilterObjectsRequest) GetObjectIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetObjectIdsOk() (*[]int64, bool) {
	if o == nil || o.ObjectIds == nil {
		return nil, false
	}
	return &o.ObjectIds, true
}

// SetObjectIds sets field value
func (o *FilterObjectsRequest) SetObjectIds(v []int64) {
	o.ObjectIds = v
}

// GetApplicationEnvironment returns the ApplicationEnvironment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterObjectsRequest) GetApplicationEnvironment() string {
	if o == nil || o.ApplicationEnvironment.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationEnvironment.Get()
}

// GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetApplicationEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationEnvironment.Get(), o.ApplicationEnvironment.IsSet()
}

// HasApplicationEnvironment returns a boolean if a field has been set.
func (o *FilterObjectsRequest) HasApplicationEnvironment() bool {
	if o != nil && o.ApplicationEnvironment.IsSet() {
		return true
	}

	return false
}

// SetApplicationEnvironment gets a reference to the given NullableString and assigns it to the ApplicationEnvironment field.
func (o *FilterObjectsRequest) SetApplicationEnvironment(v string) {
	o.ApplicationEnvironment.Set(&v)
}
// SetApplicationEnvironmentNil sets the value for ApplicationEnvironment to be an explicit nil
func (o *FilterObjectsRequest) SetApplicationEnvironmentNil() {
	o.ApplicationEnvironment.Set(nil)
}

// UnsetApplicationEnvironment ensures that no value is present for ApplicationEnvironment, not even an explicit nil
func (o *FilterObjectsRequest) UnsetApplicationEnvironment() {
	o.ApplicationEnvironment.Unset()
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterObjectsRequest) GetTenantIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return &o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *FilterObjectsRequest) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *FilterObjectsRequest) SetTenantIds(v []string) {
	o.TenantIds = v
}

// GetIncludeTenants returns the IncludeTenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterObjectsRequest) GetIncludeTenants() bool {
	if o == nil || o.IncludeTenants.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTenants.Get()
}

// GetIncludeTenantsOk returns a tuple with the IncludeTenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterObjectsRequest) GetIncludeTenantsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeTenants.Get(), o.IncludeTenants.IsSet()
}

// HasIncludeTenants returns a boolean if a field has been set.
func (o *FilterObjectsRequest) HasIncludeTenants() bool {
	if o != nil && o.IncludeTenants.IsSet() {
		return true
	}

	return false
}

// SetIncludeTenants gets a reference to the given NullableBool and assigns it to the IncludeTenants field.
func (o *FilterObjectsRequest) SetIncludeTenants(v bool) {
	o.IncludeTenants.Set(&v)
}
// SetIncludeTenantsNil sets the value for IncludeTenants to be an explicit nil
func (o *FilterObjectsRequest) SetIncludeTenantsNil() {
	o.IncludeTenants.Set(nil)
}

// UnsetIncludeTenants ensures that no value is present for IncludeTenants, not even an explicit nil
func (o *FilterObjectsRequest) UnsetIncludeTenants() {
	o.IncludeTenants.Unset()
}

func (o FilterObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filterType"] = o.FilterType.Get()
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.ObjectIds != nil {
		toSerialize["objectIds"] = o.ObjectIds
	}
	if o.ApplicationEnvironment.IsSet() {
		toSerialize["applicationEnvironment"] = o.ApplicationEnvironment.Get()
	}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	if o.IncludeTenants.IsSet() {
		toSerialize["includeTenants"] = o.IncludeTenants.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFilterObjectsRequest struct {
	value *FilterObjectsRequest
	isSet bool
}

func (v NullableFilterObjectsRequest) Get() *FilterObjectsRequest {
	return v.value
}

func (v *NullableFilterObjectsRequest) Set(val *FilterObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterObjectsRequest(val *FilterObjectsRequest) *NullableFilterObjectsRequest {
	return &NullableFilterObjectsRequest{value: val, isSet: true}
}

func (v NullableFilterObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


