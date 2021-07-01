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

// RdsParams Specifies rds params for the restore operation.
type RdsParams struct {
	// Entity representing the availability zone to use while restoring the DB.
	AvailabilityZoneId NullableInt64 `json:"availabilityZoneId,omitempty"`
	// The DB instance identifier to use for the restored DB. This field is required.
	DbInstanceId NullableString `json:"dbInstanceId"`
	// Entity representing the RDS option group to use while restoring the DB.
	DbOptionGroupId NullableInt64 `json:"dbOptionGroupId,omitempty"`
	// Entity representing the RDS parameter group to use while restoring the DB.
	DbParameterGroupId NullableInt64 `json:"dbParameterGroupId,omitempty"`
	// Port to use for the DB in the restored RDS instance.
	DbPort NullableInt32 `json:"dbPort,omitempty"`
	// Whether to enable auto minor version upgrade in the restored DB.
	EnableAutoMinorVersionUpgrade NullableBool `json:"enableAutoMinorVersionUpgrade,omitempty"`
	// Whether to enable copying of tags to snapshots of the DB.
	EnableCopyTagsToSnapshots NullableBool `json:"enableCopyTagsToSnapshots,omitempty"`
	// Whether to enable IAM authentication for the DB.
	EnableDbAuthentication NullableBool `json:"enableDbAuthentication,omitempty"`
	// Whether this DB will be publicly accessible or not.
	EnablePublicAccessibility NullableBool `json:"enablePublicAccessibility,omitempty"`
	// Whether this is a multi-az deployment or not.
	IsMultiAzDeployment NullableBool `json:"isMultiAzDeployment,omitempty"`
}

// NewRdsParams instantiates a new RdsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsParams(dbInstanceId NullableString) *RdsParams {
	this := RdsParams{}
	this.DbInstanceId = dbInstanceId
	return &this
}

// NewRdsParamsWithDefaults instantiates a new RdsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsParamsWithDefaults() *RdsParams {
	this := RdsParams{}
	return &this
}

// GetAvailabilityZoneId returns the AvailabilityZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetAvailabilityZoneId() int64 {
	if o == nil || o.AvailabilityZoneId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvailabilityZoneId.Get()
}

// GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetAvailabilityZoneIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityZoneId.Get(), o.AvailabilityZoneId.IsSet()
}

// HasAvailabilityZoneId returns a boolean if a field has been set.
func (o *RdsParams) HasAvailabilityZoneId() bool {
	if o != nil && o.AvailabilityZoneId.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityZoneId gets a reference to the given NullableInt64 and assigns it to the AvailabilityZoneId field.
func (o *RdsParams) SetAvailabilityZoneId(v int64) {
	o.AvailabilityZoneId.Set(&v)
}
// SetAvailabilityZoneIdNil sets the value for AvailabilityZoneId to be an explicit nil
func (o *RdsParams) SetAvailabilityZoneIdNil() {
	o.AvailabilityZoneId.Set(nil)
}

// UnsetAvailabilityZoneId ensures that no value is present for AvailabilityZoneId, not even an explicit nil
func (o *RdsParams) UnsetAvailabilityZoneId() {
	o.AvailabilityZoneId.Unset()
}

// GetDbInstanceId returns the DbInstanceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RdsParams) GetDbInstanceId() string {
	if o == nil || o.DbInstanceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DbInstanceId.Get()
}

// GetDbInstanceIdOk returns a tuple with the DbInstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetDbInstanceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DbInstanceId.Get(), o.DbInstanceId.IsSet()
}

// SetDbInstanceId sets field value
func (o *RdsParams) SetDbInstanceId(v string) {
	o.DbInstanceId.Set(&v)
}

// GetDbOptionGroupId returns the DbOptionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetDbOptionGroupId() int64 {
	if o == nil || o.DbOptionGroupId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbOptionGroupId.Get()
}

// GetDbOptionGroupIdOk returns a tuple with the DbOptionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetDbOptionGroupIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DbOptionGroupId.Get(), o.DbOptionGroupId.IsSet()
}

