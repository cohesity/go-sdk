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

// checks if the HeliosClusterTenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosClusterTenant{}

// HeliosClusterTenant Description of a Tenant and cluster related properties.
type HeliosClusterTenant struct {
	// Specifies the cluster id of the cluster.
	ClusterId NullableInt64 `json:"clusterId,omitempty"`
	// Specifies the incarnation id of the cluster.
	ClusterIncarnationId NullableInt64 `json:"clusterIncarnationId,omitempty"`
	// Specifies the region id of the cluster. Only valid for DMaaS clusters.
	RegionId NullableString `json:"regionId,omitempty"`
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

// NewHeliosClusterTenant instantiates a new HeliosClusterTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosClusterTenant() *HeliosClusterTenant {
	this := HeliosClusterTenant{}
	return &this
}

// NewHeliosClusterTenantWithDefaults instantiates a new HeliosClusterTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosClusterTenantWithDefaults() *HeliosClusterTenant {
	this := HeliosClusterTenant{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetClusterId() int64 {
	if o == nil || IsNil(o.ClusterId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasClusterId() bool {
	if o != nil && o.ClusterId.IsSet() {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given NullableInt64 and assigns it to the ClusterId field.
func (o *HeliosClusterTenant) SetClusterId(v int64) {
	o.ClusterId.Set(&v)
}
// SetClusterIdNil sets the value for ClusterId to be an explicit nil
func (o *HeliosClusterTenant) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
func (o *HeliosClusterTenant) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterIncarnationId returns the ClusterIncarnationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetClusterIncarnationId() int64 {
	if o == nil || IsNil(o.ClusterIncarnationId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterIncarnationId.Get()
}

// GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetClusterIncarnationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterIncarnationId.Get(), o.ClusterIncarnationId.IsSet()
}

// HasClusterIncarnationId returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasClusterIncarnationId() bool {
	if o != nil && o.ClusterIncarnationId.IsSet() {
		return true
	}

	return false
}

// SetClusterIncarnationId gets a reference to the given NullableInt64 and assigns it to the ClusterIncarnationId field.
func (o *HeliosClusterTenant) SetClusterIncarnationId(v int64) {
	o.ClusterIncarnationId.Set(&v)
}
// SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil
func (o *HeliosClusterTenant) SetClusterIncarnationIdNil() {
	o.ClusterIncarnationId.Set(nil)
}

// UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
func (o *HeliosClusterTenant) UnsetClusterIncarnationId() {
	o.ClusterIncarnationId.Unset()
}

// GetRegionId returns the RegionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetRegionId() string {
	if o == nil || IsNil(o.RegionId.Get()) {
		var ret string
		return ret
	}
	return *o.RegionId.Get()
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetRegionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegionId.Get(), o.RegionId.IsSet()
}

// HasRegionId returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasRegionId() bool {
	if o != nil && o.RegionId.IsSet() {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given NullableString and assigns it to the RegionId field.
func (o *HeliosClusterTenant) SetRegionId(v string) {
	o.RegionId.Set(&v)
}
// SetRegionIdNil sets the value for RegionId to be an explicit nil
func (o *HeliosClusterTenant) SetRegionIdNil() {
	o.RegionId.Set(nil)
}

// UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
func (o *HeliosClusterTenant) UnsetRegionId() {
	o.RegionId.Unset()
}

// GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetCreatedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.CreatedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtTimeMsecs.Get()
}

// GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetCreatedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtTimeMsecs.Get(), o.CreatedAtTimeMsecs.IsSet()
}

// HasCreatedAtTimeMsecs returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasCreatedAtTimeMsecs() bool {
	if o != nil && o.CreatedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedAtTimeMsecs field.
func (o *HeliosClusterTenant) SetCreatedAtTimeMsecs(v int64) {
	o.CreatedAtTimeMsecs.Set(&v)
}
// SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil
func (o *HeliosClusterTenant) SetCreatedAtTimeMsecsNil() {
	o.CreatedAtTimeMsecs.Set(nil)
}

// UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
func (o *HeliosClusterTenant) UnsetCreatedAtTimeMsecs() {
	o.CreatedAtTimeMsecs.Unset()
}

// GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetDeletedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.DeletedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.DeletedAtTimeMsecs.Get()
}

// GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetDeletedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAtTimeMsecs.Get(), o.DeletedAtTimeMsecs.IsSet()
}

// HasDeletedAtTimeMsecs returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasDeletedAtTimeMsecs() bool {
	if o != nil && o.DeletedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetDeletedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the DeletedAtTimeMsecs field.
func (o *HeliosClusterTenant) SetDeletedAtTimeMsecs(v int64) {
	o.DeletedAtTimeMsecs.Set(&v)
}
// SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil
func (o *HeliosClusterTenant) SetDeletedAtTimeMsecsNil() {
	o.DeletedAtTimeMsecs.Set(nil)
}

// UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil
func (o *HeliosClusterTenant) UnsetDeletedAtTimeMsecs() {
	o.DeletedAtTimeMsecs.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *HeliosClusterTenant) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *HeliosClusterTenant) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *HeliosClusterTenant) UnsetDescription() {
	o.Description.Unset()
}

