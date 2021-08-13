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

// SearchFileRequestParamsBase Specifies the request parameters to search for files and file folders.
type SearchFileRequestParamsBase struct {
	// Specifies the search string to filter the files. User can specify a wildcard character '*' as a suffix to a string where all files name are matched with the prefix string.
	SearchString NullableString `json:"searchString,omitempty"`
	// Specifies a list of file types. Only files within the given types will be returned.
	Types []string `json:"types,omitempty"`
	// Specifies a list of the source environments. Only files from these types of source will be returned.
	SourceEnvironments []string `json:"sourceEnvironments,omitempty"`
}

// NewSearchFileRequestParamsBase instantiates a new SearchFileRequestParamsBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchFileRequestParamsBase() *SearchFileRequestParamsBase {
	this := SearchFileRequestParamsBase{}
	return &this
}

// NewSearchFileRequestParamsBaseWithDefaults instantiates a new SearchFileRequestParamsBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchFileRequestParamsBaseWithDefaults() *SearchFileRequestParamsBase {
	this := SearchFileRequestParamsBase{}
	return &this
}

// GetSearchString returns the SearchString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchFileRequestParamsBase) GetSearchString() string {
	if o == nil || o.SearchString.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchString.Get()
}

// GetSearchStringOk returns a tuple with the SearchString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchFileRequestParamsBase) GetSearchStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchString.Get(), o.SearchString.IsSet()
}

// HasSearchString returns a boolean if a field has been set.
func (o *SearchFileRequestParamsBase) HasSearchString() bool {
	if o != nil && o.SearchString.IsSet() {
		return true
	}

	return false
}

// SetSearchString gets a reference to the given NullableString and assigns it to the SearchString field.
func (o *SearchFileRequestParamsBase) SetSearchString(v string) {
	o.SearchString.Set(&v)
}
// SetSearchStringNil sets the value for SearchString to be an explicit nil
func (o *SearchFileRequestParamsBase) SetSearchStringNil() {
	o.SearchString.Set(nil)
}

// UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
func (o *SearchFileRequestParamsBase) UnsetSearchString() {
	o.SearchString.Unset()
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchFileRequestParamsBase) GetTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchFileRequestParamsBase) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return &o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SearchFileRequestParamsBase) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *SearchFileRequestParamsBase) SetTypes(v []string) {
	o.Types = v
}

// GetSourceEnvironments returns the SourceEnvironments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchFileRequestParamsBase) GetSourceEnvironments() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SourceEnvironments
}

// GetSourceEnvironmentsOk returns a tuple with the SourceEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchFileRequestParamsBase) GetSourceEnvironmentsOk() (*[]string, bool) {
	if o == nil || o.SourceEnvironments == nil {
		return nil, false
	}
	return &o.SourceEnvironments, true
}

// HasSourceEnvironments returns a boolean if a field has been set.
func (o *SearchFileRequestParamsBase) HasSourceEnvironments() bool {
	if o != nil && o.SourceEnvironments != nil {
		return true
	}

	return false
}

// SetSourceEnvironments gets a reference to the given []string and assigns it to the SourceEnvironments field.
func (o *SearchFileRequestParamsBase) SetSourceEnvironments(v []string) {
	o.SourceEnvironments = v
}

func (o SearchFileRequestParamsBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchString.IsSet() {
		toSerialize["searchString"] = o.SearchString.Get()
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.SourceEnvironments != nil {
		toSerialize["sourceEnvironments"] = o.SourceEnvironments
	}
	return json.Marshal(toSerialize)
}

type NullableSearchFileRequestParamsBase struct {
	value *SearchFileRequestParamsBase
	isSet bool
}

func (v NullableSearchFileRequestParamsBase) Get() *SearchFileRequestParamsBase {
	return v.value
}

func (v *NullableSearchFileRequestParamsBase) Set(val *SearchFileRequestParamsBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchFileRequestParamsBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchFileRequestParamsBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchFileRequestParamsBase(val *SearchFileRequestParamsBase) *NullableSearchFileRequestParamsBase {
	return &NullableSearchFileRequestParamsBase{value: val, isSet: true}
}

func (v NullableSearchFileRequestParamsBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchFileRequestParamsBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SearchFileRequestParamsBase) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}