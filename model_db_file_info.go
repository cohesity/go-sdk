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

// DbFileInfo Specifies information about a database file.
type DbFileInfo struct {
	// Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. 'kRows' refers to a data file 'kLog' refers to a log file 'kFileStream' refers to a directory containing FILESTREAM data 'kNotSupportedType' is for information purposes only. Not supported. 'kFullText' refers to a full-text catalog.
	FileType NullableString `json:"fileType,omitempty"`
	// Specifies the full path of the database file on the SQL host machine.
	FullPath NullableString `json:"fullPath,omitempty"`
	// Specifies the last known size of the database file.
	SizeBytes NullableInt64 `json:"sizeBytes,omitempty"`
}

// NewDbFileInfo instantiates a new DbFileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbFileInfo() *DbFileInfo {
	this := DbFileInfo{}
	return &this
}

// NewDbFileInfoWithDefaults instantiates a new DbFileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbFileInfoWithDefaults() *DbFileInfo {
	this := DbFileInfo{}
	return &this
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbFileInfo) GetFileType() string {
	if o == nil || o.FileType.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbFileInfo) GetFileTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *DbFileInfo) HasFileType() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *DbFileInfo) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *DbFileInfo) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *DbFileInfo) UnsetFileType() {
	o.FileType.Unset()
}

// GetFullPath returns the FullPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbFileInfo) GetFullPath() string {
	if o == nil || o.FullPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullPath.Get()
}

// GetFullPathOk returns a tuple with the FullPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbFileInfo) GetFullPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullPath.Get(), o.FullPath.IsSet()
}

// HasFullPath returns a boolean if a field has been set.
func (o *DbFileInfo) HasFullPath() bool {
	if o != nil && o.FullPath.IsSet() {
		return true
	}

	return false
}

// SetFullPath gets a reference to the given NullableString and assigns it to the FullPath field.
func (o *DbFileInfo) SetFullPath(v string) {
	o.FullPath.Set(&v)
}
// SetFullPathNil sets the value for FullPath to be an explicit nil
func (o *DbFileInfo) SetFullPathNil() {
	o.FullPath.Set(nil)
}

// UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil
func (o *DbFileInfo) UnsetFullPath() {
	o.FullPath.Unset()
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DbFileInfo) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes.Get()
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DbFileInfo) GetSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeBytes.Get(), o.SizeBytes.IsSet()
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *DbFileInfo) HasSizeBytes() bool {
	if o != nil && o.SizeBytes.IsSet() {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given NullableInt64 and assigns it to the SizeBytes field.
func (o *DbFileInfo) SetSizeBytes(v int64) {
	o.SizeBytes.Set(&v)
}
// SetSizeBytesNil sets the value for SizeBytes to be an explicit nil
func (o *DbFileInfo) SetSizeBytesNil() {
	o.SizeBytes.Set(nil)
}

// UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
func (o *DbFileInfo) UnsetSizeBytes() {
	o.SizeBytes.Unset()
}

func (o DbFileInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileType.IsSet() {
		toSerialize["fileType"] = o.FileType.Get()
	}
	if o.FullPath.IsSet() {
		toSerialize["fullPath"] = o.FullPath.Get()
	}
	if o.SizeBytes.IsSet() {
		toSerialize["sizeBytes"] = o.SizeBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDbFileInfo struct {
	value *DbFileInfo
	isSet bool
}

func (v NullableDbFileInfo) Get() *DbFileInfo {
	return v.value
}

func (v *NullableDbFileInfo) Set(val *DbFileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbFileInfo(val *DbFileInfo) *NullableDbFileInfo {
	return &NullableDbFileInfo{value: val, isSet: true}
}

func (v NullableDbFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


