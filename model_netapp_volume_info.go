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

// NetappVolumeInfo Specifies information about a volume in a NetApp cluster.
type NetappVolumeInfo struct {
	// Specifies the containing aggregate name of this volume.
	AggregateName NullableString `json:"aggregateName,omitempty"`
	// Specifies the total capacity in bytes of this volume.
	CapacityBytes NullableInt64 `json:"capacityBytes,omitempty"`
	// Array of CIFS Shares.  Specifies the set of CIFS Shares exported for this volume.
	CifsShares []CifsShareInfo `json:"cifsShares,omitempty"`
	// Specifies the creation time of the volume specified in Unix epoch time (in microseconds).
	CreationTimeUsecs NullableInt64 `json:"creationTimeUsecs,omitempty"`
	// Array of Data Protocols.  Specifies the set of data protocols supported by this volume. 'kNfs' indicates NFS connections. 'kCifs' indicates SMB (CIFS) connections. 'kIscsi' indicates iSCSI connections. 'kFc' indicates Fiber Channel connections. 'kFcache' indicates Flex Cache connections. 'kHttp' indicates HTTP connections. 'kNdmp' indicates NDMP connections. 'kManagement' indicates non-data connections used for management purposes. 'kNvme' indicates NVMe connections.
	DataProtocols []string `json:"dataProtocols,omitempty"`
	// Specifies the name of the export policy (which defines the access permissions for the mount client) that has been assigned to this volume.
	ExportPolicyName NullableString `json:"exportPolicyName,omitempty"`
	// Specifies the junction path of this volume. This path can be used to mount this volume via protocols such as NFS.
	JunctionPath NullableString `json:"junctionPath,omitempty"`
	// Specifies the name of the NetApp Vserver that this volume belongs to.
	Name NullableString `json:"name,omitempty"`
	SecurityInfo *VolumeSecurityInfo `json:"securityInfo,omitempty"`
	// Specifies the state of this volume. Specifies the state of a NetApp Volume. 'kOnline' indicates the volume is online. Read and write access to this volume is allowed. 'kRestricted' indicates the volume is restricted. Some operations, such as parity reconstruction, are allowed, but data access is not allowed. 'kOffline' indicates the volume is offline. No access to the volume is allowed. 'kMixed' indicates the volume is in mixed state, which means its aggregates are not all in the same state.
	State NullableString `json:"state,omitempty"`
	// Specifies the NetApp type of this volume. Specifies the type of a NetApp Volume. 'kReadWrite' indicates read-write volume. 'kLoadSharing' indicates load-sharing volume. 'kDataProtection' indicates data-protection volume. 'kDataCache' indicates data-cache volume. 'kTmp' indicates temporary purpose. 'kUnknownType' indicates unknown type.
	Type NullableString `json:"type,omitempty"`
	// Specifies the total space (in bytes) used in this volume.
	UsedBytes NullableInt64 `json:"usedBytes,omitempty"`
}

// NewNetappVolumeInfo instantiates a new NetappVolumeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetappVolumeInfo() *NetappVolumeInfo {
	this := NetappVolumeInfo{}
	return &this
}

// NewNetappVolumeInfoWithDefaults instantiates a new NetappVolumeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetappVolumeInfoWithDefaults() *NetappVolumeInfo {
	this := NetappVolumeInfo{}
	return &this
}

// GetAggregateName returns the AggregateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetAggregateName() string {
	if o == nil || o.AggregateName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AggregateName.Get()
}

// GetAggregateNameOk returns a tuple with the AggregateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetAggregateNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AggregateName.Get(), o.AggregateName.IsSet()
}

// HasAggregateName returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasAggregateName() bool {
	if o != nil && o.AggregateName.IsSet() {
		return true
	}

	return false
}

// SetAggregateName gets a reference to the given NullableString and assigns it to the AggregateName field.
func (o *NetappVolumeInfo) SetAggregateName(v string) {
	o.AggregateName.Set(&v)
}
// SetAggregateNameNil sets the value for AggregateName to be an explicit nil
func (o *NetappVolumeInfo) SetAggregateNameNil() {
	o.AggregateName.Set(nil)
}

// UnsetAggregateName ensures that no value is present for AggregateName, not even an explicit nil
func (o *NetappVolumeInfo) UnsetAggregateName() {
	o.AggregateName.Unset()
}

// GetCapacityBytes returns the CapacityBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetCapacityBytes() int64 {
	if o == nil || o.CapacityBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CapacityBytes.Get()
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetCapacityBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CapacityBytes.Get(), o.CapacityBytes.IsSet()
}

