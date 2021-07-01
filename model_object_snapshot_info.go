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

// ObjectSnapshotInfo Specifies information about an object that has been backed up.
type ObjectSnapshotInfo struct {
	// Specifies the Cohesity Cluster partition id where this object is stored.
	ClusterPartitionId NullableInt64 `json:"clusterPartitionId,omitempty"`
	// Specifies the id for the Protection Job that is currently associated with the object. If the object was backed up on current Cohesity Cluster, this field contains the id for the Job that captured this backup object. If the object was backed up on a Primary Cluster and replicated to this Cohesity Cluster, a new Inactive Job is created, the object is now associated with new Inactive Job, and this field contains the id of the new Inactive Job.
	JobId NullableInt64 `json:"jobId,omitempty"`
	// Specifies the name of the Protection Job that captured the backup.
	JobName NullableString `json:"jobName,omitempty"`
	// Specifies the globally unique id of the Protection Job that backed up this object. This id is unique across Cohesity Clusters. Even if this object is replicated to a Remote Cohesity Cluster and the object is associated with a new Job, the value specified in this field does not change.
	JobUid NullableUniversalId `json:"jobUid,omitempty"`
	// Specifies the primary name of the object.
	ObjectName NullableString `json:"objectName,omitempty"`
	// Specifies the inferred OS type.
	OsType NullableString `json:"osType,omitempty"`
	RegisteredSource *ProtectionSource `json:"registeredSource,omitempty"`
	SnapshottedSource *ProtectionSource `json:"snapshottedSource,omitempty"`
	// Array of Snapshots.  Specifies all snapshot versions of this object. Each time a Job Run of a Job executes, it may create a new snapshot version of an object. This array stores the different snapshots versions of the object.
	Versions []SnapshotVersion `json:"versions,omitempty"`
	// Specifies the id of the Domain (View Box) where this object is stored.
	ViewBoxId NullableInt64 `json:"viewBoxId,omitempty"`
	// Specifies the View name where this object is stored.
	ViewName NullableString `json:"viewName,omitempty"`
}

// NewObjectSnapshotInfo instantiates a new ObjectSnapshotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSnapshotInfo() *ObjectSnapshotInfo {
	this := ObjectSnapshotInfo{}
	return &this
}

// NewObjectSnapshotInfoWithDefaults instantiates a new ObjectSnapshotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSnapshotInfoWithDefaults() *ObjectSnapshotInfo {
	this := ObjectSnapshotInfo{}
	return &this
}

// GetClusterPartitionId returns the ClusterPartitionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetClusterPartitionId() int64 {
	if o == nil || o.ClusterPartitionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ClusterPartitionId.Get()
}

// GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetClusterPartitionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClusterPartitionId.Get(), o.ClusterPartitionId.IsSet()
}

// HasClusterPartitionId returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasClusterPartitionId() bool {
	if o != nil && o.ClusterPartitionId.IsSet() {
		return true
	}

	return false
}

// SetClusterPartitionId gets a reference to the given NullableInt64 and assigns it to the ClusterPartitionId field.
func (o *ObjectSnapshotInfo) SetClusterPartitionId(v int64) {
	o.ClusterPartitionId.Set(&v)
}
// SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil
func (o *ObjectSnapshotInfo) SetClusterPartitionIdNil() {
	o.ClusterPartitionId.Set(nil)
}

// UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetClusterPartitionId() {
	o.ClusterPartitionId.Unset()
}

// GetJobId returns the JobId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetJobId() int64 {
	if o == nil || o.JobId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.JobId.Get()
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetJobIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobId.Get(), o.JobId.IsSet()
}

// HasJobId returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasJobId() bool {
	if o != nil && o.JobId.IsSet() {
		return true
	}

	return false
}

// SetJobId gets a reference to the given NullableInt64 and assigns it to the JobId field.
func (o *ObjectSnapshotInfo) SetJobId(v int64) {
	o.JobId.Set(&v)
}
// SetJobIdNil sets the value for JobId to be an explicit nil
func (o *ObjectSnapshotInfo) SetJobIdNil() {
	o.JobId.Set(nil)
}

// UnsetJobId ensures that no value is present for JobId, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetJobId() {
	o.JobId.Unset()
}

// GetJobName returns the JobName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetJobName() string {
	if o == nil || o.JobName.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobName.Get()
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetJobNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobName.Get(), o.JobName.IsSet()
}

