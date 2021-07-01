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

// RestoreKubernetesNamespacesParams struct for RestoreKubernetesNamespacesParams
type RestoreKubernetesNamespacesParams struct {
	// Backup job that needs to be used for recovering the namespace.
	BackupJobName NullableString `json:"backupJobName,omitempty"`
	ClusterEntity *EntityProto `json:"clusterEntity,omitempty"`
	// Namespace in which restore job will be created in K8s cluster.
	ManagementNamespace NullableString `json:"managementNamespace,omitempty"`
	RenameRestoredObjectParam *RenameObjectParamProto `json:"renameRestoredObjectParam,omitempty"`
}

// NewRestoreKubernetesNamespacesParams instantiates a new RestoreKubernetesNamespacesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreKubernetesNamespacesParams() *RestoreKubernetesNamespacesParams {
	this := RestoreKubernetesNamespacesParams{}
	return &this
}

// NewRestoreKubernetesNamespacesParamsWithDefaults instantiates a new RestoreKubernetesNamespacesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreKubernetesNamespacesParamsWithDefaults() *RestoreKubernetesNamespacesParams {
	this := RestoreKubernetesNamespacesParams{}
	return &this
}

// GetBackupJobName returns the BackupJobName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreKubernetesNamespacesParams) GetBackupJobName() string {
	if o == nil || o.BackupJobName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackupJobName.Get()
}

// GetBackupJobNameOk returns a tuple with the BackupJobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreKubernetesNamespacesParams) GetBackupJobNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupJobName.Get(), o.BackupJobName.IsSet()
}

// HasBackupJobName returns a boolean if a field has been set.
func (o *RestoreKubernetesNamespacesParams) HasBackupJobName() bool {
	if o != nil && o.BackupJobName.IsSet() {
		return true
	}

	return false
}

// SetBackupJobName gets a reference to the given NullableString and assigns it to the BackupJobName field.
func (o *RestoreKubernetesNamespacesParams) SetBackupJobName(v string) {
	o.BackupJobName.Set(&v)
}
// SetBackupJobNameNil sets the value for BackupJobName to be an explicit nil
func (o *RestoreKubernetesNamespacesParams) SetBackupJobNameNil() {
	o.BackupJobName.Set(nil)
}

// UnsetBackupJobName ensures that no value is present for BackupJobName, not even an explicit nil
func (o *RestoreKubernetesNamespacesParams) UnsetBackupJobName() {
	o.BackupJobName.Unset()
}

// GetClusterEntity returns the ClusterEntity field value if set, zero value otherwise.
func (o *RestoreKubernetesNamespacesParams) GetClusterEntity() EntityProto {
	if o == nil || o.ClusterEntity == nil {
		var ret EntityProto
		return ret
	}
	return *o.ClusterEntity
}

// GetClusterEntityOk returns a tuple with the ClusterEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreKubernetesNamespacesParams) GetClusterEntityOk() (*EntityProto, bool) {
	if o == nil || o.ClusterEntity == nil {
		return nil, false
	}
	return o.ClusterEntity, true
}

// HasClusterEntity returns a boolean if a field has been set.
func (o *RestoreKubernetesNamespacesParams) HasClusterEntity() bool {
	if o != nil && o.ClusterEntity != nil {
		return true
	}

	return false
}

// SetClusterEntity gets a reference to the given EntityProto and assigns it to the ClusterEntity field.
func (o *RestoreKubernetesNamespacesParams) SetClusterEntity(v EntityProto) {
	o.ClusterEntity = &v
}

// GetManagementNamespace returns the ManagementNamespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreKubernetesNamespacesParams) GetManagementNamespace() string {
	if o == nil || o.ManagementNamespace.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManagementNamespace.Get()
}

// GetManagementNamespaceOk returns a tuple with the ManagementNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreKubernetesNamespacesParams) GetManagementNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManagementNamespace.Get(), o.ManagementNamespace.IsSet()
}

// HasManagementNamespace returns a boolean if a field has been set.
func (o *RestoreKubernetesNamespacesParams) HasManagementNamespace() bool {
	if o != nil && o.ManagementNamespace.IsSet() {
		return true
	}

	return false
}

// SetManagementNamespace gets a reference to the given NullableString and assigns it to the ManagementNamespace field.
func (o *RestoreKubernetesNamespacesParams) SetManagementNamespace(v string) {
	o.ManagementNamespace.Set(&v)
}
// SetManagementNamespaceNil sets the value for ManagementNamespace to be an explicit nil
func (o *RestoreKubernetesNamespacesParams) SetManagementNamespaceNil() {
	o.ManagementNamespace.Set(nil)
}

// UnsetManagementNamespace ensures that no value is present for ManagementNamespace, not even an explicit nil
func (o *RestoreKubernetesNamespacesParams) UnsetManagementNamespace() {
	o.ManagementNamespace.Unset()
}

// GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field value if set, zero value otherwise.
func (o *RestoreKubernetesNamespacesParams) GetRenameRestoredObjectParam() RenameObjectParamProto {
	if o == nil || o.RenameRestoredObjectParam == nil {
		var ret RenameObjectParamProto
		return ret
	}
	return *o.RenameRestoredObjectParam
}

// GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreKubernetesNamespacesParams) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool) {
	if o == nil || o.RenameRestoredObjectParam == nil {
		return nil, false
	}
	return o.RenameRestoredObjectParam, true
}

// HasRenameRestoredObjectParam returns a boolean if a field has been set.
func (o *RestoreKubernetesNamespacesParams) HasRenameRestoredObjectParam() bool {
	if o != nil && o.RenameRestoredObjectParam != nil {
		return true
	}

	return false
}

// SetRenameRestoredObjectParam gets a reference to the given RenameObjectParamProto and assigns it to the RenameRestoredObjectParam field.
func (o *RestoreKubernetesNamespacesParams) SetRenameRestoredObjectParam(v RenameObjectParamProto) {
	o.RenameRestoredObjectParam = &v
}

func (o RestoreKubernetesNamespacesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupJobName.IsSet() {
		toSerialize["backupJobName"] = o.BackupJobName.Get()
	}
	if o.ClusterEntity != nil {
		toSerialize["clusterEntity"] = o.ClusterEntity
	}
	if o.ManagementNamespace.IsSet() {
		toSerialize["managementNamespace"] = o.ManagementNamespace.Get()
	}
	if o.RenameRestoredObjectParam != nil {
		toSerialize["renameRestoredObjectParam"] = o.RenameRestoredObjectParam
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreKubernetesNamespacesParams struct {
	value *RestoreKubernetesNamespacesParams
	isSet bool
}

func (v NullableRestoreKubernetesNamespacesParams) Get() *RestoreKubernetesNamespacesParams {
	return v.value
}

func (v *NullableRestoreKubernetesNamespacesParams) Set(val *RestoreKubernetesNamespacesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreKubernetesNamespacesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreKubernetesNamespacesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreKubernetesNamespacesParams(val *RestoreKubernetesNamespacesParams) *NullableRestoreKubernetesNamespacesParams {
	return &NullableRestoreKubernetesNamespacesParams{value: val, isSet: true}
}

func (v NullableRestoreKubernetesNamespacesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreKubernetesNamespacesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


