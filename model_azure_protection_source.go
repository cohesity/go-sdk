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

// AzureProtectionSource Specifies a Protection Source in Azure environment.
type AzureProtectionSource struct {
	// Specifies Application Id of the active directory of Azure account.
	ApplicationId NullableString `json:"applicationId,omitempty"`
	// Specifies Application key of the active directory of Azure account.
	ApplicationKey NullableString `json:"applicationKey,omitempty"`
	// Specifies the entity type such as 'kSubscription' if the environment is kAzure. Specifies the type of an Azure source entity. 'kSubscription' indicates a billing unit within Azure account. 'kResourceGroup' indicates a container that holds related resources. 'kVirtualMachine' indicates a Virtual Machine in Azure environment. 'kStorageAccount' represents a collection of storage containers. 'kStorageKey' indicates a key required to access the storage account. 'kStorageContainer' represents a storage container within a storage account. 'kStorageBlob' represents a storage blog within a storage container. 'kStorageResourceGroup' indicates a container that holds related storage resources. 'kNetworkSecurityGroup' represents a network security group. 'kVirtualNetwork' represents a virtual network. 'kNetworkResourceGroup' indicates a container that holds related network resources. 'kSubnet' represents a subnet within the virtual network. 'kComputeOptions' indicates the number of CPU cores and memory size available for a type of a Virtual Machine.
	AzureType NullableString `json:"azureType,omitempty"`
	// Specifies the OS type of the Protection Source of type 'kVirtualMachine' such as 'kWindows' or 'kLinux'. overrideDescription: true 'kLinux' indicates the Linux operating system. 'kWindows' indicates the Microsoft Windows operating system. 'kAix' indicates the IBM AIX operating system. 'kSolaris' indicates the Oracle Solaris operating system. 'kSapHana' indicates the Sap Hana database system developed by SAP SE. 'kOther' indicates the other types of operating system.
	HostType NullableString `json:"hostType,omitempty"`
	// Specifies a list of IP addresses for entities of type 'kVirtualMachine'.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Specifies whether VM is managed or not for entities of type 'kVirtualMachine'.
	IsManagedVm NullableBool `json:"isManagedVm,omitempty"`
	// Specifies the physical location of the resource group.
	Location NullableString `json:"location,omitempty"`
	// Specifies the amount of memory in MegaBytes of the Azure resource of type 'kComputeOptions'.
	MemoryMbytes NullableInt64 `json:"memoryMbytes,omitempty"`
	// Specifies the name of the Object set by the Cloud Provider. If the provider did not set a name for the object, this field is not set.
	Name NullableString `json:"name,omitempty"`
	// Specifies the number of CPU cores of the Azure resource of type 'kComputeOptions'.
	NumCores NullableInt32 `json:"numCores,omitempty"`
	// Specifies the Protection Source id of the registered Physical Host. If the cloud entity is protected using a Physical Agent, it must be registered as a physical host.
	PhysicalSourceId NullableInt64 `json:"physicalSourceId,omitempty"`
	// Specifies the unique Id of the resource given by the cloud provider.
	ResourceId NullableString `json:"resourceId,omitempty"`
	// Specifies the id of the \"convert and deploy\" restore task that created the entity in the cloud.  It is required to support the DR-to-cloud usecase where we replicate an on-prem entity to a cluster running in cloud, bring it up using \"convert and deploy\" mechanism, protect it using a cloud job that uses physical adapter, and convert it back to the on-prem format before replication.  Before replicating, we need to update the backup task state of the backed up entity using the on-prem entity and on-prem entity's parent. The id is used to lookup the restore entity that contains details about the on-prem entity.  It is set at the time of refreshing the cloud entity hierarchy if all the following conditions are met: Name of the current entity matches with name of any cloud entity deployed using the \"convert and deploy\" restore task. Restore entity associated with the above matched cloud entity has 'failed_over' flag set to true in its cloud extension.
	RestoreTaskId NullableInt64 `json:"restoreTaskId,omitempty"`
	// Specifies Subscription id inside a customer's Azure account. It represents sub-section within the Azure account where a customer allows us to create VMs, storage account etc.
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	// Specifies the subscription type of Azure such as 'kAzureCommercial' or 'kAzureGovCloud'. Specifies the subscription type of an Azure source entity. 'kAzureCommercial' indicates a standard Azure subscription. 'kAzureGovCloud' indicates a govt Azure subscription.
	SubscriptionType NullableString `json:"subscriptionType,omitempty"`
	// Specifies Tenant Id of the active directory of Azure account.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the type of an Azure Protection Source Object such as 'kStorageContainer', 'kVirtualMachine', 'kVirtualNetwork', etc. Specifies the type of an Azure source entity. 'kSubscription' indicates a billing unit within Azure account. 'kResourceGroup' indicates a container that holds related resources. 'kVirtualMachine' indicates a Virtual Machine in Azure environment. 'kStorageAccount' represents a collection of storage containers. 'kStorageKey' indicates a key required to access the storage account. 'kStorageContainer' represents a storage container within a storage account. 'kStorageBlob' represents a storage blog within a storage container. 'kStorageResourceGroup' indicates a container that holds related storage resources. 'kNetworkSecurityGroup' represents a network security group. 'kVirtualNetwork' represents a virtual network. 'kNetworkResourceGroup' indicates a container that holds related network resources. 'kSubnet' represents a subnet within the virtual network. 'kComputeOptions' indicates the number of CPU cores and memory size available for a type of a Virtual Machine.
	Type NullableString `json:"type,omitempty"`
}