// HasDbOptionGroupId returns a boolean if a field has been set.
func (o *RdsParams) HasDbOptionGroupId() bool {
	if o != nil && o.DbOptionGroupId.IsSet() {
		return true
	}

	return false
}

// SetDbOptionGroupId gets a reference to the given NullableInt64 and assigns it to the DbOptionGroupId field.
func (o *RdsParams) SetDbOptionGroupId(v int64) {
	o.DbOptionGroupId.Set(&v)
}
// SetDbOptionGroupIdNil sets the value for DbOptionGroupId to be an explicit nil
func (o *RdsParams) SetDbOptionGroupIdNil() {
	o.DbOptionGroupId.Set(nil)
}

// UnsetDbOptionGroupId ensures that no value is present for DbOptionGroupId, not even an explicit nil
func (o *RdsParams) UnsetDbOptionGroupId() {
	o.DbOptionGroupId.Unset()
}

// GetDbParameterGroupId returns the DbParameterGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetDbParameterGroupId() int64 {
	if o == nil || o.DbParameterGroupId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DbParameterGroupId.Get()
}

// GetDbParameterGroupIdOk returns a tuple with the DbParameterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetDbParameterGroupIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DbParameterGroupId.Get(), o.DbParameterGroupId.IsSet()
}

// HasDbParameterGroupId returns a boolean if a field has been set.
func (o *RdsParams) HasDbParameterGroupId() bool {
	if o != nil && o.DbParameterGroupId.IsSet() {
		return true
	}

	return false
}

// SetDbParameterGroupId gets a reference to the given NullableInt64 and assigns it to the DbParameterGroupId field.
func (o *RdsParams) SetDbParameterGroupId(v int64) {
	o.DbParameterGroupId.Set(&v)
}
// SetDbParameterGroupIdNil sets the value for DbParameterGroupId to be an explicit nil
func (o *RdsParams) SetDbParameterGroupIdNil() {
	o.DbParameterGroupId.Set(nil)
}

// UnsetDbParameterGroupId ensures that no value is present for DbParameterGroupId, not even an explicit nil
func (o *RdsParams) UnsetDbParameterGroupId() {
	o.DbParameterGroupId.Unset()
}

// GetDbPort returns the DbPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetDbPort() int32 {
	if o == nil || o.DbPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DbPort.Get()
}

// GetDbPortOk returns a tuple with the DbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetDbPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DbPort.Get(), o.DbPort.IsSet()
}

// HasDbPort returns a boolean if a field has been set.
func (o *RdsParams) HasDbPort() bool {
	if o != nil && o.DbPort.IsSet() {
		return true
	}

	return false
}

// SetDbPort gets a reference to the given NullableInt32 and assigns it to the DbPort field.
func (o *RdsParams) SetDbPort(v int32) {
	o.DbPort.Set(&v)
}
// SetDbPortNil sets the value for DbPort to be an explicit nil
func (o *RdsParams) SetDbPortNil() {
	o.DbPort.Set(nil)
}

// UnsetDbPort ensures that no value is present for DbPort, not even an explicit nil
func (o *RdsParams) UnsetDbPort() {
	o.DbPort.Unset()
}

// GetEnableAutoMinorVersionUpgrade returns the EnableAutoMinorVersionUpgrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetEnableAutoMinorVersionUpgrade() bool {
	if o == nil || o.EnableAutoMinorVersionUpgrade.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableAutoMinorVersionUpgrade.Get()
}

// GetEnableAutoMinorVersionUpgradeOk returns a tuple with the EnableAutoMinorVersionUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetEnableAutoMinorVersionUpgradeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableAutoMinorVersionUpgrade.Get(), o.EnableAutoMinorVersionUpgrade.IsSet()
}