// HasCapacityBytes returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasCapacityBytes() bool {
	if o != nil && o.CapacityBytes.IsSet() {
		return true
	}

	return false
}

// SetCapacityBytes gets a reference to the given NullableInt64 and assigns it to the CapacityBytes field.
func (o *NetappVolumeInfo) SetCapacityBytes(v int64) {
	o.CapacityBytes.Set(&v)
}
// SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil
func (o *NetappVolumeInfo) SetCapacityBytesNil() {
	o.CapacityBytes.Set(nil)
}

// UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
func (o *NetappVolumeInfo) UnsetCapacityBytes() {
	o.CapacityBytes.Unset()
}

// GetCifsShares returns the CifsShares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetCifsShares() []CifsShareInfo {
	if o == nil  {
		var ret []CifsShareInfo
		return ret
	}
	return o.CifsShares
}

// GetCifsSharesOk returns a tuple with the CifsShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetCifsSharesOk() (*[]CifsShareInfo, bool) {
	if o == nil || o.CifsShares == nil {
		return nil, false
	}
	return &o.CifsShares, true
}

// HasCifsShares returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasCifsShares() bool {
	if o != nil && o.CifsShares != nil {
		return true
	}

	return false
}

// SetCifsShares gets a reference to the given []CifsShareInfo and assigns it to the CifsShares field.
func (o *NetappVolumeInfo) SetCifsShares(v []CifsShareInfo) {
	o.CifsShares = v
}

// GetCreationTimeUsecs returns the CreationTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetCreationTimeUsecs() int64 {
	if o == nil || o.CreationTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreationTimeUsecs.Get()
}

// GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetCreationTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreationTimeUsecs.Get(), o.CreationTimeUsecs.IsSet()
}

// HasCreationTimeUsecs returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasCreationTimeUsecs() bool {
	if o != nil && o.CreationTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetCreationTimeUsecs gets a reference to the given NullableInt64 and assigns it to the CreationTimeUsecs field.
func (o *NetappVolumeInfo) SetCreationTimeUsecs(v int64) {
	o.CreationTimeUsecs.Set(&v)
}
// SetCreationTimeUsecsNil sets the value for CreationTimeUsecs to be an explicit nil
func (o *NetappVolumeInfo) SetCreationTimeUsecsNil() {
	o.CreationTimeUsecs.Set(nil)
}

// UnsetCreationTimeUsecs ensures that no value is present for CreationTimeUsecs, not even an explicit nil
func (o *NetappVolumeInfo) UnsetCreationTimeUsecs() {
	o.CreationTimeUsecs.Unset()
}

// GetDataProtocols returns the DataProtocols field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetDataProtocols() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DataProtocols
}

// GetDataProtocolsOk returns a tuple with the DataProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetDataProtocolsOk() (*[]string, bool) {
	if o == nil || o.DataProtocols == nil {
		return nil, false
	}
	return &o.DataProtocols, true
}

// HasDataProtocols returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasDataProtocols() bool {
	if o != nil && o.DataProtocols != nil {
		return true
	}

	return false
}

// SetDataProtocols gets a reference to the given []string and assigns it to the DataProtocols field.
func (o *NetappVolumeInfo) SetDataProtocols(v []string) {
	o.DataProtocols = v
}

// GetExportPolicyName returns the ExportPolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetExportPolicyName() string {
	if o == nil || o.ExportPolicyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExportPolicyName.Get()
}

// GetExportPolicyNameOk returns a tuple with the ExportPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetExportPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExportPolicyName.Get(), o.ExportPolicyName.IsSet()
}

// HasExportPolicyName returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasExportPolicyName() bool {
	if o != nil && o.ExportPolicyName.IsSet() {
		return true
	}

	return false
}

// SetExportPolicyName gets a reference to the given NullableString and assigns it to the ExportPolicyName field.
func (o *NetappVolumeInfo) SetExportPolicyName(v string) {
	o.ExportPolicyName.Set(&v)
}
// SetExportPolicyNameNil sets the value for ExportPolicyName to be an explicit nil
func (o *NetappVolumeInfo) SetExportPolicyNameNil() {
	o.ExportPolicyName.Set(nil)
}

// UnsetExportPolicyName ensures that no value is present for ExportPolicyName, not even an explicit nil
func (o *NetappVolumeInfo) UnsetExportPolicyName() {
	o.ExportPolicyName.Unset()
}

// GetJunctionPath returns the JunctionPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetJunctionPath() string {
	if o == nil || o.JunctionPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.JunctionPath.Get()
}

// GetJunctionPathOk returns a tuple with the JunctionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetJunctionPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JunctionPath.Get(), o.JunctionPath.IsSet()
}

