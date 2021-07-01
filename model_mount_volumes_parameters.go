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

// MountVolumesParameters Specifies the information required for mounting volumes. Only required for Restore Tasks of type 'kMountVolumes'. At a minimum, the targetSourceId must be specified for 'kMountVolumes' Restore Tasks. If only targetSourceId is specified, all disks are attached but may not be mounted. The mount target must be registered on the Cohesity Cluster. If the mount target is a VM, VMware Tools must be installed. If the mount target is a physical server, a Cohesity Agent must be be installed. See the Cohesity Dashboard help documentation for details. In the username and password fields, specify the credentials to access the mount target.
type MountVolumesParameters struct {
	// Optional setting that determines if the volumes are brought online on the mount target after attaching the disks. This field is only set for VMs. The Cohesity Cluster always attempts to mount Physical servers. If true and the mount target is a VM, to mount the volumes VMware Tools must be installed on the guest operating system of the VM and login credentials to the mount target must be specified. NOTE: If automount is configured for a Windows system, the volumes may be automatically brought online.
	BringDisksOnline NullableBool `json:"bringDisksOnline,omitempty"`
	// Specifies password of the username to access the target source.
	Password NullableString `json:"password,omitempty"`
	// Specifies the target Protection Source id where the volumes will be mounted. NOTE: The source that was backed up and the mount target must be the same type, for example if the source is a VMware VM, then the mount target must also be a VMware VM. The mount target must be registered on the Cohesity Cluster.
	TargetSourceId NullableInt64 `json:"targetSourceId,omitempty"`
	// Optional setting that determines if this will use an existing agent on the target vm to bring disks online.
	UseExistingAgent NullableBool `json:"useExistingAgent,omitempty"`
	// Specifies username to access the target source.
	Username NullableString `json:"username,omitempty"`
	// Array of Volume Names.  Optionally specify the names of volumes to mount. If none are specified, all volumes of the Server are mounted on the target. To get the names of the volumes, call the GET /public/restore/vms/volumesInformation operation.
	VolumeNames []string `json:"volumeNames,omitempty"`
}

// NewMountVolumesParameters instantiates a new MountVolumesParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMountVolumesParameters() *MountVolumesParameters {
	this := MountVolumesParameters{}
	return &this
}

// NewMountVolumesParametersWithDefaults instantiates a new MountVolumesParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountVolumesParametersWithDefaults() *MountVolumesParameters {
	this := MountVolumesParameters{}
	return &this
}

// GetBringDisksOnline returns the BringDisksOnline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetBringDisksOnline() bool {
	if o == nil || o.BringDisksOnline.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BringDisksOnline.Get()
}

// GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetBringDisksOnlineOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BringDisksOnline.Get(), o.BringDisksOnline.IsSet()
}

// HasBringDisksOnline returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasBringDisksOnline() bool {
	if o != nil && o.BringDisksOnline.IsSet() {
		return true
	}

	return false
}

// SetBringDisksOnline gets a reference to the given NullableBool and assigns it to the BringDisksOnline field.
func (o *MountVolumesParameters) SetBringDisksOnline(v bool) {
	o.BringDisksOnline.Set(&v)
}
// SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil
func (o *MountVolumesParameters) SetBringDisksOnlineNil() {
	o.BringDisksOnline.Set(nil)
}

// UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
func (o *MountVolumesParameters) UnsetBringDisksOnline() {
	o.BringDisksOnline.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *MountVolumesParameters) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *MountVolumesParameters) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *MountVolumesParameters) UnsetPassword() {
	o.Password.Unset()
}

// GetTargetSourceId returns the TargetSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetTargetSourceId() int64 {
	if o == nil || o.TargetSourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TargetSourceId.Get()
}

// GetTargetSourceIdOk returns a tuple with the TargetSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetTargetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetSourceId.Get(), o.TargetSourceId.IsSet()
}

// HasTargetSourceId returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasTargetSourceId() bool {
	if o != nil && o.TargetSourceId.IsSet() {
		return true
	}

	return false
}

// SetTargetSourceId gets a reference to the given NullableInt64 and assigns it to the TargetSourceId field.
func (o *MountVolumesParameters) SetTargetSourceId(v int64) {
	o.TargetSourceId.Set(&v)
}
// SetTargetSourceIdNil sets the value for TargetSourceId to be an explicit nil
func (o *MountVolumesParameters) SetTargetSourceIdNil() {
	o.TargetSourceId.Set(nil)
}

// UnsetTargetSourceId ensures that no value is present for TargetSourceId, not even an explicit nil
func (o *MountVolumesParameters) UnsetTargetSourceId() {
	o.TargetSourceId.Unset()
}

// GetUseExistingAgent returns the UseExistingAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetUseExistingAgent() bool {
	if o == nil || o.UseExistingAgent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseExistingAgent.Get()
}

// GetUseExistingAgentOk returns a tuple with the UseExistingAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetUseExistingAgentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseExistingAgent.Get(), o.UseExistingAgent.IsSet()
}

// HasUseExistingAgent returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasUseExistingAgent() bool {
	if o != nil && o.UseExistingAgent.IsSet() {
		return true
	}

	return false
}

// SetUseExistingAgent gets a reference to the given NullableBool and assigns it to the UseExistingAgent field.
func (o *MountVolumesParameters) SetUseExistingAgent(v bool) {
	o.UseExistingAgent.Set(&v)
}
// SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil
func (o *MountVolumesParameters) SetUseExistingAgentNil() {
	o.UseExistingAgent.Set(nil)
}

// UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
func (o *MountVolumesParameters) UnsetUseExistingAgent() {
	o.UseExistingAgent.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *MountVolumesParameters) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *MountVolumesParameters) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *MountVolumesParameters) UnsetUsername() {
	o.Username.Unset()
}

// GetVolumeNames returns the VolumeNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MountVolumesParameters) GetVolumeNames() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.VolumeNames
}

// GetVolumeNamesOk returns a tuple with the VolumeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MountVolumesParameters) GetVolumeNamesOk() (*[]string, bool) {
	if o == nil || o.VolumeNames == nil {
		return nil, false
	}
	return &o.VolumeNames, true
}

// HasVolumeNames returns a boolean if a field has been set.
func (o *MountVolumesParameters) HasVolumeNames() bool {
	if o != nil && o.VolumeNames != nil {
		return true
	}

	return false
}

// SetVolumeNames gets a reference to the given []string and assigns it to the VolumeNames field.
func (o *MountVolumesParameters) SetVolumeNames(v []string) {
	o.VolumeNames = v
}

func (o MountVolumesParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BringDisksOnline.IsSet() {
		toSerialize["bringDisksOnline"] = o.BringDisksOnline.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.TargetSourceId.IsSet() {
		toSerialize["targetSourceId"] = o.TargetSourceId.Get()
	}
	if o.UseExistingAgent.IsSet() {
		toSerialize["useExistingAgent"] = o.UseExistingAgent.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.VolumeNames != nil {
		toSerialize["volumeNames"] = o.VolumeNames
	}
	return json.Marshal(toSerialize)
}

type NullableMountVolumesParameters struct {
	value *MountVolumesParameters
	isSet bool
}

func (v NullableMountVolumesParameters) Get() *MountVolumesParameters {
	return v.value
}

func (v *NullableMountVolumesParameters) Set(val *MountVolumesParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableMountVolumesParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableMountVolumesParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountVolumesParameters(val *MountVolumesParameters) *NullableMountVolumesParameters {
	return &NullableMountVolumesParameters{value: val, isSet: true}
}

func (v NullableMountVolumesParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountVolumesParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


