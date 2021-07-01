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

// RetrieveArchiveInfoRetrievedEntity Proto to define the info about the entity that was retrieved from an archive.
type RetrieveArchiveInfoRetrievedEntity struct {
	// Number of physical bytes transferred over wire for this entity.
	BytesTransferred NullableInt64 `json:"bytesTransferred,omitempty"`
	// Time in microseconds when retrieve of this entity finished or failed.
	EndTimestampUsecs NullableInt64 `json:"endTimestampUsecs,omitempty"`
	Entity *EntityProto `json:"entity,omitempty"`
	Error *ErrorProto `json:"error,omitempty"`
	// Number of logical bytes transferred so far.
	LogicalBytesTransferred NullableInt64 `json:"logicalBytesTransferred,omitempty"`
	// Total logical size of this entity.
	LogicalSizeBytes NullableInt64 `json:"logicalSizeBytes,omitempty"`
	// The path relative to the root path of the retrieval task progress monitor of this entity progress monitor.
	ProgressMonitorTaskPath NullableString `json:"progressMonitorTaskPath,omitempty"`
	// The path relative to the root of the file system where the snapshot of this entity was retrieved/copied to.
	RelativeSnapshotDir NullableString `json:"relativeSnapshotDir,omitempty"`
	// Time in microseconds when retrieve of this entity started.
	StartTimestampUsecs NullableInt64 `json:"startTimestampUsecs,omitempty"`
	// The retrieval status of this entity.
	Status NullableInt32 `json:"status,omitempty"`
}

// NewRetrieveArchiveInfoRetrievedEntity instantiates a new RetrieveArchiveInfoRetrievedEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveArchiveInfoRetrievedEntity() *RetrieveArchiveInfoRetrievedEntity {
	this := RetrieveArchiveInfoRetrievedEntity{}
	return &this
}

// NewRetrieveArchiveInfoRetrievedEntityWithDefaults instantiates a new RetrieveArchiveInfoRetrievedEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveArchiveInfoRetrievedEntityWithDefaults() *RetrieveArchiveInfoRetrievedEntity {
	this := RetrieveArchiveInfoRetrievedEntity{}
	return &this
}

// GetBytesTransferred returns the BytesTransferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetBytesTransferred() int64 {
	if o == nil || o.BytesTransferred.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BytesTransferred.Get()
}

// GetBytesTransferredOk returns a tuple with the BytesTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetBytesTransferredOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BytesTransferred.Get(), o.BytesTransferred.IsSet()
}

// HasBytesTransferred returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasBytesTransferred() bool {
	if o != nil && o.BytesTransferred.IsSet() {
		return true
	}

	return false
}

// SetBytesTransferred gets a reference to the given NullableInt64 and assigns it to the BytesTransferred field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetBytesTransferred(v int64) {
	o.BytesTransferred.Set(&v)
}
// SetBytesTransferredNil sets the value for BytesTransferred to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetBytesTransferredNil() {
	o.BytesTransferred.Set(nil)
}

// UnsetBytesTransferred ensures that no value is present for BytesTransferred, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetBytesTransferred() {
	o.BytesTransferred.Unset()
}

// GetEndTimestampUsecs returns the EndTimestampUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetEndTimestampUsecs() int64 {
	if o == nil || o.EndTimestampUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimestampUsecs.Get()
}

// GetEndTimestampUsecsOk returns a tuple with the EndTimestampUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetEndTimestampUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimestampUsecs.Get(), o.EndTimestampUsecs.IsSet()
}

// HasEndTimestampUsecs returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasEndTimestampUsecs() bool {
	if o != nil && o.EndTimestampUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimestampUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimestampUsecs field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetEndTimestampUsecs(v int64) {
	o.EndTimestampUsecs.Set(&v)
}
// SetEndTimestampUsecsNil sets the value for EndTimestampUsecs to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetEndTimestampUsecsNil() {
	o.EndTimestampUsecs.Set(nil)
}

