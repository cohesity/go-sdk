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

// checks if the RemoteCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteCluster{}

// RemoteCluster Specifies a Remote Cluster config.
type RemoteCluster struct {
	// Specifies if the Tx clusters should be automatically registered at the Rx site.
	AutoRegisterTarget NullableBool `json:"autoRegisterTarget,omitempty"`
	// Specifies the Remote Cluster id.
	ClusterId NullableInt64 `json:"clusterId,omitempty"`
	// Specifies the Remote Cluster incarnation id.
	ClusterIncarnationId NullableInt64 `json:"clusterIncarnationId,omitempty"`
	// Specifies the Remote Cluster name.
	ClusterName NullableString `json:"clusterName,omitempty"`
	// Specifies any additional information if needed.
	Description NullableString `json:"description,omitempty"`
	// Specifies the effective AES Encryption mode negotiated between local and the remote cluster.
	EffectiveAesEncryptionMode NullableString `json:"effectiveAesEncryptionMode,omitempty"`
	// Specifies if the Remote Cluster was registered automatically or manually.
	IsAutoRegistered NullableBool `json:"isAutoRegistered,omitempty"`
	// Specifies the IP addresses of the interfaces in the local Cluster which will be used for communicating with the remote Cluster.
	LocalAddresses []string `json:"localAddresses,omitempty"`
	// Specifies if the Remote Cluster has Multi-Tenancy enabled.
	MultiTenancyEnabled NullableBool `json:"multiTenancyEnabled,omitempty"`
	// Specifies the name of the network interfaces to use for communicating with the Remote Cluster.
	NetworkInterface NullableString `json:"networkInterface,omitempty"`
	// Specifies the purpose for which the remote cluster is being registered.
	Purpose []string `json:"purpose,omitempty"`
	ReplicationParams *RemoteClusterParamsReplicationParams `json:"replicationParams,omitempty"`
	// Specifies the AES Encryption mode of the remote cluster.
	SupportedAesEncryptionMode NullableString `json:"supportedAesEncryptionMode,omitempty"`
	// Specifies the tenant Id of the Remote Cluster.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies if Tenant Storage Domain sharing is enabled on the Remote Cluster.
	TenantStorageDomainSharingEnabled NullableBool `json:"tenantStorageDomainSharingEnabled,omitempty"`
	// Specifies if TLS is enabled on the Remote Cluster.
	TlsEnabled NullableBool `json:"tlsEnabled,omitempty"`
	// Specifies the VIP or IP addresses of the Nodes on the Remote Cluster to connect with. Hostnames are not supported.
	NodeAddresses []string `json:"nodeAddresses,omitempty"`
	// Specifies the password for Cohesity user to use when connecting to the Remote Cluster.
	Password NullableString `json:"password,omitempty"`
	// Specifies the Cohesity user name used to connect to the Remote Cluster.
	Username *string `json:"username,omitempty"`
}

// NewRemoteCluster instantiates a new RemoteCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteCluster() *RemoteCluster {
	this := RemoteCluster{}
	var autoRegisterTarget bool = false
	this.AutoRegisterTarget = *NewNullableBool(&autoRegisterTarget)
	return &this
}

// NewRemoteClusterWithDefaults instantiates a new RemoteCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteClusterWithDefaults() *RemoteCluster {
	this := RemoteCluster{}
	var autoRegisterTarget bool = false
	this.AutoRegisterTarget = *NewNullableBool(&autoRegisterTarget)
	return &this
}

// GetAutoRegisterTarget returns the AutoRegisterTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetAutoRegisterTarget() bool {
	if o == nil || IsNil(o.AutoRegisterTarget.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoRegisterTarget.Get()
}

// GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetAutoRegisterTargetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRegisterTarget.Get(), o.AutoRegisterTarget.IsSet()
}

// HasAutoRegisterTarget returns a boolean if a field has been set.
func (o *RemoteCluster) HasAutoRegisterTarget() bool {
	if o != nil && o.AutoRegisterTarget.IsSet() {
		return true
	}

	return false
}

// SetAutoRegisterTarget gets a reference to the given NullableBool and assigns it to the AutoRegisterTarget field.
func (o *RemoteCluster) SetAutoRegisterTarget(v bool) {
	o.AutoRegisterTarget.Set(&v)
}
// SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil
func (o *RemoteCluster) SetAutoRegisterTargetNil() {
	o.AutoRegisterTarget.Set(nil)
}

// UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
func (o *RemoteCluster) UnsetAutoRegisterTarget() {
	o.AutoRegisterTarget.Unset()
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetClusterId() int64 {
	if o == nil || IsNil(o.ClusterId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *RemoteCluster) HasClusterId() bool {
	if o != nil && o.ClusterId.IsSet() {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given NullableInt64 and assigns it to the ClusterId field.
func (o *RemoteCluster) SetClusterId(v int64) {
	o.ClusterId.Set(&v)
}
// SetClusterIdNil sets the value for ClusterId to be an explicit nil
func (o *RemoteCluster) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
func (o *RemoteCluster) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterIncarnationId returns the ClusterIncarnationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetClusterIncarnationId() int64 {
	if o == nil || IsNil(o.ClusterIncarnationId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterIncarnationId.Get()
}

// GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetClusterIncarnationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterIncarnationId.Get(), o.ClusterIncarnationId.IsSet()
}

// HasClusterIncarnationId returns a boolean if a field has been set.
func (o *RemoteCluster) HasClusterIncarnationId() bool {
	if o != nil && o.ClusterIncarnationId.IsSet() {
		return true
	}

	return false
}

// SetClusterIncarnationId gets a reference to the given NullableInt64 and assigns it to the ClusterIncarnationId field.
func (o *RemoteCluster) SetClusterIncarnationId(v int64) {
	o.ClusterIncarnationId.Set(&v)
}
// SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil
func (o *RemoteCluster) SetClusterIncarnationIdNil() {
	o.ClusterIncarnationId.Set(nil)
}

// UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
func (o *RemoteCluster) UnsetClusterIncarnationId() {
	o.ClusterIncarnationId.Unset()
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName.Get()) {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *RemoteCluster) HasClusterName() bool {
	if o != nil && o.ClusterName.IsSet() {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given NullableString and assigns it to the ClusterName field.
func (o *RemoteCluster) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}
// SetClusterNameNil sets the value for ClusterName to be an explicit nil
func (o *RemoteCluster) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
func (o *RemoteCluster) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RemoteCluster) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RemoteCluster) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RemoteCluster) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RemoteCluster) UnsetDescription() {
	o.Description.Unset()
}

// GetEffectiveAesEncryptionMode returns the EffectiveAesEncryptionMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetEffectiveAesEncryptionMode() string {
	if o == nil || IsNil(o.EffectiveAesEncryptionMode.Get()) {
		var ret string
		return ret
	}
	return *o.EffectiveAesEncryptionMode.Get()
}

// GetEffectiveAesEncryptionModeOk returns a tuple with the EffectiveAesEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetEffectiveAesEncryptionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectiveAesEncryptionMode.Get(), o.EffectiveAesEncryptionMode.IsSet()
}

// HasEffectiveAesEncryptionMode returns a boolean if a field has been set.
func (o *RemoteCluster) HasEffectiveAesEncryptionMode() bool {
	if o != nil && o.EffectiveAesEncryptionMode.IsSet() {
		return true
	}

	return false
}

// SetEffectiveAesEncryptionMode gets a reference to the given NullableString and assigns it to the EffectiveAesEncryptionMode field.
func (o *RemoteCluster) SetEffectiveAesEncryptionMode(v string) {
	o.EffectiveAesEncryptionMode.Set(&v)
}
// SetEffectiveAesEncryptionModeNil sets the value for EffectiveAesEncryptionMode to be an explicit nil
func (o *RemoteCluster) SetEffectiveAesEncryptionModeNil() {
	o.EffectiveAesEncryptionMode.Set(nil)
}

// UnsetEffectiveAesEncryptionMode ensures that no value is present for EffectiveAesEncryptionMode, not even an explicit nil
func (o *RemoteCluster) UnsetEffectiveAesEncryptionMode() {
	o.EffectiveAesEncryptionMode.Unset()
}

// GetIsAutoRegistered returns the IsAutoRegistered field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetIsAutoRegistered() bool {
	if o == nil || IsNil(o.IsAutoRegistered.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutoRegistered.Get()
}

// GetIsAutoRegisteredOk returns a tuple with the IsAutoRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetIsAutoRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutoRegistered.Get(), o.IsAutoRegistered.IsSet()
}

