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

// VirtualDiskInformation struct for VirtualDiskInformation
type VirtualDiskInformation struct {
	// Specifies the Id of the controller bus that controls the disk.
	BusNumber NullableInt64 `json:"busNumber,omitempty"`
	// Specifies the controller type like SCSI, or IDE etc.
	ControllerType NullableString `json:"controllerType,omitempty"`
	// Specifies original disk id. This is sufficient to identify the disk information, but in some scenarios, users may specify the controller option instead.
	DiskId NullableString `json:"diskId,omitempty"`
	DiskLocation *ProtectionSource `json:"diskLocation,omitempty"`
	// Specifies size of the virtual disk in bytes.
	DiskSizeInBytes NullableInt64 `json:"diskSizeInBytes,omitempty"`
	// Specifies the original file path if applicable.
	FilePath NullableString `json:"filePath,omitempty"`
	// Specifies the list of mount points.
	MountPoints []string `json:"mountPoints,omitempty"`
	// Specifies the disk file name. This is the VMDK name and not the flat file name.
	UnitNumber NullableInt64 `json:"unitNumber,omitempty"`
}

// NewVirtualDiskInformation instantiates a new VirtualDiskInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualDiskInformation() *VirtualDiskInformation {
	this := VirtualDiskInformation{}
	return &this
}

// NewVirtualDiskInformationWithDefaults instantiates a new VirtualDiskInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualDiskInformationWithDefaults() *VirtualDiskInformation {
	this := VirtualDiskInformation{}
	return &this
}

// GetBusNumber returns the BusNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetBusNumber() int64 {
	if o == nil || o.BusNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BusNumber.Get()
}

// GetBusNumberOk returns a tuple with the BusNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetBusNumberOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BusNumber.Get(), o.BusNumber.IsSet()
}

// HasBusNumber returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasBusNumber() bool {
	if o != nil && o.BusNumber.IsSet() {
		return true
	}

	return false
}

// SetBusNumber gets a reference to the given NullableInt64 and assigns it to the BusNumber field.
func (o *VirtualDiskInformation) SetBusNumber(v int64) {
	o.BusNumber.Set(&v)
}
// SetBusNumberNil sets the value for BusNumber to be an explicit nil
func (o *VirtualDiskInformation) SetBusNumberNil() {
	o.BusNumber.Set(nil)
}

// UnsetBusNumber ensures that no value is present for BusNumber, not even an explicit nil
func (o *VirtualDiskInformation) UnsetBusNumber() {
	o.BusNumber.Unset()
}

// GetControllerType returns the ControllerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetControllerType() string {
	if o == nil || o.ControllerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControllerType.Get()
}

// GetControllerTypeOk returns a tuple with the ControllerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetControllerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ControllerType.Get(), o.ControllerType.IsSet()
}

// HasControllerType returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasControllerType() bool {
	if o != nil && o.ControllerType.IsSet() {
		return true
	}

	return false
}

// SetControllerType gets a reference to the given NullableString and assigns it to the ControllerType field.
func (o *VirtualDiskInformation) SetControllerType(v string) {
	o.ControllerType.Set(&v)
}
// SetControllerTypeNil sets the value for ControllerType to be an explicit nil
func (o *VirtualDiskInformation) SetControllerTypeNil() {
	o.ControllerType.Set(nil)
}

// UnsetControllerType ensures that no value is present for ControllerType, not even an explicit nil
func (o *VirtualDiskInformation) UnsetControllerType() {
	o.ControllerType.Unset()
}

// GetDiskId returns the DiskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetDiskId() string {
	if o == nil || o.DiskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DiskId.Get()
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetDiskIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiskId.Get(), o.DiskId.IsSet()
}

// HasDiskId returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasDiskId() bool {
	if o != nil && o.DiskId.IsSet() {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given NullableString and assigns it to the DiskId field.
func (o *VirtualDiskInformation) SetDiskId(v string) {
	o.DiskId.Set(&v)
}
// SetDiskIdNil sets the value for DiskId to be an explicit nil
func (o *VirtualDiskInformation) SetDiskIdNil() {
	o.DiskId.Set(nil)
}

// UnsetDiskId ensures that no value is present for DiskId, not even an explicit nil
func (o *VirtualDiskInformation) UnsetDiskId() {
	o.DiskId.Unset()
}

// GetDiskLocation returns the DiskLocation field value if set, zero value otherwise.
func (o *VirtualDiskInformation) GetDiskLocation() ProtectionSource {
	if o == nil || o.DiskLocation == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.DiskLocation
}

// GetDiskLocationOk returns a tuple with the DiskLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDiskInformation) GetDiskLocationOk() (*ProtectionSource, bool) {
	if o == nil || o.DiskLocation == nil {
		return nil, false
	}
	return o.DiskLocation, true
}

// HasDiskLocation returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasDiskLocation() bool {
	if o != nil && o.DiskLocation != nil {
		return true
	}

	return false
}

// SetDiskLocation gets a reference to the given ProtectionSource and assigns it to the DiskLocation field.
func (o *VirtualDiskInformation) SetDiskLocation(v ProtectionSource) {
	o.DiskLocation = &v
}

