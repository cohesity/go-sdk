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

// checks if the TdmObjectTimelineEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TdmObjectTimelineEvents{}

// TdmObjectTimelineEvents Specifies a collection of TDM object's timeline events.
type TdmObjectTimelineEvents struct {
	// Specifies the collection of the timeline events, filtered by the specified criteria.
	Events []TdmObjectTimelineEvent `json:"events,omitempty"`
}

// NewTdmObjectTimelineEvents instantiates a new TdmObjectTimelineEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTdmObjectTimelineEvents() *TdmObjectTimelineEvents {
	this := TdmObjectTimelineEvents{}
	return &this
}

// NewTdmObjectTimelineEventsWithDefaults instantiates a new TdmObjectTimelineEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTdmObjectTimelineEventsWithDefaults() *TdmObjectTimelineEvents {
	this := TdmObjectTimelineEvents{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TdmObjectTimelineEvents) GetEvents() []TdmObjectTimelineEvent {
	if o == nil {
		var ret []TdmObjectTimelineEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TdmObjectTimelineEvents) GetEventsOk() ([]TdmObjectTimelineEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TdmObjectTimelineEvents) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []TdmObjectTimelineEvent and assigns it to the Events field.
func (o *TdmObjectTimelineEvents) SetEvents(v []TdmObjectTimelineEvent) {
	o.Events = v
}

func (o TdmObjectTimelineEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TdmObjectTimelineEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableTdmObjectTimelineEvents struct {
	value *TdmObjectTimelineEvents
	isSet bool
}

func (v NullableTdmObjectTimelineEvents) Get() *TdmObjectTimelineEvents {
	return v.value
}

func (v *NullableTdmObjectTimelineEvents) Set(val *TdmObjectTimelineEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableTdmObjectTimelineEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableTdmObjectTimelineEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTdmObjectTimelineEvents(val *TdmObjectTimelineEvents) *NullableTdmObjectTimelineEvents {
	return &NullableTdmObjectTimelineEvents{value: val, isSet: true}
}

func (v NullableTdmObjectTimelineEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTdmObjectTimelineEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


