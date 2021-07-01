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

// SmbActiveFileOpensResponse Query response to SMB active file opens.
type SmbActiveFileOpensResponse struct {
	// Specifies the active opens for an SMB file in a view.
	ActiveFilePaths []SmbActiveFilePath `json:"activeFilePaths,omitempty"`
	// Specifies an opaque string to pass to get the next set of active opens. If null is returned, this response is the last set of active opens.
	Cookie NullableString `json:"cookie,omitempty"`
}

// NewSmbActiveFileOpensResponse instantiates a new SmbActiveFileOpensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbActiveFileOpensResponse() *SmbActiveFileOpensResponse {
	this := SmbActiveFileOpensResponse{}
	return &this
}

// NewSmbActiveFileOpensResponseWithDefaults instantiates a new SmbActiveFileOpensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbActiveFileOpensResponseWithDefaults() *SmbActiveFileOpensResponse {
	this := SmbActiveFileOpensResponse{}
	return &this
}

// GetActiveFilePaths returns the ActiveFilePaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbActiveFileOpensResponse) GetActiveFilePaths() []SmbActiveFilePath {
	if o == nil  {
		var ret []SmbActiveFilePath
		return ret
	}
	return o.ActiveFilePaths
}

// GetActiveFilePathsOk returns a tuple with the ActiveFilePaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbActiveFileOpensResponse) GetActiveFilePathsOk() (*[]SmbActiveFilePath, bool) {
	if o == nil || o.ActiveFilePaths == nil {
		return nil, false
	}
	return &o.ActiveFilePaths, true
}

// HasActiveFilePaths returns a boolean if a field has been set.
func (o *SmbActiveFileOpensResponse) HasActiveFilePaths() bool {
	if o != nil && o.ActiveFilePaths != nil {
		return true
	}

	return false
}

// SetActiveFilePaths gets a reference to the given []SmbActiveFilePath and assigns it to the ActiveFilePaths field.
func (o *SmbActiveFileOpensResponse) SetActiveFilePaths(v []SmbActiveFilePath) {
	o.ActiveFilePaths = v
}

// GetCookie returns the Cookie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbActiveFileOpensResponse) GetCookie() string {
	if o == nil || o.Cookie.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cookie.Get()
}

// GetCookieOk returns a tuple with the Cookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbActiveFileOpensResponse) GetCookieOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cookie.Get(), o.Cookie.IsSet()
}

// HasCookie returns a boolean if a field has been set.
func (o *SmbActiveFileOpensResponse) HasCookie() bool {
	if o != nil && o.Cookie.IsSet() {
		return true
	}

	return false
}

// SetCookie gets a reference to the given NullableString and assigns it to the Cookie field.
func (o *SmbActiveFileOpensResponse) SetCookie(v string) {
	o.Cookie.Set(&v)
}
// SetCookieNil sets the value for Cookie to be an explicit nil
func (o *SmbActiveFileOpensResponse) SetCookieNil() {
	o.Cookie.Set(nil)
}

// UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
func (o *SmbActiveFileOpensResponse) UnsetCookie() {
	o.Cookie.Unset()
}

func (o SmbActiveFileOpensResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveFilePaths != nil {
		toSerialize["activeFilePaths"] = o.ActiveFilePaths
	}
	if o.Cookie.IsSet() {
		toSerialize["cookie"] = o.Cookie.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSmbActiveFileOpensResponse struct {
	value *SmbActiveFileOpensResponse
	isSet bool
}

func (v NullableSmbActiveFileOpensResponse) Get() *SmbActiveFileOpensResponse {
	return v.value
}

func (v *NullableSmbActiveFileOpensResponse) Set(val *SmbActiveFileOpensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbActiveFileOpensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbActiveFileOpensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbActiveFileOpensResponse(val *SmbActiveFileOpensResponse) *NullableSmbActiveFileOpensResponse {
	return &NullableSmbActiveFileOpensResponse{value: val, isSet: true}
}

func (v NullableSmbActiveFileOpensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbActiveFileOpensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


