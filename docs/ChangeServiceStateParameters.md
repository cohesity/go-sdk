# ChangeServiceStateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action to take on the specified service. &#39;kStop&#39; indicates that the specified services will be stopped. &#39;kStart&#39; indicates that the specified services will be started. &#39;kRestart&#39; indicates that the specified services will be restarted. | [optional] 
**Services** | Pointer to **[]string** | Specifies the list of services to take the specified action on. If none are specified, all Cluster services will be affected. It should also be noted that some Cluster services are not stoppable and may not be affected by the action specified. &#39;kApollo&#39; is a service for reclaiming freed disk sectors on Nodes in the SnapFS distributed file system. &#39;kBridge&#39; is a service for managing the SnapFS distributed file system. &#39;kGenie&#39; is a service that is responsible for monitoring hardware health on the Cluster. &#39;kGenieGofer&#39; is a service that links the Genie service to other services on the Cluster. &#39;kMagneto&#39; is the data protection service of the Cohesity Data Platform. &#39;kIris&#39; is the service which serves REST API calls to the UI, CLI, and any scripts written by customers. &#39;kIrisProxy&#39; is a service that links the Iris service to other services on the Cluster. &#39;kScribe&#39; is the service responsible for storing filesystem metadata. &#39;kStats&#39; is the service that is responsible for retrieving and aggregating disk metrics across the Cluster. &#39;kYoda&#39; is an elastic search indexing service. &#39;kAlerts&#39; is a publisher and subscribing service for alerts. &#39;kKeychain&#39; is a service for managing disk encryption keys. &#39;kLogWatcher&#39; is a service that scans the log directory and reduces the number of logs if required. &#39;kStatsCollector&#39; is a service that periodically logs system stats. &#39;kGandalf&#39; is a distributed lock service and coordination manager. &#39;kNexus&#39; indicates the Nexus service. This is the service that is responsible for creation of Clusters and configuration of Nodes and networking. &#39;kNexusProxy&#39; is a service that links the Nexus service to other services on the Cluster. &#39;kStorageProxy&#39; is a service for accessing data on external entities. &#39;kTricorder&#39; is a diagnostic health testing service for Clusters. &#39;kRtClient&#39; is a reverse tunneling client service. &#39;kVaultProxy&#39; is a service for managing external targets that Clusters can be backed up to. &#39;kSmbProxy&#39; is an SMB protocol service. &#39;kBridgeProxy&#39; is the service that links the Bridge service to other services on the Cluster. &#39;kLibrarian&#39; is an elastic search indexing service. &#39;kGroot&#39; is a service for managing replication of SQL databases across multiple nodes in a Cluster. &#39;kEagleAgent&#39; is a service that is responsible for retrieving information on Cluster health. &#39;kAthena&#39; is a service for running distributed containerized applications on the Cohesity Data Platform. &#39;kBifrostBroker&#39; is a service for communicating with the Cohesity proxies for multitenancy. &#39;kSmb2Proxy&#39; is a new SMB protocol service. &#39;kOs&#39; can be specified in order to do a full reboot. &#39;kAtom&#39; is a service for receiving data for the Continuous Data Protection. | [optional] 

## Methods

### NewChangeServiceStateParameters

`func NewChangeServiceStateParameters() *ChangeServiceStateParameters`

NewChangeServiceStateParameters instantiates a new ChangeServiceStateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeServiceStateParametersWithDefaults

`func NewChangeServiceStateParametersWithDefaults() *ChangeServiceStateParameters`

NewChangeServiceStateParametersWithDefaults instantiates a new ChangeServiceStateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ChangeServiceStateParameters) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChangeServiceStateParameters) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChangeServiceStateParameters) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ChangeServiceStateParameters) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *ChangeServiceStateParameters) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *ChangeServiceStateParameters) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetServices

`func (o *ChangeServiceStateParameters) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ChangeServiceStateParameters) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ChangeServiceStateParameters) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ChangeServiceStateParameters) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ChangeServiceStateParameters) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ChangeServiceStateParameters) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


