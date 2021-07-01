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

// BackupJobProtoBackupSource struct for BackupJobProtoBackupSource
type BackupJobProtoBackupSource struct {
	// Source entities. NOTE: Multiple sources can be specified here for non-leaf-level entities in the hierarchy. The sources obtained after expanding these will be intersected among each other to form the final set of sources. e.g. this can be used to backup only those VMs that have both the tags 'SQL' and '3hrs'.
	Entities []EntityProto `json:"entities,omitempty"`
}

// NewBackupJobProtoBackupSource instantiates a new BackupJobProtoBackupSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobProtoBackupSource() *BackupJobProtoBackupSource {
	this := BackupJobProtoBackupSource{}
	return &this
}

// NewBackupJobProtoBackupSourceWithDefaults instantiates a new BackupJobProtoBackupSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobProtoBackupSourceWithDefaults() *BackupJobProtoBackupSource {
	this := BackupJobProtoBackupSource{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupJobProtoBackupSource) GetEntities() []EntityProto {
	if o == nil  {
		var ret []EntityProto
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupJobProtoBackupSource) GetEntitiesOk() (*[]EntityProto, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return &o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *BackupJobProtoBackupSource) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntityProto and assigns it to the Entities field.
func (o *BackupJobProtoBackupSource) SetEntities(v []EntityProto) {
	o.Entities = v
}

func (o BackupJobProtoBackupSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobProtoBackupSource struct {
	value *BackupJobProtoBackupSource
	isSet bool
}

func (v NullableBackupJobProtoBackupSource) Get() *BackupJobProtoBackupSource {
	return v.value
}

func (v *NullableBackupJobProtoBackupSource) Set(val *BackupJobProtoBackupSource) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobProtoBackupSource) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobProtoBackupSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobProtoBackupSource(val *BackupJobProtoBackupSource) *NullableBackupJobProtoBackupSource {
	return &NullableBackupJobProtoBackupSource{value: val, isSet: true}
}

func (v NullableBackupJobProtoBackupSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobProtoBackupSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


