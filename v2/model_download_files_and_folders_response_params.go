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

// checks if the DownloadFilesAndFoldersResponseParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadFilesAndFoldersResponseParams{}

// DownloadFilesAndFoldersResponseParams Specifies the parameters to create a download files and folders Recovery.
type DownloadFilesAndFoldersResponseParams struct {
	// Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents.
	Documents []DocumentObject `json:"documents,omitempty"`
	// Specifies the list of files and folders to download. Only one of filesAndFolders or documents should be used.
	FilesAndFolders []FilesAndFoldersObject `json:"filesAndFolders,omitempty"`
	// Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot.
	GlacierRetrievalType NullableString `json:"glacierRetrievalType,omitempty"`
	// Specifies the id of the Recovery.
	Id NullableString `json:"id,omitempty"`
	// Specifies the name of the recovery task. This field must be set and must be a unique name.
	Name NullableString `json:"name,omitempty"`
	Object *CommonRecoverObjectSnapshotParams `json:"object,omitempty"`
}

// NewDownloadFilesAndFoldersResponseParams instantiates a new DownloadFilesAndFoldersResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadFilesAndFoldersResponseParams() *DownloadFilesAndFoldersResponseParams {
	this := DownloadFilesAndFoldersResponseParams{}
	return &this
}

// NewDownloadFilesAndFoldersResponseParamsWithDefaults instantiates a new DownloadFilesAndFoldersResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadFilesAndFoldersResponseParamsWithDefaults() *DownloadFilesAndFoldersResponseParams {
	this := DownloadFilesAndFoldersResponseParams{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadFilesAndFoldersResponseParams) GetDocuments() []DocumentObject {
	if o == nil {
		var ret []DocumentObject
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadFilesAndFoldersResponseParams) GetDocumentsOk() ([]DocumentObject, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DocumentObject and assigns it to the Documents field.
func (o *DownloadFilesAndFoldersResponseParams) SetDocuments(v []DocumentObject) {
	o.Documents = v
}

// GetFilesAndFolders returns the FilesAndFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadFilesAndFoldersResponseParams) GetFilesAndFolders() []FilesAndFoldersObject {
	if o == nil {
		var ret []FilesAndFoldersObject
		return ret
	}
	return o.FilesAndFolders
}

// GetFilesAndFoldersOk returns a tuple with the FilesAndFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadFilesAndFoldersResponseParams) GetFilesAndFoldersOk() ([]FilesAndFoldersObject, bool) {
	if o == nil || IsNil(o.FilesAndFolders) {
		return nil, false
	}
	return o.FilesAndFolders, true
}

// HasFilesAndFolders returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasFilesAndFolders() bool {
	if o != nil && !IsNil(o.FilesAndFolders) {
		return true
	}

	return false
}

// SetFilesAndFolders gets a reference to the given []FilesAndFoldersObject and assigns it to the FilesAndFolders field.
func (o *DownloadFilesAndFoldersResponseParams) SetFilesAndFolders(v []FilesAndFoldersObject) {
	o.FilesAndFolders = v
}

// GetGlacierRetrievalType returns the GlacierRetrievalType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadFilesAndFoldersResponseParams) GetGlacierRetrievalType() string {
	if o == nil || IsNil(o.GlacierRetrievalType.Get()) {
		var ret string
		return ret
	}
	return *o.GlacierRetrievalType.Get()
}

// GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadFilesAndFoldersResponseParams) GetGlacierRetrievalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlacierRetrievalType.Get(), o.GlacierRetrievalType.IsSet()
}

// HasGlacierRetrievalType returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasGlacierRetrievalType() bool {
	if o != nil && o.GlacierRetrievalType.IsSet() {
		return true
	}

	return false
}

// SetGlacierRetrievalType gets a reference to the given NullableString and assigns it to the GlacierRetrievalType field.
func (o *DownloadFilesAndFoldersResponseParams) SetGlacierRetrievalType(v string) {
	o.GlacierRetrievalType.Set(&v)
}
// SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) SetGlacierRetrievalTypeNil() {
	o.GlacierRetrievalType.Set(nil)
}

// UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) UnsetGlacierRetrievalType() {
	o.GlacierRetrievalType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadFilesAndFoldersResponseParams) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadFilesAndFoldersResponseParams) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DownloadFilesAndFoldersResponseParams) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadFilesAndFoldersResponseParams) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadFilesAndFoldersResponseParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DownloadFilesAndFoldersResponseParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DownloadFilesAndFoldersResponseParams) UnsetName() {
	o.Name.Unset()
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DownloadFilesAndFoldersResponseParams) GetObject() CommonRecoverObjectSnapshotParams {
	if o == nil || IsNil(o.Object) {
		var ret CommonRecoverObjectSnapshotParams
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadFilesAndFoldersResponseParams) GetObjectOk() (*CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DownloadFilesAndFoldersResponseParams) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given CommonRecoverObjectSnapshotParams and assigns it to the Object field.
func (o *DownloadFilesAndFoldersResponseParams) SetObject(v CommonRecoverObjectSnapshotParams) {
	o.Object = &v
}

func (o DownloadFilesAndFoldersResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadFilesAndFoldersResponseParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.FilesAndFolders != nil {
		toSerialize["filesAndFolders"] = o.FilesAndFolders
	}
	if o.GlacierRetrievalType.IsSet() {
		toSerialize["glacierRetrievalType"] = o.GlacierRetrievalType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableDownloadFilesAndFoldersResponseParams struct {
	value *DownloadFilesAndFoldersResponseParams
	isSet bool
}

func (v NullableDownloadFilesAndFoldersResponseParams) Get() *DownloadFilesAndFoldersResponseParams {
	return v.value
}

func (v *NullableDownloadFilesAndFoldersResponseParams) Set(val *DownloadFilesAndFoldersResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadFilesAndFoldersResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadFilesAndFoldersResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadFilesAndFoldersResponseParams(val *DownloadFilesAndFoldersResponseParams) *NullableDownloadFilesAndFoldersResponseParams {
	return &NullableDownloadFilesAndFoldersResponseParams{value: val, isSet: true}
}

func (v NullableDownloadFilesAndFoldersResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadFilesAndFoldersResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


