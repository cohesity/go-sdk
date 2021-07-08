# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityByTier** | Pointer to [**[]CapacityByTier**](CapacityByTier.md) | CapacityByTier describes the capacity of each storage tier. | [optional] 
**ChassisInfo** | Pointer to [**ChassisInfo**](ChassisInfo.md) |  | [optional] 
**ClusterPartitionId** | Pointer to **NullableInt64** | ClusterPartitionId is the Id of the cluster partition to which the Node belongs. | [optional] 
**ClusterPartitionName** | Pointer to **NullableString** | ClusterPartitionName is the name of the cluster to which the Node belongs. | [optional] 
**CohesityNodeSerial** | Pointer to **NullableString** | Cohesity Node Serial Number of the Node. | [optional] 
**DiskCount** | Pointer to **NullableInt64** | DiskCount is the number of disks in a node. | [optional] 
**DiskCountByTier** | Pointer to [**[]CountByTier**](CountByTier.md) | DiskCountByTier describes the disk number of each storage tier. | [optional] 
**Id** | Pointer to **NullableInt64** | Id is the Id of the Node. | [optional] 
**Ip** | Pointer to **NullableString** | Ip is the IP address of the Node. | [optional] 
**IsAppNode** | Pointer to **NullableBool** | Whether node is app node. | [optional] 
**IsMarkedForRemoval** | Pointer to **NullableBool** | IsMarkedForRemoval specifies whether the node has been marked for removal. | [optional] 
**MaxPhysicalCapacityBytes** | Pointer to **NullableInt64** | MaxPhysicalCapacityBytes specifies the maximum physical capacity of the node in bytes. | [optional] 
**NodeHardwareInfo** | Pointer to [**NodeHardwareInfo**](NodeHardwareInfo.md) |  | [optional] 
**NodeIncarnationId** | Pointer to **NullableInt64** | NodeIncarnationId is the incarnation id  of this node. The incarnation id is changed every time the data is wiped from the node. Various services on a node is only run if incarnation id of the node matches the incarnation id of the cluster. Whenever a mismatch is detected, Nexus will stop all services and clean the data from the node. After clean operation is completed, Nexus will set the node incarnation id to cluster incarnation id and start the services. | [optional] 
**NodeSoftwareVersion** | Pointer to **NullableString** | NodeSoftwareVersion is the current version of Cohesity software installed on a node. | [optional] 
**NodeType** | Pointer to **NullableString** | Node type: StorageNode, AllFlashNode, RoboNode, AppNode, etc. | [optional] 
**OfflineMountPathsOfDisks** | Pointer to **[]string** | OfflineMountPathsOfDisks provides the corresponding mount paths for direct attached disks that are currently offline - access to these were detected to hang sometime in the past. After these disks have been fixed, their mount paths needs to be removed from the following list before these will be accessed again. | [optional] 
**RemovalReason** | Pointer to **[]string** | RemovalReason specifies the removal reason of the node. &#39;kAutoHealthCheck&#39; means the entity health is bad. &#39;kUserGracefulRemoval&#39; means user initiated a graceful removal. &#39;kUserAvoidAccess&#39; means user initiated a mark offline. &#39;kUserGracefulNodeRemoval&#39; mean users initiated graceful node removal. &#39;kUserRemoveDownNode&#39; mean user initiated graceful removal of down node. | [optional] 
**RemovalState** | Pointer to **NullableString** | RemovalState specifies the removal state of the node. &#39;kDontRemove&#39; means the state of object is functional and it is not being removed. &#39;kMarkedForRemoval&#39; means the object is being removed. &#39;kOkToRemove&#39; means the object has been removed on the Cohesity Cluster and if the object is physical, it can be removed from the Cohesity Cluster. | [optional] 
**SlotNumber** | Pointer to **NullableInt32** | Slot number occupied by this node within the chassis. | [optional] 
**Stats** | Pointer to [**NodeStats**](NodeStats.md) |  | [optional] 
**SystemDisks** | Pointer to [**[]NodeSystemDiskInfo**](NodeSystemDiskInfo.md) | SystemDisk describes the node system disks. | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityByTier

