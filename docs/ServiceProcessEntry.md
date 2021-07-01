# ServiceProcessEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessIds** | Pointer to **[]int64** | Specifies the list of process IDs associated with the Service. | [optional] 
**ServiceName** | Pointer to **NullableString** | Specifies the name of the Service. &#39;kApollo&#39; is a service for reclaiming freed disk sectors on Nodes in the SnapFS distributed file system. &#39;kBridge&#39; is a service for managing the SnapFS distributed file system. &#39;kGenie&#39; is a service that is responsible for monitoring hardware health on the Cluster. &#39;kGenieGofer&#39; is a service that links the Genie service to other services on the Cluster. &#39;kMagneto&#39; is the data protection service of the Cohesity Data Platform. &#39;kIris&#39; is the service which serves REST API calls to the UI, CLI, and any scripts written by customers. &#39;kIrisProxy&#39; is a service that links the Iris service to other services on the Cluster. &#39;kScribe&#39; is the service responsible for storing filesystem metadata. &#39;kStats&#39; is the service that is responsible for retrieving and aggregating disk metrics across the Cluster. &#39;kYoda&#39; is an elastic search indexing service. &#39;kAlerts&#39; is a publisher and subscribing service for alerts. &#39;kKeychain&#39; is a service for managing disk encryption keys. &#39;kLogWatcher&#39; is a service that scans the log directory and reduces the number of logs if required. &#39;kStatsCollector&#39; is a service that periodically logs system stats. &#39;kGandalf&#39; is a distributed lock service and coordination manager. &#39;kNexus&#39; indicates the Nexus service. This is the service that is responsible for creation of Clusters and configuration of Nodes and networking. &#39;kNexusProxy&#39; is a service that links the Nexus service to other services on the Cluster. &#39;kStorageProxy&#39; is a service for accessing data on external entities. &#39;kTricorder&#39; is a diagnostic health testing service for Clusters. &#39;kRtClient&#39; is a reverse tunneling client service. &#39;kVaultProxy&#39; is a service for managing external targets that Clusters can be backed up to. &#39;kSmbProxy&#39; is an SMB protocol service. &#39;kBridgeProxy&#39; is the service that links the Bridge service to other services on the Cluster. &#39;kLibrarian&#39; is an elastic search indexing service. &#39;kGroot&#39; is a service for managing replication of SQL databases across multiple nodes in a Cluster. &#39;kEagleAgent&#39; is a service that is responsible for retrieving information on Cluster health. &#39;kAthena&#39; is a service for running distributed containerized applications on the Cohesity Data Platform. &#39;kBifrostBroker&#39; is a service for communicating with the Cohesity proxies for multitenancy. &#39;kSmb2Proxy&#39; is a new SMB protocol service. &#39;kOs&#39; can be specified in order to do a full reboot. &#39;kAtom&#39; is a service for receiving data for the Continuous Data Protection. | [optional] 

## Methods

### NewServiceProcessEntry

`func NewServiceProcessEntry() *ServiceProcessEntry`

NewServiceProcessEntry instantiates a new ServiceProcessEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProcessEntryWithDefaults

`func NewServiceProcessEntryWithDefaults() *ServiceProcessEntry`

NewServiceProcessEntryWithDefaults instantiates a new ServiceProcessEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessIds

`func (o *ServiceProcessEntry) GetProcessIds() []int64`

GetProcessIds returns the ProcessIds field if non-nil, zero value otherwise.

### GetProcessIdsOk

`func (o *ServiceProcessEntry) GetProcessIdsOk() (*[]int64, bool)`

GetProcessIdsOk returns a tuple with the ProcessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessIds

`func (o *ServiceProcessEntry) SetProcessIds(v []int64)`

SetProcessIds sets ProcessIds field to given value.

### HasProcessIds

`func (o *ServiceProcessEntry) HasProcessIds() bool`

HasProcessIds returns a boolean if a field has been set.

### SetProcessIdsNil

`func (o *ServiceProcessEntry) SetProcessIdsNil(b bool)`

 SetProcessIdsNil sets the value for ProcessIds to be an explicit nil

### UnsetProcessIds
`func (o *ServiceProcessEntry) UnsetProcessIds()`

UnsetProcessIds ensures that no value is present for ProcessIds, not even an explicit nil
### GetServiceName

`func (o *ServiceProcessEntry) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceProcessEntry) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceProcessEntry) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceProcessEntry) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ServiceProcessEntry) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ServiceProcessEntry) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


