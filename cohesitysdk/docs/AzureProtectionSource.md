# AzureProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **NullableString** | Specifies Application Id of the active directory of Azure account. | [optional] 
**ApplicationKey** | Pointer to **NullableString** | Specifies Application key of the active directory of Azure account. | [optional] 
**AzureType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kSubscription&#39; if the environment is kAzure. Specifies the type of an Azure source entity. &#39;kSubscription&#39; indicates a billing unit within Azure account. &#39;kResourceGroup&#39; indicates a container that holds related resources. &#39;kVirtualMachine&#39; indicates a Virtual Machine in Azure environment. &#39;kStorageAccount&#39; represents a collection of storage containers. &#39;kStorageKey&#39; indicates a key required to access the storage account. &#39;kStorageContainer&#39; represents a storage container within a storage account. &#39;kStorageBlob&#39; represents a storage blog within a storage container. &#39;kStorageResourceGroup&#39; indicates a container that holds related storage resources. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kVirtualNetwork&#39; represents a virtual network. &#39;kNetworkResourceGroup&#39; indicates a container that holds related network resources. &#39;kSubnet&#39; represents a subnet within the virtual network. &#39;kComputeOptions&#39; indicates the number of CPU cores and memory size available for a type of a Virtual Machine. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the OS type of the Protection Source of type &#39;kVirtualMachine&#39; such as &#39;kWindows&#39; or &#39;kLinux&#39;. overrideDescription: true &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**IpAddresses** | Pointer to **[]string** | Specifies a list of IP addresses for entities of type &#39;kVirtualMachine&#39;. | [optional] 
**IsManagedVm** | Pointer to **NullableBool** | Specifies whether VM is managed or not for entities of type &#39;kVirtualMachine&#39;. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the physical location of the resource group. | [optional] 
**MemoryMbytes** | Pointer to **NullableInt64** | Specifies the amount of memory in MegaBytes of the Azure resource of type &#39;kComputeOptions&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Object set by the Cloud Provider. If the provider did not set a name for the object, this field is not set. | [optional] 
**NumCores** | Pointer to **NullableInt32** | Specifies the number of CPU cores of the Azure resource of type &#39;kComputeOptions&#39;. | [optional] 
**PhysicalSourceId** | Pointer to **NullableInt64** | Specifies the Protection Source id of the registered Physical Host. If the cloud entity is protected using a Physical Agent, it must be registered as a physical host. | [optional] 
**ResourceId** | Pointer to **NullableString** | Specifies the unique Id of the resource given by the cloud provider. | [optional] 
**RestoreTaskId** | Pointer to **NullableInt64** | Specifies the id of the \&quot;convert and deploy\&quot; restore task that created the entity in the cloud.  It is required to support the DR-to-cloud usecase where we replicate an on-prem entity to a cluster running in cloud, bring it up using \&quot;convert and deploy\&quot; mechanism, protect it using a cloud job that uses physical adapter, and convert it back to the on-prem format before replication.  Before replicating, we need to update the backup task state of the backed up entity using the on-prem entity and on-prem entity&#39;s parent. The id is used to lookup the restore entity that contains details about the on-prem entity.  It is set at the time of refreshing the cloud entity hierarchy if all the following conditions are met: Name of the current entity matches with name of any cloud entity deployed using the \&quot;convert and deploy\&quot; restore task. Restore entity associated with the above matched cloud entity has &#39;failed_over&#39; flag set to true in its cloud extension. | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Specifies Subscription id inside a customer&#39;s Azure account. It represents sub-section within the Azure account where a customer allows us to create VMs, storage account etc. | [optional] 
**SubscriptionType** | Pointer to **NullableString** | Specifies the subscription type of Azure such as &#39;kAzureCommercial&#39; or &#39;kAzureGovCloud&#39;. Specifies the subscription type of an Azure source entity. &#39;kAzureCommercial&#39; indicates a standard Azure subscription. &#39;kAzureGovCloud&#39; indicates a govt Azure subscription. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies Tenant Id of the active directory of Azure account. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of an Azure Protection Source Object such as &#39;kStorageContainer&#39;, &#39;kVirtualMachine&#39;, &#39;kVirtualNetwork&#39;, etc. Specifies the type of an Azure source entity. &#39;kSubscription&#39; indicates a billing unit within Azure account. &#39;kResourceGroup&#39; indicates a container that holds related resources. &#39;kVirtualMachine&#39; indicates a Virtual Machine in Azure environment. &#39;kStorageAccount&#39; represents a collection of storage containers. &#39;kStorageKey&#39; indicates a key required to access the storage account. &#39;kStorageContainer&#39; represents a storage container within a storage account. &#39;kStorageBlob&#39; represents a storage blog within a storage container. &#39;kStorageResourceGroup&#39; indicates a container that holds related storage resources. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kVirtualNetwork&#39; represents a virtual network. &#39;kNetworkResourceGroup&#39; indicates a container that holds related network resources. &#39;kSubnet&#39; represents a subnet within the virtual network. &#39;kComputeOptions&#39; indicates the number of CPU cores and memory size available for a type of a Virtual Machine. | [optional] 

