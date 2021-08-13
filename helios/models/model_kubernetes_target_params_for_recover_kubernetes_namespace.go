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

// KubernetesTargetParamsForRecoverKubernetesNamespace Specifies the parameters for recovering a Kubernetes namespace to a Kubernetes source.
type KubernetesTargetParamsForRecoverKubernetesNamespace struct {
	// Specifies the objects to be recovered.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects,omitempty"`
	// Specifies the Protection Group Runs params to recover. All the VM's that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
	RecoverProtectionGroupRunsParams []RecoverProtectionGroupRunParams `json:"recoverProtectionGroupRunsParams,omitempty"`
	// Specifies params to rename the Namespaces that are recovered. If not specified, the original names of the Namespaces are preserved. If a name collision occurs then the Namespace being recovered will overwrite the Namespace already present on the source.
	RenameRecoveredNamespacesParams NullableRecoveredOrClonedVmsRenameConfig `json:"renameRecoveredNamespacesParams,omitempty"`
	// Specifies the recovery target configuration of the Namespace recovery.
	RecoveryTargetConfig NullableKubernetesNamespaceRecoveryTargetConfig `json:"recoveryTargetConfig"`
}

// NewKubernetesTargetParamsForRecoverKubernetesNamespace instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesTargetParamsForRecoverKubernetesNamespace(recoveryTargetConfig NullableKubernetesNamespaceRecoveryTargetConfig) *KubernetesTargetParamsForRecoverKubernetesNamespace {
	this := KubernetesTargetParamsForRecoverKubernetesNamespace{}
	this.RecoveryTargetConfig = recoveryTargetConfig
	return &this
}

// NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesTargetParamsForRecoverKubernetesNamespaceWithDefaults() *KubernetesTargetParamsForRecoverKubernetesNamespace {
	this := KubernetesTargetParamsForRecoverKubernetesNamespace{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil  {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []CommonRecoverObjectSnapshotParams and assigns it to the Objects field.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams {
	if o == nil  {
		var ret []RecoverProtectionGroupRunParams
		return ret
	}
	return o.RecoverProtectionGroupRunsParams
}

// GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool) {
	if o == nil || o.RecoverProtectionGroupRunsParams == nil {
		return nil, false
	}
	return &o.RecoverProtectionGroupRunsParams, true
}

// HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasRecoverProtectionGroupRunsParams() bool {
	if o != nil && o.RecoverProtectionGroupRunsParams != nil {
		return true
	}

	return false
}

// SetRecoverProtectionGroupRunsParams gets a reference to the given []RecoverProtectionGroupRunParams and assigns it to the RecoverProtectionGroupRunsParams field.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams) {
	o.RecoverProtectionGroupRunsParams = v
}

// GetRenameRecoveredNamespacesParams returns the RenameRecoveredNamespacesParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParams() RecoveredOrClonedVmsRenameConfig {
	if o == nil || o.RenameRecoveredNamespacesParams.Get() == nil {
		var ret RecoveredOrClonedVmsRenameConfig
		return ret
	}
	return *o.RenameRecoveredNamespacesParams.Get()
}

// GetRenameRecoveredNamespacesParamsOk returns a tuple with the RenameRecoveredNamespacesParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRenameRecoveredNamespacesParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameRecoveredNamespacesParams.Get(), o.RenameRecoveredNamespacesParams.IsSet()
}

// HasRenameRecoveredNamespacesParams returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) HasRenameRecoveredNamespacesParams() bool {
	if o != nil && o.RenameRecoveredNamespacesParams.IsSet() {
		return true
	}

	return false
}

// SetRenameRecoveredNamespacesParams gets a reference to the given NullableRecoveredOrClonedVmsRenameConfig and assigns it to the RenameRecoveredNamespacesParams field.
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRenameRecoveredNamespacesParams(v RecoveredOrClonedVmsRenameConfig) {
	o.RenameRecoveredNamespacesParams.Set(&v)
}
// SetRenameRecoveredNamespacesParamsNil sets the value for RenameRecoveredNamespacesParams to be an explicit nil
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRenameRecoveredNamespacesParamsNil() {
	o.RenameRecoveredNamespacesParams.Set(nil)
}

// UnsetRenameRecoveredNamespacesParams ensures that no value is present for RenameRecoveredNamespacesParams, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) UnsetRenameRecoveredNamespacesParams() {
	o.RenameRecoveredNamespacesParams.Unset()
}

// GetRecoveryTargetConfig returns the RecoveryTargetConfig field value
// If the value is explicit nil, the zero value for KubernetesNamespaceRecoveryTargetConfig will be returned
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfig() KubernetesNamespaceRecoveryTargetConfig {
	if o == nil || o.RecoveryTargetConfig.Get() == nil {
		var ret KubernetesNamespaceRecoveryTargetConfig
		return ret
	}

	return *o.RecoveryTargetConfig.Get()
}

// GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) GetRecoveryTargetConfigOk() (*KubernetesNamespaceRecoveryTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoveryTargetConfig.Get(), o.RecoveryTargetConfig.IsSet()
}

// SetRecoveryTargetConfig sets field value
func (o *KubernetesTargetParamsForRecoverKubernetesNamespace) SetRecoveryTargetConfig(v KubernetesNamespaceRecoveryTargetConfig) {
	o.RecoveryTargetConfig.Set(&v)
}

func (o KubernetesTargetParamsForRecoverKubernetesNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.RecoverProtectionGroupRunsParams != nil {
		toSerialize["recoverProtectionGroupRunsParams"] = o.RecoverProtectionGroupRunsParams
	}
	if o.RenameRecoveredNamespacesParams.IsSet() {
		toSerialize["renameRecoveredNamespacesParams"] = o.RenameRecoveredNamespacesParams.Get()
	}
	if true {
		toSerialize["recoveryTargetConfig"] = o.RecoveryTargetConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesTargetParamsForRecoverKubernetesNamespace struct {
	value *KubernetesTargetParamsForRecoverKubernetesNamespace
	isSet bool
}

func (v NullableKubernetesTargetParamsForRecoverKubernetesNamespace) Get() *KubernetesTargetParamsForRecoverKubernetesNamespace {
	return v.value
}

func (v *NullableKubernetesTargetParamsForRecoverKubernetesNamespace) Set(val *KubernetesTargetParamsForRecoverKubernetesNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTargetParamsForRecoverKubernetesNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTargetParamsForRecoverKubernetesNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTargetParamsForRecoverKubernetesNamespace(val *KubernetesTargetParamsForRecoverKubernetesNamespace) *NullableKubernetesTargetParamsForRecoverKubernetesNamespace {
	return &NullableKubernetesTargetParamsForRecoverKubernetesNamespace{value: val, isSet: true}
}

func (v NullableKubernetesTargetParamsForRecoverKubernetesNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTargetParamsForRecoverKubernetesNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


