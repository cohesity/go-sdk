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

// Cluster Specifies the cluster details.
type Cluster struct {
	// Specifies the cluster id of the new cluster.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the incarnation id of the new cluster.
	IncarnationId NullableInt64 `json:"incarnationId,omitempty"`
	// Name of the new cluster.
	Name NullableString `json:"name,omitempty"`
	// Description of the cluster.
	Description NullableString `json:"description,omitempty"`
	// Specifies the type of the new cluster.
	Type NullableString `json:"type,omitempty"`
	// Specifies the local tenant id. Only applicable on Helios.
	LocalTenantId NullableString `json:"localTenantId,omitempty"`
	// Specifies the globally unique tenant id. Only applicable on Helios.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the region id on which this cluster is present. Only applicable on Helios for DMaaS clusters.
	RegionId NullableString `json:"regionId,omitempty"`
	RigelClusterParams NullableRigelClusterConfigParams `json:"rigelClusterParams,omitempty"`
	// Software version of the new cluster.
	SwVersion NullableString `json:"swVersion,omitempty"`
	NetworkConfig *ClusterCreateNetworkConfig `json:"networkConfig,omitempty"`
	ProxyServerConfig NullableClusterProxyServerConfig `json:"proxyServerConfig,omitempty"`
	// Specifies whether or not encryption is enabled. If encryption is enabled, all data on the Cluster will be encrypted.
	EnableEncryption NullableBool `json:"enableEncryption,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster() *Cluster {
	this := Cluster{}
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Cluster) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Cluster) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Cluster) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Cluster) UnsetId() {
	o.Id.Unset()
}

// GetIncarnationId returns the IncarnationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetIncarnationId() int64 {
	if o == nil || o.IncarnationId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IncarnationId.Get()
}

// GetIncarnationIdOk returns a tuple with the IncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetIncarnationIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncarnationId.Get(), o.IncarnationId.IsSet()
}

// HasIncarnationId returns a boolean if a field has been set.
func (o *Cluster) HasIncarnationId() bool {
	if o != nil && o.IncarnationId.IsSet() {
		return true
	}

	return false
}

// SetIncarnationId gets a reference to the given NullableInt64 and assigns it to the IncarnationId field.
func (o *Cluster) SetIncarnationId(v int64) {
	o.IncarnationId.Set(&v)
}
// SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil
func (o *Cluster) SetIncarnationIdNil() {
	o.IncarnationId.Set(nil)
}

// UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
func (o *Cluster) UnsetIncarnationId() {
	o.IncarnationId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Cluster) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Cluster) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Cluster) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Cluster) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Cluster) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Cluster) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Cluster) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Cluster) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Cluster) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Cluster) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Cluster) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Cluster) UnsetType() {
	o.Type.Unset()
}

// GetLocalTenantId returns the LocalTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetLocalTenantId() string {
	if o == nil || o.LocalTenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocalTenantId.Get()
}

// GetLocalTenantIdOk returns a tuple with the LocalTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLocalTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocalTenantId.Get(), o.LocalTenantId.IsSet()
}

// HasLocalTenantId returns a boolean if a field has been set.
func (o *Cluster) HasLocalTenantId() bool {
	if o != nil && o.LocalTenantId.IsSet() {
		return true
	}

	return false
}

// SetLocalTenantId gets a reference to the given NullableString and assigns it to the LocalTenantId field.
func (o *Cluster) SetLocalTenantId(v string) {
	o.LocalTenantId.Set(&v)
}
// SetLocalTenantIdNil sets the value for LocalTenantId to be an explicit nil
func (o *Cluster) SetLocalTenantIdNil() {
	o.LocalTenantId.Set(nil)
}

// UnsetLocalTenantId ensures that no value is present for LocalTenantId, not even an explicit nil
func (o *Cluster) UnsetLocalTenantId() {
	o.LocalTenantId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Cluster) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Cluster) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Cluster) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Cluster) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetRegionId returns the RegionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetRegionId() string {
	if o == nil || o.RegionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RegionId.Get()
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetRegionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RegionId.Get(), o.RegionId.IsSet()
}

// HasRegionId returns a boolean if a field has been set.
func (o *Cluster) HasRegionId() bool {
	if o != nil && o.RegionId.IsSet() {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given NullableString and assigns it to the RegionId field.
func (o *Cluster) SetRegionId(v string) {
	o.RegionId.Set(&v)
}
// SetRegionIdNil sets the value for RegionId to be an explicit nil
func (o *Cluster) SetRegionIdNil() {
	o.RegionId.Set(nil)
}

// UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
func (o *Cluster) UnsetRegionId() {
	o.RegionId.Unset()
}

// GetRigelClusterParams returns the RigelClusterParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetRigelClusterParams() RigelClusterConfigParams {
	if o == nil || o.RigelClusterParams.Get() == nil {
		var ret RigelClusterConfigParams
		return ret
	}
	return *o.RigelClusterParams.Get()
}

// GetRigelClusterParamsOk returns a tuple with the RigelClusterParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetRigelClusterParamsOk() (*RigelClusterConfigParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RigelClusterParams.Get(), o.RigelClusterParams.IsSet()
}

// HasRigelClusterParams returns a boolean if a field has been set.
func (o *Cluster) HasRigelClusterParams() bool {
	if o != nil && o.RigelClusterParams.IsSet() {
		return true
	}

	return false
}

// SetRigelClusterParams gets a reference to the given NullableRigelClusterConfigParams and assigns it to the RigelClusterParams field.
func (o *Cluster) SetRigelClusterParams(v RigelClusterConfigParams) {
	o.RigelClusterParams.Set(&v)
}
// SetRigelClusterParamsNil sets the value for RigelClusterParams to be an explicit nil
func (o *Cluster) SetRigelClusterParamsNil() {
	o.RigelClusterParams.Set(nil)
}

// UnsetRigelClusterParams ensures that no value is present for RigelClusterParams, not even an explicit nil
func (o *Cluster) UnsetRigelClusterParams() {
	o.RigelClusterParams.Unset()
}

// GetSwVersion returns the SwVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetSwVersion() string {
	if o == nil || o.SwVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.SwVersion.Get()
}

// GetSwVersionOk returns a tuple with the SwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetSwVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SwVersion.Get(), o.SwVersion.IsSet()
}

// HasSwVersion returns a boolean if a field has been set.
func (o *Cluster) HasSwVersion() bool {
	if o != nil && o.SwVersion.IsSet() {
		return true
	}

	return false
}

// SetSwVersion gets a reference to the given NullableString and assigns it to the SwVersion field.
func (o *Cluster) SetSwVersion(v string) {
	o.SwVersion.Set(&v)
}
// SetSwVersionNil sets the value for SwVersion to be an explicit nil
func (o *Cluster) SetSwVersionNil() {
	o.SwVersion.Set(nil)
}

// UnsetSwVersion ensures that no value is present for SwVersion, not even an explicit nil
func (o *Cluster) UnsetSwVersion() {
	o.SwVersion.Unset()
}

// GetNetworkConfig returns the NetworkConfig field value if set, zero value otherwise.
func (o *Cluster) GetNetworkConfig() ClusterCreateNetworkConfig {
	if o == nil || o.NetworkConfig == nil {
		var ret ClusterCreateNetworkConfig
		return ret
	}
	return *o.NetworkConfig
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNetworkConfigOk() (*ClusterCreateNetworkConfig, bool) {
	if o == nil || o.NetworkConfig == nil {
		return nil, false
	}
	return o.NetworkConfig, true
}

// HasNetworkConfig returns a boolean if a field has been set.
func (o *Cluster) HasNetworkConfig() bool {
	if o != nil && o.NetworkConfig != nil {
		return true
	}

	return false
}

// SetNetworkConfig gets a reference to the given ClusterCreateNetworkConfig and assigns it to the NetworkConfig field.
func (o *Cluster) SetNetworkConfig(v ClusterCreateNetworkConfig) {
	o.NetworkConfig = &v
}

// GetProxyServerConfig returns the ProxyServerConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetProxyServerConfig() ClusterProxyServerConfig {
	if o == nil || o.ProxyServerConfig.Get() == nil {
		var ret ClusterProxyServerConfig
		return ret
	}
	return *o.ProxyServerConfig.Get()
}

// GetProxyServerConfigOk returns a tuple with the ProxyServerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetProxyServerConfigOk() (*ClusterProxyServerConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProxyServerConfig.Get(), o.ProxyServerConfig.IsSet()
}

// HasProxyServerConfig returns a boolean if a field has been set.
func (o *Cluster) HasProxyServerConfig() bool {
	if o != nil && o.ProxyServerConfig.IsSet() {
		return true
	}

	return false
}

// SetProxyServerConfig gets a reference to the given NullableClusterProxyServerConfig and assigns it to the ProxyServerConfig field.
func (o *Cluster) SetProxyServerConfig(v ClusterProxyServerConfig) {
	o.ProxyServerConfig.Set(&v)
}
// SetProxyServerConfigNil sets the value for ProxyServerConfig to be an explicit nil
func (o *Cluster) SetProxyServerConfigNil() {
	o.ProxyServerConfig.Set(nil)
}

// UnsetProxyServerConfig ensures that no value is present for ProxyServerConfig, not even an explicit nil
func (o *Cluster) UnsetProxyServerConfig() {
	o.ProxyServerConfig.Unset()
}

// GetEnableEncryption returns the EnableEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetEnableEncryption() bool {
	if o == nil || o.EnableEncryption.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableEncryption.Get()
}

// GetEnableEncryptionOk returns a tuple with the EnableEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetEnableEncryptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableEncryption.Get(), o.EnableEncryption.IsSet()
}

// HasEnableEncryption returns a boolean if a field has been set.
func (o *Cluster) HasEnableEncryption() bool {
	if o != nil && o.EnableEncryption.IsSet() {
		return true
	}

	return false
}

// SetEnableEncryption gets a reference to the given NullableBool and assigns it to the EnableEncryption field.
func (o *Cluster) SetEnableEncryption(v bool) {
	o.EnableEncryption.Set(&v)
}
// SetEnableEncryptionNil sets the value for EnableEncryption to be an explicit nil
func (o *Cluster) SetEnableEncryptionNil() {
	o.EnableEncryption.Set(nil)
}

// UnsetEnableEncryption ensures that no value is present for EnableEncryption, not even an explicit nil
func (o *Cluster) UnsetEnableEncryption() {
	o.EnableEncryption.Unset()
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IncarnationId.IsSet() {
		toSerialize["incarnationId"] = o.IncarnationId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.LocalTenantId.IsSet() {
		toSerialize["localTenantId"] = o.LocalTenantId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.RegionId.IsSet() {
		toSerialize["regionId"] = o.RegionId.Get()
	}
	if o.RigelClusterParams.IsSet() {
		toSerialize["rigelClusterParams"] = o.RigelClusterParams.Get()
	}
	if o.SwVersion.IsSet() {
		toSerialize["swVersion"] = o.SwVersion.Get()
	}
	if o.NetworkConfig != nil {
		toSerialize["networkConfig"] = o.NetworkConfig
	}
	if o.ProxyServerConfig.IsSet() {
		toSerialize["proxyServerConfig"] = o.ProxyServerConfig.Get()
	}
	if o.EnableEncryption.IsSet() {
		toSerialize["enableEncryption"] = o.EnableEncryption.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Cluster) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}