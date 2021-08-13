/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// CommonTdmObjectTimelineEventsResponseParams Specifies the common parameters for the TDM object timeline events response.
type CommonTdmObjectTimelineEventsResponseParams struct {
	// Specifies the unique ID of the event.
	Id NullableString `json:"id"`
	// Specifies the time (in usecs from epoch) at which the event was created.
	CreatedAt NullableInt64 `json:"createdAt,omitempty"`
	// Specifies the user, who triggered the event.
	CreatedByUser NullableUser `json:"createdByUser,omitempty"`
	// Specifies the current status of the event.
	Status NullableString `json:"status,omitempty"`
	// Specifies the error message if the event is in failed state.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Specifies the description of the event, as provided by the user.
	Description NullableString `json:"description,omitempty"`
	// Specifies the ID of the group this event belongs to. Events with same group ID are considered to be a single timeline for the TDM object. Different group IDs mean different timelines for the TDM object.
	EventGroupId NullableString `json:"eventGroupId,omitempty"`
	// Specifies the action for the event.
	Action NullableString `json:"action"`
}

// NewCommonTdmObjectTimelineEventsResponseParams instantiates a new CommonTdmObjectTimelineEventsResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTdmObjectTimelineEventsResponseParams(id NullableString, action NullableString) *CommonTdmObjectTimelineEventsResponseParams {
	this := CommonTdmObjectTimelineEventsResponseParams{}
	this.Id = id
	this.Action = action
	return &this
}

// NewCommonTdmObjectTimelineEventsResponseParamsWithDefaults instantiates a new CommonTdmObjectTimelineEventsResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTdmObjectTimelineEventsResponseParamsWithDefaults() *CommonTdmObjectTimelineEventsResponseParams {
	this := CommonTdmObjectTimelineEventsResponseParams{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CommonTdmObjectTimelineEventsResponseParams) SetId(v string) {
	o.Id.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedAtOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedByUser() User {
	if o == nil || o.CreatedByUser.Get() == nil {
		var ret User
		return ret
	}
	return *o.CreatedByUser.Get()
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetCreatedByUserOk() (*User, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedByUser.Get(), o.CreatedByUser.IsSet()
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser.IsSet() {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given NullableUser and assigns it to the CreatedByUser field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedByUser(v User) {
	o.CreatedByUser.Set(&v)
}
// SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetCreatedByUserNil() {
	o.CreatedByUser.Set(nil)
}

// UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetCreatedByUser() {
	o.CreatedByUser.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetStatus() {
	o.Status.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetDescription() {
	o.Description.Unset()
}

// GetEventGroupId returns the EventGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTdmObjectTimelineEventsResponseParams) GetEventGroupId() string {
	if o == nil || o.EventGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EventGroupId.Get()
}

// GetEventGroupIdOk returns a tuple with the EventGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetEventGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EventGroupId.Get(), o.EventGroupId.IsSet()
}

// HasEventGroupId returns a boolean if a field has been set.
func (o *CommonTdmObjectTimelineEventsResponseParams) HasEventGroupId() bool {
	if o != nil && o.EventGroupId.IsSet() {
		return true
	}

	return false
}

// SetEventGroupId gets a reference to the given NullableString and assigns it to the EventGroupId field.
func (o *CommonTdmObjectTimelineEventsResponseParams) SetEventGroupId(v string) {
	o.EventGroupId.Set(&v)
}
// SetEventGroupIdNil sets the value for EventGroupId to be an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) SetEventGroupIdNil() {
	o.EventGroupId.Set(nil)
}

// UnsetEventGroupId ensures that no value is present for EventGroupId, not even an explicit nil
func (o *CommonTdmObjectTimelineEventsResponseParams) UnsetEventGroupId() {
	o.EventGroupId.Unset()
}

// GetAction returns the Action field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}

	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTdmObjectTimelineEventsResponseParams) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// SetAction sets field value
func (o *CommonTdmObjectTimelineEventsResponseParams) SetAction(v string) {
	o.Action.Set(&v)
}

func (o CommonTdmObjectTimelineEventsResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.CreatedByUser.IsSet() {
		toSerialize["createdByUser"] = o.CreatedByUser.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EventGroupId.IsSet() {
		toSerialize["eventGroupId"] = o.EventGroupId.Get()
	}
	if true {
		toSerialize["action"] = o.Action.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTdmObjectTimelineEventsResponseParams struct {
	value *CommonTdmObjectTimelineEventsResponseParams
	isSet bool
}

func (v NullableCommonTdmObjectTimelineEventsResponseParams) Get() *CommonTdmObjectTimelineEventsResponseParams {
	return v.value
}

func (v *NullableCommonTdmObjectTimelineEventsResponseParams) Set(val *CommonTdmObjectTimelineEventsResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTdmObjectTimelineEventsResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTdmObjectTimelineEventsResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTdmObjectTimelineEventsResponseParams(val *CommonTdmObjectTimelineEventsResponseParams) *NullableCommonTdmObjectTimelineEventsResponseParams {
	return &NullableCommonTdmObjectTimelineEventsResponseParams{value: val, isSet: true}
}

func (v NullableCommonTdmObjectTimelineEventsResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTdmObjectTimelineEventsResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonTdmObjectTimelineEventsResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}