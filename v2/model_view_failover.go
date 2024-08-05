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

// checks if the ViewFailover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewFailover{}

// ViewFailover Specifies the failover status of a view.
type ViewFailover struct {
	// Specifies if the view is ready for failover.
	IsFailoverReady NullableBool `json:"isFailoverReady,omitempty"`
	// Specifies the remote cluster id.
	RemoteClusterId NullableInt64 `json:"remoteClusterId,omitempty"`
	// Specifies the remote cluster incarnation id.
	RemoteClusterIncarnationId NullableInt64 `json:"remoteClusterIncarnationId,omitempty"`
	// Specifies the remote view id.
	RemoteViewId NullableInt64 `json:"remoteViewId,omitempty"`
	// View uid.
	ViewUid NullableString `json:"viewUid,omitempty"`
}

// NewViewFailover instantiates a new ViewFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewFailover() *ViewFailover {
	this := ViewFailover{}
	return &this
}

// NewViewFailoverWithDefaults instantiates a new ViewFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewFailoverWithDefaults() *ViewFailover {
	this := ViewFailover{}
	return &this
}

// GetIsFailoverReady returns the IsFailoverReady field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewFailover) GetIsFailoverReady() bool {
	if o == nil || IsNil(o.IsFailoverReady.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFailoverReady.Get()
}

// GetIsFailoverReadyOk returns a tuple with the IsFailoverReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewFailover) GetIsFailoverReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFailoverReady.Get(), o.IsFailoverReady.IsSet()
}

// HasIsFailoverReady returns a boolean if a field has been set.
func (o *ViewFailover) HasIsFailoverReady() bool {
	if o != nil && o.IsFailoverReady.IsSet() {
		return true
	}

	return false
}

// SetIsFailoverReady gets a reference to the given NullableBool and assigns it to the IsFailoverReady field.
func (o *ViewFailover) SetIsFailoverReady(v bool) {
	o.IsFailoverReady.Set(&v)
}
// SetIsFailoverReadyNil sets the value for IsFailoverReady to be an explicit nil
func (o *ViewFailover) SetIsFailoverReadyNil() {
	o.IsFailoverReady.Set(nil)
}

// UnsetIsFailoverReady ensures that no value is present for IsFailoverReady, not even an explicit nil
func (o *ViewFailover) UnsetIsFailoverReady() {
	o.IsFailoverReady.Unset()
}

// GetRemoteClusterId returns the RemoteClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewFailover) GetRemoteClusterId() int64 {
	if o == nil || IsNil(o.RemoteClusterId.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoteClusterId.Get()
}

// GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewFailover) GetRemoteClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteClusterId.Get(), o.RemoteClusterId.IsSet()
}

// HasRemoteClusterId returns a boolean if a field has been set.
func (o *ViewFailover) HasRemoteClusterId() bool {
	if o != nil && o.RemoteClusterId.IsSet() {
		return true
	}

	return false
}

// SetRemoteClusterId gets a reference to the given NullableInt64 and assigns it to the RemoteClusterId field.
func (o *ViewFailover) SetRemoteClusterId(v int64) {
	o.RemoteClusterId.Set(&v)
}
// SetRemoteClusterIdNil sets the value for RemoteClusterId to be an explicit nil
func (o *ViewFailover) SetRemoteClusterIdNil() {
	o.RemoteClusterId.Set(nil)
}

// UnsetRemoteClusterId ensures that no value is present for RemoteClusterId, not even an explicit nil
func (o *ViewFailover) UnsetRemoteClusterId() {
	o.RemoteClusterId.Unset()
}

// GetRemoteClusterIncarnationId returns the RemoteClusterIncarnationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewFailover) GetRemoteClusterIncarnationId() int64 {
	if o == nil || IsNil(o.RemoteClusterIncarnationId.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoteClusterIncarnationId.Get()
}

// GetRemoteClusterIncarnationIdOk returns a tuple with the RemoteClusterIncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewFailover) GetRemoteClusterIncarnationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteClusterIncarnationId.Get(), o.RemoteClusterIncarnationId.IsSet()
}

// HasRemoteClusterIncarnationId returns a boolean if a field has been set.
func (o *ViewFailover) HasRemoteClusterIncarnationId() bool {
	if o != nil && o.RemoteClusterIncarnationId.IsSet() {
		return true
	}

	return false
}

