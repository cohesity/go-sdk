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

// GcpCredentials Specifies the credentials to authenticate with Google Cloud Platform.
type GcpCredentials struct {
	// Specifies Client email address associated with the service account.
	ClientEmailAddress NullableString `json:"clientEmailAddress,omitempty"`
	// Specifies Client private associated with the service account.
	ClientPrivateKey NullableString `json:"clientPrivateKey,omitempty"`
	// Specifies the entity type such as 'kIAMUser' if the environment is kGCP. Specifies the type of a GCP source entity. 'kIAMUser' indicates a unique user within a GCP account. 'kProject' represents compute resources and storage. 'kRegion' indicates a geographical region in the global infrastructure. 'kAvailabilityZone' indicates an availability zone within a region. 'kVirtualMachine' indicates a Virtual Machine running in GCP environment. 'kVPC' indicates a virtual private cloud (VPC) network within GCP. 'kSubnet' indicates a subnet inside the VPC. 'kNetworkSecurityGroup' represents a network security group. 'kInstanceType' represents various machine types. 'kLabel' represents a label present on the instances. 'kMetaData' represents a custom metadata present on instances. 'kTag' represents a network tag on instances. 'kVPCConnector' represents a VPC connector used for serverless VPC access.
	GcpType NullableString `json:"gcpType,omitempty"`
	// Specifies Id of the project associated with Google cloud account.
	ProjectId NullableString `json:"projectId,omitempty"`
	// Specifies the VPC Network to deploy proxy VMs.
	VpcNetwork NullableString `json:"vpcNetwork,omitempty"`
	// Specifies the subnetwork to deploy proxy VMs.
	VpcSubnetwork NullableString `json:"vpcSubnetwork,omitempty"`
}

// NewGcpCredentials instantiates a new GcpCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpCredentials() *GcpCredentials {
	this := GcpCredentials{}
	return &this
}

// NewGcpCredentialsWithDefaults instantiates a new GcpCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpCredentialsWithDefaults() *GcpCredentials {
	this := GcpCredentials{}
	return &this
}

// GetClientEmailAddress returns the ClientEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetClientEmailAddress() string {
	if o == nil || o.ClientEmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientEmailAddress.Get()
}

// GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetClientEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientEmailAddress.Get(), o.ClientEmailAddress.IsSet()
}

// HasClientEmailAddress returns a boolean if a field has been set.
func (o *GcpCredentials) HasClientEmailAddress() bool {
	if o != nil && o.ClientEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetClientEmailAddress gets a reference to the given NullableString and assigns it to the ClientEmailAddress field.
func (o *GcpCredentials) SetClientEmailAddress(v string) {
	o.ClientEmailAddress.Set(&v)
}
// SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil
func (o *GcpCredentials) SetClientEmailAddressNil() {
	o.ClientEmailAddress.Set(nil)
}

// UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
func (o *GcpCredentials) UnsetClientEmailAddress() {
	o.ClientEmailAddress.Unset()
}

// GetClientPrivateKey returns the ClientPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetClientPrivateKey() string {
	if o == nil || o.ClientPrivateKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientPrivateKey.Get()
}

// GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetClientPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientPrivateKey.Get(), o.ClientPrivateKey.IsSet()
}

// HasClientPrivateKey returns a boolean if a field has been set.
func (o *GcpCredentials) HasClientPrivateKey() bool {
	if o != nil && o.ClientPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetClientPrivateKey gets a reference to the given NullableString and assigns it to the ClientPrivateKey field.
func (o *GcpCredentials) SetClientPrivateKey(v string) {
	o.ClientPrivateKey.Set(&v)
}
// SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil
func (o *GcpCredentials) SetClientPrivateKeyNil() {
	o.ClientPrivateKey.Set(nil)
}

// UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
func (o *GcpCredentials) UnsetClientPrivateKey() {
	o.ClientPrivateKey.Unset()
}

// GetGcpType returns the GcpType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetGcpType() string {
	if o == nil || o.GcpType.Get() == nil {
		var ret string
		return ret
	}
	return *o.GcpType.Get()
}

// GetGcpTypeOk returns a tuple with the GcpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetGcpTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GcpType.Get(), o.GcpType.IsSet()
}

// HasGcpType returns a boolean if a field has been set.
func (o *GcpCredentials) HasGcpType() bool {
	if o != nil && o.GcpType.IsSet() {
		return true
	}

	return false
}

