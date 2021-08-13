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

// VmwareProtectionGroupExtraParams Specifies the extra parameters which are specific to VMware object protection Group.
type VmwareProtectionGroupExtraParams struct {
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
	// Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds *[]int64 `json:"excludeObjectIds,omitempty"`
	// Array of Array of VM Tag Ids that Specify VMs to Protect. Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only 'Eng' VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the 'Eng' VM Tag id, 2221 is the 'East' VM Tag id and 3031 is the 'West' VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with 'Eng' and 'East' (an intersection). The outer array combines the list from the inner array with list of VMs tagged with 'West' (a union). The list of resulting VMs are protected by this Protection Group.
	VmTagIds [][]int64 `json:"vmTagIds,omitempty"`
	// Array of Arrays of VM Tag Ids that Specify VMs to Exclude. Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the 'Former Employees' VMs in the East and West but keep all the VMs for 'Former Employees' in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the 'Former Employee' VM Tag id, 2221 is the 'East' VM Tag id and 3031 is the 'West' VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with 'Former Employees' and 'East' (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with 'Former Employees' and 'West' (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job.
	ExcludeVmTagIds *[][]int64 `json:"excludeVmTagIds,omitempty"`
	// Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying these filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
	ExcludeFilters []VMFilter `json:"excludeFilters,omitempty"`
	// Whether to leverage the storage array based snapshots for this backup. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, backup will fallback to the default backup method.
	LeverageStorageSnapshots NullableBool `json:"leverageStorageSnapshots,omitempty"`
	// Whether to leverage the hyperflex based snapshots for this backup. To leverage hyperflex snapshots, it has to first be registered. If hyperflex based snapshots cannot be taken, backup will fallback to the default backup method.
	LeverageHyperflexSnapshots NullableBool `json:"leverageHyperflexSnapshots,omitempty"`
	// Whether to leverage the nutanix based snapshots for this backup. To leverage nutanix snapshots, it has to first be registered. If nutanix based snapshots cannot be taken, backup will fallback to the default backup method.
	LeverageNutanixSnapshots NullableBool `json:"leverageNutanixSnapshots,omitempty"`
	// Specifies whether or not to move the workload to the cloud.
	CloudMigration NullableBool `json:"cloudMigration,omitempty"`
}

// NewVmwareProtectionGroupExtraParams instantiates a new VmwareProtectionGroupExtraParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareProtectionGroupExtraParams() *VmwareProtectionGroupExtraParams {
	this := VmwareProtectionGroupExtraParams{}
	return &this
}

// NewVmwareProtectionGroupExtraParamsWithDefaults instantiates a new VmwareProtectionGroupExtraParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareProtectionGroupExtraParamsWithDefaults() *VmwareProtectionGroupExtraParams {
	this := VmwareProtectionGroupExtraParams{}
	return &this
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *VmwareProtectionGroupExtraParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetSourceName() string {
	if o == nil || o.SourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *VmwareProtectionGroupExtraParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetSourceName() {
	o.SourceName.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *VmwareProtectionGroupExtraParams) GetExcludeObjectIds() []int64 {
	if o == nil || o.ExcludeObjectIds == nil {
		var ret []int64
		return ret
	}
	return *o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareProtectionGroupExtraParams) GetExcludeObjectIdsOk() (*[]int64, bool) {
	if o == nil || o.ExcludeObjectIds == nil {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasExcludeObjectIds() bool {
	if o != nil && o.ExcludeObjectIds != nil {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *VmwareProtectionGroupExtraParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = &v
}

// GetVmTagIds returns the VmTagIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetVmTagIds() [][]int64 {
	if o == nil  {
		var ret [][]int64
		return ret
	}
	return o.VmTagIds
}

// GetVmTagIdsOk returns a tuple with the VmTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetVmTagIdsOk() (*[][]int64, bool) {
	if o == nil || o.VmTagIds == nil {
		return nil, false
	}
	return &o.VmTagIds, true
}

// HasVmTagIds returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasVmTagIds() bool {
	if o != nil && o.VmTagIds != nil {
		return true
	}

	return false
}

// SetVmTagIds gets a reference to the given [][]int64 and assigns it to the VmTagIds field.
func (o *VmwareProtectionGroupExtraParams) SetVmTagIds(v [][]int64) {
	o.VmTagIds = v
}

// GetExcludeVmTagIds returns the ExcludeVmTagIds field value if set, zero value otherwise.
func (o *VmwareProtectionGroupExtraParams) GetExcludeVmTagIds() [][]int64 {
	if o == nil || o.ExcludeVmTagIds == nil {
		var ret [][]int64
		return ret
	}
	return *o.ExcludeVmTagIds
}

// GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareProtectionGroupExtraParams) GetExcludeVmTagIdsOk() (*[][]int64, bool) {
	if o == nil || o.ExcludeVmTagIds == nil {
		return nil, false
	}
	return o.ExcludeVmTagIds, true
}

