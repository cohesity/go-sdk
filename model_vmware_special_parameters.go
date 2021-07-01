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

// VmwareSpecialParameters Specifies additional special settings applicable for a Protection Source of 'kVMware' type in a Protection Job.
type VmwareSpecialParameters struct {
	ApplicationParameters *ApplicationParameters `json:"applicationParameters,omitempty"`
	// Specifies the list of Disks to be excluded from backing up. These disks are excluded from all Protection Sources in the Protection Job.
	ExcludedDisks []DiskUnit `json:"excludedDisks,omitempty"`
	// Specifies the administrator credentials to log in to the guest Windows system of a VM that hosts the Microsoft Exchange Server. If truncateExchangeLog is set to true and the specified source is a VM, administrator credentials to log in to the guest Windows system of the VM must be provided to truncate the logs. This field is only applicable to Sources in the kVMware environment.
	VmCredentials NullableCredentials `json:"vmCredentials,omitempty"`
}

// NewVmwareSpecialParameters instantiates a new VmwareSpecialParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareSpecialParameters() *VmwareSpecialParameters {
	this := VmwareSpecialParameters{}
	return &this
}

// NewVmwareSpecialParametersWithDefaults instantiates a new VmwareSpecialParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareSpecialParametersWithDefaults() *VmwareSpecialParameters {
	this := VmwareSpecialParameters{}
	return &this
}

// GetApplicationParameters returns the ApplicationParameters field value if set, zero value otherwise.
func (o *VmwareSpecialParameters) GetApplicationParameters() ApplicationParameters {
	if o == nil || o.ApplicationParameters == nil {
		var ret ApplicationParameters
		return ret
	}
	return *o.ApplicationParameters
}

// GetApplicationParametersOk returns a tuple with the ApplicationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareSpecialParameters) GetApplicationParametersOk() (*ApplicationParameters, bool) {
	if o == nil || o.ApplicationParameters == nil {
		return nil, false
	}
	return o.ApplicationParameters, true
}

// HasApplicationParameters returns a boolean if a field has been set.
func (o *VmwareSpecialParameters) HasApplicationParameters() bool {
	if o != nil && o.ApplicationParameters != nil {
		return true
	}

	return false
}

// SetApplicationParameters gets a reference to the given ApplicationParameters and assigns it to the ApplicationParameters field.
func (o *VmwareSpecialParameters) SetApplicationParameters(v ApplicationParameters) {
	o.ApplicationParameters = &v
}

// GetExcludedDisks returns the ExcludedDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareSpecialParameters) GetExcludedDisks() []DiskUnit {
	if o == nil  {
		var ret []DiskUnit
		return ret
	}
	return o.ExcludedDisks
}

// GetExcludedDisksOk returns a tuple with the ExcludedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareSpecialParameters) GetExcludedDisksOk() (*[]DiskUnit, bool) {
	if o == nil || o.ExcludedDisks == nil {
		return nil, false
	}
	return &o.ExcludedDisks, true
}

// HasExcludedDisks returns a boolean if a field has been set.
func (o *VmwareSpecialParameters) HasExcludedDisks() bool {
	if o != nil && o.ExcludedDisks != nil {
		return true
	}

	return false
}

// SetExcludedDisks gets a reference to the given []DiskUnit and assigns it to the ExcludedDisks field.
func (o *VmwareSpecialParameters) SetExcludedDisks(v []DiskUnit) {
	o.ExcludedDisks = v
}

// GetVmCredentials returns the VmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareSpecialParameters) GetVmCredentials() Credentials {
	if o == nil || o.VmCredentials.Get() == nil {
		var ret Credentials
		return ret
	}
	return *o.VmCredentials.Get()
}

// GetVmCredentialsOk returns a tuple with the VmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareSpecialParameters) GetVmCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VmCredentials.Get(), o.VmCredentials.IsSet()
}

// HasVmCredentials returns a boolean if a field has been set.
func (o *VmwareSpecialParameters) HasVmCredentials() bool {
	if o != nil && o.VmCredentials.IsSet() {
		return true
	}

	return false
}

// SetVmCredentials gets a reference to the given NullableCredentials and assigns it to the VmCredentials field.
func (o *VmwareSpecialParameters) SetVmCredentials(v Credentials) {
	o.VmCredentials.Set(&v)
}
// SetVmCredentialsNil sets the value for VmCredentials to be an explicit nil
func (o *VmwareSpecialParameters) SetVmCredentialsNil() {
	o.VmCredentials.Set(nil)
}

// UnsetVmCredentials ensures that no value is present for VmCredentials, not even an explicit nil
func (o *VmwareSpecialParameters) UnsetVmCredentials() {
	o.VmCredentials.Unset()
}

func (o VmwareSpecialParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationParameters != nil {
		toSerialize["applicationParameters"] = o.ApplicationParameters
	}
	if o.ExcludedDisks != nil {
		toSerialize["excludedDisks"] = o.ExcludedDisks
	}
	if o.VmCredentials.IsSet() {
		toSerialize["vmCredentials"] = o.VmCredentials.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareSpecialParameters struct {
	value *VmwareSpecialParameters
	isSet bool
}

func (v NullableVmwareSpecialParameters) Get() *VmwareSpecialParameters {
	return v.value
}

func (v *NullableVmwareSpecialParameters) Set(val *VmwareSpecialParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareSpecialParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareSpecialParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareSpecialParameters(val *VmwareSpecialParameters) *NullableVmwareSpecialParameters {
	return &NullableVmwareSpecialParameters{value: val, isSet: true}
}

func (v NullableVmwareSpecialParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareSpecialParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


