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

// CassandraProtectionSource Specifies an Object representing Cassandra.
type CassandraProtectionSource struct {
	ClusterInfo *CassandraCluster `json:"clusterInfo,omitempty"`
	KeyspaceInfo *CassandraKeyspace `json:"keyspaceInfo,omitempty"`
	// Specifies the instance name of the Cassandra entity.
	Name NullableString `json:"name,omitempty"`
	// Specifies the type of the managed Object in Cassandra Protection Source. Replication strategy options for a keyspace. 'kCluster' indicates a Cassandra cluster distributed over several physical nodes. 'kKeyspace' indicates a Keyspace enclosing one or more tables. 'kTable' indicates a Table in the Cassandra environment.
	Type NullableString `json:"type,omitempty"`
	// Specifies the UUID for the Cassandra entity. Note : For each entity an ID unique within top level entity should be assigned by imanis backend. Example, UUID for a table can be the string <keyspace_name>.<table_name>
	Uuid NullableString `json:"uuid,omitempty"`
}

// NewCassandraProtectionSource instantiates a new CassandraProtectionSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraProtectionSource() *CassandraProtectionSource {
	this := CassandraProtectionSource{}
	return &this
}

// NewCassandraProtectionSourceWithDefaults instantiates a new CassandraProtectionSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraProtectionSourceWithDefaults() *CassandraProtectionSource {
	this := CassandraProtectionSource{}
	return &this
}

// GetClusterInfo returns the ClusterInfo field value if set, zero value otherwise.
func (o *CassandraProtectionSource) GetClusterInfo() CassandraCluster {
	if o == nil || o.ClusterInfo == nil {
		var ret CassandraCluster
		return ret
	}
	return *o.ClusterInfo
}

// GetClusterInfoOk returns a tuple with the ClusterInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraProtectionSource) GetClusterInfoOk() (*CassandraCluster, bool) {
	if o == nil || o.ClusterInfo == nil {
		return nil, false
	}
	return o.ClusterInfo, true
}

// HasClusterInfo returns a boolean if a field has been set.
func (o *CassandraProtectionSource) HasClusterInfo() bool {
	if o != nil && o.ClusterInfo != nil {
		return true
	}

	return false
}

// SetClusterInfo gets a reference to the given CassandraCluster and assigns it to the ClusterInfo field.
func (o *CassandraProtectionSource) SetClusterInfo(v CassandraCluster) {
	o.ClusterInfo = &v
}

// GetKeyspaceInfo returns the KeyspaceInfo field value if set, zero value otherwise.
func (o *CassandraProtectionSource) GetKeyspaceInfo() CassandraKeyspace {
	if o == nil || o.KeyspaceInfo == nil {
		var ret CassandraKeyspace
		return ret
	}
	return *o.KeyspaceInfo
}

// GetKeyspaceInfoOk returns a tuple with the KeyspaceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraProtectionSource) GetKeyspaceInfoOk() (*CassandraKeyspace, bool) {
	if o == nil || o.KeyspaceInfo == nil {
		return nil, false
	}
	return o.KeyspaceInfo, true
}

// HasKeyspaceInfo returns a boolean if a field has been set.
func (o *CassandraProtectionSource) HasKeyspaceInfo() bool {
	if o != nil && o.KeyspaceInfo != nil {
		return true
	}

	return false
}

// SetKeyspaceInfo gets a reference to the given CassandraKeyspace and assigns it to the KeyspaceInfo field.
func (o *CassandraProtectionSource) SetKeyspaceInfo(v CassandraKeyspace) {
	o.KeyspaceInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraProtectionSource) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraProtectionSource) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CassandraProtectionSource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CassandraProtectionSource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CassandraProtectionSource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CassandraProtectionSource) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraProtectionSource) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraProtectionSource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CassandraProtectionSource) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CassandraProtectionSource) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CassandraProtectionSource) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CassandraProtectionSource) UnsetType() {
	o.Type.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraProtectionSource) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraProtectionSource) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CassandraProtectionSource) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CassandraProtectionSource) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CassandraProtectionSource) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CassandraProtectionSource) UnsetUuid() {
	o.Uuid.Unset()
}

func (o CassandraProtectionSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterInfo != nil {
		toSerialize["clusterInfo"] = o.ClusterInfo
	}
	if o.KeyspaceInfo != nil {
		toSerialize["keyspaceInfo"] = o.KeyspaceInfo
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCassandraProtectionSource struct {
	value *CassandraProtectionSource
	isSet bool
}

func (v NullableCassandraProtectionSource) Get() *CassandraProtectionSource {
	return v.value
}

func (v *NullableCassandraProtectionSource) Set(val *CassandraProtectionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraProtectionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraProtectionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraProtectionSource(val *CassandraProtectionSource) *NullableCassandraProtectionSource {
	return &NullableCassandraProtectionSource{value: val, isSet: true}
}

func (v NullableCassandraProtectionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraProtectionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


