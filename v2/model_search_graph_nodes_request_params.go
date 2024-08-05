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

// checks if the SearchGraphNodesRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchGraphNodesRequestParams{}

// SearchGraphNodesRequestParams Specifies the request parameters to query nodes in the graph for a given session id.
type SearchGraphNodesRequestParams struct {
	// Specifies the aad adapter specific node filter params.
	AadParams NullableAadGraphNodeFilterParams `json:"aadParams,omitempty"`
	// Filters the nodes based on provided current node display name.
	Name NullableString `json:"name,omitempty"`
	// If set to true only root nodes would be returned. A root node refers to nodes in the graph with no incoming edges. Defaults to false.
	RootOnly NullableBool `json:"rootOnly,omitempty"`
	// Specifies the number of objects to be fetched for the specified pagination cookie.
	Count *int32 `json:"count,omitempty"`
	// If set to false the response will only return name, type and is_root fields filled in each node. If set to true all the attributes for the nodes are also returned. Defaults to true.
	IncludeAttributes NullableBool `json:"includeAttributes,omitempty"`
	// Specifies a cookie which can be passed in by the user in order to retrieve the next page of results.
	PaginationCookie NullableString `json:"paginationCookie,omitempty"`
	// Specifies the id of a Session.
	SessionId string `json:"sessionId"`
}

type _SearchGraphNodesRequestParams SearchGraphNodesRequestParams

// NewSearchGraphNodesRequestParams instantiates a new SearchGraphNodesRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchGraphNodesRequestParams(sessionId string) *SearchGraphNodesRequestParams {
	this := SearchGraphNodesRequestParams{}
	this.SessionId = sessionId
	return &this
}

// NewSearchGraphNodesRequestParamsWithDefaults instantiates a new SearchGraphNodesRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchGraphNodesRequestParamsWithDefaults() *SearchGraphNodesRequestParams {
	this := SearchGraphNodesRequestParams{}
	return &this
}

// GetAadParams returns the AadParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchGraphNodesRequestParams) GetAadParams() AadGraphNodeFilterParams {
	if o == nil || IsNil(o.AadParams.Get()) {
		var ret AadGraphNodeFilterParams
		return ret
	}
	return *o.AadParams.Get()
}

// GetAadParamsOk returns a tuple with the AadParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchGraphNodesRequestParams) GetAadParamsOk() (*AadGraphNodeFilterParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.AadParams.Get(), o.AadParams.IsSet()
}

// HasAadParams returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasAadParams() bool {
	if o != nil && o.AadParams.IsSet() {
		return true
	}

	return false
}

// SetAadParams gets a reference to the given NullableAadGraphNodeFilterParams and assigns it to the AadParams field.
func (o *SearchGraphNodesRequestParams) SetAadParams(v AadGraphNodeFilterParams) {
	o.AadParams.Set(&v)
}
// SetAadParamsNil sets the value for AadParams to be an explicit nil
func (o *SearchGraphNodesRequestParams) SetAadParamsNil() {
	o.AadParams.Set(nil)
}

// UnsetAadParams ensures that no value is present for AadParams, not even an explicit nil
func (o *SearchGraphNodesRequestParams) UnsetAadParams() {
	o.AadParams.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchGraphNodesRequestParams) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchGraphNodesRequestParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SearchGraphNodesRequestParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SearchGraphNodesRequestParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SearchGraphNodesRequestParams) UnsetName() {
	o.Name.Unset()
}

// GetRootOnly returns the RootOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchGraphNodesRequestParams) GetRootOnly() bool {
	if o == nil || IsNil(o.RootOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.RootOnly.Get()
}

// GetRootOnlyOk returns a tuple with the RootOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchGraphNodesRequestParams) GetRootOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootOnly.Get(), o.RootOnly.IsSet()
}

// HasRootOnly returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasRootOnly() bool {
	if o != nil && o.RootOnly.IsSet() {
		return true
	}

	return false
}

// SetRootOnly gets a reference to the given NullableBool and assigns it to the RootOnly field.
func (o *SearchGraphNodesRequestParams) SetRootOnly(v bool) {
	o.RootOnly.Set(&v)
}
// SetRootOnlyNil sets the value for RootOnly to be an explicit nil
func (o *SearchGraphNodesRequestParams) SetRootOnlyNil() {
	o.RootOnly.Set(nil)
}