// HasExcludeVmTagIds returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasExcludeVmTagIds() bool {
	if o != nil && o.ExcludeVmTagIds != nil {
		return true
	}

	return false
}

// SetExcludeVmTagIds gets a reference to the given [][]int64 and assigns it to the ExcludeVmTagIds field.
func (o *VmwareProtectionGroupExtraParams) SetExcludeVmTagIds(v [][]int64) {
	o.ExcludeVmTagIds = &v
}

// GetExcludeFilters returns the ExcludeFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetExcludeFilters() []VMFilter {
	if o == nil  {
		var ret []VMFilter
		return ret
	}
	return o.ExcludeFilters
}

// GetExcludeFiltersOk returns a tuple with the ExcludeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetExcludeFiltersOk() (*[]VMFilter, bool) {
	if o == nil || o.ExcludeFilters == nil {
		return nil, false
	}
	return &o.ExcludeFilters, true
}

// HasExcludeFilters returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasExcludeFilters() bool {
	if o != nil && o.ExcludeFilters != nil {
		return true
	}

	return false
}

// SetExcludeFilters gets a reference to the given []VMFilter and assigns it to the ExcludeFilters field.
func (o *VmwareProtectionGroupExtraParams) SetExcludeFilters(v []VMFilter) {
	o.ExcludeFilters = v
}

// GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetLeverageStorageSnapshots() bool {
	if o == nil || o.LeverageStorageSnapshots.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LeverageStorageSnapshots.Get()
}

// GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetLeverageStorageSnapshotsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeverageStorageSnapshots.Get(), o.LeverageStorageSnapshots.IsSet()
}

// HasLeverageStorageSnapshots returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasLeverageStorageSnapshots() bool {
	if o != nil && o.LeverageStorageSnapshots.IsSet() {
		return true
	}

	return false
}

// SetLeverageStorageSnapshots gets a reference to the given NullableBool and assigns it to the LeverageStorageSnapshots field.
func (o *VmwareProtectionGroupExtraParams) SetLeverageStorageSnapshots(v bool) {
	o.LeverageStorageSnapshots.Set(&v)
}
// SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetLeverageStorageSnapshotsNil() {
	o.LeverageStorageSnapshots.Set(nil)
}

// UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetLeverageStorageSnapshots() {
	o.LeverageStorageSnapshots.Unset()
}

// GetLeverageHyperflexSnapshots returns the LeverageHyperflexSnapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetLeverageHyperflexSnapshots() bool {
	if o == nil || o.LeverageHyperflexSnapshots.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LeverageHyperflexSnapshots.Get()
}

// GetLeverageHyperflexSnapshotsOk returns a tuple with the LeverageHyperflexSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetLeverageHyperflexSnapshotsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeverageHyperflexSnapshots.Get(), o.LeverageHyperflexSnapshots.IsSet()
}

// HasLeverageHyperflexSnapshots returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasLeverageHyperflexSnapshots() bool {
	if o != nil && o.LeverageHyperflexSnapshots.IsSet() {
		return true
	}

	return false
}

// SetLeverageHyperflexSnapshots gets a reference to the given NullableBool and assigns it to the LeverageHyperflexSnapshots field.
func (o *VmwareProtectionGroupExtraParams) SetLeverageHyperflexSnapshots(v bool) {
	o.LeverageHyperflexSnapshots.Set(&v)
}
// SetLeverageHyperflexSnapshotsNil sets the value for LeverageHyperflexSnapshots to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetLeverageHyperflexSnapshotsNil() {
	o.LeverageHyperflexSnapshots.Set(nil)
}

// UnsetLeverageHyperflexSnapshots ensures that no value is present for LeverageHyperflexSnapshots, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetLeverageHyperflexSnapshots() {
	o.LeverageHyperflexSnapshots.Unset()
}