// HasIsAutoRegistered returns a boolean if a field has been set.
func (o *RemoteCluster) HasIsAutoRegistered() bool {
	if o != nil && o.IsAutoRegistered.IsSet() {
		return true
	}

	return false
}

// SetIsAutoRegistered gets a reference to the given NullableBool and assigns it to the IsAutoRegistered field.
func (o *RemoteCluster) SetIsAutoRegistered(v bool) {
	o.IsAutoRegistered.Set(&v)
}
// SetIsAutoRegisteredNil sets the value for IsAutoRegistered to be an explicit nil
func (o *RemoteCluster) SetIsAutoRegisteredNil() {
	o.IsAutoRegistered.Set(nil)
}

// UnsetIsAutoRegistered ensures that no value is present for IsAutoRegistered, not even an explicit nil
func (o *RemoteCluster) UnsetIsAutoRegistered() {
	o.IsAutoRegistered.Unset()
}

// GetLocalAddresses returns the LocalAddresses field value if set, zero value otherwise.
func (o *RemoteCluster) GetLocalAddresses() []string {
	if o == nil || IsNil(o.LocalAddresses) {
		var ret []string
		return ret
	}
	return o.LocalAddresses
}

// GetLocalAddressesOk returns a tuple with the LocalAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetLocalAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalAddresses) {
		return nil, false
	}
	return o.LocalAddresses, true
}

// HasLocalAddresses returns a boolean if a field has been set.
func (o *RemoteCluster) HasLocalAddresses() bool {
	if o != nil && !IsNil(o.LocalAddresses) {
		return true
	}

	return false
}

// SetLocalAddresses gets a reference to the given []string and assigns it to the LocalAddresses field.
func (o *RemoteCluster) SetLocalAddresses(v []string) {
	o.LocalAddresses = v
}

// GetMultiTenancyEnabled returns the MultiTenancyEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetMultiTenancyEnabled() bool {
	if o == nil || IsNil(o.MultiTenancyEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.MultiTenancyEnabled.Get()
}

// GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetMultiTenancyEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultiTenancyEnabled.Get(), o.MultiTenancyEnabled.IsSet()
}

// HasMultiTenancyEnabled returns a boolean if a field has been set.
func (o *RemoteCluster) HasMultiTenancyEnabled() bool {
	if o != nil && o.MultiTenancyEnabled.IsSet() {
		return true
	}

	return false
}

// SetMultiTenancyEnabled gets a reference to the given NullableBool and assigns it to the MultiTenancyEnabled field.
func (o *RemoteCluster) SetMultiTenancyEnabled(v bool) {
	o.MultiTenancyEnabled.Set(&v)
}
// SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil
func (o *RemoteCluster) SetMultiTenancyEnabledNil() {
	o.MultiTenancyEnabled.Set(nil)
}

// UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
func (o *RemoteCluster) UnsetMultiTenancyEnabled() {
	o.MultiTenancyEnabled.Unset()
}

// GetNetworkInterface returns the NetworkInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetNetworkInterface() string {
	if o == nil || IsNil(o.NetworkInterface.Get()) {
		var ret string
		return ret
	}
	return *o.NetworkInterface.Get()
}

// GetNetworkInterfaceOk returns a tuple with the NetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetNetworkInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkInterface.Get(), o.NetworkInterface.IsSet()
}

// HasNetworkInterface returns a boolean if a field has been set.
func (o *RemoteCluster) HasNetworkInterface() bool {
	if o != nil && o.NetworkInterface.IsSet() {
		return true
	}

	return false
}

// SetNetworkInterface gets a reference to the given NullableString and assigns it to the NetworkInterface field.
func (o *RemoteCluster) SetNetworkInterface(v string) {
	o.NetworkInterface.Set(&v)
}
// SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil
func (o *RemoteCluster) SetNetworkInterfaceNil() {
	o.NetworkInterface.Set(nil)
}

// UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
func (o *RemoteCluster) UnsetNetworkInterface() {
	o.NetworkInterface.Unset()
}

// GetPurpose returns the Purpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetPurpose() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetPurposeOk() ([]string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *RemoteCluster) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given []string and assigns it to the Purpose field.
func (o *RemoteCluster) SetPurpose(v []string) {
	o.Purpose = v
}

