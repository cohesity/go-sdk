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

// VmDirectoryListResult VmDirectoryListResult is a struct containing information about each directory entry.
type VmDirectoryListResult struct {
	// Cookie is used for paginating results. If ReadVMDirResult is returning partial results, this field will be set. Supplying this cookie will resume listing from where this result left off.
	Cookie NullableString `json:"cookie,omitempty"`
	// Entries is the array of files and folders that are immediate children of the parent directory specified in the request.
	Entries []VmDirEntry `json:"entries,omitempty"`
}

// NewVmDirectoryListResult instantiates a new VmDirectoryListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmDirectoryListResult() *VmDirectoryListResult {
	this := VmDirectoryListResult{}
	return &this
}

// NewVmDirectoryListResultWithDefaults instantiates a new VmDirectoryListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmDirectoryListResultWithDefaults() *VmDirectoryListResult {
	this := VmDirectoryListResult{}
	return &this
}

// GetCookie returns the Cookie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmDirectoryListResult) GetCookie() string {
	if o == nil || o.Cookie.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cookie.Get()
}

// GetCookieOk returns a tuple with the Cookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmDirectoryListResult) GetCookieOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cookie.Get(), o.Cookie.IsSet()
}

// HasCookie returns a boolean if a field has been set.
func (o *VmDirectoryListResult) HasCookie() bool {
	if o != nil && o.Cookie.IsSet() {
		return true
	}

	return false
}

// SetCookie gets a reference to the given NullableString and assigns it to the Cookie field.
func (o *VmDirectoryListResult) SetCookie(v string) {
	o.Cookie.Set(&v)
}
// SetCookieNil sets the value for Cookie to be an explicit nil
func (o *VmDirectoryListResult) SetCookieNil() {
	o.Cookie.Set(nil)
}

// UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
func (o *VmDirectoryListResult) UnsetCookie() {
	o.Cookie.Unset()
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmDirectoryListResult) GetEntries() []VmDirEntry {
	if o == nil  {
		var ret []VmDirEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmDirectoryListResult) GetEntriesOk() (*[]VmDirEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *VmDirectoryListResult) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []VmDirEntry and assigns it to the Entries field.
func (o *VmDirectoryListResult) SetEntries(v []VmDirEntry) {
	o.Entries = v
}

func (o VmDirectoryListResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cookie.IsSet() {
		toSerialize["cookie"] = o.Cookie.Get()
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableVmDirectoryListResult struct {
	value *VmDirectoryListResult
	isSet bool
}

func (v NullableVmDirectoryListResult) Get() *VmDirectoryListResult {
	return v.value
}

func (v *NullableVmDirectoryListResult) Set(val *VmDirectoryListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVmDirectoryListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVmDirectoryListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmDirectoryListResult(val *VmDirectoryListResult) *NullableVmDirectoryListResult {
	return &NullableVmDirectoryListResult{value: val, isSet: true}
}

func (v NullableVmDirectoryListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmDirectoryListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


