# NodeStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOperation** | Pointer to **NullableString** | Specifies the active operation on the Node if there is one. | [optional] 
**ClusterId** | Pointer to **int64** | Specifies the Cluster ID if the Node is part of a Cluster. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the Node. | [optional] 
**InCluster** | Pointer to **NullableBool** | Specifies whether or not the Node is part of a Cluster. | [optional] 
**InMaintenanceMode** | Pointer to **NullableBool** | InMaintenanceMode is used to mark a node in maintenance mode. | [optional] 
**IncarnationId** | Pointer to **int64** | Specifies the Incarnation ID if the Node is part of a Cluster. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the Node. | [optional] 
**IsAppNode** | Pointer to **NullableBool** | Whether the node is an app node. | [optional] 
**LastUpgradeTimeSecs** | Pointer to **int64** | Specifies the time of the last upgrade in seconds since the epoch. | [optional] 
**MarkedForRemoval** | Pointer to **NullableBool** | Specifies whether or not this node is marked for removal. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing the current state of the Node. | [optional] 
**RemovalProgressList** | Pointer to [**[]ComponentRemovalProgress**](ComponentRemovalProgress.md) | Removal progress for various components which are not acked yet. | [optional] 
**RemovalReason** | Pointer to **NullableString** | Specifies the reason for the removal operation if there is a removal operation going on. | [optional] 
**Services** | Pointer to [**[]ServiceProcessEntry**](ServiceProcessEntry.md) | Specifies the list of services running on the cluster and their process IDs. | [optional] 
**ServicesAckedList** | Pointer to **[]string** | Services already acked for removal of this entity. | [optional] 
**ServicesNotAcked** | Pointer to **NullableString** | ServicesNotAcked specifies services that have not ACKed yet in string format after the node is marked for removal. | [optional] 
**ServicesNotAckedList** | Pointer to **[]string** | Services not acked yet for removal of this entity. | [optional] 
**SoftwareVersion** | Pointer to **string** | Specifies the version of the software running on the Node. | [optional] 
**Uptime** | Pointer to **NullableString** | Uptime of node. | [optional] 

## Methods

### NewNodeStatusResult

`func NewNodeStatusResult() *NodeStatusResult`

NewNodeStatusResult instantiates a new NodeStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStatusResultWithDefaults

`func NewNodeStatusResultWithDefaults() *NodeStatusResult`

NewNodeStatusResultWithDefaults instantiates a new NodeStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOperation

`func (o *NodeStatusResult) GetActiveOperation() string`

GetActiveOperation returns the ActiveOperation field if non-nil, zero value otherwise.

### GetActiveOperationOk

`func (o *NodeStatusResult) GetActiveOperationOk() (*string, bool)`

GetActiveOperationOk returns a tuple with the ActiveOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOperation

`func (o *NodeStatusResult) SetActiveOperation(v string)`

SetActiveOperation sets ActiveOperation field to given value.

### HasActiveOperation

`func (o *NodeStatusResult) HasActiveOperation() bool`

HasActiveOperation returns a boolean if a field has been set.

### SetActiveOperationNil

`func (o *NodeStatusResult) SetActiveOperationNil(b bool)`

 SetActiveOperationNil sets the value for ActiveOperation to be an explicit nil

### UnsetActiveOperation
`func (o *NodeStatusResult) UnsetActiveOperation()`

UnsetActiveOperation ensures that no value is present for ActiveOperation, not even an explicit nil
### GetClusterId

`func (o *NodeStatusResult) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NodeStatusResult) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NodeStatusResult) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NodeStatusResult) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetId

`func (o *NodeStatusResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeStatusResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeStatusResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeStatusResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeStatusResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeStatusResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInCluster

`func (o *NodeStatusResult) GetInCluster() bool`

GetInCluster returns the InCluster field if non-nil, zero value otherwise.

### GetInClusterOk

`func (o *NodeStatusResult) GetInClusterOk() (*bool, bool)`

GetInClusterOk returns a tuple with the InCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCluster

`func (o *NodeStatusResult) SetInCluster(v bool)`

SetInCluster sets InCluster field to given value.

### HasInCluster

`func (o *NodeStatusResult) HasInCluster() bool`

HasInCluster returns a boolean if a field has been set.

### SetInClusterNil

`func (o *NodeStatusResult) SetInClusterNil(b bool)`

 SetInClusterNil sets the value for InCluster to be an explicit nil

### UnsetInCluster
`func (o *NodeStatusResult) UnsetInCluster()`

UnsetInCluster ensures that no value is present for InCluster, not even an explicit nil
### GetInMaintenanceMode

`func (o *NodeStatusResult) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *NodeStatusResult) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *NodeStatusResult) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *NodeStatusResult) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### SetInMaintenanceModeNil

`func (o *NodeStatusResult) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *NodeStatusResult) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetIncarnationId

