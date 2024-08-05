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

// checks if the AadGraphNodeFilterParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AadGraphNodeFilterParams{}

// AadGraphNodeFilterParams Determines filter params that can be applied to query aad node.
type AadGraphNodeFilterParams struct {
	// Filters the nodes which matches with specified aad node types provided. Supported AAD node types - Users/Groups/Applications/ AdministrativeUnits/ServicePrincipals/DirectoryRoles/Contacts/ Devices
	NodeTypes []string `json:"nodeTypes,omitempty"`
}

// NewAadGraphNodeFilterParams instantiates a new AadGraphNodeFilterParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadGraphNodeFilterParams() *AadGraphNodeFilterParams {
	this := AadGraphNodeFilterParams{}
	return &this
}

// NewAadGraphNodeFilterParamsWithDefaults instantiates a new AadGraphNodeFilterParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadGraphNodeFilterParamsWithDefaults() *AadGraphNodeFilterParams {
	this := AadGraphNodeFilterParams{}
	return &this
}

// GetNodeTypes returns the NodeTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AadGraphNodeFilterParams) GetNodeTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NodeTypes
}

// GetNodeTypesOk returns a tuple with the NodeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AadGraphNodeFilterParams) GetNodeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeTypes) {
		return nil, false
	}
	return o.NodeTypes, true
}

// HasNodeTypes returns a boolean if a field has been set.
func (o *AadGraphNodeFilterParams) HasNodeTypes() bool {
	if o != nil && !IsNil(o.NodeTypes) {
		return true
	}

	return false
}

// SetNodeTypes gets a reference to the given []string and assigns it to the NodeTypes field.
func (o *AadGraphNodeFilterParams) SetNodeTypes(v []string) {
	o.NodeTypes = v
}

func (o AadGraphNodeFilterParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AadGraphNodeFilterParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeTypes != nil {
		toSerialize["nodeTypes"] = o.NodeTypes
	}
	return toSerialize, nil
}

type NullableAadGraphNodeFilterParams struct {
	value *AadGraphNodeFilterParams
	isSet bool
}

func (v NullableAadGraphNodeFilterParams) Get() *AadGraphNodeFilterParams {
	return v.value
}

func (v *NullableAadGraphNodeFilterParams) Set(val *AadGraphNodeFilterParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAadGraphNodeFilterParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAadGraphNodeFilterParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadGraphNodeFilterParams(val *AadGraphNodeFilterParams) *NullableAadGraphNodeFilterParams {
	return &NullableAadGraphNodeFilterParams{value: val, isSet: true}
}

func (v NullableAadGraphNodeFilterParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadGraphNodeFilterParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