// HasJobName returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasJobName() bool {
	if o != nil && o.JobName.IsSet() {
		return true
	}

	return false
}

// SetJobName gets a reference to the given NullableString and assigns it to the JobName field.
func (o *ObjectSnapshotInfo) SetJobName(v string) {
	o.JobName.Set(&v)
}
// SetJobNameNil sets the value for JobName to be an explicit nil
func (o *ObjectSnapshotInfo) SetJobNameNil() {
	o.JobName.Set(nil)
}

// UnsetJobName ensures that no value is present for JobName, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetJobName() {
	o.JobName.Unset()
}

// GetJobUid returns the JobUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetJobUid() UniversalId {
	if o == nil || o.JobUid.Get() == nil {
		var ret UniversalId
		return ret
	}
	return *o.JobUid.Get()
}

// GetJobUidOk returns a tuple with the JobUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetJobUidOk() (*UniversalId, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobUid.Get(), o.JobUid.IsSet()
}

// HasJobUid returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasJobUid() bool {
	if o != nil && o.JobUid.IsSet() {
		return true
	}

	return false
}

// SetJobUid gets a reference to the given NullableUniversalId and assigns it to the JobUid field.
func (o *ObjectSnapshotInfo) SetJobUid(v UniversalId) {
	o.JobUid.Set(&v)
}
// SetJobUidNil sets the value for JobUid to be an explicit nil
func (o *ObjectSnapshotInfo) SetJobUidNil() {
	o.JobUid.Set(nil)
}

// UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetJobUid() {
	o.JobUid.Unset()
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetObjectName() string {
	if o == nil || o.ObjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectName.Get()
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetObjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectName.Get(), o.ObjectName.IsSet()
}

// HasObjectName returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasObjectName() bool {
	if o != nil && o.ObjectName.IsSet() {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given NullableString and assigns it to the ObjectName field.
func (o *ObjectSnapshotInfo) SetObjectName(v string) {
	o.ObjectName.Set(&v)
}
// SetObjectNameNil sets the value for ObjectName to be an explicit nil
func (o *ObjectSnapshotInfo) SetObjectNameNil() {
	o.ObjectName.Set(nil)
}

// UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetObjectName() {
	o.ObjectName.Unset()
}

// GetOsType returns the OsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetOsType() string {
	if o == nil || o.OsType.Get() == nil {
		var ret string
		return ret
	}
	return *o.OsType.Get()
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OsType.Get(), o.OsType.IsSet()
}

// HasOsType returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasOsType() bool {
	if o != nil && o.OsType.IsSet() {
		return true
	}

	return false
}

// SetOsType gets a reference to the given NullableString and assigns it to the OsType field.
func (o *ObjectSnapshotInfo) SetOsType(v string) {
	o.OsType.Set(&v)
}
// SetOsTypeNil sets the value for OsType to be an explicit nil
func (o *ObjectSnapshotInfo) SetOsTypeNil() {
	o.OsType.Set(nil)
}

// UnsetOsType ensures that no value is present for OsType, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetOsType() {
	o.OsType.Unset()
}

// GetRegisteredSource returns the RegisteredSource field value if set, zero value otherwise.
func (o *ObjectSnapshotInfo) GetRegisteredSource() ProtectionSource {
	if o == nil || o.RegisteredSource == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.RegisteredSource
}

// GetRegisteredSourceOk returns a tuple with the RegisteredSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSnapshotInfo) GetRegisteredSourceOk() (*ProtectionSource, bool) {
	if o == nil || o.RegisteredSource == nil {
		return nil, false
	}
	return o.RegisteredSource, true
}

// HasRegisteredSource returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasRegisteredSource() bool {
	if o != nil && o.RegisteredSource != nil {
		return true
	}

	return false
}

// SetRegisteredSource gets a reference to the given ProtectionSource and assigns it to the RegisteredSource field.
func (o *ObjectSnapshotInfo) SetRegisteredSource(v ProtectionSource) {
	o.RegisteredSource = &v
}

// GetSnapshottedSource returns the SnapshottedSource field value if set, zero value otherwise.
func (o *ObjectSnapshotInfo) GetSnapshottedSource() ProtectionSource {
	if o == nil || o.SnapshottedSource == nil {
		var ret ProtectionSource
		return ret
	}
	return *o.SnapshottedSource
}

