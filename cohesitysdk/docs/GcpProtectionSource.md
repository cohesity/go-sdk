# GcpProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientEmailAddress** | Pointer to **NullableString** | Specifies Client email address associated with the service account. | [optional] 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies Client private associated with the service account. | [optional] 
**GcpType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kIAMUser&#39; if the environment is kGCP. Specifies the type of a GCP source entity. &#39;kIAMUser&#39; indicates a unique user within a GCP account. &#39;kProject&#39; represents compute resources and storage. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kVirtualMachine&#39; indicates a Virtual Machine running in GCP environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within GCP. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kLabel&#39; represents a label present on the instances. &#39;kMetaData&#39; represents a custom metadata present on instances. &#39;kTag&#39; represents a network tag on instances. &#39;kVPCConnector&#39; represents a VPC connector used for serverless VPC access. | [optional] 
**HostProjectId** | Pointer to **NullableString** | Specifies the host project id. It is populated in entities of type kSubnet if the subnet is part of a shared VPC. This contains the ID of host project the subnet belongs to. Populated in entities of type kProject if the project is a service project in a Shared VPC setup. This contains the ID of the host project it is attached to. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the OS type of the Protection Source of type &#39;kVirtualMachine&#39; such as &#39;kWindows&#39; or &#39;kLinux&#39;. overrideDescription: true &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**IpAddressesVM** | Pointer to **NullableString** | Specifies the IP address of the entity of type &#39;kVirtualMachine&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Object set by the Cloud Provider. If the provider did not set a name for the object, this field is not set. | [optional] 
**OwnerId** | Pointer to **NullableString** | Specifies the owner id of the resource in GCP environment. With type, name and ownerId gives a globally unique identity to the GCP entity. | [optional] 
**PhysicalSourceId** | Pointer to **NullableInt64** | Specifies the Protection Source id of the registered Physical Host. If the cloud entity is protected using a Physical Agent, it must be registered as a physical host. | [optional] 
**ProjectId** | Pointer to **NullableString** | Specifies the project Id. For the kIAMUser entity this contains the id of the project to be used to deploy proxy VMs. For entities of type kVirtualMachine this contains the id of the project the virtual machine belongs to. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region Id. For the kIAMUser entity this contains the region to be used to deploy proxy VMs. For entities of type kVirtualMachine this contains the region the virtual machine belongs to. | [optional] 
**ResourceId** | Pointer to **NullableString** | Specifies the unique Id of the resource given by the cloud provider. | [optional] 
**RestoreTaskId** | Pointer to **NullableInt64** | Specifies the id of the \&quot;convert and deploy\&quot; restore task that created the entity in the cloud.  It is required to support the DR-to-cloud usecase where we replicate an on-prem entity to a cluster running in cloud, bring it up using \&quot;convert and deploy\&quot; mechanism, protect it using a cloud job that uses physical adapter, and convert it back to the on-prem format before replication.  Before replicating, we need to update the backup task state of the backed up entity using the on-prem entity and on-prem entity&#39;s parent. The id is used to lookup the restore entity that contains details about the on-prem entity.  It is set at the time of refreshing the cloud entity hierarchy if all the following conditions are met: Name of the current entity matches with name of any cloud entity deployed using the \&quot;convert and deploy\&quot; restore task. Restore entity associated with the above matched cloud entity has &#39;failed_over&#39; flag set to true in its cloud extension. | [optional] 
**TagAttributes** | Pointer to [**[]TagAttribute**](TagAttribute.md) | Specifies the list of GCP tag attributes. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of an GCP Protection Source Object such as &#39;kIAMUser&#39;, &#39;kProject&#39;, &#39;kRegion&#39;, etc. Specifies the type of a GCP source entity. &#39;kIAMUser&#39; indicates a unique user within a GCP account. &#39;kProject&#39; represents compute resources and storage. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kVirtualMachine&#39; indicates a Virtual Machine running in GCP environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within GCP. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kLabel&#39; represents a label present on the instances. &#39;kMetaData&#39; represents a custom metadata present on instances. &#39;kTag&#39; represents a network tag on instances. &#39;kVPCConnector&#39; represents a VPC connector used for serverless VPC access. | [optional] 
**VpcNetwork** | Pointer to **NullableString** | Specifies the VPC Network to deploy proxy VMs. | [optional] 
**VpcSubnetwork** | Pointer to **NullableString** | Specifies the subnetwork to deploy proxy VMs. | [optional] 

## Methods

### NewGcpProtectionSource

`func NewGcpProtectionSource() *GcpProtectionSource`

NewGcpProtectionSource instantiates a new GcpProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpProtectionSourceWithDefaults

`func NewGcpProtectionSourceWithDefaults() *GcpProtectionSource`

NewGcpProtectionSourceWithDefaults instantiates a new GcpProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientEmailAddress

`func (o *GcpProtectionSource) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *GcpProtectionSource) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *GcpProtectionSource) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.

### HasClientEmailAddress

`func (o *GcpProtectionSource) HasClientEmailAddress() bool`

HasClientEmailAddress returns a boolean if a field has been set.

### SetClientEmailAddressNil

`func (o *GcpProtectionSource) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *GcpProtectionSource) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *GcpProtectionSource) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *GcpProtectionSource) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *GcpProtectionSource) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *GcpProtectionSource) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *GcpProtectionSource) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *GcpProtectionSource) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetGcpType

`func (o *GcpProtectionSource) GetGcpType() string`

GetGcpType returns the GcpType field if non-nil, zero value otherwise.

### GetGcpTypeOk

`func (o *GcpProtectionSource) GetGcpTypeOk() (*string, bool)`

GetGcpTypeOk returns a tuple with the GcpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpType

`func (o *GcpProtectionSource) SetGcpType(v string)`

SetGcpType sets GcpType field to given value.

### HasGcpType

`func (o *GcpProtectionSource) HasGcpType() bool`

HasGcpType returns a boolean if a field has been set.

### SetGcpTypeNil

`func (o *GcpProtectionSource) SetGcpTypeNil(b bool)`

 SetGcpTypeNil sets the value for GcpType to be an explicit nil

### UnsetGcpType
`func (o *GcpProtectionSource) UnsetGcpType()`

UnsetGcpType ensures that no value is present for GcpType, not even an explicit nil
### GetHostProjectId

`func (o *GcpProtectionSource) GetHostProjectId() string`

GetHostProjectId returns the HostProjectId field if non-nil, zero value otherwise.

### GetHostProjectIdOk

`func (o *GcpProtectionSource) GetHostProjectIdOk() (*string, bool)`

GetHostProjectIdOk returns a tuple with the HostProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProjectId

`func (o *GcpProtectionSource) SetHostProjectId(v string)`

SetHostProjectId sets HostProjectId field to given value.

### HasHostProjectId

`func (o *GcpProtectionSource) HasHostProjectId() bool`

HasHostProjectId returns a boolean if a field has been set.

### SetHostProjectIdNil

`func (o *GcpProtectionSource) SetHostProjectIdNil(b bool)`

 SetHostProjectIdNil sets the value for HostProjectId to be an explicit nil

### UnsetHostProjectId
`func (o *GcpProtectionSource) UnsetHostProjectId()`

UnsetHostProjectId ensures that no value is present for HostProjectId, not even an explicit nil
### GetHostType

`func (o *GcpProtectionSource) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *GcpProtectionSource) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *GcpProtectionSource) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *GcpProtectionSource) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *GcpProtectionSource) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *GcpProtectionSource) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetIpAddressesVM

`func (o *GcpProtectionSource) GetIpAddressesVM() string`

GetIpAddressesVM returns the IpAddressesVM field if non-nil, zero value otherwise.

### GetIpAddressesVMOk

`func (o *GcpProtectionSource) GetIpAddressesVMOk() (*string, bool)`

GetIpAddressesVMOk returns a tuple with the IpAddressesVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressesVM

`func (o *GcpProtectionSource) SetIpAddressesVM(v string)`

SetIpAddressesVM sets IpAddressesVM field to given value.

### HasIpAddressesVM

`func (o *GcpProtectionSource) HasIpAddressesVM() bool`

HasIpAddressesVM returns a boolean if a field has been set.

### SetIpAddressesVMNil

`func (o *GcpProtectionSource) SetIpAddressesVMNil(b bool)`

 SetIpAddressesVMNil sets the value for IpAddressesVM to be an explicit nil

### UnsetIpAddressesVM
`func (o *GcpProtectionSource) UnsetIpAddressesVM()`

UnsetIpAddressesVM ensures that no value is present for IpAddressesVM, not even an explicit nil
### GetName

`func (o *GcpProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *GcpProtectionSource) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GcpProtectionSource) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GcpProtectionSource) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GcpProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *GcpProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *GcpProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetPhysicalSourceId

`func (o *GcpProtectionSource) GetPhysicalSourceId() int64`

GetPhysicalSourceId returns the PhysicalSourceId field if non-nil, zero value otherwise.

### GetPhysicalSourceIdOk

`func (o *GcpProtectionSource) GetPhysicalSourceIdOk() (*int64, bool)`

GetPhysicalSourceIdOk returns a tuple with the PhysicalSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSourceId

`func (o *GcpProtectionSource) SetPhysicalSourceId(v int64)`