// GetExternalVendorMetadata returns the ExternalVendorMetadata field value if set, zero value otherwise.
func (o *HeliosClusterTenant) GetExternalVendorMetadata() ExternalVendorTenantMetadata {
	if o == nil || IsNil(o.ExternalVendorMetadata) {
		var ret ExternalVendorTenantMetadata
		return ret
	}
	return *o.ExternalVendorMetadata
}

// GetExternalVendorMetadataOk returns a tuple with the ExternalVendorMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosClusterTenant) GetExternalVendorMetadataOk() (*ExternalVendorTenantMetadata, bool) {
	if o == nil || IsNil(o.ExternalVendorMetadata) {
		return nil, false
	}
	return o.ExternalVendorMetadata, true
}

// HasExternalVendorMetadata returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasExternalVendorMetadata() bool {
	if o != nil && !IsNil(o.ExternalVendorMetadata) {
		return true
	}

	return false
}

// SetExternalVendorMetadata gets a reference to the given ExternalVendorTenantMetadata and assigns it to the ExternalVendorMetadata field.
func (o *HeliosClusterTenant) SetExternalVendorMetadata(v ExternalVendorTenantMetadata) {
	o.ExternalVendorMetadata = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HeliosClusterTenant) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HeliosClusterTenant) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HeliosClusterTenant) UnsetId() {
	o.Id.Unset()
}

// GetIsManagedOnHelios returns the IsManagedOnHelios field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetIsManagedOnHelios() bool {
	if o == nil || IsNil(o.IsManagedOnHelios.Get()) {
		var ret bool
		return ret
	}
	return *o.IsManagedOnHelios.Get()
}

// GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetIsManagedOnHeliosOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsManagedOnHelios.Get(), o.IsManagedOnHelios.IsSet()
}

// HasIsManagedOnHelios returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasIsManagedOnHelios() bool {
	if o != nil && o.IsManagedOnHelios.IsSet() {
		return true
	}

	return false
}

// SetIsManagedOnHelios gets a reference to the given NullableBool and assigns it to the IsManagedOnHelios field.
func (o *HeliosClusterTenant) SetIsManagedOnHelios(v bool) {
	o.IsManagedOnHelios.Set(&v)
}
// SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil
func (o *HeliosClusterTenant) SetIsManagedOnHeliosNil() {
	o.IsManagedOnHelios.Set(nil)
}

// UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
func (o *HeliosClusterTenant) UnsetIsManagedOnHelios() {
	o.IsManagedOnHelios.Unset()
}

// GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetLastUpdatedAtTimeMsecs() int64 {
	if o == nil || IsNil(o.LastUpdatedAtTimeMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedAtTimeMsecs.Get()
}

// GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetLastUpdatedAtTimeMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdatedAtTimeMsecs.Get(), o.LastUpdatedAtTimeMsecs.IsSet()
}

// HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasLastUpdatedAtTimeMsecs() bool {
	if o != nil && o.LastUpdatedAtTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedAtTimeMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedAtTimeMsecs field.
func (o *HeliosClusterTenant) SetLastUpdatedAtTimeMsecs(v int64) {
	o.LastUpdatedAtTimeMsecs.Set(&v)
}
// SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil
func (o *HeliosClusterTenant) SetLastUpdatedAtTimeMsecsNil() {
	o.LastUpdatedAtTimeMsecs.Set(nil)
}

// UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
func (o *HeliosClusterTenant) UnsetLastUpdatedAtTimeMsecs() {
	o.LastUpdatedAtTimeMsecs.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HeliosClusterTenant) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HeliosClusterTenant) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HeliosClusterTenant) UnsetName() {
	o.Name.Unset()
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *HeliosClusterTenant) GetNetwork() TenantNetwork {
	if o == nil || IsNil(o.Network) {
		var ret TenantNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosClusterTenant) GetNetworkOk() (*TenantNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TenantNetwork and assigns it to the Network field.
func (o *HeliosClusterTenant) SetNetwork(v TenantNetwork) {
	o.Network = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosClusterTenant) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosClusterTenant) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *HeliosClusterTenant) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *HeliosClusterTenant) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *HeliosClusterTenant) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *HeliosClusterTenant) UnsetStatus() {
	o.Status.Unset()
}

func (o HeliosClusterTenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosClusterTenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterId.IsSet() {
		toSerialize["clusterId"] = o.ClusterId.Get()
	}
	if o.ClusterIncarnationId.IsSet() {
		toSerialize["clusterIncarnationId"] = o.ClusterIncarnationId.Get()
	}
	if o.RegionId.IsSet() {
		toSerialize["regionId"] = o.RegionId.Get()
	}
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

type NullableHeliosClusterTenant struct {
	value *HeliosClusterTenant
	isSet bool
}

func (v NullableHeliosClusterTenant) Get() *HeliosClusterTenant {
	return v.value
}

func (v *NullableHeliosClusterTenant) Set(val *HeliosClusterTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosClusterTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosClusterTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosClusterTenant(val *HeliosClusterTenant) *NullableHeliosClusterTenant {
	return &NullableHeliosClusterTenant{value: val, isSet: true}
}

func (v NullableHeliosClusterTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosClusterTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