`func (o *Node) GetCapacityByTier() []CapacityByTier`

GetCapacityByTier returns the CapacityByTier field if non-nil, zero value otherwise.

### GetCapacityByTierOk

`func (o *Node) GetCapacityByTierOk() (*[]CapacityByTier, bool)`

GetCapacityByTierOk returns a tuple with the CapacityByTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityByTier

`func (o *Node) SetCapacityByTier(v []CapacityByTier)`

SetCapacityByTier sets CapacityByTier field to given value.

### HasCapacityByTier

`func (o *Node) HasCapacityByTier() bool`

HasCapacityByTier returns a boolean if a field has been set.

### SetCapacityByTierNil

`func (o *Node) SetCapacityByTierNil(b bool)`

 SetCapacityByTierNil sets the value for CapacityByTier to be an explicit nil

### UnsetCapacityByTier
`func (o *Node) UnsetCapacityByTier()`

UnsetCapacityByTier ensures that no value is present for CapacityByTier, not even an explicit nil
### GetChassisInfo

`func (o *Node) GetChassisInfo() ChassisInfo`

GetChassisInfo returns the ChassisInfo field if non-nil, zero value otherwise.

### GetChassisInfoOk

`func (o *Node) GetChassisInfoOk() (*ChassisInfo, bool)`

GetChassisInfoOk returns a tuple with the ChassisInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisInfo

`func (o *Node) SetChassisInfo(v ChassisInfo)`

SetChassisInfo sets ChassisInfo field to given value.

### HasChassisInfo

`func (o *Node) HasChassisInfo() bool`

HasChassisInfo returns a boolean if a field has been set.

### GetClusterPartitionId

`func (o *Node) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *Node) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *Node) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.

### HasClusterPartitionId

`func (o *Node) HasClusterPartitionId() bool`

HasClusterPartitionId returns a boolean if a field has been set.

### SetClusterPartitionIdNil

`func (o *Node) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *Node) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetClusterPartitionName

`func (o *Node) GetClusterPartitionName() string`

GetClusterPartitionName returns the ClusterPartitionName field if non-nil, zero value otherwise.

### GetClusterPartitionNameOk

`func (o *Node) GetClusterPartitionNameOk() (*string, bool)`

GetClusterPartitionNameOk returns a tuple with the ClusterPartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionName

`func (o *Node) SetClusterPartitionName(v string)`

SetClusterPartitionName sets ClusterPartitionName field to given value.

### HasClusterPartitionName

`func (o *Node) HasClusterPartitionName() bool`

HasClusterPartitionName returns a boolean if a field has been set.

### SetClusterPartitionNameNil

`func (o *Node) SetClusterPartitionNameNil(b bool)`

 SetClusterPartitionNameNil sets the value for ClusterPartitionName to be an explicit nil

### UnsetClusterPartitionName
`func (o *Node) UnsetClusterPartitionName()`

UnsetClusterPartitionName ensures that no value is present for ClusterPartitionName, not even an explicit nil
### GetCohesityNodeSerial

`func (o *Node) GetCohesityNodeSerial() string`

GetCohesityNodeSerial returns the CohesityNodeSerial field if non-nil, zero value otherwise.

### GetCohesityNodeSerialOk

`func (o *Node) GetCohesityNodeSerialOk() (*string, bool)`

GetCohesityNodeSerialOk returns a tuple with the CohesityNodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesityNodeSerial

`func (o *Node) SetCohesityNodeSerial(v string)`

SetCohesityNodeSerial sets CohesityNodeSerial field to given value.

### HasCohesityNodeSerial

`func (o *Node) HasCohesityNodeSerial() bool`

HasCohesityNodeSerial returns a boolean if a field has been set.

### SetCohesityNodeSerialNil

`func (o *Node) SetCohesityNodeSerialNil(b bool)`

 SetCohesityNodeSerialNil sets the value for CohesityNodeSerial to be an explicit nil

### UnsetCohesityNodeSerial
`func (o *Node) UnsetCohesityNodeSerial()`

UnsetCohesityNodeSerial ensures that no value is present for CohesityNodeSerial, not even an explicit nil
### GetDiskCount