`func (o *NodeStatusResult) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *NodeStatusResult) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *NodeStatusResult) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *NodeStatusResult) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### GetIp

`func (o *NodeStatusResult) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeStatusResult) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeStatusResult) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeStatusResult) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *NodeStatusResult) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NodeStatusResult) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsAppNode

`func (o *NodeStatusResult) GetIsAppNode() bool`

GetIsAppNode returns the IsAppNode field if non-nil, zero value otherwise.

### GetIsAppNodeOk

`func (o *NodeStatusResult) GetIsAppNodeOk() (*bool, bool)`

GetIsAppNodeOk returns a tuple with the IsAppNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppNode

`func (o *NodeStatusResult) SetIsAppNode(v bool)`

SetIsAppNode sets IsAppNode field to given value.

### HasIsAppNode

`func (o *NodeStatusResult) HasIsAppNode() bool`

HasIsAppNode returns a boolean if a field has been set.

### SetIsAppNodeNil

`func (o *NodeStatusResult) SetIsAppNodeNil(b bool)`

 SetIsAppNodeNil sets the value for IsAppNode to be an explicit nil

### UnsetIsAppNode
`func (o *NodeStatusResult) UnsetIsAppNode()`

UnsetIsAppNode ensures that no value is present for IsAppNode, not even an explicit nil
### GetLastUpgradeTimeSecs

`func (o *NodeStatusResult) GetLastUpgradeTimeSecs() int64`

GetLastUpgradeTimeSecs returns the LastUpgradeTimeSecs field if non-nil, zero value otherwise.

### GetLastUpgradeTimeSecsOk

`func (o *NodeStatusResult) GetLastUpgradeTimeSecsOk() (*int64, bool)`

GetLastUpgradeTimeSecsOk returns a tuple with the LastUpgradeTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradeTimeSecs

`func (o *NodeStatusResult) SetLastUpgradeTimeSecs(v int64)`

SetLastUpgradeTimeSecs sets LastUpgradeTimeSecs field to given value.

### HasLastUpgradeTimeSecs

`func (o *NodeStatusResult) HasLastUpgradeTimeSecs() bool`

HasLastUpgradeTimeSecs returns a boolean if a field has been set.

### GetMarkedForRemoval

`func (o *NodeStatusResult) GetMarkedForRemoval() bool`

GetMarkedForRemoval returns the MarkedForRemoval field if non-nil, zero value otherwise.

### GetMarkedForRemovalOk

`func (o *NodeStatusResult) GetMarkedForRemovalOk() (*bool, bool)`

GetMarkedForRemovalOk returns a tuple with the MarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForRemoval

`func (o *NodeStatusResult) SetMarkedForRemoval(v bool)`

SetMarkedForRemoval sets MarkedForRemoval field to given value.

### HasMarkedForRemoval

`func (o *NodeStatusResult) HasMarkedForRemoval() bool`

HasMarkedForRemoval returns a boolean if a field has been set.

### SetMarkedForRemovalNil

`func (o *NodeStatusResult) SetMarkedForRemovalNil(b bool)`

 SetMarkedForRemovalNil sets the value for MarkedForRemoval to be an explicit nil

### UnsetMarkedForRemoval
`func (o *NodeStatusResult) UnsetMarkedForRemoval()`

UnsetMarkedForRemoval ensures that no value is present for MarkedForRemoval, not even an explicit nil
### GetMessage

`func (o *NodeStatusResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeStatusResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeStatusResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeStatusResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NodeStatusResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NodeStatusResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRemovalProgressList

`func (o *NodeStatusResult) GetRemovalProgressList() []ComponentRemovalProgress`

GetRemovalProgressList returns the RemovalProgressList field if non-nil, zero value otherwise.

### GetRemovalProgressListOk

`func (o *NodeStatusResult) GetRemovalProgressListOk() (*[]ComponentRemovalProgress, bool)`

GetRemovalProgressListOk returns a tuple with the RemovalProgressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalProgressList

`func (o *NodeStatusResult) SetRemovalProgressList(v []ComponentRemovalProgress)`

SetRemovalProgressList sets RemovalProgressList field to given value.

### HasRemovalProgressList

`func (o *NodeStatusResult) HasRemovalProgressList() bool`

HasRemovalProgressList returns a boolean if a field has been set.

### SetRemovalProgressListNil

`func (o *NodeStatusResult) SetRemovalProgressListNil(b bool)`

 SetRemovalProgressListNil sets the value for RemovalProgressList to be an explicit nil

### UnsetRemovalProgressList
`func (o *NodeStatusResult) UnsetRemovalProgressList()`

UnsetRemovalProgressList ensures that no value is present for RemovalProgressList, not even an explicit nil
### GetRemovalReason

`func (o *NodeStatusResult) GetRemovalReason() string`

GetRemovalReason returns the RemovalReason field if non-nil, zero value otherwise.

### GetRemovalReasonOk

`func (o *NodeStatusResult) GetRemovalReasonOk() (*string, bool)`

GetRemovalReasonOk returns a tuple with the RemovalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalReason

`func (o *NodeStatusResult) SetRemovalReason(v string)`

SetRemovalReason sets RemovalReason field to given value.

### HasRemovalReason

`func (o *NodeStatusResult) HasRemovalReason() bool`

HasRemovalReason returns a boolean if a field has been set.

### SetRemovalReasonNil

`func (o *NodeStatusResult) SetRemovalReasonNil(b bool)`

 SetRemovalReasonNil sets the value for RemovalReason to be an explicit nil

### UnsetRemovalReason
`func (o *NodeStatusResult) UnsetRemovalReason()`

UnsetRemovalReason ensures that no value is present for RemovalReason, not even an explicit nil
### GetServices

`func (o *NodeStatusResult) GetServices() []ServiceProcessEntry`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NodeStatusResult) GetServicesOk() (*[]ServiceProcessEntry, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NodeStatusResult) SetServices(v []ServiceProcessEntry)`

