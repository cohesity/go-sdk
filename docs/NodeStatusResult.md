# NodeStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOperation** | Pointer to **NullableString** | Specifies the active operation on the Node if there is one. &#39;kNone&#39; specifies that there is no active operation on the Node. &#39;kDestroyCluster&#39; specifies that the Cluster which the Node is a part of is currently being destroyed. &#39;kUpgradeCluster&#39; specifies that the Cluster which the Node is a part of is currently being upgraded to a new software package. &#39;kRestartCluster&#39; specifies that the Cluster which the Node is a part of is currently being restarted. &#39;kCreateCluster&#39; specifies that the Node is currently being used to create a new Cluster. &#39;kExpandCluster&#39; specifies that the Node is currently being added to a Cluster or being used to assist in adding another Node to a Cluster. &#39;kUpgradeNode&#39; specifies that the Node is currently being upgraded to a new software package. &#39;kRemoveNode&#39; specifies that the Node is currently being removed from a Cluster or that it is assisting in removing another Node from a Cluster. &#39;kAddDisks&#39; specifies that the Node is being used to assist in adding disks to the Cluster. &#39;kMarkDiskOffline&#39; specifies that the Node is being use to assist in marking a disk in the Cluster as offline. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the Cluster ID if the Node is part of a Cluster. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the Node. | [optional] 
**InCluster** | Pointer to **NullableBool** | Specifies whether or not the Node is part of a Cluster. | [optional] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the Incarnation ID if the Node is part of a Cluster. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the Node. | [optional] 
**LastUpgradeTimeSecs** | Pointer to **NullableInt64** | Specifies the time of the last upgrade in seconds since the epoch. | [optional] 
**MarkedForRemoval** | Pointer to **NullableBool** | Specifies whether or not this node is marked for removal. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing the current state of the Node. | [optional] 
**RemovalReason** | Pointer to **NullableString** | Specifies the reason for the removal operation if there is a removal operation going on. &#39;kUnknown&#39; specifies that the removal reason is not known. &#39;kAutoHealthCheck&#39; specifies that an internal health check found problems with the Node. &#39;kUserGracefulRemoval&#39; specifies that the user requested a graceful removal. &#39;kUserAvoidAccess&#39; specifies that the user requested to avoid access to this Node. &#39;kUserGracefulNodeRemoval&#39; specifies that the user requested a graceful removal for all of the disks in this Node. &#39;kUserRemoveDownNode&#39; specifies that the user requested a graceful removal of the Node while it is down. | [optional] 
**Services** | Pointer to [**[]ServiceProcessEntry**](ServiceProcessEntry.md) | Specifies the list of services running on the cluster and their process Ids. | [optional] 
**SoftwareVersion** | Pointer to **NullableString** | Specifies the version of the software running on the Node. | [optional] 
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

### SetClusterIdNil

`func (o *NodeStatusResult) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *NodeStatusResult) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
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

### SetIncarnationIdNil

`func (o *NodeStatusResult) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *NodeStatusResult) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
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

### SetLastUpgradeTimeSecsNil

`func (o *NodeStatusResult) SetLastUpgradeTimeSecsNil(b bool)`

 SetLastUpgradeTimeSecsNil sets the value for LastUpgradeTimeSecs to be an explicit nil

### UnsetLastUpgradeTimeSecs
`func (o *NodeStatusResult) UnsetLastUpgradeTimeSecs()`

UnsetLastUpgradeTimeSecs ensures that no value is present for LastUpgradeTimeSecs, not even an explicit nil
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

### SetServicesNil

`func (o *NodeStatusResult) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *NodeStatusResult) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
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

### SetSoftwareVersionNil

`func (o *NodeStatusResult) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *NodeStatusResult) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
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


