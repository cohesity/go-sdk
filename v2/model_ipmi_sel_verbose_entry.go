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

// checks if the IpmiSelVerboseEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpmiSelVerboseEntry{}

// IpmiSelVerboseEntry Specifies each entry in the sel verbose response.
type IpmiSelVerboseEntry struct {
	// Specifies additional info about the event.
	EventData NullableString `json:"eventData,omitempty"`
	// Specifies whether the event occurred is assertion or deassertion.
	EventDirection NullableString `json:"eventDirection,omitempty"`
	// Specifies the type of event occurred.
	EventType NullableString `json:"eventType,omitempty"`
	// Specifies the version of the Event Message Format for the record added to SEL.
	EvmRevision NullableString `json:"evmRevision,omitempty"`
	// Corresponds to source of component that generated the record.
	GeneratorId NullableString `json:"generatorId,omitempty"`
	// Specifies a short description corresponding to the sensor event for which record is added to SEL.
	RecordDescription NullableString `json:"recordDescription,omitempty"`
	// Specifies the ID corresponding to record in SEL(System Event Log) for given node.
	RecordId NullableString `json:"recordId,omitempty"`
	// Specifies the time stamp at which the record is added to SEL.
	RecordTimestamp NullableString `json:"recordTimestamp,omitempty"`
	// Specifies the type of SEL record corresponding to the entry.
	RecordType NullableString `json:"recordType,omitempty"`
	// Specifies the sensor number corresponding to the current SEL record.
	SensorNumber NullableString `json:"sensorNumber,omitempty"`
	// Specifies the type of the sensor corresponding to the current SEL record.
	SensorType NullableString `json:"sensorType,omitempty"`
}

// NewIpmiSelVerboseEntry instantiates a new IpmiSelVerboseEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmiSelVerboseEntry() *IpmiSelVerboseEntry {
	this := IpmiSelVerboseEntry{}
	return &this
}

// NewIpmiSelVerboseEntryWithDefaults instantiates a new IpmiSelVerboseEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmiSelVerboseEntryWithDefaults() *IpmiSelVerboseEntry {
	this := IpmiSelVerboseEntry{}
	return &this
}

// GetEventData returns the EventData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetEventData() string {
	if o == nil || IsNil(o.EventData.Get()) {
		var ret string
		return ret
	}
	return *o.EventData.Get()
}

// GetEventDataOk returns a tuple with the EventData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetEventDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventData.Get(), o.EventData.IsSet()
}

// HasEventData returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasEventData() bool {
	if o != nil && o.EventData.IsSet() {
		return true
	}

	return false
}

// SetEventData gets a reference to the given NullableString and assigns it to the EventData field.
func (o *IpmiSelVerboseEntry) SetEventData(v string) {
	o.EventData.Set(&v)
}
// SetEventDataNil sets the value for EventData to be an explicit nil
func (o *IpmiSelVerboseEntry) SetEventDataNil() {
	o.EventData.Set(nil)
}

// UnsetEventData ensures that no value is present for EventData, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetEventData() {
	o.EventData.Unset()
}

// GetEventDirection returns the EventDirection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetEventDirection() string {
	if o == nil || IsNil(o.EventDirection.Get()) {
		var ret string
		return ret
	}
	return *o.EventDirection.Get()
}

// GetEventDirectionOk returns a tuple with the EventDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetEventDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventDirection.Get(), o.EventDirection.IsSet()
}

// HasEventDirection returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasEventDirection() bool {
	if o != nil && o.EventDirection.IsSet() {
		return true
	}

	return false
}

// SetEventDirection gets a reference to the given NullableString and assigns it to the EventDirection field.
func (o *IpmiSelVerboseEntry) SetEventDirection(v string) {
	o.EventDirection.Set(&v)
}
// SetEventDirectionNil sets the value for EventDirection to be an explicit nil
func (o *IpmiSelVerboseEntry) SetEventDirectionNil() {
	o.EventDirection.Set(nil)
}

// UnsetEventDirection ensures that no value is present for EventDirection, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetEventDirection() {
	o.EventDirection.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetEventType() string {
	if o == nil || IsNil(o.EventType.Get()) {
		var ret string
		return ret
	}
	return *o.EventType.Get()
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventType.Get(), o.EventType.IsSet()
}

// HasEventType returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasEventType() bool {
	if o != nil && o.EventType.IsSet() {
		return true
	}

	return false
}

// SetEventType gets a reference to the given NullableString and assigns it to the EventType field.
func (o *IpmiSelVerboseEntry) SetEventType(v string) {
	o.EventType.Set(&v)
}
// SetEventTypeNil sets the value for EventType to be an explicit nil
func (o *IpmiSelVerboseEntry) SetEventTypeNil() {
	o.EventType.Set(nil)
}

// UnsetEventType ensures that no value is present for EventType, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetEventType() {
	o.EventType.Unset()
}

// GetEvmRevision returns the EvmRevision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetEvmRevision() string {
	if o == nil || IsNil(o.EvmRevision.Get()) {
		var ret string
		return ret
	}
	return *o.EvmRevision.Get()
}