SetServices sets Services field to given value.

### HasServices

`func (o *NodeStatusResult) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetServicesAckedList

`func (o *NodeStatusResult) GetServicesAckedList() []string`

GetServicesAckedList returns the ServicesAckedList field if non-nil, zero value otherwise.

### GetServicesAckedListOk

`func (o *NodeStatusResult) GetServicesAckedListOk() (*[]string, bool)`

GetServicesAckedListOk returns a tuple with the ServicesAckedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesAckedList

`func (o *NodeStatusResult) SetServicesAckedList(v []string)`

SetServicesAckedList sets ServicesAckedList field to given value.

### HasServicesAckedList

`func (o *NodeStatusResult) HasServicesAckedList() bool`

HasServicesAckedList returns a boolean if a field has been set.

### SetServicesAckedListNil

`func (o *NodeStatusResult) SetServicesAckedListNil(b bool)`

 SetServicesAckedListNil sets the value for ServicesAckedList to be an explicit nil

### UnsetServicesAckedList
`func (o *NodeStatusResult) UnsetServicesAckedList()`

UnsetServicesAckedList ensures that no value is present for ServicesAckedList, not even an explicit nil
### GetServicesNotAcked

`func (o *NodeStatusResult) GetServicesNotAcked() string`

GetServicesNotAcked returns the ServicesNotAcked field if non-nil, zero value otherwise.

### GetServicesNotAckedOk

`func (o *NodeStatusResult) GetServicesNotAckedOk() (*string, bool)`

GetServicesNotAckedOk returns a tuple with the ServicesNotAcked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNotAcked

`func (o *NodeStatusResult) SetServicesNotAcked(v string)`

SetServicesNotAcked sets ServicesNotAcked field to given value.

### HasServicesNotAcked

`func (o *NodeStatusResult) HasServicesNotAcked() bool`

HasServicesNotAcked returns a boolean if a field has been set.

### SetServicesNotAckedNil

`func (o *NodeStatusResult) SetServicesNotAckedNil(b bool)`

 SetServicesNotAckedNil sets the value for ServicesNotAcked to be an explicit nil

### UnsetServicesNotAcked
`func (o *NodeStatusResult) UnsetServicesNotAcked()`

UnsetServicesNotAcked ensures that no value is present for ServicesNotAcked, not even an explicit nil
### GetServicesNotAckedList

`func (o *NodeStatusResult) GetServicesNotAckedList() []string`

GetServicesNotAckedList returns the ServicesNotAckedList field if non-nil, zero value otherwise.

### GetServicesNotAckedListOk

`func (o *NodeStatusResult) GetServicesNotAckedListOk() (*[]string, bool)`

GetServicesNotAckedListOk returns a tuple with the ServicesNotAckedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNotAckedList

`func (o *NodeStatusResult) SetServicesNotAckedList(v []string)`

SetServicesNotAckedList sets ServicesNotAckedList field to given value.

### HasServicesNotAckedList

`func (o *NodeStatusResult) HasServicesNotAckedList() bool`

HasServicesNotAckedList returns a boolean if a field has been set.

### SetServicesNotAckedListNil

`func (o *NodeStatusResult) SetServicesNotAckedListNil(b bool)`

 SetServicesNotAckedListNil sets the value for ServicesNotAckedList to be an explicit nil

### UnsetServicesNotAckedList
`func (o *NodeStatusResult) UnsetServicesNotAckedList()`

UnsetServicesNotAckedList ensures that no value is present for ServicesNotAckedList, not even an explicit nil
### GetSoftwareVersion

`func (o *NodeStatusResult) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NodeStatusResult) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NodeStatusResult) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NodeStatusResult) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetUptime

`func (o *NodeStatusResult) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *NodeStatusResult) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *NodeStatusResult) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *NodeStatusResult) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### SetUptimeNil

`func (o *NodeStatusResult) SetUptimeNil(b bool)`

 SetUptimeNil sets the value for Uptime to be an explicit nil

### UnsetUptime
`func (o *NodeStatusResult) UnsetUptime()`

UnsetUptime ensures that no value is present for Uptime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