// UnsetEndTimestampUsecs ensures that no value is present for EndTimestampUsecs, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetEndTimestampUsecs() {
	o.EndTimestampUsecs.Unset()
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *RetrieveArchiveInfoRetrievedEntity) GetEntity() EntityProto {
	if o == nil || o.Entity == nil {
		var ret EntityProto
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) GetEntityOk() (*EntityProto, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntityProto and assigns it to the Entity field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetEntity(v EntityProto) {
	o.Entity = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RetrieveArchiveInfoRetrievedEntity) GetError() ErrorProto {
	if o == nil || o.Error == nil {
		var ret ErrorProto
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) GetErrorOk() (*ErrorProto, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorProto and assigns it to the Error field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetError(v ErrorProto) {
	o.Error = &v
}

// GetLogicalBytesTransferred returns the LogicalBytesTransferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalBytesTransferred() int64 {
	if o == nil || o.LogicalBytesTransferred.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogicalBytesTransferred.Get()
}

// GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalBytesTransferredOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalBytesTransferred.Get(), o.LogicalBytesTransferred.IsSet()
}

// HasLogicalBytesTransferred returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasLogicalBytesTransferred() bool {
	if o != nil && o.LogicalBytesTransferred.IsSet() {
		return true
	}

	return false
}

// SetLogicalBytesTransferred gets a reference to the given NullableInt64 and assigns it to the LogicalBytesTransferred field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalBytesTransferred(v int64) {
	o.LogicalBytesTransferred.Set(&v)
}
// SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalBytesTransferredNil() {
	o.LogicalBytesTransferred.Set(nil)
}

// UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetLogicalBytesTransferred() {
	o.LogicalBytesTransferred.Unset()
}

// GetLogicalSizeBytes returns the LogicalSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalSizeBytes() int64 {
	if o == nil || o.LogicalSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogicalSizeBytes.Get()
}

// GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalSizeBytes.Get(), o.LogicalSizeBytes.IsSet()
}

// HasLogicalSizeBytes returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasLogicalSizeBytes() bool {
	if o != nil && o.LogicalSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetLogicalSizeBytes gets a reference to the given NullableInt64 and assigns it to the LogicalSizeBytes field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalSizeBytes(v int64) {
	o.LogicalSizeBytes.Set(&v)
}
// SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalSizeBytesNil() {
	o.LogicalSizeBytes.Set(nil)
}

// UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetLogicalSizeBytes() {
	o.LogicalSizeBytes.Unset()
}

// GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetProgressMonitorTaskPath() string {
	if o == nil || o.ProgressMonitorTaskPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressMonitorTaskPath.Get()
}

// GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetProgressMonitorTaskPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProgressMonitorTaskPath.Get(), o.ProgressMonitorTaskPath.IsSet()
}

// HasProgressMonitorTaskPath returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasProgressMonitorTaskPath() bool {
	if o != nil && o.ProgressMonitorTaskPath.IsSet() {
		return true
	}

	return false
}

// SetProgressMonitorTaskPath gets a reference to the given NullableString and assigns it to the ProgressMonitorTaskPath field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetProgressMonitorTaskPath(v string) {
	o.ProgressMonitorTaskPath.Set(&v)
}
// SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetProgressMonitorTaskPathNil() {
	o.ProgressMonitorTaskPath.Set(nil)
}

// UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetProgressMonitorTaskPath() {
	o.ProgressMonitorTaskPath.Unset()
}

// GetRelativeSnapshotDir returns the RelativeSnapshotDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetRelativeSnapshotDir() string {
	if o == nil || o.RelativeSnapshotDir.Get() == nil {
		var ret string
		return ret
	}
	return *o.RelativeSnapshotDir.Get()
}

// GetRelativeSnapshotDirOk returns a tuple with the RelativeSnapshotDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetRelativeSnapshotDirOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RelativeSnapshotDir.Get(), o.RelativeSnapshotDir.IsSet()
}

// HasRelativeSnapshotDir returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasRelativeSnapshotDir() bool {
	if o != nil && o.RelativeSnapshotDir.IsSet() {
		return true
	}

	return false
}

