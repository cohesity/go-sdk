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

// checks if the VdcObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VdcObject{}

// VdcObject Specifies the details about VMware Virtual datacenter.
type VdcObject struct {
	// Specifies a list of VDC Catalogs.
	Catalogs []VdcCatalog `json:"catalogs,omitempty"`
	// Specifies a list of Organization VDC Networks.
	OrgNetworks []OrgVDCNetwork `json:"orgNetworks,omitempty"`
}

// NewVdcObject instantiates a new VdcObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVdcObject() *VdcObject {
	this := VdcObject{}
	return &this
}

// NewVdcObjectWithDefaults instantiates a new VdcObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdcObjectWithDefaults() *VdcObject {
	this := VdcObject{}
	return &this
}

// GetCatalogs returns the Catalogs field value if set, zero value otherwise.
func (o *VdcObject) GetCatalogs() []VdcCatalog {
	if o == nil || IsNil(o.Catalogs) {
		var ret []VdcCatalog
		return ret
	}
	return o.Catalogs
}

// GetCatalogsOk returns a tuple with the Catalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdcObject) GetCatalogsOk() ([]VdcCatalog, bool) {
	if o == nil || IsNil(o.Catalogs) {
		return nil, false
	}
	return o.Catalogs, true
}

// HasCatalogs returns a boolean if a field has been set.
func (o *VdcObject) HasCatalogs() bool {
	if o != nil && !IsNil(o.Catalogs) {
		return true
	}

	return false
}

// SetCatalogs gets a reference to the given []VdcCatalog and assigns it to the Catalogs field.
func (o *VdcObject) SetCatalogs(v []VdcCatalog) {
	o.Catalogs = v
}

// GetOrgNetworks returns the OrgNetworks field value if set, zero value otherwise.
func (o *VdcObject) GetOrgNetworks() []OrgVDCNetwork {
	if o == nil || IsNil(o.OrgNetworks) {
		var ret []OrgVDCNetwork
		return ret
	}
	return o.OrgNetworks
}

// GetOrgNetworksOk returns a tuple with the OrgNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdcObject) GetOrgNetworksOk() ([]OrgVDCNetwork, bool) {
	if o == nil || IsNil(o.OrgNetworks) {
		return nil, false
	}
	return o.OrgNetworks, true
}

// HasOrgNetworks returns a boolean if a field has been set.
func (o *VdcObject) HasOrgNetworks() bool {
	if o != nil && !IsNil(o.OrgNetworks) {
		return true
	}

	return false
}

// SetOrgNetworks gets a reference to the given []OrgVDCNetwork and assigns it to the OrgNetworks field.
func (o *VdcObject) SetOrgNetworks(v []OrgVDCNetwork) {
	o.OrgNetworks = v
}

func (o VdcObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VdcObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Catalogs) {
		toSerialize["catalogs"] = o.Catalogs
	}
	if !IsNil(o.OrgNetworks) {
		toSerialize["orgNetworks"] = o.OrgNetworks
	}
	return toSerialize, nil
}

type NullableVdcObject struct {
	value *VdcObject
	isSet bool
}

func (v NullableVdcObject) Get() *VdcObject {
	return v.value
}

func (v *NullableVdcObject) Set(val *VdcObject) {
	v.value = val
	v.isSet = true
}

func (v NullableVdcObject) IsSet() bool {
	return v.isSet
}

func (v *NullableVdcObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdcObject(val *VdcObject) *NullableVdcObject {
	return &NullableVdcObject{value: val, isSet: true}
}

func (v NullableVdcObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdcObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