// UnsetRootOnly ensures that no value is present for RootOnly, not even an explicit nil
func (o *SearchGraphNodesRequestParams) UnsetRootOnly() {
	o.RootOnly.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchGraphNodesRequestParams) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchGraphNodesRequestParams) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SearchGraphNodesRequestParams) SetCount(v int32) {
	o.Count = &v
}

// GetIncludeAttributes returns the IncludeAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchGraphNodesRequestParams) GetIncludeAttributes() bool {
	if o == nil || IsNil(o.IncludeAttributes.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeAttributes.Get()
}

// GetIncludeAttributesOk returns a tuple with the IncludeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchGraphNodesRequestParams) GetIncludeAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeAttributes.Get(), o.IncludeAttributes.IsSet()
}

// HasIncludeAttributes returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasIncludeAttributes() bool {
	if o != nil && o.IncludeAttributes.IsSet() {
		return true
	}

	return false
}

// SetIncludeAttributes gets a reference to the given NullableBool and assigns it to the IncludeAttributes field.
func (o *SearchGraphNodesRequestParams) SetIncludeAttributes(v bool) {
	o.IncludeAttributes.Set(&v)
}
// SetIncludeAttributesNil sets the value for IncludeAttributes to be an explicit nil
func (o *SearchGraphNodesRequestParams) SetIncludeAttributesNil() {
	o.IncludeAttributes.Set(nil)
}

// UnsetIncludeAttributes ensures that no value is present for IncludeAttributes, not even an explicit nil
func (o *SearchGraphNodesRequestParams) UnsetIncludeAttributes() {
	o.IncludeAttributes.Unset()
}

// GetPaginationCookie returns the PaginationCookie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchGraphNodesRequestParams) GetPaginationCookie() string {
	if o == nil || IsNil(o.PaginationCookie.Get()) {
		var ret string
		return ret
	}
	return *o.PaginationCookie.Get()
}

// GetPaginationCookieOk returns a tuple with the PaginationCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchGraphNodesRequestParams) GetPaginationCookieOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaginationCookie.Get(), o.PaginationCookie.IsSet()
}

// HasPaginationCookie returns a boolean if a field has been set.
func (o *SearchGraphNodesRequestParams) HasPaginationCookie() bool {
	if o != nil && o.PaginationCookie.IsSet() {
		return true
	}

	return false
}

// SetPaginationCookie gets a reference to the given NullableString and assigns it to the PaginationCookie field.
func (o *SearchGraphNodesRequestParams) SetPaginationCookie(v string) {
	o.PaginationCookie.Set(&v)
}
// SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil
func (o *SearchGraphNodesRequestParams) SetPaginationCookieNil() {
	o.PaginationCookie.Set(nil)
}

// UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
func (o *SearchGraphNodesRequestParams) UnsetPaginationCookie() {
	o.PaginationCookie.Unset()
}

// GetSessionId returns the SessionId field value
func (o *SearchGraphNodesRequestParams) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *SearchGraphNodesRequestParams) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *SearchGraphNodesRequestParams) SetSessionId(v string) {
	o.SessionId = v
}

func (o SearchGraphNodesRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchGraphNodesRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AadParams.IsSet() {
		toSerialize["aadParams"] = o.AadParams.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.RootOnly.IsSet() {
		toSerialize["rootOnly"] = o.RootOnly.Get()
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.IncludeAttributes.IsSet() {
		toSerialize["includeAttributes"] = o.IncludeAttributes.Get()
	}
	if o.PaginationCookie.IsSet() {
		toSerialize["paginationCookie"] = o.PaginationCookie.Get()
	}
	toSerialize["sessionId"] = o.SessionId
	return toSerialize, nil
}

func (o *SearchGraphNodesRequestParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessionId",
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

	varSearchGraphNodesRequestParams := _SearchGraphNodesRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchGraphNodesRequestParams)

	if err != nil {
		return err
	}

	*o = SearchGraphNodesRequestParams(varSearchGraphNodesRequestParams)

	return err
}

type NullableSearchGraphNodesRequestParams struct {
	value *SearchGraphNodesRequestParams
	isSet bool
}

func (v NullableSearchGraphNodesRequestParams) Get() *SearchGraphNodesRequestParams {
	return v.value
}

func (v *NullableSearchGraphNodesRequestParams) Set(val *SearchGraphNodesRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchGraphNodesRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchGraphNodesRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchGraphNodesRequestParams(val *SearchGraphNodesRequestParams) *NullableSearchGraphNodesRequestParams {
	return &NullableSearchGraphNodesRequestParams{value: val, isSet: true}
}

func (v NullableSearchGraphNodesRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchGraphNodesRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