// SetRelativeSnapshotDir gets a reference to the given NullableString and assigns it to the RelativeSnapshotDir field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetRelativeSnapshotDir(v string) {
	o.RelativeSnapshotDir.Set(&v)
}
// SetRelativeSnapshotDirNil sets the value for RelativeSnapshotDir to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetRelativeSnapshotDirNil() {
	o.RelativeSnapshotDir.Set(nil)
}

// UnsetRelativeSnapshotDir ensures that no value is present for RelativeSnapshotDir, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetRelativeSnapshotDir() {
	o.RelativeSnapshotDir.Unset()
}

// GetStartTimestampUsecs returns the StartTimestampUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetStartTimestampUsecs() int64 {
	if o == nil || o.StartTimestampUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimestampUsecs.Get()
}

// GetStartTimestampUsecsOk returns a tuple with the StartTimestampUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetStartTimestampUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimestampUsecs.Get(), o.StartTimestampUsecs.IsSet()
}

// HasStartTimestampUsecs returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasStartTimestampUsecs() bool {
	if o != nil && o.StartTimestampUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimestampUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimestampUsecs field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetStartTimestampUsecs(v int64) {
	o.StartTimestampUsecs.Set(&v)
}
// SetStartTimestampUsecsNil sets the value for StartTimestampUsecs to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetStartTimestampUsecsNil() {
	o.StartTimestampUsecs.Set(nil)
}

// UnsetStartTimestampUsecs ensures that no value is present for StartTimestampUsecs, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetStartTimestampUsecs() {
	o.StartTimestampUsecs.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetrieveArchiveInfoRetrievedEntity) GetStatus() int32 {
	if o == nil || o.Status.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetrieveArchiveInfoRetrievedEntity) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RetrieveArchiveInfoRetrievedEntity) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *RetrieveArchiveInfoRetrievedEntity) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RetrieveArchiveInfoRetrievedEntity) UnsetStatus() {
	o.Status.Unset()
}

func (o RetrieveArchiveInfoRetrievedEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BytesTransferred.IsSet() {
		toSerialize["bytesTransferred"] = o.BytesTransferred.Get()
	}
	if o.EndTimestampUsecs.IsSet() {
		toSerialize["endTimestampUsecs"] = o.EndTimestampUsecs.Get()
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.LogicalBytesTransferred.IsSet() {
		toSerialize["logicalBytesTransferred"] = o.LogicalBytesTransferred.Get()
	}
	if o.LogicalSizeBytes.IsSet() {
		toSerialize["logicalSizeBytes"] = o.LogicalSizeBytes.Get()
	}
	if o.ProgressMonitorTaskPath.IsSet() {
		toSerialize["progressMonitorTaskPath"] = o.ProgressMonitorTaskPath.Get()
	}
	if o.RelativeSnapshotDir.IsSet() {
		toSerialize["relativeSnapshotDir"] = o.RelativeSnapshotDir.Get()
	}
	if o.StartTimestampUsecs.IsSet() {
		toSerialize["startTimestampUsecs"] = o.StartTimestampUsecs.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRetrieveArchiveInfoRetrievedEntity struct {
	value *RetrieveArchiveInfoRetrievedEntity
	isSet bool
}

func (v NullableRetrieveArchiveInfoRetrievedEntity) Get() *RetrieveArchiveInfoRetrievedEntity {
	return v.value
}

func (v *NullableRetrieveArchiveInfoRetrievedEntity) Set(val *RetrieveArchiveInfoRetrievedEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveArchiveInfoRetrievedEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveArchiveInfoRetrievedEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveArchiveInfoRetrievedEntity(val *RetrieveArchiveInfoRetrievedEntity) *NullableRetrieveArchiveInfoRetrievedEntity {
	return &NullableRetrieveArchiveInfoRetrievedEntity{value: val, isSet: true}
}

func (v NullableRetrieveArchiveInfoRetrievedEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveArchiveInfoRetrievedEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


