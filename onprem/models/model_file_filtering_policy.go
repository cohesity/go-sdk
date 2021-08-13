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

// FileFilteringPolicy Specifies a set of filters for a file based Protection Group. These values are strings which can represent a prefix or suffix. Example: '/tmp' or '*.mp4'. For file based Protection Groups, all files under prefixes specified by the 'includeFilters' list will be protected unless they are explicitly excluded by the 'excludeFilters' list.
type FileFilteringPolicy struct {
	// Specifies the list of included files for this Protection Group.
	IncludeList *[]string `json:"includeList,omitempty"`
	// Specifies the list of excluded files for this protection Protection Group. Exclude filters have a higher priority than include filters.
	ExcludeList *[]string `json:"excludeList,omitempty"`
}

// NewFileFilteringPolicy instantiates a new FileFilteringPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileFilteringPolicy() *FileFilteringPolicy {
	this := FileFilteringPolicy{}
	return &this
}

// NewFileFilteringPolicyWithDefaults instantiates a new FileFilteringPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileFilteringPolicyWithDefaults() *FileFilteringPolicy {
	this := FileFilteringPolicy{}
	return &this
}

// GetIncludeList returns the IncludeList field value if set, zero value otherwise.
func (o *FileFilteringPolicy) GetIncludeList() []string {
	if o == nil || o.IncludeList == nil {
		var ret []string
		return ret
	}
	return *o.IncludeList
}

// GetIncludeListOk returns a tuple with the IncludeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFilteringPolicy) GetIncludeListOk() (*[]string, bool) {
	if o == nil || o.IncludeList == nil {
		return nil, false
	}
	return o.IncludeList, true
}

// HasIncludeList returns a boolean if a field has been set.
func (o *FileFilteringPolicy) HasIncludeList() bool {
	if o != nil && o.IncludeList != nil {
		return true
	}

	return false
}

// SetIncludeList gets a reference to the given []string and assigns it to the IncludeList field.
func (o *FileFilteringPolicy) SetIncludeList(v []string) {
	o.IncludeList = &v
}

// GetExcludeList returns the ExcludeList field value if set, zero value otherwise.
func (o *FileFilteringPolicy) GetExcludeList() []string {
	if o == nil || o.ExcludeList == nil {
		var ret []string
		return ret
	}
	return *o.ExcludeList
}

// GetExcludeListOk returns a tuple with the ExcludeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFilteringPolicy) GetExcludeListOk() (*[]string, bool) {
	if o == nil || o.ExcludeList == nil {
		return nil, false
	}
	return o.ExcludeList, true
}

// HasExcludeList returns a boolean if a field has been set.
func (o *FileFilteringPolicy) HasExcludeList() bool {
	if o != nil && o.ExcludeList != nil {
		return true
	}

	return false
}

// SetExcludeList gets a reference to the given []string and assigns it to the ExcludeList field.
func (o *FileFilteringPolicy) SetExcludeList(v []string) {
	o.ExcludeList = &v
}

func (o FileFilteringPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeList != nil {
		toSerialize["includeList"] = o.IncludeList
	}
	if o.ExcludeList != nil {
		toSerialize["excludeList"] = o.ExcludeList
	}
	return json.Marshal(toSerialize)
}

type NullableFileFilteringPolicy struct {
	value *FileFilteringPolicy
	isSet bool
}

func (v NullableFileFilteringPolicy) Get() *FileFilteringPolicy {
	return v.value
}

func (v *NullableFileFilteringPolicy) Set(val *FileFilteringPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFilteringPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFilteringPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFilteringPolicy(val *FileFilteringPolicy) *NullableFileFilteringPolicy {
	return &NullableFileFilteringPolicy{value: val, isSet: true}
}

func (v NullableFileFilteringPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFilteringPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o FileFilteringPolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}