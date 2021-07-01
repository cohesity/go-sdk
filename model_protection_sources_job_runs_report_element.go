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

// ProtectionSourcesJobRunsReportElement Specifies a Protection Source and the Snapshots that back it up.
type ProtectionSourcesJobRunsReportElement struct {
	// Specifies the leaf Protection Source Object such as a VM.
	ProtectionSource NullableProtectionSource `json:"protectionSource,omitempty"`
	// Array of Snapshots  Specifies the Snapshots that contain backups of the Protection Source Object.
	SnapshotsInfo []ProtectionSourceSnapshotInformation `json:"snapshotsInfo,omitempty"`
}

// NewProtectionSourcesJobRunsReportElement instantiates a new ProtectionSourcesJobRunsReportElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectionSourcesJobRunsReportElement() *ProtectionSourcesJobRunsReportElement {
	this := ProtectionSourcesJobRunsReportElement{}
	return &this
}

// NewProtectionSourcesJobRunsReportElementWithDefaults instantiates a new ProtectionSourcesJobRunsReportElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectionSourcesJobRunsReportElementWithDefaults() *ProtectionSourcesJobRunsReportElement {
	this := ProtectionSourcesJobRunsReportElement{}
	return &this
}

// GetProtectionSource returns the ProtectionSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSourcesJobRunsReportElement) GetProtectionSource() ProtectionSource {
	if o == nil || o.ProtectionSource.Get() == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.ProtectionSource.Get()
}

// GetProtectionSourceOk returns a tuple with the ProtectionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSourcesJobRunsReportElement) GetProtectionSourceOk() (*ProtectionSource, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionSource.Get(), o.ProtectionSource.IsSet()
}

// HasProtectionSource returns a boolean if a field has been set.
func (o *ProtectionSourcesJobRunsReportElement) HasProtectionSource() bool {
	if o != nil && o.ProtectionSource.IsSet() {
		return true
	}

	return false
}

// SetProtectionSource gets a reference to the given NullableProtectionSource and assigns it to the ProtectionSource field.
func (o *ProtectionSourcesJobRunsReportElement) SetProtectionSource(v ProtectionSource) {
	o.ProtectionSource.Set(&v)
}
// SetProtectionSourceNil sets the value for ProtectionSource to be an explicit nil
func (o *ProtectionSourcesJobRunsReportElement) SetProtectionSourceNil() {
	o.ProtectionSource.Set(nil)
}

// UnsetProtectionSource ensures that no value is present for ProtectionSource, not even an explicit nil
func (o *ProtectionSourcesJobRunsReportElement) UnsetProtectionSource() {
	o.ProtectionSource.Unset()
}

// GetSnapshotsInfo returns the SnapshotsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectionSourcesJobRunsReportElement) GetSnapshotsInfo() []ProtectionSourceSnapshotInformation {
	if o == nil  {
		var ret []ProtectionSourceSnapshotInformation
		return ret
	}
	return o.SnapshotsInfo
}

// GetSnapshotsInfoOk returns a tuple with the SnapshotsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectionSourcesJobRunsReportElement) GetSnapshotsInfoOk() (*[]ProtectionSourceSnapshotInformation, bool) {
	if o == nil || o.SnapshotsInfo == nil {
		return nil, false
	}
	return &o.SnapshotsInfo, true
}

// HasSnapshotsInfo returns a boolean if a field has been set.
func (o *ProtectionSourcesJobRunsReportElement) HasSnapshotsInfo() bool {
	if o != nil && o.SnapshotsInfo != nil {
		return true
	}

	return false
}

// SetSnapshotsInfo gets a reference to the given []ProtectionSourceSnapshotInformation and assigns it to the SnapshotsInfo field.
func (o *ProtectionSourcesJobRunsReportElement) SetSnapshotsInfo(v []ProtectionSourceSnapshotInformation) {
	o.SnapshotsInfo = v
}

func (o ProtectionSourcesJobRunsReportElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionSource.IsSet() {
		toSerialize["protectionSource"] = o.ProtectionSource.Get()
	}
	if o.SnapshotsInfo != nil {
		toSerialize["snapshotsInfo"] = o.SnapshotsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableProtectionSourcesJobRunsReportElement struct {
	value *ProtectionSourcesJobRunsReportElement
	isSet bool
}

func (v NullableProtectionSourcesJobRunsReportElement) Get() *ProtectionSourcesJobRunsReportElement {
	return v.value
}

func (v *NullableProtectionSourcesJobRunsReportElement) Set(val *ProtectionSourcesJobRunsReportElement) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionSourcesJobRunsReportElement) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionSourcesJobRunsReportElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionSourcesJobRunsReportElement(val *ProtectionSourcesJobRunsReportElement) *NullableProtectionSourcesJobRunsReportElement {
	return &NullableProtectionSourcesJobRunsReportElement{value: val, isSet: true}
}

func (v NullableProtectionSourcesJobRunsReportElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionSourcesJobRunsReportElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