// GetEvmRevisionOk returns a tuple with the EvmRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetEvmRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvmRevision.Get(), o.EvmRevision.IsSet()
}

// HasEvmRevision returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasEvmRevision() bool {
	if o != nil && o.EvmRevision.IsSet() {
		return true
	}

	return false
}

// SetEvmRevision gets a reference to the given NullableString and assigns it to the EvmRevision field.
func (o *IpmiSelVerboseEntry) SetEvmRevision(v string) {
	o.EvmRevision.Set(&v)
}
// SetEvmRevisionNil sets the value for EvmRevision to be an explicit nil
func (o *IpmiSelVerboseEntry) SetEvmRevisionNil() {
	o.EvmRevision.Set(nil)
}

// UnsetEvmRevision ensures that no value is present for EvmRevision, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetEvmRevision() {
	o.EvmRevision.Unset()
}

// GetGeneratorId returns the GeneratorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetGeneratorId() string {
	if o == nil || IsNil(o.GeneratorId.Get()) {
		var ret string
		return ret
	}
	return *o.GeneratorId.Get()
}

// GetGeneratorIdOk returns a tuple with the GeneratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetGeneratorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeneratorId.Get(), o.GeneratorId.IsSet()
}

// HasGeneratorId returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasGeneratorId() bool {
	if o != nil && o.GeneratorId.IsSet() {
		return true
	}

	return false
}

// SetGeneratorId gets a reference to the given NullableString and assigns it to the GeneratorId field.
func (o *IpmiSelVerboseEntry) SetGeneratorId(v string) {
	o.GeneratorId.Set(&v)
}
// SetGeneratorIdNil sets the value for GeneratorId to be an explicit nil
func (o *IpmiSelVerboseEntry) SetGeneratorIdNil() {
	o.GeneratorId.Set(nil)
}

// UnsetGeneratorId ensures that no value is present for GeneratorId, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetGeneratorId() {
	o.GeneratorId.Unset()
}

// GetRecordDescription returns the RecordDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetRecordDescription() string {
	if o == nil || IsNil(o.RecordDescription.Get()) {
		var ret string
		return ret
	}
	return *o.RecordDescription.Get()
}

// GetRecordDescriptionOk returns a tuple with the RecordDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetRecordDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordDescription.Get(), o.RecordDescription.IsSet()
}

// HasRecordDescription returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasRecordDescription() bool {
	if o != nil && o.RecordDescription.IsSet() {
		return true
	}

	return false
}

// SetRecordDescription gets a reference to the given NullableString and assigns it to the RecordDescription field.
func (o *IpmiSelVerboseEntry) SetRecordDescription(v string) {
	o.RecordDescription.Set(&v)
}
// SetRecordDescriptionNil sets the value for RecordDescription to be an explicit nil
func (o *IpmiSelVerboseEntry) SetRecordDescriptionNil() {
	o.RecordDescription.Set(nil)
}

// UnsetRecordDescription ensures that no value is present for RecordDescription, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetRecordDescription() {
	o.RecordDescription.Unset()
}

// GetRecordId returns the RecordId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetRecordId() string {
	if o == nil || IsNil(o.RecordId.Get()) {
		var ret string
		return ret
	}
	return *o.RecordId.Get()
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordId.Get(), o.RecordId.IsSet()
}

// HasRecordId returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasRecordId() bool {
	if o != nil && o.RecordId.IsSet() {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given NullableString and assigns it to the RecordId field.
func (o *IpmiSelVerboseEntry) SetRecordId(v string) {
	o.RecordId.Set(&v)
}
// SetRecordIdNil sets the value for RecordId to be an explicit nil
func (o *IpmiSelVerboseEntry) SetRecordIdNil() {
	o.RecordId.Set(nil)
}

// UnsetRecordId ensures that no value is present for RecordId, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetRecordId() {
	o.RecordId.Unset()
}

// GetRecordTimestamp returns the RecordTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetRecordTimestamp() string {
	if o == nil || IsNil(o.RecordTimestamp.Get()) {
		var ret string
		return ret
	}
	return *o.RecordTimestamp.Get()
}

// GetRecordTimestampOk returns a tuple with the RecordTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetRecordTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordTimestamp.Get(), o.RecordTimestamp.IsSet()
}

// HasRecordTimestamp returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasRecordTimestamp() bool {
	if o != nil && o.RecordTimestamp.IsSet() {
		return true
	}

	return false
}

// SetRecordTimestamp gets a reference to the given NullableString and assigns it to the RecordTimestamp field.
func (o *IpmiSelVerboseEntry) SetRecordTimestamp(v string) {
	o.RecordTimestamp.Set(&v)
}
// SetRecordTimestampNil sets the value for RecordTimestamp to be an explicit nil
func (o *IpmiSelVerboseEntry) SetRecordTimestampNil() {
	o.RecordTimestamp.Set(nil)
}

// UnsetRecordTimestamp ensures that no value is present for RecordTimestamp, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetRecordTimestamp() {
	o.RecordTimestamp.Unset()
}

