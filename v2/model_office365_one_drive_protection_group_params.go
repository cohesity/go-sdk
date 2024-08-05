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

// checks if the Office365OneDriveProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Office365OneDriveProtectionGroupParams{}

// Office365OneDriveProtectionGroupParams Specifies the parameters which are specific to Office 365 OneDrive related Protection Groups.
type Office365OneDriveProtectionGroupParams struct {
	// Array of Excluded OneDrive folders. Specifies filters to match OneDrive folders which should be excluded when backing up Office 365 source. Two kinds of filters are supported. a) prefix which always starts with '/'. b) posix which always starts with empty quotes(''). Regular expressions are not supported. If not specified, all the mailboxes will be protected.
	ExcludeFolders []string `json:"excludeFolders,omitempty"`
	PreservationHoldLibraryParams *Office365PreservationHoldLibraryParams `json:"preservationHoldLibraryParams,omitempty"`
}

// NewOffice365OneDriveProtectionGroupParams instantiates a new Office365OneDriveProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365OneDriveProtectionGroupParams() *Office365OneDriveProtectionGroupParams {
	this := Office365OneDriveProtectionGroupParams{}
	return &this
}

// NewOffice365OneDriveProtectionGroupParamsWithDefaults instantiates a new Office365OneDriveProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365OneDriveProtectionGroupParamsWithDefaults() *Office365OneDriveProtectionGroupParams {
	this := Office365OneDriveProtectionGroupParams{}
	return &this
}

// GetExcludeFolders returns the ExcludeFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365OneDriveProtectionGroupParams) GetExcludeFolders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeFolders
}

// GetExcludeFoldersOk returns a tuple with the ExcludeFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365OneDriveProtectionGroupParams) GetExcludeFoldersOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeFolders) {
		return nil, false
	}
	return o.ExcludeFolders, true
}

// HasExcludeFolders returns a boolean if a field has been set.
func (o *Office365OneDriveProtectionGroupParams) HasExcludeFolders() bool {
	if o != nil && !IsNil(o.ExcludeFolders) {
		return true
	}

	return false
}

// SetExcludeFolders gets a reference to the given []string and assigns it to the ExcludeFolders field.
func (o *Office365OneDriveProtectionGroupParams) SetExcludeFolders(v []string) {
	o.ExcludeFolders = v
}

// GetPreservationHoldLibraryParams returns the PreservationHoldLibraryParams field value if set, zero value otherwise.
func (o *Office365OneDriveProtectionGroupParams) GetPreservationHoldLibraryParams() Office365PreservationHoldLibraryParams {
	if o == nil || IsNil(o.PreservationHoldLibraryParams) {
		var ret Office365PreservationHoldLibraryParams
		return ret
	}
	return *o.PreservationHoldLibraryParams
}

// GetPreservationHoldLibraryParamsOk returns a tuple with the PreservationHoldLibraryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365OneDriveProtectionGroupParams) GetPreservationHoldLibraryParamsOk() (*Office365PreservationHoldLibraryParams, bool) {
	if o == nil || IsNil(o.PreservationHoldLibraryParams) {
		return nil, false
	}
	return o.PreservationHoldLibraryParams, true
}

// HasPreservationHoldLibraryParams returns a boolean if a field has been set.
func (o *Office365OneDriveProtectionGroupParams) HasPreservationHoldLibraryParams() bool {
	if o != nil && !IsNil(o.PreservationHoldLibraryParams) {
		return true
	}

	return false
}

// SetPreservationHoldLibraryParams gets a reference to the given Office365PreservationHoldLibraryParams and assigns it to the PreservationHoldLibraryParams field.
func (o *Office365OneDriveProtectionGroupParams) SetPreservationHoldLibraryParams(v Office365PreservationHoldLibraryParams) {
	o.PreservationHoldLibraryParams = &v
}

func (o Office365OneDriveProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Office365OneDriveProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeFolders != nil {
		toSerialize["excludeFolders"] = o.ExcludeFolders
	}
	if !IsNil(o.PreservationHoldLibraryParams) {
		toSerialize["preservationHoldLibraryParams"] = o.PreservationHoldLibraryParams
	}
	return toSerialize, nil
}

type NullableOffice365OneDriveProtectionGroupParams struct {
	value *Office365OneDriveProtectionGroupParams
	isSet bool
}

func (v NullableOffice365OneDriveProtectionGroupParams) Get() *Office365OneDriveProtectionGroupParams {
	return v.value
}

func (v *NullableOffice365OneDriveProtectionGroupParams) Set(val *Office365OneDriveProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365OneDriveProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365OneDriveProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365OneDriveProtectionGroupParams(val *Office365OneDriveProtectionGroupParams) *NullableOffice365OneDriveProtectionGroupParams {
	return &NullableOffice365OneDriveProtectionGroupParams{value: val, isSet: true}
}

func (v NullableOffice365OneDriveProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365OneDriveProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


