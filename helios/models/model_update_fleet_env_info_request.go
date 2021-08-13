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

// UpdateFleetEnvInfoRequest Specifies the params to add fleet env info.
type UpdateFleetEnvInfoRequest struct {
	// Specifies the IAM role used to create instances.
	IamRole NullableString `json:"iamRole"`
	// Specifies the Region of the CE cluster.
	Region NullableString `json:"region"`
	// Specifies the VPC of the CE cluster.
	VpcId NullableString `json:"vpcId"`
	// Specifies the Subnet of the CE cluster.
	SubnetId NullableString `json:"subnetId"`
	// Specifies the security group Id.
	SecurityGroupId NullableString `json:"securityGroupId"`
}

// NewUpdateFleetEnvInfoRequest instantiates a new UpdateFleetEnvInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFleetEnvInfoRequest(iamRole NullableString, region NullableString, vpcId NullableString, subnetId NullableString, securityGroupId NullableString) *UpdateFleetEnvInfoRequest {
	this := UpdateFleetEnvInfoRequest{}
	this.IamRole = iamRole
	this.Region = region
	this.VpcId = vpcId
	this.SubnetId = subnetId
	this.SecurityGroupId = securityGroupId
	return &this
}

// NewUpdateFleetEnvInfoRequestWithDefaults instantiates a new UpdateFleetEnvInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFleetEnvInfoRequestWithDefaults() *UpdateFleetEnvInfoRequest {
	this := UpdateFleetEnvInfoRequest{}
	return &this
}

// GetIamRole returns the IamRole field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateFleetEnvInfoRequest) GetIamRole() string {
	if o == nil || o.IamRole.Get() == nil {
		var ret string
		return ret
	}

	return *o.IamRole.Get()
}

// GetIamRoleOk returns a tuple with the IamRole field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFleetEnvInfoRequest) GetIamRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IamRole.Get(), o.IamRole.IsSet()
}

// SetIamRole sets field value
func (o *UpdateFleetEnvInfoRequest) SetIamRole(v string) {
	o.IamRole.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateFleetEnvInfoRequest) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFleetEnvInfoRequest) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *UpdateFleetEnvInfoRequest) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetVpcId returns the VpcId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateFleetEnvInfoRequest) GetVpcId() string {
	if o == nil || o.VpcId.Get() == nil {
		var ret string
		return ret
	}

	return *o.VpcId.Get()
}

// GetVpcIdOk returns a tuple with the VpcId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFleetEnvInfoRequest) GetVpcIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VpcId.Get(), o.VpcId.IsSet()
}

// SetVpcId sets field value
func (o *UpdateFleetEnvInfoRequest) SetVpcId(v string) {
	o.VpcId.Set(&v)
}

// GetSubnetId returns the SubnetId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateFleetEnvInfoRequest) GetSubnetId() string {
	if o == nil || o.SubnetId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubnetId.Get()
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFleetEnvInfoRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubnetId.Get(), o.SubnetId.IsSet()
}

// SetSubnetId sets field value
func (o *UpdateFleetEnvInfoRequest) SetSubnetId(v string) {
	o.SubnetId.Set(&v)
}

// GetSecurityGroupId returns the SecurityGroupId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateFleetEnvInfoRequest) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SecurityGroupId.Get()
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFleetEnvInfoRequest) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityGroupId.Get(), o.SecurityGroupId.IsSet()
}

// SetSecurityGroupId sets field value
func (o *UpdateFleetEnvInfoRequest) SetSecurityGroupId(v string) {
	o.SecurityGroupId.Set(&v)
}

func (o UpdateFleetEnvInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iamRole"] = o.IamRole.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["vpcId"] = o.VpcId.Get()
	}
	if true {
		toSerialize["subnetId"] = o.SubnetId.Get()
	}
	if true {
		toSerialize["securityGroupId"] = o.SecurityGroupId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFleetEnvInfoRequest struct {
	value *UpdateFleetEnvInfoRequest
	isSet bool
}

func (v NullableUpdateFleetEnvInfoRequest) Get() *UpdateFleetEnvInfoRequest {
	return v.value
}

func (v *NullableUpdateFleetEnvInfoRequest) Set(val *UpdateFleetEnvInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFleetEnvInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFleetEnvInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFleetEnvInfoRequest(val *UpdateFleetEnvInfoRequest) *NullableUpdateFleetEnvInfoRequest {
	return &NullableUpdateFleetEnvInfoRequest{value: val, isSet: true}
}

func (v NullableUpdateFleetEnvInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFleetEnvInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


