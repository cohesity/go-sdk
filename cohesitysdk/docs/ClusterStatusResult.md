# ClusterStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the ID of the Cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation ID of the Cluster. | [optional] 
**CurrentOperation** | Pointer to **NullableString** | Specifies the current operation being run on the Cluster. &#39;kNone&#39; indicates that there is no current operation taking place. &#39;kDestroy&#39; indicates that the Cluster is currently being destroyed. &#39;kUpgrade&#39; indicates that the Cluster is currently being upgraded. &#39;kClean&#39; indicates that the Cluster is being cleaned. &#39;kRemoveNode&#39; indicates that a Node is being removed from the Cluster. &#39;kRestartServices&#39; indicates that the services on the Cluster are currently being restarted. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing details of the Cluster status. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Cluster. | [optional] 
**NodeStatuses** | Pointer to [**[]NodeStatusResult**](NodeStatusResult.md) | Specifies the status of each Node on the Cluster. | [optional] 
**RemovalState** | Pointer to **NullableString** | Specifies the current healing state of the Cluster. &#39;kNoRemoval&#39; indicates that there are no removal operations currently happening on the Cluster. &#39;kNodeRemoval&#39; indicates that there is a Node being removed from the Cluster. &#39;kDiskRemoval&#39; indicates that there is a Disk being removed from the Cluster. &#39;kNodeAndDiskRemoval&#39; indicates that there is a Node and a Disk being removed from the Cluster. | [optional] 
**ServicesSynced** | Pointer to **NullableBool** | Specifies whether or not the services are synced with the list of stopped services. | [optional] 
**SoftwareVersion** | Pointer to **NullableString** | Specifies the software version of the Cluster. | [optional] 
**StoppedServices** | Pointer to **[]string** | Specifies the list of stopped services on the Cluster. &#39;kApollo&#39; is a service for reclaiming freed disk sectors on Nodes in the SnapFS distributed file system. &#39;kBridge&#39; is a service for managing the SnapFS distributed file system. &#39;kGenie&#39; is a service that is responsible for monitoring hardware health on the Cluster. &#39;kGenieGofer&#39; is a service that links the Genie service to other services on the Cluster. &#39;kMagneto&#39; is the data protection service of the Cohesity Data Platform. &#39;kIris&#39; is the service which serves REST API calls to the UI, CLI, and any scripts written by customers. &#39;kIrisProxy&#39; is a service that links the Iris service to other services on the Cluster. &#39;kScribe&#39; is the service responsible for storing filesystem metadata. &#39;kStats&#39; is the service that is responsible for retrieving and aggregating disk metrics across the Cluster. &#39;kYoda&#39; is an elastic search indexing service. &#39;kAlerts&#39; is a publisher and subscribing service for alerts. &#39;kKeychain&#39; is a service for managing disk encryption keys. &#39;kLogWatcher&#39; is a service that scans the log directory and reduces the number of logs if required. &#39;kStatsCollector&#39; is a service that periodically logs system stats. &#39;kGandalf&#39; is a distributed lock service and coordination manager. &#39;kNexus&#39; indicates the Nexus service. This is the service that is responsible for creation of Clusters and configuration of Nodes and networking. &#39;kNexusProxy&#39; is a service that links the Nexus service to other services on the Cluster. &#39;kStorageProxy&#39; is a service for accessing data on external entities. &#39;kTricorder&#39; is a diagnostic health testing service for Clusters. &#39;kRtClient&#39; is a reverse tunneling client service. &#39;kVaultProxy&#39; is a service for managing external targets that Clusters can be backed up to. &#39;kSmbProxy&#39; is an SMB protocol service. &#39;kBridgeProxy&#39; is the service that links the Bridge service to other services on the Cluster. &#39;kLibrarian&#39; is an elastic search indexing service. &#39;kGroot&#39; is a service for managing replication of SQL databases across multiple nodes in a Cluster. &#39;kEagleAgent&#39; is a service that is responsible for retrieving information on Cluster health. &#39;kAthena&#39; is a service for running distributed containerized applications on the Cohesity Data Platform. &#39;kBifrostBroker&#39; is a service for communicating with the Cohesity proxies for multitenancy. &#39;kSmb2Proxy&#39; is a new SMB protocol service. &#39;kOs&#39; can be specified in order to do a full reboot. &#39;kAtom&#39; is a service for receiving data for the Continuous Data Protection. | [optional] 

## Methods

### NewClusterStatusResult

`func NewClusterStatusResult() *ClusterStatusResult`

NewClusterStatusResult instantiates a new ClusterStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusResultWithDefaults

`func NewClusterStatusResultWithDefaults() *ClusterStatusResult`

NewClusterStatusResultWithDefaults instantiates a new ClusterStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterStatusResult) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterStatusResult) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterStatusResult) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterStatusResult) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ClusterStatusResult) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ClusterStatusResult) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ClusterStatusResult) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ClusterStatusResult) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ClusterStatusResult) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ClusterStatusResult) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ClusterStatusResult) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ClusterStatusResult) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetCurrentOperation

