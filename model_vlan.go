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

// Vlan Specifies the settings of a VLAN. Its used by both Request and Response structures.
type Vlan struct {
	// Specifies whether to add the VLAN IPs to the cluster partition that already has one or more IPs from this VLAN.
	AddToClusterPartition NullableBool `json:"addToClusterPartition,omitempty"`
	// Specifies if this VLAN can be used by all tenants without explicit assignment to them. This option can only be set true for VLANs that are not assigned to any tenant.
	AllTenantAccess NullableBool `json:"allTenantAccess,omitempty"`
	// Set to true when ips are in use by Athena Apps. Note: If it is true then vlan interface can't be deleted.
	AppIpVecInUse NullableBool `json:"appIpVecInUse,omitempty"`
	// Array of Athena Apps IPs.  Specifies a list of Athena IPs in the VLAN.
	Appsips []string `json:"appsips,omitempty"`
	// Specifies a description of the VLAN.
	Description NullableString `json:"description,omitempty"`
	// Specifies the Gateway of the VLAN. It can carry V4 or V6 in case of requests, and carrises V4 in case of response.
	Gateway NullableString `json:"gateway,omitempty"`
	// Specifies the Gateway of the VLAN.
	GatewayV6 NullableString `json:"gatewayV6,omitempty"`
	// Specifies the hostname of the VLAN.
	Hostname NullableString `json:"hostname,omitempty"`
	// Specifies the id of the VLAN.
	Id NullableInt32 `json:"id,omitempty"`
	// Specifies the interface group name of the VLAN. It is in the format of <base_interface_group_name>.<vlan_id>.
	IfaceGroupName NullableString `json:"ifaceGroupName,omitempty"`
	// Specifies the interface name of the VLAN.
	InterfaceName NullableString `json:"interfaceName,omitempty"`
	// Specifies IP family. Based on this, subnet/gateway field contains V4 or V6 values. Used in Request.
	IpFamily NullableInt32 `json:"ipFamily,omitempty"`
	// IpPoolMap.  Pool IPs to program VIP followers.
	IpPoolMap map[string][]string `json:"ipPoolMap,omitempty"`
	IpRange *IpRange `json:"ipRange,omitempty"`
	// Array of IPs.  Specifies a list of IPs in the VLAN.
	Ips []string `json:"ips,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	// Specifies the subnet of the VLAN. The netmask can be specified by setting netmaskBits or netmaskIp4. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address. It can carry V4 or V6 in case of requests, and carries V4 in case of response.
	Subnet NullableSubnet `json:"subnet,omitempty"`
	// Specifies the subnet of the VLAN. The netmask can be specified by setting netmaskBits or netmaskIp4. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.
	SubnetV6 NullableSubnet `json:"subnetV6,omitempty"`
	// Optional tenant id that this vlan belongs to.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the VLAN name of the vlanId.
	VlanName NullableString `json:"vlanName,omitempty"`
}

// NewVlan instantiates a new Vlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlan() *Vlan {
	this := Vlan{}
	return &this
}

// NewVlanWithDefaults instantiates a new Vlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanWithDefaults() *Vlan {
	this := Vlan{}
	return &this
}

// GetAddToClusterPartition returns the AddToClusterPartition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetAddToClusterPartition() bool {
	if o == nil || o.AddToClusterPartition.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AddToClusterPartition.Get()
}

// GetAddToClusterPartitionOk returns a tuple with the AddToClusterPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetAddToClusterPartitionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddToClusterPartition.Get(), o.AddToClusterPartition.IsSet()
}

// HasAddToClusterPartition returns a boolean if a field has been set.
func (o *Vlan) HasAddToClusterPartition() bool {
	if o != nil && o.AddToClusterPartition.IsSet() {
		return true
	}

	return false
}

// SetAddToClusterPartition gets a reference to the given NullableBool and assigns it to the AddToClusterPartition field.
func (o *Vlan) SetAddToClusterPartition(v bool) {
	o.AddToClusterPartition.Set(&v)
}
// SetAddToClusterPartitionNil sets the value for AddToClusterPartition to be an explicit nil
func (o *Vlan) SetAddToClusterPartitionNil() {
	o.AddToClusterPartition.Set(nil)
}

// UnsetAddToClusterPartition ensures that no value is present for AddToClusterPartition, not even an explicit nil
func (o *Vlan) UnsetAddToClusterPartition() {
	o.AddToClusterPartition.Unset()
}

// GetAllTenantAccess returns the AllTenantAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetAllTenantAccess() bool {
	if o == nil || o.AllTenantAccess.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllTenantAccess.Get()
}

// GetAllTenantAccessOk returns a tuple with the AllTenantAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetAllTenantAccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllTenantAccess.Get(), o.AllTenantAccess.IsSet()
}

// HasAllTenantAccess returns a boolean if a field has been set.
func (o *Vlan) HasAllTenantAccess() bool {
	if o != nil && o.AllTenantAccess.IsSet() {
		return true
	}

	return false
}

// SetAllTenantAccess gets a reference to the given NullableBool and assigns it to the AllTenantAccess field.
func (o *Vlan) SetAllTenantAccess(v bool) {
	o.AllTenantAccess.Set(&v)
}
// SetAllTenantAccessNil sets the value for AllTenantAccess to be an explicit nil
func (o *Vlan) SetAllTenantAccessNil() {
	o.AllTenantAccess.Set(nil)
}

// UnsetAllTenantAccess ensures that no value is present for AllTenantAccess, not even an explicit nil
func (o *Vlan) UnsetAllTenantAccess() {
	o.AllTenantAccess.Unset()
}

// GetAppIpVecInUse returns the AppIpVecInUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetAppIpVecInUse() bool {
	if o == nil || o.AppIpVecInUse.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppIpVecInUse.Get()
}

// GetAppIpVecInUseOk returns a tuple with the AppIpVecInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetAppIpVecInUseOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppIpVecInUse.Get(), o.AppIpVecInUse.IsSet()
}

// HasAppIpVecInUse returns a boolean if a field has been set.
func (o *Vlan) HasAppIpVecInUse() bool {
	if o != nil && o.AppIpVecInUse.IsSet() {
		return true
	}

	return false
}

// SetAppIpVecInUse gets a reference to the given NullableBool and assigns it to the AppIpVecInUse field.
func (o *Vlan) SetAppIpVecInUse(v bool) {
	o.AppIpVecInUse.Set(&v)
}
// SetAppIpVecInUseNil sets the value for AppIpVecInUse to be an explicit nil
func (o *Vlan) SetAppIpVecInUseNil() {
	o.AppIpVecInUse.Set(nil)
}

// UnsetAppIpVecInUse ensures that no value is present for AppIpVecInUse, not even an explicit nil
func (o *Vlan) UnsetAppIpVecInUse() {
	o.AppIpVecInUse.Unset()
}

// GetAppsips returns the Appsips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetAppsips() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Appsips
}

// GetAppsipsOk returns a tuple with the Appsips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetAppsipsOk() (*[]string, bool) {
	if o == nil || o.Appsips == nil {
		return nil, false
	}
	return &o.Appsips, true
}

// HasAppsips returns a boolean if a field has been set.
func (o *Vlan) HasAppsips() bool {
	if o != nil && o.Appsips != nil {
		return true
	}

	return false
}

// SetAppsips gets a reference to the given []string and assigns it to the Appsips field.
func (o *Vlan) SetAppsips(v []string) {
	o.Appsips = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Vlan) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Vlan) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Vlan) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Vlan) UnsetDescription() {
	o.Description.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetGateway() string {
	if o == nil || o.Gateway.Get() == nil {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *Vlan) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *Vlan) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *Vlan) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *Vlan) UnsetGateway() {
	o.Gateway.Unset()
}

// GetGatewayV6 returns the GatewayV6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetGatewayV6() string {
	if o == nil || o.GatewayV6.Get() == nil {
		var ret string
		return ret
	}
	return *o.GatewayV6.Get()
}

// GetGatewayV6Ok returns a tuple with the GatewayV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetGatewayV6Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GatewayV6.Get(), o.GatewayV6.IsSet()
}

// HasGatewayV6 returns a boolean if a field has been set.
func (o *Vlan) HasGatewayV6() bool {
	if o != nil && o.GatewayV6.IsSet() {
		return true
	}

	return false
}

// SetGatewayV6 gets a reference to the given NullableString and assigns it to the GatewayV6 field.
func (o *Vlan) SetGatewayV6(v string) {
	o.GatewayV6.Set(&v)
}
// SetGatewayV6Nil sets the value for GatewayV6 to be an explicit nil
func (o *Vlan) SetGatewayV6Nil() {
	o.GatewayV6.Set(nil)
}

// UnsetGatewayV6 ensures that no value is present for GatewayV6, not even an explicit nil
func (o *Vlan) UnsetGatewayV6() {
	o.GatewayV6.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetHostname() string {
	if o == nil || o.Hostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *Vlan) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *Vlan) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *Vlan) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *Vlan) UnsetHostname() {
	o.Hostname.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Vlan) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Vlan) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Vlan) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Vlan) UnsetId() {
	o.Id.Unset()
}

// GetIfaceGroupName returns the IfaceGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetIfaceGroupName() string {
	if o == nil || o.IfaceGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IfaceGroupName.Get()
}

// GetIfaceGroupNameOk returns a tuple with the IfaceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetIfaceGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IfaceGroupName.Get(), o.IfaceGroupName.IsSet()
}

// HasIfaceGroupName returns a boolean if a field has been set.
func (o *Vlan) HasIfaceGroupName() bool {
	if o != nil && o.IfaceGroupName.IsSet() {
		return true
	}

	return false
}

// SetIfaceGroupName gets a reference to the given NullableString and assigns it to the IfaceGroupName field.
func (o *Vlan) SetIfaceGroupName(v string) {
	o.IfaceGroupName.Set(&v)
}
// SetIfaceGroupNameNil sets the value for IfaceGroupName to be an explicit nil
func (o *Vlan) SetIfaceGroupNameNil() {
	o.IfaceGroupName.Set(nil)
}

// UnsetIfaceGroupName ensures that no value is present for IfaceGroupName, not even an explicit nil
func (o *Vlan) UnsetIfaceGroupName() {
	o.IfaceGroupName.Unset()
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetInterfaceName() string {
	if o == nil || o.InterfaceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName.Get()
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetInterfaceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InterfaceName.Get(), o.InterfaceName.IsSet()
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *Vlan) HasInterfaceName() bool {
	if o != nil && o.InterfaceName.IsSet() {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given NullableString and assigns it to the InterfaceName field.
func (o *Vlan) SetInterfaceName(v string) {
	o.InterfaceName.Set(&v)
}
// SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil
func (o *Vlan) SetInterfaceNameNil() {
	o.InterfaceName.Set(nil)
}

// UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
func (o *Vlan) UnsetInterfaceName() {
	o.InterfaceName.Unset()
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetIpFamily() int32 {
	if o == nil || o.IpFamily.Get() == nil {
		var ret int32
		return ret
	}
	return *o.IpFamily.Get()
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetIpFamilyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpFamily.Get(), o.IpFamily.IsSet()
}

// HasIpFamily returns a boolean if a field has been set.
func (o *Vlan) HasIpFamily() bool {
	if o != nil && o.IpFamily.IsSet() {
		return true
	}

	return false
}

// SetIpFamily gets a reference to the given NullableInt32 and assigns it to the IpFamily field.
func (o *Vlan) SetIpFamily(v int32) {
	o.IpFamily.Set(&v)
}
// SetIpFamilyNil sets the value for IpFamily to be an explicit nil
func (o *Vlan) SetIpFamilyNil() {
	o.IpFamily.Set(nil)
}

// UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil
func (o *Vlan) UnsetIpFamily() {
	o.IpFamily.Unset()
}

// GetIpPoolMap returns the IpPoolMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetIpPoolMap() map[string][]string {
	if o == nil  {
		var ret map[string][]string
		return ret
	}
	return o.IpPoolMap
}

// GetIpPoolMapOk returns a tuple with the IpPoolMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetIpPoolMapOk() (*map[string][]string, bool) {
	if o == nil || o.IpPoolMap == nil {
		return nil, false
	}
	return &o.IpPoolMap, true
}

// HasIpPoolMap returns a boolean if a field has been set.
func (o *Vlan) HasIpPoolMap() bool {
	if o != nil && o.IpPoolMap != nil {
		return true
	}

	return false
}

// SetIpPoolMap gets a reference to the given map[string][]string and assigns it to the IpPoolMap field.
func (o *Vlan) SetIpPoolMap(v map[string][]string) {
	o.IpPoolMap = v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *Vlan) GetIpRange() IpRange {
	if o == nil || o.IpRange == nil {
		var ret IpRange
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vlan) GetIpRangeOk() (*IpRange, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *Vlan) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given IpRange and assigns it to the IpRange field.
func (o *Vlan) SetIpRange(v IpRange) {
	o.IpRange = &v
}

// GetIps returns the Ips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetIps() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetIpsOk() (*[]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return &o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *Vlan) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *Vlan) SetIps(v []string) {
	o.Ips = v
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetMtu() int32 {
	if o == nil || o.Mtu.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetMtuOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *Vlan) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *Vlan) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *Vlan) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *Vlan) UnsetMtu() {
	o.Mtu.Unset()
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetSubnet() Subnet {
	if o == nil || o.Subnet.Get() == nil {
		var ret Subnet
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetSubnetOk() (*Subnet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *Vlan) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableSubnet and assigns it to the Subnet field.
func (o *Vlan) SetSubnet(v Subnet) {
	o.Subnet.Set(&v)
}
// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *Vlan) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *Vlan) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetSubnetV6 returns the SubnetV6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetSubnetV6() Subnet {
	if o == nil || o.SubnetV6.Get() == nil {
		var ret Subnet
		return ret
	}
	return *o.SubnetV6.Get()
}

// GetSubnetV6Ok returns a tuple with the SubnetV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetSubnetV6Ok() (*Subnet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubnetV6.Get(), o.SubnetV6.IsSet()
}

// HasSubnetV6 returns a boolean if a field has been set.
func (o *Vlan) HasSubnetV6() bool {
	if o != nil && o.SubnetV6.IsSet() {
		return true
	}

	return false
}

// SetSubnetV6 gets a reference to the given NullableSubnet and assigns it to the SubnetV6 field.
func (o *Vlan) SetSubnetV6(v Subnet) {
	o.SubnetV6.Set(&v)
}
// SetSubnetV6Nil sets the value for SubnetV6 to be an explicit nil
func (o *Vlan) SetSubnetV6Nil() {
	o.SubnetV6.Set(nil)
}

// UnsetSubnetV6 ensures that no value is present for SubnetV6, not even an explicit nil
func (o *Vlan) UnsetSubnetV6() {
	o.SubnetV6.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Vlan) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Vlan) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Vlan) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Vlan) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetVlanName returns the VlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Vlan) GetVlanName() string {
	if o == nil || o.VlanName.Get() == nil {
		var ret string
		return ret
	}
	return *o.VlanName.Get()
}

// GetVlanNameOk returns a tuple with the VlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vlan) GetVlanNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanName.Get(), o.VlanName.IsSet()
}

// HasVlanName returns a boolean if a field has been set.
func (o *Vlan) HasVlanName() bool {
	if o != nil && o.VlanName.IsSet() {
		return true
	}

	return false
}

// SetVlanName gets a reference to the given NullableString and assigns it to the VlanName field.
func (o *Vlan) SetVlanName(v string) {
	o.VlanName.Set(&v)
}
// SetVlanNameNil sets the value for VlanName to be an explicit nil
func (o *Vlan) SetVlanNameNil() {
	o.VlanName.Set(nil)
}

// UnsetVlanName ensures that no value is present for VlanName, not even an explicit nil
func (o *Vlan) UnsetVlanName() {
	o.VlanName.Unset()
}

func (o Vlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddToClusterPartition.IsSet() {
		toSerialize["addToClusterPartition"] = o.AddToClusterPartition.Get()
	}
	if o.AllTenantAccess.IsSet() {
		toSerialize["allTenantAccess"] = o.AllTenantAccess.Get()
	}
	if o.AppIpVecInUse.IsSet() {
		toSerialize["appIpVecInUse"] = o.AppIpVecInUse.Get()
	}
	if o.Appsips != nil {
		toSerialize["appsips"] = o.Appsips
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.GatewayV6.IsSet() {
		toSerialize["gatewayV6"] = o.GatewayV6.Get()
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IfaceGroupName.IsSet() {
		toSerialize["ifaceGroupName"] = o.IfaceGroupName.Get()
	}
	if o.InterfaceName.IsSet() {
		toSerialize["interfaceName"] = o.InterfaceName.Get()
	}
	if o.IpFamily.IsSet() {
		toSerialize["ipFamily"] = o.IpFamily.Get()
	}
	if o.IpPoolMap != nil {
		toSerialize["ipPoolMap"] = o.IpPoolMap
	}
	if o.IpRange != nil {
		toSerialize["ipRange"] = o.IpRange
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.Subnet.IsSet() {
		toSerialize["subnet"] = o.Subnet.Get()
	}
	if o.SubnetV6.IsSet() {
		toSerialize["subnetV6"] = o.SubnetV6.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.VlanName.IsSet() {
		toSerialize["vlanName"] = o.VlanName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVlan struct {
	value *Vlan
	isSet bool
}

func (v NullableVlan) Get() *Vlan {
	return v.value
}

func (v *NullableVlan) Set(val *Vlan) {
	v.value = val
	v.isSet = true
}

func (v NullableVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlan(val *Vlan) *NullableVlan {
	return &NullableVlan{value: val, isSet: true}
}

func (v NullableVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


