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

// PhysicalTargetParamsForMountVolume Specifies the parameters for a physical recovery target.
type PhysicalTargetParamsForMountVolume struct {
	// Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	MountToOriginalTarget NullableBool `json:"mountToOriginalTarget"`
	// Specifies the configuration for mounting to the original target.
	OriginalTargetConfig NullablePhysicalMountVolumesOriginalTargetConfig `json:"originalTargetConfig,omitempty"`
	// Specifies the configuration for mounting to a new target.
	NewTargetConfig NullablePhysicalMountVolumesNewTargetConfig `json:"newTargetConfig,omitempty"`
	// Specifies whether to perform a read-only mount. Default is false.
	ReadOnlyMount NullableBool `json:"readOnlyMount,omitempty"`
	// Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.
	VolumeNames []string `json:"volumeNames,omitempty"`
	// Specifies the mapping of original volumes and mounted volumes
	MountedVolumeMapping []MountedVolumeMapping `json:"mountedVolumeMapping,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewPhysicalTargetParamsForMountVolume instantiates a new PhysicalTargetParamsForMountVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalTargetParamsForMountVolume(mountToOriginalTarget NullableBool) *PhysicalTargetParamsForMountVolume {
	this := PhysicalTargetParamsForMountVolume{}
	this.MountToOriginalTarget = mountToOriginalTarget
	return &this
}

// NewPhysicalTargetParamsForMountVolumeWithDefaults instantiates a new PhysicalTargetParamsForMountVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalTargetParamsForMountVolumeWithDefaults() *PhysicalTargetParamsForMountVolume {
	this := PhysicalTargetParamsForMountVolume{}
	return &this
}

// GetMountToOriginalTarget returns the MountToOriginalTarget field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PhysicalTargetParamsForMountVolume) GetMountToOriginalTarget() bool {
	if o == nil || o.MountToOriginalTarget.Get() == nil {
		var ret bool
		return ret
	}

	return *o.MountToOriginalTarget.Get()
}

// GetMountToOriginalTargetOk returns a tuple with the MountToOriginalTarget field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetMountToOriginalTargetOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MountToOriginalTarget.Get(), o.MountToOriginalTarget.IsSet()
}

// SetMountToOriginalTarget sets field value
func (o *PhysicalTargetParamsForMountVolume) SetMountToOriginalTarget(v bool) {
	o.MountToOriginalTarget.Set(&v)
}

// GetOriginalTargetConfig returns the OriginalTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetOriginalTargetConfig() PhysicalMountVolumesOriginalTargetConfig {
	if o == nil || o.OriginalTargetConfig.Get() == nil {
		var ret PhysicalMountVolumesOriginalTargetConfig
		return ret
	}
	return *o.OriginalTargetConfig.Get()
}

// GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetOriginalTargetConfigOk() (*PhysicalMountVolumesOriginalTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalTargetConfig.Get(), o.OriginalTargetConfig.IsSet()
}

// HasOriginalTargetConfig returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasOriginalTargetConfig() bool {
	if o != nil && o.OriginalTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalTargetConfig gets a reference to the given NullablePhysicalMountVolumesOriginalTargetConfig and assigns it to the OriginalTargetConfig field.
func (o *PhysicalTargetParamsForMountVolume) SetOriginalTargetConfig(v PhysicalMountVolumesOriginalTargetConfig) {
	o.OriginalTargetConfig.Set(&v)
}
// SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil
func (o *PhysicalTargetParamsForMountVolume) SetOriginalTargetConfigNil() {
	o.OriginalTargetConfig.Set(nil)
}

// UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
func (o *PhysicalTargetParamsForMountVolume) UnsetOriginalTargetConfig() {
	o.OriginalTargetConfig.Unset()
}

// GetNewTargetConfig returns the NewTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetNewTargetConfig() PhysicalMountVolumesNewTargetConfig {
	if o == nil || o.NewTargetConfig.Get() == nil {
		var ret PhysicalMountVolumesNewTargetConfig
		return ret
	}
	return *o.NewTargetConfig.Get()
}

// GetNewTargetConfigOk returns a tuple with the NewTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetNewTargetConfigOk() (*PhysicalMountVolumesNewTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewTargetConfig.Get(), o.NewTargetConfig.IsSet()
}

// HasNewTargetConfig returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasNewTargetConfig() bool {
	if o != nil && o.NewTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetNewTargetConfig gets a reference to the given NullablePhysicalMountVolumesNewTargetConfig and assigns it to the NewTargetConfig field.
func (o *PhysicalTargetParamsForMountVolume) SetNewTargetConfig(v PhysicalMountVolumesNewTargetConfig) {
	o.NewTargetConfig.Set(&v)
}
// SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil
func (o *PhysicalTargetParamsForMountVolume) SetNewTargetConfigNil() {
	o.NewTargetConfig.Set(nil)
}

// UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
func (o *PhysicalTargetParamsForMountVolume) UnsetNewTargetConfig() {
	o.NewTargetConfig.Unset()
}

// GetReadOnlyMount returns the ReadOnlyMount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetReadOnlyMount() bool {
	if o == nil || o.ReadOnlyMount.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnlyMount.Get()
}

// GetReadOnlyMountOk returns a tuple with the ReadOnlyMount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetReadOnlyMountOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReadOnlyMount.Get(), o.ReadOnlyMount.IsSet()
}

// HasReadOnlyMount returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasReadOnlyMount() bool {
	if o != nil && o.ReadOnlyMount.IsSet() {
		return true
	}

	return false
}

// SetReadOnlyMount gets a reference to the given NullableBool and assigns it to the ReadOnlyMount field.
func (o *PhysicalTargetParamsForMountVolume) SetReadOnlyMount(v bool) {
	o.ReadOnlyMount.Set(&v)
}
// SetReadOnlyMountNil sets the value for ReadOnlyMount to be an explicit nil
func (o *PhysicalTargetParamsForMountVolume) SetReadOnlyMountNil() {
	o.ReadOnlyMount.Set(nil)
}

// UnsetReadOnlyMount ensures that no value is present for ReadOnlyMount, not even an explicit nil
func (o *PhysicalTargetParamsForMountVolume) UnsetReadOnlyMount() {
	o.ReadOnlyMount.Unset()
}

// GetVolumeNames returns the VolumeNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetVolumeNames() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.VolumeNames
}

// GetVolumeNamesOk returns a tuple with the VolumeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetVolumeNamesOk() (*[]string, bool) {
	if o == nil || o.VolumeNames == nil {
		return nil, false
	}
	return &o.VolumeNames, true
}

// HasVolumeNames returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasVolumeNames() bool {
	if o != nil && o.VolumeNames != nil {
		return true
	}

	return false
}

// SetVolumeNames gets a reference to the given []string and assigns it to the VolumeNames field.
func (o *PhysicalTargetParamsForMountVolume) SetVolumeNames(v []string) {
	o.VolumeNames = v
}

// GetMountedVolumeMapping returns the MountedVolumeMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetMountedVolumeMapping() []MountedVolumeMapping {
	if o == nil  {
		var ret []MountedVolumeMapping
		return ret
	}
	return o.MountedVolumeMapping
}

// GetMountedVolumeMappingOk returns a tuple with the MountedVolumeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetMountedVolumeMappingOk() (*[]MountedVolumeMapping, bool) {
	if o == nil || o.MountedVolumeMapping == nil {
		return nil, false
	}
	return &o.MountedVolumeMapping, true
}

// HasMountedVolumeMapping returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasMountedVolumeMapping() bool {
	if o != nil && o.MountedVolumeMapping != nil {
		return true
	}

	return false
}

// SetMountedVolumeMapping gets a reference to the given []MountedVolumeMapping and assigns it to the MountedVolumeMapping field.
func (o *PhysicalTargetParamsForMountVolume) SetMountedVolumeMapping(v []MountedVolumeMapping) {
	o.MountedVolumeMapping = v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForMountVolume) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForMountVolume) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForMountVolume) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *PhysicalTargetParamsForMountVolume) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *PhysicalTargetParamsForMountVolume) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *PhysicalTargetParamsForMountVolume) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o PhysicalTargetParamsForMountVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mountToOriginalTarget"] = o.MountToOriginalTarget.Get()
	}
	if o.OriginalTargetConfig.IsSet() {
		toSerialize["originalTargetConfig"] = o.OriginalTargetConfig.Get()
	}
	if o.NewTargetConfig.IsSet() {
		toSerialize["newTargetConfig"] = o.NewTargetConfig.Get()
	}
	if o.ReadOnlyMount.IsSet() {
		toSerialize["readOnlyMount"] = o.ReadOnlyMount.Get()
	}
	if o.VolumeNames != nil {
		toSerialize["volumeNames"] = o.VolumeNames
	}
	if o.MountedVolumeMapping != nil {
		toSerialize["mountedVolumeMapping"] = o.MountedVolumeMapping
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalTargetParamsForMountVolume struct {
	value *PhysicalTargetParamsForMountVolume
	isSet bool
}

func (v NullablePhysicalTargetParamsForMountVolume) Get() *PhysicalTargetParamsForMountVolume {
	return v.value
}

func (v *NullablePhysicalTargetParamsForMountVolume) Set(val *PhysicalTargetParamsForMountVolume) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalTargetParamsForMountVolume) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalTargetParamsForMountVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalTargetParamsForMountVolume(val *PhysicalTargetParamsForMountVolume) *NullablePhysicalTargetParamsForMountVolume {
	return &NullablePhysicalTargetParamsForMountVolume{value: val, isSet: true}
}

func (v NullablePhysicalTargetParamsForMountVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalTargetParamsForMountVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