// NewAzureProtectionSource instantiates a new AzureProtectionSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProtectionSource() *AzureProtectionSource {
	this := AzureProtectionSource{}
	return &this
}

// NewAzureProtectionSourceWithDefaults instantiates a new AzureProtectionSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProtectionSourceWithDefaults() *AzureProtectionSource {
	this := AzureProtectionSource{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasApplicationId() bool {
	if o != nil && o.ApplicationId.IsSet() {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given NullableString and assigns it to the ApplicationId field.
func (o *AzureProtectionSource) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}
// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil
func (o *AzureProtectionSource) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
func (o *AzureProtectionSource) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetApplicationKey returns the ApplicationKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetApplicationKey() string {
	if o == nil || o.ApplicationKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationKey.Get()
}

// GetApplicationKeyOk returns a tuple with the ApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetApplicationKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationKey.Get(), o.ApplicationKey.IsSet()
}

// HasApplicationKey returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasApplicationKey() bool {
	if o != nil && o.ApplicationKey.IsSet() {
		return true
	}

	return false
}

// SetApplicationKey gets a reference to the given NullableString and assigns it to the ApplicationKey field.
func (o *AzureProtectionSource) SetApplicationKey(v string) {
	o.ApplicationKey.Set(&v)
}
// SetApplicationKeyNil sets the value for ApplicationKey to be an explicit nil
func (o *AzureProtectionSource) SetApplicationKeyNil() {
	o.ApplicationKey.Set(nil)
}

// UnsetApplicationKey ensures that no value is present for ApplicationKey, not even an explicit nil
func (o *AzureProtectionSource) UnsetApplicationKey() {
	o.ApplicationKey.Unset()
}

// GetAzureType returns the AzureType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetAzureType() string {
	if o == nil || o.AzureType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AzureType.Get()
}

// GetAzureTypeOk returns a tuple with the AzureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetAzureTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AzureType.Get(), o.AzureType.IsSet()
}

// HasAzureType returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasAzureType() bool {
	if o != nil && o.AzureType.IsSet() {
		return true
	}

	return false
}

// SetAzureType gets a reference to the given NullableString and assigns it to the AzureType field.
func (o *AzureProtectionSource) SetAzureType(v string) {
	o.AzureType.Set(&v)
}
// SetAzureTypeNil sets the value for AzureType to be an explicit nil
func (o *AzureProtectionSource) SetAzureTypeNil() {
	o.AzureType.Set(nil)
}

// UnsetAzureType ensures that no value is present for AzureType, not even an explicit nil
func (o *AzureProtectionSource) UnsetAzureType() {
	o.AzureType.Unset()
}

// GetHostType returns the HostType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetHostType() string {
	if o == nil || o.HostType.Get() == nil {
		var ret string
		return ret
	}
	return *o.HostType.Get()
}

// GetHostTypeOk returns a tuple with the HostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetHostTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HostType.Get(), o.HostType.IsSet()
}

// HasHostType returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasHostType() bool {
	if o != nil && o.HostType.IsSet() {
		return true
	}

	return false
}

// SetHostType gets a reference to the given NullableString and assigns it to the HostType field.
func (o *AzureProtectionSource) SetHostType(v string) {
	o.HostType.Set(&v)
}
// SetHostTypeNil sets the value for HostType to be an explicit nil
func (o *AzureProtectionSource) SetHostTypeNil() {
	o.HostType.Set(nil)
}

