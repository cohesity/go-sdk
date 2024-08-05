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

// checks if the RecoverAcropolisParamsDownloadFileAndFolderParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverAcropolisParamsDownloadFileAndFolderParams{}

// RecoverAcropolisParamsDownloadFileAndFolderParams Specifies the parameters to download files and folders.
type RecoverAcropolisParamsDownloadFileAndFolderParams struct {
	// Specifies the path location to download the files and folders.
	DownloadFilePath NullableString `json:"downloadFilePath,omitempty"`
	// Specifies the time upto which the download link is available.
	ExpiryTimeUsecs NullableInt64 `json:"expiryTimeUsecs,omitempty"`
	// Specifies the info about the files and folders to be recovered.
	FilesAndFolders []CommonRecoverFileAndFolderInfo `json:"filesAndFolders,omitempty"`
}

// NewRecoverAcropolisParamsDownloadFileAndFolderParams instantiates a new RecoverAcropolisParamsDownloadFileAndFolderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAcropolisParamsDownloadFileAndFolderParams() *RecoverAcropolisParamsDownloadFileAndFolderParams {
	this := RecoverAcropolisParamsDownloadFileAndFolderParams{}
	return &this
}

// NewRecoverAcropolisParamsDownloadFileAndFolderParamsWithDefaults instantiates a new RecoverAcropolisParamsDownloadFileAndFolderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAcropolisParamsDownloadFileAndFolderParamsWithDefaults() *RecoverAcropolisParamsDownloadFileAndFolderParams {
	this := RecoverAcropolisParamsDownloadFileAndFolderParams{}
	return &this
}

// GetDownloadFilePath returns the DownloadFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetDownloadFilePath() string {
	if o == nil || IsNil(o.DownloadFilePath.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadFilePath.Get()
}

// GetDownloadFilePathOk returns a tuple with the DownloadFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetDownloadFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadFilePath.Get(), o.DownloadFilePath.IsSet()
}

// HasDownloadFilePath returns a boolean if a field has been set.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) HasDownloadFilePath() bool {
	if o != nil && o.DownloadFilePath.IsSet() {
		return true
	}

	return false
}

// SetDownloadFilePath gets a reference to the given NullableString and assigns it to the DownloadFilePath field.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) SetDownloadFilePath(v string) {
	o.DownloadFilePath.Set(&v)
}
// SetDownloadFilePathNil sets the value for DownloadFilePath to be an explicit nil
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) SetDownloadFilePathNil() {
	o.DownloadFilePath.Set(nil)
}

// UnsetDownloadFilePath ensures that no value is present for DownloadFilePath, not even an explicit nil
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) UnsetDownloadFilePath() {
	o.DownloadFilePath.Unset()
}

// GetExpiryTimeUsecs returns the ExpiryTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetExpiryTimeUsecs() int64 {
	if o == nil || IsNil(o.ExpiryTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiryTimeUsecs.Get()
}

// GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetExpiryTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryTimeUsecs.Get(), o.ExpiryTimeUsecs.IsSet()
}

// HasExpiryTimeUsecs returns a boolean if a field has been set.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) HasExpiryTimeUsecs() bool {
	if o != nil && o.ExpiryTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiryTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ExpiryTimeUsecs field.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) SetExpiryTimeUsecs(v int64) {
	o.ExpiryTimeUsecs.Set(&v)
}
// SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) SetExpiryTimeUsecsNil() {
	o.ExpiryTimeUsecs.Set(nil)
}

// UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) UnsetExpiryTimeUsecs() {
	o.ExpiryTimeUsecs.Unset()
}

// GetFilesAndFolders returns the FilesAndFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetFilesAndFolders() []CommonRecoverFileAndFolderInfo {
	if o == nil {
		var ret []CommonRecoverFileAndFolderInfo
		return ret
	}
	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) GetFilesAndFoldersOk() ([]CommonRecoverFileAndFolderInfo, bool) {
	if o == nil || IsNil(o.FilesAndFolders) {
		return nil, false
	}
	return o.FilesAndFolders, true
}

// HasFilesAndFolders returns a boolean if a field has been set.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) HasFilesAndFolders() bool {
	if o != nil && !IsNil(o.FilesAndFolders) {
		return true
	}

	return false
}

// SetFilesAndFolders gets a reference to the given []CommonRecoverFileAndFolderInfo and assigns it to the FilesAndFolders field.
func (o *RecoverAcropolisParamsDownloadFileAndFolderParams) SetFilesAndFolders(v []CommonRecoverFileAndFolderInfo) {
	o.FilesAndFolders = v
}

func (o RecoverAcropolisParamsDownloadFileAndFolderParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverAcropolisParamsDownloadFileAndFolderParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadFilePath.IsSet() {
		toSerialize["downloadFilePath"] = o.DownloadFilePath.Get()
	}
	if o.ExpiryTimeUsecs.IsSet() {
		toSerialize["expiryTimeUsecs"] = o.ExpiryTimeUsecs.Get()
	}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	return toSerialize, nil
}

type NullableRecoverAcropolisParamsDownloadFileAndFolderParams struct {
	value *RecoverAcropolisParamsDownloadFileAndFolderParams
	isSet bool
}

func (v NullableRecoverAcropolisParamsDownloadFileAndFolderParams) Get() *RecoverAcropolisParamsDownloadFileAndFolderParams {
	return v.value
}

func (v *NullableRecoverAcropolisParamsDownloadFileAndFolderParams) Set(val *RecoverAcropolisParamsDownloadFileAndFolderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAcropolisParamsDownloadFileAndFolderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAcropolisParamsDownloadFileAndFolderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAcropolisParamsDownloadFileAndFolderParams(val *RecoverAcropolisParamsDownloadFileAndFolderParams) *NullableRecoverAcropolisParamsDownloadFileAndFolderParams {
	return &NullableRecoverAcropolisParamsDownloadFileAndFolderParams{value: val, isSet: true}
}

func (v NullableRecoverAcropolisParamsDownloadFileAndFolderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAcropolisParamsDownloadFileAndFolderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