// GetReplicationParams returns the ReplicationParams field value if set, zero value otherwise.
func (o *RemoteCluster) GetReplicationParams() RemoteClusterParamsReplicationParams {
	if o == nil || IsNil(o.ReplicationParams) {
		var ret RemoteClusterParamsReplicationParams
		return ret
	}
	return *o.ReplicationParams
}

// GetReplicationParamsOk returns a tuple with the ReplicationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetReplicationParamsOk() (*RemoteClusterParamsReplicationParams, bool) {
	if o == nil || IsNil(o.ReplicationParams) {
		return nil, false
	}
	return o.ReplicationParams, true
}

// HasReplicationParams returns a boolean if a field has been set.
func (o *RemoteCluster) HasReplicationParams() bool {
	if o != nil && !IsNil(o.ReplicationParams) {
		return true
	}

	return false
}

// SetReplicationParams gets a reference to the given RemoteClusterParamsReplicationParams and assigns it to the ReplicationParams field.
func (o *RemoteCluster) SetReplicationParams(v RemoteClusterParamsReplicationParams) {
	o.ReplicationParams = &v
}

// GetSupportedAesEncryptionMode returns the SupportedAesEncryptionMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetSupportedAesEncryptionMode() string {
	if o == nil || IsNil(o.SupportedAesEncryptionMode.Get()) {
		var ret string
		return ret
	}
	return *o.SupportedAesEncryptionMode.Get()
}

// GetSupportedAesEncryptionModeOk returns a tuple with the SupportedAesEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetSupportedAesEncryptionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedAesEncryptionMode.Get(), o.SupportedAesEncryptionMode.IsSet()
}

// HasSupportedAesEncryptionMode returns a boolean if a field has been set.
func (o *RemoteCluster) HasSupportedAesEncryptionMode() bool {
	if o != nil && o.SupportedAesEncryptionMode.IsSet() {
		return true
	}

	return false
}

// SetSupportedAesEncryptionMode gets a reference to the given NullableString and assigns it to the SupportedAesEncryptionMode field.
func (o *RemoteCluster) SetSupportedAesEncryptionMode(v string) {
	o.SupportedAesEncryptionMode.Set(&v)
}
// SetSupportedAesEncryptionModeNil sets the value for SupportedAesEncryptionMode to be an explicit nil
func (o *RemoteCluster) SetSupportedAesEncryptionModeNil() {
	o.SupportedAesEncryptionMode.Set(nil)
}

// UnsetSupportedAesEncryptionMode ensures that no value is present for SupportedAesEncryptionMode, not even an explicit nil
func (o *RemoteCluster) UnsetSupportedAesEncryptionMode() {
	o.SupportedAesEncryptionMode.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *RemoteCluster) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *RemoteCluster) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *RemoteCluster) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *RemoteCluster) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTenantStorageDomainSharingEnabled returns the TenantStorageDomainSharingEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetTenantStorageDomainSharingEnabled() bool {
	if o == nil || IsNil(o.TenantStorageDomainSharingEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TenantStorageDomainSharingEnabled.Get()
}

// GetTenantStorageDomainSharingEnabledOk returns a tuple with the TenantStorageDomainSharingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetTenantStorageDomainSharingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantStorageDomainSharingEnabled.Get(), o.TenantStorageDomainSharingEnabled.IsSet()
}

// HasTenantStorageDomainSharingEnabled returns a boolean if a field has been set.
func (o *RemoteCluster) HasTenantStorageDomainSharingEnabled() bool {
	if o != nil && o.TenantStorageDomainSharingEnabled.IsSet() {
		return true
	}

	return false
}

// SetTenantStorageDomainSharingEnabled gets a reference to the given NullableBool and assigns it to the TenantStorageDomainSharingEnabled field.
func (o *RemoteCluster) SetTenantStorageDomainSharingEnabled(v bool) {
	o.TenantStorageDomainSharingEnabled.Set(&v)
}
// SetTenantStorageDomainSharingEnabledNil sets the value for TenantStorageDomainSharingEnabled to be an explicit nil
func (o *RemoteCluster) SetTenantStorageDomainSharingEnabledNil() {
	o.TenantStorageDomainSharingEnabled.Set(nil)
}