`func (o *Node) GetDiskCount() int64`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *Node) GetDiskCountOk() (*int64, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *Node) SetDiskCount(v int64)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *Node) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *Node) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *Node) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetDiskCountByTier

`func (o *Node) GetDiskCountByTier() []CountByTier`

GetDiskCountByTier returns the DiskCountByTier field if non-nil, zero value otherwise.

### GetDiskCountByTierOk

`func (o *Node) GetDiskCountByTierOk() (*[]CountByTier, bool)`

GetDiskCountByTierOk returns a tuple with the DiskCountByTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCountByTier

`func (o *Node) SetDiskCountByTier(v []CountByTier)`

SetDiskCountByTier sets DiskCountByTier field to given value.

### HasDiskCountByTier

`func (o *Node) HasDiskCountByTier() bool`

HasDiskCountByTier returns a boolean if a field has been set.

### SetDiskCountByTierNil

`func (o *Node) SetDiskCountByTierNil(b bool)`

 SetDiskCountByTierNil sets the value for DiskCountByTier to be an explicit nil

### UnsetDiskCountByTier
`func (o *Node) UnsetDiskCountByTier()`

UnsetDiskCountByTier ensures that no value is present for DiskCountByTier, not even an explicit nil
### GetId

`func (o *Node) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Node) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Node) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Node) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *Node) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Node) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Node) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Node) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *Node) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Node) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsAppNode

`func (o *Node) GetIsAppNode() bool`

GetIsAppNode returns the IsAppNode field if non-nil, zero value otherwise.

### GetIsAppNodeOk

`func (o *Node) GetIsAppNodeOk() (*bool, bool)`

GetIsAppNodeOk returns a tuple with the IsAppNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppNode

`func (o *Node) SetIsAppNode(v bool)`

SetIsAppNode sets IsAppNode field to given value.

### HasIsAppNode

`func (o *Node) HasIsAppNode() bool`

HasIsAppNode returns a boolean if a field has been set.

### SetIsAppNodeNil

`func (o *Node) SetIsAppNodeNil(b bool)`

 SetIsAppNodeNil sets the value for IsAppNode to be an explicit nil

### UnsetIsAppNode
`func (o *Node) UnsetIsAppNode()`

UnsetIsAppNode ensures that no value is present for IsAppNode, not even an explicit nil
### GetIsMarkedForRemoval

`func (o *Node) GetIsMarkedForRemoval() bool`

GetIsMarkedForRemoval returns the IsMarkedForRemoval field if non-nil, zero value otherwise.

### GetIsMarkedForRemovalOk

`func (o *Node) GetIsMarkedForRemovalOk() (*bool, bool)`

GetIsMarkedForRemovalOk returns a tuple with the IsMarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForRemoval

`func (o *Node) SetIsMarkedForRemoval(v bool)`

SetIsMarkedForRemoval sets IsMarkedForRemoval field to given value.

### HasIsMarkedForRemoval

`func (o *Node) HasIsMarkedForRemoval() bool`

HasIsMarkedForRemoval returns a boolean if a field has been set.

### SetIsMarkedForRemovalNil

`func (o *Node) SetIsMarkedForRemovalNil(b bool)`

 SetIsMarkedForRemovalNil sets the value for IsMarkedForRemoval to be an explicit nil

### UnsetIsMarkedForRemoval
`func (o *Node) UnsetIsMarkedForRemoval()`

UnsetIsMarkedForRemoval ensures that no value is present for IsMarkedForRemoval, not even an explicit nil
### GetMaxPhysicalCapacityBytes

`func (o *Node) GetMaxPhysicalCapacityBytes() int64`

GetMaxPhysicalCapacityBytes returns the MaxPhysicalCapacityBytes field if non-nil, zero value otherwise.

### GetMaxPhysicalCapacityBytesOk

`func (o *Node) GetMaxPhysicalCapacityBytesOk() (*int64, bool)`

GetMaxPhysicalCapacityBytesOk returns a tuple with the MaxPhysicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhysicalCapacityBytes

`func (o *Node) SetMaxPhysicalCapacityBytes(v int64)`

SetMaxPhysicalCapacityBytes sets MaxPhysicalCapacityBytes field to given value.

### HasMaxPhysicalCapacityBytes

`func (o *Node) HasMaxPhysicalCapacityBytes() bool`

HasMaxPhysicalCapacityBytes returns a boolean if a field has been set.

### SetMaxPhysicalCapacityBytesNil

`func (o *Node) SetMaxPhysicalCapacityBytesNil(b bool)`

 SetMaxPhysicalCapacityBytesNil sets the value for MaxPhysicalCapacityBytes to be an explicit nil

### UnsetMaxPhysicalCapacityBytes
`func (o *Node) UnsetMaxPhysicalCapacityBytes()`

UnsetMaxPhysicalCapacityBytes ensures that no value is present for MaxPhysicalCapacityBytes, not even an explicit nil
### GetNodeHardwareInfo

`func (o *Node) GetNodeHardwareInfo() NodeHardwareInfo`

GetNodeHardwareInfo returns the NodeHardwareInfo field if non-nil, zero value otherwise.

### GetNodeHardwareInfoOk

`func (o *Node) GetNodeHardwareInfoOk() (*NodeHardwareInfo, bool)`

GetNodeHardwareInfoOk returns a tuple with the NodeHardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHardwareInfo

`func (o *Node) SetNodeHardwareInfo(v NodeHardwareInfo)`

SetNodeHardwareInfo sets NodeHardwareInfo field to given value.

### HasNodeHardwareInfo

`func (o *Node) HasNodeHardwareInfo() bool`

HasNodeHardwareInfo returns a boolean if a field has been set.

### GetNodeIncarnationId

`func (o *Node) GetNodeIncarnationId() int64`

GetNodeIncarnationId returns the NodeIncarnationId field if non-nil, zero value otherwise.

### GetNodeIncarnationIdOk

`func (o *Node) GetNodeIncarnationIdOk() (*int64, bool)`

GetNodeIncarnationIdOk returns a tuple with the NodeIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIncarnationId

`func (o *Node) SetNodeIncarnationId(v int64)`

SetNodeIncarnationId sets NodeIncarnationId field to given value.

### HasNodeIncarnationId

`func (o *Node) HasNodeIncarnationId() bool`

HasNodeIncarnationId returns a boolean if a field has been set.

### SetNodeIncarnationIdNil

`func (o *Node) SetNodeIncarnationIdNil(b bool)`

 SetNodeIncarnationIdNil sets the value for NodeIncarnationId to be an explicit nil

### UnsetNodeIncarnationId
`func (o *Node) UnsetNodeIncarnationId()`

UnsetNodeIncarnationId ensures that no value is present for NodeIncarnationId, not even an explicit nil
### GetNodeSoftwareVersion

`func (o *Node) GetNodeSoftwareVersion() string`

GetNodeSoftwareVersion returns the NodeSoftwareVersion field if non-nil, zero value otherwise.

### GetNodeSoftwareVersionOk

`func (o *Node) GetNodeSoftwareVersionOk() (*string, bool)`

GetNodeSoftwareVersionOk returns a tuple with the NodeSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSoftwareVersion

`func (o *Node) SetNodeSoftwareVersion(v string)`

SetNodeSoftwareVersion sets NodeSoftwareVersion field to given value.

### HasNodeSoftwareVersion

`func (o *Node) HasNodeSoftwareVersion() bool`

HasNodeSoftwareVersion returns a boolean if a field has been set.

### SetNodeSoftwareVersionNil

`func (o *Node) SetNodeSoftwareVersionNil(b bool)`

 SetNodeSoftwareVersionNil sets the value for NodeSoftwareVersion to be an explicit nil

### UnsetNodeSoftwareVersion
`func (o *Node) UnsetNodeSoftwareVersion()`

UnsetNodeSoftwareVersion ensures that no value is present for NodeSoftwareVersion, not even an explicit nil
### GetNodeType

`func (o *Node) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *Node) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *Node) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *Node) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### SetNodeTypeNil