// GetSnapshottedSourceOk returns a tuple with the SnapshottedSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSnapshotInfo) GetSnapshottedSourceOk() (*ProtectionSource, bool) {
	if o == nil || o.SnapshottedSource == nil {
		return nil, false
	}
	return o.SnapshottedSource, true
}

// HasSnapshottedSource returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasSnapshottedSource() bool {
	if o != nil && o.SnapshottedSource != nil {
		return true
	}

	return false
}

// SetSnapshottedSource gets a reference to the given ProtectionSource and assigns it to the SnapshottedSource field.
func (o *ObjectSnapshotInfo) SetSnapshottedSource(v ProtectionSource) {
	o.SnapshottedSource = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetVersions() []SnapshotVersion {
	if o == nil  {
		var ret []SnapshotVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetVersionsOk() (*[]SnapshotVersion, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []SnapshotVersion and assigns it to the Versions field.
func (o *ObjectSnapshotInfo) SetVersions(v []SnapshotVersion) {
	o.Versions = v
}

// GetViewBoxId returns the ViewBoxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetViewBoxId() int64 {
	if o == nil || o.ViewBoxId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewBoxId.Get()
}

// GetViewBoxIdOk returns a tuple with the ViewBoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetViewBoxIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewBoxId.Get(), o.ViewBoxId.IsSet()
}

// HasViewBoxId returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasViewBoxId() bool {
	if o != nil && o.ViewBoxId.IsSet() {
		return true
	}

	return false
}

// SetViewBoxId gets a reference to the given NullableInt64 and assigns it to the ViewBoxId field.
func (o *ObjectSnapshotInfo) SetViewBoxId(v int64) {
	o.ViewBoxId.Set(&v)
}
// SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil
func (o *ObjectSnapshotInfo) SetViewBoxIdNil() {
	o.ViewBoxId.Set(nil)
}

// UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetViewBoxId() {
	o.ViewBoxId.Unset()
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotInfo) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotInfo) GetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *ObjectSnapshotInfo) HasViewName() bool {
	if o != nil && o.ViewName.IsSet() {
		return true
	}

	return false
}

// SetViewName gets a reference to the given NullableString and assigns it to the ViewName field.
func (o *ObjectSnapshotInfo) SetViewName(v string) {
	o.ViewName.Set(&v)
}
// SetViewNameNil sets the value for ViewName to be an explicit nil
func (o *ObjectSnapshotInfo) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
func (o *ObjectSnapshotInfo) UnsetViewName() {
	o.ViewName.Unset()
}

func (o ObjectSnapshotInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterPartitionId.IsSet() {
		toSerialize["clusterPartitionId"] = o.ClusterPartitionId.Get()
	}
	if o.JobId.IsSet() {
		toSerialize["jobId"] = o.JobId.Get()
	}
	if o.JobName.IsSet() {
		toSerialize["jobName"] = o.JobName.Get()
	}
	if o.JobUid.IsSet() {
		toSerialize["jobUid"] = o.JobUid.Get()
	}
	if o.ObjectName.IsSet() {
		toSerialize["objectName"] = o.ObjectName.Get()
	}
	if o.OsType.IsSet() {
		toSerialize["osType"] = o.OsType.Get()
	}
	if o.RegisteredSource != nil {
		toSerialize["registeredSource"] = o.RegisteredSource
	}
	if o.SnapshottedSource != nil {
		toSerialize["snapshottedSource"] = o.SnapshottedSource
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.ViewBoxId.IsSet() {
		toSerialize["viewBoxId"] = o.ViewBoxId.Get()
	}
	if o.ViewName.IsSet() {
		toSerialize["viewName"] = o.ViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableObjectSnapshotInfo struct {
	value *ObjectSnapshotInfo
	isSet bool
}

func (v NullableObjectSnapshotInfo) Get() *ObjectSnapshotInfo {
	return v.value
}

func (v *NullableObjectSnapshotInfo) Set(val *ObjectSnapshotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSnapshotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSnapshotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSnapshotInfo(val *ObjectSnapshotInfo) *NullableObjectSnapshotInfo {
	return &NullableObjectSnapshotInfo{value: val, isSet: true}
}

func (v NullableObjectSnapshotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSnapshotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