// UnsetHostType ensures that no value is present for HostType, not even an explicit nil
func (o *AzureProtectionSource) UnsetHostType() {
	o.HostType.Unset()
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetIpAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *AzureProtectionSource) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetIsManagedVm returns the IsManagedVm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetIsManagedVm() bool {
	if o == nil || o.IsManagedVm.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsManagedVm.Get()
}

// GetIsManagedVmOk returns a tuple with the IsManagedVm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetIsManagedVmOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsManagedVm.Get(), o.IsManagedVm.IsSet()
}

// HasIsManagedVm returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasIsManagedVm() bool {
	if o != nil && o.IsManagedVm.IsSet() {
		return true
	}

	return false
}

// SetIsManagedVm gets a reference to the given NullableBool and assigns it to the IsManagedVm field.
func (o *AzureProtectionSource) SetIsManagedVm(v bool) {
	o.IsManagedVm.Set(&v)
}
// SetIsManagedVmNil sets the value for IsManagedVm to be an explicit nil
func (o *AzureProtectionSource) SetIsManagedVmNil() {
	o.IsManagedVm.Set(nil)
}

// UnsetIsManagedVm ensures that no value is present for IsManagedVm, not even an explicit nil
func (o *AzureProtectionSource) UnsetIsManagedVm() {
	o.IsManagedVm.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *AzureProtectionSource) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *AzureProtectionSource) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *AzureProtectionSource) UnsetLocation() {
	o.Location.Unset()
}

// GetMemoryMbytes returns the MemoryMbytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetMemoryMbytes() int64 {
	if o == nil || o.MemoryMbytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MemoryMbytes.Get()
}

// GetMemoryMbytesOk returns a tuple with the MemoryMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetMemoryMbytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryMbytes.Get(), o.MemoryMbytes.IsSet()
}

// HasMemoryMbytes returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasMemoryMbytes() bool {
	if o != nil && o.MemoryMbytes.IsSet() {
		return true
	}

	return false
}

// SetMemoryMbytes gets a reference to the given NullableInt64 and assigns it to the MemoryMbytes field.
func (o *AzureProtectionSource) SetMemoryMbytes(v int64) {
	o.MemoryMbytes.Set(&v)
}
// SetMemoryMbytesNil sets the value for MemoryMbytes to be an explicit nil
func (o *AzureProtectionSource) SetMemoryMbytesNil() {
	o.MemoryMbytes.Set(nil)
}

// UnsetMemoryMbytes ensures that no value is present for MemoryMbytes, not even an explicit nil
func (o *AzureProtectionSource) UnsetMemoryMbytes() {
	o.MemoryMbytes.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AzureProtectionSource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AzureProtectionSource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AzureProtectionSource) UnsetName() {
	o.Name.Unset()
}

// GetNumCores returns the NumCores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetNumCores() int32 {
	if o == nil || o.NumCores.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumCores.Get()
}

// GetNumCoresOk returns a tuple with the NumCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetNumCoresOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumCores.Get(), o.NumCores.IsSet()
}

// HasNumCores returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasNumCores() bool {
	if o != nil && o.NumCores.IsSet() {
		return true
	}

	return false
}

// SetNumCores gets a reference to the given NullableInt32 and assigns it to the NumCores field.
func (o *AzureProtectionSource) SetNumCores(v int32) {
	o.NumCores.Set(&v)
}
// SetNumCoresNil sets the value for NumCores to be an explicit nil
func (o *AzureProtectionSource) SetNumCoresNil() {
	o.NumCores.Set(nil)
}

// UnsetNumCores ensures that no value is present for NumCores, not even an explicit nil
func (o *AzureProtectionSource) UnsetNumCores() {
	o.NumCores.Unset()
}

// GetPhysicalSourceId returns the PhysicalSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetPhysicalSourceId() int64 {
	if o == nil || o.PhysicalSourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PhysicalSourceId.Get()
}

// GetPhysicalSourceIdOk returns a tuple with the PhysicalSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetPhysicalSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhysicalSourceId.Get(), o.PhysicalSourceId.IsSet()
}

// HasPhysicalSourceId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasPhysicalSourceId() bool {
	if o != nil && o.PhysicalSourceId.IsSet() {
		return true
	}

	return false
}