// SetGcpType gets a reference to the given NullableString and assigns it to the GcpType field.
func (o *GcpCredentials) SetGcpType(v string) {
	o.GcpType.Set(&v)
}
// SetGcpTypeNil sets the value for GcpType to be an explicit nil
func (o *GcpCredentials) SetGcpTypeNil() {
	o.GcpType.Set(nil)
}

// UnsetGcpType ensures that no value is present for GcpType, not even an explicit nil
func (o *GcpCredentials) UnsetGcpType() {
	o.GcpType.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetProjectId() string {
	if o == nil || o.ProjectId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *GcpCredentials) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *GcpCredentials) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *GcpCredentials) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *GcpCredentials) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetVpcNetwork returns the VpcNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetVpcNetwork() string {
	if o == nil || o.VpcNetwork.Get() == nil {
		var ret string
		return ret
	}
	return *o.VpcNetwork.Get()
}

// GetVpcNetworkOk returns a tuple with the VpcNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetVpcNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VpcNetwork.Get(), o.VpcNetwork.IsSet()
}

// HasVpcNetwork returns a boolean if a field has been set.
func (o *GcpCredentials) HasVpcNetwork() bool {
	if o != nil && o.VpcNetwork.IsSet() {
		return true
	}

	return false
}

// SetVpcNetwork gets a reference to the given NullableString and assigns it to the VpcNetwork field.
func (o *GcpCredentials) SetVpcNetwork(v string) {
	o.VpcNetwork.Set(&v)
}
// SetVpcNetworkNil sets the value for VpcNetwork to be an explicit nil
func (o *GcpCredentials) SetVpcNetworkNil() {
	o.VpcNetwork.Set(nil)
}

// UnsetVpcNetwork ensures that no value is present for VpcNetwork, not even an explicit nil
func (o *GcpCredentials) UnsetVpcNetwork() {
	o.VpcNetwork.Unset()
}

// GetVpcSubnetwork returns the VpcSubnetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GcpCredentials) GetVpcSubnetwork() string {
	if o == nil || o.VpcSubnetwork.Get() == nil {
		var ret string
		return ret
	}
	return *o.VpcSubnetwork.Get()
}

// GetVpcSubnetworkOk returns a tuple with the VpcSubnetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GcpCredentials) GetVpcSubnetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VpcSubnetwork.Get(), o.VpcSubnetwork.IsSet()
}

// HasVpcSubnetwork returns a boolean if a field has been set.
func (o *GcpCredentials) HasVpcSubnetwork() bool {
	if o != nil && o.VpcSubnetwork.IsSet() {
		return true
	}

	return false
}

// SetVpcSubnetwork gets a reference to the given NullableString and assigns it to the VpcSubnetwork field.
func (o *GcpCredentials) SetVpcSubnetwork(v string) {
	o.VpcSubnetwork.Set(&v)
}
// SetVpcSubnetworkNil sets the value for VpcSubnetwork to be an explicit nil
func (o *GcpCredentials) SetVpcSubnetworkNil() {
	o.VpcSubnetwork.Set(nil)
}

// UnsetVpcSubnetwork ensures that no value is present for VpcSubnetwork, not even an explicit nil
func (o *GcpCredentials) UnsetVpcSubnetwork() {
	o.VpcSubnetwork.Unset()
}

func (o GcpCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientEmailAddress.IsSet() {
		toSerialize["clientEmailAddress"] = o.ClientEmailAddress.Get()
	}
	if o.ClientPrivateKey.IsSet() {
		toSerialize["clientPrivateKey"] = o.ClientPrivateKey.Get()
	}
	if o.GcpType.IsSet() {
		toSerialize["gcpType"] = o.GcpType.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.VpcNetwork.IsSet() {
		toSerialize["vpcNetwork"] = o.VpcNetwork.Get()
	}
	if o.VpcSubnetwork.IsSet() {
		toSerialize["vpcSubnetwork"] = o.VpcSubnetwork.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGcpCredentials struct {
	value *GcpCredentials
	isSet bool
}

func (v NullableGcpCredentials) Get() *GcpCredentials {
	return v.value
}

func (v *NullableGcpCredentials) Set(val *GcpCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpCredentials(val *GcpCredentials) *NullableGcpCredentials {
	return &NullableGcpCredentials{value: val, isSet: true}
}

func (v NullableGcpCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


