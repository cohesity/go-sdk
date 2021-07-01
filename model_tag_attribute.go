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

// TagAttribute Specifies a VMware tag.
type TagAttribute struct {
	// Specifies the tag attribute type of GCP. Going forward, there will be an additional TagTypes for AWS as well. Specifies the type of the tag GCP entity refers to. 'kNetworkTag' indicates a network tag present on instances. 'kLabel' indicates a label (key-value pair) present on instances. 'kCustomMetadata' indicates custom Metadata (key-value pair) present on instances.
	GcpTagType NullableString `json:"gcpTagType,omitempty"`
	// Specifies the Coheisty id of the VM tag.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the VMware name of the VM tag.
	Name NullableString `json:"name,omitempty"`
	// Specifies the VMware Universally Unique Identifier (UUID) of the VM tag.
	Uuid NullableString `json:"uuid,omitempty"`
}

// NewTagAttribute instantiates a new TagAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAttribute() *TagAttribute {
	this := TagAttribute{}
	return &this
}

// NewTagAttributeWithDefaults instantiates a new TagAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAttributeWithDefaults() *TagAttribute {
	this := TagAttribute{}
	return &this
}

// GetGcpTagType returns the GcpTagType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagAttribute) GetGcpTagType() string {
	if o == nil || o.GcpTagType.Get() == nil {
		var ret string
		return ret
	}
	return *o.GcpTagType.Get()
}

// GetGcpTagTypeOk returns a tuple with the GcpTagType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagAttribute) GetGcpTagTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GcpTagType.Get(), o.GcpTagType.IsSet()
}

// HasGcpTagType returns a boolean if a field has been set.
func (o *TagAttribute) HasGcpTagType() bool {
	if o != nil && o.GcpTagType.IsSet() {
		return true
	}

	return false
}

// SetGcpTagType gets a reference to the given NullableString and assigns it to the GcpTagType field.
func (o *TagAttribute) SetGcpTagType(v string) {
	o.GcpTagType.Set(&v)
}
// SetGcpTagTypeNil sets the value for GcpTagType to be an explicit nil
func (o *TagAttribute) SetGcpTagTypeNil() {
	o.GcpTagType.Set(nil)
}

// UnsetGcpTagType ensures that no value is present for GcpTagType, not even an explicit nil
func (o *TagAttribute) UnsetGcpTagType() {
	o.GcpTagType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagAttribute) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagAttribute) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TagAttribute) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *TagAttribute) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TagAttribute) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TagAttribute) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagAttribute) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagAttribute) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TagAttribute) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TagAttribute) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TagAttribute) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TagAttribute) UnsetName() {
	o.Name.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagAttribute) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagAttribute) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *TagAttribute) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *TagAttribute) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *TagAttribute) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *TagAttribute) UnsetUuid() {
	o.Uuid.Unset()
}

func (o TagAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GcpTagType.IsSet() {
		toSerialize["gcpTagType"] = o.GcpTagType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTagAttribute struct {
	value *TagAttribute
	isSet bool
}

func (v NullableTagAttribute) Get() *TagAttribute {
	return v.value
}

func (v *NullableTagAttribute) Set(val *TagAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAttribute(val *TagAttribute) *NullableTagAttribute {
	return &NullableTagAttribute{value: val, isSet: true}
}

func (v NullableTagAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