// SetPhysicalSourceId gets a reference to the given NullableInt64 and assigns it to the PhysicalSourceId field.
func (o *AzureProtectionSource) SetPhysicalSourceId(v int64) {
	o.PhysicalSourceId.Set(&v)
}
// SetPhysicalSourceIdNil sets the value for PhysicalSourceId to be an explicit nil
func (o *AzureProtectionSource) SetPhysicalSourceIdNil() {
	o.PhysicalSourceId.Set(nil)
}

// UnsetPhysicalSourceId ensures that no value is present for PhysicalSourceId, not even an explicit nil
func (o *AzureProtectionSource) UnsetPhysicalSourceId() {
	o.PhysicalSourceId.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetResourceId() string {
	if o == nil || o.ResourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableString and assigns it to the ResourceId field.
func (o *AzureProtectionSource) SetResourceId(v string) {
	o.ResourceId.Set(&v)
}
// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *AzureProtectionSource) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *AzureProtectionSource) UnsetResourceId() {
	o.ResourceId.Unset()
}

// GetRestoreTaskId returns the RestoreTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetRestoreTaskId() int64 {
	if o == nil || o.RestoreTaskId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RestoreTaskId.Get()
}

// GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetRestoreTaskIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreTaskId.Get(), o.RestoreTaskId.IsSet()
}

// HasRestoreTaskId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasRestoreTaskId() bool {
	if o != nil && o.RestoreTaskId.IsSet() {
		return true
	}

	return false
}

// SetRestoreTaskId gets a reference to the given NullableInt64 and assigns it to the RestoreTaskId field.
func (o *AzureProtectionSource) SetRestoreTaskId(v int64) {
	o.RestoreTaskId.Set(&v)
}
// SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil
func (o *AzureProtectionSource) SetRestoreTaskIdNil() {
	o.RestoreTaskId.Set(nil)
}

// UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
func (o *AzureProtectionSource) UnsetRestoreTaskId() {
	o.RestoreTaskId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *AzureProtectionSource) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *AzureProtectionSource) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *AzureProtectionSource) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType.Get()
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionType.Get(), o.SubscriptionType.IsSet()
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given NullableString and assigns it to the SubscriptionType field.
func (o *AzureProtectionSource) SetSubscriptionType(v string) {
	o.SubscriptionType.Set(&v)
}
// SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil
func (o *AzureProtectionSource) SetSubscriptionTypeNil() {
	o.SubscriptionType.Set(nil)
}

// UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
func (o *AzureProtectionSource) UnsetSubscriptionType() {
	o.SubscriptionType.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *AzureProtectionSource) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *AzureProtectionSource) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *AzureProtectionSource) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureProtectionSource) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureProtectionSource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AzureProtectionSource) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AzureProtectionSource) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AzureProtectionSource) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AzureProtectionSource) UnsetType() {
	o.Type.Unset()
}

func (o AzureProtectionSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId.IsSet() {
		toSerialize["applicationId"] = o.ApplicationId.Get()
	}
	if o.ApplicationKey.IsSet() {
		toSerialize["applicationKey"] = o.ApplicationKey.Get()
	}
	if o.AzureType.IsSet() {
		toSerialize["azureType"] = o.AzureType.Get()
	}
	if o.HostType.IsSet() {
		toSerialize["hostType"] = o.HostType.Get()
	}
	if o.IpAddresses != nil {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if o.IsManagedVm.IsSet() {
		toSerialize["isManagedVm"] = o.IsManagedVm.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.MemoryMbytes.IsSet() {
		toSerialize["memoryMbytes"] = o.MemoryMbytes.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NumCores.IsSet() {
		toSerialize["numCores"] = o.NumCores.Get()
	}
	if o.PhysicalSourceId.IsSet() {
		toSerialize["physicalSourceId"] = o.PhysicalSourceId.Get()
	}
	if o.ResourceId.IsSet() {
		toSerialize["resourceId"] = o.ResourceId.Get()
	}
	if o.RestoreTaskId.IsSet() {
		toSerialize["restoreTaskId"] = o.RestoreTaskId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.SubscriptionType.IsSet() {
		toSerialize["subscriptionType"] = o.SubscriptionType.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureProtectionSource struct {
	value *AzureProtectionSource
	isSet bool
}

func (v NullableAzureProtectionSource) Get() *AzureProtectionSource {
	return v.value
}

func (v *NullableAzureProtectionSource) Set(val *AzureProtectionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProtectionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProtectionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProtectionSource(val *AzureProtectionSource) *NullableAzureProtectionSource {
	return &NullableAzureProtectionSource{value: val, isSet: true}
}

func (v NullableAzureProtectionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProtectionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


