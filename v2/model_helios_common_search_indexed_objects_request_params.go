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

// checks if the HeliosCommonSearchIndexedObjectsRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosCommonSearchIndexedObjectsRequestParams{}

// HeliosCommonSearchIndexedObjectsRequestParams Specifies the common params to search for global indexed objects.
type HeliosCommonSearchIndexedObjectsRequestParams struct {
	// List of Clusters Identifiers to filter from. The format is clusterId:clusterIncarnationId.
	ClusterIdentifiers []string `json:"clusterIdentifiers,omitempty"`
	// Specifies the number of indexed objects to be fetched.
	Count NullableInt32 `json:"count,omitempty"`
	// Specifies the object type to be searched for.
	ObjectType NullableString `json:"objectType"`
	// List of Regions to filter from.
	RegionIds []string `json:"regionIds,omitempty"`
	// Specifies a list of source UUIDs. Only matches found in these sources will be returned.
	SourceUUIDs []string `json:"sourceUUIDs,omitempty"`
}

type _HeliosCommonSearchIndexedObjectsRequestParams HeliosCommonSearchIndexedObjectsRequestParams

// NewHeliosCommonSearchIndexedObjectsRequestParams instantiates a new HeliosCommonSearchIndexedObjectsRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosCommonSearchIndexedObjectsRequestParams(objectType NullableString) *HeliosCommonSearchIndexedObjectsRequestParams {
	this := HeliosCommonSearchIndexedObjectsRequestParams{}
	this.ObjectType = objectType
	return &this
}

// NewHeliosCommonSearchIndexedObjectsRequestParamsWithDefaults instantiates a new HeliosCommonSearchIndexedObjectsRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosCommonSearchIndexedObjectsRequestParamsWithDefaults() *HeliosCommonSearchIndexedObjectsRequestParams {
	this := HeliosCommonSearchIndexedObjectsRequestParams{}
	return &this
}

// GetClusterIdentifiers returns the ClusterIdentifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetClusterIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ClusterIdentifiers
}

// GetClusterIdentifiersOk returns a tuple with the ClusterIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetClusterIdentifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.ClusterIdentifiers) {
		return nil, false
	}
	return o.ClusterIdentifiers, true
}

// HasClusterIdentifiers returns a boolean if a field has been set.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasClusterIdentifiers() bool {
	if o != nil && !IsNil(o.ClusterIdentifiers) {
		return true
	}

	return false
}

// SetClusterIdentifiers gets a reference to the given []string and assigns it to the ClusterIdentifiers field.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetClusterIdentifiers(v []string) {
	o.ClusterIdentifiers = v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetCount() int32 {
	if o == nil || IsNil(o.Count.Get()) {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnsetCount() {
	o.Count.Unset()
}

// GetObjectType returns the ObjectType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetObjectType() string {
	if o == nil || o.ObjectType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObjectType.Get()
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectType.Get(), o.ObjectType.IsSet()
}

// SetObjectType sets field value
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetObjectType(v string) {
	o.ObjectType.Set(&v)
}

// GetRegionIds returns the RegionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetRegionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RegionIds
}

// GetRegionIdsOk returns a tuple with the RegionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetRegionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RegionIds) {
		return nil, false
	}
	return o.RegionIds, true
}

// HasRegionIds returns a boolean if a field has been set.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasRegionIds() bool {
	if o != nil && !IsNil(o.RegionIds) {
		return true
	}

	return false
}

// SetRegionIds gets a reference to the given []string and assigns it to the RegionIds field.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetRegionIds(v []string) {
	o.RegionIds = v
}

// GetSourceUUIDs returns the SourceUUIDs field value if set, zero value otherwise.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetSourceUUIDs() []string {
	if o == nil || IsNil(o.SourceUUIDs) {
		var ret []string
		return ret
	}
	return o.SourceUUIDs
}

// GetSourceUUIDsOk returns a tuple with the SourceUUIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) GetSourceUUIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceUUIDs) {
		return nil, false
	}
	return o.SourceUUIDs, true
}

// HasSourceUUIDs returns a boolean if a field has been set.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) HasSourceUUIDs() bool {
	if o != nil && !IsNil(o.SourceUUIDs) {
		return true
	}

	return false
}

// SetSourceUUIDs gets a reference to the given []string and assigns it to the SourceUUIDs field.
func (o *HeliosCommonSearchIndexedObjectsRequestParams) SetSourceUUIDs(v []string) {
	o.SourceUUIDs = v
}

func (o HeliosCommonSearchIndexedObjectsRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosCommonSearchIndexedObjectsRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterIdentifiers != nil {
		toSerialize["clusterIdentifiers"] = o.ClusterIdentifiers
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	toSerialize["objectType"] = o.ObjectType.Get()
	if o.RegionIds != nil {
		toSerialize["regionIds"] = o.RegionIds
	}
	if !IsNil(o.SourceUUIDs) {
		toSerialize["sourceUUIDs"] = o.SourceUUIDs
	}
	return toSerialize, nil
}

func (o *HeliosCommonSearchIndexedObjectsRequestParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectType",
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

	varHeliosCommonSearchIndexedObjectsRequestParams := _HeliosCommonSearchIndexedObjectsRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeliosCommonSearchIndexedObjectsRequestParams)

	if err != nil {
		return err
	}

	*o = HeliosCommonSearchIndexedObjectsRequestParams(varHeliosCommonSearchIndexedObjectsRequestParams)

	return err
}

type NullableHeliosCommonSearchIndexedObjectsRequestParams struct {
	value *HeliosCommonSearchIndexedObjectsRequestParams
	isSet bool
}

func (v NullableHeliosCommonSearchIndexedObjectsRequestParams) Get() *HeliosCommonSearchIndexedObjectsRequestParams {
	return v.value
}

func (v *NullableHeliosCommonSearchIndexedObjectsRequestParams) Set(val *HeliosCommonSearchIndexedObjectsRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosCommonSearchIndexedObjectsRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosCommonSearchIndexedObjectsRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosCommonSearchIndexedObjectsRequestParams(val *HeliosCommonSearchIndexedObjectsRequestParams) *NullableHeliosCommonSearchIndexedObjectsRequestParams {
	return &NullableHeliosCommonSearchIndexedObjectsRequestParams{value: val, isSet: true}
}

func (v NullableHeliosCommonSearchIndexedObjectsRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosCommonSearchIndexedObjectsRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