// HasEnableAutoMinorVersionUpgrade returns a boolean if a field has been set.
func (o *RdsParams) HasEnableAutoMinorVersionUpgrade() bool {
	if o != nil && o.EnableAutoMinorVersionUpgrade.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoMinorVersionUpgrade gets a reference to the given NullableBool and assigns it to the EnableAutoMinorVersionUpgrade field.
func (o *RdsParams) SetEnableAutoMinorVersionUpgrade(v bool) {
	o.EnableAutoMinorVersionUpgrade.Set(&v)
}
// SetEnableAutoMinorVersionUpgradeNil sets the value for EnableAutoMinorVersionUpgrade to be an explicit nil
func (o *RdsParams) SetEnableAutoMinorVersionUpgradeNil() {
	o.EnableAutoMinorVersionUpgrade.Set(nil)
}

// UnsetEnableAutoMinorVersionUpgrade ensures that no value is present for EnableAutoMinorVersionUpgrade, not even an explicit nil
func (o *RdsParams) UnsetEnableAutoMinorVersionUpgrade() {
	o.EnableAutoMinorVersionUpgrade.Unset()
}

// GetEnableCopyTagsToSnapshots returns the EnableCopyTagsToSnapshots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetEnableCopyTagsToSnapshots() bool {
	if o == nil || o.EnableCopyTagsToSnapshots.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableCopyTagsToSnapshots.Get()
}

// GetEnableCopyTagsToSnapshotsOk returns a tuple with the EnableCopyTagsToSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetEnableCopyTagsToSnapshotsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableCopyTagsToSnapshots.Get(), o.EnableCopyTagsToSnapshots.IsSet()
}

// HasEnableCopyTagsToSnapshots returns a boolean if a field has been set.
func (o *RdsParams) HasEnableCopyTagsToSnapshots() bool {
	if o != nil && o.EnableCopyTagsToSnapshots.IsSet() {
		return true
	}

	return false
}

// SetEnableCopyTagsToSnapshots gets a reference to the given NullableBool and assigns it to the EnableCopyTagsToSnapshots field.
func (o *RdsParams) SetEnableCopyTagsToSnapshots(v bool) {
	o.EnableCopyTagsToSnapshots.Set(&v)
}
// SetEnableCopyTagsToSnapshotsNil sets the value for EnableCopyTagsToSnapshots to be an explicit nil
func (o *RdsParams) SetEnableCopyTagsToSnapshotsNil() {
	o.EnableCopyTagsToSnapshots.Set(nil)
}

// UnsetEnableCopyTagsToSnapshots ensures that no value is present for EnableCopyTagsToSnapshots, not even an explicit nil
func (o *RdsParams) UnsetEnableCopyTagsToSnapshots() {
	o.EnableCopyTagsToSnapshots.Unset()
}

// GetEnableDbAuthentication returns the EnableDbAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetEnableDbAuthentication() bool {
	if o == nil || o.EnableDbAuthentication.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableDbAuthentication.Get()
}

// GetEnableDbAuthenticationOk returns a tuple with the EnableDbAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetEnableDbAuthenticationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableDbAuthentication.Get(), o.EnableDbAuthentication.IsSet()
}

// HasEnableDbAuthentication returns a boolean if a field has been set.
func (o *RdsParams) HasEnableDbAuthentication() bool {
	if o != nil && o.EnableDbAuthentication.IsSet() {
		return true
	}

	return false
}

// SetEnableDbAuthentication gets a reference to the given NullableBool and assigns it to the EnableDbAuthentication field.
func (o *RdsParams) SetEnableDbAuthentication(v bool) {
	o.EnableDbAuthentication.Set(&v)
}
// SetEnableDbAuthenticationNil sets the value for EnableDbAuthentication to be an explicit nil
func (o *RdsParams) SetEnableDbAuthenticationNil() {
	o.EnableDbAuthentication.Set(nil)
}

// UnsetEnableDbAuthentication ensures that no value is present for EnableDbAuthentication, not even an explicit nil
func (o *RdsParams) UnsetEnableDbAuthentication() {
	o.EnableDbAuthentication.Unset()
}

// GetEnablePublicAccessibility returns the EnablePublicAccessibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetEnablePublicAccessibility() bool {
	if o == nil || o.EnablePublicAccessibility.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnablePublicAccessibility.Get()
}

// GetEnablePublicAccessibilityOk returns a tuple with the EnablePublicAccessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetEnablePublicAccessibilityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnablePublicAccessibility.Get(), o.EnablePublicAccessibility.IsSet()
}

// HasEnablePublicAccessibility returns a boolean if a field has been set.
func (o *RdsParams) HasEnablePublicAccessibility() bool {
	if o != nil && o.EnablePublicAccessibility.IsSet() {
		return true
	}

	return false
}