// GetDiskSizeInBytes returns the DiskSizeInBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetDiskSizeInBytes() int64 {
	if o == nil || o.DiskSizeInBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DiskSizeInBytes.Get()
}

// GetDiskSizeInBytesOk returns a tuple with the DiskSizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetDiskSizeInBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiskSizeInBytes.Get(), o.DiskSizeInBytes.IsSet()
}

// HasDiskSizeInBytes returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasDiskSizeInBytes() bool {
	if o != nil && o.DiskSizeInBytes.IsSet() {
		return true
	}

	return false
}

// SetDiskSizeInBytes gets a reference to the given NullableInt64 and assigns it to the DiskSizeInBytes field.
func (o *VirtualDiskInformation) SetDiskSizeInBytes(v int64) {
	o.DiskSizeInBytes.Set(&v)
}
// SetDiskSizeInBytesNil sets the value for DiskSizeInBytes to be an explicit nil
func (o *VirtualDiskInformation) SetDiskSizeInBytesNil() {
	o.DiskSizeInBytes.Set(nil)
}

// UnsetDiskSizeInBytes ensures that no value is present for DiskSizeInBytes, not even an explicit nil
func (o *VirtualDiskInformation) UnsetDiskSizeInBytes() {
	o.DiskSizeInBytes.Unset()
}

// GetFilePath returns the FilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetFilePath() string {
	if o == nil || o.FilePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetFilePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// HasFilePath returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasFilePath() bool {
	if o != nil && o.FilePath.IsSet() {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given NullableString and assigns it to the FilePath field.
func (o *VirtualDiskInformation) SetFilePath(v string) {
	o.FilePath.Set(&v)
}
// SetFilePathNil sets the value for FilePath to be an explicit nil
func (o *VirtualDiskInformation) SetFilePathNil() {
	o.FilePath.Set(nil)
}

// UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
func (o *VirtualDiskInformation) UnsetFilePath() {
	o.FilePath.Unset()
}

// GetMountPoints returns the MountPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetMountPoints() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.MountPoints
}

// GetMountPointsOk returns a tuple with the MountPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetMountPointsOk() (*[]string, bool) {
	if o == nil || o.MountPoints == nil {
		return nil, false
	}
	return &o.MountPoints, true
}

// HasMountPoints returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasMountPoints() bool {
	if o != nil && o.MountPoints != nil {
		return true
	}

	return false
}

// SetMountPoints gets a reference to the given []string and assigns it to the MountPoints field.
func (o *VirtualDiskInformation) SetMountPoints(v []string) {
	o.MountPoints = v
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualDiskInformation) GetUnitNumber() int64 {
	if o == nil || o.UnitNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UnitNumber.Get()
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualDiskInformation) GetUnitNumberOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitNumber.Get(), o.UnitNumber.IsSet()
}

// HasUnitNumber returns a boolean if a field has been set.
func (o *VirtualDiskInformation) HasUnitNumber() bool {
	if o != nil && o.UnitNumber.IsSet() {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given NullableInt64 and assigns it to the UnitNumber field.
func (o *VirtualDiskInformation) SetUnitNumber(v int64) {
	o.UnitNumber.Set(&v)
}
// SetUnitNumberNil sets the value for UnitNumber to be an explicit nil
func (o *VirtualDiskInformation) SetUnitNumberNil() {
	o.UnitNumber.Set(nil)
}

// UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
func (o *VirtualDiskInformation) UnsetUnitNumber() {
	o.UnitNumber.Unset()
}

func (o VirtualDiskInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusNumber.IsSet() {
		toSerialize["busNumber"] = o.BusNumber.Get()
	}
	if o.ControllerType.IsSet() {
		toSerialize["controllerType"] = o.ControllerType.Get()
	}
	if o.DiskId.IsSet() {
		toSerialize["diskId"] = o.DiskId.Get()
	}
	if o.DiskLocation != nil {
		toSerialize["diskLocation"] = o.DiskLocation
	}
	if o.DiskSizeInBytes.IsSet() {
		toSerialize["diskSizeInBytes"] = o.DiskSizeInBytes.Get()
	}
	if o.FilePath.IsSet() {
		toSerialize["filePath"] = o.FilePath.Get()
	}
	if o.MountPoints != nil {
		toSerialize["mountPoints"] = o.MountPoints
	}
	if o.UnitNumber.IsSet() {
		toSerialize["unitNumber"] = o.UnitNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualDiskInformation struct {
	value *VirtualDiskInformation
	isSet bool
}

func (v NullableVirtualDiskInformation) Get() *VirtualDiskInformation {
	return v.value
}

func (v *NullableVirtualDiskInformation) Set(val *VirtualDiskInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualDiskInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualDiskInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualDiskInformation(val *VirtualDiskInformation) *NullableVirtualDiskInformation {
	return &NullableVirtualDiskInformation{value: val, isSet: true}
}

func (v NullableVirtualDiskInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualDiskInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


