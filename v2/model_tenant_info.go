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

// checks if the TenantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantInfo{}

// TenantInfo Description of a Tenant and it's properties.
type TenantInfo struct {
	// Epoch time when tenant was created.
	CreatedAtTimeMsecs NullableInt64 `json:"createdAtTimeMsecs,omitempty"`
	// Epoch time when tenant was last updated.
	DeletedAtTimeMsecs NullableInt64 `json:"deletedAtTimeMsecs,omitempty"`
	// Description about the tenant.
	Description NullableString `json:"description,omitempty"`
	ExternalVendorMetadata *ExternalVendorTenantMetadata `json:"externalVendorMetadata,omitempty"`
	// The tenant id.
	Id NullableString `json:"id,omitempty"`
	// Flag to indicate if tenant is managed on helios
	IsManagedOnHelios NullableBool `json:"isManagedOnHelios,omitempty"`
	// Epoch time when tenant was last updated.
	LastUpdatedAtTimeMsecs NullableInt64 `json:"lastUpdatedAtTimeMsecs,omitempty"`
	// Name of the Tenant.
	Name NullableString `json:"name,omitempty"`
	Network *TenantNetwork `json:"network,omitempty"`
	// Current Status of the Tenant.
	Status NullableString `json:"status,omitempty"`
}

// NewTenantInfo instantiates a new TenantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantInfo() *TenantInfo {
	this := TenantInfo{}
	return &this
}

// NewTenantInfoWithDefaults instantiates a new TenantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantInfoWithDefaults() *TenantInfo {
	this := TenantInfo{}
	return &this
}

// GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetCreatedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.CreatedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtTimeMsecs.Get()
}

// GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetCreatedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtTimeMsecs.Get(), o.CreatedAtTimeMsecs.IsSet()
}

// HasCreatedAtTimeMsecs returns a boolean if a field has been set.
func (o *TenantInfo) HasCreatedAtTimeMsecs() bool {
	if o != nil && o.CreatedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedAtTimeMsecs field.
func (o *TenantInfo) SetCreatedAtTimeMsecs(v int64) {
	o.CreatedAtTimeMsecs.Set(&v)
}
// SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil
func (o *TenantInfo) SetCreatedAtTimeMsecsNil() {
	o.CreatedAtTimeMsecs.Set(nil)
}

// UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
func (o *TenantInfo) UnsetCreatedAtTimeMsecs() {
	o.CreatedAtTimeMsecs.Unset()
}

// GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetDeletedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.DeletedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.DeletedAtTimeMsecs.Get()
}

// GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetDeletedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAtTimeMsecs.Get(), o.DeletedAtTimeMsecs.IsSet()
}

// HasDeletedAtTimeMsecs returns a boolean if a field has been set.
func (o *TenantInfo) HasDeletedAtTimeMsecs() bool {
	if o != nil && o.DeletedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetDeletedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the DeletedAtTimeMsecs field.
func (o *TenantInfo) SetDeletedAtTimeMsecs(v int64) {
	o.DeletedAtTimeMsecs.Set(&v)
}
// SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil
func (o *TenantInfo) SetDeletedAtTimeMsecsNil() {
	o.DeletedAtTimeMsecs.Set(nil)
}

// UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil
func (o *TenantInfo) UnsetDeletedAtTimeMsecs() {
	o.DeletedAtTimeMsecs.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TenantInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TenantInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TenantInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetExternalVendorMetadata returns the ExternalVendorMetadata field value if set, zero value otherwise.
func (o *TenantInfo) GetExternalVendorMetadata() ExternalVendorTenantMetadata {
	if o == nil || IsNil(o.ExternalVendorMetadata) {
		var ret ExternalVendorTenantMetadata
		return ret
	}
	return *o.ExternalVendorMetadata
}

// GetExternalVendorMetadataOk returns a tuple with the ExternalVendorMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantInfo) GetExternalVendorMetadataOk() (*ExternalVendorTenantMetadata, bool) {
	if o == nil || IsNil(o.ExternalVendorMetadata) {
		return nil, false
	}
	return o.ExternalVendorMetadata, true
}

// HasExternalVendorMetadata returns a boolean if a field has been set.
func (o *TenantInfo) HasExternalVendorMetadata() bool {
	if o != nil && !IsNil(o.ExternalVendorMetadata) {
		return true
	}

	return false
}

// SetExternalVendorMetadata gets a reference to the given ExternalVendorTenantMetadata and assigns it to the ExternalVendorMetadata field.
func (o *TenantInfo) SetExternalVendorMetadata(v ExternalVendorTenantMetadata) {
	o.ExternalVendorMetadata = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TenantInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TenantInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TenantInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TenantInfo) UnsetId() {
	o.Id.Unset()
}

// GetIsManagedOnHelios returns the IsManagedOnHelios field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetIsManagedOnHelios() bool {
	if o == nil || IsNil(o.IsManagedOnHelios.Get()) {
		var ret bool
		return ret
	}
	return *o.IsManagedOnHelios.Get()
}

// GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetIsManagedOnHeliosOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsManagedOnHelios.Get(), o.IsManagedOnHelios.IsSet()
}

// HasIsManagedOnHelios returns a boolean if a field has been set.
func (o *TenantInfo) HasIsManagedOnHelios() bool {
	if o != nil && o.IsManagedOnHelios.IsSet() {
		return true
	}

	return false
}

// SetIsManagedOnHelios gets a reference to the given NullableBool and assigns it to the IsManagedOnHelios field.
func (o *TenantInfo) SetIsManagedOnHelios(v bool) {
	o.IsManagedOnHelios.Set(&v)
}
// SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil
func (o *TenantInfo) SetIsManagedOnHeliosNil() {
	o.IsManagedOnHelios.Set(nil)
}

// UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
func (o *TenantInfo) UnsetIsManagedOnHelios() {
	o.IsManagedOnHelios.Unset()
}

// GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetLastUpdatedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.LastUpdatedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedAtTimeMsecs.Get()
}

// GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetLastUpdatedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdatedAtTimeMsecs.Get(), o.LastUpdatedAtTimeMsecs.IsSet()
}

// HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.
func (o *TenantInfo) HasLastUpdatedAtTimeMsecs() bool {
	if o != nil && o.LastUpdatedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedAtTimeMsecs field.
func (o *TenantInfo) SetLastUpdatedAtTimeMsecs(v int64) {
	o.LastUpdatedAtTimeMsecs.Set(&v)
}
// SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil
func (o *TenantInfo) SetLastUpdatedAtTimeMsecsNil() {
	o.LastUpdatedAtTimeMsecs.Set(nil)
}

// UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
func (o *TenantInfo) UnsetLastUpdatedAtTimeMsecs() {
	o.LastUpdatedAtTimeMsecs.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TenantInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TenantInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TenantInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TenantInfo) UnsetName() {
	o.Name.Unset()
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TenantInfo) GetNetwork() TenantNetwork {
	if o == nil || IsNil(o.Network) {
		var ret TenantNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantInfo) GetNetworkOk() (*TenantNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TenantInfo) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TenantNetwork and assigns it to the Network field.
func (o *TenantInfo) SetNetwork(v TenantNetwork) {
	o.Network = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantInfo) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantInfo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TenantInfo) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TenantInfo) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TenantInfo) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TenantInfo) UnsetStatus() {
	o.Status.Unset()
}

func (o TenantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAtTimeMsecs.IsSet() {
		toSerialize["createdAtTimeMsecs"] = o.CreatedAtTimeMsecs.Get()
	}
	if o.DeletedAtTimeMsecs.IsSet() {
		toSerialize["deletedAtTimeMsecs"] = o.DeletedAtTimeMsecs.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.ExternalVendorMetadata) {
		toSerialize["externalVendorMetadata"] = o.ExternalVendorMetadata
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsManagedOnHelios.IsSet() {
		toSerialize["isManagedOnHelios"] = o.IsManagedOnHelios.Get()
	}
	if o.LastUpdatedAtTimeMsecs.IsSet() {
		toSerialize["lastUpdatedAtTimeMsecs"] = o.LastUpdatedAtTimeMsecs.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullableTenantInfo struct {
	value *TenantInfo
	isSet bool
}

func (v NullableTenantInfo) Get() *TenantInfo {
	return v.value
}

func (v *NullableTenantInfo) Set(val *TenantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantInfo(val *TenantInfo) *NullableTenantInfo {
	return &NullableTenantInfo{value: val, isSet: true}
}

func (v NullableTenantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


