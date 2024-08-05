/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubernetesRecoverFilesNewTargetConfigTargetNamespace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesRecoverFilesNewTargetConfigTargetNamespace{}

// KubernetesRecoverFilesNewTargetConfigTargetNamespace Specifies the target namespace to recover files and folders to.
type KubernetesRecoverFilesNewTargetConfigTargetNamespace struct {
	// Specifies the id of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies the id of the parent source of the target.
	ParentSourceId NullableInt64 `json:"parentSourceId,omitempty"`
	// Specifies the name of the parent source of the target.
	ParentSourceName NullableString `json:"parentSourceName,omitempty"`
}

type _KubernetesRecoverFilesNewTargetConfigTargetNamespace KubernetesRecoverFilesNewTargetConfigTargetNamespace

// NewKubernetesRecoverFilesNewTargetConfigTargetNamespace instantiates a new KubernetesRecoverFilesNewTargetConfigTargetNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesRecoverFilesNewTargetConfigTargetNamespace(id NullableInt64) *KubernetesRecoverFilesNewTargetConfigTargetNamespace {
	this := KubernetesRecoverFilesNewTargetConfigTargetNamespace{}
	this.Id = id
	return &this
}

// NewKubernetesRecoverFilesNewTargetConfigTargetNamespaceWithDefaults instantiates a new KubernetesRecoverFilesNewTargetConfigTargetNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesRecoverFilesNewTargetConfigTargetNamespaceWithDefaults() *KubernetesRecoverFilesNewTargetConfigTargetNamespace {
	this := KubernetesRecoverFilesNewTargetConfigTargetNamespace{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) UnsetName() {
	o.Name.Unset()
}

// GetParentSourceId returns the ParentSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetParentSourceId() int64 {
	if o == nil || IsNil(o.ParentSourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.ParentSourceId.Get()
}

// GetParentSourceIdOk returns a tuple with the ParentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetParentSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSourceId.Get(), o.ParentSourceId.IsSet()
}

// HasParentSourceId returns a boolean if a field has been set.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) HasParentSourceId() bool {
	if o != nil && o.ParentSourceId.IsSet() {
		return true
	}

	return false
}

// SetParentSourceId gets a reference to the given NullableInt64 and assigns it to the ParentSourceId field.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetParentSourceId(v int64) {
	o.ParentSourceId.Set(&v)
}
// SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetParentSourceIdNil() {
	o.ParentSourceId.Set(nil)
}

// UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) UnsetParentSourceId() {
	o.ParentSourceId.Unset()
}

// GetParentSourceName returns the ParentSourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetParentSourceName() string {
	if o == nil || IsNil(o.ParentSourceName.Get()) {
		var ret string
		return ret
	}
	return *o.ParentSourceName.Get()
}

// GetParentSourceNameOk returns a tuple with the ParentSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) GetParentSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSourceName.Get(), o.ParentSourceName.IsSet()
}

// HasParentSourceName returns a boolean if a field has been set.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) HasParentSourceName() bool {
	if o != nil && o.ParentSourceName.IsSet() {
		return true
	}

	return false
}

// SetParentSourceName gets a reference to the given NullableString and assigns it to the ParentSourceName field.
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetParentSourceName(v string) {
	o.ParentSourceName.Set(&v)
}
// SetParentSourceNameNil sets the value for ParentSourceName to be an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) SetParentSourceNameNil() {
	o.ParentSourceName.Set(nil)
}

// UnsetParentSourceName ensures that no value is present for ParentSourceName, not even an explicit nil
func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) UnsetParentSourceName() {
	o.ParentSourceName.Unset()
}

func (o KubernetesRecoverFilesNewTargetConfigTargetNamespace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesRecoverFilesNewTargetConfigTargetNamespace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentSourceId.IsSet() {
		toSerialize["parentSourceId"] = o.ParentSourceId.Get()
	}
	if o.ParentSourceName.IsSet() {
		toSerialize["parentSourceName"] = o.ParentSourceName.Get()
	}
	return toSerialize, nil
}

func (o *KubernetesRecoverFilesNewTargetConfigTargetNamespace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubernetesRecoverFilesNewTargetConfigTargetNamespace := _KubernetesRecoverFilesNewTargetConfigTargetNamespace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesRecoverFilesNewTargetConfigTargetNamespace)

	if err != nil {
		return err
	}

	*o = KubernetesRecoverFilesNewTargetConfigTargetNamespace(varKubernetesRecoverFilesNewTargetConfigTargetNamespace)

	return err
}

type NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace struct {
	value *KubernetesRecoverFilesNewTargetConfigTargetNamespace
	isSet bool
}

func (v NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) Get() *KubernetesRecoverFilesNewTargetConfigTargetNamespace {
	return v.value
}

func (v *NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) Set(val *KubernetesRecoverFilesNewTargetConfigTargetNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesRecoverFilesNewTargetConfigTargetNamespace(val *KubernetesRecoverFilesNewTargetConfigTargetNamespace) *NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace {
	return &NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace{value: val, isSet: true}
}

func (v NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


