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

// DownloadCftParams struct for DownloadCftParams
type DownloadCftParams struct {
	// Specifies the file name of the cloud formation template.
	FileName NullableString `json:"fileName,omitempty"`
	// Specifies the file path of the template. If passed null, \"/home/cohesity/bin\" will be considered as file path.
	FilePath NullableString `json:"filePath,omitempty"`
}

// NewDownloadCftParams instantiates a new DownloadCftParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadCftParams() *DownloadCftParams {
	this := DownloadCftParams{}
	return &this
}

// NewDownloadCftParamsWithDefaults instantiates a new DownloadCftParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadCftParamsWithDefaults() *DownloadCftParams {
	this := DownloadCftParams{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadCftParams) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadCftParams) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *DownloadCftParams) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *DownloadCftParams) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *DownloadCftParams) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *DownloadCftParams) UnsetFileName() {
	o.FileName.Unset()
}

// GetFilePath returns the FilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadCftParams) GetFilePath() string {
	if o == nil || o.FilePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadCftParams) GetFilePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// HasFilePath returns a boolean if a field has been set.
func (o *DownloadCftParams) HasFilePath() bool {
	if o != nil && o.FilePath.IsSet() {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given NullableString and assigns it to the FilePath field.
func (o *DownloadCftParams) SetFilePath(v string) {
	o.FilePath.Set(&v)
}
// SetFilePathNil sets the value for FilePath to be an explicit nil
func (o *DownloadCftParams) SetFilePathNil() {
	o.FilePath.Set(nil)
}

// UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
func (o *DownloadCftParams) UnsetFilePath() {
	o.FilePath.Unset()
}

func (o DownloadCftParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	if o.FilePath.IsSet() {
		toSerialize["filePath"] = o.FilePath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDownloadCftParams struct {
	value *DownloadCftParams
	isSet bool
}

func (v NullableDownloadCftParams) Get() *DownloadCftParams {
	return v.value
}

func (v *NullableDownloadCftParams) Set(val *DownloadCftParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadCftParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadCftParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadCftParams(val *DownloadCftParams) *NullableDownloadCftParams {
	return &NullableDownloadCftParams{value: val, isSet: true}
}

func (v NullableDownloadCftParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadCftParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