// UnsetTenantStorageDomainSharingEnabled ensures that no value is present for TenantStorageDomainSharingEnabled, not even an explicit nil
func (o *RemoteCluster) UnsetTenantStorageDomainSharingEnabled() {
	o.TenantStorageDomainSharingEnabled.Unset()
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetTlsEnabled() bool {
	if o == nil || IsNil(o.TlsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TlsEnabled.Get()
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetTlsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsEnabled.Get(), o.TlsEnabled.IsSet()
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *RemoteCluster) HasTlsEnabled() bool {
	if o != nil && o.TlsEnabled.IsSet() {
		return true
	}

	return false
}

// SetTlsEnabled gets a reference to the given NullableBool and assigns it to the TlsEnabled field.
func (o *RemoteCluster) SetTlsEnabled(v bool) {
	o.TlsEnabled.Set(&v)
}
// SetTlsEnabledNil sets the value for TlsEnabled to be an explicit nil
func (o *RemoteCluster) SetTlsEnabledNil() {
	o.TlsEnabled.Set(nil)
}

// UnsetTlsEnabled ensures that no value is present for TlsEnabled, not even an explicit nil
func (o *RemoteCluster) UnsetTlsEnabled() {
	o.TlsEnabled.Unset()
}

// GetNodeAddresses returns the NodeAddresses field value if set, zero value otherwise.
func (o *RemoteCluster) GetNodeAddresses() []string {
	if o == nil || IsNil(o.NodeAddresses) {
		var ret []string
		return ret
	}
	return o.NodeAddresses
}

// GetNodeAddressesOk returns a tuple with the NodeAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetNodeAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeAddresses) {
		return nil, false
	}
	return o.NodeAddresses, true
}

// HasNodeAddresses returns a boolean if a field has been set.
func (o *RemoteCluster) HasNodeAddresses() bool {
	if o != nil && !IsNil(o.NodeAddresses) {
		return true
	}

	return false
}

// SetNodeAddresses gets a reference to the given []string and assigns it to the NodeAddresses field.
func (o *RemoteCluster) SetNodeAddresses(v []string) {
	o.NodeAddresses = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteCluster) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteCluster) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *RemoteCluster) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *RemoteCluster) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *RemoteCluster) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *RemoteCluster) UnsetPassword() {
	o.Password.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RemoteCluster) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteCluster) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RemoteCluster) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RemoteCluster) SetUsername(v string) {
	o.Username = &v
}

func (o RemoteCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRegisterTarget.IsSet() {
		toSerialize["autoRegisterTarget"] = o.AutoRegisterTarget.Get()
	}
	if o.ClusterId.IsSet() {
		toSerialize["clusterId"] = o.ClusterId.Get()
	}
	if o.ClusterIncarnationId.IsSet() {
		toSerialize["clusterIncarnationId"] = o.ClusterIncarnationId.Get()
	}
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EffectiveAesEncryptionMode.IsSet() {
		toSerialize["effectiveAesEncryptionMode"] = o.EffectiveAesEncryptionMode.Get()
	}
	if o.IsAutoRegistered.IsSet() {
		toSerialize["isAutoRegistered"] = o.IsAutoRegistered.Get()
	}
	if !IsNil(o.LocalAddresses) {
		toSerialize["localAddresses"] = o.LocalAddresses
	}
	if o.MultiTenancyEnabled.IsSet() {
		toSerialize["multiTenancyEnabled"] = o.MultiTenancyEnabled.Get()
	}
	if o.NetworkInterface.IsSet() {
		toSerialize["networkInterface"] = o.NetworkInterface.Get()
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.ReplicationParams) {
		toSerialize["replicationParams"] = o.ReplicationParams
	}
	if o.SupportedAesEncryptionMode.IsSet() {
		toSerialize["supportedAesEncryptionMode"] = o.SupportedAesEncryptionMode.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.TenantStorageDomainSharingEnabled.IsSet() {
		toSerialize["tenantStorageDomainSharingEnabled"] = o.TenantStorageDomainSharingEnabled.Get()
	}
	if o.TlsEnabled.IsSet() {
		toSerialize["tlsEnabled"] = o.TlsEnabled.Get()
	}
	if !IsNil(o.NodeAddresses) {
		toSerialize["nodeAddresses"] = o.NodeAddresses
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableRemoteCluster struct {
	value *RemoteCluster
	isSet bool
}

func (v NullableRemoteCluster) Get() *RemoteCluster {
	return v.value
}

func (v *NullableRemoteCluster) Set(val *RemoteCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteCluster(val *RemoteCluster) *NullableRemoteCluster {
	return &NullableRemoteCluster{value: val, isSet: true}
}

func (v NullableRemoteCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


