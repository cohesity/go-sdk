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

// IsilonProtectionSource Specifies a Protection Source in Isilon OneFs environment.
type IsilonProtectionSource struct {
	AccessZone *IsilonAccessZone `json:"accessZone,omitempty"`
	Cluster *IsilonCluster `json:"cluster,omitempty"`
	MountPoint *IsilonMountPoint `json:"mountPoint,omitempty"`
	// Specifies a unique name of the Protection Source.
	Name NullableString `json:"name,omitempty"`
	// Specifies the type of the entity in an Isilon OneFs file system like 'kCluster', 'kZone', or, 'kMountPoint'. 'kCluster' indicates an Isilon OneFs Cluster. 'kZone' indicates an access zone in an Isilon OneFs Cluster. 'kMountPoint' indicates a mount point exposed by an Isilon OneFs Cluster.
	Type NullableString `json:"type,omitempty"`
}

// NewIsilonProtectionSource instantiates a new IsilonProtectionSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsilonProtectionSource() *IsilonProtectionSource {
	this := IsilonProtectionSource{}
	return &this
}

// NewIsilonProtectionSourceWithDefaults instantiates a new IsilonProtectionSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsilonProtectionSourceWithDefaults() *IsilonProtectionSource {
	this := IsilonProtectionSource{}
	return &this
}

// GetAccessZone returns the AccessZone field value if set, zero value otherwise.
func (o *IsilonProtectionSource) GetAccessZone() IsilonAccessZone {
	if o == nil || o.AccessZone == nil {
		var ret IsilonAccessZone
		return ret
	}
	return *o.AccessZone
}

// GetAccessZoneOk returns a tuple with the AccessZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonProtectionSource) GetAccessZoneOk() (*IsilonAccessZone, bool) {
	if o == nil || o.AccessZone == nil {
		return nil, false
	}
	return o.AccessZone, true
}

// HasAccessZone returns a boolean if a field has been set.
func (o *IsilonProtectionSource) HasAccessZone() bool {
	if o != nil && o.AccessZone != nil {
		return true
	}

	return false
}

// SetAccessZone gets a reference to the given IsilonAccessZone and assigns it to the AccessZone field.
func (o *IsilonProtectionSource) SetAccessZone(v IsilonAccessZone) {
	o.AccessZone = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *IsilonProtectionSource) GetCluster() IsilonCluster {
	if o == nil || o.Cluster == nil {
		var ret IsilonCluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonProtectionSource) GetClusterOk() (*IsilonCluster, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *IsilonProtectionSource) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given IsilonCluster and assigns it to the Cluster field.
func (o *IsilonProtectionSource) SetCluster(v IsilonCluster) {
	o.Cluster = &v
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise.
func (o *IsilonProtectionSource) GetMountPoint() IsilonMountPoint {
	if o == nil || o.MountPoint == nil {
		var ret IsilonMountPoint
		return ret
	}
	return *o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsilonProtectionSource) GetMountPointOk() (*IsilonMountPoint, bool) {
	if o == nil || o.MountPoint == nil {
		return nil, false
	}
	return o.MountPoint, true
}

// HasMountPoint returns a boolean if a field has been set.
func (o *IsilonProtectionSource) HasMountPoint() bool {
	if o != nil && o.MountPoint != nil {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given IsilonMountPoint and assigns it to the MountPoint field.
func (o *IsilonProtectionSource) SetMountPoint(v IsilonMountPoint) {
	o.MountPoint = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonProtectionSource) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonProtectionSource) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IsilonProtectionSource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IsilonProtectionSource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IsilonProtectionSource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IsilonProtectionSource) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonProtectionSource) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonProtectionSource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *IsilonProtectionSource) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *IsilonProtectionSource) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *IsilonProtectionSource) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *IsilonProtectionSource) UnsetType() {
	o.Type.Unset()
}

func (o IsilonProtectionSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessZone != nil {
		toSerialize["accessZone"] = o.AccessZone
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.MountPoint != nil {
		toSerialize["mountPoint"] = o.MountPoint
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIsilonProtectionSource struct {
	value *IsilonProtectionSource
	isSet bool
}

func (v NullableIsilonProtectionSource) Get() *IsilonProtectionSource {
	return v.value
}

func (v *NullableIsilonProtectionSource) Set(val *IsilonProtectionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIsilonProtectionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIsilonProtectionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsilonProtectionSource(val *IsilonProtectionSource) *NullableIsilonProtectionSource {
	return &NullableIsilonProtectionSource{value: val, isSet: true}
}

func (v NullableIsilonProtectionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsilonProtectionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