`func (o *ClusterStatusResult) GetCurrentOperation() string`

GetCurrentOperation returns the CurrentOperation field if non-nil, zero value otherwise.

### GetCurrentOperationOk

`func (o *ClusterStatusResult) GetCurrentOperationOk() (*string, bool)`

GetCurrentOperationOk returns a tuple with the CurrentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOperation

`func (o *ClusterStatusResult) SetCurrentOperation(v string)`

SetCurrentOperation sets CurrentOperation field to given value.

### HasCurrentOperation

`func (o *ClusterStatusResult) HasCurrentOperation() bool`

HasCurrentOperation returns a boolean if a field has been set.

### SetCurrentOperationNil

`func (o *ClusterStatusResult) SetCurrentOperationNil(b bool)`

 SetCurrentOperationNil sets the value for CurrentOperation to be an explicit nil

### UnsetCurrentOperation
`func (o *ClusterStatusResult) UnsetCurrentOperation()`

UnsetCurrentOperation ensures that no value is present for CurrentOperation, not even an explicit nil
### GetMessage

`func (o *ClusterStatusResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterStatusResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterStatusResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterStatusResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterStatusResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterStatusResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetName

`func (o *ClusterStatusResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStatusResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStatusResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStatusResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClusterStatusResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClusterStatusResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeStatuses

`func (o *ClusterStatusResult) GetNodeStatuses() []NodeStatusResult`

GetNodeStatuses returns the NodeStatuses field if non-nil, zero value otherwise.

### GetNodeStatusesOk

`func (o *ClusterStatusResult) GetNodeStatusesOk() (*[]NodeStatusResult, bool)`

GetNodeStatusesOk returns a tuple with the NodeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatuses

`func (o *ClusterStatusResult) SetNodeStatuses(v []NodeStatusResult)`

SetNodeStatuses sets NodeStatuses field to given value.

### HasNodeStatuses

`func (o *ClusterStatusResult) HasNodeStatuses() bool`

HasNodeStatuses returns a boolean if a field has been set.

### SetNodeStatusesNil

`func (o *ClusterStatusResult) SetNodeStatusesNil(b bool)`

 SetNodeStatusesNil sets the value for NodeStatuses to be an explicit nil

### UnsetNodeStatuses
`func (o *ClusterStatusResult) UnsetNodeStatuses()`

UnsetNodeStatuses ensures that no value is present for NodeStatuses, not even an explicit nil
### GetRemovalState

`func (o *ClusterStatusResult) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *ClusterStatusResult) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *ClusterStatusResult) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *ClusterStatusResult) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *ClusterStatusResult) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *ClusterStatusResult) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetServicesSynced

`func (o *ClusterStatusResult) GetServicesSynced() bool`

GetServicesSynced returns the ServicesSynced field if non-nil, zero value otherwise.

### GetServicesSyncedOk

`func (o *ClusterStatusResult) GetServicesSyncedOk() (*bool, bool)`

GetServicesSyncedOk returns a tuple with the ServicesSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesSynced

`func (o *ClusterStatusResult) SetServicesSynced(v bool)`

SetServicesSynced sets ServicesSynced field to given value.

### HasServicesSynced

`func (o *ClusterStatusResult) HasServicesSynced() bool`

HasServicesSynced returns a boolean if a field has been set.

### SetServicesSyncedNil

`func (o *ClusterStatusResult) SetServicesSyncedNil(b bool)`

 SetServicesSyncedNil sets the value for ServicesSynced to be an explicit nil

### UnsetServicesSynced
`func (o *ClusterStatusResult) UnsetServicesSynced()`

UnsetServicesSynced ensures that no value is present for ServicesSynced, not even an explicit nil
### GetSoftwareVersion

`func (o *ClusterStatusResult) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *ClusterStatusResult) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *ClusterStatusResult) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *ClusterStatusResult) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *ClusterStatusResult) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *ClusterStatusResult) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
### GetStoppedServices

`func (o *ClusterStatusResult) GetStoppedServices() []string`

GetStoppedServices returns the StoppedServices field if non-nil, zero value otherwise.

### GetStoppedServicesOk

`func (o *ClusterStatusResult) GetStoppedServicesOk() (*[]string, bool)`

GetStoppedServicesOk returns a tuple with the StoppedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedServices

`func (o *ClusterStatusResult) SetStoppedServices(v []string)`

SetStoppedServices sets StoppedServices field to given value.

### HasStoppedServices

`func (o *ClusterStatusResult) HasStoppedServices() bool`

HasStoppedServices returns a boolean if a field has been set.

### SetStoppedServicesNil

`func (o *ClusterStatusResult) SetStoppedServicesNil(b bool)`

 SetStoppedServicesNil sets the value for StoppedServices to be an explicit nil

### UnsetStoppedServices
`func (o *ClusterStatusResult) UnsetStoppedServices()`

UnsetStoppedServices ensures that no value is present for StoppedServices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


