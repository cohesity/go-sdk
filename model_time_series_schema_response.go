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

// TimeSeriesSchemaResponse Specifies the Apollo schema to list the data point.
type TimeSeriesSchemaResponse struct {
	// Specifies the list of the schema info for an entity.
	SchemaInfoList []SchemaInfo `json:"schemaInfoList,omitempty"`
}

// NewTimeSeriesSchemaResponse instantiates a new TimeSeriesSchemaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesSchemaResponse() *TimeSeriesSchemaResponse {
	this := TimeSeriesSchemaResponse{}
	return &this
}

// NewTimeSeriesSchemaResponseWithDefaults instantiates a new TimeSeriesSchemaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesSchemaResponseWithDefaults() *TimeSeriesSchemaResponse {
	this := TimeSeriesSchemaResponse{}
	return &this
}

// GetSchemaInfoList returns the SchemaInfoList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeSeriesSchemaResponse) GetSchemaInfoList() []SchemaInfo {
	if o == nil  {
		var ret []SchemaInfo
		return ret
	}
	return o.SchemaInfoList
}

// GetSchemaInfoListOk returns a tuple with the SchemaInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeSeriesSchemaResponse) GetSchemaInfoListOk() (*[]SchemaInfo, bool) {
	if o == nil || o.SchemaInfoList == nil {
		return nil, false
	}
	return &o.SchemaInfoList, true
}

// HasSchemaInfoList returns a boolean if a field has been set.
func (o *TimeSeriesSchemaResponse) HasSchemaInfoList() bool {
	if o != nil && o.SchemaInfoList != nil {
		return true
	}

	return false
}

// SetSchemaInfoList gets a reference to the given []SchemaInfo and assigns it to the SchemaInfoList field.
func (o *TimeSeriesSchemaResponse) SetSchemaInfoList(v []SchemaInfo) {
	o.SchemaInfoList = v
}

func (o TimeSeriesSchemaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaInfoList != nil {
		toSerialize["schemaInfoList"] = o.SchemaInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSeriesSchemaResponse struct {
	value *TimeSeriesSchemaResponse
	isSet bool
}

func (v NullableTimeSeriesSchemaResponse) Get() *TimeSeriesSchemaResponse {
	return v.value
}

func (v *NullableTimeSeriesSchemaResponse) Set(val *TimeSeriesSchemaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesSchemaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesSchemaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesSchemaResponse(val *TimeSeriesSchemaResponse) *NullableTimeSeriesSchemaResponse {
	return &NullableTimeSeriesSchemaResponse{value: val, isSet: true}
}

func (v NullableTimeSeriesSchemaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesSchemaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