## Methods

### NewAzureProtectionSource

`func NewAzureProtectionSource() *AzureProtectionSource`

NewAzureProtectionSource instantiates a new AzureProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProtectionSourceWithDefaults

`func NewAzureProtectionSourceWithDefaults() *AzureProtectionSource`

NewAzureProtectionSourceWithDefaults instantiates a new AzureProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AzureProtectionSource) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureProtectionSource) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureProtectionSource) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AzureProtectionSource) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *AzureProtectionSource) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *AzureProtectionSource) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationKey

`func (o *AzureProtectionSource) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *AzureProtectionSource) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *AzureProtectionSource) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *AzureProtectionSource) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### SetApplicationKeyNil

`func (o *AzureProtectionSource) SetApplicationKeyNil(b bool)`

 SetApplicationKeyNil sets the value for ApplicationKey to be an explicit nil

### UnsetApplicationKey
`func (o *AzureProtectionSource) UnsetApplicationKey()`

UnsetApplicationKey ensures that no value is present for ApplicationKey, not even an explicit nil
### GetAzureType

`func (o *AzureProtectionSource) GetAzureType() string`

GetAzureType returns the AzureType field if non-nil, zero value otherwise.

### GetAzureTypeOk

`func (o *AzureProtectionSource) GetAzureTypeOk() (*string, bool)`

GetAzureTypeOk returns a tuple with the AzureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureType

`func (o *AzureProtectionSource) SetAzureType(v string)`

SetAzureType sets AzureType field to given value.

### HasAzureType

`func (o *AzureProtectionSource) HasAzureType() bool`

HasAzureType returns a boolean if a field has been set.

### SetAzureTypeNil

`func (o *AzureProtectionSource) SetAzureTypeNil(b bool)`

 SetAzureTypeNil sets the value for AzureType to be an explicit nil

### UnsetAzureType
`func (o *AzureProtectionSource) UnsetAzureType()`

UnsetAzureType ensures that no value is present for AzureType, not even an explicit nil
### GetHostType

`func (o *AzureProtectionSource) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *AzureProtectionSource) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *AzureProtectionSource) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *AzureProtectionSource) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *AzureProtectionSource) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *AzureProtectionSource) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetIpAddresses

`func (o *AzureProtectionSource) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *AzureProtectionSource) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *AzureProtectionSource) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *AzureProtectionSource) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *AzureProtectionSource) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *AzureProtectionSource) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetIsManagedVm

`func (o *AzureProtectionSource) GetIsManagedVm() bool`

GetIsManagedVm returns the IsManagedVm field if non-nil, zero value otherwise.

### GetIsManagedVmOk

`func (o *AzureProtectionSource) GetIsManagedVmOk() (*bool, bool)`

GetIsManagedVmOk returns a tuple with the IsManagedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedVm

`func (o *AzureProtectionSource) SetIsManagedVm(v bool)`

SetIsManagedVm sets IsManagedVm field to given value.

### HasIsManagedVm

`func (o *AzureProtectionSource) HasIsManagedVm() bool`

HasIsManagedVm returns a boolean if a field has been set.

### SetIsManagedVmNil

`func (o *AzureProtectionSource) SetIsManagedVmNil(b bool)`

 SetIsManagedVmNil sets the value for IsManagedVm to be an explicit nil

### UnsetIsManagedVm
`func (o *AzureProtectionSource) UnsetIsManagedVm()`

UnsetIsManagedVm ensures that no value is present for IsManagedVm, not even an explicit nil
### GetLocation

`func (o *AzureProtectionSource) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureProtectionSource) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureProtectionSource) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureProtectionSource) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AzureProtectionSource) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AzureProtectionSource) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMemoryMbytes

