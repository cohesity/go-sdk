/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// Office365OutlookProtectionGroupParams Specifies the parameters which are specific to Office 365 Outlook related Protection Groups.
type Office365OutlookProtectionGroupParams struct {
	// Array of Excluded Outlook folders. Specifies filters to match Outlook folders which should be excluded when backing up Office 365 source. Two kinds of filters are supported. a) prefix which always starts with '/'. b) posix which always starts with empty quotes(''). Regular expressions are not supported. If not specified, all the mailboxes will be protected.
	ExcludeFolders []string `json:"excludeFolders,omitempty"`
}

// NewOffice365OutlookProtectionGroupParams instantiates a new Office365OutlookProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365OutlookProtectionGroupParams() *Office365OutlookProtectionGroupParams {
	this := Office365OutlookProtectionGroupParams{}
	return &this
}

// NewOffice365OutlookProtectionGroupParamsWithDefaults instantiates a new Office365OutlookProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365OutlookProtectionGroupParamsWithDefaults() *Office365OutlookProtectionGroupParams {
	this := Office365OutlookProtectionGroupParams{}
	return &this
}

// GetExcludeFolders returns the ExcludeFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365OutlookProtectionGroupParams) GetExcludeFolders() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ExcludeFolders
}

// GetExcludeFoldersOk returns a tuple with the ExcludeFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365OutlookProtectionGroupParams) GetExcludeFoldersOk() (*[]string, bool) {
	if o == nil || o.ExcludeFolders == nil {
		return nil, false
	}
	return &o.ExcludeFolders, true
}

// HasExcludeFolders returns a boolean if a field has been set.
func (o *Office365OutlookProtectionGroupParams) HasExcludeFolders() bool {
	if o != nil && o.ExcludeFolders != nil {
		return true
	}

	return false
}

// SetExcludeFolders gets a reference to the given []string and assigns it to the ExcludeFolders field.
func (o *Office365OutlookProtectionGroupParams) SetExcludeFolders(v []string) {
	o.ExcludeFolders = v
}

func (o Office365OutlookProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeFolders != nil {
		toSerialize["excludeFolders"] = o.ExcludeFolders
	}
	return json.Marshal(toSerialize)
}

type NullableOffice365OutlookProtectionGroupParams struct {
	value *Office365OutlookProtectionGroupParams
	isSet bool
}

func (v NullableOffice365OutlookProtectionGroupParams) Get() *Office365OutlookProtectionGroupParams {
	return v.value
}

func (v *NullableOffice365OutlookProtectionGroupParams) Set(val *Office365OutlookProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365OutlookProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365OutlookProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365OutlookProtectionGroupParams(val *Office365OutlookProtectionGroupParams) *NullableOffice365OutlookProtectionGroupParams {
	return &NullableOffice365OutlookProtectionGroupParams{value: val, isSet: true}
}

func (v NullableOffice365OutlookProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365OutlookProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