// GetRecordType returns the RecordType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetRecordType() string {
	if o == nil || IsNil(o.RecordType.Get()) {
		var ret string
		return ret
	}
	return *o.RecordType.Get()
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordType.Get(), o.RecordType.IsSet()
}

// HasRecordType returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasRecordType() bool {
	if o != nil && o.RecordType.IsSet() {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given NullableString and assigns it to the RecordType field.
func (o *IpmiSelVerboseEntry) SetRecordType(v string) {
	o.RecordType.Set(&v)
}
// SetRecordTypeNil sets the value for RecordType to be an explicit nil
func (o *IpmiSelVerboseEntry) SetRecordTypeNil() {
	o.RecordType.Set(nil)
}

// UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetRecordType() {
	o.RecordType.Unset()
}

// GetSensorNumber returns the SensorNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetSensorNumber() string {
	if o == nil || IsNil(o.SensorNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SensorNumber.Get()
}

// GetSensorNumberOk returns a tuple with the SensorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetSensorNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SensorNumber.Get(), o.SensorNumber.IsSet()
}

// HasSensorNumber returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasSensorNumber() bool {
	if o != nil && o.SensorNumber.IsSet() {
		return true
	}

	return false
}

// SetSensorNumber gets a reference to the given NullableString and assigns it to the SensorNumber field.
func (o *IpmiSelVerboseEntry) SetSensorNumber(v string) {
	o.SensorNumber.Set(&v)
}
// SetSensorNumberNil sets the value for SensorNumber to be an explicit nil
func (o *IpmiSelVerboseEntry) SetSensorNumberNil() {
	o.SensorNumber.Set(nil)
}

// UnsetSensorNumber ensures that no value is present for SensorNumber, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetSensorNumber() {
	o.SensorNumber.Unset()
}

// GetSensorType returns the SensorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiSelVerboseEntry) GetSensorType() string {
	if o == nil || IsNil(o.SensorType.Get()) {
		var ret string
		return ret
	}
	return *o.SensorType.Get()
}

// GetSensorTypeOk returns a tuple with the SensorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiSelVerboseEntry) GetSensorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SensorType.Get(), o.SensorType.IsSet()
}

// HasSensorType returns a boolean if a field has been set.
func (o *IpmiSelVerboseEntry) HasSensorType() bool {
	if o != nil && o.SensorType.IsSet() {
		return true
	}

	return false
}

// SetSensorType gets a reference to the given NullableString and assigns it to the SensorType field.
func (o *IpmiSelVerboseEntry) SetSensorType(v string) {
	o.SensorType.Set(&v)
}
// SetSensorTypeNil sets the value for SensorType to be an explicit nil
func (o *IpmiSelVerboseEntry) SetSensorTypeNil() {
	o.SensorType.Set(nil)
}

// UnsetSensorType ensures that no value is present for SensorType, not even an explicit nil
func (o *IpmiSelVerboseEntry) UnsetSensorType() {
	o.SensorType.Unset()
}

func (o IpmiSelVerboseEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpmiSelVerboseEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EventData.IsSet() {
		toSerialize["eventData"] = o.EventData.Get()
	}
	if o.EventDirection.IsSet() {
		toSerialize["eventDirection"] = o.EventDirection.Get()
	}
	if o.EventType.IsSet() {
		toSerialize["eventType"] = o.EventType.Get()
	}
	if o.EvmRevision.IsSet() {
		toSerialize["evmRevision"] = o.EvmRevision.Get()
	}
	if o.GeneratorId.IsSet() {
		toSerialize["generatorId"] = o.GeneratorId.Get()
	}
	if o.RecordDescription.IsSet() {
		toSerialize["recordDescription"] = o.RecordDescription.Get()
	}
	if o.RecordId.IsSet() {
		toSerialize["recordId"] = o.RecordId.Get()
	}
	if o.RecordTimestamp.IsSet() {
		toSerialize["recordTimestamp"] = o.RecordTimestamp.Get()
	}
	if o.RecordType.IsSet() {
		toSerialize["recordType"] = o.RecordType.Get()
	}
	if o.SensorNumber.IsSet() {
		toSerialize["sensorNumber"] = o.SensorNumber.Get()
	}
	if o.SensorType.IsSet() {
		toSerialize["sensorType"] = o.SensorType.Get()
	}
	return toSerialize, nil
}

type NullableIpmiSelVerboseEntry struct {
	value *IpmiSelVerboseEntry
	isSet bool
}

func (v NullableIpmiSelVerboseEntry) Get() *IpmiSelVerboseEntry {
	return v.value
}

func (v *NullableIpmiSelVerboseEntry) Set(val *IpmiSelVerboseEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmiSelVerboseEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmiSelVerboseEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmiSelVerboseEntry(val *IpmiSelVerboseEntry) *NullableIpmiSelVerboseEntry {
	return &NullableIpmiSelVerboseEntry{value: val, isSet: true}
}

func (v NullableIpmiSelVerboseEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmiSelVerboseEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