// GetLeverageNutanixSnapshots returns the LeverageNutanixSnapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetLeverageNutanixSnapshots() bool {
	if o == nil || o.LeverageNutanixSnapshots.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LeverageNutanixSnapshots.Get()
}

// GetLeverageNutanixSnapshotsOk returns a tuple with the LeverageNutanixSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetLeverageNutanixSnapshotsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeverageNutanixSnapshots.Get(), o.LeverageNutanixSnapshots.IsSet()
}

// HasLeverageNutanixSnapshots returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasLeverageNutanixSnapshots() bool {
	if o != nil && o.LeverageNutanixSnapshots.IsSet() {
		return true
	}

	return false
}

// SetLeverageNutanixSnapshots gets a reference to the given NullableBool and assigns it to the LeverageNutanixSnapshots field.
func (o *VmwareProtectionGroupExtraParams) SetLeverageNutanixSnapshots(v bool) {
	o.LeverageNutanixSnapshots.Set(&v)
}
// SetLeverageNutanixSnapshotsNil sets the value for LeverageNutanixSnapshots to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetLeverageNutanixSnapshotsNil() {
	o.LeverageNutanixSnapshots.Set(nil)
}

// UnsetLeverageNutanixSnapshots ensures that no value is present for LeverageNutanixSnapshots, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetLeverageNutanixSnapshots() {
	o.LeverageNutanixSnapshots.Unset()
}

// GetCloudMigration returns the CloudMigration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupExtraParams) GetCloudMigration() bool {
	if o == nil || o.CloudMigration.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CloudMigration.Get()
}

// GetCloudMigrationOk returns a tuple with the CloudMigration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupExtraParams) GetCloudMigrationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloudMigration.Get(), o.CloudMigration.IsSet()
}

// HasCloudMigration returns a boolean if a field has been set.
func (o *VmwareProtectionGroupExtraParams) HasCloudMigration() bool {
	if o != nil && o.CloudMigration.IsSet() {
		return true
	}

	return false
}

// SetCloudMigration gets a reference to the given NullableBool and assigns it to the CloudMigration field.
func (o *VmwareProtectionGroupExtraParams) SetCloudMigration(v bool) {
	o.CloudMigration.Set(&v)
}
// SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil
func (o *VmwareProtectionGroupExtraParams) SetCloudMigrationNil() {
	o.CloudMigration.Set(nil)
}

// UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
func (o *VmwareProtectionGroupExtraParams) UnsetCloudMigration() {
	o.CloudMigration.Unset()
}

func (o VmwareProtectionGroupExtraParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.VmTagIds != nil {
		toSerialize["vmTagIds"] = o.VmTagIds
	}
	if o.ExcludeVmTagIds != nil {
		toSerialize["excludeVmTagIds"] = o.ExcludeVmTagIds
	}
	if o.ExcludeFilters != nil {
		toSerialize["excludeFilters"] = o.ExcludeFilters
	}
	if o.LeverageStorageSnapshots.IsSet() {
		toSerialize["leverageStorageSnapshots"] = o.LeverageStorageSnapshots.Get()
	}
	if o.LeverageHyperflexSnapshots.IsSet() {
		toSerialize["leverageHyperflexSnapshots"] = o.LeverageHyperflexSnapshots.Get()
	}
	if o.LeverageNutanixSnapshots.IsSet() {
		toSerialize["leverageNutanixSnapshots"] = o.LeverageNutanixSnapshots.Get()
	}
	if o.CloudMigration.IsSet() {
		toSerialize["cloudMigration"] = o.CloudMigration.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareProtectionGroupExtraParams struct {
	value *VmwareProtectionGroupExtraParams
	isSet bool
}

func (v NullableVmwareProtectionGroupExtraParams) Get() *VmwareProtectionGroupExtraParams {
	return v.value
}

func (v *NullableVmwareProtectionGroupExtraParams) Set(val *VmwareProtectionGroupExtraParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareProtectionGroupExtraParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareProtectionGroupExtraParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareProtectionGroupExtraParams(val *VmwareProtectionGroupExtraParams) *NullableVmwareProtectionGroupExtraParams {
	return &NullableVmwareProtectionGroupExtraParams{value: val, isSet: true}
}

func (v NullableVmwareProtectionGroupExtraParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareProtectionGroupExtraParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o VmwareProtectionGroupExtraParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}