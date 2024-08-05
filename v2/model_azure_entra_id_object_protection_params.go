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

// checks if the AzureEntraIDObjectProtectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureEntraIDObjectProtectionParams{}

// AzureEntraIDObjectProtectionParams Specifies the parameters which are specific to Azure Entra ID Object Protection Groups using Azure native APIs. Atlease one of objects must be specified.
type AzureEntraIDObjectProtectionParams struct {
	// Specifies the list of object types to be excluded from protection.
	ExcludedObjectTypes []string `json:"excludedObjectTypes,omitempty"`
	// Specifies the objects to be protected.
	Objects []AzureObjectLevelParams `json:"objects,omitempty"`
}

// NewAzureEntraIDObjectProtectionParams instantiates a new AzureEntraIDObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureEntraIDObjectProtectionParams() *AzureEntraIDObjectProtectionParams {
	this := AzureEntraIDObjectProtectionParams{}
	return &this
}

// NewAzureEntraIDObjectProtectionParamsWithDefaults instantiates a new AzureEntraIDObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureEntraIDObjectProtectionParamsWithDefaults() *AzureEntraIDObjectProtectionParams {
	this := AzureEntraIDObjectProtectionParams{}
	return &this
}

// GetExcludedObjectTypes returns the ExcludedObjectTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureEntraIDObjectProtectionParams) GetExcludedObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludedObjectTypes
}

// GetExcludedObjectTypesOk returns a tuple with the ExcludedObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureEntraIDObjectProtectionParams) GetExcludedObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedObjectTypes) {
		return nil, false
	}
	return o.ExcludedObjectTypes, true
}

// HasExcludedObjectTypes returns a boolean if a field has been set.
func (o *AzureEntraIDObjectProtectionParams) HasExcludedObjectTypes() bool {
	if o != nil && !IsNil(o.ExcludedObjectTypes) {
		return true
	}

	return false
}

// SetExcludedObjectTypes gets a reference to the given []string and assigns it to the ExcludedObjectTypes field.
func (o *AzureEntraIDObjectProtectionParams) SetExcludedObjectTypes(v []string) {
	o.ExcludedObjectTypes = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AzureEntraIDObjectProtectionParams) GetObjects() []AzureObjectLevelParams {
	if o == nil || IsNil(o.Objects) {
		var ret []AzureObjectLevelParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureEntraIDObjectProtectionParams) GetObjectsOk() ([]AzureObjectLevelParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AzureEntraIDObjectProtectionParams) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []AzureObjectLevelParams and assigns it to the Objects field.
func (o *AzureEntraIDObjectProtectionParams) SetObjects(v []AzureObjectLevelParams) {
	o.Objects = v
}

func (o AzureEntraIDObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureEntraIDObjectProtectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludedObjectTypes != nil {
		toSerialize["excludedObjectTypes"] = o.ExcludedObjectTypes
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableAzureEntraIDObjectProtectionParams struct {
	value *AzureEntraIDObjectProtectionParams
	isSet bool
}

func (v NullableAzureEntraIDObjectProtectionParams) Get() *AzureEntraIDObjectProtectionParams {
	return v.value
}

func (v *NullableAzureEntraIDObjectProtectionParams) Set(val *AzureEntraIDObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureEntraIDObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureEntraIDObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureEntraIDObjectProtectionParams(val *AzureEntraIDObjectProtectionParams) *NullableAzureEntraIDObjectProtectionParams {
	return &NullableAzureEntraIDObjectProtectionParams{value: val, isSet: true}
}

func (v NullableAzureEntraIDObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureEntraIDObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