`func (o *Node) SetNodeTypeNil(b bool)`

 SetNodeTypeNil sets the value for NodeType to be an explicit nil

### UnsetNodeType
`func (o *Node) UnsetNodeType()`

UnsetNodeType ensures that no value is present for NodeType, not even an explicit nil
### GetOfflineMountPathsOfDisks

`func (o *Node) GetOfflineMountPathsOfDisks() []string`

GetOfflineMountPathsOfDisks returns the OfflineMountPathsOfDisks field if non-nil, zero value otherwise.

### GetOfflineMountPathsOfDisksOk

`func (o *Node) GetOfflineMountPathsOfDisksOk() (*[]string, bool)`

GetOfflineMountPathsOfDisksOk returns a tuple with the OfflineMountPathsOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineMountPathsOfDisks

`func (o *Node) SetOfflineMountPathsOfDisks(v []string)`

SetOfflineMountPathsOfDisks sets OfflineMountPathsOfDisks field to given value.

### HasOfflineMountPathsOfDisks

`func (o *Node) HasOfflineMountPathsOfDisks() bool`

HasOfflineMountPathsOfDisks returns a boolean if a field has been set.

### SetOfflineMountPathsOfDisksNil

`func (o *Node) SetOfflineMountPathsOfDisksNil(b bool)`

 SetOfflineMountPathsOfDisksNil sets the value for OfflineMountPathsOfDisks to be an explicit nil

