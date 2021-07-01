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

// RestoreSiteParamsDriveItem This will be set in case of partial drive recovery.
type RestoreSiteParamsDriveItem struct {
	// The path of the drive item relative to root.
	DriveItemPath NullableString `json:"driveItemPath,omitempty"`
	// The unique identifier of the item within the Drive.
	Id NullableString `json:"id,omitempty"`
	// Specify if the item is a file or not.
	IsFileItem NullableBool `json:"isFileItem,omitempty"`
}

// NewRestoreSiteParamsDriveItem instantiates a new RestoreSiteParamsDriveItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreSiteParamsDriveItem() *RestoreSiteParamsDriveItem {
	this := RestoreSiteParamsDriveItem{}
	return &this
}

// NewRestoreSiteParamsDriveItemWithDefaults instantiates a new RestoreSiteParamsDriveItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreSiteParamsDriveItemWithDefaults() *RestoreSiteParamsDriveItem {
	this := RestoreSiteParamsDriveItem{}
	return &this
}

// GetDriveItemPath returns the DriveItemPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreSiteParamsDriveItem) GetDriveItemPath() string {
	if o == nil || o.DriveItemPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.DriveItemPath.Get()
}

// GetDriveItemPathOk returns a tuple with the DriveItemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreSiteParamsDriveItem) GetDriveItemPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DriveItemPath.Get(), o.DriveItemPath.IsSet()
}

// HasDriveItemPath returns a boolean if a field has been set.
func (o *RestoreSiteParamsDriveItem) HasDriveItemPath() bool {
	if o != nil && o.DriveItemPath.IsSet() {
		return true
	}

	return false
}

// SetDriveItemPath gets a reference to the given NullableString and assigns it to the DriveItemPath field.
func (o *RestoreSiteParamsDriveItem) SetDriveItemPath(v string) {
	o.DriveItemPath.Set(&v)
}
// SetDriveItemPathNil sets the value for DriveItemPath to be an explicit nil
func (o *RestoreSiteParamsDriveItem) SetDriveItemPathNil() {
	o.DriveItemPath.Set(nil)
}

// UnsetDriveItemPath ensures that no value is present for DriveItemPath, not even an explicit nil
func (o *RestoreSiteParamsDriveItem) UnsetDriveItemPath() {
	o.DriveItemPath.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreSiteParamsDriveItem) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreSiteParamsDriveItem) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RestoreSiteParamsDriveItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *RestoreSiteParamsDriveItem) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RestoreSiteParamsDriveItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RestoreSiteParamsDriveItem) UnsetId() {
	o.Id.Unset()
}

// GetIsFileItem returns the IsFileItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreSiteParamsDriveItem) GetIsFileItem() bool {
	if o == nil || o.IsFileItem.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsFileItem.Get()
}

// GetIsFileItemOk returns a tuple with the IsFileItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreSiteParamsDriveItem) GetIsFileItemOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsFileItem.Get(), o.IsFileItem.IsSet()
}

// HasIsFileItem returns a boolean if a field has been set.
func (o *RestoreSiteParamsDriveItem) HasIsFileItem() bool {
	if o != nil && o.IsFileItem.IsSet() {
		return true
	}

	return false
}

// SetIsFileItem gets a reference to the given NullableBool and assigns it to the IsFileItem field.
func (o *RestoreSiteParamsDriveItem) SetIsFileItem(v bool) {
	o.IsFileItem.Set(&v)
}
// SetIsFileItemNil sets the value for IsFileItem to be an explicit nil
func (o *RestoreSiteParamsDriveItem) SetIsFileItemNil() {
	o.IsFileItem.Set(nil)
}

// UnsetIsFileItem ensures that no value is present for IsFileItem, not even an explicit nil
func (o *RestoreSiteParamsDriveItem) UnsetIsFileItem() {
	o.IsFileItem.Unset()
}

func (o RestoreSiteParamsDriveItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DriveItemPath.IsSet() {
		toSerialize["driveItemPath"] = o.DriveItemPath.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsFileItem.IsSet() {
		toSerialize["isFileItem"] = o.IsFileItem.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreSiteParamsDriveItem struct {
	value *RestoreSiteParamsDriveItem
	isSet bool
}

func (v NullableRestoreSiteParamsDriveItem) Get() *RestoreSiteParamsDriveItem {
	return v.value
}

func (v *NullableRestoreSiteParamsDriveItem) Set(val *RestoreSiteParamsDriveItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreSiteParamsDriveItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreSiteParamsDriveItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreSiteParamsDriveItem(val *RestoreSiteParamsDriveItem) *NullableRestoreSiteParamsDriveItem {
	return &NullableRestoreSiteParamsDriveItem{value: val, isSet: true}
}

func (v NullableRestoreSiteParamsDriveItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreSiteParamsDriveItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


