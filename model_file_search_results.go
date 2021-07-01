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

// FileSearchResults Specifies an array of found files and folders. In addition, a count is provided to indicate if additional requests must be made to get the full result.
type FileSearchResults struct {
	// Array of Files and Folders.  Specifies the list of files and folders returned by this request that match the specified search and filter criteria. The number of files returned is limited by the pageCount field.
	Files []FileSearchResult `json:"files,omitempty"`
	// Specifies cookie for resuming search if pagination is being used. For Librarian queries only.
	PaginationCookie NullableString `json:"paginationCookie,omitempty"`
	// Specifies the total number of files and folders that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result.
	TotalCount NullableInt64 `json:"totalCount,omitempty"`
}

// NewFileSearchResults instantiates a new FileSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSearchResults() *FileSearchResults {
	this := FileSearchResults{}
	return &this
}

// NewFileSearchResultsWithDefaults instantiates a new FileSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSearchResultsWithDefaults() *FileSearchResults {
	this := FileSearchResults{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileSearchResults) GetFiles() []FileSearchResult {
	if o == nil  {
		var ret []FileSearchResult
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileSearchResults) GetFilesOk() (*[]FileSearchResult, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileSearchResults) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileSearchResult and assigns it to the Files field.
func (o *FileSearchResults) SetFiles(v []FileSearchResult) {
	o.Files = v
}

// GetPaginationCookie returns the PaginationCookie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileSearchResults) GetPaginationCookie() string {
	if o == nil || o.PaginationCookie.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaginationCookie.Get()
}

// GetPaginationCookieOk returns a tuple with the PaginationCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileSearchResults) GetPaginationCookieOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaginationCookie.Get(), o.PaginationCookie.IsSet()
}

// HasPaginationCookie returns a boolean if a field has been set.
func (o *FileSearchResults) HasPaginationCookie() bool {
	if o != nil && o.PaginationCookie.IsSet() {
		return true
	}

	return false
}

// SetPaginationCookie gets a reference to the given NullableString and assigns it to the PaginationCookie field.
func (o *FileSearchResults) SetPaginationCookie(v string) {
	o.PaginationCookie.Set(&v)
}
// SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil
func (o *FileSearchResults) SetPaginationCookieNil() {
	o.PaginationCookie.Set(nil)
}

// UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
func (o *FileSearchResults) UnsetPaginationCookie() {
	o.PaginationCookie.Unset()
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileSearchResults) GetTotalCount() int64 {
	if o == nil || o.TotalCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileSearchResults) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *FileSearchResults) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt64 and assigns it to the TotalCount field.
func (o *FileSearchResults) SetTotalCount(v int64) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *FileSearchResults) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *FileSearchResults) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o FileSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.PaginationCookie.IsSet() {
		toSerialize["paginationCookie"] = o.PaginationCookie.Get()
	}
	if o.TotalCount.IsSet() {
		toSerialize["totalCount"] = o.TotalCount.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFileSearchResults struct {
	value *FileSearchResults
	isSet bool
}

func (v NullableFileSearchResults) Get() *FileSearchResults {
	return v.value
}

func (v *NullableFileSearchResults) Set(val *FileSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSearchResults(val *FileSearchResults) *NullableFileSearchResults {
	return &NullableFileSearchResults{value: val, isSet: true}
}

func (v NullableFileSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


