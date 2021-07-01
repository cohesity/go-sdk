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

// AdObjectAttributeParameters AdObjectAttributeParameters are AD attribute recovery parameters for one or more AD objects
type AdObjectAttributeParameters struct {
	// Specifies the array of source and destination object guid pairs to restore attributes.
	AdGuidPairs []AdGuidPair1 `json:"adGuidPairs,omitempty"`
	// Specifies the array of LDAP property names to excluded from 'property_vec'. Excluded property name cannot contain wildcard character '*'.  Property names are case insensitive.
	ExcludeLdapProperties []string `json:"excludeLdapProperties,omitempty"`
	// Specifies the array of LDAP property(attribute) names. The name can be standard or custom property defined in AD schema partition. The property can contain wildcard character '*'. If this array is empty, then '*' is assigned, means restore all properties except default system excluded properties. Wildcards will be expanded. If 'memberOf' property is included, group membership of the objects specified in 'guid_vec' will be restored. Property that does not exist for an object is ignored and no error info is returned for that property. Property names are case insensitive.
	LdapProperties []string `json:"ldapProperties,omitempty"`
	// Specifies the Option to merge multi-valued values vs clearing and setting values from the AD snapshot. Defaults is to overwrite production AD values from the source AD snapshot. When updating group membership (using 'memberOf' property), setting this option to true will result in group membership merge. This is useful in cases where production AD has later updates that should not be overridden while restoring an attribute.
	MergeMultiValProperties NullableBool `json:"mergeMultiValProperties,omitempty"`
}

// NewAdObjectAttributeParameters instantiates a new AdObjectAttributeParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdObjectAttributeParameters() *AdObjectAttributeParameters {
	this := AdObjectAttributeParameters{}
	return &this
}

// NewAdObjectAttributeParametersWithDefaults instantiates a new AdObjectAttributeParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdObjectAttributeParametersWithDefaults() *AdObjectAttributeParameters {
	this := AdObjectAttributeParameters{}
	return &this
}

// GetAdGuidPairs returns the AdGuidPairs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdObjectAttributeParameters) GetAdGuidPairs() []AdGuidPair1 {
	if o == nil  {
		var ret []AdGuidPair1
		return ret
	}
	return o.AdGuidPairs
}

// GetAdGuidPairsOk returns a tuple with the AdGuidPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdObjectAttributeParameters) GetAdGuidPairsOk() (*[]AdGuidPair1, bool) {
	if o == nil || o.AdGuidPairs == nil {
		return nil, false
	}
	return &o.AdGuidPairs, true
}

// HasAdGuidPairs returns a boolean if a field has been set.
func (o *AdObjectAttributeParameters) HasAdGuidPairs() bool {
	if o != nil && o.AdGuidPairs != nil {
		return true
	}

	return false
}

// SetAdGuidPairs gets a reference to the given []AdGuidPair1 and assigns it to the AdGuidPairs field.
func (o *AdObjectAttributeParameters) SetAdGuidPairs(v []AdGuidPair1) {
	o.AdGuidPairs = v
}

// GetExcludeLdapProperties returns the ExcludeLdapProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdObjectAttributeParameters) GetExcludeLdapProperties() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ExcludeLdapProperties
}

// GetExcludeLdapPropertiesOk returns a tuple with the ExcludeLdapProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdObjectAttributeParameters) GetExcludeLdapPropertiesOk() (*[]string, bool) {
	if o == nil || o.ExcludeLdapProperties == nil {
		return nil, false
	}
	return &o.ExcludeLdapProperties, true
}

// HasExcludeLdapProperties returns a boolean if a field has been set.
func (o *AdObjectAttributeParameters) HasExcludeLdapProperties() bool {
	if o != nil && o.ExcludeLdapProperties != nil {
		return true
	}

	return false
}

// SetExcludeLdapProperties gets a reference to the given []string and assigns it to the ExcludeLdapProperties field.
func (o *AdObjectAttributeParameters) SetExcludeLdapProperties(v []string) {
	o.ExcludeLdapProperties = v
}

// GetLdapProperties returns the LdapProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdObjectAttributeParameters) GetLdapProperties() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.LdapProperties
}

// GetLdapPropertiesOk returns a tuple with the LdapProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdObjectAttributeParameters) GetLdapPropertiesOk() (*[]string, bool) {
	if o == nil || o.LdapProperties == nil {
		return nil, false
	}
	return &o.LdapProperties, true
}

// HasLdapProperties returns a boolean if a field has been set.
func (o *AdObjectAttributeParameters) HasLdapProperties() bool {
	if o != nil && o.LdapProperties != nil {
		return true
	}

	return false
}

// SetLdapProperties gets a reference to the given []string and assigns it to the LdapProperties field.
func (o *AdObjectAttributeParameters) SetLdapProperties(v []string) {
	o.LdapProperties = v
}

// GetMergeMultiValProperties returns the MergeMultiValProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdObjectAttributeParameters) GetMergeMultiValProperties() bool {
	if o == nil || o.MergeMultiValProperties.Get() == nil {
		var ret bool
		return ret
	}
	return *o.MergeMultiValProperties.Get()
}

// GetMergeMultiValPropertiesOk returns a tuple with the MergeMultiValProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdObjectAttributeParameters) GetMergeMultiValPropertiesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MergeMultiValProperties.Get(), o.MergeMultiValProperties.IsSet()
}

// HasMergeMultiValProperties returns a boolean if a field has been set.
func (o *AdObjectAttributeParameters) HasMergeMultiValProperties() bool {
	if o != nil && o.MergeMultiValProperties.IsSet() {
		return true
	}

	return false
}

// SetMergeMultiValProperties gets a reference to the given NullableBool and assigns it to the MergeMultiValProperties field.
func (o *AdObjectAttributeParameters) SetMergeMultiValProperties(v bool) {
	o.MergeMultiValProperties.Set(&v)
}
// SetMergeMultiValPropertiesNil sets the value for MergeMultiValProperties to be an explicit nil
func (o *AdObjectAttributeParameters) SetMergeMultiValPropertiesNil() {
	o.MergeMultiValProperties.Set(nil)
}

// UnsetMergeMultiValProperties ensures that no value is present for MergeMultiValProperties, not even an explicit nil
func (o *AdObjectAttributeParameters) UnsetMergeMultiValProperties() {
	o.MergeMultiValProperties.Unset()
}

func (o AdObjectAttributeParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdGuidPairs != nil {
		toSerialize["adGuidPairs"] = o.AdGuidPairs
	}
	if o.ExcludeLdapProperties != nil {
		toSerialize["excludeLdapProperties"] = o.ExcludeLdapProperties
	}
	if o.LdapProperties != nil {
		toSerialize["ldapProperties"] = o.LdapProperties
	}
	if o.MergeMultiValProperties.IsSet() {
		toSerialize["mergeMultiValProperties"] = o.MergeMultiValProperties.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAdObjectAttributeParameters struct {
	value *AdObjectAttributeParameters
	isSet bool
}

func (v NullableAdObjectAttributeParameters) Get() *AdObjectAttributeParameters {
	return v.value
}

func (v *NullableAdObjectAttributeParameters) Set(val *AdObjectAttributeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAdObjectAttributeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAdObjectAttributeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdObjectAttributeParameters(val *AdObjectAttributeParameters) *NullableAdObjectAttributeParameters {
	return &NullableAdObjectAttributeParameters{value: val, isSet: true}
}

func (v NullableAdObjectAttributeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdObjectAttributeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


