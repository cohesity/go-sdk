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

// HypervRestoreParameters Specifies information needed when restoring VMs in HyperV enviroment. This field defines the HyperV specific params for restore tasks of type kRecoverVMs.
type HypervRestoreParameters struct {
	// A datastore entity where the object's files should be restored to. This field is optional if object is being restored to its original parent source. If not specified, the object's files will be restored to their original datastore locations. This field is mandatory if object is being restored to a different resource entity or to a different parent source.
	DatastoreId NullableInt64 `json:"datastoreId,omitempty"`
	// Specifies whether the network should be left in disabled state. Attached network is enabled by default. Set this flag to true to disable it.
	DisableNetwork NullableBool `json:"disableNetwork,omitempty"`
	// Specifies a network configuration to be attached to the cloned or recovered object. For kCloneVMs and kRecoverVMs tasks, original network configuration is detached if the cloned or recovered object is kept under a different parent Protection Source or a different Resource Pool. By default, for kRecoverVMs task, original network configuration is preserved if the recovered object is kept under the same parent Protection Source and the same Resource Pool. Specify this field to override the preserved network configuration or to attach a new network configuration to the cloned or recovered objects. You can get the networkId of the kNetwork object by setting includeNetworks to 'true' in the GET /public/protectionSources operation. In the response, get the id of the desired kNetwork object, the resource pool, and the registered parent Protection Source.
	NetworkId NullableInt64 `json:"networkId,omitempty"`
	// Specifies the power state of the cloned or recovered objects. By default, the cloned or recovered objects are powered off.
	PoweredOn NullableBool `json:"poweredOn,omitempty"`
	// Specifies a prefix to prepended to the source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters.
	Prefix NullableString `json:"prefix,omitempty"`
	// Specifies whether or not to preserve tags during the operation.
	PreserveTags NullableBool `json:"preserveTags,omitempty"`
	// The resource (HyperV host) to which the restored VM will be attached.  This field is optional for a kRecoverVMs task if the VMs are being restored to its original parent source. If not specified, restored VMs will be attached to its original host. This field is mandatory if the VMs are being restored to a different parent source.
	ResourceId NullableInt64 `json:"resourceId,omitempty"`
	// Specifies a suffix to appended to the original source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewHypervRestoreParameters instantiates a new HypervRestoreParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervRestoreParameters() *HypervRestoreParameters {
	this := HypervRestoreParameters{}
	return &this
}

// NewHypervRestoreParametersWithDefaults instantiates a new HypervRestoreParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervRestoreParametersWithDefaults() *HypervRestoreParameters {
	this := HypervRestoreParameters{}
	return &this
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetDatastoreId() int64 {
	if o == nil || o.DatastoreId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreId.Get()
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetDatastoreIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatastoreId.Get(), o.DatastoreId.IsSet()
}

// HasDatastoreId returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasDatastoreId() bool {
	if o != nil && o.DatastoreId.IsSet() {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given NullableInt64 and assigns it to the DatastoreId field.
func (o *HypervRestoreParameters) SetDatastoreId(v int64) {
	o.DatastoreId.Set(&v)
}
// SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil
func (o *HypervRestoreParameters) SetDatastoreIdNil() {
	o.DatastoreId.Set(nil)
}

// UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
func (o *HypervRestoreParameters) UnsetDatastoreId() {
	o.DatastoreId.Unset()
}

// GetDisableNetwork returns the DisableNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetDisableNetwork() bool {
	if o == nil || o.DisableNetwork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DisableNetwork.Get()
}

// GetDisableNetworkOk returns a tuple with the DisableNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetDisableNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisableNetwork.Get(), o.DisableNetwork.IsSet()
}

// HasDisableNetwork returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasDisableNetwork() bool {
	if o != nil && o.DisableNetwork.IsSet() {
		return true
	}

	return false
}

// SetDisableNetwork gets a reference to the given NullableBool and assigns it to the DisableNetwork field.
func (o *HypervRestoreParameters) SetDisableNetwork(v bool) {
	o.DisableNetwork.Set(&v)
}
// SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil
func (o *HypervRestoreParameters) SetDisableNetworkNil() {
	o.DisableNetwork.Set(nil)
}

// UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
func (o *HypervRestoreParameters) UnsetDisableNetwork() {
	o.DisableNetwork.Unset()
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetNetworkId() int64 {
	if o == nil || o.NetworkId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NetworkId.Get()
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetNetworkIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkId.Get(), o.NetworkId.IsSet()
}

// HasNetworkId returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasNetworkId() bool {
	if o != nil && o.NetworkId.IsSet() {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given NullableInt64 and assigns it to the NetworkId field.
func (o *HypervRestoreParameters) SetNetworkId(v int64) {
	o.NetworkId.Set(&v)
}
// SetNetworkIdNil sets the value for NetworkId to be an explicit nil
func (o *HypervRestoreParameters) SetNetworkIdNil() {
	o.NetworkId.Set(nil)
}

// UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
func (o *HypervRestoreParameters) UnsetNetworkId() {
	o.NetworkId.Unset()
}

// GetPoweredOn returns the PoweredOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetPoweredOn() bool {
	if o == nil || o.PoweredOn.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PoweredOn.Get()
}

// GetPoweredOnOk returns a tuple with the PoweredOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetPoweredOnOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PoweredOn.Get(), o.PoweredOn.IsSet()
}

// HasPoweredOn returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasPoweredOn() bool {
	if o != nil && o.PoweredOn.IsSet() {
		return true
	}

	return false
}

// SetPoweredOn gets a reference to the given NullableBool and assigns it to the PoweredOn field.
func (o *HypervRestoreParameters) SetPoweredOn(v bool) {
	o.PoweredOn.Set(&v)
}
// SetPoweredOnNil sets the value for PoweredOn to be an explicit nil
func (o *HypervRestoreParameters) SetPoweredOnNil() {
	o.PoweredOn.Set(nil)
}

// UnsetPoweredOn ensures that no value is present for PoweredOn, not even an explicit nil
func (o *HypervRestoreParameters) UnsetPoweredOn() {
	o.PoweredOn.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *HypervRestoreParameters) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *HypervRestoreParameters) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *HypervRestoreParameters) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetPreserveTags returns the PreserveTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetPreserveTags() bool {
	if o == nil || o.PreserveTags.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveTags.Get()
}

// GetPreserveTagsOk returns a tuple with the PreserveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetPreserveTagsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveTags.Get(), o.PreserveTags.IsSet()
}

// HasPreserveTags returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasPreserveTags() bool {
	if o != nil && o.PreserveTags.IsSet() {
		return true
	}

	return false
}

// SetPreserveTags gets a reference to the given NullableBool and assigns it to the PreserveTags field.
func (o *HypervRestoreParameters) SetPreserveTags(v bool) {
	o.PreserveTags.Set(&v)
}
// SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil
func (o *HypervRestoreParameters) SetPreserveTagsNil() {
	o.PreserveTags.Set(nil)
}

// UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
func (o *HypervRestoreParameters) UnsetPreserveTags() {
	o.PreserveTags.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetResourceId() int64 {
	if o == nil || o.ResourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetResourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableInt64 and assigns it to the ResourceId field.
func (o *HypervRestoreParameters) SetResourceId(v int64) {
	o.ResourceId.Set(&v)
}
// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *HypervRestoreParameters) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *HypervRestoreParameters) UnsetResourceId() {
	o.ResourceId.Unset()
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervRestoreParameters) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervRestoreParameters) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *HypervRestoreParameters) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *HypervRestoreParameters) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *HypervRestoreParameters) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *HypervRestoreParameters) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o HypervRestoreParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatastoreId.IsSet() {
		toSerialize["datastoreId"] = o.DatastoreId.Get()
	}
	if o.DisableNetwork.IsSet() {
		toSerialize["disableNetwork"] = o.DisableNetwork.Get()
	}
	if o.NetworkId.IsSet() {
		toSerialize["networkId"] = o.NetworkId.Get()
	}
	if o.PoweredOn.IsSet() {
		toSerialize["poweredOn"] = o.PoweredOn.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.PreserveTags.IsSet() {
		toSerialize["preserveTags"] = o.PreserveTags.Get()
	}
	if o.ResourceId.IsSet() {
		toSerialize["resourceId"] = o.ResourceId.Get()
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHypervRestoreParameters struct {
	value *HypervRestoreParameters
	isSet bool
}

func (v NullableHypervRestoreParameters) Get() *HypervRestoreParameters {
	return v.value
}

func (v *NullableHypervRestoreParameters) Set(val *HypervRestoreParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervRestoreParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervRestoreParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervRestoreParameters(val *HypervRestoreParameters) *NullableHypervRestoreParameters {
	return &NullableHypervRestoreParameters{value: val, isSet: true}
}

func (v NullableHypervRestoreParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervRestoreParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