// SetRemoteClusterIncarnationId gets a reference to the given NullableInt64 and assigns it to the RemoteClusterIncarnationId field.
func (o *ViewFailover) SetRemoteClusterIncarnationId(v int64) {
	o.RemoteClusterIncarnationId.Set(&v)
}
// SetRemoteClusterIncarnationIdNil sets the value for RemoteClusterIncarnationId to be an explicit nil
func (o *ViewFailover) SetRemoteClusterIncarnationIdNil() {
	o.RemoteClusterIncarnationId.Set(nil)
}

// UnsetRemoteClusterIncarnationId ensures that no value is present for RemoteClusterIncarnationId, not even an explicit nil
func (o *ViewFailover) UnsetRemoteClusterIncarnationId() {
	o.RemoteClusterIncarnationId.Unset()
}

// GetRemoteViewId returns the RemoteViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewFailover) GetRemoteViewId() int64 {
	if o == nil || IsNil(o.RemoteViewId.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoteViewId.Get()
}

// GetRemoteViewIdOk returns a tuple with the RemoteViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewFailover) GetRemoteViewIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteViewId.Get(), o.RemoteViewId.IsSet()
}

// HasRemoteViewId returns a boolean if a field has been set.
func (o *ViewFailover) HasRemoteViewId() bool {
	if o != nil && o.RemoteViewId.IsSet() {
		return true
	}

	return false
}

// SetRemoteViewId gets a reference to the given NullableInt64 and assigns it to the RemoteViewId field.
func (o *ViewFailover) SetRemoteViewId(v int64) {
	o.RemoteViewId.Set(&v)
}
// SetRemoteViewIdNil sets the value for RemoteViewId to be an explicit nil
func (o *ViewFailover) SetRemoteViewIdNil() {
	o.RemoteViewId.Set(nil)
}

// UnsetRemoteViewId ensures that no value is present for RemoteViewId, not even an explicit nil
func (o *ViewFailover) UnsetRemoteViewId() {
	o.RemoteViewId.Unset()
}

// GetViewUid returns the ViewUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewFailover) GetViewUid() string {
	if o == nil || IsNil(o.ViewUid.Get()) {
		var ret string
		return ret
	}
	return *o.ViewUid.Get()
}

// GetViewUidOk returns a tuple with the ViewUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewFailover) GetViewUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewUid.Get(), o.ViewUid.IsSet()
}

// HasViewUid returns a boolean if a field has been set.
func (o *ViewFailover) HasViewUid() bool {
	if o != nil && o.ViewUid.IsSet() {
		return true
	}

	return false
}

// SetViewUid gets a reference to the given NullableString and assigns it to the ViewUid field.
func (o *ViewFailover) SetViewUid(v string) {
	o.ViewUid.Set(&v)
}
// SetViewUidNil sets the value for ViewUid to be an explicit nil
func (o *ViewFailover) SetViewUidNil() {
	o.ViewUid.Set(nil)
}

// UnsetViewUid ensures that no value is present for ViewUid, not even an explicit nil
func (o *ViewFailover) UnsetViewUid() {
	o.ViewUid.Unset()
}

func (o ViewFailover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewFailover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsFailoverReady.IsSet() {
		toSerialize["isFailoverReady"] = o.IsFailoverReady.Get()
	}
	if o.RemoteClusterId.IsSet() {
		toSerialize["remoteClusterId"] = o.RemoteClusterId.Get()
	}
	if o.RemoteClusterIncarnationId.IsSet() {
		toSerialize["remoteClusterIncarnationId"] = o.RemoteClusterIncarnationId.Get()
	}
	if o.RemoteViewId.IsSet() {
		toSerialize["remoteViewId"] = o.RemoteViewId.Get()
	}
	if o.ViewUid.IsSet() {
		toSerialize["viewUid"] = o.ViewUid.Get()
	}
	return toSerialize, nil
}

type NullableViewFailover struct {
	value *ViewFailover
	isSet bool
}

func (v NullableViewFailover) Get() *ViewFailover {
	return v.value
}

func (v *NullableViewFailover) Set(val *ViewFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableViewFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableViewFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewFailover(val *ViewFailover) *NullableViewFailover {
	return &NullableViewFailover{value: val, isSet: true}
}

func (v NullableViewFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