// SetEnablePublicAccessibility gets a reference to the given NullableBool and assigns it to the EnablePublicAccessibility field.
func (o *RdsParams) SetEnablePublicAccessibility(v bool) {
	o.EnablePublicAccessibility.Set(&v)
}
// SetEnablePublicAccessibilityNil sets the value for EnablePublicAccessibility to be an explicit nil
func (o *RdsParams) SetEnablePublicAccessibilityNil() {
	o.EnablePublicAccessibility.Set(nil)
}

// UnsetEnablePublicAccessibility ensures that no value is present for EnablePublicAccessibility, not even an explicit nil
func (o *RdsParams) UnsetEnablePublicAccessibility() {
	o.EnablePublicAccessibility.Unset()
}

// GetIsMultiAzDeployment returns the IsMultiAzDeployment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RdsParams) GetIsMultiAzDeployment() bool {
	if o == nil || o.IsMultiAzDeployment.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsMultiAzDeployment.Get()
}

// GetIsMultiAzDeploymentOk returns a tuple with the IsMultiAzDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RdsParams) GetIsMultiAzDeploymentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsMultiAzDeployment.Get(), o.IsMultiAzDeployment.IsSet()
}

// HasIsMultiAzDeployment returns a boolean if a field has been set.
func (o *RdsParams) HasIsMultiAzDeployment() bool {
	if o != nil && o.IsMultiAzDeployment.IsSet() {
		return true
	}

	return false
}

// SetIsMultiAzDeployment gets a reference to the given NullableBool and assigns it to the IsMultiAzDeployment field.
func (o *RdsParams) SetIsMultiAzDeployment(v bool) {
	o.IsMultiAzDeployment.Set(&v)
}
// SetIsMultiAzDeploymentNil sets the value for IsMultiAzDeployment to be an explicit nil
func (o *RdsParams) SetIsMultiAzDeploymentNil() {
	o.IsMultiAzDeployment.Set(nil)
}

// UnsetIsMultiAzDeployment ensures that no value is present for IsMultiAzDeployment, not even an explicit nil
func (o *RdsParams) UnsetIsMultiAzDeployment() {
	o.IsMultiAzDeployment.Unset()
}

func (o RdsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityZoneId.IsSet() {
		toSerialize["availabilityZoneId"] = o.AvailabilityZoneId.Get()
	}
	if true {
		toSerialize["dbInstanceId"] = o.DbInstanceId.Get()
	}
	if o.DbOptionGroupId.IsSet() {
		toSerialize["dbOptionGroupId"] = o.DbOptionGroupId.Get()
	}
	if o.DbParameterGroupId.IsSet() {
		toSerialize["dbParameterGroupId"] = o.DbParameterGroupId.Get()
	}
	if o.DbPort.IsSet() {
		toSerialize["dbPort"] = o.DbPort.Get()
	}
	if o.EnableAutoMinorVersionUpgrade.IsSet() {
		toSerialize["enableAutoMinorVersionUpgrade"] = o.EnableAutoMinorVersionUpgrade.Get()
	}
	if o.EnableCopyTagsToSnapshots.IsSet() {
		toSerialize["enableCopyTagsToSnapshots"] = o.EnableCopyTagsToSnapshots.Get()
	}
	if o.EnableDbAuthentication.IsSet() {
		toSerialize["enableDbAuthentication"] = o.EnableDbAuthentication.Get()
	}
	if o.EnablePublicAccessibility.IsSet() {
		toSerialize["enablePublicAccessibility"] = o.EnablePublicAccessibility.Get()
	}
	if o.IsMultiAzDeployment.IsSet() {
		toSerialize["isMultiAzDeployment"] = o.IsMultiAzDeployment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRdsParams struct {
	value *RdsParams
	isSet bool
}

func (v NullableRdsParams) Get() *RdsParams {
	return v.value
}

func (v *NullableRdsParams) Set(val *RdsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsParams(val *RdsParams) *NullableRdsParams {
	return &NullableRdsParams{value: val, isSet: true}
}

func (v NullableRdsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


