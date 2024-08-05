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

// checks if the Office365SharePointProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Office365SharePointProtectionGroupParams{}

// Office365SharePointProtectionGroupParams Specifies the parameters which are specific to Office 365 SharePoint related Protection Groups.
type Office365SharePointProtectionGroupParams struct {
	// Array of paths to be excluded from backup. Specifies list of doclib/directory paths which should be excluded when backing up Office 365 source. supported exclusion: - doclib exclusion: whole doclib is excluded from backup. sample: /Doclib1 - directory exclusion: specified path in doclib will be excluded from backup. sample: /Doclib1/folderA/forderB Doclibs can be specified by either a) Doclib name - eg, Documents. b) Drive id of doclib - b!ZMSl2JRm0UeXLHfHR1m-iuD10p0CIV9qSa6TtgM Regular expressions are not supported. If not specified, all the doclibs within sharepoint site will be protected.
	ExcludePaths []string `json:"excludePaths,omitempty"`
	PreservationHoldLibraryParams *Office365PreservationHoldLibraryParams `json:"preservationHoldLibraryParams,omitempty"`
}

// NewOffice365SharePointProtectionGroupParams instantiates a new Office365SharePointProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365SharePointProtectionGroupParams() *Office365SharePointProtectionGroupParams {
	this := Office365SharePointProtectionGroupParams{}
	return &this
}

// NewOffice365SharePointProtectionGroupParamsWithDefaults instantiates a new Office365SharePointProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365SharePointProtectionGroupParamsWithDefaults() *Office365SharePointProtectionGroupParams {
	this := Office365SharePointProtectionGroupParams{}
	return &this
}

// GetExcludePaths returns the ExcludePaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365SharePointProtectionGroupParams) GetExcludePaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludePaths
}

// GetExcludePathsOk returns a tuple with the ExcludePaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365SharePointProtectionGroupParams) GetExcludePathsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludePaths) {
		return nil, false
	}
	return o.ExcludePaths, true
}

// HasExcludePaths returns a boolean if a field has been set.
func (o *Office365SharePointProtectionGroupParams) HasExcludePaths() bool {
	if o != nil && !IsNil(o.ExcludePaths) {
		return true
	}

	return false
}

// SetExcludePaths gets a reference to the given []string and assigns it to the ExcludePaths field.
func (o *Office365SharePointProtectionGroupParams) SetExcludePaths(v []string) {
	o.ExcludePaths = v
}

// GetPreservationHoldLibraryParams returns the PreservationHoldLibraryParams field value if set, zero value otherwise.
func (o *Office365SharePointProtectionGroupParams) GetPreservationHoldLibraryParams() Office365PreservationHoldLibraryParams {
	if o == nil || IsNil(o.PreservationHoldLibraryParams) {
		var ret Office365PreservationHoldLibraryParams
		return ret
	}
	return *o.PreservationHoldLibraryParams
}

// GetPreservationHoldLibraryParamsOk returns a tuple with the PreservationHoldLibraryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365SharePointProtectionGroupParams) GetPreservationHoldLibraryParamsOk() (*Office365PreservationHoldLibraryParams, bool) {
	if o == nil || IsNil(o.PreservationHoldLibraryParams) {
		return nil, false
	}
	return o.PreservationHoldLibraryParams, true
}

// HasPreservationHoldLibraryParams returns a boolean if a field has been set.
func (o *Office365SharePointProtectionGroupParams) HasPreservationHoldLibraryParams() bool {
	if o != nil && !IsNil(o.PreservationHoldLibraryParams) {
		return true
	}

	return false
}

// SetPreservationHoldLibraryParams gets a reference to the given Office365PreservationHoldLibraryParams and assigns it to the PreservationHoldLibraryParams field.
func (o *Office365SharePointProtectionGroupParams) SetPreservationHoldLibraryParams(v Office365PreservationHoldLibraryParams) {
	o.PreservationHoldLibraryParams = &v
}

func (o Office365SharePointProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Office365SharePointProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludePaths != nil {
		toSerialize["excludePaths"] = o.ExcludePaths
	}
	if !IsNil(o.PreservationHoldLibraryParams) {
		toSerialize["preservationHoldLibraryParams"] = o.PreservationHoldLibraryParams
	}
	return toSerialize, nil
}

type NullableOffice365SharePointProtectionGroupParams struct {
	value *Office365SharePointProtectionGroupParams
	isSet bool
}

func (v NullableOffice365SharePointProtectionGroupParams) Get() *Office365SharePointProtectionGroupParams {
	return v.value
}

func (v *NullableOffice365SharePointProtectionGroupParams) Set(val *Office365SharePointProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365SharePointProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365SharePointProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365SharePointProtectionGroupParams(val *Office365SharePointProtectionGroupParams) *NullableOffice365SharePointProtectionGroupParams {
	return &NullableOffice365SharePointProtectionGroupParams{value: val, isSet: true}
}

func (v NullableOffice365SharePointProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365SharePointProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


