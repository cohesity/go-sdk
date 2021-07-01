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

// AlertCategoryName AlertCategoryName returns alert category and its public facing string.
type AlertCategoryName struct {
	// Specifies alert category. Specifies the category of an Alert. kDisk - Alerts that are related to Disk. kNode - Alerts that are related to Node. kCluster - Alerts that are related to Cluster. kNodeHealth - Alerts that are related to Node Health. kClusterHealth - Alerts that are related to Cluster Health. kBackupRestore - Alerts that are related to Backup/Restore. kEncryption - Alerts that are related to Encryption. kArchivalRestore - Alerts that are related to Archival/Restore. kRemoteReplication - Alerts that are related to Remote Replication. kQuota - Alerts that are related to Quota. kLicense - Alerts that are related to License. kHeliosProActiveWellness - Alerts that are related to Helios ProActive Wellness. kHeliosAnalyticsJobs - Alerts that are related to Helios Analytics Jobs. kHeliosSignatureJobs - Alerts that are related to Helios Signature Jobs. kSecurity - Alerts that are related to Security. kAppsInfra - Alerts that are related to applications infra. kAntivirus - Alerts that are related to antivirus. kArchivalCopy - Alerts that are related to archival copies.
	Category NullableString `json:"category,omitempty"`
	// Specifies public facing string for alert enums.
	Name NullableString `json:"name,omitempty"`
}

// NewAlertCategoryName instantiates a new AlertCategoryName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCategoryName() *AlertCategoryName {
	this := AlertCategoryName{}
	return &this
}

// NewAlertCategoryNameWithDefaults instantiates a new AlertCategoryName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertCategoryNameWithDefaults() *AlertCategoryName {
	this := AlertCategoryName{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertCategoryName) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertCategoryName) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *AlertCategoryName) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *AlertCategoryName) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *AlertCategoryName) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *AlertCategoryName) UnsetCategory() {
	o.Category.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertCategoryName) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertCategoryName) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AlertCategoryName) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AlertCategoryName) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AlertCategoryName) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AlertCategoryName) UnsetName() {
	o.Name.Unset()
}

func (o AlertCategoryName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAlertCategoryName struct {
	value *AlertCategoryName
	isSet bool
}

func (v NullableAlertCategoryName) Get() *AlertCategoryName {
	return v.value
}

func (v *NullableAlertCategoryName) Set(val *AlertCategoryName) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCategoryName) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCategoryName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCategoryName(val *AlertCategoryName) *NullableAlertCategoryName {
	return &NullableAlertCategoryName{value: val, isSet: true}
}

func (v NullableAlertCategoryName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCategoryName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