SetPhysicalSourceId sets PhysicalSourceId field to given value.

### HasPhysicalSourceId

`func (o *GcpProtectionSource) HasPhysicalSourceId() bool`

HasPhysicalSourceId returns a boolean if a field has been set.

### SetPhysicalSourceIdNil

`func (o *GcpProtectionSource) SetPhysicalSourceIdNil(b bool)`

 SetPhysicalSourceIdNil sets the value for PhysicalSourceId to be an explicit nil

### UnsetPhysicalSourceId
`func (o *GcpProtectionSource) UnsetPhysicalSourceId()`

UnsetPhysicalSourceId ensures that no value is present for PhysicalSourceId, not even an explicit nil
### GetProjectId

`func (o *GcpProtectionSource) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpProtectionSource) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpProtectionSource) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GcpProtectionSource) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *GcpProtectionSource) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GcpProtectionSource) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetRegionId

`func (o *GcpProtectionSource) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GcpProtectionSource) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GcpProtectionSource) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *GcpProtectionSource) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *GcpProtectionSource) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *GcpProtectionSource) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetResourceId

`func (o *GcpProtectionSource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GcpProtectionSource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GcpProtectionSource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *GcpProtectionSource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *GcpProtectionSource) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *GcpProtectionSource) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRestoreTaskId

`func (o *GcpProtectionSource) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *GcpProtectionSource) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *GcpProtectionSource) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.

### HasRestoreTaskId

`func (o *GcpProtectionSource) HasRestoreTaskId() bool`

HasRestoreTaskId returns a boolean if a field has been set.

### SetRestoreTaskIdNil

`func (o *GcpProtectionSource) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *GcpProtectionSource) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetTagAttributes

`func (o *GcpProtectionSource) GetTagAttributes() []TagAttribute`

GetTagAttributes returns the TagAttributes field if non-nil, zero value otherwise.

### GetTagAttributesOk

`func (o *GcpProtectionSource) GetTagAttributesOk() (*[]TagAttribute, bool)`

GetTagAttributesOk returns a tuple with the TagAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAttributes

`func (o *GcpProtectionSource) SetTagAttributes(v []TagAttribute)`

SetTagAttributes sets TagAttributes field to given value.

### HasTagAttributes

`func (o *GcpProtectionSource) HasTagAttributes() bool`

HasTagAttributes returns a boolean if a field has been set.

### SetTagAttributesNil

`func (o *GcpProtectionSource) SetTagAttributesNil(b bool)`

 SetTagAttributesNil sets the value for TagAttributes to be an explicit nil

### UnsetTagAttributes
`func (o *GcpProtectionSource) UnsetTagAttributes()`

UnsetTagAttributes ensures that no value is present for TagAttributes, not even an explicit nil
### GetType

`func (o *GcpProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GcpProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GcpProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GcpProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GcpProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GcpProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVpcNetwork

`func (o *GcpProtectionSource) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *GcpProtectionSource) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *GcpProtectionSource) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.

### HasVpcNetwork

`func (o *GcpProtectionSource) HasVpcNetwork() bool`

HasVpcNetwork returns a boolean if a field has been set.

### SetVpcNetworkNil

`func (o *GcpProtectionSource) SetVpcNetworkNil(b bool)`

 SetVpcNetworkNil sets the value for VpcNetwork to be an explicit nil

### UnsetVpcNetwork
`func (o *GcpProtectionSource) UnsetVpcNetwork()`

UnsetVpcNetwork ensures that no value is present for VpcNetwork, not even an explicit nil
### GetVpcSubnetwork

`func (o *GcpProtectionSource) GetVpcSubnetwork() string`

GetVpcSubnetwork returns the VpcSubnetwork field if non-nil, zero value otherwise.

### GetVpcSubnetworkOk

`func (o *GcpProtectionSource) GetVpcSubnetworkOk() (*string, bool)`

GetVpcSubnetworkOk returns a tuple with the VpcSubnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcSubnetwork

`func (o *GcpProtectionSource) SetVpcSubnetwork(v string)`

SetVpcSubnetwork sets VpcSubnetwork field to given value.

### HasVpcSubnetwork

`func (o *GcpProtectionSource) HasVpcSubnetwork() bool`

HasVpcSubnetwork returns a boolean if a field has been set.

### SetVpcSubnetworkNil

`func (o *GcpProtectionSource) SetVpcSubnetworkNil(b bool)`

 SetVpcSubnetworkNil sets the value for VpcSubnetwork to be an explicit nil

### UnsetVpcSubnetwork
`func (o *GcpProtectionSource) UnsetVpcSubnetwork()`

UnsetVpcSubnetwork ensures that no value is present for VpcSubnetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


