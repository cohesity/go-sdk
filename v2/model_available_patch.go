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

// checks if the AvailablePatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePatch{}

// AvailablePatch Specifies the description of an available patch.
type AvailablePatch struct {
	// Specifies the description of the service.
	Component *string `json:"component,omitempty"`
	// Specifies the number of fixed issues.
	Count *int64 `json:"count,omitempty"`
	// Specifies the services for which their patches must be applied together.
	Dependencies []string `json:"dependencies,omitempty"`
	// Specifies the details of the issues fixed in the patch.
	FixedIssues []FixedIssue `json:"fixedIssues,omitempty"`
	// Specifies the name of the service.
	Service *string `json:"service,omitempty"`
	// Specifies the version of the patch.
	Version *string `json:"version,omitempty"`
}

// NewAvailablePatch instantiates a new AvailablePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePatch() *AvailablePatch {
	this := AvailablePatch{}
	return &this
}

// NewAvailablePatchWithDefaults instantiates a new AvailablePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePatchWithDefaults() *AvailablePatch {
	this := AvailablePatch{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AvailablePatch) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePatch) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AvailablePatch) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AvailablePatch) SetComponent(v string) {
	o.Component = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AvailablePatch) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePatch) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AvailablePatch) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *AvailablePatch) SetCount(v int64) {
	o.Count = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePatch) GetDependencies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePatch) GetDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *AvailablePatch) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *AvailablePatch) SetDependencies(v []string) {
	o.Dependencies = v
}

// GetFixedIssues returns the FixedIssues field value if set, zero value otherwise.
func (o *AvailablePatch) GetFixedIssues() []FixedIssue {
	if o == nil || IsNil(o.FixedIssues) {
		var ret []FixedIssue
		return ret
	}
	return o.FixedIssues
}

// GetFixedIssuesOk returns a tuple with the FixedIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePatch) GetFixedIssuesOk() ([]FixedIssue, bool) {
	if o == nil || IsNil(o.FixedIssues) {
		return nil, false
	}
	return o.FixedIssues, true
}

// HasFixedIssues returns a boolean if a field has been set.
func (o *AvailablePatch) HasFixedIssues() bool {
	if o != nil && !IsNil(o.FixedIssues) {
		return true
	}

	return false
}

// SetFixedIssues gets a reference to the given []FixedIssue and assigns it to the FixedIssues field.
func (o *AvailablePatch) SetFixedIssues(v []FixedIssue) {
	o.FixedIssues = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AvailablePatch) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePatch) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AvailablePatch) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *AvailablePatch) SetService(v string) {
	o.Service = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AvailablePatch) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePatch) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AvailablePatch) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AvailablePatch) SetVersion(v string) {
	o.Version = &v
}

func (o AvailablePatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.FixedIssues) {
		toSerialize["fixedIssues"] = o.FixedIssues
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAvailablePatch struct {
	value *AvailablePatch
	isSet bool
}

func (v NullableAvailablePatch) Get() *AvailablePatch {
	return v.value
}

func (v *NullableAvailablePatch) Set(val *AvailablePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePatch(val *AvailablePatch) *NullableAvailablePatch {
	return &NullableAvailablePatch{value: val, isSet: true}
}

func (v NullableAvailablePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