`func (o *AzureProtectionSource) GetMemoryMbytes() int64`

GetMemoryMbytes returns the MemoryMbytes field if non-nil, zero value otherwise.

### GetMemoryMbytesOk

`func (o *AzureProtectionSource) GetMemoryMbytesOk() (*int64, bool)`

GetMemoryMbytesOk returns a tuple with the MemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbytes

`func (o *AzureProtectionSource) SetMemoryMbytes(v int64)`

SetMemoryMbytes sets MemoryMbytes field to given value.

### HasMemoryMbytes

`func (o *AzureProtectionSource) HasMemoryMbytes() bool`

HasMemoryMbytes returns a boolean if a field has been set.

### SetMemoryMbytesNil

`func (o *AzureProtectionSource) SetMemoryMbytesNil(b bool)`

 SetMemoryMbytesNil sets the value for MemoryMbytes to be an explicit nil

### UnsetMemoryMbytes
`func (o *AzureProtectionSource) UnsetMemoryMbytes()`

UnsetMemoryMbytes ensures that no value is present for MemoryMbytes, not even an explicit nil
### GetName

`func (o *AzureProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumCores

`func (o *AzureProtectionSource) GetNumCores() int32`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *AzureProtectionSource) GetNumCoresOk() (*int32, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *AzureProtectionSource) SetNumCores(v int32)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *AzureProtectionSource) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### SetNumCoresNil

`func (o *AzureProtectionSource) SetNumCoresNil(b bool)`

 SetNumCoresNil sets the value for NumCores to be an explicit nil

### UnsetNumCores
`func (o *AzureProtectionSource) UnsetNumCores()`

UnsetNumCores ensures that no value is present for NumCores, not even an explicit nil
### GetPhysicalSourceId

`func (o *AzureProtectionSource) GetPhysicalSourceId() int64`

GetPhysicalSourceId returns the PhysicalSourceId field if non-nil, zero value otherwise.

### GetPhysicalSourceIdOk

`func (o *AzureProtectionSource) GetPhysicalSourceIdOk() (*int64, bool)`

GetPhysicalSourceIdOk returns a tuple with the PhysicalSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSourceId

`func (o *AzureProtectionSource) SetPhysicalSourceId(v int64)`

SetPhysicalSourceId sets PhysicalSourceId field to given value.

### HasPhysicalSourceId

`func (o *AzureProtectionSource) HasPhysicalSourceId() bool`

HasPhysicalSourceId returns a boolean if a field has been set.

### SetPhysicalSourceIdNil

`func (o *AzureProtectionSource) SetPhysicalSourceIdNil(b bool)`

 SetPhysicalSourceIdNil sets the value for PhysicalSourceId to be an explicit nil

### UnsetPhysicalSourceId
`func (o *AzureProtectionSource) UnsetPhysicalSourceId()`

UnsetPhysicalSourceId ensures that no value is present for PhysicalSourceId, not even an explicit nil
### GetResourceId

`func (o *AzureProtectionSource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AzureProtectionSource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AzureProtectionSource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AzureProtectionSource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AzureProtectionSource) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AzureProtectionSource) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRestoreTaskId

`func (o *AzureProtectionSource) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *AzureProtectionSource) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *AzureProtectionSource) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.

### HasRestoreTaskId

`func (o *AzureProtectionSource) HasRestoreTaskId() bool`

HasRestoreTaskId returns a boolean if a field has been set.

### SetRestoreTaskIdNil

`func (o *AzureProtectionSource) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *AzureProtectionSource) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetSubscriptionId

`func (o *AzureProtectionSource) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureProtectionSource) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureProtectionSource) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureProtectionSource) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureProtectionSource) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureProtectionSource) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionType

`func (o *AzureProtectionSource) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AzureProtectionSource) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AzureProtectionSource) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *AzureProtectionSource) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *AzureProtectionSource) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *AzureProtectionSource) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetTenantId

`func (o *AzureProtectionSource) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureProtectionSource) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureProtectionSource) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureProtectionSource) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureProtectionSource) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureProtectionSource) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *AzureProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AzureProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AzureProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AzureProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


