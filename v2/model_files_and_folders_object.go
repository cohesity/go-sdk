/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FilesAndFoldersObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesAndFoldersObject{}

// FilesAndFoldersObject Specifies a file or folder to download.
type FilesAndFoldersObject struct {
	// Specifies the absolute path of the file or folder.
	AbsolutePath NullableString `json:"absolutePath"`
	// Specifies whether the file or folder object is a directory.
	IsDirectory NullableBool `json:"isDirectory,omitempty"`
}

type _FilesAndFoldersObject FilesAndFoldersObject

// NewFilesAndFoldersObject instantiates a new FilesAndFoldersObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesAndFoldersObject(absolutePath NullableString) *FilesAndFoldersObject {
	this := FilesAndFoldersObject{}
	this.AbsolutePath = absolutePath
	return &this
}

// NewFilesAndFoldersObjectWithDefaults instantiates a new FilesAndFoldersObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesAndFoldersObjectWithDefaults() *FilesAndFoldersObject {
	this := FilesAndFoldersObject{}
	return &this
}

// GetAbsolutePath returns the AbsolutePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesAndFoldersObject) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbsolutePath.Get()
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesAndFoldersObject) GetAbsolutePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsolutePath.Get(), o.AbsolutePath.IsSet()
}

// SetAbsolutePath sets field value
func (o *FilesAndFoldersObject) SetAbsolutePath(v string) {
	o.AbsolutePath.Set(&v)
}

// GetIsDirectory returns the IsDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesAndFoldersObject) GetIsDirectory() bool {
	if o == nil || IsNil(o.IsDirectory.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDirectory.Get()
}

// GetIsDirectoryOk returns a tuple with the IsDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesAndFoldersObject) GetIsDirectoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDirectory.Get(), o.IsDirectory.IsSet()
}

// HasIsDirectory returns a boolean if a field has been set.
func (o *FilesAndFoldersObject) HasIsDirectory() bool {
	if o != nil && o.IsDirectory.IsSet() {
		return true
	}

	return false
}

// SetIsDirectory gets a reference to the given NullableBool and assigns it to the IsDirectory field.
func (o *FilesAndFoldersObject) SetIsDirectory(v bool) {
	o.IsDirectory.Set(&v)
}
// SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil
func (o *FilesAndFoldersObject) SetIsDirectoryNil() {
	o.IsDirectory.Set(nil)
}

// UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
func (o *FilesAndFoldersObject) UnsetIsDirectory() {
	o.IsDirectory.Unset()
}

func (o FilesAndFoldersObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesAndFoldersObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["absolutePath"] = o.AbsolutePath.Get()
	if o.IsDirectory.IsSet() {
		toSerialize["isDirectory"] = o.IsDirectory.Get()
	}
	return toSerialize, nil
}

func (o *FilesAndFoldersObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"absolutePath",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFilesAndFoldersObject := _FilesAndFoldersObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilesAndFoldersObject)

	if err != nil {
		return err
	}

	*o = FilesAndFoldersObject(varFilesAndFoldersObject)

	return err
}

type NullableFilesAndFoldersObject struct {
	value *FilesAndFoldersObject
	isSet bool
}

func (v NullableFilesAndFoldersObject) Get() *FilesAndFoldersObject {
	return v.value
}

func (v *NullableFilesAndFoldersObject) Set(val *FilesAndFoldersObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesAndFoldersObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesAndFoldersObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesAndFoldersObject(val *FilesAndFoldersObject) *NullableFilesAndFoldersObject {
	return &NullableFilesAndFoldersObject{value: val, isSet: true}
}

func (v NullableFilesAndFoldersObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesAndFoldersObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