// HasJunctionPath returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasJunctionPath() bool {
	if o != nil && o.JunctionPath.IsSet() {
		return true
	}

	return false
}

// SetJunctionPath gets a reference to the given NullableString and assigns it to the JunctionPath field.
func (o *NetappVolumeInfo) SetJunctionPath(v string) {
	o.JunctionPath.Set(&v)
}
// SetJunctionPathNil sets the value for JunctionPath to be an explicit nil
func (o *NetappVolumeInfo) SetJunctionPathNil() {
	o.JunctionPath.Set(nil)
}

// UnsetJunctionPath ensures that no value is present for JunctionPath, not even an explicit nil
func (o *NetappVolumeInfo) UnsetJunctionPath() {
	o.JunctionPath.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NetappVolumeInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NetappVolumeInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NetappVolumeInfo) UnsetName() {
	o.Name.Unset()
}

// GetSecurityInfo returns the SecurityInfo field value if set, zero value otherwise.
func (o *NetappVolumeInfo) GetSecurityInfo() VolumeSecurityInfo {
	if o == nil || o.SecurityInfo == nil {
		var ret VolumeSecurityInfo
		return ret
	}
	return *o.SecurityInfo
}

// GetSecurityInfoOk returns a tuple with the SecurityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetappVolumeInfo) GetSecurityInfoOk() (*VolumeSecurityInfo, bool) {
	if o == nil || o.SecurityInfo == nil {
		return nil, false
	}
	return o.SecurityInfo, true
}

// HasSecurityInfo returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasSecurityInfo() bool {
	if o != nil && o.SecurityInfo != nil {
		return true
	}

	return false
}

// SetSecurityInfo gets a reference to the given VolumeSecurityInfo and assigns it to the SecurityInfo field.
func (o *NetappVolumeInfo) SetSecurityInfo(v VolumeSecurityInfo) {
	o.SecurityInfo = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *NetappVolumeInfo) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *NetappVolumeInfo) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *NetappVolumeInfo) UnsetState() {
	o.State.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *NetappVolumeInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *NetappVolumeInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *NetappVolumeInfo) UnsetType() {
	o.Type.Unset()
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappVolumeInfo) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes.Get()
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappVolumeInfo) GetUsedBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsedBytes.Get(), o.UsedBytes.IsSet()
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *NetappVolumeInfo) HasUsedBytes() bool {
	if o != nil && o.UsedBytes.IsSet() {
		return true
	}

	return false
}

// SetUsedBytes gets a reference to the given NullableInt64 and assigns it to the UsedBytes field.
func (o *NetappVolumeInfo) SetUsedBytes(v int64) {
	o.UsedBytes.Set(&v)
}
// SetUsedBytesNil sets the value for UsedBytes to be an explicit nil
func (o *NetappVolumeInfo) SetUsedBytesNil() {
	o.UsedBytes.Set(nil)
}

// UnsetUsedBytes ensures that no value is present for UsedBytes, not even an explicit nil
func (o *NetappVolumeInfo) UnsetUsedBytes() {
	o.UsedBytes.Unset()
}

func (o NetappVolumeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregateName.IsSet() {
		toSerialize["aggregateName"] = o.AggregateName.Get()
	}
	if o.CapacityBytes.IsSet() {
		toSerialize["capacityBytes"] = o.CapacityBytes.Get()
	}
	if o.CifsShares != nil {
		toSerialize["cifsShares"] = o.CifsShares
	}
	if o.CreationTimeUsecs.IsSet() {
		toSerialize["creationTimeUsecs"] = o.CreationTimeUsecs.Get()
	}
	if o.DataProtocols != nil {
		toSerialize["dataProtocols"] = o.DataProtocols
	}
	if o.ExportPolicyName.IsSet() {
		toSerialize["exportPolicyName"] = o.ExportPolicyName.Get()
	}
	if o.JunctionPath.IsSet() {
		toSerialize["junctionPath"] = o.JunctionPath.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SecurityInfo != nil {
		toSerialize["securityInfo"] = o.SecurityInfo
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UsedBytes.IsSet() {
		toSerialize["usedBytes"] = o.UsedBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNetappVolumeInfo struct {
	value *NetappVolumeInfo
	isSet bool
}

func (v NullableNetappVolumeInfo) Get() *NetappVolumeInfo {
	return v.value
}

func (v *NullableNetappVolumeInfo) Set(val *NetappVolumeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetappVolumeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetappVolumeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetappVolumeInfo(val *NetappVolumeInfo) *NullableNetappVolumeInfo {
	return &NullableNetappVolumeInfo{value: val, isSet: true}
}

func (v NullableNetappVolumeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetappVolumeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


