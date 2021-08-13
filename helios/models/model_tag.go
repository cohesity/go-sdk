/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// Tag Tag details
type Tag struct {
	// Specifies unique id of the Tag.
	Id NullableString `json:"id,omitempty"`
	// Name of the Tag. Name has to be unique under Namespace.
	Name NullableString `json:"name"`
	// Namespace of the Tag. This is used to filter tags based on application or usecase. For example all tags related to vcenter can be put under one namespace or different departments could have their own namespaces e.g. finance/tag1 or operations/tag2 etc.
	Namespace NullableString `json:"namespace"`
	// Tenant Id to whom the Tag belongs.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Description of the Tag.
	Description NullableString `json:"description,omitempty"`
	// Specifies the timestamp in microseconds since the epoch when this Tag was created.
	CreatedTimeUsecs NullableInt32 `json:"createdTimeUsecs,omitempty"`
	// Specifies the timestamp in microseconds since the epoch when this Tag was last updated.
	LastUpdatedTimeUsecs NullableInt32 `json:"lastUpdatedTimeUsecs,omitempty"`
	// If true, Tag is marked for deletion.
	MarkedForDeletion NullableBool `json:"markedForDeletion,omitempty"`
	// Color of the tag in UI.
	UiColor NullableString `json:"uiColor,omitempty"`
	// Path of the tag for UI nesting purposes.
	UiPathElements []string `json:"uiPathElements,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(name NullableString, namespace NullableString) *Tag {
	this := Tag{}
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Tag) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Tag) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Tag) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Tag) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tag) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name.Set(&v)
}

// GetNamespace returns the Namespace field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Tag) GetNamespace() string {
	if o == nil || o.Namespace.Get() == nil {
		var ret string
		return ret
	}

	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// SetNamespace sets field value
func (o *Tag) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Tag) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Tag) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Tag) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Tag) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Tag) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Tag) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Tag) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Tag) UnsetDescription() {
	o.Description.Unset()
}

// GetCreatedTimeUsecs returns the CreatedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetCreatedTimeUsecs() int32 {
	if o == nil || o.CreatedTimeUsecs.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CreatedTimeUsecs.Get()
}

// GetCreatedTimeUsecsOk returns a tuple with the CreatedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetCreatedTimeUsecsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimeUsecs.Get(), o.CreatedTimeUsecs.IsSet()
}

// HasCreatedTimeUsecs returns a boolean if a field has been set.
func (o *Tag) HasCreatedTimeUsecs() bool {
	if o != nil && o.CreatedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimeUsecs gets a reference to the given NullableInt32 and assigns it to the CreatedTimeUsecs field.
func (o *Tag) SetCreatedTimeUsecs(v int32) {
	o.CreatedTimeUsecs.Set(&v)
}
// SetCreatedTimeUsecsNil sets the value for CreatedTimeUsecs to be an explicit nil
func (o *Tag) SetCreatedTimeUsecsNil() {
	o.CreatedTimeUsecs.Set(nil)
}

// UnsetCreatedTimeUsecs ensures that no value is present for CreatedTimeUsecs, not even an explicit nil
func (o *Tag) UnsetCreatedTimeUsecs() {
	o.CreatedTimeUsecs.Unset()
}

// GetLastUpdatedTimeUsecs returns the LastUpdatedTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetLastUpdatedTimeUsecs() int32 {
	if o == nil || o.LastUpdatedTimeUsecs.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LastUpdatedTimeUsecs.Get()
}

// GetLastUpdatedTimeUsecsOk returns a tuple with the LastUpdatedTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetLastUpdatedTimeUsecsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedTimeUsecs.Get(), o.LastUpdatedTimeUsecs.IsSet()
}

// HasLastUpdatedTimeUsecs returns a boolean if a field has been set.
func (o *Tag) HasLastUpdatedTimeUsecs() bool {
	if o != nil && o.LastUpdatedTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedTimeUsecs gets a reference to the given NullableInt32 and assigns it to the LastUpdatedTimeUsecs field.
func (o *Tag) SetLastUpdatedTimeUsecs(v int32) {
	o.LastUpdatedTimeUsecs.Set(&v)
}
// SetLastUpdatedTimeUsecsNil sets the value for LastUpdatedTimeUsecs to be an explicit nil
func (o *Tag) SetLastUpdatedTimeUsecsNil() {
	o.LastUpdatedTimeUsecs.Set(nil)
}

// UnsetLastUpdatedTimeUsecs ensures that no value is present for LastUpdatedTimeUsecs, not even an explicit nil
func (o *Tag) UnsetLastUpdatedTimeUsecs() {
	o.LastUpdatedTimeUsecs.Unset()
}

// GetMarkedForDeletion returns the MarkedForDeletion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetMarkedForDeletion() bool {
	if o == nil || o.MarkedForDeletion.Get() == nil {
		var ret bool
		return ret
	}
	return *o.MarkedForDeletion.Get()
}

// GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetMarkedForDeletionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MarkedForDeletion.Get(), o.MarkedForDeletion.IsSet()
}

// HasMarkedForDeletion returns a boolean if a field has been set.
func (o *Tag) HasMarkedForDeletion() bool {
	if o != nil && o.MarkedForDeletion.IsSet() {
		return true
	}

	return false
}

// SetMarkedForDeletion gets a reference to the given NullableBool and assigns it to the MarkedForDeletion field.
func (o *Tag) SetMarkedForDeletion(v bool) {
	o.MarkedForDeletion.Set(&v)
}
// SetMarkedForDeletionNil sets the value for MarkedForDeletion to be an explicit nil
func (o *Tag) SetMarkedForDeletionNil() {
	o.MarkedForDeletion.Set(nil)
}

// UnsetMarkedForDeletion ensures that no value is present for MarkedForDeletion, not even an explicit nil
func (o *Tag) UnsetMarkedForDeletion() {
	o.MarkedForDeletion.Unset()
}

// GetUiColor returns the UiColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetUiColor() string {
	if o == nil || o.UiColor.Get() == nil {
		var ret string
		return ret
	}
	return *o.UiColor.Get()
}

// GetUiColorOk returns a tuple with the UiColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetUiColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UiColor.Get(), o.UiColor.IsSet()
}

// HasUiColor returns a boolean if a field has been set.
func (o *Tag) HasUiColor() bool {
	if o != nil && o.UiColor.IsSet() {
		return true
	}

	return false
}

// SetUiColor gets a reference to the given NullableString and assigns it to the UiColor field.
func (o *Tag) SetUiColor(v string) {
	o.UiColor.Set(&v)
}
// SetUiColorNil sets the value for UiColor to be an explicit nil
func (o *Tag) SetUiColorNil() {
	o.UiColor.Set(nil)
}

// UnsetUiColor ensures that no value is present for UiColor, not even an explicit nil
func (o *Tag) UnsetUiColor() {
	o.UiColor.Unset()
}

// GetUiPathElements returns the UiPathElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetUiPathElements() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.UiPathElements
}

// GetUiPathElementsOk returns a tuple with the UiPathElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetUiPathElementsOk() (*[]string, bool) {
	if o == nil || o.UiPathElements == nil {
		return nil, false
	}
	return &o.UiPathElements, true
}

// HasUiPathElements returns a boolean if a field has been set.
func (o *Tag) HasUiPathElements() bool {
	if o != nil && o.UiPathElements != nil {
		return true
	}

	return false
}

// SetUiPathElements gets a reference to the given []string and assigns it to the UiPathElements field.
func (o *Tag) SetUiPathElements(v []string) {
	o.UiPathElements = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CreatedTimeUsecs.IsSet() {
		toSerialize["createdTimeUsecs"] = o.CreatedTimeUsecs.Get()
	}
	if o.LastUpdatedTimeUsecs.IsSet() {
		toSerialize["lastUpdatedTimeUsecs"] = o.LastUpdatedTimeUsecs.Get()
	}
	if o.MarkedForDeletion.IsSet() {
		toSerialize["markedForDeletion"] = o.MarkedForDeletion.Get()
	}
	if o.UiColor.IsSet() {
		toSerialize["uiColor"] = o.UiColor.Get()
	}
	if o.UiPathElements != nil {
		toSerialize["uiPathElements"] = o.UiPathElements
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