### UnsetOfflineMountPathsOfDisks
`func (o *Node) UnsetOfflineMountPathsOfDisks()`

UnsetOfflineMountPathsOfDisks ensures that no value is present for OfflineMountPathsOfDisks, not even an explicit nil
### GetRemovalReason

`func (o *Node) GetRemovalReason() []string`

GetRemovalReason returns the RemovalReason field if non-nil, zero value otherwise.

### GetRemovalReasonOk

`func (o *Node) GetRemovalReasonOk() (*[]string, bool)`

GetRemovalReasonOk returns a tuple with the RemovalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalReason

`func (o *Node) SetRemovalReason(v []string)`

SetRemovalReason sets RemovalReason field to given value.

### HasRemovalReason

`func (o *Node) HasRemovalReason() bool`

HasRemovalReason returns a boolean if a field has been set.

### SetRemovalReasonNil

`func (o *Node) SetRemovalReasonNil(b bool)`

 SetRemovalReasonNil sets the value for RemovalReason to be an explicit nil

### UnsetRemovalReason
`func (o *Node) UnsetRemovalReason()`

UnsetRemovalReason ensures that no value is present for RemovalReason, not even an explicit nil
### GetRemovalState

`func (o *Node) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *Node) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *Node) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *Node) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *Node) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *Node) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetSlotNumber

`func (o *Node) GetSlotNumber() int32`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *Node) GetSlotNumberOk() (*int32, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *Node) SetSlotNumber(v int32)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *Node) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### SetSlotNumberNil

`func (o *Node) SetSlotNumberNil(b bool)`

 SetSlotNumberNil sets the value for SlotNumber to be an explicit nil

### UnsetSlotNumber
`func (o *Node) UnsetSlotNumber()`

UnsetSlotNumber ensures that no value is present for SlotNumber, not even an explicit nil
### GetStats

`func (o *Node) GetStats() NodeStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Node) GetStatsOk() (*NodeStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Node) SetStats(v NodeStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Node) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSystemDisks

`func (o *Node) GetSystemDisks() []NodeSystemDiskInfo`

GetSystemDisks returns the SystemDisks field if non-nil, zero value otherwise.

### GetSystemDisksOk

`func (o *Node) GetSystemDisksOk() (*[]NodeSystemDiskInfo, bool)`

GetSystemDisksOk returns a tuple with the SystemDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDisks

`func (o *Node) SetSystemDisks(v []NodeSystemDiskInfo)`

SetSystemDisks sets SystemDisks field to given value.

### HasSystemDisks

`func (o *Node) HasSystemDisks() bool`

HasSystemDisks returns a boolean if a field has been set.

### SetSystemDisksNil

`func (o *Node) SetSystemDisksNil(b bool)`

 SetSystemDisksNil sets the value for SystemDisks to be an explicit nil

### UnsetSystemDisks
`func (o *Node) UnsetSystemDisks()`

UnsetSystemDisks ensures that no value is present for SystemDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


